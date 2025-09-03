package singboxmobile

// NOTE: This is a minimal gomobile-friendly API surface meant to be wrapped
// by Objective-C/Swift via gomobile bind. It is a placeholder that you can
// extend to actually start sing-box and bridge packets to NEPacketTunnelFlow.

// Outbound is implemented on the Swift side. It receives outbound IP packets
// from the core and logs diagnostic lines.
type Outbound interface {
    WritePacket(p []byte)
    Log(line string)
}

// Engine represents a running core instance.
type Engine struct {
    out     Outbound
    running bool
    core    interface{}
}

// NewEngine creates an Engine bound to Swift-side Outbound.
func NewEngine(out Outbound) *Engine { return &Engine{out: out} }

// New provides a zero-arg constructor so gobind always has a simple
// exported function to bind even if interfaces are not mapped.
func New() *Engine { return &Engine{} }

// Version exposes a simple exported symbol so gobind always detects this
// package as bindable even if toolchain heuristics are picky.
func Version() string { return "0.0.1" }

// Start should initialize sing-box with the provided JSON config and spawn
// goroutines that read from and write to the NEPacketTunnelFlow via the
// InboundPacket/Outbound callbacks.
//
// TODO: Replace the placeholder body with real sing-box startup and TUN
// bridging (gVisor netstack) when integrating the core module.
func (e *Engine) Start(config []byte) error {
    e.running = true
    if e.out != nil {
        e.out.Log("singboxmobile: Start called (placeholder)")
    }
    // Return an error until a real core is wired to avoid blackholing traffic.
    return newError("core not integrated")
}

// Stop terminates the core.
func (e *Engine) Stop() {
    e.running = false
    if e.out != nil {
        e.out.Log("singboxmobile: Stop")
    }
}

// InboundPacket should be called by Swift with packets read from
// NEPacketTunnelFlow. In a real implementation, feed this into sing-box's
// TUN device or netstack. Placeholder drops the packet.
func (e *Engine) InboundPacket(p []byte) {
    // no-op placeholder
}

// Helper error type with text for gomobile-friendly bridging.
type mobErr struct{ s string }

func (e mobErr) Error() string { return e.s }

func newError(s string) error { return mobErr{s: s} }

// Echo is a trivial exported function using only basic supported types.
// This helps some gomobile toolchains to detect exported names during bind.
func Echo(s string) string { return s }
