// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RerunTaskInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *RerunTaskInstancesShrinkRequest
	GetIdsShrink() *string
}

type RerunTaskInstancesShrinkRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s RerunTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RerunTaskInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *RerunTaskInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *RerunTaskInstancesShrinkRequest) SetComment(v string) *RerunTaskInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *RerunTaskInstancesShrinkRequest) SetIdsShrink(v string) *RerunTaskInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *RerunTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
