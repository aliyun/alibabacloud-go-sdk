// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeployedApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSignatureMethod(v string) *DescribeDeployedApiResponseBody
	GetAllowSignatureMethod() *string
	SetApiId(v string) *DescribeDeployedApiResponseBody
	GetApiId() *string
	SetApiName(v string) *DescribeDeployedApiResponseBody
	GetApiName() *string
	SetAuthType(v string) *DescribeDeployedApiResponseBody
	GetAuthType() *string
	SetConstantParameters(v *DescribeDeployedApiResponseBodyConstantParameters) *DescribeDeployedApiResponseBody
	GetConstantParameters() *DescribeDeployedApiResponseBodyConstantParameters
	SetCustomSystemParameters(v *DescribeDeployedApiResponseBodyCustomSystemParameters) *DescribeDeployedApiResponseBody
	GetCustomSystemParameters() *DescribeDeployedApiResponseBodyCustomSystemParameters
	SetDeployedTime(v string) *DescribeDeployedApiResponseBody
	GetDeployedTime() *string
	SetDescription(v string) *DescribeDeployedApiResponseBody
	GetDescription() *string
	SetDisableInternet(v bool) *DescribeDeployedApiResponseBody
	GetDisableInternet() *bool
	SetErrorCodeSamples(v *DescribeDeployedApiResponseBodyErrorCodeSamples) *DescribeDeployedApiResponseBody
	GetErrorCodeSamples() *DescribeDeployedApiResponseBodyErrorCodeSamples
	SetFailResultSample(v string) *DescribeDeployedApiResponseBody
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *DescribeDeployedApiResponseBody
	GetForceNonceCheck() *bool
	SetGroupId(v string) *DescribeDeployedApiResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeDeployedApiResponseBody
	GetGroupName() *string
	SetOpenIdConnectConfig(v *DescribeDeployedApiResponseBodyOpenIdConnectConfig) *DescribeDeployedApiResponseBody
	GetOpenIdConnectConfig() *DescribeDeployedApiResponseBodyOpenIdConnectConfig
	SetRegionId(v string) *DescribeDeployedApiResponseBody
	GetRegionId() *string
	SetRequestConfig(v *DescribeDeployedApiResponseBodyRequestConfig) *DescribeDeployedApiResponseBody
	GetRequestConfig() *DescribeDeployedApiResponseBodyRequestConfig
	SetRequestId(v string) *DescribeDeployedApiResponseBody
	GetRequestId() *string
	SetRequestParameters(v *DescribeDeployedApiResponseBodyRequestParameters) *DescribeDeployedApiResponseBody
	GetRequestParameters() *DescribeDeployedApiResponseBodyRequestParameters
	SetResultBodyModel(v string) *DescribeDeployedApiResponseBody
	GetResultBodyModel() *string
	SetResultDescriptions(v *DescribeDeployedApiResponseBodyResultDescriptions) *DescribeDeployedApiResponseBody
	GetResultDescriptions() *DescribeDeployedApiResponseBodyResultDescriptions
	SetResultSample(v string) *DescribeDeployedApiResponseBody
	GetResultSample() *string
	SetResultType(v string) *DescribeDeployedApiResponseBody
	GetResultType() *string
	SetServiceConfig(v *DescribeDeployedApiResponseBodyServiceConfig) *DescribeDeployedApiResponseBody
	GetServiceConfig() *DescribeDeployedApiResponseBodyServiceConfig
	SetServiceParameters(v *DescribeDeployedApiResponseBodyServiceParameters) *DescribeDeployedApiResponseBody
	GetServiceParameters() *DescribeDeployedApiResponseBodyServiceParameters
	SetServiceParametersMap(v *DescribeDeployedApiResponseBodyServiceParametersMap) *DescribeDeployedApiResponseBody
	GetServiceParametersMap() *DescribeDeployedApiResponseBodyServiceParametersMap
	SetStageName(v string) *DescribeDeployedApiResponseBody
	GetStageName() *string
	SetSystemParameters(v *DescribeDeployedApiResponseBodySystemParameters) *DescribeDeployedApiResponseBody
	GetSystemParameters() *DescribeDeployedApiResponseBodySystemParameters
	SetVisibility(v string) *DescribeDeployedApiResponseBody
	GetVisibility() *string
}

