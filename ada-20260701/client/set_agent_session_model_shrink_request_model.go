// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentSessionModelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetAgentSessionModelShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *SetAgentSessionModelShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *SetAgentSessionModelShrinkRequest
	GetParamsShrink() *string
}

type SetAgentSessionModelShrinkRequest struct {
	// The client request identifier, which is echoed in the response. It is used to associate this call and does not have idempotent semantics.
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
	// The defined business parameter structure. For more information about the core nested fields, see the supplementary description.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef","modelId":"text-model-example"}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s SetAgentSessionModelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAgentSessionModelShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetAgentSessionModelShrinkRequest) GetId() *string {
	return s.Id
}

func (s *SetAgentSessionModelShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *SetAgentSessionModelShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *SetAgentSessionModelShrinkRequest) SetId(v string) *SetAgentSessionModelShrinkRequest {
	s.Id = &v
	return s
}

func (s *SetAgentSessionModelShrinkRequest) SetJsonrpc(v string) *SetAgentSessionModelShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *SetAgentSessionModelShrinkRequest) SetParamsShrink(v string) *SetAgentSessionModelShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *SetAgentSessionModelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
