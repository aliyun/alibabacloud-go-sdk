// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *FeedbackDialogueRequest
	GetAgentKey() *string
	SetCustomerResponse(v string) *FeedbackDialogueRequest
	GetCustomerResponse() *string
	SetGoodText(v string) *FeedbackDialogueRequest
	GetGoodText() *string
	SetModifiedResponse(v string) *FeedbackDialogueRequest
	GetModifiedResponse() *string
	SetRating(v string) *FeedbackDialogueRequest
	GetRating() *string
	SetRatingTags(v []*string) *FeedbackDialogueRequest
	GetRatingTags() []*string
	SetSessionId(v string) *FeedbackDialogueRequest
	GetSessionId() *string
	SetTaskId(v string) *FeedbackDialogueRequest
	GetTaskId() *string
}

type FeedbackDialogueRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// fcb14f25c9ee41ccad33a049de8f941b_p_outbound_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// test
	CustomerResponse *string `json:"CustomerResponse,omitempty" xml:"CustomerResponse,omitempty"`
	// example:
	//
	// test
	GoodText *string `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	// example:
	//
	// test
	ModifiedResponse *string `json:"ModifiedResponse,omitempty" xml:"ModifiedResponse,omitempty"`
	// example:
	//
	// thumbsDown
	Rating     *string   `json:"Rating,omitempty" xml:"Rating,omitempty"`
	RatingTags []*string `json:"RatingTags,omitempty" xml:"RatingTags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 75bf82fa-b71b-45d7-ae40-0b00e496cd9e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FeedbackDialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s FeedbackDialogueRequest) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *FeedbackDialogueRequest) GetCustomerResponse() *string {
	return s.CustomerResponse
}

func (s *FeedbackDialogueRequest) GetGoodText() *string {
	return s.GoodText
}

func (s *FeedbackDialogueRequest) GetModifiedResponse() *string {
	return s.ModifiedResponse
}

func (s *FeedbackDialogueRequest) GetRating() *string {
	return s.Rating
}

func (s *FeedbackDialogueRequest) GetRatingTags() []*string {
	return s.RatingTags
}

func (s *FeedbackDialogueRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *FeedbackDialogueRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FeedbackDialogueRequest) SetAgentKey(v string) *FeedbackDialogueRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackDialogueRequest) SetCustomerResponse(v string) *FeedbackDialogueRequest {
	s.CustomerResponse = &v
	return s
}

func (s *FeedbackDialogueRequest) SetGoodText(v string) *FeedbackDialogueRequest {
	s.GoodText = &v
	return s
}

func (s *FeedbackDialogueRequest) SetModifiedResponse(v string) *FeedbackDialogueRequest {
	s.ModifiedResponse = &v
	return s
}

func (s *FeedbackDialogueRequest) SetRating(v string) *FeedbackDialogueRequest {
	s.Rating = &v
	return s
}

func (s *FeedbackDialogueRequest) SetRatingTags(v []*string) *FeedbackDialogueRequest {
	s.RatingTags = v
	return s
}

func (s *FeedbackDialogueRequest) SetSessionId(v string) *FeedbackDialogueRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackDialogueRequest) SetTaskId(v string) *FeedbackDialogueRequest {
	s.TaskId = &v
	return s
}

func (s *FeedbackDialogueRequest) Validate() error {
	return dara.Validate(s)
}
