// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnlysisPrompt(v string) *AsyncUploadVideoRequest
	GetAnlysisPrompt() *string
	SetSourceVideos(v []*AsyncUploadVideoRequestSourceVideos) *AsyncUploadVideoRequest
	GetSourceVideos() []*AsyncUploadVideoRequestSourceVideos
	SetWorkspaceId(v string) *AsyncUploadVideoRequest
	GetWorkspaceId() *string
}

type AsyncUploadVideoRequest struct {
	AnlysisPrompt *string `json:"AnlysisPrompt,omitempty" xml:"AnlysisPrompt,omitempty"`
	// This parameter is required.
	SourceVideos []*AsyncUploadVideoRequestSourceVideos `json:"SourceVideos,omitempty" xml:"SourceVideos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncUploadVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequest) GetAnlysisPrompt() *string {
	return s.AnlysisPrompt
}

func (s *AsyncUploadVideoRequest) GetSourceVideos() []*AsyncUploadVideoRequestSourceVideos {
	return s.SourceVideos
}

func (s *AsyncUploadVideoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadVideoRequest) SetAnlysisPrompt(v string) *AsyncUploadVideoRequest {
	s.AnlysisPrompt = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetSourceVideos(v []*AsyncUploadVideoRequestSourceVideos) *AsyncUploadVideoRequest {
	s.SourceVideos = v
	return s
}

func (s *AsyncUploadVideoRequest) SetWorkspaceId(v string) *AsyncUploadVideoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadVideoRequest) Validate() error {
	return dara.Validate(s)
}

type AsyncUploadVideoRequestSourceVideos struct {
	VideoExtraInfo *string `json:"VideoExtraInfo,omitempty" xml:"VideoExtraInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123.mp4
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	// This parameter is required.
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AsyncUploadVideoRequestSourceVideos) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequestSourceVideos) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequestSourceVideos) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *AsyncUploadVideoRequestSourceVideos) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncUploadVideoRequestSourceVideos) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AsyncUploadVideoRequestSourceVideos) SetVideoExtraInfo(v string) *AsyncUploadVideoRequestSourceVideos {
	s.VideoExtraInfo = &v
	return s
}

func (s *AsyncUploadVideoRequestSourceVideos) SetVideoName(v string) *AsyncUploadVideoRequestSourceVideos {
	s.VideoName = &v
	return s
}

func (s *AsyncUploadVideoRequestSourceVideos) SetVideoUrl(v string) *AsyncUploadVideoRequestSourceVideos {
	s.VideoUrl = &v
	return s
}

func (s *AsyncUploadVideoRequestSourceVideos) Validate() error {
	return dara.Validate(s)
}
