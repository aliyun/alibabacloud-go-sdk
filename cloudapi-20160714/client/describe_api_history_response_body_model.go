// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSignatureMethod(v string) *DescribeApiHistoryResponseBody
	GetAllowSignatureMethod() *string
	SetApiId(v string) *DescribeApiHistoryResponseBody
	GetApiId() *string
	SetApiName(v string) *DescribeApiHistoryResponseBody
	GetApiName() *string
	SetAppCodeAuthType(v string) *DescribeApiHistoryResponseBody
	GetAppCodeAuthType() *string
	SetAuthType(v string) *DescribeApiHistoryResponseBody
	GetAuthType() *string
	SetBackendConfig(v *DescribeApiHistoryResponseBodyBackendConfig) *DescribeApiHistoryResponseBody
	GetBackendConfig() *DescribeApiHistoryResponseBodyBackendConfig
	SetBackendEnable(v bool) *DescribeApiHistoryResponseBody
	GetBackendEnable() *bool
	SetConstantParameters(v *DescribeApiHistoryResponseBodyConstantParameters) *DescribeApiHistoryResponseBody
	GetConstantParameters() *DescribeApiHistoryResponseBodyConstantParameters
	SetCustomSystemParameters(v *DescribeApiHistoryResponseBodyCustomSystemParameters) *DescribeApiHistoryResponseBody
	GetCustomSystemParameters() *DescribeApiHistoryResponseBodyCustomSystemParameters
	SetDeployedTime(v string) *DescribeApiHistoryResponseBody
	GetDeployedTime() *string
	SetDescription(v string) *DescribeApiHistoryResponseBody
	GetDescription() *string
	SetDisableInternet(v bool) *DescribeApiHistoryResponseBody
	GetDisableInternet() *bool
	SetErrorCodeSamples(v *DescribeApiHistoryResponseBodyErrorCodeSamples) *DescribeApiHistoryResponseBody
	GetErrorCodeSamples() *DescribeApiHistoryResponseBodyErrorCodeSamples
	SetFailResultSample(v string) *DescribeApiHistoryResponseBody
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *DescribeApiHistoryResponseBody
	GetForceNonceCheck() *bool
	SetGroupId(v string) *DescribeApiHistoryResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeApiHistoryResponseBody
	GetGroupName() *string
	SetHistoryVersion(v string) *DescribeApiHistoryResponseBody
	GetHistoryVersion() *string
	SetOpenIdConnectConfig(v *DescribeApiHistoryResponseBodyOpenIdConnectConfig) *DescribeApiHistoryResponseBody
	GetOpenIdConnectConfig() *DescribeApiHistoryResponseBodyOpenIdConnectConfig
	SetRegionId(v string) *DescribeApiHistoryResponseBody
	GetRegionId() *string
	SetRequestConfig(v *DescribeApiHistoryResponseBodyRequestConfig) *DescribeApiHistoryResponseBody
	GetRequestConfig() *DescribeApiHistoryResponseBodyRequestConfig
	SetRequestId(v string) *DescribeApiHistoryResponseBody
	GetRequestId() *string
	SetRequestParameters(v *DescribeApiHistoryResponseBodyRequestParameters) *DescribeApiHistoryResponseBody
	GetRequestParameters() *DescribeApiHistoryResponseBodyRequestParameters
	SetResultBodyModel(v string) *DescribeApiHistoryResponseBody
	GetResultBodyModel() *string
	SetResultDescriptions(v *DescribeApiHistoryResponseBodyResultDescriptions) *DescribeApiHistoryResponseBody
	GetResultDescriptions() *DescribeApiHistoryResponseBodyResultDescriptions
	SetResultSample(v string) *DescribeApiHistoryResponseBody
	GetResultSample() *string
	SetResultType(v string) *DescribeApiHistoryResponseBody
	GetResultType() *string
	SetServiceConfig(v *DescribeApiHistoryResponseBodyServiceConfig) *DescribeApiHistoryResponseBody
	GetServiceConfig() *DescribeApiHistoryResponseBodyServiceConfig
	SetServiceParameters(v *DescribeApiHistoryResponseBodyServiceParameters) *DescribeApiHistoryResponseBody
	GetServiceParameters() *DescribeApiHistoryResponseBodyServiceParameters
	SetServiceParametersMap(v *DescribeApiHistoryResponseBodyServiceParametersMap) *DescribeApiHistoryResponseBody
	GetServiceParametersMap() *DescribeApiHistoryResponseBodyServiceParametersMap
	SetStageName(v string) *DescribeApiHistoryResponseBody
	GetStageName() *string
	SetStatus(v string) *DescribeApiHistoryResponseBody
	GetStatus() *string
	SetSystemParameters(v *DescribeApiHistoryResponseBodySystemParameters) *DescribeApiHistoryResponseBody
	GetSystemParameters() *DescribeApiHistoryResponseBodySystemParameters
	SetVisibility(v string) *DescribeApiHistoryResponseBody
	GetVisibility() *string
	SetWebSocketApiType(v string) *DescribeApiHistoryResponseBody
	GetWebSocketApiType() *string
}

