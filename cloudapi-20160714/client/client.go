// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
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
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AbolishApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AbolishApiResponse) SetBody(v *AbolishApiResponseBody) *AbolishApiResponse {
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *AddTrafficSpecialControlResponse) SetBody(v *AddTrafficSpecialControlResponseBody) *AddTrafficSpecialControlResponse {
	s.Body = v
	return s
}

type AttachPluginRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AttachPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachPluginRequest) GoString() string {
	return s.String()
}

func (s *AttachPluginRequest) SetApiId(v string) *AttachPluginRequest {
	s.ApiId = &v
	return s
}

func (s *AttachPluginRequest) SetApiIds(v string) *AttachPluginRequest {
	s.ApiIds = &v
	return s
}

func (s *AttachPluginRequest) SetGroupId(v string) *AttachPluginRequest {
	s.GroupId = &v
	return s
}

func (s *AttachPluginRequest) SetPluginId(v string) *AttachPluginRequest {
	s.PluginId = &v
	return s
}

func (s *AttachPluginRequest) SetSecurityToken(v string) *AttachPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *AttachPluginRequest) SetStageName(v string) *AttachPluginRequest {
	s.StageName = &v
	return s
}

type AttachPluginResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachPluginResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPluginResponseBody) SetRequestId(v string) *AttachPluginResponseBody {
	s.RequestId = &v
	return s
}

type AttachPluginResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AttachPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachPluginResponse) GoString() string {
	return s.String()
}

func (s *AttachPluginResponse) SetHeaders(v map[string]*string) *AttachPluginResponse {
	s.Headers = v
	return s
}

func (s *AttachPluginResponse) SetBody(v *AttachPluginResponseBody) *AttachPluginResponse {
	s.Body = v
	return s
}

type BatchAbolishApisRequest struct {
	Api           []*BatchAbolishApisRequestApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Repeated"`
	SecurityToken *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchAbolishApisRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchAbolishApisRequest) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisRequest) SetApi(v []*BatchAbolishApisRequestApi) *BatchAbolishApisRequest {
	s.Api = v
	return s
}

func (s *BatchAbolishApisRequest) SetSecurityToken(v string) *BatchAbolishApisRequest {
	s.SecurityToken = &v
	return s
}

type BatchAbolishApisRequestApi struct {
	ApiUid  *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s BatchAbolishApisRequestApi) String() string {
	return tea.Prettify(s)
}

func (s BatchAbolishApisRequestApi) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisRequestApi) SetApiUid(v string) *BatchAbolishApisRequestApi {
	s.ApiUid = &v
	return s
}

func (s *BatchAbolishApisRequestApi) SetGroupId(v string) *BatchAbolishApisRequestApi {
	s.GroupId = &v
	return s
}

func (s *BatchAbolishApisRequestApi) SetStageId(v string) *BatchAbolishApisRequestApi {
	s.StageId = &v
	return s
}

type BatchAbolishApisResponseBody struct {
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchAbolishApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchAbolishApisResponseBody) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisResponseBody) SetOperationId(v string) *BatchAbolishApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *BatchAbolishApisResponseBody) SetRequestId(v string) *BatchAbolishApisResponseBody {
	s.RequestId = &v
	return s
}

type BatchAbolishApisResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchAbolishApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchAbolishApisResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchAbolishApisResponse) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisResponse) SetHeaders(v map[string]*string) *BatchAbolishApisResponse {
	s.Headers = v
	return s
}

func (s *BatchAbolishApisResponse) SetBody(v *BatchAbolishApisResponseBody) *BatchAbolishApisResponse {
	s.Body = v
	return s
}

type BatchDeployApisRequest struct {
	Api           []*BatchDeployApisRequestApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Repeated"`
	Description   *string                      `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string                      `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string                      `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s BatchDeployApisRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeployApisRequest) GoString() string {
	return s.String()
}

func (s *BatchDeployApisRequest) SetApi(v []*BatchDeployApisRequestApi) *BatchDeployApisRequest {
	s.Api = v
	return s
}

func (s *BatchDeployApisRequest) SetDescription(v string) *BatchDeployApisRequest {
	s.Description = &v
	return s
}

func (s *BatchDeployApisRequest) SetSecurityToken(v string) *BatchDeployApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchDeployApisRequest) SetStageName(v string) *BatchDeployApisRequest {
	s.StageName = &v
	return s
}

type BatchDeployApisRequestApi struct {
	ApiUid  *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s BatchDeployApisRequestApi) String() string {
	return tea.Prettify(s)
}

func (s BatchDeployApisRequestApi) GoString() string {
	return s.String()
}

func (s *BatchDeployApisRequestApi) SetApiUid(v string) *BatchDeployApisRequestApi {
	s.ApiUid = &v
	return s
}

func (s *BatchDeployApisRequestApi) SetGroupId(v string) *BatchDeployApisRequestApi {
	s.GroupId = &v
	return s
}

type BatchDeployApisResponseBody struct {
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchDeployApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeployApisResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeployApisResponseBody) SetOperationId(v string) *BatchDeployApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *BatchDeployApisResponseBody) SetRequestId(v string) *BatchDeployApisResponseBody {
	s.RequestId = &v
	return s
}

type BatchDeployApisResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchDeployApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeployApisResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeployApisResponse) GoString() string {
	return s.String()
}

func (s *BatchDeployApisResponse) SetHeaders(v map[string]*string) *BatchDeployApisResponse {
	s.Headers = v
	return s
}

func (s *BatchDeployApisResponse) SetBody(v *BatchDeployApisResponseBody) *BatchDeployApisResponse {
	s.Body = v
	return s
}

type CreateApiRequest struct {
	AllowSignatureMethod *string `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AppCodeAuthType      *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ConstantParameters   *string `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableInternet      *bool   `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples     *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	FailResultSample     *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	ForceNonceCheck      *bool   `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OpenIdConnectConfig  *string `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty"`
	RequestConfig        *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParameters    *string `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty"`
	ResultBodyModel      *string `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	ResultDescriptions   *string `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty"`
	ResultSample         *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType           *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	SystemParameters     *string `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty"`
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

func (s *CreateApiRequest) SetAppCodeAuthType(v string) *CreateApiRequest {
	s.AppCodeAuthType = &v
	return s
}

func (s *CreateApiRequest) SetAuthType(v string) *CreateApiRequest {
	s.AuthType = &v
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

func (s *CreateApiRequest) SetDisableInternet(v bool) *CreateApiRequest {
	s.DisableInternet = &v
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

func (s *CreateApiRequest) SetForceNonceCheck(v bool) *CreateApiRequest {
	s.ForceNonceCheck = &v
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

func (s *CreateApiRequest) SetRequestParameters(v string) *CreateApiRequest {
	s.RequestParameters = &v
	return s
}

func (s *CreateApiRequest) SetResultBodyModel(v string) *CreateApiRequest {
	s.ResultBodyModel = &v
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

func (s *CreateApiRequest) SetSystemParameters(v string) *CreateApiRequest {
	s.SystemParameters = &v
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateApiResponse) SetBody(v *CreateApiResponseBody) *CreateApiResponse {
	s.Body = v
	return s
}

type CreateApiGroupRequest struct {
	BasePath      *string                     `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	Description   *string                     `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupName     *string                     `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId    *string                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string                     `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*CreateApiGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateApiGroupRequest) SetBasePath(v string) *CreateApiGroupRequest {
	s.BasePath = &v
	return s
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

func (s *CreateApiGroupRequest) SetTag(v []*CreateApiGroupRequestTag) *CreateApiGroupRequest {
	s.Tag = v
	return s
}

type CreateApiGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApiGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateApiGroupRequestTag) SetKey(v string) *CreateApiGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateApiGroupRequestTag) SetValue(v string) *CreateApiGroupRequestTag {
	s.Value = &v
	return s
}

type CreateApiGroupResponseBody struct {
	BasePath     *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubDomain    *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	TagStatus    *bool   `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
}

func (s CreateApiGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApiGroupResponseBody) SetBasePath(v string) *CreateApiGroupResponseBody {
	s.BasePath = &v
	return s
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

func (s *CreateApiGroupResponseBody) SetInstanceId(v string) *CreateApiGroupResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateApiGroupResponseBody) SetInstanceType(v string) *CreateApiGroupResponseBody {
	s.InstanceType = &v
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

func (s *CreateApiGroupResponseBody) SetTagStatus(v bool) *CreateApiGroupResponseBody {
	s.TagStatus = &v
	return s
}

type CreateApiGroupResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateApiStageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateApiStageVariableResponse) SetBody(v *CreateApiStageVariableResponseBody) *CreateApiStageVariableResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	AppName       *string                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description   *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Source        *string                `json:"Source,omitempty" xml:"Source,omitempty"`
	Tag           []*CreateAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateAppRequest) SetSource(v string) *CreateAppRequest {
	s.Source = &v
	return s
}

func (s *CreateAppRequest) SetTag(v []*CreateAppRequestTag) *CreateAppRequest {
	s.Tag = v
	return s
}

type CreateAppRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAppRequestTag) SetKey(v string) *CreateAppRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAppRequestTag) SetValue(v string) *CreateAppRequestTag {
	s.Value = &v
	return s
}

type CreateAppResponseBody struct {
	AppId     *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagStatus *bool   `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
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

func (s *CreateAppResponseBody) SetTagStatus(v bool) *CreateAppResponseBody {
	s.TagStatus = &v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateInstanceRequest struct {
	AutoPay      *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ChargeType   *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Duration     *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	HttpsPolicy  *string `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
	ZoneId       *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoPay(v bool) *CreateInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetDuration(v int32) *CreateInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpsPolicy(v string) *CreateInstanceRequest {
	s.HttpsPolicy = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceSpec(v string) *CreateInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateInstanceRequest) SetPricingCycle(v string) *CreateInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateInstanceRequest) SetToken(v string) *CreateInstanceRequest {
	s.Token = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateIntranetDomainRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateIntranetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIntranetDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateIntranetDomainRequest) SetGroupId(v string) *CreateIntranetDomainRequest {
	s.GroupId = &v
	return s
}

func (s *CreateIntranetDomainRequest) SetSecurityToken(v string) *CreateIntranetDomainRequest {
	s.SecurityToken = &v
	return s
}

type CreateIntranetDomainResponseBody struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntranetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIntranetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntranetDomainResponseBody) SetDomainName(v string) *CreateIntranetDomainResponseBody {
	s.DomainName = &v
	return s
}

func (s *CreateIntranetDomainResponseBody) SetRequestId(v string) *CreateIntranetDomainResponseBody {
	s.RequestId = &v
	return s
}

type CreateIntranetDomainResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIntranetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateIntranetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIntranetDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateIntranetDomainResponse) SetHeaders(v map[string]*string) *CreateIntranetDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateIntranetDomainResponse) SetBody(v *CreateIntranetDomainResponseBody) *CreateIntranetDomainResponse {
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
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CidrIp *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
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

func (s *CreateIpControlRequestIpControlPolicys) SetCidrIp(v string) *CreateIpControlRequestIpControlPolicys {
	s.CidrIp = &v
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateLogConfigResponse) SetBody(v *CreateLogConfigResponseBody) *CreateLogConfigResponse {
	s.Body = v
	return s
}

type CreateModelRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName   *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Schema      *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s CreateModelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) SetDescription(v string) *CreateModelRequest {
	s.Description = &v
	return s
}

func (s *CreateModelRequest) SetGroupId(v string) *CreateModelRequest {
	s.GroupId = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetSchema(v string) *CreateModelRequest {
	s.Schema = &v
	return s
}

type CreateModelResponseBody struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelId      *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName    *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelRef     *string `json:"ModelRef,omitempty" xml:"ModelRef,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s CreateModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelResponseBody) SetCreatedTime(v string) *CreateModelResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *CreateModelResponseBody) SetDescription(v string) *CreateModelResponseBody {
	s.Description = &v
	return s
}

func (s *CreateModelResponseBody) SetGroupId(v string) *CreateModelResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateModelResponseBody) SetModelId(v string) *CreateModelResponseBody {
	s.ModelId = &v
	return s
}

func (s *CreateModelResponseBody) SetModelName(v string) *CreateModelResponseBody {
	s.ModelName = &v
	return s
}

func (s *CreateModelResponseBody) SetModelRef(v string) *CreateModelResponseBody {
	s.ModelRef = &v
	return s
}

func (s *CreateModelResponseBody) SetModifiedTime(v string) *CreateModelResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *CreateModelResponseBody) SetRegionId(v string) *CreateModelResponseBody {
	s.RegionId = &v
	return s
}

func (s *CreateModelResponseBody) SetRequestId(v string) *CreateModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelResponseBody) SetSchema(v string) *CreateModelResponseBody {
	s.Schema = &v
	return s
}

type CreateModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateModelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateModelResponse) GoString() string {
	return s.String()
}

func (s *CreateModelResponse) SetHeaders(v map[string]*string) *CreateModelResponse {
	s.Headers = v
	return s
}

func (s *CreateModelResponse) SetBody(v *CreateModelResponseBody) *CreateModelResponse {
	s.Body = v
	return s
}

type CreateMonitorGroupRequest struct {
	Auth              *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	RawMonitorGroupId *int64  `json:"RawMonitorGroupId,omitempty" xml:"RawMonitorGroupId,omitempty"`
	SecurityToken     *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateMonitorGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupRequest) SetAuth(v string) *CreateMonitorGroupRequest {
	s.Auth = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetGroupId(v string) *CreateMonitorGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetRawMonitorGroupId(v int64) *CreateMonitorGroupRequest {
	s.RawMonitorGroupId = &v
	return s
}

func (s *CreateMonitorGroupRequest) SetSecurityToken(v string) *CreateMonitorGroupRequest {
	s.SecurityToken = &v
	return s
}

type CreateMonitorGroupResponseBody struct {
	MonitorGroupId *int64  `json:"MonitorGroupId,omitempty" xml:"MonitorGroupId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMonitorGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupResponseBody) SetMonitorGroupId(v int64) *CreateMonitorGroupResponseBody {
	s.MonitorGroupId = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetRequestId(v string) *CreateMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateMonitorGroupResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMonitorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMonitorGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMonitorGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupResponse) SetHeaders(v map[string]*string) *CreateMonitorGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorGroupResponse) SetBody(v *CreateMonitorGroupResponseBody) *CreateMonitorGroupResponse {
	s.Body = v
	return s
}

type CreatePluginRequest struct {
	Description   *string                   `json:"Description,omitempty" xml:"Description,omitempty"`
	PluginData    *string                   `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	PluginName    *string                   `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	PluginType    *string                   `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	SecurityToken *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*CreatePluginRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePluginRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePluginRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginRequest) SetDescription(v string) *CreatePluginRequest {
	s.Description = &v
	return s
}

func (s *CreatePluginRequest) SetPluginData(v string) *CreatePluginRequest {
	s.PluginData = &v
	return s
}

func (s *CreatePluginRequest) SetPluginName(v string) *CreatePluginRequest {
	s.PluginName = &v
	return s
}

func (s *CreatePluginRequest) SetPluginType(v string) *CreatePluginRequest {
	s.PluginType = &v
	return s
}

func (s *CreatePluginRequest) SetSecurityToken(v string) *CreatePluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreatePluginRequest) SetTag(v []*CreatePluginRequestTag) *CreatePluginRequest {
	s.Tag = v
	return s
}

type CreatePluginRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePluginRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreatePluginRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePluginRequestTag) SetKey(v string) *CreatePluginRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePluginRequestTag) SetValue(v string) *CreatePluginRequestTag {
	s.Value = &v
	return s
}

type CreatePluginResponseBody struct {
	PluginId  *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagStatus *bool   `json:"TagStatus,omitempty" xml:"TagStatus,omitempty"`
}

func (s CreatePluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePluginResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePluginResponseBody) SetPluginId(v string) *CreatePluginResponseBody {
	s.PluginId = &v
	return s
}

func (s *CreatePluginResponseBody) SetRequestId(v string) *CreatePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePluginResponseBody) SetTagStatus(v bool) *CreatePluginResponseBody {
	s.TagStatus = &v
	return s
}

type CreatePluginResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePluginResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePluginResponse) GoString() string {
	return s.String()
}

func (s *CreatePluginResponse) SetHeaders(v map[string]*string) *CreatePluginResponse {
	s.Headers = v
	return s
}

func (s *CreatePluginResponse) SetBody(v *CreatePluginResponseBody) *CreatePluginResponse {
	s.Body = v
	return s
}

type CreateSignatureRequest struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureKey    *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	SignatureName   *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	SignatureSecret *string `json:"SignatureSecret,omitempty" xml:"SignatureSecret,omitempty"`
}

func (s CreateSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureRequest) GoString() string {
	return s.String()
}

func (s *CreateSignatureRequest) SetSecurityToken(v string) *CreateSignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateSignatureRequest) SetSignatureKey(v string) *CreateSignatureRequest {
	s.SignatureKey = &v
	return s
}

func (s *CreateSignatureRequest) SetSignatureName(v string) *CreateSignatureRequest {
	s.SignatureName = &v
	return s
}

func (s *CreateSignatureRequest) SetSignatureSecret(v string) *CreateSignatureRequest {
	s.SignatureSecret = &v
	return s
}

type CreateSignatureResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s CreateSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponseBody) SetRequestId(v string) *CreateSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSignatureResponseBody) SetSignatureId(v string) *CreateSignatureResponseBody {
	s.SignatureId = &v
	return s
}

func (s *CreateSignatureResponseBody) SetSignatureName(v string) *CreateSignatureResponseBody {
	s.SignatureName = &v
	return s
}

type CreateSignatureResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSignatureResponse) GoString() string {
	return s.String()
}

func (s *CreateSignatureResponse) SetHeaders(v map[string]*string) *CreateSignatureResponse {
	s.Headers = v
	return s
}

func (s *CreateSignatureResponse) SetBody(v *CreateSignatureResponseBody) *CreateSignatureResponse {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAllTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteApiResponse) SetBody(v *DeleteApiResponseBody) *DeleteApiResponse {
	s.Body = v
	return s
}

type DeleteApiGroupRequest struct {
	GroupId       *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string                     `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DeleteApiGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DeleteApiGroupRequest) SetTag(v []*DeleteApiGroupRequestTag) *DeleteApiGroupRequest {
	s.Tag = v
	return s
}

type DeleteApiGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteApiGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DeleteApiGroupRequestTag) GoString() string {
	return s.String()
}

func (s *DeleteApiGroupRequestTag) SetKey(v string) *DeleteApiGroupRequestTag {
	s.Key = &v
	return s
}

func (s *DeleteApiGroupRequestTag) SetValue(v string) *DeleteApiGroupRequestTag {
	s.Value = &v
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteApiStageVariableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteApiStageVariableResponse) SetBody(v *DeleteApiStageVariableResponseBody) *DeleteApiStageVariableResponse {
	s.Body = v
	return s
}

type DeleteAppRequest struct {
	AppId         *int64                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DeleteAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DeleteAppRequest) SetTag(v []*DeleteAppRequestTag) *DeleteAppRequest {
	s.Tag = v
	return s
}

type DeleteAppRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteAppRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppRequestTag) GoString() string {
	return s.String()
}

func (s *DeleteAppRequestTag) SetKey(v string) *DeleteAppRequestTag {
	s.Key = &v
	return s
}

func (s *DeleteAppRequestTag) SetValue(v string) *DeleteAppRequestTag {
	s.Value = &v
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDomainCertificateResponse) SetBody(v *DeleteDomainCertificateResponseBody) *DeleteDomainCertificateResponse {
	s.Body = v
	return s
}

type DeleteInstanceRequest struct {
	InstanceId *string                     `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Tag        []*DeleteInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DeleteInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetTag(v []*DeleteInstanceRequestTag) *DeleteInstanceRequest {
	s.Tag = v
	return s
}

type DeleteInstanceRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequestTag) SetKey(v string) *DeleteInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *DeleteInstanceRequestTag) SetValue(v string) *DeleteInstanceRequestTag {
	s.Value = &v
	return s
}

type DeleteInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponse) SetHeaders(v map[string]*string) *DeleteInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceResponse) SetBody(v *DeleteInstanceResponseBody) *DeleteInstanceResponse {
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteLogConfigResponse) SetBody(v *DeleteLogConfigResponseBody) *DeleteLogConfigResponse {
	s.Body = v
	return s
}

type DeleteModelRequest struct {
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) SetGroupId(v string) *DeleteModelRequest {
	s.GroupId = &v
	return s
}

func (s *DeleteModelRequest) SetModelName(v string) *DeleteModelRequest {
	s.ModelName = &v
	return s
}

type DeleteModelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteModelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteModelResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelResponse) SetHeaders(v map[string]*string) *DeleteModelResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelResponse) SetBody(v *DeleteModelResponseBody) *DeleteModelResponse {
	s.Body = v
	return s
}

type DeletePluginRequest struct {
	PluginId      *string                   `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DeletePluginRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DeletePluginRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePluginRequest) GoString() string {
	return s.String()
}

func (s *DeletePluginRequest) SetPluginId(v string) *DeletePluginRequest {
	s.PluginId = &v
	return s
}

func (s *DeletePluginRequest) SetSecurityToken(v string) *DeletePluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeletePluginRequest) SetTag(v []*DeletePluginRequestTag) *DeletePluginRequest {
	s.Tag = v
	return s
}

type DeletePluginRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeletePluginRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DeletePluginRequestTag) GoString() string {
	return s.String()
}

func (s *DeletePluginRequestTag) SetKey(v string) *DeletePluginRequestTag {
	s.Key = &v
	return s
}

func (s *DeletePluginRequestTag) SetValue(v string) *DeletePluginRequestTag {
	s.Value = &v
	return s
}

type DeletePluginResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePluginResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePluginResponseBody) SetRequestId(v string) *DeletePluginResponseBody {
	s.RequestId = &v
	return s
}

type DeletePluginResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePluginResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePluginResponse) GoString() string {
	return s.String()
}

func (s *DeletePluginResponse) SetHeaders(v map[string]*string) *DeletePluginResponse {
	s.Headers = v
	return s
}

func (s *DeletePluginResponse) SetBody(v *DeletePluginResponseBody) *DeletePluginResponse {
	s.Body = v
	return s
}

type DeleteSignatureRequest struct {
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
}

func (s DeleteSignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureRequest) GoString() string {
	return s.String()
}

func (s *DeleteSignatureRequest) SetSecurityToken(v string) *DeleteSignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteSignatureRequest) SetSignatureId(v string) *DeleteSignatureRequest {
	s.SignatureId = &v
	return s
}

type DeleteSignatureResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponseBody) SetRequestId(v string) *DeleteSignatureResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSignatureResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSignatureResponse) GoString() string {
	return s.String()
}

func (s *DeleteSignatureResponse) SetHeaders(v map[string]*string) *DeleteSignatureResponse {
	s.Headers = v
	return s
}

func (s *DeleteSignatureResponse) SetBody(v *DeleteSignatureResponseBody) *DeleteSignatureResponse {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTrafficSpecialControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeployApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeployApiResponse) SetBody(v *DeployApiResponseBody) *DeployApiResponse {
	s.Body = v
	return s
}

type DescribeAbolishApiTaskRequest struct {
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAbolishApiTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAbolishApiTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskRequest) SetOperationUid(v string) *DescribeAbolishApiTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeAbolishApiTaskRequest) SetSecurityToken(v string) *DescribeAbolishApiTaskRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAbolishApiTaskResponseBody struct {
	ApiAbolishResults *DescribeAbolishApiTaskResponseBodyApiAbolishResults `json:"ApiAbolishResults,omitempty" xml:"ApiAbolishResults,omitempty" type:"Struct"`
	RequestId         *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAbolishApiTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAbolishApiTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponseBody) SetApiAbolishResults(v *DescribeAbolishApiTaskResponseBodyApiAbolishResults) *DescribeAbolishApiTaskResponseBody {
	s.ApiAbolishResults = v
	return s
}

func (s *DescribeAbolishApiTaskResponseBody) SetRequestId(v string) *DescribeAbolishApiTaskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAbolishApiTaskResponseBodyApiAbolishResults struct {
	ApiAbolishResult []*DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult `json:"ApiAbolishResult,omitempty" xml:"ApiAbolishResult,omitempty" type:"Repeated"`
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResults) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResults) SetApiAbolishResult(v []*DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) *DescribeAbolishApiTaskResponseBodyApiAbolishResults {
	s.ApiAbolishResult = v
	return s
}

type DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult struct {
	AbolishStatus *string `json:"AbolishStatus,omitempty" xml:"AbolishStatus,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiUid        *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	ErrorMsg      *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	StageId       *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetAbolishStatus(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.AbolishStatus = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetApiName(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.ApiName = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetApiUid(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetErrorMsg(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetGroupId(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.GroupId = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetGroupName(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.GroupName = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetStageId(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.StageId = &v
	return s
}

func (s *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult) SetStageName(v string) *DescribeAbolishApiTaskResponseBodyApiAbolishResultsApiAbolishResult {
	s.StageName = &v
	return s
}

type DescribeAbolishApiTaskResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAbolishApiTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAbolishApiTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAbolishApiTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponse) SetHeaders(v map[string]*string) *DescribeAbolishApiTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeAbolishApiTaskResponse) SetBody(v *DescribeAbolishApiTaskResponseBody) *DescribeAbolishApiTaskResponse {
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
	AllowSignatureMethod   *string                                        `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiId                  *string                                        `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                *string                                        `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AppCodeAuthType        *string                                        `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	AuthType               *string                                        `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ConstantParameters     *DescribeApiResponseBodyConstantParameters     `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	CreatedTime            *string                                        `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CustomSystemParameters *DescribeApiResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	DeployedInfos          *DescribeApiResponseBodyDeployedInfos          `json:"DeployedInfos,omitempty" xml:"DeployedInfos,omitempty" type:"Struct"`
	Description            *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableInternet        *bool                                          `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples       *DescribeApiResponseBodyErrorCodeSamples       `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample       *string                                        `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	ForceNonceCheck        *bool                                          `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	GroupId                *string                                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName              *string                                        `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Mock                   *string                                        `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockResult             *string                                        `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	ModifiedTime           *string                                        `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	OpenIdConnectConfig    *DescribeApiResponseBodyOpenIdConnectConfig    `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	RegionId               *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestConfig          *DescribeApiResponseBodyRequestConfig          `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	RequestId              *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestParameters      *DescribeApiResponseBodyRequestParameters      `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	ResultBodyModel        *string                                        `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	ResultDescriptions     *DescribeApiResponseBodyResultDescriptions     `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	ResultSample           *string                                        `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType             *string                                        `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceConfig          *DescribeApiResponseBodyServiceConfig          `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	ServiceParameters      *DescribeApiResponseBodyServiceParameters      `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty" type:"Struct"`
	ServiceParametersMap   *DescribeApiResponseBodyServiceParametersMap   `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty" type:"Struct"`
	SystemParameters       *DescribeApiResponseBodySystemParameters       `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility             *string                                        `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	WebSocketApiType       *string                                        `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
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

func (s *DescribeApiResponseBody) SetAppCodeAuthType(v string) *DescribeApiResponseBody {
	s.AppCodeAuthType = &v
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

func (s *DescribeApiResponseBody) SetDisableInternet(v bool) *DescribeApiResponseBody {
	s.DisableInternet = &v
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

func (s *DescribeApiResponseBody) SetForceNonceCheck(v bool) *DescribeApiResponseBody {
	s.ForceNonceCheck = &v
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

func (s *DescribeApiResponseBody) SetRequestParameters(v *DescribeApiResponseBodyRequestParameters) *DescribeApiResponseBody {
	s.RequestParameters = v
	return s
}

func (s *DescribeApiResponseBody) SetResultBodyModel(v string) *DescribeApiResponseBody {
	s.ResultBodyModel = &v
	return s
}

func (s *DescribeApiResponseBody) SetResultDescriptions(v *DescribeApiResponseBodyResultDescriptions) *DescribeApiResponseBody {
	s.ResultDescriptions = v
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

func (s *DescribeApiResponseBody) SetServiceParameters(v *DescribeApiResponseBodyServiceParameters) *DescribeApiResponseBody {
	s.ServiceParameters = v
	return s
}

func (s *DescribeApiResponseBody) SetServiceParametersMap(v *DescribeApiResponseBodyServiceParametersMap) *DescribeApiResponseBody {
	s.ServiceParametersMap = v
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

type DescribeApiResponseBodyRequestConfig struct {
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	BodyModel           *string `json:"BodyModel,omitempty" xml:"BodyModel,omitempty"`
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

func (s *DescribeApiResponseBodyRequestConfig) SetBodyModel(v string) *DescribeApiResponseBodyRequestConfig {
	s.BodyModel = &v
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

type DescribeApiResponseBodyRequestParameters struct {
	RequestParameter []*DescribeApiResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestParameters) SetRequestParameter(v []*DescribeApiResponseBodyRequestParametersRequestParameter) *DescribeApiResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

type DescribeApiResponseBodyRequestParametersRequestParameter struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	ArrayItemsType    *string `json:"ArrayItemsType,omitempty" xml:"ArrayItemsType,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *int32  `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
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

func (s DescribeApiResponseBodyRequestParametersRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetApiParameterName(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetArrayItemsType(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.ArrayItemsType = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetDefaultValue(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetDemoValue(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetDescription(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetDocOrder(v int32) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.DocOrder = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetDocShow(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.DocShow = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetEnumValue(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetJsonScheme(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetLocation(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetMaxLength(v int64) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetMaxValue(v int64) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetMinLength(v int64) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.MinLength = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetMinValue(v int64) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.MinValue = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetParameterType(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetRegularExpression(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) SetRequired(v string) *DescribeApiResponseBodyRequestParametersRequestParameter {
	s.Required = &v
	return s
}

type DescribeApiResponseBodyResultDescriptions struct {
	ResultDescription []*DescribeApiResponseBodyResultDescriptionsResultDescription `json:"ResultDescription,omitempty" xml:"ResultDescription,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyResultDescriptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyResultDescriptions) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyResultDescriptions) SetResultDescription(v []*DescribeApiResponseBodyResultDescriptionsResultDescription) *DescribeApiResponseBodyResultDescriptions {
	s.ResultDescription = v
	return s
}

type DescribeApiResponseBodyResultDescriptionsResultDescription struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasChild    *bool   `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Mandatory   *bool   `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApiResponseBodyResultDescriptionsResultDescription) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyResultDescriptionsResultDescription) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetDescription(v string) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Description = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetHasChild(v bool) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.HasChild = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetId(v string) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Id = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetKey(v string) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Key = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetMandatory(v bool) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Mandatory = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetName(v string) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Name = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetPid(v string) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Pid = &v
	return s
}

func (s *DescribeApiResponseBodyResultDescriptionsResultDescription) SetType(v string) *DescribeApiResponseBodyResultDescriptionsResultDescription {
	s.Type = &v
	return s
}

type DescribeApiResponseBodyServiceConfig struct {
	AoneAppName           *string                                                    `json:"AoneAppName,omitempty" xml:"AoneAppName,omitempty"`
	ContentTypeCatagory   *string                                                    `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	ContentTypeValue      *string                                                    `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	FunctionComputeConfig *DescribeApiResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	Mock                  *string                                                    `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockHeaders           *DescribeApiResponseBodyServiceConfigMockHeaders           `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Struct"`
	MockResult            *string                                                    `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	MockStatusCode        *int32                                                     `json:"MockStatusCode,omitempty" xml:"MockStatusCode,omitempty"`
	OssConfig             *DescribeApiResponseBodyServiceConfigOssConfig             `json:"OssConfig,omitempty" xml:"OssConfig,omitempty" type:"Struct"`
	ServiceAddress        *string                                                    `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceHttpMethod     *string                                                    `json:"ServiceHttpMethod,omitempty" xml:"ServiceHttpMethod,omitempty"`
	ServicePath           *string                                                    `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	ServiceProtocol       *string                                                    `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout        *int32                                                     `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
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

func (s *DescribeApiResponseBodyServiceConfig) SetMockHeaders(v *DescribeApiResponseBodyServiceConfigMockHeaders) *DescribeApiResponseBodyServiceConfig {
	s.MockHeaders = v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetMockResult(v string) *DescribeApiResponseBodyServiceConfig {
	s.MockResult = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetMockStatusCode(v int32) *DescribeApiResponseBodyServiceConfig {
	s.MockStatusCode = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfig) SetOssConfig(v *DescribeApiResponseBodyServiceConfigOssConfig) *DescribeApiResponseBodyServiceConfig {
	s.OssConfig = v
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

func (s *DescribeApiResponseBodyServiceConfig) SetServiceTimeout(v int32) *DescribeApiResponseBodyServiceConfig {
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
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	ContentTypeValue    *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	FcBaseUrl           *string `json:"FcBaseUrl,omitempty" xml:"FcBaseUrl,omitempty"`
	FcType              *string `json:"FcType,omitempty" xml:"FcType,omitempty"`
	FunctionName        *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Method              *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OnlyBusinessPath    *bool   `json:"OnlyBusinessPath,omitempty" xml:"OnlyBusinessPath,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Qualifier           *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn             *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName         *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetContentTypeCatagory(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.ContentTypeCatagory = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetContentTypeValue(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.ContentTypeValue = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetFcBaseUrl(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FcBaseUrl = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetFcType(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FcType = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetFunctionName(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetMethod(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.Method = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetOnlyBusinessPath(v bool) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.OnlyBusinessPath = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetPath(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.Path = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetQualifier(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.Qualifier = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetRegionId(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.RegionId = &v
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

type DescribeApiResponseBodyServiceConfigMockHeaders struct {
	MockHeader []*DescribeApiResponseBodyServiceConfigMockHeadersMockHeader `json:"MockHeader,omitempty" xml:"MockHeader,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceConfigMockHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigMockHeaders) SetMockHeader(v []*DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) *DescribeApiResponseBodyServiceConfigMockHeaders {
	s.MockHeader = v
	return s
}

type DescribeApiResponseBodyServiceConfigMockHeadersMockHeader struct {
	HeaderName  *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderName(v string) *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderValue(v string) *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderValue = &v
	return s
}

type DescribeApiResponseBodyServiceConfigOssConfig struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	BucketName  *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigOssConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigOssConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) SetAction(v string) *DescribeApiResponseBodyServiceConfigOssConfig {
	s.Action = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) SetBucketName(v string) *DescribeApiResponseBodyServiceConfigOssConfig {
	s.BucketName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) SetKey(v string) *DescribeApiResponseBodyServiceConfigOssConfig {
	s.Key = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) SetOssRegionId(v string) *DescribeApiResponseBodyServiceConfigOssConfig {
	s.OssRegionId = &v
	return s
}

type DescribeApiResponseBodyServiceConfigVpcConfig struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcScheme  *string `json:"VpcScheme,omitempty" xml:"VpcScheme,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
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

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) SetVpcScheme(v string) *DescribeApiResponseBodyServiceConfigVpcConfig {
	s.VpcScheme = &v
	return s
}

type DescribeApiResponseBodyServiceParameters struct {
	ServiceParameter []*DescribeApiResponseBodyServiceParametersServiceParameter `json:"ServiceParameter,omitempty" xml:"ServiceParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParameters) SetServiceParameter(v []*DescribeApiResponseBodyServiceParametersServiceParameter) *DescribeApiResponseBodyServiceParameters {
	s.ServiceParameter = v
	return s
}

type DescribeApiResponseBodyServiceParametersServiceParameter struct {
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodyServiceParametersServiceParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersServiceParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) SetLocation(v string) *DescribeApiResponseBodyServiceParametersServiceParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) SetParameterType(v string) *DescribeApiResponseBodyServiceParametersServiceParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) SetServiceParameterName(v string) *DescribeApiResponseBodyServiceParametersServiceParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiResponseBodyServiceParametersMap struct {
	ServiceParameterMap []*DescribeApiResponseBodyServiceParametersMapServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceParametersMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersMap) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersMap) SetServiceParameterMap(v []*DescribeApiResponseBodyServiceParametersMapServiceParameterMap) *DescribeApiResponseBodyServiceParametersMap {
	s.ServiceParameterMap = v
	return s
}

type DescribeApiResponseBodyServiceParametersMapServiceParameterMap struct {
	RequestParameterName *string `json:"RequestParameterName,omitempty" xml:"RequestParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodyServiceParametersMapServiceParameterMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersMapServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) SetRequestParameterName(v string) *DescribeApiResponseBodyServiceParametersMapServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) SetServiceParameterName(v string) *DescribeApiResponseBodyServiceParametersMapServiceParameterMap {
	s.ServiceParameterName = &v
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApiId              *string                                       `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName            *string                                       `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType           *string                                       `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	DeployedTime       *string                                       `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description        *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableInternet    *bool                                         `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples   *DescribeApiDocResponseBodyErrorCodeSamples   `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample   *string                                       `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	ForceNonceCheck    *bool                                         `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	GroupId            *string                                       `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName          *string                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	RegionId           *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestConfig      *DescribeApiDocResponseBodyRequestConfig      `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestParameters  *DescribeApiDocResponseBodyRequestParameters  `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	ResultDescriptions *DescribeApiDocResponseBodyResultDescriptions `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	ResultSample       *string                                       `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType         *string                                       `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	StageName          *string                                       `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility         *string                                       `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
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

func (s *DescribeApiDocResponseBody) SetAuthType(v string) *DescribeApiDocResponseBody {
	s.AuthType = &v
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

func (s *DescribeApiDocResponseBody) SetDisableInternet(v bool) *DescribeApiDocResponseBody {
	s.DisableInternet = &v
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

func (s *DescribeApiDocResponseBody) SetForceNonceCheck(v bool) *DescribeApiDocResponseBody {
	s.ForceNonceCheck = &v
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

func (s *DescribeApiDocResponseBody) SetRegionId(v string) *DescribeApiDocResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestConfig(v *DescribeApiDocResponseBodyRequestConfig) *DescribeApiDocResponseBody {
	s.RequestConfig = v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestId(v string) *DescribeApiDocResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetRequestParameters(v *DescribeApiDocResponseBodyRequestParameters) *DescribeApiDocResponseBody {
	s.RequestParameters = v
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

func (s *DescribeApiDocResponseBody) SetStageName(v string) *DescribeApiDocResponseBody {
	s.StageName = &v
	return s
}

func (s *DescribeApiDocResponseBody) SetVisibility(v string) *DescribeApiDocResponseBody {
	s.Visibility = &v
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

type DescribeApiDocResponseBodyRequestConfig struct {
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	RequestHttpMethod   *string `json:"RequestHttpMethod,omitempty" xml:"RequestHttpMethod,omitempty"`
	RequestMode         *string `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	RequestPath         *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	RequestProtocol     *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeApiDocResponseBodyRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetPostBodyDescription(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetRequestHttpMethod(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.RequestHttpMethod = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetRequestMode(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.RequestMode = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetRequestPath(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.RequestPath = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetRequestProtocol(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.RequestProtocol = &v
	return s
}

type DescribeApiDocResponseBodyRequestParameters struct {
	RequestParameter []*DescribeApiDocResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestParameters) SetRequestParameter(v []*DescribeApiDocResponseBodyRequestParametersRequestParameter) *DescribeApiDocResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

type DescribeApiDocResponseBodyRequestParametersRequestParameter struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	ArrayItemsType    *string `json:"ArrayItemsType,omitempty" xml:"ArrayItemsType,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *int32  `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
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

func (s DescribeApiDocResponseBodyRequestParametersRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetApiParameterName(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetArrayItemsType(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.ArrayItemsType = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetDefaultValue(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetDemoValue(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetDescription(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetDocOrder(v int32) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.DocOrder = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetDocShow(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.DocShow = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetEnumValue(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetJsonScheme(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetLocation(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetMaxLength(v int64) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetMaxValue(v int64) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetMinLength(v int64) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.MinLength = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetMinValue(v int64) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.MinValue = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetParameterType(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetRegularExpression(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) SetRequired(v string) *DescribeApiDocResponseBodyRequestParametersRequestParameter {
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

type DescribeApiDocResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiDocResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeApiDocResponse) SetBody(v *DescribeApiDocResponseBody) *DescribeApiDocResponse {
	s.Body = v
	return s
}

type DescribeApiGroupRequest struct {
	GroupId       *string                       `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DescribeApiGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupRequest) SetGroupId(v string) *DescribeApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupRequest) SetSecurityToken(v string) *DescribeApiGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiGroupRequest) SetTag(v []*DescribeApiGroupRequestTag) *DescribeApiGroupRequest {
	s.Tag = v
	return s
}

type DescribeApiGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupRequestTag) SetKey(v string) *DescribeApiGroupRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeApiGroupRequestTag) SetValue(v string) *DescribeApiGroupRequestTag {
	s.Value = &v
	return s
}

type DescribeApiGroupResponseBody struct {
	BasePath             *string                                    `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	BillingStatus        *string                                    `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	ClassicVpcSubDomain  *string                                    `json:"ClassicVpcSubDomain,omitempty" xml:"ClassicVpcSubDomain,omitempty"`
	CmsMonitorGroup      *string                                    `json:"CmsMonitorGroup,omitempty" xml:"CmsMonitorGroup,omitempty"`
	CompatibleFlags      *string                                    `json:"CompatibleFlags,omitempty" xml:"CompatibleFlags,omitempty"`
	CreatedTime          *string                                    `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	CustomDomains        *DescribeApiGroupResponseBodyCustomDomains `json:"CustomDomains,omitempty" xml:"CustomDomains,omitempty" type:"Struct"`
	CustomTraceConfig    *string                                    `json:"CustomTraceConfig,omitempty" xml:"CustomTraceConfig,omitempty"`
	CustomerConfigs      *string                                    `json:"CustomerConfigs,omitempty" xml:"CustomerConfigs,omitempty"`
	DefaultDomain        *string                                    `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	Description          *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId              *string                                    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName            *string                                    `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpsPolicy          *string                                    `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	IllegalStatus        *string                                    `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	InstanceId           *string                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType         *string                                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceVipList      *string                                    `json:"InstanceVipList,omitempty" xml:"InstanceVipList,omitempty"`
	Ipv6Status           *string                                    `json:"Ipv6Status,omitempty" xml:"Ipv6Status,omitempty"`
	ModifiedTime         *string                                    `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PassthroughHeaders   *string                                    `json:"PassthroughHeaders,omitempty" xml:"PassthroughHeaders,omitempty"`
	RegionId             *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId            *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RpcPattern           *string                                    `json:"RpcPattern,omitempty" xml:"RpcPattern,omitempty"`
	StageItems           *DescribeApiGroupResponseBodyStageItems    `json:"StageItems,omitempty" xml:"StageItems,omitempty" type:"Struct"`
	Status               *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	SubDomain            *string                                    `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	TrafficLimit         *int32                                     `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
	UserLogConfig        *string                                    `json:"UserLogConfig,omitempty" xml:"UserLogConfig,omitempty"`
	VpcDomain            *string                                    `json:"VpcDomain,omitempty" xml:"VpcDomain,omitempty"`
	VpcSlbIntranetDomain *string                                    `json:"VpcSlbIntranetDomain,omitempty" xml:"VpcSlbIntranetDomain,omitempty"`
}

func (s DescribeApiGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBody) SetBasePath(v string) *DescribeApiGroupResponseBody {
	s.BasePath = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetBillingStatus(v string) *DescribeApiGroupResponseBody {
	s.BillingStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetClassicVpcSubDomain(v string) *DescribeApiGroupResponseBody {
	s.ClassicVpcSubDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCmsMonitorGroup(v string) *DescribeApiGroupResponseBody {
	s.CmsMonitorGroup = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCompatibleFlags(v string) *DescribeApiGroupResponseBody {
	s.CompatibleFlags = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCreatedTime(v string) *DescribeApiGroupResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomDomains(v *DescribeApiGroupResponseBodyCustomDomains) *DescribeApiGroupResponseBody {
	s.CustomDomains = v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomTraceConfig(v string) *DescribeApiGroupResponseBody {
	s.CustomTraceConfig = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetCustomerConfigs(v string) *DescribeApiGroupResponseBody {
	s.CustomerConfigs = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetDefaultDomain(v string) *DescribeApiGroupResponseBody {
	s.DefaultDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetDescription(v string) *DescribeApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetGroupId(v string) *DescribeApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetGroupName(v string) *DescribeApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetHttpsPolicy(v string) *DescribeApiGroupResponseBody {
	s.HttpsPolicy = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetIllegalStatus(v string) *DescribeApiGroupResponseBody {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetInstanceId(v string) *DescribeApiGroupResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetInstanceType(v string) *DescribeApiGroupResponseBody {
	s.InstanceType = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetInstanceVipList(v string) *DescribeApiGroupResponseBody {
	s.InstanceVipList = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetIpv6Status(v string) *DescribeApiGroupResponseBody {
	s.Ipv6Status = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetModifiedTime(v string) *DescribeApiGroupResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetPassthroughHeaders(v string) *DescribeApiGroupResponseBody {
	s.PassthroughHeaders = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetRegionId(v string) *DescribeApiGroupResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetRequestId(v string) *DescribeApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetRpcPattern(v string) *DescribeApiGroupResponseBody {
	s.RpcPattern = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetStageItems(v *DescribeApiGroupResponseBodyStageItems) *DescribeApiGroupResponseBody {
	s.StageItems = v
	return s
}

func (s *DescribeApiGroupResponseBody) SetStatus(v string) *DescribeApiGroupResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetSubDomain(v string) *DescribeApiGroupResponseBody {
	s.SubDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetTrafficLimit(v int32) *DescribeApiGroupResponseBody {
	s.TrafficLimit = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetUserLogConfig(v string) *DescribeApiGroupResponseBody {
	s.UserLogConfig = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetVpcDomain(v string) *DescribeApiGroupResponseBody {
	s.VpcDomain = &v
	return s
}

func (s *DescribeApiGroupResponseBody) SetVpcSlbIntranetDomain(v string) *DescribeApiGroupResponseBody {
	s.VpcSlbIntranetDomain = &v
	return s
}

type DescribeApiGroupResponseBodyCustomDomains struct {
	DomainItem []*DescribeApiGroupResponseBodyCustomDomainsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupResponseBodyCustomDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupResponseBodyCustomDomains) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyCustomDomains) SetDomainItem(v []*DescribeApiGroupResponseBodyCustomDomainsDomainItem) *DescribeApiGroupResponseBodyCustomDomains {
	s.DomainItem = v
	return s
}

type DescribeApiGroupResponseBodyCustomDomainsDomainItem struct {
	BindStageName          *string `json:"BindStageName,omitempty" xml:"BindStageName,omitempty"`
	CertificateId          *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	CertificateName        *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CustomDomainType       *string `json:"CustomDomainType,omitempty" xml:"CustomDomainType,omitempty"`
	DomainBindingStatus    *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	DomainCNAMEStatus      *string `json:"DomainCNAMEStatus,omitempty" xml:"DomainCNAMEStatus,omitempty"`
	DomainLegalStatus      *string `json:"DomainLegalStatus,omitempty" xml:"DomainLegalStatus,omitempty"`
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainRemark           *string `json:"DomainRemark,omitempty" xml:"DomainRemark,omitempty"`
	DomainWebSocketStatus  *string `json:"DomainWebSocketStatus,omitempty" xml:"DomainWebSocketStatus,omitempty"`
	WildcardDomainPatterns *string `json:"WildcardDomainPatterns,omitempty" xml:"WildcardDomainPatterns,omitempty"`
}

func (s DescribeApiGroupResponseBodyCustomDomainsDomainItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupResponseBodyCustomDomainsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetBindStageName(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.BindStageName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCertificateId(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CertificateId = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCertificateName(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CertificateName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetCustomDomainType(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.CustomDomainType = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainBindingStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainBindingStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainCNAMEStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainCNAMEStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainLegalStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainLegalStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainName(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainName = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainRemark(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainRemark = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetDomainWebSocketStatus(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.DomainWebSocketStatus = &v
	return s
}

func (s *DescribeApiGroupResponseBodyCustomDomainsDomainItem) SetWildcardDomainPatterns(v string) *DescribeApiGroupResponseBodyCustomDomainsDomainItem {
	s.WildcardDomainPatterns = &v
	return s
}

type DescribeApiGroupResponseBodyStageItems struct {
	StageInfo []*DescribeApiGroupResponseBodyStageItemsStageInfo `json:"StageInfo,omitempty" xml:"StageInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupResponseBodyStageItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupResponseBodyStageItems) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyStageItems) SetStageInfo(v []*DescribeApiGroupResponseBodyStageItemsStageInfo) *DescribeApiGroupResponseBodyStageItems {
	s.StageInfo = v
	return s
}

type DescribeApiGroupResponseBodyStageItemsStageInfo struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	StageId     *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName   *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiGroupResponseBodyStageItemsStageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupResponseBodyStageItemsStageInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) SetDescription(v string) *DescribeApiGroupResponseBodyStageItemsStageInfo {
	s.Description = &v
	return s
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) SetStageId(v string) *DescribeApiGroupResponseBodyStageItemsStageInfo {
	s.StageId = &v
	return s
}

func (s *DescribeApiGroupResponseBodyStageItemsStageInfo) SetStageName(v string) *DescribeApiGroupResponseBodyStageItemsStageInfo {
	s.StageName = &v
	return s
}

type DescribeApiGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupResponse) SetHeaders(v map[string]*string) *DescribeApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupResponse) SetBody(v *DescribeApiGroupResponseBody) *DescribeApiGroupResponse {
	s.Body = v
	return s
}

type DescribeApiGroupVpcWhitelistRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiGroupVpcWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupVpcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupVpcWhitelistRequest) SetGroupId(v string) *DescribeApiGroupVpcWhitelistRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistRequest) SetSecurityToken(v string) *DescribeApiGroupVpcWhitelistRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApiGroupVpcWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VpcIds    *string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty"`
}

func (s DescribeApiGroupVpcWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupVpcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) SetRequestId(v string) *DescribeApiGroupVpcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) SetVpcIds(v string) *DescribeApiGroupVpcWhitelistResponseBody {
	s.VpcIds = &v
	return s
}

type DescribeApiGroupVpcWhitelistResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiGroupVpcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiGroupVpcWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupVpcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupVpcWhitelistResponse) SetHeaders(v map[string]*string) *DescribeApiGroupVpcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponse) SetBody(v *DescribeApiGroupVpcWhitelistResponseBody) *DescribeApiGroupVpcWhitelistResponse {
	s.Body = v
	return s
}

type DescribeApiGroupsRequest struct {
	EnableTagAuth *bool                          `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	GroupId       *string                        `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                        `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId    *string                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber    *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string                        `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sort          *string                        `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Tag           []*DescribeApiGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsRequest) SetEnableTagAuth(v bool) *DescribeApiGroupsRequest {
	s.EnableTagAuth = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetGroupId(v string) *DescribeApiGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetGroupName(v string) *DescribeApiGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetInstanceId(v string) *DescribeApiGroupsRequest {
	s.InstanceId = &v
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

func (s *DescribeApiGroupsRequest) SetSort(v string) *DescribeApiGroupsRequest {
	s.Sort = &v
	return s
}

func (s *DescribeApiGroupsRequest) SetTag(v []*DescribeApiGroupsRequestTag) *DescribeApiGroupsRequest {
	s.Tag = v
	return s
}

type DescribeApiGroupsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiGroupsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsRequestTag) SetKey(v string) *DescribeApiGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeApiGroupsRequestTag) SetValue(v string) *DescribeApiGroupsRequestTag {
	s.Value = &v
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
	BasePath      *string                                                               `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	BillingStatus *string                                                               `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	CreatedTime   *string                                                               `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description   *string                                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string                                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HttpsPolicy   *string                                                               `json:"HttpsPolicy,omitempty" xml:"HttpsPolicy,omitempty"`
	IllegalStatus *string                                                               `json:"IllegalStatus,omitempty" xml:"IllegalStatus,omitempty"`
	InstanceId    *string                                                               `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType  *string                                                               `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	ModifiedTime  *string                                                               `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId      *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubDomain     *string                                                               `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	Tags          *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TrafficLimit  *int32                                                                `json:"TrafficLimit,omitempty" xml:"TrafficLimit,omitempty"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetBasePath(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.BasePath = &v
	return s
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

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetHttpsPolicy(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.HttpsPolicy = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetIllegalStatus(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.IllegalStatus = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetInstanceId(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetInstanceType(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.InstanceType = &v
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

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetTags(v *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.Tags = v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute) SetTrafficLimit(v int32) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttribute {
	s.TrafficLimit = &v
	return s
}

type DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags struct {
	TagInfo []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags) SetTagInfo(v []*DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTags {
	s.TagInfo = v
	return s
}

type DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) SetKey(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo) SetValue(v string) *DescribeApiGroupsResponseBodyApiGroupAttributesApiGroupAttributeTagsTagInfo {
	s.Value = &v
	return s
}

type DescribeApiGroupsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeApiGroupsResponse) SetBody(v *DescribeApiGroupsResponseBody) *DescribeApiGroupsResponse {
	s.Body = v
	return s
}

type DescribeApiHistoriesRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiHistoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesRequest) SetApiId(v string) *DescribeApiHistoriesRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetApiName(v string) *DescribeApiHistoriesRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetGroupId(v string) *DescribeApiHistoriesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetPageNumber(v int32) *DescribeApiHistoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetPageSize(v int32) *DescribeApiHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetSecurityToken(v string) *DescribeApiHistoriesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiHistoriesRequest) SetStageName(v string) *DescribeApiHistoriesRequest {
	s.StageName = &v
	return s
}

type DescribeApiHistoriesResponseBody struct {
	ApiHisItems *DescribeApiHistoriesResponseBodyApiHisItems `json:"ApiHisItems,omitempty" xml:"ApiHisItems,omitempty" type:"Struct"`
	PageNumber  *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiHistoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponseBody) SetApiHisItems(v *DescribeApiHistoriesResponseBodyApiHisItems) *DescribeApiHistoriesResponseBody {
	s.ApiHisItems = v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetPageNumber(v int32) *DescribeApiHistoriesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetPageSize(v int32) *DescribeApiHistoriesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetRequestId(v string) *DescribeApiHistoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBody) SetTotalCount(v int32) *DescribeApiHistoriesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiHistoriesResponseBodyApiHisItems struct {
	ApiHisItem []*DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem `json:"ApiHisItem,omitempty" xml:"ApiHisItem,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoriesResponseBodyApiHisItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoriesResponseBodyApiHisItems) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponseBodyApiHisItems) SetApiHisItem(v []*DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) *DescribeApiHistoriesResponseBodyApiHisItems {
	s.ApiHisItem = v
	return s
}

type DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem struct {
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

func (s DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetApiId(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetApiName(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetDeployedTime(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.DeployedTime = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetDescription(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetGroupId(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.GroupId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetGroupName(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.GroupName = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetHistoryVersion(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetRegionId(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.RegionId = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetStageName(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.StageName = &v
	return s
}

func (s *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem) SetStatus(v string) *DescribeApiHistoriesResponseBodyApiHisItemsApiHisItem {
	s.Status = &v
	return s
}

type DescribeApiHistoriesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiHistoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponse) SetHeaders(v map[string]*string) *DescribeApiHistoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiHistoriesResponse) SetBody(v *DescribeApiHistoriesResponseBody) *DescribeApiHistoriesResponse {
	s.Body = v
	return s
}

type DescribeApiHistoryRequest struct {
	ApiId          *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryRequest) SetApiId(v string) *DescribeApiHistoryRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiHistoryRequest) SetGroupId(v string) *DescribeApiHistoryRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiHistoryRequest) SetHistoryVersion(v string) *DescribeApiHistoryRequest {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeApiHistoryRequest) SetSecurityToken(v string) *DescribeApiHistoryRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiHistoryRequest) SetStageName(v string) *DescribeApiHistoryRequest {
	s.StageName = &v
	return s
}

type DescribeApiHistoryResponseBody struct {
	AllowSignatureMethod   *string                                               `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiId                  *string                                               `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                *string                                               `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AppCodeAuthType        *string                                               `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	AuthType               *string                                               `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ConstantParameters     *DescribeApiHistoryResponseBodyConstantParameters     `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	CustomSystemParameters *DescribeApiHistoryResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	DeployedTime           *string                                               `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description            *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableInternet        *bool                                                 `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples       *DescribeApiHistoryResponseBodyErrorCodeSamples       `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample       *string                                               `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	ForceNonceCheck        *bool                                                 `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	GroupId                *string                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName              *string                                               `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	HistoryVersion         *string                                               `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	OpenIdConnectConfig    *DescribeApiHistoryResponseBodyOpenIdConnectConfig    `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	RegionId               *string                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestConfig          *DescribeApiHistoryResponseBodyRequestConfig          `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	RequestId              *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestParameters      *DescribeApiHistoryResponseBodyRequestParameters      `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	ResultBodyModel        *string                                               `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	ResultDescriptions     *DescribeApiHistoryResponseBodyResultDescriptions     `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	ResultSample           *string                                               `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType             *string                                               `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceConfig          *DescribeApiHistoryResponseBodyServiceConfig          `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	ServiceParameters      *DescribeApiHistoryResponseBodyServiceParameters      `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty" type:"Struct"`
	ServiceParametersMap   *DescribeApiHistoryResponseBodyServiceParametersMap   `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty" type:"Struct"`
	StageName              *string                                               `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Status                 *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	SystemParameters       *DescribeApiHistoryResponseBodySystemParameters       `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility             *string                                               `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	WebSocketApiType       *string                                               `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s DescribeApiHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBody) SetAllowSignatureMethod(v string) *DescribeApiHistoryResponseBody {
	s.AllowSignatureMethod = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetApiId(v string) *DescribeApiHistoryResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetApiName(v string) *DescribeApiHistoryResponseBody {
	s.ApiName = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetAppCodeAuthType(v string) *DescribeApiHistoryResponseBody {
	s.AppCodeAuthType = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetAuthType(v string) *DescribeApiHistoryResponseBody {
	s.AuthType = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetConstantParameters(v *DescribeApiHistoryResponseBodyConstantParameters) *DescribeApiHistoryResponseBody {
	s.ConstantParameters = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetCustomSystemParameters(v *DescribeApiHistoryResponseBodyCustomSystemParameters) *DescribeApiHistoryResponseBody {
	s.CustomSystemParameters = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetDeployedTime(v string) *DescribeApiHistoryResponseBody {
	s.DeployedTime = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetDescription(v string) *DescribeApiHistoryResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetDisableInternet(v bool) *DescribeApiHistoryResponseBody {
	s.DisableInternet = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetErrorCodeSamples(v *DescribeApiHistoryResponseBodyErrorCodeSamples) *DescribeApiHistoryResponseBody {
	s.ErrorCodeSamples = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetFailResultSample(v string) *DescribeApiHistoryResponseBody {
	s.FailResultSample = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetForceNonceCheck(v bool) *DescribeApiHistoryResponseBody {
	s.ForceNonceCheck = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetGroupId(v string) *DescribeApiHistoryResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetGroupName(v string) *DescribeApiHistoryResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetHistoryVersion(v string) *DescribeApiHistoryResponseBody {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetOpenIdConnectConfig(v *DescribeApiHistoryResponseBodyOpenIdConnectConfig) *DescribeApiHistoryResponseBody {
	s.OpenIdConnectConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetRegionId(v string) *DescribeApiHistoryResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetRequestConfig(v *DescribeApiHistoryResponseBodyRequestConfig) *DescribeApiHistoryResponseBody {
	s.RequestConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetRequestId(v string) *DescribeApiHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetRequestParameters(v *DescribeApiHistoryResponseBodyRequestParameters) *DescribeApiHistoryResponseBody {
	s.RequestParameters = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetResultBodyModel(v string) *DescribeApiHistoryResponseBody {
	s.ResultBodyModel = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetResultDescriptions(v *DescribeApiHistoryResponseBodyResultDescriptions) *DescribeApiHistoryResponseBody {
	s.ResultDescriptions = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetResultSample(v string) *DescribeApiHistoryResponseBody {
	s.ResultSample = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetResultType(v string) *DescribeApiHistoryResponseBody {
	s.ResultType = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetServiceConfig(v *DescribeApiHistoryResponseBodyServiceConfig) *DescribeApiHistoryResponseBody {
	s.ServiceConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetServiceParameters(v *DescribeApiHistoryResponseBodyServiceParameters) *DescribeApiHistoryResponseBody {
	s.ServiceParameters = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetServiceParametersMap(v *DescribeApiHistoryResponseBodyServiceParametersMap) *DescribeApiHistoryResponseBody {
	s.ServiceParametersMap = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetStageName(v string) *DescribeApiHistoryResponseBody {
	s.StageName = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetStatus(v string) *DescribeApiHistoryResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetSystemParameters(v *DescribeApiHistoryResponseBodySystemParameters) *DescribeApiHistoryResponseBody {
	s.SystemParameters = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetVisibility(v string) *DescribeApiHistoryResponseBody {
	s.Visibility = &v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetWebSocketApiType(v string) *DescribeApiHistoryResponseBody {
	s.WebSocketApiType = &v
	return s
}

type DescribeApiHistoryResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeApiHistoryResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyConstantParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyConstantParameters) SetConstantParameter(v []*DescribeApiHistoryResponseBodyConstantParametersConstantParameter) *DescribeApiHistoryResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

type DescribeApiHistoryResponseBodyConstantParametersConstantParameter struct {
	ConstantValue        *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyConstantParametersConstantParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) SetConstantValue(v string) *DescribeApiHistoryResponseBodyConstantParametersConstantParameter {
	s.ConstantValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) SetDescription(v string) *DescribeApiHistoryResponseBodyConstantParametersConstantParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) SetLocation(v string) *DescribeApiHistoryResponseBodyConstantParametersConstantParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) SetServiceParameterName(v string) *DescribeApiHistoryResponseBodyConstantParametersConstantParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiHistoryResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyCustomSystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeApiHistoryResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

type DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) SetDemoValue(v string) *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) SetDescription(v string) *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) SetLocation(v string) *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) SetParameterName(v string) *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) SetServiceParameterName(v string) *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiHistoryResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamples) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeApiHistoryResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

type DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) SetCode(v string) *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Code = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) SetDescription(v string) *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) SetMessage(v string) *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample {
	s.Message = &v
	return s
}

type DescribeApiHistoryResponseBodyOpenIdConnectConfig struct {
	IdTokenParamName *string `json:"IdTokenParamName,omitempty" xml:"IdTokenParamName,omitempty"`
	OpenIdApiType    *string `json:"OpenIdApiType,omitempty" xml:"OpenIdApiType,omitempty"`
	PublicKey        *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	PublicKeyId      *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
}

func (s DescribeApiHistoryResponseBodyOpenIdConnectConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) SetIdTokenParamName(v string) *DescribeApiHistoryResponseBodyOpenIdConnectConfig {
	s.IdTokenParamName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) SetOpenIdApiType(v string) *DescribeApiHistoryResponseBodyOpenIdConnectConfig {
	s.OpenIdApiType = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) SetPublicKey(v string) *DescribeApiHistoryResponseBodyOpenIdConnectConfig {
	s.PublicKey = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) SetPublicKeyId(v string) *DescribeApiHistoryResponseBodyOpenIdConnectConfig {
	s.PublicKeyId = &v
	return s
}

type DescribeApiHistoryResponseBodyRequestConfig struct {
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	BodyModel           *string `json:"BodyModel,omitempty" xml:"BodyModel,omitempty"`
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	RequestHttpMethod   *string `json:"RequestHttpMethod,omitempty" xml:"RequestHttpMethod,omitempty"`
	RequestMode         *string `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	RequestPath         *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	RequestProtocol     *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeApiHistoryResponseBodyRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetBodyModel(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.BodyModel = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetPostBodyDescription(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetRequestHttpMethod(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.RequestHttpMethod = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetRequestMode(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.RequestMode = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetRequestPath(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.RequestPath = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetRequestProtocol(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.RequestProtocol = &v
	return s
}

type DescribeApiHistoryResponseBodyRequestParameters struct {
	RequestParameter []*DescribeApiHistoryResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyRequestParameters) SetRequestParameter(v []*DescribeApiHistoryResponseBodyRequestParametersRequestParameter) *DescribeApiHistoryResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

type DescribeApiHistoryResponseBodyRequestParametersRequestParameter struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	ArrayItemsType    *string `json:"ArrayItemsType,omitempty" xml:"ArrayItemsType,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *int32  `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
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

func (s DescribeApiHistoryResponseBodyRequestParametersRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetApiParameterName(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetArrayItemsType(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.ArrayItemsType = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetDefaultValue(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.DefaultValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetDemoValue(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetDescription(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetDocOrder(v int32) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.DocOrder = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetDocShow(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.DocShow = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetEnumValue(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.EnumValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetJsonScheme(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.JsonScheme = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetLocation(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetMaxLength(v int64) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.MaxLength = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetMaxValue(v int64) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.MaxValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetMinLength(v int64) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.MinLength = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetMinValue(v int64) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.MinValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetParameterType(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetRegularExpression(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.RegularExpression = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) SetRequired(v string) *DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	s.Required = &v
	return s
}

type DescribeApiHistoryResponseBodyResultDescriptions struct {
	ResultDescription []*DescribeApiHistoryResponseBodyResultDescriptionsResultDescription `json:"ResultDescription,omitempty" xml:"ResultDescription,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyResultDescriptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyResultDescriptions) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyResultDescriptions) SetResultDescription(v []*DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) *DescribeApiHistoryResponseBodyResultDescriptions {
	s.ResultDescription = v
	return s
}

type DescribeApiHistoryResponseBodyResultDescriptionsResultDescription struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasChild    *bool   `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Mandatory   *bool   `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetDescription(v string) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetHasChild(v bool) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.HasChild = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetId(v string) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Id = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetKey(v string) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Key = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetMandatory(v bool) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Mandatory = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetName(v string) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Name = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetPid(v string) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Pid = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) SetType(v string) *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	s.Type = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceConfig struct {
	ContentTypeCatagory   *string                                                           `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	ContentTypeValue      *string                                                           `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	FunctionComputeConfig *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	Mock                  *string                                                           `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockHeaders           *DescribeApiHistoryResponseBodyServiceConfigMockHeaders           `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Struct"`
	MockResult            *string                                                           `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	MockStatusCode        *int32                                                            `json:"MockStatusCode,omitempty" xml:"MockStatusCode,omitempty"`
	OssConfig             *DescribeApiHistoryResponseBodyServiceConfigOssConfig             `json:"OssConfig,omitempty" xml:"OssConfig,omitempty" type:"Struct"`
	ServiceAddress        *string                                                           `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceHttpMethod     *string                                                           `json:"ServiceHttpMethod,omitempty" xml:"ServiceHttpMethod,omitempty"`
	ServicePath           *string                                                           `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	ServiceProtocol       *string                                                           `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout        *int32                                                            `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable      *string                                                           `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	VpcConfig             *DescribeApiHistoryResponseBodyServiceConfigVpcConfig             `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
	VpcId                 *string                                                           `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetContentTypeCatagory(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ContentTypeCatagory = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetContentTypeValue(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ContentTypeValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetFunctionComputeConfig(v *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) *DescribeApiHistoryResponseBodyServiceConfig {
	s.FunctionComputeConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetMock(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.Mock = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetMockHeaders(v *DescribeApiHistoryResponseBodyServiceConfigMockHeaders) *DescribeApiHistoryResponseBodyServiceConfig {
	s.MockHeaders = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetMockResult(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.MockResult = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetMockStatusCode(v int32) *DescribeApiHistoryResponseBodyServiceConfig {
	s.MockStatusCode = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetOssConfig(v *DescribeApiHistoryResponseBodyServiceConfigOssConfig) *DescribeApiHistoryResponseBodyServiceConfig {
	s.OssConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetServiceAddress(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ServiceAddress = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetServiceHttpMethod(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ServiceHttpMethod = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetServicePath(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ServicePath = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetServiceProtocol(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ServiceProtocol = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetServiceTimeout(v int32) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetServiceVpcEnable(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ServiceVpcEnable = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetVpcConfig(v *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) *DescribeApiHistoryResponseBodyServiceConfig {
	s.VpcConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetVpcId(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.VpcId = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig struct {
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	ContentTypeValue    *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	FcBaseUrl           *string `json:"FcBaseUrl,omitempty" xml:"FcBaseUrl,omitempty"`
	FcType              *string `json:"FcType,omitempty" xml:"FcType,omitempty"`
	FunctionName        *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Method              *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OnlyBusinessPath    *bool   `json:"OnlyBusinessPath,omitempty" xml:"OnlyBusinessPath,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Qualifier           *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn             *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName         *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetContentTypeCatagory(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.ContentTypeCatagory = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetContentTypeValue(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.ContentTypeValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetFcBaseUrl(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.FcBaseUrl = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetFcType(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.FcType = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetFunctionName(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetMethod(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.Method = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetOnlyBusinessPath(v bool) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.OnlyBusinessPath = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetPath(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.Path = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetQualifier(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.Qualifier = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetRegionId(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.RegionId = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetRoleArn(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) SetServiceName(v string) *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	s.ServiceName = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceConfigMockHeaders struct {
	MockHeader []*DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader `json:"MockHeader,omitempty" xml:"MockHeader,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeaders) SetMockHeader(v []*DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) *DescribeApiHistoryResponseBodyServiceConfigMockHeaders {
	s.MockHeader = v
	return s
}

type DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader struct {
	HeaderName  *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderName(v string) *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderValue(v string) *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderValue = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceConfigOssConfig struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	BucketName  *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigOssConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigOssConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) SetAction(v string) *DescribeApiHistoryResponseBodyServiceConfigOssConfig {
	s.Action = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) SetBucketName(v string) *DescribeApiHistoryResponseBodyServiceConfigOssConfig {
	s.BucketName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) SetKey(v string) *DescribeApiHistoryResponseBodyServiceConfigOssConfig {
	s.Key = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) SetOssRegionId(v string) *DescribeApiHistoryResponseBodyServiceConfigOssConfig {
	s.OssRegionId = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceConfigVpcConfig struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcScheme  *string `json:"VpcScheme,omitempty" xml:"VpcScheme,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) SetInstanceId(v string) *DescribeApiHistoryResponseBodyServiceConfigVpcConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) SetName(v string) *DescribeApiHistoryResponseBodyServiceConfigVpcConfig {
	s.Name = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) SetPort(v int32) *DescribeApiHistoryResponseBodyServiceConfigVpcConfig {
	s.Port = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) SetVpcId(v string) *DescribeApiHistoryResponseBodyServiceConfigVpcConfig {
	s.VpcId = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) SetVpcScheme(v string) *DescribeApiHistoryResponseBodyServiceConfigVpcConfig {
	s.VpcScheme = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceParameters struct {
	ServiceParameter []*DescribeApiHistoryResponseBodyServiceParametersServiceParameter `json:"ServiceParameter,omitempty" xml:"ServiceParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyServiceParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParameters) SetServiceParameter(v []*DescribeApiHistoryResponseBodyServiceParametersServiceParameter) *DescribeApiHistoryResponseBodyServiceParameters {
	s.ServiceParameter = v
	return s
}

type DescribeApiHistoryResponseBodyServiceParametersServiceParameter struct {
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceParametersServiceParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParametersServiceParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) SetLocation(v string) *DescribeApiHistoryResponseBodyServiceParametersServiceParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) SetParameterType(v string) *DescribeApiHistoryResponseBodyServiceParametersServiceParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) SetServiceParameterName(v string) *DescribeApiHistoryResponseBodyServiceParametersServiceParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiHistoryResponseBodyServiceParametersMap struct {
	ServiceParameterMap []*DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyServiceParametersMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParametersMap) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMap) SetServiceParameterMap(v []*DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) *DescribeApiHistoryResponseBodyServiceParametersMap {
	s.ServiceParameterMap = v
	return s
}

type DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap struct {
	RequestParameterName *string `json:"RequestParameterName,omitempty" xml:"RequestParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) SetRequestParameterName(v string) *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) SetServiceParameterName(v string) *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiHistoryResponseBodySystemParameters struct {
	SystemParameter []*DescribeApiHistoryResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodySystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodySystemParameters) SetSystemParameter(v []*DescribeApiHistoryResponseBodySystemParametersSystemParameter) *DescribeApiHistoryResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

type DescribeApiHistoryResponseBodySystemParametersSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodySystemParametersSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) SetDemoValue(v string) *DescribeApiHistoryResponseBodySystemParametersSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) SetDescription(v string) *DescribeApiHistoryResponseBodySystemParametersSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) SetLocation(v string) *DescribeApiHistoryResponseBodySystemParametersSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) SetParameterName(v string) *DescribeApiHistoryResponseBodySystemParametersSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) SetServiceParameterName(v string) *DescribeApiHistoryResponseBodySystemParametersSystemParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeApiHistoryResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponse) SetHeaders(v map[string]*string) *DescribeApiHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiHistoryResponse) SetBody(v *DescribeApiHistoryResponseBody) *DescribeApiHistoryResponse {
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
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiIpControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeApiIpControlsResponse) SetBody(v *DescribeApiIpControlsResponseBody) *DescribeApiIpControlsResponse {
	s.Body = v
	return s
}

type DescribeApiLatencyDataRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiLatencyDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataRequest) SetApiId(v string) *DescribeApiLatencyDataRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetEndTime(v string) *DescribeApiLatencyDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetGroupId(v string) *DescribeApiLatencyDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetSecurityToken(v string) *DescribeApiLatencyDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetStageName(v string) *DescribeApiLatencyDataRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiLatencyDataRequest) SetStartTime(v string) *DescribeApiLatencyDataRequest {
	s.StartTime = &v
	return s
}

type DescribeApiLatencyDataResponseBody struct {
	CallLatencys *DescribeApiLatencyDataResponseBodyCallLatencys `json:"CallLatencys,omitempty" xml:"CallLatencys,omitempty" type:"Struct"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiLatencyDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponseBody) SetCallLatencys(v *DescribeApiLatencyDataResponseBodyCallLatencys) *DescribeApiLatencyDataResponseBody {
	s.CallLatencys = v
	return s
}

func (s *DescribeApiLatencyDataResponseBody) SetRequestId(v string) *DescribeApiLatencyDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiLatencyDataResponseBodyCallLatencys struct {
	MonitorItem []*DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiLatencyDataResponseBodyCallLatencys) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyDataResponseBodyCallLatencys) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencys) SetMonitorItem(v []*DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) *DescribeApiLatencyDataResponseBodyCallLatencys {
	s.MonitorItem = v
	return s
}

type DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem struct {
	ItemTime  *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) SetItemTime(v string) *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem) SetItemValue(v string) *DescribeApiLatencyDataResponseBodyCallLatencysMonitorItem {
	s.ItemValue = &v
	return s
}

type DescribeApiLatencyDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiLatencyDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiLatencyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiLatencyDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiLatencyDataResponse) SetHeaders(v map[string]*string) *DescribeApiLatencyDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiLatencyDataResponse) SetBody(v *DescribeApiLatencyDataResponseBody) *DescribeApiLatencyDataResponse {
	s.Body = v
	return s
}

type DescribeApiMarketAttributesRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApiMarketAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMarketAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketAttributesRequest) SetApiId(v string) *DescribeApiMarketAttributesRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiMarketAttributesRequest) SetGroupId(v string) *DescribeApiMarketAttributesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiMarketAttributesRequest) SetSecurityToken(v string) *DescribeApiMarketAttributesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApiMarketAttributesResponseBody struct {
	ApiId              *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	MarketChargingMode *string `json:"MarketChargingMode,omitempty" xml:"MarketChargingMode,omitempty"`
	NeedCharging       *string `json:"NeedCharging,omitempty" xml:"NeedCharging,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiMarketAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMarketAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketAttributesResponseBody) SetApiId(v string) *DescribeApiMarketAttributesResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) SetMarketChargingMode(v string) *DescribeApiMarketAttributesResponseBody {
	s.MarketChargingMode = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) SetNeedCharging(v string) *DescribeApiMarketAttributesResponseBody {
	s.NeedCharging = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) SetRequestId(v string) *DescribeApiMarketAttributesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiMarketAttributesResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiMarketAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiMarketAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiMarketAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketAttributesResponse) SetHeaders(v map[string]*string) *DescribeApiMarketAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiMarketAttributesResponse) SetBody(v *DescribeApiMarketAttributesResponseBody) *DescribeApiMarketAttributesResponse {
	s.Body = v
	return s
}

type DescribeApiQpsDataRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiQpsDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataRequest) SetApiId(v string) *DescribeApiQpsDataRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetEndTime(v string) *DescribeApiQpsDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetGroupId(v string) *DescribeApiQpsDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetSecurityToken(v string) *DescribeApiQpsDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetStageName(v string) *DescribeApiQpsDataRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiQpsDataRequest) SetStartTime(v string) *DescribeApiQpsDataRequest {
	s.StartTime = &v
	return s
}

type DescribeApiQpsDataResponseBody struct {
	CallFails     *DescribeApiQpsDataResponseBodyCallFails     `json:"CallFails,omitempty" xml:"CallFails,omitempty" type:"Struct"`
	CallSuccesses *DescribeApiQpsDataResponseBodyCallSuccesses `json:"CallSuccesses,omitempty" xml:"CallSuccesses,omitempty" type:"Struct"`
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiQpsDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBody) SetCallFails(v *DescribeApiQpsDataResponseBodyCallFails) *DescribeApiQpsDataResponseBody {
	s.CallFails = v
	return s
}

func (s *DescribeApiQpsDataResponseBody) SetCallSuccesses(v *DescribeApiQpsDataResponseBodyCallSuccesses) *DescribeApiQpsDataResponseBody {
	s.CallSuccesses = v
	return s
}

func (s *DescribeApiQpsDataResponseBody) SetRequestId(v string) *DescribeApiQpsDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiQpsDataResponseBodyCallFails struct {
	MonitorItem []*DescribeApiQpsDataResponseBodyCallFailsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiQpsDataResponseBodyCallFails) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallFails) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallFails) SetMonitorItem(v []*DescribeApiQpsDataResponseBodyCallFailsMonitorItem) *DescribeApiQpsDataResponseBodyCallFails {
	s.MonitorItem = v
	return s
}

type DescribeApiQpsDataResponseBodyCallFailsMonitorItem struct {
	ItemTime  *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiQpsDataResponseBodyCallFailsMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallFailsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) SetItemTime(v string) *DescribeApiQpsDataResponseBodyCallFailsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallFailsMonitorItem) SetItemValue(v string) *DescribeApiQpsDataResponseBodyCallFailsMonitorItem {
	s.ItemValue = &v
	return s
}

type DescribeApiQpsDataResponseBodyCallSuccesses struct {
	MonitorItem []*DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiQpsDataResponseBodyCallSuccesses) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallSuccesses) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallSuccesses) SetMonitorItem(v []*DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) *DescribeApiQpsDataResponseBodyCallSuccesses {
	s.MonitorItem = v
	return s
}

type DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem struct {
	ItemTime  *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) SetItemTime(v string) *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem) SetItemValue(v string) *DescribeApiQpsDataResponseBodyCallSuccessesMonitorItem {
	s.ItemValue = &v
	return s
}

type DescribeApiQpsDataResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiQpsDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiQpsDataResponse) SetHeaders(v map[string]*string) *DescribeApiQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiQpsDataResponse) SetBody(v *DescribeApiQpsDataResponseBody) *DescribeApiQpsDataResponse {
	s.Body = v
	return s
}

type DescribeApiSignaturesRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiSignaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiSignaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesRequest) SetApiIds(v string) *DescribeApiSignaturesRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetGroupId(v string) *DescribeApiSignaturesRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetPageNumber(v int32) *DescribeApiSignaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetPageSize(v int32) *DescribeApiSignaturesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetSecurityToken(v string) *DescribeApiSignaturesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiSignaturesRequest) SetStageName(v string) *DescribeApiSignaturesRequest {
	s.StageName = &v
	return s
}

type DescribeApiSignaturesResponseBody struct {
	ApiSignatures *DescribeApiSignaturesResponseBodyApiSignatures `json:"ApiSignatures,omitempty" xml:"ApiSignatures,omitempty" type:"Struct"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiSignaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponseBody) SetApiSignatures(v *DescribeApiSignaturesResponseBodyApiSignatures) *DescribeApiSignaturesResponseBody {
	s.ApiSignatures = v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetPageNumber(v int32) *DescribeApiSignaturesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetPageSize(v int32) *DescribeApiSignaturesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetRequestId(v string) *DescribeApiSignaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiSignaturesResponseBody) SetTotalCount(v int32) *DescribeApiSignaturesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiSignaturesResponseBodyApiSignatures struct {
	ApiSignatureItem []*DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem `json:"ApiSignatureItem,omitempty" xml:"ApiSignatureItem,omitempty" type:"Repeated"`
}

func (s DescribeApiSignaturesResponseBodyApiSignatures) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiSignaturesResponseBodyApiSignatures) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponseBodyApiSignatures) SetApiSignatureItem(v []*DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) *DescribeApiSignaturesResponseBodyApiSignatures {
	s.ApiSignatureItem = v
	return s
}

type DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BoundTime     *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetApiId(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetApiName(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetBoundTime(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetSignatureId(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.SignatureId = &v
	return s
}

func (s *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem) SetSignatureName(v string) *DescribeApiSignaturesResponseBodyApiSignaturesApiSignatureItem {
	s.SignatureName = &v
	return s
}

type DescribeApiSignaturesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiSignaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiSignaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiSignaturesResponse) SetHeaders(v map[string]*string) *DescribeApiSignaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiSignaturesResponse) SetBody(v *DescribeApiSignaturesResponseBody) *DescribeApiSignaturesResponse {
	s.Body = v
	return s
}

type DescribeApiTrafficControlsRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiTrafficControlsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficControlsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsRequest) SetApiIds(v string) *DescribeApiTrafficControlsRequest {
	s.ApiIds = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetGroupId(v string) *DescribeApiTrafficControlsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetPageNumber(v int32) *DescribeApiTrafficControlsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetPageSize(v int32) *DescribeApiTrafficControlsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetSecurityToken(v string) *DescribeApiTrafficControlsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiTrafficControlsRequest) SetStageName(v string) *DescribeApiTrafficControlsRequest {
	s.StageName = &v
	return s
}

type DescribeApiTrafficControlsResponseBody struct {
	ApiTrafficControls *DescribeApiTrafficControlsResponseBodyApiTrafficControls `json:"ApiTrafficControls,omitempty" xml:"ApiTrafficControls,omitempty" type:"Struct"`
	PageNumber         *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount         *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApiTrafficControlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficControlsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponseBody) SetApiTrafficControls(v *DescribeApiTrafficControlsResponseBodyApiTrafficControls) *DescribeApiTrafficControlsResponseBody {
	s.ApiTrafficControls = v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetPageNumber(v int32) *DescribeApiTrafficControlsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetPageSize(v int32) *DescribeApiTrafficControlsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetRequestId(v string) *DescribeApiTrafficControlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBody) SetTotalCount(v int32) *DescribeApiTrafficControlsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApiTrafficControlsResponseBodyApiTrafficControls struct {
	ApiTrafficControlItem []*DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem `json:"ApiTrafficControlItem,omitempty" xml:"ApiTrafficControlItem,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControls) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControls) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControls) SetApiTrafficControlItem(v []*DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) *DescribeApiTrafficControlsResponseBodyApiTrafficControls {
	s.ApiTrafficControlItem = v
	return s
}

type DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem struct {
	ApiId              *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName            *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	BoundTime          *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	TrafficControlId   *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
	TrafficControlName *string `json:"TrafficControlName,omitempty" xml:"TrafficControlName,omitempty"`
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetApiId(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.ApiId = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetApiName(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.ApiName = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetBoundTime(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetTrafficControlId(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.TrafficControlId = &v
	return s
}

func (s *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem) SetTrafficControlName(v string) *DescribeApiTrafficControlsResponseBodyApiTrafficControlsApiTrafficControlItem {
	s.TrafficControlName = &v
	return s
}

type DescribeApiTrafficControlsResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiTrafficControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiTrafficControlsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficControlsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficControlsResponse) SetHeaders(v map[string]*string) *DescribeApiTrafficControlsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiTrafficControlsResponse) SetBody(v *DescribeApiTrafficControlsResponseBody) *DescribeApiTrafficControlsResponse {
	s.Body = v
	return s
}

type DescribeApiTrafficDataRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeApiTrafficDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataRequest) SetApiId(v string) *DescribeApiTrafficDataRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetEndTime(v string) *DescribeApiTrafficDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetGroupId(v string) *DescribeApiTrafficDataRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetSecurityToken(v string) *DescribeApiTrafficDataRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetStageName(v string) *DescribeApiTrafficDataRequest {
	s.StageName = &v
	return s
}

func (s *DescribeApiTrafficDataRequest) SetStartTime(v string) *DescribeApiTrafficDataRequest {
	s.StartTime = &v
	return s
}

type DescribeApiTrafficDataResponseBody struct {
	CallDownloads *DescribeApiTrafficDataResponseBodyCallDownloads `json:"CallDownloads,omitempty" xml:"CallDownloads,omitempty" type:"Struct"`
	CallUploads   *DescribeApiTrafficDataResponseBodyCallUploads   `json:"CallUploads,omitempty" xml:"CallUploads,omitempty" type:"Struct"`
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiTrafficDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBody) SetCallDownloads(v *DescribeApiTrafficDataResponseBodyCallDownloads) *DescribeApiTrafficDataResponseBody {
	s.CallDownloads = v
	return s
}

func (s *DescribeApiTrafficDataResponseBody) SetCallUploads(v *DescribeApiTrafficDataResponseBodyCallUploads) *DescribeApiTrafficDataResponseBody {
	s.CallUploads = v
	return s
}

func (s *DescribeApiTrafficDataResponseBody) SetRequestId(v string) *DescribeApiTrafficDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApiTrafficDataResponseBodyCallDownloads struct {
	MonitorItem []*DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficDataResponseBodyCallDownloads) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallDownloads) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloads) SetMonitorItem(v []*DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) *DescribeApiTrafficDataResponseBodyCallDownloads {
	s.MonitorItem = v
	return s
}

type DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem struct {
	ItemTime  *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) SetItemTime(v string) *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem) SetItemValue(v string) *DescribeApiTrafficDataResponseBodyCallDownloadsMonitorItem {
	s.ItemValue = &v
	return s
}

type DescribeApiTrafficDataResponseBodyCallUploads struct {
	MonitorItem []*DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem `json:"MonitorItem,omitempty" xml:"MonitorItem,omitempty" type:"Repeated"`
}

func (s DescribeApiTrafficDataResponseBodyCallUploads) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallUploads) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallUploads) SetMonitorItem(v []*DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) *DescribeApiTrafficDataResponseBodyCallUploads {
	s.MonitorItem = v
	return s
}

type DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem struct {
	ItemTime  *string `json:"ItemTime,omitempty" xml:"ItemTime,omitempty"`
	ItemValue *string `json:"ItemValue,omitempty" xml:"ItemValue,omitempty"`
}

func (s DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) SetItemTime(v string) *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem {
	s.ItemTime = &v
	return s
}

func (s *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem) SetItemValue(v string) *DescribeApiTrafficDataResponseBodyCallUploadsMonitorItem {
	s.ItemValue = &v
	return s
}

type DescribeApiTrafficDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApiTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApiTrafficDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeApiTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiTrafficDataResponse) SetBody(v *DescribeApiTrafficDataResponseBody) *DescribeApiTrafficDataResponse {
	s.Body = v
	return s
}

type DescribeApisRequest struct {
	ApiId         *string                   `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string                   `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	CatalogId     *string                   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	EnableTagAuth *bool                     `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	GroupId       *string                   `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DescribeApisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	Visibility    *string                   `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
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

func (s *DescribeApisRequest) SetCatalogId(v string) *DescribeApisRequest {
	s.CatalogId = &v
	return s
}

func (s *DescribeApisRequest) SetEnableTagAuth(v bool) *DescribeApisRequest {
	s.EnableTagAuth = &v
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

func (s *DescribeApisRequest) SetTag(v []*DescribeApisRequestTag) *DescribeApisRequest {
	s.Tag = v
	return s
}

func (s *DescribeApisRequest) SetVisibility(v string) *DescribeApisRequest {
	s.Visibility = &v
	return s
}

type DescribeApisRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApisRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeApisRequestTag) SetKey(v string) *DescribeApisRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeApisRequestTag) SetValue(v string) *DescribeApisRequestTag {
	s.Value = &v
	return s
}

type DescribeApisResponseBody struct {
	ApiSummarys *DescribeApisResponseBodyApiSummarys `json:"ApiSummarys,omitempty" xml:"ApiSummarys,omitempty" type:"Struct"`
	PageNumber  *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBody) SetApiSummarys(v *DescribeApisResponseBodyApiSummarys) *DescribeApisResponseBody {
	s.ApiSummarys = v
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

type DescribeApisResponseBodyApiSummarys struct {
	ApiSummary []*DescribeApisResponseBodyApiSummarysApiSummary `json:"ApiSummary,omitempty" xml:"ApiSummary,omitempty" type:"Repeated"`
}

func (s DescribeApisResponseBodyApiSummarys) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarys) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarys) SetApiSummary(v []*DescribeApisResponseBodyApiSummarysApiSummary) *DescribeApisResponseBodyApiSummarys {
	s.ApiSummary = v
	return s
}

type DescribeApisResponseBodyApiSummarysApiSummary struct {
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

func (s DescribeApisResponseBodyApiSummarysApiSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisResponseBodyApiSummarysApiSummary) GoString() string {
	return s.String()
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetApiId(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ApiId = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetApiName(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ApiName = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetCreatedTime(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.CreatedTime = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetDescription(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.Description = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetGroupId(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.GroupId = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetGroupName(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.GroupName = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetModifiedTime(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetRegionId(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.RegionId = &v
	return s
}

func (s *DescribeApisResponseBodyApiSummarysApiSummary) SetVisibility(v string) *DescribeApisResponseBodyApiSummarysApiSummary {
	s.Visibility = &v
	return s
}

type DescribeApisResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeApisResponse) SetBody(v *DescribeApisResponseBody) *DescribeApisResponse {
	s.Body = v
	return s
}

type DescribeApisByAppRequest struct {
	// API名称
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// API的ID
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// APP的ID
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// API的请求HTTP Method
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// 当前页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页条目
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// API请求路径
	Path          *string `json:"Path,omitempty" xml:"Path,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeApisByAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByAppRequest) SetApiName(v string) *DescribeApisByAppRequest {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByAppRequest) SetApiUid(v string) *DescribeApisByAppRequest {
	s.ApiUid = &v
	return s
}

func (s *DescribeApisByAppRequest) SetAppId(v int64) *DescribeApisByAppRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApisByAppRequest) SetMethod(v string) *DescribeApisByAppRequest {
	s.Method = &v
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

func (s *DescribeApisByAppRequest) SetPath(v string) *DescribeApisByAppRequest {
	s.Path = &v
	return s
}

func (s *DescribeApisByAppRequest) SetSecurityToken(v string) *DescribeApisByAppRequest {
	s.SecurityToken = &v
	return s
}

type DescribeApisByAppResponseBody struct {
	AppApiRelationInfos *DescribeApisByAppResponseBodyAppApiRelationInfos `json:"AppApiRelationInfos,omitempty" xml:"AppApiRelationInfos,omitempty" type:"Struct"`
	// 当前页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页条目
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总条目数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// API的ID
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// API名称
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// 授权有效时间
	AuthVaildTime *string `json:"AuthVaildTime,omitempty" xml:"AuthVaildTime,omitempty"`
	// 授权来源
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	// 授权时间
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 分组ID
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// 分组名称
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// API的请求HTTP Method
	Method   *string `json:"Method,omitempty" xml:"Method,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// API的请求路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 地区ID
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 环境名称
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
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

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetAuthVaildTime(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.AuthVaildTime = &v
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

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetMethod(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Method = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetOperator(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Operator = &v
	return s
}

func (s *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo) SetPath(v string) *DescribeApisByAppResponseBodyAppApiRelationInfosAppApiRelationInfo {
	s.Path = &v
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApisByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApisByIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeApisByIpControlResponse) SetBody(v *DescribeApisByIpControlResponseBody) *DescribeApisByIpControlResponse {
	s.Body = v
	return s
}

type DescribeApisBySignatureRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
}

func (s DescribeApisBySignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisBySignatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureRequest) SetPageNumber(v int32) *DescribeApisBySignatureRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisBySignatureRequest) SetPageSize(v int32) *DescribeApisBySignatureRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisBySignatureRequest) SetSecurityToken(v string) *DescribeApisBySignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisBySignatureRequest) SetSignatureId(v string) *DescribeApisBySignatureRequest {
	s.SignatureId = &v
	return s
}

type DescribeApisBySignatureResponseBody struct {
	ApiInfos   *DescribeApisBySignatureResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisBySignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisBySignatureResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponseBody) SetApiInfos(v *DescribeApisBySignatureResponseBodyApiInfos) *DescribeApisBySignatureResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetPageNumber(v int32) *DescribeApisBySignatureResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetPageSize(v int32) *DescribeApisBySignatureResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetRequestId(v string) *DescribeApisBySignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBody) SetTotalCount(v int32) *DescribeApisBySignatureResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisBySignatureResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisBySignatureResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisBySignatureResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisBySignatureResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponseBodyApiInfos) SetApiInfo(v []*DescribeApisBySignatureResponseBodyApiInfosApiInfo) *DescribeApisBySignatureResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApisBySignatureResponseBodyApiInfosApiInfo struct {
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

func (s DescribeApisBySignatureResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisBySignatureResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetBoundTime(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.BoundTime = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisBySignatureResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisBySignatureResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeApisBySignatureResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApisBySignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApisBySignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisBySignatureResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisBySignatureResponse) SetHeaders(v map[string]*string) *DescribeApisBySignatureResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisBySignatureResponse) SetBody(v *DescribeApisBySignatureResponseBody) *DescribeApisBySignatureResponse {
	s.Body = v
	return s
}

type DescribeApisByTrafficControlRequest struct {
	PageNumber       *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s DescribeApisByTrafficControlRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByTrafficControlRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlRequest) SetPageNumber(v int32) *DescribeApisByTrafficControlRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) SetPageSize(v int32) *DescribeApisByTrafficControlRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) SetSecurityToken(v string) *DescribeApisByTrafficControlRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeApisByTrafficControlRequest) SetTrafficControlId(v string) *DescribeApisByTrafficControlRequest {
	s.TrafficControlId = &v
	return s
}

type DescribeApisByTrafficControlResponseBody struct {
	ApiInfos   *DescribeApisByTrafficControlResponseBodyApiInfos `json:"ApiInfos,omitempty" xml:"ApiInfos,omitempty" type:"Struct"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeApisByTrafficControlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByTrafficControlResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponseBody) SetApiInfos(v *DescribeApisByTrafficControlResponseBodyApiInfos) *DescribeApisByTrafficControlResponseBody {
	s.ApiInfos = v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetPageNumber(v int32) *DescribeApisByTrafficControlResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetPageSize(v int32) *DescribeApisByTrafficControlResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetRequestId(v string) *DescribeApisByTrafficControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBody) SetTotalCount(v int32) *DescribeApisByTrafficControlResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeApisByTrafficControlResponseBodyApiInfos struct {
	ApiInfo []*DescribeApisByTrafficControlResponseBodyApiInfosApiInfo `json:"ApiInfo,omitempty" xml:"ApiInfo,omitempty" type:"Repeated"`
}

func (s DescribeApisByTrafficControlResponseBodyApiInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByTrafficControlResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfos) SetApiInfo(v []*DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) *DescribeApisByTrafficControlResponseBodyApiInfos {
	s.ApiInfo = v
	return s
}

type DescribeApisByTrafficControlResponseBodyApiInfosApiInfo struct {
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

func (s DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetApiId(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.ApiId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetApiName(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.ApiName = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetBoundTime(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.BoundTime = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetDescription(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.Description = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetGroupId(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.GroupId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetGroupName(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.GroupName = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetRegionId(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetStageName(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.StageName = &v
	return s
}

func (s *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo) SetVisibility(v string) *DescribeApisByTrafficControlResponseBodyApiInfosApiInfo {
	s.Visibility = &v
	return s
}

type DescribeApisByTrafficControlResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeApisByTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeApisByTrafficControlResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApisByTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisByTrafficControlResponse) SetHeaders(v map[string]*string) *DescribeApisByTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisByTrafficControlResponse) SetBody(v *DescribeApisByTrafficControlResponseBody) *DescribeApisByTrafficControlResponse {
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
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAppResponse) SetBody(v *DescribeAppResponseBody) *DescribeAppResponse {
	s.Body = v
	return s
}

type DescribeAppAttributesRequest struct {
	AppCode       *string                            `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppId         *int64                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppKey        *string                            `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName       *string                            `json:"AppName,omitempty" xml:"AppName,omitempty"`
	EnableTagAuth *bool                              `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	PageNumber    *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string                            `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Sort          *string                            `json:"Sort,omitempty" xml:"Sort,omitempty"`
	Tag           []*DescribeAppAttributesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAppAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesRequest) SetAppCode(v string) *DescribeAppAttributesRequest {
	s.AppCode = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetAppId(v int64) *DescribeAppAttributesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetAppKey(v string) *DescribeAppAttributesRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetAppName(v string) *DescribeAppAttributesRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetEnableTagAuth(v bool) *DescribeAppAttributesRequest {
	s.EnableTagAuth = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetPageNumber(v int32) *DescribeAppAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetPageSize(v int32) *DescribeAppAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetSecurityToken(v string) *DescribeAppAttributesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetSort(v string) *DescribeAppAttributesRequest {
	s.Sort = &v
	return s
}

func (s *DescribeAppAttributesRequest) SetTag(v []*DescribeAppAttributesRequestTag) *DescribeAppAttributesRequest {
	s.Tag = v
	return s
}

type DescribeAppAttributesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppAttributesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesRequestTag) SetKey(v string) *DescribeAppAttributesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAppAttributesRequestTag) SetValue(v string) *DescribeAppAttributesRequestTag {
	s.Value = &v
	return s
}

type DescribeAppAttributesResponseBody struct {
	Apps       *DescribeAppAttributesResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Struct"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBody) SetApps(v *DescribeAppAttributesResponseBodyApps) *DescribeAppAttributesResponseBody {
	s.Apps = v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetPageNumber(v int32) *DescribeAppAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetPageSize(v int32) *DescribeAppAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetRequestId(v string) *DescribeAppAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetTotalCount(v int32) *DescribeAppAttributesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAppAttributesResponseBodyApps struct {
	AppAttribute []*DescribeAppAttributesResponseBodyAppsAppAttribute `json:"AppAttribute,omitempty" xml:"AppAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAppAttributesResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyApps) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyApps) SetAppAttribute(v []*DescribeAppAttributesResponseBodyAppsAppAttribute) *DescribeAppAttributesResponseBodyApps {
	s.AppAttribute = v
	return s
}

type DescribeAppAttributesResponseBodyAppsAppAttribute struct {
	AppId        *int64                                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName      *string                                                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CreatedTime  *string                                                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Tags         *DescribeAppAttributesResponseBodyAppsAppAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeAppAttributesResponseBodyAppsAppAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyAppsAppAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetAppId(v int64) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.AppId = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetAppName(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.AppName = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetCreatedTime(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetDescription(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.Description = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetModifiedTime(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetTags(v *DescribeAppAttributesResponseBodyAppsAppAttributeTags) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.Tags = v
	return s
}

type DescribeAppAttributesResponseBodyAppsAppAttributeTags struct {
	TagInfo []*DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTags) SetTagInfo(v []*DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) *DescribeAppAttributesResponseBodyAppsAppAttributeTags {
	s.TagInfo = v
	return s
}

type DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) SetKey(v string) *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) SetValue(v string) *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo {
	s.Value = &v
	return s
}

type DescribeAppAttributesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppAttributesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponse) SetHeaders(v map[string]*string) *DescribeAppAttributesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppAttributesResponse) SetBody(v *DescribeAppAttributesResponseBody) *DescribeAppAttributesResponse {
	s.Body = v
	return s
}

type DescribeAppSecurityRequest struct {
	AppId         *int64                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string                          `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DescribeAppSecurityRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAppSecurityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityRequest) SetAppId(v int64) *DescribeAppSecurityRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppSecurityRequest) SetSecurityToken(v string) *DescribeAppSecurityRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppSecurityRequest) SetTag(v []*DescribeAppSecurityRequestTag) *DescribeAppSecurityRequest {
	s.Tag = v
	return s
}

type DescribeAppSecurityRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppSecurityRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppSecurityRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityRequestTag) SetKey(v string) *DescribeAppSecurityRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAppSecurityRequestTag) SetValue(v string) *DescribeAppSecurityRequestTag {
	s.Value = &v
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppSecurityResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAppSecurityResponse) SetBody(v *DescribeAppSecurityResponseBody) *DescribeAppSecurityResponse {
	s.Body = v
	return s
}

type DescribeAppsRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppOwner      *int64  `json:"AppOwner,omitempty" xml:"AppOwner,omitempty"`
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

func (s *DescribeAppsRequest) SetAppOwner(v int64) *DescribeAppsRequest {
	s.AppOwner = &v
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
	AppItem []*DescribeAppsResponseBodyAppsAppItem `json:"AppItem,omitempty" xml:"AppItem,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyApps) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyApps) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyApps) SetAppItem(v []*DescribeAppsResponseBodyAppsAppItem) *DescribeAppsResponseBodyApps {
	s.AppItem = v
	return s
}

type DescribeAppsResponseBodyAppsAppItem struct {
	AppId       *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName     *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeAppsResponseBodyAppsAppItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyAppsAppItem) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppsAppItem) SetAppId(v int64) *DescribeAppsResponseBodyAppsAppItem {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsAppItem) SetAppName(v string) *DescribeAppsResponseBodyAppsAppItem {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppsAppItem) SetDescription(v string) *DescribeAppsResponseBodyAppsAppItem {
	s.Description = &v
	return s
}

type DescribeAppsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeAuthorizedApisRequest struct {
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAuthorizedApisRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedApisRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisRequest) SetAppId(v int64) *DescribeAuthorizedApisRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAuthorizedApisRequest) SetPageNumber(v int32) *DescribeAuthorizedApisRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuthorizedApisRequest) SetPageSize(v int32) *DescribeAuthorizedApisRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthorizedApisRequest) SetSecurityToken(v string) *DescribeAuthorizedApisRequest {
	s.SecurityToken = &v
	return s
}

type DescribeAuthorizedApisResponseBody struct {
	AuthorizedApis *DescribeAuthorizedApisResponseBodyAuthorizedApis `json:"AuthorizedApis,omitempty" xml:"AuthorizedApis,omitempty" type:"Struct"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuthorizedApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponseBody) SetAuthorizedApis(v *DescribeAuthorizedApisResponseBodyAuthorizedApis) *DescribeAuthorizedApisResponseBody {
	s.AuthorizedApis = v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetPageNumber(v int32) *DescribeAuthorizedApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetPageSize(v int32) *DescribeAuthorizedApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetRequestId(v string) *DescribeAuthorizedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBody) SetTotalCount(v int32) *DescribeAuthorizedApisResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuthorizedApisResponseBodyAuthorizedApis struct {
	AuthorizedApi []*DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi `json:"AuthorizedApi,omitempty" xml:"AuthorizedApi,omitempty" type:"Repeated"`
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApis) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApis) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApis) SetAuthorizedApi(v []*DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) *DescribeAuthorizedApisResponseBodyAuthorizedApis {
	s.AuthorizedApi = v
	return s
}

type DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi struct {
	ApiId               *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName             *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthVaildTime       *string `json:"AuthVaildTime,omitempty" xml:"AuthVaildTime,omitempty"`
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	AuthorizedTime      *string `json:"AuthorizedTime,omitempty" xml:"AuthorizedTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId             *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName           *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Operator            *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName           *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetApiId(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.ApiId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetApiName(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.ApiName = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetAuthVaildTime(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.AuthVaildTime = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetAuthorizationSource(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetAuthorizedTime(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.AuthorizedTime = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetDescription(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.Description = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetGroupId(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.GroupId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetGroupName(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.GroupName = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetOperator(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.Operator = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetRegionId(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.RegionId = &v
	return s
}

func (s *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi) SetStageName(v string) *DescribeAuthorizedApisResponseBodyAuthorizedApisAuthorizedApi {
	s.StageName = &v
	return s
}

type DescribeAuthorizedApisResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuthorizedApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuthorizedApisResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedApisResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedApisResponse) SetHeaders(v map[string]*string) *DescribeAuthorizedApisResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthorizedApisResponse) SetBody(v *DescribeAuthorizedApisResponseBody) *DescribeAuthorizedApisResponse {
	s.Body = v
	return s
}

type DescribeAuthorizedAppsRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppOwnerId    *int64  `json:"AppOwnerId,omitempty" xml:"AppOwnerId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAuthorizedAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedAppsRequest) SetApiId(v string) *DescribeAuthorizedAppsRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetAppId(v int64) *DescribeAuthorizedAppsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetAppName(v string) *DescribeAuthorizedAppsRequest {
	s.AppName = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetAppOwnerId(v int64) *DescribeAuthorizedAppsRequest {
	s.AppOwnerId = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetGroupId(v string) *DescribeAuthorizedAppsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetPageNumber(v int32) *DescribeAuthorizedAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetPageSize(v int32) *DescribeAuthorizedAppsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetSecurityToken(v string) *DescribeAuthorizedAppsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuthorizedAppsRequest) SetStageName(v string) *DescribeAuthorizedAppsRequest {
	s.StageName = &v
	return s
}

type DescribeAuthorizedAppsResponseBody struct {
	AuthorizedApps *DescribeAuthorizedAppsResponseBodyAuthorizedApps `json:"AuthorizedApps,omitempty" xml:"AuthorizedApps,omitempty" type:"Struct"`
	PageNumber     *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAuthorizedAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedAppsResponseBody) SetAuthorizedApps(v *DescribeAuthorizedAppsResponseBodyAuthorizedApps) *DescribeAuthorizedAppsResponseBody {
	s.AuthorizedApps = v
	return s
}

func (s *DescribeAuthorizedAppsResponseBody) SetPageNumber(v int32) *DescribeAuthorizedAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBody) SetPageSize(v int32) *DescribeAuthorizedAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBody) SetRequestId(v string) *DescribeAuthorizedAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBody) SetTotalCount(v int32) *DescribeAuthorizedAppsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAuthorizedAppsResponseBodyAuthorizedApps struct {
	AuthorizedApp []*DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp `json:"AuthorizedApp,omitempty" xml:"AuthorizedApp,omitempty" type:"Repeated"`
}

func (s DescribeAuthorizedAppsResponseBodyAuthorizedApps) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedAppsResponseBodyAuthorizedApps) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedApps) SetAuthorizedApp(v []*DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) *DescribeAuthorizedAppsResponseBodyAuthorizedApps {
	s.AuthorizedApp = v
	return s
}

type DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp struct {
	AppId               *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName             *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AuthVaildTime       *string `json:"AuthVaildTime,omitempty" xml:"AuthVaildTime,omitempty"`
	AuthorizationSource *string `json:"AuthorizationSource,omitempty" xml:"AuthorizationSource,omitempty"`
	AuthorizedTime      *string `json:"AuthorizedTime,omitempty" xml:"AuthorizedTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Operator            *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	StageName           *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetAppId(v int64) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.AppId = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetAppName(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.AppName = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetAuthVaildTime(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.AuthVaildTime = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetAuthorizationSource(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.AuthorizationSource = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetAuthorizedTime(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.AuthorizedTime = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetDescription(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.Description = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetOperator(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.Operator = &v
	return s
}

func (s *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp) SetStageName(v string) *DescribeAuthorizedAppsResponseBodyAuthorizedAppsAuthorizedApp {
	s.StageName = &v
	return s
}

type DescribeAuthorizedAppsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAuthorizedAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAuthorizedAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAuthorizedAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAuthorizedAppsResponse) SetHeaders(v map[string]*string) *DescribeAuthorizedAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAuthorizedAppsResponse) SetBody(v *DescribeAuthorizedAppsResponseBody) *DescribeAuthorizedAppsResponse {
	s.Body = v
	return s
}

type DescribeDeployApiTaskRequest struct {
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDeployApiTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployApiTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskRequest) SetOperationUid(v string) *DescribeDeployApiTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeDeployApiTaskRequest) SetSecurityToken(v string) *DescribeDeployApiTaskRequest {
	s.SecurityToken = &v
	return s
}

type DescribeDeployApiTaskResponseBody struct {
	DeployedResults *DescribeDeployApiTaskResponseBodyDeployedResults `json:"DeployedResults,omitempty" xml:"DeployedResults,omitempty" type:"Struct"`
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeployApiTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployApiTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponseBody) SetDeployedResults(v *DescribeDeployApiTaskResponseBodyDeployedResults) *DescribeDeployApiTaskResponseBody {
	s.DeployedResults = v
	return s
}

func (s *DescribeDeployApiTaskResponseBody) SetRequestId(v string) *DescribeDeployApiTaskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeployApiTaskResponseBodyDeployedResults struct {
	DeployedResult []*DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult `json:"DeployedResult,omitempty" xml:"DeployedResult,omitempty" type:"Repeated"`
}

func (s DescribeDeployApiTaskResponseBodyDeployedResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployApiTaskResponseBodyDeployedResults) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResults) SetDeployedResult(v []*DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) *DescribeDeployApiTaskResponseBodyDeployedResults {
	s.DeployedResult = v
	return s
}

type DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult struct {
	ApiUid         *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	DeployedStatus *string `json:"DeployedStatus,omitempty" xml:"DeployedStatus,omitempty"`
	ErrorMsg       *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	StageName      *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetApiUid(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetDeployedStatus(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.DeployedStatus = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetErrorMsg(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetGroupId(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult) SetStageName(v string) *DescribeDeployApiTaskResponseBodyDeployedResultsDeployedResult {
	s.StageName = &v
	return s
}

type DescribeDeployApiTaskResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeployApiTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeployApiTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployApiTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeployApiTaskResponse) SetHeaders(v map[string]*string) *DescribeDeployApiTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeployApiTaskResponse) SetBody(v *DescribeDeployApiTaskResponseBody) *DescribeDeployApiTaskResponse {
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
	AllowSignatureMethod   *string                                                `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiId                  *string                                                `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName                *string                                                `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AuthType               *string                                                `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ConstantParameters     *DescribeDeployedApiResponseBodyConstantParameters     `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	CustomSystemParameters *DescribeDeployedApiResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	DeployedTime           *string                                                `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description            *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableInternet        *bool                                                  `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples       *DescribeDeployedApiResponseBodyErrorCodeSamples       `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	FailResultSample       *string                                                `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	ForceNonceCheck        *bool                                                  `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	GroupId                *string                                                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName              *string                                                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OpenIdConnectConfig    *DescribeDeployedApiResponseBodyOpenIdConnectConfig    `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	RegionId               *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestConfig          *DescribeDeployedApiResponseBodyRequestConfig          `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	RequestId              *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RequestParameters      *DescribeDeployedApiResponseBodyRequestParameters      `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	ResultBodyModel        *string                                                `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	ResultDescriptions     *DescribeDeployedApiResponseBodyResultDescriptions     `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	ResultSample           *string                                                `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType             *string                                                `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	ServiceConfig          *DescribeDeployedApiResponseBodyServiceConfig          `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	ServiceParameters      *DescribeDeployedApiResponseBodyServiceParameters      `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty" type:"Struct"`
	ServiceParametersMap   *DescribeDeployedApiResponseBodyServiceParametersMap   `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty" type:"Struct"`
	StageName              *string                                                `json:"StageName,omitempty" xml:"StageName,omitempty"`
	SystemParameters       *DescribeDeployedApiResponseBodySystemParameters       `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	Visibility             *string                                                `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeDeployedApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBody) SetAllowSignatureMethod(v string) *DescribeDeployedApiResponseBody {
	s.AllowSignatureMethod = &v
	return s
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

func (s *DescribeDeployedApiResponseBody) SetConstantParameters(v *DescribeDeployedApiResponseBodyConstantParameters) *DescribeDeployedApiResponseBody {
	s.ConstantParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetCustomSystemParameters(v *DescribeDeployedApiResponseBodyCustomSystemParameters) *DescribeDeployedApiResponseBody {
	s.CustomSystemParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetDeployedTime(v string) *DescribeDeployedApiResponseBody {
	s.DeployedTime = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetDescription(v string) *DescribeDeployedApiResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetDisableInternet(v bool) *DescribeDeployedApiResponseBody {
	s.DisableInternet = &v
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

func (s *DescribeDeployedApiResponseBody) SetForceNonceCheck(v bool) *DescribeDeployedApiResponseBody {
	s.ForceNonceCheck = &v
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

func (s *DescribeDeployedApiResponseBody) SetOpenIdConnectConfig(v *DescribeDeployedApiResponseBodyOpenIdConnectConfig) *DescribeDeployedApiResponseBody {
	s.OpenIdConnectConfig = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRegionId(v string) *DescribeDeployedApiResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestConfig(v *DescribeDeployedApiResponseBodyRequestConfig) *DescribeDeployedApiResponseBody {
	s.RequestConfig = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestId(v string) *DescribeDeployedApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetRequestParameters(v *DescribeDeployedApiResponseBodyRequestParameters) *DescribeDeployedApiResponseBody {
	s.RequestParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetResultBodyModel(v string) *DescribeDeployedApiResponseBody {
	s.ResultBodyModel = &v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetResultDescriptions(v *DescribeDeployedApiResponseBodyResultDescriptions) *DescribeDeployedApiResponseBody {
	s.ResultDescriptions = v
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

func (s *DescribeDeployedApiResponseBody) SetServiceConfig(v *DescribeDeployedApiResponseBodyServiceConfig) *DescribeDeployedApiResponseBody {
	s.ServiceConfig = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceParameters(v *DescribeDeployedApiResponseBodyServiceParameters) *DescribeDeployedApiResponseBody {
	s.ServiceParameters = v
	return s
}

func (s *DescribeDeployedApiResponseBody) SetServiceParametersMap(v *DescribeDeployedApiResponseBodyServiceParametersMap) *DescribeDeployedApiResponseBody {
	s.ServiceParametersMap = v
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

type DescribeDeployedApiResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyCustomSystemParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeDeployedApiResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

type DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter struct {
	DemoValue            *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) SetDemoValue(v string) *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) SetDescription(v string) *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) SetLocation(v string) *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.Location = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) SetParameterName(v string) *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter {
	s.ParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter {
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

type DescribeDeployedApiResponseBodyOpenIdConnectConfig struct {
	IdTokenParamName *string `json:"IdTokenParamName,omitempty" xml:"IdTokenParamName,omitempty"`
	OpenIdApiType    *string `json:"OpenIdApiType,omitempty" xml:"OpenIdApiType,omitempty"`
	PublicKey        *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	PublicKeyId      *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
}

func (s DescribeDeployedApiResponseBodyOpenIdConnectConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) SetIdTokenParamName(v string) *DescribeDeployedApiResponseBodyOpenIdConnectConfig {
	s.IdTokenParamName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) SetOpenIdApiType(v string) *DescribeDeployedApiResponseBodyOpenIdConnectConfig {
	s.OpenIdApiType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) SetPublicKey(v string) *DescribeDeployedApiResponseBodyOpenIdConnectConfig {
	s.PublicKey = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) SetPublicKeyId(v string) *DescribeDeployedApiResponseBodyOpenIdConnectConfig {
	s.PublicKeyId = &v
	return s
}

type DescribeDeployedApiResponseBodyRequestConfig struct {
	BodyFormat          *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	BodyModel           *string `json:"BodyModel,omitempty" xml:"BodyModel,omitempty"`
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	RequestHttpMethod   *string `json:"RequestHttpMethod,omitempty" xml:"RequestHttpMethod,omitempty"`
	RequestMode         *string `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	RequestPath         *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	RequestProtocol     *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeDeployedApiResponseBodyRequestConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetBodyModel(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.BodyModel = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetPostBodyDescription(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.PostBodyDescription = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetRequestHttpMethod(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.RequestHttpMethod = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetRequestMode(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.RequestMode = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetRequestPath(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.RequestPath = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) SetRequestProtocol(v string) *DescribeDeployedApiResponseBodyRequestConfig {
	s.RequestProtocol = &v
	return s
}

type DescribeDeployedApiResponseBodyRequestParameters struct {
	RequestParameter []*DescribeDeployedApiResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestParameters) SetRequestParameter(v []*DescribeDeployedApiResponseBodyRequestParametersRequestParameter) *DescribeDeployedApiResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

type DescribeDeployedApiResponseBodyRequestParametersRequestParameter struct {
	ApiParameterName  *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	ArrayItemsType    *string `json:"ArrayItemsType,omitempty" xml:"ArrayItemsType,omitempty"`
	DefaultValue      *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	DemoValue         *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocOrder          *int32  `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
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

func (s DescribeDeployedApiResponseBodyRequestParametersRequestParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetApiParameterName(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.ApiParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetArrayItemsType(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.ArrayItemsType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetDefaultValue(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.DefaultValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetDemoValue(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.DemoValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetDescription(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetDocOrder(v int32) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.DocOrder = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetDocShow(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.DocShow = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetEnumValue(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.EnumValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetJsonScheme(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.JsonScheme = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetLocation(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.Location = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetMaxLength(v int64) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.MaxLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetMaxValue(v int64) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.MaxValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetMinLength(v int64) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.MinLength = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetMinValue(v int64) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.MinValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetParameterType(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetRegularExpression(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.RegularExpression = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) SetRequired(v string) *DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	s.Required = &v
	return s
}

type DescribeDeployedApiResponseBodyResultDescriptions struct {
	ResultDescription []*DescribeDeployedApiResponseBodyResultDescriptionsResultDescription `json:"ResultDescription,omitempty" xml:"ResultDescription,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyResultDescriptions) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyResultDescriptions) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyResultDescriptions) SetResultDescription(v []*DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) *DescribeDeployedApiResponseBodyResultDescriptions {
	s.ResultDescription = v
	return s
}

type DescribeDeployedApiResponseBodyResultDescriptionsResultDescription struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HasChild    *bool   `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Mandatory   *bool   `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetDescription(v string) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetHasChild(v bool) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.HasChild = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetId(v string) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Id = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetKey(v string) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Key = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetMandatory(v bool) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Mandatory = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetName(v string) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Name = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetPid(v string) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Pid = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) SetType(v string) *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	s.Type = &v
	return s
}

type DescribeDeployedApiResponseBodyServiceConfig struct {
	FunctionComputeConfig *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	Mock                  *string                                                            `json:"Mock,omitempty" xml:"Mock,omitempty"`
	MockHeaders           *DescribeDeployedApiResponseBodyServiceConfigMockHeaders           `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Struct"`
	MockResult            *string                                                            `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	MockStatusCode        *int32                                                             `json:"MockStatusCode,omitempty" xml:"MockStatusCode,omitempty"`
	ServiceAddress        *string                                                            `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	ServiceHttpMethod     *string                                                            `json:"ServiceHttpMethod,omitempty" xml:"ServiceHttpMethod,omitempty"`
	ServicePath           *string                                                            `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	ServiceProtocol       *string                                                            `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	ServiceTimeout        *int32                                                             `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	ServiceVpcEnable      *string                                                            `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	VpcConfig             *DescribeDeployedApiResponseBodyServiceConfigVpcConfig             `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
	VpcId                 *string                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetFunctionComputeConfig(v *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) *DescribeDeployedApiResponseBodyServiceConfig {
	s.FunctionComputeConfig = v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetMock(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.Mock = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetMockHeaders(v *DescribeDeployedApiResponseBodyServiceConfigMockHeaders) *DescribeDeployedApiResponseBodyServiceConfig {
	s.MockHeaders = v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetMockResult(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.MockResult = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetMockStatusCode(v int32) *DescribeDeployedApiResponseBodyServiceConfig {
	s.MockStatusCode = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetServiceAddress(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.ServiceAddress = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetServiceHttpMethod(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.ServiceHttpMethod = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetServicePath(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.ServicePath = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetServiceProtocol(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.ServiceProtocol = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetServiceTimeout(v int32) *DescribeDeployedApiResponseBodyServiceConfig {
	s.ServiceTimeout = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetServiceVpcEnable(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.ServiceVpcEnable = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetVpcConfig(v *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) *DescribeDeployedApiResponseBodyServiceConfig {
	s.VpcConfig = v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) SetVpcId(v string) *DescribeDeployedApiResponseBodyServiceConfig {
	s.VpcId = &v
	return s
}

type DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig struct {
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	ContentTypeValue    *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	FcBaseUrl           *string `json:"FcBaseUrl,omitempty" xml:"FcBaseUrl,omitempty"`
	FcType              *string `json:"FcType,omitempty" xml:"FcType,omitempty"`
	FunctionName        *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	Method              *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OnlyBusinessPath    *bool   `json:"OnlyBusinessPath,omitempty" xml:"OnlyBusinessPath,omitempty"`
	Path                *string `json:"Path,omitempty" xml:"Path,omitempty"`
	Qualifier           *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RoleArn             *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	ServiceName         *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetContentTypeCatagory(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.ContentTypeCatagory = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetContentTypeValue(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.ContentTypeValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetFcBaseUrl(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FcBaseUrl = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetFcType(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FcType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetFunctionName(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetMethod(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.Method = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetOnlyBusinessPath(v bool) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.OnlyBusinessPath = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetPath(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.Path = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetQualifier(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.Qualifier = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetRegionId(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.RegionId = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetRoleArn(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) SetServiceName(v string) *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	s.ServiceName = &v
	return s
}

type DescribeDeployedApiResponseBodyServiceConfigMockHeaders struct {
	MockHeader []*DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader `json:"MockHeader,omitempty" xml:"MockHeader,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeaders) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeaders) SetMockHeader(v []*DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) *DescribeDeployedApiResponseBodyServiceConfigMockHeaders {
	s.MockHeader = v
	return s
}

type DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader struct {
	HeaderName  *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderName(v string) *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderValue(v string) *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderValue = &v
	return s
}

type DescribeDeployedApiResponseBodyServiceConfigVpcConfig struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port       *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	VpcId      *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigVpcConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) SetInstanceId(v string) *DescribeDeployedApiResponseBodyServiceConfigVpcConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) SetName(v string) *DescribeDeployedApiResponseBodyServiceConfigVpcConfig {
	s.Name = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) SetPort(v int32) *DescribeDeployedApiResponseBodyServiceConfigVpcConfig {
	s.Port = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) SetVpcId(v string) *DescribeDeployedApiResponseBodyServiceConfigVpcConfig {
	s.VpcId = &v
	return s
}

type DescribeDeployedApiResponseBodyServiceParameters struct {
	ServiceParameter []*DescribeDeployedApiResponseBodyServiceParametersServiceParameter `json:"ServiceParameter,omitempty" xml:"ServiceParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyServiceParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParameters) SetServiceParameter(v []*DescribeDeployedApiResponseBodyServiceParametersServiceParameter) *DescribeDeployedApiResponseBodyServiceParameters {
	s.ServiceParameter = v
	return s
}

type DescribeDeployedApiResponseBodyServiceParametersServiceParameter struct {
	Location             *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ParameterType        *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceParametersServiceParameter) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParametersServiceParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) SetLocation(v string) *DescribeDeployedApiResponseBodyServiceParametersServiceParameter {
	s.Location = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) SetParameterType(v string) *DescribeDeployedApiResponseBodyServiceParametersServiceParameter {
	s.ParameterType = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyServiceParametersServiceParameter {
	s.ServiceParameterName = &v
	return s
}

type DescribeDeployedApiResponseBodyServiceParametersMap struct {
	ServiceParameterMap []*DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyServiceParametersMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParametersMap) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMap) SetServiceParameterMap(v []*DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) *DescribeDeployedApiResponseBodyServiceParametersMap {
	s.ServiceParameterMap = v
	return s
}

type DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap struct {
	RequestParameterName *string `json:"RequestParameterName,omitempty" xml:"RequestParameterName,omitempty"`
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) SetRequestParameterName(v string) *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap {
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeployedApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDeployedApiResponse) SetBody(v *DescribeDeployedApiResponseBody) *DescribeDeployedApiResponse {
	s.Body = v
	return s
}

type DescribeDeployedApisRequest struct {
	ApiId         *string                           `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string                           `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	EnableTagAuth *bool                             `json:"EnableTagAuth,omitempty" xml:"EnableTagAuth,omitempty"`
	GroupId       *string                           `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	PageNumber    *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string                           `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string                           `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Tag           []*DescribeDeployedApisRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeDeployedApisRequest) SetEnableTagAuth(v bool) *DescribeDeployedApisRequest {
	s.EnableTagAuth = &v
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

func (s *DescribeDeployedApisRequest) SetTag(v []*DescribeDeployedApisRequestTag) *DescribeDeployedApisRequest {
	s.Tag = v
	return s
}

type DescribeDeployedApisRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDeployedApisRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisRequestTag) SetKey(v string) *DescribeDeployedApisRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDeployedApisRequestTag) SetValue(v string) *DescribeDeployedApisRequestTag {
	s.Value = &v
	return s
}

type DescribeDeployedApisResponseBody struct {
	DeployedApis *DescribeDeployedApisResponseBodyDeployedApis `json:"DeployedApis,omitempty" xml:"DeployedApis,omitempty" type:"Struct"`
	PageNumber   *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeployedApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBody) SetDeployedApis(v *DescribeDeployedApisResponseBodyDeployedApis) *DescribeDeployedApisResponseBody {
	s.DeployedApis = v
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

type DescribeDeployedApisResponseBodyDeployedApis struct {
	DeployedApiItem []*DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem `json:"DeployedApiItem,omitempty" xml:"DeployedApiItem,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApisResponseBodyDeployedApis) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponseBodyDeployedApis) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBodyDeployedApis) SetDeployedApiItem(v []*DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) *DescribeDeployedApisResponseBodyDeployedApis {
	s.DeployedApiItem = v
	return s
}

type DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem struct {
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

func (s DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetApiId(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.ApiId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetApiName(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.ApiName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetDeployedTime(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.DeployedTime = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetDescription(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.Description = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetGroupId(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.GroupId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetGroupName(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.GroupName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetRegionId(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.RegionId = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetStageName(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.StageName = &v
	return s
}

func (s *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem) SetVisibility(v string) *DescribeDeployedApisResponseBodyDeployedApisDeployedApiItem {
	s.Visibility = &v
	return s
}

type DescribeDeployedApisResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDeployedApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" xml:"CertificatePrivateKey,omitempty"`
	DomainBindingStatus   *string `json:"DomainBindingStatus,omitempty" xml:"DomainBindingStatus,omitempty"`
	DomainCNAMEStatus     *string `json:"DomainCNAMEStatus,omitempty" xml:"DomainCNAMEStatus,omitempty"`
	DomainLegalStatus     *string `json:"DomainLegalStatus,omitempty" xml:"DomainLegalStatus,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainRemark          *string `json:"DomainRemark,omitempty" xml:"DomainRemark,omitempty"`
	DomainWebSocketStatus *string `json:"DomainWebSocketStatus,omitempty" xml:"DomainWebSocketStatus,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
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

func (s *DescribeDomainResponseBody) SetCertificatePrivateKey(v string) *DescribeDomainResponseBody {
	s.CertificatePrivateKey = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainBindingStatus(v string) *DescribeDomainResponseBody {
	s.DomainBindingStatus = &v
	return s
}

func (s *DescribeDomainResponseBody) SetDomainCNAMEStatus(v string) *DescribeDomainResponseBody {
	s.DomainCNAMEStatus = &v
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

func (s *DescribeDomainResponseBody) SetRequestId(v string) *DescribeDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainResponseBody) SetSubDomain(v string) *DescribeDomainResponseBody {
	s.SubDomain = &v
	return s
}

type DescribeDomainResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeDomainResponse) SetBody(v *DescribeDomainResponseBody) *DescribeDomainResponse {
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
	ApiHisItems *DescribeHistoryApisResponseBodyApiHisItems `json:"ApiHisItems,omitempty" xml:"ApiHisItems,omitempty" type:"Struct"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHistoryApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBody) SetApiHisItems(v *DescribeHistoryApisResponseBodyApiHisItems) *DescribeHistoryApisResponseBody {
	s.ApiHisItems = v
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

type DescribeHistoryApisResponseBodyApiHisItems struct {
	ApiHisItem []*DescribeHistoryApisResponseBodyApiHisItemsApiHisItem `json:"ApiHisItem,omitempty" xml:"ApiHisItem,omitempty" type:"Repeated"`
}

func (s DescribeHistoryApisResponseBodyApiHisItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponseBodyApiHisItems) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBodyApiHisItems) SetApiHisItem(v []*DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) *DescribeHistoryApisResponseBodyApiHisItems {
	s.ApiHisItem = v
	return s
}

type DescribeHistoryApisResponseBodyApiHisItemsApiHisItem struct {
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

func (s DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) GoString() string {
	return s.String()
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetApiId(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.ApiId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetApiName(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.ApiName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetDeployedTime(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.DeployedTime = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetDescription(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.Description = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetGroupId(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.GroupId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetGroupName(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.GroupName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetHistoryVersion(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.HistoryVersion = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetRegionId(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.RegionId = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetStageName(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.StageName = &v
	return s
}

func (s *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem) SetStatus(v string) *DescribeHistoryApisResponseBodyApiHisItemsApiHisItem {
	s.Status = &v
	return s
}

type DescribeHistoryApisResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeHistoryApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpControlPolicyItemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeIpControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeLogConfigResponse) SetBody(v *DescribeLogConfigResponseBody) *DescribeLogConfigResponse {
	s.Body = v
	return s
}

type DescribeMarketRemainsQuotaRequest struct {
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeMarketRemainsQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMarketRemainsQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeMarketRemainsQuotaRequest) SetDomainName(v string) *DescribeMarketRemainsQuotaRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeMarketRemainsQuotaRequest) SetSecurityToken(v string) *DescribeMarketRemainsQuotaRequest {
	s.SecurityToken = &v
	return s
}

type DescribeMarketRemainsQuotaResponseBody struct {
	RemainsQuota *int64  `json:"RemainsQuota,omitempty" xml:"RemainsQuota,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMarketRemainsQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMarketRemainsQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMarketRemainsQuotaResponseBody) SetRemainsQuota(v int64) *DescribeMarketRemainsQuotaResponseBody {
	s.RemainsQuota = &v
	return s
}

func (s *DescribeMarketRemainsQuotaResponseBody) SetRequestId(v string) *DescribeMarketRemainsQuotaResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMarketRemainsQuotaResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMarketRemainsQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMarketRemainsQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMarketRemainsQuotaResponse) GoString() string {
	return s.String()
}

func (s *DescribeMarketRemainsQuotaResponse) SetHeaders(v map[string]*string) *DescribeMarketRemainsQuotaResponse {
	s.Headers = v
	return s
}

func (s *DescribeMarketRemainsQuotaResponse) SetBody(v *DescribeMarketRemainsQuotaResponseBody) *DescribeMarketRemainsQuotaResponse {
	s.Body = v
	return s
}

type DescribeModelsRequest struct {
	GroupId    *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelId    *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName  *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeModelsRequest) SetGroupId(v string) *DescribeModelsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeModelsRequest) SetModelId(v string) *DescribeModelsRequest {
	s.ModelId = &v
	return s
}

func (s *DescribeModelsRequest) SetModelName(v string) *DescribeModelsRequest {
	s.ModelName = &v
	return s
}

func (s *DescribeModelsRequest) SetPageNumber(v int32) *DescribeModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelsRequest) SetPageSize(v int32) *DescribeModelsRequest {
	s.PageSize = &v
	return s
}

type DescribeModelsResponseBody struct {
	ModelDetails *DescribeModelsResponseBodyModelDetails `json:"ModelDetails,omitempty" xml:"ModelDetails,omitempty" type:"Struct"`
	PageNumber   *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBody) SetModelDetails(v *DescribeModelsResponseBodyModelDetails) *DescribeModelsResponseBody {
	s.ModelDetails = v
	return s
}

func (s *DescribeModelsResponseBody) SetPageNumber(v int32) *DescribeModelsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelsResponseBody) SetPageSize(v int32) *DescribeModelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeModelsResponseBody) SetRequestId(v string) *DescribeModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelsResponseBody) SetTotalCount(v int32) *DescribeModelsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeModelsResponseBodyModelDetails struct {
	ModelDetail []*DescribeModelsResponseBodyModelDetailsModelDetail `json:"ModelDetail,omitempty" xml:"ModelDetail,omitempty" type:"Repeated"`
}

func (s DescribeModelsResponseBodyModelDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsResponseBodyModelDetails) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBodyModelDetails) SetModelDetail(v []*DescribeModelsResponseBodyModelDetailsModelDetail) *DescribeModelsResponseBodyModelDetails {
	s.ModelDetail = v
	return s
}

type DescribeModelsResponseBodyModelDetailsModelDetail struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelId      *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	ModelName    *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelRef     *string `json:"ModelRef,omitempty" xml:"ModelRef,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeModelsResponseBodyModelDetailsModelDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsResponseBodyModelDetailsModelDetail) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetCreatedTime(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.CreatedTime = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetDescription(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.Description = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetGroupId(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.GroupId = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModelId(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModelId = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModelName(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModelName = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModelRef(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModelRef = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModifiedTime(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetSchema(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.Schema = &v
	return s
}

type DescribeModelsResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponse) SetHeaders(v map[string]*string) *DescribeModelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelsResponse) SetBody(v *DescribeModelsResponseBody) *DescribeModelsResponse {
	s.Body = v
	return s
}

type DescribePluginSchemasRequest struct {
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePluginSchemasRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginSchemasRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasRequest) SetLanguage(v string) *DescribePluginSchemasRequest {
	s.Language = &v
	return s
}

func (s *DescribePluginSchemasRequest) SetSecurityToken(v string) *DescribePluginSchemasRequest {
	s.SecurityToken = &v
	return s
}

type DescribePluginSchemasResponseBody struct {
	PluginSchemas *DescribePluginSchemasResponseBodyPluginSchemas `json:"PluginSchemas,omitempty" xml:"PluginSchemas,omitempty" type:"Struct"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePluginSchemasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponseBody) SetPluginSchemas(v *DescribePluginSchemasResponseBodyPluginSchemas) *DescribePluginSchemasResponseBody {
	s.PluginSchemas = v
	return s
}

func (s *DescribePluginSchemasResponseBody) SetRequestId(v string) *DescribePluginSchemasResponseBody {
	s.RequestId = &v
	return s
}

type DescribePluginSchemasResponseBodyPluginSchemas struct {
	PluginSchema []*DescribePluginSchemasResponseBodyPluginSchemasPluginSchema `json:"PluginSchema,omitempty" xml:"PluginSchema,omitempty" type:"Repeated"`
}

func (s DescribePluginSchemasResponseBodyPluginSchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginSchemasResponseBodyPluginSchemas) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponseBodyPluginSchemas) SetPluginSchema(v []*DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) *DescribePluginSchemasResponseBodyPluginSchemas {
	s.PluginSchema = v
	return s
}

type DescribePluginSchemasResponseBodyPluginSchemasPluginSchema struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentId     *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SupportClassic *bool   `json:"SupportClassic,omitempty" xml:"SupportClassic,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetDescription(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.Description = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetDocumentId(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.DocumentId = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetName(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.Name = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetSupportClassic(v bool) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.SupportClassic = &v
	return s
}

func (s *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema) SetTitle(v string) *DescribePluginSchemasResponseBodyPluginSchemasPluginSchema {
	s.Title = &v
	return s
}

type DescribePluginSchemasResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePluginSchemasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePluginSchemasResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginSchemasResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginSchemasResponse) SetHeaders(v map[string]*string) *DescribePluginSchemasResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginSchemasResponse) SetBody(v *DescribePluginSchemasResponseBody) *DescribePluginSchemasResponse {
	s.Body = v
	return s
}

type DescribePluginTemplatesRequest struct {
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	PluginName    *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePluginTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesRequest) SetLanguage(v string) *DescribePluginTemplatesRequest {
	s.Language = &v
	return s
}

func (s *DescribePluginTemplatesRequest) SetPluginName(v string) *DescribePluginTemplatesRequest {
	s.PluginName = &v
	return s
}

func (s *DescribePluginTemplatesRequest) SetSecurityToken(v string) *DescribePluginTemplatesRequest {
	s.SecurityToken = &v
	return s
}

type DescribePluginTemplatesResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates *DescribePluginTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Struct"`
}

func (s DescribePluginTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponseBody) SetRequestId(v string) *DescribePluginTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginTemplatesResponseBody) SetTemplates(v *DescribePluginTemplatesResponseBodyTemplates) *DescribePluginTemplatesResponseBody {
	s.Templates = v
	return s
}

type DescribePluginTemplatesResponseBodyTemplates struct {
	Template []*DescribePluginTemplatesResponseBodyTemplatesTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s DescribePluginTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponseBodyTemplates) SetTemplate(v []*DescribePluginTemplatesResponseBodyTemplatesTemplate) *DescribePluginTemplatesResponseBodyTemplates {
	s.Template = v
	return s
}

type DescribePluginTemplatesResponseBodyTemplatesTemplate struct {
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentAnchor *string `json:"DocumentAnchor,omitempty" xml:"DocumentAnchor,omitempty"`
	DocumentId     *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	Sample         *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePluginTemplatesResponseBodyTemplatesTemplate) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginTemplatesResponseBodyTemplatesTemplate) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetDescription(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.Description = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetDocumentAnchor(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.DocumentAnchor = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetDocumentId(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.DocumentId = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetSample(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.Sample = &v
	return s
}

func (s *DescribePluginTemplatesResponseBodyTemplatesTemplate) SetTitle(v string) *DescribePluginTemplatesResponseBodyTemplatesTemplate {
	s.Title = &v
	return s
}

type DescribePluginTemplatesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePluginTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePluginTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginTemplatesResponse) SetHeaders(v map[string]*string) *DescribePluginTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginTemplatesResponse) SetBody(v *DescribePluginTemplatesResponseBody) *DescribePluginTemplatesResponse {
	s.Body = v
	return s
}

type DescribePluginsRequest struct {
	PageNumber    *int32                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PluginId      *string                      `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	PluginName    *string                      `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	PluginType    *string                      `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	SecurityToken *string                      `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*DescribePluginsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePluginsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginsRequest) SetPageNumber(v int32) *DescribePluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsRequest) SetPageSize(v int32) *DescribePluginsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsRequest) SetPluginId(v string) *DescribePluginsRequest {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsRequest) SetPluginName(v string) *DescribePluginsRequest {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsRequest) SetPluginType(v string) *DescribePluginsRequest {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsRequest) SetSecurityToken(v string) *DescribePluginsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginsRequest) SetTag(v []*DescribePluginsRequestTag) *DescribePluginsRequest {
	s.Tag = v
	return s
}

type DescribePluginsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePluginsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribePluginsRequestTag) SetKey(v string) *DescribePluginsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribePluginsRequestTag) SetValue(v string) *DescribePluginsRequestTag {
	s.Value = &v
	return s
}

type DescribePluginsResponseBody struct {
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Plugins    *DescribePluginsResponseBodyPlugins `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Struct"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBody) SetPageNumber(v int32) *DescribePluginsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsResponseBody) SetPageSize(v int32) *DescribePluginsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsResponseBody) SetPlugins(v *DescribePluginsResponseBodyPlugins) *DescribePluginsResponseBody {
	s.Plugins = v
	return s
}

func (s *DescribePluginsResponseBody) SetRequestId(v string) *DescribePluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginsResponseBody) SetTotalCount(v int32) *DescribePluginsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePluginsResponseBodyPlugins struct {
	PluginAttribute []*DescribePluginsResponseBodyPluginsPluginAttribute `json:"PluginAttribute,omitempty" xml:"PluginAttribute,omitempty" type:"Repeated"`
}

func (s DescribePluginsResponseBodyPlugins) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsResponseBodyPlugins) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPlugins) SetPluginAttribute(v []*DescribePluginsResponseBodyPluginsPluginAttribute) *DescribePluginsResponseBodyPlugins {
	s.PluginAttribute = v
	return s
}

type DescribePluginsResponseBodyPluginsPluginAttribute struct {
	CreatedTime  *string                                                `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *int32                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string                                                `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PluginData   *string                                                `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	PluginId     *string                                                `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	PluginName   *string                                                `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	PluginType   *string                                                `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	RegionId     *int32                                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Tags         *DescribePluginsResponseBodyPluginsPluginAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribePluginsResponseBodyPluginsPluginAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsResponseBodyPluginsPluginAttribute) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetCreatedTime(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetDescription(v int32) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.Description = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetModifiedTime(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginData(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginData = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginId(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginName(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginType(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetRegionId(v int32) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetTags(v *DescribePluginsResponseBodyPluginsPluginAttributeTags) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.Tags = v
	return s
}

type DescribePluginsResponseBodyPluginsPluginAttributeTags struct {
	TagInfo []*DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTags) SetTagInfo(v []*DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) *DescribePluginsResponseBodyPluginsPluginAttributeTags {
	s.TagInfo = v
	return s
}

type DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) SetKey(v string) *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) SetValue(v string) *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo {
	s.Value = &v
	return s
}

type DescribePluginsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePluginsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePluginsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponse) SetHeaders(v map[string]*string) *DescribePluginsResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginsResponse) SetBody(v *DescribePluginsResponseBody) *DescribePluginsResponse {
	s.Body = v
	return s
}

type DescribePluginsByApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribePluginsByApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiRequest) SetApiId(v string) *DescribePluginsByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetGroupId(v string) *DescribePluginsByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetSecurityToken(v string) *DescribePluginsByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginsByApiRequest) SetStageName(v string) *DescribePluginsByApiRequest {
	s.StageName = &v
	return s
}

type DescribePluginsByApiResponseBody struct {
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Plugins    *DescribePluginsByApiResponseBodyPlugins `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Struct"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginsByApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponseBody) SetPageNumber(v int32) *DescribePluginsByApiResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetPageSize(v int32) *DescribePluginsByApiResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetPlugins(v *DescribePluginsByApiResponseBodyPlugins) *DescribePluginsByApiResponseBody {
	s.Plugins = v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetRequestId(v string) *DescribePluginsByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginsByApiResponseBody) SetTotalCount(v int32) *DescribePluginsByApiResponseBody {
	s.TotalCount = &v
	return s
}

type DescribePluginsByApiResponseBodyPlugins struct {
	PluginAttribute []*DescribePluginsByApiResponseBodyPluginsPluginAttribute `json:"PluginAttribute,omitempty" xml:"PluginAttribute,omitempty" type:"Repeated"`
}

func (s DescribePluginsByApiResponseBodyPlugins) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsByApiResponseBodyPlugins) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponseBodyPlugins) SetPluginAttribute(v []*DescribePluginsByApiResponseBodyPluginsPluginAttribute) *DescribePluginsByApiResponseBodyPlugins {
	s.PluginAttribute = v
	return s
}

type DescribePluginsByApiResponseBodyPluginsPluginAttribute struct {
	CreatedTime  *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PluginData   *string `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	PluginId     *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	PluginName   *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	PluginType   *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribePluginsByApiResponseBodyPluginsPluginAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsByApiResponseBodyPluginsPluginAttribute) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetCreatedTime(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetDescription(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.Description = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetModifiedTime(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginData(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginData = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginId(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginName(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetPluginType(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsByApiResponseBodyPluginsPluginAttribute) SetRegionId(v string) *DescribePluginsByApiResponseBodyPluginsPluginAttribute {
	s.RegionId = &v
	return s
}

type DescribePluginsByApiResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePluginsByApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePluginsByApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePluginsByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribePluginsByApiResponse) SetHeaders(v map[string]*string) *DescribePluginsByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribePluginsByApiResponse) SetBody(v *DescribePluginsByApiResponseBody) *DescribePluginsByApiResponse {
	s.Body = v
	return s
}

type DescribePurchasedApiGroupRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribePurchasedApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupRequest) SetGroupId(v string) *DescribePurchasedApiGroupRequest {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupRequest) SetSecurityToken(v string) *DescribePurchasedApiGroupRequest {
	s.SecurityToken = &v
	return s
}

type DescribePurchasedApiGroupResponseBody struct {
	Description   *string                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Domains       *DescribePurchasedApiGroupResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Struct"`
	GroupId       *string                                       `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string                                       `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	PurchasedTime *string                                       `json:"PurchasedTime,omitempty" xml:"PurchasedTime,omitempty"`
	RegionId      *string                                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status        *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePurchasedApiGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponseBody) SetDescription(v string) *DescribePurchasedApiGroupResponseBody {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetDomains(v *DescribePurchasedApiGroupResponseBodyDomains) *DescribePurchasedApiGroupResponseBody {
	s.Domains = v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetGroupId(v string) *DescribePurchasedApiGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetGroupName(v string) *DescribePurchasedApiGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetPurchasedTime(v string) *DescribePurchasedApiGroupResponseBody {
	s.PurchasedTime = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetRegionId(v string) *DescribePurchasedApiGroupResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetRequestId(v string) *DescribePurchasedApiGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePurchasedApiGroupResponseBody) SetStatus(v string) *DescribePurchasedApiGroupResponseBody {
	s.Status = &v
	return s
}

type DescribePurchasedApiGroupResponseBodyDomains struct {
	DomainItem []*DescribePurchasedApiGroupResponseBodyDomainsDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApiGroupResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponseBodyDomains) SetDomainItem(v []*DescribePurchasedApiGroupResponseBodyDomainsDomainItem) *DescribePurchasedApiGroupResponseBodyDomains {
	s.DomainItem = v
	return s
}

type DescribePurchasedApiGroupResponseBodyDomainsDomainItem struct {
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribePurchasedApiGroupResponseBodyDomainsDomainItem) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupResponseBodyDomainsDomainItem) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponseBodyDomainsDomainItem) SetDomainName(v string) *DescribePurchasedApiGroupResponseBodyDomainsDomainItem {
	s.DomainName = &v
	return s
}

type DescribePurchasedApiGroupResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePurchasedApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePurchasedApiGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApiGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApiGroupResponse) SetHeaders(v map[string]*string) *DescribePurchasedApiGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribePurchasedApiGroupResponse) SetBody(v *DescribePurchasedApiGroupResponseBody) *DescribePurchasedApiGroupResponse {
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
	BillingType    *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	Description    *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExpireTime     *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InvokeTimesMax *int64  `json:"InvokeTimesMax,omitempty" xml:"InvokeTimesMax,omitempty"`
	InvokeTimesNow *int64  `json:"InvokeTimesNow,omitempty" xml:"InvokeTimesNow,omitempty"`
	PurchasedTime  *string `json:"PurchasedTime,omitempty" xml:"PurchasedTime,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetDescription(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetExpireTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.ExpireTime = &v
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

func (s *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute) SetPurchasedTime(v string) *DescribePurchasedApiGroupsResponseBodyPurchasedApiGroupAttributesPurchasedApiGroupAttribute {
	s.PurchasedTime = &v
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
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePurchasedApiGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PurchasedApis *DescribePurchasedApisResponseBodyPurchasedApis `json:"PurchasedApis,omitempty" xml:"PurchasedApis,omitempty" type:"Struct"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePurchasedApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponseBody) SetPageNumber(v int32) *DescribePurchasedApisResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePurchasedApisResponseBody) SetPageSize(v int32) *DescribePurchasedApisResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePurchasedApisResponseBody) SetPurchasedApis(v *DescribePurchasedApisResponseBodyPurchasedApis) *DescribePurchasedApisResponseBody {
	s.PurchasedApis = v
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

type DescribePurchasedApisResponseBodyPurchasedApis struct {
	PurchasedApi []*DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi `json:"PurchasedApi,omitempty" xml:"PurchasedApi,omitempty" type:"Repeated"`
}

func (s DescribePurchasedApisResponseBodyPurchasedApis) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponseBodyPurchasedApis) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponseBodyPurchasedApis) SetPurchasedApi(v []*DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) *DescribePurchasedApisResponseBodyPurchasedApis {
	s.PurchasedApi = v
	return s
}

type DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName       *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	DeployedTime  *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	PurchasedTime *string `json:"PurchasedTime,omitempty" xml:"PurchasedTime,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	Visibility    *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) String() string {
	return tea.Prettify(s)
}

func (s DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) GoString() string {
	return s.String()
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetApiId(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.ApiId = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetApiName(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.ApiName = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetDeployedTime(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.DeployedTime = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetDescription(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.Description = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetGroupId(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.GroupId = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetGroupName(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.GroupName = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetModifiedTime(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetPurchasedTime(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.PurchasedTime = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetRegionId(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.RegionId = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetStageName(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.StageName = &v
	return s
}

func (s *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi) SetVisibility(v string) *DescribePurchasedApisResponseBodyPurchasedApisPurchasedApi {
	s.Visibility = &v
	return s
}

type DescribePurchasedApisResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePurchasedApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribePurchasedApisResponse) SetBody(v *DescribePurchasedApisResponseBody) *DescribePurchasedApisResponse {
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSignaturesRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s DescribeSignaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesRequest) SetPageNumber(v int32) *DescribeSignaturesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSignaturesRequest) SetPageSize(v int32) *DescribeSignaturesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSignaturesRequest) SetSecurityToken(v string) *DescribeSignaturesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSignaturesRequest) SetSignatureId(v string) *DescribeSignaturesRequest {
	s.SignatureId = &v
	return s
}

func (s *DescribeSignaturesRequest) SetSignatureName(v string) *DescribeSignaturesRequest {
	s.SignatureName = &v
	return s
}

type DescribeSignaturesResponseBody struct {
	PageNumber     *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignatureInfos *DescribeSignaturesResponseBodySignatureInfos `json:"SignatureInfos,omitempty" xml:"SignatureInfos,omitempty" type:"Struct"`
	TotalCount     *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSignaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponseBody) SetPageNumber(v int32) *DescribeSignaturesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSignaturesResponseBody) SetPageSize(v int32) *DescribeSignaturesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSignaturesResponseBody) SetRequestId(v string) *DescribeSignaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignaturesResponseBody) SetSignatureInfos(v *DescribeSignaturesResponseBodySignatureInfos) *DescribeSignaturesResponseBody {
	s.SignatureInfos = v
	return s
}

func (s *DescribeSignaturesResponseBody) SetTotalCount(v int32) *DescribeSignaturesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSignaturesResponseBodySignatureInfos struct {
	SignatureInfo []*DescribeSignaturesResponseBodySignatureInfosSignatureInfo `json:"SignatureInfo,omitempty" xml:"SignatureInfo,omitempty" type:"Repeated"`
}

func (s DescribeSignaturesResponseBodySignatureInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesResponseBodySignatureInfos) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponseBodySignatureInfos) SetSignatureInfo(v []*DescribeSignaturesResponseBodySignatureInfosSignatureInfo) *DescribeSignaturesResponseBodySignatureInfos {
	s.SignatureInfo = v
	return s
}

type DescribeSignaturesResponseBodySignatureInfosSignatureInfo struct {
	CreatedTime     *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	ModifiedTime    *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SignatureId     *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureKey    *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	SignatureName   *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	SignatureSecret *string `json:"SignatureSecret,omitempty" xml:"SignatureSecret,omitempty"`
}

func (s DescribeSignaturesResponseBodySignatureInfosSignatureInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesResponseBodySignatureInfosSignatureInfo) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetCreatedTime(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.CreatedTime = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetModifiedTime(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetRegionId(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureId(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureId = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureKey(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureKey = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureName(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureName = &v
	return s
}

func (s *DescribeSignaturesResponseBodySignatureInfosSignatureInfo) SetSignatureSecret(v string) *DescribeSignaturesResponseBodySignatureInfosSignatureInfo {
	s.SignatureSecret = &v
	return s
}

type DescribeSignaturesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSignaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSignaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesResponse) SetHeaders(v map[string]*string) *DescribeSignaturesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignaturesResponse) SetBody(v *DescribeSignaturesResponseBody) *DescribeSignaturesResponse {
	s.Body = v
	return s
}

type DescribeSignaturesByApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeSignaturesByApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiRequest) SetApiId(v string) *DescribeSignaturesByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) SetGroupId(v string) *DescribeSignaturesByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) SetSecurityToken(v string) *DescribeSignaturesByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeSignaturesByApiRequest) SetStageName(v string) *DescribeSignaturesByApiRequest {
	s.StageName = &v
	return s
}

type DescribeSignaturesByApiResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Signatures *DescribeSignaturesByApiResponseBodySignatures `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Struct"`
}

func (s DescribeSignaturesByApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponseBody) SetRequestId(v string) *DescribeSignaturesByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBody) SetSignatures(v *DescribeSignaturesByApiResponseBodySignatures) *DescribeSignaturesByApiResponseBody {
	s.Signatures = v
	return s
}

type DescribeSignaturesByApiResponseBodySignatures struct {
	SignatureItem []*DescribeSignaturesByApiResponseBodySignaturesSignatureItem `json:"SignatureItem,omitempty" xml:"SignatureItem,omitempty" type:"Repeated"`
}

func (s DescribeSignaturesByApiResponseBodySignatures) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesByApiResponseBodySignatures) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponseBodySignatures) SetSignatureItem(v []*DescribeSignaturesByApiResponseBodySignaturesSignatureItem) *DescribeSignaturesByApiResponseBodySignatures {
	s.SignatureItem = v
	return s
}

type DescribeSignaturesByApiResponseBodySignaturesSignatureItem struct {
	BoundTime     *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s DescribeSignaturesByApiResponseBodySignaturesSignatureItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesByApiResponseBodySignaturesSignatureItem) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) SetBoundTime(v string) *DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) SetSignatureId(v string) *DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	s.SignatureId = &v
	return s
}

func (s *DescribeSignaturesByApiResponseBodySignaturesSignatureItem) SetSignatureName(v string) *DescribeSignaturesByApiResponseBodySignaturesSignatureItem {
	s.SignatureName = &v
	return s
}

type DescribeSignaturesByApiResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSignaturesByApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSignaturesByApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSignaturesByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeSignaturesByApiResponse) SetHeaders(v map[string]*string) *DescribeSignaturesByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeSignaturesByApiResponse) SetBody(v *DescribeSignaturesByApiResponseBody) *DescribeSignaturesByApiResponse {
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
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SystemParams *DescribeSystemParametersResponseBodySystemParams `json:"SystemParams,omitempty" xml:"SystemParams,omitempty" type:"Struct"`
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

func (s *DescribeSystemParametersResponseBody) SetSystemParams(v *DescribeSystemParametersResponseBodySystemParams) *DescribeSystemParametersResponseBody {
	s.SystemParams = v
	return s
}

type DescribeSystemParametersResponseBodySystemParams struct {
	SystemParamItem []*DescribeSystemParametersResponseBodySystemParamsSystemParamItem `json:"SystemParamItem,omitempty" xml:"SystemParamItem,omitempty" type:"Repeated"`
}

func (s DescribeSystemParametersResponseBodySystemParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersResponseBodySystemParams) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBodySystemParams) SetSystemParamItem(v []*DescribeSystemParametersResponseBodySystemParamsSystemParamItem) *DescribeSystemParametersResponseBodySystemParams {
	s.SystemParamItem = v
	return s
}

type DescribeSystemParametersResponseBodySystemParamsSystemParamItem struct {
	DemoValue   *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ParamName   *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	ParamType   *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
}

func (s DescribeSystemParametersResponseBodySystemParamsSystemParamItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSystemParametersResponseBodySystemParamsSystemParamItem) GoString() string {
	return s.String()
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetDemoValue(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.DemoValue = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetDescription(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.Description = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetParamName(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.ParamName = &v
	return s
}

func (s *DescribeSystemParametersResponseBodySystemParamsSystemParamItem) SetParamType(v string) *DescribeSystemParametersResponseBodySystemParamsSystemParamItem {
	s.ParamType = &v
	return s
}

type DescribeSystemParametersResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSystemParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSystemParametersResponse) SetBody(v *DescribeSystemParametersResponseBody) *DescribeSystemParametersResponse {
	s.Body = v
	return s
}

type DescribeTrafficControlsRequest struct {
	ApiId              *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
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

func (s *DescribeTrafficControlsRequest) SetApiId(v string) *DescribeTrafficControlsRequest {
	s.ApiId = &v
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
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTrafficControlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeTrafficControlsResponse) SetBody(v *DescribeTrafficControlsResponseBody) *DescribeTrafficControlsResponse {
	s.Body = v
	return s
}

type DescribeTrafficControlsByApiRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeTrafficControlsByApiRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsByApiRequest) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiRequest) SetApiId(v string) *DescribeTrafficControlsByApiRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) SetGroupId(v string) *DescribeTrafficControlsByApiRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) SetSecurityToken(v string) *DescribeTrafficControlsByApiRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeTrafficControlsByApiRequest) SetStageName(v string) *DescribeTrafficControlsByApiRequest {
	s.StageName = &v
	return s
}

type DescribeTrafficControlsByApiResponseBody struct {
	RequestId           *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficControlItems *DescribeTrafficControlsByApiResponseBodyTrafficControlItems `json:"TrafficControlItems,omitempty" xml:"TrafficControlItems,omitempty" type:"Struct"`
}

func (s DescribeTrafficControlsByApiResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponseBody) SetRequestId(v string) *DescribeTrafficControlsByApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBody) SetTrafficControlItems(v *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) *DescribeTrafficControlsByApiResponseBody {
	s.TrafficControlItems = v
	return s
}

type DescribeTrafficControlsByApiResponseBodyTrafficControlItems struct {
	TrafficControlItem []*DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem `json:"TrafficControlItem,omitempty" xml:"TrafficControlItem,omitempty" type:"Repeated"`
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItems) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItems) SetTrafficControlItem(v []*DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) *DescribeTrafficControlsByApiResponseBodyTrafficControlItems {
	s.TrafficControlItem = v
	return s
}

type DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem struct {
	BoundTime              *string `json:"BoundTime,omitempty" xml:"BoundTime,omitempty"`
	TrafficControlItemId   *string `json:"TrafficControlItemId,omitempty" xml:"TrafficControlItemId,omitempty"`
	TrafficControlItemName *string `json:"TrafficControlItemName,omitempty" xml:"TrafficControlItemName,omitempty"`
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) SetBoundTime(v string) *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	s.BoundTime = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) SetTrafficControlItemId(v string) *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	s.TrafficControlItemId = &v
	return s
}

func (s *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem) SetTrafficControlItemName(v string) *DescribeTrafficControlsByApiResponseBodyTrafficControlItemsTrafficControlItem {
	s.TrafficControlItemName = &v
	return s
}

type DescribeTrafficControlsByApiResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeTrafficControlsByApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTrafficControlsByApiResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTrafficControlsByApiResponse) GoString() string {
	return s.String()
}

func (s *DescribeTrafficControlsByApiResponse) SetHeaders(v map[string]*string) *DescribeTrafficControlsByApiResponse {
	s.Headers = v
	return s
}

func (s *DescribeTrafficControlsByApiResponse) SetBody(v *DescribeTrafficControlsByApiResponseBody) *DescribeTrafficControlsByApiResponse {
	s.Body = v
	return s
}

type DescribeUpdateVpcInfoTaskRequest struct {
	OperationUid  *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskRequest) SetOperationUid(v string) *DescribeUpdateVpcInfoTaskRequest {
	s.OperationUid = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskRequest) SetSecurityToken(v string) *DescribeUpdateVpcInfoTaskRequest {
	s.SecurityToken = &v
	return s
}

type DescribeUpdateVpcInfoTaskResponseBody struct {
	ApiUpdateVpcInfoResults *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults `json:"ApiUpdateVpcInfoResults,omitempty" xml:"ApiUpdateVpcInfoResults,omitempty" type:"Struct"`
	RequestId               *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) SetApiUpdateVpcInfoResults(v *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) *DescribeUpdateVpcInfoTaskResponseBody {
	s.ApiUpdateVpcInfoResults = v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBody) SetRequestId(v string) *DescribeUpdateVpcInfoTaskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults struct {
	ApiUpdateVpcInfoResult []*DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult `json:"ApiUpdateVpcInfoResult,omitempty" xml:"ApiUpdateVpcInfoResult,omitempty" type:"Repeated"`
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults) SetApiUpdateVpcInfoResult(v []*DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResults {
	s.ApiUpdateVpcInfoResult = v
	return s
}

type DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult struct {
	ApiName      *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	ApiUid       *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	ErrorMsg     *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	StageId      *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	StageName    *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	UpdateStatus *string `json:"UpdateStatus,omitempty" xml:"UpdateStatus,omitempty"`
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetApiName(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.ApiName = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetApiUid(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.ApiUid = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetErrorMsg(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetGroupId(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.GroupId = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetGroupName(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.GroupName = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetStageId(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.StageId = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetStageName(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.StageName = &v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult) SetUpdateStatus(v string) *DescribeUpdateVpcInfoTaskResponseBodyApiUpdateVpcInfoResultsApiUpdateVpcInfoResult {
	s.UpdateStatus = &v
	return s
}

type DescribeUpdateVpcInfoTaskResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUpdateVpcInfoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUpdateVpcInfoTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUpdateVpcInfoTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpdateVpcInfoTaskResponse) SetHeaders(v map[string]*string) *DescribeUpdateVpcInfoTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpdateVpcInfoTaskResponse) SetBody(v *DescribeUpdateVpcInfoTaskResponseBody) *DescribeUpdateVpcInfoTaskResponse {
	s.Body = v
	return s
}

type DescribeVpcAccessesRequest struct {
	// VPC授权名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 当前页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页展示条目
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Vpc授权ID
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
}

func (s DescribeVpcAccessesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcAccessesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesRequest) SetName(v string) *DescribeVpcAccessesRequest {
	s.Name = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetPageNumber(v int32) *DescribeVpcAccessesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetPageSize(v int32) *DescribeVpcAccessesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetSecurityToken(v string) *DescribeVpcAccessesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeVpcAccessesRequest) SetVpcAccessId(v string) *DescribeVpcAccessesRequest {
	s.VpcAccessId = &v
	return s
}

type DescribeVpcAccessesResponseBody struct {
	PageNumber          *int32                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId           *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount          *int32                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VpcAccessAttributes *DescribeVpcAccessesResponseBodyVpcAccessAttributes `json:"VpcAccessAttributes,omitempty" xml:"VpcAccessAttributes,omitempty" type:"Struct"`
}

func (s DescribeVpcAccessesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcAccessesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBody) SetPageNumber(v int32) *DescribeVpcAccessesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetPageSize(v int32) *DescribeVpcAccessesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetRequestId(v string) *DescribeVpcAccessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetTotalCount(v int32) *DescribeVpcAccessesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcAccessesResponseBody) SetVpcAccessAttributes(v *DescribeVpcAccessesResponseBodyVpcAccessAttributes) *DescribeVpcAccessesResponseBody {
	s.VpcAccessAttributes = v
	return s
}

type DescribeVpcAccessesResponseBodyVpcAccessAttributes struct {
	VpcAccessAttribute []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute `json:"VpcAccessAttribute,omitempty" xml:"VpcAccessAttribute,omitempty" type:"Repeated"`
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributes) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributes) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributes) SetVpcAccessAttribute(v []*DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) *DescribeVpcAccessesResponseBodyVpcAccessAttributes {
	s.VpcAccessAttribute = v
	return s
}

type DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute struct {
	// VPC授权的创建时间
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// VPC授权的描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// VPC中的后端服务信息
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// VPC授权名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// VPC中的后端服务端口
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// 地域id
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// vpc授权ID
	VpcAccessId *string `json:"VpcAccessId,omitempty" xml:"VpcAccessId,omitempty"`
	// VPC的ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetCreatedTime(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetDescription(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Description = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetInstanceId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetName(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Name = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetPort(v int32) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.Port = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetRegionId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetVpcAccessId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.VpcAccessId = &v
	return s
}

func (s *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute) SetVpcId(v string) *DescribeVpcAccessesResponseBodyVpcAccessAttributesVpcAccessAttribute {
	s.VpcId = &v
	return s
}

type DescribeVpcAccessesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeVpcAccessesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeVpcAccessesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVpcAccessesResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpcAccessesResponse) SetHeaders(v map[string]*string) *DescribeVpcAccessesResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpcAccessesResponse) SetBody(v *DescribeVpcAccessesResponseBody) *DescribeVpcAccessesResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	Language      *string `json:"Language,omitempty" xml:"Language,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetLanguage(v string) *DescribeZonesRequest {
	s.Language = &v
	return s
}

func (s *DescribeZonesRequest) SetSecurityToken(v string) *DescribeZonesRequest {
	s.SecurityToken = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v *DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	Zone []*DescribeZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetZone(v []*DescribeZonesResponseBodyZonesZone) *DescribeZonesResponseBodyZones {
	s.Zone = v
	return s
}

type DescribeZonesResponseBodyZonesZone struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZonesZone) SetLocalName(v string) *DescribeZonesResponseBodyZonesZone {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZonesZone) SetZoneId(v string) *DescribeZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type DryRunSwaggerRequest struct {
	Data            *string                `json:"Data,omitempty" xml:"Data,omitempty"`
	DataFormat      *string                `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	GlobalCondition map[string]interface{} `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	GroupId         *string                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Overwrite       *bool                  `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	SecurityToken   *string                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DryRunSwaggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerRequest) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerRequest) SetData(v string) *DryRunSwaggerRequest {
	s.Data = &v
	return s
}

func (s *DryRunSwaggerRequest) SetDataFormat(v string) *DryRunSwaggerRequest {
	s.DataFormat = &v
	return s
}

func (s *DryRunSwaggerRequest) SetGlobalCondition(v map[string]interface{}) *DryRunSwaggerRequest {
	s.GlobalCondition = v
	return s
}

func (s *DryRunSwaggerRequest) SetGroupId(v string) *DryRunSwaggerRequest {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerRequest) SetOverwrite(v bool) *DryRunSwaggerRequest {
	s.Overwrite = &v
	return s
}

func (s *DryRunSwaggerRequest) SetSecurityToken(v string) *DryRunSwaggerRequest {
	s.SecurityToken = &v
	return s
}

type DryRunSwaggerShrinkRequest struct {
	Data                  *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataFormat            *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	GlobalConditionShrink *string `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Overwrite             *bool   `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DryRunSwaggerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerShrinkRequest) SetData(v string) *DryRunSwaggerShrinkRequest {
	s.Data = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetDataFormat(v string) *DryRunSwaggerShrinkRequest {
	s.DataFormat = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetGlobalConditionShrink(v string) *DryRunSwaggerShrinkRequest {
	s.GlobalConditionShrink = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetGroupId(v string) *DryRunSwaggerShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetOverwrite(v bool) *DryRunSwaggerShrinkRequest {
	s.Overwrite = &v
	return s
}

func (s *DryRunSwaggerShrinkRequest) SetSecurityToken(v string) *DryRunSwaggerShrinkRequest {
	s.SecurityToken = &v
	return s
}

type DryRunSwaggerResponseBody struct {
	Failed          *DryRunSwaggerResponseBodyFailed       `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Struct"`
	GlobalCondition *string                                `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	ModelFailed     *DryRunSwaggerResponseBodyModelFailed  `json:"ModelFailed,omitempty" xml:"ModelFailed,omitempty" type:"Struct"`
	ModelSuccess    *DryRunSwaggerResponseBodyModelSuccess `json:"ModelSuccess,omitempty" xml:"ModelSuccess,omitempty" type:"Struct"`
	RequestId       *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *DryRunSwaggerResponseBodySuccess      `json:"Success,omitempty" xml:"Success,omitempty" type:"Struct"`
}

func (s DryRunSwaggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBody) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBody) SetFailed(v *DryRunSwaggerResponseBodyFailed) *DryRunSwaggerResponseBody {
	s.Failed = v
	return s
}

func (s *DryRunSwaggerResponseBody) SetGlobalCondition(v string) *DryRunSwaggerResponseBody {
	s.GlobalCondition = &v
	return s
}

func (s *DryRunSwaggerResponseBody) SetModelFailed(v *DryRunSwaggerResponseBodyModelFailed) *DryRunSwaggerResponseBody {
	s.ModelFailed = v
	return s
}

func (s *DryRunSwaggerResponseBody) SetModelSuccess(v *DryRunSwaggerResponseBodyModelSuccess) *DryRunSwaggerResponseBody {
	s.ModelSuccess = v
	return s
}

func (s *DryRunSwaggerResponseBody) SetRequestId(v string) *DryRunSwaggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DryRunSwaggerResponseBody) SetSuccess(v *DryRunSwaggerResponseBodySuccess) *DryRunSwaggerResponseBody {
	s.Success = v
	return s
}

type DryRunSwaggerResponseBodyFailed struct {
	ApiImportSwaggerFailed []*DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed `json:"ApiImportSwaggerFailed,omitempty" xml:"ApiImportSwaggerFailed,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodyFailed) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodyFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyFailed) SetApiImportSwaggerFailed(v []*DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) *DryRunSwaggerResponseBodyFailed {
	s.ApiImportSwaggerFailed = v
	return s
}

type DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed struct {
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) SetErrorMsg(v string) *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.ErrorMsg = &v
	return s
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) SetHttpMethod(v string) *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.HttpMethod = &v
	return s
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) SetPath(v string) *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.Path = &v
	return s
}

type DryRunSwaggerResponseBodyModelFailed struct {
	ApiImportModelFailed []*DryRunSwaggerResponseBodyModelFailedApiImportModelFailed `json:"ApiImportModelFailed,omitempty" xml:"ApiImportModelFailed,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodyModelFailed) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelFailed) SetApiImportModelFailed(v []*DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) *DryRunSwaggerResponseBodyModelFailed {
	s.ApiImportModelFailed = v
	return s
}

type DryRunSwaggerResponseBodyModelFailedApiImportModelFailed struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) SetErrorMsg(v string) *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ErrorMsg = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) SetGroupId(v string) *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) SetModelName(v string) *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ModelName = &v
	return s
}

type DryRunSwaggerResponseBodyModelSuccess struct {
	ApiImportModelSuccess []*DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess `json:"ApiImportModelSuccess,omitempty" xml:"ApiImportModelSuccess,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodyModelSuccess) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelSuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelSuccess) SetApiImportModelSuccess(v []*DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) *DryRunSwaggerResponseBodyModelSuccess {
	s.ApiImportModelSuccess = v
	return s
}

type DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess struct {
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName      *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelOperation *string `json:"ModelOperation,omitempty" xml:"ModelOperation,omitempty"`
	ModelUid       *string `json:"ModelUid,omitempty" xml:"ModelUid,omitempty"`
}

func (s DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetGroupId(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelName(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelName = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelOperation(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelOperation = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelUid(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelUid = &v
	return s
}

type DryRunSwaggerResponseBodySuccess struct {
	ApiDryRunSwaggerSuccess []*DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess `json:"ApiDryRunSwaggerSuccess,omitempty" xml:"ApiDryRunSwaggerSuccess,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodySuccess) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodySuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodySuccess) SetApiDryRunSwaggerSuccess(v []*DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) *DryRunSwaggerResponseBodySuccess {
	s.ApiDryRunSwaggerSuccess = v
	return s
}

type DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess struct {
	ApiOperation *string `json:"ApiOperation,omitempty" xml:"ApiOperation,omitempty"`
	ApiSwagger   *string `json:"ApiSwagger,omitempty" xml:"ApiSwagger,omitempty"`
	ApiUid       *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	HttpMethod   *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetApiOperation(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.ApiOperation = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetApiSwagger(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.ApiSwagger = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetApiUid(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.ApiUid = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetHttpMethod(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.HttpMethod = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetPath(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.Path = &v
	return s
}

type DryRunSwaggerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DryRunSwaggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DryRunSwaggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DryRunSwaggerResponse) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponse) SetHeaders(v map[string]*string) *DryRunSwaggerResponse {
	s.Headers = v
	return s
}

func (s *DryRunSwaggerResponse) SetBody(v *DryRunSwaggerResponseBody) *DryRunSwaggerResponse {
	s.Body = v
	return s
}

type ImportSwaggerRequest struct {
	Data            *string                `json:"Data,omitempty" xml:"Data,omitempty"`
	DataFormat      *string                `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DryRun          *bool                  `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	GlobalCondition map[string]interface{} `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	GroupId         *string                `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Overwrite       *bool                  `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	SecurityToken   *string                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ImportSwaggerRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerRequest) GoString() string {
	return s.String()
}

func (s *ImportSwaggerRequest) SetData(v string) *ImportSwaggerRequest {
	s.Data = &v
	return s
}

func (s *ImportSwaggerRequest) SetDataFormat(v string) *ImportSwaggerRequest {
	s.DataFormat = &v
	return s
}

func (s *ImportSwaggerRequest) SetDryRun(v bool) *ImportSwaggerRequest {
	s.DryRun = &v
	return s
}

func (s *ImportSwaggerRequest) SetGlobalCondition(v map[string]interface{}) *ImportSwaggerRequest {
	s.GlobalCondition = v
	return s
}

func (s *ImportSwaggerRequest) SetGroupId(v string) *ImportSwaggerRequest {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerRequest) SetOverwrite(v bool) *ImportSwaggerRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportSwaggerRequest) SetSecurityToken(v string) *ImportSwaggerRequest {
	s.SecurityToken = &v
	return s
}

type ImportSwaggerShrinkRequest struct {
	Data                  *string `json:"Data,omitempty" xml:"Data,omitempty"`
	DataFormat            *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	DryRun                *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	GlobalConditionShrink *string `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Overwrite             *bool   `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ImportSwaggerShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportSwaggerShrinkRequest) SetData(v string) *ImportSwaggerShrinkRequest {
	s.Data = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetDataFormat(v string) *ImportSwaggerShrinkRequest {
	s.DataFormat = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetDryRun(v bool) *ImportSwaggerShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetGlobalConditionShrink(v string) *ImportSwaggerShrinkRequest {
	s.GlobalConditionShrink = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetGroupId(v string) *ImportSwaggerShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetOverwrite(v bool) *ImportSwaggerShrinkRequest {
	s.Overwrite = &v
	return s
}

func (s *ImportSwaggerShrinkRequest) SetSecurityToken(v string) *ImportSwaggerShrinkRequest {
	s.SecurityToken = &v
	return s
}

type ImportSwaggerResponseBody struct {
	Failed       *ImportSwaggerResponseBodyFailed       `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Struct"`
	ModelFailed  *ImportSwaggerResponseBodyModelFailed  `json:"ModelFailed,omitempty" xml:"ModelFailed,omitempty" type:"Struct"`
	ModelSuccess *ImportSwaggerResponseBodyModelSuccess `json:"ModelSuccess,omitempty" xml:"ModelSuccess,omitempty" type:"Struct"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *ImportSwaggerResponseBodySuccess      `json:"Success,omitempty" xml:"Success,omitempty" type:"Struct"`
}

func (s ImportSwaggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBody) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBody) SetFailed(v *ImportSwaggerResponseBodyFailed) *ImportSwaggerResponseBody {
	s.Failed = v
	return s
}

func (s *ImportSwaggerResponseBody) SetModelFailed(v *ImportSwaggerResponseBodyModelFailed) *ImportSwaggerResponseBody {
	s.ModelFailed = v
	return s
}

func (s *ImportSwaggerResponseBody) SetModelSuccess(v *ImportSwaggerResponseBodyModelSuccess) *ImportSwaggerResponseBody {
	s.ModelSuccess = v
	return s
}

func (s *ImportSwaggerResponseBody) SetRequestId(v string) *ImportSwaggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportSwaggerResponseBody) SetSuccess(v *ImportSwaggerResponseBodySuccess) *ImportSwaggerResponseBody {
	s.Success = v
	return s
}

type ImportSwaggerResponseBodyFailed struct {
	ApiImportSwaggerFailed []*ImportSwaggerResponseBodyFailedApiImportSwaggerFailed `json:"ApiImportSwaggerFailed,omitempty" xml:"ApiImportSwaggerFailed,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodyFailed) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodyFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyFailed) SetApiImportSwaggerFailed(v []*ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) *ImportSwaggerResponseBodyFailed {
	s.ApiImportSwaggerFailed = v
	return s
}

type ImportSwaggerResponseBodyFailedApiImportSwaggerFailed struct {
	ErrorMsg   *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	Path       *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) SetErrorMsg(v string) *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.ErrorMsg = &v
	return s
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) SetHttpMethod(v string) *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.HttpMethod = &v
	return s
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) SetPath(v string) *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.Path = &v
	return s
}

type ImportSwaggerResponseBodyModelFailed struct {
	ApiImportModelFailed []*ImportSwaggerResponseBodyModelFailedApiImportModelFailed `json:"ApiImportModelFailed,omitempty" xml:"ApiImportModelFailed,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodyModelFailed) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelFailed) SetApiImportModelFailed(v []*ImportSwaggerResponseBodyModelFailedApiImportModelFailed) *ImportSwaggerResponseBodyModelFailed {
	s.ApiImportModelFailed = v
	return s
}

type ImportSwaggerResponseBodyModelFailedApiImportModelFailed struct {
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s ImportSwaggerResponseBodyModelFailedApiImportModelFailed) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelFailedApiImportModelFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) SetErrorMsg(v string) *ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ErrorMsg = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) SetGroupId(v string) *ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) SetModelName(v string) *ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ModelName = &v
	return s
}

type ImportSwaggerResponseBodyModelSuccess struct {
	ApiImportModelSuccess []*ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess `json:"ApiImportModelSuccess,omitempty" xml:"ApiImportModelSuccess,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodyModelSuccess) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelSuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelSuccess) SetApiImportModelSuccess(v []*ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) *ImportSwaggerResponseBodyModelSuccess {
	s.ApiImportModelSuccess = v
	return s
}

type ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess struct {
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName      *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelOperation *string `json:"ModelOperation,omitempty" xml:"ModelOperation,omitempty"`
	ModelUid       *string `json:"ModelUid,omitempty" xml:"ModelUid,omitempty"`
}

func (s ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetGroupId(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelName(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelName = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelOperation(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelOperation = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelUid(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelUid = &v
	return s
}

type ImportSwaggerResponseBodySuccess struct {
	ApiImportSwaggerSuccess []*ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess `json:"ApiImportSwaggerSuccess,omitempty" xml:"ApiImportSwaggerSuccess,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodySuccess) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodySuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodySuccess) SetApiImportSwaggerSuccess(v []*ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) *ImportSwaggerResponseBodySuccess {
	s.ApiImportSwaggerSuccess = v
	return s
}

type ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess struct {
	ApiOperation *string `json:"ApiOperation,omitempty" xml:"ApiOperation,omitempty"`
	ApiUid       *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	HttpMethod   *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	Path         *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetApiOperation(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.ApiOperation = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetApiUid(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.ApiUid = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetHttpMethod(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.HttpMethod = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetPath(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.Path = &v
	return s
}

type ImportSwaggerResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportSwaggerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportSwaggerResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportSwaggerResponse) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponse) SetHeaders(v map[string]*string) *ImportSwaggerResponse {
	s.Headers = v
	return s
}

func (s *ImportSwaggerResponse) SetBody(v *ImportSwaggerResponseBody) *ImportSwaggerResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyApiRequest struct {
	AllowSignatureMethod *string `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	ApiId                *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	ApiName              *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	AppCodeAuthType      *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	AuthType             *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	ConstantParameters   *string `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisableInternet      *bool   `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples     *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	FailResultSample     *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	ForceNonceCheck      *bool   `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	GroupId              *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	OpenIdConnectConfig  *string `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty"`
	RequestConfig        *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParameters    *string `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty"`
	ResultBodyModel      *string `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	ResultDescriptions   *string `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty"`
	ResultSample         *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	ResultType           *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	SystemParameters     *string `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty"`
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

func (s *ModifyApiRequest) SetAppCodeAuthType(v string) *ModifyApiRequest {
	s.AppCodeAuthType = &v
	return s
}

func (s *ModifyApiRequest) SetAuthType(v string) *ModifyApiRequest {
	s.AuthType = &v
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

func (s *ModifyApiRequest) SetDisableInternet(v bool) *ModifyApiRequest {
	s.DisableInternet = &v
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

func (s *ModifyApiRequest) SetForceNonceCheck(v bool) *ModifyApiRequest {
	s.ForceNonceCheck = &v
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

func (s *ModifyApiRequest) SetRequestParameters(v string) *ModifyApiRequest {
	s.RequestParameters = &v
	return s
}

func (s *ModifyApiRequest) SetResultBodyModel(v string) *ModifyApiRequest {
	s.ResultBodyModel = &v
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

func (s *ModifyApiRequest) SetSystemParameters(v string) *ModifyApiRequest {
	s.SystemParameters = &v
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyApiResponse) SetBody(v *ModifyApiResponseBody) *ModifyApiResponse {
	s.Body = v
	return s
}

type ModifyApiGroupRequest struct {
	BasePath           *string                     `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	CompatibleFlags    *string                     `json:"CompatibleFlags,omitempty" xml:"CompatibleFlags,omitempty"`
	CustomTraceConfig  *string                     `json:"CustomTraceConfig,omitempty" xml:"CustomTraceConfig,omitempty"`
	CustomerConfigs    *string                     `json:"CustomerConfigs,omitempty" xml:"CustomerConfigs,omitempty"`
	DefaultDomain      *string                     `json:"DefaultDomain,omitempty" xml:"DefaultDomain,omitempty"`
	Description        *string                     `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId            *string                     `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName          *string                     `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	PassthroughHeaders *string                     `json:"PassthroughHeaders,omitempty" xml:"PassthroughHeaders,omitempty"`
	RpcPattern         *string                     `json:"RpcPattern,omitempty" xml:"RpcPattern,omitempty"`
	SecurityToken      *string                     `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag                []*ModifyApiGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	UserLogConfig      *string                     `json:"UserLogConfig,omitempty" xml:"UserLogConfig,omitempty"`
}

func (s ModifyApiGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupRequest) SetBasePath(v string) *ModifyApiGroupRequest {
	s.BasePath = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCompatibleFlags(v string) *ModifyApiGroupRequest {
	s.CompatibleFlags = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCustomTraceConfig(v string) *ModifyApiGroupRequest {
	s.CustomTraceConfig = &v
	return s
}

func (s *ModifyApiGroupRequest) SetCustomerConfigs(v string) *ModifyApiGroupRequest {
	s.CustomerConfigs = &v
	return s
}

func (s *ModifyApiGroupRequest) SetDefaultDomain(v string) *ModifyApiGroupRequest {
	s.DefaultDomain = &v
	return s
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

func (s *ModifyApiGroupRequest) SetPassthroughHeaders(v string) *ModifyApiGroupRequest {
	s.PassthroughHeaders = &v
	return s
}

func (s *ModifyApiGroupRequest) SetRpcPattern(v string) *ModifyApiGroupRequest {
	s.RpcPattern = &v
	return s
}

func (s *ModifyApiGroupRequest) SetSecurityToken(v string) *ModifyApiGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiGroupRequest) SetTag(v []*ModifyApiGroupRequestTag) *ModifyApiGroupRequest {
	s.Tag = v
	return s
}

func (s *ModifyApiGroupRequest) SetUserLogConfig(v string) *ModifyApiGroupRequest {
	s.UserLogConfig = &v
	return s
}

type ModifyApiGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyApiGroupRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupRequestTag) SetKey(v string) *ModifyApiGroupRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyApiGroupRequestTag) SetValue(v string) *ModifyApiGroupRequestTag {
	s.Value = &v
	return s
}

type ModifyApiGroupResponseBody struct {
	BasePath    *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
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

func (s *ModifyApiGroupResponseBody) SetBasePath(v string) *ModifyApiGroupResponseBody {
	s.BasePath = &v
	return s
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
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyApiGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyApiGroupResponse) SetBody(v *ModifyApiGroupResponseBody) *ModifyApiGroupResponse {
	s.Body = v
	return s
}

type ModifyApiGroupVpcWhitelistRequest struct {
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcIds        *string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty"`
}

func (s ModifyApiGroupVpcWhitelistRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupVpcWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupVpcWhitelistRequest) SetGroupId(v string) *ModifyApiGroupVpcWhitelistRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistRequest) SetSecurityToken(v string) *ModifyApiGroupVpcWhitelistRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiGroupVpcWhitelistRequest) SetVpcIds(v string) *ModifyApiGroupVpcWhitelistRequest {
	s.VpcIds = &v
	return s
}

type ModifyApiGroupVpcWhitelistResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyApiGroupVpcWhitelistResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupVpcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupVpcWhitelistResponseBody) SetRequestId(v string) *ModifyApiGroupVpcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

type ModifyApiGroupVpcWhitelistResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyApiGroupVpcWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyApiGroupVpcWhitelistResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApiGroupVpcWhitelistResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiGroupVpcWhitelistResponse) SetHeaders(v map[string]*string) *ModifyApiGroupVpcWhitelistResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiGroupVpcWhitelistResponse) SetBody(v *ModifyApiGroupVpcWhitelistResponseBody) *ModifyApiGroupVpcWhitelistResponse {
	s.Body = v
	return s
}

type ModifyAppRequest struct {
	AppId         *int64                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName       *string                `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Description   *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string                `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*ModifyAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ModifyAppRequest) SetTag(v []*ModifyAppRequestTag) *ModifyAppRequest {
	s.Tag = v
	return s
}

type ModifyAppRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyAppRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyAppRequestTag) SetKey(v string) *ModifyAppRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyAppRequestTag) SetValue(v string) *ModifyAppRequestTag {
	s.Value = &v
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyInstanceSpecRequest struct {
	AutoPay      *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	Token        *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ModifyInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecRequest) SetAutoPay(v bool) *ModifyInstanceSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceId(v string) *ModifyInstanceSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetInstanceSpec(v string) *ModifyInstanceSpecRequest {
	s.InstanceSpec = &v
	return s
}

func (s *ModifyInstanceSpecRequest) SetToken(v string) *ModifyInstanceSpecRequest {
	s.Token = &v
	return s
}

type ModifyInstanceSpecResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponseBody) SetRequestId(v string) *ModifyInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceSpecResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceSpecResponse) SetBody(v *ModifyInstanceSpecResponseBody) *ModifyInstanceSpecResponse {
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyLogConfigResponse) SetBody(v *ModifyLogConfigResponseBody) *ModifyLogConfigResponse {
	s.Body = v
	return s
}

type ModifyModelRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId      *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	ModelName    *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	NewModelName *string `json:"NewModelName,omitempty" xml:"NewModelName,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s ModifyModelRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyModelRequest) GoString() string {
	return s.String()
}

func (s *ModifyModelRequest) SetDescription(v string) *ModifyModelRequest {
	s.Description = &v
	return s
}

func (s *ModifyModelRequest) SetGroupId(v string) *ModifyModelRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyModelRequest) SetModelName(v string) *ModifyModelRequest {
	s.ModelName = &v
	return s
}

func (s *ModifyModelRequest) SetNewModelName(v string) *ModifyModelRequest {
	s.NewModelName = &v
	return s
}

func (s *ModifyModelRequest) SetSchema(v string) *ModifyModelRequest {
	s.Schema = &v
	return s
}

type ModifyModelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyModelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelResponseBody) SetRequestId(v string) *ModifyModelResponseBody {
	s.RequestId = &v
	return s
}

type ModifyModelResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyModelResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyModelResponse) GoString() string {
	return s.String()
}

func (s *ModifyModelResponse) SetHeaders(v map[string]*string) *ModifyModelResponse {
	s.Headers = v
	return s
}

func (s *ModifyModelResponse) SetBody(v *ModifyModelResponseBody) *ModifyModelResponse {
	s.Body = v
	return s
}

type ModifyPluginRequest struct {
	Description   *string                   `json:"Description,omitempty" xml:"Description,omitempty"`
	PluginData    *string                   `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	PluginId      *string                   `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	PluginName    *string                   `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	SecurityToken *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*ModifyPluginRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ModifyPluginRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyPluginRequest) GoString() string {
	return s.String()
}

func (s *ModifyPluginRequest) SetDescription(v string) *ModifyPluginRequest {
	s.Description = &v
	return s
}

func (s *ModifyPluginRequest) SetPluginData(v string) *ModifyPluginRequest {
	s.PluginData = &v
	return s
}

func (s *ModifyPluginRequest) SetPluginId(v string) *ModifyPluginRequest {
	s.PluginId = &v
	return s
}

func (s *ModifyPluginRequest) SetPluginName(v string) *ModifyPluginRequest {
	s.PluginName = &v
	return s
}

func (s *ModifyPluginRequest) SetSecurityToken(v string) *ModifyPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyPluginRequest) SetTag(v []*ModifyPluginRequestTag) *ModifyPluginRequest {
	s.Tag = v
	return s
}

type ModifyPluginRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyPluginRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ModifyPluginRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyPluginRequestTag) SetKey(v string) *ModifyPluginRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyPluginRequestTag) SetValue(v string) *ModifyPluginRequestTag {
	s.Value = &v
	return s
}

type ModifyPluginResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyPluginResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyPluginResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyPluginResponseBody) SetRequestId(v string) *ModifyPluginResponseBody {
	s.RequestId = &v
	return s
}

type ModifyPluginResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyPluginResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyPluginResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyPluginResponse) GoString() string {
	return s.String()
}

func (s *ModifyPluginResponse) SetHeaders(v map[string]*string) *ModifyPluginResponse {
	s.Headers = v
	return s
}

func (s *ModifyPluginResponse) SetBody(v *ModifyPluginResponseBody) *ModifyPluginResponse {
	s.Body = v
	return s
}

type ModifySignatureRequest struct {
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureId     *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureKey    *string `json:"SignatureKey,omitempty" xml:"SignatureKey,omitempty"`
	SignatureName   *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
	SignatureSecret *string `json:"SignatureSecret,omitempty" xml:"SignatureSecret,omitempty"`
}

func (s ModifySignatureRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySignatureRequest) GoString() string {
	return s.String()
}

func (s *ModifySignatureRequest) SetSecurityToken(v string) *ModifySignatureRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureId(v string) *ModifySignatureRequest {
	s.SignatureId = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureKey(v string) *ModifySignatureRequest {
	s.SignatureKey = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureName(v string) *ModifySignatureRequest {
	s.SignatureName = &v
	return s
}

func (s *ModifySignatureRequest) SetSignatureSecret(v string) *ModifySignatureRequest {
	s.SignatureSecret = &v
	return s
}

type ModifySignatureResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	SignatureName *string `json:"SignatureName,omitempty" xml:"SignatureName,omitempty"`
}

func (s ModifySignatureResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySignatureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySignatureResponseBody) SetRequestId(v string) *ModifySignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySignatureResponseBody) SetSignatureId(v string) *ModifySignatureResponseBody {
	s.SignatureId = &v
	return s
}

func (s *ModifySignatureResponseBody) SetSignatureName(v string) *ModifySignatureResponseBody {
	s.SignatureName = &v
	return s
}

type ModifySignatureResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySignatureResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySignatureResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySignatureResponse) GoString() string {
	return s.String()
}

func (s *ModifySignatureResponse) SetHeaders(v map[string]*string) *ModifySignatureResponse {
	s.Headers = v
	return s
}

func (s *ModifySignatureResponse) SetBody(v *ModifySignatureResponseBody) *ModifySignatureResponse {
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyTrafficControlResponse) SetBody(v *ModifyTrafficControlResponseBody) *ModifyTrafficControlResponse {
	s.Body = v
	return s
}

type OpenApiGatewayServiceResponseBody struct {
	OrderId   *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenApiGatewayServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenApiGatewayServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenApiGatewayServiceResponseBody) SetOrderId(v string) *OpenApiGatewayServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *OpenApiGatewayServiceResponseBody) SetRequestId(v string) *OpenApiGatewayServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenApiGatewayServiceResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *OpenApiGatewayServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s OpenApiGatewayServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenApiGatewayServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenApiGatewayServiceResponse) SetHeaders(v map[string]*string) *OpenApiGatewayServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenApiGatewayServiceResponse) SetBody(v *OpenApiGatewayServiceResponseBody) *OpenApiGatewayServiceResponse {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReactivateDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ReactivateDomainResponse) SetBody(v *ReactivateDomainResponseBody) *ReactivateDomainResponse {
	s.Body = v
	return s
}

type RemoveApisAuthoritiesRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveApisAuthoritiesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveApisAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *RemoveApisAuthoritiesRequest) SetApiIds(v string) *RemoveApisAuthoritiesRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveApisAuthoritiesRequest) SetAppId(v int64) *RemoveApisAuthoritiesRequest {
	s.AppId = &v
	return s
}

func (s *RemoveApisAuthoritiesRequest) SetDescription(v string) *RemoveApisAuthoritiesRequest {
	s.Description = &v
	return s
}

func (s *RemoveApisAuthoritiesRequest) SetGroupId(v string) *RemoveApisAuthoritiesRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveApisAuthoritiesRequest) SetSecurityToken(v string) *RemoveApisAuthoritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveApisAuthoritiesRequest) SetStageName(v string) *RemoveApisAuthoritiesRequest {
	s.StageName = &v
	return s
}

type RemoveApisAuthoritiesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveApisAuthoritiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveApisAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveApisAuthoritiesResponseBody) SetRequestId(v string) *RemoveApisAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

type RemoveApisAuthoritiesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveApisAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveApisAuthoritiesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveApisAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *RemoveApisAuthoritiesResponse) SetHeaders(v map[string]*string) *RemoveApisAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *RemoveApisAuthoritiesResponse) SetBody(v *RemoveApisAuthoritiesResponseBody) *RemoveApisAuthoritiesResponse {
	s.Body = v
	return s
}

type RemoveAppsAuthoritiesRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveAppsAuthoritiesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppsAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *RemoveAppsAuthoritiesRequest) SetApiId(v string) *RemoveAppsAuthoritiesRequest {
	s.ApiId = &v
	return s
}

func (s *RemoveAppsAuthoritiesRequest) SetAppIds(v string) *RemoveAppsAuthoritiesRequest {
	s.AppIds = &v
	return s
}

func (s *RemoveAppsAuthoritiesRequest) SetGroupId(v string) *RemoveAppsAuthoritiesRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveAppsAuthoritiesRequest) SetSecurityToken(v string) *RemoveAppsAuthoritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveAppsAuthoritiesRequest) SetStageName(v string) *RemoveAppsAuthoritiesRequest {
	s.StageName = &v
	return s
}

type RemoveAppsAuthoritiesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAppsAuthoritiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppsAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppsAuthoritiesResponseBody) SetRequestId(v string) *RemoveAppsAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

type RemoveAppsAuthoritiesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAppsAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAppsAuthoritiesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppsAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppsAuthoritiesResponse) SetHeaders(v map[string]*string) *RemoveAppsAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppsAuthoritiesResponse) SetBody(v *RemoveAppsAuthoritiesResponseBody) *RemoveAppsAuthoritiesResponse {
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
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveIpControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveIpControlPolicyItemResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *RemoveIpControlPolicyItemResponse) SetBody(v *RemoveIpControlPolicyItemResponseBody) *RemoveIpControlPolicyItemResponse {
	s.Body = v
	return s
}

type RemoveSignatureApisRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s RemoveSignatureApisRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveSignatureApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveSignatureApisRequest) SetApiIds(v string) *RemoveSignatureApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetGroupId(v string) *RemoveSignatureApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetSecurityToken(v string) *RemoveSignatureApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetSignatureId(v string) *RemoveSignatureApisRequest {
	s.SignatureId = &v
	return s
}

func (s *RemoveSignatureApisRequest) SetStageName(v string) *RemoveSignatureApisRequest {
	s.StageName = &v
	return s
}

type RemoveSignatureApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveSignatureApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSignatureApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSignatureApisResponseBody) SetRequestId(v string) *RemoveSignatureApisResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSignatureApisResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveSignatureApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSignatureApisResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSignatureApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveSignatureApisResponse) SetHeaders(v map[string]*string) *RemoveSignatureApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveSignatureApisResponse) SetBody(v *RemoveSignatureApisResponseBody) *RemoveSignatureApisResponse {
	s.Body = v
	return s
}

type RemoveTrafficControlApisRequest struct {
	ApiIds           *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName        *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s RemoveTrafficControlApisRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTrafficControlApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveTrafficControlApisRequest) SetApiIds(v string) *RemoveTrafficControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetGroupId(v string) *RemoveTrafficControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetSecurityToken(v string) *RemoveTrafficControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetStageName(v string) *RemoveTrafficControlApisRequest {
	s.StageName = &v
	return s
}

func (s *RemoveTrafficControlApisRequest) SetTrafficControlId(v string) *RemoveTrafficControlApisRequest {
	s.TrafficControlId = &v
	return s
}

type RemoveTrafficControlApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTrafficControlApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTrafficControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTrafficControlApisResponseBody) SetRequestId(v string) *RemoveTrafficControlApisResponseBody {
	s.RequestId = &v
	return s
}

type RemoveTrafficControlApisResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveTrafficControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTrafficControlApisResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTrafficControlApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveTrafficControlApisResponse) SetHeaders(v map[string]*string) *RemoveTrafficControlApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveTrafficControlApisResponse) SetBody(v *RemoveTrafficControlApisResponseBody) *RemoveTrafficControlApisResponse {
	s.Body = v
	return s
}

type RemoveVpcAccessRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedBatchWork *bool   `json:"NeedBatchWork,omitempty" xml:"NeedBatchWork,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RemoveVpcAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessRequest) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessRequest) SetInstanceId(v string) *RemoveVpcAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetNeedBatchWork(v bool) *RemoveVpcAccessRequest {
	s.NeedBatchWork = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetPort(v int32) *RemoveVpcAccessRequest {
	s.Port = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetSecurityToken(v string) *RemoveVpcAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveVpcAccessRequest) SetVpcId(v string) *RemoveVpcAccessRequest {
	s.VpcId = &v
	return s
}

type RemoveVpcAccessResponseBody struct {
	Apis      *RemoveVpcAccessResponseBodyApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVpcAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponseBody) SetApis(v *RemoveVpcAccessResponseBodyApis) *RemoveVpcAccessResponseBody {
	s.Apis = v
	return s
}

func (s *RemoveVpcAccessResponseBody) SetRequestId(v string) *RemoveVpcAccessResponseBody {
	s.RequestId = &v
	return s
}

type RemoveVpcAccessResponseBodyApis struct {
	Api []*RemoveVpcAccessResponseBodyApisApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Repeated"`
}

func (s RemoveVpcAccessResponseBodyApis) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessResponseBodyApis) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponseBodyApis) SetApi(v []*RemoveVpcAccessResponseBodyApisApi) *RemoveVpcAccessResponseBodyApis {
	s.Api = v
	return s
}

type RemoveVpcAccessResponseBodyApisApi struct {
	ApiId   *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
}

func (s RemoveVpcAccessResponseBodyApisApi) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessResponseBodyApisApi) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponseBodyApisApi) SetApiId(v string) *RemoveVpcAccessResponseBodyApisApi {
	s.ApiId = &v
	return s
}

