// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedback interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *Feedback
	GetChatId() *string
	SetFeedbackAction(v string) *Feedback
	GetFeedbackAction() *string
	SetFeedbackId(v string) *Feedback
	GetFeedbackId() *string
	SetFeedbackMessage(v string) *Feedback
	GetFeedbackMessage() *string
	SetGmtCreateTime(v string) *Feedback
	GetGmtCreateTime() *string
	SetGmtModified(v string) *Feedback
	GetGmtModified() *string
	SetOwnerId(v string) *Feedback
	GetOwnerId() *string
	SetSessionId(v string) *Feedback
	GetSessionId() *string
	SetUserId(v string) *Feedback
	GetUserId() *string
}

type Feedback struct {
	ChatId          *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	FeedbackAction  *string `json:"FeedbackAction,omitempty" xml:"FeedbackAction,omitempty"`
	FeedbackId      *string `json:"FeedbackId,omitempty" xml:"FeedbackId,omitempty"`
	FeedbackMessage *string `json:"FeedbackMessage,omitempty" xml:"FeedbackMessage,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SessionId   *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Feedback) String() string {
	return dara.Prettify(s)
}

func (s Feedback) GoString() string {
	return s.String()
}

func (s *Feedback) GetChatId() *string {
	return s.ChatId
}

func (s *Feedback) GetFeedbackAction() *string {
	return s.FeedbackAction
}

func (s *Feedback) GetFeedbackId() *string {
	return s.FeedbackId
}

func (s *Feedback) GetFeedbackMessage() *string {
	return s.FeedbackMessage
}

func (s *Feedback) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Feedback) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Feedback) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Feedback) GetSessionId() *string {
	return s.SessionId
}

func (s *Feedback) GetUserId() *string {
	return s.UserId
}

func (s *Feedback) SetChatId(v string) *Feedback {
	s.ChatId = &v
	return s
}

func (s *Feedback) SetFeedbackAction(v string) *Feedback {
	s.FeedbackAction = &v
	return s
}

func (s *Feedback) SetFeedbackId(v string) *Feedback {
	s.FeedbackId = &v
	return s
}

func (s *Feedback) SetFeedbackMessage(v string) *Feedback {
	s.FeedbackMessage = &v
	return s
}

func (s *Feedback) SetGmtCreateTime(v string) *Feedback {
	s.GmtCreateTime = &v
	return s
}

func (s *Feedback) SetGmtModified(v string) *Feedback {
	s.GmtModified = &v
	return s
}

func (s *Feedback) SetOwnerId(v string) *Feedback {
	s.OwnerId = &v
	return s
}

func (s *Feedback) SetSessionId(v string) *Feedback {
	s.SessionId = &v
	return s
}

func (s *Feedback) SetUserId(v string) *Feedback {
	s.UserId = &v
	return s
}

func (s *Feedback) Validate() error {
	return dara.Validate(s)
}
