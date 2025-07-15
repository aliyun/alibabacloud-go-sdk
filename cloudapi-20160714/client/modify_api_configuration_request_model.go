// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSignatureMethod(v string) *ModifyApiConfigurationRequest
	GetAllowSignatureMethod() *string
	SetApiId(v string) *ModifyApiConfigurationRequest
	GetApiId() *string
	SetApiName(v string) *ModifyApiConfigurationRequest
	GetApiName() *string
	SetAppCodeAuthType(v string) *ModifyApiConfigurationRequest
	GetAppCodeAuthType() *string
	SetAuthType(v string) *ModifyApiConfigurationRequest
	GetAuthType() *string
	SetBackendName(v string) *ModifyApiConfigurationRequest
	GetBackendName() *string
	SetBodyFormat(v string) *ModifyApiConfigurationRequest
	GetBodyFormat() *string
	SetBodyModel(v string) *ModifyApiConfigurationRequest
	GetBodyModel() *string
	SetContentTypeCategory(v string) *ModifyApiConfigurationRequest
	GetContentTypeCategory() *string
	SetContentTypeValue(v string) *ModifyApiConfigurationRequest
	GetContentTypeValue() *string
	SetDescription(v string) *ModifyApiConfigurationRequest
	GetDescription() *string
	SetDisableInternet(v bool) *ModifyApiConfigurationRequest
	GetDisableInternet() *bool
	SetErrorCodeSamples(v string) *ModifyApiConfigurationRequest
	GetErrorCodeSamples() *string
	SetFailResultSample(v string) *ModifyApiConfigurationRequest
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *ModifyApiConfigurationRequest
	GetForceNonceCheck() *bool
	SetFunctionComputeConfig(v string) *ModifyApiConfigurationRequest
	GetFunctionComputeConfig() *string
	SetHttpConfig(v string) *ModifyApiConfigurationRequest
	GetHttpConfig() *string
	SetMockConfig(v string) *ModifyApiConfigurationRequest
	GetMockConfig() *string
	SetModelName(v string) *ModifyApiConfigurationRequest
	GetModelName() *string
	SetOssConfig(v string) *ModifyApiConfigurationRequest
	GetOssConfig() *string
	SetPostBodyDescription(v string) *ModifyApiConfigurationRequest
	GetPostBodyDescription() *string
	SetRequestHttpMethod(v string) *ModifyApiConfigurationRequest
	GetRequestHttpMethod() *string
	SetRequestMode(v string) *ModifyApiConfigurationRequest
	GetRequestMode() *string
	SetRequestParameters(v string) *ModifyApiConfigurationRequest
	GetRequestParameters() *string
	SetRequestPath(v string) *ModifyApiConfigurationRequest
	GetRequestPath() *string
	SetRequestProtocol(v string) *ModifyApiConfigurationRequest
	GetRequestProtocol() *string
	SetResultSample(v string) *ModifyApiConfigurationRequest
	GetResultSample() *string
	SetResultType(v string) *ModifyApiConfigurationRequest
	GetResultType() *string
	SetSecurityToken(v string) *ModifyApiConfigurationRequest
	GetSecurityToken() *string
	SetServiceParameters(v string) *ModifyApiConfigurationRequest
	GetServiceParameters() *string
	SetServiceParametersMap(v string) *ModifyApiConfigurationRequest
	GetServiceParametersMap() *string
	SetServiceProtocol(v string) *ModifyApiConfigurationRequest
	GetServiceProtocol() *string
	SetServiceTimeout(v int32) *ModifyApiConfigurationRequest
	GetServiceTimeout() *int32
	SetUseBackendService(v bool) *ModifyApiConfigurationRequest
	GetUseBackendService() *bool
	SetVisibility(v string) *ModifyApiConfigurationRequest
	GetVisibility() *string
	SetVpcConfig(v string) *ModifyApiConfigurationRequest
	GetVpcConfig() *string
}

