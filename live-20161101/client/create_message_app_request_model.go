// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppConfig(v map[string]*string) *CreateMessageAppRequest
	GetAppConfig() map[string]*string
	SetAppName(v string) *CreateMessageAppRequest
	GetAppName() *string
	SetExtension(v map[string]*string) *CreateMessageAppRequest
	GetExtension() map[string]*string
}

type CreateMessageAppRequest struct {
	// The configurations of the application.
	AppConfig map[string]*string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// The name of the interactive message application. The name must be 2 to 16 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The extended fields.
	Extension map[string]*string `json:"Extension,omitempty" xml:"Extension,omitempty"`
}

func (s CreateMessageAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageAppRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageAppRequest) GetAppConfig() map[string]*string {
	return s.AppConfig
}

func (s *CreateMessageAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateMessageAppRequest) GetExtension() map[string]*string {
	return s.Extension
}

func (s *CreateMessageAppRequest) SetAppConfig(v map[string]*string) *CreateMessageAppRequest {
	s.AppConfig = v
	return s
}

func (s *CreateMessageAppRequest) SetAppName(v string) *CreateMessageAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateMessageAppRequest) SetExtension(v map[string]*string) *CreateMessageAppRequest {
	s.Extension = v
	return s
}

func (s *CreateMessageAppRequest) Validate() error {
	return dara.Validate(s)
}
