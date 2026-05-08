// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *ListAgentsRequest
	GetAgentId() *string
	SetAgentScene(v string) *ListAgentsRequest
	GetAgentScene() *string
	SetOwner(v string) *ListAgentsRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListAgentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentsRequest
	GetPageSize() *int32
	SetStatus(v int32) *ListAgentsRequest
	GetStatus() *int32
}

type ListAgentsRequest struct {
	// example:
	//
	// 840016700254633984
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// text
	AgentScene *string `json:"agentScene,omitempty" xml:"agentScene,omitempty"`
	// example:
	//
	// SYSTEM
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *ListAgentsRequest) GetAgentScene() *string {
	return s.AgentScene
}

func (s *ListAgentsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListAgentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListAgentsRequest) SetAgentId(v string) *ListAgentsRequest {
	s.AgentId = &v
	return s
}

func (s *ListAgentsRequest) SetAgentScene(v string) *ListAgentsRequest {
	s.AgentScene = &v
	return s
}

func (s *ListAgentsRequest) SetOwner(v string) *ListAgentsRequest {
	s.Owner = &v
	return s
}

func (s *ListAgentsRequest) SetPageNumber(v int32) *ListAgentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentsRequest) SetPageSize(v int32) *ListAgentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentsRequest) SetStatus(v int32) *ListAgentsRequest {
	s.Status = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
