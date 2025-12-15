// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateQueueShrinkRequest
	GetClusterId() *string
	SetQueueShrink(v string) *CreateQueueShrinkRequest
	GetQueueShrink() *string
}

type CreateQueueShrinkRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configurations of the queue to be created.
	QueueShrink *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s CreateQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateQueueShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateQueueShrinkRequest) GetQueueShrink() *string {
	return s.QueueShrink
}

func (s *CreateQueueShrinkRequest) SetClusterId(v string) *CreateQueueShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateQueueShrinkRequest) SetQueueShrink(v string) *CreateQueueShrinkRequest {
	s.QueueShrink = &v
	return s
}

func (s *CreateQueueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
