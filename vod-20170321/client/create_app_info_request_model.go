// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppInfoRequest
	GetAppName() *string
	SetDescription(v string) *CreateAppInfoRequest
	GetDescription() *string
	SetResourceGroupId(v string) *CreateAppInfoRequest
	GetResourceGroupId() *string
}

type CreateAppInfoRequest struct {
	// The name of the application. The application name must be unique.
	//
	// 	- The name can contain letters, digits, periods (.), hyphens (-), and at signs (@). The name can be up to 128 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// 	- The description can contain up to 512 characters in length.
	//
	// 	- The value must be encoded in UTF-8.
	//
	// example:
	//
	// myfirstapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzko7fsuj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateAppInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateAppInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppInfoRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateAppInfoRequest) SetAppName(v string) *CreateAppInfoRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppInfoRequest) SetDescription(v string) *CreateAppInfoRequest {
	s.Description = &v
	return s
}

func (s *CreateAppInfoRequest) SetResourceGroupId(v string) *CreateAppInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAppInfoRequest) Validate() error {
	return dara.Validate(s)
}