type DescribeDeployedApiResponseBody struct {
	// The signature method used by the client. Valid values:
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
	// example:
	//
	// 4eed13a57d4e42fbb51316be8a5329ff
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API
	//
	// example:
	//
	// weather
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The security authentication method of the API. Valid values:
	//
	// 	- **APP: Only authorized applications can call the API.**
	//
	// 	- **ANONYMOUS: The API can be anonymously called. In this mode, you must take note of the following rules:**
	//
	//     	- All users who have obtained the API service information can call this API. API Gateway does not authenticate callers and cannot set user-specific throttling policies. If you make this API public, set API-specific throttling policies.
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The constant parameters.
	ConstantParameters *DescribeDeployedApiResponseBodyConstantParameters `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	// The custom system parameters.
	CustomSystemParameters *DescribeDeployedApiResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	// The deployment time. Format: yyyy-mm-ddhh:mm:ss.
	//
	// example:
	//
	// 2022-07-25T17:47:51Z
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Api description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 	- Specifies whether to set DisableInternet to **true*	- to limit API calls to within the VPC.
	//
	// 	- If you set DisableInternet to **false**, the limit is lifted.
	//
	// >  If you do not set this parameter, the original value is not modified.
	//
	// example:
	//
	// true
	DisableInternet *bool `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	// The sample error codes returned by the backend service.
	//
	// For more information, see [ErrorCodeSample](https://help.aliyun.com/document_detail/44392.html).
	ErrorCodeSamples *DescribeDeployedApiResponseBodyErrorCodeSamples `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	// The sample error response from the backend service.
	//
	// example:
	//
	// {"errorCode":"fail","errorMessage":"param invalid"}
	FailResultSample *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	// 	- Specifies whether to set **ForceNonceCheck*	- to **true*	- to force the check of X-Ca-Nonce during the request. This is the unique identifier of the request and is generally identified by UUID. After receiving this parameter, API Gateway verifies the validity of this parameter. The same value can be used only once within 15 minutes. This helps prevent replay attacks.
	//
	// 	- If you set **ForceNonceCheck*	- to **false**, the check is not performed. The default value is false when you create an API.
	//
	// example:
	//
	// true
	ForceNonceCheck *bool `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// bc77f5b49c974437a9912ea3755cd834
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// Weather
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The configuration items of the third-party OpenID Connect authentication method.
	OpenIdConnectConfig *DescribeDeployedApiResponseBodyOpenIdConnectConfig `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	// The region to which the API group belongs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Configuration items of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestConfig](https://help.aliyun.com/document_detail/43985.html).
	RequestConfig *DescribeDeployedApiResponseBodyRequestConfig `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The parameters of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestParameter](https://help.aliyun.com/document_detail/43986.html).
	RequestParameters *DescribeDeployedApiResponseBodyRequestParameters `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	// The return description of the API.
	//
	// example:
	//
	// {}
	ResultBodyModel *string `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	// The return description of the API.
	ResultDescriptions *DescribeDeployedApiResponseBodyResultDescriptions `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	// The sample response from the backend service.
	//
	// example:
	//
	// {code: 200, message:\\"success\\", data: \\"\\"}
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// The format of the response from the backend service. Valid values: JSON, TEXT, BINARY, XML, and HTML. Default value: JSON.
	//
	// example:
	//
	// HTML
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The configuration items of API requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceConfig](https://help.aliyun.com/document_detail/43987.html).
	ServiceConfig *DescribeDeployedApiResponseBodyServiceConfig `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	// The parameters of API requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceParameter](https://help.aliyun.com/document_detail/43988.html).
	ServiceParameters *DescribeDeployedApiResponseBodyServiceParameters `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty" type:"Struct"`
	// The mappings between parameters of requests sent by the consumer to API Gateway and parameters of requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceParameterMap](https://help.aliyun.com/document_detail/43989.html).
	ServiceParametersMap *DescribeDeployedApiResponseBodyServiceParametersMap `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty" type:"Struct"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE: the pre-release environment**
	//
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// System parameters
	SystemParameters *DescribeDeployedApiResponseBodySystemParameters `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	// Specifies whether to make the API public. Valid values:
	//
	// 	- **PUBLIC**: Make the API public. If you set this parameter to PUBLIC, this API is displayed on the APIs page for all users after the API is published to the production environment.**
	//
	// 	- **PRIVATE**: Make the API private. Private APIs are not displayed in the Alibaba Cloud Marketplace after the API group to which they belong is made available.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeDeployedApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBody) GetAllowSignatureMethod() *string {
	return s.AllowSignatureMethod
}

