// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackDialogueShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *FeedbackDialogueShrinkRequest
	GetAgentKey() *string
	SetCustomerResponse(v string) *FeedbackDialogueShrinkRequest
	GetCustomerResponse() *string
	SetGoodText(v string) *FeedbackDialogueShrinkRequest
	GetGoodText() *string
	SetModifiedResponse(v string) *FeedbackDialogueShrinkRequest
	GetModifiedResponse() *string
	SetRating(v string) *FeedbackDialogueShrinkRequest
	GetRating() *string
	SetRatingTagsShrink(v string) *FeedbackDialogueShrinkRequest
	GetRatingTagsShrink() *string
	SetSessionId(v string) *FeedbackDialogueShrinkRequest
	GetSessionId() *string
	SetTaskId(v string) *FeedbackDialogueShrinkRequest
	GetTaskId() *string
}

type FeedbackDialogueShrinkRequest struct {
	// The unique ID of the workspace. For more information, see [AgentKey](https://help.aliyun.com/document_detail/2587494.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The feedback.
	//
	// example:
	//
	// test
	CustomerResponse *string `json:"CustomerResponse,omitempty" xml:"CustomerResponse,omitempty"`
	// The generated content that is considered good.
	//
	// example:
	//
	// test
	GoodText *string `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	// The modified generated result.
	//
	// example:
	//
	// test
	ModifiedResponse *string `json:"ModifiedResponse,omitempty" xml:"ModifiedResponse,omitempty"`
	// thumbsDown: Dislike, thumbsUp: Like
	//
	// example:
	//
	// thumbsDown
	Rating *string `json:"Rating,omitempty" xml:"Rating,omitempty"`
	// The tags.
	RatingTagsShrink *string `json:"RatingTags,omitempty" xml:"RatingTags,omitempty"`
	// The ID of a single-turn conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 75bf82fa-b71b-45d7-ae40-0b00e496cd9e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The ID of the page.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FeedbackDialogueShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FeedbackDialogueShrinkRequest) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *FeedbackDialogueShrinkRequest) GetCustomerResponse() *string {
	return s.CustomerResponse
}

func (s *FeedbackDialogueShrinkRequest) GetGoodText() *string {
	return s.GoodText
}

func (s *FeedbackDialogueShrinkRequest) GetModifiedResponse() *string {
	return s.ModifiedResponse
}

func (s *FeedbackDialogueShrinkRequest) GetRating() *string {
	return s.Rating
}

func (s *FeedbackDialogueShrinkRequest) GetRatingTagsShrink() *string {
	return s.RatingTagsShrink
}

func (s *FeedbackDialogueShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FeedbackDialogueShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FeedbackDialogueShrinkRequest) SetAgentKey(v string) *FeedbackDialogueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetCustomerResponse(v string) *FeedbackDialogueShrinkRequest {
	s.CustomerResponse = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetGoodText(v string) *FeedbackDialogueShrinkRequest {
	s.GoodText = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetModifiedResponse(v string) *FeedbackDialogueShrinkRequest {
	s.ModifiedResponse = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetRating(v string) *FeedbackDialogueShrinkRequest {
	s.Rating = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetRatingTagsShrink(v string) *FeedbackDialogueShrinkRequest {
	s.RatingTagsShrink = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetSessionId(v string) *FeedbackDialogueShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetTaskId(v string) *FeedbackDialogueShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) Validate() error {
	return dara.Validate(s)
}
