// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyNodePoolAttributeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ModifyNodePoolAttributeShrinkRequest
	GetBizRegionId() *string
	SetNodeCapacity(v int32) *ModifyNodePoolAttributeShrinkRequest
	GetNodeCapacity() *int32
	SetNodePoolStrategyShrink(v string) *ModifyNodePoolAttributeShrinkRequest
	GetNodePoolStrategyShrink() *string
	SetPoolId(v string) *ModifyNodePoolAttributeShrinkRequest
	GetPoolId() *string
	SetProductType(v string) *ModifyNodePoolAttributeShrinkRequest
	GetProductType() *string
}

type ModifyNodePoolAttributeShrinkRequest struct {
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// 2
	NodeCapacity           *int32  `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	NodePoolStrategyShrink *string `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty"`
	// example:
	//
	// rg-ew7va2g1wl3vm****
	PoolId *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// 产品类型。
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ModifyNodePoolAttributeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyNodePoolAttributeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyNodePoolAttributeShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ModifyNodePoolAttributeShrinkRequest) GetNodeCapacity() *int32 {
	return s.NodeCapacity
}

func (s *ModifyNodePoolAttributeShrinkRequest) GetNodePoolStrategyShrink() *string {
	return s.NodePoolStrategyShrink
}

func (s *ModifyNodePoolAttributeShrinkRequest) GetPoolId() *string {
	return s.PoolId
}

func (s *ModifyNodePoolAttributeShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetBizRegionId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodeCapacity(v int32) *ModifyNodePoolAttributeShrinkRequest {
	s.NodeCapacity = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetNodePoolStrategyShrink(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.NodePoolStrategyShrink = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetPoolId(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.PoolId = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) SetProductType(v string) *ModifyNodePoolAttributeShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyNodePoolAttributeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
