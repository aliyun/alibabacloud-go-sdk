// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *StartWorkflowInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *StartWorkflowInstancesShrinkRequest
	GetIdsShrink() *string
}

type StartWorkflowInstancesShrinkRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IDs of workflow instances.
	//
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s StartWorkflowInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartWorkflowInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *StartWorkflowInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *StartWorkflowInstancesShrinkRequest) SetComment(v string) *StartWorkflowInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *StartWorkflowInstancesShrinkRequest) SetIdsShrink(v string) *StartWorkflowInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *StartWorkflowInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
