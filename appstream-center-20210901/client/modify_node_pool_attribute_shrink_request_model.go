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
	// The region ID of the delivery group. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The number of concurrent sessions, which is the number of sessions that can be simultaneously connected to a single resource. If too many sessions are connected simultaneously, the application experience may degrade. The valid values vary depending on the resource specification. The valid values for each resource specification are as follows:
	//
	// - appstreaming.general.4c8g: 1 to 2.
	//
	// - appstreaming.general.8c16g: 1 to 4.
	//
	// - appstreaming.vgpu.8c16g.4g: 1 to 4.
	//
	// - appstreaming.vgpu.8c31g.16g: 1 to 4.
	//
	// - appstreaming.vgpu.14c93g.12g: 1 to 6.
	//
	// example:
	//
	// 2
	NodeCapacity *int32 `json:"NodeCapacity,omitempty" xml:"NodeCapacity,omitempty"`
	// The automatic scaling policy of the delivery group.
	NodePoolStrategyShrink *string `json:"NodePoolStrategy,omitempty" xml:"NodePoolStrategy,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-ew7va2g1wl3vm****
	PoolId *string `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// The product type.
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
