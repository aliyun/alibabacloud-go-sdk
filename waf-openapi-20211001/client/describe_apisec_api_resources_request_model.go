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
	// The API.
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
	// The request method of the API. Valid values:
	//
	// 	- **GET**
	//
	// 	- **POST**
	//
	// 	- **HEAD**
	//
	// 	- **PUT**
	//
	// 	- **DELETE**
	//
	// 	- **CONNECT**
	//
	// 	- **PATCH**
	//
	// 	- **OPTIONS**
	//
	// example:
	//
	// POST
	ApiMethod *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	// The API status. Valid values:
	//
	// 	- **NewbornInterface**: The API is newly added.
	//
	// 	- **OfflineInterface**: The API is inactive.
	//
	// 	- **normal**: The API is normal.
	//
	// example:
	//
	// OfflineInterface
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// The business purpose of the API.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the business purposes of APIs.
	//
	// example:
	//
	// SendMail
	ApiTag *string `json:"ApiTag,omitempty" xml:"ApiTag,omitempty"`
	// The service object. Valid values:
	//
	// 	- **PublicAPI**: public services
	//
	// 	- **ThirdpartAPI**: cooperation with third-party partners
	//
	// 	- **InternalAPI**: internal office
	//
	// example:
	//
	// innerAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// Specifies whether authentication is required. Valid values:
	//
	// 	- **0**: Authentication is required.
	//
	// 	- **1**: Authentication is not required.
	//
	// example:
	//
	// 0
	AuthFlag *string `json:"AuthFlag,omitempty" xml:"AuthFlag,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// >  This parameter is available only in hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query hybrid cloud clusters.
	//
	// example:
	//
	// 740
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
	//
	// example:
	//
	// 1683388800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether to follow the API. Valid values:
	//
	// 	- **1**: follows the API.
	//
	// 	- **0**: does not follow the API.
	//
	// example:
	//
	// 0
	Follow *int64 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-u***gr20j
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address of the API.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The remarks.
	//
	// example:
	//
	// API for logon
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The name of the sorting field. Valid values:
	//
	// 	- **allCnt**: the total number of calls to the API in the previous 30 days
	//
	// 	- **botCnt**: the number of bot-initiated requests in the previous 30 days
	//
	// 	- **crossBorderCnt**: the number of cross-border requests in the previous 30 days
	//
	// 	- **abnormalNum**: the number of API-related risks
	//
	// 	- **eventNum**: the number of API-related security events
	//
	// 	- **farthestTs**: the time when the API was first detected
	//
	// 	- **lastestTs**: the time of the most recent access to the API
	//
	// example:
	//
	// allCnt
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
	// The sensitive data type in the request.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported sensitive data types.
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
	// 	- **L1**: high sensitivity
	//
	// 	- **L2**: moderate sensitivity
	//
	// 	- **L3**: low sensitivity
	//
	// 	- **N**: non-sensitivity
	//
	// example:
	//
	// L3
	SensitiveLevel *string `json:"SensitiveLevel,omitempty" xml:"SensitiveLevel,omitempty"`
	// The sensitive data type in the response.
	//
	// >  You can call the [DescribeApisecRules](https://help.aliyun.com/document_detail/2859155.html) operation to query the supported sensitive data types.
	//
	// example:
	//
	// 1004
	SensitiveType *string `json:"SensitiveType,omitempty" xml:"SensitiveType,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp in UTC. Unit: seconds.
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
