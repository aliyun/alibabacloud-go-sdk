// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrepayInstanceSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHa(v bool) *ModifyPrepayInstanceSpecShrinkRequest
	GetHa() *bool
	SetHaResourceSpecShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest
	GetHaResourceSpecShrink() *string
	SetHaVSwitchIdsShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest
	GetHaVSwitchIdsShrink() *string
	SetHaZoneId(v string) *ModifyPrepayInstanceSpecShrinkRequest
	GetHaZoneId() *string
	SetInstanceId(v string) *ModifyPrepayInstanceSpecShrinkRequest
	GetInstanceId() *string
	SetRegion(v string) *ModifyPrepayInstanceSpecShrinkRequest
	GetRegion() *string
	SetResourceSpecShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest
	GetResourceSpecShrink() *string
}

type ModifyPrepayInstanceSpecShrinkRequest struct {
	// if can be null:
	// true
	Ha *bool `json:"Ha,omitempty" xml:"Ha,omitempty"`
	// if can be null:
	// true
	HaResourceSpecShrink *string `json:"HaResourceSpec,omitempty" xml:"HaResourceSpec,omitempty"`
	// if can be null:
	// true
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// if can be null:
	// true
	HaZoneId *string `json:"HaZoneId,omitempty" xml:"HaZoneId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	ResourceSpecShrink *string `json:"ResourceSpec,omitempty" xml:"ResourceSpec,omitempty"`
}

func (s ModifyPrepayInstanceSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrepayInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetHa() *bool {
	return s.Ha
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetHaResourceSpecShrink() *string {
	return s.HaResourceSpecShrink
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetHaVSwitchIdsShrink() *string {
	return s.HaVSwitchIdsShrink
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetHaZoneId() *string {
	return s.HaZoneId
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) GetResourceSpecShrink() *string {
	return s.ResourceSpecShrink
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHa(v bool) *ModifyPrepayInstanceSpecShrinkRequest {
	s.Ha = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHaResourceSpecShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.HaResourceSpecShrink = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHaVSwitchIdsShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetHaZoneId(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.HaZoneId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetInstanceId(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetRegion(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.Region = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) SetResourceSpecShrink(v string) *ModifyPrepayInstanceSpecShrinkRequest {
	s.ResourceSpecShrink = &v
	return s
}

func (s *ModifyPrepayInstanceSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
