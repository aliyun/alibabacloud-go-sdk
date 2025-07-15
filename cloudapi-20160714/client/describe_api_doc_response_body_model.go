// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiDocResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiDocResponseBody
	GetApiId() *string
	SetApiName(v string) *DescribeApiDocResponseBody
	GetApiName() *string
	SetAuthType(v string) *DescribeApiDocResponseBody
	GetAuthType() *string
	SetDeployedTime(v string) *DescribeApiDocResponseBody
	GetDeployedTime() *string
	SetDescription(v string) *DescribeApiDocResponseBody
	GetDescription() *string
	SetDisableInternet(v bool) *DescribeApiDocResponseBody
	GetDisableInternet() *bool
	SetErrorCodeSamples(v *DescribeApiDocResponseBodyErrorCodeSamples) *DescribeApiDocResponseBody
	GetErrorCodeSamples() *DescribeApiDocResponseBodyErrorCodeSamples
	SetFailResultSample(v string) *DescribeApiDocResponseBody
	GetFailResultSample() *string
	SetForceNonceCheck(v bool) *DescribeApiDocResponseBody
	GetForceNonceCheck() *bool
	SetGroupId(v string) *DescribeApiDocResponseBody
	GetGroupId() *string
	SetGroupName(v string) *DescribeApiDocResponseBody
	GetGroupName() *string
	SetRegionId(v string) *DescribeApiDocResponseBody
	GetRegionId() *string
	SetRequestConfig(v *DescribeApiDocResponseBodyRequestConfig) *DescribeApiDocResponseBody
	GetRequestConfig() *DescribeApiDocResponseBodyRequestConfig
	SetRequestId(v string) *DescribeApiDocResponseBody
	GetRequestId() *string
	SetRequestParameters(v *DescribeApiDocResponseBodyRequestParameters) *DescribeApiDocResponseBody
	GetRequestParameters() *DescribeApiDocResponseBodyRequestParameters
	SetResultSample(v string) *DescribeApiDocResponseBody
	GetResultSample() *string
	SetResultType(v string) *DescribeApiDocResponseBody
	GetResultType() *string
	SetStageName(v string) *DescribeApiDocResponseBody
	GetStageName() *string
	SetVisibility(v string) *DescribeApiDocResponseBody
	GetVisibility() *string
}

