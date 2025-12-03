// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowSignatureMethod(v string) *DescribeApiResponseBody
	GetAllowSignatureMethod() *string
	SetApiId(v string) *DescribeApiResponseBody
	GetApiId() *string
	SetApiName(v string) *DescribeApiResponseBody
	GetApiName() *string
	SetAppCodeAuthType(v string) *DescribeApiResponseBody
	GetAppCodeAuthType() *string
	SetAuthType(v string) *DescribeApiResponseBody
	GetAuthType() *string
	SetBackendConfig(v *DescribeApiResponseBodyBackendConfig) *DescribeApiResponseBody
	GetBackendConfig() *DescribeApiResponseBodyBackendConfig
	SetBackendEnable(v bool) *DescribeApiResponseBody
	GetBackendEnable() *bool
	SetConstantParameters(v *DescribeApiResponseBodyConstantParameters) *DescribeApiResponseBody
	GetConstantParameters() *DescribeApiResponseBodyConstantParameters
	SetCreatedTime(v string) *DescribeApiResponseBody
	GetCreatedTime() *string
	SetCustomSystemParameters(v *DescribeApiResponseBodyCustomSystemParameters) *DescribeApiResponseBody
	GetCustomSystemParameters() *DescribeApiResponseBodyCustomSystemParameters
	SetDeployedInfos(v *DescribeApiResponseBodyDeployedInfos) *DescribeApiResponseBody
	GetDeployedInfos() *DescribeApiResponseBodyDeployedInfos
	SetDescription(v string) *DescribeApiResponseBody
	GetDescription() *string
	SetDisableInternet(v bool) *DescribeApiResponseBody
	GetDisableInternet() *bool
	SetErrorCodeSamples(v *DescribeApiResponseBodyErrorCodeSamples) *DescribeApiResponseBody
	GetErrorCodeSamples() *DescribeApiResponseBodyErrorCodeSamples
	SetFailResultSample(v string) *DescribeApiResponseBody
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *DescribeApiResponseBody
	GetForceNonceCheck() *bool
	SetGroupId(v string) *DescribeApiResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeApiResponseBody
	GetGroupName() *string
	SetMock(v string) *DescribeApiResponseBody
	GetMock() *string
	SetMockResult(v string) *DescribeApiResponseBody
	GetMockResult() *string
	SetModifiedTime(v string) *DescribeApiResponseBody
	GetModifiedTime() *string
	SetOpenIdConnectConfig(v *DescribeApiResponseBodyOpenIdConnectConfig) *DescribeApiResponseBody
	GetOpenIdConnectConfig() *DescribeApiResponseBodyOpenIdConnectConfig
	SetRegionId(v string) *DescribeApiResponseBody
	GetRegionId() *string
	SetRequestConfig(v *DescribeApiResponseBodyRequestConfig) *DescribeApiResponseBody
	GetRequestConfig() *DescribeApiResponseBodyRequestConfig
	SetRequestId(v string) *DescribeApiResponseBody
	GetRequestId() *string
	SetRequestParameters(v *DescribeApiResponseBodyRequestParameters) *DescribeApiResponseBody
	GetRequestParameters() *DescribeApiResponseBodyRequestParameters
	SetResultBodyModel(v string) *DescribeApiResponseBody
	GetResultBodyModel() *string
	SetResultSample(v string) *DescribeApiResponseBody
	GetResultSample() *string
	SetResultType(v string) *DescribeApiResponseBody
	GetResultType() *string
	SetServiceConfig(v *DescribeApiResponseBodyServiceConfig) *DescribeApiResponseBody
	GetServiceConfig() *DescribeApiResponseBodyServiceConfig
	SetServiceParameters(v *DescribeApiResponseBodyServiceParameters) *DescribeApiResponseBody
	GetServiceParameters() *DescribeApiResponseBodyServiceParameters
	SetServiceParametersMap(v *DescribeApiResponseBodyServiceParametersMap) *DescribeApiResponseBody
	GetServiceParametersMap() *DescribeApiResponseBodyServiceParametersMap
	SetSystemParameters(v *DescribeApiResponseBodySystemParameters) *DescribeApiResponseBody
	GetSystemParameters() *DescribeApiResponseBodySystemParameters
	SetTagList(v *DescribeApiResponseBodyTagList) *DescribeApiResponseBody
	GetTagList() *DescribeApiResponseBodyTagList
	SetVisibility(v string) *DescribeApiResponseBody
	GetVisibility() *string
	SetWebSocketApiType(v string) *DescribeApiResponseBody
	GetWebSocketApiType() *string
}

