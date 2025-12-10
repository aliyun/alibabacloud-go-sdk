// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeRolePolicyDocument(v string) *CreateRoleRequest
	GetAssumeRolePolicyDocument() *string
	SetDescription(v string) *CreateRoleRequest
	GetDescription() *string
	SetMaxSessionDuration(v int64) *CreateRoleRequest
	GetMaxSessionDuration() *int64
	SetRoleName(v string) *CreateRoleRequest
	GetRoleName() *string
}

type CreateRoleRequest struct {
	// The document of the policy that specifies
	//
	// one or more trusted entities to assume the role. The trusted entities can be Alibaba Cloud accounts, Alibaba Cloud services, or identity providers (IdPs).
	//
	// >  RAM users cannot assume the RAM roles of trusted Alibaba Cloud services.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "Statement": [ { "Action": "sts:AssumeRole", "Effect": "Allow", "Principal": { "RAM": "acs:ram::12345678901234****:root" } } ], "Version": "1" }
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The description of the role.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The name of the role.
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

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *CreateRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoleRequest) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *CreateRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleRequest) SetAssumeRolePolicyDocument(v string) *CreateRoleRequest {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleRequest) SetDescription(v string) *CreateRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleRequest) SetMaxSessionDuration(v int64) *CreateRoleRequest {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	return dara.Validate(s)
}
