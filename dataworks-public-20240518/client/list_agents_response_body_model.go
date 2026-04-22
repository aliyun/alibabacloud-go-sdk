// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *ListAgentsResponseBodyJsonRpcResponse) *ListAgentsResponseBody
	GetJsonRpcResponse() *ListAgentsResponseBodyJsonRpcResponse
	SetRequestId(v string) *ListAgentsResponseBody
	GetRequestId() *string
}

type ListAgentsResponseBody struct {
	JsonRpcResponse *ListAgentsResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 4003B739-C33C-5297-B810-0490EEE5A767
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) GetJsonRpcResponse() *ListAgentsResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *ListAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentsResponseBody) SetJsonRpcResponse(v *ListAgentsResponseBodyJsonRpcResponse) *ListAgentsResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentsResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 70623e38-a889-4192-930a-752ffdd75f48
	Id      *string                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Jsonrpc *string                                      `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *ListAgentsResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAgentsResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *ListAgentsResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentsResponseBodyJsonRpcResponse) GetResult() *ListAgentsResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *ListAgentsResponseBodyJsonRpcResponse) SetId(v string) *ListAgentsResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponse) SetJsonrpc(v string) *ListAgentsResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponse) SetResult(v *ListAgentsResponseBodyJsonRpcResponseResult) *ListAgentsResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentsResponseBodyJsonRpcResponseResult struct {
	Agents []*ListAgentsResponseBodyJsonRpcResponseResultAgents `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 27
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgentsResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) GetAgents() []*ListAgentsResponseBodyJsonRpcResponseResultAgents {
	return s.Agents
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) SetAgents(v []*ListAgentsResponseBodyJsonRpcResponseResultAgents) *ListAgentsResponseBodyJsonRpcResponseResult {
	s.Agents = v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) SetMaxResults(v int32) *ListAgentsResponseBodyJsonRpcResponseResult {
	s.MaxResults = &v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) SetNextToken(v string) *ListAgentsResponseBodyJsonRpcResponseResult {
	s.NextToken = &v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) SetTotalCount(v int32) *ListAgentsResponseBodyJsonRpcResponseResult {
	s.TotalCount = &v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponseResult) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentsResponseBodyJsonRpcResponseResultAgents struct {
	// example:
	//
	// chat_cli_chatbi
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
}

func (s ListAgentsResponseBodyJsonRpcResponseResultAgents) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyJsonRpcResponseResultAgents) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyJsonRpcResponseResultAgents) GetAgentName() *string {
	return s.AgentName
}

func (s *ListAgentsResponseBodyJsonRpcResponseResultAgents) SetAgentName(v string) *ListAgentsResponseBodyJsonRpcResponseResultAgents {
	s.AgentName = &v
	return s
}

func (s *ListAgentsResponseBodyJsonRpcResponseResultAgents) Validate() error {
	return dara.Validate(s)
}
