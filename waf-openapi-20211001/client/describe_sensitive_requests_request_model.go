// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveRequestsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveRequestsRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveRequestsRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveRequestsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeSensitiveRequestsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSensitiveRequestsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSensitiveRequestsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveRequestsRequest
	GetResourceManagerResourceGroupId() *string
	SetSensitiveCode(v int64) *DescribeSensitiveRequestsRequest
	GetSensitiveCode() *int64
	SetSensitiveData(v string) *DescribeSensitiveRequestsRequest
	GetSensitiveData() *string
	SetStartTime(v int64) *DescribeSensitiveRequestsRequest
	GetStartTime() *int64
}

type DescribeSensitiveRequestsRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is available only in hybrid cloud scenarios. Call [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) to obtain information about hybrid cloud clusters.
	//
	// example:
	//
	// 269
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp in seconds. The time is in UTC.
	//
	// example:
	//
	// 1725966000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return for a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
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
	// rg-aekzwwkpn****5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of sensitive data.
	//
	// > Call [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) to obtain the supported sensitive data types.
	//
	// example:
	//
	// 1001
	SensitiveCode *int64 `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	// The sensitive data.
	//
	// example:
	//
	// card
	SensitiveData *string `json:"SensitiveData,omitempty" xml:"SensitiveData,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp in seconds. The time is in UTC.
	//
	// example:
	//
	// 1723392000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSensitiveRequestsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveRequestsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveRequestsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveRequestsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveRequestsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveRequestsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSensitiveRequestsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSensitiveRequestsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveRequestsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveRequestsRequest) GetSensitiveCode() *int64 {
	return s.SensitiveCode
}

func (s *DescribeSensitiveRequestsRequest) GetSensitiveData() *string {
	return s.SensitiveData
}

func (s *DescribeSensitiveRequestsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveRequestsRequest) SetClusterId(v string) *DescribeSensitiveRequestsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetEndTime(v int64) *DescribeSensitiveRequestsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetInstanceId(v string) *DescribeSensitiveRequestsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetPageNumber(v int64) *DescribeSensitiveRequestsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetPageSize(v int64) *DescribeSensitiveRequestsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetRegionId(v string) *DescribeSensitiveRequestsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveRequestsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetSensitiveCode(v int64) *DescribeSensitiveRequestsRequest {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetSensitiveData(v string) *DescribeSensitiveRequestsRequest {
	s.SensitiveData = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) SetStartTime(v int64) *DescribeSensitiveRequestsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveRequestsRequest) Validate() error {
	return dara.Validate(s)
}
