// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAgentVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentVersionsRequest
	GetNextToken() *string
}

type ListAgentVersionsRequest struct {
	// 单页最大记录数，取值范围 1 至 100，默认 20。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 查询下一页时传入上一次 ListAgentVersions 响应返回的 nextToken；首次查询不传。
	//
	// example:
	//
	// YWdlbnQtdmVyc2lvbnM6MjA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAgentVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentVersionsRequest) SetMaxResults(v int32) *ListAgentVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentVersionsRequest) SetNextToken(v string) *ListAgentVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentVersionsRequest) Validate() error {
	return dara.Validate(s)
}
