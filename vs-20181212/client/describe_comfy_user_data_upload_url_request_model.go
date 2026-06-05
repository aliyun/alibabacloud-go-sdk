// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComfyUserDataUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *DescribeComfyUserDataUploadUrlRequest
	GetContentType() *string
	SetFileMd5(v string) *DescribeComfyUserDataUploadUrlRequest
	GetFileMd5() *string
	SetFileName(v string) *DescribeComfyUserDataUploadUrlRequest
	GetFileName() *string
	SetFileSizeBytes(v int64) *DescribeComfyUserDataUploadUrlRequest
	GetFileSizeBytes() *int64
}

type DescribeComfyUserDataUploadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// application/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 54d6911ba6d59dbe68990835a409f18c
	FileMd5 *string `json:"FileMd5,omitempty" xml:"FileMd5,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// myfile
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1024
	FileSizeBytes *int64 `json:"FileSizeBytes,omitempty" xml:"FileSizeBytes,omitempty"`
}

func (s DescribeComfyUserDataUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeComfyUserDataUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *DescribeComfyUserDataUploadUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *DescribeComfyUserDataUploadUrlRequest) GetFileMd5() *string {
	return s.FileMd5
}

func (s *DescribeComfyUserDataUploadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *DescribeComfyUserDataUploadUrlRequest) GetFileSizeBytes() *int64 {
	return s.FileSizeBytes
}

func (s *DescribeComfyUserDataUploadUrlRequest) SetContentType(v string) *DescribeComfyUserDataUploadUrlRequest {
	s.ContentType = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlRequest) SetFileMd5(v string) *DescribeComfyUserDataUploadUrlRequest {
	s.FileMd5 = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlRequest) SetFileName(v string) *DescribeComfyUserDataUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlRequest) SetFileSizeBytes(v int64) *DescribeComfyUserDataUploadUrlRequest {
	s.FileSizeBytes = &v
	return s
}

func (s *DescribeComfyUserDataUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
