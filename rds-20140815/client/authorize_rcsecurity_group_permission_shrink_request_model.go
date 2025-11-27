// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRCSecurityGroupPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest
	GetDirection() *string
	SetRegionId(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityGroupPermissionsShrink(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest
	GetSecurityGroupPermissionsShrink() *string
}

type AuthorizeRCSecurityGroupPermissionShrinkRequest struct {
	// The direction of the rule. Valid values:
	//
	// 	- **ingress**: the inbound security group rule.
	//
	// 	- **egress**: the outbound security group rule.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-2ze27hs990o2hn9****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The information about the security group.
	SecurityGroupPermissionsShrink *string `json:"SecurityGroupPermissions,omitempty" xml:"SecurityGroupPermissions,omitempty"`
}

func (s AuthorizeRCSecurityGroupPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRCSecurityGroupPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) GetSecurityGroupPermissionsShrink() *string {
	return s.SecurityGroupPermissionsShrink
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) SetDirection(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest {
	s.Direction = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) SetRegionId(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) SetSecurityGroupId(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) SetSecurityGroupPermissionsShrink(v string) *AuthorizeRCSecurityGroupPermissionShrinkRequest {
	s.SecurityGroupPermissionsShrink = &v
	return s
}

func (s *AuthorizeRCSecurityGroupPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
