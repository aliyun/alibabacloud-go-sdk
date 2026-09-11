// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranscribeChatVoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TranscribeChatVoiceResponseBody
	GetCode() *string
	SetMessage(v string) *TranscribeChatVoiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *TranscribeChatVoiceResponseBody
	GetRequestId() *string
	SetText(v string) *TranscribeChatVoiceResponseBody
	GetText() *string
}

type TranscribeChatVoiceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The recognized text content.
	//
	// example:
	//
	// The weather is nice today
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s TranscribeChatVoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TranscribeChatVoiceResponseBody) GoString() string {
	return s.String()
}

func (s *TranscribeChatVoiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *TranscribeChatVoiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TranscribeChatVoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TranscribeChatVoiceResponseBody) GetText() *string {
	return s.Text
}

func (s *TranscribeChatVoiceResponseBody) SetCode(v string) *TranscribeChatVoiceResponseBody {
	s.Code = &v
	return s
}

func (s *TranscribeChatVoiceResponseBody) SetMessage(v string) *TranscribeChatVoiceResponseBody {
	s.Message = &v
	return s
}

func (s *TranscribeChatVoiceResponseBody) SetRequestId(v string) *TranscribeChatVoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *TranscribeChatVoiceResponseBody) SetText(v string) *TranscribeChatVoiceResponseBody {
	s.Text = &v
	return s
}

func (s *TranscribeChatVoiceResponseBody) Validate() error {
	return dara.Validate(s)
}
