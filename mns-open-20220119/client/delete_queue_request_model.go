// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueueName(v string) *DeleteQueueRequest
	GetQueueName() *string
}

type DeleteQueueRequest struct {
	// The name of the queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// tf-testAccMNSQueue-525478433321945943
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

func (s *DeleteQueueRequest) Validate() error {
	return dara.Validate(s)
}
