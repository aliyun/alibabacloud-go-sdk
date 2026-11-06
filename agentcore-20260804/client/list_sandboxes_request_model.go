// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSandboxesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSandboxesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSandboxesRequest
	GetNextToken() *string
	SetSearchText(v string) *ListSandboxesRequest
	GetSearchText() *string
	SetServiceEndpointId(v string) *ListSandboxesRequest
	GetServiceEndpointId() *string
	SetSessionId(v string) *ListSandboxesRequest
	GetSessionId() *string
}

type ListSandboxesRequest struct {
	// The maximum number of entries to return per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The cursor for querying the next page. When paginating, keep `workspaceId`, `agentId`, `searchText`, `sessionId`, and `maxResults` unchanged.
	//
	// example:
	//
	// next-token-1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The fragment of the Sandbox ID for a case-insensitive fuzzy search.
	//
	// example:
	//
	// 2f360
	SearchText *string `json:"searchText,omitempty" xml:"searchText,omitempty"`
	// 省略或传空白时查询智能体 latest 运行单元；传入时该服务端点必须属于本智能体且目标类型为 AGENT_VERSION，否则返回 Resource.NotFound。独立运行的服务端点只返回其自身运行单元的 Sandbox；DEFAULT 服务端点尚未接管独立运行时回落到 latest 运行单元；非 DEFAULT 且未独立运行时返回 ServiceEndpointTargetNotReady。运行单元尚未创建时返回空列表且 totalCount 为 0，不报错。
	//
	// example:
	//
	// se-7f3a92c1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
	// The fragment of the current valid session ID for a case-insensitive fuzzy search.
	//
	// example:
	//
	// session-a
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ListSandboxesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSandboxesRequest) GoString() string {
	return s.String()
}

func (s *ListSandboxesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSandboxesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSandboxesRequest) GetSearchText() *string {
	return s.SearchText
}

func (s *ListSandboxesRequest) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *ListSandboxesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSandboxesRequest) SetMaxResults(v int32) *ListSandboxesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSandboxesRequest) SetNextToken(v string) *ListSandboxesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSandboxesRequest) SetSearchText(v string) *ListSandboxesRequest {
	s.SearchText = &v
	return s
}

func (s *ListSandboxesRequest) SetServiceEndpointId(v string) *ListSandboxesRequest {
	s.ServiceEndpointId = &v
	return s
}

func (s *ListSandboxesRequest) SetSessionId(v string) *ListSandboxesRequest {
	s.SessionId = &v
	return s
}

func (s *ListSandboxesRequest) Validate() error {
	return dara.Validate(s)
}
