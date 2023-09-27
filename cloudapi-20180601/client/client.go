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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbolishApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type AddBlackListRequest struct {
	BlackContent  *string `json:"BlackContent,omitempty" xml:"BlackContent,omitempty"`
	BlackType     *string `json:"BlackType,omitempty" xml:"BlackType,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *AddBlackListRequest) SetSecurityToken(v string) *AddBlackListRequest {
	s.SecurityToken = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type AddIpControlPolicyItemRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CidrIp        *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s AddIpControlPolicyItemRequest) String() string {
	return tea.Prettify(s)
}

func (s AddIpControlPolicyItemRequest) GoString() string {
	return s.String()
}

func (s *AddIpControlPolicyItemRequest) SetAppId(v string) *AddIpControlPolicyItemRequest {
	s.AppId = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) SetCidrIp(v string) *AddIpControlPolicyItemRequest {
	s.CidrIp = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) SetIpControlId(v string) *AddIpControlPolicyItemRequest {
	s.IpControlId = &v
	return s
}

func (s *AddIpControlPolicyItemRequest) SetSecurityToken(v string) *AddIpControlPolicyItemRequest {
	s.SecurityToken = &v
	return s
}

type AddIpControlPolicyItemResponseBody struct {
	PolicyItemId *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIpControlPolicyItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddIpControlPolicyItemResponseBody) GoString() string {
	return s.String()
}

func (s *AddIpControlPolicyItemResponseBody) SetPolicyItemId(v string) *AddIpControlPolicyItemResponseBody {
	s.PolicyItemId = &v
	return s
}

func (s *AddIpControlPolicyItemResponseBody) SetRequestId(v string) *AddIpControlPolicyItemResponseBody {
	s.RequestId = &v
	return s
}

type AddIpControlPolicyItemResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddIpControlPolicyItemResponse) String() string {
	return tea.Prettify(s)
}

func (s AddIpControlPolicyItemResponse) GoString() string {
	return s.String()
}

func (s *AddIpControlPolicyItemResponse) SetHeaders(v map[string]*string) *AddIpControlPolicyItemResponse {
	s.Headers = v
	return s
}

func (s *AddIpControlPolicyItemResponse) SetStatusCode(v int32) *AddIpControlPolicyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIpControlPolicyItemResponse) SetBody(v *AddIpControlPolicyItemResponseBody) *AddIpControlPolicyItemResponse {
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AddTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateApiRequest struct {
	AllowSignatureMethod *string `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples     *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	FailResultSample     *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OpenIdConnectConfig  *string `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty"`
	RequestConfig        *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParamters     *string `json:"RequestParamters,omitempty" xml:"RequestParamters,omitempty"`
	ResultDescriptions   *string `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty"`
	ResultSample         *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType           *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	Visibility           *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	WebSocketApiType     *string `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s CreateApiRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiRequest) GoString() string {
	return s.String()
}

func (s *CreateApiRequest) SetAllowSignatureMethod(v string) *CreateApiRequest {
	s.AllowSignatureMethod = &v
	return s
}

func (s *CreateApiRequest) SetApiName(v string) *CreateApiRequest {
	s.ApiName = &v
	return s
}

func (s *CreateApiRequest) SetAuthType(v string) *CreateApiRequest {
	s.AuthType = &v
	return s
}

func (s *CreateApiRequest) SetDescription(v string) *CreateApiRequest {
	s.Description = &v
	return s
}

func (s *CreateApiRequest) SetErrorCodeSamples(v string) *CreateApiRequest {
	s.ErrorCodeSamples = &v
	return s
}

func (s *CreateApiRequest) SetFailResultSample(v string) *CreateApiRequest {
	s.FailResultSample = &v
	return s
}

func (s *CreateApiRequest) SetGroupId(v string) *CreateApiRequest {
	s.GroupId = &v
	return s
}

func (s *CreateApiRequest) SetOpenIdConnectConfig(v string) *CreateApiRequest {
	s.OpenIdConnectConfig = &v
	return s
}

func (s *CreateApiRequest) SetRequestConfig(v string) *CreateApiRequest {
	s.RequestConfig = &v
	return s
}

func (s *CreateApiRequest) SetRequestParamters(v string) *CreateApiRequest {
	s.RequestParamters = &v
	return s
}

func (s *CreateApiRequest) SetResultDescriptions(v string) *CreateApiRequest {
	s.ResultDescriptions = &v
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

func (s *CreateApiRequest) SetServiceConfig(v string) *CreateApiRequest {
	s.ServiceConfig = &v
	return s
}

func (s *CreateApiRequest) SetServiceParameters(v string) *CreateApiRequest {
	s.ServiceParameters = &v
	return s
}

func (s *CreateApiRequest) SetServiceParametersMap(v string) *CreateApiRequest {
	s.ServiceParametersMap = &v
	return s
}

func (s *CreateApiRequest) SetVisibility(v string) *CreateApiRequest {
	s.Visibility = &v
	return s
}

func (s *CreateApiRequest) SetWebSocketApiType(v string) *CreateApiRequest {
	s.WebSocketApiType = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateApiGroupRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s *CreateApiGroupRequest) SetInstanceId(v string) *CreateApiGroupRequest {
	s.InstanceId = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateApiStageVariableRequest struct {
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageId         *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageRouteModel *string `json:"StageRouteModel,omitempty" xml:"StageRouteModel,omitempty"`
	SupportRoute    *bool   `json:"SupportRoute,omitempty" xml:"SupportRoute,omitempty"`
	VariableName    *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	VariableValue   *string `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s CreateApiStageVariableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiStageVariableRequest) GoString() string {
	return s.String()
}

func (s *CreateApiStageVariableRequest) SetGroupId(v string) *CreateApiStageVariableRequest {
	s.GroupId = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetSecurityToken(v string) *CreateApiStageVariableRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetStageId(v string) *CreateApiStageVariableRequest {
	s.StageId = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetStageRouteModel(v string) *CreateApiStageVariableRequest {
	s.StageRouteModel = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetSupportRoute(v bool) *CreateApiStageVariableRequest {
	s.SupportRoute = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetVariableName(v string) *CreateApiStageVariableRequest {
	s.VariableName = &v
	return s
}

func (s *CreateApiStageVariableRequest) SetVariableValue(v string) *CreateApiStageVariableRequest {
	s.VariableValue = &v
	return s
}

type CreateApiStageVariableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApiStageVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiStageVariableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiStageVariableResponseBody) SetRequestId(v string) *CreateApiStageVariableResponseBody {
	s.RequestId = &v
	return s
}

type CreateApiStageVariableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApiStageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateApiStageVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApiStageVariableResponse) GoString() string {
	return s.String()
}

func (s *CreateApiStageVariableResponse) SetHeaders(v map[string]*string) *CreateApiStageVariableResponse {
	s.Headers = v
	return s
}

func (s *CreateApiStageVariableResponse) SetStatusCode(v int32) *CreateApiStageVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApiStageVariableResponse) SetBody(v *CreateApiStageVariableResponseBody) *CreateApiStageVariableResponse {
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateCustomizedInfoRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CsharpDemo    *string `json:"CsharpDemo,omitempty" xml:"CsharpDemo,omitempty"`
	CurlDemo      *string `json:"CurlDemo,omitempty" xml:"CurlDemo,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	JavaDemo      *string `json:"JavaDemo,omitempty" xml:"JavaDemo,omitempty"`
	ObjectcDemo   *string `json:"ObjectcDemo,omitempty" xml:"ObjectcDemo,omitempty"`
	PhpDemo       *string `json:"PhpDemo,omitempty" xml:"PhpDemo,omitempty"`
	PythonDemo    *string `json:"PythonDemo,omitempty" xml:"PythonDemo,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageId       *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s CreateCustomizedInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedInfoRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedInfoRequest) SetApiId(v string) *CreateCustomizedInfoRequest {
	s.ApiId = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetApiName(v string) *CreateCustomizedInfoRequest {
	s.ApiName = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetCsharpDemo(v string) *CreateCustomizedInfoRequest {
	s.CsharpDemo = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetCurlDemo(v string) *CreateCustomizedInfoRequest {
	s.CurlDemo = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetGroupId(v string) *CreateCustomizedInfoRequest {
	s.GroupId = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetJavaDemo(v string) *CreateCustomizedInfoRequest {
	s.JavaDemo = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetObjectcDemo(v string) *CreateCustomizedInfoRequest {
	s.ObjectcDemo = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetPhpDemo(v string) *CreateCustomizedInfoRequest {
	s.PhpDemo = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetPythonDemo(v string) *CreateCustomizedInfoRequest {
	s.PythonDemo = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetSecurityToken(v string) *CreateCustomizedInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetStageId(v string) *CreateCustomizedInfoRequest {
	s.StageId = &v
	return s
}

func (s *CreateCustomizedInfoRequest) SetStageName(v string) *CreateCustomizedInfoRequest {
	s.StageName = &v
	return s
}

type CreateCustomizedInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomizedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedInfoResponseBody) SetRequestId(v string) *CreateCustomizedInfoResponseBody {
	s.RequestId = &v
	return s
}

type CreateCustomizedInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCustomizedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCustomizedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedInfoResponse) SetHeaders(v map[string]*string) *CreateCustomizedInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedInfoResponse) SetStatusCode(v int32) *CreateCustomizedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedInfoResponse) SetBody(v *CreateCustomizedInfoResponseBody) *CreateCustomizedInfoResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	AccountQuantity *int32  `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	ExpiredOn       *string `json:"ExpiredOn,omitempty" xml:"ExpiredOn,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SkuId           *string `json:"SkuId,omitempty" xml:"SkuId,omitempty"`
	Token           *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAccountQuantity(v int32) *CreateInstanceRequest {
	s.AccountQuantity = &v
	return s
}

func (s *CreateInstanceRequest) SetExpiredOn(v string) *CreateInstanceRequest {
	s.ExpiredOn = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityToken(v string) *CreateInstanceRequest {
	s.SecurityToken = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateIpControlRequest struct {
	Description      *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	IpControlName    *string                                   `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	IpControlPolicys []*CreateIpControlRequestIpControlPolicys `json:"IpControlPolicys,omitempty" xml:"IpControlPolicys,omitempty" type:"Repeated"`
	IpControlType    *string                                   `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
	SecurityToken    *string                                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateIpControlRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIpControlRequest) GoString() string {
	return s.String()
}

func (s *CreateIpControlRequest) SetDescription(v string) *CreateIpControlRequest {
	s.Description = &v
	return s
}

func (s *CreateIpControlRequest) SetIpControlName(v string) *CreateIpControlRequest {
	s.IpControlName = &v
	return s
}

func (s *CreateIpControlRequest) SetIpControlPolicys(v []*CreateIpControlRequestIpControlPolicys) *CreateIpControlRequest {
	s.IpControlPolicys = v
	return s
}

func (s *CreateIpControlRequest) SetIpControlType(v string) *CreateIpControlRequest {
	s.IpControlType = &v
	return s
}

func (s *CreateIpControlRequest) SetSecurityToken(v string) *CreateIpControlRequest {
	s.SecurityToken = &v
	return s
}

type CreateIpControlRequestIpControlPolicys struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	IP    *string `json:"IP,omitempty" xml:"IP,omitempty"`
}

func (s CreateIpControlRequestIpControlPolicys) String() string {
	return tea.Prettify(s)
}

func (s CreateIpControlRequestIpControlPolicys) GoString() string {
	return s.String()
}

func (s *CreateIpControlRequestIpControlPolicys) SetAppId(v string) *CreateIpControlRequestIpControlPolicys {
	s.AppId = &v
	return s
}

func (s *CreateIpControlRequestIpControlPolicys) SetIP(v string) *CreateIpControlRequestIpControlPolicys {
	s.IP = &v
	return s
}

type CreateIpControlResponseBody struct {
	IpControlId *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpControlResponseBody) SetIpControlId(v string) *CreateIpControlResponseBody {
	s.IpControlId = &v
	return s
}

func (s *CreateIpControlResponseBody) SetRequestId(v string) *CreateIpControlResponseBody {
	s.RequestId = &v
	return s
}

type CreateIpControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIpControlResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIpControlResponse) GoString() string {
	return s.String()
}

func (s *CreateIpControlResponse) SetHeaders(v map[string]*string) *CreateIpControlResponse {
	s.Headers = v
	return s
}

func (s *CreateIpControlResponse) SetStatusCode(v int32) *CreateIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIpControlResponse) SetBody(v *CreateIpControlResponseBody) *CreateIpControlResponse {
	s.Body = v
	return s
}

type CreateLogConfigRequest struct {
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SlsLogStore   *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	SlsProject    *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s CreateLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLogConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateLogConfigRequest) SetLogType(v string) *CreateLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *CreateLogConfigRequest) SetSecurityToken(v string) *CreateLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLogConfigRequest) SetSlsLogStore(v string) *CreateLogConfigRequest {
	s.SlsLogStore = &v
	return s
}

func (s *CreateLogConfigRequest) SetSlsProject(v string) *CreateLogConfigRequest {
	s.SlsProject = &v
	return s
}

type CreateLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogConfigResponseBody) SetRequestId(v string) *CreateLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type CreateLogConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLogConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLogConfigResponse) SetHeaders(v map[string]*string) *CreateLogConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLogConfigResponse) SetStatusCode(v int32) *CreateLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogConfigResponse) SetBody(v *CreateLogConfigResponseBody) *CreateLogConfigResponse {
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRaceWorkForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAllTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteApiStageVariableRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageId       *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	VariableName  *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
}

