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
	// The JSON-RPC correlation ID for this reply request. The response returns this value as-is. This is different from PermissionRequestId.
	//
	// This parameter is required.
	//
	// example:
	//
	// reply-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version. Fixed value: 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The user interaction reply parameters.
	//
	// This parameter is required.
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