func (s *RemoveVpcAccessResponseBodyApisApi) SetGroupId(v string) *RemoveVpcAccessResponseBodyApisApi {
	s.GroupId = &v
	return s
}

func (s *RemoveVpcAccessResponseBodyApisApi) SetStageId(v string) *RemoveVpcAccessResponseBodyApisApi {
	s.StageId = &v
	return s
}

type RemoveVpcAccessResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveVpcAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveVpcAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessResponse) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessResponse) SetHeaders(v map[string]*string) *RemoveVpcAccessResponse {
	s.Headers = v
	return s
}

func (s *RemoveVpcAccessResponse) SetBody(v *RemoveVpcAccessResponseBody) *RemoveVpcAccessResponse {
	s.Body = v
	return s
}

type RemoveVpcAccessAndAbolishApisRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NeedBatchWork *bool   `json:"NeedBatchWork,omitempty" xml:"NeedBatchWork,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RemoveVpcAccessAndAbolishApisRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessAndAbolishApisRequest) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetInstanceId(v string) *RemoveVpcAccessAndAbolishApisRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetNeedBatchWork(v bool) *RemoveVpcAccessAndAbolishApisRequest {
	s.NeedBatchWork = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetPort(v int32) *RemoveVpcAccessAndAbolishApisRequest {
	s.Port = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetSecurityToken(v string) *RemoveVpcAccessAndAbolishApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisRequest) SetVpcId(v string) *RemoveVpcAccessAndAbolishApisRequest {
	s.VpcId = &v
	return s
}

type RemoveVpcAccessAndAbolishApisResponseBody struct {
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveVpcAccessAndAbolishApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessAndAbolishApisResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) SetOperationId(v string) *RemoveVpcAccessAndAbolishApisResponseBody {
	s.OperationId = &v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponseBody) SetRequestId(v string) *RemoveVpcAccessAndAbolishApisResponseBody {
	s.RequestId = &v
	return s
}

type RemoveVpcAccessAndAbolishApisResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveVpcAccessAndAbolishApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveVpcAccessAndAbolishApisResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveVpcAccessAndAbolishApisResponse) GoString() string {
	return s.String()
}

func (s *RemoveVpcAccessAndAbolishApisResponse) SetHeaders(v map[string]*string) *RemoveVpcAccessAndAbolishApisResponse {
	s.Headers = v
	return s
}

func (s *RemoveVpcAccessAndAbolishApisResponse) SetBody(v *RemoveVpcAccessAndAbolishApisResponseBody) *RemoveVpcAccessAndAbolishApisResponse {
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
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAppCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetAppCodeResponse) SetBody(v *ResetAppCodeResponseBody) *ResetAppCodeResponse {
	s.Body = v
	return s
}

type ResetAppSecretRequest struct {
	AppKey        *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ResetAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAppSecretRequest) GoString() string {
	return s.String()
}

func (s *ResetAppSecretRequest) SetAppKey(v string) *ResetAppSecretRequest {
	s.AppKey = &v
	return s
}

func (s *ResetAppSecretRequest) SetSecurityToken(v string) *ResetAppSecretRequest {
	s.SecurityToken = &v
	return s
}

type ResetAppSecretResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetAppSecretResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetAppSecretResponseBody) GoString() string {
	return s.String()
}

func (s *ResetAppSecretResponseBody) SetRequestId(v string) *ResetAppSecretResponseBody {
	s.RequestId = &v
	return s
}

type ResetAppSecretResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAppSecretResponse) GoString() string {
	return s.String()
}

func (s *ResetAppSecretResponse) SetHeaders(v map[string]*string) *ResetAppSecretResponse {
	s.Headers = v
	return s
}

func (s *ResetAppSecretResponse) SetBody(v *ResetAppSecretResponseBody) *ResetAppSecretResponse {
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SdkGenerateByAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SdkGenerateByGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SdkGenerateByGroupResponse) SetBody(v *SdkGenerateByGroupResponseBody) *SdkGenerateByGroupResponse {
	s.Body = v
	return s
}

type SetApisAuthoritiesRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AuthValidTime *string `json:"AuthValidTime,omitempty" xml:"AuthValidTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetApisAuthoritiesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetApisAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *SetApisAuthoritiesRequest) SetApiIds(v string) *SetApisAuthoritiesRequest {
	s.ApiIds = &v
	return s
}

