package control_test

import (
	"github.com/redpointgames/rtmp/chunk"
	"github.com/redpointgames/rtmp/control"
	"github.com/stretchr/testify/mock"
)

type MockChunker struct {
	mock.Mock
}

func (c *MockChunker) Chunk(control control.Control) (*chunk.Chunk, error) {
	args := c.Called(control)
	return args.Get(0).(*chunk.Chunk), args.Error(1)
}
