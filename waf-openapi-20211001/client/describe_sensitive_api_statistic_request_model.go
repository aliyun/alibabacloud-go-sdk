// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveApiStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveApiStatisticRequest
	GetClusterId() *string
	SetEndTime(v int64) *DescribeSensitiveApiStatisticRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveApiStatisticRequest
	GetInstanceId() *string
	SetMatchedHost(v string) *DescribeSensitiveApiStatisticRequest
	GetMatchedHost() *string
	SetPageNumber(v int64) *DescribeSensitiveApiStatisticRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSensitiveApiStatisticRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSensitiveApiStatisticRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveApiStatisticRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeSensitiveApiStatisticRequest
	GetStartTime() *int64
	SetType(v string) *DescribeSensitiveApiStatisticRequest
	GetType() *string
}

type DescribeSensitiveApiStatisticRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is available only in hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query information about hybrid cloud clusters.
	//
	// example:
	//
	// 269
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a Unix timestamp. Unit: seconds.
	//
	// > Compliance assessment supports queries for data from the past 1, 3, 6, and 12 months, or from January 1 of the previous year to the current time. Make sure that the specified time range is valid.
	//
	// example:
	//
	// 1725966000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address of the API.
	//
	// example:
	//
	// a.***.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland
	//
	// - **ap-southeast-1**: outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a Unix timestamp. Unit: seconds.
	//
	// > Compliance assessment supports queries for data from the past 1, 3, 6, and 12 months, or from January 1 of the previous year to the current time. Make sure that the specified time range is valid.
	//
	// example:
	//
	// 1672502400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The dimension for statistics. Valid values:
	//
	// - **matchedHost*	- (default): statistics by domain name
	//
	// - **apiFormat**: statistics by API
	//
	// 	Notice:
	//
	// To collect statistics by API, you must specify the **MatchedHost*	- parameter.
	//
	// example:
	//
	// matchedHost
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeSensitiveApiStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveApiStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveApiStatisticRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveApiStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveApiStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveApiStatisticRequest) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeSensitiveApiStatisticRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSensitiveApiStatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSensitiveApiStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveApiStatisticRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveApiStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveApiStatisticRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSensitiveApiStatisticRequest) SetClusterId(v string) *DescribeSensitiveApiStatisticRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetEndTime(v int64) *DescribeSensitiveApiStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetInstanceId(v string) *DescribeSensitiveApiStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetMatchedHost(v string) *DescribeSensitiveApiStatisticRequest {
	s.MatchedHost = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetPageNumber(v int64) *DescribeSensitiveApiStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetPageSize(v int64) *DescribeSensitiveApiStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetRegionId(v string) *DescribeSensitiveApiStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveApiStatisticRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetStartTime(v int64) *DescribeSensitiveApiStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) SetType(v string) *DescribeSensitiveApiStatisticRequest {
	s.Type = &v
	return s
}

func (s *DescribeSensitiveApiStatisticRequest) Validate() error {
	return dara.Validate(s)
}