func (s *SetApisAuthoritiesRequest) SetAppId(v int64) *SetApisAuthoritiesRequest {
	s.AppId = &v
	return s
}

func (s *SetApisAuthoritiesRequest) SetAuthValidTime(v string) *SetApisAuthoritiesRequest {
	s.AuthValidTime = &v
	return s
}

func (s *SetApisAuthoritiesRequest) SetDescription(v string) *SetApisAuthoritiesRequest {
	s.Description = &v
	return s
}

func (s *SetApisAuthoritiesRequest) SetGroupId(v string) *SetApisAuthoritiesRequest {
	s.GroupId = &v
	return s
}

func (s *SetApisAuthoritiesRequest) SetSecurityToken(v string) *SetApisAuthoritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetApisAuthoritiesRequest) SetStageName(v string) *SetApisAuthoritiesRequest {
	s.StageName = &v
	return s
}

type SetApisAuthoritiesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApisAuthoritiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetApisAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *SetApisAuthoritiesResponseBody) SetRequestId(v string) *SetApisAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

type SetApisAuthoritiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetApisAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetApisAuthoritiesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetApisAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *SetApisAuthoritiesResponse) SetHeaders(v map[string]*string) *SetApisAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *SetApisAuthoritiesResponse) SetBody(v *SetApisAuthoritiesResponseBody) *SetApisAuthoritiesResponse {
	s.Body = v
	return s
}