type ModifyApiConfigurationRequest struct {
	// If the **AuthType*	- parameter is set to **APP**, you must include this parameter to specify the signature algorithm. If you do not specify a value, HmacSHA256 is used by default. Valid values:
	//
	// 	- HmacSHA256
	//
	// 	- HmacSHA1,HmacSHA256
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
	// baacc592e63a4cb6a41920d9d3f91f38
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// testModifyApiName
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// If the **AuthType*	- parameter is set to **APP**, the valid values are:
	//
	// 	- **DEFAULT**: The default value that is used if no other values are passed. This value indicates that the settings of the group are used.
	//
	// 	- **DISABLE**: The authentication is disabled.
	//
	// 	- **HEADER**: AppCode can be placed in the Header parameter for authentication.
	//
	// 	- **HEADER_QUERY**: AppCode can be placed in the Header or Query parameter for authentication.
	//
	// example:
	//
	// DEFAULT
	AppCodeAuthType *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	// API安全认证类型，目前可以取值：
	//
	// - **APP**：只允许已授权的APP调用
	//
	// - **ANONYMOUS**：允许匿名调用，设置为允许匿名调用需要注意：
	//
	//      - 任何能够获取该API服务信息的人，都将能够调用该API。网关不会对调用者做身份认证，也无法设置按用户的流量控制，若开放该API请设置好按API的流量控制；
	//
	//      - AppCodeAuthType的值不会生效。
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The name of the backend service. This parameter takes effect only when the UseBackendService parameter is set to TRUE.
	//
	// example:
	//
	// testBackendService
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// This parameter takes effect only when the **RequestMode*	- parameter is set to **MAPPING**.
	//
	// The format in which data is transmitted to the server for POST and PUT requests. Valid values: **FORM*	- and **STREAM**. FORM indicates that data is transmitted in the key-value pair format. STREAM indicates that data is transmitted as byte streams.
	//
	// example:
	//
	// STREAM
	BodyFormat *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	// The body model.
	//
	// example:
	//
	// https://apigateway.aliyun.com/models/f4e7333c****40dcbaf7c9da553ccd8d/3ab61f775b****d4bc35e993****87aa8
	BodyModel *string `json:"BodyModel,omitempty" xml:"BodyModel,omitempty"`
	// The ContentType configuration of the backend request.
	//
	// 	- DEFAULT: the default configuration in API Gateway
	//
	// 	- CUSTOM: a custom configuration
	//
	// example:
	//
	// DEFAULT
	ContentTypeCategory *string `json:"ContentTypeCategory,omitempty" xml:"ContentTypeCategory,omitempty"`
	// The value of the ContentType header when the ServiceProtocol parameter is set to HTTP and the ContentTypeCatagory parameter is set to DEFAULT or CUSTOM.
	//
	// example:
	//
	// application/x-www-form-urlencoded; charset=UTF-8
	ContentTypeValue *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// TestModifyDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 	- Specifies whether to call the API only in an internal network. If the **DisableInternet*	- parameter is set to **true**, the API can be called only in an internal network.
	//
	// 	- If the **DisableInternet*	- parameter is set to **false**, the API can be called over the Internet and in an internal network.
	//
	// example:
	//
	// false
	DisableInternet *bool `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	// The sample error codes returned by the backend service.
	//
	// For more information, see [ErrorCodeSample](https://help.aliyun.com/document_detail/44392.html).
	//
	// example:
	//
	// [{"Code":"400","Message":"Missing the userId","Description":"param invalid"}]
	ErrorCodeSamples *string `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty"`
	// The sample error response from the backend service. This value is used only to generate documents. It does not affect the returned result.
	//
	// example:
	//
	// {"errorCode":"fail","errorMessage":"param invalid"}
	FailResultSample *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	// 	- Specifies whether to forcibly check X-Ca-Nonce. If the **ForceNonceCheck*	- parameter is set to **true**, X-Ca-Nonce is forcibly checked. X-Ca-Nonce is the unique identifier of the request and is generally identified by UUID. After receiving this parameter, API Gateway verifies the validity of this parameter. The same value can be used only once within 15 minutes. This helps prevent replay attacks.
	//
	// 	- If the **ForceNonceCheck*	- parameter is set to **false**, X-Ca-Nonce is not checked. If you do not modify this parameter when you modify an API, the original value is used.
	//
	// example:
	//
	// true
	ForceNonceCheck *bool `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	// The Function Compute configuration.
	//
	// example:
	//
	// {"FcType":"FCEvent","FcRegionId":"cn-hangzhou","RoleArn":"acs:ram::xxxxxxxx:role/aliyunserviceroleforapigateway","selectServiceName":"fcTest","FunctionName":"funcTest","selectFunctionName":"funcTest","Qualifier":"LATEST","Path":"","FcBaseUrl":"","ServiceName":"fcTest"}
	FunctionComputeConfig *string `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty"`
	// The HTTP configuration.
	//
	// example:
	//
	// {"serviceAddress":"http://test.api.com","servicePath":"/test/api","serviceHttpMethod":"GET"}
	HttpConfig *string `json:"HttpConfig,omitempty" xml:"HttpConfig,omitempty"`
	// The Mock configuration.
	//
	// example:
	//
	// {"MockResult":"test","MockHeaders":[{"HeaderName":"testHeader","HeaderValue":"testHeader"}],"MockStatusCode":"400"}
	MockConfig *string `json:"MockConfig,omitempty" xml:"MockConfig,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The OSS configuration.
	//
	// example:
	//
	// {"OssRegionId":"cn-hangzhou","Key":"/test.html","BucketName":"test-api-oss","Action":"GetObject"}
	OssConfig *string `json:"OssConfig,omitempty" xml:"OssConfig,omitempty"`
	// The description of the request body.
	//
	// example:
	//
	// The description of the request body
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	// The HTTP method used to make the request. Valid values: GET, POST, DELETE, PUT, HEADER, TRACE, PATCH, CONNECT, and OPTIONS.
	//
	// example:
	//
	// GET
	RequestHttpMethod *string `json:"RequestHttpMethod,omitempty" xml:"RequestHttpMethod,omitempty"`
	// The request mode. Valid values:
	//
	// 	- MAPPING: Parameters are mapped. Unknown parameters are filtered out.
	//
	// 	- PASSTHROUGH: Parameters are passed through.
	//
	// 	- MAPPING_PASSTHROUGH: Parameters are mapped. Unknown parameters are passed through.
	//
	// example:
	//
	// MAPPING
	RequestMode *string `json:"RequestMode,omitempty" xml:"RequestMode,omitempty"`
	// The parameters of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestParameter](https://help.aliyun.com/document_detail/43986.html).
	//
	// example:
	//
	// [{"ParameterLocation":{"name":"Head","orderNumber":2},"ParameterType":"String","Required":"OPTIONAL","isHide":false,"ApiParameterName":"header1","DefaultValue":"123124","Location":"Head"},{"ParameterLocation":{"name":"Head","orderNumber":2},"ParameterType":"String","Required":"REQUIRED","isHide":false,"ApiParameterName":"header2","DefaultValue":"","Location":"Head"},{"ParameterLocation":{"name":"Query","orderNumber":3},"ParameterType":"String","Required":"OPTIONAL","isHide":false,"ApiParameterName":"query1","DefaultValue":"1245","Location":"Query"},{"ApiParameterName":"CaClientIp","ParameterLocation":{"name":"Query","orderNumber":0},"Location":"Query","ParameterType":"String","Required":"REQUIRED","Description":"ClientIP"},{"ApiParameterName":"testConstant","ParameterLocation":{"name":"Head","orderNumber":0},"Location":"Head","ParameterType":"String","Required":"REQUIRED","DefaultValue":"111"}]
	RequestParameters *string `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty"`
	// The path of the API request. If the complete API URL is `http://api.a.com:8080/object/add?key1=value1&key2=value2`, the path of the API request is `/object/add`.
	//
	// example:
	//
	// /test/api
	RequestPath *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	// The protocol type supported by the API. Valid values: HTTP and HTTPS. Separate multiple values with commas (,), such as "HTTP,HTTPS".
	//
	// example:
	//
	// HTTP
	RequestProtocol *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
	// The sample response from the backend service. This value is used only to generate documents. It does not affect the returned result.
	//
	// example:
	//
	// {\\n  \\"status\\": 0,\\n  \\"data\\": {\\n    \\"count\\": 1,\\n    \\"list\\": [\\n      \\"352\\"\\n    ]\\n  },\\n  \\"message\\": \\"success\\"\\n}
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// The format of the response from the backend service. Valid values: JSON, TEXT, BINARY, XML, and HTML. This value is used only to generate documents. It does not affect the returned result.
	//
	// example:
	//
	// JSON
	ResultType    *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The parameters of API requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceParameter](https://help.aliyun.com/document_detail/43988.html).
	//
	// example:
	//
	// [{"ServiceParameterName":"header1","Location":"Head","Type":"String","ParameterCatalog":"REQUEST"},{"ServiceParameterName":"header2","Location":"Query","Type":"String","ParameterCatalog":"REQUEST"},{"ServiceParameterName":"query1","Location":"Head","Type":"String","ParameterCatalog":"REQUEST"},{"ServiceParameterName":"ipp","Location":"Query","Type":"String","ParameterCatalog":"SYSTEM"},{"ServiceParameterName":"testConstant","Location":"Head","Type":"String","ParameterCatalog":"CONSTANT"}]
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	// The mappings between parameters of requests sent by the consumer to API Gateway and parameters of requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceParameterMap](https://help.aliyun.com/document_detail/43989.html).
	//
	// example:
	//
	// [{"ServiceParameterName":"header1","RequestParameterName":"header1"},{"ServiceParameterName":"header2","RequestParameterName":"header2"},{"ServiceParameterName":"query1","RequestParameterName":"query1"},{"ServiceParameterName":"ipp","RequestParameterName":"CaClientIp"},{"ServiceParameterName":"testConstant","RequestParameterName":"testConstant"}]
	ServiceParametersMap *string `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty"`
	// The protocol that is used to access backend services. Valid values:
	//
	// 	- Http: for backend services that use HTTP or HTTPS
	//
	// 	- Vpc: for backend services that use VPC
	//
	// 	- FC: for Function Compute
	//
	// 	- OSS: for Object Storage Service
	//
	// 	- Mock: for backend services that use the Mock mode
	//
	// 	- EventBridge: for EventBridge
	//
	// You must specify the config value for the corresponding backend service.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The timeout period of the backend service. Unit: milliseconds.
	//
	// example:
	//
	// 10000
	ServiceTimeout *int32 `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	// Specifies whether to use the information about the created backend service. Valid values:
	//
	// 	- TRUE: uses the information about the created backend service.
	//
	// 	- FALSE: uses the information about the custom backend service.
	//
	// example:
	//
	// TRUE
	UseBackendService *bool `json:"UseBackendService,omitempty" xml:"UseBackendService,omitempty"`
	// Specifies whether to make the API public. Valid values:
	//
	// 	- **PUBLIC:*	- The API is public. If this parameter is set to PUBLIC, the API is displayed on the APIs page for all users after the API is published to the production environment.
	//
	// 	- **PRIVATE:*	- The API is private. Private APIs are not displayed in the Alibaba Cloud Marketplace after the API group to which they belong is made available.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
	// The VPC configuration.
	//
	// example:
	//
	// {"VpcId":"vpc-xxxxxxx","Name":"testVpc","InstanceId":"i-p0ssssss","Port":80,"servicePath":"/test/vpc","serviceHttpMethod":"HEAD"}
	VpcConfig *string `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty"`
}

