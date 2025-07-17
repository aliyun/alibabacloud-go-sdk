// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *SetSuccessTaskInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *SetSuccessTaskInstancesShrinkRequest
	GetIdsShrink() *string
}

type SetSuccessTaskInstancesShrinkRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s SetSuccessTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetSuccessTaskInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *SetSuccessTaskInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *SetSuccessTaskInstancesShrinkRequest) SetComment(v string) *SetSuccessTaskInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *SetSuccessTaskInstancesShrinkRequest) SetIdsShrink(v string) *SetSuccessTaskInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *SetSuccessTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
