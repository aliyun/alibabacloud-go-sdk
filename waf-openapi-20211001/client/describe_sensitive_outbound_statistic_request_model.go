// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSensitiveOutboundStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSensitiveOutboundStatisticRequest
	GetClusterId() *string
	SetDetectionResult(v string) *DescribeSensitiveOutboundStatisticRequest
	GetDetectionResult() *string
	SetEndTime(v int64) *DescribeSensitiveOutboundStatisticRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeSensitiveOutboundStatisticRequest
	GetInstanceId() *string
	SetOrderKey(v string) *DescribeSensitiveOutboundStatisticRequest
	GetOrderKey() *string
	SetOrderWay(v string) *DescribeSensitiveOutboundStatisticRequest
	GetOrderWay() *string
	SetPageNumber(v int64) *DescribeSensitiveOutboundStatisticRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSensitiveOutboundStatisticRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSensitiveOutboundStatisticRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeSensitiveOutboundStatisticRequest
	GetResourceManagerResourceGroupId() *string
	SetSensitiveCode(v string) *DescribeSensitiveOutboundStatisticRequest
	GetSensitiveCode() *string
	SetSensitiveLevel(v string) *DescribeSensitiveOutboundStatisticRequest
	GetSensitiveLevel() *string
	SetSensitiveType(v string) *DescribeSensitiveOutboundStatisticRequest
	GetSensitiveType() *string
	SetStartTime(v int64) *DescribeSensitiveOutboundStatisticRequest
	GetStartTime() *int64
}

type DescribeSensitiveOutboundStatisticRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is only applicable to hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the information about hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The assessment result. Valid values:
	//
	// - **report**: Outbound data risks exist.
	//
	// - **none**: No outbound data risks exist.
	//
	// example:
	//
	// report
	DetectionResult *string `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty"`
	// The end time of the query. The value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// > Compliance audit currently supports querying data for the last 1 month, last 3 months, last 6 months, last 12 months, and from January 1 of the previous year to the present. Make sure that the specified time range is valid.
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
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the field used for sorting. Valid values:
	//
	// - **total_count**: the total number of data records (default)
	//
	// - **outbound_count**: the total number of outbound data records
	//
	// example:
	//
	// total_count
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The sorting method. Valid values:
	//
	// - **desc**: descending order (default)
	//
	// - **asc**: ascending order
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The number of the page to return in a paging query. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page in a paging query. Default value: **10**, which indicates 10 entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland
	//
	// - **ap-southeast-1**: outside the Chinese mainland
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
	// The sensitive data types. Separate multiple sensitive data types with commas (,).
	//
	// > You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported sensitive data types. This parameter only supports built-in sensitive data types.
	//
	// example:
	//
	// 1000,1001
	SensitiveCode *string `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	// The sensitivity level. Valid values:
	//
	// - **high**: high sensitivity
	//
	// - **medium**: medium sensitivity
	//
	// - **low**: low sensitivity
	//
	// example:
	//
	// high
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The type of information to query. Valid values:
	//
	// - **info**: all personal information (default)
	//
	// - **sensitive**: only sensitive personal information
	//
	// example:
	//
	// info
	SensitiveType *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// The start time of the query. The value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// > Compliance audit currently supports querying data for the last 1 month, last 3 months, last 6 months, last 12 months, and from January 1 of the previous year to the present. Make sure that the specified time range is valid.
	//
	// example:
	//
	// 1672502400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSensitiveOutboundStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSensitiveOutboundStatisticRequest) GoString() string {
	return s.String()
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetDetectionResult() *string {
	return s.DetectionResult
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetSensitiveCode() *string {
	return s.SensitiveCode
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetSensitiveType() *string {
	return s.SensitiveType
}

func (s *DescribeSensitiveOutboundStatisticRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetClusterId(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetDetectionResult(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.DetectionResult = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetEndTime(v int64) *DescribeSensitiveOutboundStatisticRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetInstanceId(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetOrderKey(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.OrderKey = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetOrderWay(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetPageNumber(v int64) *DescribeSensitiveOutboundStatisticRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetPageSize(v int64) *DescribeSensitiveOutboundStatisticRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetRegionId(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetResourceManagerResourceGroupId(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetSensitiveCode(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.SensitiveCode = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetSensitiveLevel(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetSensitiveType(v string) *DescribeSensitiveOutboundStatisticRequest {
	s.SensitiveType = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) SetStartTime(v int64) *DescribeSensitiveOutboundStatisticRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSensitiveOutboundStatisticRequest) Validate() error {
	return dara.Validate(s)
}
