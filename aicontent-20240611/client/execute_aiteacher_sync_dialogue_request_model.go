// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherSyncDialogueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueRequestDialogueTasks) *ExecuteAITeacherSyncDialogueRequest
  GetDialogueTasks() []*ExecuteAITeacherSyncDialogueRequestDialogueTasks 
  SetLanguageCode(v string) *ExecuteAITeacherSyncDialogueRequest
  GetLanguageCode() *string 
  SetRecords(v []*ExecuteAITeacherSyncDialogueRequestRecords) *ExecuteAITeacherSyncDialogueRequest
  GetRecords() []*ExecuteAITeacherSyncDialogueRequestRecords 
  SetUserId(v string) *ExecuteAITeacherSyncDialogueRequest
  GetUserId() *string 
}

type ExecuteAITeacherSyncDialogueRequest struct {
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherSyncDialogueRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  // example:
  // 
  // en-gb
  LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
  Records []*ExecuteAITeacherSyncDialogueRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 886eba3702xxxxxxxxx4ba52a87a525
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequest) GetDialogueTasks() []*ExecuteAITeacherSyncDialogueRequestDialogueTasks  {
  return s.DialogueTasks
}

func (s *ExecuteAITeacherSyncDialogueRequest) GetLanguageCode() *string  {
  return s.LanguageCode
}

func (s *ExecuteAITeacherSyncDialogueRequest) GetRecords() []*ExecuteAITeacherSyncDialogueRequestRecords  {
  return s.Records
}

func (s *ExecuteAITeacherSyncDialogueRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueRequestDialogueTasks) *ExecuteAITeacherSyncDialogueRequest {
  s.DialogueTasks = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetLanguageCode(v string) *ExecuteAITeacherSyncDialogueRequest {
  s.LanguageCode = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetRecords(v []*ExecuteAITeacherSyncDialogueRequestRecords) *ExecuteAITeacherSyncDialogueRequest {
  s.Records = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) SetUserId(v string) *ExecuteAITeacherSyncDialogueRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequest) Validate() error {
  if s.DialogueTasks != nil {
    for _, item := range s.DialogueTasks {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Records != nil {
    for _, item := range s.Records {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExecuteAITeacherSyncDialogueRequestDialogueTasks struct {
  // This parameter is required.
  // 
  // example:
  // 
  // Why might some people think dog walking is a great job?
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // example:
  // 
  // 为什么有些人认为遛狗是份好差事?
  AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // They think it\\"s great because they won\\"t be stuck in an office.
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequestDialogueTasks) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequestDialogueTasks) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) GetAssistantTranslate() *string  {
  return s.AssistantTranslate
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
  s.AssistantTranslate = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherSyncDialogueRequestDialogueTasks {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestDialogueTasks) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherSyncDialogueRequestRecords struct {
  // This parameter is required.
  // 
  // example:
  // 
  // Ask Mark if he has thought about what his dream job might be.
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // example:
  // 
  // 跑题：true, 不跑题：false
  IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
  // example:
  // 
  // 扣题：true, 不扣题：false
  IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 老师：assistant；学生：user
  Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueRequestRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueRequestRecords) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) GetIsOffTopicControl() *bool  {
  return s.IsOffTopicControl
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) GetRole() *string  {
  return s.Role
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetContent(v string) *ExecuteAITeacherSyncDialogueRequestRecords {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherSyncDialogueRequestRecords {
  s.IsOffTopicControl = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueRequestRecords {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetOrder(v int32) *ExecuteAITeacherSyncDialogueRequestRecords {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) SetRole(v string) *ExecuteAITeacherSyncDialogueRequestRecords {
  s.Role = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueRequestRecords) Validate() error {
  return dara.Validate(s)
}

