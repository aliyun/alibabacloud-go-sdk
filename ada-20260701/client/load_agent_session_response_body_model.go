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
}

type LoadAgentSessionResponseBody struct {
	// The `Error` field in the frame. It is an in-stream error and is actually an object.
	//
	// example:
	//
	// {"code":-32603,"errorCode":"CLI_STREAM_FAILED","message":"The response stream was interrupted."}
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// The `Id` field in the frame, used to associate the request. Notification frames may not contain this field.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The `Jsonrpc` field in the frame. The value is `2.0`.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The `Method` field in the frame, used to distinguish notification categories.
	//
	// example:
	//
	// session/update
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The `Params` field in the frame. It is the notification payload and is actually an object.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef","update":{"sessionUpdate":"agent_message_chunk","content":{"type":"text","text":"Hello"}}}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// The request diagnostic identifier. It is not guaranteed to be included in every frame. You should also retain the diagnostic information in the response headers.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The `Result` field in the frame. It is the normal result and is actually an object.
	//
	// example:
	//
	// {"stopReason":"end_turn"}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s *LoadAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
