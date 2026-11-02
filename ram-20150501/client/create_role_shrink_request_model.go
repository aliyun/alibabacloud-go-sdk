// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowConsoleLogin(v bool) *CreateRoleShrinkRequest
	GetAllowConsoleLogin() *bool
	SetAssumeRolePolicyDocument(v string) *CreateRoleShrinkRequest
	GetAssumeRolePolicyDocument() *string
	SetDescription(v string) *CreateRoleShrinkRequest
	GetDescription() *string
	SetMaxSessionDuration(v int64) *CreateRoleShrinkRequest
	GetMaxSessionDuration() *int64
	SetRoleName(v string) *CreateRoleShrinkRequest
	GetRoleName() *string
	SetTagShrink(v string) *CreateRoleShrinkRequest
	GetTagShrink() *string
}

type CreateRoleShrinkRequest struct {
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
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRoleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleShrinkRequest) GetAllowConsoleLogin() *bool {
	return s.AllowConsoleLogin
}

func (s *CreateRoleShrinkRequest) GetAssumeRolePolicyDocument() *string {
	return s.AssumeRolePolicyDocument
}

func (s *CreateRoleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoleShrinkRequest) GetMaxSessionDuration() *int64 {
	return s.MaxSessionDuration
}

func (s *CreateRoleShrinkRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateRoleShrinkRequest) SetAllowConsoleLogin(v bool) *CreateRoleShrinkRequest {
	s.AllowConsoleLogin = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetAssumeRolePolicyDocument(v string) *CreateRoleShrinkRequest {
	s.AssumeRolePolicyDocument = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetDescription(v string) *CreateRoleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetMaxSessionDuration(v int64) *CreateRoleShrinkRequest {
	s.MaxSessionDuration = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetRoleName(v string) *CreateRoleShrinkRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleShrinkRequest) SetTagShrink(v string) *CreateRoleShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateRoleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