func (s *DescribeDeployedApiResponseBody) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeDeployedApiResponseBody) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeDeployedApiResponseBody) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeDeployedApiResponseBody) GetConstantParameters() *DescribeDeployedApiResponseBodyConstantParameters {
	return s.ConstantParameters
}

func (s *DescribeDeployedApiResponseBody) GetCustomSystemParameters() *DescribeDeployedApiResponseBodyCustomSystemParameters {
	return s.CustomSystemParameters
}

func (s *DescribeDeployedApiResponseBody) GetDeployedTime() *string {
	return s.DeployedTime
}

func (s *DescribeDeployedApiResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBody) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *DescribeDeployedApiResponseBody) GetErrorCodeSamples() *DescribeDeployedApiResponseBodyErrorCodeSamples {
	return s.ErrorCodeSamples
}

func (s *DescribeDeployedApiResponseBody) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *DescribeDeployedApiResponseBody) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *DescribeDeployedApiResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeDeployedApiResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDeployedApiResponseBody) GetOpenIdConnectConfig() *DescribeDeployedApiResponseBodyOpenIdConnectConfig {
	return s.OpenIdConnectConfig
}

func (s *DescribeDeployedApiResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeployedApiResponseBody) GetRequestConfig() *DescribeDeployedApiResponseBodyRequestConfig {
	return s.RequestConfig
}

func (s *DescribeDeployedApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeployedApiResponseBody) GetRequestParameters() *DescribeDeployedApiResponseBodyRequestParameters {
	return s.RequestParameters
}

func (s *DescribeDeployedApiResponseBody) GetResultBodyModel() *string {
	return s.ResultBodyModel
}

func (s *DescribeDeployedApiResponseBody) GetResultDescriptions() *DescribeDeployedApiResponseBodyResultDescriptions {
	return s.ResultDescriptions
}

func (s *DescribeDeployedApiResponseBody) GetResultSample() *string {
	return s.ResultSample
}

func (s *DescribeDeployedApiResponseBody) GetResultType() *string {
	return s.ResultType
}

func (s *DescribeDeployedApiResponseBody) GetServiceConfig() *DescribeDeployedApiResponseBodyServiceConfig {
	return s.ServiceConfig
}

func (s *DescribeDeployedApiResponseBody) GetServiceParameters() *DescribeDeployedApiResponseBodyServiceParameters {
	return s.ServiceParameters
}

func (s *DescribeDeployedApiResponseBody) GetServiceParametersMap() *DescribeDeployedApiResponseBodyServiceParametersMap {
	return s.ServiceParametersMap
}

func (s *DescribeDeployedApiResponseBody) GetStageName() *string {
	return s.StageName
}

func (s *DescribeDeployedApiResponseBody) GetSystemParameters() *DescribeDeployedApiResponseBodySystemParameters {
	return s.SystemParameters
}

