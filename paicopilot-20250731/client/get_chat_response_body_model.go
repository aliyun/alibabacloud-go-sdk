// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v *ChatDetail) *GetChatResponseBody
	GetAnswer() *ChatDetail
	SetChatId(v string) *GetChatResponseBody
	GetChatId() *string
	SetExtraData(v string) *GetChatResponseBody
	GetExtraData() *string
	SetGmtCreateTime(v string) *GetChatResponseBody
	GetGmtCreateTime() *string
	SetGmtModified(v string) *GetChatResponseBody
	GetGmtModified() *string
	SetMessage(v string) *GetChatResponseBody
	GetMessage() *string
	SetOwnerId(v string) *GetChatResponseBody
	GetOwnerId() *string
	SetQuestion(v *ChatDetail) *GetChatResponseBody
	GetQuestion() *ChatDetail
	SetRequestId(v string) *GetChatResponseBody
	GetRequestId() *string
	SetSessionId(v string) *GetChatResponseBody
	GetSessionId() *string
	SetStatus(v string) *GetChatResponseBody
	GetStatus() *string
	SetTitle(v string) *GetChatResponseBody
	GetTitle() *string
	SetUserId(v string) *GetChatResponseBody
	GetUserId() *string
}

type GetChatResponseBody struct {
	Answer *ChatDetail `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// chat-jkd******
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// example:
	//
	// {"solutionDetail":{"formValues":{"params":{"InstanceId":"dsw-g54******qg9"},"content":{"EcsSpec":"ecs.gn6i-c8g1.2xlarge"}},"success":true}}
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2025-12-01T17:52:05+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// how to ******
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1500******860
	OwnerId  *string     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Question *ChatDetail `json:"Question,omitempty" xml:"Question,omitempty"`
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
	// example:
	//
	// how to ******
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 2100******360
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatResponseBody) GetAnswer() *ChatDetail {
	return s.Answer
}

func (s *GetChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *GetChatResponseBody) GetExtraData() *string {
	return s.ExtraData
}

func (s *GetChatResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetChatResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetChatResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetChatResponseBody) GetQuestion() *ChatDetail {
	return s.Question
}

func (s *GetChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetChatResponseBody) GetTitle() *string {
	return s.Title
}

func (s *GetChatResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetChatResponseBody) SetAnswer(v *ChatDetail) *GetChatResponseBody {
	s.Answer = v
	return s
}

func (s *GetChatResponseBody) SetChatId(v string) *GetChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *GetChatResponseBody) SetExtraData(v string) *GetChatResponseBody {
	s.ExtraData = &v
	return s
}

func (s *GetChatResponseBody) SetGmtCreateTime(v string) *GetChatResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetChatResponseBody) SetGmtModified(v string) *GetChatResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetChatResponseBody) SetMessage(v string) *GetChatResponseBody {
	s.Message = &v
	return s
}

func (s *GetChatResponseBody) SetOwnerId(v string) *GetChatResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetChatResponseBody) SetQuestion(v *ChatDetail) *GetChatResponseBody {
	s.Question = v
	return s
}

func (s *GetChatResponseBody) SetRequestId(v string) *GetChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatResponseBody) SetSessionId(v string) *GetChatResponseBody {
	s.SessionId = &v
	return s
}

func (s *GetChatResponseBody) SetStatus(v string) *GetChatResponseBody {
	s.Status = &v
	return s
}

func (s *GetChatResponseBody) SetTitle(v string) *GetChatResponseBody {
	s.Title = &v
	return s
}

func (s *GetChatResponseBody) SetUserId(v string) *GetChatResponseBody {
	s.UserId = &v
	return s
}

func (s *GetChatResponseBody) Validate() error {
	if s.Answer != nil {
		if err := s.Answer.Validate(); err != nil {
			return err
		}
	}
	if s.Question != nil {
		if err := s.Question.Validate(); err != nil {
			return err
		}
	}
	return nil
}
