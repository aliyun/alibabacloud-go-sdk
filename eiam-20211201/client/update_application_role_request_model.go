// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *UpdateApplicationRoleRequest
	GetApplicationId() *string
	SetApplicationRoleId(v string) *UpdateApplicationRoleRequest
	GetApplicationRoleId() *string
	SetApplicationRoleName(v string) *UpdateApplicationRoleRequest
	GetApplicationRoleName() *string
	SetClientToken(v string) *UpdateApplicationRoleRequest
	GetClientToken() *string
	SetInstanceId(v string) *UpdateApplicationRoleRequest
	GetInstanceId() *string
}

type UpdateApplicationRoleRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// approle_01kh2vuo8v9splv8maak1d22rxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// 应用角色名称
	//
	// This parameter is required.
	//
	// example:
	//
	// Admin Role
	ApplicationRoleName *string `json:"ApplicationRoleName,omitempty" xml:"ApplicationRoleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateApplicationRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRoleRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdateApplicationRoleRequest) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *UpdateApplicationRoleRequest) GetApplicationRoleName() *string {
	return s.ApplicationRoleName
}

func (s *UpdateApplicationRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateApplicationRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateApplicationRoleRequest) SetApplicationId(v string) *UpdateApplicationRoleRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdateApplicationRoleRequest) SetApplicationRoleId(v string) *UpdateApplicationRoleRequest {
	s.ApplicationRoleId = &v
	return s
}

func (s *UpdateApplicationRoleRequest) SetApplicationRoleName(v string) *UpdateApplicationRoleRequest {
	s.ApplicationRoleName = &v
	return s
}

func (s *UpdateApplicationRoleRequest) SetClientToken(v string) *UpdateApplicationRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateApplicationRoleRequest) SetInstanceId(v string) *UpdateApplicationRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateApplicationRoleRequest) Validate() error {
	return dara.Validate(s)
}
