// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *CreateAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *CreateAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type CreateAgentSessionShrinkRequest struct {
	// The correlation identifier for this call. It is echoed in the response and does not provide idempotency guarantees.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version. Set the value to `2.0`.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameter object. For nested fields, see the supplementary description of request parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"clientToken":"create-review-001","_meta":{"agent":{"agentName":"review-agent-demo"},"initialConfigOptions":{"approvalMode":"default"}}}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s CreateAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *CreateAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CreateAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *CreateAgentSessionShrinkRequest) SetId(v string) *CreateAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *CreateAgentSessionShrinkRequest) SetJsonrpc(v string) *CreateAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *CreateAgentSessionShrinkRequest) SetParamsShrink(v string) *CreateAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *CreateAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
