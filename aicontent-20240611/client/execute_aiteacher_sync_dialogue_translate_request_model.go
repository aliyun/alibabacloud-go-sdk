// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherSyncDialogueTranslateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherSyncDialogueTranslateRequest
  GetDialogueTasks() []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks 
  SetRecords(v []*ExecuteAITeacherSyncDialogueTranslateRequestRecords) *ExecuteAITeacherSyncDialogueTranslateRequest
  GetRecords() []*ExecuteAITeacherSyncDialogueTranslateRequestRecords 
  SetUserId(v string) *ExecuteAITeacherSyncDialogueTranslateRequest
  GetUserId() *string 
}

type ExecuteAITeacherSyncDialogueTranslateRequest struct {
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  Records []*ExecuteAITeacherSyncDialogueTranslateRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // This parameter is required.
  // 
  // example:
  // 
  // 886eba3702xxxxxxxxx4ba52a87a525
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherSyncDialogueTranslateRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) GetDialogueTasks() []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks  {
  return s.DialogueTasks
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) GetRecords() []*ExecuteAITeacherSyncDialogueTranslateRequestRecords  {
  return s.Records
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetDialogueTasks(v []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherSyncDialogueTranslateRequest {
  s.DialogueTasks = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetRecords(v []*ExecuteAITeacherSyncDialogueTranslateRequestRecords) *ExecuteAITeacherSyncDialogueTranslateRequest {
  s.Records = v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) SetUserId(v string) *ExecuteAITeacherSyncDialogueTranslateRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequest) Validate() error {
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

type ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks struct {
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

func (s ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GetAssistantTranslate() *string  {
  return s.AssistantTranslate
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
  s.AssistantTranslate = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherSyncDialogueTranslateRequestRecords struct {
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

func (s ExecuteAITeacherSyncDialogueTranslateRequestRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherSyncDialogueTranslateRequestRecords) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) GetIsOffTopicControl() *bool  {
  return s.IsOffTopicControl
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) GetRole() *string  {
  return s.Role
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetContent(v string) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
  s.IsOffTopicControl = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetOrder(v int32) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) SetRole(v string) *ExecuteAITeacherSyncDialogueTranslateRequestRecords {
  s.Role = &v
  return s
}

func (s *ExecuteAITeacherSyncDialogueTranslateRequestRecords) Validate() error {
  return dara.Validate(s)
}