func (s ModifyApiConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyApiConfigurationRequest) GetAllowSignatureMethod() *string {
	return s.AllowSignatureMethod
}

func (s *ModifyApiConfigurationRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ModifyApiConfigurationRequest) GetApiName() *string {
	return s.ApiName
}

func (s *ModifyApiConfigurationRequest) GetAppCodeAuthType() *string {
	return s.AppCodeAuthType
}

func (s *ModifyApiConfigurationRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ModifyApiConfigurationRequest) GetBackendName() *string {
	return s.BackendName
}

func (s *ModifyApiConfigurationRequest) GetBodyFormat() *string {
	return s.BodyFormat
}

func (s *ModifyApiConfigurationRequest) GetBodyModel() *string {
	return s.BodyModel
}

func (s *ModifyApiConfigurationRequest) GetContentTypeCategory() *string {
	return s.ContentTypeCategory
}

func (s *ModifyApiConfigurationRequest) GetContentTypeValue() *string {
	return s.ContentTypeValue
}

func (s *ModifyApiConfigurationRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyApiConfigurationRequest) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *ModifyApiConfigurationRequest) GetErrorCodeSamples() *string {
	return s.ErrorCodeSamples
}

func (s *ModifyApiConfigurationRequest) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *ModifyApiConfigurationRequest) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *ModifyApiConfigurationRequest) GetFunctionComputeConfig() *string {
	return s.FunctionComputeConfig
}

