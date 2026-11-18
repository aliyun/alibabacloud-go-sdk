// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *LoadAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *LoadAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *LoadAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type LoadAgentSessionShrinkRequest struct {
	// The client call identifier used to associate the current response. It does not provide idempotency guarantees.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Set this parameter to `2.0`.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameter structure. For core nested fields, see the supplementary description.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef"}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s LoadAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *LoadAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *LoadAgentSessionShrinkRequest) SetId(v string) *LoadAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionShrinkRequest) SetJsonrpc(v string) *LoadAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *LoadAgentSessionShrinkRequest) SetParamsShrink(v string) *LoadAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *LoadAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
