// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueRefineRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBackground(v string) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetBackground() *string 
  SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetDialogueTasks() []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks 
  SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetLanguageCode() *string 
  SetRecords(v []*ExecuteAITeacherExpansionDialogueRefineRequestRecords) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetRecords() []*ExecuteAITeacherExpansionDialogueRefineRequestRecords 
  SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetRoleInfo() *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo 
  SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetStartSentence() *string 
  SetTopic(v string) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetTopic() *string 
  SetUserId(v string) *ExecuteAITeacherExpansionDialogueRefineRequest
  GetUserId() *string 
}

type ExecuteAITeacherExpansionDialogueRefineRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
  Background *string `json:"background,omitempty" xml:"background,omitempty"`
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  // example:
  // 
  // en-gb
  LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
  // This parameter is required.
  Records []*ExecuteAITeacherExpansionDialogueRefineRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // This parameter is required.
  RoleInfo *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
  // example:
  // 
  // Hello Lily, could you please come to the kitchen for a moment?
  StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // talk about your dream job.
  Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 886eba3702xxxxxxxxx4ba52a87a525
  UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRefineRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetBackground() *string  {
  return s.Background
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetDialogueTasks() []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks  {
  return s.DialogueTasks
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetLanguageCode() *string  {
  return s.LanguageCode
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetRecords() []*ExecuteAITeacherExpansionDialogueRefineRequestRecords  {
  return s.Records
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetRoleInfo() *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo  {
  return s.RoleInfo
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetStartSentence() *string  {
  return s.StartSentence
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetTopic() *string  {
  return s.Topic
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.Background = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.DialogueTasks = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.LanguageCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueRefineRequestRecords) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.Records = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.RoleInfo = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.StartSentence = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.Topic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueRefineRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequest) Validate() error {
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

type ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks struct {
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

func (s ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GetAssistantTranslate() *string  {
  return s.AssistantTranslate
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
  s.AssistantTranslate = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherExpansionDialogueRefineRequestRecords struct {
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

func (s ExecuteAITeacherExpansionDialogueRefineRequestRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRecords) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) GetIsOffTopicControl() *bool  {
  return s.IsOffTopicControl
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) GetRole() *string  {
  return s.Role
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
  s.IsOffTopicControl = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRecords {
  s.Role = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRecords) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo struct {
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

func (s ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo) Validate() error {
  return dara.Validate(s)
}

