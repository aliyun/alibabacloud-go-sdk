// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *LoadAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *LoadAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *LoadAgentSessionRequest
	GetParams() interface{}
}

type LoadAgentSessionRequest struct {
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
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s LoadAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *LoadAgentSessionRequest) GetParams() interface{} {
	return s.Params
}

func (s *LoadAgentSessionRequest) SetId(v string) *LoadAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionRequest) SetJsonrpc(v string) *LoadAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *LoadAgentSessionRequest) SetParams(v interface{}) *LoadAgentSessionRequest {
	s.Params = v
	return s
}

func (s *LoadAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
