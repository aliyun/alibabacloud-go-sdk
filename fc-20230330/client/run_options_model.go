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
	// The batch window configurations.
	BatchWindow *BatchWindow `json:"batchWindow,omitempty" xml:"batchWindow,omitempty"`
	// Whether to enable dead-letter queues. If you configure this parameter, dead-letter queues are enabled. By default, dead-letter queues are not enabled and messages are discarded when the retry policy is exhausted. Queues of Simple Message Queue (formerly MNS), ApsaraMQ for RocketMQ, and ApsaraMQ for Kafka, and EventBridge event buses can be used as dead-letter queues.
	DeadLetterQueue *DeadLetterQueue `json:"deadLetterQueue,omitempty" xml:"deadLetterQueue,omitempty"`
	// The fault tolerance policy. Valid values:
	//
	// 	- **NONE**: does not tolerate exceptions.
	//
	// 	- **ALL**: tolerates all exceptions.
	//
	// >  The default value is **NONE**.
	//
	// example:
	//
	// ALL
	ErrorsTolerance *string `json:"errorsTolerance,omitempty" xml:"errorsTolerance,omitempty"`
	// The underlying application mode when message data is pushed to Function Compute. Valid values:
	//
	// 	- **event-streaming**: the event streaming mode. In this mode, events are pushed in arrays. One or more message events are pushed to the function in batches based on your push configurations. This mode is suitable for end-to-end streaming data processing scenarios. The event streaming mode supports the following event sources: Simple Message Queue (formerly MNS), ApsaraMQ for RocketMQ, ApsaraMQ for RabbitMQ, ApsaraMQ for Kafka, ApsaraMQ for MQTT, and Data Transmission Service (DTS).
	//
	// 	- **event-driven**: the event mode. In event mode, a single message is passed to the function as event parameters at a time. Events follow the CloudEvents specifications. The event mode supports the following event sources: Default, Simple Message Queue (formerly MNS), ApsaraMQ for RocketMQ, and ApsaraMQ for RabbitMQ. In this mode, batch configurations are not supported.
	//
	// example:
	//
	// event-streaming
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The retry policy that you want to use if events fail to be pushed.
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
	if s.BatchWindow != nil {
		if err := s.BatchWindow.Validate(); err != nil {
			return err
		}
	}
	if s.DeadLetterQueue != nil {
		if err := s.DeadLetterQueue.Validate(); err != nil {
			return err
		}
	}
	if s.RetryStrategy != nil {
		if err := s.RetryStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