func (s DeleteApiStageVariableRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiStageVariableRequest) GoString() string {
	return s.String()
}

func (s *DeleteApiStageVariableRequest) SetGroupId(v string) *DeleteApiStageVariableRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteApiStageVariableRequest) SetSecurityToken(v string) *DeleteApiStageVariableRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteApiStageVariableRequest) SetStageId(v string) *DeleteApiStageVariableRequest {
	s.StageId = &v
	return s
}

func (s *DeleteApiStageVariableRequest) SetVariableName(v string) *DeleteApiStageVariableRequest {
	s.VariableName = &v
	return s
}

type DeleteApiStageVariableResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApiStageVariableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiStageVariableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiStageVariableResponseBody) SetRequestId(v string) *DeleteApiStageVariableResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApiStageVariableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApiStageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteApiStageVariableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiStageVariableResponse) GoString() string {
	return s.String()
}

func (s *DeleteApiStageVariableResponse) SetHeaders(v map[string]*string) *DeleteApiStageVariableResponse {
	s.Headers = v
	return s
}

func (s *DeleteApiStageVariableResponse) SetStatusCode(v int32) *DeleteApiStageVariableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApiStageVariableResponse) SetBody(v *DeleteApiStageVariableResponseBody) *DeleteApiStageVariableResponse {
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteIpControlRequest struct {
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteIpControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpControlRequest) GoString() string {
	return s.String()
}

func (s *DeleteIpControlRequest) SetIpControlId(v string) *DeleteIpControlRequest {
	s.IpControlId = &v
	return s
}

func (s *DeleteIpControlRequest) SetSecurityToken(v string) *DeleteIpControlRequest {
	s.SecurityToken = &v
	return s
}

type DeleteIpControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpControlResponseBody) SetRequestId(v string) *DeleteIpControlResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIpControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteIpControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIpControlResponse) GoString() string {
	return s.String()
}

func (s *DeleteIpControlResponse) SetHeaders(v map[string]*string) *DeleteIpControlResponse {
	s.Headers = v
	return s
}

func (s *DeleteIpControlResponse) SetStatusCode(v int32) *DeleteIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIpControlResponse) SetBody(v *DeleteIpControlResponseBody) *DeleteIpControlResponse {
	s.Body = v
	return s
}

