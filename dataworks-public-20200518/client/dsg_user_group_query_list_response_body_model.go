// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupQueryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DsgUserGroupQueryListResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgUserGroupQueryListResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgUserGroupQueryListResponseBody
	GetHttpStatusCode() *int32
	SetPageData(v *DsgUserGroupQueryListResponseBodyPageData) *DsgUserGroupQueryListResponseBody
	GetPageData() *DsgUserGroupQueryListResponseBodyPageData
	SetRequestId(v string) *DsgUserGroupQueryListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgUserGroupQueryListResponseBody
	GetSuccess() *bool
}

type DsgUserGroupQueryListResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The pagination information.
	PageData *DsgUserGroupQueryListResponseBodyPageData `json:"PageData,omitempty" xml:"PageData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgUserGroupQueryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryListResponseBody) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgUserGroupQueryListResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgUserGroupQueryListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgUserGroupQueryListResponseBody) GetPageData() *DsgUserGroupQueryListResponseBodyPageData {
	return s.PageData
}

func (s *DsgUserGroupQueryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgUserGroupQueryListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgUserGroupQueryListResponseBody) SetErrorCode(v string) *DsgUserGroupQueryListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBody) SetErrorMessage(v string) *DsgUserGroupQueryListResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBody) SetHttpStatusCode(v int32) *DsgUserGroupQueryListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBody) SetPageData(v *DsgUserGroupQueryListResponseBodyPageData) *DsgUserGroupQueryListResponseBody {
	s.PageData = v
	return s
}

func (s *DsgUserGroupQueryListResponseBody) SetRequestId(v string) *DsgUserGroupQueryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBody) SetSuccess(v bool) *DsgUserGroupQueryListResponseBody {
	s.Success = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DsgUserGroupQueryListResponseBodyPageData struct {
	// The user groups.
	Data []*DsgUserGroupQueryListResponseBodyPageDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of user groups returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DsgUserGroupQueryListResponseBodyPageData) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryListResponseBodyPageData) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryListResponseBodyPageData) GetData() []*DsgUserGroupQueryListResponseBodyPageDataData {
	return s.Data
}

func (s *DsgUserGroupQueryListResponseBodyPageData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgUserGroupQueryListResponseBodyPageData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgUserGroupQueryListResponseBodyPageData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DsgUserGroupQueryListResponseBodyPageData) SetData(v []*DsgUserGroupQueryListResponseBodyPageDataData) *DsgUserGroupQueryListResponseBodyPageData {
	s.Data = v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageData) SetPageNumber(v int32) *DsgUserGroupQueryListResponseBodyPageData {
	s.PageNumber = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageData) SetPageSize(v int32) *DsgUserGroupQueryListResponseBodyPageData {
	s.PageSize = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageData) SetTotalCount(v int32) *DsgUserGroupQueryListResponseBodyPageData {
	s.TotalCount = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageData) Validate() error {
	return dara.Validate(s)
}

type DsgUserGroupQueryListResponseBodyPageDataData struct {
	// The usernames in the user group.
	Accounts []*string `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// The time when the user group was created.
	//
	// example:
	//
	// 2024-05-10 17:14:44
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the user group was modified.
	//
	// example:
	//
	// 2024-05-10 17:14:44
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// 123
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// test_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the user group.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s DsgUserGroupQueryListResponseBodyPageDataData) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryListResponseBodyPageDataData) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) GetAccounts() []*string {
	return s.Accounts
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) GetId() *int32 {
	return s.Id
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) GetName() *string {
	return s.Name
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) GetOwner() *string {
	return s.Owner
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) SetAccounts(v []*string) *DsgUserGroupQueryListResponseBodyPageDataData {
	s.Accounts = v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) SetGmtCreate(v string) *DsgUserGroupQueryListResponseBodyPageDataData {
	s.GmtCreate = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) SetGmtModified(v string) *DsgUserGroupQueryListResponseBodyPageDataData {
	s.GmtModified = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) SetId(v int32) *DsgUserGroupQueryListResponseBodyPageDataData {
	s.Id = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) SetName(v string) *DsgUserGroupQueryListResponseBodyPageDataData {
	s.Name = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) SetOwner(v string) *DsgUserGroupQueryListResponseBodyPageDataData {
	s.Owner = &v
	return s
}

func (s *DsgUserGroupQueryListResponseBodyPageDataData) Validate() error {
	return dara.Validate(s)
}
