// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStorageConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceStorageConfigRequest
	GetBranchName() *string
	SetInstanceName(v string) *DescribeInstanceStorageConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceStorageConfigRequest
	GetRegionId() *string
}

type DescribeInstanceStorageConfigRequest struct {
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

func (s DescribeInstanceStorageConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStorageConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStorageConfigRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceStorageConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceStorageConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceStorageConfigRequest) SetBranchName(v string) *DescribeInstanceStorageConfigRequest {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceStorageConfigRequest) SetInstanceName(v string) *DescribeInstanceStorageConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceStorageConfigRequest) SetRegionId(v string) *DescribeInstanceStorageConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceStorageConfigRequest) Validate() error {
	return dara.Validate(s)
}
