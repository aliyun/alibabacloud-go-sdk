// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSignatureMethod(v string) *CreateApiRequest
	GetAllowSignatureMethod() *string
	SetApiName(v string) *CreateApiRequest
	GetApiName() *string
	SetAppCodeAuthType(v string) *CreateApiRequest
	GetAppCodeAuthType() *string
	SetAuthType(v string) *CreateApiRequest
	GetAuthType() *string
	SetBackendEnable(v bool) *CreateApiRequest
	GetBackendEnable() *bool
	SetBackendId(v string) *CreateApiRequest
	GetBackendId() *string
	SetConstantParameters(v string) *CreateApiRequest
	GetConstantParameters() *string
	SetDescription(v string) *CreateApiRequest
	GetDescription() *string
	SetDisableInternet(v bool) *CreateApiRequest
	GetDisableInternet() *bool
	SetErrorCodeSamples(v string) *CreateApiRequest
	GetErrorCodeSamples() *string
	SetFailResultSample(v string) *CreateApiRequest
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *CreateApiRequest
	GetForceNonceCheck() *bool
	SetGroupId(v string) *CreateApiRequest
	GetGroupId() *string
	SetOpenIdConnectConfig(v string) *CreateApiRequest
	GetOpenIdConnectConfig() *string
	SetRequestConfig(v string) *CreateApiRequest
	GetRequestConfig() *string
	SetRequestParameters(v string) *CreateApiRequest
	GetRequestParameters() *string
	SetResultBodyModel(v string) *CreateApiRequest
	GetResultBodyModel() *string
	SetResultDescriptions(v string) *CreateApiRequest
	GetResultDescriptions() *string
	SetResultSample(v string) *CreateApiRequest
	GetResultSample() *string
	SetResultType(v string) *CreateApiRequest
	GetResultType() *string
	SetSecurityToken(v string) *CreateApiRequest
	GetSecurityToken() *string
	SetServiceConfig(v string) *CreateApiRequest
	GetServiceConfig() *string
	SetServiceParameters(v string) *CreateApiRequest
	GetServiceParameters() *string
	SetServiceParametersMap(v string) *CreateApiRequest
	GetServiceParametersMap() *string
	SetSystemParameters(v string) *CreateApiRequest
	GetSystemParameters() *string
	SetTag(v []*CreateApiRequestTag) *CreateApiRequest
	GetTag() []*CreateApiRequestTag
	SetVisibility(v string) *CreateApiRequest
	GetVisibility() *string
	SetWebSocketApiType(v string) *CreateApiRequest
	GetWebSocketApiType() *string
}

