// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *ReplyAgentSessionResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *ReplyAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReplyAgentSessionResponseBody
	GetSuccess() *bool
}

type ReplyAgentSessionResponseBody struct {
	// The JSON-RPC response. For core embedded fields, see the supplementary description.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"accepted":true}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The ID of the current request, used for troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the current API call is successful. This parameter cannot replace the business acceptance result.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReplyAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *ReplyAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplyAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReplyAgentSessionResponseBody) SetJsonRpcResponse(v interface{}) *ReplyAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *ReplyAgentSessionResponseBody) SetRequestId(v string) *ReplyAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplyAgentSessionResponseBody) SetSuccess(v bool) *ReplyAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *ReplyAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
