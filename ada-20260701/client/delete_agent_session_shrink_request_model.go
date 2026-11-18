// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *DeleteAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *DeleteAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type DeleteAgentSessionShrinkRequest struct {
	// The correlation identifier for this call, which is echoed in the response. This parameter does not provide idempotence guarantees.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version. Set this parameter to `2.0`.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The business parameter object. For nested fields, refer to the supplementary description of request parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef"}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s DeleteAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *DeleteAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *DeleteAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *DeleteAgentSessionShrinkRequest) SetId(v string) *DeleteAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *DeleteAgentSessionShrinkRequest) SetJsonrpc(v string) *DeleteAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *DeleteAgentSessionShrinkRequest) SetParamsShrink(v string) *DeleteAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *DeleteAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