type SetAppsAuthoritiesRequest struct {
	ApiId         *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	AppIds        *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	AuthValidTime *string `json:"AuthValidTime,omitempty" xml:"AuthValidTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetAppsAuthoritiesRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAppsAuthoritiesRequest) GoString() string {
	return s.String()
}

func (s *SetAppsAuthoritiesRequest) SetApiId(v string) *SetAppsAuthoritiesRequest {
	s.ApiId = &v
	return s
}

func (s *SetAppsAuthoritiesRequest) SetAppIds(v string) *SetAppsAuthoritiesRequest {
	s.AppIds = &v
	return s
}

func (s *SetAppsAuthoritiesRequest) SetAuthValidTime(v string) *SetAppsAuthoritiesRequest {
	s.AuthValidTime = &v
	return s
}

func (s *SetAppsAuthoritiesRequest) SetDescription(v string) *SetAppsAuthoritiesRequest {
	s.Description = &v
	return s
}

func (s *SetAppsAuthoritiesRequest) SetGroupId(v string) *SetAppsAuthoritiesRequest {
	s.GroupId = &v
	return s
}

func (s *SetAppsAuthoritiesRequest) SetSecurityToken(v string) *SetAppsAuthoritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetAppsAuthoritiesRequest) SetStageName(v string) *SetAppsAuthoritiesRequest {
	s.StageName = &v
	return s
}

type SetAppsAuthoritiesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAppsAuthoritiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAppsAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *SetAppsAuthoritiesResponseBody) SetRequestId(v string) *SetAppsAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

type SetAppsAuthoritiesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAppsAuthoritiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAppsAuthoritiesResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAppsAuthoritiesResponse) GoString() string {
	return s.String()
}

func (s *SetAppsAuthoritiesResponse) SetHeaders(v map[string]*string) *SetAppsAuthoritiesResponse {
	s.Headers = v
	return s
}

func (s *SetAppsAuthoritiesResponse) SetBody(v *SetAppsAuthoritiesResponseBody) *SetAppsAuthoritiesResponse {
	s.Body = v
	return s
}

type SetDomainRequest struct {
	BindStageName    *string `json:"BindStageName,omitempty" xml:"BindStageName,omitempty"`
	CustomDomainType *string `json:"CustomDomainType,omitempty" xml:"CustomDomainType,omitempty"`
	DomainName       *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IsForce          *bool   `json:"IsForce,omitempty" xml:"IsForce,omitempty"`
}

func (s SetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDomainRequest) SetBindStageName(v string) *SetDomainRequest {
	s.BindStageName = &v
	return s
}

func (s *SetDomainRequest) SetCustomDomainType(v string) *SetDomainRequest {
	s.CustomDomainType = &v
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

func (s *SetDomainRequest) SetIsForce(v bool) *SetDomainRequest {
	s.IsForce = &v
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetDomainResponse) SetBody(v *SetDomainResponseBody) *SetDomainResponse {
	s.Body = v
	return s
}

type SetDomainCertificateRequest struct {
	CaCertificateBody     *string `json:"CaCertificateBody,omitempty" xml:"CaCertificateBody,omitempty"`
	CertificateBody       *string `json:"CertificateBody,omitempty" xml:"CertificateBody,omitempty"`
	CertificateName       *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	CertificatePrivateKey *string `json:"CertificatePrivateKey,omitempty" xml:"CertificatePrivateKey,omitempty"`
	DomainName            *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId               *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken         *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDomainCertificateRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SetDomainCertificateRequest) SetCaCertificateBody(v string) *SetDomainCertificateRequest {
	s.CaCertificateBody = &v
	return s
}

func (s *SetDomainCertificateRequest) SetCertificateBody(v string) *SetDomainCertificateRequest {
	s.CertificateBody = &v
	return s
}

func (s *SetDomainCertificateRequest) SetCertificateName(v string) *SetDomainCertificateRequest {
	s.CertificateName = &v
	return s
}

func (s *SetDomainCertificateRequest) SetCertificatePrivateKey(v string) *SetDomainCertificateRequest {
	s.CertificatePrivateKey = &v
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
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetDomainWebSocketStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetIpControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SetIpControlApisResponse) SetBody(v *SetIpControlApisResponseBody) *SetIpControlApisResponse {
	s.Body = v
	return s
}

type SetSignatureApisRequest struct {
	ApiIds        *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SignatureId   *string `json:"SignatureId,omitempty" xml:"SignatureId,omitempty"`
	StageName     *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s SetSignatureApisRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSignatureApisRequest) GoString() string {
	return s.String()
}

func (s *SetSignatureApisRequest) SetApiIds(v string) *SetSignatureApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetSignatureApisRequest) SetGroupId(v string) *SetSignatureApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetSignatureApisRequest) SetSecurityToken(v string) *SetSignatureApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetSignatureApisRequest) SetSignatureId(v string) *SetSignatureApisRequest {
	s.SignatureId = &v
	return s
}

func (s *SetSignatureApisRequest) SetStageName(v string) *SetSignatureApisRequest {
	s.StageName = &v
	return s
}

type SetSignatureApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSignatureApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetSignatureApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetSignatureApisResponseBody) SetRequestId(v string) *SetSignatureApisResponseBody {
	s.RequestId = &v
	return s
}

type SetSignatureApisResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetSignatureApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetSignatureApisResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSignatureApisResponse) GoString() string {
	return s.String()
}

func (s *SetSignatureApisResponse) SetHeaders(v map[string]*string) *SetSignatureApisResponse {
	s.Headers = v
	return s
}

func (s *SetSignatureApisResponse) SetBody(v *SetSignatureApisResponseBody) *SetSignatureApisResponse {
	s.Body = v
	return s
}

type SetTrafficControlApisRequest struct {
	ApiIds           *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	GroupId          *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken    *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StageName        *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	TrafficControlId *string `json:"TrafficControlId,omitempty" xml:"TrafficControlId,omitempty"`
}

func (s SetTrafficControlApisRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTrafficControlApisRequest) GoString() string {
	return s.String()
}

func (s *SetTrafficControlApisRequest) SetApiIds(v string) *SetTrafficControlApisRequest {
	s.ApiIds = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetGroupId(v string) *SetTrafficControlApisRequest {
	s.GroupId = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetSecurityToken(v string) *SetTrafficControlApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetStageName(v string) *SetTrafficControlApisRequest {
	s.StageName = &v
	return s
}

func (s *SetTrafficControlApisRequest) SetTrafficControlId(v string) *SetTrafficControlApisRequest {
	s.TrafficControlId = &v
	return s
}

type SetTrafficControlApisResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTrafficControlApisResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTrafficControlApisResponseBody) GoString() string {
	return s.String()
}

func (s *SetTrafficControlApisResponseBody) SetRequestId(v string) *SetTrafficControlApisResponseBody {
	s.RequestId = &v
	return s
}

type SetTrafficControlApisResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetTrafficControlApisResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetTrafficControlApisResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTrafficControlApisResponse) GoString() string {
	return s.String()
}

func (s *SetTrafficControlApisResponse) SetHeaders(v map[string]*string) *SetTrafficControlApisResponse {
	s.Headers = v
	return s
}

func (s *SetTrafficControlApisResponse) SetBody(v *SetTrafficControlApisResponseBody) *SetTrafficControlApisResponse {
	s.Body = v
	return s
}

type SetVpcAccessRequest struct {
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SetVpcAccessRequest) String() string {
	return tea.Prettify(s)
}

func (s SetVpcAccessRequest) GoString() string {
	return s.String()
}

func (s *SetVpcAccessRequest) SetDescription(v string) *SetVpcAccessRequest {
	s.Description = &v
	return s
}

func (s *SetVpcAccessRequest) SetInstanceId(v string) *SetVpcAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *SetVpcAccessRequest) SetName(v string) *SetVpcAccessRequest {
	s.Name = &v
	return s
}

func (s *SetVpcAccessRequest) SetPort(v int32) *SetVpcAccessRequest {
	s.Port = &v
	return s
}

func (s *SetVpcAccessRequest) SetSecurityToken(v string) *SetVpcAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetVpcAccessRequest) SetVpcId(v string) *SetVpcAccessRequest {
	s.VpcId = &v
	return s
}

type SetVpcAccessResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetVpcAccessResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetVpcAccessResponseBody) GoString() string {
	return s.String()
}

func (s *SetVpcAccessResponseBody) SetRequestId(v string) *SetVpcAccessResponseBody {
	s.RequestId = &v
	return s
}

type SetVpcAccessResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetVpcAccessResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetVpcAccessResponse) String() string {
	return tea.Prettify(s)
}

func (s SetVpcAccessResponse) GoString() string {
	return s.String()
}

func (s *SetVpcAccessResponse) SetHeaders(v map[string]*string) *SetVpcAccessResponse {
	s.Headers = v
	return s
}

func (s *SetVpcAccessResponse) SetBody(v *SetVpcAccessResponseBody) *SetVpcAccessResponse {
	s.Body = v
	return s
}

type SetWildcardDomainPatternsRequest struct {
	DomainName             *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	GroupId                *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SecurityToken          *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	WildcardDomainPatterns *string `json:"WildcardDomainPatterns,omitempty" xml:"WildcardDomainPatterns,omitempty"`
}

func (s SetWildcardDomainPatternsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetWildcardDomainPatternsRequest) GoString() string {
	return s.String()
}

func (s *SetWildcardDomainPatternsRequest) SetDomainName(v string) *SetWildcardDomainPatternsRequest {
	s.DomainName = &v
	return s
}

func (s *SetWildcardDomainPatternsRequest) SetGroupId(v string) *SetWildcardDomainPatternsRequest {
	s.GroupId = &v
	return s
}

func (s *SetWildcardDomainPatternsRequest) SetSecurityToken(v string) *SetWildcardDomainPatternsRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetWildcardDomainPatternsRequest) SetWildcardDomainPatterns(v string) *SetWildcardDomainPatternsRequest {
	s.WildcardDomainPatterns = &v
	return s
}

type SetWildcardDomainPatternsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetWildcardDomainPatternsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetWildcardDomainPatternsResponseBody) GoString() string {
	return s.String()
}

func (s *SetWildcardDomainPatternsResponseBody) SetRequestId(v string) *SetWildcardDomainPatternsResponseBody {
	s.RequestId = &v
	return s
}

type SetWildcardDomainPatternsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetWildcardDomainPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetWildcardDomainPatternsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetWildcardDomainPatternsResponse) GoString() string {
	return s.String()
}

func (s *SetWildcardDomainPatternsResponse) SetHeaders(v map[string]*string) *SetWildcardDomainPatternsResponse {
	s.Headers = v
	return s
}

func (s *SetWildcardDomainPatternsResponse) SetBody(v *SetWildcardDomainPatternsResponseBody) *SetWildcardDomainPatternsResponse {
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
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchApiResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SwitchApiResponse) SetBody(v *SwitchApiResponseBody) *SwitchApiResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	ResourceId    []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType  *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string                   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	Tag           []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetSecurityToken(v string) *TagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All           *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId    []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType  *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string   `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	TagKey        []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetSecurityToken(v string) *UntagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
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
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AbolishApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) AddIpControlPolicyItemWithOptions(request *AddIpControlPolicyItemRequest, runtime *util.RuntimeOptions) (_result *AddIpControlPolicyItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppId"] = request.AppId
	query["CidrIp"] = request.CidrIp
	query["IpControlId"] = request.IpControlId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddIpControlPolicyItem"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["SecurityToken"] = request.SecurityToken
	query["SpecialKey"] = request.SpecialKey
	query["SpecialType"] = request.SpecialType
	query["TrafficControlId"] = request.TrafficControlId
	query["TrafficValue"] = request.TrafficValue
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AddTrafficSpecialControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) AttachPluginWithOptions(request *AttachPluginRequest, runtime *util.RuntimeOptions) (_result *AttachPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["PluginId"] = request.PluginId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachPlugin"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachPlugin(request *AttachPluginRequest) (_result *AttachPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachPluginResponse{}
	_body, _err := client.AttachPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchAbolishApisWithOptions(request *BatchAbolishApisRequest, runtime *util.RuntimeOptions) (_result *BatchAbolishApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Api"] = request.Api
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchAbolishApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchAbolishApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchAbolishApis(request *BatchAbolishApisRequest) (_result *BatchAbolishApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchAbolishApisResponse{}
	_body, _err := client.BatchAbolishApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeployApisWithOptions(request *BatchDeployApisRequest, runtime *util.RuntimeOptions) (_result *BatchDeployApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Api"] = request.Api
	query["Description"] = request.Description
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeployApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeployApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeployApis(request *BatchDeployApisRequest) (_result *BatchDeployApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchDeployApisResponse{}
	_body, _err := client.BatchDeployApisWithOptions(request, runtime)
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
	query["AllowSignatureMethod"] = request.AllowSignatureMethod
	query["ApiName"] = request.ApiName
	query["AppCodeAuthType"] = request.AppCodeAuthType
	query["AuthType"] = request.AuthType
	query["ConstantParameters"] = request.ConstantParameters
	query["Description"] = request.Description
	query["DisableInternet"] = request.DisableInternet
	query["ErrorCodeSamples"] = request.ErrorCodeSamples
	query["FailResultSample"] = request.FailResultSample
	query["ForceNonceCheck"] = request.ForceNonceCheck
	query["GroupId"] = request.GroupId
	query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	query["RequestConfig"] = request.RequestConfig
	query["RequestParameters"] = request.RequestParameters
	query["ResultBodyModel"] = request.ResultBodyModel
	query["ResultDescriptions"] = request.ResultDescriptions
	query["ResultSample"] = request.ResultSample
	query["ResultType"] = request.ResultType
	query["SecurityToken"] = request.SecurityToken
	query["ServiceConfig"] = request.ServiceConfig
	query["ServiceParameters"] = request.ServiceParameters
	query["ServiceParametersMap"] = request.ServiceParametersMap
	query["SystemParameters"] = request.SystemParameters
	query["Visibility"] = request.Visibility
	query["WebSocketApiType"] = request.WebSocketApiType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["BasePath"] = request.BasePath
	query["Description"] = request.Description
	query["GroupName"] = request.GroupName
	query["InstanceId"] = request.InstanceId
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageId"] = request.StageId
	query["StageRouteModel"] = request.StageRouteModel
	query["SupportRoute"] = request.SupportRoute
	query["VariableName"] = request.VariableName
	query["VariableValue"] = request.VariableValue
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApiStageVariable"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["AppName"] = request.AppName
	query["Description"] = request.Description
	query["SecurityToken"] = request.SecurityToken
	query["Source"] = request.Source
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoPay"] = request.AutoPay
	query["ChargeType"] = request.ChargeType
	query["Duration"] = request.Duration
	query["HttpsPolicy"] = request.HttpsPolicy
	query["InstanceName"] = request.InstanceName
	query["InstanceSpec"] = request.InstanceSpec
	query["PricingCycle"] = request.PricingCycle
	query["Token"] = request.Token
	query["ZoneId"] = request.ZoneId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateIntranetDomainWithOptions(request *CreateIntranetDomainRequest, runtime *util.RuntimeOptions) (_result *CreateIntranetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIntranetDomain"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIntranetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateIntranetDomain(request *CreateIntranetDomainRequest) (_result *CreateIntranetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIntranetDomainResponse{}
	_body, _err := client.CreateIntranetDomainWithOptions(request, runtime)
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
	query["Description"] = request.Description
	query["IpControlName"] = request.IpControlName
	query["IpControlPolicys"] = request.IpControlPolicys
	query["IpControlType"] = request.IpControlType
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIpControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["LogType"] = request.LogType
	query["SecurityToken"] = request.SecurityToken
	query["SlsLogStore"] = request.SlsLogStore
	query["SlsProject"] = request.SlsProject
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLogConfig"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) CreateModelWithOptions(request *CreateModelRequest, runtime *util.RuntimeOptions) (_result *CreateModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["ModelName"] = request.ModelName
	query["Schema"] = request.Schema
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateModel"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateModel(request *CreateModelRequest) (_result *CreateModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateModelResponse{}
	_body, _err := client.CreateModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMonitorGroupWithOptions(request *CreateMonitorGroupRequest, runtime *util.RuntimeOptions) (_result *CreateMonitorGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Auth"] = request.Auth
	query["GroupId"] = request.GroupId
	query["RawMonitorGroupId"] = request.RawMonitorGroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMonitorGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMonitorGroup(request *CreateMonitorGroupRequest) (_result *CreateMonitorGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMonitorGroupResponse{}
	_body, _err := client.CreateMonitorGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePluginWithOptions(request *CreatePluginRequest, runtime *util.RuntimeOptions) (_result *CreatePluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["PluginData"] = request.PluginData
	query["PluginName"] = request.PluginName
	query["PluginType"] = request.PluginType
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePlugin"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePlugin(request *CreatePluginRequest) (_result *CreatePluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePluginResponse{}
	_body, _err := client.CreatePluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSignatureWithOptions(request *CreateSignatureRequest, runtime *util.RuntimeOptions) (_result *CreateSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["SecurityToken"] = request.SecurityToken
	query["SignatureKey"] = request.SignatureKey
	query["SignatureName"] = request.SignatureName
	query["SignatureSecret"] = request.SignatureSecret
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSignature"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSignature(request *CreateSignatureRequest) (_result *CreateSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSignatureResponse{}
	_body, _err := client.CreateSignatureWithOptions(request, runtime)
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
	query["ApiDefault"] = request.ApiDefault
	query["AppDefault"] = request.AppDefault
	query["Description"] = request.Description
	query["SecurityToken"] = request.SecurityToken
	query["TrafficControlName"] = request.TrafficControlName
	query["TrafficControlUnit"] = request.TrafficControlUnit
	query["UserDefault"] = request.UserDefault
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTrafficControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["SecurityToken"] = request.SecurityToken
	query["TrafficControlId"] = request.TrafficControlId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAllTrafficSpecialControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApiGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageId"] = request.StageId
	query["VariableName"] = request.VariableName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApiStageVariable"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["AppId"] = request.AppId
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApp"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["CertificateId"] = request.CertificateId
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomainCertificate"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteInstanceWithOptions(request *DeleteInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstance"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstance(request *DeleteInstanceRequest) (_result *DeleteInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceResponse{}
	_body, _err := client.DeleteInstanceWithOptions(request, runtime)
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
	query["IpControlId"] = request.IpControlId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIpControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["LogType"] = request.LogType
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLogConfig"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DeleteModelWithOptions(request *DeleteModelRequest, runtime *util.RuntimeOptions) (_result *DeleteModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["ModelName"] = request.ModelName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteModel"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteModel(request *DeleteModelRequest) (_result *DeleteModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteModelResponse{}
	_body, _err := client.DeleteModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePluginWithOptions(request *DeletePluginRequest, runtime *util.RuntimeOptions) (_result *DeletePluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PluginId"] = request.PluginId
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePlugin"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePlugin(request *DeletePluginRequest) (_result *DeletePluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePluginResponse{}
	_body, _err := client.DeletePluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSignatureWithOptions(request *DeleteSignatureRequest, runtime *util.RuntimeOptions) (_result *DeleteSignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["SecurityToken"] = request.SecurityToken
	query["SignatureId"] = request.SignatureId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSignature"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSignature(request *DeleteSignatureRequest) (_result *DeleteSignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSignatureResponse{}
	_body, _err := client.DeleteSignatureWithOptions(request, runtime)
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
	query["SecurityToken"] = request.SecurityToken
	query["TrafficControlId"] = request.TrafficControlId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["SecurityToken"] = request.SecurityToken
	query["SpecialKey"] = request.SpecialKey
	query["SpecialType"] = request.SpecialType
	query["TrafficControlId"] = request.TrafficControlId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTrafficSpecialControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiId"] = request.ApiId
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeAbolishApiTaskWithOptions(request *DescribeAbolishApiTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeAbolishApiTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OperationUid"] = request.OperationUid
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAbolishApiTask"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAbolishApiTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAbolishApiTask(request *DescribeAbolishApiTaskRequest) (_result *DescribeAbolishApiTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAbolishApiTaskResponse{}
	_body, _err := client.DescribeAbolishApiTaskWithOptions(request, runtime)
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
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiDoc"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeApiGroupWithOptions(request *DescribeApiGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeApiGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (_result *DescribeApiGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiGroupResponse{}
	_body, _err := client.DescribeApiGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiGroupVpcWhitelistWithOptions(request *DescribeApiGroupVpcWhitelistRequest, runtime *util.RuntimeOptions) (_result *DescribeApiGroupVpcWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiGroupVpcWhitelist"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiGroupVpcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiGroupVpcWhitelist(request *DescribeApiGroupVpcWhitelistRequest) (_result *DescribeApiGroupVpcWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiGroupVpcWhitelistResponse{}
	_body, _err := client.DescribeApiGroupVpcWhitelistWithOptions(request, runtime)
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
	query["EnableTagAuth"] = request.EnableTagAuth
	query["GroupId"] = request.GroupId
	query["GroupName"] = request.GroupName
	query["InstanceId"] = request.InstanceId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["Sort"] = request.Sort
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiGroups"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeApiHistoriesWithOptions(request *DescribeApiHistoriesRequest, runtime *util.RuntimeOptions) (_result *DescribeApiHistoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["ApiName"] = request.ApiName
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiHistories"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiHistoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiHistories(request *DescribeApiHistoriesRequest) (_result *DescribeApiHistoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiHistoriesResponse{}
	_body, _err := client.DescribeApiHistoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiHistoryWithOptions(request *DescribeApiHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeApiHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["HistoryVersion"] = request.HistoryVersion
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiHistory"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiHistory(request *DescribeApiHistoryRequest) (_result *DescribeApiHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiHistoryResponse{}
	_body, _err := client.DescribeApiHistoryWithOptions(request, runtime)
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
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiIpControls"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeApiLatencyDataWithOptions(request *DescribeApiLatencyDataRequest, runtime *util.RuntimeOptions) (_result *DescribeApiLatencyDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["EndTime"] = request.EndTime
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiLatencyData"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiLatencyDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiLatencyData(request *DescribeApiLatencyDataRequest) (_result *DescribeApiLatencyDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiLatencyDataResponse{}
	_body, _err := client.DescribeApiLatencyDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiMarketAttributesWithOptions(request *DescribeApiMarketAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeApiMarketAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiMarketAttributes"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiMarketAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiMarketAttributes(request *DescribeApiMarketAttributesRequest) (_result *DescribeApiMarketAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiMarketAttributesResponse{}
	_body, _err := client.DescribeApiMarketAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiQpsDataWithOptions(request *DescribeApiQpsDataRequest, runtime *util.RuntimeOptions) (_result *DescribeApiQpsDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["EndTime"] = request.EndTime
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiQpsData"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiQpsDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiQpsData(request *DescribeApiQpsDataRequest) (_result *DescribeApiQpsDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiQpsDataResponse{}
	_body, _err := client.DescribeApiQpsDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiSignaturesWithOptions(request *DescribeApiSignaturesRequest, runtime *util.RuntimeOptions) (_result *DescribeApiSignaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiSignatures"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiSignaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiSignatures(request *DescribeApiSignaturesRequest) (_result *DescribeApiSignaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiSignaturesResponse{}
	_body, _err := client.DescribeApiSignaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiTrafficControlsWithOptions(request *DescribeApiTrafficControlsRequest, runtime *util.RuntimeOptions) (_result *DescribeApiTrafficControlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiTrafficControls"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiTrafficControlsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiTrafficControls(request *DescribeApiTrafficControlsRequest) (_result *DescribeApiTrafficControlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiTrafficControlsResponse{}
	_body, _err := client.DescribeApiTrafficControlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiTrafficDataWithOptions(request *DescribeApiTrafficDataRequest, runtime *util.RuntimeOptions) (_result *DescribeApiTrafficDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["EndTime"] = request.EndTime
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["StartTime"] = request.StartTime
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApiTrafficData"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApiTrafficDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiTrafficData(request *DescribeApiTrafficDataRequest) (_result *DescribeApiTrafficDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiTrafficDataResponse{}
	_body, _err := client.DescribeApiTrafficDataWithOptions(request, runtime)
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
	query["ApiId"] = request.ApiId
	query["ApiName"] = request.ApiName
	query["CatalogId"] = request.CatalogId
	query["EnableTagAuth"] = request.EnableTagAuth
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	query["Visibility"] = request.Visibility
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiName"] = request.ApiName
	query["ApiUid"] = request.ApiUid
	query["AppId"] = request.AppId
	query["Method"] = request.Method
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["Path"] = request.Path
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisByApp"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["IpControlId"] = request.IpControlId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisByIpControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeApisBySignatureWithOptions(request *DescribeApisBySignatureRequest, runtime *util.RuntimeOptions) (_result *DescribeApisBySignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["SignatureId"] = request.SignatureId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisBySignature"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisBySignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApisBySignature(request *DescribeApisBySignatureRequest) (_result *DescribeApisBySignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisBySignatureResponse{}
	_body, _err := client.DescribeApisBySignatureWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApisByTrafficControlWithOptions(request *DescribeApisByTrafficControlRequest, runtime *util.RuntimeOptions) (_result *DescribeApisByTrafficControlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["TrafficControlId"] = request.TrafficControlId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApisByTrafficControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApisByTrafficControlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApisByTrafficControl(request *DescribeApisByTrafficControlRequest) (_result *DescribeApisByTrafficControlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApisByTrafficControlResponse{}
	_body, _err := client.DescribeApisByTrafficControlWithOptions(request, runtime)
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
	query["AppId"] = request.AppId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApp"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeAppAttributesWithOptions(request *DescribeAppAttributesRequest, runtime *util.RuntimeOptions) (_result *DescribeAppAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppCode"] = request.AppCode
	query["AppId"] = request.AppId
	query["AppKey"] = request.AppKey
	query["AppName"] = request.AppName
	query["EnableTagAuth"] = request.EnableTagAuth
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["Sort"] = request.Sort
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppAttributes"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppAttributesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppAttributes(request *DescribeAppAttributesRequest) (_result *DescribeAppAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppAttributesResponse{}
	_body, _err := client.DescribeAppAttributesWithOptions(request, runtime)
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
	query["AppId"] = request.AppId
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppSecurity"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["AppId"] = request.AppId
	query["AppOwner"] = request.AppOwner
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApps"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeAuthorizedApisWithOptions(request *DescribeAuthorizedApisRequest, runtime *util.RuntimeOptions) (_result *DescribeAuthorizedApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppId"] = request.AppId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuthorizedApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuthorizedApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuthorizedApis(request *DescribeAuthorizedApisRequest) (_result *DescribeAuthorizedApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuthorizedApisResponse{}
	_body, _err := client.DescribeAuthorizedApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAuthorizedAppsWithOptions(request *DescribeAuthorizedAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAuthorizedAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["AppId"] = request.AppId
	query["AppName"] = request.AppName
	query["AppOwnerId"] = request.AppOwnerId
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAuthorizedApps"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAuthorizedAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAuthorizedApps(request *DescribeAuthorizedAppsRequest) (_result *DescribeAuthorizedAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAuthorizedAppsResponse{}
	_body, _err := client.DescribeAuthorizedAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDeployApiTaskWithOptions(request *DescribeDeployApiTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeDeployApiTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OperationUid"] = request.OperationUid
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeployApiTask"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeployApiTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDeployApiTask(request *DescribeDeployApiTaskRequest) (_result *DescribeDeployApiTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeployApiTaskResponse{}
	_body, _err := client.DescribeDeployApiTaskWithOptions(request, runtime)
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
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeployedApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiId"] = request.ApiId
	query["ApiName"] = request.ApiName
	query["EnableTagAuth"] = request.EnableTagAuth
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeployedApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDomain"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeHistoryApisWithOptions(request *DescribeHistoryApisRequest, runtime *util.RuntimeOptions) (_result *DescribeHistoryApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["ApiName"] = request.ApiName
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeHistoryApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["IpControlId"] = request.IpControlId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["PolicyItemId"] = request.PolicyItemId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpControlPolicyItems"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["IpControlId"] = request.IpControlId
	query["IpControlName"] = request.IpControlName
	query["IpControlType"] = request.IpControlType
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeIpControls"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["LogType"] = request.LogType
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeLogConfig"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeMarketRemainsQuotaWithOptions(request *DescribeMarketRemainsQuotaRequest, runtime *util.RuntimeOptions) (_result *DescribeMarketRemainsQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DomainName"] = request.DomainName
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMarketRemainsQuota"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMarketRemainsQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMarketRemainsQuota(request *DescribeMarketRemainsQuotaRequest) (_result *DescribeMarketRemainsQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMarketRemainsQuotaResponse{}
	_body, _err := client.DescribeMarketRemainsQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModelsWithOptions(request *DescribeModelsRequest, runtime *util.RuntimeOptions) (_result *DescribeModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["ModelId"] = request.ModelId
	query["ModelName"] = request.ModelName
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeModels"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModels(request *DescribeModelsRequest) (_result *DescribeModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModelsResponse{}
	_body, _err := client.DescribeModelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePluginSchemasWithOptions(request *DescribePluginSchemasRequest, runtime *util.RuntimeOptions) (_result *DescribePluginSchemasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Language"] = request.Language
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePluginSchemas"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePluginSchemasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePluginSchemas(request *DescribePluginSchemasRequest) (_result *DescribePluginSchemasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePluginSchemasResponse{}
	_body, _err := client.DescribePluginSchemasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePluginTemplatesWithOptions(request *DescribePluginTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribePluginTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Language"] = request.Language
	query["PluginName"] = request.PluginName
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePluginTemplates"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePluginTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePluginTemplates(request *DescribePluginTemplatesRequest) (_result *DescribePluginTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePluginTemplatesResponse{}
	_body, _err := client.DescribePluginTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePluginsWithOptions(request *DescribePluginsRequest, runtime *util.RuntimeOptions) (_result *DescribePluginsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["PluginId"] = request.PluginId
	query["PluginName"] = request.PluginName
	query["PluginType"] = request.PluginType
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePlugins"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePluginsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePlugins(request *DescribePluginsRequest) (_result *DescribePluginsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePluginsResponse{}
	_body, _err := client.DescribePluginsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePluginsByApiWithOptions(request *DescribePluginsByApiRequest, runtime *util.RuntimeOptions) (_result *DescribePluginsByApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePluginsByApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePluginsByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePluginsByApi(request *DescribePluginsByApiRequest) (_result *DescribePluginsByApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePluginsByApiResponse{}
	_body, _err := client.DescribePluginsByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePurchasedApiGroupWithOptions(request *DescribePurchasedApiGroupRequest, runtime *util.RuntimeOptions) (_result *DescribePurchasedApiGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApiGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePurchasedApiGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePurchasedApiGroup(request *DescribePurchasedApiGroupRequest) (_result *DescribePurchasedApiGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePurchasedApiGroupResponse{}
	_body, _err := client.DescribePurchasedApiGroupWithOptions(request, runtime)
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
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApiGroups"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiId"] = request.ApiId
	query["ApiName"] = request.ApiName
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["Visibility"] = request.Visibility
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePurchasedApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Language"] = request.Language
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeSignaturesWithOptions(request *DescribeSignaturesRequest, runtime *util.RuntimeOptions) (_result *DescribeSignaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["SignatureId"] = request.SignatureId
	query["SignatureName"] = request.SignatureName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSignatures"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSignaturesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSignatures(request *DescribeSignaturesRequest) (_result *DescribeSignaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSignaturesResponse{}
	_body, _err := client.DescribeSignaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSignaturesByApiWithOptions(request *DescribeSignaturesByApiRequest, runtime *util.RuntimeOptions) (_result *DescribeSignaturesByApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSignaturesByApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSignaturesByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSignaturesByApi(request *DescribeSignaturesByApiRequest) (_result *DescribeSignaturesByApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSignaturesByApiResponse{}
	_body, _err := client.DescribeSignaturesByApiWithOptions(request, runtime)
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
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSystemParameters"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeTrafficControlsWithOptions(request *DescribeTrafficControlsRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficControlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["TrafficControlId"] = request.TrafficControlId
	query["TrafficControlName"] = request.TrafficControlName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrafficControls"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) DescribeTrafficControlsByApiWithOptions(request *DescribeTrafficControlsByApiRequest, runtime *util.RuntimeOptions) (_result *DescribeTrafficControlsByApiResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTrafficControlsByApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTrafficControlsByApiResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTrafficControlsByApi(request *DescribeTrafficControlsByApiRequest) (_result *DescribeTrafficControlsByApiResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTrafficControlsByApiResponse{}
	_body, _err := client.DescribeTrafficControlsByApiWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUpdateVpcInfoTaskWithOptions(request *DescribeUpdateVpcInfoTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeUpdateVpcInfoTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["OperationUid"] = request.OperationUid
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUpdateVpcInfoTask"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUpdateVpcInfoTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUpdateVpcInfoTask(request *DescribeUpdateVpcInfoTaskRequest) (_result *DescribeUpdateVpcInfoTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUpdateVpcInfoTaskResponse{}
	_body, _err := client.DescribeUpdateVpcInfoTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeVpcAccessesWithOptions(request *DescribeVpcAccessesRequest, runtime *util.RuntimeOptions) (_result *DescribeVpcAccessesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Name"] = request.Name
	query["PageNumber"] = request.PageNumber
	query["PageSize"] = request.PageSize
	query["SecurityToken"] = request.SecurityToken
	query["VpcAccessId"] = request.VpcAccessId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVpcAccesses"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVpcAccessesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeVpcAccesses(request *DescribeVpcAccessesRequest) (_result *DescribeVpcAccessesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVpcAccessesResponse{}
	_body, _err := client.DescribeVpcAccessesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Language"] = request.Language
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DryRunSwaggerWithOptions(tmpReq *DryRunSwaggerRequest, runtime *util.RuntimeOptions) (_result *DryRunSwaggerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DryRunSwaggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GlobalCondition)) {
		request.GlobalConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalCondition, tea.String("GlobalCondition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["DataFormat"] = request.DataFormat
	query["GlobalCondition"] = request.GlobalConditionShrink
	query["GroupId"] = request.GroupId
	query["Overwrite"] = request.Overwrite
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("DryRunSwagger"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DryRunSwaggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DryRunSwagger(request *DryRunSwaggerRequest) (_result *DryRunSwaggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DryRunSwaggerResponse{}
	_body, _err := client.DryRunSwaggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportSwaggerWithOptions(tmpReq *ImportSwaggerRequest, runtime *util.RuntimeOptions) (_result *ImportSwaggerResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ImportSwaggerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.GlobalCondition)) {
		request.GlobalConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GlobalCondition, tea.String("GlobalCondition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	query["DataFormat"] = request.DataFormat
	query["DryRun"] = request.DryRun
	query["GlobalCondition"] = request.GlobalConditionShrink
	query["GroupId"] = request.GroupId
	query["Overwrite"] = request.Overwrite
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportSwagger"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportSwaggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportSwagger(request *ImportSwaggerRequest) (_result *ImportSwaggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportSwaggerResponse{}
	_body, _err := client.ImportSwaggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["NextToken"] = request.NextToken
	query["ResourceId"] = request.ResourceId
	query["ResourceType"] = request.ResourceType
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
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
	query["AllowSignatureMethod"] = request.AllowSignatureMethod
	query["ApiId"] = request.ApiId
	query["ApiName"] = request.ApiName
	query["AppCodeAuthType"] = request.AppCodeAuthType
	query["AuthType"] = request.AuthType
	query["ConstantParameters"] = request.ConstantParameters
	query["Description"] = request.Description
	query["DisableInternet"] = request.DisableInternet
	query["ErrorCodeSamples"] = request.ErrorCodeSamples
	query["FailResultSample"] = request.FailResultSample
	query["ForceNonceCheck"] = request.ForceNonceCheck
	query["GroupId"] = request.GroupId
	query["OpenIdConnectConfig"] = request.OpenIdConnectConfig
	query["RequestConfig"] = request.RequestConfig
	query["RequestParameters"] = request.RequestParameters
	query["ResultBodyModel"] = request.ResultBodyModel
	query["ResultDescriptions"] = request.ResultDescriptions
	query["ResultSample"] = request.ResultSample
	query["ResultType"] = request.ResultType
	query["SecurityToken"] = request.SecurityToken
	query["ServiceConfig"] = request.ServiceConfig
	query["ServiceParameters"] = request.ServiceParameters
	query["ServiceParametersMap"] = request.ServiceParametersMap
	query["SystemParameters"] = request.SystemParameters
	query["Visibility"] = request.Visibility
	query["WebSocketApiType"] = request.WebSocketApiType
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["BasePath"] = request.BasePath
	query["CompatibleFlags"] = request.CompatibleFlags
	query["CustomTraceConfig"] = request.CustomTraceConfig
	query["CustomerConfigs"] = request.CustomerConfigs
	query["DefaultDomain"] = request.DefaultDomain
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["GroupName"] = request.GroupName
	query["PassthroughHeaders"] = request.PassthroughHeaders
	query["RpcPattern"] = request.RpcPattern
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	query["UserLogConfig"] = request.UserLogConfig
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyApiGroupVpcWhitelistWithOptions(request *ModifyApiGroupVpcWhitelistRequest, runtime *util.RuntimeOptions) (_result *ModifyApiGroupVpcWhitelistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["VpcIds"] = request.VpcIds
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApiGroupVpcWhitelist"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApiGroupVpcWhitelistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApiGroupVpcWhitelist(request *ModifyApiGroupVpcWhitelistRequest) (_result *ModifyApiGroupVpcWhitelistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApiGroupVpcWhitelistResponse{}
	_body, _err := client.ModifyApiGroupVpcWhitelistWithOptions(request, runtime)
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
	query["AppId"] = request.AppId
	query["AppName"] = request.AppName
	query["Description"] = request.Description
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApp"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyInstanceSpecWithOptions(request *ModifyInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AutoPay"] = request.AutoPay
	query["InstanceId"] = request.InstanceId
	query["InstanceSpec"] = request.InstanceSpec
	query["Token"] = request.Token
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceSpec"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceSpec(request *ModifyInstanceSpecRequest) (_result *ModifyInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceSpecResponse{}
	_body, _err := client.ModifyInstanceSpecWithOptions(request, runtime)
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
	query["Description"] = request.Description
	query["IpControlId"] = request.IpControlId
	query["IpControlName"] = request.IpControlName
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIpControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["AppId"] = request.AppId
	query["CidrIp"] = request.CidrIp
	query["IpControlId"] = request.IpControlId
	query["PolicyItemId"] = request.PolicyItemId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyIpControlPolicyItem"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["LogType"] = request.LogType
	query["SecurityToken"] = request.SecurityToken
	query["SlsLogStore"] = request.SlsLogStore
	query["SlsProject"] = request.SlsProject
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLogConfig"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ModifyModelWithOptions(request *ModifyModelRequest, runtime *util.RuntimeOptions) (_result *ModifyModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["ModelName"] = request.ModelName
	query["NewModelName"] = request.NewModelName
	query["Schema"] = request.Schema
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyModel"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyModel(request *ModifyModelRequest) (_result *ModifyModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyModelResponse{}
	_body, _err := client.ModifyModelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyPluginWithOptions(request *ModifyPluginRequest, runtime *util.RuntimeOptions) (_result *ModifyPluginResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["PluginData"] = request.PluginData
	query["PluginId"] = request.PluginId
	query["PluginName"] = request.PluginName
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyPlugin"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyPluginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyPlugin(request *ModifyPluginRequest) (_result *ModifyPluginResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyPluginResponse{}
	_body, _err := client.ModifyPluginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySignatureWithOptions(request *ModifySignatureRequest, runtime *util.RuntimeOptions) (_result *ModifySignatureResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["SecurityToken"] = request.SecurityToken
	query["SignatureId"] = request.SignatureId
	query["SignatureKey"] = request.SignatureKey
	query["SignatureName"] = request.SignatureName
	query["SignatureSecret"] = request.SignatureSecret
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySignature"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySignatureResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySignature(request *ModifySignatureRequest) (_result *ModifySignatureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySignatureResponse{}
	_body, _err := client.ModifySignatureWithOptions(request, runtime)
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
	query["ApiDefault"] = request.ApiDefault
	query["AppDefault"] = request.AppDefault
	query["Description"] = request.Description
	query["SecurityToken"] = request.SecurityToken
	query["TrafficControlId"] = request.TrafficControlId
	query["TrafficControlName"] = request.TrafficControlName
	query["TrafficControlUnit"] = request.TrafficControlUnit
	query["UserDefault"] = request.UserDefault
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTrafficControl"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) OpenApiGatewayServiceWithOptions(runtime *util.RuntimeOptions) (_result *OpenApiGatewayServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("OpenApiGatewayService"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenApiGatewayServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OpenApiGatewayService() (_result *OpenApiGatewayServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenApiGatewayServiceResponse{}
	_body, _err := client.OpenApiGatewayServiceWithOptions(runtime)
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
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ReactivateDomain"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RemoveApisAuthoritiesWithOptions(request *RemoveApisAuthoritiesRequest, runtime *util.RuntimeOptions) (_result *RemoveApisAuthoritiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["AppId"] = request.AppId
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveApisAuthorities"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveApisAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveApisAuthorities(request *RemoveApisAuthoritiesRequest) (_result *RemoveApisAuthoritiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveApisAuthoritiesResponse{}
	_body, _err := client.RemoveApisAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAppsAuthoritiesWithOptions(request *RemoveAppsAuthoritiesRequest, runtime *util.RuntimeOptions) (_result *RemoveAppsAuthoritiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["AppIds"] = request.AppIds
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAppsAuthorities"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAppsAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAppsAuthorities(request *RemoveAppsAuthoritiesRequest) (_result *RemoveAppsAuthoritiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAppsAuthoritiesResponse{}
	_body, _err := client.RemoveAppsAuthoritiesWithOptions(request, runtime)
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
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["IpControlId"] = request.IpControlId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveIpControlApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["IpControlId"] = request.IpControlId
	query["PolicyItemIds"] = request.PolicyItemIds
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveIpControlPolicyItem"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) RemoveSignatureApisWithOptions(request *RemoveSignatureApisRequest, runtime *util.RuntimeOptions) (_result *RemoveSignatureApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["SignatureId"] = request.SignatureId
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSignatureApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSignatureApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSignatureApis(request *RemoveSignatureApisRequest) (_result *RemoveSignatureApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveSignatureApisResponse{}
	_body, _err := client.RemoveSignatureApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTrafficControlApisWithOptions(request *RemoveTrafficControlApisRequest, runtime *util.RuntimeOptions) (_result *RemoveTrafficControlApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["TrafficControlId"] = request.TrafficControlId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTrafficControlApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTrafficControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTrafficControlApis(request *RemoveTrafficControlApisRequest) (_result *RemoveTrafficControlApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTrafficControlApisResponse{}
	_body, _err := client.RemoveTrafficControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveVpcAccessWithOptions(request *RemoveVpcAccessRequest, runtime *util.RuntimeOptions) (_result *RemoveVpcAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["NeedBatchWork"] = request.NeedBatchWork
	query["Port"] = request.Port
	query["SecurityToken"] = request.SecurityToken
	query["VpcId"] = request.VpcId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveVpcAccess"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveVpcAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveVpcAccess(request *RemoveVpcAccessRequest) (_result *RemoveVpcAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveVpcAccessResponse{}
	_body, _err := client.RemoveVpcAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveVpcAccessAndAbolishApisWithOptions(request *RemoveVpcAccessAndAbolishApisRequest, runtime *util.RuntimeOptions) (_result *RemoveVpcAccessAndAbolishApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["InstanceId"] = request.InstanceId
	query["NeedBatchWork"] = request.NeedBatchWork
	query["Port"] = request.Port
	query["SecurityToken"] = request.SecurityToken
	query["VpcId"] = request.VpcId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveVpcAccessAndAbolishApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveVpcAccessAndAbolishApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveVpcAccessAndAbolishApis(request *RemoveVpcAccessAndAbolishApisRequest) (_result *RemoveVpcAccessAndAbolishApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveVpcAccessAndAbolishApisResponse{}
	_body, _err := client.RemoveVpcAccessAndAbolishApisWithOptions(request, runtime)
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
	query["AppCode"] = request.AppCode
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAppCode"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) ResetAppSecretWithOptions(request *ResetAppSecretRequest, runtime *util.RuntimeOptions) (_result *ResetAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["AppKey"] = request.AppKey
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetAppSecret"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetAppSecretResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAppSecret(request *ResetAppSecretRequest) (_result *ResetAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAppSecretResponse{}
	_body, _err := client.ResetAppSecretWithOptions(request, runtime)
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
	query["AppId"] = request.AppId
	query["Language"] = request.Language
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SdkGenerateByApp"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["GroupId"] = request.GroupId
	query["Language"] = request.Language
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SdkGenerateByGroup"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) SetApisAuthoritiesWithOptions(request *SetApisAuthoritiesRequest, runtime *util.RuntimeOptions) (_result *SetApisAuthoritiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["AppId"] = request.AppId
	query["AuthValidTime"] = request.AuthValidTime
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetApisAuthorities"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetApisAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetApisAuthorities(request *SetApisAuthoritiesRequest) (_result *SetApisAuthoritiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetApisAuthoritiesResponse{}
	_body, _err := client.SetApisAuthoritiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAppsAuthoritiesWithOptions(request *SetAppsAuthoritiesRequest, runtime *util.RuntimeOptions) (_result *SetAppsAuthoritiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiId"] = request.ApiId
	query["AppIds"] = request.AppIds
	query["AuthValidTime"] = request.AuthValidTime
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetAppsAuthorities"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetAppsAuthoritiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAppsAuthorities(request *SetAppsAuthoritiesRequest) (_result *SetAppsAuthoritiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAppsAuthoritiesResponse{}
	_body, _err := client.SetAppsAuthoritiesWithOptions(request, runtime)
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
	query["BindStageName"] = request.BindStageName
	query["CustomDomainType"] = request.CustomDomainType
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["IsForce"] = request.IsForce
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomain"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["CaCertificateBody"] = request.CaCertificateBody
	query["CertificateBody"] = request.CertificateBody
	query["CertificateName"] = request.CertificateName
	query["CertificatePrivateKey"] = request.CertificatePrivateKey
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomainCertificate"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ActionValue"] = request.ActionValue
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDomainWebSocketStatus"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["IpControlId"] = request.IpControlId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetIpControlApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) SetSignatureApisWithOptions(request *SetSignatureApisRequest, runtime *util.RuntimeOptions) (_result *SetSignatureApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["SignatureId"] = request.SignatureId
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetSignatureApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetSignatureApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSignatureApis(request *SetSignatureApisRequest) (_result *SetSignatureApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSignatureApisResponse{}
	_body, _err := client.SetSignatureApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetTrafficControlApisWithOptions(request *SetTrafficControlApisRequest, runtime *util.RuntimeOptions) (_result *SetTrafficControlApisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ApiIds"] = request.ApiIds
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	query["TrafficControlId"] = request.TrafficControlId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTrafficControlApis"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTrafficControlApisResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetTrafficControlApis(request *SetTrafficControlApisRequest) (_result *SetTrafficControlApisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTrafficControlApisResponse{}
	_body, _err := client.SetTrafficControlApisWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetVpcAccessWithOptions(request *SetVpcAccessRequest, runtime *util.RuntimeOptions) (_result *SetVpcAccessResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["Description"] = request.Description
	query["InstanceId"] = request.InstanceId
	query["Name"] = request.Name
	query["Port"] = request.Port
	query["SecurityToken"] = request.SecurityToken
	query["VpcId"] = request.VpcId
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetVpcAccess"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetVpcAccessResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetVpcAccess(request *SetVpcAccessRequest) (_result *SetVpcAccessResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetVpcAccessResponse{}
	_body, _err := client.SetVpcAccessWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetWildcardDomainPatternsWithOptions(request *SetWildcardDomainPatternsRequest, runtime *util.RuntimeOptions) (_result *SetWildcardDomainPatternsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["DomainName"] = request.DomainName
	query["GroupId"] = request.GroupId
	query["SecurityToken"] = request.SecurityToken
	query["WildcardDomainPatterns"] = request.WildcardDomainPatterns
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SetWildcardDomainPatterns"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetWildcardDomainPatternsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetWildcardDomainPatterns(request *SetWildcardDomainPatternsRequest) (_result *SetWildcardDomainPatternsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetWildcardDomainPatternsResponse{}
	_body, _err := client.SetWildcardDomainPatternsWithOptions(request, runtime)
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
	query["ApiId"] = request.ApiId
	query["Description"] = request.Description
	query["GroupId"] = request.GroupId
	query["HistoryVersion"] = request.HistoryVersion
	query["SecurityToken"] = request.SecurityToken
	query["StageName"] = request.StageName
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchApi"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
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

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["ResourceId"] = request.ResourceId
	query["ResourceType"] = request.ResourceType
	query["SecurityToken"] = request.SecurityToken
	query["Tag"] = request.Tag
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	query["All"] = request.All
	query["ResourceId"] = request.ResourceId
	query["ResourceType"] = request.ResourceType
	query["SecurityToken"] = request.SecurityToken
	query["TagKey"] = request.TagKey
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  util.ToMap(request),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2016-07-14"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
