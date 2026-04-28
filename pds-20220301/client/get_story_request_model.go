// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoverImageThumbnailProcess(v string) *GetStoryRequest
	GetCoverImageThumbnailProcess() *string
	SetCoverVideoThumbnailProcess(v string) *GetStoryRequest
	GetCoverVideoThumbnailProcess() *string
	SetDriveId(v string) *GetStoryRequest
	GetDriveId() *string
	SetImageThumbnailProcess(v string) *GetStoryRequest
	GetImageThumbnailProcess() *string
	SetImageUrlProcess(v string) *GetStoryRequest
	GetImageUrlProcess() *string
	SetStoryId(v string) *GetStoryRequest
	GetStoryId() *string
	SetUrlExpireSec(v int64) *GetStoryRequest
	GetUrlExpireSec() *int64
	SetVideoThumbnailProcess(v string) *GetStoryRequest
	GetVideoThumbnailProcess() *string
}

type GetStoryRequest struct {
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	CoverImageThumbnailProcess *string `json:"cover_image_thumbnail_process,omitempty" xml:"cover_image_thumbnail_process,omitempty"`
	// Deprecated
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	CoverVideoThumbnailProcess *string `json:"cover_video_thumbnail_process,omitempty" xml:"cover_video_thumbnail_process,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	ImageUrlProcess *string `json:"image_url_process,omitempty" xml:"image_url_process,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 900
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
	// Deprecated
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s GetStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStoryRequest) GoString() string {
	return s.String()
}

func (s *GetStoryRequest) GetCoverImageThumbnailProcess() *string {
	return s.CoverImageThumbnailProcess
}

func (s *GetStoryRequest) GetCoverVideoThumbnailProcess() *string {
	return s.CoverVideoThumbnailProcess
}

func (s *GetStoryRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetStoryRequest) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *GetStoryRequest) GetImageUrlProcess() *string {
	return s.ImageUrlProcess
}

func (s *GetStoryRequest) GetStoryId() *string {
	return s.StoryId
}

func (s *GetStoryRequest) GetUrlExpireSec() *int64 {
	return s.UrlExpireSec
}

func (s *GetStoryRequest) GetVideoThumbnailProcess() *string {
	return s.VideoThumbnailProcess
}

func (s *GetStoryRequest) SetCoverImageThumbnailProcess(v string) *GetStoryRequest {
	s.CoverImageThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) SetCoverVideoThumbnailProcess(v string) *GetStoryRequest {
	s.CoverVideoThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) SetDriveId(v string) *GetStoryRequest {
	s.DriveId = &v
	return s
}

func (s *GetStoryRequest) SetImageThumbnailProcess(v string) *GetStoryRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) SetImageUrlProcess(v string) *GetStoryRequest {
	s.ImageUrlProcess = &v
	return s
}

func (s *GetStoryRequest) SetStoryId(v string) *GetStoryRequest {
	s.StoryId = &v
	return s
}

func (s *GetStoryRequest) SetUrlExpireSec(v int64) *GetStoryRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *GetStoryRequest) SetVideoThumbnailProcess(v string) *GetStoryRequest {
	s.VideoThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) Validate() error {
	return dara.Validate(s)
}