type DeleteLogConfigRequest struct {
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogConfigRequest) SetLogType(v string) *DeleteLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *DeleteLogConfigRequest) SetSecurityToken(v string) *DeleteLogConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogConfigResponseBody) SetRequestId(v string) *DeleteLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLogConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLogConfigResponse) SetHeaders(v map[string]*string) *DeleteLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLogConfigResponse) SetStatusCode(v int32) *DeleteLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLogConfigResponse) SetBody(v *DeleteLogConfigResponseBody) *DeleteLogConfigResponse {
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeployApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	SupportMock   *string `json:"SupportMock,omitempty" xml:"SupportMock,omitempty"`
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

func (s *DeployApiRequest) SetSupportMock(v string) *DeployApiRequest {
	s.SupportMock = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AllowSignatureMethod    *string                                         `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
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
	WebSocketApiType        *string                                         `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s DescribeApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBody) SetAllowSignatureMethod(v string) *DescribeApiResponseBody {
	s.AllowSignatureMethod = &v
	return s
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

func (s *DescribeApiResponseBody) SetWebSocketApiType(v string) *DescribeApiResponseBody {
	s.WebSocketApiType = &v
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
	RequestMode         *string `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
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

func (s *DescribeApiResponseBodyRequestConfig) SetRequestMode(v string) *DescribeApiResponseBodyRequestConfig {
	s.RequestMode = &v
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
	AoneAppName           *string                                                    `json:"AoneAppName,omitempty" xml:"AoneAppName,omitempty"`
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

func (s *DescribeApiResponseBodyServiceConfig) SetAoneAppName(v string) *DescribeApiResponseBodyServiceConfig {
	s.AoneAppName = &v
	return s
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApiId                   *string                                     `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                 *string                                     `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BodyFormat              *string                                     `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	DeployedTime            *string                                     `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description             *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples        *DescribeApiDocResponseBodyErrorCodeSamples `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample        *string                                     `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	GroupId                 *string                                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName               *string                                     `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpMethod              *string                                     `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol            *string                                     `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Mock                    *string                                     `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult              *string                                     `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	OriginResultDescription *string                                     `json:"OriginResultDescription,omitempty" xml:"OriginResultDescription,omitempty"`
	Path                    *string                                     `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters          *DescribeApiDocResponseBodyPathParameters   `json:"PathParameters,omitempty" xml:"PathParameters,omitempty" type:"Struct"`
	PostBodyDescription     *string                                     `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType            *string                                     `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RegionId                *string                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestBody             *DescribeApiDocResponseBodyRequestBody      `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	RequestHeaders          *DescribeApiDocResponseBodyRequestHeaders   `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Struct"`
	RequestId               *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestMode             *string                                     `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	RequestQueries          *DescribeApiDocResponseBodyRequestQueries   `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty" type:"Struct"`
	ResultSample            *string                                     `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType              *string                                     `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceTimeout          *int32                                      `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable        *string                                     `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	StageName               *string                                     `json:"StageName,omitempty" xml:"StageName,omitempty"`
	VpcName                 *string                                     `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
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

func (s *DescribeApiDocResponseBody) SetMock(v string) *DescribeApiDocResponseBody {
	s.Mock = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetMockResult(v string) *DescribeApiDocResponseBody {
	s.MockResult = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetOriginResultDescription(v string) *DescribeApiDocResponseBody {
	s.OriginResultDescription = &v
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

func (s *DescribeApiDocResponseBody) SetRequestMode(v string) *DescribeApiDocResponseBody {
	s.RequestMode = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestQueries(v *DescribeApiDocResponseBodyRequestQueries) *DescribeApiDocResponseBody {
	s.RequestQueries = v
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

func (s *DescribeApiDocResponseBody) SetServiceTimeout(v int32) *DescribeApiDocResponseBody {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetServiceVpcEnable(v string) *DescribeApiDocResponseBody {
	s.ServiceVpcEnable = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetStageName(v string) *DescribeApiDocResponseBody {
	s.StageName = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetVpcName(v string) *DescribeApiDocResponseBody {
	s.VpcName = &v
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

type DescribeApiDocResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiDocsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiErrorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	StageItems    *DescribeApiGroupDetailResponseBodyStageItems  `json:"StageItems,omitempty" xml:"StageItems,omitempty" type:"Struct"`
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

func (s *DescribeApiGroupDetailResponseBody) SetStageItems(v *DescribeApiGroupDetailResponseBodyStageItems) *DescribeApiGroupDetailResponseBody {
	s.StageItems = v
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
	CertificateId         *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName       *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainBindingStatus   *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	DomainLegalStatus     *string `json:"DomainLegalStatus,omitempty" xml:"DomainLegalStatus,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution  *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
	DomainRemark          *string `json:"DomainRemark,omitempty" xml:"DomainRemark,omitempty"`
	DomainWebSocketStatus *string `json:"DomainWebSocketStatus,omitempty" xml:"DomainWebSocketStatus,omitempty"`
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

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainBindingStatus(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainBindingStatus = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainLegalStatus(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainLegalStatus = &v
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

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainRemark(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainRemark = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem) SetDomainWebSocketStatus(v string) *DescribeApiGroupDetailResponseBodyDomainItemsDomainItem {
	s.DomainWebSocketStatus = &v
	return s
}

type DescribeApiGroupDetailResponseBodyStageItems struct {
	StageInfo []*DescribeApiGroupDetailResponseBodyStageItemsStageInfo `json:"StageInfo,omitempty" xml:"StageInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupDetailResponseBodyStageItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailResponseBodyStageItems) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailResponseBodyStageItems) SetStageInfo(v []*DescribeApiGroupDetailResponseBodyStageItemsStageInfo) *DescribeApiGroupDetailResponseBodyStageItems {
	s.StageInfo = v
	return s
}

type DescribeApiGroupDetailResponseBodyStageItemsStageInfo struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StageId     *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName   *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiGroupDetailResponseBodyStageItemsStageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailResponseBodyStageItemsStageInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailResponseBodyStageItemsStageInfo) SetDescription(v string) *DescribeApiGroupDetailResponseBodyStageItemsStageInfo {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyStageItemsStageInfo) SetStageId(v string) *DescribeApiGroupDetailResponseBodyStageItemsStageInfo {
	s.StageId = &v
	return s
}

func (s *DescribeApiGroupDetailResponseBodyStageItemsStageInfo) SetStageName(v string) *DescribeApiGroupDetailResponseBodyStageItemsStageInfo {
	s.StageName = &v
	return s
}

type DescribeApiGroupDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeApiGroupDetailForConsumerRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiGroupDetailForConsumerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailForConsumerRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailForConsumerRequest) SetApiId(v string) *DescribeApiGroupDetailForConsumerRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerRequest) SetGroupId(v string) *DescribeApiGroupDetailForConsumerRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerRequest) SetSecurityToken(v string) *DescribeApiGroupDetailForConsumerRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerRequest) SetStageName(v string) *DescribeApiGroupDetailForConsumerRequest {
	s.StageName = &v
	return s
}

type DescribeApiGroupDetailForConsumerResponseBody struct {
	BillingStatus *string                                                   `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	CreatedTime   *string                                                   `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainItems   *DescribeApiGroupDetailForConsumerResponseBodyDomainItems `json:"DomainItems,omitempty" xml:"DomainItems,omitempty" type:"Struct"`
	GroupId       *string                                                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                                   `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IllegalStatus *string                                                   `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	ModifiedTime  *string                                                   `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Purchased     *string                                                   `json:"Purchased,omitempty" xml:"Purchased,omitempty"`
	RegionId      *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId     *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	SubDomain     *string                                                   `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	TrafficLimit  *int32                                                    `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
}

func (s DescribeApiGroupDetailForConsumerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailForConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetBillingStatus(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.BillingStatus = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetCreatedTime(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetDescription(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetDomainItems(v *DescribeApiGroupDetailForConsumerResponseBodyDomainItems) *DescribeApiGroupDetailForConsumerResponseBody {
	s.DomainItems = v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetGroupId(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetGroupName(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetIllegalStatus(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetModifiedTime(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetPurchased(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.Purchased = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetRegionId(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetRequestId(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetStatus(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetSubDomain(v string) *DescribeApiGroupDetailForConsumerResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBody) SetTrafficLimit(v int32) *DescribeApiGroupDetailForConsumerResponseBody {
	s.TrafficLimit = &v
	return s
}

type DescribeApiGroupDetailForConsumerResponseBodyDomainItems struct {
	DomainItem []*DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupDetailForConsumerResponseBodyDomainItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailForConsumerResponseBodyDomainItems) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailForConsumerResponseBodyDomainItems) SetDomainItem(v []*DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) *DescribeApiGroupDetailForConsumerResponseBodyDomainItems {
	s.DomainItem = v
	return s
}

type DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem struct {
	CertificateId        *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName      *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainName           *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
	DomainStatus         *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
}

func (s DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) SetCertificateId(v string) *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem {
	s.CertificateId = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) SetCertificateName(v string) *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem {
	s.CertificateName = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) SetDomainName(v string) *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem {
	s.DomainName = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) SetDomainNameResolution(v string) *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem {
	s.DomainNameResolution = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem) SetDomainStatus(v string) *DescribeApiGroupDetailForConsumerResponseBodyDomainItemsDomainItem {
	s.DomainStatus = &v
	return s
}

type DescribeApiGroupDetailForConsumerResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiGroupDetailForConsumerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiGroupDetailForConsumerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupDetailForConsumerResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupDetailForConsumerResponse) SetHeaders(v map[string]*string) *DescribeApiGroupDetailForConsumerResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponse) SetStatusCode(v int32) *DescribeApiGroupDetailForConsumerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiGroupDetailForConsumerResponse) SetBody(v *DescribeApiGroupDetailForConsumerResponseBody) *DescribeApiGroupDetailForConsumerResponse {
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeApiIpControlsRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiIpControlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiIpControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsRequest) SetApiIds(v string) *DescribeApiIpControlsRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetGroupId(v string) *DescribeApiIpControlsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetPageNumber(v int32) *DescribeApiIpControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetPageSize(v int32) *DescribeApiIpControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetSecurityToken(v string) *DescribeApiIpControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiIpControlsRequest) SetStageName(v string) *DescribeApiIpControlsRequest {
	s.StageName = &v
	return s
}

type DescribeApiIpControlsResponseBody struct {
	ApiIpControls *DescribeApiIpControlsResponseBodyApiIpControls `json:"ApiIpControls,omitempty" xml:"ApiIpControls,omitempty" type:"Struct"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiIpControlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiIpControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponseBody) SetApiIpControls(v *DescribeApiIpControlsResponseBodyApiIpControls) *DescribeApiIpControlsResponseBody {
	s.ApiIpControls = v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetPageNumber(v int32) *DescribeApiIpControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetPageSize(v int32) *DescribeApiIpControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetRequestId(v string) *DescribeApiIpControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiIpControlsResponseBody) SetTotalCount(v int32) *DescribeApiIpControlsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiIpControlsResponseBodyApiIpControls struct {
	ApiIpControlItem []*DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem `json:"ApiIpControlItem,omitempty" xml:"ApiIpControlItem,omitempty" type:"Repeated"`
}

func (s DescribeApiIpControlsResponseBodyApiIpControls) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiIpControlsResponseBodyApiIpControls) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponseBodyApiIpControls) SetApiIpControlItem(v []*DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) *DescribeApiIpControlsResponseBodyApiIpControls {
	s.ApiIpControlItem = v
	return s
}

type DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BoundTime     *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
}

func (s DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetApiId(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetApiName(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetBoundTime(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetIpControlId(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.IpControlId = &v
	return s
}

func (s *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem) SetIpControlName(v string) *DescribeApiIpControlsResponseBodyApiIpControlsApiIpControlItem {
	s.IpControlName = &v
	return s
}

type DescribeApiIpControlsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiIpControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiIpControlsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiIpControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiIpControlsResponse) SetHeaders(v map[string]*string) *DescribeApiIpControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiIpControlsResponse) SetStatusCode(v int32) *DescribeApiIpControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiIpControlsResponse) SetBody(v *DescribeApiIpControlsResponseBody) *DescribeApiIpControlsResponse {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiLatencyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiQpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
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

func (s *DescribeApiRulesRequest) SetApiName(v string) *DescribeApiRulesRequest {
	s.ApiName = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeApiStageDetailRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageId       *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s DescribeApiStageDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailRequest) SetGroupId(v string) *DescribeApiStageDetailRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiStageDetailRequest) SetSecurityToken(v string) *DescribeApiStageDetailRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiStageDetailRequest) SetStageId(v string) *DescribeApiStageDetailRequest {
	s.StageId = &v
	return s
}

type DescribeApiStageDetailResponseBody struct {
	CreatedTime  *string                                      `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string                                      `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModifiedTime *string                                      `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StageId      *string                                      `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName    *string                                      `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Variables    *DescribeApiStageDetailResponseBodyVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Struct"`
}

func (s DescribeApiStageDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponseBody) SetCreatedTime(v string) *DescribeApiStageDetailResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetDescription(v string) *DescribeApiStageDetailResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetGroupId(v string) *DescribeApiStageDetailResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetModifiedTime(v string) *DescribeApiStageDetailResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetRequestId(v string) *DescribeApiStageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetStageId(v string) *DescribeApiStageDetailResponseBody {
	s.StageId = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetStageName(v string) *DescribeApiStageDetailResponseBody {
	s.StageName = &v
	return s
}

func (s *DescribeApiStageDetailResponseBody) SetVariables(v *DescribeApiStageDetailResponseBodyVariables) *DescribeApiStageDetailResponseBody {
	s.Variables = v
	return s
}

type DescribeApiStageDetailResponseBodyVariables struct {
	VariableItem []*DescribeApiStageDetailResponseBodyVariablesVariableItem `json:"VariableItem,omitempty" xml:"VariableItem,omitempty" type:"Repeated"`
}

func (s DescribeApiStageDetailResponseBodyVariables) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponseBodyVariables) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponseBodyVariables) SetVariableItem(v []*DescribeApiStageDetailResponseBodyVariablesVariableItem) *DescribeApiStageDetailResponseBodyVariables {
	s.VariableItem = v
	return s
}

type DescribeApiStageDetailResponseBodyVariablesVariableItem struct {
	StageRouteModel *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel `json:"StageRouteModel,omitempty" xml:"StageRouteModel,omitempty" type:"Struct"`
	SupportRoute    *bool                                                                   `json:"SupportRoute,omitempty" xml:"SupportRoute,omitempty"`
	VariableName    *string                                                                 `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
	VariableValue   *string                                                                 `json:"VariableValue,omitempty" xml:"VariableValue,omitempty"`
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItem) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItem) SetStageRouteModel(v *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) *DescribeApiStageDetailResponseBodyVariablesVariableItem {
	s.StageRouteModel = v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItem) SetSupportRoute(v bool) *DescribeApiStageDetailResponseBodyVariablesVariableItem {
	s.SupportRoute = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItem) SetVariableName(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItem {
	s.VariableName = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItem) SetVariableValue(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItem {
	s.VariableValue = &v
	return s
}

type DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel struct {
	Location             *string                                                                           `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterCatalog     *string                                                                           `json:"ParameterCatalog,omitempty" xml:"ParameterCatalog,omitempty"`
	ParameterType        *string                                                                           `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	RouteMatchSymbol     *string                                                                           `json:"RouteMatchSymbol,omitempty" xml:"RouteMatchSymbol,omitempty"`
	RouteRules           *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules `json:"RouteRules,omitempty" xml:"RouteRules,omitempty" type:"Struct"`
	ServiceParameterName *string                                                                           `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) SetLocation(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel {
	s.Location = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) SetParameterCatalog(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel {
	s.ParameterCatalog = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) SetParameterType(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) SetRouteMatchSymbol(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel {
	s.RouteMatchSymbol = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) SetRouteRules(v *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel {
	s.RouteRules = v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel) SetServiceParameterName(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModel {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules struct {
	RouteRuleItem []*DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem `json:"RouteRuleItem,omitempty" xml:"RouteRuleItem,omitempty" type:"Repeated"`
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules) SetRouteRuleItem(v []*DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRules {
	s.RouteRuleItem = v
	return s
}

type DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem struct {
	ConditionValue *string `json:"ConditionValue,omitempty" xml:"ConditionValue,omitempty"`
	MaxValue       *int64  `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue       *int64  `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	ResultValue    *string `json:"ResultValue,omitempty" xml:"ResultValue,omitempty"`
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) SetConditionValue(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem {
	s.ConditionValue = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) SetMaxValue(v int64) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) SetMinValue(v int64) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem {
	s.MinValue = &v
	return s
}

func (s *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem) SetResultValue(v string) *DescribeApiStageDetailResponseBodyVariablesVariableItemStageRouteModelRouteRulesRouteRuleItem {
	s.ResultValue = &v
	return s
}

type DescribeApiStageDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiStageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiStageDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiStageDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiStageDetailResponse) SetHeaders(v map[string]*string) *DescribeApiStageDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiStageDetailResponse) SetStatusCode(v int32) *DescribeApiStageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiStageDetailResponse) SetBody(v *DescribeApiStageDetailResponseBody) *DescribeApiStageDetailResponse {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApiTrafficResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApisByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeApisByIpControlRequest struct {
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApisByIpControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByIpControlRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlRequest) SetIpControlId(v string) *DescribeApisByIpControlRequest {
	s.IpControlId = &v
	return s
}

func (s *DescribeApisByIpControlRequest) SetPageNumber(v int32) *DescribeApisByIpControlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByIpControlRequest) SetPageSize(v int32) *DescribeApisByIpControlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByIpControlRequest) SetSecurityToken(v string) *DescribeApisByIpControlRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApisByIpControlResponseBody struct {
	ApiInfos   *DescribeApisByIpControlResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByIpControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponseBody) SetApiInfos(v *DescribeApisByIpControlResponseBodyApiInfos) *DescribeApisByIpControlResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetPageNumber(v int32) *DescribeApisByIpControlResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetPageSize(v int32) *DescribeApisByIpControlResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetRequestId(v string) *DescribeApisByIpControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBody) SetTotalCount(v int32) *DescribeApisByIpControlResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisByIpControlResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisByIpControlResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByIpControlResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByIpControlResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponseBodyApiInfos) SetApiInfo(v []*DescribeApisByIpControlResponseBodyApiInfosApiInfo) *DescribeApisByIpControlResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApisByIpControlResponseBodyApiInfosApiInfo struct {
	ApiId       *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName     *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BoundTime   *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName   *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility  *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApisByIpControlResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByIpControlResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetBoundTime(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.BoundTime = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByIpControlResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisByIpControlResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeApisByIpControlResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApisByIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApisByIpControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByIpControlResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByIpControlResponse) SetHeaders(v map[string]*string) *DescribeApisByIpControlResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByIpControlResponse) SetStatusCode(v int32) *DescribeApisByIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisByIpControlResponse) SetBody(v *DescribeApisByIpControlResponseBody) *DescribeApisByIpControlResponse {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApisByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeApisForConsoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppSecuritiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *DescribeAppSecurityRequest) SetSecurityToken(v string) *DescribeAppSecurityRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAppSecurityResponseBody struct {
	AppCode      *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
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

func (s *DescribeAppSecurityResponseBody) SetAppCode(v string) *DescribeAppSecurityResponseBody {
	s.AppCode = &v
	return s
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeAppsRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
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

func (s *DescribeAppsRequest) SetAppName(v string) *DescribeAppsRequest {
	s.AppName = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppsByApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AppOwnerId    *int64  `json:"AppOwnerId,omitempty" xml:"AppOwnerId,omitempty"`
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

func (s *DescribeAppsForProviderRequest) SetAppOwnerId(v int64) *DescribeAppsForProviderRequest {
	s.AppOwnerId = &v
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppsForProviderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBlackListsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	FailResultSample      *string                                               `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
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
	RequestMode           *string                                               `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
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

func (s *DescribeDeployedApiResponseBody) SetFailResultSample(v string) *DescribeDeployedApiResponseBody {
	s.FailResultSample = &v
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

func (s *DescribeDeployedApiResponseBody) SetRequestMode(v string) *DescribeDeployedApiResponseBody {
	s.RequestMode = &v
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeployedApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeployedApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CertificateBody       *string `json:"CertificateBody,omitempty" xml:"CertificateBody,omitempty"`
	CertificateId         *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName       *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	DomainBindingStatus   *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	DomainLegalStatus     *string `json:"DomainLegalStatus,omitempty" xml:"DomainLegalStatus,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainNameResolution  *string `json:"DomainNameResolution,omitempty" xml:"DomainNameResolution,omitempty"`
	DomainRemark          *string `json:"DomainRemark,omitempty" xml:"DomainRemark,omitempty"`
	DomainWebSocketStatus *string `json:"DomainWebSocketStatus,omitempty" xml:"DomainWebSocketStatus,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PrivateKey            *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain             *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
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

func (s *DescribeDomainResponseBody) SetDomainBindingStatus(v string) *DescribeDomainResponseBody {
	s.DomainBindingStatus = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainLegalStatus(v string) *DescribeDomainResponseBody {
	s.DomainLegalStatus = &v
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

func (s *DescribeDomainResponseBody) SetDomainRemark(v string) *DescribeDomainResponseBody {
	s.DomainRemark = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainWebSocketStatus(v string) *DescribeDomainResponseBody {
	s.DomainWebSocketStatus = &v
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDomainResolutionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AllowSignatureMethod    *string                                               `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiId                   *string                                               `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                 *string                                               `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType                *string                                               `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BodyFormat              *string                                               `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	ConstantParameters      *DescribeHistoryApiResponseBodyConstantParameters     `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	CustomSystemParameters  *DescribeHistoryApiResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	DeployedTime            *string                                               `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description             *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples        *DescribeHistoryApiResponseBodyErrorCodeSamples       `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample        *string                                               `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	FunctionComputeConfig   *DescribeHistoryApiResponseBodyFunctionComputeConfig  `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	GroupId                 *string                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName               *string                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HistoryVersion          *string                                               `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	HttpMethod              *string                                               `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	HttpProtocol            *string                                               `json:"HttpProtocol,omitempty" xml:"HttpProtocol,omitempty"`
	Mock                    *string                                               `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult              *string                                               `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	OpenIdConnectConfig     *DescribeHistoryApiResponseBodyOpenIdConnectConfig    `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	OriginResultDescription *string                                               `json:"OriginResultDescription,omitempty" xml:"OriginResultDescription,omitempty"`
	Path                    *string                                               `json:"Path,omitempty" xml:"Path,omitempty"`
	PathParameters          *DescribeHistoryApiResponseBodyPathParameters         `json:"PathParameters,omitempty" xml:"PathParameters,omitempty" type:"Struct"`
	PostBodyDescription     *string                                               `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	PostBodyType            *string                                               `json:"PostBodyType,omitempty" xml:"PostBodyType,omitempty"`
	RegionId                *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestBody             *DescribeHistoryApiResponseBodyRequestBody            `json:"RequestBody,omitempty" xml:"RequestBody,omitempty" type:"Struct"`
	RequestHeaders          *DescribeHistoryApiResponseBodyRequestHeaders         `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Struct"`
	RequestId               *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestMode             *string                                               `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	RequestQueries          *DescribeHistoryApiResponseBodyRequestQueries         `json:"RequestQueries,omitempty" xml:"RequestQueries,omitempty" type:"Struct"`
	ResultSample            *string                                               `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType              *string                                               `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceAddress          *string                                               `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceFCEnable         *string                                               `json:"ServiceFCEnable,omitempty" xml:"ServiceFCEnable,omitempty"`
	ServiceProtocol         *string                                               `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout          *int32                                                `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable        *string                                               `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	Status                  *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemParameters        *DescribeHistoryApiResponseBodySystemParameters       `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility              *string                                               `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	VpcName                 *string                                               `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeHistoryApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBody) SetAllowSignatureMethod(v string) *DescribeHistoryApiResponseBody {
	s.AllowSignatureMethod = &v
	return s
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

func (s *DescribeHistoryApiResponseBody) SetCustomSystemParameters(v *DescribeHistoryApiResponseBodyCustomSystemParameters) *DescribeHistoryApiResponseBody {
	s.CustomSystemParameters = v
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

func (s *DescribeHistoryApiResponseBody) SetFailResultSample(v string) *DescribeHistoryApiResponseBody {
	s.FailResultSample = &v
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

func (s *DescribeHistoryApiResponseBody) SetOpenIdConnectConfig(v *DescribeHistoryApiResponseBodyOpenIdConnectConfig) *DescribeHistoryApiResponseBody {
	s.OpenIdConnectConfig = v
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

func (s *DescribeHistoryApiResponseBody) SetRequestMode(v string) *DescribeHistoryApiResponseBody {
	s.RequestMode = &v
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

type DescribeHistoryApiResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApiResponseBodyCustomSystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeHistoryApiResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

type DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) SetDemoValue(v string) *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) SetDescription(v string) *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) SetLocation(v string) *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) SetParameterName(v string) *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter) SetServiceParameterName(v string) *DescribeHistoryApiResponseBodyCustomSystemParametersCustomSystemParameter {
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

type DescribeHistoryApiResponseBodyOpenIdConnectConfig struct {
	IdTokenParamName *string `json:"IdTokenParamName,omitempty" xml:"IdTokenParamName,omitempty"`
	OpenIdApiType    *string `json:"OpenIdApiType,omitempty" xml:"OpenIdApiType,omitempty"`
	PublicKey        *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	PublicKeyId      *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
}

func (s DescribeHistoryApiResponseBodyOpenIdConnectConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApiResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApiResponseBodyOpenIdConnectConfig) SetIdTokenParamName(v string) *DescribeHistoryApiResponseBodyOpenIdConnectConfig {
	s.IdTokenParamName = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyOpenIdConnectConfig) SetOpenIdApiType(v string) *DescribeHistoryApiResponseBodyOpenIdConnectConfig {
	s.OpenIdApiType = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyOpenIdConnectConfig) SetPublicKey(v string) *DescribeHistoryApiResponseBodyOpenIdConnectConfig {
	s.PublicKey = &v
	return s
}

func (s *DescribeHistoryApiResponseBodyOpenIdConnectConfig) SetPublicKeyId(v string) *DescribeHistoryApiResponseBodyOpenIdConnectConfig {
	s.PublicKeyId = &v
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHistoryApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber    *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *DescribeHistoryApisRequest) SetPageNumber(v string) *DescribeHistoryApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHistoryApisRequest) SetPageSize(v string) *DescribeHistoryApisRequest {
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeHistoryApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribeIpControlPolicyItemsRequest struct {
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyItemId  *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeIpControlPolicyItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlPolicyItemsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsRequest) SetIpControlId(v string) *DescribeIpControlPolicyItemsRequest {
	s.IpControlId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetPageNumber(v int32) *DescribeIpControlPolicyItemsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetPageSize(v int32) *DescribeIpControlPolicyItemsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetPolicyItemId(v string) *DescribeIpControlPolicyItemsRequest {
	s.PolicyItemId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsRequest) SetSecurityToken(v string) *DescribeIpControlPolicyItemsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeIpControlPolicyItemsResponseBody struct {
	IpControlPolicyItems *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems `json:"IpControlPolicyItems,omitempty" xml:"IpControlPolicyItems,omitempty" type:"Struct"`
	PageNumber           *int32                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId            *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount           *int32                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpControlPolicyItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetIpControlPolicyItems(v *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) *DescribeIpControlPolicyItemsResponseBody {
	s.IpControlPolicyItems = v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetPageNumber(v int32) *DescribeIpControlPolicyItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetPageSize(v int32) *DescribeIpControlPolicyItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetRequestId(v string) *DescribeIpControlPolicyItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBody) SetTotalCount(v int32) *DescribeIpControlPolicyItemsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems struct {
	IpControlPolicyItem []*DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem `json:"IpControlPolicyItem,omitempty" xml:"IpControlPolicyItem,omitempty" type:"Repeated"`
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems) SetIpControlPolicyItem(v []*DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItems {
	s.IpControlPolicyItem = v
	return s
}

type DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem struct {
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CidrIp       *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PolicyItemId *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetAppId(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.AppId = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetCidrIp(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.CidrIp = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetCreateTime(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.CreateTime = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetModifiedTime(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem) SetPolicyItemId(v string) *DescribeIpControlPolicyItemsResponseBodyIpControlPolicyItemsIpControlPolicyItem {
	s.PolicyItemId = &v
	return s
}

type DescribeIpControlPolicyItemsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIpControlPolicyItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpControlPolicyItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlPolicyItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpControlPolicyItemsResponse) SetHeaders(v map[string]*string) *DescribeIpControlPolicyItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpControlPolicyItemsResponse) SetStatusCode(v int32) *DescribeIpControlPolicyItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpControlPolicyItemsResponse) SetBody(v *DescribeIpControlPolicyItemsResponseBody) *DescribeIpControlPolicyItemsResponse {
	s.Body = v
	return s
}

type DescribeIpControlsRequest struct {
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	IpControlType *string `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeIpControlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsRequest) SetIpControlId(v string) *DescribeIpControlsRequest {
	s.IpControlId = &v
	return s
}

func (s *DescribeIpControlsRequest) SetIpControlName(v string) *DescribeIpControlsRequest {
	s.IpControlName = &v
	return s
}

func (s *DescribeIpControlsRequest) SetIpControlType(v string) *DescribeIpControlsRequest {
	s.IpControlType = &v
	return s
}

func (s *DescribeIpControlsRequest) SetPageNumber(v int32) *DescribeIpControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlsRequest) SetPageSize(v int32) *DescribeIpControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlsRequest) SetSecurityToken(v string) *DescribeIpControlsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeIpControlsResponseBody struct {
	IpControlInfos *DescribeIpControlsResponseBodyIpControlInfos `json:"IpControlInfos,omitempty" xml:"IpControlInfos,omitempty" type:"Struct"`
	PageNumber     *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpControlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponseBody) SetIpControlInfos(v *DescribeIpControlsResponseBodyIpControlInfos) *DescribeIpControlsResponseBody {
	s.IpControlInfos = v
	return s
}

func (s *DescribeIpControlsResponseBody) SetPageNumber(v int32) *DescribeIpControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpControlsResponseBody) SetPageSize(v int32) *DescribeIpControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpControlsResponseBody) SetRequestId(v string) *DescribeIpControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpControlsResponseBody) SetTotalCount(v int32) *DescribeIpControlsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeIpControlsResponseBodyIpControlInfos struct {
	IpControlInfo []*DescribeIpControlsResponseBodyIpControlInfosIpControlInfo `json:"IpControlInfo,omitempty" xml:"IpControlInfo,omitempty" type:"Repeated"`
}

func (s DescribeIpControlsResponseBodyIpControlInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlsResponseBodyIpControlInfos) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponseBodyIpControlInfos) SetIpControlInfo(v []*DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) *DescribeIpControlsResponseBodyIpControlInfos {
	s.IpControlInfo = v
	return s
}

type DescribeIpControlsResponseBodyIpControlInfosIpControlInfo struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	IpControlType *string `json:"IpControlType,omitempty" xml:"IpControlType,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetCreateTime(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetDescription(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.Description = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetIpControlId(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.IpControlId = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetIpControlName(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.IpControlName = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetIpControlType(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.IpControlType = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetModifiedTime(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo) SetRegionId(v string) *DescribeIpControlsResponseBodyIpControlInfosIpControlInfo {
	s.RegionId = &v
	return s
}

type DescribeIpControlsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeIpControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeIpControlsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeIpControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIpControlsResponse) SetHeaders(v map[string]*string) *DescribeIpControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIpControlsResponse) SetStatusCode(v int32) *DescribeIpControlsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIpControlsResponse) SetBody(v *DescribeIpControlsResponseBody) *DescribeIpControlsResponse {
	s.Body = v
	return s
}

type DescribeLogConfigRequest struct {
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigRequest) SetLogType(v string) *DescribeLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogConfigRequest) SetSecurityToken(v string) *DescribeLogConfigRequest {
	s.SecurityToken = &v
	return s
}

type DescribeLogConfigResponseBody struct {
	LogInfos  *DescribeLogConfigResponseBodyLogInfos `json:"LogInfos,omitempty" xml:"LogInfos,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponseBody) SetLogInfos(v *DescribeLogConfigResponseBodyLogInfos) *DescribeLogConfigResponseBody {
	s.LogInfos = v
	return s
}

func (s *DescribeLogConfigResponseBody) SetRequestId(v string) *DescribeLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type DescribeLogConfigResponseBodyLogInfos struct {
	LogInfo []*DescribeLogConfigResponseBodyLogInfosLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogConfigResponseBodyLogInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogConfigResponseBodyLogInfos) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponseBodyLogInfos) SetLogInfo(v []*DescribeLogConfigResponseBodyLogInfosLogInfo) *DescribeLogConfigResponseBodyLogInfos {
	s.LogInfo = v
	return s
}

type DescribeLogConfigResponseBodyLogInfosLogInfo struct {
	LogType     *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SlsLogStore *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	SlsProject  *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s DescribeLogConfigResponseBodyLogInfosLogInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogConfigResponseBodyLogInfosLogInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetLogType(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.LogType = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetRegionId(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetSlsLogStore(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.SlsLogStore = &v
	return s
}

func (s *DescribeLogConfigResponseBodyLogInfosLogInfo) SetSlsProject(v string) *DescribeLogConfigResponseBodyLogInfosLogInfo {
	s.SlsProject = &v
	return s
}

type DescribeLogConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigResponse) SetHeaders(v map[string]*string) *DescribeLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogConfigResponse) SetStatusCode(v int32) *DescribeLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogConfigResponse) SetBody(v *DescribeLogConfigResponseBody) *DescribeLogConfigResponse {
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePurchasedApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePurchasedApiGroupDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DescribePurchasedApiGroupsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePurchasedApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
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

func (s *DescribePurchasedApisResponseBodyApiInfosApiInfo) SetDeployedTime(v string) *DescribePurchasedApisResponseBodyApiInfosApiInfo {
	s.DeployedTime = &v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribePurchasedApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRaceWorkForInnerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	LocalName      *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetLocalName(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRulesByApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSecretKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSystemParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSystemParamsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTrafficControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetApiMethodsRequest struct {
	ApiPath       *string `json:"ApiPath,omitempty" xml:"ApiPath,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s GetApiMethodsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApiMethodsRequest) GoString() string {
	return s.String()
}

func (s *GetApiMethodsRequest) SetApiPath(v string) *GetApiMethodsRequest {
	s.ApiPath = &v
	return s
}

func (s *GetApiMethodsRequest) SetGroupId(v string) *GetApiMethodsRequest {
	s.GroupId = &v
	return s
}

func (s *GetApiMethodsRequest) SetSecurityToken(v string) *GetApiMethodsRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetApiMethodsRequest) SetStageName(v string) *GetApiMethodsRequest {
	s.StageName = &v
	return s
}

type GetApiMethodsResponseBody struct {
	Methods   []*string `json:"Methods,omitempty" xml:"Methods,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApiMethodsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApiMethodsResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiMethodsResponseBody) SetMethods(v []*string) *GetApiMethodsResponseBody {
	s.Methods = v
	return s
}

func (s *GetApiMethodsResponseBody) SetRequestId(v string) *GetApiMethodsResponseBody {
	s.RequestId = &v
	return s
}

type GetApiMethodsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApiMethodsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetApiMethodsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApiMethodsResponse) GoString() string {
	return s.String()
}

func (s *GetApiMethodsResponse) SetHeaders(v map[string]*string) *GetApiMethodsResponse {
	s.Headers = v
	return s
}

func (s *GetApiMethodsResponse) SetStatusCode(v int32) *GetApiMethodsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiMethodsResponse) SetBody(v *GetApiMethodsResponseBody) *GetApiMethodsResponse {
	s.Body = v
	return s
}

type GetCustomizedInfoRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageId       *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s GetCustomizedInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizedInfoRequest) GoString() string {
	return s.String()
}

func (s *GetCustomizedInfoRequest) SetApiId(v string) *GetCustomizedInfoRequest {
	s.ApiId = &v
	return s
}

func (s *GetCustomizedInfoRequest) SetApiName(v string) *GetCustomizedInfoRequest {
	s.ApiName = &v
	return s
}

func (s *GetCustomizedInfoRequest) SetGroupId(v string) *GetCustomizedInfoRequest {
	s.GroupId = &v
	return s
}

func (s *GetCustomizedInfoRequest) SetSecurityToken(v string) *GetCustomizedInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetCustomizedInfoRequest) SetStageId(v string) *GetCustomizedInfoRequest {
	s.StageId = &v
	return s
}

func (s *GetCustomizedInfoRequest) SetStageName(v string) *GetCustomizedInfoRequest {
	s.StageName = &v
	return s
}

type GetCustomizedInfoResponseBody struct {
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkDemos  *GetCustomizedInfoResponseBodySdkDemos `json:"SdkDemos,omitempty" xml:"SdkDemos,omitempty" type:"Struct"`
}

func (s GetCustomizedInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizedInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomizedInfoResponseBody) SetRequestId(v string) *GetCustomizedInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomizedInfoResponseBody) SetSdkDemos(v *GetCustomizedInfoResponseBodySdkDemos) *GetCustomizedInfoResponseBody {
	s.SdkDemos = v
	return s
}

type GetCustomizedInfoResponseBodySdkDemos struct {
	SdkDemo []*GetCustomizedInfoResponseBodySdkDemosSdkDemo `json:"SdkDemo,omitempty" xml:"SdkDemo,omitempty" type:"Repeated"`
}

func (s GetCustomizedInfoResponseBodySdkDemos) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizedInfoResponseBodySdkDemos) GoString() string {
	return s.String()
}

func (s *GetCustomizedInfoResponseBodySdkDemos) SetSdkDemo(v []*GetCustomizedInfoResponseBodySdkDemosSdkDemo) *GetCustomizedInfoResponseBodySdkDemos {
	s.SdkDemo = v
	return s
}

type GetCustomizedInfoResponseBodySdkDemosSdkDemo struct {
	CallDemo *string `json:"CallDemo,omitempty" xml:"CallDemo,omitempty"`
	IdeKey   *string `json:"IdeKey,omitempty" xml:"IdeKey,omitempty"`
}

func (s GetCustomizedInfoResponseBodySdkDemosSdkDemo) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizedInfoResponseBodySdkDemosSdkDemo) GoString() string {
	return s.String()
}

func (s *GetCustomizedInfoResponseBodySdkDemosSdkDemo) SetCallDemo(v string) *GetCustomizedInfoResponseBodySdkDemosSdkDemo {
	s.CallDemo = &v
	return s
}

func (s *GetCustomizedInfoResponseBodySdkDemosSdkDemo) SetIdeKey(v string) *GetCustomizedInfoResponseBodySdkDemosSdkDemo {
	s.IdeKey = &v
	return s
}

type GetCustomizedInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCustomizedInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCustomizedInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCustomizedInfoResponse) GoString() string {
	return s.String()
}

func (s *GetCustomizedInfoResponse) SetHeaders(v map[string]*string) *GetCustomizedInfoResponse {
	s.Headers = v
	return s
}

func (s *GetCustomizedInfoResponse) SetStatusCode(v int32) *GetCustomizedInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomizedInfoResponse) SetBody(v *GetCustomizedInfoResponseBody) *GetCustomizedInfoResponse {
	s.Body = v
	return s
}

type ModifyApiRequest struct {
	AllowSignatureMethod *string `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiId                *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ErrorCodeSamples     *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	FailResultSample     *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OpenIdConnectConfig  *string `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty"`
	RequestConfig        *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParamters     *string `json:"RequestParamters,omitempty" xml:"RequestParamters,omitempty"`
	ResultDescriptions   *string `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty"`
	ResultSample         *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType           *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	Visibility           *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	WebSocketApiType     *string `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s ModifyApiRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiRequest) SetAllowSignatureMethod(v string) *ModifyApiRequest {
	s.AllowSignatureMethod = &v
	return s
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

func (s *ModifyApiRequest) SetDescription(v string) *ModifyApiRequest {
	s.Description = &v
	return s
}

func (s *ModifyApiRequest) SetErrorCodeSamples(v string) *ModifyApiRequest {
	s.ErrorCodeSamples = &v
	return s
}

func (s *ModifyApiRequest) SetFailResultSample(v string) *ModifyApiRequest {
	s.FailResultSample = &v
	return s
}

func (s *ModifyApiRequest) SetGroupId(v string) *ModifyApiRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiRequest) SetOpenIdConnectConfig(v string) *ModifyApiRequest {
	s.OpenIdConnectConfig = &v
	return s
}

func (s *ModifyApiRequest) SetRequestConfig(v string) *ModifyApiRequest {
	s.RequestConfig = &v
	return s
}

func (s *ModifyApiRequest) SetRequestParamters(v string) *ModifyApiRequest {
	s.RequestParamters = &v
	return s
}

func (s *ModifyApiRequest) SetResultDescriptions(v string) *ModifyApiRequest {
	s.ResultDescriptions = &v
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

func (s *ModifyApiRequest) SetServiceConfig(v string) *ModifyApiRequest {
	s.ServiceConfig = &v
	return s
}

func (s *ModifyApiRequest) SetServiceParameters(v string) *ModifyApiRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ModifyApiRequest) SetServiceParametersMap(v string) *ModifyApiRequest {
	s.ServiceParametersMap = &v
	return s
}

func (s *ModifyApiRequest) SetVisibility(v string) *ModifyApiRequest {
	s.Visibility = &v
	return s
}

func (s *ModifyApiRequest) SetWebSocketApiType(v string) *ModifyApiRequest {
	s.WebSocketApiType = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ModifyIpControlRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	IpControlName *string `json:"IpControlName,omitempty" xml:"IpControlName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyIpControlRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpControlRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpControlRequest) SetDescription(v string) *ModifyIpControlRequest {
	s.Description = &v
	return s
}

func (s *ModifyIpControlRequest) SetIpControlId(v string) *ModifyIpControlRequest {
	s.IpControlId = &v
	return s
}

func (s *ModifyIpControlRequest) SetIpControlName(v string) *ModifyIpControlRequest {
	s.IpControlName = &v
	return s
}

func (s *ModifyIpControlRequest) SetSecurityToken(v string) *ModifyIpControlRequest {
	s.SecurityToken = &v
	return s
}

type ModifyIpControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpControlResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpControlResponseBody) SetRequestId(v string) *ModifyIpControlResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpControlResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpControlResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpControlResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpControlResponse) SetHeaders(v map[string]*string) *ModifyIpControlResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpControlResponse) SetStatusCode(v int32) *ModifyIpControlResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpControlResponse) SetBody(v *ModifyIpControlResponseBody) *ModifyIpControlResponse {
	s.Body = v
	return s
}

type ModifyIpControlPolicyItemRequest struct {
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CidrIp        *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	PolicyItemId  *string `json:"PolicyItemId,omitempty" xml:"PolicyItemId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyIpControlPolicyItemRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpControlPolicyItemRequest) GoString() string {
	return s.String()
}

func (s *ModifyIpControlPolicyItemRequest) SetAppId(v string) *ModifyIpControlPolicyItemRequest {
	s.AppId = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetCidrIp(v string) *ModifyIpControlPolicyItemRequest {
	s.CidrIp = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetIpControlId(v string) *ModifyIpControlPolicyItemRequest {
	s.IpControlId = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetPolicyItemId(v string) *ModifyIpControlPolicyItemRequest {
	s.PolicyItemId = &v
	return s
}

func (s *ModifyIpControlPolicyItemRequest) SetSecurityToken(v string) *ModifyIpControlPolicyItemRequest {
	s.SecurityToken = &v
	return s
}

type ModifyIpControlPolicyItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpControlPolicyItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpControlPolicyItemResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpControlPolicyItemResponseBody) SetRequestId(v string) *ModifyIpControlPolicyItemResponseBody {
	s.RequestId = &v
	return s
}

type ModifyIpControlPolicyItemResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyIpControlPolicyItemResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyIpControlPolicyItemResponse) GoString() string {
	return s.String()
}

func (s *ModifyIpControlPolicyItemResponse) SetHeaders(v map[string]*string) *ModifyIpControlPolicyItemResponse {
	s.Headers = v
	return s
}

func (s *ModifyIpControlPolicyItemResponse) SetStatusCode(v int32) *ModifyIpControlPolicyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIpControlPolicyItemResponse) SetBody(v *ModifyIpControlPolicyItemResponseBody) *ModifyIpControlPolicyItemResponse {
	s.Body = v
	return s
}

type ModifyLogConfigRequest struct {
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SlsLogStore   *string `json:"SlsLogStore,omitempty" xml:"SlsLogStore,omitempty"`
	SlsProject    *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
}

func (s ModifyLogConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyLogConfigRequest) SetLogType(v string) *ModifyLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *ModifyLogConfigRequest) SetSecurityToken(v string) *ModifyLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLogConfigRequest) SetSlsLogStore(v string) *ModifyLogConfigRequest {
	s.SlsLogStore = &v
	return s
}

func (s *ModifyLogConfigRequest) SetSlsProject(v string) *ModifyLogConfigRequest {
	s.SlsProject = &v
	return s
}

type ModifyLogConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLogConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLogConfigResponseBody) SetRequestId(v string) *ModifyLogConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLogConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyLogConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLogConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyLogConfigResponse) SetHeaders(v map[string]*string) *ModifyLogConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyLogConfigResponse) SetStatusCode(v int32) *ModifyLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLogConfigResponse) SetBody(v *ModifyLogConfigResponseBody) *ModifyLogConfigResponse {
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ReactivateDomainRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReactivateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s ReactivateDomainRequest) GoString() string {
	return s.String()
}

func (s *ReactivateDomainRequest) SetDomainName(v string) *ReactivateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *ReactivateDomainRequest) SetGroupId(v string) *ReactivateDomainRequest {
	s.GroupId = &v
	return s
}

func (s *ReactivateDomainRequest) SetSecurityToken(v string) *ReactivateDomainRequest {
	s.SecurityToken = &v
	return s
}

type ReactivateDomainResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReactivateDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReactivateDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ReactivateDomainResponseBody) SetRequestId(v string) *ReactivateDomainResponseBody {
	s.RequestId = &v
	return s
}

type ReactivateDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReactivateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReactivateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s ReactivateDomainResponse) GoString() string {
	return s.String()
}

func (s *ReactivateDomainResponse) SetHeaders(v map[string]*string) *ReactivateDomainResponse {
	s.Headers = v
	return s
}

func (s *ReactivateDomainResponse) SetStatusCode(v int32) *ReactivateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *ReactivateDomainResponse) SetBody(v *ReactivateDomainResponseBody) *ReactivateDomainResponse {
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoverApiFromHistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoveryApiDefineFromHistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RecoveryApiFromHistoricalResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAccessPermissionByApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAccessPermissionByAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAllBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveApiRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAppsFromApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveBlackListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RemoveIpControlApisRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveIpControlApisRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveIpControlApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveIpControlApisRequest) SetApiIds(v string) *RemoveIpControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetGroupId(v string) *RemoveIpControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetIpControlId(v string) *RemoveIpControlApisRequest {
	s.IpControlId = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetSecurityToken(v string) *RemoveIpControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveIpControlApisRequest) SetStageName(v string) *RemoveIpControlApisRequest {
	s.StageName = &v
	return s
}

type RemoveIpControlApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIpControlApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveIpControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIpControlApisResponseBody) SetRequestId(v string) *RemoveIpControlApisResponseBody {
	s.RequestId = &v
	return s
}

type RemoveIpControlApisResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveIpControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveIpControlApisResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveIpControlApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveIpControlApisResponse) SetHeaders(v map[string]*string) *RemoveIpControlApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveIpControlApisResponse) SetStatusCode(v int32) *RemoveIpControlApisResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIpControlApisResponse) SetBody(v *RemoveIpControlApisResponseBody) *RemoveIpControlApisResponse {
	s.Body = v
	return s
}

type RemoveIpControlPolicyItemRequest struct {
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	PolicyItemIds *string `json:"PolicyItemIds,omitempty" xml:"PolicyItemIds,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RemoveIpControlPolicyItemRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveIpControlPolicyItemRequest) GoString() string {
	return s.String()
}

func (s *RemoveIpControlPolicyItemRequest) SetIpControlId(v string) *RemoveIpControlPolicyItemRequest {
	s.IpControlId = &v
	return s
}

func (s *RemoveIpControlPolicyItemRequest) SetPolicyItemIds(v string) *RemoveIpControlPolicyItemRequest {
	s.PolicyItemIds = &v
	return s
}

func (s *RemoveIpControlPolicyItemRequest) SetSecurityToken(v string) *RemoveIpControlPolicyItemRequest {
	s.SecurityToken = &v
	return s
}

type RemoveIpControlPolicyItemResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveIpControlPolicyItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveIpControlPolicyItemResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveIpControlPolicyItemResponseBody) SetRequestId(v string) *RemoveIpControlPolicyItemResponseBody {
	s.RequestId = &v
	return s
}

type RemoveIpControlPolicyItemResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveIpControlPolicyItemResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveIpControlPolicyItemResponse) GoString() string {
	return s.String()
}

func (s *RemoveIpControlPolicyItemResponse) SetHeaders(v map[string]*string) *RemoveIpControlPolicyItemResponse {
	s.Headers = v
	return s
}

func (s *RemoveIpControlPolicyItemResponse) SetStatusCode(v int32) *RemoveIpControlPolicyItemResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIpControlPolicyItemResponse) SetBody(v *RemoveIpControlPolicyItemResponseBody) *RemoveIpControlPolicyItemResponse {
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetAppKeySecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ResetCustomizedRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageId       *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s ResetCustomizedRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetCustomizedRequest) GoString() string {
	return s.String()
}

func (s *ResetCustomizedRequest) SetApiId(v string) *ResetCustomizedRequest {
	s.ApiId = &v
	return s
}

func (s *ResetCustomizedRequest) SetApiName(v string) *ResetCustomizedRequest {
	s.ApiName = &v
	return s
}

func (s *ResetCustomizedRequest) SetGroupId(v string) *ResetCustomizedRequest {
	s.GroupId = &v
	return s
}

func (s *ResetCustomizedRequest) SetSecurityToken(v string) *ResetCustomizedRequest {
	s.SecurityToken = &v
	return s
}

func (s *ResetCustomizedRequest) SetStageId(v string) *ResetCustomizedRequest {
	s.StageId = &v
	return s
}

func (s *ResetCustomizedRequest) SetStageName(v string) *ResetCustomizedRequest {
	s.StageName = &v
	return s
}

type ResetCustomizedResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SdkDemos  *ResetCustomizedResponseBodySdkDemos `json:"SdkDemos,omitempty" xml:"SdkDemos,omitempty" type:"Struct"`
}

func (s ResetCustomizedResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetCustomizedResponseBody) GoString() string {
	return s.String()
}

func (s *ResetCustomizedResponseBody) SetRequestId(v string) *ResetCustomizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetCustomizedResponseBody) SetSdkDemos(v *ResetCustomizedResponseBodySdkDemos) *ResetCustomizedResponseBody {
	s.SdkDemos = v
	return s
}

type ResetCustomizedResponseBodySdkDemos struct {
	SdkDemo []*ResetCustomizedResponseBodySdkDemosSdkDemo `json:"SdkDemo,omitempty" xml:"SdkDemo,omitempty" type:"Repeated"`
}

func (s ResetCustomizedResponseBodySdkDemos) String() string {
	return tea.Prettify(s)
}

func (s ResetCustomizedResponseBodySdkDemos) GoString() string {
	return s.String()
}

func (s *ResetCustomizedResponseBodySdkDemos) SetSdkDemo(v []*ResetCustomizedResponseBodySdkDemosSdkDemo) *ResetCustomizedResponseBodySdkDemos {
	s.SdkDemo = v
	return s
}

type ResetCustomizedResponseBodySdkDemosSdkDemo struct {
	CallDemo *string `json:"CallDemo,omitempty" xml:"CallDemo,omitempty"`
	IdeKey   *string `json:"IdeKey,omitempty" xml:"IdeKey,omitempty"`
}

func (s ResetCustomizedResponseBodySdkDemosSdkDemo) String() string {
	return tea.Prettify(s)
}

func (s ResetCustomizedResponseBodySdkDemosSdkDemo) GoString() string {
	return s.String()
}

func (s *ResetCustomizedResponseBodySdkDemosSdkDemo) SetCallDemo(v string) *ResetCustomizedResponseBodySdkDemosSdkDemo {
	s.CallDemo = &v
	return s
}

func (s *ResetCustomizedResponseBodySdkDemosSdkDemo) SetIdeKey(v string) *ResetCustomizedResponseBodySdkDemosSdkDemo {
	s.IdeKey = &v
	return s
}

type ResetCustomizedResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetCustomizedResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetCustomizedResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetCustomizedResponse) GoString() string {
	return s.String()
}

func (s *ResetCustomizedResponse) SetHeaders(v map[string]*string) *ResetCustomizedResponse {
	s.Headers = v
	return s
}

func (s *ResetCustomizedResponse) SetStatusCode(v int32) *ResetCustomizedResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetCustomizedResponse) SetBody(v *ResetCustomizedResponseBody) *ResetCustomizedResponse {
	s.Body = v
	return s
}

type SdkGenerateRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SdkGenerateRequest) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateRequest) GoString() string {
	return s.String()
}

func (s *SdkGenerateRequest) SetAppId(v int64) *SdkGenerateRequest {
	s.AppId = &v
	return s
}

func (s *SdkGenerateRequest) SetGroupId(v string) *SdkGenerateRequest {
	s.GroupId = &v
	return s
}

func (s *SdkGenerateRequest) SetLanguage(v string) *SdkGenerateRequest {
	s.Language = &v
	return s
}

func (s *SdkGenerateRequest) SetSecurityToken(v string) *SdkGenerateRequest {
	s.SecurityToken = &v
	return s
}

type SdkGenerateResponseBody struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdkGenerateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGenerateResponseBody) SetDownloadLink(v string) *SdkGenerateResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *SdkGenerateResponseBody) SetRequestId(v string) *SdkGenerateResponseBody {
	s.RequestId = &v
	return s
}

type SdkGenerateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SdkGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SdkGenerateResponse) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateResponse) GoString() string {
	return s.String()
}

func (s *SdkGenerateResponse) SetHeaders(v map[string]*string) *SdkGenerateResponse {
	s.Headers = v
	return s
}

func (s *SdkGenerateResponse) SetStatusCode(v int32) *SdkGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGenerateResponse) SetBody(v *SdkGenerateResponseBody) *SdkGenerateResponse {
	s.Body = v
	return s
}

type SdkGenerateByAppRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SdkGenerateByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateByAppRequest) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppRequest) SetAppId(v int64) *SdkGenerateByAppRequest {
	s.AppId = &v
	return s
}

