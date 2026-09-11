// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEtlJobLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeEtlJobLogsRequest
	GetDtsJobId() *string
	SetRegionId(v string) *DescribeEtlJobLogsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeEtlJobLogsRequest
	GetResourceGroupId() *string
}

type DescribeEtlJobLogsRequest struct {
	// The ID of the ETL task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vd*********2ls
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The region ID of the DTS instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the list of Alibaba Cloud regions.
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
}

func (s DescribeEtlJobLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEtlJobLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEtlJobLogsRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeEtlJobLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEtlJobLogsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeEtlJobLogsRequest) SetDtsJobId(v string) *DescribeEtlJobLogsRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeEtlJobLogsRequest) SetRegionId(v string) *DescribeEtlJobLogsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEtlJobLogsRequest) SetResourceGroupId(v string) *DescribeEtlJobLogsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeEtlJobLogsRequest) Validate() error {
	return dara.Validate(s)
}
