// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLiveInputShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputSettingsShrink(v string) *CreateMediaLiveInputShrinkRequest
	GetInputSettingsShrink() *string
	SetName(v string) *CreateMediaLiveInputShrinkRequest
	GetName() *string
	SetSecurityGroupIdsShrink(v string) *CreateMediaLiveInputShrinkRequest
	GetSecurityGroupIdsShrink() *string
	SetType(v string) *CreateMediaLiveInputShrinkRequest
	GetType() *string
}

type CreateMediaLiveInputShrinkRequest struct {
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
	// The input type. Valid values: RTMP_PUSH, RTMP_PULL, SRT_PUSH, SRT_PULL, and MEDIA_CONNECT.
	//
	// This parameter is required.
	//
	// example:
	//
	// RTMP_PUSH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateMediaLiveInputShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLiveInputShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaLiveInputShrinkRequest) GetInputSettingsShrink() *string {
	return s.InputSettingsShrink
}

func (s *CreateMediaLiveInputShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateMediaLiveInputShrinkRequest) GetSecurityGroupIdsShrink() *string {
	return s.SecurityGroupIdsShrink
}

func (s *CreateMediaLiveInputShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateMediaLiveInputShrinkRequest) SetInputSettingsShrink(v string) *CreateMediaLiveInputShrinkRequest {
	s.InputSettingsShrink = &v
	return s
}

func (s *CreateMediaLiveInputShrinkRequest) SetName(v string) *CreateMediaLiveInputShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateMediaLiveInputShrinkRequest) SetSecurityGroupIdsShrink(v string) *CreateMediaLiveInputShrinkRequest {
	s.SecurityGroupIdsShrink = &v
	return s
}

func (s *CreateMediaLiveInputShrinkRequest) SetType(v string) *CreateMediaLiveInputShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateMediaLiveInputShrinkRequest) Validate() error {
	return dara.Validate(s)
}
