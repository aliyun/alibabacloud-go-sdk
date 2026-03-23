// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListTicketsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListTicketsResponseBody
	GetCode() *int32
	SetCurrentPage(v int32) *ListTicketsResponseBody
	GetCurrentPage() *int32
	SetList(v []*ListTicketsResponseBodyList) *ListTicketsResponseBody
	GetList() []*ListTicketsResponseBodyList
	SetMessage(v string) *ListTicketsResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListTicketsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTicketsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTicketsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *ListTicketsResponseBody
	GetTotal() *int32
}

type ListTicketsResponseBody struct {
	AccessDeniedDetail *string                        `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage        *int32                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List               []*ListTicketsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	Message            *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize           *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
	Total              *int32                         `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTicketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListTicketsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTicketsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTicketsResponseBody) GetList() []*ListTicketsResponseBodyList {
	return s.List
}

func (s *ListTicketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTicketsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListTicketsResponseBody) SetAccessDeniedDetail(v string) *ListTicketsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListTicketsResponseBody) SetCode(v int32) *ListTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketsResponseBody) SetCurrentPage(v int32) *ListTicketsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTicketsResponseBody) SetList(v []*ListTicketsResponseBodyList) *ListTicketsResponseBody {
	s.List = v
	return s
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketsResponseBody) SetPageSize(v int32) *ListTicketsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBody) SetRequestId(v string) *ListTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketsResponseBody) SetSuccess(v bool) *ListTicketsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTicketsResponseBody) SetTotal(v int32) *ListTicketsResponseBody {
	s.Total = &v
	return s
}

func (s *ListTicketsResponseBody) Validate() error {
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

type ListTicketsResponseBodyList struct {
	AddTime      *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListTicketsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyList) GetAddTime() *string {
	return s.AddTime
}

func (s *ListTicketsResponseBodyList) GetId() *string {
	return s.Id
}

func (s *ListTicketsResponseBodyList) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *ListTicketsResponseBodyList) GetTitle() *string {
	return s.Title
}

func (s *ListTicketsResponseBodyList) SetAddTime(v string) *ListTicketsResponseBodyList {
	s.AddTime = &v
	return s
}

func (s *ListTicketsResponseBodyList) SetId(v string) *ListTicketsResponseBodyList {
	s.Id = &v
	return s
}

func (s *ListTicketsResponseBodyList) SetTicketStatus(v string) *ListTicketsResponseBodyList {
	s.TicketStatus = &v
	return s
}

func (s *ListTicketsResponseBodyList) SetTitle(v string) *ListTicketsResponseBodyList {
	s.Title = &v
	return s
}

func (s *ListTicketsResponseBodyList) Validate() error {
	return dara.Validate(s)
}
