// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDownloadFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *DownloadFileRequest
	GetDriveId() *string
	SetFileId(v string) *DownloadFileRequest
	GetFileId() *string
	SetImageThumbnailProcess(v string) *DownloadFileRequest
	GetImageThumbnailProcess() *string
	SetOfficeThumbnailProcess(v string) *DownloadFileRequest
	GetOfficeThumbnailProcess() *string
	SetShareId(v string) *DownloadFileRequest
	GetShareId() *string
	SetVideoThumbnailProcess(v string) *DownloadFileRequest
	GetVideoThumbnailProcess() *string
}

type DownloadFileRequest struct {
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
	// The method used to generate the thumbnail of an image. If this parameter is specified, you are redirected to the URL of the generated thumbnail.
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The method used to generate the thumbnail of a document. If this parameter is specified, you are redirected to the URL of the generated thumbnail.
	//
	// example:
	//
	// image/resize,w_200
	OfficeThumbnailProcess *string `json:"office_thumbnail_process,omitempty" xml:"office_thumbnail_process,omitempty"`
	// The share ID. If you want to manage a file by using a share link, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The method used to generate the thumbnail of a video. If this parameter is specified, you are redirected to the URL of the generated thumbnail.
	//
	// example:
	//
	// video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s DownloadFileRequest) String() string {
	return dara.Prettify(s)
}

func (s DownloadFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *DownloadFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *DownloadFileRequest) GetImageThumbnailProcess() *string {
	return s.ImageThumbnailProcess
}

func (s *DownloadFileRequest) GetOfficeThumbnailProcess() *string {
	return s.OfficeThumbnailProcess
}

func (s *DownloadFileRequest) GetShareId() *string {
	return s.ShareId
}

func (s *DownloadFileRequest) GetVideoThumbnailProcess() *string {
	return s.VideoThumbnailProcess
}

func (s *DownloadFileRequest) SetDriveId(v string) *DownloadFileRequest {
	s.DriveId = &v
	return s
}

func (s *DownloadFileRequest) SetFileId(v string) *DownloadFileRequest {
	s.FileId = &v
	return s
}

func (s *DownloadFileRequest) SetImageThumbnailProcess(v string) *DownloadFileRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *DownloadFileRequest) SetOfficeThumbnailProcess(v string) *DownloadFileRequest {
	s.OfficeThumbnailProcess = &v
	return s
}

func (s *DownloadFileRequest) SetShareId(v string) *DownloadFileRequest {
	s.ShareId = &v
	return s
}

func (s *DownloadFileRequest) SetVideoThumbnailProcess(v string) *DownloadFileRequest {
	s.VideoThumbnailProcess = &v
	return s
}

func (s *DownloadFileRequest) Validate() error {
	return dara.Validate(s)
}
