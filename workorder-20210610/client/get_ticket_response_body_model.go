// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetTicketResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetTicketResponseBody
	GetCode() *int32
	SetData(v *GetTicketResponseBodyData) *GetTicketResponseBody
	GetData() *GetTicketResponseBodyData
	SetMessage(v string) *GetTicketResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetTicketResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetTicketResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetTicketResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTicketResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetTicketResponseBody
	GetTotalCount() *int64
}

type GetTicketResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The return code of the request result.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned after the call succeeds.
	Data *GetTicketResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. If success is set to false, the message is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number of the paging query parameter
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a pagination query parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C499BB0F-630D-5BE6-B3EA-5FCD95B85503
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is normal.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of query results, that is, the total number of my ticket records.
	//
	// example:
	//
	// 108
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetTicketResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTicketResponseBody) GetData() *GetTicketResponseBodyData {
	return s.Data
}

func (s *GetTicketResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTicketResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetTicketResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTicketResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTicketResponseBody) SetAccessDeniedDetail(v string) *GetTicketResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetTicketResponseBody) SetCode(v int32) *GetTicketResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketResponseBody) SetData(v *GetTicketResponseBodyData) *GetTicketResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketResponseBody) SetMessage(v string) *GetTicketResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketResponseBody) SetPageNumber(v int32) *GetTicketResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTicketResponseBody) SetPageSize(v int32) *GetTicketResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTicketResponseBody) SetRequestId(v string) *GetTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketResponseBody) SetSuccess(v bool) *GetTicketResponseBody {
	s.Success = &v
	return s
}

func (s *GetTicketResponseBody) SetTotalCount(v int64) *GetTicketResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTicketResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTicketResponseBodyData struct {
	// The ID of the ticket issue category.
	//
	// example:
	//
	// 7161
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The timestamp when the ticket was created.
	//
	// example:
	//
	// 1623396736000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The uid of the ticket creator.
	//
	// example:
	//
	// 1168025
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the ticket. Only pure text format is supported.
	//
	// example:
	//
	// Why ECS renewal failed?
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Urgency enumeration value, 1 for general problem, 2 for urgent problem
	Severity *GetTicketResponseBodyDataSeverity `json:"Severity,omitempty" xml:"Severity,omitempty" type:"Struct"`
	// The status of the ticket. The reference values are as follows 1, "assigned", "Pending Response", 2, "handling", "handling", 3, "wait_feedback", "Pending feedback", 4: "feedback", "Feedback", 5, "wait_confirm", "To be confirmed", 6, "confirmed", "Completed"
	Status *GetTicketResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
	// Work Order Number
	//
	// example:
	//
	// 0005PYGCW
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
	// The title of the ticket.
	//
	// example:
	//
	// Why ECS renewal failed?
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTicketResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *GetTicketResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetTicketResponseBodyData) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetTicketResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetTicketResponseBodyData) GetSeverity() *GetTicketResponseBodyDataSeverity {
	return s.Severity
}

func (s *GetTicketResponseBodyData) GetStatus() *GetTicketResponseBodyDataStatus {
	return s.Status
}

func (s *GetTicketResponseBodyData) GetTicketId() *string {
	return s.TicketId
}

func (s *GetTicketResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetTicketResponseBodyData) SetCategoryId(v string) *GetTicketResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreateTime(v int64) *GetTicketResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTicketResponseBodyData) SetCreatorId(v string) *GetTicketResponseBodyData {
	s.CreatorId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetDescription(v string) *GetTicketResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTicketResponseBodyData) SetSeverity(v *GetTicketResponseBodyDataSeverity) *GetTicketResponseBodyData {
	s.Severity = v
	return s
}

func (s *GetTicketResponseBodyData) SetStatus(v *GetTicketResponseBodyDataStatus) *GetTicketResponseBodyData {
	s.Status = v
	return s
}

func (s *GetTicketResponseBodyData) SetTicketId(v string) *GetTicketResponseBodyData {
	s.TicketId = &v
	return s
}

func (s *GetTicketResponseBodyData) SetTitle(v string) *GetTicketResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetTicketResponseBodyData) Validate() error {
	if s.Severity != nil {
		if err := s.Severity.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTicketResponseBodyDataSeverity struct {
	// Enumeration Text
	//
	// example:
	//
	// Common
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Enumeration
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTicketResponseBodyDataSeverity) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyDataSeverity) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyDataSeverity) GetLabel() *string {
	return s.Label
}

func (s *GetTicketResponseBodyDataSeverity) GetValue() *string {
	return s.Value
}

func (s *GetTicketResponseBodyDataSeverity) SetLabel(v string) *GetTicketResponseBodyDataSeverity {
	s.Label = &v
	return s
}

func (s *GetTicketResponseBodyDataSeverity) SetValue(v string) *GetTicketResponseBodyDataSeverity {
	s.Value = &v
	return s
}

func (s *GetTicketResponseBodyDataSeverity) Validate() error {
	return dara.Validate(s)
}

type GetTicketResponseBodyDataStatus struct {
	// Status Enumeration Text
	//
	// example:
	//
	// Completed
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// State enumeration value
	//
	// example:
	//
	// 6
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTicketResponseBodyDataStatus) String() string {
	return dara.Prettify(s)
}

func (s GetTicketResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *GetTicketResponseBodyDataStatus) GetLabel() *string {
	return s.Label
}

func (s *GetTicketResponseBodyDataStatus) GetValue() *string {
	return s.Value
}

func (s *GetTicketResponseBodyDataStatus) SetLabel(v string) *GetTicketResponseBodyDataStatus {
	s.Label = &v
	return s
}

func (s *GetTicketResponseBodyDataStatus) SetValue(v string) *GetTicketResponseBodyDataStatus {
	s.Value = &v
	return s
}

func (s *GetTicketResponseBodyDataStatus) Validate() error {
	return dara.Validate(s)
}
