// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotTurnByClientMessageIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientMessageId(v string) *GetCopilotTurnByClientMessageIdRequest
	GetClientMessageId() *string
	SetConversationId(v string) *GetCopilotTurnByClientMessageIdRequest
	GetConversationId() *string
	SetLang(v string) *GetCopilotTurnByClientMessageIdRequest
	GetLang() *string
	SetRegionId(v string) *GetCopilotTurnByClientMessageIdRequest
	GetRegionId() *string
	SetTraceId(v string) *GetCopilotTurnByClientMessageIdRequest
	GetTraceId() *string
}

type GetCopilotTurnByClientMessageIdRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The language of the returned message. Valid values:
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
	// - cn-hangzhou: Select this value if your assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Select this value if your assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trace ID used to track the service request.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetCopilotTurnByClientMessageIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotTurnByClientMessageIdRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotTurnByClientMessageIdRequest) GetClientMessageId() *string {
	return s.ClientMessageId
}

func (s *GetCopilotTurnByClientMessageIdRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetCopilotTurnByClientMessageIdRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCopilotTurnByClientMessageIdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCopilotTurnByClientMessageIdRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *GetCopilotTurnByClientMessageIdRequest) SetClientMessageId(v string) *GetCopilotTurnByClientMessageIdRequest {
	s.ClientMessageId = &v
	return s
}

func (s *GetCopilotTurnByClientMessageIdRequest) SetConversationId(v string) *GetCopilotTurnByClientMessageIdRequest {
	s.ConversationId = &v
	return s
}

func (s *GetCopilotTurnByClientMessageIdRequest) SetLang(v string) *GetCopilotTurnByClientMessageIdRequest {
	s.Lang = &v
	return s
}

func (s *GetCopilotTurnByClientMessageIdRequest) SetRegionId(v string) *GetCopilotTurnByClientMessageIdRequest {
	s.RegionId = &v
	return s
}

func (s *GetCopilotTurnByClientMessageIdRequest) SetTraceId(v string) *GetCopilotTurnByClientMessageIdRequest {
	s.TraceId = &v
	return s
}

func (s *GetCopilotTurnByClientMessageIdRequest) Validate() error {
	return dara.Validate(s)
}
