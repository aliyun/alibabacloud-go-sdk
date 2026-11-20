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
  // The collection of conversation content for dialogue tasks.
  // 
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherSyncDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  // The collection of conversation records.
  Records []*ExecuteAITeacherSyncDialogueTranslateRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // The user ID.
  // 
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
  // The AI conversation content.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Why might some people think dog walking is a great job?
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // The translation of the AI conversation content.
  // 
  // example:
  // 
  // Why do some people think walking dogs is a good job?
  AssistantTranslate *string `json:"assistantTranslate,omitempty" xml:"assistantTranslate,omitempty"`
  // The sort order.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
  // The user conversation content.
  // 
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
  // The conversation content.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Ask Mark if he has thought about what his dream job might be.
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // Specifies whether the content is off topic for flow control purposes. Valid values:
  // 
  // - true: Off topic.
  // 
  // - false: On topic.
  // 
  // example:
  // 
  // true
  IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
  // Specifies whether the content is on topic. Valid values:
  // 
  // - true: On topic.
  // 
  // - false: Off topic.
  // 
  // example:
  // 
  // true
  IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
  // The sort order.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
  // The role.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // AI: assistant; User: user
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

