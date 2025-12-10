// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRoleResponseBody
	GetRequestId() *string
	SetRole(v *UpdateRoleResponseBodyRole) *UpdateRoleResponseBody
	GetRole() *UpdateRoleResponseBodyRole
}

type UpdateRoleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM role.
	Role *UpdateRoleResponseBodyRole `json:"Role,omitempty" xml:"Role,omitempty" type:"Struct"`
}

func (s UpdateRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRoleResponseBody) GetRole() *UpdateRoleResponseBodyRole {
	return s.Role
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRole(v *UpdateRoleResponseBodyRole) *UpdateRoleResponseBody {
	s.Role = v
	return s
}

func (s *UpdateRoleResponseBody) Validate() error {
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRoleResponseBodyRole struct {
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/ECSAdmin
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The trust policy of the RAM role.
	//
	// example:
	//
	// { \\"Statement\\": [ { \\"Action\\": \\"sts:AssumeRole\\", \\"Effect\\": \\"Allow\\", \\"Principal\\": { \\"RAM\\": \\"acs:ram::12345678901234****:root\\" } } ], \\"Version\\": \\"1\\" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The time when the RAM role was created.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the RAM role.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session time of the RAM role.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The ID of the RAM role.
	//
	// example:
	//
	// 90123456789****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the RAM role after authorization.
	//
	// example:
	//
	// ECSAdmin@role.123456.onaliyunservice.com
	RolePrincipalName *string `json:"RolePrincipalName,omitempty" xml:"RolePrincipalName,omitempty"`
	// The time when the RAM role was updated.
	//
	// example:
	//
	// 2016-01-23T12:33:18Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateRoleResponseBodyRole) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleResponseBodyRole) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBodyRole) GetArn() *string {
	return s.Arn
}

func (s *UpdateRoleResponseBodyRole) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *UpdateRoleResponseBodyRole) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateRoleResponseBodyRole) GetDescription() *string {
	return s.Description
}

func (s *UpdateRoleResponseBodyRole) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *UpdateRoleResponseBodyRole) GetRoleId() *string {
	return s.RoleId
}

func (s *UpdateRoleResponseBodyRole) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateRoleResponseBodyRole) GetRolePrincipalName() *string {
	return s.RolePrincipalName
}

func (s *UpdateRoleResponseBodyRole) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateRoleResponseBodyRole) SetArn(v string) *UpdateRoleResponseBodyRole {
	s.Arn = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetAssumeRolePolicyDocument(v string) *UpdateRoleResponseBodyRole {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetCreateDate(v string) *UpdateRoleResponseBodyRole {
	s.CreateDate = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetDescription(v string) *UpdateRoleResponseBodyRole {
	s.Description = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetMaxSessionDuration(v int64) *UpdateRoleResponseBodyRole {
	s.MaxSessionDuration = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetRoleId(v string) *UpdateRoleResponseBodyRole {
	s.RoleId = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetRoleName(v string) *UpdateRoleResponseBodyRole {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetRolePrincipalName(v string) *UpdateRoleResponseBodyRole {
	s.RolePrincipalName = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) SetUpdateDate(v string) *UpdateRoleResponseBodyRole {
	s.UpdateDate = &v
	return s
}

func (s *UpdateRoleResponseBodyRole) Validate() error {
	return dara.Validate(s)
}
