// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ChatResponseBody
	GetCode() *string
	SetContent(v []*ChatResponseBodyContent) *ChatResponseBody
	GetContent() []*ChatResponseBodyContent
	SetCreated(v string) *ChatResponseBody
	GetCreated() *string
	SetCreatedAt(v string) *ChatResponseBody
	GetCreatedAt() *string
	SetHttpStatusCode(v int32) *ChatResponseBody
	GetHttpStatusCode() *int32
	SetId(v string) *ChatResponseBody
	GetId() *string
	SetMessage(v string) *ChatResponseBody
	GetMessage() *string
	SetObject(v string) *ChatResponseBody
	GetObject() *string
	SetRequestId(v string) *ChatResponseBody
	GetRequestId() *string
	SetRole(v string) *ChatResponseBody
	GetRole() *string
	SetSequenceNumber(v string) *ChatResponseBody
	GetSequenceNumber() *string
	SetSessionId(v string) *ChatResponseBody
	GetSessionId() *string
	SetStatus(v string) *ChatResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ChatResponseBody
	GetSuccess() *bool
	SetText(v string) *ChatResponseBody
	GetText() *string
	SetTraceId(v string) *ChatResponseBody
	GetTraceId() *string
	SetType(v string) *ChatResponseBody
	GetType() *string
}

type ChatResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The content block list (included only when Status is completed).
	Content []*ChatResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// （已废弃）创建时间戳
	//
	// example:
	//
	// 1773380609
	Created *string `json:"Created,omitempty" xml:"Created,omitempty"`
	// The creation timestamp (Unix seconds).
	//
	// example:
	//
	// 1773380609
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The unique message identifier.
	//
	// example:
	//
	// msg_xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error details (returned on failure).
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The event object type.
	//
	// example:
	//
	// response
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA12****-****-****-****-****E5C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role (user / assistant / system / tool).
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The event sequence number (an incrementing integer in string format, used to guarantee ordering). Note: When StreamOptions filters out certain event types, the filtered events still consume sequence numbers. Therefore, the sequence numbers received by the client may not be contiguous.
	//
	// example:
	//
	// 1
	SequenceNumber *string `json:"SequenceNumber,omitempty" xml:"SequenceNumber,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 176405663****961
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The reply status (created / in_progress / completed).
	//
	// example:
	//
	// in_progress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The incremental text content (included in Object=content events).
	//
	// [_single.resp.200.props.Created.desc](Deprecated) The creation timestamp.
	//
	// example:
	//
	// "hello"
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 0a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// The message type (reasoning (model thinking process) / message (formal reply)).
	//
	// example:
	//
	// message / text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBody) GoString() string {
	return s.String()
}

func (s *ChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatResponseBody) GetContent() []*ChatResponseBodyContent {
	return s.Content
}

func (s *ChatResponseBody) GetCreated() *string {
	return s.Created
}

func (s *ChatResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ChatResponseBody) GetId() *string {
	return s.Id
}

func (s *ChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatResponseBody) GetObject() *string {
	return s.Object
}

func (s *ChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatResponseBody) GetRole() *string {
	return s.Role
}

func (s *ChatResponseBody) GetSequenceNumber() *string {
	return s.SequenceNumber
}

func (s *ChatResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ChatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatResponseBody) GetText() *string {
	return s.Text
}

func (s *ChatResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ChatResponseBody) GetType() *string {
	return s.Type
}

func (s *ChatResponseBody) SetCode(v string) *ChatResponseBody {
	s.Code = &v
	return s
}

func (s *ChatResponseBody) SetContent(v []*ChatResponseBodyContent) *ChatResponseBody {
	s.Content = v
	return s
}

func (s *ChatResponseBody) SetCreated(v string) *ChatResponseBody {
	s.Created = &v
	return s
}

func (s *ChatResponseBody) SetCreatedAt(v string) *ChatResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *ChatResponseBody) SetHttpStatusCode(v int32) *ChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ChatResponseBody) SetId(v string) *ChatResponseBody {
	s.Id = &v
	return s
}

func (s *ChatResponseBody) SetMessage(v string) *ChatResponseBody {
	s.Message = &v
	return s
}

func (s *ChatResponseBody) SetObject(v string) *ChatResponseBody {
	s.Object = &v
	return s
}

func (s *ChatResponseBody) SetRequestId(v string) *ChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatResponseBody) SetRole(v string) *ChatResponseBody {
	s.Role = &v
	return s
}

func (s *ChatResponseBody) SetSequenceNumber(v string) *ChatResponseBody {
	s.SequenceNumber = &v
	return s
}

func (s *ChatResponseBody) SetSessionId(v string) *ChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *ChatResponseBody) SetStatus(v string) *ChatResponseBody {
	s.Status = &v
	return s
}

func (s *ChatResponseBody) SetSuccess(v bool) *ChatResponseBody {
	s.Success = &v
	return s
}

func (s *ChatResponseBody) SetText(v string) *ChatResponseBody {
	s.Text = &v
	return s
}

func (s *ChatResponseBody) SetTraceId(v string) *ChatResponseBody {
	s.TraceId = &v
	return s
}

func (s *ChatResponseBody) SetType(v string) *ChatResponseBody {
	s.Type = &v
	return s
}

func (s *ChatResponseBody) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChatResponseBodyContent struct {
	// The structured data (such as tool invocation). Example: {"call_id":"call_xxx","name":"get_weather", "output":"Tool output details in text format"}.
	//
	// example:
	//
	// {"call_id":"call_xxx","name":"get_weather","arguments":"{"city":"Beijing"}"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The object type.
	//
	// example:
	//
	// ""
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The text content.
	//
	// example:
	//
	// 您好
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The content type ("text" / "data").
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ChatResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ChatResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ChatResponseBodyContent) GetData() map[string]interface{} {
	return s.Data
}

func (s *ChatResponseBodyContent) GetObject() *string {
	return s.Object
}

func (s *ChatResponseBodyContent) GetText() *string {
	return s.Text
}

func (s *ChatResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *ChatResponseBodyContent) SetData(v map[string]interface{}) *ChatResponseBodyContent {
	s.Data = v
	return s
}

func (s *ChatResponseBodyContent) SetObject(v string) *ChatResponseBodyContent {
	s.Object = &v
	return s
}

func (s *ChatResponseBodyContent) SetText(v string) *ChatResponseBodyContent {
	s.Text = &v
	return s
}

func (s *ChatResponseBodyContent) SetType(v string) *ChatResponseBodyContent {
	s.Type = &v
	return s
}

func (s *ChatResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
