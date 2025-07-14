// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateAppSecurityGroupRequest
	GetAppId() *string
	SetSecurityGroupId(v string) *UpdateAppSecurityGroupRequest
	GetSecurityGroupId() *string
}

type UpdateAppSecurityGroupRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 017f39b8-dfa4-4e16-a84b-1dcee4b1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s UpdateAppSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppSecurityGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAppSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateAppSecurityGroupRequest) SetAppId(v string) *UpdateAppSecurityGroupRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAppSecurityGroupRequest) SetSecurityGroupId(v string) *UpdateAppSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateAppSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
