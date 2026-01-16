// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterNodePoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBodyShrink(v string) *ScaleClusterNodePoolShrinkRequest
	GetBodyShrink() *string
	SetClusterId(v string) *ScaleClusterNodePoolShrinkRequest
	GetClusterId() *string
	SetNodepoolId(v string) *ScaleClusterNodePoolShrinkRequest
	GetNodepoolId() *string
}

type ScaleClusterNodePoolShrinkRequest struct {
	BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// np2201da356f5245cf8faa522a8a4c8224
	NodepoolId *string `json:"NodepoolId,omitempty" xml:"NodepoolId,omitempty"`
}

func (s ScaleClusterNodePoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *ScaleClusterNodePoolShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ScaleClusterNodePoolShrinkRequest) GetNodepoolId() *string {
	return s.NodepoolId
}

func (s *ScaleClusterNodePoolShrinkRequest) SetBodyShrink(v string) *ScaleClusterNodePoolShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *ScaleClusterNodePoolShrinkRequest) SetClusterId(v string) *ScaleClusterNodePoolShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ScaleClusterNodePoolShrinkRequest) SetNodepoolId(v string) *ScaleClusterNodePoolShrinkRequest {
	s.NodepoolId = &v
	return s
}

func (s *ScaleClusterNodePoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}
