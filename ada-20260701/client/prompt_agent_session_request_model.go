// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *PromptAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *PromptAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *PromptAgentSessionRequest
	GetParams() interface{}
}

type PromptAgentSessionRequest struct {
	// The client call identifier used to associate the current response. This parameter does not provide idempotency guarantees.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Set this parameter to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The structure of the business parameters. For core nested fields, refer to the supplementary instructions.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef","prompt":[{"type":"text","text":"Explain the issue of accessing the first item when a Python list is empty, and provide modification suggestions."}]}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s PromptAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *PromptAgentSessionRequest) GetParams() interface{} {
	return s.Params
}

func (s *PromptAgentSessionRequest) SetId(v string) *PromptAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionRequest) SetJsonrpc(v string) *PromptAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *PromptAgentSessionRequest) SetParams(v interface{}) *PromptAgentSessionRequest {
	s.Params = v
	return s
}

func (s *PromptAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
