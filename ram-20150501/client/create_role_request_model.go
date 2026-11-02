// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowConsoleLogin(v bool) *CreateRoleRequest
	GetAllowConsoleLogin() *bool
	SetAssumeRolePolicyDocument(v string) *CreateRoleRequest
	GetAssumeRolePolicyDocument() *string
	SetDescription(v string) *CreateRoleRequest
	GetDescription() *string
	SetMaxSessionDuration(v int64) *CreateRoleRequest
	GetMaxSessionDuration() *int64
	SetRoleName(v string) *CreateRoleRequest
	GetRoleName() *string
	SetTag(v []*CreateRoleRequestTag) *CreateRoleRequest
	GetTag() []*CreateRoleRequestTag
}

type CreateRoleRequest struct {
	// Specifies whether the RAM role is allowed to log on to the console. Valid values:
	//
	// - true: allowed
	//
	// - false: not allowed
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	AllowConsoleLogin *bool `json:"AllowConsoleLogin,omitempty" xml:"AllowConsoleLogin,omitempty"`
	// The trust policy. Specifies one or more principals that are allowed to assume the RAM role. The principals can be Alibaba Cloud accounts, Alibaba Cloud services, or identity providers.
	//
	// >Resource Access Management (RAM) users cannot assume RAM roles whose trusted entity is an Alibaba Cloud service.
	//
	// example:
	//
	// {"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"RAM":"acs:ram::123456789012****:root"}}],"Version":"1"}
	AssumeRolePolicyDocument *string `json:"AssumeRolePolicyDocument,omitempty" xml:"AssumeRolePolicyDocument,omitempty"`
	// The description of the RAM role.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// ECS management role
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session duration of the RAM role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If this parameter is left empty, the default value is used.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The tags.
	Tag []*CreateRoleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetAllowConsoleLogin() *bool {
	return s.AllowConsoleLogin
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

func (s *CreateRoleRequest) GetTag() []*CreateRoleRequestTag {
	return s.Tag
}

func (s *CreateRoleRequest) SetAllowConsoleLogin(v bool) *CreateRoleRequest {
	s.AllowConsoleLogin = &v
	return s
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

func (s *CreateRoleRequest) SetTag(v []*CreateRoleRequestTag) *CreateRoleRequest {
	s.Tag = v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRoleRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRoleRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRoleRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRoleRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRoleRequestTag) SetKey(v string) *CreateRoleRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRoleRequestTag) SetValue(v string) *CreateRoleRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRoleRequestTag) Validate() error {
	return dara.Validate(s)
}
