// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueTranslateRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBackground(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetBackground() *string 
  SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetDialogueTasks() []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks 
  SetRecords(v []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetRecords() []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords 
  SetRoleInfo(v *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetRoleInfo() *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo 
  SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetStartSentence() *string 
  SetTopic(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetTopic() *string 
  SetUserId(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest
  GetUserId() *string 
}

type ExecuteAITeacherExpansionDialogueTranslateRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // In this dialogue, you will be playing the role of Lily, a young girl. I will be Jane, Lily\\"s mother. We are in the kitchen, where I am preparing dinner. I am asking you about your food preferences, specifically if you like meat, fish, and milk. You like meat and milk, but you don\\"t like fish because of its smell. I explain to you the nutritional benefits of these foods and suggest alternatives for the ones you don\\"t like. Finally, I invite you to start eating.
  Background *string `json:"background,omitempty" xml:"background,omitempty"`
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  Records []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // This parameter is required.
  RoleInfo *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
  // example:
  // 
  // Hello Lily, could you please come to the kitchen for a moment?
  StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // talk about food.
  Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 886eba3702xxxxxxxxx4ba52a87a525
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetBackground() *string  {
  return s.Background
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetDialogueTasks() []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks  {
  return s.DialogueTasks
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetRecords() []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords  {
  return s.Records
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetRoleInfo() *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo  {
  return s.RoleInfo
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetStartSentence() *string  {
  return s.StartSentence
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetTopic() *string  {
  return s.Topic
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.Background = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.DialogueTasks = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueTranslateRequestRecords) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.Records = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.RoleInfo = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.StartSentence = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.Topic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueTranslateRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequest) Validate() error {
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
  if s.RoleInfo != nil {
    if err := s.RoleInfo.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks struct {
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

func (s ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GetAssistantTranslate() *string  {
  return s.AssistantTranslate
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
  s.AssistantTranslate = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestDialogueTasks) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherExpansionDialogueTranslateRequestRecords struct {
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

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GetIsOffTopicControl() *bool  {
  return s.IsOffTopicControl
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) GetRole() *string  {
  return s.Role
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
  s.IsOffTopicControl = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRecords {
  s.Role = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRecords) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo struct {
  // This parameter is required.
  // 
  // example:
  // 
  // Jane, a caring mother
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // Lily, a friendly student
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueTranslateRequestRoleInfo) Validate() error {
  return dara.Validate(s)
}

