// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketNotesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListTicketNotesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListTicketNotesResponseBody
	GetCode() *int32
	SetData(v *ListTicketNotesResponseBodyData) *ListTicketNotesResponseBody
	GetData() *ListTicketNotesResponseBodyData
	SetMessage(v string) *ListTicketNotesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTicketNotesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTicketNotesResponseBody
	GetSuccess() *bool
}

type ListTicketNotesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *int32                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListTicketNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CA6204AC-6AA9-4CFA-9310-7DFD20C19EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTicketNotesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListTicketNotesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTicketNotesResponseBody) GetData() *ListTicketNotesResponseBodyData {
	return s.Data
}

func (s *ListTicketNotesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketNotesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketNotesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTicketNotesResponseBody) SetAccessDeniedDetail(v string) *ListTicketNotesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetCode(v int32) *ListTicketNotesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetData(v *ListTicketNotesResponseBodyData) *ListTicketNotesResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketNotesResponseBody) SetMessage(v string) *ListTicketNotesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetRequestId(v string) *ListTicketNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketNotesResponseBody) SetSuccess(v bool) *ListTicketNotesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTicketNotesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketNotesResponseBodyData struct {
	List []*ListTicketNotesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListTicketNotesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyData) GetList() []*ListTicketNotesResponseBodyDataList {
	return s.List
}

func (s *ListTicketNotesResponseBodyData) SetList(v []*ListTicketNotesResponseBodyDataList) *ListTicketNotesResponseBodyData {
	s.List = v
	return s
}

func (s *ListTicketNotesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketNotesResponseBodyDataList struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// false
	FromOfficial *bool `json:"FromOfficial,omitempty" xml:"FromOfficial,omitempty"`
	// example:
	//
	// 1586920240
	GmtCreated *int32 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// CT24GHXT
	NoteId *string `json:"NoteId,omitempty" xml:"NoteId,omitempty"`
}

func (s ListTicketNotesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTicketNotesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTicketNotesResponseBodyDataList) GetContent() *string {
	return s.Content
}

func (s *ListTicketNotesResponseBodyDataList) GetFromOfficial() *bool {
	return s.FromOfficial
}

func (s *ListTicketNotesResponseBodyDataList) GetGmtCreated() *int32 {
	return s.GmtCreated
}

func (s *ListTicketNotesResponseBodyDataList) GetNoteId() *string {
	return s.NoteId
}

func (s *ListTicketNotesResponseBodyDataList) SetContent(v string) *ListTicketNotesResponseBodyDataList {
	s.Content = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) SetFromOfficial(v bool) *ListTicketNotesResponseBodyDataList {
	s.FromOfficial = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) SetGmtCreated(v int32) *ListTicketNotesResponseBodyDataList {
	s.GmtCreated = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) SetNoteId(v string) *ListTicketNotesResponseBodyDataList {
	s.NoteId = &v
	return s
}

func (s *ListTicketNotesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
