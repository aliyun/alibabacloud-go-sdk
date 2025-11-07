// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlStrategyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlStrategyListShrink(v string) *ModifyControlStrategyShrinkRequest
	GetControlStrategyListShrink() *string
	SetProductType(v string) *ModifyControlStrategyShrinkRequest
	GetProductType() *string
	SetRegionId(v string) *ModifyControlStrategyShrinkRequest
	GetRegionId() *string
}

type ModifyControlStrategyShrinkRequest struct {
	// List of security alarm rules.
	ControlStrategyListShrink *string `json:"ControlStrategyList,omitempty" xml:"ControlStrategyList,omitempty"`
	// Product type, currently only supports **ANT_CLOUD_AUTH*	- (Financial-grade Real Person), all others are phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Region ID of the intelligent access gateway instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyControlStrategyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlStrategyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlStrategyShrinkRequest) GetControlStrategyListShrink() *string {
	return s.ControlStrategyListShrink
}

func (s *ModifyControlStrategyShrinkRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyControlStrategyShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyControlStrategyShrinkRequest) SetControlStrategyListShrink(v string) *ModifyControlStrategyShrinkRequest {
	s.ControlStrategyListShrink = &v
	return s
}

func (s *ModifyControlStrategyShrinkRequest) SetProductType(v string) *ModifyControlStrategyShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyControlStrategyShrinkRequest) SetRegionId(v string) *ModifyControlStrategyShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyControlStrategyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
