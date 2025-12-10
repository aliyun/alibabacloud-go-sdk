// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityGroupPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionsShrink(v string) *CreateSecurityGroupPermissionsShrinkRequest
	GetPermissionsShrink() *string
	SetSecurityGroupId(v string) *CreateSecurityGroupPermissionsShrinkRequest
	GetSecurityGroupId() *string
}

type CreateSecurityGroupPermissionsShrinkRequest struct {
	// The security group rules.
	//
	// This parameter is required.
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The IDs of the security groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s CreateSecurityGroupPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityGroupPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityGroupPermissionsShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *CreateSecurityGroupPermissionsShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateSecurityGroupPermissionsShrinkRequest) SetPermissionsShrink(v string) *CreateSecurityGroupPermissionsShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *CreateSecurityGroupPermissionsShrinkRequest) SetSecurityGroupId(v string) *CreateSecurityGroupPermissionsShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateSecurityGroupPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
