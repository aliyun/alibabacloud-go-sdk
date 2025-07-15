// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppConfig(v map[string]*string) *UpdateMessageAppRequest
	GetAppConfig() map[string]*string
	SetAppId(v string) *UpdateMessageAppRequest
	GetAppId() *string
	SetAppName(v string) *UpdateMessageAppRequest
	GetAppName() *string
	SetExtension(v map[string]*string) *UpdateMessageAppRequest
	GetExtension() map[string]*string
}

type UpdateMessageAppRequest struct {
	// The configurations of the application.
	AppConfig map[string]*string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
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
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s UpdateMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageAppRequest) GoString() string {
	return s.String()
}

func (s *UpdateMessageAppRequest) GetAppConfig() map[string]*string {
	return s.AppConfig
}

func (s *UpdateMessageAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMessageAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateMessageAppRequest) GetExtension() map[string]*string {
	return s.Extension
}

func (s *UpdateMessageAppRequest) SetAppConfig(v map[string]*string) *UpdateMessageAppRequest {
	s.AppConfig = v
	return s
}

func (s *UpdateMessageAppRequest) SetAppId(v string) *UpdateMessageAppRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMessageAppRequest) SetAppName(v string) *UpdateMessageAppRequest {
	s.AppName = &v
	return s
}

func (s *UpdateMessageAppRequest) SetExtension(v map[string]*string) *UpdateMessageAppRequest {
	s.Extension = v
	return s
}

func (s *UpdateMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
