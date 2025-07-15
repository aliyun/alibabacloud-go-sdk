// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppConfigShrink(v string) *UpdateMessageAppShrinkRequest
	GetAppConfigShrink() *string
	SetAppId(v string) *UpdateMessageAppShrinkRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMessageAppShrinkRequest
	GetAppName() *string
	SetExtensionShrink(v string) *UpdateMessageAppShrinkRequest
	GetExtensionShrink() *string
}

type UpdateMessageAppShrinkRequest struct {
	// The configurations of the application.
	AppConfigShrink *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The ID of the interactive messaging application.
	//
	// This parameter is required.
	//
	// example:
	//
	// VKL3***
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the interactive messaging application.
	//
	// example:
	//
	// testApp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The extended field.
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s UpdateMessageAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageAppShrinkRequest) GetAppConfigShrink() *string {
	return s.AppConfigShrink
}

func (s *UpdateMessageAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMessageAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMessageAppShrinkRequest) GetExtensionShrink() *string {
	return s.ExtensionShrink
}

func (s *UpdateMessageAppShrinkRequest) SetAppConfigShrink(v string) *UpdateMessageAppShrinkRequest {
	s.AppConfigShrink = &v
	return s
}

func (s *UpdateMessageAppShrinkRequest) SetAppId(v string) *UpdateMessageAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMessageAppShrinkRequest) SetAppName(v string) *UpdateMessageAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMessageAppShrinkRequest) SetExtensionShrink(v string) *UpdateMessageAppShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *UpdateMessageAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
