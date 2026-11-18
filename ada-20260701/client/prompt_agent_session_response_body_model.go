// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPromptAgentSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetError(v interface{}) *PromptAgentSessionResponseBody
	GetError() interface{}
	SetId(v string) *PromptAgentSessionResponseBody
	GetId() *string
	SetJsonrpc(v string) *PromptAgentSessionResponseBody
	GetJsonrpc() *string
	SetMethod(v string) *PromptAgentSessionResponseBody
	GetMethod() *string
	SetParams(v interface{}) *PromptAgentSessionResponseBody
	GetParams() interface{}
	SetRequestId(v string) *PromptAgentSessionResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *PromptAgentSessionResponseBody
	GetResult() interface{}
}

type PromptAgentSessionResponseBody struct {
	// The Error field in the frame. This is an in-stream error, which is an object.
	//
	// example:
	//
	// {"code":-32603,"errorCode":"CLI_STREAM_FAILED","message":"The response stream was interrupted."}
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// The Id field in the frame, used to associate the request. Notification frames may not contain this field.
	//
	// example:
	//
	// request-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The Jsonrpc field in the frame. The value is 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The Method field in the frame, used to distinguish notification categories.
	//
	// example:
	//
	// session/update
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The Params field in the frame. This is the notification payload, which is an object.
	//
	// example:
	//
	// {"sessionId":"0123456789abcdef0123456789abcdef","update":{"sessionUpdate":"agent_message_chunk","content":{"type":"text","text":"You should first check whether the list is empty."}},"_meta":{"offset":12,"eventEpoch":"epoch-example","promptId":"prompt-example"}}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// The request diagnostic identifier. This identifier is not guaranteed to be included in every frame. You should also retain the diagnostic information in the response headers.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Result field in the frame. This is the normal result, which is an object.
	//
	// example:
	//
	// {"stopReason":"end_turn"}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PromptAgentSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionResponseBody) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionResponseBody) GetError() interface{} {
	return s.Error
}

func (s *PromptAgentSessionResponseBody) GetId() *string {
	return s.Id
}

func (s *PromptAgentSessionResponseBody) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *PromptAgentSessionResponseBody) GetMethod() *string {
	return s.Method
}

func (s *PromptAgentSessionResponseBody) GetParams() interface{} {
	return s.Params
}

func (s *PromptAgentSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PromptAgentSessionResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *PromptAgentSessionResponseBody) SetError(v interface{}) *PromptAgentSessionResponseBody {
	s.Error = v
	return s
}

func (s *PromptAgentSessionResponseBody) SetId(v string) *PromptAgentSessionResponseBody {
	s.Id = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetJsonrpc(v string) *PromptAgentSessionResponseBody {
	s.Jsonrpc = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetMethod(v string) *PromptAgentSessionResponseBody {
	s.Method = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetParams(v interface{}) *PromptAgentSessionResponseBody {
	s.Params = v
	return s
}

func (s *PromptAgentSessionResponseBody) SetRequestId(v string) *PromptAgentSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *PromptAgentSessionResponseBody) SetResult(v interface{}) *PromptAgentSessionResponseBody {
	s.Result = v
	return s
}

func (s *PromptAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
