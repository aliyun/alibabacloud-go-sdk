// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateServiceLinkedRoleResponseBody
	GetRequestId() *string
	SetRole(v *CreateServiceLinkedRoleResponseBodyRole) *CreateServiceLinkedRoleResponseBody
	GetRole() *CreateServiceLinkedRoleResponseBodyRole
}

type CreateServiceLinkedRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE58D7CF-03BC-432A-B42D-BC3390C8C2E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the role.
	Role *CreateServiceLinkedRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleResponseBody) GetRole() *CreateServiceLinkedRoleResponseBodyRole {
	return s.Role
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetRole(v *CreateServiceLinkedRoleResponseBodyRole) *CreateServiceLinkedRoleResponseBody {
	s.Role = v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) Validate() error {
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceLinkedRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	//
	// example:
	//
	// acs:ram::177242285274****:role/aliyunserviceroleforpolardb
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the trust policy for the role.
	//
	// example:
	//
	// {\\"Statement\\":[{\\"Action\\":\\"sts:AssumeRole\\",\\"Effect\\":\\"Allow\\",\\"Principal\\":{\\"Service\\":[\\"polardb.aliyuncs.com\\"]}}],\\"Version\\":\\"1\\"}
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the role was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-30T08:14:16Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the role.
	//
	// example:
	//
	// Service Linked Role for PolarDB. PolarDB will use this role to access your resources in other services.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the role is a service-linked role. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsServiceLinkedRole *bool `json:"IsServiceLinkedRole,omitempty" xml:"IsServiceLinkedRole,omitempty"`
	// The ID of the role.
	//
	// example:
	//
	// 32833240981067****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// AliyunServiceRoleForPolarDB
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The role name that uses a domain name as the suffix.
	//
	// example:
	//
	// AliyunServiceRoleForPolarDB@role.test.onaliyunservice.com
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBodyRole) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetArn() *string {
	return s.Arn
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetDescription() *string {
	return s.Description
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetIsServiceLinkedRole() *bool {
	return s.IsServiceLinkedRole
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetRoleId() *string {
	return s.RoleId
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateServiceLinkedRoleResponseBodyRole) GetRolePrincipalName() *string {
	return s.RolePrincipalName
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetArn(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetCreateDate(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetDescription(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetIsServiceLinkedRole(v bool) *CreateServiceLinkedRoleResponseBodyRole {
	s.IsServiceLinkedRole = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetRoleId(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetRoleName(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) SetRolePrincipalName(v string) *CreateServiceLinkedRoleResponseBodyRole {
	s.RolePrincipalName = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyRole) Validate() error {
	return dara.Validate(s)
}
