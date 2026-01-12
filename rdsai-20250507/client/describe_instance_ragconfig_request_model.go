// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRAGConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DescribeInstanceRAGConfigRequest
	GetInstanceName() *string
	SetRegionId(v string) *DescribeInstanceRAGConfigRequest
	GetRegionId() *string
}

type DescribeInstanceRAGConfigRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The operation that you want to perform. Set the value to **DescribeInstanceRAGConfig**.
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

func (s *DescribeInstanceRAGConfigRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeInstanceRAGConfigRequest) GetRegionId() *string {
	return s.RegionId
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
