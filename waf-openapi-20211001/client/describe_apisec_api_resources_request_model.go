// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecApiResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiFormat(v string) *DescribeApisecApiResourcesRequest
	GetApiFormat() *string
	SetApiId(v string) *DescribeApisecApiResourcesRequest
	GetApiId() *string
	SetApiMethod(v string) *DescribeApisecApiResourcesRequest
	GetApiMethod() *string
	SetApiStatus(v string) *DescribeApisecApiResourcesRequest
	GetApiStatus() *string
	SetApiTag(v string) *DescribeApisecApiResourcesRequest
	GetApiTag() *string
	SetApiType(v string) *DescribeApisecApiResourcesRequest
	GetApiType() *string
	SetAuthFlag(v string) *DescribeApisecApiResourcesRequest
	GetAuthFlag() *string
	SetClusterId(v string) *DescribeApisecApiResourcesRequest
	GetClusterId() *string
	SetEndTime(v string) *DescribeApisecApiResourcesRequest
	GetEndTime() *string
	SetFollow(v int64) *DescribeApisecApiResourcesRequest
	GetFollow() *int64
	SetInstanceId(v string) *DescribeApisecApiResourcesRequest
	GetInstanceId() *string
	SetMatchedHost(v string) *DescribeApisecApiResourcesRequest
	GetMatchedHost() *string
	SetNote(v string) *DescribeApisecApiResourcesRequest
	GetNote() *string
	SetOrderKey(v string) *DescribeApisecApiResourcesRequest
	GetOrderKey() *string
	SetOrderWay(v string) *DescribeApisecApiResourcesRequest
	GetOrderWay() *string
	SetPageNumber(v int64) *DescribeApisecApiResourcesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecApiResourcesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecApiResourcesRequest
	GetRegionId() *string
	SetRequestSensitiveType(v string) *DescribeApisecApiResourcesRequest
	GetRequestSensitiveType() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecApiResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetSensitiveLevel(v string) *DescribeApisecApiResourcesRequest
	GetSensitiveLevel() *string
	SetSensitiveType(v string) *DescribeApisecApiResourcesRequest
	GetSensitiveType() *string
	SetStartTime(v string) *DescribeApisecApiResourcesRequest
	GetStartTime() *string
}

