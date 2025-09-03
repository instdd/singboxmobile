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

// The methods Start/Stop/InboundPacket are provided in build-tagged files:
// - engine_stub.go       (default, no with_singbox): returns explanatory error
// - engine_singbox.go    (with_singbox): starts real sing-box core

// Helper error type with text for gomobile-friendly bridging.
type mobErr struct{ s string }

func (e mobErr) Error() string { return e.s }

func newError(s string) error { return mobErr{s: s} }

// Echo is a trivial exported function using only basic supported types.
// This helps some gomobile toolchains to detect exported names during bind.
func Echo(s string) string { return s }
