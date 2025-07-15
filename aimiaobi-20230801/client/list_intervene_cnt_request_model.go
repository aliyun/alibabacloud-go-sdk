// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterveneCntRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListInterveneCntRequest
	GetAgentKey() *string
	SetPageIndex(v int32) *ListInterveneCntRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListInterveneCntRequest
	GetPageSize() *int32
}

type ListInterveneCntRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
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

func (s ListInterveneCntRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInterveneCntRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneCntRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListInterveneCntRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListInterveneCntRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInterveneCntRequest) SetAgentKey(v string) *ListInterveneCntRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneCntRequest) SetPageIndex(v int32) *ListInterveneCntRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneCntRequest) SetPageSize(v int32) *ListInterveneCntRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterveneCntRequest) Validate() error {
	return dara.Validate(s)
}
