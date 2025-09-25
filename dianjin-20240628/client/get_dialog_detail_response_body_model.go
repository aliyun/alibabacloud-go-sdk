// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDialogDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetDialogDetailResponseBody
	GetCost() *int64
	SetData(v *GetDialogDetailResponseBodyData) *GetDialogDetailResponseBody
	GetData() *GetDialogDetailResponseBodyData
	SetDataType(v string) *GetDialogDetailResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetDialogDetailResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetDialogDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDialogDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDialogDetailResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetDialogDetailResponseBody
	GetTime() *string
}

type GetDialogDetailResponseBody struct {
	// example:
	//
	// null
	Cost *int64                           `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetDialogDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 5E3FBAF1-17AF-53B7-AF0A-CDCEEB6DE658
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetDialogDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDialogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetDialogDetailResponseBody) GetData() *GetDialogDetailResponseBodyData {
	return s.Data
}

func (s *GetDialogDetailResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetDialogDetailResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetDialogDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDialogDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDialogDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDialogDetailResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetDialogDetailResponseBody) SetCost(v int64) *GetDialogDetailResponseBody {
	s.Cost = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetData(v *GetDialogDetailResponseBodyData) *GetDialogDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetDialogDetailResponseBody) SetDataType(v string) *GetDialogDetailResponseBody {
	s.DataType = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetErrCode(v string) *GetDialogDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetMessage(v string) *GetDialogDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetRequestId(v string) *GetDialogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetSuccess(v bool) *GetDialogDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetDialogDetailResponseBody) SetTime(v string) *GetDialogDetailResponseBody {
	s.Time = &v
	return s
}

func (s *GetDialogDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDialogDetailResponseBodyData struct {
	DialogueList []*GetDialogDetailResponseBodyDataDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 10
	TotalDialogTurns *int32 `json:"totalDialogTurns,omitempty" xml:"totalDialogTurns,omitempty"`
	// example:
	//
	// 5
	ValidDialogTurns *int32 `json:"validDialogTurns,omitempty" xml:"validDialogTurns,omitempty"`
}

func (s GetDialogDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDialogDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponseBodyData) GetDialogueList() []*GetDialogDetailResponseBodyDataDialogueList {
	return s.DialogueList
}

func (s *GetDialogDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDialogDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDialogDetailResponseBodyData) GetTotalDialogTurns() *int32 {
	return s.TotalDialogTurns
}

func (s *GetDialogDetailResponseBodyData) GetValidDialogTurns() *int32 {
	return s.ValidDialogTurns
}

func (s *GetDialogDetailResponseBodyData) SetDialogueList(v []*GetDialogDetailResponseBodyDataDialogueList) *GetDialogDetailResponseBodyData {
	s.DialogueList = v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetGmtCreate(v string) *GetDialogDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetStatus(v string) *GetDialogDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetTotalDialogTurns(v int32) *GetDialogDetailResponseBodyData {
	s.TotalDialogTurns = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) SetValidDialogTurns(v int32) *GetDialogDetailResponseBodyData {
	s.ValidDialogTurns = &v
	return s
}

func (s *GetDialogDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetDialogDetailResponseBodyDataDialogueList struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 123761283
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// BOT
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// true
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// example:
	//
	// 1742869659849
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 193874634xxx
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	IntentName *string `json:"intentName,omitempty" xml:"intentName,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetDialogDetailResponseBodyDataDialogueList) String() string {
	return dara.Prettify(s)
}

func (s GetDialogDetailResponseBodyDataDialogueList) GoString() string {
	return s.String()
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetContent() *string {
	return s.Content
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetHangUpDialog() *bool {
	return s.HangUpDialog
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetId() *int64 {
	return s.Id
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetIntentCode() *string {
	return s.IntentCode
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetIntentName() *string {
	return s.IntentName
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetRole() *string {
	return s.Role
}

func (s *GetDialogDetailResponseBodyDataDialogueList) GetType() *string {
	return s.Type
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetContent(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.Content = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetCustomerId(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.CustomerId = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetCustomerServiceId(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.CustomerServiceId = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetCustomerServiceType(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.CustomerServiceType = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetHangUpDialog(v bool) *GetDialogDetailResponseBodyDataDialogueList {
	s.HangUpDialog = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetId(v int64) *GetDialogDetailResponseBodyDataDialogueList {
	s.Id = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetIntentCode(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.IntentCode = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetIntentName(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.IntentName = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetRole(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.Role = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) SetType(v string) *GetDialogDetailResponseBodyDataDialogueList {
	s.Type = &v
	return s
}

func (s *GetDialogDetailResponseBodyDataDialogueList) Validate() error {
	return dara.Validate(s)
}
