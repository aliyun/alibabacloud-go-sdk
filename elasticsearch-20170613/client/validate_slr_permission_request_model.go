// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateSlrPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ValidateSlrPermissionRequest
	GetClientToken() *string
	SetRolename(v string) *ValidateSlrPermissionRequest
	GetRolename() *string
}

type ValidateSlrPermissionRequest struct {
	// A client-generated token used to ensure the idempotence of the request. The value must be unique across different requests and cannot exceed 64 ASCII characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the service-linked role. Valid values:
	//
	// <props="china">- AliyunServiceRoleForElasticsearchOps: used to perform elastic scaling tasks for clusters- AliyunServiceRoleForElasticsearchCollector: used to create and manage Beats collectors
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForElasticsearchCollector
	Rolename *string `json:"rolename,omitempty" xml:"rolename,omitempty"`
}

func (s ValidateSlrPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateSlrPermissionRequest) GoString() string {
	return s.String()
}

func (s *ValidateSlrPermissionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValidateSlrPermissionRequest) GetRolename() *string {
	return s.Rolename
}

func (s *ValidateSlrPermissionRequest) SetClientToken(v string) *ValidateSlrPermissionRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateSlrPermissionRequest) SetRolename(v string) *ValidateSlrPermissionRequest {
	s.Rolename = &v
	return s
}

func (s *ValidateSlrPermissionRequest) Validate() error {
	return dara.Validate(s)
}
