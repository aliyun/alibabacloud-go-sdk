// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachNodesShrinkRequest
	GetClusterId() *string
	SetComputeNodeShrink(v string) *AttachNodesShrinkRequest
	GetComputeNodeShrink() *string
	SetQueueName(v string) *AttachNodesShrinkRequest
	GetQueueName() *string
}

type AttachNodesShrinkRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-xxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of compute nodes.
	//
	// This parameter is required.
	ComputeNodeShrink *string `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty"`
	// The name of the queue to which the instance is to be added.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AttachNodesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachNodesShrinkRequest) GetComputeNodeShrink() *string {
	return s.ComputeNodeShrink
}

func (s *AttachNodesShrinkRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *AttachNodesShrinkRequest) SetClusterId(v string) *AttachNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachNodesShrinkRequest) SetComputeNodeShrink(v string) *AttachNodesShrinkRequest {
	s.ComputeNodeShrink = &v
	return s
}

func (s *AttachNodesShrinkRequest) SetQueueName(v string) *AttachNodesShrinkRequest {
	s.QueueName = &v
	return s
}

func (s *AttachNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
