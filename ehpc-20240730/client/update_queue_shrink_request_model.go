// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQueueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateQueueShrinkRequest
	GetClusterId() *string
	SetQueueShrink(v string) *UpdateQueueShrinkRequest
	GetQueueShrink() *string
}

type UpdateQueueShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The information about the queue to be updated.
	QueueShrink *string `json:"Queue,omitempty" xml:"Queue,omitempty"`
}

func (s UpdateQueueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQueueShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateQueueShrinkRequest) GetQueueShrink() *string {
	return s.QueueShrink
}

func (s *UpdateQueueShrinkRequest) SetClusterId(v string) *UpdateQueueShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueShrinkRequest) SetQueueShrink(v string) *UpdateQueueShrinkRequest {
	s.QueueShrink = &v
	return s
}

func (s *UpdateQueueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
