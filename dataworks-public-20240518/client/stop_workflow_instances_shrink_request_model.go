// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWorkflowInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *StopWorkflowInstancesShrinkRequest
	GetComment() *string
	SetIdsShrink(v string) *StopWorkflowInstancesShrinkRequest
	GetIdsShrink() *string
}

type StopWorkflowInstancesShrinkRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The workflow instance IDs.
	//
	// This parameter is required.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s StopWorkflowInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopWorkflowInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopWorkflowInstancesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *StopWorkflowInstancesShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *StopWorkflowInstancesShrinkRequest) SetComment(v string) *StopWorkflowInstancesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *StopWorkflowInstancesShrinkRequest) SetIdsShrink(v string) *StopWorkflowInstancesShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *StopWorkflowInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