func (s *ModifyApiConfigurationRequest) GetHttpConfig() *string {
	return s.HttpConfig
}

func (s *ModifyApiConfigurationRequest) GetMockConfig() *string {
	return s.MockConfig
}

func (s *ModifyApiConfigurationRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ModifyApiConfigurationRequest) GetOssConfig() *string {
	return s.OssConfig
}

func (s *ModifyApiConfigurationRequest) GetPostBodyDescription() *string {
	return s.PostBodyDescription
}

func (s *ModifyApiConfigurationRequest) GetRequestHttpMethod() *string {
	return s.RequestHttpMethod
}

func (s *ModifyApiConfigurationRequest) GetRequestMode() *string {
	return s.RequestMode
}

func (s *ModifyApiConfigurationRequest) GetRequestParameters() *string {
	return s.RequestParameters
}

func (s *ModifyApiConfigurationRequest) GetRequestPath() *string {
	return s.RequestPath
}

func (s *ModifyApiConfigurationRequest) GetRequestProtocol() *string {
	return s.RequestProtocol
}

func (s *ModifyApiConfigurationRequest) GetResultSample() *string {
	return s.ResultSample
}

func (s *ModifyApiConfigurationRequest) GetResultType() *string {
	return s.ResultType
}

func (s *ModifyApiConfigurationRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyApiConfigurationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ModifyApiConfigurationRequest) GetServiceParametersMap() *string {
	return s.ServiceParametersMap
}

