// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListTicketsResponseBody
	GetMessage() *string
	SetPage(v *ListTicketsResponseBodyPage) *ListTicketsResponseBody
	GetPage() *ListTicketsResponseBodyPage
	SetRequestId(v string) *ListTicketsResponseBody
	GetRequestId() *string
	SetResult(v []*ListTicketsResponseBodyResult) *ListTicketsResponseBody
	GetResult() []*ListTicketsResponseBodyResult
	SetStatusCode(v int32) *ListTicketsResponseBody
	GetStatusCode() *int32
}

type ListTicketsResponseBody struct {
	// example:
	//
	// success
	Message *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListTicketsResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListTicketsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListTicketsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTicketsResponseBody) GetPage() *ListTicketsResponseBodyPage {
	return s.Page
}

func (s *ListTicketsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTicketsResponseBody) GetResult() []*ListTicketsResponseBodyResult {
	return s.Result
}

func (s *ListTicketsResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTicketsResponseBody) SetMessage(v string) *ListTicketsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTicketsResponseBody) SetPage(v *ListTicketsResponseBodyPage) *ListTicketsResponseBody {
	s.Page = v
	return s
}

func (s *ListTicketsResponseBody) SetRequestId(v string) *ListTicketsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTicketsResponseBody) SetResult(v []*ListTicketsResponseBodyResult) *ListTicketsResponseBody {
	s.Result = v
	return s
}

