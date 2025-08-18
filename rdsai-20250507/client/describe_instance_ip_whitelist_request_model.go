// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeInstanceIpWhitelistRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceIpWhitelistRequest
	GetRegionId() *string
}

type DescribeInstanceIpWhitelistRequest struct {
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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

func (s *DescribeInstanceIpWhitelistRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceIpWhitelistRequest) GetRegionId() *string {
	return s.RegionId
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
