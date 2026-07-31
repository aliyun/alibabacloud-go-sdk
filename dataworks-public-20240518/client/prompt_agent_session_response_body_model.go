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
	SetTimestamp(v int64) *PromptAgentSessionResponseBody
	GetTimestamp() *int64
}

type PromptAgentSessionResponseBody struct {
	// The SSE frame error message. The returned content conforms to the open-source Agent Client Protocol (ACP) specification. For more information, visit: https://agentclientprotocol.com/protocol/prompt-turn.
	//
	// example:
	//
	// {"code": 400, "errorCode": "0x50000000001", "message": "not exist session", "data": null}
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// The ID passed by the requester. The value is returned as-is.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Fixed value: 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The SSE method. The returned content conforms to the open-source Agent Client Protocol (ACP) specification. For more information, visit: https://agentclientprotocol.com/protocol/prompt-turn.
	//
	// example:
	//
	// session/update
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The SSE params. The returned content conforms to the open-source Agent Client Protocol (ACP) specification. For more information, visit: https://agentclientprotocol.com/protocol/prompt-turn.
	//
	// example:
	//
	// {"sessionId":"af4f5ef8-e8f5-481c-ad1f-94886c6c0aed","update":{"sessionUpdate":"agent_message_chunk","content":{"type":"text","text":"hello world"}}}
	Params interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D5D70885-7CC7-594A-80C7-2EF1B00FFB4B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The SSE frame result set. The returned content conforms to the open-source Agent Client Protocol (ACP) specification. For more information, visit: https://agentclientprotocol.com/protocol/prompt-turn.
	//
	// example:
	//
	// {"stopReason":"end_turn"}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1747447032
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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

func (s *PromptAgentSessionResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
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

func (s *PromptAgentSessionResponseBody) SetTimestamp(v int64) *PromptAgentSessionResponseBody {
	s.Timestamp = &v
	return s
}

func (s *PromptAgentSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
