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
	// The direction of the security group rules that you want to delete. Valid values:
	//
	// 	- **ingress**: inbound security group rules.
	//
	// 	- **egress**: outbound security group rules.
	//
	// >  You can specify security group rules only in the same direction in a request.
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
	// The IDs of the security group rules that you want to delete.
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
