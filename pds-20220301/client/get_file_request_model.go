// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *GetFileRequest
	GetDriveId() *string
	SetFields(v string) *GetFileRequest
	GetFields() *string
	SetFileId(v string) *GetFileRequest
	GetFileId() *string
	SetShareId(v string) *GetFileRequest
	GetShareId() *string
	SetThumbnailProcesses(v map[string]*ImageProcess) *GetFileRequest
	GetThumbnailProcesses() map[string]*ImageProcess
	SetUrlExpireSec(v int32) *GetFileRequest
	GetUrlExpireSec() *int32
}

type GetFileRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The fields to return.
	//
	// 1.  If this parameter is set to \\*, all fields of the file except the fields that must be specified are returned.
	//
	// 2.  If only specific fields are required, you can specify the following fields: url, thumbnail, exif, cropping_suggestion, characteristic_hash, video_metadata, and video_preview_metadata. If multiple fields are required, separate them with commas (,). Example: url,thumbnail.
	//
	// 3.  The investigation_info field is returned only if it is specified.
	//
	// By default, all fields except the fields that must be specified are returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID. If you want to share a file, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The thumbnail configurations. Up to five thumbnails can be returned at a time. The value contains key-value pairs. You can customize the keys. The URL of a thumbnail is returned based on the key.
	ThumbnailProcesses map[string]*ImageProcess `json:"thumbnail_processes,omitempty" xml:"thumbnail_processes,omitempty"`
	// The time when the file expires. Unit: seconds. Valid values: 10 to 14400.
	//
	// example:
	//
	// 100
	UrlExpireSec *int32 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s GetFileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetFileRequest) GetFields() *string {
	return s.Fields
}

func (s *GetFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetFileRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetFileRequest) GetThumbnailProcesses() map[string]*ImageProcess {
	return s.ThumbnailProcesses
}

func (s *GetFileRequest) GetUrlExpireSec() *int32 {
	return s.UrlExpireSec
}

func (s *GetFileRequest) SetDriveId(v string) *GetFileRequest {
	s.DriveId = &v
	return s
}

func (s *GetFileRequest) SetFields(v string) *GetFileRequest {
	s.Fields = &v
	return s
}

func (s *GetFileRequest) SetFileId(v string) *GetFileRequest {
	s.FileId = &v
	return s
}

func (s *GetFileRequest) SetShareId(v string) *GetFileRequest {
	s.ShareId = &v
	return s
}

func (s *GetFileRequest) SetThumbnailProcesses(v map[string]*ImageProcess) *GetFileRequest {
	s.ThumbnailProcesses = v
	return s
}

func (s *GetFileRequest) SetUrlExpireSec(v int32) *GetFileRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *GetFileRequest) Validate() error {
	return dara.Validate(s)
}