func (s *ListTicketsResponseBody) SetStatusCode(v int32) *ListTicketsResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListTicketsResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketsResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTicketsResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTicketsResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTicketsResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListTicketsResponseBodyPage) SetPageNumber(v int32) *ListTicketsResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListTicketsResponseBodyPage) SetPageSize(v int32) *ListTicketsResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListTicketsResponseBodyPage) SetTotal(v int32) *ListTicketsResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListTicketsResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListTicketsResponseBodyResult struct {
	// example:
	//
	// false
	Action *bool `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// ***
	AssignedHandler *string `json:"AssignedHandler,omitempty" xml:"AssignedHandler,omitempty"`
	// example:
	//
	// ***
	ChargesRemark *string `json:"ChargesRemark,omitempty" xml:"ChargesRemark,omitempty"`
	// example:
	//
	// ***
	CompleteRemark *string                                 `json:"CompleteRemark,omitempty" xml:"CompleteRemark,omitempty"`
	Dialogs        []*ListTicketsResponseBodyResultDialogs `json:"Dialogs,omitempty" xml:"Dialogs,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtCalled *string `json:"GmtCalled,omitempty" xml:"GmtCalled,omitempty"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtDelayed *string `json:"GmtDelayed,omitempty" xml:"GmtDelayed,omitempty"`
	// example:
	//
	// 2023-01-09 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2023***93975
	GroupKey *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// example:
	//
	// 45
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsAcceptedCharges *bool `json:"IsAcceptedCharges,omitempty" xml:"IsAcceptedCharges,omitempty"`
	// example:
	//
	// true
	IsDelayed *bool `json:"IsDelayed,omitempty" xml:"IsDelayed,omitempty"`
	// example:
	//
	// false
	IsNeedCallback *bool `json:"IsNeedCallback,omitempty" xml:"IsNeedCallback,omitempty"`
	// example:
	//
	// false
	IsNeedCharges *bool                    `json:"IsNeedCharges,omitempty" xml:"IsNeedCharges,omitempty"`
	Operations    []map[string]interface{} `json:"Operations,omitempty" xml:"Operations,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// 101
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTicketsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyResult) GetAction() *bool {
	return s.Action
}

func (s *ListTicketsResponseBodyResult) GetAssignedHandler() *string {
	return s.AssignedHandler
}

func (s *ListTicketsResponseBodyResult) GetChargesRemark() *string {
	return s.ChargesRemark
}

func (s *ListTicketsResponseBodyResult) GetCompleteRemark() *string {
	return s.CompleteRemark
}

func (s *ListTicketsResponseBodyResult) GetDialogs() []*ListTicketsResponseBodyResultDialogs {
	return s.Dialogs
}

func (s *ListTicketsResponseBodyResult) GetGmtCalled() *string {
	return s.GmtCalled
}

func (s *ListTicketsResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListTicketsResponseBodyResult) GetGmtDelayed() *string {
	return s.GmtDelayed
}

func (s *ListTicketsResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListTicketsResponseBodyResult) GetGroupKey() *string {
	return s.GroupKey
}

func (s *ListTicketsResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *ListTicketsResponseBodyResult) GetIsAcceptedCharges() *bool {
	return s.IsAcceptedCharges
}

func (s *ListTicketsResponseBodyResult) GetIsDelayed() *bool {
	return s.IsDelayed
}

func (s *ListTicketsResponseBodyResult) GetIsNeedCallback() *bool {
	return s.IsNeedCallback
}

func (s *ListTicketsResponseBodyResult) GetIsNeedCharges() *bool {
	return s.IsNeedCharges
}

func (s *ListTicketsResponseBodyResult) GetOperations() []map[string]interface{} {
	return s.Operations
}

func (s *ListTicketsResponseBodyResult) GetRemark() *string {
	return s.Remark
}

func (s *ListTicketsResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *ListTicketsResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListTicketsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListTicketsResponseBodyResult) SetAction(v bool) *ListTicketsResponseBodyResult {
	s.Action = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetAssignedHandler(v string) *ListTicketsResponseBodyResult {
	s.AssignedHandler = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetChargesRemark(v string) *ListTicketsResponseBodyResult {
	s.ChargesRemark = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetCompleteRemark(v string) *ListTicketsResponseBodyResult {
	s.CompleteRemark = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetDialogs(v []*ListTicketsResponseBodyResultDialogs) *ListTicketsResponseBodyResult {
	s.Dialogs = v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtCalled(v string) *ListTicketsResponseBodyResult {
	s.GmtCalled = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtCreate(v string) *ListTicketsResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtDelayed(v string) *ListTicketsResponseBodyResult {
	s.GmtDelayed = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGmtModified(v string) *ListTicketsResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetGroupKey(v string) *ListTicketsResponseBodyResult {
	s.GroupKey = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetId(v int64) *ListTicketsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsAcceptedCharges(v bool) *ListTicketsResponseBodyResult {
	s.IsAcceptedCharges = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsDelayed(v bool) *ListTicketsResponseBodyResult {
	s.IsDelayed = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsNeedCallback(v bool) *ListTicketsResponseBodyResult {
	s.IsNeedCallback = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetIsNeedCharges(v bool) *ListTicketsResponseBodyResult {
	s.IsNeedCharges = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetOperations(v []map[string]interface{}) *ListTicketsResponseBodyResult {
	s.Operations = v
	return s
}

func (s *ListTicketsResponseBodyResult) SetRemark(v string) *ListTicketsResponseBodyResult {
	s.Remark = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetRoomNo(v string) *ListTicketsResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetStatus(v string) *ListTicketsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListTicketsResponseBodyResult) SetType(v string) *ListTicketsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListTicketsResponseBodyResult) Validate() error {
	if s.Dialogs != nil {
		for _, item := range s.Dialogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTicketsResponseBodyResultDialogs struct {
	Answer   *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	Question *string `json:"Question,omitempty" xml:"Question,omitempty"`
}

func (s ListTicketsResponseBodyResultDialogs) String() string {
	return dara.Prettify(s)
}

func (s ListTicketsResponseBodyResultDialogs) GoString() string {
	return s.String()
}

func (s *ListTicketsResponseBodyResultDialogs) GetAnswer() *string {
	return s.Answer
}

func (s *ListTicketsResponseBodyResultDialogs) GetQuestion() *string {
	return s.Question
}

func (s *ListTicketsResponseBodyResultDialogs) SetAnswer(v string) *ListTicketsResponseBodyResultDialogs {
	s.Answer = &v
	return s
}

func (s *ListTicketsResponseBodyResultDialogs) SetQuestion(v string) *ListTicketsResponseBodyResultDialogs {
	s.Question = &v
	return s
}

func (s *ListTicketsResponseBodyResultDialogs) Validate() error {
	return dara.Validate(s)
}
