// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetChatListResponseBodyData) *GetChatListResponseBody
	GetData() *GetChatListResponseBodyData
	SetRequestId(v string) *GetChatListResponseBody
	GetRequestId() *string
}

type GetChatListResponseBody struct {
	// Returned data structure.
	Data *GetChatListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetChatListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatListResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatListResponseBody) GetData() *GetChatListResponseBodyData {
	return s.Data
}

func (s *GetChatListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatListResponseBody) SetData(v *GetChatListResponseBodyData) *GetChatListResponseBody {
	s.Data = v
	return s
}

func (s *GetChatListResponseBody) SetRequestId(v string) *GetChatListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatListResponseBodyData struct {
	// Q&A list.
	ChatList []*ChatItem `json:"chatList,omitempty" xml:"chatList,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 21
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 3
	TotalPage *int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s GetChatListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetChatListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetChatListResponseBodyData) GetChatList() []*ChatItem {
	return s.ChatList
}

func (s *GetChatListResponseBodyData) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetChatListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetChatListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetChatListResponseBodyData) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *GetChatListResponseBodyData) SetChatList(v []*ChatItem) *GetChatListResponseBodyData {
	s.ChatList = v
	return s
}

func (s *GetChatListResponseBodyData) SetCurrentPage(v int64) *GetChatListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *GetChatListResponseBodyData) SetPageSize(v int64) *GetChatListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetChatListResponseBodyData) SetTotal(v int64) *GetChatListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetChatListResponseBodyData) SetTotalPage(v int64) *GetChatListResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *GetChatListResponseBodyData) Validate() error {
	if s.ChatList != nil {
		for _, item := range s.ChatList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
