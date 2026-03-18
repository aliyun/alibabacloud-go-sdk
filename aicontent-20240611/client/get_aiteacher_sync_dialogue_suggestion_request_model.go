// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITeacherSyncDialogueSuggestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogueTasks(v []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) *GetAITeacherSyncDialogueSuggestionRequest
	GetDialogueTasks() []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks
	SetLanguageCode(v string) *GetAITeacherSyncDialogueSuggestionRequest
	GetLanguageCode() *string
	SetRecords(v []*GetAITeacherSyncDialogueSuggestionRequestRecords) *GetAITeacherSyncDialogueSuggestionRequest
	GetRecords() []*GetAITeacherSyncDialogueSuggestionRequestRecords
	SetUserId(v string) *GetAITeacherSyncDialogueSuggestionRequest
	GetUserId() *string
}

type GetAITeacherSyncDialogueSuggestionRequest struct {
	// This parameter is required.
	DialogueTasks []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks `json:"dialogueTasks,omitempty" xml:"dialogueTasks,omitempty" type:"Repeated"`
	// example:
	//
	// en-gb
	LanguageCode *string `json:"languageCode,omitempty" xml:"languageCode,omitempty"`
	// This parameter is required.
	Records []*GetAITeacherSyncDialogueSuggestionRequestRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 886eba3702xxxxxxxxx4ba52a87a525
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetAITeacherSyncDialogueSuggestionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequest) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) GetDialogueTasks() []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	return s.DialogueTasks
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) GetRecords() []*GetAITeacherSyncDialogueSuggestionRequestRecords {
	return s.Records
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetDialogueTasks(v []*GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) *GetAITeacherSyncDialogueSuggestionRequest {
	s.DialogueTasks = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetLanguageCode(v string) *GetAITeacherSyncDialogueSuggestionRequest {
	s.LanguageCode = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetRecords(v []*GetAITeacherSyncDialogueSuggestionRequestRecords) *GetAITeacherSyncDialogueSuggestionRequest {
	s.Records = v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) SetUserId(v string) *GetAITeacherSyncDialogueSuggestionRequest {
	s.UserId = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequest) Validate() error {
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

type GetAITeacherSyncDialogueSuggestionRequestDialogueTasks struct {
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

func (s GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GetAssistant() *string {
	return s.Assistant
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GetAssistantTranslate() *string {
	return s.AssistantTranslate
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GetOrder() *int32 {
	return s.Order
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) GetUser() *string {
	return s.User
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetAssistant(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.Assistant = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetAssistantTranslate(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.AssistantTranslate = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetOrder(v int32) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.Order = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) SetUser(v string) *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks {
	s.User = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestDialogueTasks) Validate() error {
	return dara.Validate(s)
}

type GetAITeacherSyncDialogueSuggestionRequestRecords struct {
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

func (s GetAITeacherSyncDialogueSuggestionRequestRecords) String() string {
	return dara.Prettify(s)
}

func (s GetAITeacherSyncDialogueSuggestionRequestRecords) GoString() string {
	return s.String()
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) GetContent() *string {
	return s.Content
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) GetIsOffTopicControl() *bool {
	return s.IsOffTopicControl
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) GetIsOnTopic() *bool {
	return s.IsOnTopic
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) GetOrder() *int32 {
	return s.Order
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) GetRole() *string {
	return s.Role
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetContent(v string) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Content = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetIsOffTopicControl(v bool) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.IsOffTopicControl = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetIsOnTopic(v bool) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.IsOnTopic = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetOrder(v int32) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Order = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) SetRole(v string) *GetAITeacherSyncDialogueSuggestionRequestRecords {
	s.Role = &v
	return s
}

func (s *GetAITeacherSyncDialogueSuggestionRequestRecords) Validate() error {
	return dara.Validate(s)
}