func (s *SdkGenerateByAppRequest) SetLanguage(v string) *SdkGenerateByAppRequest {
	s.Language = &v
	return s
}

func (s *SdkGenerateByAppRequest) SetSecurityToken(v string) *SdkGenerateByAppRequest {
	s.SecurityToken = &v
	return s
}

type SdkGenerateByAppResponseBody struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdkGenerateByAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateByAppResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppResponseBody) SetDownloadLink(v string) *SdkGenerateByAppResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *SdkGenerateByAppResponseBody) SetRequestId(v string) *SdkGenerateByAppResponseBody {
	s.RequestId = &v
	return s
}

type SdkGenerateByAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SdkGenerateByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SdkGenerateByAppResponse) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateByAppResponse) GoString() string {
	return s.String()
}

func (s *SdkGenerateByAppResponse) SetHeaders(v map[string]*string) *SdkGenerateByAppResponse {
	s.Headers = v
	return s
}

func (s *SdkGenerateByAppResponse) SetStatusCode(v int32) *SdkGenerateByAppResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGenerateByAppResponse) SetBody(v *SdkGenerateByAppResponseBody) *SdkGenerateByAppResponse {
	s.Body = v
	return s
}

type SdkGenerateByGroupRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SdkGenerateByGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateByGroupRequest) GoString() string {
	return s.String()
}

