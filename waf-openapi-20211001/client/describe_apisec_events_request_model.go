// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiFormat(v string) *DescribeApisecEventsRequest
	GetApiFormat() *string
	SetApiId(v string) *DescribeApisecEventsRequest
	GetApiId() *string
	SetApiTag(v string) *DescribeApisecEventsRequest
	GetApiTag() *string
	SetAttackIp(v string) *DescribeApisecEventsRequest
	GetAttackIp() *string
	SetClusterId(v string) *DescribeApisecEventsRequest
	GetClusterId() *string
	SetEndTs(v int64) *DescribeApisecEventsRequest
	GetEndTs() *int64
	SetEventId(v string) *DescribeApisecEventsRequest
	GetEventId() *string
	SetEventLevel(v string) *DescribeApisecEventsRequest
	GetEventLevel() *string
	SetEventTag(v string) *DescribeApisecEventsRequest
	GetEventTag() *string
	SetInstanceId(v string) *DescribeApisecEventsRequest
	GetInstanceId() *string
	SetMatchedHost(v string) *DescribeApisecEventsRequest
	GetMatchedHost() *string
	SetOrderKey(v string) *DescribeApisecEventsRequest
	GetOrderKey() *string
	SetOrderWay(v string) *DescribeApisecEventsRequest
	GetOrderWay() *string
	SetOrigin(v string) *DescribeApisecEventsRequest
	GetOrigin() *string
	SetPageNumber(v int64) *DescribeApisecEventsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecEventsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecEventsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecEventsRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTs(v int64) *DescribeApisecEventsRequest
	GetStartTs() *int64
	SetUserStatus(v string) *DescribeApisecEventsRequest
	GetUserStatus() *string
}

type DescribeApisecEventsRequest struct {
	// The API.
	//
	// example:
	//
	// /apisec/v1/register.php
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the event-related API.
	//
	// example:
	//
	// 820b860***6205da93b935b28
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The business purpose of the API.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the business purposes of APIs.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The Attack source IP.
	//
	// example:
	//
	// 42.224.*.*
	AttackIp *string `json:"AttackIp,omitempty" xml:"AttackIp,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// >  This parameter is available only in hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683703260
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// The ID of the API security event.
	//
	// example:
	//
	// 18ba94fea9***e66ba0557b7b91
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The severity level of the event. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// low
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The type of the event.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported event types.
	//
	// example:
	//
	// ObtainSensitiveUnauthorized
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-5y***d31
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address of the API.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The name of the sorting field. Valid values:
	//
	// 	- **allCnt**: the number of attacks
	//
	// 	- **startTs**: the start time of the event
	//
	// 	- **endTs**: the end time of the event
	//
	// example:
	//
	// startTs
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The sorting method. Valid values:
	//
	// 	- **desc*	- (default): descending order
	//
	// 	- **asc**: ascending order
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The source of the event type. Valid values:
	//
	// 	- **custom**
	//
	// 	- **default**
	//
	// example:
	//
	// default
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
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
	// The region ID of the WAF instance. Value:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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
	// The beginning of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683648000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The event status. Valid values:
	//
	// 	- **toBeConfirmed**
	//
	// 	- **confirmed**
	//
	// 	- **ignored**
	//
	// example:
	//
	// Ignore
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeApisecEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventsRequest) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecEventsRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecEventsRequest) GetApiTag() *string {
	return s.ApiTag
}

func (s *DescribeApisecEventsRequest) GetAttackIp() *string {
	return s.AttackIp
}

func (s *DescribeApisecEventsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecEventsRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *DescribeApisecEventsRequest) GetEventId() *string {
	return s.EventId
}

func (s *DescribeApisecEventsRequest) GetEventLevel() *string {
	return s.EventLevel
}

func (s *DescribeApisecEventsRequest) GetEventTag() *string {
	return s.EventTag
}

func (s *DescribeApisecEventsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecEventsRequest) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecEventsRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *DescribeApisecEventsRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeApisecEventsRequest) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeApisecEventsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecEventsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecEventsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecEventsRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *DescribeApisecEventsRequest) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeApisecEventsRequest) SetApiFormat(v string) *DescribeApisecEventsRequest {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetApiId(v string) *DescribeApisecEventsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetApiTag(v string) *DescribeApisecEventsRequest {
	s.ApiTag = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetAttackIp(v string) *DescribeApisecEventsRequest {
	s.AttackIp = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetClusterId(v string) *DescribeApisecEventsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetEndTs(v int64) *DescribeApisecEventsRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetEventId(v string) *DescribeApisecEventsRequest {
	s.EventId = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetEventLevel(v string) *DescribeApisecEventsRequest {
	s.EventLevel = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetEventTag(v string) *DescribeApisecEventsRequest {
	s.EventTag = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetInstanceId(v string) *DescribeApisecEventsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetMatchedHost(v string) *DescribeApisecEventsRequest {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetOrderKey(v string) *DescribeApisecEventsRequest {
	s.OrderKey = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetOrderWay(v string) *DescribeApisecEventsRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetOrigin(v string) *DescribeApisecEventsRequest {
	s.Origin = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetPageNumber(v int64) *DescribeApisecEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetPageSize(v int64) *DescribeApisecEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetRegionId(v string) *DescribeApisecEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecEventsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetStartTs(v int64) *DescribeApisecEventsRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeApisecEventsRequest) SetUserStatus(v string) *DescribeApisecEventsRequest {
	s.UserStatus = &v
	return s
}

func (s *DescribeApisecEventsRequest) Validate() error {
	return dara.Validate(s)
}
