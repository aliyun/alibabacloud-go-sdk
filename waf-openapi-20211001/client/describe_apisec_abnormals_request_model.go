// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecAbnormalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalId(v string) *DescribeApisecAbnormalsRequest
	GetAbnormalId() *string
	SetAbnormalLevel(v string) *DescribeApisecAbnormalsRequest
	GetAbnormalLevel() *string
	SetAbnormalTag(v string) *DescribeApisecAbnormalsRequest
	GetAbnormalTag() *string
	SetApiFormat(v string) *DescribeApisecAbnormalsRequest
	GetApiFormat() *string
	SetApiId(v string) *DescribeApisecAbnormalsRequest
	GetApiId() *string
	SetApiTag(v string) *DescribeApisecAbnormalsRequest
	GetApiTag() *string
	SetClusterId(v string) *DescribeApisecAbnormalsRequest
	GetClusterId() *string
	SetEndTime(v string) *DescribeApisecAbnormalsRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeApisecAbnormalsRequest
	GetInstanceId() *string
	SetMatchedHost(v string) *DescribeApisecAbnormalsRequest
	GetMatchedHost() *string
	SetOrderKey(v string) *DescribeApisecAbnormalsRequest
	GetOrderKey() *string
	SetOrderWay(v string) *DescribeApisecAbnormalsRequest
	GetOrderWay() *string
	SetOrigin(v string) *DescribeApisecAbnormalsRequest
	GetOrigin() *string
	SetPageNumber(v int64) *DescribeApisecAbnormalsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecAbnormalsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecAbnormalsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecAbnormalsRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v string) *DescribeApisecAbnormalsRequest
	GetStartTime() *string
	SetUserStatus(v string) *DescribeApisecAbnormalsRequest
	GetUserStatus() *string
}

type DescribeApisecAbnormalsRequest struct {
	// The ID of the security risk.
	//
	// example:
	//
	// 29c6401****99a2bad3943e26d8
	AbnormalId *string `json:"AbnormalId,omitempty" xml:"AbnormalId,omitempty"`
	// The severity level of the security risk. Valid values:
	//
	// - **high**: High.
	//
	// - **medium**: Medium.
	//
	// - **low**: Low.
	//
	// example:
	//
	// medium
	AbnormalLevel *string `json:"AbnormalLevel,omitempty" xml:"AbnormalLevel,omitempty"`
	// The type of the security risk.
	//
	// > Call [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) to query the supported risk types.
	//
	// example:
	//
	// LackOfSpeedLimit
	AbnormalTag *string `json:"AbnormalTag,omitempty" xml:"AbnormalTag,omitempty"`
	// The path of the API that is associated with the security risk.
	//
	// example:
	//
	// /api/users/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API that is associated with the security risk.
	//
	// example:
	//
	// bd9efb8ad******d9ca6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The business purpose of the API.
	//
	// > Call [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) to query the supported business purposes.
	//
	// example:
	//
	// RegisterAPI
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The ID of the hybrid cloud WAF cluster.
	//
	// > This parameter is required only for hybrid cloud scenarios. Call [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) to query the IDs of hybrid cloud WAF clusters.
	//
	// example:
	//
	// 546
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1684382100
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-z***9g301
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address that the API resides on.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The field by which the query results are sorted. Valid values:
	//
	// - **discoverTime*	- (default): The time when the risk was first detected.
	//
	// - **abnormalLevel**: The risk level.
	//
	// - **latestDiscoverTime**: The most recent time when the risk was detected.
	//
	// example:
	//
	// firstTime
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The sorting order. Valid values:
	//
	// - **desc*	- (default): Descending order.
	//
	// - **asc**: Ascending order.
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The source of the risk detection rule. Valid values:
	//
	// - **custom**: Custom rule.
	//
	// - **default**: Built-in rule.
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: Chinese mainland.
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
	// The start of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1684252800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The handling status of the security risk. Valid values:
	//
	// - **toBeConfirmed**: To be confirmed.
	//
	// - **confirmed**: Confirmed.
	//
	// - **toBeFixed**: To be fixed.
	//
	// - **fixed**: Fixed (manually verified).
	//
	// - **ignored**: Ignored.
	//
	// - **toBeVerified**: To be verified by the system.
	//
	// - **notFixed**: Verification failed.
	//
	// - **systemFixed**: Fixed (verified by the system).
	//
	// example:
	//
	// Confirmed
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s DescribeApisecAbnormalsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecAbnormalsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecAbnormalsRequest) GetAbnormalId() *string {
	return s.AbnormalId
}

func (s *DescribeApisecAbnormalsRequest) GetAbnormalLevel() *string {
	return s.AbnormalLevel
}

func (s *DescribeApisecAbnormalsRequest) GetAbnormalTag() *string {
	return s.AbnormalTag
}

func (s *DescribeApisecAbnormalsRequest) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecAbnormalsRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecAbnormalsRequest) GetApiTag() *string {
	return s.ApiTag
}

func (s *DescribeApisecAbnormalsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecAbnormalsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApisecAbnormalsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecAbnormalsRequest) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecAbnormalsRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *DescribeApisecAbnormalsRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeApisecAbnormalsRequest) GetOrigin() *string {
	return s.Origin
}

func (s *DescribeApisecAbnormalsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecAbnormalsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecAbnormalsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecAbnormalsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecAbnormalsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApisecAbnormalsRequest) GetUserStatus() *string {
	return s.UserStatus
}

func (s *DescribeApisecAbnormalsRequest) SetAbnormalId(v string) *DescribeApisecAbnormalsRequest {
	s.AbnormalId = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetAbnormalLevel(v string) *DescribeApisecAbnormalsRequest {
	s.AbnormalLevel = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetAbnormalTag(v string) *DescribeApisecAbnormalsRequest {
	s.AbnormalTag = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetApiFormat(v string) *DescribeApisecAbnormalsRequest {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetApiId(v string) *DescribeApisecAbnormalsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetApiTag(v string) *DescribeApisecAbnormalsRequest {
	s.ApiTag = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetClusterId(v string) *DescribeApisecAbnormalsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetEndTime(v string) *DescribeApisecAbnormalsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetInstanceId(v string) *DescribeApisecAbnormalsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetMatchedHost(v string) *DescribeApisecAbnormalsRequest {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetOrderKey(v string) *DescribeApisecAbnormalsRequest {
	s.OrderKey = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetOrderWay(v string) *DescribeApisecAbnormalsRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetOrigin(v string) *DescribeApisecAbnormalsRequest {
	s.Origin = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetPageNumber(v int64) *DescribeApisecAbnormalsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetPageSize(v int64) *DescribeApisecAbnormalsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetRegionId(v string) *DescribeApisecAbnormalsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecAbnormalsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetStartTime(v string) *DescribeApisecAbnormalsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) SetUserStatus(v string) *DescribeApisecAbnormalsRequest {
	s.UserStatus = &v
	return s
}

func (s *DescribeApisecAbnormalsRequest) Validate() error {
	return dara.Validate(s)
}
