// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationRoleRequest
	GetApplicationId() *string
	SetApplicationRoleName(v string) *CreateApplicationRoleRequest
	GetApplicationRoleName() *string
	SetApplicationRoleValue(v string) *CreateApplicationRoleRequest
	GetApplicationRoleValue() *string
	SetClientToken(v string) *CreateApplicationRoleRequest
	GetClientToken() *string
	SetInstanceId(v string) *CreateApplicationRoleRequest
	GetInstanceId() *string
}

type CreateApplicationRoleRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 管理员角色
	ApplicationRoleName *string `json:"ApplicationRoleName,omitempty" xml:"ApplicationRoleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// admin_role
	ApplicationRoleValue *string `json:"ApplicationRoleValue,omitempty" xml:"ApplicationRoleValue,omitempty"`
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

func (s CreateApplicationRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRoleRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationRoleRequest) GetApplicationRoleName() *string {
	return s.ApplicationRoleName
}

func (s *CreateApplicationRoleRequest) GetApplicationRoleValue() *string {
	return s.ApplicationRoleValue
}

func (s *CreateApplicationRoleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApplicationRoleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApplicationRoleRequest) SetApplicationId(v string) *CreateApplicationRoleRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationRoleRequest) SetApplicationRoleName(v string) *CreateApplicationRoleRequest {
	s.ApplicationRoleName = &v
	return s
}

func (s *CreateApplicationRoleRequest) SetApplicationRoleValue(v string) *CreateApplicationRoleRequest {
	s.ApplicationRoleValue = &v
	return s
}

func (s *CreateApplicationRoleRequest) SetClientToken(v string) *CreateApplicationRoleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationRoleRequest) SetInstanceId(v string) *CreateApplicationRoleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApplicationRoleRequest) Validate() error {
	return dara.Validate(s)
}
