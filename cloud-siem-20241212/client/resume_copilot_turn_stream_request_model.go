// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCopilotTurnStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *ResumeCopilotTurnStreamRequest
	GetConversationId() *string
	SetCursorByRun(v map[string]*int64) *ResumeCopilotTurnStreamRequest
	GetCursorByRun() map[string]*int64
	SetLang(v string) *ResumeCopilotTurnStreamRequest
	GetLang() *string
	SetRegionId(v string) *ResumeCopilotTurnStreamRequest
	GetRegionId() *string
	SetTraceId(v string) *ResumeCopilotTurnStreamRequest
	GetTraceId() *string
	SetTurnId(v string) *ResumeCopilotTurnStreamRequest
	GetTurnId() *string
}

type ResumeCopilotTurnStreamRequest struct {
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
	CursorByRun map[string]*int64 `json:"CursorByRun,omitempty" xml:"CursorByRun,omitempty"`
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

func (s ResumeCopilotTurnStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeCopilotTurnStreamRequest) GoString() string {
	return s.String()
}

func (s *ResumeCopilotTurnStreamRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ResumeCopilotTurnStreamRequest) GetCursorByRun() map[string]*int64 {
	return s.CursorByRun
}

func (s *ResumeCopilotTurnStreamRequest) GetLang() *string {
	return s.Lang
}

func (s *ResumeCopilotTurnStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResumeCopilotTurnStreamRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *ResumeCopilotTurnStreamRequest) GetTurnId() *string {
	return s.TurnId
}

func (s *ResumeCopilotTurnStreamRequest) SetConversationId(v string) *ResumeCopilotTurnStreamRequest {
	s.ConversationId = &v
	return s
}

func (s *ResumeCopilotTurnStreamRequest) SetCursorByRun(v map[string]*int64) *ResumeCopilotTurnStreamRequest {
	s.CursorByRun = v
	return s
}

func (s *ResumeCopilotTurnStreamRequest) SetLang(v string) *ResumeCopilotTurnStreamRequest {
	s.Lang = &v
	return s
}

func (s *ResumeCopilotTurnStreamRequest) SetRegionId(v string) *ResumeCopilotTurnStreamRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeCopilotTurnStreamRequest) SetTraceId(v string) *ResumeCopilotTurnStreamRequest {
	s.TraceId = &v
	return s
}

func (s *ResumeCopilotTurnStreamRequest) SetTurnId(v string) *ResumeCopilotTurnStreamRequest {
	s.TurnId = &v
	return s
}

func (s *ResumeCopilotTurnStreamRequest) Validate() error {
	return dara.Validate(s)
}
