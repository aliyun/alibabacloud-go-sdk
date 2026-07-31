// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSemanticFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *UploadSemanticFileRequest
	GetContentType() *string
	SetFileName(v string) *UploadSemanticFileRequest
	GetFileName() *string
	SetSizeBytes(v int64) *UploadSemanticFileRequest
	GetSizeBytes() *int64
}

type UploadSemanticFileRequest struct {
	// The MIME type of the object to upload. Maximum length: 128 characters. This value is included in the signature of UploadUrl. Use the same Content-Type when you perform the PUT request.
	//
	// This parameter is required.
	//
	// example:
	//
	// application/pdf
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The original file name of the reference file to upload. Maximum length: 255 characters. When FileId is used for singleTableFile, only CSV or XLSX files are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// reference.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The size of the file to upload, in bytes. This value is recorded as attachment metadata. Specify the actual file size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s UploadSemanticFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSemanticFileRequest) GoString() string {
	return s.String()
}

func (s *UploadSemanticFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *UploadSemanticFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadSemanticFileRequest) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *UploadSemanticFileRequest) SetContentType(v string) *UploadSemanticFileRequest {
	s.ContentType = &v
	return s
}

func (s *UploadSemanticFileRequest) SetFileName(v string) *UploadSemanticFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadSemanticFileRequest) SetSizeBytes(v int64) *UploadSemanticFileRequest {
	s.SizeBytes = &v
	return s
}

func (s *UploadSemanticFileRequest) Validate() error {
	return dara.Validate(s)
}
