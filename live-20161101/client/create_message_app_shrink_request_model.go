// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppConfigShrink(v string) *CreateMessageAppShrinkRequest
	GetAppConfigShrink() *string
	SetAppName(v string) *CreateMessageAppShrinkRequest
	GetAppName() *string
	SetExtensionShrink(v string) *CreateMessageAppShrinkRequest
	GetExtensionShrink() *string
}

type CreateMessageAppShrinkRequest struct {
	// The configurations of the application.
	AppConfigShrink *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The name of the interactive message application. The name must be 2 to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The extended fields.
	ExtensionShrink *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateMessageAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageAppShrinkRequest) GetAppConfigShrink() *string {
	return s.AppConfigShrink
}

func (s *CreateMessageAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateMessageAppShrinkRequest) GetExtensionShrink() *string {
	return s.ExtensionShrink
}

func (s *CreateMessageAppShrinkRequest) SetAppConfigShrink(v string) *CreateMessageAppShrinkRequest {
	s.AppConfigShrink = &v
	return s
}

func (s *CreateMessageAppShrinkRequest) SetAppName(v string) *CreateMessageAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateMessageAppShrinkRequest) SetExtensionShrink(v string) *CreateMessageAppShrinkRequest {
	s.ExtensionShrink = &v
	return s
}

func (s *CreateMessageAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
