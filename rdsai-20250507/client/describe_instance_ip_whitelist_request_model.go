// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceIpWhitelistRequest
	GetBranchName() *string
	SetGroupName(v string) *DescribeInstanceIpWhitelistRequest
	GetGroupName() *string
	SetInstanceName(v string) *DescribeInstanceIpWhitelistRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceIpWhitelistRequest
	GetRegionId() *string
}

type DescribeInstanceIpWhitelistRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The group name.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceIpWhitelistRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceIpWhitelistRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeInstanceIpWhitelistRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceIpWhitelistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceIpWhitelistRequest) SetBranchName(v string) *DescribeInstanceIpWhitelistRequest {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceIpWhitelistRequest) SetGroupName(v string) *DescribeInstanceIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeInstanceIpWhitelistRequest) SetInstanceName(v string) *DescribeInstanceIpWhitelistRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceIpWhitelistRequest) SetRegionId(v string) *DescribeInstanceIpWhitelistRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
