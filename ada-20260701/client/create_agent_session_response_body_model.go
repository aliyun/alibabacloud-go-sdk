// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *CreateAgentSessionResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *CreateAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgentSessionResponseBody
	GetSuccess() *bool
}

type CreateAgentSessionResponseBody struct {
	// The response object. The business result is located in `result`.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"sessionId":"0123456789abcdef0123456789abcdef","attached":false}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The server-side request ID. It is used to troubleshoot this call.
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

func (s CreateAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgentSessionResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *CreateAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgentSessionResponseBody) SetJsonRpcResponse(v interface{}) *CreateAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *CreateAgentSessionResponseBody) SetRequestId(v string) *CreateAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgentSessionResponseBody) SetSuccess(v bool) *CreateAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
