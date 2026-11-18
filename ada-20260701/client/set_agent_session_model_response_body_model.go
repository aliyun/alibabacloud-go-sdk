// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentSessionModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonRpcResponse(v interface{}) *SetAgentSessionModelResponseBody
	GetJsonRpcResponse() interface{}
	SetRequestId(v string) *SetAgentSessionModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetAgentSessionModelResponseBody
	GetSuccess() *bool
}

type SetAgentSessionModelResponseBody struct {
	// The defined JSON-RPC response structure. For more information about the core nested fields, see the supplementary description.
	//
	// example:
	//
	// {"jsonrpc":"2.0","id":"request-001","result":{"modelId":"text-model-example","thinkingLevel":"off","thinkingLevels":["off","low","medium","high"]}}
	JsonRpcResponse interface{} `json:"JsonRpcResponse,omitempty" xml:"JsonRpcResponse,omitempty"`
	// The ID of the current request, which is used for troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the setting is successful. The value is true if the setting is successful. Read the actual model and thinking level from the business result.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAgentSessionModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAgentSessionModelResponseBody) GoString() string {
	return s.String()
}

func (s *SetAgentSessionModelResponseBody) GetJsonRpcResponse() interface{} {
	return s.JsonRpcResponse
}

func (s *SetAgentSessionModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAgentSessionModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetAgentSessionModelResponseBody) SetJsonRpcResponse(v interface{}) *SetAgentSessionModelResponseBody {
	s.JsonRpcResponse = v
	return s
}

func (s *SetAgentSessionModelResponseBody) SetRequestId(v string) *SetAgentSessionModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAgentSessionModelResponseBody) SetSuccess(v bool) *SetAgentSessionModelResponseBody {
	s.Success = &v
	return s
}

func (s *SetAgentSessionModelResponseBody) Validate() error {
	return dara.Validate(s)
}
