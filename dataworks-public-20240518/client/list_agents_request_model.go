// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentsRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentsRequest
	GetJsonrpc() *string
	SetParams(v *ListAgentsRequestParams) *ListAgentsRequest
	GetParams() *ListAgentsRequestParams
}

type ListAgentsRequest struct {
	// The request ID passed in by the caller. The value is returned as-is in the response.
	//
	// example:
	//
	// 4as3dasf654a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Fixed value: 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The parameters for this request.
	Params *ListAgentsRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentsRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentsRequest) GetParams() *ListAgentsRequestParams {
	return s.Params
}

func (s *ListAgentsRequest) SetId(v string) *ListAgentsRequest {
	s.Id = &v
	return s
}

func (s *ListAgentsRequest) SetJsonrpc(v string) *ListAgentsRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentsRequest) SetParams(v *ListAgentsRequestParams) *ListAgentsRequest {
	s.Params = v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentsRequestParams struct {
	// The exact agent name to filter by. If not specified, all agents are returned.
	//
	// example:
	//
	// chat_cli_chatbi
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The number of entries per page. Default value: 50.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page, following the Alibaba Cloud OpenAPI convention. Do not specify this parameter for the first page. For subsequent pages, pass in the nextToken value returned in the previous response.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAgentsRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequestParams) GoString() string {
	return s.String()
}

func (s *ListAgentsRequestParams) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentsRequestParams) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentsRequestParams) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentsRequestParams) SetAgentName(v string) *ListAgentsRequestParams {
	s.AgentName = &v
	return s
}

func (s *ListAgentsRequestParams) SetMaxResults(v int32) *ListAgentsRequestParams {
	s.MaxResults = &v
	return s
}

func (s *ListAgentsRequestParams) SetNextToken(v string) *ListAgentsRequestParams {
	s.NextToken = &v
	return s
}

func (s *ListAgentsRequestParams) Validate() error {
	return dara.Validate(s)
}
