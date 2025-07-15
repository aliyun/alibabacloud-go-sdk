// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSignatureMethod(v string) *ModifyApiRequest
	GetAllowSignatureMethod() *string
	SetApiId(v string) *ModifyApiRequest
	GetApiId() *string
	SetApiName(v string) *ModifyApiRequest
	GetApiName() *string
	SetAppCodeAuthType(v string) *ModifyApiRequest
	GetAppCodeAuthType() *string
	SetAuthType(v string) *ModifyApiRequest
	GetAuthType() *string
	SetBackendEnable(v bool) *ModifyApiRequest
	GetBackendEnable() *bool
	SetBackendId(v string) *ModifyApiRequest
	GetBackendId() *string
	SetConstantParameters(v string) *ModifyApiRequest
	GetConstantParameters() *string
	SetDescription(v string) *ModifyApiRequest
	GetDescription() *string
	SetDisableInternet(v bool) *ModifyApiRequest
	GetDisableInternet() *bool
	SetErrorCodeSamples(v string) *ModifyApiRequest
	GetErrorCodeSamples() *string
	SetFailResultSample(v string) *ModifyApiRequest
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *ModifyApiRequest
	GetForceNonceCheck() *bool
	SetGroupId(v string) *ModifyApiRequest
	GetGroupId() *string
	SetOpenIdConnectConfig(v string) *ModifyApiRequest
	GetOpenIdConnectConfig() *string
	SetRequestConfig(v string) *ModifyApiRequest
	GetRequestConfig() *string
	SetRequestParameters(v string) *ModifyApiRequest
	GetRequestParameters() *string
	SetResultBodyModel(v string) *ModifyApiRequest
	GetResultBodyModel() *string
	SetResultDescriptions(v string) *ModifyApiRequest
	GetResultDescriptions() *string
	SetResultSample(v string) *ModifyApiRequest
	GetResultSample() *string
	SetResultType(v string) *ModifyApiRequest
	GetResultType() *string
	SetSecurityToken(v string) *ModifyApiRequest
	GetSecurityToken() *string
	SetServiceConfig(v string) *ModifyApiRequest
	GetServiceConfig() *string
	SetServiceParameters(v string) *ModifyApiRequest
	GetServiceParameters() *string
	SetServiceParametersMap(v string) *ModifyApiRequest
	GetServiceParametersMap() *string
	SetSystemParameters(v string) *ModifyApiRequest
	GetSystemParameters() *string
	SetVisibility(v string) *ModifyApiRequest
	GetVisibility() *string
	SetWebSocketApiType(v string) *ModifyApiRequest
	GetWebSocketApiType() *string
}

