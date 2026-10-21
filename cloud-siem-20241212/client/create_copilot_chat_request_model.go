// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientConversationId(v string) *CreateCopilotChatRequest
	GetClientConversationId() *string
	SetClientMessageId(v string) *CreateCopilotChatRequest
	GetClientMessageId() *string
	SetConversationId(v string) *CreateCopilotChatRequest
	GetConversationId() *string
	SetCreateConversation(v *CreateCopilotChatRequestCreateConversation) *CreateCopilotChatRequest
	GetCreateConversation() *CreateCopilotChatRequestCreateConversation
	SetLang(v string) *CreateCopilotChatRequest
	GetLang() *string
	SetMessages(v []*CreateCopilotChatRequestMessages) *CreateCopilotChatRequest
	GetMessages() []*CreateCopilotChatRequestMessages
	SetModel(v string) *CreateCopilotChatRequest
	GetModel() *string
	SetPlanMode(v string) *CreateCopilotChatRequest
	GetPlanMode() *string
	SetProjectId(v string) *CreateCopilotChatRequest
	GetProjectId() *string
	SetRegionId(v string) *CreateCopilotChatRequest
	GetRegionId() *string
	SetReplaceTurnId(v string) *CreateCopilotChatRequest
	GetReplaceTurnId() *string
	SetRouteTarget(v string) *CreateCopilotChatRequest
	GetRouteTarget() *string
	SetSource(v string) *CreateCopilotChatRequest
	GetSource() *string
	SetStream(v bool) *CreateCopilotChatRequest
	GetStream() *bool
	SetThinkingMode(v string) *CreateCopilotChatRequest
	GetThinkingMode() *string
	SetTraceId(v string) *CreateCopilotChatRequest
	GetTraceId() *string
}

type CreateCopilotChatRequest struct {
	// The client conversation ID.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ClientConversationId *string `json:"ClientConversationId,omitempty" xml:"ClientConversationId,omitempty"`
	// The client message ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 366932fb-9a83-4a2d-94dd-c925b35f****
	ClientMessageId *string `json:"ClientMessageId,omitempty" xml:"ClientMessageId,omitempty"`
	// The conversation ID.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The conversation creation configuration.
	CreateConversation *CreateCopilotChatRequestCreateConversation `json:"CreateConversation,omitempty" xml:"CreateConversation,omitempty" type:"Struct"`
	// The language of the returned messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The list of messages.
	//
	// This parameter is required.
	Messages []*CreateCopilotChatRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The model identifier.
	//
	// example:
	//
	// qwen3.7-max
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The planning configuration.
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
	// The region of the data management center for threat analysis. Specify the management center based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside in regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the conversation turn to replace.
	//
	// example:
	//
	// dc18b894-19a3-45ee-8768-a59478e0****
	ReplaceTurnId *string `json:"ReplaceTurnId,omitempty" xml:"ReplaceTurnId,omitempty"`
	// The route target identifier.
	//
	// example:
	//
	// security-agent-route-****
	RouteTarget *string `json:"RouteTarget,omitempty" xml:"RouteTarget,omitempty"`
	// The request source.
	//
	// example:
	//
	// console
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Specifies whether to enable streaming responses.
	//
	// example:
	//
	// true
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// The deep thinking configuration.
	//
	// example:
	//
	// enabled
	ThinkingMode *string `json:"ThinkingMode,omitempty" xml:"ThinkingMode,omitempty"`
	// The trace ID, which is used to query specific call information.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateCopilotChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotChatRequest) GoString() string {
	return s.String()
}

func (s *CreateCopilotChatRequest) GetClientConversationId() *string {
	return s.ClientConversationId
}

func (s *CreateCopilotChatRequest) GetClientMessageId() *string {
	return s.ClientMessageId
}

func (s *CreateCopilotChatRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateCopilotChatRequest) GetCreateConversation() *CreateCopilotChatRequestCreateConversation {
	return s.CreateConversation
}

func (s *CreateCopilotChatRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateCopilotChatRequest) GetMessages() []*CreateCopilotChatRequestMessages {
	return s.Messages
}

func (s *CreateCopilotChatRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCopilotChatRequest) GetPlanMode() *string {
	return s.PlanMode
}

func (s *CreateCopilotChatRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateCopilotChatRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCopilotChatRequest) GetReplaceTurnId() *string {
	return s.ReplaceTurnId
}

func (s *CreateCopilotChatRequest) GetRouteTarget() *string {
	return s.RouteTarget
}

func (s *CreateCopilotChatRequest) GetSource() *string {
	return s.Source
}

func (s *CreateCopilotChatRequest) GetStream() *bool {
	return s.Stream
}

func (s *CreateCopilotChatRequest) GetThinkingMode() *string {
	return s.ThinkingMode
}

func (s *CreateCopilotChatRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateCopilotChatRequest) SetClientConversationId(v string) *CreateCopilotChatRequest {
	s.ClientConversationId = &v
	return s
}

func (s *CreateCopilotChatRequest) SetClientMessageId(v string) *CreateCopilotChatRequest {
	s.ClientMessageId = &v
	return s
}

