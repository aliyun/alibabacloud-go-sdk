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
	// The content of the files involved in the current Q&A. It is recommended to use the return value of the ContextualRetrieval interface as input.
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// Yes, the history of conversations and tool invocations. The latest message is at the end (index n-1), and the oldest message is at the beginning (index 0).
	//
	// It must be in the form of user-assistant pairs, with a total count of 2*n+1, and the length of the latest question should not exceed 1000 characters.
	//
	// The length of the historical conversation is limited to 100.
	//
	// This parameter is required.
	MessagesShrink *string `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// Project name. For how to obtain it, see [Creating a Project](https://help.aliyun.com/zh/imm/getting-started/create-a-project-1?spm=a2c4g.11186623.help-menu-search-62354.d_0).
	//
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
