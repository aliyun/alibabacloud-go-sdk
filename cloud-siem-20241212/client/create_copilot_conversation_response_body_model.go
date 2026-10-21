// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiVersion(v string) *CreateCopilotConversationResponseBody
	GetApiVersion() *string
	SetConversation(v *CreateCopilotConversationResponseBodyConversation) *CreateCopilotConversationResponseBody
	GetConversation() *CreateCopilotConversationResponseBodyConversation
	SetRequestId(v string) *CreateCopilotConversationResponseBody
	GetRequestId() *string
}

type CreateCopilotConversationResponseBody struct {
	// The API version number. The value is 1.0.0.
	//
	// example:
	//
	// v1.2
	ApiVersion *string `json:"ApiVersion,omitempty" xml:"ApiVersion,omitempty"`
	// The conversation timestamp.
	Conversation *CreateCopilotConversationResponseBodyConversation `json:"Conversation,omitempty" xml:"Conversation,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCopilotConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotConversationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCopilotConversationResponseBody) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CreateCopilotConversationResponseBody) GetConversation() *CreateCopilotConversationResponseBodyConversation {
	return s.Conversation
}

func (s *CreateCopilotConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCopilotConversationResponseBody) SetApiVersion(v string) *CreateCopilotConversationResponseBody {
	s.ApiVersion = &v
	return s
}

func (s *CreateCopilotConversationResponseBody) SetConversation(v *CreateCopilotConversationResponseBodyConversation) *CreateCopilotConversationResponseBody {
	s.Conversation = v
	return s
}

func (s *CreateCopilotConversationResponseBody) SetRequestId(v string) *CreateCopilotConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCopilotConversationResponseBody) Validate() error {
	if s.Conversation != nil {
		if err := s.Conversation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCopilotConversationResponseBodyConversation struct {
	// The security operations Agent session ID.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The creation time. The time follows the ISO 8601 standard in UTC.
	//
	// example:
	//
	// 2026-01-31T10:10:20Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The returned conversation information.
	//
	// example:
	//
	// 1
	Feedback *int64 `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The model configuration.
	//
	// example:
	//
	// qwen3.7-max
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether planning is enabled. Valid values: disable, enable, force. Default value: enable.
	//
	// example:
	//
	// enabled
	PlanMode *string `json:"PlanMode,omitempty" xml:"PlanMode,omitempty"`
	// The security project ID.
	//
	// example:
	//
	// project-incident-****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the scheduled task. The ID is system-generated and globally unique.
	//
	// example:
	//
	// task-e34f770da4f04f6db21c2819****
	ScheduledTaskId *string `json:"ScheduledTaskId,omitempty" xml:"ScheduledTaskId,omitempty"`
	// The status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The thinking mode configuration.
	//
	// example:
	//
	// false
	ThinkingMode *string `json:"ThinkingMode,omitempty" xml:"ThinkingMode,omitempty"`
	// The session title.
	//
	// example:
	//
	// Security incident investigation
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-11-27 16:02:28
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The session visibility.
	//
	// example:
	//
	// []
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s CreateCopilotConversationResponseBodyConversation) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotConversationResponseBodyConversation) GoString() string {
	return s.String()
}

func (s *CreateCopilotConversationResponseBodyConversation) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateCopilotConversationResponseBodyConversation) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCopilotConversationResponseBodyConversation) GetFeedback() *int64 {
	return s.Feedback
}

func (s *CreateCopilotConversationResponseBodyConversation) GetModel() *string {
	return s.Model
}

func (s *CreateCopilotConversationResponseBodyConversation) GetPlanMode() *string {
	return s.PlanMode
}

func (s *CreateCopilotConversationResponseBodyConversation) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateCopilotConversationResponseBodyConversation) GetScheduledTaskId() *string {
	return s.ScheduledTaskId
}

func (s *CreateCopilotConversationResponseBodyConversation) GetStatus() *string {
	return s.Status
}

func (s *CreateCopilotConversationResponseBodyConversation) GetThinkingMode() *string {
	return s.ThinkingMode
}

func (s *CreateCopilotConversationResponseBodyConversation) GetTitle() *string {
	return s.Title
}

func (s *CreateCopilotConversationResponseBodyConversation) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCopilotConversationResponseBodyConversation) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateCopilotConversationResponseBodyConversation) SetConversationId(v string) *CreateCopilotConversationResponseBodyConversation {
	s.ConversationId = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetCreatedAt(v string) *CreateCopilotConversationResponseBodyConversation {
	s.CreatedAt = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetFeedback(v int64) *CreateCopilotConversationResponseBodyConversation {
	s.Feedback = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetModel(v string) *CreateCopilotConversationResponseBodyConversation {
	s.Model = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetPlanMode(v string) *CreateCopilotConversationResponseBodyConversation {
	s.PlanMode = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetProjectId(v string) *CreateCopilotConversationResponseBodyConversation {
	s.ProjectId = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetScheduledTaskId(v string) *CreateCopilotConversationResponseBodyConversation {
	s.ScheduledTaskId = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetStatus(v string) *CreateCopilotConversationResponseBodyConversation {
	s.Status = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetThinkingMode(v string) *CreateCopilotConversationResponseBodyConversation {
	s.ThinkingMode = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetTitle(v string) *CreateCopilotConversationResponseBodyConversation {
	s.Title = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetUpdatedAt(v string) *CreateCopilotConversationResponseBodyConversation {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) SetVisibility(v string) *CreateCopilotConversationResponseBodyConversation {
	s.Visibility = &v
	return s
}

func (s *CreateCopilotConversationResponseBodyConversation) Validate() error {
	return dara.Validate(s)
}
