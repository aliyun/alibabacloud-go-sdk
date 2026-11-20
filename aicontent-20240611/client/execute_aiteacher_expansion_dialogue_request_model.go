// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteAITeacherExpansionDialogueRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBackground(v string) *ExecuteAITeacherExpansionDialogueRequest
  GetBackground() *string 
  SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRequest
  GetDialogueTasks() []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks 
  SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRequest
  GetLanguageCode() *string 
  SetRecords(v []*ExecuteAITeacherExpansionDialogueRequestRecords) *ExecuteAITeacherExpansionDialogueRequest
  GetRecords() []*ExecuteAITeacherExpansionDialogueRequestRecords 
  SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRequest
  GetRoleInfo() *ExecuteAITeacherExpansionDialogueRequestRoleInfo 
  SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRequest
  GetStartSentence() *string 
  SetTopic(v string) *ExecuteAITeacherExpansionDialogueRequest
  GetTopic() *string 
  SetUserId(v string) *ExecuteAITeacherExpansionDialogueRequest
  GetUserId() *string 
}

type ExecuteAITeacherExpansionDialogueRequest struct {
  // The conversation background description.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
  Background *string `json:"background,omitempty" xml:"background,omitempty"`
  // The collection of conversation task content.
  // 
  // This parameter is required.
  DialogueTasks []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
  // The language code.
  // 
  // example:
  // 
  // en-gb
  LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
  // The collection of conversation records.
  Records []*ExecuteAITeacherExpansionDialogueRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
  // The role settings.
  // 
  // This parameter is required.
  RoleInfo *ExecuteAITeacherExpansionDialogueRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
  // The opening remark.
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
  // Let\\"s talk about traffic rules.
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

func (s ExecuteAITeacherExpansionDialogueRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequest) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetBackground() *string  {
  return s.Background
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetDialogueTasks() []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks  {
  return s.DialogueTasks
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetLanguageCode() *string  {
  return s.LanguageCode
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetRecords() []*ExecuteAITeacherExpansionDialogueRequestRecords  {
  return s.Records
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetRoleInfo() *ExecuteAITeacherExpansionDialogueRequestRoleInfo  {
  return s.RoleInfo
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetStartSentence() *string  {
  return s.StartSentence
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetTopic() *string  {
  return s.Topic
}

func (s *ExecuteAITeacherExpansionDialogueRequest) GetUserId() *string  {
  return s.UserId
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetBackground(v string) *ExecuteAITeacherExpansionDialogueRequest {
  s.Background = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetDialogueTasks(v []*ExecuteAITeacherExpansionDialogueRequestDialogueTasks) *ExecuteAITeacherExpansionDialogueRequest {
  s.DialogueTasks = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetLanguageCode(v string) *ExecuteAITeacherExpansionDialogueRequest {
  s.LanguageCode = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetRecords(v []*ExecuteAITeacherExpansionDialogueRequestRecords) *ExecuteAITeacherExpansionDialogueRequest {
  s.Records = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetRoleInfo(v *ExecuteAITeacherExpansionDialogueRequestRoleInfo) *ExecuteAITeacherExpansionDialogueRequest {
  s.RoleInfo = v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetStartSentence(v string) *ExecuteAITeacherExpansionDialogueRequest {
  s.StartSentence = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetTopic(v string) *ExecuteAITeacherExpansionDialogueRequest {
  s.Topic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) SetUserId(v string) *ExecuteAITeacherExpansionDialogueRequest {
  s.UserId = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequest) Validate() error {
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

type ExecuteAITeacherExpansionDialogueRequestDialogueTasks struct {
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

func (s ExecuteAITeacherExpansionDialogueRequestDialogueTasks) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GetAssistantTranslate() *string  {
  return s.AssistantTranslate
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetAssistantTranslate(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
  s.AssistantTranslate = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) SetUser(v string) *ExecuteAITeacherExpansionDialogueRequestDialogueTasks {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestDialogueTasks) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherExpansionDialogueRequestRecords struct {
  // The conversation content.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Ask Mark if he has thought about what his dream job might be.
  Content *string `json:"content,omitempty" xml:"content,omitempty"`
  // Specifies whether the response is off-topic, which also serves for flow control. This parameter reflects whether the student\\"s response complies with the settings of the conversation task. The value of this parameter is derived from the return results of previous conversation turns. In addition, this parameter serves for flow control. If the user goes off-topic more than twice, the value is forcibly set to true to force a task switch. Valid values:
  // 
  // - true: Off-topic.
  // 
  // - false: On-topic.
  // 
  // example:
  // 
  // true
  IsOffTopicControl *bool `json:"isOffTopicControl,omitempty" xml:"isOffTopicControl,omitempty"`
  // Specifies whether the response is on-topic. Valid values:
  // 
  // - true: On-topic.
  // 
  // - false: Off-topic.
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

func (s ExecuteAITeacherExpansionDialogueRequestRecords) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestRecords) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) GetContent() *string  {
  return s.Content
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) GetIsOffTopicControl() *bool  {
  return s.IsOffTopicControl
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) GetIsOnTopic() *bool  {
  return s.IsOnTopic
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) GetOrder() *int32  {
  return s.Order
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) GetRole() *string  {
  return s.Role
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetContent(v string) *ExecuteAITeacherExpansionDialogueRequestRecords {
  s.Content = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetIsOffTopicControl(v bool) *ExecuteAITeacherExpansionDialogueRequestRecords {
  s.IsOffTopicControl = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetIsOnTopic(v bool) *ExecuteAITeacherExpansionDialogueRequestRecords {
  s.IsOnTopic = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetOrder(v int32) *ExecuteAITeacherExpansionDialogueRequestRecords {
  s.Order = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) SetRole(v string) *ExecuteAITeacherExpansionDialogueRequestRecords {
  s.Role = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRecords) Validate() error {
  return dara.Validate(s)
}

type ExecuteAITeacherExpansionDialogueRequestRoleInfo struct {
  // The name of the AI.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Alex
  Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
  // The name of the user.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // Jamie
  User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ExecuteAITeacherExpansionDialogueRequestRoleInfo) String() string {
  return dara.Prettify(s)
}

func (s ExecuteAITeacherExpansionDialogueRequestRoleInfo) GoString() string {
  return s.String()
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) GetAssistant() *string  {
  return s.Assistant
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) GetUser() *string  {
  return s.User
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) SetAssistant(v string) *ExecuteAITeacherExpansionDialogueRequestRoleInfo {
  s.Assistant = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) SetUser(v string) *ExecuteAITeacherExpansionDialogueRequestRoleInfo {
  s.User = &v
  return s
}

func (s *ExecuteAITeacherExpansionDialogueRequestRoleInfo) Validate() error {
  return dara.Validate(s)
}

