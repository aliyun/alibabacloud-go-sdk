// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRoleResponseBody
	GetRequestId() *string
	SetRole(v *CreateRoleResponseBodyRole) *CreateRoleResponseBody
	GetRole() *CreateRoleResponseBodyRole
}

type CreateRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the role.
	Role *CreateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s CreateRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoleResponseBody) GetRole() *CreateRoleResponseBodyRole {
	return s.Role
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) SetRole(v *CreateRoleResponseBodyRole) *CreateRoleResponseBody {
	s.Role = v
	return s
}

func (s *CreateRoleResponseBody) Validate() error {
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The document of the policy in which the entity that can assume the role is specified.
	//
	// example:
	//
	// { \\"Statement\\": [ { \\"Action\\": \\"sts:AssumeRole\\", \\"Effect\\": \\"Allow\\", \\"Principal\\": { \\"RAM\\": \\"acs:ram::12345678901234****:root\\" } } ], \\"Version\\": \\"1\\" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the role.
	//
	// example:
	//
	// 90123456789****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the role after authorization.
	//
	// example:
	//
	// ECSAdmin@role.123456.onaliyunservice.com
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
}

func (s CreateRoleResponseBodyRole) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBodyRole) GetArn() *string {
	return s.Arn
}

func (s *CreateRoleResponseBodyRole) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *CreateRoleResponseBodyRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateRoleResponseBodyRole) GetDescription() *string {
	return s.Description
}

func (s *CreateRoleResponseBodyRole) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *CreateRoleResponseBodyRole) GetRoleId() *string {
	return s.RoleId
}

func (s *CreateRoleResponseBodyRole) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleResponseBodyRole) GetRolePrincipalName() *string {
	return s.RolePrincipalName
}

func (s *CreateRoleResponseBodyRole) SetArn(v string) *CreateRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *CreateRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetCreateDate(v string) *CreateRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetDescription(v string) *CreateRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetMaxSessionDuration(v int64) *CreateRoleResponseBodyRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRoleId(v string) *CreateRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRoleName(v string) *CreateRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *CreateRoleResponseBodyRole) SetRolePrincipalName(v string) *CreateRoleResponseBodyRole {
	s.RolePrincipalName = &v
	return s
}

func (s *CreateRoleResponseBodyRole) Validate() error {
	return dara.Validate(s)
}
