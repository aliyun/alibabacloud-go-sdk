// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSessionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ListAgentSessionsShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *ListAgentSessionsShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *ListAgentSessionsShrinkRequest
	GetParamsShrink() *string
}

type ListAgentSessionsShrinkRequest struct {
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
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ListAgentSessionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSessionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAgentSessionsShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ListAgentSessionsShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ListAgentSessionsShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *ListAgentSessionsShrinkRequest) SetId(v string) *ListAgentSessionsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListAgentSessionsShrinkRequest) SetJsonrpc(v string) *ListAgentSessionsShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ListAgentSessionsShrinkRequest) SetParamsShrink(v string) *ListAgentSessionsShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *ListAgentSessionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
