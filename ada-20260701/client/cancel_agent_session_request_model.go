// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CancelAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *CancelAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *CancelAgentSessionRequest
	GetParams() interface{}
}

type CancelAgentSessionRequest struct {
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
	// {"sessionId":"0123456789abcdef0123456789abcdef","_meta":{"promptId":"prompt-example-001"}}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s CancelAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *CancelAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *CancelAgentSessionRequest) GetParams() interface{} {
	return s.Params
}

func (s *CancelAgentSessionRequest) SetId(v string) *CancelAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *CancelAgentSessionRequest) SetJsonrpc(v string) *CancelAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *CancelAgentSessionRequest) SetParams(v interface{}) *CancelAgentSessionRequest {
	s.Params = v
	return s
}

func (s *CancelAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
