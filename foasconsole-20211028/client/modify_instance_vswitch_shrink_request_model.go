// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceVswitchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVSwitchIdsShrink(v string) *ModifyInstanceVswitchShrinkRequest
	GetHaVSwitchIdsShrink() *string
	SetInstanceId(v string) *ModifyInstanceVswitchShrinkRequest
	GetInstanceId() *string
	SetVSwitchIdsShrink(v string) *ModifyInstanceVswitchShrinkRequest
	GetVSwitchIdsShrink() *string
}

type ModifyInstanceVswitchShrinkRequest struct {
	HaVSwitchIdsShrink *string `json:"HaVSwitchIds,omitempty" xml:"HaVSwitchIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sc_flinkserverless_public_cn-7e22ae****
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VSwitchIdsShrink *string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
}

func (s ModifyInstanceVswitchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceVswitchShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceVswitchShrinkRequest) GetHaVSwitchIdsShrink() *string {
	return s.HaVSwitchIdsShrink
}

func (s *ModifyInstanceVswitchShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceVswitchShrinkRequest) GetVSwitchIdsShrink() *string {
	return s.VSwitchIdsShrink
}

func (s *ModifyInstanceVswitchShrinkRequest) SetHaVSwitchIdsShrink(v string) *ModifyInstanceVswitchShrinkRequest {
	s.HaVSwitchIdsShrink = &v
	return s
}

func (s *ModifyInstanceVswitchShrinkRequest) SetInstanceId(v string) *ModifyInstanceVswitchShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceVswitchShrinkRequest) SetVSwitchIdsShrink(v string) *ModifyInstanceVswitchShrinkRequest {
	s.VSwitchIdsShrink = &v
	return s
}

func (s *ModifyInstanceVswitchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
