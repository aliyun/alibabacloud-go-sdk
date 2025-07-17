// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ResumeTaskInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *ResumeTaskInstancesShrinkRequest
	GetIdsShrink() *string
}

type ResumeTaskInstancesShrinkRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s ResumeTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ResumeTaskInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *ResumeTaskInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ResumeTaskInstancesShrinkRequest) SetComment(v string) *ResumeTaskInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *ResumeTaskInstancesShrinkRequest) SetIdsShrink(v string) *ResumeTaskInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ResumeTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
