// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLiveInputShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputId(v string) *UpdateMediaLiveInputShrinkRequest
	GetInputId() *string
	SetInputSettingsShrink(v string) *UpdateMediaLiveInputShrinkRequest
	GetInputSettingsShrink() *string
	SetName(v string) *UpdateMediaLiveInputShrinkRequest
	GetName() *string
	SetSecurityGroupIdsShrink(v string) *UpdateMediaLiveInputShrinkRequest
	GetSecurityGroupIdsShrink() *string
}

type UpdateMediaLiveInputShrinkRequest struct {
	// The ID of the input.
	//
	// This parameter is required.
	//
	// example:
	//
	// SEGK5KA6KYKAWQQH
	InputId *string `json:"InputId,omitempty" xml:"InputId,omitempty"`
	// The input settings. An input can have up to two sources: primary and backup sources.
	//
	// This parameter is required.
	InputSettingsShrink *string `json:"InputSettings,omitempty" xml:"InputSettings,omitempty"`
	// The name of the input. Letters, digits, hyphens (-), and underscores (_) are supported. It can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// myinput
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the security groups to be associated with the input. This parameter is required for PUSH inputs.
	//
	// example:
	//
	// ["G6G4X5T4SZYPSTT5"]
	SecurityGroupIdsShrink *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
}

func (s UpdateMediaLiveInputShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLiveInputShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMediaLiveInputShrinkRequest) GetInputId() *string {
	return s.InputId
}

func (s *UpdateMediaLiveInputShrinkRequest) GetInputSettingsShrink() *string {
	return s.InputSettingsShrink
}

func (s *UpdateMediaLiveInputShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMediaLiveInputShrinkRequest) GetSecurityGroupIdsShrink() *string {
	return s.SecurityGroupIdsShrink
}

func (s *UpdateMediaLiveInputShrinkRequest) SetInputId(v string) *UpdateMediaLiveInputShrinkRequest {
	s.InputId = &v
	return s
}

func (s *UpdateMediaLiveInputShrinkRequest) SetInputSettingsShrink(v string) *UpdateMediaLiveInputShrinkRequest {
	s.InputSettingsShrink = &v
	return s
}

func (s *UpdateMediaLiveInputShrinkRequest) SetName(v string) *UpdateMediaLiveInputShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateMediaLiveInputShrinkRequest) SetSecurityGroupIdsShrink(v string) *UpdateMediaLiveInputShrinkRequest {
	s.SecurityGroupIdsShrink = &v
	return s
}

func (s *UpdateMediaLiveInputShrinkRequest) Validate() error {
	return dara.Validate(s)
}
