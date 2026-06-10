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
	// The key for the business space. If this parameter is not set, the system uses the default business space. You can obtain the key from the business management page of your primary account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The feedback rating for the response. This parameter corresponds to `FeedbackType` in the session history API.
	//
	// Enumerated values: \\"good\\" (a positive rating) and \\"bad\\" (a negative rating).
	//
	// example:
	//
	// good
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The detailed feedback content. You can provide this as a raw string or as a JSON string. If you use a JSON string, the \\"content\\" field corresponds to \\"FeedbackUserInfo\\" and the \\"feedbackLabels\\" field corresponds to \\"FeedbackLabels\\" in the session history.
	//
	// example:
	//
	// {
	//
	//     "content": "我觉得这个答案不对，需要依据事实回答。",
	//
	//     "feedbackLabels": [
	//
	//         "存在事实性错误"
	//
	//     ]
	//
	// }
	FeedbackContent *string `json:"FeedbackContent,omitempty" xml:"FeedbackContent,omitempty"`
	// The unique identifier of the chatbot instance.
	//
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unique identifier of a single message within the session.
	//
	// example:
	//
	// 5ca40988-4f99-47ad-ac96-9060d0f81db9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The unique identifier for the session. The instant messaging (IM) system uses this ID to track the conversation.
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
