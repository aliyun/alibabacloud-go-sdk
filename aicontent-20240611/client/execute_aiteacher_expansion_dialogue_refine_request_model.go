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
  // The conversation background description.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
  Background *string `json:"background,omitempty" xml:"background,omitempty"`
  // The conversation task content collection.
  // 
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherExpansionDialogueRefineRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  // The language code.
  // 
  // example:
  // 
  // en-gb
  LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
  // The conversation records.
  // 
  // This parameter is required.
  Records []*ExecuteAITeacherExpansionDialogueRefineRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // The role settings.
  // 
  // This parameter is required.
  RoleInfo *ExecuteAITeacherExpansionDialogueRefineRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
  // The opening sentence.
  // 
  // example:
  // 
  // Hello Lily, could you please come to the kitchen for a moment?
  StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
  // The topic.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // talk about your dream job.
  Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
  // The user ID.
  // 
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
  // The sequence number.
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
  // The conversation content.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Ask Mark if he has thought about what his dream job might be.
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // Specifies whether the response is off topic and controls the conversation flow. This parameter indicates whether the student\\"s response meets the requirements of the conversation task. The value of this parameter is derived from the return results of previous conversation turns. Additionally, this parameter controls the conversation flow. If the response is off topic more than twice, the value is forcibly set to true to switch the task. Valid values:
  // 
  // - true: Off topic.
  // 
  // - false: On topic.
  // 
  // example:
  // 
  // true
  IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
  // Specifies whether the response is on topic. Valid values:
  // 
  // - true: On topic.
  // 
  // - false: Off topic.
  // 
  // example:
  // 
  // true
  IsOnTopic *bool `json:"isOnTopic,omitempty" xml:"isOnTopic,omitempty"`
  // The sequence number.
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
  // The name of the AI.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Jane, a caring mother
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // The name of the user.
  // 
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

