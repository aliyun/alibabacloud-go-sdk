// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTaskInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *SuspendTaskInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *SuspendTaskInstancesShrinkRequest
	GetIdsShrink() *string
}

type SuspendTaskInstancesShrinkRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s SuspendTaskInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendTaskInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SuspendTaskInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *SuspendTaskInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *SuspendTaskInstancesShrinkRequest) SetComment(v string) *SuspendTaskInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *SuspendTaskInstancesShrinkRequest) SetIdsShrink(v string) *SuspendTaskInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *SuspendTaskInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