type DescribeApiDocResponseBody struct {
	// The ID of the API.
	//
	// example:
	//
	// b24be7e59a104e52bffbf432cc9272af
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The name of the API
	//
	// example:
	//
	// ObtainKeywordQRCodeAddress
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The security authentication method. Valid values: APP, ANONYMOUS, and APPOPENID, indicating respectively Alibaba Cloud application authentication, anonymous authentication, and third-party OpenID Connect account authentication.
	//
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The publishing time.
	//
	// example:
	//
	// 2022-07-13T16:00:33Z
	DeployedTime *string `json:"DeployedTime,omitempty" xml:"DeployedTime,omitempty"`
	// The API description.
	//
	// example:
	//
	// Lynk\\&Co Digital Mall OMS-UAT
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
	ErrorCodeSamples *DescribeApiDocResponseBodyErrorCodeSamples `json:"ErrorCodeSamples,omitempty" xml:"ErrorCodeSamples,omitempty" type:"Struct"`
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
	// f51d08c5b7c84342905544ebaec26d35
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the API group.
	//
	// example:
	//
	// Member Age Transaction Service
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The region ID of the API group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The returned API frontend definition. It is an array consisting of RequestConfig data.
	RequestConfig *DescribeApiDocResponseBodyRequestConfig `json:"RequestConfig,omitempty" xml:"RequestConfig,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F253FB5F-9AE1-5DDA-99B5-46BE00A3719E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned frontend input parameters in the API. It is an array consisting of RequestParameter data.
	RequestParameters *DescribeApiDocResponseBodyRequestParameters `json:"RequestParameters,omitempty" xml:"RequestParameters,omitempty" type:"Struct"`
	// The sample response.
	//
	// example:
	//
	// {\\n  \\"status\\": 0,\\n  \\"data\\": {\\n    \\"count\\": 1,\\n    \\"list\\": [\\n      \\"352\\"\\n    ]\\n  },\\n  \\"message\\": \\"success\\"\\n}
	ResultSample *string `json:"ResultSample,omitempty" xml:"ResultSample,omitempty"`
	// The return value type.
	//
	// example:
	//
	// JSON
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// Indicates whether the API is public. Valid values: PUBLIC and PRIVATE.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s DescribeApiDocResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBody) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiDocResponseBody) GetApiName() *string {
	return s.ApiName
}

func (s *DescribeApiDocResponseBody) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeApiDocResponseBody) GetDeployedTime() *string {
	return s.DeployedTime
}

func (s *DescribeApiDocResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiDocResponseBody) GetDisableInternet() *bool {
	return s.DisableInternet
}

func (s *DescribeApiDocResponseBody) GetErrorCodeSamples() *DescribeApiDocResponseBodyErrorCodeSamples {
	return s.ErrorCodeSamples
}

func (s *DescribeApiDocResponseBody) GetFailResultSample() *string {
	return s.FailResultSample
}

func (s *DescribeApiDocResponseBody) GetForceNonceCheck() *bool {
	return s.ForceNonceCheck
}

func (s *DescribeApiDocResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeApiDocResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeApiDocResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApiDocResponseBody) GetRequestConfig() *DescribeApiDocResponseBodyRequestConfig {
	return s.RequestConfig
}

func (s *DescribeApiDocResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiDocResponseBody) GetRequestParameters() *DescribeApiDocResponseBodyRequestParameters {
	return s.RequestParameters
}

func (s *DescribeApiDocResponseBody) GetResultSample() *string {
	return s.ResultSample
}

func (s *DescribeApiDocResponseBody) GetResultType() *string {
	return s.ResultType
}

func (s *DescribeApiDocResponseBody) GetStageName() *string {
	return s.StageName
}

func (s *DescribeApiDocResponseBody) GetVisibility() *string {
	return s.Visibility
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

func (s *DescribeApiDocResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiDocResponseBodyErrorCodeSamples struct {
	ErrorCodeSample []*DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample `json:"ErrorCodeSample,omitempty" xml:"ErrorCodeSample,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyErrorCodeSamples) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponseBodyErrorCodeSamples) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyErrorCodeSamples) GetErrorCodeSample() []*DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample {
	return s.ErrorCodeSample
}

func (s *DescribeApiDocResponseBodyErrorCodeSamples) SetErrorCodeSample(v []*DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) *DescribeApiDocResponseBodyErrorCodeSamples {
	s.ErrorCodeSample = v
	return s
}

func (s *DescribeApiDocResponseBodyErrorCodeSamples) Validate() error {
	return dara.Validate(s)
}

type DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample struct {
	// The returned error code.
	//
	// example:
	//
	// Error
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error description.
	//
	// example:
	//
	// Unauthorized
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) GetCode() *string {
	return s.Code
}

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) GetMessage() *string {
	return s.Message
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

func (s *DescribeApiDocResponseBodyErrorCodeSamplesErrorCodeSample) Validate() error {
	return dara.Validate(s)
}

type DescribeApiDocResponseBodyRequestConfig struct {
	// This parameter takes effect only when the RequestMode parameter is set to MAPPING.********
	//
	// The server data transmission method used for POST and PUT requests. Valid values: FORM and STREAM. FORM indicates that data in key-value pairs is transmitted as forms. STREAM indicates that data is transmitted as byte streams.
	//
	// example:
	//
	// STREAM
	BodyFormat *string `json:"BodyFormat,omitempty" xml:"BodyFormat,omitempty"`
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
	// The protocol type supported by the API. Valid values: HTTP and HTTPS. Separate multiple values with commas (,), such as "HTTP,HTTPS".
	//
	// example:
	//
	// HTTP
	RequestProtocol *string `json:"RequestProtocol,omitempty" xml:"RequestProtocol,omitempty"`
}

func (s DescribeApiDocResponseBodyRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestConfig) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetBodyFormat() *string {
	return s.BodyFormat
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetEscapePathParam() *bool {
	return s.EscapePathParam
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetPostBodyDescription() *string {
	return s.PostBodyDescription
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetRequestHttpMethod() *string {
	return s.RequestHttpMethod
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetRequestMode() *string {
	return s.RequestMode
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetRequestPath() *string {
	return s.RequestPath
}

func (s *DescribeApiDocResponseBodyRequestConfig) GetRequestProtocol() *string {
	return s.RequestProtocol
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetBodyFormat(v string) *DescribeApiDocResponseBodyRequestConfig {
	s.BodyFormat = &v
	return s
}

func (s *DescribeApiDocResponseBodyRequestConfig) SetEscapePathParam(v bool) *DescribeApiDocResponseBodyRequestConfig {
	s.EscapePathParam = &v
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

func (s *DescribeApiDocResponseBodyRequestConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeApiDocResponseBodyRequestParameters struct {
	RequestParameter []*DescribeApiDocResponseBodyRequestParametersRequestParameter `json:"RequestParameter,omitempty" xml:"RequestParameter,omitempty" type:"Repeated"`
}

func (s DescribeApiDocResponseBodyRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestParameters) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestParameters) GetRequestParameter() []*DescribeApiDocResponseBodyRequestParametersRequestParameter {
	return s.RequestParameter
}

func (s *DescribeApiDocResponseBodyRequestParameters) SetRequestParameter(v []*DescribeApiDocResponseBodyRequestParametersRequestParameter) *DescribeApiDocResponseBodyRequestParameters {
	s.RequestParameter = v
	return s
}

func (s *DescribeApiDocResponseBodyRequestParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeApiDocResponseBodyRequestParametersRequestParameter struct {
	// The name of the parameter in the API request.
	//
	// example:
	//
	// Length
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
	// Parameters
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
	// The maximum length.
	//
	// example:
	//
	// 123456
	MaxLength *int64 `json:"MaxLength,omitempty" xml:"MaxLength,omitempty"`
	// The maximum value.
	//
	// example:
	//
	// 200
	MaxValue *int64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// The minimum length.
	//
	// example:
	//
	// 2
	MinLength *int64 `json:"MinLength,omitempty" xml:"MinLength,omitempty"`
	// The minimum value.
	//
	// example:
	//
	// 123456
	MinValue *int64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// The data type of the parameter.
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
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// OPTIONAL
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s DescribeApiDocResponseBodyRequestParametersRequestParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiDocResponseBodyRequestParametersRequestParameter) GoString() string {
	return s.String()
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetApiParameterName() *string {
	return s.ApiParameterName
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetArrayItemsType() *string {
	return s.ArrayItemsType
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetDemoValue() *string {
	return s.DemoValue
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetDescription() *string {
	return s.Description
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetDocOrder() *int32 {
	return s.DocOrder
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetDocShow() *string {
	return s.DocShow
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetEnumValue() *string {
	return s.EnumValue
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetJsonScheme() *string {
	return s.JsonScheme
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetLocation() *string {
	return s.Location
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetMaxLength() *int64 {
	return s.MaxLength
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetMaxValue() *int64 {
	return s.MaxValue
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetMinLength() *int64 {
	return s.MinLength
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetMinValue() *int64 {
	return s.MinValue
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetParameterType() *string {
	return s.ParameterType
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetRegularExpression() *string {
	return s.RegularExpression
}

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) GetRequired() *string {
	return s.Required
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

func (s *DescribeApiDocResponseBodyRequestParametersRequestParameter) Validate() error {
	return dara.Validate(s)
}
