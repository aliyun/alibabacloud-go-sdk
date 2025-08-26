// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskSubmitShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCounts(v int32) *PodcastTaskSubmitShrinkRequest
	GetCounts() *int32
	SetFileUrlsShrink(v string) *PodcastTaskSubmitShrinkRequest
	GetFileUrlsShrink() *string
	SetSourceLang(v string) *PodcastTaskSubmitShrinkRequest
	GetSourceLang() *string
	SetText(v string) *PodcastTaskSubmitShrinkRequest
	GetText() *string
	SetTopic(v string) *PodcastTaskSubmitShrinkRequest
	GetTopic() *string
	SetVoicesShrink(v string) *PodcastTaskSubmitShrinkRequest
	GetVoicesShrink() *string
	SetWorkspaceId(v string) *PodcastTaskSubmitShrinkRequest
	GetWorkspaceId() *string
}

type PodcastTaskSubmitShrinkRequest struct {
	// example:
	//
	// 2
	Counts         *int32  `json:"counts,omitempty" xml:"counts,omitempty"`
	FileUrlsShrink *string `json:"fileUrls,omitempty" xml:"fileUrls,omitempty"`
	SourceLang     *string `json:"sourceLang,omitempty" xml:"sourceLang,omitempty"`
	Text           *string `json:"text,omitempty" xml:"text,omitempty"`
	Topic          *string `json:"topic,omitempty" xml:"topic,omitempty"`
	VoicesShrink   *string `json:"voices,omitempty" xml:"voices,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PodcastTaskSubmitShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskSubmitShrinkRequest) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitShrinkRequest) GetCounts() *int32 {
	return s.Counts
}

func (s *PodcastTaskSubmitShrinkRequest) GetFileUrlsShrink() *string {
	return s.FileUrlsShrink
}

func (s *PodcastTaskSubmitShrinkRequest) GetSourceLang() *string {
	return s.SourceLang
}

func (s *PodcastTaskSubmitShrinkRequest) GetText() *string {
	return s.Text
}

func (s *PodcastTaskSubmitShrinkRequest) GetTopic() *string {
	return s.Topic
}

func (s *PodcastTaskSubmitShrinkRequest) GetVoicesShrink() *string {
	return s.VoicesShrink
}

func (s *PodcastTaskSubmitShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PodcastTaskSubmitShrinkRequest) SetCounts(v int32) *PodcastTaskSubmitShrinkRequest {
	s.Counts = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetFileUrlsShrink(v string) *PodcastTaskSubmitShrinkRequest {
	s.FileUrlsShrink = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetSourceLang(v string) *PodcastTaskSubmitShrinkRequest {
	s.SourceLang = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetText(v string) *PodcastTaskSubmitShrinkRequest {
	s.Text = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetTopic(v string) *PodcastTaskSubmitShrinkRequest {
	s.Topic = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetVoicesShrink(v string) *PodcastTaskSubmitShrinkRequest {
	s.VoicesShrink = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) SetWorkspaceId(v string) *PodcastTaskSubmitShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PodcastTaskSubmitShrinkRequest) Validate() error {
	return dara.Validate(s)
}