type CreateApiRequest struct {
	// The type of the two-way communication API.
	//
	// 	- **COMMON**: normal APIs
	//
	// 	- **REGISTER**: registered APIs
	//
	// 	- **UNREGISTER**: unregistered APIs
	//
	// 	- **NOTIFY**: downstream notification APIs
	//
	// example:
	//
	// HmacSHA256
	AllowSignatureMethod *string `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	// The name of the API that you want to create. The name must be unique within the API group. The name must be 4 to 50 characters in length. It must start with a letter and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// ApiName
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The IDof the backend service
	//
	// example:
	//
	// HEADER
	AppCodeAuthType *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	// The configuration items of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestConfig](https://help.aliyun.com/document_detail/43985.html).
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Specifies whether to enable backend services.
	//
	// example:
	//
	// true
	BackendEnable *bool `json:"BackendEnable,omitempty" xml:"BackendEnable,omitempty"`
	// Specifies whether to enable backend services.
	//
	// example:
	//
	// a0305308908c4740aba9cbfd63ba99b7
	BackendId          *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	ConstantParameters *string `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty"`
	// The description of the API. The description can be up to 180 characters in length.
	//
	// example:
	//
	// Api description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// If **AuthType*	- is set to **APP**, the valid values are:
	//
	// 	- **DEFAULT**: The default value that is used if no other values are passed. This value means that the setting of the group is used.
	//
	// 	- **DISABLE**: The authentication is disabled.
	//
	// 	- **HEADER**: AppCode can be placed in the Header parameter for authentication.
	//
	// 	- **HEADER_QUERY**: AppCode can be placed in the Header or Query parameter for authentication.
	//
	// example:
	//
	// true
	DisableInternet  *bool   `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	FailResultSample *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	// 	- Specifies whether to set **DisableInternet*	- to **true*	- to limit API calls to within the VPC.
	//
	// 	- If you set **DisableInternet*	- to **false**, the limit is lifted. The default value is false when you create an API.
	//
	// example:
	//
	// true
	ForceNonceCheck *bool `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 08ae4aa0f95e4321849ee57f4e0b3077
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// If the **AuthType*	- is **APP*	- authentication, you need to pass this value to specify the signature algorithm. If you do not specify this parameter, the default value HmacSHA256 is used. Valid values:
	//
	// 	- HmacSHA256
	//
	// 	- HmacSHA1,HmacSHA256
	//
	// example:
	//
	// {\\"openIdApiType\\":null,\\"idTokenParamName\\":null,\\"publicKeyId\\":null,\\"publicKey\\":null}
	OpenIdConnectConfig *string `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty"`
	// The configuration items of API requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceConfig](https://help.aliyun.com/document_detail/43987.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"RequestProtocol":"HTTP","RequestHttpMethod":"GET","RequestPath":"/v3/getUserTest/[userId]","BodyFormat":"FORM","PostBodyDescription":""}
	RequestConfig     *string `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty"`
	RequestParameters *string `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty"`
	// 	- Specifies whether to set **ForceNonceCheck*	- to **true*	- to force the check of X-Ca-Nonce during the request. This is the unique identifier of the request and is generally identified by UUID. After receiving this parameter, API Gateway verifies the validity of this parameter. The same value can be used only once within 15 minutes. This helps prevent replay attacks.
	//
	// 	- If you set **ForceNonceCheck*	- to **false**, the check is not performed. The default value is false when you create an API.
	//
	// example:
	//
	// {}
	ResultBodyModel    *string `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	ResultDescriptions *string `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty"`
	ResultSample       *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// The sample response from the backend service.
	//
	// example:
	//
	// HTML
	ResultType    *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The parameters of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestParameter](https://help.aliyun.com/document_detail/43986.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {"ServiceProtocol":"HTTP","ServiceHttpMethod":"GET","ServiceAddress":"http://www.customerdomain.com","ServiceTimeout":"1000","ServicePath":"/v3/getUserTest/[userId]"}
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	SystemParameters     *string `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty"`
	// The list of tags.
	Tag []*CreateApiRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to make the API public. Valid values:
	//
	// 	- **PUBLIC**: Make the API public. If you set this parameter to PUBLIC, this API is displayed on the APIs page for all users after the API is published to the production environment.
	//
	// 	- **PRIVATE**: Make the API private. Private APIs are not displayed in the Alibaba Cloud Marketplace after the API group to which they belong is made available.
	//
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The return description of the API.
	//
	// example:
	//
	// COMMON
	WebSocketApiType *string `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s CreateApiRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiRequest) GoString() string {
	return s.String()
}

func (s *CreateApiRequest) GetAllowSignatureMethod() *string {
	return s.AllowSignatureMethod
}

func (s *CreateApiRequest) GetApiName() *string {
	return s.ApiName
}

func (s *CreateApiRequest) GetAppCodeAuthType() *string {
	return s.AppCodeAuthType
}

func (s *CreateApiRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateApiRequest) GetBackendEnable() *bool {
	return s.BackendEnable
}

func (s *CreateApiRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *CreateApiRequest) GetConstantParameters() *string {
	return s.ConstantParameters
}

func (s *CreateApiRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiRequest) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *CreateApiRequest) GetErrorCodeSamples() *string {
	return s.ErrorCodeSamples
}

func (s *CreateApiRequest) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *CreateApiRequest) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *CreateApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateApiRequest) GetOpenIdConnectConfig() *string {
	return s.OpenIdConnectConfig
}

func (s *CreateApiRequest) GetRequestConfig() *string {
	return s.RequestConfig
}

func (s *CreateApiRequest) GetRequestParameters() *string {
	return s.RequestParameters
}

func (s *CreateApiRequest) GetResultBodyModel() *string {
	return s.ResultBodyModel
}

func (s *CreateApiRequest) GetResultDescriptions() *string {
	return s.ResultDescriptions
}

func (s *CreateApiRequest) GetResultSample() *string {
	return s.ResultSample
}

func (s *CreateApiRequest) GetResultType() *string {
	return s.ResultType
}

func (s *CreateApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateApiRequest) GetServiceConfig() *string {
	return s.ServiceConfig
}

func (s *CreateApiRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *CreateApiRequest) GetServiceParametersMap() *string {
	return s.ServiceParametersMap
}

func (s *CreateApiRequest) GetSystemParameters() *string {
	return s.SystemParameters
}

func (s *CreateApiRequest) GetTag() []*CreateApiRequestTag {
	return s.Tag
}

func (s *CreateApiRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateApiRequest) GetWebSocketApiType() *string {
	return s.WebSocketApiType
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

func (s *CreateApiRequest) SetBackendEnable(v bool) *CreateApiRequest {
	s.BackendEnable = &v
	return s
}

func (s *CreateApiRequest) SetBackendId(v string) *CreateApiRequest {
	s.BackendId = &v
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

func (s *CreateApiRequest) SetTag(v []*CreateApiRequestTag) *CreateApiRequest {
	s.Tag = v
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

func (s *CreateApiRequest) Validate() error {
	return dara.Validate(s)
}

type CreateApiRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApiRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateApiRequestTag) GoString() string {
	return s.String()
}

func (s *CreateApiRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateApiRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateApiRequestTag) SetKey(v string) *CreateApiRequestTag {
	s.Key = &v
	return s
}

func (s *CreateApiRequestTag) SetValue(v string) *CreateApiRequestTag {
	s.Value = &v
	return s
}

func (s *CreateApiRequestTag) Validate() error {
	return dara.Validate(s)
}
