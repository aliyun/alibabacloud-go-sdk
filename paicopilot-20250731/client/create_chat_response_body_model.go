// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v *ChatDetail) *CreateChatResponseBody
	GetAnswer() *ChatDetail
	SetChatId(v string) *CreateChatResponseBody
	GetChatId() *string
	SetErrorCode(v string) *CreateChatResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateChatResponseBody
	GetErrorMessage() *string
	SetGmtCreateTime(v string) *CreateChatResponseBody
	GetGmtCreateTime() *string
	SetRequestId(v string) *CreateChatResponseBody
	GetRequestId() *string
	SetSessionId(v string) *CreateChatResponseBody
	GetSessionId() *string
	SetStatus(v string) *CreateChatResponseBody
	GetStatus() *string
}

type CreateChatResponseBody struct {
	Answer *ChatDetail `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// chat-jkd******
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// example:
	//
	// 2187322
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Stream error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 44553E9A-******-37ADC33FE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// se-dss******
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatResponseBody) GetAnswer() *ChatDetail {
	return s.Answer
}

func (s *CreateChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreateChatResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateChatResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateChatResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CreateChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateChatResponseBody) SetAnswer(v *ChatDetail) *CreateChatResponseBody {
	s.Answer = v
	return s
}

func (s *CreateChatResponseBody) SetChatId(v string) *CreateChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateChatResponseBody) SetErrorCode(v string) *CreateChatResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateChatResponseBody) SetErrorMessage(v string) *CreateChatResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateChatResponseBody) SetGmtCreateTime(v string) *CreateChatResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *CreateChatResponseBody) SetRequestId(v string) *CreateChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatResponseBody) SetSessionId(v string) *CreateChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateChatResponseBody) SetStatus(v string) *CreateChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreateChatResponseBody) Validate() error {
	if s.Answer != nil {
		if err := s.Answer.Validate(); err != nil {
			return err
		}
	}
	return nil
}
