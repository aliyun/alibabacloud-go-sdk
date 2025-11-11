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
	SetReferenceVideo(v *AsyncUploadVideoRequestReferenceVideo) *AsyncUploadVideoRequest
	GetReferenceVideo() *AsyncUploadVideoRequestReferenceVideo
	SetSourceVideos(v []*AsyncUploadVideoRequestSourceVideos) *AsyncUploadVideoRequest
	GetSourceVideos() []*AsyncUploadVideoRequestSourceVideos
	SetSplitInterval(v int32) *AsyncUploadVideoRequest
	GetSplitInterval() *int32
	SetWorkspaceId(v string) *AsyncUploadVideoRequest
	GetWorkspaceId() *string
}

type AsyncUploadVideoRequest struct {
	AnlysisPrompt  *string                                `json:"AnlysisPrompt,omitempty" xml:"AnlysisPrompt,omitempty"`
	ReferenceVideo *AsyncUploadVideoRequestReferenceVideo `json:"ReferenceVideo,omitempty" xml:"ReferenceVideo,omitempty" type:"Struct"`
	// This parameter is required.
	SourceVideos  []*AsyncUploadVideoRequestSourceVideos `json:"SourceVideos,omitempty" xml:"SourceVideos,omitempty" type:"Repeated"`
	SplitInterval *int32                                 `json:"SplitInterval,omitempty" xml:"SplitInterval,omitempty"`
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

func (s *AsyncUploadVideoRequest) GetReferenceVideo() *AsyncUploadVideoRequestReferenceVideo {
	return s.ReferenceVideo
}

func (s *AsyncUploadVideoRequest) GetSourceVideos() []*AsyncUploadVideoRequestSourceVideos {
	return s.SourceVideos
}

func (s *AsyncUploadVideoRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *AsyncUploadVideoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadVideoRequest) SetAnlysisPrompt(v string) *AsyncUploadVideoRequest {
	s.AnlysisPrompt = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetReferenceVideo(v *AsyncUploadVideoRequestReferenceVideo) *AsyncUploadVideoRequest {
	s.ReferenceVideo = v
	return s
}

func (s *AsyncUploadVideoRequest) SetSourceVideos(v []*AsyncUploadVideoRequestSourceVideos) *AsyncUploadVideoRequest {
	s.SourceVideos = v
	return s
}

func (s *AsyncUploadVideoRequest) SetSplitInterval(v int32) *AsyncUploadVideoRequest {
	s.SplitInterval = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetWorkspaceId(v string) *AsyncUploadVideoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadVideoRequest) Validate() error {
	if s.ReferenceVideo != nil {
		if err := s.ReferenceVideo.Validate(); err != nil {
			return err
		}
	}
	if s.SourceVideos != nil {
		for _, item := range s.SourceVideos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AsyncUploadVideoRequestReferenceVideo struct {
	VideoExtraInfo *string `json:"VideoExtraInfo,omitempty" xml:"VideoExtraInfo,omitempty"`
	VideoName      *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	VideoUrl       *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AsyncUploadVideoRequestReferenceVideo) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequestReferenceVideo) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequestReferenceVideo) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *AsyncUploadVideoRequestReferenceVideo) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncUploadVideoRequestReferenceVideo) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AsyncUploadVideoRequestReferenceVideo) SetVideoExtraInfo(v string) *AsyncUploadVideoRequestReferenceVideo {
	s.VideoExtraInfo = &v
	return s
}

func (s *AsyncUploadVideoRequestReferenceVideo) SetVideoName(v string) *AsyncUploadVideoRequestReferenceVideo {
	s.VideoName = &v
	return s
}

func (s *AsyncUploadVideoRequestReferenceVideo) SetVideoUrl(v string) *AsyncUploadVideoRequestReferenceVideo {
	s.VideoUrl = &v
	return s
}

func (s *AsyncUploadVideoRequestReferenceVideo) Validate() error {
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
