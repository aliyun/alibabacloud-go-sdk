// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSecurityGroupShrinkRequest
	GetDescription() *string
	SetPermissionsShrink(v string) *CreateSecurityGroupShrinkRequest
	GetPermissionsShrink() *string
	SetSecurityGroupName(v string) *CreateSecurityGroupShrinkRequest
	GetSecurityGroupName() *string
}

type CreateSecurityGroupShrinkRequest struct {
	// The description of the security group. The description must be 2 to 256 characters in length. It must start with a letter and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rules of the security group. Valid values of N: 0 to 200.
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The name of the security group. The name must be 2 to 128 characters in length, and can contain digits, colons (:), underscores (_), and hyphens (-). It must start with a letter and cannot start with `http://` or `https://`. By default, this parameter is empty.
	//
	// example:
	//
	// Dcdn1:2_3-4
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s CreateSecurityGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSecurityGroupShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *CreateSecurityGroupShrinkRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *CreateSecurityGroupShrinkRequest) SetDescription(v string) *CreateSecurityGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSecurityGroupShrinkRequest) SetPermissionsShrink(v string) *CreateSecurityGroupShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *CreateSecurityGroupShrinkRequest) SetSecurityGroupName(v string) *CreateSecurityGroupShrinkRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateSecurityGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
