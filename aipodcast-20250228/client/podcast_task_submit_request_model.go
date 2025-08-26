// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPodcastTaskSubmitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCounts(v int32) *PodcastTaskSubmitRequest
	GetCounts() *int32
	SetFileUrls(v []*string) *PodcastTaskSubmitRequest
	GetFileUrls() []*string
	SetSourceLang(v string) *PodcastTaskSubmitRequest
	GetSourceLang() *string
	SetText(v string) *PodcastTaskSubmitRequest
	GetText() *string
	SetTopic(v string) *PodcastTaskSubmitRequest
	GetTopic() *string
	SetVoices(v []*string) *PodcastTaskSubmitRequest
	GetVoices() []*string
	SetWorkspaceId(v string) *PodcastTaskSubmitRequest
	GetWorkspaceId() *string
}

type PodcastTaskSubmitRequest struct {
	// example:
	//
	// 2
	Counts     *int32    `json:"counts,omitempty" xml:"counts,omitempty"`
	FileUrls   []*string `json:"fileUrls,omitempty" xml:"fileUrls,omitempty" type:"Repeated"`
	SourceLang *string   `json:"sourceLang,omitempty" xml:"sourceLang,omitempty"`
	Text       *string   `json:"text,omitempty" xml:"text,omitempty"`
	Topic      *string   `json:"topic,omitempty" xml:"topic,omitempty"`
	Voices     []*string `json:"voices,omitempty" xml:"voices,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-ep8ba0dr6seiddxx
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PodcastTaskSubmitRequest) String() string {
	return dara.Prettify(s)
}

func (s PodcastTaskSubmitRequest) GoString() string {
	return s.String()
}

func (s *PodcastTaskSubmitRequest) GetCounts() *int32 {
	return s.Counts
}

func (s *PodcastTaskSubmitRequest) GetFileUrls() []*string {
	return s.FileUrls
}

func (s *PodcastTaskSubmitRequest) GetSourceLang() *string {
	return s.SourceLang
}

func (s *PodcastTaskSubmitRequest) GetText() *string {
	return s.Text
}

func (s *PodcastTaskSubmitRequest) GetTopic() *string {
	return s.Topic
}

func (s *PodcastTaskSubmitRequest) GetVoices() []*string {
	return s.Voices
}

func (s *PodcastTaskSubmitRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PodcastTaskSubmitRequest) SetCounts(v int32) *PodcastTaskSubmitRequest {
	s.Counts = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetFileUrls(v []*string) *PodcastTaskSubmitRequest {
	s.FileUrls = v
	return s
}

func (s *PodcastTaskSubmitRequest) SetSourceLang(v string) *PodcastTaskSubmitRequest {
	s.SourceLang = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetText(v string) *PodcastTaskSubmitRequest {
	s.Text = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetTopic(v string) *PodcastTaskSubmitRequest {
	s.Topic = &v
	return s
}

func (s *PodcastTaskSubmitRequest) SetVoices(v []*string) *PodcastTaskSubmitRequest {
	s.Voices = v
	return s
}

func (s *PodcastTaskSubmitRequest) SetWorkspaceId(v string) *PodcastTaskSubmitRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PodcastTaskSubmitRequest) Validate() error {
	return dara.Validate(s)
}
