// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *StopTaskInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *StopTaskInstancesShrinkRequest
	GetIdsShrink() *string
}

type StopTaskInstancesShrinkRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s StopTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopTaskInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *StopTaskInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *StopTaskInstancesShrinkRequest) SetComment(v string) *StopTaskInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *StopTaskInstancesShrinkRequest) SetIdsShrink(v string) *StopTaskInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *StopTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
