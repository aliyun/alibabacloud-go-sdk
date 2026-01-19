// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationRole(v *GetApplicationRoleResponseBodyApplicationRole) *GetApplicationRoleResponseBody
	GetApplicationRole() *GetApplicationRoleResponseBodyApplicationRole
	SetRequestId(v string) *GetApplicationRoleResponseBody
	GetRequestId() *string
}

type GetApplicationRoleResponseBody struct {
	ApplicationRole *GetApplicationRoleResponseBodyApplicationRole `json:"ApplicationRole,omitempty" xml:"ApplicationRole,omitempty" type:"Struct"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationRoleResponseBody) GetApplicationRole() *GetApplicationRoleResponseBodyApplicationRole {
	return s.ApplicationRole
}

func (s *GetApplicationRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationRoleResponseBody) SetApplicationRole(v *GetApplicationRoleResponseBodyApplicationRole) *GetApplicationRoleResponseBody {
	s.ApplicationRole = v
	return s
}

func (s *GetApplicationRoleResponseBody) SetRequestId(v string) *GetApplicationRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationRoleResponseBody) Validate() error {
	if s.ApplicationRole != nil {
		if err := s.ApplicationRole.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationRoleResponseBodyApplicationRole struct {
	// 应用唯一标识
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用角色的唯一标识
	//
	// example:
	//
	// approle_01kh2vuo8v9splv8maak1d22rxxxx
	ApplicationRoleId *string `json:"ApplicationRoleId,omitempty" xml:"ApplicationRoleId,omitempty"`
	// 应用角色名称
	//
	// example:
	//
	// Admin Role
	ApplicationRoleName *string `json:"ApplicationRoleName,omitempty" xml:"ApplicationRoleName,omitempty"`
	// 应用角色值
	//
	// example:
	//
	// admin_role
	ApplicationRoleValue *string `json:"ApplicationRoleValue,omitempty" xml:"ApplicationRoleValue,omitempty"`
	// 应用角色描述
	//
	// example:
	//
	// Admin Role Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// EIAM 实例唯一标识
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationRoleResponseBodyApplicationRole) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRoleResponseBodyApplicationRole) GoString() string {
	return s.String()
}

func (s *GetApplicationRoleResponseBodyApplicationRole) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationRoleResponseBodyApplicationRole) GetApplicationRoleId() *string {
	return s.ApplicationRoleId
}

func (s *GetApplicationRoleResponseBodyApplicationRole) GetApplicationRoleName() *string {
	return s.ApplicationRoleName
}

func (s *GetApplicationRoleResponseBodyApplicationRole) GetApplicationRoleValue() *string {
	return s.ApplicationRoleValue
}

func (s *GetApplicationRoleResponseBodyApplicationRole) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationRoleResponseBodyApplicationRole) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationRoleResponseBodyApplicationRole) SetApplicationId(v string) *GetApplicationRoleResponseBodyApplicationRole {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationRoleResponseBodyApplicationRole) SetApplicationRoleId(v string) *GetApplicationRoleResponseBodyApplicationRole {
	s.ApplicationRoleId = &v
	return s
}

func (s *GetApplicationRoleResponseBodyApplicationRole) SetApplicationRoleName(v string) *GetApplicationRoleResponseBodyApplicationRole {
	s.ApplicationRoleName = &v
	return s
}

func (s *GetApplicationRoleResponseBodyApplicationRole) SetApplicationRoleValue(v string) *GetApplicationRoleResponseBodyApplicationRole {
	s.ApplicationRoleValue = &v
	return s
}

func (s *GetApplicationRoleResponseBodyApplicationRole) SetDescription(v string) *GetApplicationRoleResponseBodyApplicationRole {
	s.Description = &v
	return s
}

func (s *GetApplicationRoleResponseBodyApplicationRole) SetInstanceId(v string) *GetApplicationRoleResponseBodyApplicationRole {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationRoleResponseBodyApplicationRole) Validate() error {
	return dara.Validate(s)
}
