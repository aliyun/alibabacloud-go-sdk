// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunQuickWritingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticlesShrink(v string) *RunQuickWritingShrinkRequest
	GetArticlesShrink() *string
	SetPrompt(v string) *RunQuickWritingShrinkRequest
	GetPrompt() *string
	SetSearchSourcesShrink(v string) *RunQuickWritingShrinkRequest
	GetSearchSourcesShrink() *string
	SetTaskId(v string) *RunQuickWritingShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunQuickWritingShrinkRequest
	GetWorkspaceId() *string
}

type RunQuickWritingShrinkRequest struct {
	ArticlesShrink *string `json:"Articles,omitempty" xml:"Articles,omitempty"`
	// This parameter is required.
	Prompt              *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	SearchSourcesShrink *string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunQuickWritingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunQuickWritingShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunQuickWritingShrinkRequest) GetArticlesShrink() *string {
	return s.ArticlesShrink
}

func (s *RunQuickWritingShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunQuickWritingShrinkRequest) GetSearchSourcesShrink() *string {
	return s.SearchSourcesShrink
}

func (s *RunQuickWritingShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunQuickWritingShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunQuickWritingShrinkRequest) SetArticlesShrink(v string) *RunQuickWritingShrinkRequest {
	s.ArticlesShrink = &v
	return s
}

func (s *RunQuickWritingShrinkRequest) SetPrompt(v string) *RunQuickWritingShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunQuickWritingShrinkRequest) SetSearchSourcesShrink(v string) *RunQuickWritingShrinkRequest {
	s.SearchSourcesShrink = &v
	return s
}

func (s *RunQuickWritingShrinkRequest) SetTaskId(v string) *RunQuickWritingShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunQuickWritingShrinkRequest) SetWorkspaceId(v string) *RunQuickWritingShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunQuickWritingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
