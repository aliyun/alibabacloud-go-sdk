// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateQueueRequest
	GetClusterId() *string
	SetQueue(v *QueueTemplate) *CreateQueueRequest
	GetQueue() *QueueTemplate
}

type CreateQueueRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configurations of the queue to be created.
	Queue *QueueTemplate `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s CreateQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateQueueRequest) GetQueue() *QueueTemplate {
	return s.Queue
}

func (s *CreateQueueRequest) SetClusterId(v string) *CreateQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateQueueRequest) SetQueue(v *QueueTemplate) *CreateQueueRequest {
	s.Queue = v
	return s
}

func (s *CreateQueueRequest) Validate() error {
	if s.Queue != nil {
		if err := s.Queue.Validate(); err != nil {
			return err
		}
	}
	return nil
}
