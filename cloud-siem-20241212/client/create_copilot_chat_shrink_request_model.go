// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientConversationId(v string) *CreateCopilotChatShrinkRequest
	GetClientConversationId() *string
	SetClientMessageId(v string) *CreateCopilotChatShrinkRequest
	GetClientMessageId() *string
	SetConversationId(v string) *CreateCopilotChatShrinkRequest
	GetConversationId() *string
	SetCreateConversationShrink(v string) *CreateCopilotChatShrinkRequest
	GetCreateConversationShrink() *string
	SetLang(v string) *CreateCopilotChatShrinkRequest
	GetLang() *string
	SetMessagesShrink(v string) *CreateCopilotChatShrinkRequest
	GetMessagesShrink() *string
	SetModel(v string) *CreateCopilotChatShrinkRequest
	GetModel() *string
	SetPlanMode(v string) *CreateCopilotChatShrinkRequest
	GetPlanMode() *string
	SetProjectId(v string) *CreateCopilotChatShrinkRequest
	GetProjectId() *string
	SetRegionId(v string) *CreateCopilotChatShrinkRequest
	GetRegionId() *string
	SetReplaceTurnId(v string) *CreateCopilotChatShrinkRequest
	GetReplaceTurnId() *string
	SetRouteTarget(v string) *CreateCopilotChatShrinkRequest
	GetRouteTarget() *string
	SetSource(v string) *CreateCopilotChatShrinkRequest
	GetSource() *string
	SetStream(v bool) *CreateCopilotChatShrinkRequest
	GetStream() *bool
	SetThinkingMode(v string) *CreateCopilotChatShrinkRequest
	GetThinkingMode() *string
	SetTraceId(v string) *CreateCopilotChatShrinkRequest
	GetTraceId() *string
}

type CreateCopilotChatShrinkRequest struct {
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
	CreateConversationShrink *string `json:"CreateConversation,omitempty" xml:"CreateConversation,omitempty"`
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
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
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

func (s CreateCopilotChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCopilotChatShrinkRequest) GetClientConversationId() *string {
	return s.ClientConversationId
}

func (s *CreateCopilotChatShrinkRequest) GetClientMessageId() *string {
	return s.ClientMessageId
}

func (s *CreateCopilotChatShrinkRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CreateCopilotChatShrinkRequest) GetCreateConversationShrink() *string {
	return s.CreateConversationShrink
}

func (s *CreateCopilotChatShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateCopilotChatShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *CreateCopilotChatShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCopilotChatShrinkRequest) GetPlanMode() *string {
	return s.PlanMode
}

func (s *CreateCopilotChatShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateCopilotChatShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCopilotChatShrinkRequest) GetReplaceTurnId() *string {
	return s.ReplaceTurnId
}

func (s *CreateCopilotChatShrinkRequest) GetRouteTarget() *string {
	return s.RouteTarget
}

func (s *CreateCopilotChatShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *CreateCopilotChatShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *CreateCopilotChatShrinkRequest) GetThinkingMode() *string {
	return s.ThinkingMode
}

func (s *CreateCopilotChatShrinkRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateCopilotChatShrinkRequest) SetClientConversationId(v string) *CreateCopilotChatShrinkRequest {
	s.ClientConversationId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetClientMessageId(v string) *CreateCopilotChatShrinkRequest {
	s.ClientMessageId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetConversationId(v string) *CreateCopilotChatShrinkRequest {
	s.ConversationId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetCreateConversationShrink(v string) *CreateCopilotChatShrinkRequest {
	s.CreateConversationShrink = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetLang(v string) *CreateCopilotChatShrinkRequest {
	s.Lang = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetMessagesShrink(v string) *CreateCopilotChatShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetModel(v string) *CreateCopilotChatShrinkRequest {
	s.Model = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetPlanMode(v string) *CreateCopilotChatShrinkRequest {
	s.PlanMode = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetProjectId(v string) *CreateCopilotChatShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetRegionId(v string) *CreateCopilotChatShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetReplaceTurnId(v string) *CreateCopilotChatShrinkRequest {
	s.ReplaceTurnId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetRouteTarget(v string) *CreateCopilotChatShrinkRequest {
	s.RouteTarget = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetSource(v string) *CreateCopilotChatShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetStream(v bool) *CreateCopilotChatShrinkRequest {
	s.Stream = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetThinkingMode(v string) *CreateCopilotChatShrinkRequest {
	s.ThinkingMode = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) SetTraceId(v string) *CreateCopilotChatShrinkRequest {
	s.TraceId = &v
	return s
}

func (s *CreateCopilotChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
