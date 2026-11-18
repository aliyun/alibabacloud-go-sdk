// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateAppInfoRequest
	GetAppId() *string
	SetAppName(v string) *UpdateAppInfoRequest
	GetAppName() *string
	SetDescription(v string) *UpdateAppInfoRequest
	GetDescription() *string
	SetStatus(v string) *UpdateAppInfoRequest
	GetStatus() *string
}

type UpdateAppInfoRequest struct {
	// The application ID. You can obtain the value from the AppId parameter returned by the [Create an application](https://help.aliyun.com/document_detail/113266.html) or [List applications](https://help.aliyun.com/document_detail/114000.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The new name of the application.
	//
	// - The name can be up to 128 characters in length and can contain only Chinese characters, uppercase and lowercase letters, digits, periods (.), hyphens (-), and at signs (@).
	//
	// - The name must be encoded in UTF-8.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The new description of the application.
	//
	// - The description can be up to 512 characters in length.
	//
	// - The description must be encoded in UTF-8.
	//
	// example:
	//
	// my first app.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new status of the application. Valid values:
	//
	// - **Normal**: Enabled.
	//
	// - **Disable**: Disabled.
	//
	// example:
	//
	// Disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAppInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAppInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAppInfoRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateAppInfoRequest) SetAppId(v string) *UpdateAppInfoRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppInfoRequest) SetAppName(v string) *UpdateAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAppInfoRequest) SetDescription(v string) *UpdateAppInfoRequest {
	s.Description = &v
	return s
}

func (s *UpdateAppInfoRequest) SetStatus(v string) *UpdateAppInfoRequest {
	s.Status = &v
	return s
}

func (s *UpdateAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
