// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ReplyAgentSessionShrinkRequest
	GetId() *string
	SetJsonrpc(v string) *ReplyAgentSessionShrinkRequest
	GetJsonrpc() *string
	SetParamsShrink(v string) *ReplyAgentSessionShrinkRequest
	GetParamsShrink() *string
}

type ReplyAgentSessionShrinkRequest struct {
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
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
}

func (s ReplyAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ReplyAgentSessionShrinkRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ReplyAgentSessionShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *ReplyAgentSessionShrinkRequest) SetId(v string) *ReplyAgentSessionShrinkRequest {
	s.Id = &v
	return s
}

func (s *ReplyAgentSessionShrinkRequest) SetJsonrpc(v string) *ReplyAgentSessionShrinkRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ReplyAgentSessionShrinkRequest) SetParamsShrink(v string) *ReplyAgentSessionShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *ReplyAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
