// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileDownloadCallbackInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *FileDownloadCallbackInfo
	GetBucket() *string
	SetDomainId(v string) *FileDownloadCallbackInfo
	GetDomainId() *string
	SetDriveId(v string) *FileDownloadCallbackInfo
	GetDriveId() *string
	SetFileId(v string) *FileDownloadCallbackInfo
	GetFileId() *string
	SetObject(v string) *FileDownloadCallbackInfo
	GetObject() *string
	SetUserId(v string) *FileDownloadCallbackInfo
	GetUserId() *string
}

type FileDownloadCallbackInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// ccp-bj1-bj-1234
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// This parameter is required.
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// This parameter is required.
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// This parameter is required.
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// This parameter is required.
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FileDownloadCallbackInfo) String() string {
	return dara.Prettify(s)
}

func (s FileDownloadCallbackInfo) GoString() string {
	return s.String()
}

func (s *FileDownloadCallbackInfo) GetBucket() *string {
	return s.Bucket
}

func (s *FileDownloadCallbackInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *FileDownloadCallbackInfo) GetDriveId() *string {
	return s.DriveId
}

func (s *FileDownloadCallbackInfo) GetFileId() *string {
	return s.FileId
}

func (s *FileDownloadCallbackInfo) GetObject() *string {
	return s.Object
}

func (s *FileDownloadCallbackInfo) GetUserId() *string {
	return s.UserId
}

func (s *FileDownloadCallbackInfo) SetBucket(v string) *FileDownloadCallbackInfo {
	s.Bucket = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetDomainId(v string) *FileDownloadCallbackInfo {
	s.DomainId = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetDriveId(v string) *FileDownloadCallbackInfo {
	s.DriveId = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetFileId(v string) *FileDownloadCallbackInfo {
	s.FileId = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetObject(v string) *FileDownloadCallbackInfo {
	s.Object = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetUserId(v string) *FileDownloadCallbackInfo {
	s.UserId = &v
	return s
}

func (s *FileDownloadCallbackInfo) Validate() error {
	return dara.Validate(s)
}
