// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *GetAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type GetAgentSessionShrinkRequest struct {
	// The correlation identifier for this call, which is echoed in the response. It does not provide idempotency guarantees.
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

func (s GetAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *GetAgentSessionShrinkRequest) SetId(v string) *GetAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionShrinkRequest) SetJsonrpc(v string) *GetAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionShrinkRequest) SetParamsShrink(v string) *GetAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *GetAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
