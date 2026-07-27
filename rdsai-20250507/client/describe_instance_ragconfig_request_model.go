// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRAGConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBranchName(v string) *DescribeInstanceRAGConfigRequest
	GetBranchName() *string
	SetInstanceName(v string) *DescribeInstanceRAGConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceRAGConfigRequest
	GetRegionId() *string
}

type DescribeInstanceRAGConfigRequest struct {
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

func (s DescribeInstanceRAGConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRAGConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRAGConfigRequest) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeInstanceRAGConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceRAGConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceRAGConfigRequest) SetBranchName(v string) *DescribeInstanceRAGConfigRequest {
	s.BranchName = &v
	return s
}

func (s *DescribeInstanceRAGConfigRequest) SetInstanceName(v string) *DescribeInstanceRAGConfigRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceRAGConfigRequest) SetRegionId(v string) *DescribeInstanceRAGConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRAGConfigRequest) Validate() error {
	return dara.Validate(s)
}
