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
	// The user feedback comment.
	//
	// example:
	//
	// good
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The user satisfaction level. Valid values:
	//
	// - 1: satisfied
	//
	// - -1: not satisfied
	//
	// - 0: cancel the evaluation
	//
	// example:
	//
	// 1
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// e356c91c-8220-425c-9d86-********
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The primary key ID.
	//
	// example:
	//
	// 5243231*****
	TurnId *string `json:"TurnId,omitempty" xml:"TurnId,omitempty"`
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
