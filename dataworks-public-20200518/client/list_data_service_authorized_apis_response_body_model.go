// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceAuthorizedApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataServiceAuthorizedApisResponseBodyData) *ListDataServiceAuthorizedApisResponseBody
	GetData() *ListDataServiceAuthorizedApisResponseBodyData
	SetErrorCode(v string) *ListDataServiceAuthorizedApisResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataServiceAuthorizedApisResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListDataServiceAuthorizedApisResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDataServiceAuthorizedApisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServiceAuthorizedApisResponseBody
	GetSuccess() *bool
}

type ListDataServiceAuthorizedApisResponseBody struct {
	// The information about the APIs that you are authorized to access.
	Data *ListDataServiceAuthorizedApisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
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

func (s ListDataServiceAuthorizedApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedApisResponseBody) GetData() *ListDataServiceAuthorizedApisResponseBodyData {
	return s.Data
}

func (s *ListDataServiceAuthorizedApisResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataServiceAuthorizedApisResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataServiceAuthorizedApisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServiceAuthorizedApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServiceAuthorizedApisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServiceAuthorizedApisResponseBody) SetData(v *ListDataServiceAuthorizedApisResponseBodyData) *ListDataServiceAuthorizedApisResponseBody {
	s.Data = v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBody) SetErrorCode(v string) *ListDataServiceAuthorizedApisResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBody) SetErrorMessage(v string) *ListDataServiceAuthorizedApisResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBody) SetHttpStatusCode(v int32) *ListDataServiceAuthorizedApisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBody) SetRequestId(v string) *ListDataServiceAuthorizedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBody) SetSuccess(v bool) *ListDataServiceAuthorizedApisResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceAuthorizedApisResponseBodyData struct {
	// The APIs that you are authorized to access.
	ApiAuthorizedList []*ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList `json:"ApiAuthorizedList,omitempty" xml:"ApiAuthorizedList,omitempty" type:"Repeated"`
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

func (s ListDataServiceAuthorizedApisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) GetApiAuthorizedList() []*ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	return s.ApiAuthorizedList
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) SetApiAuthorizedList(v []*ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) *ListDataServiceAuthorizedApisResponseBodyData {
	s.ApiAuthorizedList = v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) SetPageNumber(v int32) *ListDataServiceAuthorizedApisResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) SetPageSize(v int32) *ListDataServiceAuthorizedApisResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) SetTotalCount(v int32) *ListDataServiceAuthorizedApisResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList struct {
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
	// My API Name
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
	// The time when the access permissions on the API were granted.
	//
	// example:
	//
	// 2020-06-23T00:21:01+0800
	GrantCreatedTime *string `json:"GrantCreatedTime,omitempty" xml:"GrantCreatedTime,omitempty"`
	// The expiration time of the access permissions granted on the API.
	//
	// example:
	//
	// 2020-06-24T00:21:01+0800
	GrantEndTime *string `json:"GrantEndTime,omitempty" xml:"GrantEndTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who granted the access permissions on the API.
	//
	// example:
	//
	// 23456
	GrantOperatorId *string `json:"GrantOperatorId,omitempty" xml:"GrantOperatorId,omitempty"`
	// The group ID.
	//
	// example:
	//
	// abcde123456789
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The time when the API was last updated.
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

func (s ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GoString() string {
	return s.String()
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetApiPath() *string {
	return s.ApiPath
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetApiStatus() *int32 {
	return s.ApiStatus
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetGrantCreatedTime() *string {
	return s.GrantCreatedTime
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetGrantEndTime() *string {
	return s.GrantEndTime
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetGrantOperatorId() *string {
	return s.GrantOperatorId
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetApiId(v int64) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetApiName(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.ApiName = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetApiPath(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.ApiPath = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetApiStatus(v int32) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.ApiStatus = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetCreatedTime(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.CreatedTime = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetCreatorId(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.CreatorId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetGrantCreatedTime(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.GrantCreatedTime = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetGrantEndTime(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.GrantEndTime = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetGrantOperatorId(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.GrantOperatorId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetGroupId(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.GroupId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetModifiedTime(v string) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.ModifiedTime = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetProjectId(v int64) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) SetTenantId(v int64) *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceAuthorizedApisResponseBodyDataApiAuthorizedList) Validate() error {
	return dara.Validate(s)
}
