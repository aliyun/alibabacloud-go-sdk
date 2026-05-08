// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUploadTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetOssUploadTokenRequest
	GetFileName() *string
	SetFileType(v string) *GetOssUploadTokenRequest
	GetFileType() *string
	SetUploadType(v int32) *GetOssUploadTokenRequest
	GetUploadType() *int32
}

type GetOssUploadTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 8021678.png
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ProductImage
	FileType   *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	UploadType *int32  `json:"uploadType,omitempty" xml:"uploadType,omitempty"`
}

func (s GetOssUploadTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOssUploadTokenRequest) GoString() string {
	return s.String()
}

func (s *GetOssUploadTokenRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetOssUploadTokenRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetOssUploadTokenRequest) GetUploadType() *int32 {
	return s.UploadType
}

func (s *GetOssUploadTokenRequest) SetFileName(v string) *GetOssUploadTokenRequest {
	s.FileName = &v
	return s
}

func (s *GetOssUploadTokenRequest) SetFileType(v string) *GetOssUploadTokenRequest {
	s.FileType = &v
	return s
}

func (s *GetOssUploadTokenRequest) SetUploadType(v int32) *GetOssUploadTokenRequest {
	s.UploadType = &v
	return s
}

func (s *GetOssUploadTokenRequest) Validate() error {
	return dara.Validate(s)
}
