// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChat interface {
	dara.Model
	String() string
	GoString() string
	SetAnswer(v *ChatDetail) *Chat
	GetAnswer() *ChatDetail
	SetChatId(v string) *Chat
	GetChatId() *string
	SetExtraData(v string) *Chat
	GetExtraData() *string
	SetGmtCreateTime(v string) *Chat
	GetGmtCreateTime() *string
	SetGmtModified(v string) *Chat
	GetGmtModified() *string
	SetMessage(v string) *Chat
	GetMessage() *string
	SetOwnerId(v string) *Chat
	GetOwnerId() *string
	SetQuestion(v *ChatDetail) *Chat
	GetQuestion() *ChatDetail
	SetSessionId(v string) *Chat
	GetSessionId() *string
	SetStatus(v string) *Chat
	GetStatus() *string
	SetTitle(v string) *Chat
	GetTitle() *string
	SetUserId(v string) *Chat
	GetUserId() *string
}

type Chat struct {
	Answer    *ChatDetail `json:"Answer,omitempty" xml:"Answer,omitempty"`
	ChatId    *string     `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	ExtraData *string     `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtModified *string     `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Message     *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	OwnerId     *string     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Question    *ChatDetail `json:"Question,omitempty" xml:"Question,omitempty"`
	SessionId   *string     `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Status      *string     `json:"Status,omitempty" xml:"Status,omitempty"`
	Title       *string     `json:"Title,omitempty" xml:"Title,omitempty"`
	UserId      *string     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Chat) String() string {
	return dara.Prettify(s)
}

func (s Chat) GoString() string {
	return s.String()
}

func (s *Chat) GetAnswer() *ChatDetail {
	return s.Answer
}

func (s *Chat) GetChatId() *string {
	return s.ChatId
}

func (s *Chat) GetExtraData() *string {
	return s.ExtraData
}

func (s *Chat) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Chat) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Chat) GetMessage() *string {
	return s.Message
}

func (s *Chat) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Chat) GetQuestion() *ChatDetail {
	return s.Question
}

func (s *Chat) GetSessionId() *string {
	return s.SessionId
}

func (s *Chat) GetStatus() *string {
	return s.Status
}

func (s *Chat) GetTitle() *string {
	return s.Title
}

func (s *Chat) GetUserId() *string {
	return s.UserId
}

func (s *Chat) SetAnswer(v *ChatDetail) *Chat {
	s.Answer = v
	return s
}

func (s *Chat) SetChatId(v string) *Chat {
	s.ChatId = &v
	return s
}

func (s *Chat) SetExtraData(v string) *Chat {
	s.ExtraData = &v
	return s
}

func (s *Chat) SetGmtCreateTime(v string) *Chat {
	s.GmtCreateTime = &v
	return s
}

func (s *Chat) SetGmtModified(v string) *Chat {
	s.GmtModified = &v
	return s
}

func (s *Chat) SetMessage(v string) *Chat {
	s.Message = &v
	return s
}

func (s *Chat) SetOwnerId(v string) *Chat {
	s.OwnerId = &v
	return s
}

func (s *Chat) SetQuestion(v *ChatDetail) *Chat {
	s.Question = v
	return s
}

func (s *Chat) SetSessionId(v string) *Chat {
	s.SessionId = &v
	return s
}

func (s *Chat) SetStatus(v string) *Chat {
	s.Status = &v
	return s
}

func (s *Chat) SetTitle(v string) *Chat {
	s.Title = &v
	return s
}

func (s *Chat) SetUserId(v string) *Chat {
	s.UserId = &v
	return s
}

func (s *Chat) Validate() error {
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
