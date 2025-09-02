// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataServiceApplicationsResponseBodyData) *ListDataServiceApplicationsResponseBody
	GetData() *ListDataServiceApplicationsResponseBodyData
	SetErrorCode(v string) *ListDataServiceApplicationsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataServiceApplicationsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListDataServiceApplicationsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDataServiceApplicationsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceApplicationsResponseBody
	GetSuccess() *bool
}

type ListDataServiceApplicationsResponseBody struct {
	// The returned data.
	Data *ListDataServiceApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Normal
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
	// 0000-ABCD-EFG***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApplicationsResponseBody) GetData() *ListDataServiceApplicationsResponseBodyData {
	return s.Data
}

func (s *ListDataServiceApplicationsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServiceApplicationsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServiceApplicationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApplicationsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceApplicationsResponseBody) SetData(v *ListDataServiceApplicationsResponseBodyData) *ListDataServiceApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServiceApplicationsResponseBody) SetErrorCode(v string) *ListDataServiceApplicationsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBody) SetErrorMessage(v string) *ListDataServiceApplicationsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBody) SetHttpStatusCode(v int32) *ListDataServiceApplicationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBody) SetRequestId(v string) *ListDataServiceApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBody) SetSuccess(v bool) *ListDataServiceApplicationsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApplicationsResponseBodyData struct {
	// The basic information of the applications.
	Applications []*ListDataServiceApplicationsResponseBodyDataApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The page number. The value of this parameter is the same as that of the PageNumber parameter in the request.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServiceApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServiceApplicationsResponseBodyData) GetApplications() []*ListDataServiceApplicationsResponseBodyDataApplications {
	return s.Applications
}

func (s *ListDataServiceApplicationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceApplicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApplicationsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceApplicationsResponseBodyData) SetApplications(v []*ListDataServiceApplicationsResponseBodyDataApplications) *ListDataServiceApplicationsResponseBodyData {
	s.Applications = v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyData) SetPageNumber(v int32) *ListDataServiceApplicationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyData) SetPageSize(v int32) *ListDataServiceApplicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyData) SetTotalCount(v int32) *ListDataServiceApplicationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApplicationsResponseBodyDataApplications struct {
	// The application ID.
	//
	// example:
	//
	// 20000
	ApplicationId *int64 `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// My application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceApplicationsResponseBodyDataApplications) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApplicationsResponseBodyDataApplications) GoString() string {
	return s.String()
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) GetApplicationId() *int64 {
	return s.ApplicationId
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) SetApplicationId(v int64) *ListDataServiceApplicationsResponseBodyDataApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) SetApplicationName(v string) *ListDataServiceApplicationsResponseBodyDataApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) SetProjectId(v int64) *ListDataServiceApplicationsResponseBodyDataApplications {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApplicationsResponseBodyDataApplications) Validate() error {
	return dara.Validate(s)
}
