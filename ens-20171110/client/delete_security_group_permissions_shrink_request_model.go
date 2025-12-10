// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityGroupPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPermissionsShrink(v string) *DeleteSecurityGroupPermissionsShrinkRequest
	GetPermissionsShrink() *string
	SetSecurityGroupId(v string) *DeleteSecurityGroupPermissionsShrinkRequest
	GetSecurityGroupId() *string
}

type DeleteSecurityGroupPermissionsShrinkRequest struct {
	// The security group rules.
	//
	// This parameter is required.
	PermissionsShrink *string `json:"Permissions,omitempty" xml:"Permissions,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityGroupPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupPermissionsShrinkRequest) GetPermissionsShrink() *string {
	return s.PermissionsShrink
}

func (s *DeleteSecurityGroupPermissionsShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeleteSecurityGroupPermissionsShrinkRequest) SetPermissionsShrink(v string) *DeleteSecurityGroupPermissionsShrinkRequest {
	s.PermissionsShrink = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsShrinkRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupPermissionsShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DeleteSecurityGroupPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
