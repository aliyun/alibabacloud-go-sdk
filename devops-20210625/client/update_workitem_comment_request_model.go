// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkitemCommentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommentId(v int64) *UpdateWorkitemCommentRequest
	GetCommentId() *int64
	SetContent(v string) *UpdateWorkitemCommentRequest
	GetContent() *string
	SetFormatType(v string) *UpdateWorkitemCommentRequest
	GetFormatType() *string
	SetWorkitemIdentifier(v string) *UpdateWorkitemCommentRequest
	GetWorkitemIdentifier() *string
}

type UpdateWorkitemCommentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1964584
	CommentId *int64 `json:"commentId,omitempty" xml:"commentId,omitempty"`
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MARKDOWN/RICHTEXT
	FormatType *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s UpdateWorkitemCommentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkitemCommentRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkitemCommentRequest) GetCommentId() *int64 {
	return s.CommentId
}

func (s *UpdateWorkitemCommentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateWorkitemCommentRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *UpdateWorkitemCommentRequest) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *UpdateWorkitemCommentRequest) SetCommentId(v int64) *UpdateWorkitemCommentRequest {
	s.CommentId = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) SetContent(v string) *UpdateWorkitemCommentRequest {
	s.Content = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) SetFormatType(v string) *UpdateWorkitemCommentRequest {
	s.FormatType = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) SetWorkitemIdentifier(v string) *UpdateWorkitemCommentRequest {
	s.WorkitemIdentifier = &v
	return s
}

func (s *UpdateWorkitemCommentRequest) Validate() error {
	return dara.Validate(s)
}
