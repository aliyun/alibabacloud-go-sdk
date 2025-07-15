// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListHotSourcesRequest
	GetAgentKey() *string
	SetMaxResults(v int32) *ListHotSourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHotSourcesRequest
	GetNextToken() *string
}

type ListHotSourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxxxx_p_efm
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 66
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 下一页的token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListHotSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListHotSourcesRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListHotSourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotSourcesRequest) SetAgentKey(v string) *ListHotSourcesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotSourcesRequest) SetMaxResults(v int32) *ListHotSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotSourcesRequest) SetNextToken(v string) *ListHotSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotSourcesRequest) Validate() error {
	return dara.Validate(s)
}
