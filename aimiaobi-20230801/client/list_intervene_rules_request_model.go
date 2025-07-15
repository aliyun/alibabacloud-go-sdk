// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListInterveneRulesRequest
	GetAgentKey() *string
	SetPageIndex(v int32) *ListInterveneRulesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListInterveneRulesRequest
	GetPageSize() *int32
}

type ListInterveneRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c160c841c8e54295bf2f441432785944_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneRulesRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListInterveneRulesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListInterveneRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterveneRulesRequest) SetAgentKey(v string) *ListInterveneRulesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneRulesRequest) SetPageIndex(v int32) *ListInterveneRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneRulesRequest) SetPageSize(v int32) *ListInterveneRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterveneRulesRequest) Validate() error {
	return dara.Validate(s)
}
