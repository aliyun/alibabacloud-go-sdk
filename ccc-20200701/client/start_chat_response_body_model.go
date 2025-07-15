// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartChatResponseBody
	GetCode() *string
	SetData(v *StartChatResponseBodyData) *StartChatResponseBody
	GetData() *StartChatResponseBodyData
	SetHttpStatusCode(v int32) *StartChatResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartChatResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartChatResponseBody
	GetRequestId() *string
}

type StartChatResponseBody struct {
	// example:
	//
	// OK
	Code *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *StartChatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartChatResponseBody) GoString() string {
	return s.String()
}

func (s *StartChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartChatResponseBody) GetData() *StartChatResponseBodyData {
	return s.Data
}

func (s *StartChatResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartChatResponseBody) SetCode(v string) *StartChatResponseBody {
	s.Code = &v
	return s
}

func (s *StartChatResponseBody) SetData(v *StartChatResponseBodyData) *StartChatResponseBody {
	s.Data = v
	return s
}

func (s *StartChatResponseBody) SetHttpStatusCode(v int32) *StartChatResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartChatResponseBody) SetMessage(v string) *StartChatResponseBody {
	s.Message = &v
	return s
}

func (s *StartChatResponseBody) SetRequestId(v string) *StartChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartChatResponseBody) Validate() error {
	return dara.Validate(s)
}

type StartChatResponseBodyData struct {
	// example:
	//
	// $23086709$EAUNIT
	ChatConversationId *string `json:"ChatConversationId,omitempty" xml:"ChatConversationId,omitempty"`
	// example:
	//
	// chat-525523618219921408
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s StartChatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartChatResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartChatResponseBodyData) GetChatConversationId() *string {
	return s.ChatConversationId
}

func (s *StartChatResponseBodyData) GetJobId() *string {
	return s.JobId
}

func (s *StartChatResponseBodyData) SetChatConversationId(v string) *StartChatResponseBodyData {
	s.ChatConversationId = &v
	return s
}

func (s *StartChatResponseBodyData) SetJobId(v string) *StartChatResponseBodyData {
	s.JobId = &v
	return s
}

func (s *StartChatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
