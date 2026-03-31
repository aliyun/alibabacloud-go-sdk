// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The trust policy that specifies one or more trusted entities to assume the RAM role. The trusted entities can be Alibaba Cloud accounts, Alibaba Cloud services, or identity providers (IdPs).
	//
	// >  RAM users cannot assume the RAM roles of trusted Alibaba Cloud services.
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
	// ECS administrator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum session time of the RAM role.
	//
	// Valid values: 3600 to 43200. Unit: seconds. Default value: 3600.
	//
	// If you do not specify this parameter, the default value is used.
	//
	// example:
	//
	// 3600
	MaxSessionDuration *int64 `json:"MaxSessionDuration,omitempty" xml:"MaxSessionDuration,omitempty"`
	// The name of the RAM role.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, periods (.), and hyphens (-).
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
