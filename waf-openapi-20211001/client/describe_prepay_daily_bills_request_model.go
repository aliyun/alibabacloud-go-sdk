// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrepayDailyBillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribePrepayDailyBillsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribePrepayDailyBillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePrepayDailyBillsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePrepayDailyBillsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribePrepayDailyBillsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribePrepayDailyBillsRequest struct {
	// ID of the WAF instance.
	//
	// > To view your WAF instance ID, call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Page number of the returned list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribePrepayDailyBillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrepayDailyBillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrepayDailyBillsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePrepayDailyBillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePrepayDailyBillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePrepayDailyBillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePrepayDailyBillsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribePrepayDailyBillsRequest) SetInstanceId(v string) *DescribePrepayDailyBillsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePrepayDailyBillsRequest) SetPageNumber(v int32) *DescribePrepayDailyBillsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePrepayDailyBillsRequest) SetPageSize(v int32) *DescribePrepayDailyBillsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePrepayDailyBillsRequest) SetRegionId(v string) *DescribePrepayDailyBillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePrepayDailyBillsRequest) SetResourceManagerResourceGroupId(v string) *DescribePrepayDailyBillsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePrepayDailyBillsRequest) Validate() error {
	return dara.Validate(s)
}
