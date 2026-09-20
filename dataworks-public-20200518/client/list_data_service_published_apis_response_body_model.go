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
	// The published API information returned.
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
	// The request ID, which is the unique identifier for the request.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyData struct {
	// The list of published API information.
	Apis []*ListDataServicePublishedApisResponseBodyDataApis `json:"Apis,omitempty" xml:"Apis,omitempty" type:"Repeated"`
	// The page number, which is consistent with the PageNumber in the request.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
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
	if s.Apis != nil {
		for _, item := range s.Apis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyDataApis struct {
	// The ID of the API.
	//
	// example:
	//
	// 10002
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The type of the API. Valid values:
	//
	// - 0: wizard API.
	//
	// - 1: script API.
	//
	// - 2: registration API.
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
	// The Alibaba Cloud ID of the creator.
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
	// The Alibaba Cloud ID of the user who last edited the API.
	//
	// example:
	//
	// 2345678
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The API protocol. Valid values:
	//
	// - 0: HTTP.
	//
	// - 1: HTTPS.
	Protocols []*int32 `json:"Protocols,omitempty" xml:"Protocols,omitempty" type:"Repeated"`
	// The details of the registration API. This is returned only for registration APIs.
	RegistrationDetails *ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails `json:"RegistrationDetails,omitempty" xml:"RegistrationDetails,omitempty" type:"Struct"`
	// The request method of the API. Valid values:
	//
	// - 0: GET.
	//
	// - 1: POST.
	//
	// - 2: PUT.
	//
	// - 3: DELETE.
	//
	// Wizard and script APIs support GET and POST. Registration APIs support GET, POST, PUT, and DELETE.
	//
	// example:
	//
	// 0
	RequestMethod *int32 `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	// example:
	//
	// 0
	ResponseContentType *int32 `json:"ResponseContentType,omitempty" xml:"ResponseContentType,omitempty"`
	// The details of the script API. This is returned only for script APIs.
	ScriptDetails *ListDataServicePublishedApisResponseBodyDataApisScriptDetails `json:"ScriptDetails,omitempty" xml:"ScriptDetails,omitempty" type:"Struct"`
	// The SQL mode. Valid values: 0 (basic SQL) and 1 (advanced SQL).
	//
	// example:
	//
	// 0
	SqlMode *int32 `json:"SqlMode,omitempty" xml:"SqlMode,omitempty"`
	// The status of the API. Valid values:
	//
	// - 0: unpublished.
	//
	// - 1: published.
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
	// The timeout period, in milliseconds (ms).
	//
	// example:
	//
	// 10000
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The visibility range. Valid values:
	//
	// - 0: workspace.
	//
	// - 1: private.
	//
	// example:
	//
	// 0
	VisibleRange *int32 `json:"VisibleRange,omitempty" xml:"VisibleRange,omitempty"`
	// The details of the wizard API. This is returned only for wizard APIs.
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

func (s *ListDataServicePublishedApisResponseBodyDataApis) GetSqlMode() *int32 {
	return s.SqlMode
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

func (s *ListDataServicePublishedApisResponseBodyDataApis) SetSqlMode(v int32) *ListDataServicePublishedApisResponseBodyDataApis {
	s.SqlMode = &v
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
	if s.RegistrationDetails != nil {
		if err := s.RegistrationDetails.Validate(); err != nil {
			return err
		}
	}
	if s.ScriptDetails != nil {
		if err := s.ScriptDetails.Validate(); err != nil {
			return err
		}
	}
	if s.WizardDetails != nil {
		if err := s.WizardDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyDataApisRegistrationDetails struct {
	// The sample of an error response.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
	// The list of error codes for the registration API.
	RegistrationErrorCodes []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationErrorCodes `json:"RegistrationErrorCodes,omitempty" xml:"RegistrationErrorCodes,omitempty" type:"Repeated"`
	// The list of request parameters for the registration API.
	RegistrationRequestParameters []*ListDataServicePublishedApisResponseBodyDataApisRegistrationDetailsRegistrationRequestParameters `json:"RegistrationRequestParameters,omitempty" xml:"RegistrationRequestParameters,omitempty" type:"Repeated"`
	// The return data type of the API. Valid values:
	//
	// - 0: JSON.
	//
	// - 1: XML.
	//
	// Wizard and script APIs support JSON. Registration APIs support JSON and XML.
	//
	// example:
	//
	// 0
	ServiceContentType *int32 `json:"ServiceContentType,omitempty" xml:"ServiceContentType,omitempty"`
	// The backend service address.
	//
	// example:
	//
	// http://www.abc.com
	ServiceHost *string `json:"ServiceHost,omitempty" xml:"ServiceHost,omitempty"`
	// The backend service path.
	//
	// example:
	//
	// /index
	ServicePath *string `json:"ServicePath,omitempty" xml:"ServicePath,omitempty"`
	// The description of the backend request body content.
	//
	// example:
	//
	// {"abc":1}
	ServiceRequestBodyDescription *string `json:"ServiceRequestBodyDescription,omitempty" xml:"ServiceRequestBodyDescription,omitempty"`
	// The sample of a successful response.
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
	if s.RegistrationErrorCodes != nil {
		for _, item := range s.RegistrationErrorCodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RegistrationRequestParameters != nil {
		for _, item := range s.RegistrationRequestParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
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
	// The fault Solutions.
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
	// The example value.
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
	// The data type. Valid values:
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
	// The parameter name.
	//
	// example:
	//
	// name1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The operator. Valid values:
	//
	// example:
	//
	// 0
	ParameterOperator *int32 `json:"ParameterOperator,omitempty" xml:"ParameterOperator,omitempty"`
	// The parameter position. Valid values:
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
	// The sample of an error response.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
	// Indicates whether the response is paginated.
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
	// The data source information for the script API.
	ScriptConnection *ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection `json:"ScriptConnection,omitempty" xml:"ScriptConnection,omitempty" type:"Struct"`
	// The list of error codes for the script API.
	ScriptErrorCodes []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptErrorCodes `json:"ScriptErrorCodes,omitempty" xml:"ScriptErrorCodes,omitempty" type:"Repeated"`
	// The list of request parameters for the script API.
	ScriptRequestParameters []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptRequestParameters `json:"ScriptRequestParameters,omitempty" xml:"ScriptRequestParameters,omitempty" type:"Repeated"`
	// The list of response parameters for the script API.
	ScriptResponseParameters []*ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptResponseParameters `json:"ScriptResponseParameters,omitempty" xml:"ScriptResponseParameters,omitempty" type:"Repeated"`
	// The sample of a successful response.
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
	if s.ScriptConnection != nil {
		if err := s.ScriptConnection.Validate(); err != nil {
			return err
		}
	}
	if s.ScriptErrorCodes != nil {
		for _, item := range s.ScriptErrorCodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScriptRequestParameters != nil {
		for _, item := range s.ScriptRequestParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScriptResponseParameters != nil {
		for _, item := range s.ScriptResponseParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyDataApisScriptDetailsScriptConnection struct {
	// The data source ID.
	//
	// example:
	//
	// 123
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The table name of the data source.
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
	// The fault Solutions.
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
	// The example value.
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
	// The data type. Valid values:
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
	// The parameter name.
	//
	// example:
	//
	// param1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The operator. Valid values:
	//
	// example:
	//
	// 0
	ParameterOperator *int32 `json:"ParameterOperator,omitempty" xml:"ParameterOperator,omitempty"`
	// The parameter position. Valid values:
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
	// The example value.
	//
	// example:
	//
	// example2
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The data type. Valid values:
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
	// The parameter name.
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
	// The sample of an error response.
	//
	// example:
	//
	// {"success": false}
	FailedResultSample *string `json:"FailedResultSample,omitempty" xml:"FailedResultSample,omitempty"`
	// Indicates whether the response is paginated.
	//
	// example:
	//
	// true
	IsPagedResponse *bool `json:"IsPagedResponse,omitempty" xml:"IsPagedResponse,omitempty"`
	// The sample of a successful response.
	//
	// example:
	//
	// {"success": true}
	SuccessfulResultSample *string `json:"SuccessfulResultSample,omitempty" xml:"SuccessfulResultSample,omitempty"`
	// The data source information of the wizard API.
	WizardConnection *ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection `json:"WizardConnection,omitempty" xml:"WizardConnection,omitempty" type:"Struct"`
	// The list of error codes for the wizard API.
	WizardErrorCodes []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardErrorCodes `json:"WizardErrorCodes,omitempty" xml:"WizardErrorCodes,omitempty" type:"Repeated"`
	// The list of request parameters for the wizard API.
	WizardRequestParameters []*ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardRequestParameters `json:"WizardRequestParameters,omitempty" xml:"WizardRequestParameters,omitempty" type:"Repeated"`
	// The list of response parameters for the wizard API.
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
	if s.WizardConnection != nil {
		if err := s.WizardConnection.Validate(); err != nil {
			return err
		}
	}
	if s.WizardErrorCodes != nil {
		for _, item := range s.WizardErrorCodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WizardRequestParameters != nil {
		for _, item := range s.WizardRequestParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WizardResponseParameters != nil {
		for _, item := range s.WizardResponseParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyDataApisWizardDetailsWizardConnection struct {
	// The data source ID.
	//
	// example:
	//
	// 12354
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The table name of the data source.
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
	// The fault Solutions.
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
	// The example value.
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
	// The data type. Valid values:
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
	// The operator. Valid values: 0 (Equal), 1 (Like), 2 (Const), and 3 (In). Wizard mode APIs support Equal, Like, and In. Script mode APIs support Equal. Registered APIs support Equal and Const.
	//
	// example:
	//
	// 0
	ParameterOperator *int32 `json:"ParameterOperator,omitempty" xml:"ParameterOperator,omitempty"`
	// The position of the parameter. Valid values: 0 (Path), 1 (Query), 2 (Head), and 3 (Body). Wizard and script APIs support only Query. For registered APIs, the GET and DELETE methods support Query and Head, and the PUT and POST methods support Query, Head, and Body.
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
	// The example value.
	//
	// example:
	//
	// example2
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The data type. Valid values:
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
