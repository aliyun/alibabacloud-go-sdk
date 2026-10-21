// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCopilotTurnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *CancelCopilotTurnRequest
	GetConversationId() *string
	SetLang(v string) *CancelCopilotTurnRequest
	GetLang() *string
	SetRegionId(v string) *CancelCopilotTurnRequest
	GetRegionId() *string
	SetTraceId(v string) *CancelCopilotTurnRequest
	GetTraceId() *string
	SetTurnId(v string) *CancelCopilotTurnRequest
	GetTurnId() *string
}

type CancelCopilotTurnRequest struct {
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the data management center for threat analysis is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trace ID used to track service requests.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The conversation turn ID, obtained from the streaming response of CreateCopilotChat.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc18b894-19a3-45ee-8768-a59478e0****
	TurnId *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s CancelCopilotTurnRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCopilotTurnRequest) GoString() string {
	return s.String()
}

func (s *CancelCopilotTurnRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CancelCopilotTurnRequest) GetLang() *string {
	return s.Lang
}

func (s *CancelCopilotTurnRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelCopilotTurnRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *CancelCopilotTurnRequest) GetTurnId() *string {
	return s.TurnId
}

func (s *CancelCopilotTurnRequest) SetConversationId(v string) *CancelCopilotTurnRequest {
	s.ConversationId = &v
	return s
}

func (s *CancelCopilotTurnRequest) SetLang(v string) *CancelCopilotTurnRequest {
	s.Lang = &v
	return s
}

func (s *CancelCopilotTurnRequest) SetRegionId(v string) *CancelCopilotTurnRequest {
	s.RegionId = &v
	return s
}

func (s *CancelCopilotTurnRequest) SetTraceId(v string) *CancelCopilotTurnRequest {
	s.TraceId = &v
	return s
}

func (s *CancelCopilotTurnRequest) SetTurnId(v string) *CancelCopilotTurnRequest {
	s.TurnId = &v
	return s
}

func (s *CancelCopilotTurnRequest) Validate() error {
	return dara.Validate(s)
}
