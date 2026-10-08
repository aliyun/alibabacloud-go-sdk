// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRCSecurityGroupPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *RevokeRCSecurityGroupPermissionRequest
	GetDirection() *string
	SetRegionId(v string) *RevokeRCSecurityGroupPermissionRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *RevokeRCSecurityGroupPermissionRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleIdList(v []*string) *RevokeRCSecurityGroupPermissionRequest
	GetSecurityGroupRuleIdList() []*string
}

type RevokeRCSecurityGroupPermissionRequest struct {
	Direction               *string   `json:"Direction,omitempty" xml:"Direction,omitempty"`
	RegionId                *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroupId         *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupRuleIdList []*string `json:"SecurityGroupRuleIdList,omitempty" xml:"SecurityGroupRuleIdList,omitempty" type:"Repeated"`
}

func (s RevokeRCSecurityGroupPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeRCSecurityGroupPermissionRequest) GoString() string {
	return s.String()
}

func (s *RevokeRCSecurityGroupPermissionRequest) GetDirection() *string {
	return s.Direction
}

func (s *RevokeRCSecurityGroupPermissionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeRCSecurityGroupPermissionRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RevokeRCSecurityGroupPermissionRequest) GetSecurityGroupRuleIdList() []*string {
	return s.SecurityGroupRuleIdList
}

func (s *RevokeRCSecurityGroupPermissionRequest) SetDirection(v string) *RevokeRCSecurityGroupPermissionRequest {
	s.Direction = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionRequest) SetRegionId(v string) *RevokeRCSecurityGroupPermissionRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionRequest) SetSecurityGroupId(v string) *RevokeRCSecurityGroupPermissionRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionRequest) SetSecurityGroupRuleIdList(v []*string) *RevokeRCSecurityGroupPermissionRequest {
	s.SecurityGroupRuleIdList = v
	return s
}

func (s *RevokeRCSecurityGroupPermissionRequest) Validate() error {
	return dara.Validate(s)
}
