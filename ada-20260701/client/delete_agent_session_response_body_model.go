// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *DeleteAgentSessionResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *DeleteAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgentSessionResponseBody
	GetSuccess() *bool
}

type DeleteAgentSessionResponseBody struct {
	// The response object. The business result is located in `result`.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"sessionId":"0123456789abcdef0123456789abcdef"}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The server-side request ID, which is used to troubleshoot this call.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Determine the final execution state based on the description of this operation.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentSessionResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *DeleteAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgentSessionResponseBody) SetJsonRpcResponse(v interface{}) *DeleteAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *DeleteAgentSessionResponseBody) SetRequestId(v string) *DeleteAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentSessionResponseBody) SetSuccess(v bool) *DeleteAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
