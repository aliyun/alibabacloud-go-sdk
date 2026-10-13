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
	SetData(v []*ListTicketsResponseBodyData) *ListTicketsResponseBody
	GetData() []*ListTicketsResponseBodyData
	SetMessage(v string) *ListTicketsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListTicketsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTicketsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTicketsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTicketsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListTicketsResponseBody
	GetTotalCount() *int64
}

type ListTicketsResponseBody struct {
	// The details of the permission error. If you have the required permissions, an empty string is returned.
	//
	// example:
	//
	// {\\"AuthAction\\":\\"ram:ListTickets\\",\\"AuthPrincipalDisplayName\\":\\"2146216584788xxxxx\\",\\"AuthPrincipalOwnerId\\":\\"1135850448xxxxx\\",\\"AuthPrincipalType\\":\\"SubUser\\",\\"NoPermissionType\\":\\"ImplicitDeny\\",\\"PolicyType\\":\\"AccountLevelIdentityBasedPolicy\\"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the request.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of your tickets.
	Data []*ListTicketsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message. This parameter is returned only when the value of success is false.
	//
	// example:
	//
	// 入参解析失败
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The unique ID of the request. A unique ID is returned for each request.
	//
	// example:
	//
	// AC0AB2EC-AFBC-44BA-AE77-132A5A1EC0AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. A value of true indicates that the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of query results, which is the total number of your tickets.
	//
	// example:
	//
	// 99
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListTicketsResponseBody) GetData() []*ListTicketsResponseBodyData {
	return s.Data
}

func (s *ListTicketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
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

func (s *ListTicketsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTicketsResponseBody) SetAccessDeniedDetail(v string) *ListTicketsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListTicketsResponseBody) SetCode(v int32) *ListTicketsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTicketsResponseBody) SetData(v []*ListTicketsResponseBodyData) *ListTicketsResponseBody {
	s.Data = v
	return s
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketsResponseBody) SetPageNumber(v int32) *ListTicketsResponseBody {
	s.PageNumber = &v
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

func (s *ListTicketsResponseBody) SetTotalCount(v int64) *ListTicketsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTicketsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketsResponseBodyData struct {
	// The ticket status. Valid values: assigned (Pending response), dealing (In progress), wait_feedback (Waiting for feedback), feedback (Feedback received), wait_confirm (Pending confirmation), and confirmed (Completed).
	Status *ListTicketsResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	// The ticket ID.
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// The ticket title.
	//
	// example:
	//
	// Why did ECS renewal fail?
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListTicketsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyData) GetStatus() *ListTicketsResponseBodyDataStatus {
	return s.Status
}

func (s *ListTicketsResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *ListTicketsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *ListTicketsResponseBodyData) SetStatus(v *ListTicketsResponseBodyDataStatus) *ListTicketsResponseBodyData {
	s.Status = v
	return s
}

func (s *ListTicketsResponseBodyData) SetTicketId(v string) *ListTicketsResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *ListTicketsResponseBodyData) SetTitle(v string) *ListTicketsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListTicketsResponseBodyData) Validate() error {
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketsResponseBodyDataStatus struct {
	// The description of the status. Example: Completed.
	//
	// example:
	//
	// Completed
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The value of the status. For example, a value of 6 indicates that the ticket is completed.
	//
	// example:
	//
	// 6
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTicketsResponseBodyDataStatus) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyDataStatus) GetLabel() *string {
	return s.Label
}

func (s *ListTicketsResponseBodyDataStatus) GetValue() *string {
	return s.Value
}

func (s *ListTicketsResponseBodyDataStatus) SetLabel(v string) *ListTicketsResponseBodyDataStatus {
	s.Label = &v
	return s
}

func (s *ListTicketsResponseBodyDataStatus) SetValue(v string) *ListTicketsResponseBodyDataStatus {
	s.Value = &v
	return s
}

func (s *ListTicketsResponseBodyDataStatus) Validate() error {
	return dara.Validate(s)
}