type DescribeApiResponseBody struct {
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
	// 8afff6c8c4c6447abb035812e4d66b65
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API, which is unique in the group.
	//
	// example:
	//
	// ApiName
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
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
	// HEADER
	AppCodeAuthType *string `json:"AppCodeAuthType,omitempty" xml:"AppCodeAuthType,omitempty"`
	// The security authentication method of the API. Valid values:
	//
	// 	- **APP**: Only authorized applications can call the API.
	//
	// 	- **ANONYMOUS**: The API can be anonymously called. In this mode, you must take note of the following rules:
	//
	//     	- All users who have obtained the API service information can call this API. API Gateway does not authenticate callers and cannot set user-specific throttling policies. If you make this API public, set API-specific throttling policies.
	//
	//     	- We recommend that you do not make the API whose security authentication method is ANONYMOUS available in Alibaba Cloud Marketplace because API Gateway cannot meter calls on the caller or limit the number of calls on the API. If you want to make the API group to which the API belongs available in Alibaba Cloud Marketplace, we recommend that you move the API to another group, set its type to PRIVATE, or set its security authentication method to APP.
	//
	// 	- **APPOPENID**: The OpenID Connect account authentication method is used. Only applications authorized by OpenID Connect can call the API. If this method is selected, the OpenIdConnectConfig parameter is required.
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Backend configurations
	BackendConfig *DescribeApiResponseBodyBackendConfig `json:"BackendConfig,omitempty" xml:"BackendConfig,omitempty" type:"Struct"`
	// Specifies whether to enable backend services.
	//
	// example:
	//
	// true
	BackendEnable *bool `json:"BackendEnable,omitempty" xml:"BackendEnable,omitempty"`
	// System parameters sent by API Gateway to the backend service
	ConstantParameters *DescribeApiResponseBodyConstantParameters `json:"ConstantParameters,omitempty" xml:"ConstantParameters,omitempty" type:"Struct"`
	// The creation time of the API.
	//
	// example:
	//
	// 2016-07-28T09:50:43Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Custom system parameters
	CustomSystemParameters *DescribeApiResponseBodyCustomSystemParameters `json:"CustomSystemParameters,omitempty" xml:"CustomSystemParameters,omitempty" type:"Struct"`
	// The API publishing status.
	DeployedInfos *DescribeApiResponseBodyDeployedInfos `json:"DeployedInfos,omitempty" xml:"DeployedInfos,omitempty" type:"Struct"`
	// The description of the API.
	//
	// example:
	//
	// Api description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to limit API calls to within the VPC. Valid values:
	//
	// 	- **true**: Only API calls from the VPC are supported.
	//
	// 	- **false**: API calls from the VPC and Internet are both supported.
	//
	// example:
	//
	// false
	DisableInternet *bool `json:"DisableInternet,omitempty" xml:"DisableInternet,omitempty"`
	// The sample error codes returned by the backend service.
	ErrorCodeSamples *DescribeApiResponseBodyErrorCodeSamples `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
	// The sample error response from the backend service.
	//
	// example:
	//
	// 400
	FailResultSample *string `json:"FailResultSample,omitempty" xml:"FailResultSample,omitempty"`
	// Specifies whether to carry the header : X-Ca-Nonce when calling an API. This is the unique identifier of the request and is generally identified by UUID. After receiving this parameter, API Gateway verifies the validity of this parameter. The same value can be used only once within 15 minutes. This helps prevent reply attacks. Valid values:
	//
	// 	- **true**: This field is forcibly checked when an API is requested to prevent replay attacks.
	//
	// 	- **false**: This field is not checked.
	//
	// example:
	//
	// true
	ForceNonceCheck *bool `json:"ForceNonceCheck,omitempty" xml:"ForceNonceCheck,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 08ae4aa0f95e4321849ee57f4e0b3077
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// ApiTest
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Specifies whether to enable the Mock mode. Valid values:
	//
	// 	- OPEN: The Mock mode is enabled.
	//
	// 	- CLOSED: The Mock mode is not enabled.
	//
	// example:
	//
	// CLOSED
	Mock *string `json:"Mock,omitempty" xml:"Mock,omitempty"`
	// The result returned for service mocking.
	//
	// example:
	//
	// test result
	MockResult *string `json:"MockResult,omitempty" xml:"MockResult,omitempty"`
	// The last modification time of the API.
	//
	// example:
	//
	// 2016-07-28T13:13:12Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Configuration items of the third-party OpenID Connect authentication method
	OpenIdConnectConfig *DescribeApiResponseBodyOpenIdConnectConfig `json:"OpenIdConnectConfig,omitempty" xml:"OpenIdConnectConfig,omitempty" type:"Struct"`
	// The region ID of the API.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration items of API requests sent by the consumer to API Gateway.
	RequestConfig *DescribeApiResponseBodyRequestConfig `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// D0FF585F-7966-40CF-BC60-75DB070B23D5<
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The parameters of API requests sent by the consumer to API Gateway.
	RequestParameters *DescribeApiResponseBodyRequestParameters `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	// The returned description of the API.
	//
	// example:
	//
	// {}
	ResultBodyModel *string `json:"ResultBodyModel,omitempty" xml:"ResultBodyModel,omitempty"`
	// The sample response from the backend service.
	//
	// example:
	//
	// 200
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// The format of the response from the backend service. Valid values: JSON, TEXT, BINARY, XML, and HTML.
	//
	// example:
	//
	// JSON
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The configuration items of API requests that API Gateway sends to the backend service.
	ServiceConfig *DescribeApiResponseBodyServiceConfig `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty" type:"Struct"`
	// The parameters of API requests sent by API Gateway to the backend service.
	ServiceParameters *DescribeApiResponseBodyServiceParameters `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty" type:"Struct"`
	// The mappings between parameters of requests sent by the consumer to API Gateway and parameters of requests sent by API Gateway to the backend service.
	ServiceParametersMap *DescribeApiResponseBodyServiceParametersMap `json:"ServiceParametersMap,omitempty" xml:"ServiceParametersMap,omitempty" type:"Struct"`
	// System parameters sent by API Gateway to the backend service
	SystemParameters *DescribeApiResponseBodySystemParameters `json:"SystemParameters,omitempty" xml:"SystemParameters,omitempty" type:"Struct"`
	// Tag List.
	TagList *DescribeApiResponseBodyTagList `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Struct"`
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
	// The type of the two-way communication API.
	//
	// 	- **COMMON**: common API
	//
	// 	- **REGISTER**: registered API
	//
	// 	- **UNREGISTER**: unregistered API
	//
	// 	- **NOTIFY**: downstream notification API
	//
	// example:
	//
	// COMMON
	WebSocketApiType *string `json:"WebSocketApiType,omitempty" xml:"WebSocketApiType,omitempty"`
}

