// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadVideoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnlysisPrompt(v string) *AsyncUploadVideoShrinkRequest
	GetAnlysisPrompt() *string
	SetSourceVideosShrink(v string) *AsyncUploadVideoShrinkRequest
	GetSourceVideosShrink() *string
	SetWorkspaceId(v string) *AsyncUploadVideoShrinkRequest
	GetWorkspaceId() *string
}

type AsyncUploadVideoShrinkRequest struct {
	AnlysisPrompt *string `json:"AnlysisPrompt,omitempty" xml:"AnlysisPrompt,omitempty"`
	// This parameter is required.
	SourceVideosShrink *string `json:"SourceVideos,omitempty" xml:"SourceVideos,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncUploadVideoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoShrinkRequest) GetAnlysisPrompt() *string {
	return s.AnlysisPrompt
}

func (s *AsyncUploadVideoShrinkRequest) GetSourceVideosShrink() *string {
	return s.SourceVideosShrink
}

func (s *AsyncUploadVideoShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadVideoShrinkRequest) SetAnlysisPrompt(v string) *AsyncUploadVideoShrinkRequest {
	s.AnlysisPrompt = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetSourceVideosShrink(v string) *AsyncUploadVideoShrinkRequest {
	s.SourceVideosShrink = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetWorkspaceId(v string) *AsyncUploadVideoShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
