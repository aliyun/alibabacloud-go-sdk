// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFullProcessListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeFullProcessListRequest
	GetDtsJobId() *string
	SetRegionId(v string) *DescribeFullProcessListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeFullProcessListRequest
	GetResourceGroupId() *string
	SetZeroEtlJob(v bool) *DescribeFullProcessListRequest
	GetZeroEtlJob() *bool
}

type DescribeFullProcessListRequest struct {
	// The ID of the migration, synchronization, or change tracking task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i03e3zty16i****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region ID. Specify this parameter to indicate the region where the instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether the node is a zero-ETL seamless integration node. Valid values:
	//
	// - **true**: The node is a zero-ETL seamless integration node.
	//
	// - **false**: The node is not a zero-ETL seamless integration node.
	//
	// example:
	//
	// true
	ZeroEtlJob *bool `json:"ZeroEtlJob,omitempty" xml:"ZeroEtlJob,omitempty"`
}

func (s DescribeFullProcessListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFullProcessListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFullProcessListRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeFullProcessListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFullProcessListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeFullProcessListRequest) GetZeroEtlJob() *bool {
	return s.ZeroEtlJob
}

func (s *DescribeFullProcessListRequest) SetDtsJobId(v string) *DescribeFullProcessListRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeFullProcessListRequest) SetRegionId(v string) *DescribeFullProcessListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFullProcessListRequest) SetResourceGroupId(v string) *DescribeFullProcessListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeFullProcessListRequest) SetZeroEtlJob(v bool) *DescribeFullProcessListRequest {
	s.ZeroEtlJob = &v
	return s
}

func (s *DescribeFullProcessListRequest) Validate() error {
	return dara.Validate(s)
}