type DescribeApiHistoryResponseBody struct {
	// If **AuthType*	- is set to **APP**, this value must be passed to specify the signature algorithm. If you do not specify a value, HmacSHA256 is used by default. Valid values:
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
	// bebf996e4b3d445d83078094b72b0d91
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// Backstage_MengMeng Broadcast_Seven Niu Cloud Push Stream Callback_Official
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The AppCode authentication type supported. Valid values:
	//
	// 	- DEFAULT: supported after being made available in Alibaba Cloud Marketplace
	//
	// 	- DISABLE: not supported.
	//
	// 	- HEADER : supported only in the Header parameter
	//
	// 	- HEADER_QUERY : supported in the Header or Query parameter.
	//
	// example:
	//
	// HEADER
	AppCodeAuthType *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
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
	// Backend configurations
	BackendConfig *DescribeApiHistoryResponseBodyBackendConfig `json:"BackendConfig,omitempty" xml:"BackendConfig,omitempty" type:"Struct"`
	// Specifies whether to enable backend services.
	//
	// example:
	//
	// true
	BackendEnable *bool `json:"BackendEnable,omitempty" xml:"BackendEnable,omitempty"`
	// The constant parameters.
	ConstantParameters *DescribeApiHistoryResponseBodyConstantParameters `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	// The custom system parameters.
	CustomSystemParameters *DescribeApiHistoryResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	// The publishing time (UTC) of the API.
	//
	// example:
	//
	// 2021-06-1113:57:38
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// Queries weather based on the region name
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 	- Specifies whether to set **DisableInternet*	- to **true*	- to limit API calls to within the VPC.
	//
	// 	- If you set **DisableInternet*	- to **false**, the limit is lifted. The default value is false when you create an API.
	//
	// example:
	//
	// true
	DisableInternet *bool `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	// The sample error codes returned by the backend service.
	//
	// For more information, see [ErrorCodeSample](https://help.aliyun.com/document_detail/44392.html).
	ErrorCodeSamples *DescribeApiHistoryResponseBodyErrorCodeSamples `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
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
	// cfb6ef799bf54fffabb0f02019ad2581
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// dev_api
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The historical version number.
	//
	// example:
	//
	// 20211022134156663
	HistoryVersion *string `json:"HistoryVersion,omitempty" xml:"HistoryVersion,omitempty"`
	// The configuration items of the third-party OpenID Connect authentication method.
	OpenIdConnectConfig *DescribeApiHistoryResponseBodyOpenIdConnectConfig `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	// The region where the API is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration items of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestConfig](https://help.aliyun.com/document_detail/43985.html).
	RequestConfig *DescribeApiHistoryResponseBodyRequestConfig `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The parameters of API requests sent by the consumer to API Gateway.
	//
	// For more information, see [RequestParameter](https://help.aliyun.com/document_detail/43986.html).
	RequestParameters *DescribeApiHistoryResponseBodyRequestParameters `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	// The return description of the API.
	//
	// example:
	//
	// {}
	ResultBodyModel *string `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	// The return description of the API.
	ResultDescriptions *DescribeApiHistoryResponseBodyResultDescriptions `json:"ResultDescriptions,omitempty" xml:"ResultDescriptions,omitempty" type:"Struct"`
	// The sample response.
	//
	// example:
	//
	// {\\n  \\"status\\": 0,\\n  \\"data\\": {\\n    \\"count\\": 1,\\n    \\"list\\": [\\n      \\"352\\"\\n    ]\\n  },\\n  \\"message\\": \\"success\\"\\n}
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// The type of the data to return.
	//
	// example:
	//
	// JSON
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The information about a backend service call.
	ServiceConfig *DescribeApiHistoryResponseBodyServiceConfig `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	// The parameters of API requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceParameter](https://help.aliyun.com/document_detail/43988.html).
	ServiceParameters *DescribeApiHistoryResponseBodyServiceParameters `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty" type:"Struct"`
	// The mappings between parameters of requests sent by the consumer to API Gateway and parameters of requests sent by API Gateway to the backend service.
	//
	// For more information, see [ServiceParameterMap](https://help.aliyun.com/document_detail/43989.html).
	ServiceParametersMap *DescribeApiHistoryResponseBodyServiceParametersMap `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty" type:"Struct"`
	// The environment in which the API is requested. Valid values:
	//
	// 	- **RELEASE**: the production environment
	//
	// 	- **PRE**: the pre-release environment
	//
	// 	- **TEST**: the test environment
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// The invocation status of the API.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The common parameters of the APIs, in JSON format.
	SystemParameters *DescribeApiHistoryResponseBodySystemParameters `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	// Specifies whether to make the API public. Valid values:
	//
	// 	- **PUBLIC**: Make the API public. If you set this parameter to PUBLIC, this API is displayed on the APIs page for all users after the API is published to the production environment.
	//
	// 	- **PRIVATE**: Make the API private. Private APIs are not displayed in the Alibaba Cloud Marketplace after the API group to which they belong is made available.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
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
	// COMMON
	WebSocketApiType *string `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s DescribeApiHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBody) GetAllowSignatureMethod() *string {
	return s.AllowSignatureMethod
}

