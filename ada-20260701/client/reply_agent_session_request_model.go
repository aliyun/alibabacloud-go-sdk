// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ReplyAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *ReplyAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v interface{}) *ReplyAgentSessionRequest
	GetParams() interface{}
}

type ReplyAgentSessionRequest struct {
	// The client call identifier used to associate the current response. This parameter does not provide idempotence guarantees.
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
	// {"sessionId":"0123456789abcdef0123456789abcdef","permissionRequestId":"permission-example","outcome":{"outcome":"selected","optionId":"allow_once"}}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ReplyAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *ReplyAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ReplyAgentSessionRequest) GetParams() interface{} {
	return s.Params
}

func (s *ReplyAgentSessionRequest) SetId(v string) *ReplyAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *ReplyAgentSessionRequest) SetJsonrpc(v string) *ReplyAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ReplyAgentSessionRequest) SetParams(v interface{}) *ReplyAgentSessionRequest {
	s.Params = v
	return s
}

func (s *ReplyAgentSessionRequest) Validate() error {
	return dara.Validate(s)
}
