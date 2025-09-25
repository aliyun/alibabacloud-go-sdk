// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealtimeDialogAssistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCost(v int64) *RealtimeDialogAssistResponseBody
	GetCost() *int64
	SetData(v *RealtimeDialogAssistResponseBodyData) *RealtimeDialogAssistResponseBody
	GetData() *RealtimeDialogAssistResponseBodyData
	SetDataType(v string) *RealtimeDialogAssistResponseBody
	GetDataType() *string
	SetErrCode(v string) *RealtimeDialogAssistResponseBody
	GetErrCode() *string
	SetMessage(v string) *RealtimeDialogAssistResponseBody
	GetMessage() *string
	SetRequestId(v string) *RealtimeDialogAssistResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RealtimeDialogAssistResponseBody
	GetSuccess() *bool
	SetTime(v string) *RealtimeDialogAssistResponseBody
	GetTime() *string
}

type RealtimeDialogAssistResponseBody struct {
	// example:
	//
	// null
	Cost *int64                                `json:"cost,omitempty" xml:"cost,omitempty"`
	Data *RealtimeDialogAssistResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 2024-04-24 11:54:34
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
}

func (s RealtimeDialogAssistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistResponseBody) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistResponseBody) GetCost() *int64 {
	return s.Cost
}

func (s *RealtimeDialogAssistResponseBody) GetData() *RealtimeDialogAssistResponseBodyData {
	return s.Data
}

func (s *RealtimeDialogAssistResponseBody) GetDataType() *string {
	return s.DataType
}

func (s *RealtimeDialogAssistResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RealtimeDialogAssistResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RealtimeDialogAssistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RealtimeDialogAssistResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RealtimeDialogAssistResponseBody) GetTime() *string {
	return s.Time
}

