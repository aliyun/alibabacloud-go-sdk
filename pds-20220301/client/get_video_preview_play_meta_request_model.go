// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoPreviewPlayMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetVideoPreviewPlayMetaRequest
	GetCategory() *string
	SetDriveId(v string) *GetVideoPreviewPlayMetaRequest
	GetDriveId() *string
	SetFileId(v string) *GetVideoPreviewPlayMetaRequest
	GetFileId() *string
	SetShareId(v string) *GetVideoPreviewPlayMetaRequest
	GetShareId() *string
}

type GetVideoPreviewPlayMetaRequest struct {
	// The preview type. You must enable the corresponding video transcoding feature. Valid values:
	//
	// 	- live_transcoding: previews a live stream while transcoding is in progress.
	//
	// 	- quick_video: previews a video while transcoding is in progress.
	//
	// 	- offline_audio: previews a piece of audio after the audio is transcoded offline.
	//
	// 	- offline_video: previews a video after the video is transcoded offline.
	//
	// This parameter is required.
	//
	// example:
	//
	// live_transcoding
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID. If you want to manage a file by using a sharing link, carry the `x-share-token` header in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify at least either `share_id` or `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetVideoPreviewPlayMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoPreviewPlayMetaRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayMetaRequest) GetCategory() *string {
	return s.Category
}

func (s *GetVideoPreviewPlayMetaRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetVideoPreviewPlayMetaRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetVideoPreviewPlayMetaRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetVideoPreviewPlayMetaRequest) SetCategory(v string) *GetVideoPreviewPlayMetaRequest {
	s.Category = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) SetDriveId(v string) *GetVideoPreviewPlayMetaRequest {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) SetFileId(v string) *GetVideoPreviewPlayMetaRequest {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) SetShareId(v string) *GetVideoPreviewPlayMetaRequest {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) Validate() error {
	return dara.Validate(s)
}
