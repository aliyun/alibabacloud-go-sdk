// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *ListTagsRequest
	GetDriveId() *string
	SetImageThumbnailProcess(v string) *ListTagsRequest
	GetImageThumbnailProcess() *string
	SetVideoThumbnailProcess(v string) *ListTagsRequest
	GetVideoThumbnailProcess() *string
}

type ListTagsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The method that is used to generate the thumbnail of an image.
	//
	// example:
	//
	// image/resize,w_200
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The method that is used to generate the thumbnail of a video.
	//
	// example:
	//
	// video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *ListTagsRequest) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *ListTagsRequest) GetVideoThumbnailProcess() *string {
	return s.VideoThumbnailProcess
}

func (s *ListTagsRequest) SetDriveId(v string) *ListTagsRequest {
	s.DriveId = &v
	return s
}

func (s *ListTagsRequest) SetImageThumbnailProcess(v string) *ListTagsRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *ListTagsRequest) SetVideoThumbnailProcess(v string) *ListTagsRequest {
	s.VideoThumbnailProcess = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}
