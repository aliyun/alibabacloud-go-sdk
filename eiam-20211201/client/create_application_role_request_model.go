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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application role.
	//
	// This parameter is required.
	//
	// example:
	//
	// Admin Role
	ApplicationRoleName *string `json:"ApplicationRoleName,omitempty" xml:"ApplicationRoleName,omitempty"`
	// The value of the application role.
	//
	// This parameter is required.
	//
	// example:
	//
	// admin_role
	ApplicationRoleValue *string `json:"ApplicationRoleValue,omitempty" xml:"ApplicationRoleValue,omitempty"`
	// A client token used to ensure the idempotence of the request. Generate a unique value for this parameter from your client. The client token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
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
