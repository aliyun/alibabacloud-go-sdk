// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataServiceApisResponseBodyData) *ListDataServiceApisResponseBody
	GetData() *ListDataServiceApisResponseBodyData
	SetErrorCode(v string) *ListDataServiceApisResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataServiceApisResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListDataServiceApisResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDataServiceApisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceApisResponseBody
	GetSuccess() *bool
}

type ListDataServiceApisResponseBody struct {
	// The data returned.
	Data *ListDataServiceApisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBody) GetData() *ListDataServiceApisResponseBodyData {
	return s.Data
}

func (s *ListDataServiceApisResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServiceApisResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServiceApisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceApisResponseBody) SetData(v *ListDataServiceApisResponseBodyData) *ListDataServiceApisResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServiceApisResponseBody) SetErrorCode(v string) *ListDataServiceApisResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServiceApisResponseBody) SetErrorMessage(v string) *ListDataServiceApisResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServiceApisResponseBody) SetHttpStatusCode(v int32) *ListDataServiceApisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApisResponseBody) SetRequestId(v string) *ListDataServiceApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApisResponseBody) SetSuccess(v bool) *ListDataServiceApisResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceApisResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyData struct {
	// The list of APIs in the development state.
	Apis []*ListDataServiceApisResponseBodyDataApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	// The page number. The value of this parameter is the same as that of the PageNumber parameter in the request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceApisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyData) GetApis() []*ListDataServiceApisResponseBodyDataApis {
	return s.Apis
}

func (s *ListDataServiceApisResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceApisResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApisResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceApisResponseBodyData) SetApis(v []*ListDataServiceApisResponseBodyDataApis) *ListDataServiceApisResponseBodyData {
	s.Apis = v
	return s
}

