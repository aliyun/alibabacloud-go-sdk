// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AbolishApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AbolishApiRequest) String() string {
	return tea.Prettify(s)
}

func (s AbolishApiRequest) GoString() string {
	return s.String()
}

func (s *AbolishApiRequest) SetApiId(v string) *AbolishApiRequest {
	s.ApiId = &v
	return s
}

func (s *AbolishApiRequest) SetGroupId(v string) *AbolishApiRequest {
	s.GroupId = &v
	return s
}

func (s *AbolishApiRequest) SetSecurityToken(v string) *AbolishApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *AbolishApiRequest) SetStageName(v string) *AbolishApiRequest {
	s.StageName = &v
	return s
}

type AbolishApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbolishApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbolishApiResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishApiResponseBody) SetRequestId(v string) *AbolishApiResponseBody {
	s.RequestId = &v
	return s
}

type AbolishApiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishApiResponse) String() string {
	return tea.Prettify(s)
}

func (s AbolishApiResponse) GoString() string {
	return s.String()
}

func (s *AbolishApiResponse) SetHeaders(v map[string]*string) *AbolishApiResponse {
	s.Headers = v
	return s
}

func (s *AbolishApiResponse) SetStatusCode(v int32) *AbolishApiResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishApiResponse) SetBody(v *AbolishApiResponseBody) *AbolishApiResponse {
	s.Body = v
	return s
}

type AbolishApiForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AbolishApiForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s AbolishApiForInnerRequest) GoString() string {
	return s.String()
}

func (s *AbolishApiForInnerRequest) SetAliUid(v int64) *AbolishApiForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *AbolishApiForInnerRequest) SetApiId(v string) *AbolishApiForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *AbolishApiForInnerRequest) SetGroupId(v string) *AbolishApiForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *AbolishApiForInnerRequest) SetSecurityToken(v string) *AbolishApiForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *AbolishApiForInnerRequest) SetStageName(v string) *AbolishApiForInnerRequest {
	s.StageName = &v
	return s
}

type AbolishApiForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AbolishApiForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbolishApiForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *AbolishApiForInnerResponseBody) SetRequestId(v string) *AbolishApiForInnerResponseBody {
	s.RequestId = &v
	return s
}

type AbolishApiForInnerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishApiForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishApiForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s AbolishApiForInnerResponse) GoString() string {
	return s.String()
}

func (s *AbolishApiForInnerResponse) SetHeaders(v map[string]*string) *AbolishApiForInnerResponse {
	s.Headers = v
	return s
}

func (s *AbolishApiForInnerResponse) SetStatusCode(v int32) *AbolishApiForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishApiForInnerResponse) SetBody(v *AbolishApiForInnerResponseBody) *AbolishApiForInnerResponse {
	s.Body = v
	return s
}

type AddBlackListRequest struct {
	BlackContent *string `json:"BlackContent,omitempty" xml:"BlackContent,omitempty"`
	BlackType    *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s AddBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBlackListRequest) GoString() string {
	return s.String()
}

func (s *AddBlackListRequest) SetBlackContent(v string) *AddBlackListRequest {
	s.BlackContent = &v
	return s
}

func (s *AddBlackListRequest) SetBlackType(v string) *AddBlackListRequest {
	s.BlackType = &v
	return s
}

func (s *AddBlackListRequest) SetDescription(v string) *AddBlackListRequest {
	s.Description = &v
	return s
}

type AddBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *AddBlackListResponseBody) SetRequestId(v string) *AddBlackListResponseBody {
	s.RequestId = &v
	return s
}

type AddBlackListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBlackListResponse) GoString() string {
	return s.String()
}

func (s *AddBlackListResponse) SetHeaders(v map[string]*string) *AddBlackListResponse {
	s.Headers = v
	return s
}

func (s *AddBlackListResponse) SetStatusCode(v int32) *AddBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBlackListResponse) SetBody(v *AddBlackListResponseBody) *AddBlackListResponse {
	s.Body = v
	return s
}

type AddTrafficSpecialControlRequest struct {
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SpecialKey       *string `json:"SpecialKey,omitempty" xml:"SpecialKey,omitempty"`
	SpecialType      *string `json:"SpecialType,omitempty" xml:"SpecialType,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	TrafficValue     *int32  `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s AddTrafficSpecialControlRequest) String() string {
	return tea.Prettify(s)
}

func (s AddTrafficSpecialControlRequest) GoString() string {
	return s.String()
}

func (s *AddTrafficSpecialControlRequest) SetSecurityToken(v string) *AddTrafficSpecialControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetSpecialKey(v string) *AddTrafficSpecialControlRequest {
	s.SpecialKey = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetSpecialType(v string) *AddTrafficSpecialControlRequest {
	s.SpecialType = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetTrafficControlId(v string) *AddTrafficSpecialControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *AddTrafficSpecialControlRequest) SetTrafficValue(v int32) *AddTrafficSpecialControlRequest {
	s.TrafficValue = &v
	return s
}

type AddTrafficSpecialControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTrafficSpecialControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddTrafficSpecialControlResponseBody) GoString() string {
	return s.String()
}

func (s *AddTrafficSpecialControlResponseBody) SetRequestId(v string) *AddTrafficSpecialControlResponseBody {
	s.RequestId = &v
	return s
}

type AddTrafficSpecialControlResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTrafficSpecialControlResponse) String() string {
	return tea.Prettify(s)
}

func (s AddTrafficSpecialControlResponse) GoString() string {
	return s.String()
}

func (s *AddTrafficSpecialControlResponse) SetHeaders(v map[string]*string) *AddTrafficSpecialControlResponse {
	s.Headers = v
	return s
}

func (s *AddTrafficSpecialControlResponse) SetStatusCode(v int32) *AddTrafficSpecialControlResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTrafficSpecialControlResponse) SetBody(v *AddTrafficSpecialControlResponseBody) *AddTrafficSpecialControlResponse {
	s.Body = v
	return s
}

type CheckAccountForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckAccountForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountForInnerRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountForInnerRequest) SetAliUid(v int64) *CheckAccountForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *CheckAccountForInnerRequest) SetSecurityToken(v string) *CheckAccountForInnerRequest {
	s.SecurityToken = &v
	return s
}

type CheckAccountForInnerResponseBody struct {
	CheckResult *bool   `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAccountForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAccountForInnerResponseBody) SetCheckResult(v bool) *CheckAccountForInnerResponseBody {
	s.CheckResult = &v
	return s
}

func (s *CheckAccountForInnerResponseBody) SetRequestId(v string) *CheckAccountForInnerResponseBody {
	s.RequestId = &v
	return s
}

type CheckAccountForInnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAccountForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAccountForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAccountForInnerResponse) GoString() string {
	return s.String()
}

func (s *CheckAccountForInnerResponse) SetHeaders(v map[string]*string) *CheckAccountForInnerResponse {
	s.Headers = v
	return s
}

func (s *CheckAccountForInnerResponse) SetStatusCode(v int32) *CheckAccountForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAccountForInnerResponse) SetBody(v *CheckAccountForInnerResponseBody) *CheckAccountForInnerResponse {
	s.Body = v
	return s
}

type CheckAoneAppAuditRequest struct {
	AoneAppName   *string `json:"AoneAppName,omitempty" xml:"AoneAppName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckAoneAppAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckAoneAppAuditRequest) GoString() string {
	return s.String()
}

func (s *CheckAoneAppAuditRequest) SetAoneAppName(v string) *CheckAoneAppAuditRequest {
	s.AoneAppName = &v
	return s
}

func (s *CheckAoneAppAuditRequest) SetSecurityToken(v string) *CheckAoneAppAuditRequest {
	s.SecurityToken = &v
	return s
}

type CheckAoneAppAuditResponseBody struct {
	CheckResult *bool   `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckAoneAppAuditResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckAoneAppAuditResponseBody) GoString() string {
	return s.String()
}

func (s *CheckAoneAppAuditResponseBody) SetCheckResult(v bool) *CheckAoneAppAuditResponseBody {
	s.CheckResult = &v
	return s
}

func (s *CheckAoneAppAuditResponseBody) SetRequestId(v string) *CheckAoneAppAuditResponseBody {
	s.RequestId = &v
	return s
}

type CheckAoneAppAuditResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckAoneAppAuditResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckAoneAppAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckAoneAppAuditResponse) GoString() string {
	return s.String()
}

func (s *CheckAoneAppAuditResponse) SetHeaders(v map[string]*string) *CheckAoneAppAuditResponse {
	s.Headers = v
	return s
}

func (s *CheckAoneAppAuditResponse) SetStatusCode(v int32) *CheckAoneAppAuditResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckAoneAppAuditResponse) SetBody(v *CheckAoneAppAuditResponseBody) *CheckAoneAppAuditResponse {
	s.Body = v
	return s
}

type CopyConsumerOpenForInnerRequest struct {
	CopyList      *string `json:"CopyList,omitempty" xml:"CopyList,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CopyConsumerOpenForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyConsumerOpenForInnerRequest) GoString() string {
	return s.String()
}

func (s *CopyConsumerOpenForInnerRequest) SetCopyList(v string) *CopyConsumerOpenForInnerRequest {
	s.CopyList = &v
	return s
}

func (s *CopyConsumerOpenForInnerRequest) SetSecurityToken(v string) *CopyConsumerOpenForInnerRequest {
	s.SecurityToken = &v
	return s
}

type CopyConsumerOpenForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyConsumerOpenForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyConsumerOpenForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *CopyConsumerOpenForInnerResponseBody) SetRequestId(v string) *CopyConsumerOpenForInnerResponseBody {
	s.RequestId = &v
	return s
}

type CopyConsumerOpenForInnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyConsumerOpenForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyConsumerOpenForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyConsumerOpenForInnerResponse) GoString() string {
	return s.String()
}

func (s *CopyConsumerOpenForInnerResponse) SetHeaders(v map[string]*string) *CopyConsumerOpenForInnerResponse {
	s.Headers = v
	return s
}

func (s *CopyConsumerOpenForInnerResponse) SetStatusCode(v int32) *CopyConsumerOpenForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyConsumerOpenForInnerResponse) SetBody(v *CopyConsumerOpenForInnerResponseBody) *CopyConsumerOpenForInnerResponse {
	s.Body = v
	return s
}

type CreateApiRequest struct {
	ApiName             *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType            *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	ConstantParameters  *string `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HttpMethod          *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol        *string `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters      *string `json:"PathParameters,omitempty" xml:"PathParameters,omitempty"`
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType        *string `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RequestBody         *string `json:"RequestBody,omitempty" xml:"RequestBody,omitempty"`
	RequestHeaders      *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	RequestQueries      *string `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty"`
	ResultSample        *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType          *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceAddress      *string `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceProtocol     *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout      *int32  `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	SystemParameters    *string `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty"`
	Visibility          *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiRequest) GoString() string {
	return s.String()
}

func (s *CreateApiRequest) SetApiName(v string) *CreateApiRequest {
	s.ApiName = &v
	return s
}

func (s *CreateApiRequest) SetAuthType(v string) *CreateApiRequest {
	s.AuthType = &v
	return s
}

func (s *CreateApiRequest) SetBodyFormat(v string) *CreateApiRequest {
	s.BodyFormat = &v
	return s
}

func (s *CreateApiRequest) SetConstantParameters(v string) *CreateApiRequest {
	s.ConstantParameters = &v
	return s
}

func (s *CreateApiRequest) SetDescription(v string) *CreateApiRequest {
	s.Description = &v
	return s
}

func (s *CreateApiRequest) SetGroupId(v string) *CreateApiRequest {
	s.GroupId = &v
	return s
}

func (s *CreateApiRequest) SetHttpMethod(v string) *CreateApiRequest {
	s.HttpMethod = &v
	return s
}

func (s *CreateApiRequest) SetHttpProtocol(v string) *CreateApiRequest {
	s.HttpProtocol = &v
	return s
}

func (s *CreateApiRequest) SetPath(v string) *CreateApiRequest {
	s.Path = &v
	return s
}

func (s *CreateApiRequest) SetPathParameters(v string) *CreateApiRequest {
	s.PathParameters = &v
	return s
}

func (s *CreateApiRequest) SetPostBodyDescription(v string) *CreateApiRequest {
	s.PostBodyDescription = &v
	return s
}

func (s *CreateApiRequest) SetPostBodyType(v string) *CreateApiRequest {
	s.PostBodyType = &v
	return s
}

func (s *CreateApiRequest) SetRequestBody(v string) *CreateApiRequest {
	s.RequestBody = &v
	return s
}

func (s *CreateApiRequest) SetRequestHeaders(v string) *CreateApiRequest {
	s.RequestHeaders = &v
	return s
}

func (s *CreateApiRequest) SetRequestQueries(v string) *CreateApiRequest {
	s.RequestQueries = &v
	return s
}

func (s *CreateApiRequest) SetResultSample(v string) *CreateApiRequest {
	s.ResultSample = &v
	return s
}

func (s *CreateApiRequest) SetResultType(v string) *CreateApiRequest {
	s.ResultType = &v
	return s
}

func (s *CreateApiRequest) SetSecurityToken(v string) *CreateApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateApiRequest) SetServiceAddress(v string) *CreateApiRequest {
	s.ServiceAddress = &v
	return s
}

func (s *CreateApiRequest) SetServiceProtocol(v string) *CreateApiRequest {
	s.ServiceProtocol = &v
	return s
}

func (s *CreateApiRequest) SetServiceTimeout(v int32) *CreateApiRequest {
	s.ServiceTimeout = &v
	return s
}

func (s *CreateApiRequest) SetSystemParameters(v string) *CreateApiRequest {
	s.SystemParameters = &v
	return s
}

func (s *CreateApiRequest) SetVisibility(v string) *CreateApiRequest {
	s.Visibility = &v
	return s
}

type CreateApiResponseBody struct {
	ApiId     *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiResponseBody) SetApiId(v string) *CreateApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *CreateApiResponseBody) SetRequestId(v string) *CreateApiResponseBody {
	s.RequestId = &v
	return s
}

type CreateApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiResponse) GoString() string {
	return s.String()
}

func (s *CreateApiResponse) SetHeaders(v map[string]*string) *CreateApiResponse {
	s.Headers = v
	return s
}

func (s *CreateApiResponse) SetStatusCode(v int32) *CreateApiResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiResponse) SetBody(v *CreateApiResponseBody) *CreateApiResponse {
	s.Body = v
	return s
}

type CreateApiForInnerRequest struct {
	AliUid               *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestConfig        *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParamters     *string `json:"RequestParamters,omitempty" xml:"RequestParamters,omitempty"`
	ResultSample         *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType           *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	Visibility           *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateApiForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiForInnerRequest) GoString() string {
	return s.String()
}

func (s *CreateApiForInnerRequest) SetAliUid(v int64) *CreateApiForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *CreateApiForInnerRequest) SetApiName(v string) *CreateApiForInnerRequest {
	s.ApiName = &v
	return s
}

func (s *CreateApiForInnerRequest) SetAuthType(v string) *CreateApiForInnerRequest {
	s.AuthType = &v
	return s
}

func (s *CreateApiForInnerRequest) SetDescription(v string) *CreateApiForInnerRequest {
	s.Description = &v
	return s
}

func (s *CreateApiForInnerRequest) SetGroupId(v string) *CreateApiForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *CreateApiForInnerRequest) SetRequestConfig(v string) *CreateApiForInnerRequest {
	s.RequestConfig = &v
	return s
}

func (s *CreateApiForInnerRequest) SetRequestParamters(v string) *CreateApiForInnerRequest {
	s.RequestParamters = &v
	return s
}

func (s *CreateApiForInnerRequest) SetResultSample(v string) *CreateApiForInnerRequest {
	s.ResultSample = &v
	return s
}

func (s *CreateApiForInnerRequest) SetResultType(v string) *CreateApiForInnerRequest {
	s.ResultType = &v
	return s
}

func (s *CreateApiForInnerRequest) SetSecurityToken(v string) *CreateApiForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateApiForInnerRequest) SetServiceConfig(v string) *CreateApiForInnerRequest {
	s.ServiceConfig = &v
	return s
}

func (s *CreateApiForInnerRequest) SetServiceParameters(v string) *CreateApiForInnerRequest {
	s.ServiceParameters = &v
	return s
}

func (s *CreateApiForInnerRequest) SetServiceParametersMap(v string) *CreateApiForInnerRequest {
	s.ServiceParametersMap = &v
	return s
}

func (s *CreateApiForInnerRequest) SetVisibility(v string) *CreateApiForInnerRequest {
	s.Visibility = &v
	return s
}

type CreateApiForInnerResponseBody struct {
	ApiId     *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiForInnerResponseBody) SetApiId(v string) *CreateApiForInnerResponseBody {
	s.ApiId = &v
	return s
}

func (s *CreateApiForInnerResponseBody) SetRequestId(v string) *CreateApiForInnerResponseBody {
	s.RequestId = &v
	return s
}

type CreateApiForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiForInnerResponse) GoString() string {
	return s.String()
}

func (s *CreateApiForInnerResponse) SetHeaders(v map[string]*string) *CreateApiForInnerResponse {
	s.Headers = v
	return s
}

func (s *CreateApiForInnerResponse) SetStatusCode(v int32) *CreateApiForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiForInnerResponse) SetBody(v *CreateApiForInnerResponseBody) *CreateApiForInnerResponse {
	s.Body = v
	return s
}

type CreateApiGroupRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateApiGroupRequest) SetDescription(v string) *CreateApiGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateApiGroupRequest) SetGroupName(v string) *CreateApiGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateApiGroupRequest) SetSecurityToken(v string) *CreateApiGroupRequest {
	s.SecurityToken = &v
	return s
}

type CreateApiGroupResponseBody struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain   *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s CreateApiGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiGroupResponseBody) SetDescription(v string) *CreateApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetGroupId(v string) *CreateApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetGroupName(v string) *CreateApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetRequestId(v string) *CreateApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetSubDomain(v string) *CreateApiGroupResponseBody {
	s.SubDomain = &v
	return s
}

type CreateApiGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateApiGroupResponse) SetHeaders(v map[string]*string) *CreateApiGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateApiGroupResponse) SetStatusCode(v int32) *CreateApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiGroupResponse) SetBody(v *CreateApiGroupResponseBody) *CreateApiGroupResponse {
	s.Body = v
	return s
}

type CreateApiGroupForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateApiGroupForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupForInnerRequest) GoString() string {
	return s.String()
}

func (s *CreateApiGroupForInnerRequest) SetAliUid(v int64) *CreateApiGroupForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *CreateApiGroupForInnerRequest) SetDescription(v string) *CreateApiGroupForInnerRequest {
	s.Description = &v
	return s
}

func (s *CreateApiGroupForInnerRequest) SetGroupName(v string) *CreateApiGroupForInnerRequest {
	s.GroupName = &v
	return s
}

func (s *CreateApiGroupForInnerRequest) SetSecurityToken(v string) *CreateApiGroupForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateApiGroupForInnerRequest) SetSource(v string) *CreateApiGroupForInnerRequest {
	s.Source = &v
	return s
}

type CreateApiGroupForInnerResponseBody struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain   *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s CreateApiGroupForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiGroupForInnerResponseBody) SetDescription(v string) *CreateApiGroupForInnerResponseBody {
	s.Description = &v
	return s
}

func (s *CreateApiGroupForInnerResponseBody) SetGroupId(v string) *CreateApiGroupForInnerResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateApiGroupForInnerResponseBody) SetGroupName(v string) *CreateApiGroupForInnerResponseBody {
	s.GroupName = &v
	return s
}

func (s *CreateApiGroupForInnerResponseBody) SetRequestId(v string) *CreateApiGroupForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApiGroupForInnerResponseBody) SetSubDomain(v string) *CreateApiGroupForInnerResponseBody {
	s.SubDomain = &v
	return s
}

type CreateApiGroupForInnerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApiGroupForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApiGroupForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupForInnerResponse) GoString() string {
	return s.String()
}

func (s *CreateApiGroupForInnerResponse) SetHeaders(v map[string]*string) *CreateApiGroupForInnerResponse {
	s.Headers = v
	return s
}

func (s *CreateApiGroupForInnerResponse) SetStatusCode(v int32) *CreateApiGroupForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiGroupForInnerResponse) SetBody(v *CreateApiGroupForInnerResponseBody) *CreateApiGroupForInnerResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetSecurityToken(v string) *CreateAppRequest {
	s.SecurityToken = &v
	return s
}

type CreateAppResponseBody struct {
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetAppId(v int64) *CreateAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateAppForBackendRequest struct {
	AliUid      *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateAppForBackendRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppForBackendRequest) GoString() string {
	return s.String()
}

func (s *CreateAppForBackendRequest) SetAliUid(v int64) *CreateAppForBackendRequest {
	s.AliUid = &v
	return s
}

func (s *CreateAppForBackendRequest) SetAppName(v string) *CreateAppForBackendRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppForBackendRequest) SetDescription(v string) *CreateAppForBackendRequest {
	s.Description = &v
	return s
}

func (s *CreateAppForBackendRequest) SetSource(v string) *CreateAppForBackendRequest {
	s.Source = &v
	return s
}

type CreateAppForBackendResponseBody struct {
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppForBackendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppForBackendResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppForBackendResponseBody) SetAppId(v int64) *CreateAppForBackendResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppForBackendResponseBody) SetRequestId(v string) *CreateAppForBackendResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppForBackendResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppForBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppForBackendResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppForBackendResponse) GoString() string {
	return s.String()
}

func (s *CreateAppForBackendResponse) SetHeaders(v map[string]*string) *CreateAppForBackendResponse {
	s.Headers = v
	return s
}

func (s *CreateAppForBackendResponse) SetStatusCode(v int32) *CreateAppForBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppForBackendResponse) SetBody(v *CreateAppForBackendResponseBody) *CreateAppForBackendResponse {
	s.Body = v
	return s
}

type CreateAppForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppCode       *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppSecret     *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Extend        *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateAppForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppForInnerRequest) GoString() string {
	return s.String()
}

func (s *CreateAppForInnerRequest) SetAliUid(v int64) *CreateAppForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *CreateAppForInnerRequest) SetAppCode(v string) *CreateAppForInnerRequest {
	s.AppCode = &v
	return s
}

func (s *CreateAppForInnerRequest) SetAppKey(v string) *CreateAppForInnerRequest {
	s.AppKey = &v
	return s
}

func (s *CreateAppForInnerRequest) SetAppName(v string) *CreateAppForInnerRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppForInnerRequest) SetAppSecret(v string) *CreateAppForInnerRequest {
	s.AppSecret = &v
	return s
}

func (s *CreateAppForInnerRequest) SetDescription(v string) *CreateAppForInnerRequest {
	s.Description = &v
	return s
}

func (s *CreateAppForInnerRequest) SetExtend(v string) *CreateAppForInnerRequest {
	s.Extend = &v
	return s
}

func (s *CreateAppForInnerRequest) SetSecurityToken(v string) *CreateAppForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAppForInnerRequest) SetSource(v string) *CreateAppForInnerRequest {
	s.Source = &v
	return s
}

type CreateAppForInnerResponseBody struct {
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppForInnerResponseBody) SetAppId(v int64) *CreateAppForInnerResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateAppForInnerResponseBody) SetRequestId(v string) *CreateAppForInnerResponseBody {
	s.RequestId = &v
	return s
}

type CreateAppForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppForInnerResponse) GoString() string {
	return s.String()
}

func (s *CreateAppForInnerResponse) SetHeaders(v map[string]*string) *CreateAppForInnerResponse {
	s.Headers = v
	return s
}

func (s *CreateAppForInnerResponse) SetStatusCode(v int32) *CreateAppForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppForInnerResponse) SetBody(v *CreateAppForInnerResponseBody) *CreateAppForInnerResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	AccountQuantity       *int64  `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	AlarmQuota            *int64  `json:"AlarmQuota,omitempty" xml:"AlarmQuota,omitempty"`
	AliUid                *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId                 *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BillingType           *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CloudMarketInstanceId *string `json:"CloudMarketInstanceId,omitempty" xml:"CloudMarketInstanceId,omitempty"`
	ExpiredOn             *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	InstanceAttributes    *string `json:"InstanceAttributes,omitempty" xml:"InstanceAttributes,omitempty"`
	SkuId                 *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Token                 *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAccountQuantity(v int64) *CreateInstanceRequest {
	s.AccountQuantity = &v
	return s
}

func (s *CreateInstanceRequest) SetAlarmQuota(v int64) *CreateInstanceRequest {
	s.AlarmQuota = &v
	return s
}

func (s *CreateInstanceRequest) SetAliUid(v int64) *CreateInstanceRequest {
	s.AliUid = &v
	return s
}

func (s *CreateInstanceRequest) SetAppId(v int64) *CreateInstanceRequest {
	s.AppId = &v
	return s
}

func (s *CreateInstanceRequest) SetBillingType(v string) *CreateInstanceRequest {
	s.BillingType = &v
	return s
}

func (s *CreateInstanceRequest) SetCloudMarketInstanceId(v string) *CreateInstanceRequest {
	s.CloudMarketInstanceId = &v
	return s
}

func (s *CreateInstanceRequest) SetExpiredOn(v string) *CreateInstanceRequest {
	s.ExpiredOn = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceAttributes(v string) *CreateInstanceRequest {
	s.InstanceAttributes = &v
	return s
}

func (s *CreateInstanceRequest) SetSkuId(v string) *CreateInstanceRequest {
	s.SkuId = &v
	return s
}

func (s *CreateInstanceRequest) SetToken(v string) *CreateInstanceRequest {
	s.Token = &v
	return s
}

type CreateInstanceResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateRaceWorkForInnerRequest struct {
	CommodityCode    *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Keywords         *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	LogoUrl          *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	WorkName         *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
}

func (s CreateRaceWorkForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRaceWorkForInnerRequest) GoString() string {
	return s.String()
}

func (s *CreateRaceWorkForInnerRequest) SetCommodityCode(v string) *CreateRaceWorkForInnerRequest {
	s.CommodityCode = &v
	return s
}

func (s *CreateRaceWorkForInnerRequest) SetGroupId(v string) *CreateRaceWorkForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *CreateRaceWorkForInnerRequest) SetKeywords(v string) *CreateRaceWorkForInnerRequest {
	s.Keywords = &v
	return s
}

func (s *CreateRaceWorkForInnerRequest) SetLogoUrl(v string) *CreateRaceWorkForInnerRequest {
	s.LogoUrl = &v
	return s
}

func (s *CreateRaceWorkForInnerRequest) SetSecurityToken(v string) *CreateRaceWorkForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateRaceWorkForInnerRequest) SetShortDescription(v string) *CreateRaceWorkForInnerRequest {
	s.ShortDescription = &v
	return s
}

func (s *CreateRaceWorkForInnerRequest) SetWorkName(v string) *CreateRaceWorkForInnerRequest {
	s.WorkName = &v
	return s
}

type CreateRaceWorkForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRaceWorkForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRaceWorkForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRaceWorkForInnerResponseBody) SetRequestId(v string) *CreateRaceWorkForInnerResponseBody {
	s.RequestId = &v
	return s
}

type CreateRaceWorkForInnerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRaceWorkForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRaceWorkForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRaceWorkForInnerResponse) GoString() string {
	return s.String()
}

func (s *CreateRaceWorkForInnerResponse) SetHeaders(v map[string]*string) *CreateRaceWorkForInnerResponse {
	s.Headers = v
	return s
}

func (s *CreateRaceWorkForInnerResponse) SetStatusCode(v int32) *CreateRaceWorkForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRaceWorkForInnerResponse) SetBody(v *CreateRaceWorkForInnerResponseBody) *CreateRaceWorkForInnerResponse {
	s.Body = v
	return s
}

type CreateSecretKeyRequest struct {
	SecretKey     *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	SecretKeyName *string `json:"SecretKeyName,omitempty" xml:"SecretKeyName,omitempty"`
	SecretValue   *string `json:"SecretValue,omitempty" xml:"SecretValue,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateSecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretKeyRequest) SetSecretKey(v string) *CreateSecretKeyRequest {
	s.SecretKey = &v
	return s
}

func (s *CreateSecretKeyRequest) SetSecretKeyName(v string) *CreateSecretKeyRequest {
	s.SecretKeyName = &v
	return s
}

func (s *CreateSecretKeyRequest) SetSecretValue(v string) *CreateSecretKeyRequest {
	s.SecretValue = &v
	return s
}

func (s *CreateSecretKeyRequest) SetSecurityToken(v string) *CreateSecretKeyRequest {
	s.SecurityToken = &v
	return s
}

type CreateSecretKeyResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretKeyId   *string `json:"SecretKeyId,omitempty" xml:"SecretKeyId,omitempty"`
	SecretKeyName *string `json:"SecretKeyName,omitempty" xml:"SecretKeyName,omitempty"`
}

func (s CreateSecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretKeyResponseBody) SetRequestId(v string) *CreateSecretKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecretKeyResponseBody) SetSecretKeyId(v string) *CreateSecretKeyResponseBody {
	s.SecretKeyId = &v
	return s
}

func (s *CreateSecretKeyResponseBody) SetSecretKeyName(v string) *CreateSecretKeyResponseBody {
	s.SecretKeyName = &v
	return s
}

type CreateSecretKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretKeyResponse) SetHeaders(v map[string]*string) *CreateSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretKeyResponse) SetStatusCode(v int32) *CreateSecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretKeyResponse) SetBody(v *CreateSecretKeyResponseBody) *CreateSecretKeyResponse {
	s.Body = v
	return s
}

type CreateTrafficControlRequest struct {
	ApiDefault         *int32  `json:"ApiDefault,omitempty" xml:"ApiDefault,omitempty"`
	AppDefault         *int32  `json:"AppDefault,omitempty" xml:"AppDefault,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
	TrafficControlUnit *string `json:"TrafficControlUnit,omitempty" xml:"TrafficControlUnit,omitempty"`
	UserDefault        *int32  `json:"UserDefault,omitempty" xml:"UserDefault,omitempty"`
}

func (s CreateTrafficControlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlRequest) SetApiDefault(v int32) *CreateTrafficControlRequest {
	s.ApiDefault = &v
	return s
}

func (s *CreateTrafficControlRequest) SetAppDefault(v int32) *CreateTrafficControlRequest {
	s.AppDefault = &v
	return s
}

func (s *CreateTrafficControlRequest) SetDescription(v string) *CreateTrafficControlRequest {
	s.Description = &v
	return s
}

func (s *CreateTrafficControlRequest) SetSecurityToken(v string) *CreateTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateTrafficControlRequest) SetTrafficControlName(v string) *CreateTrafficControlRequest {
	s.TrafficControlName = &v
	return s
}

func (s *CreateTrafficControlRequest) SetTrafficControlUnit(v string) *CreateTrafficControlRequest {
	s.TrafficControlUnit = &v
	return s
}

func (s *CreateTrafficControlRequest) SetUserDefault(v int32) *CreateTrafficControlRequest {
	s.UserDefault = &v
	return s
}

type CreateTrafficControlResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s CreateTrafficControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlResponseBody) SetRequestId(v string) *CreateTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficControlResponseBody) SetTrafficControlId(v string) *CreateTrafficControlResponseBody {
	s.TrafficControlId = &v
	return s
}

type CreateTrafficControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficControlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlResponse) SetHeaders(v map[string]*string) *CreateTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficControlResponse) SetStatusCode(v int32) *CreateTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficControlResponse) SetBody(v *CreateTrafficControlResponseBody) *CreateTrafficControlResponse {
	s.Body = v
	return s
}

type CreateUserWhiteListRequest struct {
	AoneId        *string `json:"AoneId,omitempty" xml:"AoneId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	LimitCount    *int32  `json:"LimitCount,omitempty" xml:"LimitCount,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid           *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateUserWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserWhiteListRequest) GoString() string {
	return s.String()
}

func (s *CreateUserWhiteListRequest) SetAoneId(v string) *CreateUserWhiteListRequest {
	s.AoneId = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetDescription(v string) *CreateUserWhiteListRequest {
	s.Description = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetEntityId(v string) *CreateUserWhiteListRequest {
	s.EntityId = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetLimitCount(v int32) *CreateUserWhiteListRequest {
	s.LimitCount = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetSecurityToken(v string) *CreateUserWhiteListRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetType(v string) *CreateUserWhiteListRequest {
	s.Type = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetUid(v int64) *CreateUserWhiteListRequest {
	s.Uid = &v
	return s
}

func (s *CreateUserWhiteListRequest) SetValue(v string) *CreateUserWhiteListRequest {
	s.Value = &v
	return s
}

type CreateUserWhiteListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserWhiteListResponseBody) SetRequestId(v string) *CreateUserWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type CreateUserWhiteListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserWhiteListResponse) GoString() string {
	return s.String()
}

func (s *CreateUserWhiteListResponse) SetHeaders(v map[string]*string) *CreateUserWhiteListResponse {
	s.Headers = v
	return s
}

func (s *CreateUserWhiteListResponse) SetStatusCode(v int32) *CreateUserWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserWhiteListResponse) SetBody(v *CreateUserWhiteListResponseBody) *CreateUserWhiteListResponse {
	s.Body = v
	return s
}

type DeleteAllTrafficSpecialControlRequest struct {
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DeleteAllTrafficSpecialControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllTrafficSpecialControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteAllTrafficSpecialControlRequest) SetSecurityToken(v string) *DeleteAllTrafficSpecialControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAllTrafficSpecialControlRequest) SetTrafficControlId(v string) *DeleteAllTrafficSpecialControlRequest {
	s.TrafficControlId = &v
	return s
}

type DeleteAllTrafficSpecialControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAllTrafficSpecialControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllTrafficSpecialControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllTrafficSpecialControlResponseBody) SetRequestId(v string) *DeleteAllTrafficSpecialControlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAllTrafficSpecialControlResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAllTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAllTrafficSpecialControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAllTrafficSpecialControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteAllTrafficSpecialControlResponse) SetHeaders(v map[string]*string) *DeleteAllTrafficSpecialControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteAllTrafficSpecialControlResponse) SetStatusCode(v int32) *DeleteAllTrafficSpecialControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAllTrafficSpecialControlResponse) SetBody(v *DeleteAllTrafficSpecialControlResponseBody) *DeleteAllTrafficSpecialControlResponse {
	s.Body = v
	return s
}

type DeleteApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiRequest) SetApiId(v string) *DeleteApiRequest {
	s.ApiId = &v
	return s
}

func (s *DeleteApiRequest) SetGroupId(v string) *DeleteApiRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteApiRequest) SetSecurityToken(v string) *DeleteApiRequest {
	s.SecurityToken = &v
	return s
}

type DeleteApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiResponseBody) SetRequestId(v string) *DeleteApiResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiResponse) SetHeaders(v map[string]*string) *DeleteApiResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiResponse) SetStatusCode(v int32) *DeleteApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiResponse) SetBody(v *DeleteApiResponseBody) *DeleteApiResponse {
	s.Body = v
	return s
}

type DeleteApiForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteApiForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiForInnerRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiForInnerRequest) SetAliUid(v int64) *DeleteApiForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *DeleteApiForInnerRequest) SetApiId(v string) *DeleteApiForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *DeleteApiForInnerRequest) SetGroupId(v string) *DeleteApiForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteApiForInnerRequest) SetSecurityToken(v string) *DeleteApiForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DeleteApiForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiForInnerResponseBody) SetRequestId(v string) *DeleteApiForInnerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApiForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiForInnerResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiForInnerResponse) SetHeaders(v map[string]*string) *DeleteApiForInnerResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiForInnerResponse) SetStatusCode(v int32) *DeleteApiForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiForInnerResponse) SetBody(v *DeleteApiForInnerResponseBody) *DeleteApiForInnerResponse {
	s.Body = v
	return s
}

type DeleteApiGroupRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiGroupRequest) SetGroupId(v string) *DeleteApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteApiGroupRequest) SetSecurityToken(v string) *DeleteApiGroupRequest {
	s.SecurityToken = &v
	return s
}

type DeleteApiGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiGroupResponseBody) SetRequestId(v string) *DeleteApiGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApiGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApiGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiGroupResponse) SetHeaders(v map[string]*string) *DeleteApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiGroupResponse) SetStatusCode(v int32) *DeleteApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiGroupResponse) SetBody(v *DeleteApiGroupResponseBody) *DeleteApiGroupResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) SetAppId(v int64) *DeleteAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRequest) SetSecurityToken(v string) *DeleteAppRequest {
	s.SecurityToken = &v
	return s
}

type DeleteAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppResponse) SetStatusCode(v int32) *DeleteAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppResponse) SetBody(v *DeleteAppResponseBody) *DeleteAppResponse {
	s.Body = v
	return s
}

type DeleteAppForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteAppForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppForInnerRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppForInnerRequest) SetAliUid(v int64) *DeleteAppForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *DeleteAppForInnerRequest) SetAppId(v int64) *DeleteAppForInnerRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppForInnerRequest) SetSecurityToken(v string) *DeleteAppForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DeleteAppForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppForInnerResponseBody) SetRequestId(v string) *DeleteAppForInnerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppForInnerResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppForInnerResponse) SetHeaders(v map[string]*string) *DeleteAppForInnerResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppForInnerResponse) SetStatusCode(v int32) *DeleteAppForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppForInnerResponse) SetBody(v *DeleteAppForInnerResponseBody) *DeleteAppForInnerResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainRequest) SetGroupId(v string) *DeleteDomainRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDomainRequest) SetSecurityToken(v string) *DeleteDomainRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponseBody) SetRequestId(v string) *DeleteDomainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainResponse) SetBody(v *DeleteDomainResponseBody) *DeleteDomainResponse {
	s.Body = v
	return s
}

type DeleteDomainCertificateRequest struct {
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainCertificateRequest) SetCertificateId(v string) *DeleteDomainCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *DeleteDomainCertificateRequest) SetDomainName(v string) *DeleteDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainCertificateRequest) SetGroupId(v string) *DeleteDomainCertificateRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteDomainCertificateRequest) SetSecurityToken(v string) *DeleteDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

type DeleteDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDomainCertificateResponseBody) SetRequestId(v string) *DeleteDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDomainCertificateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainCertificateResponse) SetHeaders(v map[string]*string) *DeleteDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainCertificateResponse) SetStatusCode(v int32) *DeleteDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDomainCertificateResponse) SetBody(v *DeleteDomainCertificateResponseBody) *DeleteDomainCertificateResponse {
	s.Body = v
	return s
}

type DeleteSecretKeyRequest struct {
	SecretKeyId   *string `json:"SecretKeyId,omitempty" xml:"SecretKeyId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteSecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretKeyRequest) SetSecretKeyId(v string) *DeleteSecretKeyRequest {
	s.SecretKeyId = &v
	return s
}

func (s *DeleteSecretKeyRequest) SetSecurityToken(v string) *DeleteSecretKeyRequest {
	s.SecurityToken = &v
	return s
}

type DeleteSecretKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretKeyResponseBody) SetRequestId(v string) *DeleteSecretKeyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecretKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretKeyResponse) SetHeaders(v map[string]*string) *DeleteSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretKeyResponse) SetStatusCode(v int32) *DeleteSecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretKeyResponse) SetBody(v *DeleteSecretKeyResponseBody) *DeleteSecretKeyResponse {
	s.Body = v
	return s
}

type DeleteTrafficControlRequest struct {
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DeleteTrafficControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlRequest) SetSecurityToken(v string) *DeleteTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteTrafficControlRequest) SetTrafficControlId(v string) *DeleteTrafficControlRequest {
	s.TrafficControlId = &v
	return s
}

type DeleteTrafficControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlResponseBody) SetRequestId(v string) *DeleteTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrafficControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlResponse) SetHeaders(v map[string]*string) *DeleteTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficControlResponse) SetStatusCode(v int32) *DeleteTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficControlResponse) SetBody(v *DeleteTrafficControlResponseBody) *DeleteTrafficControlResponse {
	s.Body = v
	return s
}

type DeleteTrafficSpecialControlRequest struct {
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SpecialKey       *string `json:"SpecialKey,omitempty" xml:"SpecialKey,omitempty"`
	SpecialType      *string `json:"SpecialType,omitempty" xml:"SpecialType,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DeleteTrafficSpecialControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficSpecialControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficSpecialControlRequest) SetSecurityToken(v string) *DeleteTrafficSpecialControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) SetSpecialKey(v string) *DeleteTrafficSpecialControlRequest {
	s.SpecialKey = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) SetSpecialType(v string) *DeleteTrafficSpecialControlRequest {
	s.SpecialType = &v
	return s
}

func (s *DeleteTrafficSpecialControlRequest) SetTrafficControlId(v string) *DeleteTrafficSpecialControlRequest {
	s.TrafficControlId = &v
	return s
}

type DeleteTrafficSpecialControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTrafficSpecialControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficSpecialControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTrafficSpecialControlResponseBody) SetRequestId(v string) *DeleteTrafficSpecialControlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTrafficSpecialControlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficSpecialControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTrafficSpecialControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficSpecialControlResponse) SetHeaders(v map[string]*string) *DeleteTrafficSpecialControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficSpecialControlResponse) SetStatusCode(v int32) *DeleteTrafficSpecialControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficSpecialControlResponse) SetBody(v *DeleteTrafficSpecialControlResponseBody) *DeleteTrafficSpecialControlResponse {
	s.Body = v
	return s
}

type DeleteUserWhiteListByTypeRequest struct {
	EntityId      *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid           *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DeleteUserWhiteListByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserWhiteListByTypeRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserWhiteListByTypeRequest) SetEntityId(v string) *DeleteUserWhiteListByTypeRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteUserWhiteListByTypeRequest) SetSecurityToken(v string) *DeleteUserWhiteListByTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteUserWhiteListByTypeRequest) SetType(v string) *DeleteUserWhiteListByTypeRequest {
	s.Type = &v
	return s
}

func (s *DeleteUserWhiteListByTypeRequest) SetUid(v int64) *DeleteUserWhiteListByTypeRequest {
	s.Uid = &v
	return s
}

type DeleteUserWhiteListByTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserWhiteListByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserWhiteListByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserWhiteListByTypeResponseBody) SetRequestId(v string) *DeleteUserWhiteListByTypeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserWhiteListByTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserWhiteListByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserWhiteListByTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserWhiteListByTypeResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserWhiteListByTypeResponse) SetHeaders(v map[string]*string) *DeleteUserWhiteListByTypeResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserWhiteListByTypeResponse) SetStatusCode(v int32) *DeleteUserWhiteListByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserWhiteListByTypeResponse) SetBody(v *DeleteUserWhiteListByTypeResponseBody) *DeleteUserWhiteListByTypeResponse {
	s.Body = v
	return s
}

type DeployApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DeployApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployApiRequest) GoString() string {
	return s.String()
}

func (s *DeployApiRequest) SetApiId(v string) *DeployApiRequest {
	s.ApiId = &v
	return s
}

func (s *DeployApiRequest) SetDescription(v string) *DeployApiRequest {
	s.Description = &v
	return s
}

func (s *DeployApiRequest) SetGroupId(v string) *DeployApiRequest {
	s.GroupId = &v
	return s
}

func (s *DeployApiRequest) SetSecurityToken(v string) *DeployApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployApiRequest) SetStageName(v string) *DeployApiRequest {
	s.StageName = &v
	return s
}

type DeployApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApiResponseBody) SetRequestId(v string) *DeployApiResponseBody {
	s.RequestId = &v
	return s
}

type DeployApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployApiResponse) GoString() string {
	return s.String()
}

func (s *DeployApiResponse) SetHeaders(v map[string]*string) *DeployApiResponse {
	s.Headers = v
	return s
}

func (s *DeployApiResponse) SetStatusCode(v int32) *DeployApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApiResponse) SetBody(v *DeployApiResponseBody) *DeployApiResponse {
	s.Body = v
	return s
}

type DeployApiForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DeployApiForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployApiForInnerRequest) GoString() string {
	return s.String()
}

func (s *DeployApiForInnerRequest) SetAliUid(v int64) *DeployApiForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *DeployApiForInnerRequest) SetApiId(v string) *DeployApiForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *DeployApiForInnerRequest) SetDescription(v string) *DeployApiForInnerRequest {
	s.Description = &v
	return s
}

func (s *DeployApiForInnerRequest) SetGroupId(v string) *DeployApiForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *DeployApiForInnerRequest) SetSecurityToken(v string) *DeployApiForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeployApiForInnerRequest) SetStageName(v string) *DeployApiForInnerRequest {
	s.StageName = &v
	return s
}

type DeployApiForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApiForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployApiForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApiForInnerResponseBody) SetRequestId(v string) *DeployApiForInnerResponseBody {
	s.RequestId = &v
	return s
}

type DeployApiForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApiForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApiForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployApiForInnerResponse) GoString() string {
	return s.String()
}

func (s *DeployApiForInnerResponse) SetHeaders(v map[string]*string) *DeployApiForInnerResponse {
	s.Headers = v
	return s
}

func (s *DeployApiForInnerResponse) SetStatusCode(v int32) *DeployApiForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApiForInnerResponse) SetBody(v *DeployApiForInnerResponseBody) *DeployApiForInnerResponse {
	s.Body = v
	return s
}

type DescribeApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiRequest) SetApiId(v string) *DescribeApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiRequest) SetGroupId(v string) *DescribeApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiRequest) SetSecurityToken(v string) *DescribeApiRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApiResponseBody struct {
	ApiId                   *string                                         `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                 *string                                         `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType                *string                                         `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ConstantParameters      *DescribeApiResponseBodyConstantParameters      `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	CreatedTime             *string                                         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CustomSystemParameters  *DescribeApiResponseBodyCustomSystemParameters  `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	DeployedInfos           *DescribeApiResponseBodyDeployedInfos           `json:"DeployedInfos,omitempty" xml:"DeployedInfos,omitempty" type:"Struct"`
	Description             *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples        *DescribeApiResponseBodyErrorCodeSamples        `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample        *string                                         `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	GroupId                 *string                                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName               *string                                         `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Mock                    *string                                         `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult              *string                                         `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	ModifiedTime            *string                                         `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OpenIdConnectConfig     *DescribeApiResponseBodyOpenIdConnectConfig     `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	OriginResultDescription *string                                         `json:"OriginResultDescription,omitempty" xml:"OriginResultDescription,omitempty"`
	ParametersMapObject     *DescribeApiResponseBodyParametersMapObject     `json:"ParametersMapObject,omitempty" xml:"ParametersMapObject,omitempty" type:"Struct"`
	RegionId                *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestConfig           *DescribeApiResponseBodyRequestConfig           `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	RequestId               *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestParametersObject *DescribeApiResponseBodyRequestParametersObject `json:"RequestParametersObject,omitempty" xml:"RequestParametersObject,omitempty" type:"Struct"`
	ResultSample            *string                                         `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType              *string                                         `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceConfig           *DescribeApiResponseBodyServiceConfig           `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	ServiceParametersObject *DescribeApiResponseBodyServiceParametersObject `json:"ServiceParametersObject,omitempty" xml:"ServiceParametersObject,omitempty" type:"Struct"`
	SystemParameters        *DescribeApiResponseBodySystemParameters        `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility              *string                                         `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBody) SetApiId(v string) *DescribeApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeApiResponseBody) SetApiName(v string) *DescribeApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribeApiResponseBody) SetAuthType(v string) *DescribeApiResponseBody {
	s.AuthType = &v
	return s
}

func (s *DescribeApiResponseBody) SetConstantParameters(v *DescribeApiResponseBodyConstantParameters) *DescribeApiResponseBody {
	s.ConstantParameters = v
	return s
}

func (s *DescribeApiResponseBody) SetCreatedTime(v string) *DescribeApiResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiResponseBody) SetCustomSystemParameters(v *DescribeApiResponseBodyCustomSystemParameters) *DescribeApiResponseBody {
	s.CustomSystemParameters = v
	return s
}

func (s *DescribeApiResponseBody) SetDeployedInfos(v *DescribeApiResponseBodyDeployedInfos) *DescribeApiResponseBody {
	s.DeployedInfos = v
	return s
}

func (s *DescribeApiResponseBody) SetDescription(v string) *DescribeApiResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBody) SetErrorCodeSamples(v *DescribeApiResponseBodyErrorCodeSamples) *DescribeApiResponseBody {
	s.ErrorCodeSamples = v
	return s
}

func (s *DescribeApiResponseBody) SetFailResultSample(v string) *DescribeApiResponseBody {
	s.FailResultSample = &v
	return s
}

func (s *DescribeApiResponseBody) SetGroupId(v string) *DescribeApiResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiResponseBody) SetGroupName(v string) *DescribeApiResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiResponseBody) SetMock(v string) *DescribeApiResponseBody {
	s.Mock = &v
	return s
}

func (s *DescribeApiResponseBody) SetMockResult(v string) *DescribeApiResponseBody {
	s.MockResult = &v
	return s
}

func (s *DescribeApiResponseBody) SetModifiedTime(v string) *DescribeApiResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiResponseBody) SetOpenIdConnectConfig(v *DescribeApiResponseBodyOpenIdConnectConfig) *DescribeApiResponseBody {
	s.OpenIdConnectConfig = v
	return s
}

func (s *DescribeApiResponseBody) SetOriginResultDescription(v string) *DescribeApiResponseBody {
	s.OriginResultDescription = &v
	return s
}

func (s *DescribeApiResponseBody) SetParametersMapObject(v *DescribeApiResponseBodyParametersMapObject) *DescribeApiResponseBody {
	s.ParametersMapObject = v
	return s
}

func (s *DescribeApiResponseBody) SetRegionId(v string) *DescribeApiResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiResponseBody) SetRequestConfig(v *DescribeApiResponseBodyRequestConfig) *DescribeApiResponseBody {
	s.RequestConfig = v
	return s
}

func (s *DescribeApiResponseBody) SetRequestId(v string) *DescribeApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiResponseBody) SetRequestParametersObject(v *DescribeApiResponseBodyRequestParametersObject) *DescribeApiResponseBody {
	s.RequestParametersObject = v
	return s
}

func (s *DescribeApiResponseBody) SetResultSample(v string) *DescribeApiResponseBody {
	s.ResultSample = &v
	return s
}

func (s *DescribeApiResponseBody) SetResultType(v string) *DescribeApiResponseBody {
	s.ResultType = &v
	return s
}

func (s *DescribeApiResponseBody) SetServiceConfig(v *DescribeApiResponseBodyServiceConfig) *DescribeApiResponseBody {
	s.ServiceConfig = v
	return s
}

func (s *DescribeApiResponseBody) SetServiceParametersObject(v *DescribeApiResponseBodyServiceParametersObject) *DescribeApiResponseBody {
	s.ServiceParametersObject = v
	return s
}

func (s *DescribeApiResponseBody) SetSystemParameters(v *DescribeApiResponseBodySystemParameters) *DescribeApiResponseBody {
	s.SystemParameters = v
	return s
}

func (s *DescribeApiResponseBody) SetVisibility(v string) *DescribeApiResponseBody {
	s.Visibility = &v
	return s
}

type DescribeApiResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeApiResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyConstantParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyConstantParameters) SetConstantParameter(v []*DescribeApiResponseBodyConstantParametersConstantParameter) *DescribeApiResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

type DescribeApiResponseBodyConstantParametersConstantParameter struct {
	ConstantValue        *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodyConstantParametersConstantParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) SetConstantValue(v string) *DescribeApiResponseBodyConstantParametersConstantParameter {
	s.ConstantValue = &v
	return s
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) SetDescription(v string) *DescribeApiResponseBodyConstantParametersConstantParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) SetLocation(v string) *DescribeApiResponseBodyConstantParametersConstantParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) SetServiceParameterName(v string) *DescribeApiResponseBodyConstantParametersConstantParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyCustomSystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeApiResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

type DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) SetDemoValue(v string) *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) SetDescription(v string) *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) SetLocation(v string) *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) SetParameterName(v string) *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) SetServiceParameterName(v string) *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiResponseBodyDeployedInfos struct {
	DeployedInfo []*DescribeApiResponseBodyDeployedInfosDeployedInfo `json:"DeployedInfo,omitempty" xml:"DeployedInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyDeployedInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyDeployedInfos) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyDeployedInfos) SetDeployedInfo(v []*DescribeApiResponseBodyDeployedInfosDeployedInfo) *DescribeApiResponseBodyDeployedInfos {
	s.DeployedInfo = v
	return s
}

type DescribeApiResponseBodyDeployedInfosDeployedInfo struct {
	DeployedStatus   *string `json:"DeployedStatus,omitempty" xml:"DeployedStatus,omitempty"`
	EffectiveVersion *string `json:"EffectiveVersion,omitempty" xml:"EffectiveVersion,omitempty"`
	StageName        *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiResponseBodyDeployedInfosDeployedInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyDeployedInfosDeployedInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) SetDeployedStatus(v string) *DescribeApiResponseBodyDeployedInfosDeployedInfo {
	s.DeployedStatus = &v
	return s
}

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) SetEffectiveVersion(v string) *DescribeApiResponseBodyDeployedInfosDeployedInfo {
	s.EffectiveVersion = &v
	return s
}

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) SetStageName(v string) *DescribeApiResponseBodyDeployedInfosDeployedInfo {
	s.StageName = &v
	return s
}

type DescribeApiResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyErrorCodeSamples) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeApiResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

type DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Model       *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) SetCode(v string) *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Code = &v
	return s
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) SetDescription(v string) *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) SetMessage(v string) *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Message = &v
	return s
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) SetModel(v string) *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Model = &v
	return s
}

type DescribeApiResponseBodyOpenIdConnectConfig struct {
	IdTokenParamName *string `json:"IdTokenParamName,omitempty" xml:"IdTokenParamName,omitempty"`
	OpenIdApiType    *string `json:"OpenIdApiType,omitempty" xml:"OpenIdApiType,omitempty"`
	PublicKey        *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	PublicKeyId      *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
}

func (s DescribeApiResponseBodyOpenIdConnectConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) SetIdTokenParamName(v string) *DescribeApiResponseBodyOpenIdConnectConfig {
	s.IdTokenParamName = &v
	return s
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) SetOpenIdApiType(v string) *DescribeApiResponseBodyOpenIdConnectConfig {
	s.OpenIdApiType = &v
	return s
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) SetPublicKey(v string) *DescribeApiResponseBodyOpenIdConnectConfig {
	s.PublicKey = &v
	return s
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) SetPublicKeyId(v string) *DescribeApiResponseBodyOpenIdConnectConfig {
	s.PublicKeyId = &v
	return s
}

type DescribeApiResponseBodyParametersMapObject struct {
	ServiceParameterMap []*DescribeApiResponseBodyParametersMapObjectServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyParametersMapObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyParametersMapObject) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyParametersMapObject) SetServiceParameterMap(v []*DescribeApiResponseBodyParametersMapObjectServiceParameterMap) *DescribeApiResponseBodyParametersMapObject {
	s.ServiceParameterMap = v
	return s
}

type DescribeApiResponseBodyParametersMapObjectServiceParameterMap struct {
	RequestParameterName *string `json:"RequestParameterName,omitempty" xml:"RequestParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodyParametersMapObjectServiceParameterMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyParametersMapObjectServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyParametersMapObjectServiceParameterMap) SetRequestParameterName(v string) *DescribeApiResponseBodyParametersMapObjectServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyParametersMapObjectServiceParameterMap) SetServiceParameterName(v string) *DescribeApiResponseBodyParametersMapObjectServiceParameterMap {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiResponseBodyRequestConfig struct {
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	RequestHttpMethod   *string `json:"RequestHttpMethod,omitempty" xml:"RequestHttpMethod,omitempty"`
	RequestPath         *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	RequestProtocol     *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeApiResponseBodyRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeApiResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiResponseBodyRequestConfig) SetPostBodyDescription(v string) *DescribeApiResponseBodyRequestConfig {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeApiResponseBodyRequestConfig) SetRequestHttpMethod(v string) *DescribeApiResponseBodyRequestConfig {
	s.RequestHttpMethod = &v
	return s
}

func (s *DescribeApiResponseBodyRequestConfig) SetRequestPath(v string) *DescribeApiResponseBodyRequestConfig {
	s.RequestPath = &v
	return s
}

func (s *DescribeApiResponseBodyRequestConfig) SetRequestProtocol(v string) *DescribeApiResponseBodyRequestConfig {
	s.RequestProtocol = &v
	return s
}

type DescribeApiResponseBodyRequestParametersObject struct {
	RequestParam []*DescribeApiResponseBodyRequestParametersObjectRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyRequestParametersObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyRequestParametersObject) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestParametersObject) SetRequestParam(v []*DescribeApiResponseBodyRequestParametersObjectRequestParam) *DescribeApiResponseBodyRequestParametersObject {
	s.RequestParam = v
	return s
}

type DescribeApiResponseBodyRequestParametersObjectRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow           *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	Location          *string `json:"Location,omitempty" xml:"Location,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeApiResponseBodyRequestParametersObjectRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyRequestParametersObjectRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetApiParameterName(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetDefaultValue(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetDemoValue(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetDescription(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetDocOrder(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetDocShow(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetEnumValue(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetJsonScheme(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetLocation(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetMaxLength(v int64) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetMaxValue(v int64) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetMinLength(v int64) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetMinValue(v int64) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetParameterType(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetRegularExpression(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersObjectRequestParam) SetRequired(v string) *DescribeApiResponseBodyRequestParametersObjectRequestParam {
	s.Required = &v
	return s
}

type DescribeApiResponseBodyServiceConfig struct {
	ContentTypeCatagory   *string                                                    `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	ContentTypeValue      *string                                                    `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	FunctionComputeConfig *DescribeApiResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	Mock                  *string                                                    `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult            *string                                                    `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	ServiceAddress        *string                                                    `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceHttpMethod     *string                                                    `json:"ServiceHttpMethod,omitempty" xml:"ServiceHttpMethod,omitempty"`
	ServicePath           *string                                                    `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	ServiceProtocol       *string                                                    `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout        *string                                                    `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable      *string                                                    `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	VpcConfig             *DescribeApiResponseBodyServiceConfigVpcConfig             `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
}

func (s DescribeApiResponseBodyServiceConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfig) SetContentTypeCatagory(v string) *DescribeApiResponseBodyServiceConfig {
	s.ContentTypeCatagory = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetContentTypeValue(v string) *DescribeApiResponseBodyServiceConfig {
	s.ContentTypeValue = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetFunctionComputeConfig(v *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) *DescribeApiResponseBodyServiceConfig {
	s.FunctionComputeConfig = v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetMock(v string) *DescribeApiResponseBodyServiceConfig {
	s.Mock = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetMockResult(v string) *DescribeApiResponseBodyServiceConfig {
	s.MockResult = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetServiceAddress(v string) *DescribeApiResponseBodyServiceConfig {
	s.ServiceAddress = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetServiceHttpMethod(v string) *DescribeApiResponseBodyServiceConfig {
	s.ServiceHttpMethod = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetServicePath(v string) *DescribeApiResponseBodyServiceConfig {
	s.ServicePath = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetServiceProtocol(v string) *DescribeApiResponseBodyServiceConfig {
	s.ServiceProtocol = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetServiceTimeout(v string) *DescribeApiResponseBodyServiceConfig {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetServiceVpcEnable(v string) *DescribeApiResponseBodyServiceConfig {
	s.ServiceVpcEnable = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetVpcConfig(v *DescribeApiResponseBodyServiceConfigVpcConfig) *DescribeApiResponseBodyServiceConfig {
	s.VpcConfig = v
	return s
}

type DescribeApiResponseBodyServiceConfigFunctionComputeConfig struct {
	FcRegionId   *string `json:"FcRegionId,omitempty" xml:"FcRegionId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	RoleArn      *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetFcRegionId(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FcRegionId = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetFunctionName(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetRoleArn(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetServiceName(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.ServiceName = &v
	return s
}

type DescribeApiResponseBodyServiceConfigVpcConfig struct {
	Id         *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) SetId(v string) *DescribeApiResponseBodyServiceConfigVpcConfig {
	s.Id = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) SetInstanceId(v string) *DescribeApiResponseBodyServiceConfigVpcConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) SetName(v string) *DescribeApiResponseBodyServiceConfigVpcConfig {
	s.Name = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) SetPort(v int32) *DescribeApiResponseBodyServiceConfigVpcConfig {
	s.Port = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) SetVpcId(v string) *DescribeApiResponseBodyServiceConfigVpcConfig {
	s.VpcId = &v
	return s
}

type DescribeApiResponseBodyServiceParametersObject struct {
	ServiceParam []*DescribeApiResponseBodyServiceParametersObjectServiceParam `json:"ServiceParam,omitempty" xml:"ServiceParam,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceParametersObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersObject) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersObject) SetServiceParam(v []*DescribeApiResponseBodyServiceParametersObjectServiceParam) *DescribeApiResponseBodyServiceParametersObject {
	s.ServiceParam = v
	return s
}

type DescribeApiResponseBodyServiceParametersObjectServiceParam struct {
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
	Type                 *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApiResponseBodyServiceParametersObjectServiceParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersObjectServiceParam) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersObjectServiceParam) SetLocation(v string) *DescribeApiResponseBodyServiceParametersObjectServiceParam {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersObjectServiceParam) SetServiceParameterName(v string) *DescribeApiResponseBodyServiceParametersObjectServiceParam {
	s.ServiceParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersObjectServiceParam) SetType(v string) *DescribeApiResponseBodyServiceParametersObjectServiceParam {
	s.Type = &v
	return s
}

type DescribeApiResponseBodySystemParameters struct {
	SystemParameter []*DescribeApiResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodySystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodySystemParameters) SetSystemParameter(v []*DescribeApiResponseBodySystemParametersSystemParameter) *DescribeApiResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

type DescribeApiResponseBodySystemParametersSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodySystemParametersSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) SetDemoValue(v string) *DescribeApiResponseBodySystemParametersSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) SetDescription(v string) *DescribeApiResponseBodySystemParametersSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) SetLocation(v string) *DescribeApiResponseBodySystemParametersSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) SetParameterName(v string) *DescribeApiResponseBodySystemParametersSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) SetServiceParameterName(v string) *DescribeApiResponseBodySystemParametersSystemParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiResponse) SetHeaders(v map[string]*string) *DescribeApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiResponse) SetStatusCode(v int32) *DescribeApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiResponse) SetBody(v *DescribeApiResponseBody) *DescribeApiResponse {
	s.Body = v
	return s
}

type DescribeApiDocRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiDocRequest) SetApiId(v string) *DescribeApiDocRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocRequest) SetGroupId(v string) *DescribeApiDocRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocRequest) SetSecurityToken(v string) *DescribeApiDocRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiDocRequest) SetStageName(v string) *DescribeApiDocRequest {
	s.StageName = &v
	return s
}

type DescribeApiDocResponseBody struct {
	ApiId               *string                                       `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string                                       `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BodyFormat          *string                                       `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	DeployedTime        *string                                       `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description         *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples    *DescribeApiDocResponseBodyErrorCodeSamples   `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample    *string                                       `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	GroupId             *string                                       `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpMethod          *string                                       `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol        *string                                       `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Path                *string                                       `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters      *DescribeApiDocResponseBodyPathParameters     `json:"PathParameters,omitempty" xml:"PathParameters,omitempty" type:"Struct"`
	PostBodyDescription *string                                       `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType        *string                                       `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RegionId            *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestBody         *DescribeApiDocResponseBodyRequestBody        `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	RequestHeaders      *DescribeApiDocResponseBodyRequestHeaders     `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Struct"`
	RequestId           *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestQueries      *DescribeApiDocResponseBodyRequestQueries     `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty" type:"Struct"`
	ResultDescriptions  *DescribeApiDocResponseBodyResultDescriptions `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	ResultSample        *string                                       `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType          *string                                       `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SdkDemos            *DescribeApiDocResponseBodySdkDemos           `json:"SdkDemos,omitempty" xml:"SdkDemos,omitempty" type:"Struct"`
	ServiceTimeout      *int32                                        `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	StageName           *string                                       `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBody) SetApiId(v string) *DescribeApiDocResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetApiName(v string) *DescribeApiDocResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetBodyFormat(v string) *DescribeApiDocResponseBody {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetDeployedTime(v string) *DescribeApiDocResponseBody {
	s.DeployedTime = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetDescription(v string) *DescribeApiDocResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetErrorCodeSamples(v *DescribeApiDocResponseBodyErrorCodeSamples) *DescribeApiDocResponseBody {
	s.ErrorCodeSamples = v
	return s
}

func (s *DescribeApiDocResponseBody) SetFailResultSample(v string) *DescribeApiDocResponseBody {
	s.FailResultSample = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetGroupId(v string) *DescribeApiDocResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetGroupName(v string) *DescribeApiDocResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetHttpMethod(v string) *DescribeApiDocResponseBody {
	s.HttpMethod = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetHttpProtocol(v string) *DescribeApiDocResponseBody {
	s.HttpProtocol = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetPath(v string) *DescribeApiDocResponseBody {
	s.Path = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetPathParameters(v *DescribeApiDocResponseBodyPathParameters) *DescribeApiDocResponseBody {
	s.PathParameters = v
	return s
}

func (s *DescribeApiDocResponseBody) SetPostBodyDescription(v string) *DescribeApiDocResponseBody {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetPostBodyType(v string) *DescribeApiDocResponseBody {
	s.PostBodyType = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetRegionId(v string) *DescribeApiDocResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestBody(v *DescribeApiDocResponseBodyRequestBody) *DescribeApiDocResponseBody {
	s.RequestBody = v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestHeaders(v *DescribeApiDocResponseBodyRequestHeaders) *DescribeApiDocResponseBody {
	s.RequestHeaders = v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestId(v string) *DescribeApiDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestQueries(v *DescribeApiDocResponseBodyRequestQueries) *DescribeApiDocResponseBody {
	s.RequestQueries = v
	return s
}

func (s *DescribeApiDocResponseBody) SetResultDescriptions(v *DescribeApiDocResponseBodyResultDescriptions) *DescribeApiDocResponseBody {
	s.ResultDescriptions = v
	return s
}

func (s *DescribeApiDocResponseBody) SetResultSample(v string) *DescribeApiDocResponseBody {
	s.ResultSample = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetResultType(v string) *DescribeApiDocResponseBody {
	s.ResultType = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetSdkDemos(v *DescribeApiDocResponseBodySdkDemos) *DescribeApiDocResponseBody {
	s.SdkDemos = v
	return s
}

func (s *DescribeApiDocResponseBody) SetServiceTimeout(v int32) *DescribeApiDocResponseBody {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetStageName(v string) *DescribeApiDocResponseBody {
	s.StageName = &v
	return s
}

type DescribeApiDocResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyErrorCodeSamples) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeApiDocResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

type DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) SetCode(v string) *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Code = &v
	return s
}

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) SetDescription(v string) *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) SetMessage(v string) *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Message = &v
	return s
}

type DescribeApiDocResponseBodyPathParameters struct {
	PathParameter []*DescribeApiDocResponseBodyPathParametersPathParameter `json:"PathParameter,omitempty" xml:"PathParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyPathParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyPathParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyPathParameters) SetPathParameter(v []*DescribeApiDocResponseBodyPathParametersPathParameter) *DescribeApiDocResponseBodyPathParameters {
	s.PathParameter = v
	return s
}

type DescribeApiDocResponseBodyPathParametersPathParameter struct {
	ApiParameterName *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DemoValue        *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeApiDocResponseBodyPathParametersPathParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyPathParametersPathParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyPathParametersPathParameter) SetApiParameterName(v string) *DescribeApiDocResponseBodyPathParametersPathParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiDocResponseBodyPathParametersPathParameter) SetDemoValue(v string) *DescribeApiDocResponseBodyPathParametersPathParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyPathParametersPathParameter) SetDescription(v string) *DescribeApiDocResponseBodyPathParametersPathParameter {
	s.Description = &v
	return s
}

type DescribeApiDocResponseBodyRequestBody struct {
	RequestParam []*DescribeApiDocResponseBodyRequestBodyRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyRequestBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestBody) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestBody) SetRequestParam(v []*DescribeApiDocResponseBodyRequestBodyRequestParam) *DescribeApiDocResponseBodyRequestBody {
	s.RequestParam = v
	return s
}

type DescribeApiDocResponseBodyRequestBodyRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeApiDocResponseBodyRequestBodyRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestBodyRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetApiParameterName(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetDefaultValue(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetDemoValue(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetDescription(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetEnumValue(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetJsonScheme(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetMaxLength(v int64) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetMaxValue(v int64) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetMinLength(v int64) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetMinValue(v int64) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetParameterType(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetRegularExpression(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestBodyRequestParam) SetRequired(v string) *DescribeApiDocResponseBodyRequestBodyRequestParam {
	s.Required = &v
	return s
}

type DescribeApiDocResponseBodyRequestHeaders struct {
	RequestParam []*DescribeApiDocResponseBodyRequestHeadersRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestHeaders) SetRequestParam(v []*DescribeApiDocResponseBodyRequestHeadersRequestParam) *DescribeApiDocResponseBodyRequestHeaders {
	s.RequestParam = v
	return s
}

type DescribeApiDocResponseBodyRequestHeadersRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeApiDocResponseBodyRequestHeadersRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestHeadersRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetApiParameterName(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetDefaultValue(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetDemoValue(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetDescription(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetEnumValue(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetJsonScheme(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetMaxLength(v int64) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetMaxValue(v int64) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetMinLength(v int64) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetMinValue(v int64) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetParameterType(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetRegularExpression(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestHeadersRequestParam) SetRequired(v string) *DescribeApiDocResponseBodyRequestHeadersRequestParam {
	s.Required = &v
	return s
}

type DescribeApiDocResponseBodyRequestQueries struct {
	RequestParam []*DescribeApiDocResponseBodyRequestQueriesRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyRequestQueries) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestQueries) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestQueries) SetRequestParam(v []*DescribeApiDocResponseBodyRequestQueriesRequestParam) *DescribeApiDocResponseBodyRequestQueries {
	s.RequestParam = v
	return s
}

type DescribeApiDocResponseBodyRequestQueriesRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeApiDocResponseBodyRequestQueriesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestQueriesRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetApiParameterName(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetDefaultValue(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetDemoValue(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetDescription(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetEnumValue(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetJsonScheme(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetMaxLength(v int64) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetMaxValue(v int64) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetMinLength(v int64) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetMinValue(v int64) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetParameterType(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetRegularExpression(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestQueriesRequestParam) SetRequired(v string) *DescribeApiDocResponseBodyRequestQueriesRequestParam {
	s.Required = &v
	return s
}

type DescribeApiDocResponseBodyResultDescriptions struct {
	ResultDescription []*DescribeApiDocResponseBodyResultDescriptionsResultDescription `json:"ResultDescription,omitempty" xml:"ResultDescription,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyResultDescriptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyResultDescriptions) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyResultDescriptions) SetResultDescription(v []*DescribeApiDocResponseBodyResultDescriptionsResultDescription) *DescribeApiDocResponseBodyResultDescriptions {
	s.ResultDescription = v
	return s
}

type DescribeApiDocResponseBodyResultDescriptionsResultDescription struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasChild    *bool   `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Mandatory   *bool   `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApiDocResponseBodyResultDescriptionsResultDescription) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyResultDescriptionsResultDescription) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetDescription(v string) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetHasChild(v bool) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.HasChild = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetId(v string) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Id = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetKey(v string) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Key = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetMandatory(v bool) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Mandatory = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetName(v string) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Name = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetPid(v string) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Pid = &v
	return s
}

func (s *DescribeApiDocResponseBodyResultDescriptionsResultDescription) SetType(v string) *DescribeApiDocResponseBodyResultDescriptionsResultDescription {
	s.Type = &v
	return s
}

type DescribeApiDocResponseBodySdkDemos struct {
	SdkDemo []*DescribeApiDocResponseBodySdkDemosSdkDemo `json:"SdkDemo,omitempty" xml:"SdkDemo,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodySdkDemos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodySdkDemos) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodySdkDemos) SetSdkDemo(v []*DescribeApiDocResponseBodySdkDemosSdkDemo) *DescribeApiDocResponseBodySdkDemos {
	s.SdkDemo = v
	return s
}

type DescribeApiDocResponseBodySdkDemosSdkDemo struct {
	CallDemo *string `json:"CallDemo,omitempty" xml:"CallDemo,omitempty"`
	IdeKey   *string `json:"IdeKey,omitempty" xml:"IdeKey,omitempty"`
}

func (s DescribeApiDocResponseBodySdkDemosSdkDemo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodySdkDemosSdkDemo) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodySdkDemosSdkDemo) SetCallDemo(v string) *DescribeApiDocResponseBodySdkDemosSdkDemo {
	s.CallDemo = &v
	return s
}

func (s *DescribeApiDocResponseBodySdkDemosSdkDemo) SetIdeKey(v string) *DescribeApiDocResponseBodySdkDemosSdkDemo {
	s.IdeKey = &v
	return s
}

type DescribeApiDocResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiDocResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiDocResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponse) SetHeaders(v map[string]*string) *DescribeApiDocResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiDocResponse) SetStatusCode(v int32) *DescribeApiDocResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiDocResponse) SetBody(v *DescribeApiDocResponseBody) *DescribeApiDocResponse {
	s.Body = v
	return s
}

type DescribeApiDocsRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsRequest) SetApiId(v string) *DescribeApiDocsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocsRequest) SetApiName(v string) *DescribeApiDocsRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApiDocsRequest) SetGroupId(v string) *DescribeApiDocsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocsRequest) SetPageNumber(v int32) *DescribeApiDocsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiDocsRequest) SetPageSize(v int32) *DescribeApiDocsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiDocsRequest) SetSecurityToken(v string) *DescribeApiDocsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiDocsRequest) SetStageName(v string) *DescribeApiDocsRequest {
	s.StageName = &v
	return s
}

type DescribeApiDocsResponseBody struct {
	ApiInfos   *DescribeApiDocsResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiDocsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsResponseBody) SetApiInfos(v *DescribeApiDocsResponseBodyApiInfos) *DescribeApiDocsResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApiDocsResponseBody) SetPageNumber(v int32) *DescribeApiDocsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiDocsResponseBody) SetPageSize(v int32) *DescribeApiDocsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiDocsResponseBody) SetRequestId(v string) *DescribeApiDocsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiDocsResponseBody) SetTotalCount(v int32) *DescribeApiDocsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiDocsResponseBodyApiInfos struct {
	ApiInfo []*DescribeApiDocsResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiDocsResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsResponseBodyApiInfos) SetApiInfo(v []*DescribeApiDocsResponseBodyApiInfosApiInfo) *DescribeApiDocsResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApiDocsResponseBodyApiInfosApiInfo struct {
	ApiId            *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName          *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	DeployedTime     *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupDescription *string `json:"GroupDescription,omitempty" xml:"GroupDescription,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName        *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocsResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetDeployedTime(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.DeployedTime = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetGroupDescription(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.GroupDescription = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApiDocsResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApiDocsResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

type DescribeApiDocsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiDocsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiDocsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsResponse) SetHeaders(v map[string]*string) *DescribeApiDocsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiDocsResponse) SetStatusCode(v int32) *DescribeApiDocsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiDocsResponse) SetBody(v *DescribeApiDocsResponseBody) *DescribeApiDocsResponse {
	s.Body = v
	return s
}

type DescribeApiDocsForBackendRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocsForBackendRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsForBackendRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsForBackendRequest) SetAliUid(v int64) *DescribeApiDocsForBackendRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetApiId(v string) *DescribeApiDocsForBackendRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetApiName(v string) *DescribeApiDocsForBackendRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetGroupId(v string) *DescribeApiDocsForBackendRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetPageNumber(v int32) *DescribeApiDocsForBackendRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetPageSize(v int32) *DescribeApiDocsForBackendRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetSecurityToken(v string) *DescribeApiDocsForBackendRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiDocsForBackendRequest) SetStageName(v string) *DescribeApiDocsForBackendRequest {
	s.StageName = &v
	return s
}

type DescribeApiDocsForBackendResponseBody struct {
	ApiInfos   *DescribeApiDocsForBackendResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiDocsForBackendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsForBackendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsForBackendResponseBody) SetApiInfos(v *DescribeApiDocsForBackendResponseBodyApiInfos) *DescribeApiDocsForBackendResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApiDocsForBackendResponseBody) SetPageNumber(v int32) *DescribeApiDocsForBackendResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBody) SetPageSize(v int32) *DescribeApiDocsForBackendResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBody) SetRequestId(v string) *DescribeApiDocsForBackendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBody) SetTotalCount(v int32) *DescribeApiDocsForBackendResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiDocsForBackendResponseBodyApiInfos struct {
	ApiInfo []*DescribeApiDocsForBackendResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiDocsForBackendResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsForBackendResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfos) SetApiInfo(v []*DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) *DescribeApiDocsForBackendResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApiDocsForBackendResponseBodyApiInfosApiInfo struct {
	ApiId            *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName          *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	DeployedTime     *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupDescription *string `json:"GroupDescription,omitempty" xml:"GroupDescription,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName        *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetDeployedTime(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.DeployedTime = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetGroupDescription(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.GroupDescription = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApiDocsForBackendResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

type DescribeApiDocsForBackendResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiDocsForBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiDocsForBackendResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocsForBackendResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiDocsForBackendResponse) SetHeaders(v map[string]*string) *DescribeApiDocsForBackendResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiDocsForBackendResponse) SetStatusCode(v int32) *DescribeApiDocsForBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiDocsForBackendResponse) SetBody(v *DescribeApiDocsForBackendResponseBody) *DescribeApiDocsForBackendResponse {
	s.Body = v
	return s
}

type DescribeApiErrorRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiErrorRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorRequest) SetApiId(v string) *DescribeApiErrorRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiErrorRequest) SetEndTime(v string) *DescribeApiErrorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiErrorRequest) SetGroupId(v string) *DescribeApiErrorRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiErrorRequest) SetSecurityToken(v string) *DescribeApiErrorRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiErrorRequest) SetStartTime(v string) *DescribeApiErrorRequest {
	s.StartTime = &v
	return s
}

type DescribeApiErrorResponseBody struct {
	ClientErrors *DescribeApiErrorResponseBodyClientErrors `json:"ClientErrors,omitempty" xml:"ClientErrors,omitempty" type:"Struct"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerErrors *DescribeApiErrorResponseBodyServerErrors `json:"ServerErrors,omitempty" xml:"ServerErrors,omitempty" type:"Struct"`
}

func (s DescribeApiErrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorResponseBody) SetClientErrors(v *DescribeApiErrorResponseBodyClientErrors) *DescribeApiErrorResponseBody {
	s.ClientErrors = v
	return s
}

func (s *DescribeApiErrorResponseBody) SetRequestId(v string) *DescribeApiErrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiErrorResponseBody) SetServerErrors(v *DescribeApiErrorResponseBodyServerErrors) *DescribeApiErrorResponseBody {
	s.ServerErrors = v
	return s
}

type DescribeApiErrorResponseBodyClientErrors struct {
	ClientError []*DescribeApiErrorResponseBodyClientErrorsClientError `json:"ClientError,omitempty" xml:"ClientError,omitempty" type:"Repeated"`
}

func (s DescribeApiErrorResponseBodyClientErrors) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorResponseBodyClientErrors) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorResponseBodyClientErrors) SetClientError(v []*DescribeApiErrorResponseBodyClientErrorsClientError) *DescribeApiErrorResponseBodyClientErrors {
	s.ClientError = v
	return s
}

type DescribeApiErrorResponseBodyClientErrorsClientError struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiErrorResponseBodyClientErrorsClientError) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorResponseBodyClientErrorsClientError) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorResponseBodyClientErrorsClientError) SetTime(v string) *DescribeApiErrorResponseBodyClientErrorsClientError {
	s.Time = &v
	return s
}

func (s *DescribeApiErrorResponseBodyClientErrorsClientError) SetValue(v string) *DescribeApiErrorResponseBodyClientErrorsClientError {
	s.Value = &v
	return s
}

type DescribeApiErrorResponseBodyServerErrors struct {
	ServerError []*DescribeApiErrorResponseBodyServerErrorsServerError `json:"ServerError,omitempty" xml:"ServerError,omitempty" type:"Repeated"`
}

func (s DescribeApiErrorResponseBodyServerErrors) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorResponseBodyServerErrors) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorResponseBodyServerErrors) SetServerError(v []*DescribeApiErrorResponseBodyServerErrorsServerError) *DescribeApiErrorResponseBodyServerErrors {
	s.ServerError = v
	return s
}

type DescribeApiErrorResponseBodyServerErrorsServerError struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiErrorResponseBodyServerErrorsServerError) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorResponseBodyServerErrorsServerError) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorResponseBodyServerErrorsServerError) SetTime(v string) *DescribeApiErrorResponseBodyServerErrorsServerError {
	s.Time = &v
	return s
}

func (s *DescribeApiErrorResponseBodyServerErrorsServerError) SetValue(v string) *DescribeApiErrorResponseBodyServerErrorsServerError {
	s.Value = &v
	return s
}

type DescribeApiErrorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiErrorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiErrorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiErrorResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiErrorResponse) SetHeaders(v map[string]*string) *DescribeApiErrorResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiErrorResponse) SetStatusCode(v int32) *DescribeApiErrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiErrorResponse) SetBody(v *DescribeApiErrorResponseBody) *DescribeApiErrorResponse {
	s.Body = v
	return s
}

type DescribeApiGroupDetailRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailRequest) SetGroupId(v string) *DescribeApiGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupDetailRequest) SetSecurityToken(v string) *DescribeApiGroupDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApiGroupDetailResponseBody struct {
	BillingStatus *string                                        `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	CreatedTime   *string                                        `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainItems   *DescribeApiGroupDetailResponseBodyDomainItems `json:"DomainItems,omitempty" xml:"DomainItems,omitempty" type:"Struct"`
	GroupId       *string                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                        `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IllegalStatus *string                                        `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	ModifiedTime  *string                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId     *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	SubDomain     *string                                        `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	TrafficLimit  *int32                                         `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
}

func (s DescribeApiGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailResponseBody) SetBillingStatus(v string) *DescribeApiGroupDetailResponseBody {
	s.BillingStatus = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetCreatedTime(v string) *DescribeApiGroupDetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetDescription(v string) *DescribeApiGroupDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetDomainItems(v *DescribeApiGroupDetailResponseBodyDomainItems) *DescribeApiGroupDetailResponseBody {
	s.DomainItems = v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetGroupId(v string) *DescribeApiGroupDetailResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetGroupName(v string) *DescribeApiGroupDetailResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetIllegalStatus(v string) *DescribeApiGroupDetailResponseBody {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetModifiedTime(v string) *DescribeApiGroupDetailResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetRegionId(v string) *DescribeApiGroupDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetRequestId(v string) *DescribeApiGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetStatus(v string) *DescribeApiGroupDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetSubDomain(v string) *DescribeApiGroupDetailResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBody) SetTrafficLimit(v int32) *DescribeApiGroupDetailResponseBody {
	s.TrafficLimit = &v
	return s
}

type DescribeApiGroupDetailResponseBodyDomainItems struct {
	DomainItem []*DescribeApiGroupDetailResponseBodyDomainItemsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupDetailResponseBodyDomainItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailResponseBodyDomainItems) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailResponseBodyDomainItems) SetDomainItem(v []*DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) *DescribeApiGroupDetailResponseBodyDomainItems {
	s.DomainItem = v
	return s
}

type DescribeApiGroupDetailResponseBodyDomainItemsDomainItem struct {
	CertificateId        *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName      *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
	DomainStatus         *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
}

func (s DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetCertificateId(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.CertificateId = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetCertificateName(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.CertificateName = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainName(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainName = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainNameResolution(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainNameResolution = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainStatus(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainStatus = &v
	return s
}

type DescribeApiGroupDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailResponse) SetHeaders(v map[string]*string) *DescribeApiGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupDetailResponse) SetStatusCode(v int32) *DescribeApiGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGroupDetailResponse) SetBody(v *DescribeApiGroupDetailResponseBody) *DescribeApiGroupDetailResponse {
	s.Body = v
	return s
}

type DescribeApiGroupsRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsRequest) SetGroupId(v string) *DescribeApiGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetGroupName(v string) *DescribeApiGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetPageNumber(v int32) *DescribeApiGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetPageSize(v int32) *DescribeApiGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetSecurityToken(v string) *DescribeApiGroupsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApiGroupsResponseBody struct {
	ApiGroupAttributes *DescribeApiGroupsResponseBodyApiGroupAttributes `json:"ApiGroupAttributes,omitempty" xml:"ApiGroupAttributes,omitempty" type:"Struct"`
	PageNumber         *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBody) SetApiGroupAttributes(v *DescribeApiGroupsResponseBodyApiGroupAttributes) *DescribeApiGroupsResponseBody {
	s.ApiGroupAttributes = v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetPageNumber(v int32) *DescribeApiGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetPageSize(v int32) *DescribeApiGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetRequestId(v string) *DescribeApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupsResponseBody) SetTotalCount(v int32) *DescribeApiGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiGroupsResponseBodyApiGroupAttributes struct {
	ApiGroupAttribute []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute `json:"ApiGroupAttribute,omitempty" xml:"ApiGroupAttribute,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributes) SetApiGroupAttribute(v []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) *DescribeApiGroupsResponseBodyApiGroupAttributes {
	s.ApiGroupAttribute = v
	return s
}

type DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute struct {
	BillingStatus *string `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	CreatedTime   *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IllegalStatus *string `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubDomain     *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	TrafficLimit  *int32  `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetBillingStatus(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.BillingStatus = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetCreatedTime(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetDescription(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetGroupId(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetGroupName(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetIllegalStatus(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetModifiedTime(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetRegionId(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetSubDomain(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.SubDomain = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetTrafficLimit(v int32) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.TrafficLimit = &v
	return s
}

type DescribeApiGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponse) SetHeaders(v map[string]*string) *DescribeApiGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupsResponse) SetStatusCode(v int32) *DescribeApiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGroupsResponse) SetBody(v *DescribeApiGroupsResponseBody) *DescribeApiGroupsResponse {
	s.Body = v
	return s
}

type DescribeApiLatencyRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiLatencyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyRequest) SetApiId(v string) *DescribeApiLatencyRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiLatencyRequest) SetEndTime(v string) *DescribeApiLatencyRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiLatencyRequest) SetGroupId(v string) *DescribeApiLatencyRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiLatencyRequest) SetSecurityToken(v string) *DescribeApiLatencyRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiLatencyRequest) SetStartTime(v string) *DescribeApiLatencyRequest {
	s.StartTime = &v
	return s
}

type DescribeApiLatencyResponseBody struct {
	Latencys  *DescribeApiLatencyResponseBodyLatencys `json:"Latencys,omitempty" xml:"Latencys,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiLatencyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyResponseBody) SetLatencys(v *DescribeApiLatencyResponseBodyLatencys) *DescribeApiLatencyResponseBody {
	s.Latencys = v
	return s
}

func (s *DescribeApiLatencyResponseBody) SetRequestId(v string) *DescribeApiLatencyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiLatencyResponseBodyLatencys struct {
	Latency []*DescribeApiLatencyResponseBodyLatencysLatency `json:"Latency,omitempty" xml:"Latency,omitempty" type:"Repeated"`
}

func (s DescribeApiLatencyResponseBodyLatencys) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyResponseBodyLatencys) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyResponseBodyLatencys) SetLatency(v []*DescribeApiLatencyResponseBodyLatencysLatency) *DescribeApiLatencyResponseBodyLatencys {
	s.Latency = v
	return s
}

type DescribeApiLatencyResponseBodyLatencysLatency struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiLatencyResponseBodyLatencysLatency) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyResponseBodyLatencysLatency) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyResponseBodyLatencysLatency) SetTime(v string) *DescribeApiLatencyResponseBodyLatencysLatency {
	s.Time = &v
	return s
}

func (s *DescribeApiLatencyResponseBodyLatencysLatency) SetValue(v string) *DescribeApiLatencyResponseBodyLatencysLatency {
	s.Value = &v
	return s
}

type DescribeApiLatencyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiLatencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiLatencyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyResponse) SetHeaders(v map[string]*string) *DescribeApiLatencyResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiLatencyResponse) SetStatusCode(v int32) *DescribeApiLatencyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiLatencyResponse) SetBody(v *DescribeApiLatencyResponseBody) *DescribeApiLatencyResponse {
	s.Body = v
	return s
}

type DescribeApiMarketInstanceRequest struct {
	AliUid        *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiMarketInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMarketInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketInstanceRequest) SetAliUid(v string) *DescribeApiMarketInstanceRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeApiMarketInstanceRequest) SetGroupId(v string) *DescribeApiMarketInstanceRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiMarketInstanceRequest) SetSecurityToken(v string) *DescribeApiMarketInstanceRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApiMarketInstanceResponseBody struct {
	InstanceAttributes *string `json:"InstanceAttributes,omitempty" xml:"InstanceAttributes,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiMarketInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMarketInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketInstanceResponseBody) SetInstanceAttributes(v string) *DescribeApiMarketInstanceResponseBody {
	s.InstanceAttributes = &v
	return s
}

func (s *DescribeApiMarketInstanceResponseBody) SetRequestId(v string) *DescribeApiMarketInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiMarketInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiMarketInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiMarketInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMarketInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketInstanceResponse) SetHeaders(v map[string]*string) *DescribeApiMarketInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiMarketInstanceResponse) SetStatusCode(v int32) *DescribeApiMarketInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiMarketInstanceResponse) SetBody(v *DescribeApiMarketInstanceResponseBody) *DescribeApiMarketInstanceResponse {
	s.Body = v
	return s
}

type DescribeApiQpsRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiQpsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsRequest) SetApiId(v string) *DescribeApiQpsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiQpsRequest) SetEndTime(v string) *DescribeApiQpsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiQpsRequest) SetGroupId(v string) *DescribeApiQpsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiQpsRequest) SetSecurityToken(v string) *DescribeApiQpsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiQpsRequest) SetStartTime(v string) *DescribeApiQpsRequest {
	s.StartTime = &v
	return s
}

type DescribeApiQpsResponseBody struct {
	Fails     *DescribeApiQpsResponseBodyFails     `json:"Fails,omitempty" xml:"Fails,omitempty" type:"Struct"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Successes *DescribeApiQpsResponseBodySuccesses `json:"Successes,omitempty" xml:"Successes,omitempty" type:"Struct"`
}

func (s DescribeApiQpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsResponseBody) SetFails(v *DescribeApiQpsResponseBodyFails) *DescribeApiQpsResponseBody {
	s.Fails = v
	return s
}

func (s *DescribeApiQpsResponseBody) SetRequestId(v string) *DescribeApiQpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiQpsResponseBody) SetSuccesses(v *DescribeApiQpsResponseBodySuccesses) *DescribeApiQpsResponseBody {
	s.Successes = v
	return s
}

type DescribeApiQpsResponseBodyFails struct {
	Fail []*DescribeApiQpsResponseBodyFailsFail `json:"Fail,omitempty" xml:"Fail,omitempty" type:"Repeated"`
}

func (s DescribeApiQpsResponseBodyFails) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsResponseBodyFails) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsResponseBodyFails) SetFail(v []*DescribeApiQpsResponseBodyFailsFail) *DescribeApiQpsResponseBodyFails {
	s.Fail = v
	return s
}

type DescribeApiQpsResponseBodyFailsFail struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiQpsResponseBodyFailsFail) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsResponseBodyFailsFail) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsResponseBodyFailsFail) SetTime(v string) *DescribeApiQpsResponseBodyFailsFail {
	s.Time = &v
	return s
}

func (s *DescribeApiQpsResponseBodyFailsFail) SetValue(v string) *DescribeApiQpsResponseBodyFailsFail {
	s.Value = &v
	return s
}

type DescribeApiQpsResponseBodySuccesses struct {
	Success []*DescribeApiQpsResponseBodySuccessesSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Repeated"`
}

func (s DescribeApiQpsResponseBodySuccesses) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsResponseBodySuccesses) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsResponseBodySuccesses) SetSuccess(v []*DescribeApiQpsResponseBodySuccessesSuccess) *DescribeApiQpsResponseBodySuccesses {
	s.Success = v
	return s
}

type DescribeApiQpsResponseBodySuccessesSuccess struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiQpsResponseBodySuccessesSuccess) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsResponseBodySuccessesSuccess) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsResponseBodySuccessesSuccess) SetTime(v string) *DescribeApiQpsResponseBodySuccessesSuccess {
	s.Time = &v
	return s
}

func (s *DescribeApiQpsResponseBodySuccessesSuccess) SetValue(v string) *DescribeApiQpsResponseBodySuccessesSuccess {
	s.Value = &v
	return s
}

type DescribeApiQpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiQpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiQpsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsResponse) SetHeaders(v map[string]*string) *DescribeApiQpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiQpsResponse) SetStatusCode(v int32) *DescribeApiQpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiQpsResponse) SetBody(v *DescribeApiQpsResponseBody) *DescribeApiQpsResponse {
	s.Body = v
	return s
}

type DescribeApiRulesRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiRulesRequest) SetApiIds(v string) *DescribeApiRulesRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiRulesRequest) SetGroupId(v string) *DescribeApiRulesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiRulesRequest) SetPageNumber(v int32) *DescribeApiRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiRulesRequest) SetPageSize(v int32) *DescribeApiRulesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiRulesRequest) SetRuleType(v string) *DescribeApiRulesRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeApiRulesRequest) SetSecurityToken(v string) *DescribeApiRulesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiRulesRequest) SetStageName(v string) *DescribeApiRulesRequest {
	s.StageName = &v
	return s
}

type DescribeApiRulesResponseBody struct {
	ApiRules   *DescribeApiRulesResponseBodyApiRules `json:"ApiRules,omitempty" xml:"ApiRules,omitempty" type:"Struct"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiRulesResponseBody) SetApiRules(v *DescribeApiRulesResponseBodyApiRules) *DescribeApiRulesResponseBody {
	s.ApiRules = v
	return s
}

func (s *DescribeApiRulesResponseBody) SetPageNumber(v int32) *DescribeApiRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiRulesResponseBody) SetPageSize(v int32) *DescribeApiRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiRulesResponseBody) SetRequestId(v string) *DescribeApiRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiRulesResponseBody) SetTotalCount(v int32) *DescribeApiRulesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiRulesResponseBodyApiRules struct {
	ApiRule []*DescribeApiRulesResponseBodyApiRulesApiRule `json:"ApiRule,omitempty" xml:"ApiRule,omitempty" type:"Repeated"`
}

func (s DescribeApiRulesResponseBodyApiRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiRulesResponseBodyApiRules) GoString() string {
	return s.String()
}

func (s *DescribeApiRulesResponseBodyApiRules) SetApiRule(v []*DescribeApiRulesResponseBodyApiRulesApiRule) *DescribeApiRulesResponseBodyApiRules {
	s.ApiRule = v
	return s
}

type DescribeApiRulesResponseBodyApiRulesApiRule struct {
	ApiId       *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName     *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeApiRulesResponseBodyApiRulesApiRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiRulesResponseBodyApiRulesApiRule) GoString() string {
	return s.String()
}

func (s *DescribeApiRulesResponseBodyApiRulesApiRule) SetApiId(v string) *DescribeApiRulesResponseBodyApiRulesApiRule {
	s.ApiId = &v
	return s
}

func (s *DescribeApiRulesResponseBodyApiRulesApiRule) SetApiName(v string) *DescribeApiRulesResponseBodyApiRulesApiRule {
	s.ApiName = &v
	return s
}

func (s *DescribeApiRulesResponseBodyApiRulesApiRule) SetCreatedTime(v string) *DescribeApiRulesResponseBodyApiRulesApiRule {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiRulesResponseBodyApiRulesApiRule) SetRuleId(v string) *DescribeApiRulesResponseBodyApiRulesApiRule {
	s.RuleId = &v
	return s
}

func (s *DescribeApiRulesResponseBodyApiRulesApiRule) SetRuleName(v string) *DescribeApiRulesResponseBodyApiRulesApiRule {
	s.RuleName = &v
	return s
}

func (s *DescribeApiRulesResponseBodyApiRulesApiRule) SetRuleType(v string) *DescribeApiRulesResponseBodyApiRulesApiRule {
	s.RuleType = &v
	return s
}

type DescribeApiRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiRulesResponse) SetHeaders(v map[string]*string) *DescribeApiRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiRulesResponse) SetStatusCode(v int32) *DescribeApiRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiRulesResponse) SetBody(v *DescribeApiRulesResponseBody) *DescribeApiRulesResponse {
	s.Body = v
	return s
}

type DescribeApiTrafficRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiTrafficRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficRequest) SetApiId(v string) *DescribeApiTrafficRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiTrafficRequest) SetEndTime(v string) *DescribeApiTrafficRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiTrafficRequest) SetGroupId(v string) *DescribeApiTrafficRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiTrafficRequest) SetSecurityToken(v string) *DescribeApiTrafficRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiTrafficRequest) SetStartTime(v string) *DescribeApiTrafficRequest {
	s.StartTime = &v
	return s
}

type DescribeApiTrafficResponseBody struct {
	Downloads *DescribeApiTrafficResponseBodyDownloads `json:"Downloads,omitempty" xml:"Downloads,omitempty" type:"Struct"`
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Uploads   *DescribeApiTrafficResponseBodyUploads   `json:"Uploads,omitempty" xml:"Uploads,omitempty" type:"Struct"`
}

func (s DescribeApiTrafficResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficResponseBody) SetDownloads(v *DescribeApiTrafficResponseBodyDownloads) *DescribeApiTrafficResponseBody {
	s.Downloads = v
	return s
}

func (s *DescribeApiTrafficResponseBody) SetRequestId(v string) *DescribeApiTrafficResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiTrafficResponseBody) SetUploads(v *DescribeApiTrafficResponseBodyUploads) *DescribeApiTrafficResponseBody {
	s.Uploads = v
	return s
}

type DescribeApiTrafficResponseBodyDownloads struct {
	Download []*DescribeApiTrafficResponseBodyDownloadsDownload `json:"Download,omitempty" xml:"Download,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficResponseBodyDownloads) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficResponseBodyDownloads) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficResponseBodyDownloads) SetDownload(v []*DescribeApiTrafficResponseBodyDownloadsDownload) *DescribeApiTrafficResponseBodyDownloads {
	s.Download = v
	return s
}

type DescribeApiTrafficResponseBodyDownloadsDownload struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiTrafficResponseBodyDownloadsDownload) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficResponseBodyDownloadsDownload) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficResponseBodyDownloadsDownload) SetTime(v string) *DescribeApiTrafficResponseBodyDownloadsDownload {
	s.Time = &v
	return s
}

func (s *DescribeApiTrafficResponseBodyDownloadsDownload) SetValue(v string) *DescribeApiTrafficResponseBodyDownloadsDownload {
	s.Value = &v
	return s
}

type DescribeApiTrafficResponseBodyUploads struct {
	Upload []*DescribeApiTrafficResponseBodyUploadsUpload `json:"Upload,omitempty" xml:"Upload,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficResponseBodyUploads) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficResponseBodyUploads) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficResponseBodyUploads) SetUpload(v []*DescribeApiTrafficResponseBodyUploadsUpload) *DescribeApiTrafficResponseBodyUploads {
	s.Upload = v
	return s
}

type DescribeApiTrafficResponseBodyUploadsUpload struct {
	Time  *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiTrafficResponseBodyUploadsUpload) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficResponseBodyUploadsUpload) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficResponseBodyUploadsUpload) SetTime(v string) *DescribeApiTrafficResponseBodyUploadsUpload {
	s.Time = &v
	return s
}

func (s *DescribeApiTrafficResponseBodyUploadsUpload) SetValue(v string) *DescribeApiTrafficResponseBodyUploadsUpload {
	s.Value = &v
	return s
}

type DescribeApiTrafficResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiTrafficResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficResponse) SetHeaders(v map[string]*string) *DescribeApiTrafficResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiTrafficResponse) SetStatusCode(v int32) *DescribeApiTrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiTrafficResponse) SetBody(v *DescribeApiTrafficResponseBody) *DescribeApiTrafficResponse {
	s.Body = v
	return s
}

type DescribeApisRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisRequest) SetApiId(v string) *DescribeApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisRequest) SetApiName(v string) *DescribeApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApisRequest) SetGroupId(v string) *DescribeApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApisRequest) SetPageNumber(v int32) *DescribeApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisRequest) SetPageSize(v int32) *DescribeApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisRequest) SetSecurityToken(v string) *DescribeApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisRequest) SetVisibility(v string) *DescribeApisRequest {
	s.Visibility = &v
	return s
}

type DescribeApisResponseBody struct {
	ApiInfos   *DescribeApisResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBody) SetApiInfos(v *DescribeApisResponseBodyApiInfos) *DescribeApisResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisResponseBody) SetPageNumber(v int32) *DescribeApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisResponseBody) SetPageSize(v int32) *DescribeApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisResponseBody) SetRequestId(v string) *DescribeApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisResponseBody) SetTotalCount(v int32) *DescribeApisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiInfos) SetApiInfo(v []*DescribeApisResponseBodyApiInfosApiInfo) *DescribeApisResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApisResponseBodyApiInfosApiInfo struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Visibility   *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetCreatedTime(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetModifiedTime(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeApisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisResponse) SetHeaders(v map[string]*string) *DescribeApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisResponse) SetStatusCode(v int32) *DescribeApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisResponse) SetBody(v *DescribeApisResponseBody) *DescribeApisResponse {
	s.Body = v
	return s
}

type DescribeApisByAppRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApisByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppRequest) SetAppId(v int64) *DescribeApisByAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApisByAppRequest) SetPageNumber(v int32) *DescribeApisByAppRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByAppRequest) SetPageSize(v int32) *DescribeApisByAppRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByAppRequest) SetSecurityToken(v string) *DescribeApisByAppRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApisByAppResponseBody struct {
	AppApiRelationInfos *DescribeApisByAppResponseBodyAppApiRelationInfos `json:"AppApiRelationInfos,omitempty" xml:"AppApiRelationInfos,omitempty" type:"Struct"`
	PageNumber          *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponseBody) SetAppApiRelationInfos(v *DescribeApisByAppResponseBodyAppApiRelationInfos) *DescribeApisByAppResponseBody {
	s.AppApiRelationInfos = v
	return s
}

func (s *DescribeApisByAppResponseBody) SetPageNumber(v int32) *DescribeApisByAppResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByAppResponseBody) SetPageSize(v int32) *DescribeApisByAppResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByAppResponseBody) SetRequestId(v string) *DescribeApisByAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByAppResponseBody) SetTotalCount(v int32) *DescribeApisByAppResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisByAppResponseBodyAppApiRelationInfos struct {
	AppApiRelationInfo []*DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo `json:"AppApiRelationInfo,omitempty" xml:"AppApiRelationInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfos) SetAppApiRelationInfo(v []*DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) *DescribeApisByAppResponseBodyAppApiRelationInfos {
	s.AppApiRelationInfo = v
	return s
}

type DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo struct {
	ApiId               *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	CreatedTime         *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Operator            *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName           *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetApiId(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetApiName(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAuthorizationSource(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetCreatedTime(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetDescription(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetGroupId(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetGroupName(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetOperator(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Operator = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetRegionId(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetStageName(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.StageName = &v
	return s
}

type DescribeApisByAppResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppResponse) SetHeaders(v map[string]*string) *DescribeApisByAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByAppResponse) SetStatusCode(v int32) *DescribeApisByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByAppResponse) SetBody(v *DescribeApisByAppResponseBody) *DescribeApisByAppResponse {
	s.Body = v
	return s
}

type DescribeApisByRuleRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApisByRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByRuleRequest) SetPageNumber(v int32) *DescribeApisByRuleRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByRuleRequest) SetPageSize(v int32) *DescribeApisByRuleRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByRuleRequest) SetRuleId(v string) *DescribeApisByRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeApisByRuleRequest) SetRuleType(v string) *DescribeApisByRuleRequest {
	s.RuleType = &v
	return s
}

func (s *DescribeApisByRuleRequest) SetSecurityToken(v string) *DescribeApisByRuleRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApisByRuleResponseBody struct {
	ApiInfos   *DescribeApisByRuleResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByRuleResponseBody) SetApiInfos(v *DescribeApisByRuleResponseBodyApiInfos) *DescribeApisByRuleResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisByRuleResponseBody) SetPageNumber(v int32) *DescribeApisByRuleResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByRuleResponseBody) SetPageSize(v int32) *DescribeApisByRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByRuleResponseBody) SetRequestId(v string) *DescribeApisByRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByRuleResponseBody) SetTotalCount(v int32) *DescribeApisByRuleResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisByRuleResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisByRuleResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByRuleResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByRuleResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByRuleResponseBodyApiInfos) SetApiInfo(v []*DescribeApisByRuleResponseBodyApiInfosApiInfo) *DescribeApisByRuleResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApisByRuleResponseBodyApiInfosApiInfo struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName    *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility   *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisByRuleResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByRuleResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetCreatedTime(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetModifiedTime(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByRuleResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisByRuleResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeApisByRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisByRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByRuleResponse) SetHeaders(v map[string]*string) *DescribeApisByRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByRuleResponse) SetStatusCode(v int32) *DescribeApisByRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByRuleResponse) SetBody(v *DescribeApisByRuleResponseBody) *DescribeApisByRuleResponse {
	s.Body = v
	return s
}

type DescribeApisForConsoleRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisForConsoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleRequest) SetApiId(v string) *DescribeApisForConsoleRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetApiName(v string) *DescribeApisForConsoleRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetGroupId(v string) *DescribeApisForConsoleRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetPageNumber(v int32) *DescribeApisForConsoleRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetPageSize(v int32) *DescribeApisForConsoleRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetSecurityToken(v string) *DescribeApisForConsoleRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetStageName(v string) *DescribeApisForConsoleRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApisForConsoleRequest) SetVisibility(v string) *DescribeApisForConsoleRequest {
	s.Visibility = &v
	return s
}

type DescribeApisForConsoleResponseBody struct {
	ApiInfos   *DescribeApisForConsoleResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisForConsoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleResponseBody) SetApiInfos(v *DescribeApisForConsoleResponseBodyApiInfos) *DescribeApisForConsoleResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisForConsoleResponseBody) SetPageNumber(v int32) *DescribeApisForConsoleResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisForConsoleResponseBody) SetPageSize(v int32) *DescribeApisForConsoleResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisForConsoleResponseBody) SetRequestId(v string) *DescribeApisForConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisForConsoleResponseBody) SetTotalCount(v int32) *DescribeApisForConsoleResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisForConsoleResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisForConsoleResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisForConsoleResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleResponseBodyApiInfos) SetApiInfo(v []*DescribeApisForConsoleResponseBodyApiInfosApiInfo) *DescribeApisForConsoleResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApisForConsoleResponseBodyApiInfosApiInfo struct {
	ApiId         *string                                                         `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string                                                         `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CreatedTime   *string                                                         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	DeployedInfos *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos `json:"DeployedInfos,omitempty" xml:"DeployedInfos,omitempty" type:"Struct"`
	Description   *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string                                                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                                         `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifiedTime  *string                                                         `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Visibility    *string                                                         `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisForConsoleResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetCreatedTime(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetDeployedInfos(v *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.DeployedInfos = v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetModifiedTime(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos struct {
	DeployedInfo []*DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo `json:"DeployedInfo,omitempty" xml:"DeployedInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos) SetDeployedInfo(v []*DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo) *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfos {
	s.DeployedInfo = v
	return s
}

type DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo struct {
	DeployedStatus   *string `json:"DeployedStatus,omitempty" xml:"DeployedStatus,omitempty"`
	EffectiveVersion *string `json:"EffectiveVersion,omitempty" xml:"EffectiveVersion,omitempty"`
	StageName        *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo) SetDeployedStatus(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo {
	s.DeployedStatus = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo) SetEffectiveVersion(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo {
	s.EffectiveVersion = &v
	return s
}

func (s *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo) SetStageName(v string) *DescribeApisForConsoleResponseBodyApiInfosApiInfoDeployedInfosDeployedInfo {
	s.StageName = &v
	return s
}

type DescribeApisForConsoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisForConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisForConsoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisForConsoleResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisForConsoleResponse) SetHeaders(v map[string]*string) *DescribeApisForConsoleResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisForConsoleResponse) SetStatusCode(v int32) *DescribeApisForConsoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisForConsoleResponse) SetBody(v *DescribeApisForConsoleResponseBody) *DescribeApisForConsoleResponse {
	s.Body = v
	return s
}

type DescribeAppRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppRequest) SetAppId(v int64) *DescribeAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppRequest) SetSecurityToken(v string) *DescribeAppRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAppResponseBody struct {
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBody) SetAppId(v int64) *DescribeAppResponseBody {
	s.AppId = &v
	return s
}

func (s *DescribeAppResponseBody) SetAppName(v string) *DescribeAppResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeAppResponseBody) SetCreatedTime(v string) *DescribeAppResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppResponseBody) SetDescription(v string) *DescribeAppResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBody) SetModifiedTime(v string) *DescribeAppResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppResponseBody) SetRequestId(v string) *DescribeAppResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppResponse) SetHeaders(v map[string]*string) *DescribeAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppResponse) SetStatusCode(v int32) *DescribeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppResponse) SetBody(v *DescribeAppResponseBody) *DescribeAppResponse {
	s.Body = v
	return s
}

type DescribeAppSecuritiesRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppSecuritiesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecuritiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesRequest) SetAppId(v int64) *DescribeAppSecuritiesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppSecuritiesRequest) SetSecurityToken(v string) *DescribeAppSecuritiesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAppSecuritiesResponseBody struct {
	AppSecuritys *DescribeAppSecuritiesResponseBodyAppSecuritys `json:"AppSecuritys,omitempty" xml:"AppSecuritys,omitempty" type:"Struct"`
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppSecuritiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecuritiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponseBody) SetAppSecuritys(v *DescribeAppSecuritiesResponseBodyAppSecuritys) *DescribeAppSecuritiesResponseBody {
	s.AppSecuritys = v
	return s
}

func (s *DescribeAppSecuritiesResponseBody) SetRequestId(v string) *DescribeAppSecuritiesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppSecuritiesResponseBodyAppSecuritys struct {
	AppSecurity []*DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity `json:"AppSecurity,omitempty" xml:"AppSecurity,omitempty" type:"Repeated"`
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritys) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritys) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritys) SetAppSecurity(v []*DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) *DescribeAppSecuritiesResponseBodyAppSecuritys {
	s.AppSecurity = v
	return s
}

type DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppKey       *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppSecret    *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetAppCode(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.AppCode = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetAppKey(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.AppKey = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetAppSecret(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.AppSecret = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetCreatedTime(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity) SetModifiedTime(v string) *DescribeAppSecuritiesResponseBodyAppSecuritysAppSecurity {
	s.ModifiedTime = &v
	return s
}

type DescribeAppSecuritiesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppSecuritiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppSecuritiesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecuritiesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesResponse) SetHeaders(v map[string]*string) *DescribeAppSecuritiesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppSecuritiesResponse) SetStatusCode(v int32) *DescribeAppSecuritiesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppSecuritiesResponse) SetBody(v *DescribeAppSecuritiesResponseBody) *DescribeAppSecuritiesResponse {
	s.Body = v
	return s
}

type DescribeAppSecurityRequest struct {
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s DescribeAppSecurityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityRequest) SetAppKey(v string) *DescribeAppSecurityRequest {
	s.AppKey = &v
	return s
}

type DescribeAppSecurityResponseBody struct {
	AppKey       *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppSecret    *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppSecurityResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityResponseBody) SetAppKey(v string) *DescribeAppSecurityResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetAppSecret(v string) *DescribeAppSecurityResponseBody {
	s.AppSecret = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetCreatedTime(v string) *DescribeAppSecurityResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetModifiedTime(v string) *DescribeAppSecurityResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppSecurityResponseBody) SetRequestId(v string) *DescribeAppSecurityResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppSecurityResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppSecurityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityResponse) SetHeaders(v map[string]*string) *DescribeAppSecurityResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppSecurityResponse) SetStatusCode(v int32) *DescribeAppSecurityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppSecurityResponse) SetBody(v *DescribeAppSecurityResponseBody) *DescribeAppSecurityResponse {
	s.Body = v
	return s
}

type DescribeAppSecurityForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppSecurityForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityForInnerRequest) SetAliUid(v int64) *DescribeAppSecurityForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeAppSecurityForInnerRequest) SetAppId(v int64) *DescribeAppSecurityForInnerRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppSecurityForInnerRequest) SetSecurityToken(v string) *DescribeAppSecurityForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAppSecurityForInnerResponseBody struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppKey       *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppSecret    *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppSecurityForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityForInnerResponseBody) SetAppCode(v string) *DescribeAppSecurityForInnerResponseBody {
	s.AppCode = &v
	return s
}

func (s *DescribeAppSecurityForInnerResponseBody) SetAppKey(v string) *DescribeAppSecurityForInnerResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeAppSecurityForInnerResponseBody) SetAppSecret(v string) *DescribeAppSecurityForInnerResponseBody {
	s.AppSecret = &v
	return s
}

func (s *DescribeAppSecurityForInnerResponseBody) SetCreatedTime(v string) *DescribeAppSecurityForInnerResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppSecurityForInnerResponseBody) SetModifiedTime(v string) *DescribeAppSecurityForInnerResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppSecurityForInnerResponseBody) SetRequestId(v string) *DescribeAppSecurityForInnerResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppSecurityForInnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppSecurityForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppSecurityForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityForInnerResponse) SetHeaders(v map[string]*string) *DescribeAppSecurityForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppSecurityForInnerResponse) SetStatusCode(v int32) *DescribeAppSecurityForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppSecurityForInnerResponse) SetBody(v *DescribeAppSecurityForInnerResponseBody) *DescribeAppSecurityForInnerResponse {
	s.Body = v
	return s
}

type DescribeAppsRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) SetAppId(v int64) *DescribeAppsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppsRequest) SetPageNumber(v int32) *DescribeAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsRequest) SetPageSize(v int32) *DescribeAppsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsRequest) SetSecurityToken(v string) *DescribeAppsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAppsResponseBody struct {
	Apps       *DescribeAppsResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Struct"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetApps(v *DescribeAppsResponseBodyApps) *DescribeAppsResponseBody {
	s.Apps = v
	return s
}

func (s *DescribeAppsResponseBody) SetPageNumber(v int32) *DescribeAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsResponseBody) SetPageSize(v int32) *DescribeAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalCount(v int32) *DescribeAppsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppsResponseBodyApps struct {
	App []*DescribeAppsResponseBodyAppsApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyApps) SetApp(v []*DescribeAppsResponseBodyAppsApp) *DescribeAppsResponseBodyApps {
	s.App = v
	return s
}

type DescribeAppsResponseBodyAppsApp struct {
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeAppsResponseBodyAppsApp) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyAppsApp) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppsApp) SetAppId(v int64) *DescribeAppsResponseBodyAppsApp {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsApp) SetAppName(v string) *DescribeAppsResponseBodyAppsApp {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsApp) SetCreatedTime(v string) *DescribeAppsResponseBodyAppsApp {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsApp) SetDescription(v string) *DescribeAppsResponseBodyAppsApp {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsApp) SetModifiedTime(v string) *DescribeAppsResponseBodyAppsApp {
	s.ModifiedTime = &v
	return s
}

type DescribeAppsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetStatusCode(v int32) *DescribeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeAppsByApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAppsByApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiRequest) SetApiId(v string) *DescribeAppsByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeAppsByApiRequest) SetGroupId(v string) *DescribeAppsByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAppsByApiRequest) SetPageNumber(v int32) *DescribeAppsByApiRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsByApiRequest) SetPageSize(v int32) *DescribeAppsByApiRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsByApiRequest) SetSecurityToken(v string) *DescribeAppsByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppsByApiRequest) SetStageName(v string) *DescribeAppsByApiRequest {
	s.StageName = &v
	return s
}

type DescribeAppsByApiResponseBody struct {
	AppApiRelationInfos *DescribeAppsByApiResponseBodyAppApiRelationInfos `json:"AppApiRelationInfos,omitempty" xml:"AppApiRelationInfos,omitempty" type:"Struct"`
	PageNumber          *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsByApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiResponseBody) SetAppApiRelationInfos(v *DescribeAppsByApiResponseBodyAppApiRelationInfos) *DescribeAppsByApiResponseBody {
	s.AppApiRelationInfos = v
	return s
}

func (s *DescribeAppsByApiResponseBody) SetPageNumber(v int32) *DescribeAppsByApiResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsByApiResponseBody) SetPageSize(v int32) *DescribeAppsByApiResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsByApiResponseBody) SetRequestId(v string) *DescribeAppsByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsByApiResponseBody) SetTotalCount(v int32) *DescribeAppsByApiResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppsByApiResponseBodyAppApiRelationInfos struct {
	AppApiRelationInfo []*DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo `json:"AppApiRelationInfo,omitempty" xml:"AppApiRelationInfo,omitempty" type:"Repeated"`
}

func (s DescribeAppsByApiResponseBodyAppApiRelationInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsByApiResponseBodyAppApiRelationInfos) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfos) SetAppApiRelationInfo(v []*DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) *DescribeAppsByApiResponseBodyAppApiRelationInfos {
	s.AppApiRelationInfo = v
	return s
}

type DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo struct {
	AppId               *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	CreatedTime         *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Operator            *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	StageName           *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAppId(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AppId = &v
	return s
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAppName(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AppName = &v
	return s
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAuthorizationSource(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetCreatedTime(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetDescription(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Description = &v
	return s
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetOperator(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Operator = &v
	return s
}

func (s *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo) SetStageName(v string) *DescribeAppsByApiResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.StageName = &v
	return s
}

type DescribeAppsByApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsByApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppsByApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiResponse) SetHeaders(v map[string]*string) *DescribeAppsByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsByApiResponse) SetStatusCode(v int32) *DescribeAppsByApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsByApiResponse) SetBody(v *DescribeAppsByApiResponseBody) *DescribeAppsByApiResponse {
	s.Body = v
	return s
}

type DescribeAppsForProviderRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppOwner      *string `json:"AppOwner,omitempty" xml:"AppOwner,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppsForProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsForProviderRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsForProviderRequest) SetAppId(v int64) *DescribeAppsForProviderRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppsForProviderRequest) SetAppOwner(v string) *DescribeAppsForProviderRequest {
	s.AppOwner = &v
	return s
}

func (s *DescribeAppsForProviderRequest) SetPageNumber(v int32) *DescribeAppsForProviderRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsForProviderRequest) SetPageSize(v int32) *DescribeAppsForProviderRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsForProviderRequest) SetSecurityToken(v string) *DescribeAppsForProviderRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAppsForProviderResponseBody struct {
	Apps       *DescribeAppsForProviderResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Struct"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppsForProviderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsForProviderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsForProviderResponseBody) SetApps(v *DescribeAppsForProviderResponseBodyApps) *DescribeAppsForProviderResponseBody {
	s.Apps = v
	return s
}

func (s *DescribeAppsForProviderResponseBody) SetPageNumber(v int32) *DescribeAppsForProviderResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppsForProviderResponseBody) SetPageSize(v int32) *DescribeAppsForProviderResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppsForProviderResponseBody) SetRequestId(v string) *DescribeAppsForProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsForProviderResponseBody) SetTotalCount(v int32) *DescribeAppsForProviderResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppsForProviderResponseBodyApps struct {
	App []*DescribeAppsForProviderResponseBodyAppsApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s DescribeAppsForProviderResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsForProviderResponseBodyApps) GoString() string {
	return s.String()
}

func (s *DescribeAppsForProviderResponseBodyApps) SetApp(v []*DescribeAppsForProviderResponseBodyAppsApp) *DescribeAppsForProviderResponseBodyApps {
	s.App = v
	return s
}

type DescribeAppsForProviderResponseBodyAppsApp struct {
	AppId        *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeAppsForProviderResponseBodyAppsApp) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsForProviderResponseBodyAppsApp) GoString() string {
	return s.String()
}

func (s *DescribeAppsForProviderResponseBodyAppsApp) SetAppId(v int64) *DescribeAppsForProviderResponseBodyAppsApp {
	s.AppId = &v
	return s
}

func (s *DescribeAppsForProviderResponseBodyAppsApp) SetAppName(v string) *DescribeAppsForProviderResponseBodyAppsApp {
	s.AppName = &v
	return s
}

func (s *DescribeAppsForProviderResponseBodyAppsApp) SetCreateTime(v string) *DescribeAppsForProviderResponseBodyAppsApp {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppsForProviderResponseBodyAppsApp) SetDescription(v string) *DescribeAppsForProviderResponseBodyAppsApp {
	s.Description = &v
	return s
}

func (s *DescribeAppsForProviderResponseBodyAppsApp) SetModifiedTime(v string) *DescribeAppsForProviderResponseBodyAppsApp {
	s.ModifiedTime = &v
	return s
}

type DescribeAppsForProviderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsForProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppsForProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsForProviderResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsForProviderResponse) SetHeaders(v map[string]*string) *DescribeAppsForProviderResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsForProviderResponse) SetStatusCode(v int32) *DescribeAppsForProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsForProviderResponse) SetBody(v *DescribeAppsForProviderResponseBody) *DescribeAppsForProviderResponse {
	s.Body = v
	return s
}

type DescribeAvailableVipsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Vip           *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
}

func (s DescribeAvailableVipsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableVipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableVipsRequest) SetSecurityToken(v string) *DescribeAvailableVipsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAvailableVipsRequest) SetVip(v string) *DescribeAvailableVipsRequest {
	s.Vip = &v
	return s
}

type DescribeAvailableVipsResponseBody struct {
	AvailableVips *DescribeAvailableVipsResponseBodyAvailableVips `json:"AvailableVips,omitempty" xml:"AvailableVips,omitempty" type:"Struct"`
}

func (s DescribeAvailableVipsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableVipsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableVipsResponseBody) SetAvailableVips(v *DescribeAvailableVipsResponseBodyAvailableVips) *DescribeAvailableVipsResponseBody {
	s.AvailableVips = v
	return s
}

type DescribeAvailableVipsResponseBodyAvailableVips struct {
	AvailableVip []*string `json:"AvailableVip,omitempty" xml:"AvailableVip,omitempty" type:"Repeated"`
}

func (s DescribeAvailableVipsResponseBodyAvailableVips) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableVipsResponseBodyAvailableVips) GoString() string {
	return s.String()
}

func (s *DescribeAvailableVipsResponseBodyAvailableVips) SetAvailableVip(v []*string) *DescribeAvailableVipsResponseBodyAvailableVips {
	s.AvailableVip = v
	return s
}

type DescribeAvailableVipsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableVipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableVipsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableVipsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableVipsResponse) SetHeaders(v map[string]*string) *DescribeAvailableVipsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableVipsResponse) SetStatusCode(v int32) *DescribeAvailableVipsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableVipsResponse) SetBody(v *DescribeAvailableVipsResponseBody) *DescribeAvailableVipsResponse {
	s.Body = v
	return s
}

type DescribeBidByUserIdForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeBidByUserIdForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBidByUserIdForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeBidByUserIdForInnerRequest) SetAliUid(v int64) *DescribeBidByUserIdForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeBidByUserIdForInnerRequest) SetSecurityToken(v string) *DescribeBidByUserIdForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DescribeBidByUserIdForInnerResponseBody struct {
	Bid       *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBidByUserIdForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBidByUserIdForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBidByUserIdForInnerResponseBody) SetBid(v string) *DescribeBidByUserIdForInnerResponseBody {
	s.Bid = &v
	return s
}

func (s *DescribeBidByUserIdForInnerResponseBody) SetRequestId(v string) *DescribeBidByUserIdForInnerResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBidByUserIdForInnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBidByUserIdForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBidByUserIdForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBidByUserIdForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeBidByUserIdForInnerResponse) SetHeaders(v map[string]*string) *DescribeBidByUserIdForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeBidByUserIdForInnerResponse) SetStatusCode(v int32) *DescribeBidByUserIdForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBidByUserIdForInnerResponse) SetBody(v *DescribeBidByUserIdForInnerResponseBody) *DescribeBidByUserIdForInnerResponse {
	s.Body = v
	return s
}

type DescribeBlackListsRequest struct {
	BlackType     *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeBlackListsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackListsRequest) SetBlackType(v string) *DescribeBlackListsRequest {
	s.BlackType = &v
	return s
}

func (s *DescribeBlackListsRequest) SetPageNumber(v int32) *DescribeBlackListsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBlackListsRequest) SetPageSize(v int32) *DescribeBlackListsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBlackListsRequest) SetSecurityToken(v string) *DescribeBlackListsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeBlackListsResponseBody struct {
	BlackLists *DescribeBlackListsResponseBodyBlackLists `json:"BlackLists,omitempty" xml:"BlackLists,omitempty" type:"Struct"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBlackListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlackListsResponseBody) SetBlackLists(v *DescribeBlackListsResponseBodyBlackLists) *DescribeBlackListsResponseBody {
	s.BlackLists = v
	return s
}

func (s *DescribeBlackListsResponseBody) SetPageNumber(v int32) *DescribeBlackListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeBlackListsResponseBody) SetPageSize(v int32) *DescribeBlackListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeBlackListsResponseBody) SetRequestId(v string) *DescribeBlackListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlackListsResponseBody) SetTotalCount(v int32) *DescribeBlackListsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeBlackListsResponseBodyBlackLists struct {
	BlackList []*DescribeBlackListsResponseBodyBlackListsBlackList `json:"BlackList,omitempty" xml:"BlackList,omitempty" type:"Repeated"`
}

func (s DescribeBlackListsResponseBodyBlackLists) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListsResponseBodyBlackLists) GoString() string {
	return s.String()
}

func (s *DescribeBlackListsResponseBodyBlackLists) SetBlackList(v []*DescribeBlackListsResponseBodyBlackListsBlackList) *DescribeBlackListsResponseBodyBlackLists {
	s.BlackList = v
	return s
}

type DescribeBlackListsResponseBodyBlackListsBlackList struct {
	BlackContent *string `json:"BlackContent,omitempty" xml:"BlackContent,omitempty"`
	BlackType    *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
}

func (s DescribeBlackListsResponseBodyBlackListsBlackList) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListsResponseBodyBlackListsBlackList) GoString() string {
	return s.String()
}

func (s *DescribeBlackListsResponseBodyBlackListsBlackList) SetBlackContent(v string) *DescribeBlackListsResponseBodyBlackListsBlackList {
	s.BlackContent = &v
	return s
}

func (s *DescribeBlackListsResponseBodyBlackListsBlackList) SetBlackType(v string) *DescribeBlackListsResponseBodyBlackListsBlackList {
	s.BlackType = &v
	return s
}

func (s *DescribeBlackListsResponseBodyBlackListsBlackList) SetCreateTime(v string) *DescribeBlackListsResponseBodyBlackListsBlackList {
	s.CreateTime = &v
	return s
}

func (s *DescribeBlackListsResponseBodyBlackListsBlackList) SetDescription(v string) *DescribeBlackListsResponseBodyBlackListsBlackList {
	s.Description = &v
	return s
}

func (s *DescribeBlackListsResponseBodyBlackListsBlackList) SetModifiedTime(v string) *DescribeBlackListsResponseBodyBlackListsBlackList {
	s.ModifiedTime = &v
	return s
}

type DescribeBlackListsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBlackListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBlackListsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBlackListsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBlackListsResponse) SetHeaders(v map[string]*string) *DescribeBlackListsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBlackListsResponse) SetStatusCode(v int32) *DescribeBlackListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBlackListsResponse) SetBody(v *DescribeBlackListsResponseBody) *DescribeBlackListsResponse {
	s.Body = v
	return s
}

type DescribeDeployedApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeDeployedApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiRequest) SetApiId(v string) *DescribeDeployedApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApiRequest) SetGroupId(v string) *DescribeDeployedApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApiRequest) SetSecurityToken(v string) *DescribeDeployedApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDeployedApiRequest) SetStageName(v string) *DescribeDeployedApiRequest {
	s.StageName = &v
	return s
}

type DescribeDeployedApiResponseBody struct {
	ApiId                 *string                                               `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName               *string                                               `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType              *string                                               `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BodyFormat            *string                                               `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	ConstantParameters    *DescribeDeployedApiResponseBodyConstantParameters    `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	DeployedTime          *string                                               `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	ErrorCodeSamples      *DescribeDeployedApiResponseBodyErrorCodeSamples      `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FunctionComputeConfig *DescribeDeployedApiResponseBodyFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	GroupId               *string                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName             *string                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpMethod            *string                                               `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol          *string                                               `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Path                  *string                                               `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters        *DescribeDeployedApiResponseBodyPathParameters        `json:"PathParameters,omitempty" xml:"PathParameters,omitempty" type:"Struct"`
	PostBodyDescription   *string                                               `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType          *string                                               `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RegionId              *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestBody           *DescribeDeployedApiResponseBodyRequestBody           `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	RequestHeaders        *DescribeDeployedApiResponseBodyRequestHeaders        `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Struct"`
	RequestId             *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestQueries        *DescribeDeployedApiResponseBodyRequestQueries        `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty" type:"Struct"`
	ResultSample          *string                                               `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType            *string                                               `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceAddress        *string                                               `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceFCEnable       *string                                               `json:"ServiceFCEnable,omitempty" xml:"ServiceFCEnable,omitempty"`
	ServiceProtocol       *string                                               `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout        *int32                                                `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable      *string                                               `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	StageName             *string                                               `json:"StageName,omitempty" xml:"StageName,omitempty"`
	SystemParameters      *DescribeDeployedApiResponseBodySystemParameters      `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility            *string                                               `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VpcName               *string                                               `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeDeployedApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBody) SetApiId(v string) *DescribeDeployedApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetApiName(v string) *DescribeDeployedApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetAuthType(v string) *DescribeDeployedApiResponseBody {
	s.AuthType = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetBodyFormat(v string) *DescribeDeployedApiResponseBody {
	s.BodyFormat = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetConstantParameters(v *DescribeDeployedApiResponseBodyConstantParameters) *DescribeDeployedApiResponseBody {
	s.ConstantParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetDeployedTime(v string) *DescribeDeployedApiResponseBody {
	s.DeployedTime = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetErrorCodeSamples(v *DescribeDeployedApiResponseBodyErrorCodeSamples) *DescribeDeployedApiResponseBody {
	s.ErrorCodeSamples = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetFunctionComputeConfig(v *DescribeDeployedApiResponseBodyFunctionComputeConfig) *DescribeDeployedApiResponseBody {
	s.FunctionComputeConfig = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetGroupId(v string) *DescribeDeployedApiResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetGroupName(v string) *DescribeDeployedApiResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetHttpMethod(v string) *DescribeDeployedApiResponseBody {
	s.HttpMethod = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetHttpProtocol(v string) *DescribeDeployedApiResponseBody {
	s.HttpProtocol = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetPath(v string) *DescribeDeployedApiResponseBody {
	s.Path = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetPathParameters(v *DescribeDeployedApiResponseBodyPathParameters) *DescribeDeployedApiResponseBody {
	s.PathParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetPostBodyDescription(v string) *DescribeDeployedApiResponseBody {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetPostBodyType(v string) *DescribeDeployedApiResponseBody {
	s.PostBodyType = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRegionId(v string) *DescribeDeployedApiResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestBody(v *DescribeDeployedApiResponseBodyRequestBody) *DescribeDeployedApiResponseBody {
	s.RequestBody = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestHeaders(v *DescribeDeployedApiResponseBodyRequestHeaders) *DescribeDeployedApiResponseBody {
	s.RequestHeaders = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestId(v string) *DescribeDeployedApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestQueries(v *DescribeDeployedApiResponseBodyRequestQueries) *DescribeDeployedApiResponseBody {
	s.RequestQueries = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetResultSample(v string) *DescribeDeployedApiResponseBody {
	s.ResultSample = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetResultType(v string) *DescribeDeployedApiResponseBody {
	s.ResultType = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceAddress(v string) *DescribeDeployedApiResponseBody {
	s.ServiceAddress = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceFCEnable(v string) *DescribeDeployedApiResponseBody {
	s.ServiceFCEnable = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceProtocol(v string) *DescribeDeployedApiResponseBody {
	s.ServiceProtocol = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceTimeout(v int32) *DescribeDeployedApiResponseBody {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceVpcEnable(v string) *DescribeDeployedApiResponseBody {
	s.ServiceVpcEnable = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetStageName(v string) *DescribeDeployedApiResponseBody {
	s.StageName = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetSystemParameters(v *DescribeDeployedApiResponseBodySystemParameters) *DescribeDeployedApiResponseBody {
	s.SystemParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetVisibility(v string) *DescribeDeployedApiResponseBody {
	s.Visibility = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetVpcName(v string) *DescribeDeployedApiResponseBody {
	s.VpcName = &v
	return s
}

type DescribeDeployedApiResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeDeployedApiResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyConstantParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyConstantParameters) SetConstantParameter(v []*DescribeDeployedApiResponseBodyConstantParametersConstantParameter) *DescribeDeployedApiResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

type DescribeDeployedApiResponseBodyConstantParametersConstantParameter struct {
	ConstantValue        *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyConstantParametersConstantParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) SetConstantValue(v string) *DescribeDeployedApiResponseBodyConstantParametersConstantParameter {
	s.ConstantValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) SetDescription(v string) *DescribeDeployedApiResponseBodyConstantParametersConstantParameter {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) SetLocation(v string) *DescribeDeployedApiResponseBodyConstantParametersConstantParameter {
	s.Location = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyConstantParametersConstantParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamples) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeDeployedApiResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

type DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) SetCode(v string) *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Code = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) SetDescription(v string) *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) SetMessage(v string) *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Message = &v
	return s
}

type DescribeDeployedApiResponseBodyFunctionComputeConfig struct {
	FcRegionId   *string `json:"FcRegionId,omitempty" xml:"FcRegionId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	RoleArn      *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyFunctionComputeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyFunctionComputeConfig) SetFcRegionId(v string) *DescribeDeployedApiResponseBodyFunctionComputeConfig {
	s.FcRegionId = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyFunctionComputeConfig) SetFunctionName(v string) *DescribeDeployedApiResponseBodyFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyFunctionComputeConfig) SetRoleArn(v string) *DescribeDeployedApiResponseBodyFunctionComputeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyFunctionComputeConfig) SetServiceName(v string) *DescribeDeployedApiResponseBodyFunctionComputeConfig {
	s.ServiceName = &v
	return s
}

type DescribeDeployedApiResponseBodyPathParameters struct {
	PathParameter []*DescribeDeployedApiResponseBodyPathParametersPathParameter `json:"PathParameter,omitempty" xml:"PathParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyPathParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyPathParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyPathParameters) SetPathParameter(v []*DescribeDeployedApiResponseBodyPathParametersPathParameter) *DescribeDeployedApiResponseBodyPathParameters {
	s.PathParameter = v
	return s
}

type DescribeDeployedApiResponseBodyPathParametersPathParameter struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyPathParametersPathParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyPathParametersPathParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyPathParametersPathParameter) SetApiParameterName(v string) *DescribeDeployedApiResponseBodyPathParametersPathParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyPathParametersPathParameter) SetDemoValue(v string) *DescribeDeployedApiResponseBodyPathParametersPathParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyPathParametersPathParameter) SetDescription(v string) *DescribeDeployedApiResponseBodyPathParametersPathParameter {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyPathParametersPathParameter) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyPathParametersPathParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponseBodyRequestBody struct {
	RequestParam []*DescribeDeployedApiResponseBodyRequestBodyRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyRequestBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestBody) SetRequestParam(v []*DescribeDeployedApiResponseBodyRequestBodyRequestParam) *DescribeDeployedApiResponseBodyRequestBody {
	s.RequestParam = v
	return s
}

type DescribeDeployedApiResponseBodyRequestBodyRequestParam struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder             *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow              *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue            *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme           *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue             *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength            *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue             *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression    *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required             *string `json:"Required,omitempty" xml:"Required,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyRequestBodyRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestBodyRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetApiParameterName(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetDefaultValue(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetDemoValue(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetDescription(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetDocOrder(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetDocShow(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetEnumValue(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetJsonScheme(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetMaxLength(v int64) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetMaxValue(v int64) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetMinLength(v int64) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetMinValue(v int64) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetParameterType(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetRegularExpression(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetRequired(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.Required = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestBodyRequestParam) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyRequestBodyRequestParam {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponseBodyRequestHeaders struct {
	RequestParam []*DescribeDeployedApiResponseBodyRequestHeadersRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestHeaders) SetRequestParam(v []*DescribeDeployedApiResponseBodyRequestHeadersRequestParam) *DescribeDeployedApiResponseBodyRequestHeaders {
	s.RequestParam = v
	return s
}

type DescribeDeployedApiResponseBodyRequestHeadersRequestParam struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder             *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow              *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue            *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme           *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue             *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength            *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue             *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression    *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required             *string `json:"Required,omitempty" xml:"Required,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyRequestHeadersRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestHeadersRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetApiParameterName(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetDefaultValue(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetDemoValue(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetDescription(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetDocOrder(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetDocShow(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetEnumValue(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetJsonScheme(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetMaxLength(v int64) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetMaxValue(v int64) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetMinLength(v int64) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetMinValue(v int64) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetParameterType(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetRegularExpression(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetRequired(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.Required = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestHeadersRequestParam) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyRequestHeadersRequestParam {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponseBodyRequestQueries struct {
	RequestParam []*DescribeDeployedApiResponseBodyRequestQueriesRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyRequestQueries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestQueries) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestQueries) SetRequestParam(v []*DescribeDeployedApiResponseBodyRequestQueriesRequestParam) *DescribeDeployedApiResponseBodyRequestQueries {
	s.RequestParam = v
	return s
}

type DescribeDeployedApiResponseBodyRequestQueriesRequestParam struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder             *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow              *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue            *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme           *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue             *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength            *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue             *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression    *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required             *string `json:"Required,omitempty" xml:"Required,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyRequestQueriesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestQueriesRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetApiParameterName(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetDefaultValue(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetDemoValue(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetDescription(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetDocOrder(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetDocShow(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetEnumValue(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetJsonScheme(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetMaxLength(v int64) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetMaxValue(v int64) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetMinLength(v int64) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetMinValue(v int64) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetParameterType(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetRegularExpression(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetRequired(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.Required = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestQueriesRequestParam) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyRequestQueriesRequestParam {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponseBodySystemParameters struct {
	SystemParameter []*DescribeDeployedApiResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodySystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodySystemParameters) SetSystemParameter(v []*DescribeDeployedApiResponseBodySystemParametersSystemParameter) *DescribeDeployedApiResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

type DescribeDeployedApiResponseBodySystemParametersSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodySystemParametersSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) SetDemoValue(v string) *DescribeDeployedApiResponseBodySystemParametersSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) SetDescription(v string) *DescribeDeployedApiResponseBodySystemParametersSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) SetLocation(v string) *DescribeDeployedApiResponseBodySystemParametersSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) SetParameterName(v string) *DescribeDeployedApiResponseBodySystemParametersSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodySystemParametersSystemParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeployedApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeployedApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponse) SetHeaders(v map[string]*string) *DescribeDeployedApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployedApiResponse) SetStatusCode(v int32) *DescribeDeployedApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeployedApiResponse) SetBody(v *DescribeDeployedApiResponseBody) *DescribeDeployedApiResponse {
	s.Body = v
	return s
}

type DescribeDeployedApisRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeDeployedApisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisRequest) SetApiId(v string) *DescribeDeployedApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetApiName(v string) *DescribeDeployedApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetGroupId(v string) *DescribeDeployedApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetPageNumber(v int32) *DescribeDeployedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetPageSize(v int32) *DescribeDeployedApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetSecurityToken(v string) *DescribeDeployedApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDeployedApisRequest) SetStageName(v string) *DescribeDeployedApisRequest {
	s.StageName = &v
	return s
}

type DescribeDeployedApisResponseBody struct {
	ApiInfos   *DescribeDeployedApisResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeployedApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBody) SetApiInfos(v *DescribeDeployedApisResponseBodyApiInfos) *DescribeDeployedApisResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetPageNumber(v int32) *DescribeDeployedApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetPageSize(v int32) *DescribeDeployedApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetRequestId(v string) *DescribeDeployedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeployedApisResponseBody) SetTotalCount(v int32) *DescribeDeployedApisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDeployedApisResponseBodyApiInfos struct {
	ApiInfo []*DescribeDeployedApisResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApisResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBodyApiInfos) SetApiInfo(v []*DescribeDeployedApisResponseBodyApiInfosApiInfo) *DescribeDeployedApisResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeDeployedApisResponseBodyApiInfosApiInfo struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName    *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility   *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeDeployedApisResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetDeployedTime(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.DeployedTime = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeDeployedApisResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeDeployedApisResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeployedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeployedApisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponse) SetHeaders(v map[string]*string) *DescribeDeployedApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployedApisResponse) SetStatusCode(v int32) *DescribeDeployedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeployedApisResponse) SetBody(v *DescribeDeployedApisResponseBody) *DescribeDeployedApisResponse {
	s.Body = v
	return s
}

type DescribeDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRequest) SetDomainName(v string) *DescribeDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainRequest) SetGroupId(v string) *DescribeDomainRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainRequest) SetSecurityToken(v string) *DescribeDomainRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDomainResponseBody struct {
	CertificateBody      *string `json:"CertificateBody,omitempty" xml:"CertificateBody,omitempty"`
	CertificateId        *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName      *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
	DomainStatus         *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PrivateKey           *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain            *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s DescribeDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponseBody) SetCertificateBody(v string) *DescribeDomainResponseBody {
	s.CertificateBody = &v
	return s
}

func (s *DescribeDomainResponseBody) SetCertificateId(v string) *DescribeDomainResponseBody {
	s.CertificateId = &v
	return s
}

func (s *DescribeDomainResponseBody) SetCertificateName(v string) *DescribeDomainResponseBody {
	s.CertificateName = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainName(v string) *DescribeDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainNameResolution(v string) *DescribeDomainResponseBody {
	s.DomainNameResolution = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainStatus(v string) *DescribeDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *DescribeDomainResponseBody) SetGroupId(v string) *DescribeDomainResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainResponseBody) SetPrivateKey(v string) *DescribeDomainResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *DescribeDomainResponseBody) SetRequestId(v string) *DescribeDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResponseBody) SetSubDomain(v string) *DescribeDomainResponseBody {
	s.SubDomain = &v
	return s
}

type DescribeDomainResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResponse) SetHeaders(v map[string]*string) *DescribeDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResponse) SetStatusCode(v int32) *DescribeDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResponse) SetBody(v *DescribeDomainResponseBody) *DescribeDomainResponse {
	s.Body = v
	return s
}

type DescribeDomainResolutionRequest struct {
	DomainNames   *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDomainResolutionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolutionRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolutionRequest) SetDomainNames(v string) *DescribeDomainResolutionRequest {
	s.DomainNames = &v
	return s
}

func (s *DescribeDomainResolutionRequest) SetGroupId(v string) *DescribeDomainResolutionRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainResolutionRequest) SetSecurityToken(v string) *DescribeDomainResolutionRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDomainResolutionResponseBody struct {
	DomainResolutions *DescribeDomainResolutionResponseBodyDomainResolutions `json:"DomainResolutions,omitempty" xml:"DomainResolutions,omitempty" type:"Struct"`
	GroupId           *string                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainResolutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolutionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolutionResponseBody) SetDomainResolutions(v *DescribeDomainResolutionResponseBodyDomainResolutions) *DescribeDomainResolutionResponseBody {
	s.DomainResolutions = v
	return s
}

func (s *DescribeDomainResolutionResponseBody) SetGroupId(v string) *DescribeDomainResolutionResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDomainResolutionResponseBody) SetRequestId(v string) *DescribeDomainResolutionResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDomainResolutionResponseBodyDomainResolutions struct {
	DomainResolution []*DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution `json:"DomainResolution,omitempty" xml:"DomainResolution,omitempty" type:"Repeated"`
}

func (s DescribeDomainResolutionResponseBodyDomainResolutions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolutionResponseBodyDomainResolutions) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolutionResponseBodyDomainResolutions) SetDomainResolution(v []*DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution) *DescribeDomainResolutionResponseBodyDomainResolutions {
	s.DomainResolution = v
	return s
}

type DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution struct {
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
}

func (s DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution) SetDomainName(v string) *DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution) SetDomainNameResolution(v string) *DescribeDomainResolutionResponseBodyDomainResolutionsDomainResolution {
	s.DomainNameResolution = &v
	return s
}

type DescribeDomainResolutionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainResolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainResolutionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDomainResolutionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainResolutionResponse) SetHeaders(v map[string]*string) *DescribeDomainResolutionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainResolutionResponse) SetStatusCode(v int32) *DescribeDomainResolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainResolutionResponse) SetBody(v *DescribeDomainResolutionResponseBody) *DescribeDomainResolutionResponse {
	s.Body = v
	return s
}

type DescribeHistoryApiRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeHistoryApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiRequest) SetApiId(v string) *DescribeHistoryApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApiRequest) SetGroupId(v string) *DescribeHistoryApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApiRequest) SetHistoryVersion(v string) *DescribeHistoryApiRequest {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeHistoryApiRequest) SetSecurityToken(v string) *DescribeHistoryApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryApiRequest) SetStageName(v string) *DescribeHistoryApiRequest {
	s.StageName = &v
	return s
}

type DescribeHistoryApiResponseBody struct {
	ApiId                   *string                                              `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                 *string                                              `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType                *string                                              `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BodyFormat              *string                                              `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	ConstantParameters      *DescribeHistoryApiResponseBodyConstantParameters    `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	DeployedTime            *string                                              `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description             *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples        *DescribeHistoryApiResponseBodyErrorCodeSamples      `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FunctionComputeConfig   *DescribeHistoryApiResponseBodyFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	GroupId                 *string                                              `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName               *string                                              `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HistoryVersion          *string                                              `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	HttpMethod              *string                                              `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol            *string                                              `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Mock                    *string                                              `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult              *string                                              `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	OriginResultDescription *string                                              `json:"OriginResultDescription,omitempty" xml:"OriginResultDescription,omitempty"`
	Path                    *string                                              `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters          *DescribeHistoryApiResponseBodyPathParameters        `json:"PathParameters,omitempty" xml:"PathParameters,omitempty" type:"Struct"`
	PostBodyDescription     *string                                              `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType            *string                                              `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RegionId                *string                                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestBody             *DescribeHistoryApiResponseBodyRequestBody           `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	RequestHeaders          *DescribeHistoryApiResponseBodyRequestHeaders        `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Struct"`
	RequestId               *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestQueries          *DescribeHistoryApiResponseBodyRequestQueries        `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty" type:"Struct"`
	ResultSample            *string                                              `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType              *string                                              `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceAddress          *string                                              `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceFCEnable         *string                                              `json:"ServiceFCEnable,omitempty" xml:"ServiceFCEnable,omitempty"`
	ServiceProtocol         *string                                              `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout          *int32                                               `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable        *string                                              `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	Status                  *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemParameters        *DescribeHistoryApiResponseBodySystemParameters      `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility              *string                                              `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VpcName                 *string                                              `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeHistoryApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBody) SetApiId(v string) *DescribeHistoryApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetApiName(v string) *DescribeHistoryApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetAuthType(v string) *DescribeHistoryApiResponseBody {
	s.AuthType = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetBodyFormat(v string) *DescribeHistoryApiResponseBody {
	s.BodyFormat = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetConstantParameters(v *DescribeHistoryApiResponseBodyConstantParameters) *DescribeHistoryApiResponseBody {
	s.ConstantParameters = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetDeployedTime(v string) *DescribeHistoryApiResponseBody {
	s.DeployedTime = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetDescription(v string) *DescribeHistoryApiResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetErrorCodeSamples(v *DescribeHistoryApiResponseBodyErrorCodeSamples) *DescribeHistoryApiResponseBody {
	s.ErrorCodeSamples = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetFunctionComputeConfig(v *DescribeHistoryApiResponseBodyFunctionComputeConfig) *DescribeHistoryApiResponseBody {
	s.FunctionComputeConfig = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetGroupId(v string) *DescribeHistoryApiResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetGroupName(v string) *DescribeHistoryApiResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetHistoryVersion(v string) *DescribeHistoryApiResponseBody {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetHttpMethod(v string) *DescribeHistoryApiResponseBody {
	s.HttpMethod = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetHttpProtocol(v string) *DescribeHistoryApiResponseBody {
	s.HttpProtocol = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetMock(v string) *DescribeHistoryApiResponseBody {
	s.Mock = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetMockResult(v string) *DescribeHistoryApiResponseBody {
	s.MockResult = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetOriginResultDescription(v string) *DescribeHistoryApiResponseBody {
	s.OriginResultDescription = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetPath(v string) *DescribeHistoryApiResponseBody {
	s.Path = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetPathParameters(v *DescribeHistoryApiResponseBodyPathParameters) *DescribeHistoryApiResponseBody {
	s.PathParameters = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetPostBodyDescription(v string) *DescribeHistoryApiResponseBody {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetPostBodyType(v string) *DescribeHistoryApiResponseBody {
	s.PostBodyType = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetRegionId(v string) *DescribeHistoryApiResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetRequestBody(v *DescribeHistoryApiResponseBodyRequestBody) *DescribeHistoryApiResponseBody {
	s.RequestBody = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetRequestHeaders(v *DescribeHistoryApiResponseBodyRequestHeaders) *DescribeHistoryApiResponseBody {
	s.RequestHeaders = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetRequestId(v string) *DescribeHistoryApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetRequestQueries(v *DescribeHistoryApiResponseBodyRequestQueries) *DescribeHistoryApiResponseBody {
	s.RequestQueries = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetResultSample(v string) *DescribeHistoryApiResponseBody {
	s.ResultSample = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetResultType(v string) *DescribeHistoryApiResponseBody {
	s.ResultType = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetServiceAddress(v string) *DescribeHistoryApiResponseBody {
	s.ServiceAddress = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetServiceFCEnable(v string) *DescribeHistoryApiResponseBody {
	s.ServiceFCEnable = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetServiceProtocol(v string) *DescribeHistoryApiResponseBody {
	s.ServiceProtocol = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetServiceTimeout(v int32) *DescribeHistoryApiResponseBody {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetServiceVpcEnable(v string) *DescribeHistoryApiResponseBody {
	s.ServiceVpcEnable = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetStatus(v string) *DescribeHistoryApiResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetSystemParameters(v *DescribeHistoryApiResponseBodySystemParameters) *DescribeHistoryApiResponseBody {
	s.SystemParameters = v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetVisibility(v string) *DescribeHistoryApiResponseBody {
	s.Visibility = &v
	return s
}

func (s *DescribeHistoryApiResponseBody) SetVpcName(v string) *DescribeHistoryApiResponseBody {
	s.VpcName = &v
	return s
}

type DescribeHistoryApiResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeHistoryApiResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyConstantParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyConstantParameters) SetConstantParameter(v []*DescribeHistoryApiResponseBodyConstantParametersConstantParameter) *DescribeHistoryApiResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

type DescribeHistoryApiResponseBodyConstantParametersConstantParameter struct {
	ConstantValue        *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyConstantParametersConstantParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyConstantParametersConstantParameter) SetConstantValue(v string) *DescribeHistoryApiResponseBodyConstantParametersConstantParameter {
	s.ConstantValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyConstantParametersConstantParameter) SetDescription(v string) *DescribeHistoryApiResponseBodyConstantParametersConstantParameter {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyConstantParametersConstantParameter) SetLocation(v string) *DescribeHistoryApiResponseBodyConstantParametersConstantParameter {
	s.Location = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyConstantParametersConstantParameter) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodyConstantParametersConstantParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeHistoryApiResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyErrorCodeSamples) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeHistoryApiResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

type DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample) SetCode(v string) *DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Code = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample) SetDescription(v string) *DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample) SetMessage(v string) *DescribeHistoryApiResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Message = &v
	return s
}

type DescribeHistoryApiResponseBodyFunctionComputeConfig struct {
	FcRegionId   *string `json:"FcRegionId,omitempty" xml:"FcRegionId,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	RoleArn      *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName  *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyFunctionComputeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyFunctionComputeConfig) SetFcRegionId(v string) *DescribeHistoryApiResponseBodyFunctionComputeConfig {
	s.FcRegionId = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyFunctionComputeConfig) SetFunctionName(v string) *DescribeHistoryApiResponseBodyFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyFunctionComputeConfig) SetRoleArn(v string) *DescribeHistoryApiResponseBodyFunctionComputeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyFunctionComputeConfig) SetServiceName(v string) *DescribeHistoryApiResponseBodyFunctionComputeConfig {
	s.ServiceName = &v
	return s
}

type DescribeHistoryApiResponseBodyPathParameters struct {
	PathParameter []*DescribeHistoryApiResponseBodyPathParametersPathParameter `json:"PathParameter,omitempty" xml:"PathParameter,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyPathParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyPathParameters) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyPathParameters) SetPathParameter(v []*DescribeHistoryApiResponseBodyPathParametersPathParameter) *DescribeHistoryApiResponseBodyPathParameters {
	s.PathParameter = v
	return s
}

type DescribeHistoryApiResponseBodyPathParametersPathParameter struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyPathParametersPathParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyPathParametersPathParameter) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyPathParametersPathParameter) SetApiParameterName(v string) *DescribeHistoryApiResponseBodyPathParametersPathParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyPathParametersPathParameter) SetDemoValue(v string) *DescribeHistoryApiResponseBodyPathParametersPathParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyPathParametersPathParameter) SetDescription(v string) *DescribeHistoryApiResponseBodyPathParametersPathParameter {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyPathParametersPathParameter) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodyPathParametersPathParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeHistoryApiResponseBodyRequestBody struct {
	RequestParam []*DescribeHistoryApiResponseBodyRequestBodyRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyRequestBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyRequestBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyRequestBody) SetRequestParam(v []*DescribeHistoryApiResponseBodyRequestBodyRequestParam) *DescribeHistoryApiResponseBodyRequestBody {
	s.RequestParam = v
	return s
}

type DescribeHistoryApiResponseBodyRequestBodyRequestParam struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder             *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow              *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue            *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme           *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue             *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength            *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue             *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression    *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required             *string `json:"Required,omitempty" xml:"Required,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyRequestBodyRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyRequestBodyRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetApiParameterName(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetDefaultValue(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetDemoValue(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetDescription(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetDocOrder(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetDocShow(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetEnumValue(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetJsonScheme(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetMaxLength(v int64) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetMaxValue(v int64) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetMinLength(v int64) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetMinValue(v int64) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetParameterType(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetRegularExpression(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetRequired(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.Required = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestBodyRequestParam) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodyRequestBodyRequestParam {
	s.ServiceParameterName = &v
	return s
}

type DescribeHistoryApiResponseBodyRequestHeaders struct {
	RequestParam []*DescribeHistoryApiResponseBodyRequestHeadersRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyRequestHeaders) SetRequestParam(v []*DescribeHistoryApiResponseBodyRequestHeadersRequestParam) *DescribeHistoryApiResponseBodyRequestHeaders {
	s.RequestParam = v
	return s
}

type DescribeHistoryApiResponseBodyRequestHeadersRequestParam struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder             *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow              *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue            *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme           *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue             *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength            *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue             *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression    *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required             *string `json:"Required,omitempty" xml:"Required,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyRequestHeadersRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyRequestHeadersRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetApiParameterName(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetDefaultValue(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetDemoValue(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetDescription(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetDocOrder(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetDocShow(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetEnumValue(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetJsonScheme(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetMaxLength(v int64) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetMaxValue(v int64) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetMinLength(v int64) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetMinValue(v int64) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetParameterType(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetRegularExpression(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetRequired(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.Required = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestHeadersRequestParam) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodyRequestHeadersRequestParam {
	s.ServiceParameterName = &v
	return s
}

type DescribeHistoryApiResponseBodyRequestQueries struct {
	RequestParam []*DescribeHistoryApiResponseBodyRequestQueriesRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyRequestQueries) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyRequestQueries) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyRequestQueries) SetRequestParam(v []*DescribeHistoryApiResponseBodyRequestQueriesRequestParam) *DescribeHistoryApiResponseBodyRequestQueries {
	s.RequestParam = v
	return s
}

type DescribeHistoryApiResponseBodyRequestQueriesRequestParam struct {
	ApiParameterName     *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue         *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder             *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow              *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue            *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme           *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength            *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue             *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength            *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue             *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression    *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required             *string `json:"Required,omitempty" xml:"Required,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyRequestQueriesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyRequestQueriesRequestParam) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetApiParameterName(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetDefaultValue(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetDemoValue(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetDescription(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetDocOrder(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetDocShow(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetEnumValue(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetJsonScheme(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetMaxLength(v int64) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetMaxValue(v int64) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetMinLength(v int64) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetMinValue(v int64) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetParameterType(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetRegularExpression(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetRequired(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.Required = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyRequestQueriesRequestParam) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodyRequestQueriesRequestParam {
	s.ServiceParameterName = &v
	return s
}

type DescribeHistoryApiResponseBodySystemParameters struct {
	SystemParameter []*DescribeHistoryApiResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodySystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodySystemParameters) SetSystemParameter(v []*DescribeHistoryApiResponseBodySystemParametersSystemParameter) *DescribeHistoryApiResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

type DescribeHistoryApiResponseBodySystemParametersSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodySystemParametersSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodySystemParametersSystemParameter) SetDemoValue(v string) *DescribeHistoryApiResponseBodySystemParametersSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodySystemParametersSystemParameter) SetDescription(v string) *DescribeHistoryApiResponseBodySystemParametersSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodySystemParametersSystemParameter) SetLocation(v string) *DescribeHistoryApiResponseBodySystemParametersSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeHistoryApiResponseBodySystemParametersSystemParameter) SetParameterName(v string) *DescribeHistoryApiResponseBodySystemParametersSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodySystemParametersSystemParameter) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodySystemParametersSystemParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeHistoryApiResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponse) SetHeaders(v map[string]*string) *DescribeHistoryApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryApiResponse) SetStatusCode(v int32) *DescribeHistoryApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryApiResponse) SetBody(v *DescribeHistoryApiResponseBody) *DescribeHistoryApiResponse {
	s.Body = v
	return s
}

type DescribeHistoryApisRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeHistoryApisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisRequest) SetApiId(v string) *DescribeHistoryApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetApiName(v string) *DescribeHistoryApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetGroupId(v string) *DescribeHistoryApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetPageNumber(v int32) *DescribeHistoryApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetPageSize(v int32) *DescribeHistoryApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetSecurityToken(v string) *DescribeHistoryApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetStageName(v string) *DescribeHistoryApisRequest {
	s.StageName = &v
	return s
}

type DescribeHistoryApisResponseBody struct {
	ApiInfos   *DescribeHistoryApisResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBody) SetApiInfos(v *DescribeHistoryApisResponseBodyApiInfos) *DescribeHistoryApisResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetPageNumber(v int32) *DescribeHistoryApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetPageSize(v int32) *DescribeHistoryApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetRequestId(v string) *DescribeHistoryApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHistoryApisResponseBody) SetTotalCount(v int32) *DescribeHistoryApisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeHistoryApisResponseBodyApiInfos struct {
	ApiInfo []*DescribeHistoryApisResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApisResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBodyApiInfos) SetApiInfo(v []*DescribeHistoryApisResponseBodyApiInfosApiInfo) *DescribeHistoryApisResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeHistoryApisResponseBodyApiInfosApiInfo struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName        *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	DeployedTime   *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHistoryApisResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetDeployedTime(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.DeployedTime = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetHistoryVersion(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiInfosApiInfo) SetStatus(v string) *DescribeHistoryApisResponseBodyApiInfosApiInfo {
	s.Status = &v
	return s
}

type DescribeHistoryApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryApisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponse) SetHeaders(v map[string]*string) *DescribeHistoryApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryApisResponse) SetStatusCode(v int32) *DescribeHistoryApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryApisResponse) SetBody(v *DescribeHistoryApisResponseBody) *DescribeHistoryApisResponse {
	s.Body = v
	return s
}

type DescribeModelsForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelId       *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName     *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeModelsForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelsForInnerRequest) SetAliUid(v int64) *DescribeModelsForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeModelsForInnerRequest) SetGroupId(v string) *DescribeModelsForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeModelsForInnerRequest) SetModelId(v string) *DescribeModelsForInnerRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeModelsForInnerRequest) SetModelName(v string) *DescribeModelsForInnerRequest {
	s.ModelName = &v
	return s
}

func (s *DescribeModelsForInnerRequest) SetPageNumber(v int32) *DescribeModelsForInnerRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelsForInnerRequest) SetPageSize(v int32) *DescribeModelsForInnerRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeModelsForInnerRequest) SetSecurityToken(v string) *DescribeModelsForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DescribeModelsForInnerResponseBody struct {
	ModelDetails *DescribeModelsForInnerResponseBodyModelDetails `json:"ModelDetails,omitempty" xml:"ModelDetails,omitempty" type:"Struct"`
	PageNumber   *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeModelsForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelsForInnerResponseBody) SetModelDetails(v *DescribeModelsForInnerResponseBodyModelDetails) *DescribeModelsForInnerResponseBody {
	s.ModelDetails = v
	return s
}

func (s *DescribeModelsForInnerResponseBody) SetPageNumber(v int32) *DescribeModelsForInnerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelsForInnerResponseBody) SetPageSize(v int32) *DescribeModelsForInnerResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeModelsForInnerResponseBody) SetRequestId(v string) *DescribeModelsForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelsForInnerResponseBody) SetTotalCount(v int32) *DescribeModelsForInnerResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeModelsForInnerResponseBodyModelDetails struct {
	ModelDetail []*DescribeModelsForInnerResponseBodyModelDetailsModelDetail `json:"ModelDetail,omitempty" xml:"ModelDetail,omitempty" type:"Repeated"`
}

func (s DescribeModelsForInnerResponseBodyModelDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsForInnerResponseBodyModelDetails) GoString() string {
	return s.String()
}

func (s *DescribeModelsForInnerResponseBodyModelDetails) SetModelDetail(v []*DescribeModelsForInnerResponseBodyModelDetailsModelDetail) *DescribeModelsForInnerResponseBodyModelDetails {
	s.ModelDetail = v
	return s
}

type DescribeModelsForInnerResponseBodyModelDetailsModelDetail struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName    *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelRef     *string `json:"ModelRef,omitempty" xml:"ModelRef,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeModelsForInnerResponseBodyModelDetailsModelDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsForInnerResponseBodyModelDetailsModelDetail) GoString() string {
	return s.String()
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetCreatedTime(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.CreatedTime = &v
	return s
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetDescription(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.Description = &v
	return s
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetGroupId(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.GroupId = &v
	return s
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetModelName(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.ModelName = &v
	return s
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetModelRef(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.ModelRef = &v
	return s
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetModifiedTime(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeModelsForInnerResponseBodyModelDetailsModelDetail) SetSchema(v string) *DescribeModelsForInnerResponseBodyModelDetailsModelDetail {
	s.Schema = &v
	return s
}

type DescribeModelsForInnerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelsForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelsForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelsForInnerResponse) SetHeaders(v map[string]*string) *DescribeModelsForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelsForInnerResponse) SetStatusCode(v int32) *DescribeModelsForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelsForInnerResponse) SetBody(v *DescribeModelsForInnerResponseBody) *DescribeModelsForInnerResponse {
	s.Body = v
	return s
}

type DescribePurchasedApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePurchasedApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiRequest) SetApiId(v string) *DescribePurchasedApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribePurchasedApiRequest) SetGroupId(v string) *DescribePurchasedApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiRequest) SetSecurityToken(v string) *DescribePurchasedApiRequest {
	s.SecurityToken = &v
	return s
}

type DescribePurchasedApiResponseBody struct {
	ApiId               *string                                         `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string                                         `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BodyFormat          *string                                         `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	CreatedTime         *string                                         `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description         *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string                                         `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string                                         `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpMethod          *string                                         `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol        *string                                         `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Mock                *string                                         `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult          *string                                         `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	ModifiedTime        *string                                         `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Path                *string                                         `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters      *DescribePurchasedApiResponseBodyPathParameters `json:"PathParameters,omitempty" xml:"PathParameters,omitempty" type:"Struct"`
	PostBodyDescription *string                                         `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType        *string                                         `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RegionId            *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestBody         *DescribePurchasedApiResponseBodyRequestBody    `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	RequestHeaders      *DescribePurchasedApiResponseBodyRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Struct"`
	RequestId           *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestQueries      *DescribePurchasedApiResponseBodyRequestQueries `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty" type:"Struct"`
	ResultSample        *string                                         `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType          *string                                         `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	Visibility          *string                                         `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribePurchasedApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBody) SetApiId(v string) *DescribePurchasedApiResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetApiName(v string) *DescribePurchasedApiResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetBodyFormat(v string) *DescribePurchasedApiResponseBody {
	s.BodyFormat = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetCreatedTime(v string) *DescribePurchasedApiResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetDescription(v string) *DescribePurchasedApiResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetGroupId(v string) *DescribePurchasedApiResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetGroupName(v string) *DescribePurchasedApiResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetHttpMethod(v string) *DescribePurchasedApiResponseBody {
	s.HttpMethod = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetHttpProtocol(v string) *DescribePurchasedApiResponseBody {
	s.HttpProtocol = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetMock(v string) *DescribePurchasedApiResponseBody {
	s.Mock = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetMockResult(v string) *DescribePurchasedApiResponseBody {
	s.MockResult = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetModifiedTime(v string) *DescribePurchasedApiResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetPath(v string) *DescribePurchasedApiResponseBody {
	s.Path = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetPathParameters(v *DescribePurchasedApiResponseBodyPathParameters) *DescribePurchasedApiResponseBody {
	s.PathParameters = v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetPostBodyDescription(v string) *DescribePurchasedApiResponseBody {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetPostBodyType(v string) *DescribePurchasedApiResponseBody {
	s.PostBodyType = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetRegionId(v string) *DescribePurchasedApiResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetRequestBody(v *DescribePurchasedApiResponseBodyRequestBody) *DescribePurchasedApiResponseBody {
	s.RequestBody = v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetRequestHeaders(v *DescribePurchasedApiResponseBodyRequestHeaders) *DescribePurchasedApiResponseBody {
	s.RequestHeaders = v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetRequestId(v string) *DescribePurchasedApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetRequestQueries(v *DescribePurchasedApiResponseBodyRequestQueries) *DescribePurchasedApiResponseBody {
	s.RequestQueries = v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetResultSample(v string) *DescribePurchasedApiResponseBody {
	s.ResultSample = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetResultType(v string) *DescribePurchasedApiResponseBody {
	s.ResultType = &v
	return s
}

func (s *DescribePurchasedApiResponseBody) SetVisibility(v string) *DescribePurchasedApiResponseBody {
	s.Visibility = &v
	return s
}

type DescribePurchasedApiResponseBodyPathParameters struct {
	PathParameter []*DescribePurchasedApiResponseBodyPathParametersPathParameter `json:"PathParameter,omitempty" xml:"PathParameter,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiResponseBodyPathParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyPathParameters) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyPathParameters) SetPathParameter(v []*DescribePurchasedApiResponseBodyPathParametersPathParameter) *DescribePurchasedApiResponseBodyPathParameters {
	s.PathParameter = v
	return s
}

type DescribePurchasedApiResponseBodyPathParametersPathParameter struct {
	ApiParameterName *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DemoValue        *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribePurchasedApiResponseBodyPathParametersPathParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyPathParametersPathParameter) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyPathParametersPathParameter) SetApiParameterName(v string) *DescribePurchasedApiResponseBodyPathParametersPathParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyPathParametersPathParameter) SetDemoValue(v string) *DescribePurchasedApiResponseBodyPathParametersPathParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyPathParametersPathParameter) SetDescription(v string) *DescribePurchasedApiResponseBodyPathParametersPathParameter {
	s.Description = &v
	return s
}

type DescribePurchasedApiResponseBodyRequestBody struct {
	RequestParam []*DescribePurchasedApiResponseBodyRequestBodyRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiResponseBodyRequestBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyRequestBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyRequestBody) SetRequestParam(v []*DescribePurchasedApiResponseBodyRequestBodyRequestParam) *DescribePurchasedApiResponseBodyRequestBody {
	s.RequestParam = v
	return s
}

type DescribePurchasedApiResponseBodyRequestBodyRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow           *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribePurchasedApiResponseBodyRequestBodyRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyRequestBodyRequestParam) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetApiParameterName(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetDefaultValue(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetDemoValue(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetDescription(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetDocOrder(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetDocShow(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetEnumValue(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetJsonScheme(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetMaxLength(v int64) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetMaxValue(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetMinLength(v int64) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetMinValue(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetParameterType(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetRegularExpression(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestBodyRequestParam) SetRequired(v string) *DescribePurchasedApiResponseBodyRequestBodyRequestParam {
	s.Required = &v
	return s
}

type DescribePurchasedApiResponseBodyRequestHeaders struct {
	RequestParam []*DescribePurchasedApiResponseBodyRequestHeadersRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiResponseBodyRequestHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyRequestHeaders) SetRequestParam(v []*DescribePurchasedApiResponseBodyRequestHeadersRequestParam) *DescribePurchasedApiResponseBodyRequestHeaders {
	s.RequestParam = v
	return s
}

type DescribePurchasedApiResponseBodyRequestHeadersRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow           *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribePurchasedApiResponseBodyRequestHeadersRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyRequestHeadersRequestParam) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetApiParameterName(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetDefaultValue(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetDemoValue(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetDescription(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetDocOrder(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetDocShow(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetEnumValue(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetJsonScheme(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetMaxLength(v int64) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetMaxValue(v int64) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetMinLength(v int64) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetMinValue(v int64) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetParameterType(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetRegularExpression(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestHeadersRequestParam) SetRequired(v string) *DescribePurchasedApiResponseBodyRequestHeadersRequestParam {
	s.Required = &v
	return s
}

type DescribePurchasedApiResponseBodyRequestQueries struct {
	RequestParam []*DescribePurchasedApiResponseBodyRequestQueriesRequestParam `json:"RequestParam,omitempty" xml:"RequestParam,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiResponseBodyRequestQueries) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyRequestQueries) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyRequestQueries) SetRequestParam(v []*DescribePurchasedApiResponseBodyRequestQueriesRequestParam) *DescribePurchasedApiResponseBodyRequestQueries {
	s.RequestParam = v
	return s
}

type DescribePurchasedApiResponseBodyRequestQueriesRequestParam struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *string `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	DocShow           *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	EnumValue         *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	JsonScheme        *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	MaxLength         *int64  `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	MaxValue          *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinLength         *int64  `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	MinValue          *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ParameterType     *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	Required          *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribePurchasedApiResponseBodyRequestQueriesRequestParam) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponseBodyRequestQueriesRequestParam) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetApiParameterName(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.ApiParameterName = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetDefaultValue(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.DefaultValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetDemoValue(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.DemoValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetDescription(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetDocOrder(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.DocOrder = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetDocShow(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.DocShow = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetEnumValue(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.EnumValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetJsonScheme(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.JsonScheme = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetMaxLength(v int64) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.MaxLength = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetMaxValue(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.MaxValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetMinLength(v int64) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.MinLength = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetMinValue(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.MinValue = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetParameterType(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.ParameterType = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetRegularExpression(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.RegularExpression = &v
	return s
}

func (s *DescribePurchasedApiResponseBodyRequestQueriesRequestParam) SetRequired(v string) *DescribePurchasedApiResponseBodyRequestQueriesRequestParam {
	s.Required = &v
	return s
}

type DescribePurchasedApiResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiResponse) SetHeaders(v map[string]*string) *DescribePurchasedApiResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApiResponse) SetStatusCode(v int32) *DescribePurchasedApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApiResponse) SetBody(v *DescribePurchasedApiResponseBody) *DescribePurchasedApiResponse {
	s.Body = v
	return s
}

type DescribePurchasedApiGroupDetailRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePurchasedApiGroupDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupDetailRequest) SetGroupId(v string) *DescribePurchasedApiGroupDetailRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailRequest) SetSecurityToken(v string) *DescribePurchasedApiGroupDetailRequest {
	s.SecurityToken = &v
	return s
}

type DescribePurchasedApiGroupDetailResponseBody struct {
	CreatedTime  *string                                                 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainItems  *DescribePurchasedApiGroupDetailResponseBodyDomainItems `json:"DomainItems,omitempty" xml:"DomainItems,omitempty" type:"Struct"`
	GroupId      *string                                                 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string                                                 `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifiedTime *string                                                 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId     *string                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePurchasedApiGroupDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetCreatedTime(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetDescription(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetDomainItems(v *DescribePurchasedApiGroupDetailResponseBodyDomainItems) *DescribePurchasedApiGroupDetailResponseBody {
	s.DomainItems = v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetGroupId(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetGroupName(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetModifiedTime(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetRegionId(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetRequestId(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponseBody) SetStatus(v string) *DescribePurchasedApiGroupDetailResponseBody {
	s.Status = &v
	return s
}

type DescribePurchasedApiGroupDetailResponseBodyDomainItems struct {
	DomainItem []*DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiGroupDetailResponseBodyDomainItems) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupDetailResponseBodyDomainItems) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupDetailResponseBodyDomainItems) SetDomainItem(v []*DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem) *DescribePurchasedApiGroupDetailResponseBodyDomainItems {
	s.DomainItem = v
	return s
}

type DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainName(v string) *DescribePurchasedApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainName = &v
	return s
}

type DescribePurchasedApiGroupDetailResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApiGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApiGroupDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupDetailResponse) SetHeaders(v map[string]*string) *DescribePurchasedApiGroupDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponse) SetStatusCode(v int32) *DescribePurchasedApiGroupDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApiGroupDetailResponse) SetBody(v *DescribePurchasedApiGroupDetailResponseBody) *DescribePurchasedApiGroupDetailResponse {
	s.Body = v
	return s
}

type DescribePurchasedApiGroupsRequest struct {
	GroupIds      *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePurchasedApiGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsRequest) SetGroupIds(v string) *DescribePurchasedApiGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *DescribePurchasedApiGroupsRequest) SetPageNumber(v int32) *DescribePurchasedApiGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApiGroupsRequest) SetPageSize(v int32) *DescribePurchasedApiGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApiGroupsRequest) SetSecurityToken(v string) *DescribePurchasedApiGroupsRequest {
	s.SecurityToken = &v
	return s
}

type DescribePurchasedApiGroupsResponseBody struct {
	PageNumber                  *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                    *int32                                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PurchasedApiGroupAttributes *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes `json:"PurchasedApiGroupAttributes,omitempty" xml:"PurchasedApiGroupAttributes,omitempty" type:"Struct"`
	RequestId                   *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                  *int32                                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurchasedApiGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponseBody) SetPageNumber(v int32) *DescribePurchasedApiGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetPageSize(v int32) *DescribePurchasedApiGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetPurchasedApiGroupAttributes(v *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) *DescribePurchasedApiGroupsResponseBody {
	s.PurchasedApiGroupAttributes = v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetRequestId(v string) *DescribePurchasedApiGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBody) SetTotalCount(v int32) *DescribePurchasedApiGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes struct {
	PurchasedApiGroupAttribute []*DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute `json:"PurchasedApiGroupAttribute,omitempty" xml:"PurchasedApiGroupAttribute,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes) SetPurchasedApiGroupAttribute(v []*DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributes {
	s.PurchasedApiGroupAttribute = v
	return s
}

type DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute struct {
	BillingType      *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CreatedTime      *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ExpireTime       *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GroupDescription *string `json:"GroupDescription,omitempty" xml:"GroupDescription,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName        *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InvokeTimesMax   *int64  `json:"InvokeTimesMax,omitempty" xml:"InvokeTimesMax,omitempty"`
	InvokeTimesNow   *int64  `json:"InvokeTimesNow,omitempty" xml:"InvokeTimesNow,omitempty"`
	ModifiedTime     *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetBillingType(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.BillingType = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetCreatedTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetExpireTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetGroupDescription(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.GroupDescription = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetGroupId(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetGroupName(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetInvokeTimesMax(v int64) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.InvokeTimesMax = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetInvokeTimesNow(v int64) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.InvokeTimesNow = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetModifiedTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetRegionId(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetStatus(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.Status = &v
	return s
}

type DescribePurchasedApiGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApiGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupsResponse) SetHeaders(v map[string]*string) *DescribePurchasedApiGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApiGroupsResponse) SetStatusCode(v int32) *DescribePurchasedApiGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponse) SetBody(v *DescribePurchasedApiGroupsResponseBody) *DescribePurchasedApiGroupsResponse {
	s.Body = v
	return s
}

type DescribePurchasedApisRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribePurchasedApisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisRequest) SetApiId(v string) *DescribePurchasedApisRequest {
	s.ApiId = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetApiName(v string) *DescribePurchasedApisRequest {
	s.ApiName = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetGroupId(v string) *DescribePurchasedApisRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetPageNumber(v int32) *DescribePurchasedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetPageSize(v int32) *DescribePurchasedApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetSecurityToken(v string) *DescribePurchasedApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetStageName(v string) *DescribePurchasedApisRequest {
	s.StageName = &v
	return s
}

func (s *DescribePurchasedApisRequest) SetVisibility(v string) *DescribePurchasedApisRequest {
	s.Visibility = &v
	return s
}

type DescribePurchasedApisResponseBody struct {
	ApiInfos   *DescribePurchasedApisResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurchasedApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponseBody) SetApiInfos(v *DescribePurchasedApisResponseBodyApiInfos) *DescribePurchasedApisResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribePurchasedApisResponseBody) SetPageNumber(v int32) *DescribePurchasedApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApisResponseBody) SetPageSize(v int32) *DescribePurchasedApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApisResponseBody) SetRequestId(v string) *DescribePurchasedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApisResponseBody) SetTotalCount(v int32) *DescribePurchasedApisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePurchasedApisResponseBodyApiInfos struct {
	ApiInfo []*DescribePurchasedApisResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApisResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponseBodyApiInfos) SetApiInfo(v []*DescribePurchasedApisResponseBodyApiInfosApiInfo) *DescribePurchasedApisResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribePurchasedApisResponseBodyApiInfosApiInfo struct {
	ApiId        *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName    *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility   *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribePurchasedApisResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetCreatedTime(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetModifiedTime(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribePurchasedApisResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePurchasedApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePurchasedApisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponse) SetHeaders(v map[string]*string) *DescribePurchasedApisResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApisResponse) SetStatusCode(v int32) *DescribePurchasedApisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePurchasedApisResponse) SetBody(v *DescribePurchasedApisResponseBody) *DescribePurchasedApisResponse {
	s.Body = v
	return s
}

type DescribeRaceWorkForInnerRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRaceWorkForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorkForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorkForInnerRequest) SetGroupId(v string) *DescribeRaceWorkForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeRaceWorkForInnerRequest) SetSecurityToken(v string) *DescribeRaceWorkForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRaceWorkForInnerResponseBody struct {
	CommodityCode    *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Keywords         *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	LogoUrl          *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	ModifiedTime     *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	WorkName         *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
}

func (s DescribeRaceWorkForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorkForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorkForInnerResponseBody) SetCommodityCode(v string) *DescribeRaceWorkForInnerResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetCreateTime(v string) *DescribeRaceWorkForInnerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetGroupId(v string) *DescribeRaceWorkForInnerResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetKeywords(v string) *DescribeRaceWorkForInnerResponseBody {
	s.Keywords = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetLogoUrl(v string) *DescribeRaceWorkForInnerResponseBody {
	s.LogoUrl = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetModifiedTime(v string) *DescribeRaceWorkForInnerResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetRequestId(v string) *DescribeRaceWorkForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetShortDescription(v string) *DescribeRaceWorkForInnerResponseBody {
	s.ShortDescription = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponseBody) SetWorkName(v string) *DescribeRaceWorkForInnerResponseBody {
	s.WorkName = &v
	return s
}

type DescribeRaceWorkForInnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRaceWorkForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRaceWorkForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorkForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorkForInnerResponse) SetHeaders(v map[string]*string) *DescribeRaceWorkForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeRaceWorkForInnerResponse) SetStatusCode(v int32) *DescribeRaceWorkForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRaceWorkForInnerResponse) SetBody(v *DescribeRaceWorkForInnerResponseBody) *DescribeRaceWorkForInnerResponse {
	s.Body = v
	return s
}

type DescribeRaceWorksForInnerRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRaceWorksForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorksForInnerRequest) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorksForInnerRequest) SetPageNumber(v int32) *DescribeRaceWorksForInnerRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRaceWorksForInnerRequest) SetPageSize(v int32) *DescribeRaceWorksForInnerRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRaceWorksForInnerRequest) SetSecurityToken(v string) *DescribeRaceWorksForInnerRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRaceWorksForInnerResponseBody struct {
	ApiWorks   *DescribeRaceWorksForInnerResponseBodyApiWorks `json:"ApiWorks,omitempty" xml:"ApiWorks,omitempty" type:"Struct"`
	PageNumber *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRaceWorksForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorksForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorksForInnerResponseBody) SetApiWorks(v *DescribeRaceWorksForInnerResponseBodyApiWorks) *DescribeRaceWorksForInnerResponseBody {
	s.ApiWorks = v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBody) SetPageNumber(v int32) *DescribeRaceWorksForInnerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBody) SetPageSize(v int32) *DescribeRaceWorksForInnerResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBody) SetRequestId(v string) *DescribeRaceWorksForInnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBody) SetTotalCount(v int32) *DescribeRaceWorksForInnerResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeRaceWorksForInnerResponseBodyApiWorks struct {
	ApiWork []*DescribeRaceWorksForInnerResponseBodyApiWorksApiWork `json:"ApiWork,omitempty" xml:"ApiWork,omitempty" type:"Repeated"`
}

func (s DescribeRaceWorksForInnerResponseBodyApiWorks) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorksForInnerResponseBodyApiWorks) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorks) SetApiWork(v []*DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) *DescribeRaceWorksForInnerResponseBodyApiWorks {
	s.ApiWork = v
	return s
}

type DescribeRaceWorksForInnerResponseBodyApiWorksApiWork struct {
	CommodityCode    *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Keywords         *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	LogoUrl          *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	ModifiedTime     *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
	WorkName         *string `json:"WorkName,omitempty" xml:"WorkName,omitempty"`
}

func (s DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetCommodityCode(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.CommodityCode = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetCreateTime(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.CreateTime = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetGroupId(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.GroupId = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetKeywords(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.Keywords = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetLogoUrl(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.LogoUrl = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetModifiedTime(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetShortDescription(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.ShortDescription = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork) SetWorkName(v string) *DescribeRaceWorksForInnerResponseBodyApiWorksApiWork {
	s.WorkName = &v
	return s
}

type DescribeRaceWorksForInnerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRaceWorksForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRaceWorksForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRaceWorksForInnerResponse) GoString() string {
	return s.String()
}

func (s *DescribeRaceWorksForInnerResponse) SetHeaders(v map[string]*string) *DescribeRaceWorksForInnerResponse {
	s.Headers = v
	return s
}

func (s *DescribeRaceWorksForInnerResponse) SetStatusCode(v int32) *DescribeRaceWorksForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRaceWorksForInnerResponse) SetBody(v *DescribeRaceWorksForInnerResponseBody) *DescribeRaceWorksForInnerResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetLanguage(v string) *DescribeRegionsRequest {
	s.Language = &v
	return s
}

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   *DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v *DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Region []*DescribeRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetRegion(v []*DescribeRegionsResponseBodyRegionsRegion) *DescribeRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseBodyRegionsRegion struct {
	EndPoint  *string `json:"EndPoint,omitempty" xml:"EndPoint,omitempty"`
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetEndPoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.EndPoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeRulesByApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeRulesByApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesByApiRequest) SetApiId(v string) *DescribeRulesByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeRulesByApiRequest) SetGroupId(v string) *DescribeRulesByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeRulesByApiRequest) SetSecurityToken(v string) *DescribeRulesByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRulesByApiRequest) SetStageName(v string) *DescribeRulesByApiRequest {
	s.StageName = &v
	return s
}

type DescribeRulesByApiResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     *DescribeRulesByApiResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeRulesByApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesByApiResponseBody) SetRequestId(v string) *DescribeRulesByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesByApiResponseBody) SetRules(v *DescribeRulesByApiResponseBodyRules) *DescribeRulesByApiResponseBody {
	s.Rules = v
	return s
}

type DescribeRulesByApiResponseBodyRules struct {
	Rule []*DescribeRulesByApiResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeRulesByApiResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesByApiResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeRulesByApiResponseBodyRules) SetRule(v []*DescribeRulesByApiResponseBodyRulesRule) *DescribeRulesByApiResponseBodyRules {
	s.Rule = v
	return s
}

type DescribeRulesByApiResponseBodyRulesRule struct {
	CreatedTime *int32  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	RuleId      *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName    *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeRulesByApiResponseBodyRulesRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesByApiResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRulesByApiResponseBodyRulesRule) SetCreatedTime(v int32) *DescribeRulesByApiResponseBodyRulesRule {
	s.CreatedTime = &v
	return s
}

func (s *DescribeRulesByApiResponseBodyRulesRule) SetRuleId(v string) *DescribeRulesByApiResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeRulesByApiResponseBodyRulesRule) SetRuleName(v string) *DescribeRulesByApiResponseBodyRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeRulesByApiResponseBodyRulesRule) SetRuleType(v string) *DescribeRulesByApiResponseBodyRulesRule {
	s.RuleType = &v
	return s
}

type DescribeRulesByApiResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRulesByApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRulesByApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulesByApiResponse) SetHeaders(v map[string]*string) *DescribeRulesByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulesByApiResponse) SetStatusCode(v int32) *DescribeRulesByApiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulesByApiResponse) SetBody(v *DescribeRulesByApiResponseBody) *DescribeRulesByApiResponse {
	s.Body = v
	return s
}

type DescribeSecretKeysRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecretKeyId   *string `json:"SecretKeyId,omitempty" xml:"SecretKeyId,omitempty"`
	SecretKeyName *string `json:"SecretKeyName,omitempty" xml:"SecretKeyName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSecretKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretKeysRequest) SetPageNumber(v int32) *DescribeSecretKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecretKeysRequest) SetPageSize(v int32) *DescribeSecretKeysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecretKeysRequest) SetSecretKeyId(v string) *DescribeSecretKeysRequest {
	s.SecretKeyId = &v
	return s
}

func (s *DescribeSecretKeysRequest) SetSecretKeyName(v string) *DescribeSecretKeysRequest {
	s.SecretKeyName = &v
	return s
}

func (s *DescribeSecretKeysRequest) SetSecurityToken(v string) *DescribeSecretKeysRequest {
	s.SecurityToken = &v
	return s
}

type DescribeSecretKeysResponseBody struct {
	PageNumber *int32                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretKeys *DescribeSecretKeysResponseBodySecretKeys `json:"SecretKeys,omitempty" xml:"SecretKeys,omitempty" type:"Struct"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecretKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecretKeysResponseBody) SetPageNumber(v int32) *DescribeSecretKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecretKeysResponseBody) SetPageSize(v int32) *DescribeSecretKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSecretKeysResponseBody) SetRequestId(v string) *DescribeSecretKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretKeysResponseBody) SetSecretKeys(v *DescribeSecretKeysResponseBodySecretKeys) *DescribeSecretKeysResponseBody {
	s.SecretKeys = v
	return s
}

func (s *DescribeSecretKeysResponseBody) SetTotalCount(v int32) *DescribeSecretKeysResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSecretKeysResponseBodySecretKeys struct {
	SecretKey []*DescribeSecretKeysResponseBodySecretKeysSecretKey `json:"SecretKey,omitempty" xml:"SecretKey,omitempty" type:"Repeated"`
}

func (s DescribeSecretKeysResponseBodySecretKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretKeysResponseBodySecretKeys) GoString() string {
	return s.String()
}

func (s *DescribeSecretKeysResponseBodySecretKeys) SetSecretKey(v []*DescribeSecretKeysResponseBodySecretKeysSecretKey) *DescribeSecretKeysResponseBodySecretKeys {
	s.SecretKey = v
	return s
}

type DescribeSecretKeysResponseBodySecretKeysSecretKey struct {
	CreatedTime    *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ModifiedTime   *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecretKey      *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	SecretKeyId    *string `json:"SecretKeyId,omitempty" xml:"SecretKeyId,omitempty"`
	SecretKeyName  *string `json:"SecretKeyName,omitempty" xml:"SecretKeyName,omitempty"`
	SecretKeyValue *string `json:"SecretKeyValue,omitempty" xml:"SecretKeyValue,omitempty"`
}

func (s DescribeSecretKeysResponseBodySecretKeysSecretKey) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretKeysResponseBodySecretKeysSecretKey) GoString() string {
	return s.String()
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetCreatedTime(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.CreatedTime = &v
	return s
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetModifiedTime(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetRegionId(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.RegionId = &v
	return s
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetSecretKey(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.SecretKey = &v
	return s
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetSecretKeyId(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.SecretKeyId = &v
	return s
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetSecretKeyName(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.SecretKeyName = &v
	return s
}

func (s *DescribeSecretKeysResponseBodySecretKeysSecretKey) SetSecretKeyValue(v string) *DescribeSecretKeysResponseBodySecretKeysSecretKey {
	s.SecretKeyValue = &v
	return s
}

type DescribeSecretKeysResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecretKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecretKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecretKeysResponse) SetHeaders(v map[string]*string) *DescribeSecretKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecretKeysResponse) SetStatusCode(v int32) *DescribeSecretKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecretKeysResponse) SetBody(v *DescribeSecretKeysResponseBody) *DescribeSecretKeysResponse {
	s.Body = v
	return s
}

type DescribeSystemParametersRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSystemParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersRequest) SetSecurityToken(v string) *DescribeSystemParametersRequest {
	s.SecurityToken = &v
	return s
}

type DescribeSystemParametersResponseBody struct {
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemParameters *DescribeSystemParametersResponseBodySystemParameters `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
}

func (s DescribeSystemParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBody) SetRequestId(v string) *DescribeSystemParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemParametersResponseBody) SetSystemParameters(v *DescribeSystemParametersResponseBodySystemParameters) *DescribeSystemParametersResponseBody {
	s.SystemParameters = v
	return s
}

type DescribeSystemParametersResponseBodySystemParameters struct {
	SystemParameter []*DescribeSystemParametersResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeSystemParametersResponseBodySystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBodySystemParameters) SetSystemParameter(v []*DescribeSystemParametersResponseBodySystemParametersSystemParameter) *DescribeSystemParametersResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

type DescribeSystemParametersResponseBodySystemParametersSystemParameter struct {
	DemoValue   *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ParamName   *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeSystemParametersResponseBodySystemParametersSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBodySystemParametersSystemParameter) SetDemoValue(v string) *DescribeSystemParametersResponseBodySystemParametersSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParametersSystemParameter) SetDescription(v string) *DescribeSystemParametersResponseBodySystemParametersSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParametersSystemParameter) SetParamName(v string) *DescribeSystemParametersResponseBodySystemParametersSystemParameter {
	s.ParamName = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParametersSystemParameter) SetParamType(v string) *DescribeSystemParametersResponseBodySystemParametersSystemParameter {
	s.ParamType = &v
	return s
}

type DescribeSystemParametersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponse) SetHeaders(v map[string]*string) *DescribeSystemParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemParametersResponse) SetStatusCode(v int32) *DescribeSystemParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemParametersResponse) SetBody(v *DescribeSystemParametersResponseBody) *DescribeSystemParametersResponse {
	s.Body = v
	return s
}

type DescribeSystemParamsRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeSystemParamsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParamsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSystemParamsRequest) SetSecurityToken(v string) *DescribeSystemParamsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeSystemParamsResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemParams *DescribeSystemParamsResponseBodySystemParams `json:"SystemParams,omitempty" xml:"SystemParams,omitempty" type:"Struct"`
}

func (s DescribeSystemParamsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParamsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSystemParamsResponseBody) SetRequestId(v string) *DescribeSystemParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSystemParamsResponseBody) SetSystemParams(v *DescribeSystemParamsResponseBodySystemParams) *DescribeSystemParamsResponseBody {
	s.SystemParams = v
	return s
}

type DescribeSystemParamsResponseBodySystemParams struct {
	SystemParam []*DescribeSystemParamsResponseBodySystemParamsSystemParam `json:"SystemParam,omitempty" xml:"SystemParam,omitempty" type:"Repeated"`
}

func (s DescribeSystemParamsResponseBodySystemParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParamsResponseBodySystemParams) GoString() string {
	return s.String()
}

func (s *DescribeSystemParamsResponseBodySystemParams) SetSystemParam(v []*DescribeSystemParamsResponseBodySystemParamsSystemParam) *DescribeSystemParamsResponseBodySystemParams {
	s.SystemParam = v
	return s
}

type DescribeSystemParamsResponseBodySystemParamsSystemParam struct {
	DemoValue   *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ParamName   *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeSystemParamsResponseBodySystemParamsSystemParam) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParamsResponseBodySystemParamsSystemParam) GoString() string {
	return s.String()
}

func (s *DescribeSystemParamsResponseBodySystemParamsSystemParam) SetDemoValue(v string) *DescribeSystemParamsResponseBodySystemParamsSystemParam {
	s.DemoValue = &v
	return s
}

func (s *DescribeSystemParamsResponseBodySystemParamsSystemParam) SetDescription(v string) *DescribeSystemParamsResponseBodySystemParamsSystemParam {
	s.Description = &v
	return s
}

func (s *DescribeSystemParamsResponseBodySystemParamsSystemParam) SetParamName(v string) *DescribeSystemParamsResponseBodySystemParamsSystemParam {
	s.ParamName = &v
	return s
}

func (s *DescribeSystemParamsResponseBodySystemParamsSystemParam) SetParamType(v string) *DescribeSystemParamsResponseBodySystemParamsSystemParam {
	s.ParamType = &v
	return s
}

type DescribeSystemParamsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSystemParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSystemParamsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSystemParamsResponse) SetHeaders(v map[string]*string) *DescribeSystemParamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSystemParamsResponse) SetStatusCode(v int32) *DescribeSystemParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSystemParamsResponse) SetBody(v *DescribeSystemParamsResponseBody) *DescribeSystemParamsResponse {
	s.Body = v
	return s
}

type DescribeTrafficControlsRequest struct {
	ApiUid             *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber         *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName          *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	TrafficControlId   *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
}

func (s DescribeTrafficControlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsRequest) SetApiUid(v string) *DescribeTrafficControlsRequest {
	s.ApiUid = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetGroupId(v string) *DescribeTrafficControlsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetPageNumber(v int32) *DescribeTrafficControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetPageSize(v int32) *DescribeTrafficControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetSecurityToken(v string) *DescribeTrafficControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetStageName(v string) *DescribeTrafficControlsRequest {
	s.StageName = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetTrafficControlId(v string) *DescribeTrafficControlsRequest {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeTrafficControlsRequest) SetTrafficControlName(v string) *DescribeTrafficControlsRequest {
	s.TrafficControlName = &v
	return s
}

type DescribeTrafficControlsResponseBody struct {
	PageNumber      *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId       *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount      *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TrafficControls *DescribeTrafficControlsResponseBodyTrafficControls `json:"TrafficControls,omitempty" xml:"TrafficControls,omitempty" type:"Struct"`
}

func (s DescribeTrafficControlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBody) SetPageNumber(v int32) *DescribeTrafficControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetPageSize(v int32) *DescribeTrafficControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetRequestId(v string) *DescribeTrafficControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetTotalCount(v int32) *DescribeTrafficControlsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTrafficControlsResponseBody) SetTrafficControls(v *DescribeTrafficControlsResponseBodyTrafficControls) *DescribeTrafficControlsResponseBody {
	s.TrafficControls = v
	return s
}

type DescribeTrafficControlsResponseBodyTrafficControls struct {
	TrafficControl []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl `json:"TrafficControl,omitempty" xml:"TrafficControl,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControls) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControls) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControls) SetTrafficControl(v []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) *DescribeTrafficControlsResponseBodyTrafficControls {
	s.TrafficControl = v
	return s
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl struct {
	ApiDefault         *int32                                                                           `json:"ApiDefault,omitempty" xml:"ApiDefault,omitempty"`
	AppDefault         *int32                                                                           `json:"AppDefault,omitempty" xml:"AppDefault,omitempty"`
	CreatedTime        *string                                                                          `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description        *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime       *string                                                                          `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	SpecialPolicies    *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies `json:"SpecialPolicies,omitempty" xml:"SpecialPolicies,omitempty" type:"Struct"`
	TrafficControlId   *string                                                                          `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	TrafficControlName *string                                                                          `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
	TrafficControlUnit *string                                                                          `json:"TrafficControlUnit,omitempty" xml:"TrafficControlUnit,omitempty"`
	UserDefault        *int32                                                                           `json:"UserDefault,omitempty" xml:"UserDefault,omitempty"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetApiDefault(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.ApiDefault = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetAppDefault(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.AppDefault = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetCreatedTime(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.CreatedTime = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetDescription(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.Description = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetModifiedTime(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetSpecialPolicies(v *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.SpecialPolicies = v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetTrafficControlId(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetTrafficControlName(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.TrafficControlName = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetTrafficControlUnit(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.TrafficControlUnit = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl) SetUserDefault(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControl {
	s.UserDefault = &v
	return s
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies struct {
	SpecialPolicy []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy `json:"SpecialPolicy,omitempty" xml:"SpecialPolicy,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies) SetSpecialPolicy(v []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPolicies {
	s.SpecialPolicy = v
	return s
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy struct {
	SpecialType *string                                                                                               `json:"SpecialType,omitempty" xml:"SpecialType,omitempty"`
	Specials    *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials `json:"Specials,omitempty" xml:"Specials,omitempty" type:"Struct"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) SetSpecialType(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy {
	s.SpecialType = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy) SetSpecials(v *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicy {
	s.Specials = v
	return s
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials struct {
	Special []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial `json:"Special,omitempty" xml:"Special,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials) SetSpecial(v []*DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecials {
	s.Special = v
	return s
}

type DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial struct {
	SpecialKey   *string `json:"SpecialKey,omitempty" xml:"SpecialKey,omitempty"`
	TrafficValue *int32  `json:"TrafficValue,omitempty" xml:"TrafficValue,omitempty"`
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) SetSpecialKey(v string) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial {
	s.SpecialKey = &v
	return s
}

func (s *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial) SetTrafficValue(v int32) *DescribeTrafficControlsResponseBodyTrafficControlsTrafficControlSpecialPoliciesSpecialPolicySpecialsSpecial {
	s.TrafficValue = &v
	return s
}

type DescribeTrafficControlsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTrafficControlsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTrafficControlsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsResponse) SetHeaders(v map[string]*string) *DescribeTrafficControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficControlsResponse) SetStatusCode(v int32) *DescribeTrafficControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTrafficControlsResponse) SetBody(v *DescribeTrafficControlsResponseBody) *DescribeTrafficControlsResponse {
	s.Body = v
	return s
}

type DescribeUserWhiteListsRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid           *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeUserWhiteListsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWhiteListsRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserWhiteListsRequest) SetPageNumber(v int32) *DescribeUserWhiteListsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserWhiteListsRequest) SetPageSize(v int32) *DescribeUserWhiteListsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUserWhiteListsRequest) SetSecurityToken(v string) *DescribeUserWhiteListsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeUserWhiteListsRequest) SetType(v string) *DescribeUserWhiteListsRequest {
	s.Type = &v
	return s
}

func (s *DescribeUserWhiteListsRequest) SetUid(v int64) *DescribeUserWhiteListsRequest {
	s.Uid = &v
	return s
}

type DescribeUserWhiteListsResponseBody struct {
	PageNumber         *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	UserWhiteListInfos *DescribeUserWhiteListsResponseBodyUserWhiteListInfos `json:"UserWhiteListInfos,omitempty" xml:"UserWhiteListInfos,omitempty" type:"Struct"`
}

func (s DescribeUserWhiteListsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWhiteListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserWhiteListsResponseBody) SetPageNumber(v int32) *DescribeUserWhiteListsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBody) SetPageSize(v int32) *DescribeUserWhiteListsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBody) SetRequestId(v string) *DescribeUserWhiteListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBody) SetTotalCount(v int32) *DescribeUserWhiteListsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBody) SetUserWhiteListInfos(v *DescribeUserWhiteListsResponseBodyUserWhiteListInfos) *DescribeUserWhiteListsResponseBody {
	s.UserWhiteListInfos = v
	return s
}

type DescribeUserWhiteListsResponseBodyUserWhiteListInfos struct {
	UserWhiteListInfo []*DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo `json:"UserWhiteListInfo,omitempty" xml:"UserWhiteListInfo,omitempty" type:"Repeated"`
}

func (s DescribeUserWhiteListsResponseBodyUserWhiteListInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWhiteListsResponseBodyUserWhiteListInfos) GoString() string {
	return s.String()
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfos) SetUserWhiteListInfo(v []*DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) *DescribeUserWhiteListsResponseBodyUserWhiteListInfos {
	s.UserWhiteListInfo = v
	return s
}

type DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo struct {
	AoneId       *string `json:"AoneId,omitempty" xml:"AoneId,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid          *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetAoneId(v string) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.AoneId = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetCreateTime(v string) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetDescription(v string) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.Description = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetId(v int64) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.Id = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetModifiedTime(v string) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetType(v string) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.Type = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetUid(v int64) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.Uid = &v
	return s
}

func (s *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo) SetValue(v string) *DescribeUserWhiteListsResponseBodyUserWhiteListInfosUserWhiteListInfo {
	s.Value = &v
	return s
}

type DescribeUserWhiteListsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserWhiteListsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserWhiteListsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserWhiteListsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserWhiteListsResponse) SetHeaders(v map[string]*string) *DescribeUserWhiteListsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserWhiteListsResponse) SetStatusCode(v int32) *DescribeUserWhiteListsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserWhiteListsResponse) SetBody(v *DescribeUserWhiteListsResponseBody) *DescribeUserWhiteListsResponse {
	s.Body = v
	return s
}

type IsIncludedByUserWhitelistRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IsIncludedByUserWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s IsIncludedByUserWhitelistRequest) GoString() string {
	return s.String()
}

func (s *IsIncludedByUserWhitelistRequest) SetSecurityToken(v string) *IsIncludedByUserWhitelistRequest {
	s.SecurityToken = &v
	return s
}

func (s *IsIncludedByUserWhitelistRequest) SetType(v string) *IsIncludedByUserWhitelistRequest {
	s.Type = &v
	return s
}

type IsIncludedByUserWhitelistResponseBody struct {
	IsIncluded *bool   `json:"IsIncluded,omitempty" xml:"IsIncluded,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IsIncludedByUserWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IsIncludedByUserWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *IsIncludedByUserWhitelistResponseBody) SetIsIncluded(v bool) *IsIncludedByUserWhitelistResponseBody {
	s.IsIncluded = &v
	return s
}

func (s *IsIncludedByUserWhitelistResponseBody) SetRequestId(v string) *IsIncludedByUserWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type IsIncludedByUserWhitelistResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsIncludedByUserWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsIncludedByUserWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s IsIncludedByUserWhitelistResponse) GoString() string {
	return s.String()
}

func (s *IsIncludedByUserWhitelistResponse) SetHeaders(v map[string]*string) *IsIncludedByUserWhitelistResponse {
	s.Headers = v
	return s
}

func (s *IsIncludedByUserWhitelistResponse) SetStatusCode(v int32) *IsIncludedByUserWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *IsIncludedByUserWhitelistResponse) SetBody(v *IsIncludedByUserWhitelistResponseBody) *IsIncludedByUserWhitelistResponse {
	s.Body = v
	return s
}

type ModifyApiRequest struct {
	ApiId               *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType            *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	ConstantParameters  *string `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HttpMethod          *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol        *string `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters      *string `json:"PathParameters,omitempty" xml:"PathParameters,omitempty"`
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType        *string `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RequestBody         *string `json:"RequestBody,omitempty" xml:"RequestBody,omitempty"`
	RequestHeaders      *string `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty"`
	RequestQueries      *string `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty"`
	ResultSample        *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType          *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceAddress      *string `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceProtocol     *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout      *int32  `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	SystemParameters    *string `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty"`
	Visibility          *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ModifyApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiRequest) SetApiId(v string) *ModifyApiRequest {
	s.ApiId = &v
	return s
}

func (s *ModifyApiRequest) SetApiName(v string) *ModifyApiRequest {
	s.ApiName = &v
	return s
}

func (s *ModifyApiRequest) SetAuthType(v string) *ModifyApiRequest {
	s.AuthType = &v
	return s
}

func (s *ModifyApiRequest) SetBodyFormat(v string) *ModifyApiRequest {
	s.BodyFormat = &v
	return s
}

func (s *ModifyApiRequest) SetConstantParameters(v string) *ModifyApiRequest {
	s.ConstantParameters = &v
	return s
}

func (s *ModifyApiRequest) SetDescription(v string) *ModifyApiRequest {
	s.Description = &v
	return s
}

func (s *ModifyApiRequest) SetGroupId(v string) *ModifyApiRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiRequest) SetHttpMethod(v string) *ModifyApiRequest {
	s.HttpMethod = &v
	return s
}

func (s *ModifyApiRequest) SetHttpProtocol(v string) *ModifyApiRequest {
	s.HttpProtocol = &v
	return s
}

func (s *ModifyApiRequest) SetPath(v string) *ModifyApiRequest {
	s.Path = &v
	return s
}

func (s *ModifyApiRequest) SetPathParameters(v string) *ModifyApiRequest {
	s.PathParameters = &v
	return s
}

func (s *ModifyApiRequest) SetPostBodyDescription(v string) *ModifyApiRequest {
	s.PostBodyDescription = &v
	return s
}

func (s *ModifyApiRequest) SetPostBodyType(v string) *ModifyApiRequest {
	s.PostBodyType = &v
	return s
}

func (s *ModifyApiRequest) SetRequestBody(v string) *ModifyApiRequest {
	s.RequestBody = &v
	return s
}

func (s *ModifyApiRequest) SetRequestHeaders(v string) *ModifyApiRequest {
	s.RequestHeaders = &v
	return s
}

func (s *ModifyApiRequest) SetRequestQueries(v string) *ModifyApiRequest {
	s.RequestQueries = &v
	return s
}

func (s *ModifyApiRequest) SetResultSample(v string) *ModifyApiRequest {
	s.ResultSample = &v
	return s
}

func (s *ModifyApiRequest) SetResultType(v string) *ModifyApiRequest {
	s.ResultType = &v
	return s
}

func (s *ModifyApiRequest) SetSecurityToken(v string) *ModifyApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiRequest) SetServiceAddress(v string) *ModifyApiRequest {
	s.ServiceAddress = &v
	return s
}

func (s *ModifyApiRequest) SetServiceProtocol(v string) *ModifyApiRequest {
	s.ServiceProtocol = &v
	return s
}

func (s *ModifyApiRequest) SetServiceTimeout(v int32) *ModifyApiRequest {
	s.ServiceTimeout = &v
	return s
}

func (s *ModifyApiRequest) SetSystemParameters(v string) *ModifyApiRequest {
	s.SystemParameters = &v
	return s
}

func (s *ModifyApiRequest) SetVisibility(v string) *ModifyApiRequest {
	s.Visibility = &v
	return s
}

type ModifyApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiResponseBody) SetRequestId(v string) *ModifyApiResponseBody {
	s.RequestId = &v
	return s
}

type ModifyApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiResponse) SetHeaders(v map[string]*string) *ModifyApiResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiResponse) SetStatusCode(v int32) *ModifyApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiResponse) SetBody(v *ModifyApiResponseBody) *ModifyApiResponse {
	s.Body = v
	return s
}

type ModifyApiForInnerRequest struct {
	AliUid               *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId                *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestConfig        *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParamters     *string `json:"RequestParamters,omitempty" xml:"RequestParamters,omitempty"`
	ResultSample         *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType           *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	Visibility           *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ModifyApiForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiForInnerRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiForInnerRequest) SetAliUid(v int64) *ModifyApiForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetApiId(v string) *ModifyApiForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetApiName(v string) *ModifyApiForInnerRequest {
	s.ApiName = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetAuthType(v string) *ModifyApiForInnerRequest {
	s.AuthType = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetDescription(v string) *ModifyApiForInnerRequest {
	s.Description = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetGroupId(v string) *ModifyApiForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetRequestConfig(v string) *ModifyApiForInnerRequest {
	s.RequestConfig = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetRequestParamters(v string) *ModifyApiForInnerRequest {
	s.RequestParamters = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetResultSample(v string) *ModifyApiForInnerRequest {
	s.ResultSample = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetResultType(v string) *ModifyApiForInnerRequest {
	s.ResultType = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetSecurityToken(v string) *ModifyApiForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetServiceConfig(v string) *ModifyApiForInnerRequest {
	s.ServiceConfig = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetServiceParameters(v string) *ModifyApiForInnerRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetServiceParametersMap(v string) *ModifyApiForInnerRequest {
	s.ServiceParametersMap = &v
	return s
}

func (s *ModifyApiForInnerRequest) SetVisibility(v string) *ModifyApiForInnerRequest {
	s.Visibility = &v
	return s
}

type ModifyApiForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiForInnerResponseBody) SetRequestId(v string) *ModifyApiForInnerResponseBody {
	s.RequestId = &v
	return s
}

type ModifyApiForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiForInnerResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiForInnerResponse) SetHeaders(v map[string]*string) *ModifyApiForInnerResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiForInnerResponse) SetStatusCode(v int32) *ModifyApiForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiForInnerResponse) SetBody(v *ModifyApiForInnerResponseBody) *ModifyApiForInnerResponse {
	s.Body = v
	return s
}

type ModifyApiGroupRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupRequest) SetDescription(v string) *ModifyApiGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyApiGroupRequest) SetGroupId(v string) *ModifyApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupRequest) SetGroupName(v string) *ModifyApiGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyApiGroupRequest) SetSecurityToken(v string) *ModifyApiGroupRequest {
	s.SecurityToken = &v
	return s
}

type ModifyApiGroupResponseBody struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain   *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s ModifyApiGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupResponseBody) SetDescription(v string) *ModifyApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *ModifyApiGroupResponseBody) SetGroupId(v string) *ModifyApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupResponseBody) SetGroupName(v string) *ModifyApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *ModifyApiGroupResponseBody) SetRequestId(v string) *ModifyApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApiGroupResponseBody) SetSubDomain(v string) *ModifyApiGroupResponseBody {
	s.SubDomain = &v
	return s
}

type ModifyApiGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupResponse) SetHeaders(v map[string]*string) *ModifyApiGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiGroupResponse) SetStatusCode(v int32) *ModifyApiGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiGroupResponse) SetBody(v *ModifyApiGroupResponseBody) *ModifyApiGroupResponse {
	s.Body = v
	return s
}

type ModifyApiMarketInstanceAttributeRequest struct {
	AliUid             *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	GroupId            *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	InstanceAttributes *string `json:"InstanceAttributes,omitempty" xml:"InstanceAttributes,omitempty"`
}

func (s ModifyApiMarketInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiMarketInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiMarketInstanceAttributeRequest) SetAliUid(v string) *ModifyApiMarketInstanceAttributeRequest {
	s.AliUid = &v
	return s
}

func (s *ModifyApiMarketInstanceAttributeRequest) SetGroupId(v string) *ModifyApiMarketInstanceAttributeRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiMarketInstanceAttributeRequest) SetInstanceAttributes(v string) *ModifyApiMarketInstanceAttributeRequest {
	s.InstanceAttributes = &v
	return s
}

type ModifyApiMarketInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiMarketInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiMarketInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiMarketInstanceAttributeResponseBody) SetRequestId(v string) *ModifyApiMarketInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyApiMarketInstanceAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiMarketInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiMarketInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiMarketInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiMarketInstanceAttributeResponse) SetHeaders(v map[string]*string) *ModifyApiMarketInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiMarketInstanceAttributeResponse) SetStatusCode(v int32) *ModifyApiMarketInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiMarketInstanceAttributeResponse) SetBody(v *ModifyApiMarketInstanceAttributeResponseBody) *ModifyApiMarketInstanceAttributeResponse {
	s.Body = v
	return s
}

type ModifyAppRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) SetAppId(v int64) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetDescription(v string) *ModifyAppRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppRequest) SetSecurityToken(v string) *ModifyAppRequest {
	s.SecurityToken = &v
	return s
}

type ModifyAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppResponse) SetStatusCode(v int32) *ModifyAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyAppForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Extend        *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyAppForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppForInnerRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppForInnerRequest) SetAliUid(v int64) *ModifyAppForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *ModifyAppForInnerRequest) SetAppId(v int64) *ModifyAppForInnerRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppForInnerRequest) SetAppName(v string) *ModifyAppForInnerRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppForInnerRequest) SetDescription(v string) *ModifyAppForInnerRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppForInnerRequest) SetExtend(v string) *ModifyAppForInnerRequest {
	s.Extend = &v
	return s
}

func (s *ModifyAppForInnerRequest) SetSecurityToken(v string) *ModifyAppForInnerRequest {
	s.SecurityToken = &v
	return s
}

type ModifyAppForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppForInnerResponseBody) SetRequestId(v string) *ModifyAppForInnerResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppForInnerResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppForInnerResponse) SetHeaders(v map[string]*string) *ModifyAppForInnerResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppForInnerResponse) SetStatusCode(v int32) *ModifyAppForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppForInnerResponse) SetBody(v *ModifyAppForInnerResponseBody) *ModifyAppForInnerResponse {
	s.Body = v
	return s
}

type ModifyGroupAuthAppCodeForBackendRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AuthAppCode   *string `json:"AuthAppCode,omitempty" xml:"AuthAppCode,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyGroupAuthAppCodeForBackendRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupAuthAppCodeForBackendRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupAuthAppCodeForBackendRequest) SetAliUid(v int64) *ModifyGroupAuthAppCodeForBackendRequest {
	s.AliUid = &v
	return s
}

func (s *ModifyGroupAuthAppCodeForBackendRequest) SetAuthAppCode(v string) *ModifyGroupAuthAppCodeForBackendRequest {
	s.AuthAppCode = &v
	return s
}

func (s *ModifyGroupAuthAppCodeForBackendRequest) SetGroupId(v string) *ModifyGroupAuthAppCodeForBackendRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyGroupAuthAppCodeForBackendRequest) SetSecurityToken(v string) *ModifyGroupAuthAppCodeForBackendRequest {
	s.SecurityToken = &v
	return s
}

type ModifyGroupAuthAppCodeForBackendResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupAuthAppCodeForBackendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupAuthAppCodeForBackendResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupAuthAppCodeForBackendResponseBody) SetRequestId(v string) *ModifyGroupAuthAppCodeForBackendResponseBody {
	s.RequestId = &v
	return s
}

type ModifyGroupAuthAppCodeForBackendResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGroupAuthAppCodeForBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGroupAuthAppCodeForBackendResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupAuthAppCodeForBackendResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupAuthAppCodeForBackendResponse) SetHeaders(v map[string]*string) *ModifyGroupAuthAppCodeForBackendResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupAuthAppCodeForBackendResponse) SetStatusCode(v int32) *ModifyGroupAuthAppCodeForBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGroupAuthAppCodeForBackendResponse) SetBody(v *ModifyGroupAuthAppCodeForBackendResponseBody) *ModifyGroupAuthAppCodeForBackendResponse {
	s.Body = v
	return s
}

type ModifySecretKeyRequest struct {
	SecretKey     *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	SecretKeyId   *string `json:"SecretKeyId,omitempty" xml:"SecretKeyId,omitempty"`
	SecretKeyName *string `json:"SecretKeyName,omitempty" xml:"SecretKeyName,omitempty"`
	SecretValue   *string `json:"SecretValue,omitempty" xml:"SecretValue,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifySecretKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecretKeyRequest) GoString() string {
	return s.String()
}

func (s *ModifySecretKeyRequest) SetSecretKey(v string) *ModifySecretKeyRequest {
	s.SecretKey = &v
	return s
}

func (s *ModifySecretKeyRequest) SetSecretKeyId(v string) *ModifySecretKeyRequest {
	s.SecretKeyId = &v
	return s
}

func (s *ModifySecretKeyRequest) SetSecretKeyName(v string) *ModifySecretKeyRequest {
	s.SecretKeyName = &v
	return s
}

func (s *ModifySecretKeyRequest) SetSecretValue(v string) *ModifySecretKeyRequest {
	s.SecretValue = &v
	return s
}

func (s *ModifySecretKeyRequest) SetSecurityToken(v string) *ModifySecretKeyRequest {
	s.SecurityToken = &v
	return s
}

type ModifySecretKeyResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecretKeyId   *string `json:"SecretKeyId,omitempty" xml:"SecretKeyId,omitempty"`
	SecretKeyName *string `json:"SecretKeyName,omitempty" xml:"SecretKeyName,omitempty"`
}

func (s ModifySecretKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecretKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecretKeyResponseBody) SetRequestId(v string) *ModifySecretKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecretKeyResponseBody) SetSecretKeyId(v string) *ModifySecretKeyResponseBody {
	s.SecretKeyId = &v
	return s
}

func (s *ModifySecretKeyResponseBody) SetSecretKeyName(v string) *ModifySecretKeyResponseBody {
	s.SecretKeyName = &v
	return s
}

type ModifySecretKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySecretKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecretKeyResponse) GoString() string {
	return s.String()
}

func (s *ModifySecretKeyResponse) SetHeaders(v map[string]*string) *ModifySecretKeyResponse {
	s.Headers = v
	return s
}

func (s *ModifySecretKeyResponse) SetStatusCode(v int32) *ModifySecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecretKeyResponse) SetBody(v *ModifySecretKeyResponseBody) *ModifySecretKeyResponse {
	s.Body = v
	return s
}

type ModifyTrafficControlRequest struct {
	ApiDefault         *int32  `json:"ApiDefault,omitempty" xml:"ApiDefault,omitempty"`
	AppDefault         *int32  `json:"AppDefault,omitempty" xml:"AppDefault,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TrafficControlId   *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
	TrafficControlUnit *string `json:"TrafficControlUnit,omitempty" xml:"TrafficControlUnit,omitempty"`
	UserDefault        *int32  `json:"UserDefault,omitempty" xml:"UserDefault,omitempty"`
}

func (s ModifyTrafficControlRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *ModifyTrafficControlRequest) SetApiDefault(v int32) *ModifyTrafficControlRequest {
	s.ApiDefault = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetAppDefault(v int32) *ModifyTrafficControlRequest {
	s.AppDefault = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetDescription(v string) *ModifyTrafficControlRequest {
	s.Description = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetSecurityToken(v string) *ModifyTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetTrafficControlId(v string) *ModifyTrafficControlRequest {
	s.TrafficControlId = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetTrafficControlName(v string) *ModifyTrafficControlRequest {
	s.TrafficControlName = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetTrafficControlUnit(v string) *ModifyTrafficControlRequest {
	s.TrafficControlUnit = &v
	return s
}

func (s *ModifyTrafficControlRequest) SetUserDefault(v int32) *ModifyTrafficControlRequest {
	s.UserDefault = &v
	return s
}

type ModifyTrafficControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTrafficControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrafficControlResponseBody) SetRequestId(v string) *ModifyTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTrafficControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrafficControlResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrafficControlResponse) SetHeaders(v map[string]*string) *ModifyTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrafficControlResponse) SetStatusCode(v int32) *ModifyTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrafficControlResponse) SetBody(v *ModifyTrafficControlResponseBody) *ModifyTrafficControlResponse {
	s.Body = v
	return s
}

type ModifyUserWhiteListValueByTypeRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Uid           *int64  `json:"Uid,omitempty" xml:"Uid,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyUserWhiteListValueByTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserWhiteListValueByTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserWhiteListValueByTypeRequest) SetDescription(v string) *ModifyUserWhiteListValueByTypeRequest {
	s.Description = &v
	return s
}

func (s *ModifyUserWhiteListValueByTypeRequest) SetSecurityToken(v string) *ModifyUserWhiteListValueByTypeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyUserWhiteListValueByTypeRequest) SetType(v string) *ModifyUserWhiteListValueByTypeRequest {
	s.Type = &v
	return s
}

func (s *ModifyUserWhiteListValueByTypeRequest) SetUid(v int64) *ModifyUserWhiteListValueByTypeRequest {
	s.Uid = &v
	return s
}

func (s *ModifyUserWhiteListValueByTypeRequest) SetValue(v string) *ModifyUserWhiteListValueByTypeRequest {
	s.Value = &v
	return s
}

type ModifyUserWhiteListValueByTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserWhiteListValueByTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserWhiteListValueByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserWhiteListValueByTypeResponseBody) SetRequestId(v string) *ModifyUserWhiteListValueByTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserWhiteListValueByTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserWhiteListValueByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserWhiteListValueByTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserWhiteListValueByTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserWhiteListValueByTypeResponse) SetHeaders(v map[string]*string) *ModifyUserWhiteListValueByTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserWhiteListValueByTypeResponse) SetStatusCode(v int32) *ModifyUserWhiteListValueByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserWhiteListValueByTypeResponse) SetBody(v *ModifyUserWhiteListValueByTypeResponseBody) *ModifyUserWhiteListValueByTypeResponse {
	s.Body = v
	return s
}

type ReactivateDomainForBackendRequest struct {
	Aliuid        *int64  `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReactivateDomainForBackendRequest) String() string {
	return tea.Prettify(s)
}

func (s ReactivateDomainForBackendRequest) GoString() string {
	return s.String()
}

func (s *ReactivateDomainForBackendRequest) SetAliuid(v int64) *ReactivateDomainForBackendRequest {
	s.Aliuid = &v
	return s
}

func (s *ReactivateDomainForBackendRequest) SetDomainName(v string) *ReactivateDomainForBackendRequest {
	s.DomainName = &v
	return s
}

func (s *ReactivateDomainForBackendRequest) SetGroupId(v string) *ReactivateDomainForBackendRequest {
	s.GroupId = &v
	return s
}

func (s *ReactivateDomainForBackendRequest) SetSecurityToken(v string) *ReactivateDomainForBackendRequest {
	s.SecurityToken = &v
	return s
}

type ReactivateDomainForBackendResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReactivateDomainForBackendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReactivateDomainForBackendResponseBody) GoString() string {
	return s.String()
}

func (s *ReactivateDomainForBackendResponseBody) SetRequestId(v string) *ReactivateDomainForBackendResponseBody {
	s.RequestId = &v
	return s
}

type ReactivateDomainForBackendResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReactivateDomainForBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReactivateDomainForBackendResponse) String() string {
	return tea.Prettify(s)
}

func (s ReactivateDomainForBackendResponse) GoString() string {
	return s.String()
}

func (s *ReactivateDomainForBackendResponse) SetHeaders(v map[string]*string) *ReactivateDomainForBackendResponse {
	s.Headers = v
	return s
}

func (s *ReactivateDomainForBackendResponse) SetStatusCode(v int32) *ReactivateDomainForBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *ReactivateDomainForBackendResponse) SetBody(v *ReactivateDomainForBackendResponseBody) *ReactivateDomainForBackendResponse {
	s.Body = v
	return s
}

type RecoverApiFromHistoricalRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RecoverApiFromHistoricalRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverApiFromHistoricalRequest) GoString() string {
	return s.String()
}

func (s *RecoverApiFromHistoricalRequest) SetApiId(v string) *RecoverApiFromHistoricalRequest {
	s.ApiId = &v
	return s
}

func (s *RecoverApiFromHistoricalRequest) SetGroupId(v string) *RecoverApiFromHistoricalRequest {
	s.GroupId = &v
	return s
}

func (s *RecoverApiFromHistoricalRequest) SetHistoryVersion(v string) *RecoverApiFromHistoricalRequest {
	s.HistoryVersion = &v
	return s
}

func (s *RecoverApiFromHistoricalRequest) SetSecurityToken(v string) *RecoverApiFromHistoricalRequest {
	s.SecurityToken = &v
	return s
}

func (s *RecoverApiFromHistoricalRequest) SetStageName(v string) *RecoverApiFromHistoricalRequest {
	s.StageName = &v
	return s
}

type RecoverApiFromHistoricalResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverApiFromHistoricalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverApiFromHistoricalResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverApiFromHistoricalResponseBody) SetRequestId(v string) *RecoverApiFromHistoricalResponseBody {
	s.RequestId = &v
	return s
}

type RecoverApiFromHistoricalResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoverApiFromHistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoverApiFromHistoricalResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverApiFromHistoricalResponse) GoString() string {
	return s.String()
}

func (s *RecoverApiFromHistoricalResponse) SetHeaders(v map[string]*string) *RecoverApiFromHistoricalResponse {
	s.Headers = v
	return s
}

func (s *RecoverApiFromHistoricalResponse) SetStatusCode(v int32) *RecoverApiFromHistoricalResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoverApiFromHistoricalResponse) SetBody(v *RecoverApiFromHistoricalResponseBody) *RecoverApiFromHistoricalResponse {
	s.Body = v
	return s
}

type RecoveryApiDefineFromHistoricalRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RecoveryApiDefineFromHistoricalRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoveryApiDefineFromHistoricalRequest) GoString() string {
	return s.String()
}

func (s *RecoveryApiDefineFromHistoricalRequest) SetApiId(v string) *RecoveryApiDefineFromHistoricalRequest {
	s.ApiId = &v
	return s
}

func (s *RecoveryApiDefineFromHistoricalRequest) SetGroupId(v string) *RecoveryApiDefineFromHistoricalRequest {
	s.GroupId = &v
	return s
}

func (s *RecoveryApiDefineFromHistoricalRequest) SetHistoryVersion(v string) *RecoveryApiDefineFromHistoricalRequest {
	s.HistoryVersion = &v
	return s
}

func (s *RecoveryApiDefineFromHistoricalRequest) SetSecurityToken(v string) *RecoveryApiDefineFromHistoricalRequest {
	s.SecurityToken = &v
	return s
}

func (s *RecoveryApiDefineFromHistoricalRequest) SetStageName(v string) *RecoveryApiDefineFromHistoricalRequest {
	s.StageName = &v
	return s
}

type RecoveryApiDefineFromHistoricalResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoveryApiDefineFromHistoricalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoveryApiDefineFromHistoricalResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryApiDefineFromHistoricalResponseBody) SetRequestId(v string) *RecoveryApiDefineFromHistoricalResponseBody {
	s.RequestId = &v
	return s
}

type RecoveryApiDefineFromHistoricalResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoveryApiDefineFromHistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoveryApiDefineFromHistoricalResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoveryApiDefineFromHistoricalResponse) GoString() string {
	return s.String()
}

func (s *RecoveryApiDefineFromHistoricalResponse) SetHeaders(v map[string]*string) *RecoveryApiDefineFromHistoricalResponse {
	s.Headers = v
	return s
}

func (s *RecoveryApiDefineFromHistoricalResponse) SetStatusCode(v int32) *RecoveryApiDefineFromHistoricalResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryApiDefineFromHistoricalResponse) SetBody(v *RecoveryApiDefineFromHistoricalResponseBody) *RecoveryApiDefineFromHistoricalResponse {
	s.Body = v
	return s
}

type RecoveryApiFromHistoricalRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RecoveryApiFromHistoricalRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoveryApiFromHistoricalRequest) GoString() string {
	return s.String()
}

func (s *RecoveryApiFromHistoricalRequest) SetApiId(v string) *RecoveryApiFromHistoricalRequest {
	s.ApiId = &v
	return s
}

func (s *RecoveryApiFromHistoricalRequest) SetGroupId(v string) *RecoveryApiFromHistoricalRequest {
	s.GroupId = &v
	return s
}

func (s *RecoveryApiFromHistoricalRequest) SetHistoryVersion(v string) *RecoveryApiFromHistoricalRequest {
	s.HistoryVersion = &v
	return s
}

func (s *RecoveryApiFromHistoricalRequest) SetSecurityToken(v string) *RecoveryApiFromHistoricalRequest {
	s.SecurityToken = &v
	return s
}

func (s *RecoveryApiFromHistoricalRequest) SetStageName(v string) *RecoveryApiFromHistoricalRequest {
	s.StageName = &v
	return s
}

type RecoveryApiFromHistoricalResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoveryApiFromHistoricalResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoveryApiFromHistoricalResponseBody) GoString() string {
	return s.String()
}

func (s *RecoveryApiFromHistoricalResponseBody) SetRequestId(v string) *RecoveryApiFromHistoricalResponseBody {
	s.RequestId = &v
	return s
}

type RecoveryApiFromHistoricalResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RecoveryApiFromHistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RecoveryApiFromHistoricalResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoveryApiFromHistoricalResponse) GoString() string {
	return s.String()
}

func (s *RecoveryApiFromHistoricalResponse) SetHeaders(v map[string]*string) *RecoveryApiFromHistoricalResponse {
	s.Headers = v
	return s
}

func (s *RecoveryApiFromHistoricalResponse) SetStatusCode(v int32) *RecoveryApiFromHistoricalResponse {
	s.StatusCode = &v
	return s
}

func (s *RecoveryApiFromHistoricalResponse) SetBody(v *RecoveryApiFromHistoricalResponseBody) *RecoveryApiFromHistoricalResponse {
	s.Body = v
	return s
}

type RefreshDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RefreshDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshDomainRequest) GoString() string {
	return s.String()
}

func (s *RefreshDomainRequest) SetDomainName(v string) *RefreshDomainRequest {
	s.DomainName = &v
	return s
}

func (s *RefreshDomainRequest) SetGroupId(v string) *RefreshDomainRequest {
	s.GroupId = &v
	return s
}

func (s *RefreshDomainRequest) SetSecurityToken(v string) *RefreshDomainRequest {
	s.SecurityToken = &v
	return s
}

type RefreshDomainResponseBody struct {
	CertificateId        *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName      *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
	DomainStatus         *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain            *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s RefreshDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshDomainResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDomainResponseBody) SetCertificateId(v string) *RefreshDomainResponseBody {
	s.CertificateId = &v
	return s
}

func (s *RefreshDomainResponseBody) SetCertificateName(v string) *RefreshDomainResponseBody {
	s.CertificateName = &v
	return s
}

func (s *RefreshDomainResponseBody) SetDomainName(v string) *RefreshDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *RefreshDomainResponseBody) SetDomainNameResolution(v string) *RefreshDomainResponseBody {
	s.DomainNameResolution = &v
	return s
}

func (s *RefreshDomainResponseBody) SetDomainStatus(v string) *RefreshDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *RefreshDomainResponseBody) SetGroupId(v string) *RefreshDomainResponseBody {
	s.GroupId = &v
	return s
}

func (s *RefreshDomainResponseBody) SetRequestId(v string) *RefreshDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDomainResponseBody) SetSubDomain(v string) *RefreshDomainResponseBody {
	s.SubDomain = &v
	return s
}

type RefreshDomainResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RefreshDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RefreshDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshDomainResponse) GoString() string {
	return s.String()
}

func (s *RefreshDomainResponse) SetHeaders(v map[string]*string) *RefreshDomainResponse {
	s.Headers = v
	return s
}

func (s *RefreshDomainResponse) SetStatusCode(v int32) *RefreshDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshDomainResponse) SetBody(v *RefreshDomainResponseBody) *RefreshDomainResponse {
	s.Body = v
	return s
}

type RemoveAccessPermissionByApisRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveAccessPermissionByApisRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByApisRequest) SetApiIds(v string) *RemoveAccessPermissionByApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveAccessPermissionByApisRequest) SetAppId(v int64) *RemoveAccessPermissionByApisRequest {
	s.AppId = &v
	return s
}

func (s *RemoveAccessPermissionByApisRequest) SetDescription(v string) *RemoveAccessPermissionByApisRequest {
	s.Description = &v
	return s
}

func (s *RemoveAccessPermissionByApisRequest) SetGroupId(v string) *RemoveAccessPermissionByApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveAccessPermissionByApisRequest) SetSecurityToken(v string) *RemoveAccessPermissionByApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveAccessPermissionByApisRequest) SetStageName(v string) *RemoveAccessPermissionByApisRequest {
	s.StageName = &v
	return s
}

type RemoveAccessPermissionByApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAccessPermissionByApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByApisResponseBody) SetRequestId(v string) *RemoveAccessPermissionByApisResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAccessPermissionByApisResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAccessPermissionByApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAccessPermissionByApisResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByApisResponse) SetHeaders(v map[string]*string) *RemoveAccessPermissionByApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveAccessPermissionByApisResponse) SetStatusCode(v int32) *RemoveAccessPermissionByApisResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAccessPermissionByApisResponse) SetBody(v *RemoveAccessPermissionByApisResponseBody) *RemoveAccessPermissionByApisResponse {
	s.Body = v
	return s
}

type RemoveAccessPermissionByAppsRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveAccessPermissionByAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByAppsRequest) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByAppsRequest) SetApiId(v string) *RemoveAccessPermissionByAppsRequest {
	s.ApiId = &v
	return s
}

func (s *RemoveAccessPermissionByAppsRequest) SetAppIds(v string) *RemoveAccessPermissionByAppsRequest {
	s.AppIds = &v
	return s
}

func (s *RemoveAccessPermissionByAppsRequest) SetGroupId(v string) *RemoveAccessPermissionByAppsRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveAccessPermissionByAppsRequest) SetSecurityToken(v string) *RemoveAccessPermissionByAppsRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveAccessPermissionByAppsRequest) SetStageName(v string) *RemoveAccessPermissionByAppsRequest {
	s.StageName = &v
	return s
}

type RemoveAccessPermissionByAppsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAccessPermissionByAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByAppsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByAppsResponseBody) SetRequestId(v string) *RemoveAccessPermissionByAppsResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAccessPermissionByAppsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAccessPermissionByAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAccessPermissionByAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByAppsResponse) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByAppsResponse) SetHeaders(v map[string]*string) *RemoveAccessPermissionByAppsResponse {
	s.Headers = v
	return s
}

func (s *RemoveAccessPermissionByAppsResponse) SetStatusCode(v int32) *RemoveAccessPermissionByAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAccessPermissionByAppsResponse) SetBody(v *RemoveAccessPermissionByAppsResponseBody) *RemoveAccessPermissionByAppsResponse {
	s.Body = v
	return s
}

type RemoveAccessPermissionByAppsForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveAccessPermissionByAppsForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByAppsForInnerRequest) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetAliUid(v int64) *RemoveAccessPermissionByAppsForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetApiId(v string) *RemoveAccessPermissionByAppsForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetAppIds(v string) *RemoveAccessPermissionByAppsForInnerRequest {
	s.AppIds = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetGroupId(v string) *RemoveAccessPermissionByAppsForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetSecurityToken(v string) *RemoveAccessPermissionByAppsForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetSource(v string) *RemoveAccessPermissionByAppsForInnerRequest {
	s.Source = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerRequest) SetStageName(v string) *RemoveAccessPermissionByAppsForInnerRequest {
	s.StageName = &v
	return s
}

type RemoveAccessPermissionByAppsForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAccessPermissionByAppsForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByAppsForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByAppsForInnerResponseBody) SetRequestId(v string) *RemoveAccessPermissionByAppsForInnerResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAccessPermissionByAppsForInnerResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAccessPermissionByAppsForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAccessPermissionByAppsForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAccessPermissionByAppsForInnerResponse) GoString() string {
	return s.String()
}

func (s *RemoveAccessPermissionByAppsForInnerResponse) SetHeaders(v map[string]*string) *RemoveAccessPermissionByAppsForInnerResponse {
	s.Headers = v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerResponse) SetStatusCode(v int32) *RemoveAccessPermissionByAppsForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAccessPermissionByAppsForInnerResponse) SetBody(v *RemoveAccessPermissionByAppsForInnerResponseBody) *RemoveAccessPermissionByAppsForInnerResponse {
	s.Body = v
	return s
}

type RemoveAllBlackListRequest struct {
	BlackType     *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveAllBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAllBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveAllBlackListRequest) SetBlackType(v string) *RemoveAllBlackListRequest {
	s.BlackType = &v
	return s
}

func (s *RemoveAllBlackListRequest) SetSecurityToken(v string) *RemoveAllBlackListRequest {
	s.SecurityToken = &v
	return s
}

type RemoveAllBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAllBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAllBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAllBlackListResponseBody) SetRequestId(v string) *RemoveAllBlackListResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAllBlackListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAllBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAllBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAllBlackListResponse) GoString() string {
	return s.String()
}

func (s *RemoveAllBlackListResponse) SetHeaders(v map[string]*string) *RemoveAllBlackListResponse {
	s.Headers = v
	return s
}

func (s *RemoveAllBlackListResponse) SetStatusCode(v int32) *RemoveAllBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAllBlackListResponse) SetBody(v *RemoveAllBlackListResponseBody) *RemoveAllBlackListResponse {
	s.Body = v
	return s
}

type RemoveApiRuleRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveApiRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveApiRuleRequest) GoString() string {
	return s.String()
}

func (s *RemoveApiRuleRequest) SetApiIds(v string) *RemoveApiRuleRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveApiRuleRequest) SetGroupId(v string) *RemoveApiRuleRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveApiRuleRequest) SetRuleId(v string) *RemoveApiRuleRequest {
	s.RuleId = &v
	return s
}

func (s *RemoveApiRuleRequest) SetRuleType(v string) *RemoveApiRuleRequest {
	s.RuleType = &v
	return s
}

func (s *RemoveApiRuleRequest) SetSecurityToken(v string) *RemoveApiRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveApiRuleRequest) SetStageName(v string) *RemoveApiRuleRequest {
	s.StageName = &v
	return s
}

type RemoveApiRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveApiRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveApiRuleResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApiRuleResponseBody) SetRequestId(v string) *RemoveApiRuleResponseBody {
	s.RequestId = &v
	return s
}

type RemoveApiRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApiRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApiRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveApiRuleResponse) GoString() string {
	return s.String()
}

func (s *RemoveApiRuleResponse) SetHeaders(v map[string]*string) *RemoveApiRuleResponse {
	s.Headers = v
	return s
}

func (s *RemoveApiRuleResponse) SetStatusCode(v int32) *RemoveApiRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApiRuleResponse) SetBody(v *RemoveApiRuleResponseBody) *RemoveApiRuleResponse {
	s.Body = v
	return s
}

type RemoveAppsFromApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveAppsFromApiRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppsFromApiRequest) GoString() string {
	return s.String()
}

func (s *RemoveAppsFromApiRequest) SetApiId(v string) *RemoveAppsFromApiRequest {
	s.ApiId = &v
	return s
}

func (s *RemoveAppsFromApiRequest) SetAppIds(v string) *RemoveAppsFromApiRequest {
	s.AppIds = &v
	return s
}

func (s *RemoveAppsFromApiRequest) SetGroupId(v string) *RemoveAppsFromApiRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveAppsFromApiRequest) SetSecurityToken(v string) *RemoveAppsFromApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveAppsFromApiRequest) SetStageName(v string) *RemoveAppsFromApiRequest {
	s.StageName = &v
	return s
}

type RemoveAppsFromApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAppsFromApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppsFromApiResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppsFromApiResponseBody) SetRequestId(v string) *RemoveAppsFromApiResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAppsFromApiResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAppsFromApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveAppsFromApiResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppsFromApiResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppsFromApiResponse) SetHeaders(v map[string]*string) *RemoveAppsFromApiResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppsFromApiResponse) SetStatusCode(v int32) *RemoveAppsFromApiResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAppsFromApiResponse) SetBody(v *RemoveAppsFromApiResponseBody) *RemoveAppsFromApiResponse {
	s.Body = v
	return s
}

type RemoveBlackListRequest struct {
	BlackContent  *string `json:"BlackContent,omitempty" xml:"BlackContent,omitempty"`
	BlackType     *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveBlackListRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveBlackListRequest) GoString() string {
	return s.String()
}

func (s *RemoveBlackListRequest) SetBlackContent(v string) *RemoveBlackListRequest {
	s.BlackContent = &v
	return s
}

func (s *RemoveBlackListRequest) SetBlackType(v string) *RemoveBlackListRequest {
	s.BlackType = &v
	return s
}

func (s *RemoveBlackListRequest) SetSecurityToken(v string) *RemoveBlackListRequest {
	s.SecurityToken = &v
	return s
}

type RemoveBlackListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveBlackListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveBlackListResponseBody) SetRequestId(v string) *RemoveBlackListResponseBody {
	s.RequestId = &v
	return s
}

type RemoveBlackListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveBlackListResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveBlackListResponse) GoString() string {
	return s.String()
}

func (s *RemoveBlackListResponse) SetHeaders(v map[string]*string) *RemoveBlackListResponse {
	s.Headers = v
	return s
}

func (s *RemoveBlackListResponse) SetStatusCode(v int32) *RemoveBlackListResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveBlackListResponse) SetBody(v *RemoveBlackListResponseBody) *RemoveBlackListResponse {
	s.Body = v
	return s
}

type ResetAppCodeRequest struct {
	AppCode       *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetAppCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAppCodeRequest) GoString() string {
	return s.String()
}

func (s *ResetAppCodeRequest) SetAppCode(v string) *ResetAppCodeRequest {
	s.AppCode = &v
	return s
}

func (s *ResetAppCodeRequest) SetSecurityToken(v string) *ResetAppCodeRequest {
	s.SecurityToken = &v
	return s
}

type ResetAppCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAppCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppCodeResponseBody) SetRequestId(v string) *ResetAppCodeResponseBody {
	s.RequestId = &v
	return s
}

type ResetAppCodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAppCodeResponse) GoString() string {
	return s.String()
}

func (s *ResetAppCodeResponse) SetHeaders(v map[string]*string) *ResetAppCodeResponse {
	s.Headers = v
	return s
}

func (s *ResetAppCodeResponse) SetStatusCode(v int32) *ResetAppCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppCodeResponse) SetBody(v *ResetAppCodeResponseBody) *ResetAppCodeResponse {
	s.Body = v
	return s
}

type ResetAppCodeForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppCode       *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	NewAppCode    *string `json:"NewAppCode,omitempty" xml:"NewAppCode,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetAppCodeForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAppCodeForInnerRequest) GoString() string {
	return s.String()
}

func (s *ResetAppCodeForInnerRequest) SetAliUid(v int64) *ResetAppCodeForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *ResetAppCodeForInnerRequest) SetAppCode(v string) *ResetAppCodeForInnerRequest {
	s.AppCode = &v
	return s
}

func (s *ResetAppCodeForInnerRequest) SetNewAppCode(v string) *ResetAppCodeForInnerRequest {
	s.NewAppCode = &v
	return s
}

func (s *ResetAppCodeForInnerRequest) SetSecurityToken(v string) *ResetAppCodeForInnerRequest {
	s.SecurityToken = &v
	return s
}

type ResetAppCodeForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppCodeForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAppCodeForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppCodeForInnerResponseBody) SetRequestId(v string) *ResetAppCodeForInnerResponseBody {
	s.RequestId = &v
	return s
}

type ResetAppCodeForInnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppCodeForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppCodeForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAppCodeForInnerResponse) GoString() string {
	return s.String()
}

func (s *ResetAppCodeForInnerResponse) SetHeaders(v map[string]*string) *ResetAppCodeForInnerResponse {
	s.Headers = v
	return s
}

func (s *ResetAppCodeForInnerResponse) SetStatusCode(v int32) *ResetAppCodeForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppCodeForInnerResponse) SetBody(v *ResetAppCodeForInnerResponseBody) *ResetAppCodeForInnerResponse {
	s.Body = v
	return s
}

type ResetAppKeySecretRequest struct {
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetAppKeySecretRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAppKeySecretRequest) GoString() string {
	return s.String()
}

func (s *ResetAppKeySecretRequest) SetAppKey(v string) *ResetAppKeySecretRequest {
	s.AppKey = &v
	return s
}

func (s *ResetAppKeySecretRequest) SetSecurityToken(v string) *ResetAppKeySecretRequest {
	s.SecurityToken = &v
	return s
}

type ResetAppKeySecretResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppKeySecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAppKeySecretResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppKeySecretResponseBody) SetRequestId(v string) *ResetAppKeySecretResponseBody {
	s.RequestId = &v
	return s
}

type ResetAppKeySecretResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetAppKeySecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetAppKeySecretResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAppKeySecretResponse) GoString() string {
	return s.String()
}

func (s *ResetAppKeySecretResponse) SetHeaders(v map[string]*string) *ResetAppKeySecretResponse {
	s.Headers = v
	return s
}

func (s *ResetAppKeySecretResponse) SetStatusCode(v int32) *ResetAppKeySecretResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetAppKeySecretResponse) SetBody(v *ResetAppKeySecretResponseBody) *ResetAppKeySecretResponse {
	s.Body = v
	return s
}

type ResetSecretByAppKeyForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	NewAppKey     *string `json:"NewAppKey,omitempty" xml:"NewAppKey,omitempty"`
	NewAppSecret  *string `json:"NewAppSecret,omitempty" xml:"NewAppSecret,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetSecretByAppKeyForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSecretByAppKeyForInnerRequest) GoString() string {
	return s.String()
}

func (s *ResetSecretByAppKeyForInnerRequest) SetAliUid(v int64) *ResetSecretByAppKeyForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *ResetSecretByAppKeyForInnerRequest) SetAppKey(v string) *ResetSecretByAppKeyForInnerRequest {
	s.AppKey = &v
	return s
}

func (s *ResetSecretByAppKeyForInnerRequest) SetNewAppKey(v string) *ResetSecretByAppKeyForInnerRequest {
	s.NewAppKey = &v
	return s
}

func (s *ResetSecretByAppKeyForInnerRequest) SetNewAppSecret(v string) *ResetSecretByAppKeyForInnerRequest {
	s.NewAppSecret = &v
	return s
}

func (s *ResetSecretByAppKeyForInnerRequest) SetSecurityToken(v string) *ResetSecretByAppKeyForInnerRequest {
	s.SecurityToken = &v
	return s
}

type ResetSecretByAppKeyForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetSecretByAppKeyForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSecretByAppKeyForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSecretByAppKeyForInnerResponseBody) SetRequestId(v string) *ResetSecretByAppKeyForInnerResponseBody {
	s.RequestId = &v
	return s
}

type ResetSecretByAppKeyForInnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetSecretByAppKeyForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetSecretByAppKeyForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSecretByAppKeyForInnerResponse) GoString() string {
	return s.String()
}

func (s *ResetSecretByAppKeyForInnerResponse) SetHeaders(v map[string]*string) *ResetSecretByAppKeyForInnerResponse {
	s.Headers = v
	return s
}

func (s *ResetSecretByAppKeyForInnerResponse) SetStatusCode(v int32) *ResetSecretByAppKeyForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSecretByAppKeyForInnerResponse) SetBody(v *ResetSecretByAppKeyForInnerResponseBody) *ResetSecretByAppKeyForInnerResponse {
	s.Body = v
	return s
}

type SetAccessPermissionByApisRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetAccessPermissionByApisRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionByApisRequest) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionByApisRequest) SetApiIds(v string) *SetAccessPermissionByApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetAccessPermissionByApisRequest) SetAppId(v int64) *SetAccessPermissionByApisRequest {
	s.AppId = &v
	return s
}

func (s *SetAccessPermissionByApisRequest) SetDescription(v string) *SetAccessPermissionByApisRequest {
	s.Description = &v
	return s
}

func (s *SetAccessPermissionByApisRequest) SetGroupId(v string) *SetAccessPermissionByApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetAccessPermissionByApisRequest) SetSecurityToken(v string) *SetAccessPermissionByApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetAccessPermissionByApisRequest) SetStageName(v string) *SetAccessPermissionByApisRequest {
	s.StageName = &v
	return s
}

type SetAccessPermissionByApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessPermissionByApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionByApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionByApisResponseBody) SetRequestId(v string) *SetAccessPermissionByApisResponseBody {
	s.RequestId = &v
	return s
}

type SetAccessPermissionByApisResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessPermissionByApisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessPermissionByApisResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionByApisResponse) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionByApisResponse) SetHeaders(v map[string]*string) *SetAccessPermissionByApisResponse {
	s.Headers = v
	return s
}

func (s *SetAccessPermissionByApisResponse) SetStatusCode(v int32) *SetAccessPermissionByApisResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessPermissionByApisResponse) SetBody(v *SetAccessPermissionByApisResponseBody) *SetAccessPermissionByApisResponse {
	s.Body = v
	return s
}

type SetAccessPermissionByGroupForBackendRequest struct {
	AppId   *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s SetAccessPermissionByGroupForBackendRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionByGroupForBackendRequest) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionByGroupForBackendRequest) SetAppId(v int64) *SetAccessPermissionByGroupForBackendRequest {
	s.AppId = &v
	return s
}

func (s *SetAccessPermissionByGroupForBackendRequest) SetGroupId(v string) *SetAccessPermissionByGroupForBackendRequest {
	s.GroupId = &v
	return s
}

type SetAccessPermissionByGroupForBackendResponseBody struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessPermissionByGroupForBackendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionByGroupForBackendResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionByGroupForBackendResponseBody) SetInstanceId(v string) *SetAccessPermissionByGroupForBackendResponseBody {
	s.InstanceId = &v
	return s
}

func (s *SetAccessPermissionByGroupForBackendResponseBody) SetRequestId(v string) *SetAccessPermissionByGroupForBackendResponseBody {
	s.RequestId = &v
	return s
}

type SetAccessPermissionByGroupForBackendResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessPermissionByGroupForBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessPermissionByGroupForBackendResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionByGroupForBackendResponse) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionByGroupForBackendResponse) SetHeaders(v map[string]*string) *SetAccessPermissionByGroupForBackendResponse {
	s.Headers = v
	return s
}

func (s *SetAccessPermissionByGroupForBackendResponse) SetStatusCode(v int32) *SetAccessPermissionByGroupForBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessPermissionByGroupForBackendResponse) SetBody(v *SetAccessPermissionByGroupForBackendResponseBody) *SetAccessPermissionByGroupForBackendResponse {
	s.Body = v
	return s
}

type SetAccessPermissionsRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetAccessPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionsRequest) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionsRequest) SetApiId(v string) *SetAccessPermissionsRequest {
	s.ApiId = &v
	return s
}

func (s *SetAccessPermissionsRequest) SetAppIds(v string) *SetAccessPermissionsRequest {
	s.AppIds = &v
	return s
}

func (s *SetAccessPermissionsRequest) SetDescription(v string) *SetAccessPermissionsRequest {
	s.Description = &v
	return s
}

func (s *SetAccessPermissionsRequest) SetGroupId(v string) *SetAccessPermissionsRequest {
	s.GroupId = &v
	return s
}

func (s *SetAccessPermissionsRequest) SetSecurityToken(v string) *SetAccessPermissionsRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetAccessPermissionsRequest) SetStageName(v string) *SetAccessPermissionsRequest {
	s.StageName = &v
	return s
}

type SetAccessPermissionsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionsResponseBody) SetRequestId(v string) *SetAccessPermissionsResponseBody {
	s.RequestId = &v
	return s
}

type SetAccessPermissionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionsResponse) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionsResponse) SetHeaders(v map[string]*string) *SetAccessPermissionsResponse {
	s.Headers = v
	return s
}

func (s *SetAccessPermissionsResponse) SetStatusCode(v int32) *SetAccessPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessPermissionsResponse) SetBody(v *SetAccessPermissionsResponseBody) *SetAccessPermissionsResponse {
	s.Body = v
	return s
}

type SetAccessPermissionsForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetAccessPermissionsForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionsForInnerRequest) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionsForInnerRequest) SetAliUid(v int64) *SetAccessPermissionsForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetApiId(v string) *SetAccessPermissionsForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetAppIds(v string) *SetAccessPermissionsForInnerRequest {
	s.AppIds = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetDescription(v string) *SetAccessPermissionsForInnerRequest {
	s.Description = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetGroupId(v string) *SetAccessPermissionsForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetSecurityToken(v string) *SetAccessPermissionsForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetSource(v string) *SetAccessPermissionsForInnerRequest {
	s.Source = &v
	return s
}

func (s *SetAccessPermissionsForInnerRequest) SetStageName(v string) *SetAccessPermissionsForInnerRequest {
	s.StageName = &v
	return s
}

type SetAccessPermissionsForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAccessPermissionsForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionsForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionsForInnerResponseBody) SetRequestId(v string) *SetAccessPermissionsForInnerResponseBody {
	s.RequestId = &v
	return s
}

type SetAccessPermissionsForInnerResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessPermissionsForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessPermissionsForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAccessPermissionsForInnerResponse) GoString() string {
	return s.String()
}

func (s *SetAccessPermissionsForInnerResponse) SetHeaders(v map[string]*string) *SetAccessPermissionsForInnerResponse {
	s.Headers = v
	return s
}

func (s *SetAccessPermissionsForInnerResponse) SetStatusCode(v int32) *SetAccessPermissionsForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessPermissionsForInnerResponse) SetBody(v *SetAccessPermissionsForInnerResponseBody) *SetAccessPermissionsForInnerResponse {
	s.Body = v
	return s
}

type SetApiRuleRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleType      *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetApiRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s SetApiRuleRequest) GoString() string {
	return s.String()
}

func (s *SetApiRuleRequest) SetApiIds(v string) *SetApiRuleRequest {
	s.ApiIds = &v
	return s
}

func (s *SetApiRuleRequest) SetGroupId(v string) *SetApiRuleRequest {
	s.GroupId = &v
	return s
}

func (s *SetApiRuleRequest) SetRuleId(v string) *SetApiRuleRequest {
	s.RuleId = &v
	return s
}

func (s *SetApiRuleRequest) SetRuleType(v string) *SetApiRuleRequest {
	s.RuleType = &v
	return s
}

func (s *SetApiRuleRequest) SetSecurityToken(v string) *SetApiRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetApiRuleRequest) SetStageName(v string) *SetApiRuleRequest {
	s.StageName = &v
	return s
}

type SetApiRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApiRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetApiRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SetApiRuleResponseBody) SetRequestId(v string) *SetApiRuleResponseBody {
	s.RequestId = &v
	return s
}

type SetApiRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetApiRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetApiRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s SetApiRuleResponse) GoString() string {
	return s.String()
}

func (s *SetApiRuleResponse) SetHeaders(v map[string]*string) *SetApiRuleResponse {
	s.Headers = v
	return s
}

func (s *SetApiRuleResponse) SetStatusCode(v int32) *SetApiRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetApiRuleResponse) SetBody(v *SetApiRuleResponseBody) *SetApiRuleResponse {
	s.Body = v
	return s
}

type SetDomainRequest struct {
	CertificateBody *string `json:"CertificateBody,omitempty" xml:"CertificateBody,omitempty"`
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDomainRequest) SetCertificateBody(v string) *SetDomainRequest {
	s.CertificateBody = &v
	return s
}

func (s *SetDomainRequest) SetCertificateName(v string) *SetDomainRequest {
	s.CertificateName = &v
	return s
}

func (s *SetDomainRequest) SetDomainName(v string) *SetDomainRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainRequest) SetGroupId(v string) *SetDomainRequest {
	s.GroupId = &v
	return s
}

func (s *SetDomainRequest) SetPrivateKey(v string) *SetDomainRequest {
	s.PrivateKey = &v
	return s
}

func (s *SetDomainRequest) SetSecurityToken(v string) *SetDomainRequest {
	s.SecurityToken = &v
	return s
}

type SetDomainResponseBody struct {
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain    *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s SetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainResponseBody) SetDomainName(v string) *SetDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *SetDomainResponseBody) SetDomainStatus(v string) *SetDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *SetDomainResponseBody) SetGroupId(v string) *SetDomainResponseBody {
	s.GroupId = &v
	return s
}

func (s *SetDomainResponseBody) SetRequestId(v string) *SetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainResponseBody) SetSubDomain(v string) *SetDomainResponseBody {
	s.SubDomain = &v
	return s
}

type SetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDomainResponse) SetHeaders(v map[string]*string) *SetDomainResponse {
	s.Headers = v
	return s
}

func (s *SetDomainResponse) SetStatusCode(v int32) *SetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainResponse) SetBody(v *SetDomainResponseBody) *SetDomainResponse {
	s.Body = v
	return s
}

type SetDomainCertificateRequest struct {
	CertificateBody *string `json:"CertificateBody,omitempty" xml:"CertificateBody,omitempty"`
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainName      *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PrivateKey      *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDomainCertificateRequest) SetCertificateBody(v string) *SetDomainCertificateRequest {
	s.CertificateBody = &v
	return s
}

func (s *SetDomainCertificateRequest) SetCertificateName(v string) *SetDomainCertificateRequest {
	s.CertificateName = &v
	return s
}

func (s *SetDomainCertificateRequest) SetDomainName(v string) *SetDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainCertificateRequest) SetGroupId(v string) *SetDomainCertificateRequest {
	s.GroupId = &v
	return s
}

func (s *SetDomainCertificateRequest) SetPrivateKey(v string) *SetDomainCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *SetDomainCertificateRequest) SetSecurityToken(v string) *SetDomainCertificateRequest {
	s.SecurityToken = &v
	return s
}

type SetDomainCertificateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainCertificateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainCertificateResponseBody) SetRequestId(v string) *SetDomainCertificateResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainCertificateResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainCertificateResponse) GoString() string {
	return s.String()
}

func (s *SetDomainCertificateResponse) SetHeaders(v map[string]*string) *SetDomainCertificateResponse {
	s.Headers = v
	return s
}

func (s *SetDomainCertificateResponse) SetStatusCode(v int32) *SetDomainCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainCertificateResponse) SetBody(v *SetDomainCertificateResponseBody) *SetDomainCertificateResponse {
	s.Body = v
	return s
}

type SetDomainForBackendRequest struct {
	CertificateBody       *string `json:"CertificateBody,omitempty" xml:"CertificateBody,omitempty"`
	CertificateName       *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" xml:"CertificatePrivateKey,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDomainForBackendRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainForBackendRequest) GoString() string {
	return s.String()
}

func (s *SetDomainForBackendRequest) SetCertificateBody(v string) *SetDomainForBackendRequest {
	s.CertificateBody = &v
	return s
}

func (s *SetDomainForBackendRequest) SetCertificateName(v string) *SetDomainForBackendRequest {
	s.CertificateName = &v
	return s
}

func (s *SetDomainForBackendRequest) SetCertificatePrivateKey(v string) *SetDomainForBackendRequest {
	s.CertificatePrivateKey = &v
	return s
}

func (s *SetDomainForBackendRequest) SetDomainName(v string) *SetDomainForBackendRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainForBackendRequest) SetGroupId(v string) *SetDomainForBackendRequest {
	s.GroupId = &v
	return s
}

func (s *SetDomainForBackendRequest) SetSecurityToken(v string) *SetDomainForBackendRequest {
	s.SecurityToken = &v
	return s
}

type SetDomainForBackendResponseBody struct {
	DomainBindingStatus *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	DomainName          *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain           *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s SetDomainForBackendResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainForBackendResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainForBackendResponseBody) SetDomainBindingStatus(v string) *SetDomainForBackendResponseBody {
	s.DomainBindingStatus = &v
	return s
}

func (s *SetDomainForBackendResponseBody) SetDomainName(v string) *SetDomainForBackendResponseBody {
	s.DomainName = &v
	return s
}

func (s *SetDomainForBackendResponseBody) SetGroupId(v string) *SetDomainForBackendResponseBody {
	s.GroupId = &v
	return s
}

func (s *SetDomainForBackendResponseBody) SetRequestId(v string) *SetDomainForBackendResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainForBackendResponseBody) SetSubDomain(v string) *SetDomainForBackendResponseBody {
	s.SubDomain = &v
	return s
}

type SetDomainForBackendResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainForBackendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainForBackendResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainForBackendResponse) GoString() string {
	return s.String()
}

func (s *SetDomainForBackendResponse) SetHeaders(v map[string]*string) *SetDomainForBackendResponse {
	s.Headers = v
	return s
}

func (s *SetDomainForBackendResponse) SetStatusCode(v int32) *SetDomainForBackendResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainForBackendResponse) SetBody(v *SetDomainForBackendResponseBody) *SetDomainForBackendResponse {
	s.Body = v
	return s
}

type SetMockIntegrationRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Mock          *string `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult    *string `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetMockIntegrationRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMockIntegrationRequest) GoString() string {
	return s.String()
}

func (s *SetMockIntegrationRequest) SetApiId(v string) *SetMockIntegrationRequest {
	s.ApiId = &v
	return s
}

func (s *SetMockIntegrationRequest) SetGroupId(v string) *SetMockIntegrationRequest {
	s.GroupId = &v
	return s
}

func (s *SetMockIntegrationRequest) SetMock(v string) *SetMockIntegrationRequest {
	s.Mock = &v
	return s
}

func (s *SetMockIntegrationRequest) SetMockResult(v string) *SetMockIntegrationRequest {
	s.MockResult = &v
	return s
}

func (s *SetMockIntegrationRequest) SetSecurityToken(v string) *SetMockIntegrationRequest {
	s.SecurityToken = &v
	return s
}

type SetMockIntegrationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetMockIntegrationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMockIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *SetMockIntegrationResponseBody) SetRequestId(v string) *SetMockIntegrationResponseBody {
	s.RequestId = &v
	return s
}

type SetMockIntegrationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMockIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMockIntegrationResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMockIntegrationResponse) GoString() string {
	return s.String()
}

func (s *SetMockIntegrationResponse) SetHeaders(v map[string]*string) *SetMockIntegrationResponse {
	s.Headers = v
	return s
}

func (s *SetMockIntegrationResponse) SetStatusCode(v int32) *SetMockIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMockIntegrationResponse) SetBody(v *SetMockIntegrationResponseBody) *SetMockIntegrationResponse {
	s.Body = v
	return s
}

type SetPurchasedApiGroupStatusRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	BillingStatus *string `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	CloseOrder    *bool   `json:"CloseOrder,omitempty" xml:"CloseOrder,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s SetPurchasedApiGroupStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPurchasedApiGroupStatusRequest) GoString() string {
	return s.String()
}

func (s *SetPurchasedApiGroupStatusRequest) SetAliUid(v int64) *SetPurchasedApiGroupStatusRequest {
	s.AliUid = &v
	return s
}

func (s *SetPurchasedApiGroupStatusRequest) SetBillingStatus(v string) *SetPurchasedApiGroupStatusRequest {
	s.BillingStatus = &v
	return s
}

func (s *SetPurchasedApiGroupStatusRequest) SetCloseOrder(v bool) *SetPurchasedApiGroupStatusRequest {
	s.CloseOrder = &v
	return s
}

func (s *SetPurchasedApiGroupStatusRequest) SetGroupId(v string) *SetPurchasedApiGroupStatusRequest {
	s.GroupId = &v
	return s
}

type SetPurchasedApiGroupStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPurchasedApiGroupStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPurchasedApiGroupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetPurchasedApiGroupStatusResponseBody) SetRequestId(v string) *SetPurchasedApiGroupStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetPurchasedApiGroupStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPurchasedApiGroupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPurchasedApiGroupStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPurchasedApiGroupStatusResponse) GoString() string {
	return s.String()
}

func (s *SetPurchasedApiGroupStatusResponse) SetHeaders(v map[string]*string) *SetPurchasedApiGroupStatusResponse {
	s.Headers = v
	return s
}

func (s *SetPurchasedApiGroupStatusResponse) SetStatusCode(v int32) *SetPurchasedApiGroupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPurchasedApiGroupStatusResponse) SetBody(v *SetPurchasedApiGroupStatusResponseBody) *SetPurchasedApiGroupStatusResponse {
	s.Body = v
	return s
}

type SwitchApiRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SwitchApiRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchApiRequest) GoString() string {
	return s.String()
}

func (s *SwitchApiRequest) SetApiId(v string) *SwitchApiRequest {
	s.ApiId = &v
	return s
}

func (s *SwitchApiRequest) SetDescription(v string) *SwitchApiRequest {
	s.Description = &v
	return s
}

func (s *SwitchApiRequest) SetGroupId(v string) *SwitchApiRequest {
	s.GroupId = &v
	return s
}

func (s *SwitchApiRequest) SetHistoryVersion(v string) *SwitchApiRequest {
	s.HistoryVersion = &v
	return s
}

func (s *SwitchApiRequest) SetSecurityToken(v string) *SwitchApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchApiRequest) SetStageName(v string) *SwitchApiRequest {
	s.StageName = &v
	return s
}

type SwitchApiResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchApiResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchApiResponseBody) SetRequestId(v string) *SwitchApiResponseBody {
	s.RequestId = &v
	return s
}

type SwitchApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchApiResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchApiResponse) GoString() string {
	return s.String()
}

func (s *SwitchApiResponse) SetHeaders(v map[string]*string) *SwitchApiResponse {
	s.Headers = v
	return s
}

func (s *SwitchApiResponse) SetStatusCode(v int32) *SwitchApiResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchApiResponse) SetBody(v *SwitchApiResponseBody) *SwitchApiResponse {
	s.Body = v
	return s
}

type SwitchApiForInnerRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SwitchApiForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchApiForInnerRequest) GoString() string {
	return s.String()
}

func (s *SwitchApiForInnerRequest) SetApiId(v string) *SwitchApiForInnerRequest {
	s.ApiId = &v
	return s
}

func (s *SwitchApiForInnerRequest) SetDescription(v string) *SwitchApiForInnerRequest {
	s.Description = &v
	return s
}

func (s *SwitchApiForInnerRequest) SetGroupId(v string) *SwitchApiForInnerRequest {
	s.GroupId = &v
	return s
}

func (s *SwitchApiForInnerRequest) SetHistoryVersion(v string) *SwitchApiForInnerRequest {
	s.HistoryVersion = &v
	return s
}

func (s *SwitchApiForInnerRequest) SetSecurityToken(v string) *SwitchApiForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchApiForInnerRequest) SetStageName(v string) *SwitchApiForInnerRequest {
	s.StageName = &v
	return s
}

type SwitchApiForInnerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchApiForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchApiForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchApiForInnerResponseBody) SetRequestId(v string) *SwitchApiForInnerResponseBody {
	s.RequestId = &v
	return s
}

type SwitchApiForInnerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchApiForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchApiForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchApiForInnerResponse) GoString() string {
	return s.String()
}

func (s *SwitchApiForInnerResponse) SetHeaders(v map[string]*string) *SwitchApiForInnerResponse {
	s.Headers = v
	return s
}

func (s *SwitchApiForInnerResponse) SetStatusCode(v int32) *SwitchApiForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchApiForInnerResponse) SetBody(v *SwitchApiForInnerResponseBody) *SwitchApiForInnerResponse {
	s.Body = v
	return s
}

type SynCreateAppForInnerRequest struct {
	AliUid        *int64  `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AppKey        *int64  `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppSecret     *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Source        *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s SynCreateAppForInnerRequest) String() string {
	return tea.Prettify(s)
}

func (s SynCreateAppForInnerRequest) GoString() string {
	return s.String()
}

func (s *SynCreateAppForInnerRequest) SetAliUid(v int64) *SynCreateAppForInnerRequest {
	s.AliUid = &v
	return s
}

func (s *SynCreateAppForInnerRequest) SetAppKey(v int64) *SynCreateAppForInnerRequest {
	s.AppKey = &v
	return s
}

func (s *SynCreateAppForInnerRequest) SetAppName(v string) *SynCreateAppForInnerRequest {
	s.AppName = &v
	return s
}

func (s *SynCreateAppForInnerRequest) SetAppSecret(v string) *SynCreateAppForInnerRequest {
	s.AppSecret = &v
	return s
}

func (s *SynCreateAppForInnerRequest) SetDescription(v string) *SynCreateAppForInnerRequest {
	s.Description = &v
	return s
}

func (s *SynCreateAppForInnerRequest) SetSecurityToken(v string) *SynCreateAppForInnerRequest {
	s.SecurityToken = &v
	return s
}

func (s *SynCreateAppForInnerRequest) SetSource(v string) *SynCreateAppForInnerRequest {
	s.Source = &v
	return s
}

type SynCreateAppForInnerResponseBody struct {
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SynCreateAppForInnerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SynCreateAppForInnerResponseBody) GoString() string {
	return s.String()
}

func (s *SynCreateAppForInnerResponseBody) SetAppId(v int64) *SynCreateAppForInnerResponseBody {
	s.AppId = &v
	return s
}

func (s *SynCreateAppForInnerResponseBody) SetRequestId(v string) *SynCreateAppForInnerResponseBody {
	s.RequestId = &v
	return s
}

type SynCreateAppForInnerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SynCreateAppForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SynCreateAppForInnerResponse) String() string {
	return tea.Prettify(s)
}

func (s SynCreateAppForInnerResponse) GoString() string {
	return s.String()
}

func (s *SynCreateAppForInnerResponse) SetHeaders(v map[string]*string) *SynCreateAppForInnerResponse {
	s.Headers = v
	return s
}

func (s *SynCreateAppForInnerResponse) SetStatusCode(v int32) *SynCreateAppForInnerResponse {
	s.StatusCode = &v
	return s
}

func (s *SynCreateAppForInnerResponse) SetBody(v *SynCreateAppForInnerResponseBody) *SynCreateAppForInnerResponse {
	s.Body = v
	return s
}

type TagResourcesSystemTagsRequest struct {
	ResourceId    []*string                           `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType  *string                             `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Scope         *string                             `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SecurityToken *string                             `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*TagResourcesSystemTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TagOwnerBid   *string                             `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	TagOwnerUid   *int64                              `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s TagResourcesSystemTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsRequest) SetResourceId(v []*string) *TagResourcesSystemTagsRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetResourceType(v string) *TagResourcesSystemTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetScope(v string) *TagResourcesSystemTagsRequest {
	s.Scope = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetSecurityToken(v string) *TagResourcesSystemTagsRequest {
	s.SecurityToken = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTag(v []*TagResourcesSystemTagsRequestTag) *TagResourcesSystemTagsRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTagOwnerBid(v string) *TagResourcesSystemTagsRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *TagResourcesSystemTagsRequest) SetTagOwnerUid(v int64) *TagResourcesSystemTagsRequest {
	s.TagOwnerUid = &v
	return s
}

type TagResourcesSystemTagsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesSystemTagsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsRequestTag) SetKey(v string) *TagResourcesSystemTagsRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesSystemTagsRequestTag) SetValue(v string) *TagResourcesSystemTagsRequestTag {
	s.Value = &v
	return s
}

type TagResourcesSystemTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesSystemTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsResponseBody) SetRequestId(v string) *TagResourcesSystemTagsResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesSystemTagsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesSystemTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesSystemTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesSystemTagsResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesSystemTagsResponse) SetHeaders(v map[string]*string) *TagResourcesSystemTagsResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesSystemTagsResponse) SetStatusCode(v int32) *TagResourcesSystemTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesSystemTagsResponse) SetBody(v *TagResourcesSystemTagsResponseBody) *TagResourcesSystemTagsResponse {
	s.Body = v
	return s
}

type UntagResourcesSystemTagsRequest struct {
	All           *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId    []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TagKey        []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	TagOwnerBid   *string   `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	TagOwnerUid   *int64    `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s UntagResourcesSystemTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesSystemTagsRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsRequest) SetAll(v bool) *UntagResourcesSystemTagsRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetResourceId(v []*string) *UntagResourcesSystemTagsRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetResourceType(v string) *UntagResourcesSystemTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetSecurityToken(v string) *UntagResourcesSystemTagsRequest {
	s.SecurityToken = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagKey(v []*string) *UntagResourcesSystemTagsRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagOwnerBid(v string) *UntagResourcesSystemTagsRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *UntagResourcesSystemTagsRequest) SetTagOwnerUid(v int64) *UntagResourcesSystemTagsRequest {
	s.TagOwnerUid = &v
	return s
}

type UntagResourcesSystemTagsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesSystemTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesSystemTagsResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsResponseBody) SetRequestId(v string) *UntagResourcesSystemTagsResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesSystemTagsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesSystemTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesSystemTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesSystemTagsResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesSystemTagsResponse) SetHeaders(v map[string]*string) *UntagResourcesSystemTagsResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesSystemTagsResponse) SetStatusCode(v int32) *UntagResourcesSystemTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesSystemTagsResponse) SetBody(v *UntagResourcesSystemTagsResponseBody) *UntagResourcesSystemTagsResponse {
	s.Body = v
	return s
}

type VipMigrationRequest struct {
	NewVip        *string `json:"NewVip,omitempty" xml:"NewVip,omitempty"`
	OriginalVip   *string `json:"OriginalVip,omitempty" xml:"OriginalVip,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s VipMigrationRequest) String() string {
	return tea.Prettify(s)
}

func (s VipMigrationRequest) GoString() string {
	return s.String()
}

func (s *VipMigrationRequest) SetNewVip(v string) *VipMigrationRequest {
	s.NewVip = &v
	return s
}

func (s *VipMigrationRequest) SetOriginalVip(v string) *VipMigrationRequest {
	s.OriginalVip = &v
	return s
}

func (s *VipMigrationRequest) SetSecurityToken(v string) *VipMigrationRequest {
	s.SecurityToken = &v
	return s
}

type VipMigrationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s VipMigrationResponse) String() string {
	return tea.Prettify(s)
}

func (s VipMigrationResponse) GoString() string {
	return s.String()
}

func (s *VipMigrationResponse) SetHeaders(v map[string]*string) *VipMigrationResponse {
	s.Headers = v
	return s
}

func (s *VipMigrationResponse) SetStatusCode(v int32) *VipMigrationResponse {
	s.StatusCode = &v
	return s
}

type VipMigrationByUidRequest struct {
	NewVip        *string `json:"NewVip,omitempty" xml:"NewVip,omitempty"`
	OriginalVip   *string `json:"OriginalVip,omitempty" xml:"OriginalVip,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s VipMigrationByUidRequest) String() string {
	return tea.Prettify(s)
}

func (s VipMigrationByUidRequest) GoString() string {
	return s.String()
}

func (s *VipMigrationByUidRequest) SetNewVip(v string) *VipMigrationByUidRequest {
	s.NewVip = &v
	return s
}

func (s *VipMigrationByUidRequest) SetOriginalVip(v string) *VipMigrationByUidRequest {
	s.OriginalVip = &v
	return s
}

func (s *VipMigrationByUidRequest) SetSecurityToken(v string) *VipMigrationByUidRequest {
	s.SecurityToken = &v
	return s
}

type VipMigrationByUidResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VipMigrationByUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VipMigrationByUidResponseBody) GoString() string {
	return s.String()
}

func (s *VipMigrationByUidResponseBody) SetRequestId(v string) *VipMigrationByUidResponseBody {
	s.RequestId = &v
	return s
}

type VipMigrationByUidResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VipMigrationByUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VipMigrationByUidResponse) String() string {
	return tea.Prettify(s)
}

func (s VipMigrationByUidResponse) GoString() string {
	return s.String()
}

func (s *VipMigrationByUidResponse) SetHeaders(v map[string]*string) *VipMigrationByUidResponse {
	s.Headers = v
	return s
}

func (s *VipMigrationByUidResponse) SetStatusCode(v int32) *VipMigrationByUidResponse {
	s.StatusCode = &v
	return s
}

func (s *VipMigrationByUidResponse) SetBody(v *VipMigrationByUidResponseBody) *VipMigrationByUidResponse {
	s.Body = v
	return s
}

type VpcAddAppServerRequest struct {
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s VpcAddAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcAddAppServerRequest) GoString() string {
	return s.String()
}

func (s *VpcAddAppServerRequest) SetAddressPoolId(v string) *VpcAddAppServerRequest {
	s.AddressPoolId = &v
	return s
}

func (s *VpcAddAppServerRequest) SetAppId(v string) *VpcAddAppServerRequest {
	s.AppId = &v
	return s
}

func (s *VpcAddAppServerRequest) SetSecurityToken(v string) *VpcAddAppServerRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcAddAppServerRequest) SetServerIp(v string) *VpcAddAppServerRequest {
	s.ServerIp = &v
	return s
}

func (s *VpcAddAppServerRequest) SetToken(v string) *VpcAddAppServerRequest {
	s.Token = &v
	return s
}

type VpcAddAppServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcAddAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcAddAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *VpcAddAppServerResponseBody) SetRequestId(v string) *VpcAddAppServerResponseBody {
	s.RequestId = &v
	return s
}

type VpcAddAppServerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VpcAddAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VpcAddAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcAddAppServerResponse) GoString() string {
	return s.String()
}

func (s *VpcAddAppServerResponse) SetHeaders(v map[string]*string) *VpcAddAppServerResponse {
	s.Headers = v
	return s
}

func (s *VpcAddAppServerResponse) SetStatusCode(v int32) *VpcAddAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcAddAppServerResponse) SetBody(v *VpcAddAppServerResponseBody) *VpcAddAppServerResponse {
	s.Body = v
	return s
}

type VpcCreateAddressPoolRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s VpcCreateAddressPoolRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcCreateAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *VpcCreateAddressPoolRequest) SetAppId(v string) *VpcCreateAddressPoolRequest {
	s.AppId = &v
	return s
}

func (s *VpcCreateAddressPoolRequest) SetDescription(v string) *VpcCreateAddressPoolRequest {
	s.Description = &v
	return s
}

func (s *VpcCreateAddressPoolRequest) SetName(v string) *VpcCreateAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *VpcCreateAddressPoolRequest) SetSecurityToken(v string) *VpcCreateAddressPoolRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcCreateAddressPoolRequest) SetToken(v string) *VpcCreateAddressPoolRequest {
	s.Token = &v
	return s
}

type VpcCreateAddressPoolResponseBody struct {
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcCreateAddressPoolResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcCreateAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *VpcCreateAddressPoolResponseBody) SetAddressPoolId(v string) *VpcCreateAddressPoolResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *VpcCreateAddressPoolResponseBody) SetRequestId(v string) *VpcCreateAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

type VpcCreateAddressPoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VpcCreateAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VpcCreateAddressPoolResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcCreateAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *VpcCreateAddressPoolResponse) SetHeaders(v map[string]*string) *VpcCreateAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *VpcCreateAddressPoolResponse) SetStatusCode(v int32) *VpcCreateAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcCreateAddressPoolResponse) SetBody(v *VpcCreateAddressPoolResponseBody) *VpcCreateAddressPoolResponse {
	s.Body = v
	return s
}

type VpcQueryAppServersRequest struct {
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
}

func (s VpcQueryAppServersRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcQueryAppServersRequest) GoString() string {
	return s.String()
}

func (s *VpcQueryAppServersRequest) SetAddressPoolId(v string) *VpcQueryAppServersRequest {
	s.AddressPoolId = &v
	return s
}

func (s *VpcQueryAppServersRequest) SetAppId(v string) *VpcQueryAppServersRequest {
	s.AppId = &v
	return s
}

func (s *VpcQueryAppServersRequest) SetServerIp(v string) *VpcQueryAppServersRequest {
	s.ServerIp = &v
	return s
}

type VpcQueryAppServersResponseBody struct {
	AppServerInfos *VpcQueryAppServersResponseBodyAppServerInfos `json:"AppServerInfos,omitempty" xml:"AppServerInfos,omitempty" type:"Struct"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcQueryAppServersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcQueryAppServersResponseBody) GoString() string {
	return s.String()
}

func (s *VpcQueryAppServersResponseBody) SetAppServerInfos(v *VpcQueryAppServersResponseBodyAppServerInfos) *VpcQueryAppServersResponseBody {
	s.AppServerInfos = v
	return s
}

func (s *VpcQueryAppServersResponseBody) SetRequestId(v string) *VpcQueryAppServersResponseBody {
	s.RequestId = &v
	return s
}

type VpcQueryAppServersResponseBodyAppServerInfos struct {
	AppServerInfo []*VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo `json:"AppServerInfo,omitempty" xml:"AppServerInfo,omitempty" type:"Repeated"`
}

func (s VpcQueryAppServersResponseBodyAppServerInfos) String() string {
	return tea.Prettify(s)
}

func (s VpcQueryAppServersResponseBodyAppServerInfos) GoString() string {
	return s.String()
}

func (s *VpcQueryAppServersResponseBodyAppServerInfos) SetAppServerInfo(v []*VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) *VpcQueryAppServersResponseBodyAppServerInfos {
	s.AppServerInfo = v
	return s
}

type VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo struct {
	AddressPoolId   *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServerIp        *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	ServerMappingIp *string `json:"ServerMappingIp,omitempty" xml:"ServerMappingIp,omitempty"`
	ServerType      *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId           *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) String() string {
	return tea.Prettify(s)
}

func (s VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) GoString() string {
	return s.String()
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetAddressPoolId(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.AddressPoolId = &v
	return s
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetAppId(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.AppId = &v
	return s
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetServerIp(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.ServerIp = &v
	return s
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetServerMappingIp(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.ServerMappingIp = &v
	return s
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetServerType(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.ServerType = &v
	return s
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetStatus(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.Status = &v
	return s
}

func (s *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo) SetVpcId(v string) *VpcQueryAppServersResponseBodyAppServerInfosAppServerInfo {
	s.VpcId = &v
	return s
}

type VpcQueryAppServersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VpcQueryAppServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VpcQueryAppServersResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcQueryAppServersResponse) GoString() string {
	return s.String()
}

func (s *VpcQueryAppServersResponse) SetHeaders(v map[string]*string) *VpcQueryAppServersResponse {
	s.Headers = v
	return s
}

func (s *VpcQueryAppServersResponse) SetStatusCode(v int32) *VpcQueryAppServersResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcQueryAppServersResponse) SetBody(v *VpcQueryAppServersResponseBody) *VpcQueryAppServersResponse {
	s.Body = v
	return s
}

type VpcRegisterAppRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s VpcRegisterAppRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcRegisterAppRequest) GoString() string {
	return s.String()
}

func (s *VpcRegisterAppRequest) SetAppId(v string) *VpcRegisterAppRequest {
	s.AppId = &v
	return s
}

func (s *VpcRegisterAppRequest) SetDescription(v string) *VpcRegisterAppRequest {
	s.Description = &v
	return s
}

func (s *VpcRegisterAppRequest) SetName(v string) *VpcRegisterAppRequest {
	s.Name = &v
	return s
}

func (s *VpcRegisterAppRequest) SetSecurityToken(v string) *VpcRegisterAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcRegisterAppRequest) SetToken(v string) *VpcRegisterAppRequest {
	s.Token = &v
	return s
}

type VpcRegisterAppResponseBody struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcRegisterAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcRegisterAppResponseBody) GoString() string {
	return s.String()
}

func (s *VpcRegisterAppResponseBody) SetAppId(v string) *VpcRegisterAppResponseBody {
	s.AppId = &v
	return s
}

func (s *VpcRegisterAppResponseBody) SetRequestId(v string) *VpcRegisterAppResponseBody {
	s.RequestId = &v
	return s
}

type VpcRegisterAppResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VpcRegisterAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VpcRegisterAppResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcRegisterAppResponse) GoString() string {
	return s.String()
}

func (s *VpcRegisterAppResponse) SetHeaders(v map[string]*string) *VpcRegisterAppResponse {
	s.Headers = v
	return s
}

func (s *VpcRegisterAppResponse) SetStatusCode(v int32) *VpcRegisterAppResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcRegisterAppResponse) SetBody(v *VpcRegisterAppResponseBody) *VpcRegisterAppResponse {
	s.Body = v
	return s
}

type VpcRemoveAppServerRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServerIp      *string `json:"ServerIp,omitempty" xml:"ServerIp,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s VpcRemoveAppServerRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcRemoveAppServerRequest) GoString() string {
	return s.String()
}

func (s *VpcRemoveAppServerRequest) SetAppId(v string) *VpcRemoveAppServerRequest {
	s.AppId = &v
	return s
}

func (s *VpcRemoveAppServerRequest) SetSecurityToken(v string) *VpcRemoveAppServerRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcRemoveAppServerRequest) SetServerIp(v string) *VpcRemoveAppServerRequest {
	s.ServerIp = &v
	return s
}

func (s *VpcRemoveAppServerRequest) SetToken(v string) *VpcRemoveAppServerRequest {
	s.Token = &v
	return s
}

type VpcRemoveAppServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcRemoveAppServerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcRemoveAppServerResponseBody) GoString() string {
	return s.String()
}

func (s *VpcRemoveAppServerResponseBody) SetRequestId(v string) *VpcRemoveAppServerResponseBody {
	s.RequestId = &v
	return s
}

type VpcRemoveAppServerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VpcRemoveAppServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VpcRemoveAppServerResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcRemoveAppServerResponse) GoString() string {
	return s.String()
}

func (s *VpcRemoveAppServerResponse) SetHeaders(v map[string]*string) *VpcRemoveAppServerResponse {
	s.Headers = v
	return s
}

func (s *VpcRemoveAppServerResponse) SetStatusCode(v int32) *VpcRemoveAppServerResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcRemoveAppServerResponse) SetBody(v *VpcRemoveAppServerResponseBody) *VpcRemoveAppServerResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-qingdao":                  tea.String("apigateway.cn-qingdao.aliyuncs.com"),
		"cn-beijing":                  tea.String("apigateway.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("apigateway.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":                tea.String("apigateway.cn-huhehaote.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("apigateway.cn-wulanchabu.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("apigateway.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":                 tea.String("apigateway.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("apigateway.cn-shenzhen.aliyuncs.com"),
		"cn-heyuan":                   tea.String("apigateway.cn-heyuan.aliyuncs.com"),
		"cn-guangzhou":                tea.String("apigateway.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":                  tea.String("apigateway.cn-chengdu.aliyuncs.com"),
		"cn-hongkong":                 tea.String("apigateway.cn-hongkong.aliyuncs.com"),
		"ap-northeast-1":              tea.String("apigateway.ap-northeast-1.aliyuncs.com"),
		"ap-southeast-1":              tea.String("apigateway.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":              tea.String("apigateway.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":              tea.String("apigateway.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":              tea.String("apigateway.ap-southeast-5.aliyuncs.com"),
		"ap-southeast-6":              tea.String("apigateway.ap-southeast-6.aliyuncs.com"),
		"ap-southeast-7":              tea.String("apigateway.ap-southeast-7.aliyuncs.com"),
		"us-east-1":                   tea.String("apigateway.us-east-1.aliyuncs.com"),
		"us-west-1":                   tea.String("apigateway.us-west-1.aliyuncs.com"),
		"eu-west-1":                   tea.String("apigateway.eu-west-1.aliyuncs.com"),
		"eu-central-1":                tea.String("apigateway.eu-central-1.aliyuncs.com"),
		"ap-south-1":                  tea.String("apigateway.ap-south-1.aliyuncs.com"),
		"me-east-1":                   tea.String("apigateway.me-east-1.aliyuncs.com"),
		"me-central-1":                tea.String("apigateway.me-central-1.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("apigateway.cn-hangzhou-finance.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("apigateway.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("apigateway.cn-shenzhen-finance-1.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("apigateway.cn-north-2-gov-1.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("apigateway.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("apigateway.cn-beijing-finance-1.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("apigateway.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("apigateway.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("apigateway.aliyuncs.com"),
		"cn-edge-1":                   tea.String("apigateway.aliyuncs.com"),
		"cn-fujian":                   tea.String("apigateway.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("apigateway.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("apigateway.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("apigateway.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("apigateway.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("apigateway.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("apigateway.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("apigateway.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("apigateway.cn-shanghai-inner.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("apigateway.aliyuncs.com"),
		"cn-wuhan":                    tea.String("apigateway.aliyuncs.com"),
		"cn-yushanfang":               tea.String("apigateway.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("apigateway.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("apigateway.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("apigateway.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("apigateway.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("apigateway.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("apigateway.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbolishApiWithOptions(request *AbolishApiRequest, runtime *util.RuntimeOptions) (_result *AbolishApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbolishApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbolishApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbolishApi(request *AbolishApiRequest) (_result *AbolishApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbolishApiResponse{}
	_body, _err := client.AbolishApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbolishApiForInnerWithOptions(request *AbolishApiForInnerRequest, runtime *util.RuntimeOptions) (_result *AbolishApiForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbolishApiForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AbolishApiForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbolishApiForInner(request *AbolishApiForInnerRequest) (_result *AbolishApiForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AbolishApiForInnerResponse{}
	_body, _err := client.AbolishApiForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddBlackListWithOptions(request *AddBlackListRequest, runtime *util.RuntimeOptions) (_result *AddBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackContent)) {
		query["BlackContent"] = request.BlackContent
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBlackList"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddBlackList(request *AddBlackListRequest) (_result *AddBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBlackListResponse{}
	_body, _err := client.AddBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddTrafficSpecialControlWithOptions(request *AddTrafficSpecialControlRequest, runtime *util.RuntimeOptions) (_result *AddTrafficSpecialControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialKey)) {
		query["SpecialKey"] = request.SpecialKey
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialType)) {
		query["SpecialType"] = request.SpecialType
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlId)) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficValue)) {
		query["TrafficValue"] = request.TrafficValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTrafficSpecialControl"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddTrafficSpecialControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddTrafficSpecialControl(request *AddTrafficSpecialControlRequest) (_result *AddTrafficSpecialControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddTrafficSpecialControlResponse{}
	_body, _err := client.AddTrafficSpecialControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAccountForInnerWithOptions(request *CheckAccountForInnerRequest, runtime *util.RuntimeOptions) (_result *CheckAccountForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAccountForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAccountForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAccountForInner(request *CheckAccountForInnerRequest) (_result *CheckAccountForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAccountForInnerResponse{}
	_body, _err := client.CheckAccountForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckAoneAppAuditWithOptions(request *CheckAoneAppAuditRequest, runtime *util.RuntimeOptions) (_result *CheckAoneAppAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AoneAppName)) {
		query["AoneAppName"] = request.AoneAppName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckAoneAppAudit"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckAoneAppAuditResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckAoneAppAudit(request *CheckAoneAppAuditRequest) (_result *CheckAoneAppAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckAoneAppAuditResponse{}
	_body, _err := client.CheckAoneAppAuditWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyConsumerOpenForInnerWithOptions(request *CopyConsumerOpenForInnerRequest, runtime *util.RuntimeOptions) (_result *CopyConsumerOpenForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CopyList)) {
		query["CopyList"] = request.CopyList
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyConsumerOpenForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyConsumerOpenForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyConsumerOpenForInner(request *CopyConsumerOpenForInnerRequest) (_result *CopyConsumerOpenForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyConsumerOpenForInnerResponse{}
	_body, _err := client.CopyConsumerOpenForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApiWithOptions(request *CreateApiRequest, runtime *util.RuntimeOptions) (_result *CreateApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.BodyFormat)) {
		query["BodyFormat"] = request.BodyFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ConstantParameters)) {
		query["ConstantParameters"] = request.ConstantParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HttpMethod)) {
		query["HttpMethod"] = request.HttpMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HttpProtocol)) {
		query["HttpProtocol"] = request.HttpProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.PathParameters)) {
		query["PathParameters"] = request.PathParameters
	}

	if !tea.BoolValue(util.IsUnset(request.PostBodyDescription)) {
		query["PostBodyDescription"] = request.PostBodyDescription
	}

	if !tea.BoolValue(util.IsUnset(request.PostBodyType)) {
		query["PostBodyType"] = request.PostBodyType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestBody)) {
		query["RequestBody"] = request.RequestBody
	}

	if !tea.BoolValue(util.IsUnset(request.RequestHeaders)) {
		query["RequestHeaders"] = request.RequestHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.RequestQueries)) {
		query["RequestQueries"] = request.RequestQueries
	}

	if !tea.BoolValue(util.IsUnset(request.ResultSample)) {
		query["ResultSample"] = request.ResultSample
	}

	if !tea.BoolValue(util.IsUnset(request.ResultType)) {
		query["ResultType"] = request.ResultType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceAddress)) {
		query["ServiceAddress"] = request.ServiceAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceProtocol)) {
		query["ServiceProtocol"] = request.ServiceProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceTimeout)) {
		query["ServiceTimeout"] = request.ServiceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.SystemParameters)) {
		query["SystemParameters"] = request.SystemParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApi(request *CreateApiRequest) (_result *CreateApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApiResponse{}
	_body, _err := client.CreateApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApiForInnerWithOptions(request *CreateApiForInnerRequest, runtime *util.RuntimeOptions) (_result *CreateApiForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestConfig)) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestParamters)) {
		query["RequestParamters"] = request.RequestParamters
	}

	if !tea.BoolValue(util.IsUnset(request.ResultSample)) {
		query["ResultSample"] = request.ResultSample
	}

	if !tea.BoolValue(util.IsUnset(request.ResultType)) {
		query["ResultType"] = request.ResultType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceConfig)) {
		query["ServiceConfig"] = request.ServiceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParametersMap)) {
		query["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApiForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApiForInner(request *CreateApiForInnerRequest) (_result *CreateApiForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApiForInnerResponse{}
	_body, _err := client.CreateApiForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApiGroupWithOptions(request *CreateApiGroupRequest, runtime *util.RuntimeOptions) (_result *CreateApiGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiGroup"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApiGroup(request *CreateApiGroupRequest) (_result *CreateApiGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApiGroupResponse{}
	_body, _err := client.CreateApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApiGroupForInnerWithOptions(request *CreateApiGroupForInnerRequest, runtime *util.RuntimeOptions) (_result *CreateApiGroupForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiGroupForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApiGroupForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApiGroupForInner(request *CreateApiGroupForInnerRequest) (_result *CreateApiGroupForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApiGroupForInnerResponse{}
	_body, _err := client.CreateApiGroupForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(request *CreateAppRequest, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp(request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppForBackendWithOptions(request *CreateAppForBackendRequest, runtime *util.RuntimeOptions) (_result *CreateAppForBackendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppForBackend"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppForBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppForBackend(request *CreateAppForBackendRequest) (_result *CreateAppForBackendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppForBackendResponse{}
	_body, _err := client.CreateAppForBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppForInnerWithOptions(request *CreateAppForInnerRequest, runtime *util.RuntimeOptions) (_result *CreateAppForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["AppSecret"] = request.AppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		query["Extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppForInner(request *CreateAppForInnerRequest) (_result *CreateAppForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppForInnerResponse{}
	_body, _err := client.CreateAppForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountQuantity)) {
		query["AccountQuantity"] = request.AccountQuantity
	}

	if !tea.BoolValue(util.IsUnset(request.AlarmQuota)) {
		query["AlarmQuota"] = request.AlarmQuota
	}

	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BillingType)) {
		query["BillingType"] = request.BillingType
	}

	if !tea.BoolValue(util.IsUnset(request.CloudMarketInstanceId)) {
		query["CloudMarketInstanceId"] = request.CloudMarketInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpiredOn)) {
		query["ExpiredOn"] = request.ExpiredOn
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAttributes)) {
		query["InstanceAttributes"] = request.InstanceAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.SkuId)) {
		query["SkuId"] = request.SkuId
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRaceWorkForInnerWithOptions(request *CreateRaceWorkForInnerRequest, runtime *util.RuntimeOptions) (_result *CreateRaceWorkForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		query["CommodityCode"] = request.CommodityCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		query["Keywords"] = request.Keywords
	}

	if !tea.BoolValue(util.IsUnset(request.LogoUrl)) {
		query["LogoUrl"] = request.LogoUrl
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ShortDescription)) {
		query["ShortDescription"] = request.ShortDescription
	}

	if !tea.BoolValue(util.IsUnset(request.WorkName)) {
		query["WorkName"] = request.WorkName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRaceWorkForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRaceWorkForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRaceWorkForInner(request *CreateRaceWorkForInnerRequest) (_result *CreateRaceWorkForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRaceWorkForInnerResponse{}
	_body, _err := client.CreateRaceWorkForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecretKeyWithOptions(request *CreateSecretKeyRequest, runtime *util.RuntimeOptions) (_result *CreateSecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		query["SecretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKeyName)) {
		query["SecretKeyName"] = request.SecretKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.SecretValue)) {
		query["SecretValue"] = request.SecretValue
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecretKey"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecretKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecretKey(request *CreateSecretKeyRequest) (_result *CreateSecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecretKeyResponse{}
	_body, _err := client.CreateSecretKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTrafficControlWithOptions(request *CreateTrafficControlRequest, runtime *util.RuntimeOptions) (_result *CreateTrafficControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDefault)) {
		query["ApiDefault"] = request.ApiDefault
	}

	if !tea.BoolValue(util.IsUnset(request.AppDefault)) {
		query["AppDefault"] = request.AppDefault
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlName)) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlUnit)) {
		query["TrafficControlUnit"] = request.TrafficControlUnit
	}

	if !tea.BoolValue(util.IsUnset(request.UserDefault)) {
		query["UserDefault"] = request.UserDefault
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrafficControl"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTrafficControl(request *CreateTrafficControlRequest) (_result *CreateTrafficControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTrafficControlResponse{}
	_body, _err := client.CreateTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWhiteListWithOptions(request *CreateUserWhiteListRequest, runtime *util.RuntimeOptions) (_result *CreateUserWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AoneId)) {
		query["AoneId"] = request.AoneId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.LimitCount)) {
		query["LimitCount"] = request.LimitCount
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserWhiteList"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserWhiteListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserWhiteList(request *CreateUserWhiteListRequest) (_result *CreateUserWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserWhiteListResponse{}
	_body, _err := client.CreateUserWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAllTrafficSpecialControlWithOptions(request *DeleteAllTrafficSpecialControlRequest, runtime *util.RuntimeOptions) (_result *DeleteAllTrafficSpecialControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlId)) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAllTrafficSpecialControl"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAllTrafficSpecialControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAllTrafficSpecialControl(request *DeleteAllTrafficSpecialControlRequest) (_result *DeleteAllTrafficSpecialControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAllTrafficSpecialControlResponse{}
	_body, _err := client.DeleteAllTrafficSpecialControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApiWithOptions(request *DeleteApiRequest, runtime *util.RuntimeOptions) (_result *DeleteApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApi(request *DeleteApiRequest) (_result *DeleteApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApiResponse{}
	_body, _err := client.DeleteApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApiForInnerWithOptions(request *DeleteApiForInnerRequest, runtime *util.RuntimeOptions) (_result *DeleteApiForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApiForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApiForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApiForInner(request *DeleteApiForInnerRequest) (_result *DeleteApiForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApiForInnerResponse{}
	_body, _err := client.DeleteApiForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteApiGroupWithOptions(request *DeleteApiGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteApiGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApiGroup"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (_result *DeleteApiGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApiGroupResponse{}
	_body, _err := client.DeleteApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(request *DeleteAppRequest, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(request *DeleteAppRequest) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppForInnerWithOptions(request *DeleteAppForInnerRequest, runtime *util.RuntimeOptions) (_result *DeleteAppForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppForInner(request *DeleteAppForInnerRequest) (_result *DeleteAppForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppForInnerResponse{}
	_body, _err := client.DeleteAppForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDomainCertificateWithOptions(request *DeleteDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *DeleteDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateId)) {
		query["CertificateId"] = request.CertificateId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainCertificate"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDomainCertificate(request *DeleteDomainCertificateRequest) (_result *DeleteDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDomainCertificateResponse{}
	_body, _err := client.DeleteDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecretKeyWithOptions(request *DeleteSecretKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretKeyId)) {
		query["SecretKeyId"] = request.SecretKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecretKey"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecretKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecretKey(request *DeleteSecretKeyRequest) (_result *DeleteSecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretKeyResponse{}
	_body, _err := client.DeleteSecretKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrafficControlWithOptions(request *DeleteTrafficControlRequest, runtime *util.RuntimeOptions) (_result *DeleteTrafficControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlId)) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficControl"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrafficControl(request *DeleteTrafficControlRequest) (_result *DeleteTrafficControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrafficControlResponse{}
	_body, _err := client.DeleteTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTrafficSpecialControlWithOptions(request *DeleteTrafficSpecialControlRequest, runtime *util.RuntimeOptions) (_result *DeleteTrafficSpecialControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialKey)) {
		query["SpecialKey"] = request.SpecialKey
	}

	if !tea.BoolValue(util.IsUnset(request.SpecialType)) {
		query["SpecialType"] = request.SpecialType
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlId)) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficSpecialControl"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTrafficSpecialControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTrafficSpecialControl(request *DeleteTrafficSpecialControlRequest) (_result *DeleteTrafficSpecialControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTrafficSpecialControlResponse{}
	_body, _err := client.DeleteTrafficSpecialControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWhiteListByTypeWithOptions(request *DeleteUserWhiteListByTypeRequest, runtime *util.RuntimeOptions) (_result *DeleteUserWhiteListByTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityId)) {
		query["EntityId"] = request.EntityId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserWhiteListByType"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserWhiteListByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUserWhiteListByType(request *DeleteUserWhiteListByTypeRequest) (_result *DeleteUserWhiteListByTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserWhiteListByTypeResponse{}
	_body, _err := client.DeleteUserWhiteListByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployApiWithOptions(request *DeployApiRequest, runtime *util.RuntimeOptions) (_result *DeployApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployApi(request *DeployApiRequest) (_result *DeployApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployApiResponse{}
	_body, _err := client.DeployApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployApiForInnerWithOptions(request *DeployApiForInnerRequest, runtime *util.RuntimeOptions) (_result *DeployApiForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApiForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployApiForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployApiForInner(request *DeployApiForInnerRequest) (_result *DeployApiForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployApiForInnerResponse{}
	_body, _err := client.DeployApiForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiWithOptions(request *DescribeApiRequest, runtime *util.RuntimeOptions) (_result *DescribeApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApi(request *DescribeApiRequest) (_result *DescribeApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiResponse{}
	_body, _err := client.DescribeApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiDocWithOptions(request *DescribeApiDocRequest, runtime *util.RuntimeOptions) (_result *DescribeApiDocResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiDoc"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiDocResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiDoc(request *DescribeApiDocRequest) (_result *DescribeApiDocResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiDocResponse{}
	_body, _err := client.DescribeApiDocWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiDocsWithOptions(request *DescribeApiDocsRequest, runtime *util.RuntimeOptions) (_result *DescribeApiDocsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiDocs"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiDocsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiDocs(request *DescribeApiDocsRequest) (_result *DescribeApiDocsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiDocsResponse{}
	_body, _err := client.DescribeApiDocsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiDocsForBackendWithOptions(request *DescribeApiDocsForBackendRequest, runtime *util.RuntimeOptions) (_result *DescribeApiDocsForBackendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiDocsForBackend"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiDocsForBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiDocsForBackend(request *DescribeApiDocsForBackendRequest) (_result *DescribeApiDocsForBackendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiDocsForBackendResponse{}
	_body, _err := client.DescribeApiDocsForBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiErrorWithOptions(request *DescribeApiErrorRequest, runtime *util.RuntimeOptions) (_result *DescribeApiErrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiError"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiErrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiError(request *DescribeApiErrorRequest) (_result *DescribeApiErrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiErrorResponse{}
	_body, _err := client.DescribeApiErrorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiGroupDetailWithOptions(request *DescribeApiGroupDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeApiGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiGroupDetail"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiGroupDetail(request *DescribeApiGroupDetailRequest) (_result *DescribeApiGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiGroupDetailResponse{}
	_body, _err := client.DescribeApiGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiGroupsWithOptions(request *DescribeApiGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeApiGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiGroups"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (_result *DescribeApiGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiGroupsResponse{}
	_body, _err := client.DescribeApiGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiLatencyWithOptions(request *DescribeApiLatencyRequest, runtime *util.RuntimeOptions) (_result *DescribeApiLatencyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiLatency"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiLatencyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiLatency(request *DescribeApiLatencyRequest) (_result *DescribeApiLatencyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiLatencyResponse{}
	_body, _err := client.DescribeApiLatencyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiMarketInstanceWithOptions(request *DescribeApiMarketInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeApiMarketInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiMarketInstance"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiMarketInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiMarketInstance(request *DescribeApiMarketInstanceRequest) (_result *DescribeApiMarketInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiMarketInstanceResponse{}
	_body, _err := client.DescribeApiMarketInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiQpsWithOptions(request *DescribeApiQpsRequest, runtime *util.RuntimeOptions) (_result *DescribeApiQpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiQps"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiQpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiQps(request *DescribeApiQpsRequest) (_result *DescribeApiQpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiQpsResponse{}
	_body, _err := client.DescribeApiQpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiRulesWithOptions(request *DescribeApiRulesRequest, runtime *util.RuntimeOptions) (_result *DescribeApiRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiIds)) {
		query["ApiIds"] = request.ApiIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiRules"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiRules(request *DescribeApiRulesRequest) (_result *DescribeApiRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiRulesResponse{}
	_body, _err := client.DescribeApiRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiTrafficWithOptions(request *DescribeApiTrafficRequest, runtime *util.RuntimeOptions) (_result *DescribeApiTrafficResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiTraffic"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiTrafficResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiTraffic(request *DescribeApiTrafficRequest) (_result *DescribeApiTrafficResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiTrafficResponse{}
	_body, _err := client.DescribeApiTrafficWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApisWithOptions(request *DescribeApisRequest, runtime *util.RuntimeOptions) (_result *DescribeApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApis"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApis(request *DescribeApisRequest) (_result *DescribeApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisResponse{}
	_body, _err := client.DescribeApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApisByAppWithOptions(request *DescribeApisByAppRequest, runtime *util.RuntimeOptions) (_result *DescribeApisByAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisByApp"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApisByApp(request *DescribeApisByAppRequest) (_result *DescribeApisByAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisByAppResponse{}
	_body, _err := client.DescribeApisByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApisByRuleWithOptions(request *DescribeApisByRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeApisByRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisByRule"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisByRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApisByRule(request *DescribeApisByRuleRequest) (_result *DescribeApisByRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisByRuleResponse{}
	_body, _err := client.DescribeApisByRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApisForConsoleWithOptions(request *DescribeApisForConsoleRequest, runtime *util.RuntimeOptions) (_result *DescribeApisForConsoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisForConsole"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisForConsoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApisForConsole(request *DescribeApisForConsoleRequest) (_result *DescribeApisForConsoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisForConsoleResponse{}
	_body, _err := client.DescribeApisForConsoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppWithOptions(request *DescribeAppRequest, runtime *util.RuntimeOptions) (_result *DescribeAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApp"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApp(request *DescribeAppRequest) (_result *DescribeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppResponse{}
	_body, _err := client.DescribeAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppSecuritiesWithOptions(request *DescribeAppSecuritiesRequest, runtime *util.RuntimeOptions) (_result *DescribeAppSecuritiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppSecurities"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppSecuritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppSecurities(request *DescribeAppSecuritiesRequest) (_result *DescribeAppSecuritiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppSecuritiesResponse{}
	_body, _err := client.DescribeAppSecuritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppSecurityWithOptions(request *DescribeAppSecurityRequest, runtime *util.RuntimeOptions) (_result *DescribeAppSecurityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppSecurity"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppSecurityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppSecurity(request *DescribeAppSecurityRequest) (_result *DescribeAppSecurityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppSecurityResponse{}
	_body, _err := client.DescribeAppSecurityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppSecurityForInnerWithOptions(request *DescribeAppSecurityForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeAppSecurityForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppSecurityForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppSecurityForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppSecurityForInner(request *DescribeAppSecurityForInnerRequest) (_result *DescribeAppSecurityForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppSecurityForInnerResponse{}
	_body, _err := client.DescribeAppSecurityForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApps"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppsByApiWithOptions(request *DescribeAppsByApiRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsByApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppsByApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppsByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppsByApi(request *DescribeAppsByApiRequest) (_result *DescribeAppsByApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppsByApiResponse{}
	_body, _err := client.DescribeAppsByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppsForProviderWithOptions(request *DescribeAppsForProviderRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsForProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppOwner)) {
		query["AppOwner"] = request.AppOwner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppsForProvider"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppsForProviderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppsForProvider(request *DescribeAppsForProviderRequest) (_result *DescribeAppsForProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppsForProviderResponse{}
	_body, _err := client.DescribeAppsForProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableVipsWithOptions(request *DescribeAvailableVipsRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableVipsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Vip)) {
		query["Vip"] = request.Vip
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableVips"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableVipsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableVips(request *DescribeAvailableVipsRequest) (_result *DescribeAvailableVipsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableVipsResponse{}
	_body, _err := client.DescribeAvailableVipsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBidByUserIdForInnerWithOptions(request *DescribeBidByUserIdForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeBidByUserIdForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBidByUserIdForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBidByUserIdForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBidByUserIdForInner(request *DescribeBidByUserIdForInnerRequest) (_result *DescribeBidByUserIdForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBidByUserIdForInnerResponse{}
	_body, _err := client.DescribeBidByUserIdForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBlackListsWithOptions(request *DescribeBlackListsRequest, runtime *util.RuntimeOptions) (_result *DescribeBlackListsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBlackLists"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBlackListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBlackLists(request *DescribeBlackListsRequest) (_result *DescribeBlackListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeBlackListsResponse{}
	_body, _err := client.DescribeBlackListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeployedApiWithOptions(request *DescribeDeployedApiRequest, runtime *util.RuntimeOptions) (_result *DescribeDeployedApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeployedApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeployedApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeployedApi(request *DescribeDeployedApiRequest) (_result *DescribeDeployedApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeployedApiResponse{}
	_body, _err := client.DescribeDeployedApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeployedApisWithOptions(request *DescribeDeployedApisRequest, runtime *util.RuntimeOptions) (_result *DescribeDeployedApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeployedApis"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeployedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeployedApis(request *DescribeDeployedApisRequest) (_result *DescribeDeployedApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeployedApisResponse{}
	_body, _err := client.DescribeDeployedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainWithOptions(request *DescribeDomainRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomain"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomain(request *DescribeDomainRequest) (_result *DescribeDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainResponse{}
	_body, _err := client.DescribeDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDomainResolutionWithOptions(request *DescribeDomainResolutionRequest, runtime *util.RuntimeOptions) (_result *DescribeDomainResolutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainNames)) {
		query["DomainNames"] = request.DomainNames
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomainResolution"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDomainResolutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDomainResolution(request *DescribeDomainResolutionRequest) (_result *DescribeDomainResolutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDomainResolutionResponse{}
	_body, _err := client.DescribeDomainResolutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHistoryApiWithOptions(request *DescribeHistoryApiRequest, runtime *util.RuntimeOptions) (_result *DescribeHistoryApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryVersion)) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHistoryApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHistoryApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHistoryApi(request *DescribeHistoryApiRequest) (_result *DescribeHistoryApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHistoryApiResponse{}
	_body, _err := client.DescribeHistoryApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeHistoryApisWithOptions(request *DescribeHistoryApisRequest, runtime *util.RuntimeOptions) (_result *DescribeHistoryApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHistoryApis"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeHistoryApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeHistoryApis(request *DescribeHistoryApisRequest) (_result *DescribeHistoryApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeHistoryApisResponse{}
	_body, _err := client.DescribeHistoryApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModelsForInnerWithOptions(request *DescribeModelsForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeModelsForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelId)) {
		query["ModelId"] = request.ModelId
	}

	if !tea.BoolValue(util.IsUnset(request.ModelName)) {
		query["ModelName"] = request.ModelName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeModelsForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeModelsForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModelsForInner(request *DescribeModelsForInnerRequest) (_result *DescribeModelsForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModelsForInnerResponse{}
	_body, _err := client.DescribeModelsForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedApiWithOptions(request *DescribePurchasedApiRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedApi(request *DescribePurchasedApiRequest) (_result *DescribePurchasedApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedApiResponse{}
	_body, _err := client.DescribePurchasedApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedApiGroupDetailWithOptions(request *DescribePurchasedApiGroupDetailRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedApiGroupDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApiGroupDetail"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedApiGroupDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedApiGroupDetail(request *DescribePurchasedApiGroupDetailRequest) (_result *DescribePurchasedApiGroupDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedApiGroupDetailResponse{}
	_body, _err := client.DescribePurchasedApiGroupDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedApiGroupsWithOptions(request *DescribePurchasedApiGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedApiGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupIds)) {
		query["GroupIds"] = request.GroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApiGroups"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedApiGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedApiGroups(request *DescribePurchasedApiGroupsRequest) (_result *DescribePurchasedApiGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedApiGroupsResponse{}
	_body, _err := client.DescribePurchasedApiGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedApisWithOptions(request *DescribePurchasedApisRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApis"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedApis(request *DescribePurchasedApisRequest) (_result *DescribePurchasedApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedApisResponse{}
	_body, _err := client.DescribePurchasedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRaceWorkForInnerWithOptions(request *DescribeRaceWorkForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeRaceWorkForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRaceWorkForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRaceWorkForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRaceWorkForInner(request *DescribeRaceWorkForInnerRequest) (_result *DescribeRaceWorkForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRaceWorkForInnerResponse{}
	_body, _err := client.DescribeRaceWorkForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRaceWorksForInnerWithOptions(request *DescribeRaceWorksForInnerRequest, runtime *util.RuntimeOptions) (_result *DescribeRaceWorksForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRaceWorksForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRaceWorksForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRaceWorksForInner(request *DescribeRaceWorksForInnerRequest) (_result *DescribeRaceWorksForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRaceWorksForInnerResponse{}
	_body, _err := client.DescribeRaceWorksForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Language)) {
		query["Language"] = request.Language
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRulesByApiWithOptions(request *DescribeRulesByApiRequest, runtime *util.RuntimeOptions) (_result *DescribeRulesByApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRulesByApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRulesByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRulesByApi(request *DescribeRulesByApiRequest) (_result *DescribeRulesByApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRulesByApiResponse{}
	_body, _err := client.DescribeRulesByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecretKeysWithOptions(request *DescribeSecretKeysRequest, runtime *util.RuntimeOptions) (_result *DescribeSecretKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKeyId)) {
		query["SecretKeyId"] = request.SecretKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKeyName)) {
		query["SecretKeyName"] = request.SecretKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecretKeys"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecretKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecretKeys(request *DescribeSecretKeysRequest) (_result *DescribeSecretKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecretKeysResponse{}
	_body, _err := client.DescribeSecretKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSystemParametersWithOptions(request *DescribeSystemParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeSystemParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSystemParameters"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSystemParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSystemParameters(request *DescribeSystemParametersRequest) (_result *DescribeSystemParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSystemParametersResponse{}
	_body, _err := client.DescribeSystemParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSystemParamsWithOptions(request *DescribeSystemParamsRequest, runtime *util.RuntimeOptions) (_result *DescribeSystemParamsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSystemParams"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSystemParamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSystemParams(request *DescribeSystemParamsRequest) (_result *DescribeSystemParamsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSystemParamsResponse{}
	_body, _err := client.DescribeSystemParamsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTrafficControlsWithOptions(request *DescribeTrafficControlsRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficControlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiUid)) {
		query["ApiUid"] = request.ApiUid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlId)) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlName)) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrafficControls"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrafficControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrafficControls(request *DescribeTrafficControlsRequest) (_result *DescribeTrafficControlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficControlsResponse{}
	_body, _err := client.DescribeTrafficControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserWhiteListsWithOptions(request *DescribeUserWhiteListsRequest, runtime *util.RuntimeOptions) (_result *DescribeUserWhiteListsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserWhiteLists"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserWhiteListsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserWhiteLists(request *DescribeUserWhiteListsRequest) (_result *DescribeUserWhiteListsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserWhiteListsResponse{}
	_body, _err := client.DescribeUserWhiteListsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IsIncludedByUserWhitelistWithOptions(request *IsIncludedByUserWhitelistRequest, runtime *util.RuntimeOptions) (_result *IsIncludedByUserWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("IsIncludedByUserWhitelist"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IsIncludedByUserWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IsIncludedByUserWhitelist(request *IsIncludedByUserWhitelistRequest) (_result *IsIncludedByUserWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IsIncludedByUserWhitelistResponse{}
	_body, _err := client.IsIncludedByUserWhitelistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyApiWithOptions(request *ModifyApiRequest, runtime *util.RuntimeOptions) (_result *ModifyApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.BodyFormat)) {
		query["BodyFormat"] = request.BodyFormat
	}

	if !tea.BoolValue(util.IsUnset(request.ConstantParameters)) {
		query["ConstantParameters"] = request.ConstantParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HttpMethod)) {
		query["HttpMethod"] = request.HttpMethod
	}

	if !tea.BoolValue(util.IsUnset(request.HttpProtocol)) {
		query["HttpProtocol"] = request.HttpProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.PathParameters)) {
		query["PathParameters"] = request.PathParameters
	}

	if !tea.BoolValue(util.IsUnset(request.PostBodyDescription)) {
		query["PostBodyDescription"] = request.PostBodyDescription
	}

	if !tea.BoolValue(util.IsUnset(request.PostBodyType)) {
		query["PostBodyType"] = request.PostBodyType
	}

	if !tea.BoolValue(util.IsUnset(request.RequestBody)) {
		query["RequestBody"] = request.RequestBody
	}

	if !tea.BoolValue(util.IsUnset(request.RequestHeaders)) {
		query["RequestHeaders"] = request.RequestHeaders
	}

	if !tea.BoolValue(util.IsUnset(request.RequestQueries)) {
		query["RequestQueries"] = request.RequestQueries
	}

	if !tea.BoolValue(util.IsUnset(request.ResultSample)) {
		query["ResultSample"] = request.ResultSample
	}

	if !tea.BoolValue(util.IsUnset(request.ResultType)) {
		query["ResultType"] = request.ResultType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceAddress)) {
		query["ServiceAddress"] = request.ServiceAddress
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceProtocol)) {
		query["ServiceProtocol"] = request.ServiceProtocol
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceTimeout)) {
		query["ServiceTimeout"] = request.ServiceTimeout
	}

	if !tea.BoolValue(util.IsUnset(request.SystemParameters)) {
		query["SystemParameters"] = request.SystemParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApi(request *ModifyApiRequest) (_result *ModifyApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiResponse{}
	_body, _err := client.ModifyApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyApiForInnerWithOptions(request *ModifyApiForInnerRequest, runtime *util.RuntimeOptions) (_result *ModifyApiForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.ApiName)) {
		query["ApiName"] = request.ApiName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthType)) {
		query["AuthType"] = request.AuthType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RequestConfig)) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestParamters)) {
		query["RequestParamters"] = request.RequestParamters
	}

	if !tea.BoolValue(util.IsUnset(request.ResultSample)) {
		query["ResultSample"] = request.ResultSample
	}

	if !tea.BoolValue(util.IsUnset(request.ResultType)) {
		query["ResultType"] = request.ResultType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceConfig)) {
		query["ServiceConfig"] = request.ServiceConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParameters)) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceParametersMap)) {
		query["ServiceParametersMap"] = request.ServiceParametersMap
	}

	if !tea.BoolValue(util.IsUnset(request.Visibility)) {
		query["Visibility"] = request.Visibility
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApiForInner(request *ModifyApiForInnerRequest) (_result *ModifyApiForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiForInnerResponse{}
	_body, _err := client.ModifyApiForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyApiGroupWithOptions(request *ModifyApiGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyApiGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiGroup"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApiGroup(request *ModifyApiGroupRequest) (_result *ModifyApiGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiGroupResponse{}
	_body, _err := client.ModifyApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyApiMarketInstanceAttributeWithOptions(request *ModifyApiMarketInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *ModifyApiMarketInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAttributes)) {
		query["InstanceAttributes"] = request.InstanceAttributes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiMarketInstanceAttribute"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiMarketInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApiMarketInstanceAttribute(request *ModifyApiMarketInstanceAttributeRequest) (_result *ModifyApiMarketInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiMarketInstanceAttributeResponse{}
	_body, _err := client.ModifyApiMarketInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApp"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppForInnerWithOptions(request *ModifyAppForInnerRequest, runtime *util.RuntimeOptions) (_result *ModifyAppForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Extend)) {
		query["Extend"] = request.Extend
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppForInner(request *ModifyAppForInnerRequest) (_result *ModifyAppForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppForInnerResponse{}
	_body, _err := client.ModifyAppForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyGroupAuthAppCodeForBackendWithOptions(request *ModifyGroupAuthAppCodeForBackendRequest, runtime *util.RuntimeOptions) (_result *ModifyGroupAuthAppCodeForBackendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AuthAppCode)) {
		query["AuthAppCode"] = request.AuthAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGroupAuthAppCodeForBackend"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGroupAuthAppCodeForBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyGroupAuthAppCodeForBackend(request *ModifyGroupAuthAppCodeForBackendRequest) (_result *ModifyGroupAuthAppCodeForBackendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGroupAuthAppCodeForBackendResponse{}
	_body, _err := client.ModifyGroupAuthAppCodeForBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecretKeyWithOptions(request *ModifySecretKeyRequest, runtime *util.RuntimeOptions) (_result *ModifySecretKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SecretKey)) {
		query["SecretKey"] = request.SecretKey
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKeyId)) {
		query["SecretKeyId"] = request.SecretKeyId
	}

	if !tea.BoolValue(util.IsUnset(request.SecretKeyName)) {
		query["SecretKeyName"] = request.SecretKeyName
	}

	if !tea.BoolValue(util.IsUnset(request.SecretValue)) {
		query["SecretValue"] = request.SecretValue
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecretKey"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecretKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecretKey(request *ModifySecretKeyRequest) (_result *ModifySecretKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecretKeyResponse{}
	_body, _err := client.ModifySecretKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTrafficControlWithOptions(request *ModifyTrafficControlRequest, runtime *util.RuntimeOptions) (_result *ModifyTrafficControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiDefault)) {
		query["ApiDefault"] = request.ApiDefault
	}

	if !tea.BoolValue(util.IsUnset(request.AppDefault)) {
		query["AppDefault"] = request.AppDefault
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlId)) {
		query["TrafficControlId"] = request.TrafficControlId
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlName)) {
		query["TrafficControlName"] = request.TrafficControlName
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficControlUnit)) {
		query["TrafficControlUnit"] = request.TrafficControlUnit
	}

	if !tea.BoolValue(util.IsUnset(request.UserDefault)) {
		query["UserDefault"] = request.UserDefault
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTrafficControl"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTrafficControl(request *ModifyTrafficControlRequest) (_result *ModifyTrafficControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTrafficControlResponse{}
	_body, _err := client.ModifyTrafficControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserWhiteListValueByTypeWithOptions(request *ModifyUserWhiteListValueByTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyUserWhiteListValueByTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Uid)) {
		query["Uid"] = request.Uid
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUserWhiteListValueByType"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserWhiteListValueByTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserWhiteListValueByType(request *ModifyUserWhiteListValueByTypeRequest) (_result *ModifyUserWhiteListValueByTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserWhiteListValueByTypeResponse{}
	_body, _err := client.ModifyUserWhiteListValueByTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReactivateDomainForBackendWithOptions(request *ReactivateDomainForBackendRequest, runtime *util.RuntimeOptions) (_result *ReactivateDomainForBackendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Aliuid)) {
		query["Aliuid"] = request.Aliuid
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReactivateDomainForBackend"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReactivateDomainForBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReactivateDomainForBackend(request *ReactivateDomainForBackendRequest) (_result *ReactivateDomainForBackendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReactivateDomainForBackendResponse{}
	_body, _err := client.ReactivateDomainForBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverApiFromHistoricalWithOptions(request *RecoverApiFromHistoricalRequest, runtime *util.RuntimeOptions) (_result *RecoverApiFromHistoricalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryVersion)) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoverApiFromHistorical"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoverApiFromHistoricalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverApiFromHistorical(request *RecoverApiFromHistoricalRequest) (_result *RecoverApiFromHistoricalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoverApiFromHistoricalResponse{}
	_body, _err := client.RecoverApiFromHistoricalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoveryApiDefineFromHistoricalWithOptions(request *RecoveryApiDefineFromHistoricalRequest, runtime *util.RuntimeOptions) (_result *RecoveryApiDefineFromHistoricalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryVersion)) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoveryApiDefineFromHistorical"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoveryApiDefineFromHistoricalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoveryApiDefineFromHistorical(request *RecoveryApiDefineFromHistoricalRequest) (_result *RecoveryApiDefineFromHistoricalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoveryApiDefineFromHistoricalResponse{}
	_body, _err := client.RecoveryApiDefineFromHistoricalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoveryApiFromHistoricalWithOptions(request *RecoveryApiFromHistoricalRequest, runtime *util.RuntimeOptions) (_result *RecoveryApiFromHistoricalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryVersion)) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RecoveryApiFromHistorical"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RecoveryApiFromHistoricalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoveryApiFromHistorical(request *RecoveryApiFromHistoricalRequest) (_result *RecoveryApiFromHistoricalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoveryApiFromHistoricalResponse{}
	_body, _err := client.RecoveryApiFromHistoricalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshDomainWithOptions(request *RefreshDomainRequest, runtime *util.RuntimeOptions) (_result *RefreshDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshDomain"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshDomain(request *RefreshDomainRequest) (_result *RefreshDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RefreshDomainResponse{}
	_body, _err := client.RefreshDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAccessPermissionByApisWithOptions(request *RemoveAccessPermissionByApisRequest, runtime *util.RuntimeOptions) (_result *RemoveAccessPermissionByApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiIds)) {
		query["ApiIds"] = request.ApiIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAccessPermissionByApis"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAccessPermissionByApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAccessPermissionByApis(request *RemoveAccessPermissionByApisRequest) (_result *RemoveAccessPermissionByApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAccessPermissionByApisResponse{}
	_body, _err := client.RemoveAccessPermissionByApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAccessPermissionByAppsWithOptions(request *RemoveAccessPermissionByAppsRequest, runtime *util.RuntimeOptions) (_result *RemoveAccessPermissionByAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAccessPermissionByApps"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAccessPermissionByAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAccessPermissionByApps(request *RemoveAccessPermissionByAppsRequest) (_result *RemoveAccessPermissionByAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAccessPermissionByAppsResponse{}
	_body, _err := client.RemoveAccessPermissionByAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAccessPermissionByAppsForInnerWithOptions(request *RemoveAccessPermissionByAppsForInnerRequest, runtime *util.RuntimeOptions) (_result *RemoveAccessPermissionByAppsForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAccessPermissionByAppsForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAccessPermissionByAppsForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAccessPermissionByAppsForInner(request *RemoveAccessPermissionByAppsForInnerRequest) (_result *RemoveAccessPermissionByAppsForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAccessPermissionByAppsForInnerResponse{}
	_body, _err := client.RemoveAccessPermissionByAppsForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAllBlackListWithOptions(request *RemoveAllBlackListRequest, runtime *util.RuntimeOptions) (_result *RemoveAllBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAllBlackList"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAllBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAllBlackList(request *RemoveAllBlackListRequest) (_result *RemoveAllBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAllBlackListResponse{}
	_body, _err := client.RemoveAllBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveApiRuleWithOptions(request *RemoveApiRuleRequest, runtime *util.RuntimeOptions) (_result *RemoveApiRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiIds)) {
		query["ApiIds"] = request.ApiIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveApiRule"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveApiRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveApiRule(request *RemoveApiRuleRequest) (_result *RemoveApiRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveApiRuleResponse{}
	_body, _err := client.RemoveApiRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAppsFromApiWithOptions(request *RemoveAppsFromApiRequest, runtime *util.RuntimeOptions) (_result *RemoveAppsFromApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAppsFromApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAppsFromApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAppsFromApi(request *RemoveAppsFromApiRequest) (_result *RemoveAppsFromApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAppsFromApiResponse{}
	_body, _err := client.RemoveAppsFromApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveBlackListWithOptions(request *RemoveBlackListRequest, runtime *util.RuntimeOptions) (_result *RemoveBlackListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BlackContent)) {
		query["BlackContent"] = request.BlackContent
	}

	if !tea.BoolValue(util.IsUnset(request.BlackType)) {
		query["BlackType"] = request.BlackType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveBlackList"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveBlackListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveBlackList(request *RemoveBlackListRequest) (_result *RemoveBlackListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveBlackListResponse{}
	_body, _err := client.RemoveBlackListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAppCodeWithOptions(request *ResetAppCodeRequest, runtime *util.RuntimeOptions) (_result *ResetAppCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAppCode"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAppCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAppCode(request *ResetAppCodeRequest) (_result *ResetAppCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAppCodeResponse{}
	_body, _err := client.ResetAppCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAppCodeForInnerWithOptions(request *ResetAppCodeForInnerRequest, runtime *util.RuntimeOptions) (_result *ResetAppCodeForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppCode)) {
		query["AppCode"] = request.AppCode
	}

	if !tea.BoolValue(util.IsUnset(request.NewAppCode)) {
		query["NewAppCode"] = request.NewAppCode
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAppCodeForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAppCodeForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAppCodeForInner(request *ResetAppCodeForInnerRequest) (_result *ResetAppCodeForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAppCodeForInnerResponse{}
	_body, _err := client.ResetAppCodeForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAppKeySecretWithOptions(request *ResetAppKeySecretRequest, runtime *util.RuntimeOptions) (_result *ResetAppKeySecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAppKeySecret"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAppKeySecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAppKeySecret(request *ResetAppKeySecretRequest) (_result *ResetAppKeySecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAppKeySecretResponse{}
	_body, _err := client.ResetAppKeySecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSecretByAppKeyForInnerWithOptions(request *ResetSecretByAppKeyForInnerRequest, runtime *util.RuntimeOptions) (_result *ResetSecretByAppKeyForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.NewAppKey)) {
		query["NewAppKey"] = request.NewAppKey
	}

	if !tea.BoolValue(util.IsUnset(request.NewAppSecret)) {
		query["NewAppSecret"] = request.NewAppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSecretByAppKeyForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSecretByAppKeyForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSecretByAppKeyForInner(request *ResetSecretByAppKeyForInnerRequest) (_result *ResetSecretByAppKeyForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSecretByAppKeyForInnerResponse{}
	_body, _err := client.ResetSecretByAppKeyForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAccessPermissionByApisWithOptions(request *SetAccessPermissionByApisRequest, runtime *util.RuntimeOptions) (_result *SetAccessPermissionByApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiIds)) {
		query["ApiIds"] = request.ApiIds
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccessPermissionByApis"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccessPermissionByApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAccessPermissionByApis(request *SetAccessPermissionByApisRequest) (_result *SetAccessPermissionByApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccessPermissionByApisResponse{}
	_body, _err := client.SetAccessPermissionByApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAccessPermissionByGroupForBackendWithOptions(request *SetAccessPermissionByGroupForBackendRequest, runtime *util.RuntimeOptions) (_result *SetAccessPermissionByGroupForBackendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccessPermissionByGroupForBackend"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccessPermissionByGroupForBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAccessPermissionByGroupForBackend(request *SetAccessPermissionByGroupForBackendRequest) (_result *SetAccessPermissionByGroupForBackendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccessPermissionByGroupForBackendResponse{}
	_body, _err := client.SetAccessPermissionByGroupForBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAccessPermissionsWithOptions(request *SetAccessPermissionsRequest, runtime *util.RuntimeOptions) (_result *SetAccessPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccessPermissions"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccessPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAccessPermissions(request *SetAccessPermissionsRequest) (_result *SetAccessPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccessPermissionsResponse{}
	_body, _err := client.SetAccessPermissionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAccessPermissionsForInnerWithOptions(request *SetAccessPermissionsForInnerRequest, runtime *util.RuntimeOptions) (_result *SetAccessPermissionsForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.AppIds)) {
		query["AppIds"] = request.AppIds
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAccessPermissionsForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAccessPermissionsForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAccessPermissionsForInner(request *SetAccessPermissionsForInnerRequest) (_result *SetAccessPermissionsForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAccessPermissionsForInnerResponse{}
	_body, _err := client.SetAccessPermissionsForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetApiRuleWithOptions(request *SetApiRuleRequest, runtime *util.RuntimeOptions) (_result *SetApiRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiIds)) {
		query["ApiIds"] = request.ApiIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetApiRule"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetApiRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetApiRule(request *SetApiRuleRequest) (_result *SetApiRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetApiRuleResponse{}
	_body, _err := client.SetApiRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainWithOptions(request *SetDomainRequest, runtime *util.RuntimeOptions) (_result *SetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateBody)) {
		query["CertificateBody"] = request.CertificateBody
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateName)) {
		query["CertificateName"] = request.CertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomain"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomain(request *SetDomainRequest) (_result *SetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainResponse{}
	_body, _err := client.SetDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainCertificateWithOptions(request *SetDomainCertificateRequest, runtime *util.RuntimeOptions) (_result *SetDomainCertificateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateBody)) {
		query["CertificateBody"] = request.CertificateBody
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateName)) {
		query["CertificateName"] = request.CertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateKey)) {
		query["PrivateKey"] = request.PrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomainCertificate"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDomainCertificateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainCertificate(request *SetDomainCertificateRequest) (_result *SetDomainCertificateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainCertificateResponse{}
	_body, _err := client.SetDomainCertificateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDomainForBackendWithOptions(request *SetDomainForBackendRequest, runtime *util.RuntimeOptions) (_result *SetDomainForBackendResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CertificateBody)) {
		query["CertificateBody"] = request.CertificateBody
	}

	if !tea.BoolValue(util.IsUnset(request.CertificateName)) {
		query["CertificateName"] = request.CertificateName
	}

	if !tea.BoolValue(util.IsUnset(request.CertificatePrivateKey)) {
		query["CertificatePrivateKey"] = request.CertificatePrivateKey
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		query["DomainName"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomainForBackend"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDomainForBackendResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainForBackend(request *SetDomainForBackendRequest) (_result *SetDomainForBackendResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainForBackendResponse{}
	_body, _err := client.SetDomainForBackendWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetMockIntegrationWithOptions(request *SetMockIntegrationRequest, runtime *util.RuntimeOptions) (_result *SetMockIntegrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Mock)) {
		query["Mock"] = request.Mock
	}

	if !tea.BoolValue(util.IsUnset(request.MockResult)) {
		query["MockResult"] = request.MockResult
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetMockIntegration"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetMockIntegrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetMockIntegration(request *SetMockIntegrationRequest) (_result *SetMockIntegrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMockIntegrationResponse{}
	_body, _err := client.SetMockIntegrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetPurchasedApiGroupStatusWithOptions(request *SetPurchasedApiGroupStatusRequest, runtime *util.RuntimeOptions) (_result *SetPurchasedApiGroupStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.BillingStatus)) {
		query["BillingStatus"] = request.BillingStatus
	}

	if !tea.BoolValue(util.IsUnset(request.CloseOrder)) {
		query["CloseOrder"] = request.CloseOrder
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPurchasedApiGroupStatus"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPurchasedApiGroupStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPurchasedApiGroupStatus(request *SetPurchasedApiGroupStatusRequest) (_result *SetPurchasedApiGroupStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPurchasedApiGroupStatusResponse{}
	_body, _err := client.SetPurchasedApiGroupStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchApiWithOptions(request *SwitchApiRequest, runtime *util.RuntimeOptions) (_result *SwitchApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryVersion)) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchApi"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchApi(request *SwitchApiRequest) (_result *SwitchApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchApiResponse{}
	_body, _err := client.SwitchApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchApiForInnerWithOptions(request *SwitchApiForInnerRequest, runtime *util.RuntimeOptions) (_result *SwitchApiForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiId)) {
		query["ApiId"] = request.ApiId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryVersion)) {
		query["HistoryVersion"] = request.HistoryVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchApiForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchApiForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchApiForInner(request *SwitchApiForInnerRequest) (_result *SwitchApiForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchApiForInnerResponse{}
	_body, _err := client.SwitchApiForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SynCreateAppForInnerWithOptions(request *SynCreateAppForInnerRequest, runtime *util.RuntimeOptions) (_result *SynCreateAppForInnerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AliUid)) {
		query["AliUid"] = request.AliUid
	}

	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		query["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.AppSecret)) {
		query["AppSecret"] = request.AppSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SynCreateAppForInner"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SynCreateAppForInnerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SynCreateAppForInner(request *SynCreateAppForInnerRequest) (_result *SynCreateAppForInnerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SynCreateAppForInnerResponse{}
	_body, _err := client.SynCreateAppForInnerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesSystemTagsWithOptions(request *TagResourcesSystemTagsRequest, runtime *util.RuntimeOptions) (_result *TagResourcesSystemTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TagOwnerBid)) {
		query["TagOwnerBid"] = request.TagOwnerBid
	}

	if !tea.BoolValue(util.IsUnset(request.TagOwnerUid)) {
		query["TagOwnerUid"] = request.TagOwnerUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResourcesSystemTags"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesSystemTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResourcesSystemTags(request *TagResourcesSystemTagsRequest) (_result *TagResourcesSystemTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesSystemTagsResponse{}
	_body, _err := client.TagResourcesSystemTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesSystemTagsWithOptions(request *UntagResourcesSystemTagsRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesSystemTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	if !tea.BoolValue(util.IsUnset(request.TagOwnerBid)) {
		query["TagOwnerBid"] = request.TagOwnerBid
	}

	if !tea.BoolValue(util.IsUnset(request.TagOwnerUid)) {
		query["TagOwnerUid"] = request.TagOwnerUid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResourcesSystemTags"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesSystemTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResourcesSystemTags(request *UntagResourcesSystemTagsRequest) (_result *UntagResourcesSystemTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesSystemTagsResponse{}
	_body, _err := client.UntagResourcesSystemTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VipMigrationWithOptions(request *VipMigrationRequest, runtime *util.RuntimeOptions) (_result *VipMigrationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewVip)) {
		query["NewVip"] = request.NewVip
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalVip)) {
		query["OriginalVip"] = request.OriginalVip
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VipMigration"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &VipMigrationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VipMigration(request *VipMigrationRequest) (_result *VipMigrationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VipMigrationResponse{}
	_body, _err := client.VipMigrationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VipMigrationByUidWithOptions(request *VipMigrationByUidRequest, runtime *util.RuntimeOptions) (_result *VipMigrationByUidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewVip)) {
		query["NewVip"] = request.NewVip
	}

	if !tea.BoolValue(util.IsUnset(request.OriginalVip)) {
		query["OriginalVip"] = request.OriginalVip
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VipMigrationByUid"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VipMigrationByUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VipMigrationByUid(request *VipMigrationByUidRequest) (_result *VipMigrationByUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VipMigrationByUidResponse{}
	_body, _err := client.VipMigrationByUidWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcAddAppServerWithOptions(request *VpcAddAppServerRequest, runtime *util.RuntimeOptions) (_result *VpcAddAppServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressPoolId)) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIp)) {
		query["ServerIp"] = request.ServerIp
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcAddAppServer"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcAddAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcAddAppServer(request *VpcAddAppServerRequest) (_result *VpcAddAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcAddAppServerResponse{}
	_body, _err := client.VpcAddAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcCreateAddressPoolWithOptions(request *VpcCreateAddressPoolRequest, runtime *util.RuntimeOptions) (_result *VpcCreateAddressPoolResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcCreateAddressPool"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcCreateAddressPoolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcCreateAddressPool(request *VpcCreateAddressPoolRequest) (_result *VpcCreateAddressPoolResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcCreateAddressPoolResponse{}
	_body, _err := client.VpcCreateAddressPoolWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcQueryAppServersWithOptions(request *VpcQueryAppServersRequest, runtime *util.RuntimeOptions) (_result *VpcQueryAppServersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressPoolId)) {
		query["AddressPoolId"] = request.AddressPoolId
	}

	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIp)) {
		query["ServerIp"] = request.ServerIp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcQueryAppServers"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcQueryAppServersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcQueryAppServers(request *VpcQueryAppServersRequest) (_result *VpcQueryAppServersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcQueryAppServersResponse{}
	_body, _err := client.VpcQueryAppServersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcRegisterAppWithOptions(request *VpcRegisterAppRequest, runtime *util.RuntimeOptions) (_result *VpcRegisterAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcRegisterApp"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcRegisterAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcRegisterApp(request *VpcRegisterAppRequest) (_result *VpcRegisterAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcRegisterAppResponse{}
	_body, _err := client.VpcRegisterAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcRemoveAppServerWithOptions(request *VpcRemoveAppServerRequest, runtime *util.RuntimeOptions) (_result *VpcRemoveAppServerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServerIp)) {
		query["ServerIp"] = request.ServerIp
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcRemoveAppServer"),
		Version:     tea.String("2016-02-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcRemoveAppServerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcRemoveAppServer(request *VpcRemoveAppServerRequest) (_result *VpcRemoveAppServerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcRemoveAppServerResponse{}
	_body, _err := client.VpcRemoveAppServerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