func (s *SdkGenerateByGroupRequest) SetGroupId(v string) *SdkGenerateByGroupRequest {
	s.GroupId = &v
	return s
}

func (s *SdkGenerateByGroupRequest) SetLanguage(v string) *SdkGenerateByGroupRequest {
	s.Language = &v
	return s
}

func (s *SdkGenerateByGroupRequest) SetSecurityToken(v string) *SdkGenerateByGroupRequest {
	s.SecurityToken = &v
	return s
}

type SdkGenerateByGroupResponseBody struct {
	DownloadLink *string `json:"DownloadLink,omitempty" xml:"DownloadLink,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SdkGenerateByGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateByGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SdkGenerateByGroupResponseBody) SetDownloadLink(v string) *SdkGenerateByGroupResponseBody {
	s.DownloadLink = &v
	return s
}

func (s *SdkGenerateByGroupResponseBody) SetRequestId(v string) *SdkGenerateByGroupResponseBody {
	s.RequestId = &v
	return s
}

type SdkGenerateByGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SdkGenerateByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SdkGenerateByGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SdkGenerateByGroupResponse) GoString() string {
	return s.String()
}

func (s *SdkGenerateByGroupResponse) SetHeaders(v map[string]*string) *SdkGenerateByGroupResponse {
	s.Headers = v
	return s
}

func (s *SdkGenerateByGroupResponse) SetStatusCode(v int32) *SdkGenerateByGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SdkGenerateByGroupResponse) SetBody(v *SdkGenerateByGroupResponseBody) *SdkGenerateByGroupResponse {
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAccessPermissionByApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetAccessPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetApiRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DomainBindingStatus   *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	DomainLegalStatus     *string `json:"DomainLegalStatus,omitempty" xml:"DomainLegalStatus,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainRemark          *string `json:"DomainRemark,omitempty" xml:"DomainRemark,omitempty"`
	DomainWebSocketStatus *string `json:"DomainWebSocketStatus,omitempty" xml:"DomainWebSocketStatus,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain             *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
}

func (s SetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainResponseBody) SetDomainBindingStatus(v string) *SetDomainResponseBody {
	s.DomainBindingStatus = &v
	return s
}

func (s *SetDomainResponseBody) SetDomainLegalStatus(v string) *SetDomainResponseBody {
	s.DomainLegalStatus = &v
	return s
}

func (s *SetDomainResponseBody) SetDomainName(v string) *SetDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *SetDomainResponseBody) SetDomainRemark(v string) *SetDomainResponseBody {
	s.DomainRemark = &v
	return s
}

func (s *SetDomainResponseBody) SetDomainWebSocketStatus(v string) *SetDomainResponseBody {
	s.DomainWebSocketStatus = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SetDomainWebSocketStatusRequest struct {
	ActionValue   *string `json:"ActionValue,omitempty" xml:"ActionValue,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDomainWebSocketStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainWebSocketStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDomainWebSocketStatusRequest) SetActionValue(v string) *SetDomainWebSocketStatusRequest {
	s.ActionValue = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetDomainName(v string) *SetDomainWebSocketStatusRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetGroupId(v string) *SetDomainWebSocketStatusRequest {
	s.GroupId = &v
	return s
}

func (s *SetDomainWebSocketStatusRequest) SetSecurityToken(v string) *SetDomainWebSocketStatusRequest {
	s.SecurityToken = &v
	return s
}

type SetDomainWebSocketStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainWebSocketStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDomainWebSocketStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainWebSocketStatusResponseBody) SetRequestId(v string) *SetDomainWebSocketStatusResponseBody {
	s.RequestId = &v
	return s
}

type SetDomainWebSocketStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDomainWebSocketStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDomainWebSocketStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDomainWebSocketStatusResponse) GoString() string {
	return s.String()
}

func (s *SetDomainWebSocketStatusResponse) SetHeaders(v map[string]*string) *SetDomainWebSocketStatusResponse {
	s.Headers = v
	return s
}

func (s *SetDomainWebSocketStatusResponse) SetStatusCode(v int32) *SetDomainWebSocketStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainWebSocketStatusResponse) SetBody(v *SetDomainWebSocketStatusResponseBody) *SetDomainWebSocketStatusResponse {
	s.Body = v
	return s
}

type SetIpControlApisRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IpControlId   *string `json:"IpControlId,omitempty" xml:"IpControlId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetIpControlApisRequest) String() string {
	return tea.Prettify(s)
}

func (s SetIpControlApisRequest) GoString() string {
	return s.String()
}

func (s *SetIpControlApisRequest) SetApiIds(v string) *SetIpControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetIpControlApisRequest) SetGroupId(v string) *SetIpControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetIpControlApisRequest) SetIpControlId(v string) *SetIpControlApisRequest {
	s.IpControlId = &v
	return s
}

func (s *SetIpControlApisRequest) SetSecurityToken(v string) *SetIpControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetIpControlApisRequest) SetStageName(v string) *SetIpControlApisRequest {
	s.StageName = &v
	return s
}

type SetIpControlApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetIpControlApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetIpControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetIpControlApisResponseBody) SetRequestId(v string) *SetIpControlApisResponseBody {
	s.RequestId = &v
	return s
}

type SetIpControlApisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetIpControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetIpControlApisResponse) String() string {
	return tea.Prettify(s)
}

func (s SetIpControlApisResponse) GoString() string {
	return s.String()
}

func (s *SetIpControlApisResponse) SetHeaders(v map[string]*string) *SetIpControlApisResponse {
	s.Headers = v
	return s
}

func (s *SetIpControlApisResponse) SetStatusCode(v int32) *SetIpControlApisResponse {
	s.StatusCode = &v
	return s
}

func (s *SetIpControlApisResponse) SetBody(v *SetIpControlApisResponseBody) *SetIpControlApisResponse {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetMockIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type VpcDescribeAccessesRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort  *string `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s VpcDescribeAccessesRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcDescribeAccessesRequest) GoString() string {
	return s.String()
}

func (s *VpcDescribeAccessesRequest) SetInstanceId(v string) *VpcDescribeAccessesRequest {
	s.InstanceId = &v
	return s
}

func (s *VpcDescribeAccessesRequest) SetInstancePort(v string) *VpcDescribeAccessesRequest {
	s.InstancePort = &v
	return s
}

func (s *VpcDescribeAccessesRequest) SetName(v string) *VpcDescribeAccessesRequest {
	s.Name = &v
	return s
}

func (s *VpcDescribeAccessesRequest) SetPageNumber(v int32) *VpcDescribeAccessesRequest {
	s.PageNumber = &v
	return s
}

func (s *VpcDescribeAccessesRequest) SetPageSize(v int32) *VpcDescribeAccessesRequest {
	s.PageSize = &v
	return s
}

func (s *VpcDescribeAccessesRequest) SetSecurityToken(v string) *VpcDescribeAccessesRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcDescribeAccessesRequest) SetVpcId(v string) *VpcDescribeAccessesRequest {
	s.VpcId = &v
	return s
}

type VpcDescribeAccessesResponseBody struct {
	PageNumber          *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcAccessAttributes *VpcDescribeAccessesResponseBodyVpcAccessAttributes `json:"VpcAccessAttributes,omitempty" xml:"VpcAccessAttributes,omitempty" type:"Struct"`
}

func (s VpcDescribeAccessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcDescribeAccessesResponseBody) GoString() string {
	return s.String()
}

