// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentSessionsRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentSessionsRequest
	GetJsonrpc() *string
	SetParams(v *ListAgentSessionsRequestParams) *ListAgentSessionsRequest
	GetParams() *ListAgentSessionsRequestParams
}

type ListAgentSessionsRequest struct {
	// The correlation identifier for this call. This identifier is echoed in the response and does not provide idempotency guarantees.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version. Set this parameter to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameter object. The public fields are only AgentName, MaxResults, and NextToken.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"AgentName":"review-agent-demo","MaxResults":20}
	Params *ListAgentSessionsRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s ListAgentSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionsRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionsRequest) GetParams() *ListAgentSessionsRequestParams {
	return s.Params
}

func (s *ListAgentSessionsRequest) SetId(v string) *ListAgentSessionsRequest {
	s.Id = &v
	return s
}

func (s *ListAgentSessionsRequest) SetJsonrpc(v string) *ListAgentSessionsRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionsRequest) SetParams(v *ListAgentSessionsRequestParams) *ListAgentSessionsRequest {
	s.Params = v
	return s
}

func (s *ListAgentSessionsRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsRequestParams struct {
	// Optional. Filters sessions by agent name. The maximum length is 64 characters.
	//
	// example:
	//
	// review-agent-demo
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 50. Keep the page size and filter conditions consistent for subsequent pages.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The NextToken returned from the previous page. Do not pass this parameter on the first page. In the first phase, the default uses a native opaque cursor. Keep the same filter conditions and pass the token back as is.
	//
	// example:
	//
	// next-token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAgentSessionsRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsRequestParams) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsRequestParams) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentSessionsRequestParams) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSessionsRequestParams) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSessionsRequestParams) SetAgentName(v string) *ListAgentSessionsRequestParams {
	s.AgentName = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetMaxResults(v int32) *ListAgentSessionsRequestParams {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSessionsRequestParams) SetNextToken(v string) *ListAgentSessionsRequestParams {
	s.NextToken = &v
	return s
}

func (s *ListAgentSessionsRequestParams) Validate() error {
	return dara.Validate(s)
}