type DescribeApisecApiResourcesRequest struct {
	// The API endpoint path used to filter the query results.
	//
	// example:
	//
	// /auth/login
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API.
	//
	// example:
	//
	// 867ade***24ee6e205b8da82b8f84
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The HTTP request method of the API. Valid values: **GET**, **POST**, **HEAD**, **PUT**, **DELETE**, **CONNECT**, **PATCH**, and **OPTIONS**.
	//
	// example:
	//
	// POST
	ApiMethod *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	// The lifecycle status of the API. Valid values:
	//
	// - **NewbornInterface**: newly discovered.
	//
	// - **OfflineInterface**: inactive.
	//
	// - **normal**: active.
	//
	// example:
	//
	// OfflineInterface
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// The business purpose of the API.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to obtain the supported business purposes.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The type of service that the API serves. Valid values:
	//
	// - **PublicAPI**: public-facing service.
	//
	// - **ThirdpartAPI**: third-party service.
	//
	// - **InternalAPI**: internal service.
	//
	// example:
	//
	// innerAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// Indicates whether the API requires authentication. Valid values:
	//
	// - **0**: The API requires authentication.
	//
	// - **1**: The API does not require authentication.
	//
	// example:
	//
	// 0
	AuthFlag *string `json:"AuthFlag,omitempty" xml:"AuthFlag,omitempty"`
	// The ID of the Hybrid Cloud WAF cluster.
	//
	// > This parameter is available only for hybrid cloud scenarios. Call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to obtain information about Hybrid Cloud WAF clusters.
	//
	// example:
	//
	// 740
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1683388800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the API is followed. Valid values:
	//
	// - **1**: The API is followed.
	//
	// - **0**: The API is not followed.
	//
	// example:
	//
	// 0
	Follow *int64 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-u***gr20j
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address of the API used to filter the query results.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The remarks of the API asset used to filter the query results.
	//
	// example:
	//
	// loginApi
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The field by which to sort the results. Valid values:
	//
	// - **allCnt**: sorts by the total number of requests in the last 30 days.
	//
	// - **botCnt**: sorts by the number of bot requests in the last 30 days.
	//
	// - **crossBorderCnt**: sorts by the number of cross-border requests in the last 30 days.
	//
	// - **abnormalNum**: sorts by the number of threats associated with the API.
	//
	// - **eventNum**: sorts by the number of security events associated with the API.
	//
	// - **farthestTs**: sorts by the time when the API was first discovered.
	//
	// - **lastestTs**: sorts by the time of the most recent access.
	//
	// example:
	//
	// allCnt
	OrderKey *string `json:"OrderKey,omitempty" xml:"OrderKey,omitempty"`
	// The sort order. Valid values:
	//
	// - **desc**: descending order (default).
	//
	// - **asc**: ascending order.
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
	// The region ID of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of sensitive data in the request.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to obtain the supported sensitive data types.
	//
	// example:
	//
	// 1004,1005
	RequestSensitiveType *string `json:"RequestSensitiveType,omitempty" xml:"RequestSensitiveType,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The sensitivity level of the API. Valid values:
	//
	// - **L1**: High.
	//
	// - **L2**: Medium.
	//
	// - **L3**: Low.
	//
	// - **N**: Non-sensitive.
	//
	// example:
	//
	// L3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The type of sensitive data in the response.
	//
	// > Call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to obtain the supported sensitive data types.
	//
	// example:
	//
	// 1004
	SensitiveType *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1681833600
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApisecApiResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecApiResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecApiResourcesRequest) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecApiResourcesRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecApiResourcesRequest) GetApiMethod() *string {
	return s.ApiMethod
}

func (s *DescribeApisecApiResourcesRequest) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *DescribeApisecApiResourcesRequest) GetApiTag() *string {
	return s.ApiTag
}

func (s *DescribeApisecApiResourcesRequest) GetApiType() *string {
	return s.ApiType
}

func (s *DescribeApisecApiResourcesRequest) GetAuthFlag() *string {
	return s.AuthFlag
}

func (s *DescribeApisecApiResourcesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecApiResourcesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeApisecApiResourcesRequest) GetFollow() *int64 {
	return s.Follow
}

func (s *DescribeApisecApiResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecApiResourcesRequest) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecApiResourcesRequest) GetNote() *string {
	return s.Note
}

func (s *DescribeApisecApiResourcesRequest) GetOrderKey() *string {
	return s.OrderKey
}

func (s *DescribeApisecApiResourcesRequest) GetOrderWay() *string {
	return s.OrderWay
}

func (s *DescribeApisecApiResourcesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecApiResourcesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecApiResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecApiResourcesRequest) GetRequestSensitiveType() *string {
	return s.RequestSensitiveType
}

func (s *DescribeApisecApiResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecApiResourcesRequest) GetSensitiveLevel() *string {
	return s.SensitiveLevel
}

func (s *DescribeApisecApiResourcesRequest) GetSensitiveType() *string {
	return s.SensitiveType
}

func (s *DescribeApisecApiResourcesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeApisecApiResourcesRequest) SetApiFormat(v string) *DescribeApisecApiResourcesRequest {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetApiId(v string) *DescribeApisecApiResourcesRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetApiMethod(v string) *DescribeApisecApiResourcesRequest {
	s.ApiMethod = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetApiStatus(v string) *DescribeApisecApiResourcesRequest {
	s.ApiStatus = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetApiTag(v string) *DescribeApisecApiResourcesRequest {
	s.ApiTag = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetApiType(v string) *DescribeApisecApiResourcesRequest {
	s.ApiType = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetAuthFlag(v string) *DescribeApisecApiResourcesRequest {
	s.AuthFlag = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetClusterId(v string) *DescribeApisecApiResourcesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetEndTime(v string) *DescribeApisecApiResourcesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetFollow(v int64) *DescribeApisecApiResourcesRequest {
	s.Follow = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetInstanceId(v string) *DescribeApisecApiResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetMatchedHost(v string) *DescribeApisecApiResourcesRequest {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetNote(v string) *DescribeApisecApiResourcesRequest {
	s.Note = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetOrderKey(v string) *DescribeApisecApiResourcesRequest {
	s.OrderKey = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetOrderWay(v string) *DescribeApisecApiResourcesRequest {
	s.OrderWay = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetPageNumber(v int64) *DescribeApisecApiResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetPageSize(v int64) *DescribeApisecApiResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetRegionId(v string) *DescribeApisecApiResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetRequestSensitiveType(v string) *DescribeApisecApiResourcesRequest {
	s.RequestSensitiveType = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecApiResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetSensitiveLevel(v string) *DescribeApisecApiResourcesRequest {
	s.SensitiveLevel = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetSensitiveType(v string) *DescribeApisecApiResourcesRequest {
	s.SensitiveType = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) SetStartTime(v string) *DescribeApisecApiResourcesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeApisecApiResourcesRequest) Validate() error {
	return dara.Validate(s)
}
