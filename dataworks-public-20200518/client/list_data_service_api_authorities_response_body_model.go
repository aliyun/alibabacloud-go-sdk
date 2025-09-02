// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiAuthoritiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataServiceApiAuthoritiesResponseBodyData) *ListDataServiceApiAuthoritiesResponseBody
	GetData() *ListDataServiceApiAuthoritiesResponseBodyData
	SetErrorCode(v string) *ListDataServiceApiAuthoritiesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataServiceApiAuthoritiesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListDataServiceApiAuthoritiesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDataServiceApiAuthoritiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceApiAuthoritiesResponseBody
	GetSuccess() *bool
}

type ListDataServiceApiAuthoritiesResponseBody struct {
	// The APIs on which other users are granted the access permissions.
	Data *ListDataServiceApiAuthoritiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServiceApiAuthoritiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiAuthoritiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiAuthoritiesResponseBody) GetData() *ListDataServiceApiAuthoritiesResponseBodyData {
	return s.Data
}

func (s *ListDataServiceApiAuthoritiesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServiceApiAuthoritiesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServiceApiAuthoritiesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceApiAuthoritiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceApiAuthoritiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceApiAuthoritiesResponseBody) SetData(v *ListDataServiceApiAuthoritiesResponseBodyData) *ListDataServiceApiAuthoritiesResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBody) SetErrorCode(v string) *ListDataServiceApiAuthoritiesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBody) SetErrorMessage(v string) *ListDataServiceApiAuthoritiesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBody) SetHttpStatusCode(v int32) *ListDataServiceApiAuthoritiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBody) SetRequestId(v string) *ListDataServiceApiAuthoritiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBody) SetSuccess(v bool) *ListDataServiceApiAuthoritiesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiAuthoritiesResponseBodyData struct {
	// The APIs on which other users are granted the access permissions.
	ApiAuthorizationList []*ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList `json:"ApiAuthorizationList,omitempty" xml:"ApiAuthorizationList,omitempty" type:"Repeated"`
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

func (s ListDataServiceApiAuthoritiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiAuthoritiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) GetApiAuthorizationList() []*ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	return s.ApiAuthorizationList
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) SetApiAuthorizationList(v []*ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) *ListDataServiceApiAuthoritiesResponseBodyData {
	s.ApiAuthorizationList = v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) SetPageNumber(v int32) *ListDataServiceApiAuthoritiesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) SetPageSize(v int32) *ListDataServiceApiAuthoritiesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) SetTotalCount(v int32) *ListDataServiceApiAuthoritiesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList struct {
	// The API ID.
	//
	// example:
	//
	// 10002
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
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
	// The status of the API. Valid values: 0 and 1. The value 0 indicates that the API is not published. The value 1 indicates that the API is published.
	//
	// example:
	//
	// 0
	ApiStatus *int32 `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// The authorization records.
	AuthorizationRecords []*ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords `json:"AuthorizationRecords,omitempty" xml:"AuthorizationRecords,omitempty" type:"Repeated"`
	// The time when the API was created.
	//
	// example:
	//
	// 2020-06-23T00:21:01+0800
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the API owner.
	//
	// example:
	//
	// 12345
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
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
	// The workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetApiPath() *string {
	return s.ApiPath
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetApiStatus() *int32 {
	return s.ApiStatus
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetAuthorizationRecords() []*ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords {
	return s.AuthorizationRecords
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetApiId(v int64) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetApiName(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetApiPath(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.ApiPath = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetApiStatus(v int32) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.ApiStatus = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetAuthorizationRecords(v []*ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.AuthorizationRecords = v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetCreatedTime(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetCreatorId(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.CreatorId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetGroupId(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.GroupId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetModifiedTime(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.ModifiedTime = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetProjectId(v int64) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) SetTenantId(v int64) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationList) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords struct {
	// The time when the access permissions on the API were granted to other users.
	//
	// example:
	//
	// 2020-06-23T00:21:01+0800
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the API owner.
	//
	// example:
	//
	// 12345
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The end time of the validity period of the authorization.
	//
	// example:
	//
	// 2020-06-24T00:21:01+0800
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the workspace to which the access permissions on the API are granted.
	//
	// example:
	//
	// 10004
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) SetCreatedTime(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) SetCreatorId(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords {
	s.CreatorId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) SetEndTime(v string) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords {
	s.EndTime = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) SetProjectId(v int64) *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceApiAuthoritiesResponseBodyDataApiAuthorizationListAuthorizationRecords) Validate() error {
	return dara.Validate(s)
}
