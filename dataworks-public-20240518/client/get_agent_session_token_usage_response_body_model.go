// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionTokenUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) *GetAgentSessionTokenUsageResponseBody
	GetJsonRpcResponse() *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse
	SetRequestId(v string) *GetAgentSessionTokenUsageResponseBody
	GetRequestId() *string
}

type GetAgentSessionTokenUsageResponseBody struct {
	JsonRpcResponse *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// B313938A-4475-599B-98EB-A0875019FD5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentSessionTokenUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageResponseBody) GetJsonRpcResponse() *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse {
	return s.JsonRpcResponse
}

func (s *GetAgentSessionTokenUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSessionTokenUsageResponseBody) SetJsonRpcResponse(v *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) *GetAgentSessionTokenUsageResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBody) SetRequestId(v string) *GetAgentSessionTokenUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBody) Validate() error {
	if s.JsonRpcResponse != nil {
		if err := s.JsonRpcResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSessionTokenUsageResponseBodyJsonRpcResponse struct {
	// example:
	//
	// 8212598228302533855
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2.0
	Jsonrpc *string                                                     `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	Result  *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) GetResult() *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult {
	return s.Result
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) SetId(v string) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse {
	s.Id = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) SetJsonrpc(v string) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) SetResult(v *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse {
	s.Result = v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponse) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult struct {
	// example:
	//
	// 4000
	CachedTokens *int64 `json:"CachedTokens,omitempty" xml:"CachedTokens,omitempty"`
	// example:
	//
	// 2000
	CompletionTokens *int64 `json:"CompletionTokens,omitempty" xml:"CompletionTokens,omitempty"`
	// example:
	//
	// 1000
	PromptTokens *int64 `json:"PromptTokens,omitempty" xml:"PromptTokens,omitempty"`
	// example:
	//
	// 3000
	ThoughtsTokens *int64 `json:"ThoughtsTokens,omitempty" xml:"ThoughtsTokens,omitempty"`
	// example:
	//
	// 2000
	TotalTokens *int64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) GoString() string {
	return s.String()
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) GetCachedTokens() *int64 {
	return s.CachedTokens
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) GetCompletionTokens() *int64 {
	return s.CompletionTokens
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) GetPromptTokens() *int64 {
	return s.PromptTokens
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) GetThoughtsTokens() *int64 {
	return s.ThoughtsTokens
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) SetCachedTokens(v int64) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult {
	s.CachedTokens = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) SetCompletionTokens(v int64) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult {
	s.CompletionTokens = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) SetPromptTokens(v int64) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult {
	s.PromptTokens = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) SetThoughtsTokens(v int64) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult {
	s.ThoughtsTokens = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) SetTotalTokens(v int64) *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult {
	s.TotalTokens = &v
	return s
}

func (s *GetAgentSessionTokenUsageResponseBodyJsonRpcResponseResult) Validate() error {
	return dara.Validate(s)
}
