// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSessionModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *GetAgentSessionModelResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *GetAgentSessionModelResponseBody
	GetRequestId() *string
}

type GetAgentSessionModelResponseBody struct {
	// The defined JSON-RPC response structure. For core nested fields, see the supplementary description.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"modelId":"text-model-example","thinkingLevel":"off","thinkingLevels":["off","low","medium","high"]}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The ID of the request, which is used for troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAgentSessionModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSessionModelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSessionModelResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *GetAgentSessionModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSessionModelResponseBody) SetJsonRpcResponse(v interface{}) *GetAgentSessionModelResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *GetAgentSessionModelResponseBody) SetRequestId(v string) *GetAgentSessionModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSessionModelResponseBody) Validate() error {
	return dara.Validate(s)
}
