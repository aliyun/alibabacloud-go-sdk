// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v string) *DescribeApisecEventsRequest
	GetAccount() *string
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
	SetEventScope(v string) *DescribeApisecEventsRequest
	GetEventScope() *string
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
	// The account that you want to use to filter events.
	//
	// example:
	//
	// 1818743389962696
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The path of the API that is associated with the security event.
	//
	// example:
	//
	// /apisec/v1/***.php
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API.
	//
	// example:
	//
	// 820b860***6205da93b935b28
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The business purpose of the API.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported business purposes.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The IP address of the attacker that you want to use to filter events.
	//
	// example:
	//
	// 42.224.*.*
	AttackIp *string `json:"AttackIp,omitempty" xml:"AttackIp,omitempty"`
	// The ID of the hybrid cloud WAF cluster.
	//
	// > This parameter is required only in hybrid cloud scenarios. Call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the IDs of hybrid cloud WAF clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp. Unit: seconds.
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
	// - **high**: high severity.
	//
	// - **medium**: medium severity.
	//
	// - **low**: low severity.
	//
	// example:
	//
	// low
	EventLevel *string `json:"EventLevel,omitempty" xml:"EventLevel,omitempty"`
	// The dimension by which security events are categorized. Valid values:
	//
	// - **ip**: IP security event. This is the default value.
	//
	// - **account**: account security event.
	//
	// example:
	//
	// ip
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The event type.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported event types.
	//
	// example:
	//
	// ObtainSensitiveUnauthorized
	EventTag *string `json:"EventTag,omitempty" xml:"EventTag,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-5y***d31
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address that is protected by WAF.
	//
	// example:
	//
	// a.***.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The field that is used to sort the query results. Valid values:
	//
	// - **allCnt**: the number of attacks.
	//
	// - **startTs**: the start time of the event.
	//
	// - **endTs**: the end time of the event.
	//
	// example:
	//
	// startTs
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The order in which the query results are sorted. Valid values:
	//
	// - **desc**: descending order. This is the default value.
	//
	// - **asc**: ascending order.
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The source of the event type. Valid values:
	//
	// - **custom**: a user-defined event type.
	//
	// - **default**: a built-in event type.
	//
	// example:
	//
	// default
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683648000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// The handling status of the event. Valid values:
	//
	// - **toBeConfirmed**: pending confirmation.
	//
	// - **confirmed**: confirmed but not yet handled.
	//
	// - **actioned**: handled.
	//
	// - **ignored**: ignored.
	//
	// example:
	//
	// ignored
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeApisecEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecEventsRequest) GetAccount() *string {
	return s.Account
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

func (s *DescribeApisecEventsRequest) GetEventScope() *string {
	return s.EventScope
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

func (s *DescribeApisecEventsRequest) SetAccount(v string) *DescribeApisecEventsRequest {
	s.Account = &v
	return s
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

func (s *DescribeApisecEventsRequest) SetEventScope(v string) *DescribeApisecEventsRequest {
	s.EventScope = &v
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
