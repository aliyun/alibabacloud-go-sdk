// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *GetCopilotConversationRequest
	GetConversationId() *string
	SetLang(v string) *GetCopilotConversationRequest
	GetLang() *string
	SetRegionId(v string) *GetCopilotConversationRequest
	GetRegionId() *string
	SetTraceId(v string) *GetCopilotConversationRequest
	GetTraceId() *string
}

type GetCopilotConversationRequest struct {
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the data management center for threat analysis is located. Select the region of the management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Select this value if your assets are located in the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Select this value if your assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trace ID used to precisely query invocation information.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetCopilotConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotConversationRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotConversationRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetCopilotConversationRequest) GetLang() *string {
	return s.Lang
}

func (s *GetCopilotConversationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCopilotConversationRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *GetCopilotConversationRequest) SetConversationId(v string) *GetCopilotConversationRequest {
	s.ConversationId = &v
	return s
}

func (s *GetCopilotConversationRequest) SetLang(v string) *GetCopilotConversationRequest {
	s.Lang = &v
	return s
}

func (s *GetCopilotConversationRequest) SetRegionId(v string) *GetCopilotConversationRequest {
	s.RegionId = &v
	return s
}

func (s *GetCopilotConversationRequest) SetTraceId(v string) *GetCopilotConversationRequest {
	s.TraceId = &v
	return s
}

func (s *GetCopilotConversationRequest) Validate() error {
	return dara.Validate(s)
}
