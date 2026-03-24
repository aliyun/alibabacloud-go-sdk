// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecApiResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeApisecApiResourcesResponseBodyData) *DescribeApisecApiResourcesResponseBody
	GetData() []*DescribeApisecApiResourcesResponseBodyData
	SetRequestId(v string) *DescribeApisecApiResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeApisecApiResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeApisecApiResourcesResponseBody struct {
	// The list of API assets.
	Data []*DescribeApisecApiResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisecApiResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecApiResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisecApiResourcesResponseBody) GetData() []*DescribeApisecApiResourcesResponseBodyData {
	return s.Data
}

func (s *DescribeApisecApiResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApisecApiResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeApisecApiResourcesResponseBody) SetData(v []*DescribeApisecApiResourcesResponseBodyData) *DescribeApisecApiResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApisecApiResourcesResponseBody) SetRequestId(v string) *DescribeApisecApiResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBody) SetTotalCount(v int64) *DescribeApisecApiResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApisecApiResourcesResponseBodyData struct {
	// The number of threats associated with the API.
	//
	// example:
	//
	// 2
	AbnormalNum *int64 `json:"AbnormalNum,omitempty" xml:"AbnormalNum,omitempty"`
	// The number of account security events associated with the API.
	//
	// example:
	//
	// 1
	AccountEventNum *int64 `json:"AccountEventNum,omitempty" xml:"AccountEventNum,omitempty"`
	// The total number of requests in the last 30 days.
	//
	// example:
	//
	// 1683388800
	AllCnt *int64 `json:"AllCnt,omitempty" xml:"AllCnt,omitempty"`
	// The API endpoint path.
	//
	// example:
	//
	// /v1/etl/finddatabyvid
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API.
	//
	// example:
	//
	// 197b52abcd81d6a8bd4***e477
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The detailed information about the API. The value is a JSON string that contains the following fields:
	//
	// - **param_num**: the number of API parameters.
	//
	// - **request_method**: the request method.
	//
	// - **protocol**: the request protocol.
	//
	// - **api_url**: the request URL.
	//
	// - **poc_payload**: the request.
	//
	// - **request**: the request sample.
	//
	// - **response**: the response sample.
	//
	// - **param**: the request parameters.
	//
	// > This parameter is returned only when you specify the **ApiId*	- request parameter.
	ApiInfo *string `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty"`
	// The HTTP request method of the API. Valid values: **GET**, **POST**, **HEAD**, **PUT**, **DELETE**, **CONNECT**, **PATCH**, and **OPTIONS**.
	//
	// example:
	//
	// POST
	ApiMethod *string `json:"ApiMethod,omitempty" xml:"ApiMethod,omitempty"`
	// The sensitive data classification of the API. The value is a JSON string that contains the following fields:
	//
	// - **request_sensitive_list**: the list of sensitive data types in the request.
	//
	// - **response_sensitive_list**: the list of sensitive data types in the response.
	//
	// - **sensitive_list**: the list of sensitive data types.
	//
	// - **sensitive_level**: the sensitivity level.
	//
	// example:
	//
	// {
	//
	//     "sensitive_list": ["1003","1005"],
	//
	//     "sensitive_level": "L2",
	//
	//     "request_sensitive_list": ["1003"],
	//
	//     "response_sensitive_list": ["1005"]
	//
	// }
	ApiSensitive *string `json:"ApiSensitive,omitempty" xml:"ApiSensitive,omitempty"`
	// The types of sensitive data detected in the API request. The value is a JSON array of sensitive data type IDs.
	//
	// example:
	//
	// ["1002","1005"]
	ApiSensitiveRequest *string `json:"ApiSensitiveRequest,omitempty" xml:"ApiSensitiveRequest,omitempty"`
	// The types of sensitive data detected in the API response. The value is a JSON array of sensitive data type IDs.
	//
	// example:
	//
	// ["1002","1005"]
	ApiSensitiveResponse *string `json:"ApiSensitiveResponse,omitempty" xml:"ApiSensitiveResponse,omitempty"`
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
	// NewbornInterface
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
	// PublicAPI
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// Indicates whether the API requires authentication. Valid values:
	//
	// - **0**: The API requires authentication.
	//
	// - **1**: The API does not require authentication.
	//
	// example:
	//
	// 1
	AuthFlag *string `json:"AuthFlag,omitempty" xml:"AuthFlag,omitempty"`
	// The number of bot requests in the last 30 days.
	//
	// example:
	//
	// 2
	BotCnt *int64 `json:"BotCnt,omitempty" xml:"BotCnt,omitempty"`
	// The number of cross-border requests in the last 30 days.
	//
	// example:
	//
	// 2
	CrossBorderCnt *int64 `json:"CrossBorderCnt,omitempty" xml:"CrossBorderCnt,omitempty"`
	// The number of security events associated with the API.
	//
	// example:
	//
	// 2
	EventNum *int64 `json:"EventNum,omitempty" xml:"EventNum,omitempty"`
	// Deprecated
	//
	// The list of API samples.
	Examples []*string `json:"Examples,omitempty" xml:"Examples,omitempty" type:"Repeated"`
	// The time when the API was first discovered. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683388800
	FarthestTs *int64 `json:"FarthestTs,omitempty" xml:"FarthestTs,omitempty"`
	// Indicates whether the API is followed. Valid values:
	//
	// - **1**: The API is followed.
	//
	// - **0**: The API is not followed.
	//
	// example:
	//
	// 1
	Follow *int32 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The time of the most recent access to the API. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1683388800
	LastestTs *int64 `json:"LastestTs,omitempty" xml:"LastestTs,omitempty"`
	// The domain name or IP address that the API resides on.
	//
	// example:
	//
	// a.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The remarks of the API asset.
	//
	// example:
	//
	// loginApi
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The list of protected objects associated with the API.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s DescribeApisecApiResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecApiResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetAbnormalNum() *int64 {
	return s.AbnormalNum
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetAccountEventNum() *int64 {
	return s.AccountEventNum
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetAllCnt() *int64 {
	return s.AllCnt
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiInfo() *string {
	return s.ApiInfo
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiMethod() *string {
	return s.ApiMethod
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiSensitive() *string {
	return s.ApiSensitive
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiSensitiveRequest() *string {
	return s.ApiSensitiveRequest
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiSensitiveResponse() *string {
	return s.ApiSensitiveResponse
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiTag() *string {
	return s.ApiTag
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetApiType() *string {
	return s.ApiType
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetAuthFlag() *string {
	return s.AuthFlag
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetBotCnt() *int64 {
	return s.BotCnt
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetCrossBorderCnt() *int64 {
	return s.CrossBorderCnt
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetEventNum() *int64 {
	return s.EventNum
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetExamples() []*string {
	return s.Examples
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetFarthestTs() *int64 {
	return s.FarthestTs
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetFollow() *int32 {
	return s.Follow
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetLastestTs() *int64 {
	return s.LastestTs
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *DescribeApisecApiResourcesResponseBodyData) GetResources() []*string {
	return s.Resources
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetAbnormalNum(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.AbnormalNum = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetAccountEventNum(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.AccountEventNum = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetAllCnt(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.AllCnt = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiFormat(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiFormat = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiId(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiId = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiInfo(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiInfo = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiMethod(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiMethod = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiSensitive(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiSensitive = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiSensitiveRequest(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiSensitiveRequest = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiSensitiveResponse(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiSensitiveResponse = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiStatus(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiStatus = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiTag(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiTag = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetApiType(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.ApiType = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetAuthFlag(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.AuthFlag = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetBotCnt(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.BotCnt = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetCrossBorderCnt(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.CrossBorderCnt = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetEventNum(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.EventNum = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetExamples(v []*string) *DescribeApisecApiResourcesResponseBodyData {
	s.Examples = v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetFarthestTs(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.FarthestTs = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetFollow(v int32) *DescribeApisecApiResourcesResponseBodyData {
	s.Follow = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetLastestTs(v int64) *DescribeApisecApiResourcesResponseBodyData {
	s.LastestTs = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetMatchedHost(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetNote(v string) *DescribeApisecApiResourcesResponseBodyData {
	s.Note = &v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) SetResources(v []*string) *DescribeApisecApiResourcesResponseBodyData {
	s.Resources = v
	return s
}

func (s *DescribeApisecApiResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
