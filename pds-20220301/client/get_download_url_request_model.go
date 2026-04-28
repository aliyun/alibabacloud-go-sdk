// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDriveId(v string) *GetDownloadUrlRequest
	GetDriveId() *string
	SetExpireSec(v int32) *GetDownloadUrlRequest
	GetExpireSec() *int32
	SetFileId(v string) *GetDownloadUrlRequest
	GetFileId() *string
	SetFileName(v string) *GetDownloadUrlRequest
	GetFileName() *string
	SetResponseContentType(v string) *GetDownloadUrlRequest
	GetResponseContentType() *string
	SetShareId(v string) *GetDownloadUrlRequest
	GetShareId() *string
}

type GetDownloadUrlRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The validity period of the download URL. Maximum value: 115200. Default value: 900. Unit: seconds.
	//
	// example:
	//
	// 100
	ExpireSec *int32 `json:"expire_sec,omitempty" xml:"expire_sec,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The name of the file. The name can be up to 1,024 characters in length.
	//
	// example:
	//
	// 1.txt
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// example:
	//
	// video/mp4
	ResponseContentType *string `json:"response_content_type,omitempty" xml:"response_content_type,omitempty"`
	// The share ID. If you want to manage a file by using a sharing link, carry the `x-share-token` header in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify at least either `share_id` or `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *GetDownloadUrlRequest) GetExpireSec() *int32 {
	return s.ExpireSec
}

func (s *GetDownloadUrlRequest) GetFileId() *string {
	return s.FileId
}

func (s *GetDownloadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetDownloadUrlRequest) GetResponseContentType() *string {
	return s.ResponseContentType
}

func (s *GetDownloadUrlRequest) GetShareId() *string {
	return s.ShareId
}

func (s *GetDownloadUrlRequest) SetDriveId(v string) *GetDownloadUrlRequest {
	s.DriveId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetExpireSec(v int32) *GetDownloadUrlRequest {
	s.ExpireSec = &v
	return s
}

func (s *GetDownloadUrlRequest) SetFileId(v string) *GetDownloadUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetFileName(v string) *GetDownloadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetDownloadUrlRequest) SetResponseContentType(v string) *GetDownloadUrlRequest {
	s.ResponseContentType = &v
	return s
}

func (s *GetDownloadUrlRequest) SetShareId(v string) *GetDownloadUrlRequest {
	s.ShareId = &v
	return s
}

func (s *GetDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
