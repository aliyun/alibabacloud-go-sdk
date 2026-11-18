// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionModelShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionModelShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *GetAgentSessionModelShrinkRequest
	GetParamsShrink() *string
}

type GetAgentSessionModelShrinkRequest struct {
	// The client request identifier, which is echoed in the response. It is used to associate this call and does not have idempotent semantics.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Set the value to `2.0`.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The defined business parameter structure. For core nested fields, see the supplementary description.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef"}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s GetAgentSessionModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionModelShrinkRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionModelShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionModelShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *GetAgentSessionModelShrinkRequest) SetId(v string) *GetAgentSessionModelShrinkRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionModelShrinkRequest) SetJsonrpc(v string) *GetAgentSessionModelShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionModelShrinkRequest) SetParamsShrink(v string) *GetAgentSessionModelShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *GetAgentSessionModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
