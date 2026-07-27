// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceSSLRequest
	GetBranchName() *string
	SetInstanceName(v string) *DescribeInstanceSSLRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceSSLRequest
	GetRegionId() *string
}

type DescribeInstanceSSLRequest struct {
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The instance ID of the AI application.
	//
	// This parameter is required.
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

func (s DescribeInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceSSLRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceSSLRequest) SetBranchName(v string) *DescribeInstanceSSLRequest {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetInstanceName(v string) *DescribeInstanceSSLRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceSSLRequest) SetRegionId(v string) *DescribeInstanceSSLRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
