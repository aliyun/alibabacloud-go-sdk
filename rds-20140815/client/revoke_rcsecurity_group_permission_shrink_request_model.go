// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRCSecurityGroupPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *RevokeRCSecurityGroupPermissionShrinkRequest
	GetDirection() *string
	SetRegionId(v string) *RevokeRCSecurityGroupPermissionShrinkRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *RevokeRCSecurityGroupPermissionShrinkRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleIdListShrink(v string) *RevokeRCSecurityGroupPermissionShrinkRequest
	GetSecurityGroupRuleIdListShrink() *string
}

type RevokeRCSecurityGroupPermissionShrinkRequest struct {
	Direction                     *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId               *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupRuleIdListShrink *string `json:"SecurityGroupRuleIdList,omitempty" xml:"SecurityGroupRuleIdList,omitempty"`
}

func (s RevokeRCSecurityGroupPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeRCSecurityGroupPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) GetSecurityGroupRuleIdListShrink() *string {
	return s.SecurityGroupRuleIdListShrink
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) SetDirection(v string) *RevokeRCSecurityGroupPermissionShrinkRequest {
	s.Direction = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) SetRegionId(v string) *RevokeRCSecurityGroupPermissionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) SetSecurityGroupId(v string) *RevokeRCSecurityGroupPermissionShrinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) SetSecurityGroupRuleIdListShrink(v string) *RevokeRCSecurityGroupPermissionShrinkRequest {
	s.SecurityGroupRuleIdListShrink = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
