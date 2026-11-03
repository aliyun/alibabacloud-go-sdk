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
	SetMeta(v *PromptAgentSessionResponseBodyMeta) *PromptAgentSessionResponseBody
	GetMeta() *PromptAgentSessionResponseBodyMeta
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
	// The error message in the SSE frame. The returned content complies with the open-source Agent Client Protocol (ACP) specification. For more information, see https://agentclientprotocol.com/protocol/prompt-turn.
	//
	// example:
	//
	// {"code": 400, "errorCode": "0x50000000001", "message": "not exist session", "data": null}
	Error interface{} `json:"Error,omitempty" xml:"Error,omitempty"`
	// The ID passed by the requester. The value is returned as is.
	//
	// example:
	//
	// 631478864897630XXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC version. Set the value to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// Carried only by verified native events. The client saves the EventCursor after fully applying the event, and uses EventId and EventEpoch to resume the subscription to GetAgentSessionEvents after a disconnection. Common subscription errors do not contain proof of the native execution terminal state.
	Meta *PromptAgentSessionResponseBodyMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	// The method of the SSE. The returned content complies with the open-source Agent Client Protocol (ACP) specification. For more information, see https://agentclientprotocol.com/protocol/prompt-turn.
	//
	// example:
	//
	// session/update
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The params of the SSE. The returned content complies with the open-source Agent Client Protocol (ACP) specification. For more information, see https://agentclientprotocol.com/protocol/prompt-turn.
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
	// The result set of the SSE frame. The returned content complies with the open-source Agent Client Protocol (ACP) specification. For more information, see https://agentclientprotocol.com/protocol/prompt-turn.
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

func (s *PromptAgentSessionResponseBody) GetMeta() *PromptAgentSessionResponseBodyMeta {
	return s.Meta
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

func (s *PromptAgentSessionResponseBody) SetMeta(v *PromptAgentSessionResponseBodyMeta) *PromptAgentSessionResponseBody {
	s.Meta = v
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
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PromptAgentSessionResponseBodyMeta struct {
	// The resumption cursor of the current native event. Carried only by verified native events. The client saves the cursor after fully processing the event, and uses the EventId and EventEpoch in it to resume the subscription to GetAgentSessionEvents after a disconnection without resending the prompt.
	EventCursor *PromptAgentSessionResponseBodyMetaEventCursor `json:"EventCursor,omitempty" xml:"EventCursor,omitempty" type:"Struct"`
}

func (s PromptAgentSessionResponseBodyMeta) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionResponseBodyMeta) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionResponseBodyMeta) GetEventCursor() *PromptAgentSessionResponseBodyMetaEventCursor {
	return s.EventCursor
}

func (s *PromptAgentSessionResponseBodyMeta) SetEventCursor(v *PromptAgentSessionResponseBodyMetaEventCursor) *PromptAgentSessionResponseBodyMeta {
	s.EventCursor = v
	return s
}

func (s *PromptAgentSessionResponseBodyMeta) Validate() error {
	if s.EventCursor != nil {
		if err := s.EventCursor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PromptAgentSessionResponseBodyMetaEventCursor struct {
	// The event stream generation identifier returned when the current prompt is accepted. It is paired with EventId during resumption. Pass this value as is to the EventEpoch query parameter of GetAgentSessionEvents to verify the same event stream.
	//
	// example:
	//
	// f7c6a27a-09f6-4bc9-b924-0c360351373c
	EventEpoch *string `json:"EventEpoch,omitempty" xml:"EventEpoch,omitempty"`
	// The integer ID of the current native event, returned as a decimal string. The client saves it after fully processing the event, and passes this value as the Last-Event-Id HTTP header of GetAgentSessionEvents during resumption to continue receiving subsequent events.
	//
	// example:
	//
	// 4
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The type of the current native event, obtained from the type field of the native event. The examples do not constitute a fixed enumeration.
	//
	// example:
	//
	// session_update
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The execution identifier returned when the current prompt is accepted. It is used to associate the current execution with its events.
	//
	// example:
	//
	// p1
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// The identifier of the current DataAgent session. It is used to associate the event with its session.
	//
	// example:
	//
	// 79fea58a-45f5-45c6-9169-c59beac0d5b6
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s PromptAgentSessionResponseBodyMetaEventCursor) String() string {
	return dara.Prettify(s)
}

func (s PromptAgentSessionResponseBodyMetaEventCursor) GoString() string {
	return s.String()
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) GetEventEpoch() *string {
	return s.EventEpoch
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) GetEventId() *string {
	return s.EventId
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) GetEventType() *string {
	return s.EventType
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) GetPromptId() *string {
	return s.PromptId
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) GetSessionId() *string {
	return s.SessionId
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) SetEventEpoch(v string) *PromptAgentSessionResponseBodyMetaEventCursor {
	s.EventEpoch = &v
	return s
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) SetEventId(v string) *PromptAgentSessionResponseBodyMetaEventCursor {
	s.EventId = &v
	return s
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) SetEventType(v string) *PromptAgentSessionResponseBodyMetaEventCursor {
	s.EventType = &v
	return s
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) SetPromptId(v string) *PromptAgentSessionResponseBodyMetaEventCursor {
	s.PromptId = &v
	return s
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) SetSessionId(v string) *PromptAgentSessionResponseBodyMetaEventCursor {
	s.SessionId = &v
	return s
}

func (s *PromptAgentSessionResponseBodyMetaEventCursor) Validate() error {
	return dara.Validate(s)
}
