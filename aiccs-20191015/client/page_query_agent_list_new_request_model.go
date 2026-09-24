// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListNewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *PageQueryAgentListNewRequest
	GetAgentId() *string
	SetAgentName(v string) *PageQueryAgentListNewRequest
	GetAgentName() *string
	SetInboundConfigurableOnly(v bool) *PageQueryAgentListNewRequest
	GetInboundConfigurableOnly() *bool
	SetIsAvailable(v bool) *PageQueryAgentListNewRequest
	GetIsAvailable() *bool
	SetPageIndex(v int64) *PageQueryAgentListNewRequest
	GetPageIndex() *int64
	SetPageNo(v int64) *PageQueryAgentListNewRequest
	GetPageNo() *int64
	SetPageSize(v int64) *PageQueryAgentListNewRequest
	GetPageSize() *int64
	SetServiceDirection(v string) *PageQueryAgentListNewRequest
	GetServiceDirection() *string
	SetTemplateId(v int64) *PageQueryAgentListNewRequest
	GetTemplateId() *int64
}

type PageQueryAgentListNewRequest struct {
	// Agent ID
	//
	// example:
	//
	// 12345
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The agent name.
	//
	// example:
	//
	// Intelligent Customer Service Assistant
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// Specifies whether to return only candidate agents that are configurable for inbound calls.
	//
	// example:
	//
	// true
	InboundConfigurableOnly *bool `json:"InboundConfigurableOnly,omitempty" xml:"InboundConfigurableOnly,omitempty"`
	// Specifies whether the agent is available for outbound calls. A value of True indicates that the current deployment branch of the agent has a published version and is available for outbound calls.
	//
	// example:
	//
	// true
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// The page number. This parameter is deprecated. Use PageNo instead.
	//
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The service direction.
	//
	// example:
	//
	// Sample value
	ServiceDirection *string `json:"ServiceDirection,omitempty" xml:"ServiceDirection,omitempty"`
	// The source template ID.
	//
	// example:
	//
	// 23
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s PageQueryAgentListNewRequest) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListNewRequest) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListNewRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *PageQueryAgentListNewRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *PageQueryAgentListNewRequest) GetInboundConfigurableOnly() *bool {
	return s.InboundConfigurableOnly
}

func (s *PageQueryAgentListNewRequest) GetIsAvailable() *bool {
	return s.IsAvailable
}

func (s *PageQueryAgentListNewRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *PageQueryAgentListNewRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageQueryAgentListNewRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageQueryAgentListNewRequest) GetServiceDirection() *string {
	return s.ServiceDirection
}

func (s *PageQueryAgentListNewRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *PageQueryAgentListNewRequest) SetAgentId(v string) *PageQueryAgentListNewRequest {
	s.AgentId = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetAgentName(v string) *PageQueryAgentListNewRequest {
	s.AgentName = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetInboundConfigurableOnly(v bool) *PageQueryAgentListNewRequest {
	s.InboundConfigurableOnly = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetIsAvailable(v bool) *PageQueryAgentListNewRequest {
	s.IsAvailable = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetPageIndex(v int64) *PageQueryAgentListNewRequest {
	s.PageIndex = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetPageNo(v int64) *PageQueryAgentListNewRequest {
	s.PageNo = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetPageSize(v int64) *PageQueryAgentListNewRequest {
	s.PageSize = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetServiceDirection(v string) *PageQueryAgentListNewRequest {
	s.ServiceDirection = &v
	return s
}

func (s *PageQueryAgentListNewRequest) SetTemplateId(v int64) *PageQueryAgentListNewRequest {
	s.TemplateId = &v
	return s
}

func (s *PageQueryAgentListNewRequest) Validate() error {
	return dara.Validate(s)
}
