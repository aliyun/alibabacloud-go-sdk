// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *FeedbackRequest
	GetAgentKey() *string
	SetFeedback(v string) *FeedbackRequest
	GetFeedback() *string
	SetFeedbackContent(v string) *FeedbackRequest
	GetFeedbackContent() *string
	SetInstanceId(v string) *FeedbackRequest
	GetInstanceId() *string
	SetMessageId(v string) *FeedbackRequest
	GetMessageId() *string
	SetSessionId(v string) *FeedbackRequest
	GetSessionId() *string
}

type FeedbackRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// good
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// 这个回答很棒
	FeedbackContent *string `json:"FeedbackContent,omitempty" xml:"FeedbackContent,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 5ca40988-4f99-47ad-ac96-9060d0f81db9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 7c3cec23cc8940bc9db4a318c8f4f0aa
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s FeedbackRequest) GoString() string {
	return s.String()
}

func (s *FeedbackRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *FeedbackRequest) GetFeedback() *string {
	return s.Feedback
}

func (s *FeedbackRequest) GetFeedbackContent() *string {
	return s.FeedbackContent
}

func (s *FeedbackRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FeedbackRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *FeedbackRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FeedbackRequest) SetAgentKey(v string) *FeedbackRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackRequest) SetFeedback(v string) *FeedbackRequest {
	s.Feedback = &v
	return s
}

func (s *FeedbackRequest) SetFeedbackContent(v string) *FeedbackRequest {
	s.FeedbackContent = &v
	return s
}

func (s *FeedbackRequest) SetInstanceId(v string) *FeedbackRequest {
	s.InstanceId = &v
	return s
}

func (s *FeedbackRequest) SetMessageId(v string) *FeedbackRequest {
	s.MessageId = &v
	return s
}

func (s *FeedbackRequest) SetSessionId(v string) *FeedbackRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackRequest) Validate() error {
	return dara.Validate(s)
}