func (s *VpcDescribeAccessesResponseBody) SetPageNumber(v int32) *VpcDescribeAccessesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *VpcDescribeAccessesResponseBody) SetPageSize(v int32) *VpcDescribeAccessesResponseBody {
	s.PageSize = &v
	return s
}

func (s *VpcDescribeAccessesResponseBody) SetRequestId(v string) *VpcDescribeAccessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *VpcDescribeAccessesResponseBody) SetTotalCount(v int32) *VpcDescribeAccessesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *VpcDescribeAccessesResponseBody) SetVpcAccessAttributes(v *VpcDescribeAccessesResponseBodyVpcAccessAttributes) *VpcDescribeAccessesResponseBody {
	s.VpcAccessAttributes = v
	return s
}

type VpcDescribeAccessesResponseBodyVpcAccessAttributes struct {
	VpcAccessAttribute []*VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute `json:"VpcAccessAttribute,omitempty" xml:"VpcAccessAttribute,omitempty" type:"Repeated"`
}

func (s VpcDescribeAccessesResponseBodyVpcAccessAttributes) String() string {
	return tea.Prettify(s)
}

func (s VpcDescribeAccessesResponseBodyVpcAccessAttributes) GoString() string {
	return s.String()
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributes) SetVpcAccessAttribute(v []*VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) *VpcDescribeAccessesResponseBodyVpcAccessAttributes {
	s.VpcAccessAttribute = v
	return s
}

type VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) String() string {
	return tea.Prettify(s)
}

func (s VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GoString() string {
	return s.String()
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetCreatedTime(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.CreatedTime = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetId(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Id = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetInstanceId(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.InstanceId = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetModifiedTime(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetName(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Name = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetPort(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Port = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetRegionId(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.RegionId = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetStatus(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Status = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetUserId(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.UserId = &v
	return s
}

func (s *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetVpcId(v string) *VpcDescribeAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.VpcId = &v
	return s
}

type VpcDescribeAccessesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VpcDescribeAccessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VpcDescribeAccessesResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcDescribeAccessesResponse) GoString() string {
	return s.String()
}

func (s *VpcDescribeAccessesResponse) SetHeaders(v map[string]*string) *VpcDescribeAccessesResponse {
	s.Headers = v
	return s
}

func (s *VpcDescribeAccessesResponse) SetStatusCode(v int32) *VpcDescribeAccessesResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcDescribeAccessesResponse) SetBody(v *VpcDescribeAccessesResponseBody) *VpcDescribeAccessesResponse {
	s.Body = v
	return s
}

type VpcGrantAccessRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort  *string `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s VpcGrantAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcGrantAccessRequest) GoString() string {
	return s.String()
}

func (s *VpcGrantAccessRequest) SetInstanceId(v string) *VpcGrantAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *VpcGrantAccessRequest) SetInstancePort(v string) *VpcGrantAccessRequest {
	s.InstancePort = &v
	return s
}

func (s *VpcGrantAccessRequest) SetName(v string) *VpcGrantAccessRequest {
	s.Name = &v
	return s
}

func (s *VpcGrantAccessRequest) SetSecurityToken(v string) *VpcGrantAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcGrantAccessRequest) SetToken(v string) *VpcGrantAccessRequest {
	s.Token = &v
	return s
}

func (s *VpcGrantAccessRequest) SetVpcId(v string) *VpcGrantAccessRequest {
	s.VpcId = &v
	return s
}

type VpcGrantAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcGrantAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcGrantAccessResponseBody) GoString() string {
	return s.String()
}

func (s *VpcGrantAccessResponseBody) SetRequestId(v string) *VpcGrantAccessResponseBody {
	s.RequestId = &v
	return s
}

type VpcGrantAccessResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VpcGrantAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VpcGrantAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcGrantAccessResponse) GoString() string {
	return s.String()
}

func (s *VpcGrantAccessResponse) SetHeaders(v map[string]*string) *VpcGrantAccessResponse {
	s.Headers = v
	return s
}

func (s *VpcGrantAccessResponse) SetStatusCode(v int32) *VpcGrantAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcGrantAccessResponse) SetBody(v *VpcGrantAccessResponseBody) *VpcGrantAccessResponse {
	s.Body = v
	return s
}

type VpcModifyAccessRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort  *string `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s VpcModifyAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcModifyAccessRequest) GoString() string {
	return s.String()
}

func (s *VpcModifyAccessRequest) SetInstanceId(v string) *VpcModifyAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *VpcModifyAccessRequest) SetInstancePort(v string) *VpcModifyAccessRequest {
	s.InstancePort = &v
	return s
}

func (s *VpcModifyAccessRequest) SetSecurityToken(v string) *VpcModifyAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcModifyAccessRequest) SetToken(v string) *VpcModifyAccessRequest {
	s.Token = &v
	return s
}

func (s *VpcModifyAccessRequest) SetVpcId(v string) *VpcModifyAccessRequest {
	s.VpcId = &v
	return s
}

type VpcModifyAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcModifyAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcModifyAccessResponseBody) GoString() string {
	return s.String()
}

func (s *VpcModifyAccessResponseBody) SetRequestId(v string) *VpcModifyAccessResponseBody {
	s.RequestId = &v
	return s
}

type VpcModifyAccessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VpcModifyAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VpcModifyAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcModifyAccessResponse) GoString() string {
	return s.String()
}

func (s *VpcModifyAccessResponse) SetHeaders(v map[string]*string) *VpcModifyAccessResponse {
	s.Headers = v
	return s
}

func (s *VpcModifyAccessResponse) SetStatusCode(v int32) *VpcModifyAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcModifyAccessResponse) SetBody(v *VpcModifyAccessResponseBody) *VpcModifyAccessResponse {
	s.Body = v
	return s
}

type VpcRevokeAccessRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort  *string `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Token         *string `json:"Token,omitempty" xml:"Token,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s VpcRevokeAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s VpcRevokeAccessRequest) GoString() string {
	return s.String()
}

func (s *VpcRevokeAccessRequest) SetInstanceId(v string) *VpcRevokeAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *VpcRevokeAccessRequest) SetInstancePort(v string) *VpcRevokeAccessRequest {
	s.InstancePort = &v
	return s
}

func (s *VpcRevokeAccessRequest) SetSecurityToken(v string) *VpcRevokeAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *VpcRevokeAccessRequest) SetToken(v string) *VpcRevokeAccessRequest {
	s.Token = &v
	return s
}

func (s *VpcRevokeAccessRequest) SetVpcId(v string) *VpcRevokeAccessRequest {
	s.VpcId = &v
	return s
}

type VpcRevokeAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VpcRevokeAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VpcRevokeAccessResponseBody) GoString() string {
	return s.String()
}

func (s *VpcRevokeAccessResponseBody) SetRequestId(v string) *VpcRevokeAccessResponseBody {
	s.RequestId = &v
	return s
}

type VpcRevokeAccessResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *VpcRevokeAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s VpcRevokeAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s VpcRevokeAccessResponse) GoString() string {
	return s.String()
}

func (s *VpcRevokeAccessResponse) SetHeaders(v map[string]*string) *VpcRevokeAccessResponse {
	s.Headers = v
	return s
}

func (s *VpcRevokeAccessResponse) SetStatusCode(v int32) *VpcRevokeAccessResponse {
	s.StatusCode = &v
	return s
}

