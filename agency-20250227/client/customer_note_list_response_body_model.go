// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerNoteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CustomerNoteListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CustomerNoteListResponseBody
	GetCode() *string
	SetData(v *CustomerNoteListResponseBodyData) *CustomerNoteListResponseBody
	GetData() *CustomerNoteListResponseBodyData
	SetHttpStatusCode(v int32) *CustomerNoteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CustomerNoteListResponseBody
	GetMessage() *string
	SetMsg(v string) *CustomerNoteListResponseBody
	GetMsg() *string
	SetRequestId(v string) *CustomerNoteListResponseBody
	GetRequestId() *string
}

type CustomerNoteListResponseBody struct {
	// The details of the access denied error returned by the POP API when the caller does not have the required RAM permissions.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The result code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CustomerNoteListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned by the POP API.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The prompt message, same as message.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CustomerNoteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerNoteListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CustomerNoteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerNoteListResponseBody) GetData() *CustomerNoteListResponseBodyData {
	return s.Data
}

func (s *CustomerNoteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CustomerNoteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CustomerNoteListResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CustomerNoteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerNoteListResponseBody) SetAccessDeniedDetail(v string) *CustomerNoteListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CustomerNoteListResponseBody) SetCode(v string) *CustomerNoteListResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerNoteListResponseBody) SetData(v *CustomerNoteListResponseBodyData) *CustomerNoteListResponseBody {
	s.Data = v
	return s
}

func (s *CustomerNoteListResponseBody) SetHttpStatusCode(v int32) *CustomerNoteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CustomerNoteListResponseBody) SetMessage(v string) *CustomerNoteListResponseBody {
	s.Message = &v
	return s
}

func (s *CustomerNoteListResponseBody) SetMsg(v string) *CustomerNoteListResponseBody {
	s.Msg = &v
	return s
}

func (s *CustomerNoteListResponseBody) SetRequestId(v string) *CustomerNoteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerNoteListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CustomerNoteListResponseBodyData struct {
	// The details of the access denied error returned by the POP API when the caller does not have the required RAM permissions.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The result code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*CustomerNoteListResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned by the POP API.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The prompt message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The prompt message, same as message.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The pagination information compatible with the legacy interface (deprecated).
	PageInfo *CustomerNoteListResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 16
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CustomerNoteListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomerNoteListResponseBodyData) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CustomerNoteListResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *CustomerNoteListResponseBodyData) GetData() []*CustomerNoteListResponseBodyDataData {
	return s.Data
}

func (s *CustomerNoteListResponseBodyData) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CustomerNoteListResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *CustomerNoteListResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *CustomerNoteListResponseBodyData) GetPageInfo() *CustomerNoteListResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *CustomerNoteListResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CustomerNoteListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CustomerNoteListResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerNoteListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *CustomerNoteListResponseBodyData) SetAccessDeniedDetail(v string) *CustomerNoteListResponseBodyData {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetCode(v string) *CustomerNoteListResponseBodyData {
	s.Code = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetData(v []*CustomerNoteListResponseBodyDataData) *CustomerNoteListResponseBodyData {
	s.Data = v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetHttpStatusCode(v int32) *CustomerNoteListResponseBodyData {
	s.HttpStatusCode = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetMessage(v string) *CustomerNoteListResponseBodyData {
	s.Message = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetMsg(v string) *CustomerNoteListResponseBodyData {
	s.Msg = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetPageInfo(v *CustomerNoteListResponseBodyDataPageInfo) *CustomerNoteListResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetPageNo(v int32) *CustomerNoteListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetPageSize(v int32) *CustomerNoteListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetRequestId(v string) *CustomerNoteListResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) SetTotal(v int32) *CustomerNoteListResponseBodyData {
	s.Total = &v
	return s
}

func (s *CustomerNoteListResponseBodyData) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CustomerNoteListResponseBodyDataData struct {
	// The name of the contact.
	//
	// example:
	//
	// 张三
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The UID of the creator.
	//
	// example:
	//
	// 1757916424103619
	Creator *int64 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The logon name of the creator.
	//
	// example:
	//
	// 张三
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The creation time in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-07-17 12:18:23
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The content of the note.
	//
	// example:
	//
	// 讨论技术方案
	NoteContent *string `json:"NoteContent,omitempty" xml:"NoteContent,omitempty"`
	// The ID of the note.
	//
	// example:
	//
	// 2932252
	NoteId *int64 `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
	// The type of the note (CUSTOMER).
	//
	// example:
	//
	// CUSTOMER
	NoteType *string `json:"NoteType,omitempty" xml:"NoteType,omitempty"`
	// The label of the note type.
	//
	// example:
	//
	// 客户
	NoteTypeLabel *string `json:"NoteTypeLabel,omitempty" xml:"NoteTypeLabel,omitempty"`
	// The touch date (timestamp).
	//
	// example:
	//
	// 1784266662000
	TouchDate *string `json:"TouchDate,omitempty" xml:"TouchDate,omitempty"`
}

func (s CustomerNoteListResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *CustomerNoteListResponseBodyDataData) GetContactName() *string {
	return s.ContactName
}

func (s *CustomerNoteListResponseBodyDataData) GetCreator() *int64 {
	return s.Creator
}

func (s *CustomerNoteListResponseBodyDataData) GetCreatorName() *string {
	return s.CreatorName
}

func (s *CustomerNoteListResponseBodyDataData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CustomerNoteListResponseBodyDataData) GetNoteContent() *string {
	return s.NoteContent
}

func (s *CustomerNoteListResponseBodyDataData) GetNoteId() *int64 {
	return s.NoteId
}

func (s *CustomerNoteListResponseBodyDataData) GetNoteType() *string {
	return s.NoteType
}

func (s *CustomerNoteListResponseBodyDataData) GetNoteTypeLabel() *string {
	return s.NoteTypeLabel
}

func (s *CustomerNoteListResponseBodyDataData) GetTouchDate() *string {
	return s.TouchDate
}

func (s *CustomerNoteListResponseBodyDataData) SetContactName(v string) *CustomerNoteListResponseBodyDataData {
	s.ContactName = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetCreator(v int64) *CustomerNoteListResponseBodyDataData {
	s.Creator = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetCreatorName(v string) *CustomerNoteListResponseBodyDataData {
	s.CreatorName = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetGmtCreate(v string) *CustomerNoteListResponseBodyDataData {
	s.GmtCreate = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetNoteContent(v string) *CustomerNoteListResponseBodyDataData {
	s.NoteContent = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetNoteId(v int64) *CustomerNoteListResponseBodyDataData {
	s.NoteId = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetNoteType(v string) *CustomerNoteListResponseBodyDataData {
	s.NoteType = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetNoteTypeLabel(v string) *CustomerNoteListResponseBodyDataData {
	s.NoteTypeLabel = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) SetTouchDate(v string) *CustomerNoteListResponseBodyDataData {
	s.TouchDate = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

type CustomerNoteListResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CustomerNoteListResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s CustomerNoteListResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *CustomerNoteListResponseBodyDataPageInfo) GetPage() *int32 {
	return s.Page
}

func (s *CustomerNoteListResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CustomerNoteListResponseBodyDataPageInfo) GetTotal() *int32 {
	return s.Total
}

func (s *CustomerNoteListResponseBodyDataPageInfo) SetPage(v int32) *CustomerNoteListResponseBodyDataPageInfo {
	s.Page = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataPageInfo) SetPageSize(v int32) *CustomerNoteListResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataPageInfo) SetTotal(v int32) *CustomerNoteListResponseBodyDataPageInfo {
	s.Total = &v
	return s
}

func (s *CustomerNoteListResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}
