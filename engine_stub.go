//go:build !with_singbox

package singboxmobile

// This file provides the default (no-singbox) implementation used when the
// build tag `with_singbox` is NOT supplied. It allows the gomobile binding to
// compile and integrate on iOS without pulling large dependencies.

// Start returns an explanatory error when sing-box is not compiled in.
func (e *Engine) Start(config []byte) error {
    if e.out != nil {
        e.out.Log("singboxmobile: Start called without with_singbox build tag")
    }
    return newError("with_singbox build tag required (engine not linked)")
}

func (e *Engine) InboundPacket(p []byte) {
    // No-op in stub implementation.
}

func (e *Engine) Stop() {
    if e.out != nil {
        e.out.Log("singboxmobile: Stop (stub)")
    }
}