func (s *DescribeApiHistoryResponseBody) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiHistoryResponseBody) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiHistoryResponseBody) GetAppCodeAuthType() *string {
	return s.AppCodeAuthType
}

func (s *DescribeApiHistoryResponseBody) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeApiHistoryResponseBody) GetBackendConfig() *DescribeApiHistoryResponseBodyBackendConfig {
	return s.BackendConfig
}

func (s *DescribeApiHistoryResponseBody) GetBackendEnable() *bool {
	return s.BackendEnable
}

func (s *DescribeApiHistoryResponseBody) GetConstantParameters() *DescribeApiHistoryResponseBodyConstantParameters {
	return s.ConstantParameters
}

func (s *DescribeApiHistoryResponseBody) GetCustomSystemParameters() *DescribeApiHistoryResponseBodyCustomSystemParameters {
	return s.CustomSystemParameters
}

func (s *DescribeApiHistoryResponseBody) GetDeployedTime() *string {
	return s.DeployedTime
}

func (s *DescribeApiHistoryResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBody) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *DescribeApiHistoryResponseBody) GetErrorCodeSamples() *DescribeApiHistoryResponseBodyErrorCodeSamples {
	return s.ErrorCodeSamples
}

func (s *DescribeApiHistoryResponseBody) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *DescribeApiHistoryResponseBody) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *DescribeApiHistoryResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiHistoryResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiHistoryResponseBody) GetHistoryVersion() *string {
	return s.HistoryVersion
}

func (s *DescribeApiHistoryResponseBody) GetOpenIdConnectConfig() *DescribeApiHistoryResponseBodyOpenIdConnectConfig {
	return s.OpenIdConnectConfig
}

func (s *DescribeApiHistoryResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiHistoryResponseBody) GetRequestConfig() *DescribeApiHistoryResponseBodyRequestConfig {
	return s.RequestConfig
}

func (s *DescribeApiHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiHistoryResponseBody) GetRequestParameters() *DescribeApiHistoryResponseBodyRequestParameters {
	return s.RequestParameters
}

func (s *DescribeApiHistoryResponseBody) GetResultBodyModel() *string {
	return s.ResultBodyModel
}

func (s *DescribeApiHistoryResponseBody) GetResultDescriptions() *DescribeApiHistoryResponseBodyResultDescriptions {
	return s.ResultDescriptions
}

func (s *DescribeApiHistoryResponseBody) GetResultSample() *string {
	return s.ResultSample
}

func (s *DescribeApiHistoryResponseBody) GetResultType() *string {
	return s.ResultType
}

func (s *DescribeApiHistoryResponseBody) GetServiceConfig() *DescribeApiHistoryResponseBodyServiceConfig {
	return s.ServiceConfig
}

func (s *DescribeApiHistoryResponseBody) GetServiceParameters() *DescribeApiHistoryResponseBodyServiceParameters {
	return s.ServiceParameters
}

func (s *DescribeApiHistoryResponseBody) GetServiceParametersMap() *DescribeApiHistoryResponseBodyServiceParametersMap {
	return s.ServiceParametersMap
}

func (s *DescribeApiHistoryResponseBody) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiHistoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeApiHistoryResponseBody) GetSystemParameters() *DescribeApiHistoryResponseBodySystemParameters {
	return s.SystemParameters
}