func (s *VpcRevokeAccessResponse) SetBody(v *VpcRevokeAccessResponseBody) *VpcRevokeAccessResponse {
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
		"cn-qingdao":            tea.String("apigateway.cn-qingdao.aliyuncs.com"),
		"cn-beijing":            tea.String("apigateway.cn-beijing.aliyuncs.com"),
		"cn-chengdu":            tea.String("apigateway.cn-chengdu.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("apigateway.cn-zhangjiakou.aliyuncs.com"),
		"cn-huhehaote":          tea.String("apigateway.cn-huhehaote.aliyuncs.com"),
		"cn-hangzhou":           tea.String("apigateway.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("apigateway.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("apigateway.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":           tea.String("apigateway.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":        tea.String("apigateway.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("apigateway.ap-southeast-2.aliyuncs.com"),
		"ap-southeast-3":        tea.String("apigateway.ap-southeast-3.aliyuncs.com"),
		"ap-southeast-5":        tea.String("apigateway.ap-southeast-5.aliyuncs.com"),
		"ap-northeast-1":        tea.String("apigateway.ap-northeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("apigateway.eu-west-1.aliyuncs.com"),
		"us-west-1":             tea.String("apigateway.us-west-1.aliyuncs.com"),
		"us-east-1":             tea.String("apigateway.us-east-1.aliyuncs.com"),
		"eu-central-1":          tea.String("apigateway.eu-central-1.aliyuncs.com"),
		"me-east-1":             tea.String("apigateway.me-east-1.aliyuncs.com"),
		"ap-south-1":            tea.String("apigateway.ap-south-1.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("apigateway.cn-north-2-gov-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("apigateway.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("apigateway.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("apigateway.aliyuncs.com"),
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
		Version:     tea.String("2018-06-01"),
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddBlackList"),
		Version:     tea.String("2018-06-01"),
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

func (client *Client) AddIpControlPolicyItemWithOptions(request *AddIpControlPolicyItemRequest, runtime *util.RuntimeOptions) (_result *AddIpControlPolicyItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CidrIp)) {
		query["CidrIp"] = request.CidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIpControlPolicyItem"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddIpControlPolicyItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddIpControlPolicyItem(request *AddIpControlPolicyItemRequest) (_result *AddIpControlPolicyItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddIpControlPolicyItemResponse{}
	_body, _err := client.AddIpControlPolicyItemWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) CreateApiWithOptions(request *CreateApiRequest, runtime *util.RuntimeOptions) (_result *CreateApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowSignatureMethod)) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
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

	if !tea.BoolValue(util.IsUnset(request.ErrorCodeSamples)) {
		query["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !tea.BoolValue(util.IsUnset(request.FailResultSample)) {
		query["FailResultSample"] = request.FailResultSample
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenIdConnectConfig)) {
		query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestConfig)) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestParamters)) {
		query["RequestParamters"] = request.RequestParamters
	}

	if !tea.BoolValue(util.IsUnset(request.ResultDescriptions)) {
		query["ResultDescriptions"] = request.ResultDescriptions
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

	if !tea.BoolValue(util.IsUnset(request.WebSocketApiType)) {
		query["WebSocketApiType"] = request.WebSocketApiType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApi"),
		Version:     tea.String("2018-06-01"),
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

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiGroup"),
		Version:     tea.String("2018-06-01"),
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

func (client *Client) CreateApiStageVariableWithOptions(request *CreateApiStageVariableRequest, runtime *util.RuntimeOptions) (_result *CreateApiStageVariableResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	if !tea.BoolValue(util.IsUnset(request.StageRouteModel)) {
		query["StageRouteModel"] = request.StageRouteModel
	}

	if !tea.BoolValue(util.IsUnset(request.SupportRoute)) {
		query["SupportRoute"] = request.SupportRoute
	}

	if !tea.BoolValue(util.IsUnset(request.VariableName)) {
		query["VariableName"] = request.VariableName
	}

	if !tea.BoolValue(util.IsUnset(request.VariableValue)) {
		query["VariableValue"] = request.VariableValue
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiStageVariable"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApiStageVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApiStageVariable(request *CreateApiStageVariableRequest) (_result *CreateApiStageVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApiStageVariableResponse{}
	_body, _err := client.CreateApiStageVariableWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) CreateCustomizedInfoWithOptions(request *CreateCustomizedInfoRequest, runtime *util.RuntimeOptions) (_result *CreateCustomizedInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.CsharpDemo)) {
		query["CsharpDemo"] = request.CsharpDemo
	}

	if !tea.BoolValue(util.IsUnset(request.CurlDemo)) {
		query["CurlDemo"] = request.CurlDemo
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.JavaDemo)) {
		query["JavaDemo"] = request.JavaDemo
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectcDemo)) {
		query["ObjectcDemo"] = request.ObjectcDemo
	}

	if !tea.BoolValue(util.IsUnset(request.PhpDemo)) {
		query["PhpDemo"] = request.PhpDemo
	}

	if !tea.BoolValue(util.IsUnset(request.PythonDemo)) {
		query["PythonDemo"] = request.PythonDemo
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomizedInfo"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomizedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCustomizedInfo(request *CreateCustomizedInfoRequest) (_result *CreateCustomizedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCustomizedInfoResponse{}
	_body, _err := client.CreateCustomizedInfoWithOptions(request, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.ExpiredOn)) {
		query["ExpiredOn"] = request.ExpiredOn
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) CreateIpControlWithOptions(request *CreateIpControlRequest, runtime *util.RuntimeOptions) (_result *CreateIpControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlName)) {
		query["IpControlName"] = request.IpControlName
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlPolicys)) {
		query["IpControlPolicys"] = request.IpControlPolicys
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlType)) {
		query["IpControlType"] = request.IpControlType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpControl"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIpControl(request *CreateIpControlRequest) (_result *CreateIpControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIpControlResponse{}
	_body, _err := client.CreateIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLogConfigWithOptions(request *CreateLogConfigRequest, runtime *util.RuntimeOptions) (_result *CreateLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStore)) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogConfig"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLogConfig(request *CreateLogConfigRequest) (_result *CreateLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLogConfigResponse{}
	_body, _err := client.CreateLogConfigWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DeleteApiStageVariableWithOptions(request *DeleteApiStageVariableRequest, runtime *util.RuntimeOptions) (_result *DeleteApiStageVariableResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	if !tea.BoolValue(util.IsUnset(request.VariableName)) {
		query["VariableName"] = request.VariableName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApiStageVariable"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApiStageVariableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApiStageVariable(request *DeleteApiStageVariableRequest) (_result *DeleteApiStageVariableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApiStageVariableResponse{}
	_body, _err := client.DeleteApiStageVariableWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DeleteIpControlWithOptions(request *DeleteIpControlRequest, runtime *util.RuntimeOptions) (_result *DeleteIpControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpControl"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteIpControl(request *DeleteIpControlRequest) (_result *DeleteIpControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIpControlResponse{}
	_body, _err := client.DeleteIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLogConfigWithOptions(request *DeleteLogConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogConfig"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLogConfig(request *DeleteLogConfigRequest) (_result *DeleteLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLogConfigResponse{}
	_body, _err := client.DeleteLogConfigWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

	if !tea.BoolValue(util.IsUnset(request.SupportMock)) {
		query["SupportMock"] = request.SupportMock
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApi"),
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DescribeApiGroupDetailForConsumerWithOptions(request *DescribeApiGroupDetailForConsumerRequest, runtime *util.RuntimeOptions) (_result *DescribeApiGroupDetailForConsumerResponse, _err error) {
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
		Action:      tea.String("DescribeApiGroupDetailForConsumer"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiGroupDetailForConsumerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiGroupDetailForConsumer(request *DescribeApiGroupDetailForConsumerRequest) (_result *DescribeApiGroupDetailForConsumerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiGroupDetailForConsumerResponse{}
	_body, _err := client.DescribeApiGroupDetailForConsumerWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DescribeApiIpControlsWithOptions(request *DescribeApiIpControlsRequest, runtime *util.RuntimeOptions) (_result *DescribeApiIpControlsResponse, _err error) {
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
		Action:      tea.String("DescribeApiIpControls"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiIpControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiIpControls(request *DescribeApiIpControlsRequest) (_result *DescribeApiIpControlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiIpControlsResponse{}
	_body, _err := client.DescribeApiIpControlsWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DescribeApiStageDetailWithOptions(request *DescribeApiStageDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeApiStageDetailResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiStageDetail"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiStageDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiStageDetail(request *DescribeApiStageDetailRequest) (_result *DescribeApiStageDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiStageDetailResponse{}
	_body, _err := client.DescribeApiStageDetailWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DescribeApisByIpControlWithOptions(request *DescribeApisByIpControlRequest, runtime *util.RuntimeOptions) (_result *DescribeApisByIpControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
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
		Action:      tea.String("DescribeApisByIpControl"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisByIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApisByIpControl(request *DescribeApisByIpControlRequest) (_result *DescribeApisByIpControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisByIpControlResponse{}
	_body, _err := client.DescribeApisByIpControlWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppSecurity"),
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

	if !tea.BoolValue(util.IsUnset(request.AppOwnerId)) {
		query["AppOwnerId"] = request.AppOwnerId
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) DescribeIpControlPolicyItemsWithOptions(request *DescribeIpControlPolicyItemsRequest, runtime *util.RuntimeOptions) (_result *DescribeIpControlPolicyItemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyItemId)) {
		query["PolicyItemId"] = request.PolicyItemId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpControlPolicyItems"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpControlPolicyItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpControlPolicyItems(request *DescribeIpControlPolicyItemsRequest) (_result *DescribeIpControlPolicyItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpControlPolicyItemsResponse{}
	_body, _err := client.DescribeIpControlPolicyItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeIpControlsWithOptions(request *DescribeIpControlsRequest, runtime *util.RuntimeOptions) (_result *DescribeIpControlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlName)) {
		query["IpControlName"] = request.IpControlName
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlType)) {
		query["IpControlType"] = request.IpControlType
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
		Action:      tea.String("DescribeIpControls"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeIpControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeIpControls(request *DescribeIpControlsRequest) (_result *DescribeIpControlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeIpControlsResponse{}
	_body, _err := client.DescribeIpControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeLogConfigWithOptions(request *DescribeLogConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogConfig"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeLogConfig(request *DescribeLogConfigRequest) (_result *DescribeLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeLogConfigResponse{}
	_body, _err := client.DescribeLogConfigWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApis"),
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) GetApiMethodsWithOptions(request *GetApiMethodsRequest, runtime *util.RuntimeOptions) (_result *GetApiMethodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApiPath)) {
		query["ApiPath"] = request.ApiPath
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
		Action:      tea.String("GetApiMethods"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApiMethodsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApiMethods(request *GetApiMethodsRequest) (_result *GetApiMethodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApiMethodsResponse{}
	_body, _err := client.GetApiMethodsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCustomizedInfoWithOptions(request *GetCustomizedInfoRequest, runtime *util.RuntimeOptions) (_result *GetCustomizedInfoResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCustomizedInfo"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCustomizedInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCustomizedInfo(request *GetCustomizedInfoRequest) (_result *GetCustomizedInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCustomizedInfoResponse{}
	_body, _err := client.GetCustomizedInfoWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AllowSignatureMethod)) {
		query["AllowSignatureMethod"] = request.AllowSignatureMethod
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

	if !tea.BoolValue(util.IsUnset(request.ErrorCodeSamples)) {
		query["ErrorCodeSamples"] = request.ErrorCodeSamples
	}

	if !tea.BoolValue(util.IsUnset(request.FailResultSample)) {
		query["FailResultSample"] = request.FailResultSample
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.OpenIdConnectConfig)) {
		query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestConfig)) {
		query["RequestConfig"] = request.RequestConfig
	}

	if !tea.BoolValue(util.IsUnset(request.RequestParamters)) {
		query["RequestParamters"] = request.RequestParamters
	}

	if !tea.BoolValue(util.IsUnset(request.ResultDescriptions)) {
		query["ResultDescriptions"] = request.ResultDescriptions
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

	if !tea.BoolValue(util.IsUnset(request.WebSocketApiType)) {
		query["WebSocketApiType"] = request.WebSocketApiType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApi"),
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) ModifyIpControlWithOptions(request *ModifyIpControlRequest, runtime *util.RuntimeOptions) (_result *ModifyIpControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlName)) {
		query["IpControlName"] = request.IpControlName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIpControl"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIpControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpControl(request *ModifyIpControlRequest) (_result *ModifyIpControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpControlResponse{}
	_body, _err := client.ModifyIpControlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyIpControlPolicyItemWithOptions(request *ModifyIpControlPolicyItemRequest, runtime *util.RuntimeOptions) (_result *ModifyIpControlPolicyItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CidrIp)) {
		query["CidrIp"] = request.CidrIp
	}

	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyItemId)) {
		query["PolicyItemId"] = request.PolicyItemId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIpControlPolicyItem"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyIpControlPolicyItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyIpControlPolicyItem(request *ModifyIpControlPolicyItemRequest) (_result *ModifyIpControlPolicyItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyIpControlPolicyItemResponse{}
	_body, _err := client.ModifyIpControlPolicyItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyLogConfigWithOptions(request *ModifyLogConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyLogConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SlsLogStore)) {
		query["SlsLogStore"] = request.SlsLogStore
	}

	if !tea.BoolValue(util.IsUnset(request.SlsProject)) {
		query["SlsProject"] = request.SlsProject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLogConfig"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyLogConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyLogConfig(request *ModifyLogConfigRequest) (_result *ModifyLogConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLogConfigResponse{}
	_body, _err := client.ModifyLogConfigWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) ReactivateDomainWithOptions(request *ReactivateDomainRequest, runtime *util.RuntimeOptions) (_result *ReactivateDomainResponse, _err error) {
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
		Action:      tea.String("ReactivateDomain"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReactivateDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReactivateDomain(request *ReactivateDomainRequest) (_result *ReactivateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReactivateDomainResponse{}
	_body, _err := client.ReactivateDomainWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) RemoveIpControlApisWithOptions(request *RemoveIpControlApisRequest, runtime *util.RuntimeOptions) (_result *RemoveIpControlApisResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
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
		Action:      tea.String("RemoveIpControlApis"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveIpControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveIpControlApis(request *RemoveIpControlApisRequest) (_result *RemoveIpControlApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveIpControlApisResponse{}
	_body, _err := client.RemoveIpControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveIpControlPolicyItemWithOptions(request *RemoveIpControlPolicyItemRequest, runtime *util.RuntimeOptions) (_result *RemoveIpControlPolicyItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyItemIds)) {
		query["PolicyItemIds"] = request.PolicyItemIds
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveIpControlPolicyItem"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveIpControlPolicyItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveIpControlPolicyItem(request *RemoveIpControlPolicyItemRequest) (_result *RemoveIpControlPolicyItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveIpControlPolicyItemResponse{}
	_body, _err := client.RemoveIpControlPolicyItemWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) ResetCustomizedWithOptions(request *ResetCustomizedRequest, runtime *util.RuntimeOptions) (_result *ResetCustomizedResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StageId)) {
		query["StageId"] = request.StageId
	}

	if !tea.BoolValue(util.IsUnset(request.StageName)) {
		query["StageName"] = request.StageName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetCustomized"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetCustomizedResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetCustomized(request *ResetCustomizedRequest) (_result *ResetCustomizedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetCustomizedResponse{}
	_body, _err := client.ResetCustomizedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SdkGenerateWithOptions(request *SdkGenerateRequest, runtime *util.RuntimeOptions) (_result *SdkGenerateResponse, _err error) {
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
		Action:      tea.String("SdkGenerate"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SdkGenerateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SdkGenerate(request *SdkGenerateRequest) (_result *SdkGenerateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SdkGenerateResponse{}
	_body, _err := client.SdkGenerateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SdkGenerateByAppWithOptions(request *SdkGenerateByAppRequest, runtime *util.RuntimeOptions) (_result *SdkGenerateByAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

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
		Action:      tea.String("SdkGenerateByApp"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SdkGenerateByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SdkGenerateByApp(request *SdkGenerateByAppRequest) (_result *SdkGenerateByAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SdkGenerateByAppResponse{}
	_body, _err := client.SdkGenerateByAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SdkGenerateByGroupWithOptions(request *SdkGenerateByGroupRequest, runtime *util.RuntimeOptions) (_result *SdkGenerateByGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

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
		Action:      tea.String("SdkGenerateByGroup"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SdkGenerateByGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SdkGenerateByGroup(request *SdkGenerateByGroupRequest) (_result *SdkGenerateByGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SdkGenerateByGroupResponse{}
	_body, _err := client.SdkGenerateByGroupWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) SetDomainWebSocketStatusWithOptions(request *SetDomainWebSocketStatusRequest, runtime *util.RuntimeOptions) (_result *SetDomainWebSocketStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionValue)) {
		query["ActionValue"] = request.ActionValue
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
		Action:      tea.String("SetDomainWebSocketStatus"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDomainWebSocketStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDomainWebSocketStatus(request *SetDomainWebSocketStatusRequest) (_result *SetDomainWebSocketStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDomainWebSocketStatusResponse{}
	_body, _err := client.SetDomainWebSocketStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetIpControlApisWithOptions(request *SetIpControlApisRequest, runtime *util.RuntimeOptions) (_result *SetIpControlApisResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.IpControlId)) {
		query["IpControlId"] = request.IpControlId
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
		Action:      tea.String("SetIpControlApis"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetIpControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetIpControlApis(request *SetIpControlApisRequest) (_result *SetIpControlApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetIpControlApisResponse{}
	_body, _err := client.SetIpControlApisWithOptions(request, runtime)
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
		Version:     tea.String("2018-06-01"),
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
		Version:     tea.String("2018-06-01"),
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

func (client *Client) VpcDescribeAccessesWithOptions(request *VpcDescribeAccessesRequest, runtime *util.RuntimeOptions) (_result *VpcDescribeAccessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstancePort)) {
		query["InstancePort"] = request.InstancePort
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
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

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcDescribeAccesses"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcDescribeAccessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcDescribeAccesses(request *VpcDescribeAccessesRequest) (_result *VpcDescribeAccessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcDescribeAccessesResponse{}
	_body, _err := client.VpcDescribeAccessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcGrantAccessWithOptions(request *VpcGrantAccessRequest, runtime *util.RuntimeOptions) (_result *VpcGrantAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstancePort)) {
		query["InstancePort"] = request.InstancePort
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

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcGrantAccess"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcGrantAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcGrantAccess(request *VpcGrantAccessRequest) (_result *VpcGrantAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcGrantAccessResponse{}
	_body, _err := client.VpcGrantAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcModifyAccessWithOptions(request *VpcModifyAccessRequest, runtime *util.RuntimeOptions) (_result *VpcModifyAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstancePort)) {
		query["InstancePort"] = request.InstancePort
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcModifyAccess"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcModifyAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcModifyAccess(request *VpcModifyAccessRequest) (_result *VpcModifyAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcModifyAccessResponse{}
	_body, _err := client.VpcModifyAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VpcRevokeAccessWithOptions(request *VpcRevokeAccessRequest, runtime *util.RuntimeOptions) (_result *VpcRevokeAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstancePort)) {
		query["InstancePort"] = request.InstancePort
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VpcRevokeAccess"),
		Version:     tea.String("2018-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &VpcRevokeAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VpcRevokeAccess(request *VpcRevokeAccessRequest) (_result *VpcRevokeAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &VpcRevokeAccessResponse{}
	_body, _err := client.VpcRevokeAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
