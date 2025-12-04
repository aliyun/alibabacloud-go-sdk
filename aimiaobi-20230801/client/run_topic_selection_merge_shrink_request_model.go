// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTopicSelectionMergeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunTopicSelectionMergeShrinkRequest
	GetPrompt() *string
	SetTopicsShrink(v string) *RunTopicSelectionMergeShrinkRequest
	GetTopicsShrink() *string
	SetWorkspaceId(v string) *RunTopicSelectionMergeShrinkRequest
	GetWorkspaceId() *string
}

type RunTopicSelectionMergeShrinkRequest struct {
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	TopicsShrink *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTopicSelectionMergeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTopicSelectionMergeShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTopicSelectionMergeShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunTopicSelectionMergeShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *RunTopicSelectionMergeShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTopicSelectionMergeShrinkRequest) SetPrompt(v string) *RunTopicSelectionMergeShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunTopicSelectionMergeShrinkRequest) SetTopicsShrink(v string) *RunTopicSelectionMergeShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *RunTopicSelectionMergeShrinkRequest) SetWorkspaceId(v string) *RunTopicSelectionMergeShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTopicSelectionMergeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
