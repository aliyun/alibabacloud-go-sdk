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
	SetTag(v []*CreateRoleRequestTag) *CreateRoleRequest
	GetTag() []*CreateRoleRequestTag
}

type CreateRoleRequest struct {
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
	Tag []*CreateRoleRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateRoleRequest) GetTag() []*CreateRoleRequestTag {
	return s.Tag
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
	// The key of the tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
