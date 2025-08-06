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
	// The ID of the application.
	//
	// 	- Default value: **app-1000000**.
	//
	// 	- For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// 	- The name can contain up to 128 characters in length, including Chinese letters, digits, and periods (.), dash (-), and at character (@).
	//
	// 	- The name can contain only UTF-8 characters.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// 	- The description can contain up to 512 characters in length.
	//
	// 	- The description can contain only UTF-8 characters.
	//
	// example:
	//
	// my first app.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Disable**
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
