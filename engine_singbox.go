//go:build with_singbox

package singboxmobile

import (
    "context"
    "encoding/json"
    "time"

    boxpkg "github.com/sagernet/sing-box/box"
    "github.com/sagernet/sing-box/option"
)

// Engine implementation backed by sing-box. This requires building with the
// `with_singbox` build tag and having the sing-box module available.

type coreRunner struct {
    box    *boxpkg.Box
    cancel context.CancelFunc
}

func (e *Engine) Start(config []byte) error {
    if e.running {
        return nil
    }
    if e.out != nil {
        e.out.Log("singboxmobile: initializing sing-box core")
    }

    // Parse JSON config produced by the iOS app (ConfigBuilder).
    var opts option.Options
    if err := json.Unmarshal(config, &opts); err != nil {
        return err
    }

    // Create sing-box instance.
    ctx, cancel := context.WithCancel(context.Background())
    eb, err := boxpkg.New(opts)
    if err != nil {
        cancel()
        return err
    }
    // Start core services.
    if err := eb.Start(); err != nil {
        cancel()
        _ = eb.Close()
        return err
    }

    // Store state and mark as running.
    e.running = true
    // Stash runner in a private field via closure
    e.core = &coreRunner{box: eb, cancel: cancel}
    if e.out != nil {
        e.out.Log("singboxmobile: core started")
    }
    return nil
}

// InboundPacket currently relies on sing-box TUN inbound to pick up packets
// from the OS. For a full in-process bridge to NEPacketTunnelFlow you need a
// custom TUN adapter; that is outside the current scope. This method is kept
// for API compatibility and future extension.
func (e *Engine) InboundPacket(p []byte) {
    // No direct injection without custom TUN adapter; left intentionally empty.
}

func (e *Engine) Stop() {
    if !e.running {
        return
    }
    if e.out != nil {
        e.out.Log("singboxmobile: stopping core")
    }
    if e.core != nil {
        e.core.cancel()
        // Give the core a brief moment to settle.
        time.Sleep(50 * time.Millisecond)
        _ = e.core.box.Close()
        e.core = nil
    }
    e.running = false
}