func (s *RealtimeDialogAssistResponseBody) SetCost(v int64) *RealtimeDialogAssistResponseBody {
	s.Cost = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetData(v *RealtimeDialogAssistResponseBodyData) *RealtimeDialogAssistResponseBody {
	s.Data = v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetDataType(v string) *RealtimeDialogAssistResponseBody {
	s.DataType = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetErrCode(v string) *RealtimeDialogAssistResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetMessage(v string) *RealtimeDialogAssistResponseBody {
	s.Message = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetRequestId(v string) *RealtimeDialogAssistResponseBody {
	s.RequestId = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetSuccess(v bool) *RealtimeDialogAssistResponseBody {
	s.Success = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) SetTime(v string) *RealtimeDialogAssistResponseBody {
	s.Time = &v
	return s
}

func (s *RealtimeDialogAssistResponseBody) Validate() error {
	return dara.Validate(s)
}

type RealtimeDialogAssistResponseBodyData struct {
	AnalysisProcess   *string                                                  `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	AssistScripts     []*RealtimeDialogAssistResponseBodyDataAssistScripts     `json:"assistScripts,omitempty" xml:"assistScripts,omitempty" type:"Repeated"`
	AssistSop         []*RealtimeDialogAssistResponseBodyDataAssistSop         `json:"assistSop,omitempty" xml:"assistSop,omitempty" type:"Repeated"`
	ConversationModel []*RealtimeDialogAssistResponseBodyDataConversationModel `json:"conversationModel,omitempty" xml:"conversationModel,omitempty" type:"Repeated"`
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1915593248420413441
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s RealtimeDialogAssistResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistResponseBodyData) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistResponseBodyData) GetAnalysisProcess() *string {
	return s.AnalysisProcess
}

func (s *RealtimeDialogAssistResponseBodyData) GetAssistScripts() []*RealtimeDialogAssistResponseBodyDataAssistScripts {
	return s.AssistScripts
}

func (s *RealtimeDialogAssistResponseBodyData) GetAssistSop() []*RealtimeDialogAssistResponseBodyDataAssistSop {
	return s.AssistSop
}

func (s *RealtimeDialogAssistResponseBodyData) GetConversationModel() []*RealtimeDialogAssistResponseBodyDataConversationModel {
	return s.ConversationModel
}

func (s *RealtimeDialogAssistResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *RealtimeDialogAssistResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *RealtimeDialogAssistResponseBodyData) SetAnalysisProcess(v string) *RealtimeDialogAssistResponseBodyData {
	s.AnalysisProcess = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyData) SetAssistScripts(v []*RealtimeDialogAssistResponseBodyDataAssistScripts) *RealtimeDialogAssistResponseBodyData {
	s.AssistScripts = v
	return s
}

func (s *RealtimeDialogAssistResponseBodyData) SetAssistSop(v []*RealtimeDialogAssistResponseBodyDataAssistSop) *RealtimeDialogAssistResponseBodyData {
	s.AssistSop = v
	return s
}

func (s *RealtimeDialogAssistResponseBodyData) SetConversationModel(v []*RealtimeDialogAssistResponseBodyDataConversationModel) *RealtimeDialogAssistResponseBodyData {
	s.ConversationModel = v
	return s
}

func (s *RealtimeDialogAssistResponseBodyData) SetRequestId(v string) *RealtimeDialogAssistResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyData) SetSessionId(v string) *RealtimeDialogAssistResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RealtimeDialogAssistResponseBodyDataAssistScripts struct {
	AssistScript *string `json:"assistScript,omitempty" xml:"assistScript,omitempty"`
	// example:
	//
	// 1920005488515465216
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	// example:
	//
	// null
	IntentLabels *string `json:"intentLabels,omitempty" xml:"intentLabels,omitempty"`
	IntentName   *string `json:"intentName,omitempty" xml:"intentName,omitempty"`
	IsDefault    *bool   `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
}

func (s RealtimeDialogAssistResponseBodyDataAssistScripts) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistResponseBodyDataAssistScripts) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) GetAssistScript() *string {
	return s.AssistScript
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) GetIntentCode() *string {
	return s.IntentCode
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) GetIntentLabels() *string {
	return s.IntentLabels
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) GetIntentName() *string {
	return s.IntentName
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) SetAssistScript(v string) *RealtimeDialogAssistResponseBodyDataAssistScripts {
	s.AssistScript = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) SetIntentCode(v string) *RealtimeDialogAssistResponseBodyDataAssistScripts {
	s.IntentCode = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) SetIntentLabels(v string) *RealtimeDialogAssistResponseBodyDataAssistScripts {
	s.IntentLabels = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) SetIntentName(v string) *RealtimeDialogAssistResponseBodyDataAssistScripts {
	s.IntentName = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) SetIsDefault(v bool) *RealtimeDialogAssistResponseBodyDataAssistScripts {
	s.IsDefault = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistScripts) Validate() error {
	return dara.Validate(s)
}

type RealtimeDialogAssistResponseBodyDataAssistSop struct {
	// example:
	//
	// XXX
	AssistSop *string `json:"assistSop,omitempty" xml:"assistSop,omitempty"`
	// example:
	//
	// XXX
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	// example:
	//
	// XXX
	IntentName *string `json:"intentName,omitempty" xml:"intentName,omitempty"`
	IsDefault  *bool   `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
}

func (s RealtimeDialogAssistResponseBodyDataAssistSop) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistResponseBodyDataAssistSop) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) GetAssistSop() *string {
	return s.AssistSop
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) GetIntentCode() *string {
	return s.IntentCode
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) GetIntentName() *string {
	return s.IntentName
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) SetAssistSop(v string) *RealtimeDialogAssistResponseBodyDataAssistSop {
	s.AssistSop = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) SetIntentCode(v string) *RealtimeDialogAssistResponseBodyDataAssistSop {
	s.IntentCode = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) SetIntentName(v string) *RealtimeDialogAssistResponseBodyDataAssistSop {
	s.IntentName = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) SetIsDefault(v bool) *RealtimeDialogAssistResponseBodyDataAssistSop {
	s.IsDefault = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataAssistSop) Validate() error {
	return dara.Validate(s)
}

type RealtimeDialogAssistResponseBodyDataConversationModel struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 98457834685635
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// example:
	//
	// 1374683645635
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// example:
	//
	// 0
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RealtimeDialogAssistResponseBodyDataConversationModel) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistResponseBodyDataConversationModel) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) GetContent() *string {
	return s.Content
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) GetCustomerId() *string {
	return s.CustomerId
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) GetRole() *string {
	return s.Role
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) GetType() *string {
	return s.Type
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) SetContent(v string) *RealtimeDialogAssistResponseBodyDataConversationModel {
	s.Content = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) SetCustomerId(v string) *RealtimeDialogAssistResponseBodyDataConversationModel {
	s.CustomerId = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) SetCustomerServiceId(v string) *RealtimeDialogAssistResponseBodyDataConversationModel {
	s.CustomerServiceId = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) SetCustomerServiceType(v string) *RealtimeDialogAssistResponseBodyDataConversationModel {
	s.CustomerServiceType = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) SetRole(v string) *RealtimeDialogAssistResponseBodyDataConversationModel {
	s.Role = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) SetType(v string) *RealtimeDialogAssistResponseBodyDataConversationModel {
	s.Type = &v
	return s
}

func (s *RealtimeDialogAssistResponseBodyDataConversationModel) Validate() error {
	return dara.Validate(s)
}
