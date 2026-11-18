// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *GetAgentSessionResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *GetAgentSessionResponseBody
	GetRequestId() *string
}

type GetAgentSessionResponseBody struct {
	// The response object. The business result is located in `result`.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"sessionId":"0123456789abcdef0123456789abcdef","title":"Code review","displayName":"Code review","agentName":"review-agent-demo","createdAtMs":1793232000000,"lastMessageAtMs":1793232060000}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The server-side request ID, which is used to troubleshoot this call.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSessionResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *GetAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSessionResponseBody) SetJsonRpcResponse(v interface{}) *GetAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *GetAgentSessionResponseBody) SetRequestId(v string) *GetAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
