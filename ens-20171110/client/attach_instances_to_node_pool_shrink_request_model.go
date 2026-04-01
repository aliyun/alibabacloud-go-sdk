// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachInstancesToNodePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AttachInstancesToNodePoolShrinkRequest
	GetClusterId() *string
	SetInstancesShrink(v string) *AttachInstancesToNodePoolShrinkRequest
	GetInstancesShrink() *string
	SetNodepoolId(v string) *AttachInstancesToNodePoolShrinkRequest
	GetNodepoolId() *string
}

type AttachInstancesToNodePoolShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// np68mi5y1dd748ky37ojo2kqdrz
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s AttachInstancesToNodePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachInstancesToNodePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachInstancesToNodePoolShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AttachInstancesToNodePoolShrinkRequest) GetInstancesShrink() *string {
	return s.InstancesShrink
}

func (s *AttachInstancesToNodePoolShrinkRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *AttachInstancesToNodePoolShrinkRequest) SetClusterId(v string) *AttachInstancesToNodePoolShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *AttachInstancesToNodePoolShrinkRequest) SetInstancesShrink(v string) *AttachInstancesToNodePoolShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *AttachInstancesToNodePoolShrinkRequest) SetNodepoolId(v string) *AttachInstancesToNodePoolShrinkRequest {
	s.NodepoolId = &v
	return s
}

func (s *AttachInstancesToNodePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
