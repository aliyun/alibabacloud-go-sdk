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
	SetReferenceVideoShrink(v string) *AsyncUploadVideoShrinkRequest
	GetReferenceVideoShrink() *string
	SetSourceVideosShrink(v string) *AsyncUploadVideoShrinkRequest
	GetSourceVideosShrink() *string
	SetSplitInterval(v int32) *AsyncUploadVideoShrinkRequest
	GetSplitInterval() *int32
	SetWorkspaceId(v string) *AsyncUploadVideoShrinkRequest
	GetWorkspaceId() *string
}

type AsyncUploadVideoShrinkRequest struct {
	AnlysisPrompt        *string `json:"AnlysisPrompt,omitempty" xml:"AnlysisPrompt,omitempty"`
	ReferenceVideoShrink *string `json:"ReferenceVideo,omitempty" xml:"ReferenceVideo,omitempty"`
	// This parameter is required.
	SourceVideosShrink *string `json:"SourceVideos,omitempty" xml:"SourceVideos,omitempty"`
	SplitInterval      *int32  `json:"SplitInterval,omitempty" xml:"SplitInterval,omitempty"`
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

func (s *AsyncUploadVideoShrinkRequest) GetReferenceVideoShrink() *string {
	return s.ReferenceVideoShrink
}

func (s *AsyncUploadVideoShrinkRequest) GetSourceVideosShrink() *string {
	return s.SourceVideosShrink
}

func (s *AsyncUploadVideoShrinkRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *AsyncUploadVideoShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadVideoShrinkRequest) SetAnlysisPrompt(v string) *AsyncUploadVideoShrinkRequest {
	s.AnlysisPrompt = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetReferenceVideoShrink(v string) *AsyncUploadVideoShrinkRequest {
	s.ReferenceVideoShrink = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetSourceVideosShrink(v string) *AsyncUploadVideoShrinkRequest {
	s.SourceVideosShrink = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetSplitInterval(v int32) *AsyncUploadVideoShrinkRequest {
	s.SplitInterval = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetWorkspaceId(v string) *AsyncUploadVideoShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
