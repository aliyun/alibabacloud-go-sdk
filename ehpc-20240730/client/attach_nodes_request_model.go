// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachNodesRequest
	GetClusterId() *string
	SetComputeNode(v *AttachNodesRequestComputeNode) *AttachNodesRequest
	GetComputeNode() *AttachNodesRequestComputeNode
	SetQueueName(v string) *AttachNodesRequest
	GetQueueName() *string
}

type AttachNodesRequest struct {
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
	ComputeNode *AttachNodesRequestComputeNode `json:"ComputeNode,omitempty" xml:"ComputeNode,omitempty" type:"Struct"`
	// The name of the queue to which the instance is to be added.
	//
	// example:
	//
	// comp
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AttachNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachNodesRequest) GoString() string {
	return s.String()
}

func (s *AttachNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachNodesRequest) GetComputeNode() *AttachNodesRequestComputeNode {
	return s.ComputeNode
}

func (s *AttachNodesRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *AttachNodesRequest) SetClusterId(v string) *AttachNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachNodesRequest) SetComputeNode(v *AttachNodesRequestComputeNode) *AttachNodesRequest {
	s.ComputeNode = v
	return s
}

func (s *AttachNodesRequest) SetQueueName(v string) *AttachNodesRequest {
	s.QueueName = &v
	return s
}

func (s *AttachNodesRequest) Validate() error {
	if s.ComputeNode != nil {
		if err := s.ComputeNode.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AttachNodesRequestComputeNode struct {
	// The image ID. This image will be used to replace the original system disk image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-bp1h765oyqyxxxxxxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The IDs of ECS instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s AttachNodesRequestComputeNode) String() string {
	return dara.Prettify(s)
}

func (s AttachNodesRequestComputeNode) GoString() string {
	return s.String()
}

func (s *AttachNodesRequestComputeNode) GetImageId() *string {
	return s.ImageId
}

func (s *AttachNodesRequestComputeNode) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *AttachNodesRequestComputeNode) SetImageId(v string) *AttachNodesRequestComputeNode {
	s.ImageId = &v
	return s
}

func (s *AttachNodesRequestComputeNode) SetInstanceIds(v []*string) *AttachNodesRequestComputeNode {
	s.InstanceIds = v
	return s
}

func (s *AttachNodesRequestComputeNode) Validate() error {
	return dara.Validate(s)
}
