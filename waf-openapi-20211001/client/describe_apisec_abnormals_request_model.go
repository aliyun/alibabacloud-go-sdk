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
	// The risk ID.
	//
	// example:
	//
	// 29c6401****99a2bad3943e26d8
	AbnormalId *string `json:"AbnormalId,omitempty" xml:"AbnormalId,omitempty"`
	// The risk level. Valid values:
	//
	// - **high**: high risk.
	//
	// - **medium**: medium risk.
	//
	// - **low**: low risk.
	//
	// example:
	//
	// medium
	AbnormalLevel *string `json:"AbnormalLevel,omitempty" xml:"AbnormalLevel,omitempty"`
	// The risk type.
	//
	// > You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported risk types.
	//
	// example:
	//
	// LackOfSpeedLimit
	AbnormalTag *string `json:"AbnormalTag,omitempty" xml:"AbnormalTag,omitempty"`
	// The API operation associated with the risk.
	//
	// example:
	//
	// /api/users/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API associated with the risk.
	//
	// example:
	//
	// bd9efb8ad******d9ca6
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The business purpose of the API.
	//
	// > You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported business purposes.
	//
	// example:
	//
	// RegisterAPI
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter applies only to hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query hybrid cloud cluster information.
	//
	// example:
	//
	// 546
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end time of the query. The value is a UNIX timestamp (UTC). Unit: seconds.
	//
	// example:
	//
	// 1684382100
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Instance ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-z***9g301
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address to which the API operation belongs.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The name of the field used for sorting. Valid values:
	//
	// - **discoverTime**: the time when the risk was first detected. This is the default value.
	//
	// - **abnormalLevel**: the risk level.
	//
	// - **latestDiscoverTime**: the time when the risk was most recently detected.
	//
	// example:
	//
	// firstTime
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The sort order. Valid values:
	//
	// - **desc**: descending order. This is the default value.
	//
	// - **asc**: ascending order.
	//
	// example:
	//
	// desc
	OrderWay *string `json:"OrderWay,omitempty" xml:"OrderWay,omitempty"`
	// The source of the risk type. Valid values:
	//
	// - **custom**: custom.
	//
	// - **default**: built-in.
	//
	// example:
	//
	// custom
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The page number to return in a paging query. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paging query. Default value: **10**, which indicates 10 entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The start time of the query. The value is a UNIX timestamp (UTC). Unit: seconds.
	//
	// example:
	//
	// 1684252800
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The risk status. Valid values:
	//
	// - **toBeConfirmed**: to be confirmed.
	//
	// - **confirmed**: confirmed.
	//
	// - **toBeFixed**: to be fixed.
	//
	// - **fixed**: fixed (manually verified).
	//
	// - **ignored**: ignored.
	//
	// - **toBeVerified**: to be verified by the system.
	//
	// - **notFixed**: not fixed after verification.
	//
	// - **systemFixed**: fixed (verified by the system).
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
