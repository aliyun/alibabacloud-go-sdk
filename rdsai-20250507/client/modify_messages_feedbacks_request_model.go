// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMessagesFeedbacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ModifyMessagesFeedbacksRequest
	GetContent() *string
	SetMessageId(v string) *ModifyMessagesFeedbacksRequest
	GetMessageId() *string
	SetRating(v string) *ModifyMessagesFeedbacksRequest
	GetRating() *string
	SetWorkspaceId(v string) *ModifyMessagesFeedbacksRequest
	GetWorkspaceId() *string
}

type ModifyMessagesFeedbacksRequest struct {
	// The feedback content.
	//
	// example:
	//
	// Unable to understand context; irrelevant answer
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The message ID.
	//
	// example:
	//
	// yy9rkn6q-js75-0dka-0cc2-6b5o86uj****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The rating.
	//
	// example:
	//
	// like
	Rating *string `json:"Rating,omitempty" xml:"Rating,omitempty"`
	// The ContextDB workspace ID. Required only for ContextDB Manager App requests.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ModifyMessagesFeedbacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMessagesFeedbacksRequest) GoString() string {
	return s.String()
}

func (s *ModifyMessagesFeedbacksRequest) GetContent() *string {
	return s.Content
}

func (s *ModifyMessagesFeedbacksRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *ModifyMessagesFeedbacksRequest) GetRating() *string {
	return s.Rating
}

func (s *ModifyMessagesFeedbacksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyMessagesFeedbacksRequest) SetContent(v string) *ModifyMessagesFeedbacksRequest {
	s.Content = &v
	return s
}

func (s *ModifyMessagesFeedbacksRequest) SetMessageId(v string) *ModifyMessagesFeedbacksRequest {
	s.MessageId = &v
	return s
}

func (s *ModifyMessagesFeedbacksRequest) SetRating(v string) *ModifyMessagesFeedbacksRequest {
	s.Rating = &v
	return s
}

func (s *ModifyMessagesFeedbacksRequest) SetWorkspaceId(v string) *ModifyMessagesFeedbacksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyMessagesFeedbacksRequest) Validate() error {
	return dara.Validate(s)
}
