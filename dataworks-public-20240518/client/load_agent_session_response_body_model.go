// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoadAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetError(v interface{}) *LoadAgentSessionResponseBody
	GetError() interface{}
	SetId(v string) *LoadAgentSessionResponseBody
	GetId() *string
	SetJsonrpc(v string) *LoadAgentSessionResponseBody
	GetJsonrpc() *string
	SetMethod(v string) *LoadAgentSessionResponseBody
	GetMethod() *string
	SetParams(v interface{}) *LoadAgentSessionResponseBody
	GetParams() interface{}
	SetRequestId(v string) *LoadAgentSessionResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *LoadAgentSessionResponseBody
	GetResult() interface{}
	SetTimestamp(v int64) *LoadAgentSessionResponseBody
	GetTimestamp() *int64
}

type LoadAgentSessionResponseBody struct {
	// The error object of the SSE frame. This field is present when an error occurs.
	//
	// example:
	//
	// {"code": 400, "errorCode": "0x50000000001", "message": "not exist session", "data": null}
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// The client-generated request ID, returned from the request.
	//
	// example:
	//
	// 676303114031776
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. The value is `2.0`.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The method of the SSE frame.
	//
	// example:
	//
	// session/update
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The parameters of the SSE frame.
	//
	// example:
	//
	// {"sessionId":"af4f5ef8-e8f5-481c-ad1f-94886c6c0aed","update":{"sessionUpdate":"agent_message_chunk","content":{"type":"text","text":"hello world"}}}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// The unique request ID generated for this request.
	//
	// example:
	//
	// 0D41C608-0C60-5EB0-B986-1460909CF642
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result object of the SSE frame. This field is present when the operation is successful.
	//
	// example:
	//
	// {"stopReason":"end_turn"}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1769479322828
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s LoadAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LoadAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *LoadAgentSessionResponseBody) GetError() interface{} {
	return s.Error
}

func (s *LoadAgentSessionResponseBody) GetId() *string {
	return s.Id
}

func (s *LoadAgentSessionResponseBody) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *LoadAgentSessionResponseBody) GetMethod() *string {
	return s.Method
}

func (s *LoadAgentSessionResponseBody) GetParams() interface{} {
	return s.Params
}

func (s *LoadAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LoadAgentSessionResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *LoadAgentSessionResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *LoadAgentSessionResponseBody) SetError(v interface{}) *LoadAgentSessionResponseBody {
	s.Error = v
	return s
}

func (s *LoadAgentSessionResponseBody) SetId(v string) *LoadAgentSessionResponseBody {
	s.Id = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetJsonrpc(v string) *LoadAgentSessionResponseBody {
	s.Jsonrpc = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetMethod(v string) *LoadAgentSessionResponseBody {
	s.Method = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetParams(v interface{}) *LoadAgentSessionResponseBody {
	s.Params = v
	return s
}

func (s *LoadAgentSessionResponseBody) SetRequestId(v string) *LoadAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *LoadAgentSessionResponseBody) SetResult(v interface{}) *LoadAgentSessionResponseBody {
	s.Result = v
	return s
}

func (s *LoadAgentSessionResponseBody) SetTimestamp(v int64) *LoadAgentSessionResponseBody {
	s.Timestamp = &v
	return s
}

func (s *LoadAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
