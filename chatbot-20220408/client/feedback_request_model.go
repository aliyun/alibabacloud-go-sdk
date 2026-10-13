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
	// The business space key. If this parameter is not specified, the default business space is accessed. You can obtain the key value on the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The feedback type, which corresponds to the FeedbackType field in the session history API.
	//
	// Valid values:
	//
	// - good: Like.
	//
	// - bad: Dislike.
	//
	// example:
	//
	// good
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The content of the like or dislike feedback. You can directly enter the feedback content or pass in a JSON string as shown in the example. The content field corresponds to FeedbackUserInfo in the session history, and feedbackLabels corresponds to FeedbackLabels in the session history.
	//
	// example:
	//
	// {
	//
	//     "content": "I think this answer is incorrect. It should be based on facts.",
	//
	//     "feedbackLabels": [
	//
	//         "Contains factual errors"
	//
	//     ]
	//
	// }
	FeedbackContent *string `json:"FeedbackContent,omitempty" xml:"FeedbackContent,omitempty"`
	// The unique identifier of the chatbot.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID that uniquely identifies a single message in a session window.
	//
	// example:
	//
	// 5ca40988-4f99-47ad-ac96-9060d0f81db9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The session ID that uniquely identifies an IM session.
	//
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
