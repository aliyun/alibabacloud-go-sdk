// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *GetAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *GetAgentSessionRequest
	GetParams() interface{}
}

type GetAgentSessionRequest struct {
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
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s GetAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *GetAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *GetAgentSessionRequest) GetParams() interface{} {
	return s.Params
}

func (s *GetAgentSessionRequest) SetId(v string) *GetAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *GetAgentSessionRequest) SetJsonrpc(v string) *GetAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *GetAgentSessionRequest) SetParams(v interface{}) *GetAgentSessionRequest {
	s.Params = v
	return s
}

func (s *GetAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
