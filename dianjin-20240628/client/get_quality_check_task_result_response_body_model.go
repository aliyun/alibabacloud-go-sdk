// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityCheckTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *GetQualityCheckTaskResultResponseBody
	GetCost() *int64
	SetData(v *GetQualityCheckTaskResultResponseBodyData) *GetQualityCheckTaskResultResponseBody
	GetData() *GetQualityCheckTaskResultResponseBodyData
	SetDataType(v string) *GetQualityCheckTaskResultResponseBody
	GetDataType() *string
	SetErrCode(v string) *GetQualityCheckTaskResultResponseBody
	GetErrCode() *string
	SetMessage(v string) *GetQualityCheckTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetQualityCheckTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityCheckTaskResultResponseBody
	GetSuccess() *bool
	SetTime(v string) *GetQualityCheckTaskResultResponseBody
	GetTime() *string
}

type GetQualityCheckTaskResultResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                     `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *GetQualityCheckTaskResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 67C7021A-D268-553D-8C15-A087B9604028
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2024-01-01 00:00:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *GetQualityCheckTaskResultResponseBody) GetData() *GetQualityCheckTaskResultResponseBodyData {
	return s.Data
}

func (s *GetQualityCheckTaskResultResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *GetQualityCheckTaskResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetQualityCheckTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityCheckTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityCheckTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityCheckTaskResultResponseBody) GetTime() *string {
	return s.Time
}

func (s *GetQualityCheckTaskResultResponseBody) SetCost(v int64) *GetQualityCheckTaskResultResponseBody {
	s.Cost = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetData(v *GetQualityCheckTaskResultResponseBodyData) *GetQualityCheckTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetDataType(v string) *GetQualityCheckTaskResultResponseBody {
	s.DataType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetErrCode(v string) *GetQualityCheckTaskResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetMessage(v string) *GetQualityCheckTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetRequestId(v string) *GetQualityCheckTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetSuccess(v bool) *GetQualityCheckTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) SetTime(v string) *GetQualityCheckTaskResultResponseBody {
	s.Time = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetQualityCheckTaskResultResponseBodyData struct {
	ConversationList *GetQualityCheckTaskResultResponseBodyDataConversationList `json:"conversationList,omitempty" xml:"conversationList,omitempty" type:"Struct"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtStart         *string                                                      `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	QualityCheckList []*GetQualityCheckTaskResultResponseBodyDataQualityCheckList `json:"qualityCheckList,omitempty" xml:"qualityCheckList,omitempty" type:"Repeated"`
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1703557101831
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetConversationList() *GetQualityCheckTaskResultResponseBodyDataConversationList {
	return s.ConversationList
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetGmtStart() *string {
	return s.GmtStart
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetQualityCheckList() []*GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	return s.QualityCheckList
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetQualityCheckTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetConversationList(v *GetQualityCheckTaskResultResponseBodyDataConversationList) *GetQualityCheckTaskResultResponseBodyData {
	s.ConversationList = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetGmtCreate(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetGmtEnd(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.GmtEnd = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetGmtStart(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.GmtStart = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetQualityCheckList(v []*GetQualityCheckTaskResultResponseBodyDataQualityCheckList) *GetQualityCheckTaskResultResponseBodyData {
	s.QualityCheckList = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetStatus(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) SetTaskId(v string) *GetQualityCheckTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetQualityCheckTaskResultResponseBodyDataConversationList struct {
	// example:
	//
	// 1
	CallType *string `json:"callType,omitempty" xml:"callType,omitempty"`
	// example:
	//
	// 234234
	CustomerId   *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	CustomerName *string `json:"customerName,omitempty" xml:"customerName,omitempty"`
	// example:
	//
	// 23984763826
	CustomerServiceId   *string                                                                  `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	CustomerServiceName *string                                                                  `json:"customerServiceName,omitempty" xml:"customerServiceName,omitempty"`
	DialogueList        []*GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList `json:"dialogueList,omitempty" xml:"dialogueList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-27 11:23:20
	GmtService *string `json:"gmtService,omitempty" xml:"gmtService,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetCallType() *string {
	return s.CallType
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetCustomerName() *string {
	return s.CustomerName
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetCustomerServiceName() *string {
	return s.CustomerServiceName
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetDialogueList() []*GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	return s.DialogueList
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) GetGmtService() *string {
	return s.GmtService
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCallType(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CallType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerName(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerName = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerServiceId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerServiceId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetCustomerServiceName(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.CustomerServiceName = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetDialogueList(v []*GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.DialogueList = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) SetGmtService(v string) *GetQualityCheckTaskResultResponseBodyDataConversationList {
	s.GmtService = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationList) Validate() error {
	return dara.Validate(s)
}

type GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList struct {
	// example:
	//
	// 0
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-09-27 11:23:20
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// null
	CustomerId        *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetBegin() *int32 {
	return s.Begin
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetContent() *string {
	return s.Content
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetEnd() *int32 {
	return s.End
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetId() *int32 {
	return s.Id
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetRole() *string {
	return s.Role
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) GetType() *string {
	return s.Type
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetBegin(v int32) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Begin = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetBeginTime(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.BeginTime = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetContent(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Content = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetCustomerId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.CustomerId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetCustomerServiceId(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.CustomerServiceId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetCustomerServiceType(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.CustomerServiceType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetEnd(v int32) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.End = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetId(v int32) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Id = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetRole(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Role = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) SetType(v string) *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList {
	s.Type = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataConversationListDialogueList) Validate() error {
	return dara.Validate(s)
}

type GetQualityCheckTaskResultResponseBodyDataQualityCheckList struct {
	BizType          *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	CheckExplanation *string `json:"checkExplanation,omitempty" xml:"checkExplanation,omitempty"`
	// example:
	//
	// PASSED
	CheckPassed  *string `json:"checkPassed,omitempty" xml:"checkPassed,omitempty"`
	CheckProcess *string `json:"checkProcess,omitempty" xml:"checkProcess,omitempty"`
	// example:
	//
	// HIT
	Checked *string `json:"checked,omitempty" xml:"checked,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	GmtEnd *string `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	GmtStart *string `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// example:
	//
	// 0
	Mode           *string                                                                    `json:"mode,omitempty" xml:"mode,omitempty"`
	OriginDialogue []*GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue `json:"originDialogue,omitempty" xml:"originDialogue,omitempty" type:"Repeated"`
	// example:
	//
	// warning_customers
	QualityGroupId  *string `json:"qualityGroupId,omitempty" xml:"qualityGroupId,omitempty"`
	RuleDescription *string `json:"ruleDescription,omitempty" xml:"ruleDescription,omitempty"`
	// example:
	//
	// wcm_start
	RuleId     *string       `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	RuleType   *string       `json:"ruleType,omitempty" xml:"ruleType,omitempty"`
	SubNodeCol []interface{} `json:"subNodeCol,omitempty" xml:"subNodeCol,omitempty" type:"Repeated"`
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckList) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetBizType() *string {
	return s.BizType
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetCheckExplanation() *string {
	return s.CheckExplanation
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetCheckPassed() *string {
	return s.CheckPassed
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetCheckProcess() *string {
	return s.CheckProcess
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetChecked() *string {
	return s.Checked
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetGmtEnd() *string {
	return s.GmtEnd
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetGmtStart() *string {
	return s.GmtStart
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetMode() *string {
	return s.Mode
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetOriginDialogue() []*GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	return s.OriginDialogue
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetQualityGroupId() *string {
	return s.QualityGroupId
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetRuleDescription() *string {
	return s.RuleDescription
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetRuleId() *string {
	return s.RuleId
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) GetSubNodeCol() []interface{} {
	return s.SubNodeCol
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetBizType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.BizType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetCheckExplanation(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.CheckExplanation = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetCheckPassed(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.CheckPassed = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetCheckProcess(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.CheckProcess = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetChecked(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.Checked = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetGmtEnd(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.GmtEnd = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetGmtStart(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.GmtStart = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetMode(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.Mode = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetOriginDialogue(v []*GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.OriginDialogue = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetQualityGroupId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.QualityGroupId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetRuleDescription(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.RuleDescription = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetRuleId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.RuleId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetRuleType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.RuleType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) SetSubNodeCol(v []interface{}) *GetQualityCheckTaskResultResponseBodyDataQualityCheckList {
	s.SubNodeCol = v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckList) Validate() error {
	return dara.Validate(s)
}

type GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue struct {
	// example:
	//
	// 0
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2024-05-23 14:57:50
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	Content   *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// xxx
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 23876432
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// TEXT
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) String() string {
	return dara.Prettify(s)
}

func (s GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GoString() string {
	return s.String()
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetBegin() *int32 {
	return s.Begin
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetBeginTime() *string {
	return s.BeginTime
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetContent() *string {
	return s.Content
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetEnd() *int32 {
	return s.End
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetId() *int32 {
	return s.Id
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetRole() *string {
	return s.Role
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) GetType() *string {
	return s.Type
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetBegin(v int32) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Begin = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetBeginTime(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.BeginTime = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetContent(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Content = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetCustomerId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.CustomerId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetCustomerServiceId(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.CustomerServiceId = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetCustomerServiceType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.CustomerServiceType = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetEnd(v int32) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.End = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetId(v int32) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Id = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetRole(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Role = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) SetType(v string) *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue {
	s.Type = &v
	return s
}

func (s *GetQualityCheckTaskResultResponseBodyDataQualityCheckListOriginDialogue) Validate() error {
	return dara.Validate(s)
}
