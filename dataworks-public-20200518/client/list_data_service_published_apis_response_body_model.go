// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServicePublishedApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataServicePublishedApisResponseBodyData) *ListDataServicePublishedApisResponseBody
	GetData() *ListDataServicePublishedApisResponseBodyData
	SetErrorCode(v string) *ListDataServicePublishedApisResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataServicePublishedApisResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListDataServicePublishedApisResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDataServicePublishedApisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServicePublishedApisResponseBody
	GetSuccess() *bool
}

type ListDataServicePublishedApisResponseBody struct {
	// The data returned.
	Data *ListDataServicePublishedApisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListDataServicePublishedApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBody) GetData() *ListDataServicePublishedApisResponseBodyData {
	return s.Data
}

func (s *ListDataServicePublishedApisResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServicePublishedApisResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServicePublishedApisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServicePublishedApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServicePublishedApisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServicePublishedApisResponseBody) SetData(v *ListDataServicePublishedApisResponseBodyData) *ListDataServicePublishedApisResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetErrorCode(v string) *ListDataServicePublishedApisResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetErrorMessage(v string) *ListDataServicePublishedApisResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetHttpStatusCode(v int32) *ListDataServicePublishedApisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetRequestId(v string) *ListDataServicePublishedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetSuccess(v bool) *ListDataServicePublishedApisResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyData struct {
	// The information about the APIs in the published state.
	Apis []*ListDataServicePublishedApisResponseBodyDataApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	// The page number. The value of this parameter is the same as that of the PageNumber parameter in the request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 10.
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

func (s ListDataServicePublishedApisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyData) GetApis() []*ListDataServicePublishedApisResponseBodyDataApis {
	return s.Apis
}

func (s *ListDataServicePublishedApisResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServicePublishedApisResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServicePublishedApisResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServicePublishedApisResponseBodyData) SetApis(v []*ListDataServicePublishedApisResponseBodyDataApis) *ListDataServicePublishedApisResponseBodyData {
	s.Apis = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyData) SetPageNumber(v int32) *ListDataServicePublishedApisResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyData) SetPageSize(v int32) *ListDataServicePublishedApisResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyData) SetTotalCount(v int32) *ListDataServicePublishedApisResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApis struct {
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
	// The description.
	//
	// example:
	//
	// Test API description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group ID.
	//
	// example:
	//
	// ab123
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
	// The protocol used by the API. Valid values: 0 and 1. The value 0 indicates HTTP. The value 1 indicates HTTPS.
	Protocols []*int32 `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The details of the API generated by registration. This parameter is returned only if the API is generated by registration.
	RegistrationDetails *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails `json:"RegistrationDetails,omitempty" xml:"RegistrationDetails,omitempty" type:"Struct"`
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
	ScriptDetails *ListDataServicePublishedApisResponseBodyDataApisScriptDetails `json:"ScriptDetails,omitempty" xml:"ScriptDetails,omitempty" type:"Struct"`
	// The status of the API. Valid values: 0 and 1. The value 0 indicates that the API is not published. The value 1 indicates that the API is published.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
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
	WizardDetails *ListDataServicePublishedApisResponseBodyDataApisWizardDetails `json:"WizardDetails,omitempty" xml:"WizardDetails,omitempty" type:"Struct"`
}

func (s ListDataServicePublishedApisResponseBodyDataApis) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApis) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetApiMode() *int32 {
	return s.ApiMode
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetApiPath() *string {
	return s.ApiPath
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetDescription() *string {
	return s.Description
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetOperatorId() *string {
	return s.OperatorId
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetProtocols() []*int32 {
	return s.Protocols
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetRegistrationDetails() *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	return s.RegistrationDetails
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetRequestMethod() *int32 {
	return s.RequestMethod
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetResponseContentType() *int32 {
	return s.ResponseContentType
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetScriptDetails() *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	return s.ScriptDetails
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetVisibleRange() *int32 {
	return s.VisibleRange
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetWizardDetails() *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	return s.WizardDetails
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetApiId(v int64) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ApiId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetApiMode(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ApiMode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetApiName(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ApiName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetApiPath(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ApiPath = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetCreatedTime(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetCreatorId(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.CreatorId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetDescription(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.Description = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetGroupId(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.GroupId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetModifiedTime(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ModifiedTime = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetOperatorId(v string) *ListDataServicePublishedApisResponseBodyDataApis {
	s.OperatorId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetProjectId(v int64) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ProjectId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetProtocols(v []*int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.Protocols = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetRegistrationDetails(v *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) *ListDataServicePublishedApisResponseBodyDataApis {
	s.RegistrationDetails = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetRequestMethod(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.RequestMethod = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetResponseContentType(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ResponseContentType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetScriptDetails(v *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) *ListDataServicePublishedApisResponseBodyDataApis {
	s.ScriptDetails = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetStatus(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.Status = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetTenantId(v int64) *ListDataServicePublishedApisResponseBodyDataApis {
	s.TenantId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetTimeout(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.Timeout = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetVisibleRange(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.VisibleRange = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetWizardDetails(v *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) *ListDataServicePublishedApisResponseBodyDataApis {
	s.WizardDetails = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApis) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails struct {
	// The sample error response of the API.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
	// The error codes returned for the API generated by registration.
	RegistrationErrorCodes []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes `json:"RegistrationErrorCodes,omitempty" xml:"RegistrationErrorCodes,omitempty" type:"Repeated"`
	// The request parameters of the API generated by registration.
	RegistrationRequestParameters []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters `json:"RegistrationRequestParameters,omitempty" xml:"RegistrationRequestParameters,omitempty" type:"Repeated"`
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

func (s ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetFailedResultSample() *string {
	return s.FailedResultSample
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetRegistrationErrorCodes() []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	return s.RegistrationErrorCodes
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetRegistrationRequestParameters() []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	return s.RegistrationRequestParameters
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetServiceContentType() *int32 {
	return s.ServiceContentType
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetServiceHost() *string {
	return s.ServiceHost
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetServicePath() *string {
	return s.ServicePath
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetServiceRequestBodyDescription() *string {
	return s.ServiceRequestBodyDescription
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) GetSuccessfulResultSample() *string {
	return s.SuccessfulResultSample
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetFailedResultSample(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.FailedResultSample = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetRegistrationErrorCodes(v []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.RegistrationErrorCodes = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetRegistrationRequestParameters(v []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.RegistrationRequestParameters = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetServiceContentType(v int32) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.ServiceContentType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetServiceHost(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.ServiceHost = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetServicePath(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.ServicePath = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetServiceRequestBodyDescription(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.ServiceRequestBodyDescription = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) SetSuccessfulResultSample(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails {
	s.SuccessfulResultSample = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes struct {
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

func (s ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) GetErrorSolution() *string {
	return s.ErrorSolution
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) SetErrorCode(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) SetErrorMessage(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) SetErrorSolution(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes {
	s.ErrorSolution = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters struct {
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
	// Indicates whether the request parameter is required.
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

func (s ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterOperator() *int32 {
	return s.ParameterOperator
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) GetParameterPosition() *int32 {
	return s.ParameterPosition
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetDefaultValue(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetExampleValue(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetIsRequiredParameter(v bool) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterDataType(v int32) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterDescription(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterName(v string) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterOperator(v int32) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterOperator = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) SetParameterPosition(v int32) *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters {
	s.ParameterPosition = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisScriptDetails struct {
	// The sample error response of the API.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
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
	ScriptConnection *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection `json:"ScriptConnection,omitempty" xml:"ScriptConnection,omitempty" type:"Struct"`
	// The error codes returned for the API generated in script mode.
	ScriptErrorCodes []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes `json:"ScriptErrorCodes,omitempty" xml:"ScriptErrorCodes,omitempty" type:"Repeated"`
	// The request parameters of the API generated in script mode.
	ScriptRequestParameters []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters `json:"ScriptRequestParameters,omitempty" xml:"ScriptRequestParameters,omitempty" type:"Repeated"`
	// The response parameters of the API generated in script mode.
	ScriptResponseParameters []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters `json:"ScriptResponseParameters,omitempty" xml:"ScriptResponseParameters,omitempty" type:"Repeated"`
	// The sample success response of the API.
	//
	// example:
	//
	// {"success": true}
	SuccessfulResultSample *string `json:"SuccessfulResultSample,omitempty" xml:"SuccessfulResultSample,omitempty"`
}

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetFailedResultSample() *string {
	return s.FailedResultSample
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetIsPagedResponse() *bool {
	return s.IsPagedResponse
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetScript() *string {
	return s.Script
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetScriptConnection() *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection {
	return s.ScriptConnection
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetScriptErrorCodes() []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes {
	return s.ScriptErrorCodes
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetScriptRequestParameters() []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	return s.ScriptRequestParameters
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetScriptResponseParameters() []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	return s.ScriptResponseParameters
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) GetSuccessfulResultSample() *string {
	return s.SuccessfulResultSample
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetFailedResultSample(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.FailedResultSample = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetIsPagedResponse(v bool) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.IsPagedResponse = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetScript(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.Script = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetScriptConnection(v *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.ScriptConnection = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetScriptErrorCodes(v []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.ScriptErrorCodes = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetScriptRequestParameters(v []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.ScriptRequestParameters = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetScriptResponseParameters(v []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.ScriptResponseParameters = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) SetSuccessfulResultSample(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetails {
	s.SuccessfulResultSample = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection struct {
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

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) GetTableName() *string {
	return s.TableName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) SetConnectionId(v int64) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection {
	s.ConnectionId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) SetTableName(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection {
	s.TableName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes struct {
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

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) GetErrorSolution() *string {
	return s.ErrorSolution
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) SetErrorCode(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) SetErrorMessage(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) SetErrorSolution(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes {
	s.ErrorSolution = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters struct {
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
	// Indicates whether the request parameter is required.
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
	// Advanced scripts also support the following data types:
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

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterOperator() *int32 {
	return s.ParameterOperator
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) GetParameterPosition() *int32 {
	return s.ParameterPosition
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetDefaultValue(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetExampleValue(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetIsRequiredParameter(v bool) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterDataType(v int32) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterDescription(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterName(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterOperator(v int32) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterOperator = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) SetParameterPosition(v int32) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters {
	s.ParameterPosition = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters struct {
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

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetExampleValue(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetParameterDataType(v int32) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetParameterDescription(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) SetParameterName(v string) *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisWizardDetails struct {
	// The sample error response of the API.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
	// Indicates whether the entries are returned by page.
	//
	// example:
	//
	// true
	IsPagedResponse *bool `json:"IsPagedResponse,omitempty" xml:"IsPagedResponse,omitempty"`
	// The sample success response of the API.
	//
	// example:
	//
	// {"success": true}
	SuccessfulResultSample *string `json:"SuccessfulResultSample,omitempty" xml:"SuccessfulResultSample,omitempty"`
	// The data source information about the API generated in wizard mode.
	WizardConnection *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection `json:"WizardConnection,omitempty" xml:"WizardConnection,omitempty" type:"Struct"`
	// The error codes returned for the API generated in wizard mode.
	WizardErrorCodes []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes `json:"WizardErrorCodes,omitempty" xml:"WizardErrorCodes,omitempty" type:"Repeated"`
	// The request parameters of the API generated in wizard mode.
	WizardRequestParameters []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters `json:"WizardRequestParameters,omitempty" xml:"WizardRequestParameters,omitempty" type:"Repeated"`
	// The response parameters of the API generated in wizard mode.
	WizardResponseParameters []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters `json:"WizardResponseParameters,omitempty" xml:"WizardResponseParameters,omitempty" type:"Repeated"`
}

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetFailedResultSample() *string {
	return s.FailedResultSample
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetIsPagedResponse() *bool {
	return s.IsPagedResponse
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetSuccessfulResultSample() *string {
	return s.SuccessfulResultSample
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetWizardConnection() *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection {
	return s.WizardConnection
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetWizardErrorCodes() []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes {
	return s.WizardErrorCodes
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetWizardRequestParameters() []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	return s.WizardRequestParameters
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) GetWizardResponseParameters() []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	return s.WizardResponseParameters
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetFailedResultSample(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.FailedResultSample = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetIsPagedResponse(v bool) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.IsPagedResponse = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetSuccessfulResultSample(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.SuccessfulResultSample = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetWizardConnection(v *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.WizardConnection = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetWizardErrorCodes(v []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.WizardErrorCodes = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetWizardRequestParameters(v []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.WizardRequestParameters = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) SetWizardResponseParameters(v []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) *ListDataServicePublishedApisResponseBodyDataApisWizardDetails {
	s.WizardResponseParameters = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection struct {
	// The ID of the data source.
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

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) GetTableName() *string {
	return s.TableName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) SetConnectionId(v int64) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection {
	s.ConnectionId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) SetTableName(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection {
	s.TableName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes struct {
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

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) GetErrorSolution() *string {
	return s.ErrorSolution
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) SetErrorCode(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) SetErrorMessage(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) SetErrorSolution(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes {
	s.ErrorSolution = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters struct {
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
	// Indicates whether the request parameter is required.
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

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetIsRequiredParameter() *bool {
	return s.IsRequiredParameter
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterOperator() *int32 {
	return s.ParameterOperator
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) GetParameterPosition() *int32 {
	return s.ParameterPosition
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetDefaultValue(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.DefaultValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetExampleValue(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetIsRequiredParameter(v bool) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.IsRequiredParameter = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterDataType(v int32) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterDescription(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterName(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterOperator(v int32) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterOperator = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) SetParameterPosition(v int32) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters {
	s.ParameterPosition = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters) Validate() error {
	return dara.Validate(s)
}

type ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters struct {
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

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetExampleValue(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetParameterDataType(v int32) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetParameterDescription(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) SetParameterName(v string) *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardResponseParameters) Validate() error {
	return dara.Validate(s)
}