func (s *DescribeDeployedApiResponseBody) GetVisibility() *string {
	return s.Visibility
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

func (s *DescribeDeployedApiResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeDeployedApiResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyConstantParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyConstantParameters) GetConstantParameter() []*DescribeDeployedApiResponseBodyConstantParametersConstantParameter {
	return s.ConstantParameter
}

func (s *DescribeDeployedApiResponseBodyConstantParameters) SetConstantParameter(v []*DescribeDeployedApiResponseBodyConstantParametersConstantParameter) *DescribeDeployedApiResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

func (s *DescribeDeployedApiResponseBodyConstantParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyConstantParametersConstantParameter struct {
	// The constant value.
	//
	// example:
	//
	// constance
	ConstantValue *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	// The description.
	//
	// example:
	//
	// 123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the backend service parameter.
	//
	// example:
	//
	// constance
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyConstantParametersConstantParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) GetConstantValue() *string {
	return s.ConstantValue
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeDeployedApiResponseBodyConstantParametersConstantParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyCustomSystemParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParameters) GetCustomSystemParameter() []*DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter {
	return s.CustomSystemParameter
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeDeployedApiResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter struct {
	// Example
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The description.
	//
	// example:
	//
	// 123
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the custom system parameter.
	//
	// example:
	//
	// appid
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The name of the corresponding backend parameter.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeDeployedApiResponseBodyCustomSystemParametersCustomSystemParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamples) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamples) GetErrorCodeSample() []*DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample {
	return s.ErrorCodeSample
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeDeployedApiResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamples) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample struct {
	// The error code.
	//
	// example:
	//
	// Error
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) GetCode() *string {
	return s.Code
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) GetMessage() *string {
	return s.Message
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

func (s *DescribeDeployedApiResponseBodyErrorCodeSamplesErrorCodeSample) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyOpenIdConnectConfig struct {
	// The name of the parameter that corresponds to the token.
	//
	// example:
	//
	// xxx
	IdTokenParamName *string `json:"IdTokenParamName,omitempty" xml:"IdTokenParamName,omitempty"`
	// The configuration of OpenID Connect authentication. Valid values:
	//
	// 	- **IDTOKEN: indicates the APIs that are called by clients to obtain tokens. If you specify this value, the PublicKeyId parameter and the PublicKey parameter are required.**
	//
	// 	- **BUSINESS: indicates business APIs. Tokens are used to call the business APIs. If you specify this value, the IdTokenParamName parameter is required.
	//
	// example:
	//
	// IDTOKEN
	OpenIdApiType *string `json:"OpenIdApiType,omitempty" xml:"OpenIdApiType,omitempty"`
	// The public key of the API.
	//
	// example:
	//
	// EB1837F8693CCED0BF750B3AD48467BEB569E780A14591CF92
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
	// The ID of the public key.
	//
	// example:
	//
	// 88483727556929326703309904351185815489
	PublicKeyId *string `json:"PublicKeyId,omitempty" xml:"PublicKeyId,omitempty"`
}

func (s DescribeDeployedApiResponseBodyOpenIdConnectConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) GetIdTokenParamName() *string {
	return s.IdTokenParamName
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) GetOpenIdApiType() *string {
	return s.OpenIdApiType
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) GetPublicKey() *string {
	return s.PublicKey
}

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) GetPublicKeyId() *string {
	return s.PublicKeyId
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

func (s *DescribeDeployedApiResponseBodyOpenIdConnectConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyRequestConfig struct {
	// This parameter takes effect only when the RequestMode parameter is set to MAPPING.
	//
	// The server data transmission method used for POST and PUT requests. Valid values: FORM and STREAM. FORM indicates that data in key-value pairs is transmitted as forms. STREAM indicates that data is transmitted as byte streams.
	//
	// example:
	//
	// STREAM
	BodyFormat *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	// The body model.
	//
	// example:
	//
	// https://apigateway.aliyun.com/models/3a240a127dccXXXXXXXX947b4095/9e2df550e85b4121a79XXXXXxaab
	BodyModel *string `json:"BodyModel,omitempty" xml:"BodyModel,omitempty"`
	// The description of the request body.
	//
	// example:
	//
	// fwefwef
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	// The HTTP method used to make the request. Valid values: GET, POST, DELETE, PUT, HEADER, TRACE, PATCH, CONNECT, and OPTIONS.
	//
	// example:
	//
	// POST
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
	// The API request path. If the complete API URL is `http://api.a.com:8080/object/add?key1=value1&key2=value2`, the API request path is ` /object/add  `.
	//
	// example:
	//
	// /api/billing/test/[type]
	RequestPath *string `json:"RequestPath,omitempty" xml:"RequestPath,omitempty"`
	// The protocol type supported by the API. Valid values: HTTP, HTTPS, and WebSocket. Separate multiple values with commas (,), such as "HTTP,HTTPS".
	//
	// example:
	//
	// HTTP
	RequestProtocol *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeDeployedApiResponseBodyRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetBodyFormat() *string {
	return s.BodyFormat
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetBodyModel() *string {
	return s.BodyModel
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetPostBodyDescription() *string {
	return s.PostBodyDescription
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetRequestHttpMethod() *string {
	return s.RequestHttpMethod
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetRequestMode() *string {
	return s.RequestMode
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetRequestPath() *string {
	return s.RequestPath
}

func (s *DescribeDeployedApiResponseBodyRequestConfig) GetRequestProtocol() *string {
	return s.RequestProtocol
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

func (s *DescribeDeployedApiResponseBodyRequestConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyRequestParameters struct {
	RequestParameter []*DescribeDeployedApiResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestParameters) GetRequestParameter() []*DescribeDeployedApiResponseBodyRequestParametersRequestParameter {
	return s.RequestParameter
}

func (s *DescribeDeployedApiResponseBodyRequestParameters) SetRequestParameter(v []*DescribeDeployedApiResponseBodyRequestParametersRequestParameter) *DescribeDeployedApiResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

func (s *DescribeDeployedApiResponseBodyRequestParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyRequestParametersRequestParameter struct {
	// The name of the API parameter.
	//
	// example:
	//
	// age
	ApiParameterName *string `json:"ApiParameterName,omitempty" xml:"ApiParameterName,omitempty"`
	// The type of the array element.
	//
	// example:
	//
	// String
	ArrayItemsType *string `json:"ArrayItemsType,omitempty" xml:"ArrayItemsType,omitempty"`
	// The default value.
	//
	// example:
	//
	// 20
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Example
	//
	// example:
	//
	// 20
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// Description
	//
	// example:
	//
	// parameter description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The order in the document.
	//
	// example:
	//
	// 0
	DocOrder *int32 `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	// Specifies whether the document is public. Valid values: PUBLIC and PRIVATE.
	//
	// example:
	//
	// PUBLIC
	DocShow *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	// The hash values that can be entered when ParameterType is set to Int, Long, Float, Double, or String. Separate different values with commas (,), such as 1,2,3,4,9 or A,B,C,E,F.
	//
	// example:
	//
	// boy,girl
	EnumValue *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
	// JSON scheme
	//
	// example:
	//
	// {}
	JsonScheme *string `json:"JsonScheme,omitempty" xml:"JsonScheme,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The maximum parameter length when ParameterType is set to String.
	//
	// example:
	//
	// 123456
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The maximum parameter value when ParameterType is set to Int, Long, Float, or Double.
	//
	// example:
	//
	// 123456
	MaxValue *int64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum parameter length when ParameterType is set to String.
	//
	// example:
	//
	// 123456
	MinLength *int64 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// The minimum parameter value when ParameterType is set to Int, Long, Float, or Double.
	//
	// example:
	//
	// 123456
	MinValue *int64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The type of a request parameter. Valid values: String, Int, Long, Float, Double, and Boolean.
	//
	// example:
	//
	// String
	ParameterType *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	// The regular expression used for parameter validation when ParameterType is set to String.
	//
	// example:
	//
	// xxx
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	// Indicates whether the parameter is required. Valid values: REQUIRED and OPTIONAL.
	//
	// example:
	//
	// OPTIONAL
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeDeployedApiResponseBodyRequestParametersRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetApiParameterName() *string {
	return s.ApiParameterName
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetArrayItemsType() *string {
	return s.ArrayItemsType
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetDocOrder() *int32 {
	return s.DocOrder
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetDocShow() *string {
	return s.DocShow
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetJsonScheme() *string {
	return s.JsonScheme
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetMinLength() *int64 {
	return s.MinLength
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetMinValue() *int64 {
	return s.MinValue
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetRegularExpression() *string {
	return s.RegularExpression
}

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) GetRequired() *string {
	return s.Required
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

func (s *DescribeDeployedApiResponseBodyRequestParametersRequestParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyResultDescriptions struct {
	ResultDescription []*DescribeDeployedApiResponseBodyResultDescriptionsResultDescription `json:"ResultDescription,omitempty" xml:"ResultDescription,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyResultDescriptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyResultDescriptions) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyResultDescriptions) GetResultDescription() []*DescribeDeployedApiResponseBodyResultDescriptionsResultDescription {
	return s.ResultDescription
}

func (s *DescribeDeployedApiResponseBodyResultDescriptions) SetResultDescription(v []*DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) *DescribeDeployedApiResponseBodyResultDescriptions {
	s.ResultDescription = v
	return s
}

func (s *DescribeDeployedApiResponseBodyResultDescriptions) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyResultDescriptionsResultDescription struct {
	// The description.
	//
	// example:
	//
	// result description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether a subnode exists.
	//
	// example:
	//
	// false
	HasChild *bool `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	// The ID of the result.
	//
	// example:
	//
	// id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The primary key of the result.
	//
	// example:
	//
	// DEMO
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Specifies whether the parameter is required.
	//
	// example:
	//
	// true
	Mandatory *bool `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	// The name of the result.
	//
	// example:
	//
	// fwqf
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent node.
	//
	// example:
	//
	// pid
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The type of the result.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetHasChild() *bool {
	return s.HasChild
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetId() *string {
	return s.Id
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetKey() *string {
	return s.Key
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetMandatory() *bool {
	return s.Mandatory
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetName() *string {
	return s.Name
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetPid() *string {
	return s.Pid
}

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) GetType() *string {
	return s.Type
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

func (s *DescribeDeployedApiResponseBodyResultDescriptionsResultDescription) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceConfig struct {
	// Backend configuration items when the backend service is Function Compute
	FunctionComputeConfig *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	// Specifies whether to enable the Mock mode. Valid values:
	//
	// 	- **TRUE: The Mock mode is enabled.**
	//
	// 	- **FALSE: The Mock mode is not enabled.
	//
	// example:
	//
	// TRUE
	Mock *string `json:"Mock,omitempty" xml:"Mock,omitempty"`
	// The simulated Headers.
	MockHeaders *DescribeDeployedApiResponseBodyServiceConfigMockHeaders `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Struct"`
	// The result returned when the Mock mode is enabled.
	//
	// example:
	//
	// test result
	MockResult *string `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	// The status code returned for service mocking.
	//
	// example:
	//
	// 200
	MockStatusCode *int32 `json:"MockStatusCode,omitempty" xml:"MockStatusCode,omitempty"`
	// The URL used to call the back-end service. If the complete back-end service URL is `http://api.a.com:8080/object/add?key1=value1&key2=value2`, the value of ServiceAddress is **http://api.a.com:8080**.``
	//
	// example:
	//
	// http://api.a.com:8080
	ServiceAddress *string `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	// The HTTP method used to call a backend service. Valid values: GET, POST, DELETE, PUT, HEADER, TRACE, PATCH, CONNECT, and OPTIONS.
	//
	// example:
	//
	// POST
	ServiceHttpMethod *string `json:"ServiceHttpMethod,omitempty" xml:"ServiceHttpMethod,omitempty"`
	// example:
	//
	// /object/add
	ServicePath *string `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	// The backend service protocol. Currently, only HTTP, HTTPS, and FunctionCompute are supported.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The timeout period of the backend service, in millisecond.
	//
	// example:
	//
	// 1000
	ServiceTimeout *int32 `json:"ServiceTimeout,omitempty" xml:"ServiceTimeout,omitempty"`
	// Specifies whether to enable the VPC channel. Valid values:
	//
	// 	- **TRUE**: The VPC channel is enabled. You must create the corresponding VPC access authorization before you can enable a VPC channel.
	//
	// 	- **FALSE**: The VPC channel is not enabled.
	//
	// example:
	//
	// TRUE
	ServiceVpcEnable *string `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	// Configuration items related to VPC channels
	VpcConfig *DescribeDeployedApiResponseBodyServiceConfigVpcConfig `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6kg9x8sx2tbxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetFunctionComputeConfig() *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig {
	return s.FunctionComputeConfig
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetMock() *string {
	return s.Mock
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetMockHeaders() *DescribeDeployedApiResponseBodyServiceConfigMockHeaders {
	return s.MockHeaders
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetMockResult() *string {
	return s.MockResult
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetMockStatusCode() *int32 {
	return s.MockStatusCode
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetServiceHttpMethod() *string {
	return s.ServiceHttpMethod
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetServicePath() *string {
	return s.ServicePath
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetServiceTimeout() *int32 {
	return s.ServiceTimeout
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetServiceVpcEnable() *string {
	return s.ServiceVpcEnable
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetVpcConfig() *DescribeDeployedApiResponseBodyServiceConfigVpcConfig {
	return s.VpcConfig
}

func (s *DescribeDeployedApiResponseBodyServiceConfig) GetVpcId() *string {
	return s.VpcId
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

func (s *DescribeDeployedApiResponseBodyServiceConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig struct {
	// The ContentType header type used when you call the backend service over HTTP.
	//
	// 	- **DEFAULT: the default header type in API Gateway.**
	//
	// 	- **CUSTOM: a custom header type.**
	//
	// 	- **CLIENT: the ContentType header type of the client.
	//
	// example:
	//
	// DEFAULT
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	// The value of the ContentType header when the ServiceProtocol parameter is set to HTTP and the ContentTypeCatagory parameter is set to DEFAULT or CUSTOM.
	//
	// example:
	//
	// application/x-www-form-urlencoded; charset=UTF-8
	ContentTypeValue *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	// The root path of Function Compute.
	//
	// example:
	//
	// https://122xxxxxxx.fc.aliyun.com/2016xxxx/proxy/testSxxx.xxx/testHttp/
	FcBaseUrl *string `json:"FcBaseUrl,omitempty" xml:"FcBaseUrl,omitempty"`
	// The type of the Function Compute instance.
	//
	// example:
	//
	// HttpTrigger
	FcType *string `json:"FcType,omitempty" xml:"FcType,omitempty"`
	// The function name defined in Function Compute.
	//
	// example:
	//
	// domain_business_control
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The request method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The backend only receives the service path.
	//
	// example:
	//
	// false
	OnlyBusinessPath *bool `json:"OnlyBusinessPath,omitempty" xml:"OnlyBusinessPath,omitempty"`
	// The API request path.
	//
	// example:
	//
	// /api/offline/cacheData
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The alias of the function.
	//
	// example:
	//
	// 2
	Qualifier *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
	// The region where the API is located.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role to be assumed by API Gateway to access Function Compute.
	//
	// example:
	//
	// acs:ram::111***:role/aliyunserviceroleforsas
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The service name defined in Function Compute.
	//
	// example:
	//
	// fcservicename
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetContentTypeCatagory() *string {
	return s.ContentTypeCatagory
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetContentTypeValue() *string {
	return s.ContentTypeValue
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetFcBaseUrl() *string {
	return s.FcBaseUrl
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetFcType() *string {
	return s.FcType
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetMethod() *string {
	return s.Method
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetOnlyBusinessPath() *bool {
	return s.OnlyBusinessPath
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetPath() *string {
	return s.Path
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetQualifier() *string {
	return s.Qualifier
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) GetServiceName() *string {
	return s.ServiceName
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

func (s *DescribeDeployedApiResponseBodyServiceConfigFunctionComputeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceConfigMockHeaders struct {
	MockHeader []*DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader `json:"MockHeader,omitempty" xml:"MockHeader,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeaders) GetMockHeader() []*DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader {
	return s.MockHeader
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeaders) SetMockHeader(v []*DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) *DescribeDeployedApiResponseBodyServiceConfigMockHeaders {
	s.MockHeader = v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader struct {
	// The name of the HTTP header parameter.
	//
	// example:
	//
	// Content-Type
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The value of the HTTP header parameter.
	//
	// example:
	//
	// 86400
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderName(v string) *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderValue(v string) *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderValue = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceConfigMockHeadersMockHeader) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceConfigVpcConfig struct {
	// The IDs of the ELB and SLB instances in the VPC.
	//
	// example:
	//
	// i-bp1h497hkijewv2***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC access authorization.
	//
	// example:
	//
	// glmall-app-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number that corresponds to the instance.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2zeafsc3fygk1***
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) GetName() *string {
	return s.Name
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
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

func (s *DescribeDeployedApiResponseBodyServiceConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceParameters struct {
	ServiceParameter []*DescribeDeployedApiResponseBodyServiceParametersServiceParameter `json:"ServiceParameter,omitempty" xml:"ServiceParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyServiceParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParameters) GetServiceParameter() []*DescribeDeployedApiResponseBodyServiceParametersServiceParameter {
	return s.ServiceParameter
}

func (s *DescribeDeployedApiResponseBodyServiceParameters) SetServiceParameter(v []*DescribeDeployedApiResponseBodyServiceParametersServiceParameter) *DescribeDeployedApiResponseBodyServiceParameters {
	s.ServiceParameter = v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceParametersServiceParameter struct {
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The data type of the back-end service parameter.
	//
	// example:
	//
	// String
	ParameterType *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	// The name of the backend service parameter.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceParametersServiceParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParametersServiceParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeDeployedApiResponseBodyServiceParametersServiceParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceParametersMap struct {
	ServiceParameterMap []*DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodyServiceParametersMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParametersMap) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMap) GetServiceParameterMap() []*DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap {
	return s.ServiceParameterMap
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMap) SetServiceParameterMap(v []*DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) *DescribeDeployedApiResponseBodyServiceParametersMap {
	s.ServiceParameterMap = v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMap) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap struct {
	// The name of the front-end input parameter.
	//
	// example:
	//
	// sex
	RequestParameterName *string `json:"RequestParameterName,omitempty" xml:"RequestParameterName,omitempty"`
	// The name of the backend service parameter.
	//
	// example:
	//
	// sex
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) GetRequestParameterName() *string {
	return s.RequestParameterName
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) GetServiceParameterName() *string {
	return s.ServiceParameterName
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) SetRequestParameterName(v string) *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) SetServiceParameterName(v string) *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap {
	s.ServiceParameterName = &v
	return s
}

func (s *DescribeDeployedApiResponseBodyServiceParametersMapServiceParameterMap) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodySystemParameters struct {
	SystemParameter []*DescribeDeployedApiResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeDeployedApiResponseBodySystemParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodySystemParameters) GetSystemParameter() []*DescribeDeployedApiResponseBodySystemParametersSystemParameter {
	return s.SystemParameter
}

func (s *DescribeDeployedApiResponseBodySystemParameters) SetSystemParameter(v []*DescribeDeployedApiResponseBodySystemParametersSystemParameter) *DescribeDeployedApiResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

func (s *DescribeDeployedApiResponseBodySystemParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeDeployedApiResponseBodySystemParametersSystemParameter struct {
	// Examples
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The description.
	//
	// example:
	//
	// QueryParamDTO
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The name of the system parameter. Valid values: CaClientIp, CaDomain, CaRequestHandleTime, CaAppId, CaRequestId, CaHttpSchema, and CaProxy.
	//
	// example:
	//
	// CaClientIp
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The name of the corresponding backend parameter.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeDeployedApiResponseBodySystemParametersSystemParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeployedApiResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeDeployedApiResponseBodySystemParametersSystemParameter) Validate() error {
	return dara.Validate(s)
}
