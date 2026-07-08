// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayBillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribePostpayBillsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribePostpayBillsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribePostpayBillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePostpayBillsRequest
	GetNextToken() *string
	SetPeriodType(v string) *DescribePostpayBillsRequest
	GetPeriodType() *string
	SetRegionId(v string) *DescribePostpayBillsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribePostpayBillsRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribePostpayBillsRequest
	GetStartTime() *int64
}

type DescribePostpayBillsRequest struct {
	// The end time. The value is a UNIX timestamp (UTC). Unit: seconds.
	//
	// example:
	//
	// 1779195599
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-zz11sr5****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 24
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. If a next page exists, this field has a return value.
	//
	// > If this parameter has a return value, a next page exists. You can use the returned **NextToken*	- as a request parameter to obtain the data on the next page. Repeat this process until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The epoch type for billing information. Valid values:
	//
	// - **hour**: Billing information is calculated on an hourly basis.
	//
	// example:
	//
	// hour
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The region where the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The start time. The value is a UNIX timestamp (UTC). Unit: seconds.
	//
	// example:
	//
	// 1779120000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePostpayBillsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayBillsRequest) GoString() string {
	return s.String()
}

func (s *DescribePostpayBillsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribePostpayBillsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePostpayBillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePostpayBillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePostpayBillsRequest) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribePostpayBillsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePostpayBillsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribePostpayBillsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribePostpayBillsRequest) SetEndTime(v int64) *DescribePostpayBillsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetInstanceId(v string) *DescribePostpayBillsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetMaxResults(v int32) *DescribePostpayBillsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetNextToken(v string) *DescribePostpayBillsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetPeriodType(v string) *DescribePostpayBillsRequest {
	s.PeriodType = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetRegionId(v string) *DescribePostpayBillsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetResourceManagerResourceGroupId(v string) *DescribePostpayBillsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePostpayBillsRequest) SetStartTime(v int64) *DescribePostpayBillsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePostpayBillsRequest) Validate() error {
	return dara.Validate(s)
}