func (s *CreateCopilotChatRequest) SetConversationId(v string) *CreateCopilotChatRequest {
	s.ConversationId = &v
	return s
}

func (s *CreateCopilotChatRequest) SetCreateConversation(v *CreateCopilotChatRequestCreateConversation) *CreateCopilotChatRequest {
	s.CreateConversation = v
	return s
}

func (s *CreateCopilotChatRequest) SetLang(v string) *CreateCopilotChatRequest {
	s.Lang = &v
	return s
}

func (s *CreateCopilotChatRequest) SetMessages(v []*CreateCopilotChatRequestMessages) *CreateCopilotChatRequest {
	s.Messages = v
	return s
}

func (s *CreateCopilotChatRequest) SetModel(v string) *CreateCopilotChatRequest {
	s.Model = &v
	return s
}

func (s *CreateCopilotChatRequest) SetPlanMode(v string) *CreateCopilotChatRequest {
	s.PlanMode = &v
	return s
}

func (s *CreateCopilotChatRequest) SetProjectId(v string) *CreateCopilotChatRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCopilotChatRequest) SetRegionId(v string) *CreateCopilotChatRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCopilotChatRequest) SetReplaceTurnId(v string) *CreateCopilotChatRequest {
	s.ReplaceTurnId = &v
	return s
}

func (s *CreateCopilotChatRequest) SetRouteTarget(v string) *CreateCopilotChatRequest {
	s.RouteTarget = &v
	return s
}

func (s *CreateCopilotChatRequest) SetSource(v string) *CreateCopilotChatRequest {
	s.Source = &v
	return s
}

func (s *CreateCopilotChatRequest) SetStream(v bool) *CreateCopilotChatRequest {
	s.Stream = &v
	return s
}

func (s *CreateCopilotChatRequest) SetThinkingMode(v string) *CreateCopilotChatRequest {
	s.ThinkingMode = &v
	return s
}

func (s *CreateCopilotChatRequest) SetTraceId(v string) *CreateCopilotChatRequest {
	s.TraceId = &v
	return s
}

func (s *CreateCopilotChatRequest) Validate() error {
	if s.CreateConversation != nil {
		if err := s.CreateConversation.Validate(); err != nil {
			return err
		}
	}
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCopilotChatRequestCreateConversation struct {
	// The model identifier.
	//
	// example:
	//
	// qwen3.7-max
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The planning configuration.
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
	// The deep thinking configuration.
	//
	// example:
	//
	// enabled
	ThinkingMode *string `json:"ThinkingMode,omitempty" xml:"ThinkingMode,omitempty"`
	// The conversation title.
	//
	// example:
	//
	// Security incident investigation
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateCopilotChatRequestCreateConversation) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotChatRequestCreateConversation) GoString() string {
	return s.String()
}

func (s *CreateCopilotChatRequestCreateConversation) GetModel() *string {
	return s.Model
}

func (s *CreateCopilotChatRequestCreateConversation) GetPlanMode() *string {
	return s.PlanMode
}

func (s *CreateCopilotChatRequestCreateConversation) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateCopilotChatRequestCreateConversation) GetThinkingMode() *string {
	return s.ThinkingMode
}

func (s *CreateCopilotChatRequestCreateConversation) GetTitle() *string {
	return s.Title
}

func (s *CreateCopilotChatRequestCreateConversation) SetModel(v string) *CreateCopilotChatRequestCreateConversation {
	s.Model = &v
	return s
}

func (s *CreateCopilotChatRequestCreateConversation) SetPlanMode(v string) *CreateCopilotChatRequestCreateConversation {
	s.PlanMode = &v
	return s
}

func (s *CreateCopilotChatRequestCreateConversation) SetProjectId(v string) *CreateCopilotChatRequestCreateConversation {
	s.ProjectId = &v
	return s
}

func (s *CreateCopilotChatRequestCreateConversation) SetThinkingMode(v string) *CreateCopilotChatRequestCreateConversation {
	s.ThinkingMode = &v
	return s
}

func (s *CreateCopilotChatRequestCreateConversation) SetTitle(v string) *CreateCopilotChatRequestCreateConversation {
	s.Title = &v
	return s
}

func (s *CreateCopilotChatRequestCreateConversation) Validate() error {
	return dara.Validate(s)
}

type CreateCopilotChatRequestMessages struct {
	// The message content.
	//
	// example:
	//
	// Investigate abnormal login activities in the last hour
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The message role.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateCopilotChatRequestMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotChatRequestMessages) GoString() string {
	return s.String()
}

func (s *CreateCopilotChatRequestMessages) GetContent() *string {
	return s.Content
}

func (s *CreateCopilotChatRequestMessages) GetRole() *string {
	return s.Role
}

func (s *CreateCopilotChatRequestMessages) SetContent(v string) *CreateCopilotChatRequestMessages {
	s.Content = &v
	return s
}

func (s *CreateCopilotChatRequestMessages) SetRole(v string) *CreateCopilotChatRequestMessages {
	s.Role = &v
	return s
}

func (s *CreateCopilotChatRequestMessages) Validate() error {
	return dara.Validate(s)
}
