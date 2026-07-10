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
	// The list of security alert rules.
	ControlStrategyListShrink *string `json:"ControlStrategyList,omitempty" xml:"ControlStrategyList,omitempty"`
	// The product type. Currently, only **ANT_CLOUD_AUTH*	- (financial-grade ID Verification) is supported. All other types have been discontinued.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID of the Smart Access Gateway instance.
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
