// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealTimeDialogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysis(v bool) *RealTimeDialogRequest
	GetAnalysis() *bool
	SetBizType(v string) *RealTimeDialogRequest
	GetBizType() *string
	SetConversationModel(v []*RealTimeDialogRequestConversationModel) *RealTimeDialogRequest
	GetConversationModel() []*RealTimeDialogRequestConversationModel
	SetDialogMemoryTurns(v int32) *RealTimeDialogRequest
	GetDialogMemoryTurns() *int32
	SetMetaData(v map[string]interface{}) *RealTimeDialogRequest
	GetMetaData() map[string]interface{}
	SetOpType(v string) *RealTimeDialogRequest
	GetOpType() *string
	SetRecommend(v bool) *RealTimeDialogRequest
	GetRecommend() *bool
	SetScriptContentPlayed(v string) *RealTimeDialogRequest
	GetScriptContentPlayed() *string
	SetSessionId(v string) *RealTimeDialogRequest
	GetSessionId() *string
	SetStream(v bool) *RealTimeDialogRequest
	GetStream() *bool
	SetUserVad(v bool) *RealTimeDialogRequest
	GetUserVad() *bool
}

type RealTimeDialogRequest struct {
	// Specifies whether to perform analysis.
	//
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// The business type. The default value is mixIntentChat.
	//
	// example:
	//
	// mixIntentChat
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// The list of conversations.
	//
	// This parameter is required.
	ConversationModel []*RealTimeDialogRequestConversationModel `json:"conversationModel,omitempty" xml:"conversationModel,omitempty" type:"Repeated"`
	// The number of historical conversation turns to include.
	//
	// example:
	//
	// 3
	DialogMemoryTurns *int32 `json:"dialogMemoryTurns,omitempty" xml:"dialogMemoryTurns,omitempty"`
	// The metadata used to encapsulate prompts.
	//
	// example:
	//
	// {
	//
	//       "phoneTailNumber": "机主尾号：98X1",
	//
	//       "preScreeningQuota": "预审额度：3万",
	//
	//       "generalInterest": "平台一般利息：20.4%"
	//
	//     }
	MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// The operation type. Only common and hierarchical are supported.
	//
	// example:
	//
	// common
	OpType *string `json:"opType,omitempty" xml:"opType,omitempty"`
	// The recommended intent.
	//
	// example:
	//
	// false
	Recommend *bool `json:"recommend,omitempty" xml:"recommend,omitempty"`
	// The part of the previous script from the customer service representative that has been played.
	//
	// example:
	//
	// 你好，我是
	ScriptContentPlayed *string `json:"scriptContentPlayed,omitempty" xml:"scriptContentPlayed,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 237645726354
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// Specifies whether to return the response in a stream.
	//
	// example:
	//
	// false
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// Specifies whether the user interrupted the conversation.
	//
	// example:
	//
	// true
	UserVad *bool `json:"userVad,omitempty" xml:"userVad,omitempty"`
}

func (s RealTimeDialogRequest) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogRequest) GoString() string {
	return s.String()
}

func (s *RealTimeDialogRequest) GetAnalysis() *bool {
	return s.Analysis
}

func (s *RealTimeDialogRequest) GetBizType() *string {
	return s.BizType
}

func (s *RealTimeDialogRequest) GetConversationModel() []*RealTimeDialogRequestConversationModel {
	return s.ConversationModel
}

func (s *RealTimeDialogRequest) GetDialogMemoryTurns() *int32 {
	return s.DialogMemoryTurns
}

func (s *RealTimeDialogRequest) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *RealTimeDialogRequest) GetOpType() *string {
	return s.OpType
}

func (s *RealTimeDialogRequest) GetRecommend() *bool {
	return s.Recommend
}

func (s *RealTimeDialogRequest) GetScriptContentPlayed() *string {
	return s.ScriptContentPlayed
}

func (s *RealTimeDialogRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RealTimeDialogRequest) GetStream() *bool {
	return s.Stream
}

func (s *RealTimeDialogRequest) GetUserVad() *bool {
	return s.UserVad
}

func (s *RealTimeDialogRequest) SetAnalysis(v bool) *RealTimeDialogRequest {
	s.Analysis = &v
	return s
}

func (s *RealTimeDialogRequest) SetBizType(v string) *RealTimeDialogRequest {
	s.BizType = &v
	return s
}

func (s *RealTimeDialogRequest) SetConversationModel(v []*RealTimeDialogRequestConversationModel) *RealTimeDialogRequest {
	s.ConversationModel = v
	return s
}

func (s *RealTimeDialogRequest) SetDialogMemoryTurns(v int32) *RealTimeDialogRequest {
	s.DialogMemoryTurns = &v
	return s
}

func (s *RealTimeDialogRequest) SetMetaData(v map[string]interface{}) *RealTimeDialogRequest {
	s.MetaData = v
	return s
}

func (s *RealTimeDialogRequest) SetOpType(v string) *RealTimeDialogRequest {
	s.OpType = &v
	return s
}

func (s *RealTimeDialogRequest) SetRecommend(v bool) *RealTimeDialogRequest {
	s.Recommend = &v
	return s
}

func (s *RealTimeDialogRequest) SetScriptContentPlayed(v string) *RealTimeDialogRequest {
	s.ScriptContentPlayed = &v
	return s
}

func (s *RealTimeDialogRequest) SetSessionId(v string) *RealTimeDialogRequest {
	s.SessionId = &v
	return s
}

func (s *RealTimeDialogRequest) SetStream(v bool) *RealTimeDialogRequest {
	s.Stream = &v
	return s
}

func (s *RealTimeDialogRequest) SetUserVad(v bool) *RealTimeDialogRequest {
	s.UserVad = &v
	return s
}

func (s *RealTimeDialogRequest) Validate() error {
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

type RealTimeDialogRequestConversationModel struct {
	// The start time of the sentence, in milliseconds, relative to the start of the session.
	//
	// example:
	//
	// 5
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// The start time of this sentence.
	//
	// example:
	//
	// 2024-11-08 09:51:16
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// The specific content of the conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 人工客服
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The unique ID of the conversation role. This parameter is **required**.
	//
	// example:
	//
	// 98457834685635
	CustomerId *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	// The ID of the customer service representative. This parameter is **required**.
	//
	// example:
	//
	// 1374683645635
	CustomerServiceId *string `json:"customerServiceId,omitempty" xml:"customerServiceId,omitempty"`
	// The type of the customer service representative. 0: bot, 1: human.
	//
	// example:
	//
	// 0
	CustomerServiceType *string `json:"customerServiceType,omitempty" xml:"customerServiceType,omitempty"`
	// The end time of the sentence, in milliseconds, relative to the start of the session.
	//
	// example:
	//
	// 10
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The intent code.
	//
	// example:
	//
	// 198379874354
	IntentionCode *string `json:"intentionCode,omitempty" xml:"intentionCode,omitempty"`
	// The role. 0 indicates the customer, and 1 indicates the customer service representative.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Role *int32 `json:"role,omitempty" xml:"role,omitempty"`
	// The type of the conversation content. Valid values: text, audio, and image.
	//
	// This parameter is required.
	//
	// example:
	//
	// audio
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RealTimeDialogRequestConversationModel) String() string {
	return dara.Prettify(s)
}

func (s RealTimeDialogRequestConversationModel) GoString() string {
	return s.String()
}

func (s *RealTimeDialogRequestConversationModel) GetBegin() *int32 {
	return s.Begin
}

func (s *RealTimeDialogRequestConversationModel) GetBeginTime() *string {
	return s.BeginTime
}

func (s *RealTimeDialogRequestConversationModel) GetContent() *string {
	return s.Content
}

func (s *RealTimeDialogRequestConversationModel) GetCustomerId() *string {
	return s.CustomerId
}

func (s *RealTimeDialogRequestConversationModel) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *RealTimeDialogRequestConversationModel) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *RealTimeDialogRequestConversationModel) GetEnd() *int32 {
	return s.End
}

func (s *RealTimeDialogRequestConversationModel) GetIntentionCode() *string {
	return s.IntentionCode
}

func (s *RealTimeDialogRequestConversationModel) GetRole() *int32 {
	return s.Role
}

func (s *RealTimeDialogRequestConversationModel) GetType() *string {
	return s.Type
}

func (s *RealTimeDialogRequestConversationModel) SetBegin(v int32) *RealTimeDialogRequestConversationModel {
	s.Begin = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetBeginTime(v string) *RealTimeDialogRequestConversationModel {
	s.BeginTime = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetContent(v string) *RealTimeDialogRequestConversationModel {
	s.Content = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetCustomerId(v string) *RealTimeDialogRequestConversationModel {
	s.CustomerId = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetCustomerServiceId(v string) *RealTimeDialogRequestConversationModel {
	s.CustomerServiceId = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetCustomerServiceType(v string) *RealTimeDialogRequestConversationModel {
	s.CustomerServiceType = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetEnd(v int32) *RealTimeDialogRequestConversationModel {
	s.End = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetIntentionCode(v string) *RealTimeDialogRequestConversationModel {
	s.IntentionCode = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetRole(v int32) *RealTimeDialogRequestConversationModel {
	s.Role = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) SetType(v string) *RealTimeDialogRequestConversationModel {
	s.Type = &v
	return s
}

func (s *RealTimeDialogRequestConversationModel) Validate() error {
	return dara.Validate(s)
}
