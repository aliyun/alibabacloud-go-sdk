// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotTurnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetCopilotTurnRequest
	GetConversationId() *string
	SetLang(v string) *GetCopilotTurnRequest
	GetLang() *string
	SetRegionId(v string) *GetCopilotTurnRequest
	GetRegionId() *string
	SetTraceId(v string) *GetCopilotTurnRequest
	GetTraceId() *string
	SetTurnId(v string) *GetCopilotTurnRequest
	GetTurnId() *string
}

type GetCopilotTurnRequest struct {
	// The conversation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
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
	// The region of the data management center for threat analysis. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trace ID used to accurately query call information.
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

func (s GetCopilotTurnRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotTurnRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotTurnRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetCopilotTurnRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCopilotTurnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCopilotTurnRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *GetCopilotTurnRequest) GetTurnId() *string {
	return s.TurnId
}

func (s *GetCopilotTurnRequest) SetConversationId(v string) *GetCopilotTurnRequest {
	s.ConversationId = &v
	return s
}

func (s *GetCopilotTurnRequest) SetLang(v string) *GetCopilotTurnRequest {
	s.Lang = &v
	return s
}

func (s *GetCopilotTurnRequest) SetRegionId(v string) *GetCopilotTurnRequest {
	s.RegionId = &v
	return s
}

func (s *GetCopilotTurnRequest) SetTraceId(v string) *GetCopilotTurnRequest {
	s.TraceId = &v
	return s
}

func (s *GetCopilotTurnRequest) SetTurnId(v string) *GetCopilotTurnRequest {
	s.TurnId = &v
	return s
}

func (s *GetCopilotTurnRequest) Validate() error {
	return dara.Validate(s)
}
