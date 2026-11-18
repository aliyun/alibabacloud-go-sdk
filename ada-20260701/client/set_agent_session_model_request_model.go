// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentSessionModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetAgentSessionModelRequest
	GetId() *string
	SetJsonrpc(v string) *SetAgentSessionModelRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *SetAgentSessionModelRequest
	GetParams() interface{}
}

type SetAgentSessionModelRequest struct {
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
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s SetAgentSessionModelRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAgentSessionModelRequest) GoString() string {
	return s.String()
}

func (s *SetAgentSessionModelRequest) GetId() *string {
	return s.Id
}

func (s *SetAgentSessionModelRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *SetAgentSessionModelRequest) GetParams() interface{} {
	return s.Params
}

func (s *SetAgentSessionModelRequest) SetId(v string) *SetAgentSessionModelRequest {
	s.Id = &v
	return s
}

func (s *SetAgentSessionModelRequest) SetJsonrpc(v string) *SetAgentSessionModelRequest {
	s.Jsonrpc = &v
	return s
}

func (s *SetAgentSessionModelRequest) SetParams(v interface{}) *SetAgentSessionModelRequest {
	s.Params = v
	return s
}

func (s *SetAgentSessionModelRequest) Validate() error {
	return dara.Validate(s)
}