func (s *ListDataServiceApisResponseBodyData) SetPageNumber(v int32) *ListDataServiceApisResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceApisResponseBodyData) SetPageSize(v int32) *ListDataServiceApisResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApisResponseBodyData) SetTotalCount(v int32) *ListDataServiceApisResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApisResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApis struct {
	// The API ID.
	//
	// example:
	//
	// 10002
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The type of the API. Valid values: 0, 1, and 2. The value 0 indicates that the API is generated in wizard mode. The value 1 indicates that the API is generated in script mode. The value 2 indicates that the API is generated by registration.
	//
	// example:
	//
	// 0
	ApiMode *int32 `json:"ApiMode,omitempty" xml:"ApiMode,omitempty"`
	// The name of the API.
	//
	// example:
	//
	// My API name
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The path of the API.
	//
	// example:
	//
	// /test/1
	ApiPath *string `json:"ApiPath,omitempty" xml:"ApiPath,omitempty"`
	// The time when the API was created.
	//
	// example:
	//
	// 2020-06-23T00:21:01+0800
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the creator of the API.
	//
	// example:
	//
	// 1234567
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the API.
	//
	// example:
	//
	// Test API description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// 0
	FolderId *int64 `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// abcde123456789
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the API was last modified.
	//
	// example:
	//
	// 2020-06-23T00:21:01+0800
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who last modified the API.
	//
	// example:
	//
	// 2345678
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of fields.
	Protocols []*int32 `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The details of the API generated by registration. This parameter is returned only if the API is generated by registration.
	RegistrationDetails *ListDataServiceApisResponseBodyDataApisRegistrationDetails `json:"RegistrationDetails,omitempty" xml:"RegistrationDetails,omitempty" type:"Struct"`
	// The request method of the API. Valid values: 0, 1, 2, and 3. The value 0 indicates the GET method. The value 1 indicates the POST method. The value 2 indicates the PUT method. The value 3 indicates the DELETE method. APIs generated in wizard or script mode support the GET and POST methods. APIs generated by registration support the GET, POST, PUT, and DELETE methods.
	//
	// example:
	//
	// 0
	RequestMethod *int32 `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	// The format in which the response of the API request is returned. Valid values: 0 and 1. The value 0 indicates the JSON format. The value 1 indicates the XML format. APIs generated in wizard or script mode support the JSON format. APIs generated by registration support the JSON and XML formats.
	//
	// example:
	//
	// 0
	ResponseContentType *int32 `json:"ResponseContentType,omitempty" xml:"ResponseContentType,omitempty"`
	// The details of the API generated in script mode. This parameter is returned only if the API is generated in script mode.
	ScriptDetails *ListDataServiceApisResponseBodyDataApisScriptDetails `json:"ScriptDetails,omitempty" xml:"ScriptDetails,omitempty" type:"Struct"`
	// The status of the API. Valid values: 0 and 1. The value 0 indicates that the API is not published. The value 1 indicates that the API is published.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The timeout period of the API request. Unit: milliseconds.
	//
	// example:
	//
	// 10000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The scope in which the API is visible. Valid values: 0 and 1. The value 0 indicates that the API is visible within the workspace. The value 1 indicates that the API is visible only to its owner.
	//
	// example:
	//
	// 0
	VisibleRange *int32 `json:"VisibleRange,omitempty" xml:"VisibleRange,omitempty"`
	// The details of the API generated in wizard mode. This parameter is returned only if the API is generated in wizard mode.
	WizardDetails *ListDataServiceApisResponseBodyDataApisWizardDetails `json:"WizardDetails,omitempty" xml:"WizardDetails,omitempty" type:"Struct"`
}

func (s ListDataServiceApisResponseBodyDataApis) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApis) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApis) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetApiMode() *int32 {
	return s.ApiMode
}

func (s *ListDataServiceApisResponseBodyDataApis) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceApisResponseBodyDataApis) GetApiPath() *string {
	return s.ApiPath
}

func (s *ListDataServiceApisResponseBodyDataApis) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServiceApisResponseBodyDataApis) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetDescription() *string {
	return s.Description
}

func (s *ListDataServiceApisResponseBodyDataApis) GetFolderId() *int64 {
	return s.FolderId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDataServiceApisResponseBodyDataApis) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetProtocols() []*int32 {
	return s.Protocols
}

func (s *ListDataServiceApisResponseBodyDataApis) GetRegistrationDetails() *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	return s.RegistrationDetails
}

func (s *ListDataServiceApisResponseBodyDataApis) GetRequestMethod() *int32 {
	return s.RequestMethod
}

func (s *ListDataServiceApisResponseBodyDataApis) GetResponseContentType() *int32 {
	return s.ResponseContentType
}

func (s *ListDataServiceApisResponseBodyDataApis) GetScriptDetails() *ListDataServiceApisResponseBodyDataApisScriptDetails {
	return s.ScriptDetails
}

func (s *ListDataServiceApisResponseBodyDataApis) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataServiceApisResponseBodyDataApis) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceApisResponseBodyDataApis) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListDataServiceApisResponseBodyDataApis) GetVisibleRange() *int32 {
	return s.VisibleRange
}

func (s *ListDataServiceApisResponseBodyDataApis) GetWizardDetails() *ListDataServiceApisResponseBodyDataApisWizardDetails {
	return s.WizardDetails
}

func (s *ListDataServiceApisResponseBodyDataApis) SetApiId(v int64) *ListDataServiceApisResponseBodyDataApis {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetApiMode(v int32) *ListDataServiceApisResponseBodyDataApis {
	s.ApiMode = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetApiName(v string) *ListDataServiceApisResponseBodyDataApis {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetApiPath(v string) *ListDataServiceApisResponseBodyDataApis {
	s.ApiPath = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetCreatedTime(v string) *ListDataServiceApisResponseBodyDataApis {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetCreatorId(v string) *ListDataServiceApisResponseBodyDataApis {
	s.CreatorId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetDescription(v string) *ListDataServiceApisResponseBodyDataApis {
	s.Description = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetFolderId(v int64) *ListDataServiceApisResponseBodyDataApis {
	s.FolderId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetGroupId(v string) *ListDataServiceApisResponseBodyDataApis {
	s.GroupId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetModifiedTime(v string) *ListDataServiceApisResponseBodyDataApis {
	s.ModifiedTime = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetOperatorId(v string) *ListDataServiceApisResponseBodyDataApis {
	s.OperatorId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetProjectId(v int64) *ListDataServiceApisResponseBodyDataApis {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetProtocols(v []*int32) *ListDataServiceApisResponseBodyDataApis {
	s.Protocols = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetRegistrationDetails(v *ListDataServiceApisResponseBodyDataApisRegistrationDetails) *ListDataServiceApisResponseBodyDataApis {
	s.RegistrationDetails = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetRequestMethod(v int32) *ListDataServiceApisResponseBodyDataApis {
	s.RequestMethod = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetResponseContentType(v int32) *ListDataServiceApisResponseBodyDataApis {
	s.ResponseContentType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetScriptDetails(v *ListDataServiceApisResponseBodyDataApisScriptDetails) *ListDataServiceApisResponseBodyDataApis {
	s.ScriptDetails = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetStatus(v int32) *ListDataServiceApisResponseBodyDataApis {
	s.Status = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetTenantId(v int64) *ListDataServiceApisResponseBodyDataApis {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetTimeout(v int32) *ListDataServiceApisResponseBodyDataApis {
	s.Timeout = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetVisibleRange(v int32) *ListDataServiceApisResponseBodyDataApis {
	s.VisibleRange = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) SetWizardDetails(v *ListDataServiceApisResponseBodyDataApisWizardDetails) *ListDataServiceApisResponseBodyDataApis {
	s.WizardDetails = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApis) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisRegistrationDetails struct {
	// The sample error response of the API.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
	// The error codes returned for the API generated by registration.
	RegistrationErrorCodes []*ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes `json:"RegistrationErrorCodes,omitempty" xml:"RegistrationErrorCodes,omitempty" type:"Repeated"`
	// The request parameters of the API generated by registration.
	RegistrationRequestParameters []*ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters `json:"RegistrationRequestParameters,omitempty" xml:"RegistrationRequestParameters,omitempty" type:"Repeated"`
	// The format in which the response of the API request is returned. Valid values: 0 and 1. The value 0 indicates the JSON format. The value 1 indicates the XML format. APIs generated in wizard or script mode support the JSON format. APIs generated by registration support the JSON and XML formats.
	//
	// example:
	//
	// 0
	ServiceContentType *int32 `json:"ServiceContentType,omitempty" xml:"ServiceContentType,omitempty"`
	// The URL of the backend service.
	//
	// example:
	//
	// http://example.aliyundoc.com
	ServiceHost *string `json:"ServiceHost,omitempty" xml:"ServiceHost,omitempty"`
	// The path of the backend service.
	//
	// example:
	//
	// /index
	ServicePath *string `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	// The description of the request body initiated to call the backend service.
	//
	// example:
	//
	// {"abc":1}
	ServiceRequestBodyDescription *string `json:"ServiceRequestBodyDescription,omitempty" xml:"ServiceRequestBodyDescription,omitempty"`
	// The sample success response of the API.
	//
	// example:
	//
	// {"success": true}
	SuccessfulResultSample *string `json:"SuccessfulResultSample,omitempty" xml:"SuccessfulResultSample,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisRegistrationDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisRegistrationDetails) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetFailedResultSample() *string {
	return s.FailedResultSample
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetRegistrationErrorCodes() []*ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	return s.RegistrationErrorCodes
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetRegistrationRequestParameters() []*ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	return s.RegistrationRequestParameters
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetServiceContentType() *int32 {
	return s.ServiceContentType
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetServiceHost() *string {
	return s.ServiceHost
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetServicePath() *string {
	return s.ServicePath
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetServiceRequestBodyDescription() *string {
	return s.ServiceRequestBodyDescription
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) GetSuccessfulResultSample() *string {
	return s.SuccessfulResultSample
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetFailedResultSample(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.FailedResultSample = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetRegistrationErrorCodes(v []*ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.RegistrationErrorCodes = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetRegistrationRequestParameters(v []*ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.RegistrationRequestParameters = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetServiceContentType(v int32) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.ServiceContentType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetServiceHost(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.ServiceHost = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetServicePath(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.ServicePath = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetServiceRequestBodyDescription(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.ServiceRequestBodyDescription = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) SetSuccessfulResultSample(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetails {
	s.SuccessfulResultSample = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes struct {
	// The error code.
	//
	// example:
	//
	// 1001
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// fail to call
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The solution used to fix the error.
	//
	// example:
	//
	// retry
	ErrorSolution *string `json:"ErrorSolution,omitempty" xml:"ErrorSolution,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GetErrorSolution() *string {
	return s.ErrorSolution
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) SetErrorCode(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) SetErrorMessage(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) SetErrorSolution(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	s.ErrorSolution = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters struct {
	// The name of the associated field. This parameter is supported only if the API is generated in wizard mode.
	//
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The default value.
	//
	// example:
	//
	// default1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The sample value.
	//
	// example:
	//
	// example1
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
	IsRequiredParameter *bool `json:"IsRequiredParameter,omitempty" xml:"IsRequiredParameter,omitempty"`
	// The data type of the parameter. Valid values:
	//
	// 	- 0: String
	//
	// 	- 1: Int
	//
	// 	- 2: Long
	//
	// 	- 3: Float
	//
	// 	- 4: Double
	//
	// 	- 5: Boolean
	//
	// 	- 6: StringList
	//
	// 	- 7: IntList
	//
	// 	- 8: LongList
	//
	// 	- 9: FloatList
	//
	// 	- 10: DoubleList
	//
	// 	- 11: BooleanList
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description1
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// name1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The operator used for the value of the parameter. Valid values: 0, 1, 2, and 3. The value 0 indicates the Equal operator. The value 1 indicates the Like operator. The value 2 indicates the Const operator. The value 3 indicates the In operator. APIs generated in wizard mode support the Equal, Like, and In operators. APIs generated in script mode support the Equal operator. APIs generated by registration support the Equal and Const operators.
	//
	// example:
	//
	// 0
	ParameterOperator *int32 `json:"ParameterOperator,omitempty" xml:"ParameterOperator,omitempty"`
	// The position of the parameter. Valid values: 0, 1, 2, and 3. The value 0 indicates that the parameter is in the URL path of the request. The value 1 indicates that the parameter is in the Query parameter of the request URL. The value 2 indicates that the parameter is in the request header. The value 3 indicates that the parameter is in the request body. APIs generated in wizard or script mode support only the Query position. APIs generated by registration whose request method is GET or DELETE support the Query and Head positions. APIs generated by registration whose request method is PUT or POST support the Query, Head, and Body positions.
	//
	// example:
	//
	// 0
	ParameterPosition *int32 `json:"ParameterPosition,omitempty" xml:"ParameterPosition,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterOperator() *int32 {
	return s.ParameterOperator
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterPosition() *int32 {
	return s.ParameterPosition
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetColumnName(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ColumnName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetDefaultValue(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetExampleValue(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetIsRequiredParameter(v bool) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterDataType(v int32) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterDescription(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterName(v string) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterOperator(v int32) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterOperator = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterPosition(v int32) *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterPosition = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisScriptDetails struct {
	// Indicates whether the entries are returned by page.
	//
	// example:
	//
	// true
	IsPagedResponse *bool `json:"IsPagedResponse,omitempty" xml:"IsPagedResponse,omitempty"`
	// The SQL script.
	//
	// example:
	//
	// select a from t
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// The data source information about the API generated in script mode.
	ScriptConnection *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection `json:"ScriptConnection,omitempty" xml:"ScriptConnection,omitempty" type:"Struct"`
	// The request parameters of the API generated in script mode.
	ScriptRequestParameters []*ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters `json:"ScriptRequestParameters,omitempty" xml:"ScriptRequestParameters,omitempty" type:"Repeated"`
	// The response parameters of the API generated in script mode.
	ScriptResponseParameters []*ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters `json:"ScriptResponseParameters,omitempty" xml:"ScriptResponseParameters,omitempty" type:"Repeated"`
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetails) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) GetIsPagedResponse() *bool {
	return s.IsPagedResponse
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) GetScript() *string {
	return s.Script
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) GetScriptConnection() *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection {
	return s.ScriptConnection
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) GetScriptRequestParameters() []*ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	return s.ScriptRequestParameters
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) GetScriptResponseParameters() []*ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	return s.ScriptResponseParameters
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) SetIsPagedResponse(v bool) *ListDataServiceApisResponseBodyDataApisScriptDetails {
	s.IsPagedResponse = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) SetScript(v string) *ListDataServiceApisResponseBodyDataApisScriptDetails {
	s.Script = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) SetScriptConnection(v *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) *ListDataServiceApisResponseBodyDataApisScriptDetails {
	s.ScriptConnection = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) SetScriptRequestParameters(v []*ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) *ListDataServiceApisResponseBodyDataApisScriptDetails {
	s.ScriptRequestParameters = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) SetScriptResponseParameters(v []*ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) *ListDataServiceApisResponseBodyDataApisScriptDetails {
	s.ScriptResponseParameters = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection struct {
	// The data source ID.
	//
	// example:
	//
	// 123
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The name of the table in the data source.
	//
	// example:
	//
	// t
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) GetTableName() *string {
	return s.TableName
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) SetConnectionId(v int64) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection {
	s.ConnectionId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) SetTableName(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection {
	s.TableName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptConnection) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters struct {
	// The name of the associated field. This parameter is supported only if the API is generated in wizard mode.
	//
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The default value.
	//
	// example:
	//
	// default1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The sample value.
	//
	// example:
	//
	// example1
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
	IsRequiredParameter *bool `json:"IsRequiredParameter,omitempty" xml:"IsRequiredParameter,omitempty"`
	// The data type of the parameter. Valid values:
	//
	// 	- 0: String
	//
	// 	- 1: Int
	//
	// 	- 2: Long
	//
	// 	- 3: Float
	//
	// 	- 4: Double
	//
	// 	- 5: Boolean
	//
	// 	- 6: StringList
	//
	// 	- 7: IntList
	//
	// 	- 8: LongList
	//
	// 	- 9: FloatList
	//
	// 	- 10: DoubleList
	//
	// 	- 11: BooleanList
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description1
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// param1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The operator used for the value of the parameter. Valid values:
	//
	// 	- 0: Equal
	//
	// 	- 1: Like
	//
	// 	- 2: Const
	//
	// 	- 3: In
	//
	// APIs generated in wizard mode support the Equal, Like, and In operators. APIs generated in script mode support the Equal operator. APIs generated by registration support the Equal and Const operators.
	//
	// example:
	//
	// 0
	ParameterOperator *int32 `json:"ParameterOperator,omitempty" xml:"ParameterOperator,omitempty"`
	// The position of the parameter. Valid values:
	//
	// 	- 0: indicates that the parameter is in the URL path of the request.
	//
	// 	- 1: indicates that the parameter is in the Query parameter of the request URL.
	//
	// 	- 2: indicates that the parameter is in the request header.
	//
	// 	- 3: indicates that the parameter is in the request body.
	//
	// APIs generated in wizard or script mode support only the Query position. APIs generated by registration whose request method is GET or DELETE support the Query and Head positions. APIs generated by registration whose request method is PUT or POST support the Query, Head, and Body positions.
	//
	// example:
	//
	// 0
	ParameterPosition *int32 `json:"ParameterPosition,omitempty" xml:"ParameterPosition,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterOperator() *int32 {
	return s.ParameterOperator
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterPosition() *int32 {
	return s.ParameterPosition
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetColumnName(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ColumnName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetDefaultValue(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetExampleValue(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetIsRequiredParameter(v bool) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterDataType(v int32) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterDescription(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterName(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterOperator(v int32) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterOperator = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterPosition(v int32) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterPosition = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters struct {
	// The name of the associated field. This parameter is supported only if the API is generated in wizard mode.
	//
	// example:
	//
	// column2
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The sample value.
	//
	// example:
	//
	// example2
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The data type of the parameter. Valid values:
	//
	// 	- 0: String
	//
	// 	- 1: Int
	//
	// 	- 2: Long
	//
	// 	- 3: Float
	//
	// 	- 4: Double
	//
	// 	- 5: Boolean
	//
	// 	- 6: StringList
	//
	// 	- 7: IntList
	//
	// 	- 8: LongList
	//
	// 	- 9: FloatList
	//
	// 	- 10: DoubleList
	//
	// 	- 11: BooleanList
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description2
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// param2
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetColumnName(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ColumnName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetExampleValue(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetParameterDataType(v int32) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetParameterDescription(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetParameterName(v string) *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisScriptDetailsScriptResponseParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisWizardDetails struct {
	// Indicates whether the entries are returned by page.
	//
	// example:
	//
	// true
	IsPagedResponse *bool `json:"IsPagedResponse,omitempty" xml:"IsPagedResponse,omitempty"`
	// The data source information about the API generated in wizard mode.
	WizardConnection *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection `json:"WizardConnection,omitempty" xml:"WizardConnection,omitempty" type:"Struct"`
	// The request parameters of the API generated in wizard mode.
	WizardRequestParameters []*ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters `json:"WizardRequestParameters,omitempty" xml:"WizardRequestParameters,omitempty" type:"Repeated"`
	// The response parameters of the API generated in wizard mode.
	WizardResponseParameters []*ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters `json:"WizardResponseParameters,omitempty" xml:"WizardResponseParameters,omitempty" type:"Repeated"`
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetails) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) GetIsPagedResponse() *bool {
	return s.IsPagedResponse
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) GetWizardConnection() *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection {
	return s.WizardConnection
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) GetWizardRequestParameters() []*ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	return s.WizardRequestParameters
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) GetWizardResponseParameters() []*ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	return s.WizardResponseParameters
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) SetIsPagedResponse(v bool) *ListDataServiceApisResponseBodyDataApisWizardDetails {
	s.IsPagedResponse = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) SetWizardConnection(v *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) *ListDataServiceApisResponseBodyDataApisWizardDetails {
	s.WizardConnection = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) SetWizardRequestParameters(v []*ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) *ListDataServiceApisResponseBodyDataApisWizardDetails {
	s.WizardRequestParameters = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) SetWizardResponseParameters(v []*ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) *ListDataServiceApisResponseBodyDataApisWizardDetails {
	s.WizardResponseParameters = v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection struct {
	// The data source ID.
	//
	// example:
	//
	// 123
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The name of the table in the data source.
	//
	// example:
	//
	// t
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) GetTableName() *string {
	return s.TableName
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) SetConnectionId(v int64) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection {
	s.ConnectionId = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) SetTableName(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection {
	s.TableName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardConnection) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters struct {
	// The name of the associated field. This parameter is supported only if the API is generated in wizard mode.
	//
	// example:
	//
	// column1
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The default value.
	//
	// example:
	//
	// default1
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The sample value.
	//
	// example:
	//
	// example1
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
	IsRequiredParameter *bool `json:"IsRequiredParameter,omitempty" xml:"IsRequiredParameter,omitempty"`
	// The data type of the parameter. Valid values:
	//
	// 	- 0: String
	//
	// 	- 1: Int
	//
	// 	- 2: Long
	//
	// 	- 3: Float
	//
	// 	- 4: Double
	//
	// 	- 5: Boolean
	//
	// 	- 6: StringList
	//
	// 	- 7: IntList
	//
	// 	- 8: LongList
	//
	// 	- 9: FloatList
	//
	// 	- 10: DoubleList
	//
	// 	- 11: BooleanList
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description1
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// param1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The operator used for the value of the parameter. Valid values: 0, 1, 2, and 3. The value 0 indicates the Equal operator. The value 1 indicates the Like operator. The value 2 indicates the Const operator. The value 3 indicates the In operator. APIs generated in wizard mode support the Equal, Like, and In operators. APIs generated in script mode support the Equal operator. APIs generated by registration support the Equal and Const operators.
	//
	// example:
	//
	// 0
	ParameterOperator *int32 `json:"ParameterOperator,omitempty" xml:"ParameterOperator,omitempty"`
	// The position of the parameter. Valid values: 0, 1, 2, and 3. The value 0 indicates that the parameter is in the URL path of the request. The value 1 indicates that the parameter is in the Query parameter of the request URL. The value 2 indicates that the parameter is in the request header. The value 3 indicates that the parameter is in the request body. APIs generated in wizard or script mode support only the Query position. APIs generated by registration whose request method is GET or DELETE support the Query and Head positions. APIs generated by registration whose request method is PUT or POST support the Query, Head, and Body positions.
	//
	// example:
	//
	// 0
	ParameterPosition *int32 `json:"ParameterPosition,omitempty" xml:"ParameterPosition,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterOperator() *int32 {
	return s.ParameterOperator
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterPosition() *int32 {
	return s.ParameterPosition
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetColumnName(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ColumnName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetDefaultValue(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetExampleValue(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetIsRequiredParameter(v bool) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterDataType(v int32) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterDescription(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterName(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterOperator(v int32) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterOperator = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterPosition(v int32) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterPosition = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters struct {
	// The name of the associated field. This parameter is supported only if the API is generated in wizard mode.
	//
	// example:
	//
	// column2
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The sample value.
	//
	// example:
	//
	// example2
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The data type of the parameter. Valid values:
	//
	// 	- 0: String
	//
	// 	- 1: Int
	//
	// 	- 2: Long
	//
	// 	- 3: Float
	//
	// 	- 4: Double
	//
	// 	- 5: Boolean
	//
	// 	- 6: StringList
	//
	// 	- 7: IntList
	//
	// 	- 8: LongList
	//
	// 	- 9: FloatList
	//
	// 	- 10: DoubleList
	//
	// 	- 11: BooleanList
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description2
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// param2
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GoString() string {
	return s.String()
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetColumnName() *string {
	return s.ColumnName
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetColumnName(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ColumnName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetExampleValue(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetParameterDataType(v int32) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetParameterDescription(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetParameterName(v string) *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServiceApisResponseBodyDataApisWizardDetailsWizardResponseParameters) Validate() error {
	return dara.Validate(s)
}