func (s *DescribeApiHistoryResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApiHistoryResponseBody) GetWebSocketApiType() *string {
	return s.WebSocketApiType
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

func (s *DescribeApiHistoryResponseBody) SetBackendConfig(v *DescribeApiHistoryResponseBodyBackendConfig) *DescribeApiHistoryResponseBody {
	s.BackendConfig = v
	return s
}

func (s *DescribeApiHistoryResponseBody) SetBackendEnable(v bool) *DescribeApiHistoryResponseBody {
	s.BackendEnable = &v
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

func (s *DescribeApiHistoryResponseBody) Validate() error {
	if s.BackendConfig != nil {
		if err := s.BackendConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ConstantParameters != nil {
		if err := s.ConstantParameters.Validate(); err != nil {
			return err
		}
	}
	if s.CustomSystemParameters != nil {
		if err := s.CustomSystemParameters.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorCodeSamples != nil {
		if err := s.ErrorCodeSamples.Validate(); err != nil {
			return err
		}
	}
	if s.OpenIdConnectConfig != nil {
		if err := s.OpenIdConnectConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RequestConfig != nil {
		if err := s.RequestConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RequestParameters != nil {
		if err := s.RequestParameters.Validate(); err != nil {
			return err
		}
	}
	if s.ResultDescriptions != nil {
		if err := s.ResultDescriptions.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceConfig != nil {
		if err := s.ServiceConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceParameters != nil {
		if err := s.ServiceParameters.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceParametersMap != nil {
		if err := s.ServiceParametersMap.Validate(); err != nil {
			return err
		}
	}
	if s.SystemParameters != nil {
		if err := s.SystemParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyBackendConfig struct {
	// The ID of the backend service.
	//
	// example:
	//
	// a0305308908c4740aba9cbfd63ba99b7
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The name of the backend service.
	//
	// example:
	//
	// zmapi
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// The type of the backend service.
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
}

func (s DescribeApiHistoryResponseBodyBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyBackendConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) GetBackendName() *string {
	return s.BackendName
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) GetBackendType() *string {
	return s.BackendType
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) SetBackendId(v string) *DescribeApiHistoryResponseBodyBackendConfig {
	s.BackendId = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) SetBackendName(v string) *DescribeApiHistoryResponseBodyBackendConfig {
	s.BackendName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) SetBackendType(v string) *DescribeApiHistoryResponseBodyBackendConfig {
	s.BackendType = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyBackendConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeApiHistoryResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyConstantParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyConstantParameters) GetConstantParameter() []*DescribeApiHistoryResponseBodyConstantParametersConstantParameter {
	return s.ConstantParameter
}

func (s *DescribeApiHistoryResponseBodyConstantParameters) SetConstantParameter(v []*DescribeApiHistoryResponseBodyConstantParametersConstantParameter) *DescribeApiHistoryResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

func (s *DescribeApiHistoryResponseBodyConstantParameters) Validate() error {
	if s.ConstantParameter != nil {
		for _, item := range s.ConstantParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyConstantParametersConstantParameter struct {
	// The value of the constant parameter.
	//
	// example:
	//
	// constance
	ConstantValue *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// for_test1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The mapped parameter name in the backend service.
	//
	// example:
	//
	// constance
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyConstantParametersConstantParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) GetConstantValue() *string {
	return s.ConstantValue
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiHistoryResponseBodyConstantParametersConstantParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyCustomSystemParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParameters) GetCustomSystemParameter() []*DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter {
	return s.CustomSystemParameter
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeApiHistoryResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParameters) Validate() error {
	if s.CustomSystemParameter != nil {
		for _, item := range s.CustomSystemParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter struct {
	// The sample value.
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// balabala
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// CaClientIp
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The mapped parameter name in the backend service.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiHistoryResponseBodyCustomSystemParametersCustomSystemParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamples) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamples) GetErrorCodeSample() []*DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample {
	return s.ErrorCodeSample
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeApiHistoryResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamples) Validate() error {
	if s.ErrorCodeSample != nil {
		for _, item := range s.ErrorCodeSample {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample struct {
	// The returned error code.
	//
	// example:
	//
	// 400
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error description.
	//
	// example:
	//
	// Missing the parameter UserId
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// MissingParameter
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) GetCode() *string {
	return s.Code
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) GetMessage() *string {
	return s.Message
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

func (s *DescribeApiHistoryResponseBodyErrorCodeSamplesErrorCodeSample) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyOpenIdConnectConfig struct {
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

func (s DescribeApiHistoryResponseBodyOpenIdConnectConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) GetIdTokenParamName() *string {
	return s.IdTokenParamName
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) GetOpenIdApiType() *string {
	return s.OpenIdApiType
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) GetPublicKey() *string {
	return s.PublicKey
}

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) GetPublicKeyId() *string {
	return s.PublicKeyId
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

func (s *DescribeApiHistoryResponseBodyOpenIdConnectConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyRequestConfig struct {
	// The server data transmission method used for POST and PUT requests. Valid values: FORM and STREAM. FORM indicates that data in key-value pairs is transmitted as forms. STREAM indicates that data is transmitted as byte streams. This parameter takes effect only when the RequestMode parameter is set to MAPPING.
	//
	// example:
	//
	// STREAM
	BodyFormat *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
	// The body model.
	//
	// example:
	//
	// https://apigateway.aliyun.com/models/3a240a1XXXXXXXXd9ab1bf7e947b4095/9e2df550e85b4XXXXXXXX619eaab
	BodyModel *string `json:"BodyModel,omitempty" xml:"BodyModel,omitempty"`
	// Whether to escape the Path parameter, if true, the [param] on the Path will be treated as a regular character.
	//
	// example:
	//
	// true
	EscapePathParam *bool `json:"EscapePathParam,omitempty" xml:"EscapePathParam,omitempty"`
	// The description of the request body.
	//
	// example:
	//
	// fwefwef
	PostBodyDescription *string `json:"PostBodyDescription,omitempty" xml:"PostBodyDescription,omitempty"`
	// The HTTP method. Valid values: GET, POST, DELETE, PUT, HEADER, TRACE, PATCH, CONNECT, and OPTIONS.
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
	// API path
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

func (s DescribeApiHistoryResponseBodyRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetBodyFormat() *string {
	return s.BodyFormat
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetBodyModel() *string {
	return s.BodyModel
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetEscapePathParam() *bool {
	return s.EscapePathParam
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetPostBodyDescription() *string {
	return s.PostBodyDescription
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetRequestHttpMethod() *string {
	return s.RequestHttpMethod
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetRequestMode() *string {
	return s.RequestMode
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetRequestPath() *string {
	return s.RequestPath
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) GetRequestProtocol() *string {
	return s.RequestProtocol
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetBodyModel(v string) *DescribeApiHistoryResponseBodyRequestConfig {
	s.BodyModel = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestConfig) SetEscapePathParam(v bool) *DescribeApiHistoryResponseBodyRequestConfig {
	s.EscapePathParam = &v
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

func (s *DescribeApiHistoryResponseBodyRequestConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyRequestParameters struct {
	RequestParameter []*DescribeApiHistoryResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyRequestParameters) GetRequestParameter() []*DescribeApiHistoryResponseBodyRequestParametersRequestParameter {
	return s.RequestParameter
}

func (s *DescribeApiHistoryResponseBodyRequestParameters) SetRequestParameter(v []*DescribeApiHistoryResponseBodyRequestParametersRequestParameter) *DescribeApiHistoryResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

func (s *DescribeApiHistoryResponseBodyRequestParameters) Validate() error {
	if s.RequestParameter != nil {
		for _, item := range s.RequestParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyRequestParametersRequestParameter struct {
	// The name of the parameter in the API request.
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
	// The sample value.
	//
	// example:
	//
	// 20
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// modidyTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The order in which the parameter is sorted in the document.
	//
	// example:
	//
	// 0
	DocOrder *int32 `json:"DocOrder,omitempty" xml:"DocOrder,omitempty"`
	// Indicates whether the document is public. Valid values: **PUBLIC*	- and **PRIVATE**.
	//
	// example:
	//
	// PUBLIC
	DocShow *string `json:"DocShow,omitempty" xml:"DocShow,omitempty"`
	// The hash values that are supported when **ParameterType*	- is set to Int, Long, Float, Double, or String. Separate values with commas (,). Examples: 1,2,3,4,9 and A,B,C,E,F.
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
	// The maximum parameter length when **ParameterType*	- is set to String.
	//
	// example:
	//
	// 123456
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The maximum parameter value when **ParameterType*	- is set to Int, Long, Float, or Double.
	//
	// example:
	//
	// 123456
	MaxValue *int64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum parameter length when **ParameterType*	- is set to String.
	//
	// example:
	//
	// 123456
	MinLength *int64 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// The minimum parameter value when **ParameterType*	- is set to Int, Long, Float, or Double.
	//
	// example:
	//
	// 123456
	MinValue *int64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The data type of the parameter. Valid values: String, Int, Long, Float, Double, and Boolean.
	//
	// example:
	//
	// String
	ParameterType *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	// The regular expression that is used for parameter validation when **ParameterType*	- is set to String.
	//
	// example:
	//
	// xxx
	RegularExpression *string `json:"RegularExpression,omitempty" xml:"RegularExpression,omitempty"`
	// Indicates whether the parameter is required. Valid values: **REQUIRED*	- and **OPTIONAL**.
	//
	// example:
	//
	// OPTIONAL
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeApiHistoryResponseBodyRequestParametersRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetApiParameterName() *string {
	return s.ApiParameterName
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetArrayItemsType() *string {
	return s.ArrayItemsType
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetDocOrder() *int32 {
	return s.DocOrder
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetDocShow() *string {
	return s.DocShow
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetJsonScheme() *string {
	return s.JsonScheme
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetMinLength() *int64 {
	return s.MinLength
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetMinValue() *int64 {
	return s.MinValue
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetRegularExpression() *string {
	return s.RegularExpression
}

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) GetRequired() *string {
	return s.Required
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

func (s *DescribeApiHistoryResponseBodyRequestParametersRequestParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyResultDescriptions struct {
	ResultDescription []*DescribeApiHistoryResponseBodyResultDescriptionsResultDescription `json:"ResultDescription,omitempty" xml:"ResultDescription,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyResultDescriptions) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyResultDescriptions) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyResultDescriptions) GetResultDescription() []*DescribeApiHistoryResponseBodyResultDescriptionsResultDescription {
	return s.ResultDescription
}

func (s *DescribeApiHistoryResponseBodyResultDescriptions) SetResultDescription(v []*DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) *DescribeApiHistoryResponseBodyResultDescriptions {
	s.ResultDescription = v
	return s
}

func (s *DescribeApiHistoryResponseBodyResultDescriptions) Validate() error {
	if s.ResultDescription != nil {
		for _, item := range s.ResultDescription {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyResultDescriptionsResultDescription struct {
	// The subnode description.
	//
	// example:
	//
	// for_test1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether a subnode exists.
	//
	// example:
	//
	// true
	HasChild *bool `json:"HasChild,omitempty" xml:"HasChild,omitempty"`
	// The result ID.
	//
	// example:
	//
	// id
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The primary key of the result.
	//
	// example:
	//
	// groupName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
	Mandatory *bool `json:"Mandatory,omitempty" xml:"Mandatory,omitempty"`
	// The result name.
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
	// The result type.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetHasChild() *bool {
	return s.HasChild
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetId() *string {
	return s.Id
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetKey() *string {
	return s.Key
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetMandatory() *bool {
	return s.Mandatory
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetName() *string {
	return s.Name
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetPid() *string {
	return s.Pid
}

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) GetType() *string {
	return s.Type
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

func (s *DescribeApiHistoryResponseBodyResultDescriptionsResultDescription) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceConfig struct {
	// The ContentType header type used when you call the backend service over HTTP.
	//
	// 	- DEFAULT: the default header type in API Gateway
	//
	// 	- CUSTOM: a custom header type
	//
	// 	- CLIENT: the ContentType header type of the client
	//
	// example:
	//
	// CUSTOM
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	// The value of the ContentType header when the ServiceProtocol parameter is set to HTTP and the ContentTypeCatagory parameter is set to DEFAULT or CUSTOM.
	//
	// example:
	//
	// application/json
	ContentTypeValue *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	// Configuration items of EventBridge
	EventBridgeConfig *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig `json:"EventBridgeConfig,omitempty" xml:"EventBridgeConfig,omitempty" type:"Struct"`
	// Backend configuration items when the backend service is Function Compute
	FunctionComputeConfig *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	// Specifies whether to enable the MOCK mode. Valid values:
	//
	// 	- TRUE: The Mock mode is enabled.
	//
	// 	- FALSE: The Mock mode is not enabled.
	//
	// example:
	//
	// TRUE
	Mock *string `json:"Mock,omitempty" xml:"Mock,omitempty"`
	// The simulated Headers.
	MockHeaders *DescribeApiHistoryResponseBodyServiceConfigMockHeaders `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Struct"`
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
	// Information when the backend service is OSS
	OssConfig *DescribeApiHistoryResponseBodyServiceConfigOssConfig `json:"OssConfig,omitempty" xml:"OssConfig,omitempty" type:"Struct"`
	// The URL used to call the backend service.
	//
	// example:
	//
	// http://api.a.com:8080
	ServiceAddress *string `json:"ServiceAddress,omitempty" xml:"ServiceAddress,omitempty"`
	// The HTTP request method used when calling the backend service. Valid values: PUT, GET, POST, DELETE, PATCH, HEAD, OPTIONS, and ANY.
	//
	// example:
	//
	// POST
	ServiceHttpMethod *string `json:"ServiceHttpMethod,omitempty" xml:"ServiceHttpMethod,omitempty"`
	// The path used when you call the backend service.
	//
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
	// 	- TRUE: The VPC channel is enabled.
	//
	// 	- FALSE: The VPC channel is not enabled.
	//
	// You must create the corresponding VPC access authorization before you can enable a VPC channel.
	//
	// example:
	//
	// TRUE
	ServiceVpcEnable *string `json:"ServiceVpcEnable,omitempty" xml:"ServiceVpcEnable,omitempty"`
	// Configuration items related to VPC channels
	VpcConfig *DescribeApiHistoryResponseBodyServiceConfigVpcConfig `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-uf6kg9x8sx2tbxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetContentTypeCatagory() *string {
	return s.ContentTypeCatagory
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetContentTypeValue() *string {
	return s.ContentTypeValue
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetEventBridgeConfig() *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig {
	return s.EventBridgeConfig
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetFunctionComputeConfig() *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig {
	return s.FunctionComputeConfig
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetMock() *string {
	return s.Mock
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetMockHeaders() *DescribeApiHistoryResponseBodyServiceConfigMockHeaders {
	return s.MockHeaders
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetMockResult() *string {
	return s.MockResult
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetMockStatusCode() *int32 {
	return s.MockStatusCode
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetOssConfig() *DescribeApiHistoryResponseBodyServiceConfigOssConfig {
	return s.OssConfig
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetServiceHttpMethod() *string {
	return s.ServiceHttpMethod
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetServicePath() *string {
	return s.ServicePath
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetServiceTimeout() *int32 {
	return s.ServiceTimeout
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetServiceVpcEnable() *string {
	return s.ServiceVpcEnable
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetVpcConfig() *DescribeApiHistoryResponseBodyServiceConfigVpcConfig {
	return s.VpcConfig
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetContentTypeCatagory(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ContentTypeCatagory = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetContentTypeValue(v string) *DescribeApiHistoryResponseBodyServiceConfig {
	s.ContentTypeValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfig) SetEventBridgeConfig(v *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) *DescribeApiHistoryResponseBodyServiceConfig {
	s.EventBridgeConfig = v
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

func (s *DescribeApiHistoryResponseBodyServiceConfig) Validate() error {
	if s.EventBridgeConfig != nil {
		if err := s.EventBridgeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.FunctionComputeConfig != nil {
		if err := s.FunctionComputeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.MockHeaders != nil {
		if err := s.MockHeaders.Validate(); err != nil {
			return err
		}
	}
	if s.OssConfig != nil {
		if err := s.OssConfig.Validate(); err != nil {
			return err
		}
	}
	if s.VpcConfig != nil {
		if err := s.VpcConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig struct {
	// The ID of the region where the EventBridge instance is located.
	//
	// example:
	//
	// cn-beijing
	EventBridgeRegionId *string `json:"EventBridgeRegionId,omitempty" xml:"EventBridgeRegionId,omitempty"`
	// The event bus.
	//
	// example:
	//
	// testBus
	EventBus *string `json:"EventBus,omitempty" xml:"EventBus,omitempty"`
	// The event source of the managed rule.
	//
	// example:
	//
	// baas_driver
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The Arn that is authorized by a RAM user to EventBridge.
	//
	// example:
	//
	// acs:ram::1933122015759***:role/adminoidcaliyun
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) GetEventBridgeRegionId() *string {
	return s.EventBridgeRegionId
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) GetEventBus() *string {
	return s.EventBus
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) GetEventSource() *string {
	return s.EventSource
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) SetEventBridgeRegionId(v string) *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig {
	s.EventBridgeRegionId = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) SetEventBus(v string) *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig {
	s.EventBus = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) SetEventSource(v string) *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig {
	s.EventSource = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) SetRoleArn(v string) *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigEventBridgeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig struct {
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
	// application/json
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
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
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

func (s DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetContentTypeCatagory() *string {
	return s.ContentTypeCatagory
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetContentTypeValue() *string {
	return s.ContentTypeValue
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetFcBaseUrl() *string {
	return s.FcBaseUrl
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetFcType() *string {
	return s.FcType
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetMethod() *string {
	return s.Method
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetOnlyBusinessPath() *bool {
	return s.OnlyBusinessPath
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetPath() *string {
	return s.Path
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetQualifier() *string {
	return s.Qualifier
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) GetServiceName() *string {
	return s.ServiceName
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

func (s *DescribeApiHistoryResponseBodyServiceConfigFunctionComputeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceConfigMockHeaders struct {
	MockHeader []*DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader `json:"MockHeader,omitempty" xml:"MockHeader,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeaders) GetMockHeader() []*DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader {
	return s.MockHeader
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeaders) SetMockHeader(v []*DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) *DescribeApiHistoryResponseBodyServiceConfigMockHeaders {
	s.MockHeader = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeaders) Validate() error {
	if s.MockHeader != nil {
		for _, item := range s.MockHeader {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader struct {
	// The HTTP headers.
	//
	// example:
	//
	// Content-Type
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The values of the HTTP headers.
	//
	// example:
	//
	// 86400
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderName(v string) *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderValue(v string) *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderValue = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceConfigMockHeadersMockHeader) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceConfigOssConfig struct {
	// The operation options on OSS. Valid values:
	//
	// 	- GetObject
	//
	// 	- PostObject
	//
	// 	- DeleteObject
	//
	// 	- PutObject
	//
	// 	- HeadObject
	//
	// 	- GetObjectMeta
	//
	// 	- AppendObject
	//
	// example:
	//
	// GetObject
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The OSS bucket.
	//
	// example:
	//
	// phototest02
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The stored object or folder path.
	//
	// example:
	//
	// ENV
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the region where the OSS instance is located.
	//
	// example:
	//
	// cn-hangzhou
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigOssConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigOssConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) GetAction() *string {
	return s.Action
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) GetKey() *string {
	return s.Key
}

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) GetOssRegionId() *string {
	return s.OssRegionId
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

func (s *DescribeApiHistoryResponseBodyServiceConfigOssConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceConfigVpcConfig struct {
	// The IDs of the ELB and SLB instances in the VPC.
	//
	// example:
	//
	// i-bp1h497hkijewv2***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC.
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
	// The VPC protocol.
	//
	// example:
	//
	// HTTP
	VpcScheme *string `json:"VpcScheme,omitempty" xml:"VpcScheme,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GetName() *string {
	return s.Name
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GetPort() *int32 {
	return s.Port
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) GetVpcScheme() *string {
	return s.VpcScheme
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

func (s *DescribeApiHistoryResponseBodyServiceConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceParameters struct {
	ServiceParameter []*DescribeApiHistoryResponseBodyServiceParametersServiceParameter `json:"ServiceParameter,omitempty" xml:"ServiceParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyServiceParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParameters) GetServiceParameter() []*DescribeApiHistoryResponseBodyServiceParametersServiceParameter {
	return s.ServiceParameter
}

func (s *DescribeApiHistoryResponseBodyServiceParameters) SetServiceParameter(v []*DescribeApiHistoryResponseBodyServiceParametersServiceParameter) *DescribeApiHistoryResponseBodyServiceParameters {
	s.ServiceParameter = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParameters) Validate() error {
	if s.ServiceParameter != nil {
		for _, item := range s.ServiceParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyServiceParametersServiceParameter struct {
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The data type of the parameter. Valid values: STRING, NUMBER, and BOOLEAN.
	//
	// example:
	//
	// String
	ParameterType *string `json:"ParameterType,omitempty" xml:"ParameterType,omitempty"`
	// The mapped parameter name in the backend service.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceParametersServiceParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParametersServiceParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiHistoryResponseBodyServiceParametersServiceParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodyServiceParametersMap struct {
	ServiceParameterMap []*DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodyServiceParametersMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParametersMap) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMap) GetServiceParameterMap() []*DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap {
	return s.ServiceParameterMap
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMap) SetServiceParameterMap(v []*DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) *DescribeApiHistoryResponseBodyServiceParametersMap {
	s.ServiceParameterMap = v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMap) Validate() error {
	if s.ServiceParameterMap != nil {
		for _, item := range s.ServiceParameterMap {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap struct {
	// The corresponding frontend parameter name. The value must be contained in RequestParametersObject and match RequestParam.ApiParameterName.
	//
	// example:
	//
	// sex
	RequestParameterName *string `json:"RequestParameterName,omitempty" xml:"RequestParameterName,omitempty"`
	// The mapped parameter name in the backend service.
	//
	// example:
	//
	// sex
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) GetRequestParameterName() *string {
	return s.RequestParameterName
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) GetServiceParameterName() *string {
	return s.ServiceParameterName
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) SetRequestParameterName(v string) *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) SetServiceParameterName(v string) *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap {
	s.ServiceParameterName = &v
	return s
}

func (s *DescribeApiHistoryResponseBodyServiceParametersMapServiceParameterMap) Validate() error {
	return dara.Validate(s)
}

type DescribeApiHistoryResponseBodySystemParameters struct {
	SystemParameter []*DescribeApiHistoryResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiHistoryResponseBodySystemParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodySystemParameters) GetSystemParameter() []*DescribeApiHistoryResponseBodySystemParametersSystemParameter {
	return s.SystemParameter
}

func (s *DescribeApiHistoryResponseBodySystemParameters) SetSystemParameter(v []*DescribeApiHistoryResponseBodySystemParametersSystemParameter) *DescribeApiHistoryResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

func (s *DescribeApiHistoryResponseBodySystemParameters) Validate() error {
	if s.SystemParameter != nil {
		for _, item := range s.SystemParameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiHistoryResponseBodySystemParametersSystemParameter struct {
	// The sample value.
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The description.
	//
	// example:
	//
	// system parameters description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter location. Valid values: BODY, HEAD, QUERY, and PATH.
	//
	// example:
	//
	// HEAD
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The system parameter. Valid values: CaClientIp, CaDomain, CaRequestHandleTime, CaAppId, CaRequestId, CaHttpSchema, and CaProxy.
	//
	// example:
	//
	// CaClientIp
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The mapped parameter name in the backend service.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiHistoryResponseBodySystemParametersSystemParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoryResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiHistoryResponseBodySystemParametersSystemParameter) Validate() error {
	return dara.Validate(s)
}
