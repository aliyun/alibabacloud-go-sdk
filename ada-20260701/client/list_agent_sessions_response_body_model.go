// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *ListAgentSessionsResponseBodyJsonRpcResponse) *ListAgentSessionsResponseBody
	GetJsonRpcResponse() *ListAgentSessionsResponseBodyJsonRpcResponse
	SetRequestId(v string) *ListAgentSessionsResponseBody
	GetRequestId() *string
}

type ListAgentSessionsResponseBody struct {
	// The JSON-RPC response object. The business result is located in Result.
	//
	// example:
	//
	// {"Jsonrpc":"2.0","Id":"request-001","Result":{"MaxResults":20,"NextToken":"next-token-example","SessionList":[{"sessionId":"0123456789abcdef0123456789abcdef","title":"Code review","agentName":"review-agent-demo","createdAt":1793232000000,"updatedAt":1793232060000}]}}
	JsonRpcResponse *ListAgentSessionsResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// The server-side request ID used to troubleshoot this call.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAgentSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBody) GetJsonRpcResponse() *ListAgentSessionsResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *ListAgentSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSessionsResponseBody) SetJsonRpcResponse(v *ListAgentSessionsResponseBodyJsonRpcResponse) *ListAgentSessionsResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ListAgentSessionsResponseBody) SetRequestId(v string) *ListAgentSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentSessionsResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponse struct {
	// The Id from the request, echoed as is.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The fixed value 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The paginated result.
	//
	// example:
	//
	// {"MaxResults":20,"NextToken":"next-token-example","SessionList":[{"sessionId":"0123456789abcdef0123456789abcdef"}]}
	Result *ListAgentSessionsResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) GetResult() *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) SetId(v string) *ListAgentSessionsResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) SetJsonrpc(v string) *ListAgentSessionsResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) SetResult(v *ListAgentSessionsResponseBodyJsonRpcResponseResult) *ListAgentSessionsResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSessionsResponseBodyJsonRpcResponseResult struct {
	// The effective page size for the current page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The credential for querying the next page. Pass this value as is to Params.NextToken. Pagination ends when this value is not returned or is empty.
	//
	// example:
	//
	// next-token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The array of session summaries for the current page. If no matching records are found, an empty array [] is returned. The type is any. The fields of array elements retain the original backend names. For the structure, see the supplementary description of response parameters.
	//
	// example:
	//
	// [{"sessionId":"0123456789abcdef0123456789abcdef","title":"Code review","agentName":"review-agent-demo","createdAt":1793232000000,"updatedAt":1793232060000}]
	SessionList interface{} `json:"SessionList,omitempty" xml:"SessionList,omitempty"`
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) GetSessionList() interface{} {
	return s.SessionList
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetMaxResults(v int32) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.MaxResults = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetNextToken(v string) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.NextToken = &v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) SetSessionList(v interface{}) *ListAgentSessionsResponseBodyJsonRpcResponseResult {
	s.SessionList = v
	return s
}

func (s *ListAgentSessionsResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