func (s DescribeApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBody) GetAllowSignatureMethod() *string {
	return s.AllowSignatureMethod
}

func (s *DescribeApiResponseBody) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiResponseBody) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiResponseBody) GetAppCodeAuthType() *string {
	return s.AppCodeAuthType
}

func (s *DescribeApiResponseBody) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeApiResponseBody) GetBackendConfig() *DescribeApiResponseBodyBackendConfig {
	return s.BackendConfig
}

func (s *DescribeApiResponseBody) GetBackendEnable() *bool {
	return s.BackendEnable
}

func (s *DescribeApiResponseBody) GetConstantParameters() *DescribeApiResponseBodyConstantParameters {
	return s.ConstantParameters
}

func (s *DescribeApiResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeApiResponseBody) GetCustomSystemParameters() *DescribeApiResponseBodyCustomSystemParameters {
	return s.CustomSystemParameters
}

func (s *DescribeApiResponseBody) GetDeployedInfos() *DescribeApiResponseBodyDeployedInfos {
	return s.DeployedInfos
}

func (s *DescribeApiResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiResponseBody) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *DescribeApiResponseBody) GetErrorCodeSamples() *DescribeApiResponseBodyErrorCodeSamples {
	return s.ErrorCodeSamples
}

func (s *DescribeApiResponseBody) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *DescribeApiResponseBody) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *DescribeApiResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiResponseBody) GetMock() *string {
	return s.Mock
}

func (s *DescribeApiResponseBody) GetMockResult() *string {
	return s.MockResult
}

func (s *DescribeApiResponseBody) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeApiResponseBody) GetOpenIdConnectConfig() *DescribeApiResponseBodyOpenIdConnectConfig {
	return s.OpenIdConnectConfig
}

func (s *DescribeApiResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiResponseBody) GetRequestConfig() *DescribeApiResponseBodyRequestConfig {
	return s.RequestConfig
}

func (s *DescribeApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiResponseBody) GetRequestParameters() *DescribeApiResponseBodyRequestParameters {
	return s.RequestParameters
}

func (s *DescribeApiResponseBody) GetResultBodyModel() *string {
	return s.ResultBodyModel
}

func (s *DescribeApiResponseBody) GetResultSample() *string {
	return s.ResultSample
}

func (s *DescribeApiResponseBody) GetResultType() *string {
	return s.ResultType
}

func (s *DescribeApiResponseBody) GetServiceConfig() *DescribeApiResponseBodyServiceConfig {
	return s.ServiceConfig
}

func (s *DescribeApiResponseBody) GetServiceParameters() *DescribeApiResponseBodyServiceParameters {
	return s.ServiceParameters
}

func (s *DescribeApiResponseBody) GetServiceParametersMap() *DescribeApiResponseBodyServiceParametersMap {
	return s.ServiceParametersMap
}

func (s *DescribeApiResponseBody) GetSystemParameters() *DescribeApiResponseBodySystemParameters {
	return s.SystemParameters
}

func (s *DescribeApiResponseBody) GetTagList() *DescribeApiResponseBodyTagList {
	return s.TagList
}

