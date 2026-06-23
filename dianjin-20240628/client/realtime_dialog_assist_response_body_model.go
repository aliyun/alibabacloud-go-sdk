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
	// Time consumed
	//
	// example:
	//
	// null
	Cost *int64 `json:"cost,omitempty" xml:"cost,omitempty"`
	// Response data
	Data *RealtimeDialogAssistResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Data type
	//
	// example:
	//
	// null
	DataType *string `json:"dataType,omitempty" xml:"dataType,omitempty"`
	// Error code
	//
	// example:
	//
	// 0
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// Error message
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID. This is the system-recorded request ID. If issues arise, provide this ID to the Model Studio DianJin R\\&D team for troubleshooting.
	//
	// example:
	//
	// 67C7021A-D268-553D-8C15-A087B9604028
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Whether successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// Timestamp
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RealtimeDialogAssistResponseBodyData struct {
	// Analysis process
	//
	// example:
	//
	// 客户回答的内容与提供的意图列表描述均不匹配，没有表达出对账单、还款、天气或其他服务的具体需求或问题。
	AnalysisProcess *string `json:"analysisProcess,omitempty" xml:"analysisProcess,omitempty"`
	// List of dialog assist results
	AssistScripts []*RealtimeDialogAssistResponseBodyDataAssistScripts `json:"assistScripts,omitempty" xml:"assistScripts,omitempty" type:"Repeated"`
	// List of flow assist results
	AssistSop []*RealtimeDialogAssistResponseBodyDataAssistSop `json:"assistSop,omitempty" xml:"assistSop,omitempty" type:"Repeated"`
	// Current dialog content
	ConversationModel []*RealtimeDialogAssistResponseBodyDataConversationModel `json:"conversationModel,omitempty" xml:"conversationModel,omitempty" type:"Repeated"`
	// Whether interrupted
	//
	// example:
	//
	// true
	Interrupt *bool `json:"interrupt,omitempty" xml:"interrupt,omitempty"`
	// Unique request ID. This request ID matches the request ID in the input parameter.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Session ID
	//
	// example:
	//
	// "1915593248420413441"
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

func (s *RealtimeDialogAssistResponseBodyData) GetInterrupt() *bool {
	return s.Interrupt
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

func (s *RealtimeDialogAssistResponseBodyData) SetInterrupt(v bool) *RealtimeDialogAssistResponseBodyData {
	s.Interrupt = &v
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
	if s.AssistScripts != nil {
		for _, item := range s.AssistScripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AssistSop != nil {
		for _, item := range s.AssistSop {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConversationModel != nil {
		for _, item := range s.ConversationModel {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RealtimeDialogAssistResponseBodyDataAssistScripts struct {
	// Recommended utterance
	//
	// example:
	//
	// 可按照SOP流程回应。
	AssistScript *string `json:"assistScript,omitempty" xml:"assistScript,omitempty"`
	// Intent encoding
	//
	// example:
	//
	// "1920005488515465216"
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	// Intent labels
	//
	// example:
	//
	// null
	IntentLabels *string `json:"intentLabels,omitempty" xml:"intentLabels,omitempty"`
	// Intent name
	//
	// example:
	//
	// 礼貌问答
	IntentName *string `json:"intentName,omitempty" xml:"intentName,omitempty"`
	// Whether intent escaped
	//
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
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
	// Recommended flow
	//
	// example:
	//
	// XXX
	AssistSop *string `json:"assistSop,omitempty" xml:"assistSop,omitempty"`
	// Intent encoding
	//
	// example:
	//
	// XXX
	IntentCode *string `json:"intentCode,omitempty" xml:"intentCode,omitempty"`
	// Intent name
	//
	// example:
	//
	// XXX
	IntentName *string `json:"intentName,omitempty" xml:"intentName,omitempty"`
	// Indicates whether the intent is to escape.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
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
	// Specific content of the dialog
	//
	// example:
	//
	// 你好
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Unique identity of the dialog role
	//
	// example:
	//
	// "98457834685635"
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// Customer service ID
	//
	// example:
	//
	// "1374683645635"
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// Agent type. 0: Robot, 1: Human.
	//
	// example:
	//
	// "0"
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// Role. 0 indicates customer, 1 indicates agent.
	//
	// example:
	//
	// "0"
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// Type of dialog content
	//
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
