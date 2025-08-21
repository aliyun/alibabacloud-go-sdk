// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *ListAgentRequest
	GetAgentName() *string
	SetGoodsCodes(v string) *ListAgentRequest
	GetGoodsCodes() *string
	SetPageNumber(v int32) *ListAgentRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAgentRequest
	GetPageSize() *int32
}

type ListAgentRequest struct {
	AgentName  *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	GoodsCodes *string `json:"GoodsCodes,omitempty" xml:"GoodsCodes,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentRequest) GoString() string {
	return s.String()
}

func (s *ListAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentRequest) GetGoodsCodes() *string {
	return s.GoodsCodes
}

func (s *ListAgentRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAgentRequest) SetAgentName(v string) *ListAgentRequest {
	s.AgentName = &v
	return s
}

func (s *ListAgentRequest) SetGoodsCodes(v string) *ListAgentRequest {
	s.GoodsCodes = &v
	return s
}

func (s *ListAgentRequest) SetPageNumber(v int32) *ListAgentRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentRequest) SetPageSize(v int32) *ListAgentRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentRequest) Validate() error {
	return dara.Validate(s)
}
