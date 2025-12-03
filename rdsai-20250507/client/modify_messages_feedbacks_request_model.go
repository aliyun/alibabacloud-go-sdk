// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMessagesFeedbacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *ModifyMessagesFeedbacksRequest
	GetApiId() *string
	SetContent(v string) *ModifyMessagesFeedbacksRequest
	GetContent() *string
	SetMessageId(v string) *ModifyMessagesFeedbacksRequest
	GetMessageId() *string
	SetRating(v string) *ModifyMessagesFeedbacksRequest
	GetRating() *string
}

type ModifyMessagesFeedbacksRequest struct {
	// example:
	//
	// app-iBuGU1VxEY42zrQRQfNA****
	ApiId   *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// yy9rkn6q-js75-0dka-0cc2-6b5o86uj****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// like
	Rating *string `json:"Rating,omitempty" xml:"Rating,omitempty"`
}

func (s ModifyMessagesFeedbacksRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMessagesFeedbacksRequest) GoString() string {
	return s.String()
}

func (s *ModifyMessagesFeedbacksRequest) GetApiId() *string {
	return s.ApiId
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

func (s *ModifyMessagesFeedbacksRequest) SetApiId(v string) *ModifyMessagesFeedbacksRequest {
	s.ApiId = &v
	return s
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

func (s *ModifyMessagesFeedbacksRequest) Validate() error {
	return dara.Validate(s)
}
