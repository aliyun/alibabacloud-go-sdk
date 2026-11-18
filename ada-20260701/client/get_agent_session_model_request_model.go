// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionModelRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionModelRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *GetAgentSessionModelRequest
	GetParams() interface{}
}

type GetAgentSessionModelRequest struct {
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
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s GetAgentSessionModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionModelRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionModelRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionModelRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionModelRequest) GetParams() interface{} {
	return s.Params
}

func (s *GetAgentSessionModelRequest) SetId(v string) *GetAgentSessionModelRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionModelRequest) SetJsonrpc(v string) *GetAgentSessionModelRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionModelRequest) SetParams(v interface{}) *GetAgentSessionModelRequest {
	s.Params = v
	return s
}

func (s *GetAgentSessionModelRequest) Validate() error {
	return dara.Validate(s)
}
