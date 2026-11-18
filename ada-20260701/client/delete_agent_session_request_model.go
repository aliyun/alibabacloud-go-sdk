// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *DeleteAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *DeleteAgentSessionRequest
	GetParams() interface{}
}

type DeleteAgentSessionRequest struct {
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
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s DeleteAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *DeleteAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *DeleteAgentSessionRequest) GetParams() interface{} {
	return s.Params
}

func (s *DeleteAgentSessionRequest) SetId(v string) *DeleteAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *DeleteAgentSessionRequest) SetJsonrpc(v string) *DeleteAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *DeleteAgentSessionRequest) SetParams(v interface{}) *DeleteAgentSessionRequest {
	s.Params = v
	return s
}

func (s *DeleteAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
