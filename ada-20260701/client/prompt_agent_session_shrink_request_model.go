// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *PromptAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *PromptAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *PromptAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type PromptAgentSessionShrinkRequest struct {
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
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s PromptAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *PromptAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *PromptAgentSessionShrinkRequest) SetId(v string) *PromptAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionShrinkRequest) SetJsonrpc(v string) *PromptAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *PromptAgentSessionShrinkRequest) SetParamsShrink(v string) *PromptAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *PromptAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
