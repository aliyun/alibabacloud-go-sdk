// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIMBotsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentRuntimeId(v string) *ListIMBotsRequest
	GetAgentRuntimeId() *string
	SetBotBizType(v string) *ListIMBotsRequest
	GetBotBizType() *string
	SetBotName(v string) *ListIMBotsRequest
	GetBotName() *string
	SetPageNumber(v int32) *ListIMBotsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIMBotsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListIMBotsRequest
	GetStatus() *string
}

type ListIMBotsRequest struct {
	// The ID of the agent runtime.
	AgentRuntimeId *string `json:"agentRuntimeId,omitempty" xml:"agentRuntimeId,omitempty"`
	// The business type of the bot.
	BotBizType *string `json:"botBizType,omitempty" xml:"botBizType,omitempty"`
	// The name of the IM bot. The system performs a case-insensitive substring search.
	BotName *string `json:"botName,omitempty" xml:"botName,omitempty"`
	// The page number. Must be greater than or equal to 1. Default: 1.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Valid values: 1–100. Default: 20.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The status of the bot.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIMBotsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIMBotsRequest) GoString() string {
	return s.String()
}

func (s *ListIMBotsRequest) GetAgentRuntimeId() *string {
	return s.AgentRuntimeId
}

func (s *ListIMBotsRequest) GetBotBizType() *string {
	return s.BotBizType
}

func (s *ListIMBotsRequest) GetBotName() *string {
	return s.BotName
}

func (s *ListIMBotsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIMBotsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIMBotsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListIMBotsRequest) SetAgentRuntimeId(v string) *ListIMBotsRequest {
	s.AgentRuntimeId = &v
	return s
}

func (s *ListIMBotsRequest) SetBotBizType(v string) *ListIMBotsRequest {
	s.BotBizType = &v
	return s
}

func (s *ListIMBotsRequest) SetBotName(v string) *ListIMBotsRequest {
	s.BotName = &v
	return s
}

func (s *ListIMBotsRequest) SetPageNumber(v int32) *ListIMBotsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIMBotsRequest) SetPageSize(v int32) *ListIMBotsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIMBotsRequest) SetStatus(v string) *ListIMBotsRequest {
	s.Status = &v
	return s
}

func (s *ListIMBotsRequest) Validate() error {
	return dara.Validate(s)
}
