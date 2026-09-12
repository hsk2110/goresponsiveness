package lgc

import (
	"net"
	"sync/atomic"
)

type countingConnWrite struct {
	net.Conn
	n *uint64
}

func (c *countingConnWrite) Unwrap() net.Conn { return c.Conn }
func (c *countingConnRead) Unwrap() net.Conn  { return c.Conn }

func (c *countingConnWrite) Write(p []byte) (int, error) {
	n, err := c.Conn.Write(p)
	if n > 0 {
		atomic.AddUint64(c.n, uint64(n))
	}
	return n, err
}

type countingConnRead struct {
	net.Conn
	n *uint64
}

func (c *countingConnRead) Read(p []byte) (int, error) {
	n, err := c.Conn.Read(p)
	if n > 0 {
		atomic.AddUint64(c.n, uint64(n))
	}
	return n, err
}
