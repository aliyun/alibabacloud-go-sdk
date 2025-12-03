// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemCommentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CreateWorkitemCommentRequest
	GetContent() *string
	SetFormatType(v string) *CreateWorkitemCommentRequest
	GetFormatType() *string
	SetParentId(v string) *CreateWorkitemCommentRequest
	GetParentId() *string
	SetWorkitemIdentifier(v string) *CreateWorkitemCommentRequest
	GetWorkitemIdentifier() *string
}

type CreateWorkitemCommentRequest struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RICHTEXT/MARKDOWN
	FormatType *string `json:"formatType,omitempty" xml:"formatType,omitempty"`
	// example:
	//
	// 26842
	ParentId *string `json:"parentId,omitempty" xml:"parentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1e9903d8b3f12xxxxxf9286ef5
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemCommentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemCommentRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkitemCommentRequest) GetContent() *string {
	return s.Content
}

func (s *CreateWorkitemCommentRequest) GetFormatType() *string {
	return s.FormatType
}

func (s *CreateWorkitemCommentRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateWorkitemCommentRequest) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemCommentRequest) SetContent(v string) *CreateWorkitemCommentRequest {
	s.Content = &v
	return s
}

func (s *CreateWorkitemCommentRequest) SetFormatType(v string) *CreateWorkitemCommentRequest {
	s.FormatType = &v
	return s
}

func (s *CreateWorkitemCommentRequest) SetParentId(v string) *CreateWorkitemCommentRequest {
	s.ParentId = &v
	return s
}

func (s *CreateWorkitemCommentRequest) SetWorkitemIdentifier(v string) *CreateWorkitemCommentRequest {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemCommentRequest) Validate() error {
	return dara.Validate(s)
}
