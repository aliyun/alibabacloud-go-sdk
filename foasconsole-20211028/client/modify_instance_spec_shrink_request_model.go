// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *ModifyInstanceSpecShrinkRequest
	GetHa() *bool
	SetHaResourceSpecShrink(v string) *ModifyInstanceSpecShrinkRequest
	GetHaResourceSpecShrink() *string
	SetHaVSwitchIdsShrink(v string) *ModifyInstanceSpecShrinkRequest
	GetHaVSwitchIdsShrink() *string
	SetInstanceId(v string) *ModifyInstanceSpecShrinkRequest
	GetInstanceId() *string
	SetPromotionCode(v string) *ModifyInstanceSpecShrinkRequest
	GetPromotionCode() *string
	SetRegion(v string) *ModifyInstanceSpecShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *ModifyInstanceSpecShrinkRequest
	GetResourceSpecShrink() *string
	SetUsePromotionCode(v bool) *ModifyInstanceSpecShrinkRequest
	GetUsePromotionCode() *bool
}

type ModifyInstanceSpecShrinkRequest struct {
	// example:
	//
	// true
	Ha                   *bool   `json:"Ha,omitempty" xml:"Ha,omitempty"`
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	HaVSwitchIdsShrink   *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
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

func (s ModifyInstanceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *ModifyInstanceSpecShrinkRequest) GetHaResourceSpecShrink() *string {
	return s.HaResourceSpecShrink
}

func (s *ModifyInstanceSpecShrinkRequest) GetHaVSwitchIdsShrink() *string {
	return s.HaVSwitchIdsShrink
}

func (s *ModifyInstanceSpecShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSpecShrinkRequest) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *ModifyInstanceSpecShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyInstanceSpecShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *ModifyInstanceSpecShrinkRequest) GetUsePromotionCode() *bool {
	return s.UsePromotionCode
}

func (s *ModifyInstanceSpecShrinkRequest) SetHa(v bool) *ModifyInstanceSpecShrinkRequest {
	s.Ha = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetHaResourceSpecShrink(v string) *ModifyInstanceSpecShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetHaVSwitchIdsShrink(v string) *ModifyInstanceSpecShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetInstanceId(v string) *ModifyInstanceSpecShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetPromotionCode(v string) *ModifyInstanceSpecShrinkRequest {
	s.PromotionCode = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetRegion(v string) *ModifyInstanceSpecShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetResourceSpecShrink(v string) *ModifyInstanceSpecShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) SetUsePromotionCode(v bool) *ModifyInstanceSpecShrinkRequest {
	s.UsePromotionCode = &v
	return s
}

func (s *ModifyInstanceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
