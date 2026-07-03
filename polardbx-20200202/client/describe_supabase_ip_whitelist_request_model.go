// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSupabaseIpWhitelistRequest
	GetDBInstanceName() *string
	SetGroupName(v string) *DescribeSupabaseIpWhitelistRequest
	GetGroupName() *string
	SetRegionId(v string) *DescribeSupabaseIpWhitelistRequest
	GetRegionId() *string
}

type DescribeSupabaseIpWhitelistRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcsp-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The whitelist group name.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSupabaseIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseIpWhitelistRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseIpWhitelistRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSupabaseIpWhitelistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseIpWhitelistRequest) SetDBInstanceName(v string) *DescribeSupabaseIpWhitelistRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistRequest) SetGroupName(v string) *DescribeSupabaseIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistRequest) SetRegionId(v string) *DescribeSupabaseIpWhitelistRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
