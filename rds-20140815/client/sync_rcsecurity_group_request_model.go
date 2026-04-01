// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRCSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SyncRCSecurityGroupRequest
	GetInstanceId() *string
	SetRegionId(v string) *SyncRCSecurityGroupRequest
	GetRegionId() *string
	SetSecurityGroupId(v string) *SyncRCSecurityGroupRequest
	GetSecurityGroupId() *string
}

type SyncRCSecurityGroupRequest struct {
	// The instance ID.
	//
	// example:
	//
	// rc-i322y2t562oh7o******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp196e4d2ndjgy******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s SyncRCSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncRCSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *SyncRCSecurityGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SyncRCSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SyncRCSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SyncRCSecurityGroupRequest) SetInstanceId(v string) *SyncRCSecurityGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SyncRCSecurityGroupRequest) SetRegionId(v string) *SyncRCSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *SyncRCSecurityGroupRequest) SetSecurityGroupId(v string) *SyncRCSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *SyncRCSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
