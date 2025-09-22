// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetChatSessionListResponseBodyData) *GetChatSessionListResponseBody
	GetData() *GetChatSessionListResponseBodyData
	SetRequestId(v string) *GetChatSessionListResponseBody
	GetRequestId() *string
}

type GetChatSessionListResponseBody struct {
	// Returned data
	Data *GetChatSessionListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChatSessionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponseBody) GetData() *GetChatSessionListResponseBodyData {
	return s.Data
}

func (s *GetChatSessionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatSessionListResponseBody) SetData(v *GetChatSessionListResponseBodyData) *GetChatSessionListResponseBody {
	s.Data = v
	return s
}

func (s *GetChatSessionListResponseBody) SetRequestId(v string) *GetChatSessionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatSessionListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetChatSessionListResponseBodyData struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 分页大小。
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Session list.
	SessionList []*GetChatSessionListResponseBodyDataSessionList `json:"sessionList,omitempty" xml:"sessionList,omitempty" type:"Repeated"`
	// Total number of records.
	//
	// example:
	//
	// 21
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetChatSessionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetChatSessionListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetChatSessionListResponseBodyData) GetSessionList() []*GetChatSessionListResponseBodyDataSessionList {
	return s.SessionList
}

func (s *GetChatSessionListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetChatSessionListResponseBodyData) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *GetChatSessionListResponseBodyData) SetCurrentPage(v int64) *GetChatSessionListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetPageSize(v int64) *GetChatSessionListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetSessionList(v []*GetChatSessionListResponseBodyDataSessionList) *GetChatSessionListResponseBodyData {
	s.SessionList = v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetTotal(v int64) *GetChatSessionListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) SetTotalPage(v int64) *GetChatSessionListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *GetChatSessionListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetChatSessionListResponseBodyDataSessionList struct {
	// Report creation timestamp, in milliseconds.
	//
	// example:
	//
	// 2025-01-01T14:45:17Z
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// Folder ID, used to specify the scope of documents to be queried.
	//
	// example:
	//
	// 3493370b-4884-47dd-95ed-49038769af53
	FolderId *string `json:"folderId,omitempty" xml:"folderId,omitempty"`
	// Session name
	//
	// example:
	//
	// student_app_spelling
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Q&A session ID, used to record multiple Q&As of the same user.
	//
	// example:
	//
	// 5c748ef9-3f23-4b5a-916f-966c0d2c6dcd
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Modification time, in milliseconds.
	//
	// example:
	//
	// 2024-12-30T02:05:03Z
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// User ID of the current session.
	//
	// example:
	//
	// 12222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetChatSessionListResponseBodyDataSessionList) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionListResponseBodyDataSessionList) GoString() string {
	return s.String()
}

func (s *GetChatSessionListResponseBodyDataSessionList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetChatSessionListResponseBodyDataSessionList) GetFolderId() *string {
	return s.FolderId
}

func (s *GetChatSessionListResponseBodyDataSessionList) GetName() *string {
	return s.Name
}

func (s *GetChatSessionListResponseBodyDataSessionList) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatSessionListResponseBodyDataSessionList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetChatSessionListResponseBodyDataSessionList) GetUserId() *string {
	return s.UserId
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetCreateTime(v int64) *GetChatSessionListResponseBodyDataSessionList {
	s.CreateTime = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetFolderId(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.FolderId = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetName(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.Name = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetSessionId(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.SessionId = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetUpdateTime(v int64) *GetChatSessionListResponseBodyDataSessionList {
	s.UpdateTime = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) SetUserId(v string) *GetChatSessionListResponseBodyDataSessionList {
	s.UserId = &v
	return s
}

func (s *GetChatSessionListResponseBodyDataSessionList) Validate() error {
	return dara.Validate(s)
}
