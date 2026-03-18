// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITeacherExpansionDialogueSuggestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackground(v string) *GetAITeacherExpansionDialogueSuggestionRequest
	GetBackground() *string
	SetDialogueTasks(v []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) *GetAITeacherExpansionDialogueSuggestionRequest
	GetDialogueTasks() []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks
	SetLanguageCode(v string) *GetAITeacherExpansionDialogueSuggestionRequest
	GetLanguageCode() *string
	SetRecords(v []*GetAITeacherExpansionDialogueSuggestionRequestRecords) *GetAITeacherExpansionDialogueSuggestionRequest
	GetRecords() []*GetAITeacherExpansionDialogueSuggestionRequestRecords
	SetRoleInfo(v *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) *GetAITeacherExpansionDialogueSuggestionRequest
	GetRoleInfo() *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo
	SetStartSentence(v string) *GetAITeacherExpansionDialogueSuggestionRequest
	GetStartSentence() *string
	SetTopic(v string) *GetAITeacherExpansionDialogueSuggestionRequest
	GetTopic() *string
	SetUserId(v string) *GetAITeacherExpansionDialogueSuggestionRequest
	GetUserId() *string
}

type GetAITeacherExpansionDialogueSuggestionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// In a career counseling session, we are going to discuss our dream jobs and the responsibilities associated with them. Alex, who dreams of becoming a professional travel blogger, will share the tasks and skills required for this role, while Jamie, aspiring to be a wildlife photographer, will outline the responsibilities and challenges of capturing nature\\"s moments. Both will explore how their interests align with the practical aspects of their chosen careers, discussing the potential for travel, creativity, and the impact of their work on society and the environment.
	Background *string `json:"background,omitempty" xml:"background,omitempty"`
	// This parameter is required.
	DialogueTasks []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*GetAITeacherExpansionDialogueSuggestionRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	RoleInfo *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo `json:"roleInfo,omitempty" xml:"roleInfo,omitempty" type:"Struct"`
	// example:
	//
	// Hello Lily, could you please come to the kitchen for a moment?
	StartSentence *string `json:"startSentence,omitempty" xml:"startSentence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Let\\"s talk about traffic rules.
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequest) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetBackground() *string {
	return s.Background
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetDialogueTasks() []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	return s.DialogueTasks
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetRecords() []*GetAITeacherExpansionDialogueSuggestionRequestRecords {
	return s.Records
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetRoleInfo() *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	return s.RoleInfo
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetStartSentence() *string {
	return s.StartSentence
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetTopic() *string {
	return s.Topic
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetBackground(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Background = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetDialogueTasks(v []*GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.DialogueTasks = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetLanguageCode(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.LanguageCode = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetRecords(v []*GetAITeacherExpansionDialogueSuggestionRequestRecords) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Records = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetRoleInfo(v *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.RoleInfo = v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetStartSentence(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.StartSentence = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetTopic(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.Topic = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) SetUserId(v string) *GetAITeacherExpansionDialogueSuggestionRequest {
	s.UserId = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequest) Validate() error {
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

type GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks struct {
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

func (s GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GetAssistant() *string {
	return s.Assistant
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GetAssistantTranslate() *string {
	return s.AssistantTranslate
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GetOrder() *int32 {
	return s.Order
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) GetUser() *string {
	return s.User
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetAssistant(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetAssistantTranslate(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetOrder(v int32) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) SetUser(v string) *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks {
	s.User = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestDialogueTasks) Validate() error {
	return dara.Validate(s)
}

type GetAITeacherExpansionDialogueSuggestionRequestRecords struct {
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

func (s GetAITeacherExpansionDialogueSuggestionRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRecords) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) GetContent() *string {
	return s.Content
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) GetIsOffTopicControl() *bool {
	return s.IsOffTopicControl
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) GetIsOnTopic() *bool {
	return s.IsOnTopic
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) GetOrder() *int32 {
	return s.Order
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) GetRole() *string {
	return s.Role
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetContent(v string) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Content = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetIsOffTopicControl(v bool) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetIsOnTopic(v bool) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetOrder(v int32) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Order = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) SetRole(v string) *GetAITeacherExpansionDialogueSuggestionRequestRecords {
	s.Role = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRecords) Validate() error {
	return dara.Validate(s)
}

type GetAITeacherExpansionDialogueSuggestionRequestRoleInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// Alex
	Assistant *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Jamie
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) GoString() string {
	return s.String()
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) GetAssistant() *string {
	return s.Assistant
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) GetUser() *string {
	return s.User
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) SetAssistant(v string) *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) SetUser(v string) *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo {
	s.User = &v
	return s
}

func (s *GetAITeacherExpansionDialogueSuggestionRequestRoleInfo) Validate() error {
	return dara.Validate(s)
}
