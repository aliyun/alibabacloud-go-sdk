// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetErrorCodeSolutionsRequest struct {
	AcceptLanguage *string `json:"acceptLanguage,omitempty" xml:"acceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234-56789012
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// An error occurred while processing your request.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// oss
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetErrorCodeSolutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetErrorCodeSolutionsRequest) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsRequest) SetAcceptLanguage(v string) *GetErrorCodeSolutionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) SetErrorCode(v string) *GetErrorCodeSolutionsRequest {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) SetErrorMessage(v string) *GetErrorCodeSolutionsRequest {
	s.ErrorMessage = &v
	return s
}

func (s *GetErrorCodeSolutionsRequest) SetProduct(v string) *GetErrorCodeSolutionsRequest {
	s.Product = &v
	return s
}

type GetErrorCodeSolutionsResponseBody struct {
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Solutions []*GetErrorCodeSolutionsResponseBodySolutions `json:"solutions,omitempty" xml:"solutions,omitempty" type:"Repeated"`
}

func (s GetErrorCodeSolutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetErrorCodeSolutionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsResponseBody) SetRequestId(v string) *GetErrorCodeSolutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBody) SetSolutions(v []*GetErrorCodeSolutionsResponseBodySolutions) *GetErrorCodeSolutionsResponseBody {
	s.Solutions = v
	return s
}

type GetErrorCodeSolutionsResponseBodySolutions struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 0017-00000502
	ErrorCode    *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Product      *string `json:"product,omitempty" xml:"product,omitempty"`
	ProductName  *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// 0017-00000502
	SolutionId *string `json:"solutionId,omitempty" xml:"solutionId,omitempty"`
}

func (s GetErrorCodeSolutionsResponseBodySolutions) String() string {
	return tea.Prettify(s)
}

func (s GetErrorCodeSolutionsResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetContent(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetErrorCode(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetErrorMessage(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.ErrorMessage = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetProduct(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.Product = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetProductName(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.ProductName = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetSolutionId(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.SolutionId = &v
	return s
}

type GetErrorCodeSolutionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorCodeSolutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorCodeSolutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetErrorCodeSolutionsResponse) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsResponse) SetHeaders(v map[string]*string) *GetErrorCodeSolutionsResponse {
	s.Headers = v
	return s
}

func (s *GetErrorCodeSolutionsResponse) SetStatusCode(v int32) *GetErrorCodeSolutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorCodeSolutionsResponse) SetBody(v *GetErrorCodeSolutionsResponseBody) *GetErrorCodeSolutionsResponse {
	s.Body = v
	return s
}

type GetOwnRequestLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123E4567-E89B-12D3-A456-426614174000
	LogRequestId *string `json:"logRequestId,omitempty" xml:"logRequestId,omitempty"`
}

func (s GetOwnRequestLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogRequest) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogRequest) SetLogRequestId(v string) *GetOwnRequestLogRequest {
	s.LogRequestId = &v
	return s
}

type GetOwnRequestLogResponseBody struct {
	LogInfo   *GetOwnRequestLogResponseBodyLogInfo `json:"logInfo,omitempty" xml:"logInfo,omitempty" type:"Struct"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetOwnRequestLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBody) SetLogInfo(v *GetOwnRequestLogResponseBodyLogInfo) *GetOwnRequestLogResponseBody {
	s.LogInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBody) SetRequestId(v string) *GetOwnRequestLogResponseBody {
	s.RequestId = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfo struct {
	AuthenticationInfo *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo `json:"authenticationInfo,omitempty" xml:"authenticationInfo,omitempty" type:"Struct"`
	BasicInfo          *GetOwnRequestLogResponseBodyLogInfoBasicInfo          `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" type:"Struct"`
	CallerInfo         *GetOwnRequestLogResponseBodyLogInfoCallerInfo         `json:"callerInfo,omitempty" xml:"callerInfo,omitempty" type:"Struct"`
	Parameters         []*GetOwnRequestLogResponseBodyLogInfoParameters       `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	Responses          *GetOwnRequestLogResponseBodyLogInfoResponses          `json:"responses,omitempty" xml:"responses,omitempty" type:"Struct"`
}

func (s GetOwnRequestLogResponseBodyLogInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetAuthenticationInfo(v *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) *GetOwnRequestLogResponseBodyLogInfo {
	s.AuthenticationInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetBasicInfo(v *GetOwnRequestLogResponseBodyLogInfoBasicInfo) *GetOwnRequestLogResponseBodyLogInfo {
	s.BasicInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetCallerInfo(v *GetOwnRequestLogResponseBodyLogInfoCallerInfo) *GetOwnRequestLogResponseBodyLogInfo {
	s.CallerInfo = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetParameters(v []*GetOwnRequestLogResponseBodyLogInfoParameters) *GetOwnRequestLogResponseBodyLogInfo {
	s.Parameters = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfo) SetResponses(v *GetOwnRequestLogResponseBodyLogInfoResponses) *GetOwnRequestLogResponseBodyLogInfo {
	s.Responses = v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo struct {
	AuthenticationType *string `json:"authenticationType,omitempty" xml:"authenticationType,omitempty"`
	// example:
	//
	// HMAC-SHA256
	SignatureMethod  *string `json:"signatureMethod,omitempty" xml:"signatureMethod,omitempty"`
	SignatureVersion *string `json:"signatureVersion,omitempty" xml:"signatureVersion,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) SetAuthenticationType(v string) *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.AuthenticationType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureMethod(v string) *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureMethod = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureVersion(v string) *GetOwnRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureVersion = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfo struct {
	AccessDeniedDetail *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty" type:"Struct"`
	Api                *string                                                         `json:"api,omitempty" xml:"api,omitempty"`
	ApiDoc             *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc             `json:"apiDoc,omitempty" xml:"apiDoc,omitempty" type:"Struct"`
	ApiStyle           *string                                                         `json:"apiStyle,omitempty" xml:"apiStyle,omitempty"`
	ApiVersion         *string                                                         `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	Endpoint           *string                                                         `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	ErrorCode          *string                                                         `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMessage       *string                                                         `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	GatewayProcessTime *string                                                         `json:"gatewayProcessTime,omitempty" xml:"gatewayProcessTime,omitempty"`
	HttpMethod         *string                                                         `json:"httpMethod,omitempty" xml:"httpMethod,omitempty"`
	HttpStatusCode     *string                                                         `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	LogRequestId       *string                                                         `json:"logRequestId,omitempty" xml:"logRequestId,omitempty"`
	Product            *string                                                         `json:"product,omitempty" xml:"product,omitempty"`
	ProductName        *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName        `json:"productName,omitempty" xml:"productName,omitempty" type:"Struct"`
	RegionId           *string                                                         `json:"regionId,omitempty" xml:"regionId,omitempty"`
	RequestDuration    *string                                                         `json:"requestDuration,omitempty" xml:"requestDuration,omitempty"`
	SdkRequestTime     *string                                                         `json:"sdkRequestTime,omitempty" xml:"sdkRequestTime,omitempty"`
	ThrottlingResult   *string                                                         `json:"throttlingResult,omitempty" xml:"throttlingResult,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetAccessDeniedDetail(v *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApi(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.Api = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApiDoc(v *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiDoc = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApiStyle(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiStyle = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetApiVersion(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiVersion = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetEndpoint(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.Endpoint = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetErrorCode(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorCode = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetErrorMessage(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorMessage = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetGatewayProcessTime(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.GatewayProcessTime = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetHttpMethod(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpMethod = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetHttpStatusCode(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpStatusCode = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetLogRequestId(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.LogRequestId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetProduct(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.Product = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetProductName(v *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ProductName = v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetRegionId(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.RegionId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetRequestDuration(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.RequestDuration = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetSdkRequestTime(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.SdkRequestTime = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfo) SetThrottlingResult(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfo {
	s.ThrottlingResult = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail struct {
	AuthAction               *string `json:"authAction,omitempty" xml:"authAction,omitempty"`
	AuthPrincipalDisplayName *string `json:"authPrincipalDisplayName,omitempty" xml:"authPrincipalDisplayName,omitempty"`
	AuthPrincipalOwnerId     *string `json:"authPrincipalOwnerId,omitempty" xml:"authPrincipalOwnerId,omitempty"`
	AuthPrincipalType        *string `json:"authPrincipalType,omitempty" xml:"authPrincipalType,omitempty"`
	EncodedDiagnosticMessage *string `json:"encodedDiagnosticMessage,omitempty" xml:"encodedDiagnosticMessage,omitempty"`
	NoPermissionType         *string `json:"noPermissionType,omitempty" xml:"noPermissionType,omitempty"`
	PolicyType               *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthAction(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalType(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetNoPermissionType(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetPolicyType(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc struct {
	AlibabacloudSite *string `json:"alibabacloudSite,omitempty" xml:"alibabacloudSite,omitempty"`
	AliyunSite       *string `json:"aliyunSite,omitempty" xml:"aliyunSite,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAlibabacloudSite(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AlibabacloudSite = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAliyunSite(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AliyunSite = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName struct {
	CnName *string `json:"cnName,omitempty" xml:"cnName,omitempty"`
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) SetCnName(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.CnName = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName) SetEnName(v string) *GetOwnRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.EnName = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoCallerInfo struct {
	CallerAccountId *string `json:"callerAccountId,omitempty" xml:"callerAccountId,omitempty"`
	CallerIp        *string `json:"callerIp,omitempty" xml:"callerIp,omitempty"`
	CallerType      *string `json:"callerType,omitempty" xml:"callerType,omitempty"`
	MasterAccountId *string `json:"masterAccountId,omitempty" xml:"masterAccountId,omitempty"`
	UserAgent       *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoCallerInfo) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoCallerInfo) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetCallerAccountId(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerAccountId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetCallerIp(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerIp = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetCallerType(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerType = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetMasterAccountId(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.MasterAccountId = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoCallerInfo) SetUserAgent(v string) *GetOwnRequestLogResponseBodyLogInfoCallerInfo {
	s.UserAgent = &v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoParameters struct {
	Name     *string     `json:"name,omitempty" xml:"name,omitempty"`
	Required *bool       `json:"required,omitempty" xml:"required,omitempty"`
	Type     *string     `json:"type,omitempty" xml:"type,omitempty"`
	Value    interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoParameters) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoParameters) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetName(v string) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Name = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetRequired(v bool) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Required = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetType(v string) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Type = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoParameters) SetValue(v interface{}) *GetOwnRequestLogResponseBodyLogInfoParameters {
	s.Value = v
	return s
}

type GetOwnRequestLogResponseBodyLogInfoResponses struct {
	ResponseBody       *string `json:"responseBody,omitempty" xml:"responseBody,omitempty"`
	ResponseBodyFormat *string `json:"responseBodyFormat,omitempty" xml:"responseBodyFormat,omitempty"`
}

func (s GetOwnRequestLogResponseBodyLogInfoResponses) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponseBodyLogInfoResponses) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) SetResponseBody(v string) *GetOwnRequestLogResponseBodyLogInfoResponses {
	s.ResponseBody = &v
	return s
}

func (s *GetOwnRequestLogResponseBodyLogInfoResponses) SetResponseBodyFormat(v string) *GetOwnRequestLogResponseBodyLogInfoResponses {
	s.ResponseBodyFormat = &v
	return s
}

type GetOwnRequestLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOwnRequestLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOwnRequestLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOwnRequestLogResponse) GoString() string {
	return s.String()
}

func (s *GetOwnRequestLogResponse) SetHeaders(v map[string]*string) *GetOwnRequestLogResponse {
	s.Headers = v
	return s
}

func (s *GetOwnRequestLogResponse) SetStatusCode(v int32) *GetOwnRequestLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOwnRequestLogResponse) SetBody(v *GetOwnRequestLogResponseBody) *GetOwnRequestLogResponse {
	s.Body = v
	return s
}

type GetRequestLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123E4567-E89B-12D3-A456-426614174000
	LogRequestId *string `json:"logRequestId,omitempty" xml:"logRequestId,omitempty"`
}

func (s GetRequestLogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogRequest) GoString() string {
	return s.String()
}

func (s *GetRequestLogRequest) SetLogRequestId(v string) *GetRequestLogRequest {
	s.LogRequestId = &v
	return s
}

type GetRequestLogResponseBody struct {
	LogInfo *GetRequestLogResponseBodyLogInfo `json:"logInfo,omitempty" xml:"logInfo,omitempty" type:"Struct"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRequestLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBody) SetLogInfo(v *GetRequestLogResponseBodyLogInfo) *GetRequestLogResponseBody {
	s.LogInfo = v
	return s
}

func (s *GetRequestLogResponseBody) SetRequestId(v string) *GetRequestLogResponseBody {
	s.RequestId = &v
	return s
}

type GetRequestLogResponseBodyLogInfo struct {
	AuthenticationInfo *GetRequestLogResponseBodyLogInfoAuthenticationInfo `json:"authenticationInfo,omitempty" xml:"authenticationInfo,omitempty" type:"Struct"`
	BasicInfo          *GetRequestLogResponseBodyLogInfoBasicInfo          `json:"basicInfo,omitempty" xml:"basicInfo,omitempty" type:"Struct"`
	CallerInfo         *GetRequestLogResponseBodyLogInfoCallerInfo         `json:"callerInfo,omitempty" xml:"callerInfo,omitempty" type:"Struct"`
	Parameters         []*GetRequestLogResponseBodyLogInfoParameters       `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	Responses          *GetRequestLogResponseBodyLogInfoResponses          `json:"responses,omitempty" xml:"responses,omitempty" type:"Struct"`
}

func (s GetRequestLogResponseBodyLogInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfo) SetAuthenticationInfo(v *GetRequestLogResponseBodyLogInfoAuthenticationInfo) *GetRequestLogResponseBodyLogInfo {
	s.AuthenticationInfo = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetBasicInfo(v *GetRequestLogResponseBodyLogInfoBasicInfo) *GetRequestLogResponseBodyLogInfo {
	s.BasicInfo = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetCallerInfo(v *GetRequestLogResponseBodyLogInfoCallerInfo) *GetRequestLogResponseBodyLogInfo {
	s.CallerInfo = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetParameters(v []*GetRequestLogResponseBodyLogInfoParameters) *GetRequestLogResponseBodyLogInfo {
	s.Parameters = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfo) SetResponses(v *GetRequestLogResponseBodyLogInfoResponses) *GetRequestLogResponseBodyLogInfo {
	s.Responses = v
	return s
}

type GetRequestLogResponseBodyLogInfoAuthenticationInfo struct {
	// example:
	//
	// AK
	AuthenticationType *string `json:"authenticationType,omitempty" xml:"authenticationType,omitempty"`
	// example:
	//
	// HMAC-SHA256
	SignatureMethod *string `json:"signatureMethod,omitempty" xml:"signatureMethod,omitempty"`
	// example:
	//
	// unknown
	SignatureVersion *string `json:"signatureVersion,omitempty" xml:"signatureVersion,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoAuthenticationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoAuthenticationInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) SetAuthenticationType(v string) *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.AuthenticationType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureMethod(v string) *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureMethod = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoAuthenticationInfo) SetSignatureVersion(v string) *GetRequestLogResponseBodyLogInfoAuthenticationInfo {
	s.SignatureVersion = &v
	return s
}

type GetRequestLogResponseBodyLogInfoBasicInfo struct {
	AccessDeniedDetail *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// RunInstances
	Api    *string                                          `json:"api,omitempty" xml:"api,omitempty"`
	ApiDoc *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc `json:"apiDoc,omitempty" xml:"apiDoc,omitempty" type:"Struct"`
	// example:
	//
	// roa
	ApiStyle *string `json:"apiStyle,omitempty" xml:"apiStyle,omitempty"`
	// example:
	//
	// 2024-11-30
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// example:
	//
	// ecs.cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// IncorrectStatus.TransitRouter
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// The resource is not in a valid state for the operation.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 2025-01-21T07:43:06Z
	GatewayProcessTime *string `json:"gatewayProcessTime,omitempty" xml:"gatewayProcessTime,omitempty"`
	// example:
	//
	// GET
	HttpMethod *string `json:"httpMethod,omitempty" xml:"httpMethod,omitempty"`
	// example:
	//
	// 404
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// 123E4567-E89B-12D3-A456-426614174000
	LogRequestId *string `json:"logRequestId,omitempty" xml:"logRequestId,omitempty"`
	// example:
	//
	// Ecs
	Product     *string                                               `json:"product,omitempty" xml:"product,omitempty"`
	ProductName *GetRequestLogResponseBodyLogInfoBasicInfoProductName `json:"productName,omitempty" xml:"productName,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 188
	RequestDuration *string `json:"requestDuration,omitempty" xml:"requestDuration,omitempty"`
	// example:
	//
	// 2025-01-21T07:43:06Z
	SdkRequestTime *string `json:"sdkRequestTime,omitempty" xml:"sdkRequestTime,omitempty"`
	// example:
	//
	// FC.PASS
	ThrottlingResult *string `json:"throttlingResult,omitempty" xml:"throttlingResult,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetAccessDeniedDetail(v *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.AccessDeniedDetail = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApi(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.Api = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApiDoc(v *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiDoc = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApiStyle(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiStyle = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetApiVersion(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ApiVersion = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetEndpoint(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.Endpoint = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetErrorCode(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorCode = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetErrorMessage(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ErrorMessage = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetGatewayProcessTime(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.GatewayProcessTime = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetHttpMethod(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpMethod = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetHttpStatusCode(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetLogRequestId(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.LogRequestId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetProduct(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.Product = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetProductName(v *GetRequestLogResponseBodyLogInfoBasicInfoProductName) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ProductName = v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetRegionId(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.RegionId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetRequestDuration(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.RequestDuration = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetSdkRequestTime(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.SdkRequestTime = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfo) SetThrottlingResult(v string) *GetRequestLogResponseBodyLogInfoBasicInfo {
	s.ThrottlingResult = &v
	return s
}

type GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail struct {
	// example:
	//
	// openapiexplorer:GetRequestLog
	AuthAction *string `json:"authAction,omitempty" xml:"authAction,omitempty"`
	// example:
	//
	// 205618123456123456
	AuthPrincipalDisplayName *string `json:"authPrincipalDisplayName,omitempty" xml:"authPrincipalDisplayName,omitempty"`
	// example:
	//
	// 1001234561234567
	AuthPrincipalOwnerId *string `json:"authPrincipalOwnerId,omitempty" xml:"authPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"authPrincipalType,omitempty" xml:"authPrincipalType,omitempty"`
	// example:
	//
	// -
	EncodedDiagnosticMessage *string `json:"encodedDiagnosticMessage,omitempty" xml:"encodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"noPermissionType,omitempty" xml:"noPermissionType,omitempty"`
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthAction(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetAuthPrincipalType(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetNoPermissionType(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail) SetPolicyType(v string) *GetRequestLogResponseBodyLogInfoBasicInfoAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

type GetRequestLogResponseBodyLogInfoBasicInfoApiDoc struct {
	// example:
	//
	// https://api.alibabacloud.com/document/Ecs/2014-05-26/RunInstances
	AlibabacloudSite *string `json:"alibabacloudSite,omitempty" xml:"alibabacloudSite,omitempty"`
	// example:
	//
	// https://api.aliyun.com/document/Ecs/2014-05-26/RunInstances
	AliyunSite *string `json:"aliyunSite,omitempty" xml:"aliyunSite,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAlibabacloudSite(v string) *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AlibabacloudSite = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc) SetAliyunSite(v string) *GetRequestLogResponseBodyLogInfoBasicInfoApiDoc {
	s.AliyunSite = &v
	return s
}

type GetRequestLogResponseBodyLogInfoBasicInfoProductName struct {
	CnName *string `json:"cnName,omitempty" xml:"cnName,omitempty"`
	// example:
	//
	// Elastic Compute Service
	EnName *string `json:"enName,omitempty" xml:"enName,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoProductName) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoBasicInfoProductName) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) SetCnName(v string) *GetRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.CnName = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoBasicInfoProductName) SetEnName(v string) *GetRequestLogResponseBodyLogInfoBasicInfoProductName {
	s.EnName = &v
	return s
}

type GetRequestLogResponseBodyLogInfoCallerInfo struct {
	// example:
	//
	// 241009849925897811
	CallerAccountId *string `json:"callerAccountId,omitempty" xml:"callerAccountId,omitempty"`
	// example:
	//
	// 100.68.xxx.xxx
	CallerIp *string `json:"callerIp,omitempty" xml:"callerIp,omitempty"`
	// example:
	//
	// sub
	CallerType *string `json:"callerType,omitempty" xml:"callerType,omitempty"`
	// example:
	//
	// 1973374733454118
	MasterAccountId *string `json:"masterAccountId,omitempty" xml:"masterAccountId,omitempty"`
	// example:
	//
	// AlibabaCloud API Workbench
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoCallerInfo) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoCallerInfo) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetCallerAccountId(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerAccountId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetCallerIp(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerIp = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetCallerType(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.CallerType = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetMasterAccountId(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.MasterAccountId = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoCallerInfo) SetUserAgent(v string) *GetRequestLogResponseBodyLogInfoCallerInfo {
	s.UserAgent = &v
	return s
}

type GetRequestLogResponseBodyLogInfoParameters struct {
	// example:
	//
	// InstanceType
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// false
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// ecs.g6.large
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoParameters) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoParameters) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetName(v string) *GetRequestLogResponseBodyLogInfoParameters {
	s.Name = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetRequired(v bool) *GetRequestLogResponseBodyLogInfoParameters {
	s.Required = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetType(v string) *GetRequestLogResponseBodyLogInfoParameters {
	s.Type = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoParameters) SetValue(v interface{}) *GetRequestLogResponseBodyLogInfoParameters {
	s.Value = v
	return s
}

type GetRequestLogResponseBodyLogInfoResponses struct {
	// example:
	//
	// -
	ResponseBody *string `json:"responseBody,omitempty" xml:"responseBody,omitempty"`
	// example:
	//
	// JSON
	ResponseBodyFormat *string `json:"responseBodyFormat,omitempty" xml:"responseBodyFormat,omitempty"`
}

func (s GetRequestLogResponseBodyLogInfoResponses) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponseBodyLogInfoResponses) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponseBodyLogInfoResponses) SetResponseBody(v string) *GetRequestLogResponseBodyLogInfoResponses {
	s.ResponseBody = &v
	return s
}

func (s *GetRequestLogResponseBodyLogInfoResponses) SetResponseBodyFormat(v string) *GetRequestLogResponseBodyLogInfoResponses {
	s.ResponseBodyFormat = &v
	return s
}

type GetRequestLogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRequestLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRequestLogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRequestLogResponse) GoString() string {
	return s.String()
}

func (s *GetRequestLogResponse) SetHeaders(v map[string]*string) *GetRequestLogResponse {
	s.Headers = v
	return s
}

func (s *GetRequestLogResponse) SetStatusCode(v int32) *GetRequestLogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRequestLogResponse) SetBody(v *GetRequestLogResponseBody) *GetRequestLogResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("openapiexplorer"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// Summary:
//
// 根据提供的错误码获取对应的解决方案
//
// Description:
//
// ## 请求说明
//
// - 本接口支持通过POST或GET方法调用。
//
// - `Accept-Language`请求头必须设置为`zh-CN`或`en-US`之一，用于指定返回结果的语言类型。
//
// - 错误码格式需符合特定规则，特别是针对OSS的错误码应遵循正则表达式`[0-9]{4}-[0-9]{8}`。
//
// - 当前实现中未使用`maxResults`和`nextToken`参数。
//
// - 如果请求失败，将根据不同的错误情况返回相应的错误代码及描述信息。
//
// @param request - GetErrorCodeSolutionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErrorCodeSolutionsResponse
func (client *Client) GetErrorCodeSolutionsWithOptions(request *GetErrorCodeSolutionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetErrorCodeSolutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorCode)) {
		query["errorCode"] = request.ErrorCode
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorMessage)) {
		query["errorMessage"] = request.ErrorMessage
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetErrorCodeSolutions"),
		Version:     tea.String("2024-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/getErrorCodeSolutions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetErrorCodeSolutionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetErrorCodeSolutionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 根据提供的错误码获取对应的解决方案
//
// Description:
//
// ## 请求说明
//
// - 本接口支持通过POST或GET方法调用。
//
// - `Accept-Language`请求头必须设置为`zh-CN`或`en-US`之一，用于指定返回结果的语言类型。
//
// - 错误码格式需符合特定规则，特别是针对OSS的错误码应遵循正则表达式`[0-9]{4}-[0-9]{8}`。
//
// - 当前实现中未使用`maxResults`和`nextToken`参数。
//
// - 如果请求失败，将根据不同的错误情况返回相应的错误代码及描述信息。
//
// @param request - GetErrorCodeSolutionsRequest
//
// @return GetErrorCodeSolutionsResponse
func (client *Client) GetErrorCodeSolutions(request *GetErrorCodeSolutionsRequest) (_result *GetErrorCodeSolutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetErrorCodeSolutionsResponse{}
	_body, _err := client.GetErrorCodeSolutionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过API RequestId 查询当前账号调用OpenAPI的日志详情，用于故障排查。
//
// Description:
//
// ## 请求说明
//
// - 本接口主要用于帮助用户通过提供具体的`apiRequestId`来获取相关API请求的详细日志信息。
//
// - `apiRequestId`必须是大写形式的UUID，并且应确保该ID确实来自于您之前对某个OpenAPI的实际调用。
//
// - 如果提供的`apiRequestId`无效或者没有找到对应的日志记录，系统将返回相应的错误提示。
//
// - 在使用此接口时，请注意检查您的网络环境以及权限设置，以保证能够顺利访问到所需资源。
//
// @param request - GetOwnRequestLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOwnRequestLogResponse
func (client *Client) GetOwnRequestLogWithOptions(request *GetOwnRequestLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetOwnRequestLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		query["logRequestId"] = request.LogRequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOwnRequestLog"),
		Version:     tea.String("2024-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/getOwnRequestLog"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetOwnRequestLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetOwnRequestLogResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 通过API RequestId 查询当前账号调用OpenAPI的日志详情，用于故障排查。
//
// Description:
//
// ## 请求说明
//
// - 本接口主要用于帮助用户通过提供具体的`apiRequestId`来获取相关API请求的详细日志信息。
//
// - `apiRequestId`必须是大写形式的UUID，并且应确保该ID确实来自于您之前对某个OpenAPI的实际调用。
//
// - 如果提供的`apiRequestId`无效或者没有找到对应的日志记录，系统将返回相应的错误提示。
//
// - 在使用此接口时，请注意检查您的网络环境以及权限设置，以保证能够顺利访问到所需资源。
//
// @param request - GetOwnRequestLogRequest
//
// @return GetOwnRequestLogResponse
func (client *Client) GetOwnRequestLog(request *GetOwnRequestLogRequest) (_result *GetOwnRequestLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOwnRequestLogResponse{}
	_body, _err := client.GetOwnRequestLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通过API请求ID查询特定请求的日志详情，用于故障排查。
//
// Description:
//
// ## 请求说明
//
// - 本接口主要用于帮助用户通过提供具体的`apiRequestId`来获取相关API请求的详细日志信息。
//
// - `apiRequestId`必须是大写形式的UUID，并且应确保该ID确实来自于您之前对某个OpenAPI的实际调用。
//
// - 如果提供的`apiRequestId`无效或者没有找到对应的日志记录，系统将返回相应的错误提示。
//
// - 在使用此接口时，请注意检查您的网络环境以及权限设置，以保证能够顺利访问到所需资源。
//
// @param request - GetRequestLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRequestLogResponse
func (client *Client) GetRequestLogWithOptions(request *GetRequestLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRequestLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogRequestId)) {
		query["logRequestId"] = request.LogRequestId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRequestLog"),
		Version:     tea.String("2024-11-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/getRequestLog"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetRequestLogResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetRequestLogResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 通过API请求ID查询特定请求的日志详情，用于故障排查。
//
// Description:
//
// ## 请求说明
//
// - 本接口主要用于帮助用户通过提供具体的`apiRequestId`来获取相关API请求的详细日志信息。
//
// - `apiRequestId`必须是大写形式的UUID，并且应确保该ID确实来自于您之前对某个OpenAPI的实际调用。
//
// - 如果提供的`apiRequestId`无效或者没有找到对应的日志记录，系统将返回相应的错误提示。
//
// - 在使用此接口时，请注意检查您的网络环境以及权限设置，以保证能够顺利访问到所需资源。
//
// @param request - GetRequestLogRequest
//
// @return GetRequestLogResponse
func (client *Client) GetRequestLog(request *GetRequestLogRequest) (_result *GetRequestLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRequestLogResponse{}
	_body, _err := client.GetRequestLogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
