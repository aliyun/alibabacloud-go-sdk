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
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The evaluation result. Valid values:
	//
	// 	- **report**: Risks exist in cross-border data transfer.
	//
	// 	- **none**: No risks exist in cross-border data transfer.
	//
	// example:
	//
	// report
	DetectionResult *string `json:"DetectionResult,omitempty" xml:"DetectionResult,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// >  You can query only data of the previous month, previous 3 months, previous 6 months, previous 12 months, and data generated since January 1 of last year for compliance check. You must specify a valid time range.
	//
	// example:
	//
	// 1725966000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the sorting field. Valid values:
	//
	// 	- **total_count*	- (default): total number of data entries
	//
	// 	- **outbound_count**: total number of data entries that are transferred across borders
	//
	// example:
	//
	// total_count
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- **desc*	- (default): in descending order
	//
	// 	- **asc**: in ascending order
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The page number. Default value: **1**.
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
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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
	// The type of the sensitive data. Separate multiple types with commas (,).
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported types of sensitive data. Only built-in types of sensitive data are supported for this operation.
	//
	// example:
	//
	// 1000,1001
	SensitiveCode *string `json:"SensitiveCode,omitempty" xml:"SensitiveCode,omitempty"`
	// The sensitivity level. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The type of the information. Valid values:
	//
	// 	- **info*	- (default): full personal information
	//
	// 	- **sensitive**: sensitive personal information
	//
	// example:
	//
	// info
	SensitiveType *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// >  You can query only data of the previous month, previous 3 months, previous 6 months, previous 12 months, and data generated since January 1 of last year for compliance check. You must specify a valid time range.
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