func (s *ModifyApiConfigurationRequest) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *ModifyApiConfigurationRequest) GetServiceTimeout() *int32 {
	return s.ServiceTimeout
}

func (s *ModifyApiConfigurationRequest) GetUseBackendService() *bool {
	return s.UseBackendService
}

func (s *ModifyApiConfigurationRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *ModifyApiConfigurationRequest) GetVpcConfig() *string {
	return s.VpcConfig
}

func (s *ModifyApiConfigurationRequest) SetAllowSignatureMethod(v string) *ModifyApiConfigurationRequest {
	s.AllowSignatureMethod = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetApiId(v string) *ModifyApiConfigurationRequest {
	s.ApiId = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetApiName(v string) *ModifyApiConfigurationRequest {
	s.ApiName = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetAppCodeAuthType(v string) *ModifyApiConfigurationRequest {
	s.AppCodeAuthType = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetAuthType(v string) *ModifyApiConfigurationRequest {
	s.AuthType = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetBackendName(v string) *ModifyApiConfigurationRequest {
	s.BackendName = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetBodyFormat(v string) *ModifyApiConfigurationRequest {
	s.BodyFormat = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetBodyModel(v string) *ModifyApiConfigurationRequest {
	s.BodyModel = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetContentTypeCategory(v string) *ModifyApiConfigurationRequest {
	s.ContentTypeCategory = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetContentTypeValue(v string) *ModifyApiConfigurationRequest {
	s.ContentTypeValue = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetDescription(v string) *ModifyApiConfigurationRequest {
	s.Description = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetDisableInternet(v bool) *ModifyApiConfigurationRequest {
	s.DisableInternet = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetErrorCodeSamples(v string) *ModifyApiConfigurationRequest {
	s.ErrorCodeSamples = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetFailResultSample(v string) *ModifyApiConfigurationRequest {
	s.FailResultSample = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetForceNonceCheck(v bool) *ModifyApiConfigurationRequest {
	s.ForceNonceCheck = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetFunctionComputeConfig(v string) *ModifyApiConfigurationRequest {
	s.FunctionComputeConfig = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetHttpConfig(v string) *ModifyApiConfigurationRequest {
	s.HttpConfig = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetMockConfig(v string) *ModifyApiConfigurationRequest {
	s.MockConfig = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetModelName(v string) *ModifyApiConfigurationRequest {
	s.ModelName = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetOssConfig(v string) *ModifyApiConfigurationRequest {
	s.OssConfig = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetPostBodyDescription(v string) *ModifyApiConfigurationRequest {
	s.PostBodyDescription = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetRequestHttpMethod(v string) *ModifyApiConfigurationRequest {
	s.RequestHttpMethod = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetRequestMode(v string) *ModifyApiConfigurationRequest {
	s.RequestMode = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetRequestParameters(v string) *ModifyApiConfigurationRequest {
	s.RequestParameters = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetRequestPath(v string) *ModifyApiConfigurationRequest {
	s.RequestPath = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetRequestProtocol(v string) *ModifyApiConfigurationRequest {
	s.RequestProtocol = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetResultSample(v string) *ModifyApiConfigurationRequest {
	s.ResultSample = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetResultType(v string) *ModifyApiConfigurationRequest {
	s.ResultType = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetSecurityToken(v string) *ModifyApiConfigurationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetServiceParameters(v string) *ModifyApiConfigurationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetServiceParametersMap(v string) *ModifyApiConfigurationRequest {
	s.ServiceParametersMap = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetServiceProtocol(v string) *ModifyApiConfigurationRequest {
	s.ServiceProtocol = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetServiceTimeout(v int32) *ModifyApiConfigurationRequest {
	s.ServiceTimeout = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetUseBackendService(v bool) *ModifyApiConfigurationRequest {
	s.UseBackendService = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetVisibility(v string) *ModifyApiConfigurationRequest {
	s.Visibility = &v
	return s
}

func (s *ModifyApiConfigurationRequest) SetVpcConfig(v string) *ModifyApiConfigurationRequest {
	s.VpcConfig = &v
	return s
}

func (s *ModifyApiConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
