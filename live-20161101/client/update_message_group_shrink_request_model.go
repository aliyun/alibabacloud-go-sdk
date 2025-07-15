// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMessageGroupShrinkRequest
	GetAppId() *string
	SetExtensionShrink(v string) *UpdateMessageGroupShrinkRequest
	GetExtensionShrink() *string
	SetGroupId(v string) *UpdateMessageGroupShrinkRequest
	GetGroupId() *string
}

type UpdateMessageGroupShrinkRequest struct {
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The extended field.
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// The ID of the message group.
	//
	// This parameter is required.
	//
	// example:
	//
	// AE35-****-T95F
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UpdateMessageGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMessageGroupShrinkRequest) GetExtensionShrink() *string {
	return s.ExtensionShrink
}

func (s *UpdateMessageGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateMessageGroupShrinkRequest) SetAppId(v string) *UpdateMessageGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMessageGroupShrinkRequest) SetExtensionShrink(v string) *UpdateMessageGroupShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *UpdateMessageGroupShrinkRequest) SetGroupId(v string) *UpdateMessageGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateMessageGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
