// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *KopilotFeedbackRequest
	GetComment() *string
	SetFeedback(v string) *KopilotFeedbackRequest
	GetFeedback() *string
	SetRegionId(v string) *KopilotFeedbackRequest
	GetRegionId() *string
	SetSessionId(v string) *KopilotFeedbackRequest
	GetSessionId() *string
	SetTurnId(v string) *KopilotFeedbackRequest
	GetTurnId() *string
}

type KopilotFeedbackRequest struct {
	Comment  *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TurnId    *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
}

func (s KopilotFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotFeedbackRequest) GoString() string {
	return s.String()
}

func (s *KopilotFeedbackRequest) GetComment() *string {
	return s.Comment
}

func (s *KopilotFeedbackRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *KopilotFeedbackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotFeedbackRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotFeedbackRequest) GetTurnId() *string {
	return s.TurnId
}

func (s *KopilotFeedbackRequest) SetComment(v string) *KopilotFeedbackRequest {
	s.Comment = &v
	return s
}

func (s *KopilotFeedbackRequest) SetFeedback(v string) *KopilotFeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *KopilotFeedbackRequest) SetRegionId(v string) *KopilotFeedbackRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotFeedbackRequest) SetSessionId(v string) *KopilotFeedbackRequest {
	s.SessionId = &v
	return s
}

func (s *KopilotFeedbackRequest) SetTurnId(v string) *KopilotFeedbackRequest {
	s.TurnId = &v
	return s
}

func (s *KopilotFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
