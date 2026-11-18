// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *CancelAgentSessionResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *CancelAgentSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelAgentSessionResponseBody
	GetSuccess() *bool
}

type CancelAgentSessionResponseBody struct {
	// The response object. The business result is located in `result`.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"sessionId":"0123456789abcdef0123456789abcdef","removed":true}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The server-side request ID, which is used to troubleshoot this call.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Determine the desired state of the execution based on the description of this operation.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAgentSessionResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *CancelAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAgentSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelAgentSessionResponseBody) SetJsonRpcResponse(v interface{}) *CancelAgentSessionResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *CancelAgentSessionResponseBody) SetRequestId(v string) *CancelAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAgentSessionResponseBody) SetSuccess(v bool) *CancelAgentSessionResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
