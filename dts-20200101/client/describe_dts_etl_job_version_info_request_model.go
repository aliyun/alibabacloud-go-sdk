// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsEtlJobVersionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsInstanceId(v string) *DescribeDtsEtlJobVersionInfoRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *DescribeDtsEtlJobVersionInfoRequest
	GetDtsJobId() *string
	SetPageNumber(v int32) *DescribeDtsEtlJobVersionInfoRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDtsEtlJobVersionInfoRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDtsEtlJobVersionInfoRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDtsEtlJobVersionInfoRequest
	GetResourceGroupId() *string
}

type DescribeDtsEtlJobVersionInfoRequest struct {
	// The ID of the Data Transmission Service (DTS) instance. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsggk12iwya1a****
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the ETL task. You can call the [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) operation to query the task ID.
	//
	// example:
	//
	// l5512es7w15****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the DTS instance resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the available Alibaba Cloud regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDtsEtlJobVersionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsEtlJobVersionInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsEtlJobVersionInfoRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeDtsEtlJobVersionInfoRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsEtlJobVersionInfoRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDtsEtlJobVersionInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDtsEtlJobVersionInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsEtlJobVersionInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsEtlJobVersionInfoRequest) SetDtsInstanceId(v string) *DescribeDtsEtlJobVersionInfoRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoRequest) SetDtsJobId(v string) *DescribeDtsEtlJobVersionInfoRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoRequest) SetPageNumber(v int32) *DescribeDtsEtlJobVersionInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoRequest) SetPageSize(v int32) *DescribeDtsEtlJobVersionInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoRequest) SetRegionId(v string) *DescribeDtsEtlJobVersionInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoRequest) SetResourceGroupId(v string) *DescribeDtsEtlJobVersionInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsEtlJobVersionInfoRequest) Validate() error {
	return dara.Validate(s)
}
