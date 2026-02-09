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
	SetData(v *ListTicketsResponseBodyData) *ListTicketsResponseBody
	GetData() *ListTicketsResponseBodyData
	SetMessage(v string) *ListTicketsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTicketsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTicketsResponseBody
	GetSuccess() *bool
}

type ListTicketsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *int32                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListTicketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s *ListTicketsResponseBody) GetData() *ListTicketsResponseBodyData {
	return s.Data
}

func (s *ListTicketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTicketsResponseBody) SetAccessDeniedDetail(v string) *ListTicketsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListTicketsResponseBody) SetCode(v int32) *ListTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketsResponseBody) SetData(v *ListTicketsResponseBodyData) *ListTicketsResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
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

func (s *ListTicketsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketsResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *int32                             `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	List        []*ListTicketsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 99
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTicketsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTicketsResponseBodyData) GetList() []*ListTicketsResponseBodyDataList {
	return s.List
}

func (s *ListTicketsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListTicketsResponseBodyData) SetCurrentPage(v int32) *ListTicketsResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetList(v []*ListTicketsResponseBodyDataList) *ListTicketsResponseBodyData {
	s.List = v
	return s
}

func (s *ListTicketsResponseBodyData) SetPageSize(v int32) *ListTicketsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetTotal(v int32) *ListTicketsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListTicketsResponseBodyData) Validate() error {
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

type ListTicketsResponseBodyDataList struct {
	// example:
	//
	// 1585818240
	AddTime *int32 `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// example:
	//
	// 252448085032933023
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// example:
	//
	// CAB327A
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// assigned
	TicketStatus *string `json:"TicketStatus,omitempty" xml:"TicketStatus,omitempty"`
	Title        *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListTicketsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyDataList) GetAddTime() *int32 {
	return s.AddTime
}

func (s *ListTicketsResponseBodyDataList) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListTicketsResponseBodyDataList) GetId() *string {
	return s.Id
}

func (s *ListTicketsResponseBodyDataList) GetTicketStatus() *string {
	return s.TicketStatus
}

func (s *ListTicketsResponseBodyDataList) GetTitle() *string {
	return s.Title
}

func (s *ListTicketsResponseBodyDataList) SetAddTime(v int32) *ListTicketsResponseBodyDataList {
	s.AddTime = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetCreatorId(v string) *ListTicketsResponseBodyDataList {
	s.CreatorId = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetId(v string) *ListTicketsResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTicketStatus(v string) *ListTicketsResponseBodyDataList {
	s.TicketStatus = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) SetTitle(v string) *ListTicketsResponseBodyDataList {
	s.Title = &v
	return s
}

func (s *ListTicketsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
