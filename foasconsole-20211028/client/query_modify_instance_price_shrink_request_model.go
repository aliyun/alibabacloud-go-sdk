// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryModifyInstancePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *QueryModifyInstancePriceShrinkRequest
	GetHa() *bool
	SetHaResourceSpecShrink(v string) *QueryModifyInstancePriceShrinkRequest
	GetHaResourceSpecShrink() *string
	SetHaVSwitchIdsShrink(v string) *QueryModifyInstancePriceShrinkRequest
	GetHaVSwitchIdsShrink() *string
	SetInstanceId(v string) *QueryModifyInstancePriceShrinkRequest
	GetInstanceId() *string
	SetPromotionCode(v string) *QueryModifyInstancePriceShrinkRequest
	GetPromotionCode() *string
	SetRegion(v string) *QueryModifyInstancePriceShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *QueryModifyInstancePriceShrinkRequest
	GetResourceSpecShrink() *string
	SetUsePromotionCode(v bool) *QueryModifyInstancePriceShrinkRequest
	GetUsePromotionCode() *bool
}

type QueryModifyInstancePriceShrinkRequest struct {
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// if can be null:
	// true
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
	UsePromotionCode   *bool   `json:"UsePromotionCode,omitempty" xml:"UsePromotionCode,omitempty"`
}

func (s QueryModifyInstancePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryModifyInstancePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryModifyInstancePriceShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *QueryModifyInstancePriceShrinkRequest) GetHaResourceSpecShrink() *string {
	return s.HaResourceSpecShrink
}

func (s *QueryModifyInstancePriceShrinkRequest) GetHaVSwitchIdsShrink() *string {
	return s.HaVSwitchIdsShrink
}

func (s *QueryModifyInstancePriceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryModifyInstancePriceShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *QueryModifyInstancePriceShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *QueryModifyInstancePriceShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *QueryModifyInstancePriceShrinkRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHa(v bool) *QueryModifyInstancePriceShrinkRequest {
	s.Ha = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHaResourceSpecShrink(v string) *QueryModifyInstancePriceShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetHaVSwitchIdsShrink(v string) *QueryModifyInstancePriceShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetInstanceId(v string) *QueryModifyInstancePriceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetPromotionCode(v string) *QueryModifyInstancePriceShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetRegion(v string) *QueryModifyInstancePriceShrinkRequest {
	s.Region = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetResourceSpecShrink(v string) *QueryModifyInstancePriceShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) SetUsePromotionCode(v bool) *QueryModifyInstancePriceShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *QueryModifyInstancePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
