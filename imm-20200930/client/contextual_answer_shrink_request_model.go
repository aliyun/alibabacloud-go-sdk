// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContextualAnswerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilesShrink(v string) *ContextualAnswerShrinkRequest
	GetFilesShrink() *string
	SetMessagesShrink(v string) *ContextualAnswerShrinkRequest
	GetMessagesShrink() *string
	SetProjectName(v string) *ContextualAnswerShrinkRequest
	GetProjectName() *string
}

type ContextualAnswerShrinkRequest struct {
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// This parameter is required.
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s ContextualAnswerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ContextualAnswerShrinkRequest) GoString() string {
	return s.String()
}

func (s *ContextualAnswerShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *ContextualAnswerShrinkRequest) GetMessagesShrink() *string {
	return s.MessagesShrink
}

func (s *ContextualAnswerShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ContextualAnswerShrinkRequest) SetFilesShrink(v string) *ContextualAnswerShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *ContextualAnswerShrinkRequest) SetMessagesShrink(v string) *ContextualAnswerShrinkRequest {
	s.MessagesShrink = &v
	return s
}

func (s *ContextualAnswerShrinkRequest) SetProjectName(v string) *ContextualAnswerShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *ContextualAnswerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
