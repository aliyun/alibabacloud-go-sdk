// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCopilotTurnStreamShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *ResumeCopilotTurnStreamShrinkRequest
	GetConversationId() *string
	SetCursorByRunShrink(v string) *ResumeCopilotTurnStreamShrinkRequest
	GetCursorByRunShrink() *string
	SetLang(v string) *ResumeCopilotTurnStreamShrinkRequest
	GetLang() *string
	SetRegionId(v string) *ResumeCopilotTurnStreamShrinkRequest
	GetRegionId() *string
	SetTraceId(v string) *ResumeCopilotTurnStreamShrinkRequest
	GetTraceId() *string
	SetTurnId(v string) *ResumeCopilotTurnStreamShrinkRequest
	GetTurnId() *string
}

type ResumeCopilotTurnStreamShrinkRequest struct {
	// The conversation ID. This parameter is required for the second and subsequent turns in a multi-turn conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The native run cursor mapping.
	//
	// This parameter is required.
	CursorByRunShrink *string `json:"CursorByRun,omitempty" xml:"CursorByRun,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the data management center for threat analysis is located. You need to select the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The asset belongs to the Chinese mainland and China (Hong Kong).
	//
	// - ap-southeast-1: The asset belongs to a region outside the Chinese mainland.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The primary key ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc18b894-19a3-45ee-8768-a59478e0****
	TurnId *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s ResumeCopilotTurnStreamShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeCopilotTurnStreamShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResumeCopilotTurnStreamShrinkRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ResumeCopilotTurnStreamShrinkRequest) GetCursorByRunShrink() *string {
	return s.CursorByRunShrink
}

func (s *ResumeCopilotTurnStreamShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ResumeCopilotTurnStreamShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResumeCopilotTurnStreamShrinkRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *ResumeCopilotTurnStreamShrinkRequest) GetTurnId() *string {
	return s.TurnId
}

func (s *ResumeCopilotTurnStreamShrinkRequest) SetConversationId(v string) *ResumeCopilotTurnStreamShrinkRequest {
	s.ConversationId = &v
	return s
}

func (s *ResumeCopilotTurnStreamShrinkRequest) SetCursorByRunShrink(v string) *ResumeCopilotTurnStreamShrinkRequest {
	s.CursorByRunShrink = &v
	return s
}

func (s *ResumeCopilotTurnStreamShrinkRequest) SetLang(v string) *ResumeCopilotTurnStreamShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ResumeCopilotTurnStreamShrinkRequest) SetRegionId(v string) *ResumeCopilotTurnStreamShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeCopilotTurnStreamShrinkRequest) SetTraceId(v string) *ResumeCopilotTurnStreamShrinkRequest {
	s.TraceId = &v
	return s
}

func (s *ResumeCopilotTurnStreamShrinkRequest) SetTurnId(v string) *ResumeCopilotTurnStreamShrinkRequest {
	s.TurnId = &v
	return s
}

func (s *ResumeCopilotTurnStreamShrinkRequest) Validate() error {
	return dara.Validate(s)
}
