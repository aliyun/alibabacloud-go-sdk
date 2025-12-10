// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewAssumeRolePolicyDocument(v string) *UpdateRoleRequest
	GetNewAssumeRolePolicyDocument() *string
	SetNewDescription(v string) *UpdateRoleRequest
	GetNewDescription() *string
	SetNewMaxSessionDuration(v int64) *UpdateRoleRequest
	GetNewMaxSessionDuration() *int64
	SetRoleName(v string) *UpdateRoleRequest
	GetRoleName() *string
}

type UpdateRoleRequest struct {
	// The trust policy of the RAM role.
	//
	// example:
	//
	// { "Statement": [ { "Action": "sts:AssumeRole", "Effect": "Allow", "Principal": { "RAM": "acs:ram::12345678901234****:root" } } ], "Version": "1" }
	NewAssumeRolePolicyDocument *string `json:"NewAssumeRolePolicyDocument,omitempty" xml:"NewAssumeRolePolicyDocument,omitempty"`
	// The description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// ECS administrator
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The maximum session time of the RAM role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 3600
	NewMaxSessionDuration *int64 `json:"NewMaxSessionDuration,omitempty" xml:"NewMaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) GetNewAssumeRolePolicyDocument() *string {
	return s.NewAssumeRolePolicyDocument
}

func (s *UpdateRoleRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateRoleRequest) GetNewMaxSessionDuration() *int64 {
	return s.NewMaxSessionDuration
}

func (s *UpdateRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateRoleRequest) SetNewAssumeRolePolicyDocument(v string) *UpdateRoleRequest {
	s.NewAssumeRolePolicyDocument = &v
	return s
}

func (s *UpdateRoleRequest) SetNewDescription(v string) *UpdateRoleRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateRoleRequest) SetNewMaxSessionDuration(v int64) *UpdateRoleRequest {
	s.NewMaxSessionDuration = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleRequest) Validate() error {
	return dara.Validate(s)
}
