// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunOptions interface {
	dara.Model
	String() string
	GoString() string
	SetBatchWindow(v *BatchWindow) *RunOptions
	GetBatchWindow() *BatchWindow
	SetDeadLetterQueue(v *DeadLetterQueue) *RunOptions
	GetDeadLetterQueue() *DeadLetterQueue
	SetErrorsTolerance(v string) *RunOptions
	GetErrorsTolerance() *string
	SetMode(v string) *RunOptions
	GetMode() *string
	SetRetryStrategy(v *RetryStrategy) *RunOptions
	GetRetryStrategy() *RetryStrategy
}

type RunOptions struct {
	BatchWindow     *BatchWindow     `json:"batchWindow,omitempty" xml:"batchWindow,omitempty"`
	DeadLetterQueue *DeadLetterQueue `json:"deadLetterQueue,omitempty" xml:"deadLetterQueue,omitempty"`
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"errorsTolerance,omitempty" xml:"errorsTolerance,omitempty"`
	// example:
	//
	// event-streaming
	Mode          *string        `json:"mode,omitempty" xml:"mode,omitempty"`
	RetryStrategy *RetryStrategy `json:"retryStrategy,omitempty" xml:"retryStrategy,omitempty"`
}

func (s RunOptions) String() string {
	return dara.Prettify(s)
}

func (s RunOptions) GoString() string {
	return s.String()
}

func (s *RunOptions) GetBatchWindow() *BatchWindow {
	return s.BatchWindow
}

func (s *RunOptions) GetDeadLetterQueue() *DeadLetterQueue {
	return s.DeadLetterQueue
}

func (s *RunOptions) GetErrorsTolerance() *string {
	return s.ErrorsTolerance
}

func (s *RunOptions) GetMode() *string {
	return s.Mode
}

func (s *RunOptions) GetRetryStrategy() *RetryStrategy {
	return s.RetryStrategy
}

func (s *RunOptions) SetBatchWindow(v *BatchWindow) *RunOptions {
	s.BatchWindow = v
	return s
}

func (s *RunOptions) SetDeadLetterQueue(v *DeadLetterQueue) *RunOptions {
	s.DeadLetterQueue = v
	return s
}

func (s *RunOptions) SetErrorsTolerance(v string) *RunOptions {
	s.ErrorsTolerance = &v
	return s
}

func (s *RunOptions) SetMode(v string) *RunOptions {
	s.Mode = &v
	return s
}

func (s *RunOptions) SetRetryStrategy(v *RetryStrategy) *RunOptions {
	s.RetryStrategy = v
	return s
}

func (s *RunOptions) Validate() error {
	return dara.Validate(s)
}
