// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRealtimeDialogAssistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnalysis(v bool) *RealtimeDialogAssistRequest
	GetAnalysis() *bool
	SetBizType(v string) *RealtimeDialogAssistRequest
	GetBizType() *string
	SetConversationModel(v []*RealtimeDialogAssistRequestConversationModel) *RealtimeDialogAssistRequest
	GetConversationModel() []*RealtimeDialogAssistRequestConversationModel
	SetDialogMemoryTurns(v int32) *RealtimeDialogAssistRequest
	GetDialogMemoryTurns() *int32
	SetHangUpDialog(v bool) *RealtimeDialogAssistRequest
	GetHangUpDialog() *bool
	SetMetaData(v map[string]interface{}) *RealtimeDialogAssistRequest
	GetMetaData() map[string]interface{}
	SetRequestId(v string) *RealtimeDialogAssistRequest
	GetRequestId() *string
	SetScriptContentPlayed(v string) *RealtimeDialogAssistRequest
	GetScriptContentPlayed() *string
	SetSessionId(v string) *RealtimeDialogAssistRequest
	GetSessionId() *string
	SetUserVad(v bool) *RealtimeDialogAssistRequest
	GetUserVad() *bool
}

type RealtimeDialogAssistRequest struct {
	// example:
	//
	// false
	Analysis *bool `json:"analysis,omitempty" xml:"analysis,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dialogAssist
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// This parameter is required.
	ConversationModel []*RealtimeDialogAssistRequestConversationModel `json:"conversationModel,omitempty" xml:"conversationModel,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	DialogMemoryTurns *int32 `json:"dialogMemoryTurns,omitempty" xml:"dialogMemoryTurns,omitempty"`
	// example:
	//
	// false
	HangUpDialog *bool `json:"hangUpDialog,omitempty" xml:"hangUpDialog,omitempty"`
	// metaData
	MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0FC6636E-380A-5369-AE01-D1C15BB9B254
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 你好
	ScriptContentPlayed *string `json:"scriptContentPlayed,omitempty" xml:"scriptContentPlayed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1915593248420413441
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	UserVad   *bool   `json:"userVad,omitempty" xml:"userVad,omitempty"`
}

func (s RealtimeDialogAssistRequest) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistRequest) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistRequest) GetAnalysis() *bool {
	return s.Analysis
}

func (s *RealtimeDialogAssistRequest) GetBizType() *string {
	return s.BizType
}

func (s *RealtimeDialogAssistRequest) GetConversationModel() []*RealtimeDialogAssistRequestConversationModel {
	return s.ConversationModel
}

func (s *RealtimeDialogAssistRequest) GetDialogMemoryTurns() *int32 {
	return s.DialogMemoryTurns
}

func (s *RealtimeDialogAssistRequest) GetHangUpDialog() *bool {
	return s.HangUpDialog
}

func (s *RealtimeDialogAssistRequest) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *RealtimeDialogAssistRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *RealtimeDialogAssistRequest) GetScriptContentPlayed() *string {
	return s.ScriptContentPlayed
}

func (s *RealtimeDialogAssistRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RealtimeDialogAssistRequest) GetUserVad() *bool {
	return s.UserVad
}

func (s *RealtimeDialogAssistRequest) SetAnalysis(v bool) *RealtimeDialogAssistRequest {
	s.Analysis = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetBizType(v string) *RealtimeDialogAssistRequest {
	s.BizType = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetConversationModel(v []*RealtimeDialogAssistRequestConversationModel) *RealtimeDialogAssistRequest {
	s.ConversationModel = v
	return s
}

func (s *RealtimeDialogAssistRequest) SetDialogMemoryTurns(v int32) *RealtimeDialogAssistRequest {
	s.DialogMemoryTurns = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetHangUpDialog(v bool) *RealtimeDialogAssistRequest {
	s.HangUpDialog = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetMetaData(v map[string]interface{}) *RealtimeDialogAssistRequest {
	s.MetaData = v
	return s
}

func (s *RealtimeDialogAssistRequest) SetRequestId(v string) *RealtimeDialogAssistRequest {
	s.RequestId = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetScriptContentPlayed(v string) *RealtimeDialogAssistRequest {
	s.ScriptContentPlayed = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetSessionId(v string) *RealtimeDialogAssistRequest {
	s.SessionId = &v
	return s
}

func (s *RealtimeDialogAssistRequest) SetUserVad(v bool) *RealtimeDialogAssistRequest {
	s.UserVad = &v
	return s
}

func (s *RealtimeDialogAssistRequest) Validate() error {
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

type RealtimeDialogAssistRequestConversationModel struct {
	// example:
	//
	// 1
	Begin *int32 `json:"begin,omitempty" xml:"begin,omitempty"`
	// example:
	//
	// 2025-12-12 09:00:00
	BeginTime *string `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
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
	// 1
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Role *int32 `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// text
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s RealtimeDialogAssistRequestConversationModel) String() string {
	return dara.Prettify(s)
}

func (s RealtimeDialogAssistRequestConversationModel) GoString() string {
	return s.String()
}

func (s *RealtimeDialogAssistRequestConversationModel) GetBegin() *int32 {
	return s.Begin
}

func (s *RealtimeDialogAssistRequestConversationModel) GetBeginTime() *string {
	return s.BeginTime
}

func (s *RealtimeDialogAssistRequestConversationModel) GetContent() *string {
	return s.Content
}

func (s *RealtimeDialogAssistRequestConversationModel) GetCustomerId() *string {
	return s.CustomerId
}

func (s *RealtimeDialogAssistRequestConversationModel) GetCustomerServiceId() *string {
	return s.CustomerServiceId
}

func (s *RealtimeDialogAssistRequestConversationModel) GetCustomerServiceType() *string {
	return s.CustomerServiceType
}

func (s *RealtimeDialogAssistRequestConversationModel) GetEnd() *int32 {
	return s.End
}

func (s *RealtimeDialogAssistRequestConversationModel) GetRole() *int32 {
	return s.Role
}

func (s *RealtimeDialogAssistRequestConversationModel) GetType() *string {
	return s.Type
}

func (s *RealtimeDialogAssistRequestConversationModel) SetBegin(v int32) *RealtimeDialogAssistRequestConversationModel {
	s.Begin = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetBeginTime(v string) *RealtimeDialogAssistRequestConversationModel {
	s.BeginTime = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetContent(v string) *RealtimeDialogAssistRequestConversationModel {
	s.Content = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetCustomerId(v string) *RealtimeDialogAssistRequestConversationModel {
	s.CustomerId = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetCustomerServiceId(v string) *RealtimeDialogAssistRequestConversationModel {
	s.CustomerServiceId = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetCustomerServiceType(v string) *RealtimeDialogAssistRequestConversationModel {
	s.CustomerServiceType = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetEnd(v int32) *RealtimeDialogAssistRequestConversationModel {
	s.End = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetRole(v int32) *RealtimeDialogAssistRequestConversationModel {
	s.Role = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) SetType(v string) *RealtimeDialogAssistRequestConversationModel {
	s.Type = &v
	return s
}

func (s *RealtimeDialogAssistRequestConversationModel) Validate() error {
	return dara.Validate(s)
}
