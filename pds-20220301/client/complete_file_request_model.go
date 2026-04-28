// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrc64Hash(v string) *CompleteFileRequest
	GetCrc64Hash() *string
	SetDriveId(v string) *CompleteFileRequest
	GetDriveId() *string
	SetFileId(v string) *CompleteFileRequest
	GetFileId() *string
	SetUploadId(v string) *CompleteFileRequest
	GetUploadId() *string
}

type CompleteFileRequest struct {
	Crc64Hash *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
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
	// The upload ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C9DCFE5A82644AC7A02DB74C30C934A6
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s CompleteFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CompleteFileRequest) GoString() string {
	return s.String()
}

func (s *CompleteFileRequest) GetCrc64Hash() *string {
	return s.Crc64Hash
}

func (s *CompleteFileRequest) GetDriveId() *string {
	return s.DriveId
}

func (s *CompleteFileRequest) GetFileId() *string {
	return s.FileId
}

func (s *CompleteFileRequest) GetUploadId() *string {
	return s.UploadId
}

func (s *CompleteFileRequest) SetCrc64Hash(v string) *CompleteFileRequest {
	s.Crc64Hash = &v
	return s
}

func (s *CompleteFileRequest) SetDriveId(v string) *CompleteFileRequest {
	s.DriveId = &v
	return s
}

func (s *CompleteFileRequest) SetFileId(v string) *CompleteFileRequest {
	s.FileId = &v
	return s
}

func (s *CompleteFileRequest) SetUploadId(v string) *CompleteFileRequest {
	s.UploadId = &v
	return s
}

func (s *CompleteFileRequest) Validate() error {
	return dara.Validate(s)
}