type ModifyApiRequest struct {
	// The type of the two-way communication API. Valid values:
	//
	// 	- **COMMON**: general APIs
	//
	// 	- **REGISTER**: registered APIs
	//
	// 	- **UNREGISTER**: unregistered APIs
	//
	// 	- **NOTIFY**: downstream notification
	//
	// example:
	//
	// HmacSHA256
	AllowSignatureMethod *string `json:"AllowSignatureMethod,omitempty" xml:"AllowSignatureMethod,omitempty"`
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8afff6c8c4c6447abb035812e4d66b65
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API. The name must be unique within the API group. The name must be 4 to 50 characters in length. It must start with a letter and can contain letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// ApiName
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the backend service.
	//
	// example:
	//
	// DEFAULT
	AppCodeAuthType *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	// The configuration items of API requests sent by the consumer to API Gateway.
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Configuration Mode
	//
	// example:
	//
	// true
	BackendEnable *bool `json:"BackendEnable,omitempty" xml:"BackendEnable,omitempty"`
	// Specifies whether to enable backend services.
	//
	// example:
	//
	// 0d105f80a8f340408bd34954d4e4ff22
	BackendId          *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	ConstantParameters *string `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty"`
	// The description of the API. The description can be up to 180 characters in length.
	//
	// example:
	//
	// Api description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The AppCode authentication type supported. Valid values:
	//
	// 	- DEFAULT: supported after being made available in Alibaba Cloud Marketplace
	//
	// 	- DISABLE: not supported
	//
	// 	- HEADER : supported only in the Header parameter
	//
	// 	- HEADER_QUERY: supported in the Header or Query parameter
	//
	// example:
	//
	// true
	DisableInternet  *bool   `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	ErrorCodeSamples *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	FailResultSample *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	// 	- Specifies whether to set DisableInternet to **true*	- to limit API calls to within the VPC.
	//
	// 	- If you set DisableInternet to **false**, the limit if lifted.
	//
	// >  If you do not set this parameter, the original value is used.
	//
	// example:
	//
	// true
	ForceNonceCheck *bool `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 927d50c0f2e54b359919923d908bb015
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The client-side request signature method of the API. Valid values:
	//
	// 	- HmacSHA256
	//
	// 	- HmacSHA1,HmacSHA256
	//
	// example:
	//
	// {\\"OpenIdApiType\\":\\"IDTOKEN\\",\\"PublicKey\\":\\"lzlj1573\\",\\"IdTokenParamName\\":\\"\\",\\"PublicKeyId\\":\\"lzljorders\\"}
	OpenIdConnectConfig *string `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty"`
	// The configuration items of API requests sent by API Gateway to the backend service.
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
	// 	- If you set **ForceNonceCheck*	- to **false**, the check is not performed. If you do not set this parameter, the original value is used.
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
	// This parameter is required.
	//
	// example:
	//
	// {"ServiceProtocol":"HTTP","ServiceHttpMethod":"GET","ServiceAddress":"http://www.customerdomain.com","ServiceTimeout":"1000","ServicePath":"/v3/getUserTest/[userId]"}
	ServiceConfig        *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	ServiceParameters    *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	SystemParameters     *string `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty"`
	// Specifies whether the API is public. Valid values:
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
	// The returned description of the API.
	//
	// example:
	//
	// COMMON
	WebSocketApiType *string `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s ModifyApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiRequest) GetAllowSignatureMethod() *string {
	return s.AllowSignatureMethod
}

func (s *ModifyApiRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ModifyApiRequest) GetApiName() *string {
	return s.ApiName
}

func (s *ModifyApiRequest) GetAppCodeAuthType() *string {
	return s.AppCodeAuthType
}

func (s *ModifyApiRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ModifyApiRequest) GetBackendEnable() *bool {
	return s.BackendEnable
}

func (s *ModifyApiRequest) GetBackendId() *string {
	return s.BackendId
}

func (s *ModifyApiRequest) GetConstantParameters() *string {
	return s.ConstantParameters
}

func (s *ModifyApiRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyApiRequest) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *ModifyApiRequest) GetErrorCodeSamples() *string {
	return s.ErrorCodeSamples
}

func (s *ModifyApiRequest) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *ModifyApiRequest) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *ModifyApiRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyApiRequest) GetOpenIdConnectConfig() *string {
	return s.OpenIdConnectConfig
}

func (s *ModifyApiRequest) GetRequestConfig() *string {
	return s.RequestConfig
}

func (s *ModifyApiRequest) GetRequestParameters() *string {
	return s.RequestParameters
}

func (s *ModifyApiRequest) GetResultBodyModel() *string {
	return s.ResultBodyModel
}

func (s *ModifyApiRequest) GetResultDescriptions() *string {
	return s.ResultDescriptions
}

func (s *ModifyApiRequest) GetResultSample() *string {
	return s.ResultSample
}

func (s *ModifyApiRequest) GetResultType() *string {
	return s.ResultType
}

func (s *ModifyApiRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyApiRequest) GetServiceConfig() *string {
	return s.ServiceConfig
}

func (s *ModifyApiRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ModifyApiRequest) GetServiceParametersMap() *string {
	return s.ServiceParametersMap
}

func (s *ModifyApiRequest) GetSystemParameters() *string {
	return s.SystemParameters
}

func (s *ModifyApiRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *ModifyApiRequest) GetWebSocketApiType() *string {
	return s.WebSocketApiType
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

func (s *ModifyApiRequest) SetBackendEnable(v bool) *ModifyApiRequest {
	s.BackendEnable = &v
	return s
}

func (s *ModifyApiRequest) SetBackendId(v string) *ModifyApiRequest {
	s.BackendId = &v
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

func (s *ModifyApiRequest) Validate() error {
	return dara.Validate(s)
}