func (s *DescribeApiResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *DescribeApiResponseBody) GetWebSocketApiType() *string {
	return s.WebSocketApiType
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

func (s *DescribeApiResponseBody) SetBackendConfig(v *DescribeApiResponseBodyBackendConfig) *DescribeApiResponseBody {
	s.BackendConfig = v
	return s
}

func (s *DescribeApiResponseBody) SetBackendEnable(v bool) *DescribeApiResponseBody {
	s.BackendEnable = &v
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

func (s *DescribeApiResponseBody) SetTagList(v *DescribeApiResponseBodyTagList) *DescribeApiResponseBody {
	s.TagList = v
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

func (s *DescribeApiResponseBody) Validate() error {
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
	if s.DeployedInfos != nil {
		if err := s.DeployedInfos.Validate(); err != nil {
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
	if s.TagList != nil {
		if err := s.TagList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApiResponseBodyBackendConfig struct {
	// The ID of the backend service.
	//
	// example:
	//
	// 0038e00c3dca44fcba3a94015d8f5bbf
	BackendId *string `json:"BackendId,omitempty" xml:"BackendId,omitempty"`
	// The name of the backend service.
	//
	// example:
	//
	// testoss
	BackendName *string `json:"BackendName,omitempty" xml:"BackendName,omitempty"`
	// Backend service type
	//
	// example:
	//
	// HTTP
	BackendType *string `json:"BackendType,omitempty" xml:"BackendType,omitempty"`
}

func (s DescribeApiResponseBodyBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyBackendConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyBackendConfig) GetBackendId() *string {
	return s.BackendId
}

func (s *DescribeApiResponseBodyBackendConfig) GetBackendName() *string {
	return s.BackendName
}

func (s *DescribeApiResponseBodyBackendConfig) GetBackendType() *string {
	return s.BackendType
}

func (s *DescribeApiResponseBodyBackendConfig) SetBackendId(v string) *DescribeApiResponseBodyBackendConfig {
	s.BackendId = &v
	return s
}

func (s *DescribeApiResponseBodyBackendConfig) SetBackendName(v string) *DescribeApiResponseBodyBackendConfig {
	s.BackendName = &v
	return s
}

func (s *DescribeApiResponseBodyBackendConfig) SetBackendType(v string) *DescribeApiResponseBodyBackendConfig {
	s.BackendType = &v
	return s
}

func (s *DescribeApiResponseBodyBackendConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyConstantParameters struct {
	ConstantParameter []*DescribeApiResponseBodyConstantParametersConstantParameter `json:"ConstantParameter,omitempty" xml:"ConstantParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyConstantParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyConstantParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyConstantParameters) GetConstantParameter() []*DescribeApiResponseBodyConstantParametersConstantParameter {
	return s.ConstantParameter
}

func (s *DescribeApiResponseBodyConstantParameters) SetConstantParameter(v []*DescribeApiResponseBodyConstantParametersConstantParameter) *DescribeApiResponseBodyConstantParameters {
	s.ConstantParameter = v
	return s
}

func (s *DescribeApiResponseBodyConstantParameters) Validate() error {
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

type DescribeApiResponseBodyConstantParametersConstantParameter struct {
	// The constant parameter value.
	//
	// example:
	//
	// constance
	ConstantValue *string `json:"ConstantValue,omitempty" xml:"ConstantValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// constance
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

func (s DescribeApiResponseBodyConstantParametersConstantParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyConstantParametersConstantParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) GetConstantValue() *string {
	return s.ConstantValue
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiResponseBodyConstantParametersConstantParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyCustomSystemParameters struct {
	CustomSystemParameter []*DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter `json:"CustomSystemParameter,omitempty" xml:"CustomSystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyCustomSystemParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyCustomSystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyCustomSystemParameters) GetCustomSystemParameter() []*DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter {
	return s.CustomSystemParameter
}

func (s *DescribeApiResponseBodyCustomSystemParameters) SetCustomSystemParameter(v []*DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) *DescribeApiResponseBodyCustomSystemParameters {
	s.CustomSystemParameter = v
	return s
}

func (s *DescribeApiResponseBodyCustomSystemParameters) Validate() error {
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

type DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter struct {
	// The example value.
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// Client IP Address
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
	// The mapped parameter name in the backend service.
	//
	// example:
	//
	// clientIp
	ServiceParameterName *string `json:"ServiceParameterName,omitempty" xml:"ServiceParameterName,omitempty"`
}

func (s DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiResponseBodyCustomSystemParametersCustomSystemParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyDeployedInfos struct {
	DeployedInfo []*DescribeApiResponseBodyDeployedInfosDeployedInfo `json:"DeployedInfo,omitempty" xml:"DeployedInfo,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyDeployedInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyDeployedInfos) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyDeployedInfos) GetDeployedInfo() []*DescribeApiResponseBodyDeployedInfosDeployedInfo {
	return s.DeployedInfo
}

func (s *DescribeApiResponseBodyDeployedInfos) SetDeployedInfo(v []*DescribeApiResponseBodyDeployedInfosDeployedInfo) *DescribeApiResponseBodyDeployedInfos {
	s.DeployedInfo = v
	return s
}

func (s *DescribeApiResponseBodyDeployedInfos) Validate() error {
	if s.DeployedInfo != nil {
		for _, item := range s.DeployedInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiResponseBodyDeployedInfosDeployedInfo struct {
	// The deployment status. Valid values: DEPLOYED and NONDEPLOYED.
	//
	// example:
	//
	// DEPLOYED
	DeployedStatus *string `json:"DeployedStatus,omitempty" xml:"DeployedStatus,omitempty"`
	// The effective version.
	//
	// example:
	//
	// xxx
	EffectiveVersion *string `json:"EffectiveVersion,omitempty" xml:"EffectiveVersion,omitempty"`
	// The environment to which the API is published. Valid values: RELEASE and TEST.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s DescribeApiResponseBodyDeployedInfosDeployedInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyDeployedInfosDeployedInfo) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) GetDeployedStatus() *string {
	return s.DeployedStatus
}

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) GetStageName() *string {
	return s.StageName
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

func (s *DescribeApiResponseBodyDeployedInfosDeployedInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyErrorCodeSamples) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyErrorCodeSamples) GetErrorCodeSample() []*DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample {
	return s.ErrorCodeSample
}

func (s *DescribeApiResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeApiResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

func (s *DescribeApiResponseBodyErrorCodeSamples) Validate() error {
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

type DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample struct {
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
	// The UserId parameter is missing from the request.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// Missing the parameter UserId
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The model.
	//
	// example:
	//
	// [\\"*\\"]
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
}

func (s DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) GetCode() *string {
	return s.Code
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) GetMessage() *string {
	return s.Message
}

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) GetModel() *string {
	return s.Model
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

func (s *DescribeApiResponseBodyErrorCodeSamplesErrorCodeSample) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyOpenIdConnectConfig struct {
	// The name of the parameter that corresponds to the token.
	//
	// example:
	//
	// xxx
	IdTokenParamName *string `json:"IdTokenParamName,omitempty" xml:"IdTokenParamName,omitempty"`
	// The OpenID Connect mode. Valid values:
	//
	// 	- **IDTOKEN**: indicates the APIs that are called by clients to obtain tokens. If you specify this value, the PublicKeyId parameter and the PublicKey parameter are required.
	//
	// 	- **BUSINESS**: indicates business APIs. Tokens are used to call the business APIs. If you specify this value, the IdTokenParamName parameter is required.
	//
	// example:
	//
	// IDTOKEN
	OpenIdApiType *string `json:"OpenIdApiType,omitempty" xml:"OpenIdApiType,omitempty"`
	// The public key.
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

func (s DescribeApiResponseBodyOpenIdConnectConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyOpenIdConnectConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) GetIdTokenParamName() *string {
	return s.IdTokenParamName
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) GetOpenIdApiType() *string {
	return s.OpenIdApiType
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) GetPublicKey() *string {
	return s.PublicKey
}

func (s *DescribeApiResponseBodyOpenIdConnectConfig) GetPublicKeyId() *string {
	return s.PublicKeyId
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

func (s *DescribeApiResponseBodyOpenIdConnectConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyRequestConfig struct {
	// This parameter takes effect only when the RequestMode parameter is set to MAPPING.********
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
	// https://apigateway.aliyun.com/models/3a240a127dcc4afd9ab1bf7e947b4095/9e2df550e85b4121a79ec33e2619eaab
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
	// The HTTP method used to make the request. Valid values: GET, POST, DELETE, PUT, HEADER, TRACE, PATCH, CONNECT, and OPTIONS.
	//
	// example:
	//
	// POST
	RequestHttpMethod *string `json:"RequestHttpMethod,omitempty" xml:"RequestHttpMethod,omitempty"`
	// The request mode. Valid values: MAPPING and PASSTHROUGH.
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
	// The protocol type supported by the API. Valid values: HTTP and HTTPS. Separate multiple values with commas (,), such as "HTTP,HTTPS".
	//
	// example:
	//
	// HTTP
	RequestProtocol *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeApiResponseBodyRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestConfig) GetBodyFormat() *string {
	return s.BodyFormat
}

func (s *DescribeApiResponseBodyRequestConfig) GetBodyModel() *string {
	return s.BodyModel
}

func (s *DescribeApiResponseBodyRequestConfig) GetEscapePathParam() *bool {
	return s.EscapePathParam
}

func (s *DescribeApiResponseBodyRequestConfig) GetPostBodyDescription() *string {
	return s.PostBodyDescription
}

func (s *DescribeApiResponseBodyRequestConfig) GetRequestHttpMethod() *string {
	return s.RequestHttpMethod
}

func (s *DescribeApiResponseBodyRequestConfig) GetRequestMode() *string {
	return s.RequestMode
}

func (s *DescribeApiResponseBodyRequestConfig) GetRequestPath() *string {
	return s.RequestPath
}

func (s *DescribeApiResponseBodyRequestConfig) GetRequestProtocol() *string {
	return s.RequestProtocol
}

func (s *DescribeApiResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeApiResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiResponseBodyRequestConfig) SetBodyModel(v string) *DescribeApiResponseBodyRequestConfig {
	s.BodyModel = &v
	return s
}

func (s *DescribeApiResponseBodyRequestConfig) SetEscapePathParam(v bool) *DescribeApiResponseBodyRequestConfig {
	s.EscapePathParam = &v
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

func (s *DescribeApiResponseBodyRequestConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyRequestParameters struct {
	RequestParameter []*DescribeApiResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestParameters) GetRequestParameter() []*DescribeApiResponseBodyRequestParametersRequestParameter {
	return s.RequestParameter
}

func (s *DescribeApiResponseBodyRequestParameters) SetRequestParameter(v []*DescribeApiResponseBodyRequestParametersRequestParameter) *DescribeApiResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

func (s *DescribeApiResponseBodyRequestParameters) Validate() error {
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

type DescribeApiResponseBodyRequestParametersRequestParameter struct {
	// The parameter name.
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
	// The example value.
	//
	// example:
	//
	// 20
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// Age
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
	// The JSON Schema used for JSON validation when **ParameterType*	- is set to String.
	//
	// example:
	//
	// JSON
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

func (s DescribeApiResponseBodyRequestParametersRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetApiParameterName() *string {
	return s.ApiParameterName
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetArrayItemsType() *string {
	return s.ArrayItemsType
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetDocOrder() *int32 {
	return s.DocOrder
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetDocShow() *string {
	return s.DocShow
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetJsonScheme() *string {
	return s.JsonScheme
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetMinLength() *int64 {
	return s.MinLength
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetMinValue() *int64 {
	return s.MinValue
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetRegularExpression() *string {
	return s.RegularExpression
}

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) GetRequired() *string {
	return s.Required
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

func (s *DescribeApiResponseBodyRequestParametersRequestParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceConfig struct {
	// The application name in AONE.
	//
	// example:
	//
	// ib-blank
	AoneAppName *string `json:"AoneAppName,omitempty" xml:"AoneAppName,omitempty"`
	// The ContentType header type used when you call the backend service over HTTP.
	//
	// 	- **DEFAULT**: the default header type in API Gateway
	//
	// 	- **CUSTOM**: a custom header type
	//
	// 	- **CLIENT**: the ContentType header type of the client
	//
	// example:
	//
	// CUSTOM
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	// The value of the ContentType header when the ServiceProtocol parameter is set to HTTP and the ContentTypeCatagory parameter is set to DEFAULT or CUSTOM.
	//
	// example:
	//
	// application/x-www-form-urlencoded; charset=UTF-8
	ContentTypeValue *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	// Configuration items of EventBridge
	EventBridgeConfig *DescribeApiResponseBodyServiceConfigEventBridgeConfig `json:"EventBridgeConfig,omitempty" xml:"EventBridgeConfig,omitempty" type:"Struct"`
	// Backend configuration items when the backend service is Function Compute
	FunctionComputeConfig *DescribeApiResponseBodyServiceConfigFunctionComputeConfig `json:"FunctionComputeConfig,omitempty" xml:"FunctionComputeConfig,omitempty" type:"Struct"`
	// Specifies whether to enable the Mock mode. Valid values:
	//
	// 	- **TRUE**: The Mock mode is enabled.
	//
	// 	- **FALSE**: The Mock mode is not enabled.
	//
	// example:
	//
	// TRUE
	Mock *string `json:"Mock,omitempty" xml:"Mock,omitempty"`
	// The simulated headers.
	MockHeaders *DescribeApiResponseBodyServiceConfigMockHeaders `json:"MockHeaders,omitempty" xml:"MockHeaders,omitempty" type:"Struct"`
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
	// The information returned when the backend service is Object Storage Service (OSS).
	OssConfig *DescribeApiResponseBodyServiceConfigOssConfig `json:"OssConfig,omitempty" xml:"OssConfig,omitempty" type:"Struct"`
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
	// The protocol used by the backend service. Valid values: HTTP and HTTPS.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The timeout period of the backend service. Unit: milliseconds.
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
	VpcConfig *DescribeApiResponseBodyServiceConfigVpcConfig `json:"VpcConfig,omitempty" xml:"VpcConfig,omitempty" type:"Struct"`
}

func (s DescribeApiResponseBodyServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfig) GetAoneAppName() *string {
	return s.AoneAppName
}

func (s *DescribeApiResponseBodyServiceConfig) GetContentTypeCatagory() *string {
	return s.ContentTypeCatagory
}

func (s *DescribeApiResponseBodyServiceConfig) GetContentTypeValue() *string {
	return s.ContentTypeValue
}

func (s *DescribeApiResponseBodyServiceConfig) GetEventBridgeConfig() *DescribeApiResponseBodyServiceConfigEventBridgeConfig {
	return s.EventBridgeConfig
}

func (s *DescribeApiResponseBodyServiceConfig) GetFunctionComputeConfig() *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	return s.FunctionComputeConfig
}

func (s *DescribeApiResponseBodyServiceConfig) GetMock() *string {
	return s.Mock
}

func (s *DescribeApiResponseBodyServiceConfig) GetMockHeaders() *DescribeApiResponseBodyServiceConfigMockHeaders {
	return s.MockHeaders
}

func (s *DescribeApiResponseBodyServiceConfig) GetMockResult() *string {
	return s.MockResult
}

func (s *DescribeApiResponseBodyServiceConfig) GetMockStatusCode() *int32 {
	return s.MockStatusCode
}

func (s *DescribeApiResponseBodyServiceConfig) GetOssConfig() *DescribeApiResponseBodyServiceConfigOssConfig {
	return s.OssConfig
}

func (s *DescribeApiResponseBodyServiceConfig) GetServiceAddress() *string {
	return s.ServiceAddress
}

func (s *DescribeApiResponseBodyServiceConfig) GetServiceHttpMethod() *string {
	return s.ServiceHttpMethod
}

func (s *DescribeApiResponseBodyServiceConfig) GetServicePath() *string {
	return s.ServicePath
}

func (s *DescribeApiResponseBodyServiceConfig) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *DescribeApiResponseBodyServiceConfig) GetServiceTimeout() *int32 {
	return s.ServiceTimeout
}

func (s *DescribeApiResponseBodyServiceConfig) GetServiceVpcEnable() *string {
	return s.ServiceVpcEnable
}

func (s *DescribeApiResponseBodyServiceConfig) GetVpcConfig() *DescribeApiResponseBodyServiceConfigVpcConfig {
	return s.VpcConfig
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

func (s *DescribeApiResponseBodyServiceConfig) SetEventBridgeConfig(v *DescribeApiResponseBodyServiceConfigEventBridgeConfig) *DescribeApiResponseBodyServiceConfig {
	s.EventBridgeConfig = v
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

func (s *DescribeApiResponseBodyServiceConfig) Validate() error {
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

type DescribeApiResponseBodyServiceConfigEventBridgeConfig struct {
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
	// The event source.
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

func (s DescribeApiResponseBodyServiceConfigEventBridgeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigEventBridgeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) GetEventBridgeRegionId() *string {
	return s.EventBridgeRegionId
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) GetEventBus() *string {
	return s.EventBus
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) GetEventSource() *string {
	return s.EventSource
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) SetEventBridgeRegionId(v string) *DescribeApiResponseBodyServiceConfigEventBridgeConfig {
	s.EventBridgeRegionId = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) SetEventBus(v string) *DescribeApiResponseBodyServiceConfigEventBridgeConfig {
	s.EventBus = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) SetEventSource(v string) *DescribeApiResponseBodyServiceConfigEventBridgeConfig {
	s.EventSource = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) SetRoleArn(v string) *DescribeApiResponseBodyServiceConfigEventBridgeConfig {
	s.RoleArn = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigEventBridgeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceConfigFunctionComputeConfig struct {
	// The ContentType header type used when you call the backend service over HTTP.
	//
	// 	- **DEFAULT**: the default header type in API Gateway
	//
	// 	- **CUSTOM**: a custom header type
	//
	// 	- **CLIENT**: the ContentType header type of the client
	//
	// example:
	//
	// DEFAULT
	ContentTypeCatagory *string `json:"ContentTypeCatagory,omitempty" xml:"ContentTypeCatagory,omitempty"`
	// The value of the ContentType header when the ContentTypeCatagory parameter is set to DEFAULT or CUSTOM.
	//
	// example:
	//
	// application/x-www-form-urlencoded; charset=UTF-8
	ContentTypeValue *string `json:"ContentTypeValue,omitempty" xml:"ContentTypeValue,omitempty"`
	// The root path of Function Compute.
	//
	// example:
	//
	// https://1227****64334133.ap-southeast-1-int***al.fc.aliyuncs.com/201****-15/proxy/test****ice.LATEST/testHttp/
	FcBaseUrl *string `json:"FcBaseUrl,omitempty" xml:"FcBaseUrl,omitempty"`
	// The type of the Function Compute instance.
	//
	// example:
	//
	// HttpTrigger
	FcType    *string `json:"FcType,omitempty" xml:"FcType,omitempty"`
	FcVersion *string `json:"FcVersion,omitempty" xml:"FcVersion,omitempty"`
	// The function name defined in Function Compute.
	//
	// example:
	//
	// edge_function
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
	// The region where the Function Compute instance is located.
	//
	// example:
	//
	// cn-qingdao
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
	// The name of the trigger.
	//
	// You can specify the TriggerName or TriggerUrl parameter. The TriggerName parameter is optional.
	//
	// example:
	//
	// test1
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigFunctionComputeConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetContentTypeCatagory() *string {
	return s.ContentTypeCatagory
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetContentTypeValue() *string {
	return s.ContentTypeValue
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetFcBaseUrl() *string {
	return s.FcBaseUrl
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetFcType() *string {
	return s.FcType
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetFcVersion() *string {
	return s.FcVersion
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetMethod() *string {
	return s.Method
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetOnlyBusinessPath() *bool {
	return s.OnlyBusinessPath
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetPath() *string {
	return s.Path
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetQualifier() *string {
	return s.Qualifier
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) GetTriggerName() *string {
	return s.TriggerName
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

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetFcVersion(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.FcVersion = &v
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

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) SetTriggerName(v string) *DescribeApiResponseBodyServiceConfigFunctionComputeConfig {
	s.TriggerName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigFunctionComputeConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceConfigMockHeaders struct {
	MockHeader []*DescribeApiResponseBodyServiceConfigMockHeadersMockHeader `json:"MockHeader,omitempty" xml:"MockHeader,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceConfigMockHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigMockHeaders) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigMockHeaders) GetMockHeader() []*DescribeApiResponseBodyServiceConfigMockHeadersMockHeader {
	return s.MockHeader
}

func (s *DescribeApiResponseBodyServiceConfigMockHeaders) SetMockHeader(v []*DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) *DescribeApiResponseBodyServiceConfigMockHeaders {
	s.MockHeader = v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigMockHeaders) Validate() error {
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

type DescribeApiResponseBodyServiceConfigMockHeadersMockHeader struct {
	// The HTTP header.
	//
	// example:
	//
	// Content-Length
	HeaderName *string `json:"HeaderName,omitempty" xml:"HeaderName,omitempty"`
	// The value of the HTTP header.
	//
	// example:
	//
	// 86400
	HeaderValue *string `json:"HeaderValue,omitempty" xml:"HeaderValue,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) GetHeaderName() *string {
	return s.HeaderName
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) GetHeaderValue() *string {
	return s.HeaderValue
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderName(v string) *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) SetHeaderValue(v string) *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader {
	s.HeaderValue = &v
	return s
}

func (s *DescribeApiResponseBodyServiceConfigMockHeadersMockHeader) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceConfigOssConfig struct {
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
	// cbg-db
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The stored object or folder path.
	//
	// example:
	//
	// /folder/test.json
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the region where the OSS instance is located.
	//
	// example:
	//
	// cn-hangzhou
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigOssConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigOssConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) GetAction() *string {
	return s.Action
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) GetKey() *string {
	return s.Key
}

func (s *DescribeApiResponseBodyServiceConfigOssConfig) GetOssRegionId() *string {
	return s.OssRegionId
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

func (s *DescribeApiResponseBodyServiceConfigOssConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceConfigVpcConfig struct {
	// The ID of the ECS or SLB instance in the VPC.
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
	// The VPC protocol.
	//
	// example:
	//
	// HTTP
	VpcScheme *string `json:"VpcScheme,omitempty" xml:"VpcScheme,omitempty"`
}

func (s DescribeApiResponseBodyServiceConfigVpcConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceConfigVpcConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) GetName() *string {
	return s.Name
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) GetPort() *int32 {
	return s.Port
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) GetVpcScheme() *string {
	return s.VpcScheme
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

func (s *DescribeApiResponseBodyServiceConfigVpcConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceParameters struct {
	ServiceParameter []*DescribeApiResponseBodyServiceParametersServiceParameter `json:"ServiceParameter,omitempty" xml:"ServiceParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParameters) GetServiceParameter() []*DescribeApiResponseBodyServiceParametersServiceParameter {
	return s.ServiceParameter
}

func (s *DescribeApiResponseBodyServiceParameters) SetServiceParameter(v []*DescribeApiResponseBodyServiceParametersServiceParameter) *DescribeApiResponseBodyServiceParameters {
	s.ServiceParameter = v
	return s
}

func (s *DescribeApiResponseBodyServiceParameters) Validate() error {
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

type DescribeApiResponseBodyServiceParametersServiceParameter struct {
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

func (s DescribeApiResponseBodyServiceParametersServiceParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersServiceParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiResponseBodyServiceParametersServiceParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyServiceParametersMap struct {
	ServiceParameterMap []*DescribeApiResponseBodyServiceParametersMapServiceParameterMap `json:"ServiceParameterMap,omitempty" xml:"ServiceParameterMap,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyServiceParametersMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersMap) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersMap) GetServiceParameterMap() []*DescribeApiResponseBodyServiceParametersMapServiceParameterMap {
	return s.ServiceParameterMap
}

func (s *DescribeApiResponseBodyServiceParametersMap) SetServiceParameterMap(v []*DescribeApiResponseBodyServiceParametersMapServiceParameterMap) *DescribeApiResponseBodyServiceParametersMap {
	s.ServiceParameterMap = v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersMap) Validate() error {
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

type DescribeApiResponseBodyServiceParametersMapServiceParameterMap struct {
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

func (s DescribeApiResponseBodyServiceParametersMapServiceParameterMap) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyServiceParametersMapServiceParameterMap) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) GetRequestParameterName() *string {
	return s.RequestParameterName
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) GetServiceParameterName() *string {
	return s.ServiceParameterName
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) SetRequestParameterName(v string) *DescribeApiResponseBodyServiceParametersMapServiceParameterMap {
	s.RequestParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) SetServiceParameterName(v string) *DescribeApiResponseBodyServiceParametersMapServiceParameterMap {
	s.ServiceParameterName = &v
	return s
}

func (s *DescribeApiResponseBodyServiceParametersMapServiceParameterMap) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodySystemParameters struct {
	SystemParameter []*DescribeApiResponseBodySystemParametersSystemParameter `json:"SystemParameter,omitempty" xml:"SystemParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodySystemParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodySystemParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodySystemParameters) GetSystemParameter() []*DescribeApiResponseBodySystemParametersSystemParameter {
	return s.SystemParameter
}

func (s *DescribeApiResponseBodySystemParameters) SetSystemParameter(v []*DescribeApiResponseBodySystemParametersSystemParameter) *DescribeApiResponseBodySystemParameters {
	s.SystemParameter = v
	return s
}

func (s *DescribeApiResponseBodySystemParameters) Validate() error {
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

type DescribeApiResponseBodySystemParametersSystemParameter struct {
	// The example value.
	//
	// example:
	//
	// 192.168.1.1
	DemoValue *string `json:"DemoValue,omitempty" xml:"DemoValue,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// Client IP Address
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

func (s DescribeApiResponseBodySystemParametersSystemParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodySystemParametersSystemParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeApiResponseBodySystemParametersSystemParameter) GetServiceParameterName() *string {
	return s.ServiceParameterName
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

func (s *DescribeApiResponseBodySystemParametersSystemParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeApiResponseBodyTagList struct {
	Tag []*DescribeApiResponseBodyTagListTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeApiResponseBodyTagList) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyTagList) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyTagList) GetTag() []*DescribeApiResponseBodyTagListTag {
	return s.Tag
}

func (s *DescribeApiResponseBodyTagList) SetTag(v []*DescribeApiResponseBodyTagListTag) *DescribeApiResponseBodyTagList {
	s.Tag = v
	return s
}

func (s *DescribeApiResponseBodyTagList) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApiResponseBodyTagListTag struct {
	// Label key.
	//
	// example:
	//
	// APP
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Label value.
	//
	// example:
	//
	// value3
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeApiResponseBodyTagListTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiResponseBodyTagListTag) GoString() string {
	return s.String()
}

func (s *DescribeApiResponseBodyTagListTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeApiResponseBodyTagListTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeApiResponseBodyTagListTag) SetTagKey(v string) *DescribeApiResponseBodyTagListTag {
	s.TagKey = &v
	return s
}

func (s *DescribeApiResponseBodyTagListTag) SetTagValue(v string) *DescribeApiResponseBodyTagListTag {
	s.TagValue = &v
	return s
}

func (s *DescribeApiResponseBodyTagListTag) Validate() error {
	return dara.Validate(s)
}
