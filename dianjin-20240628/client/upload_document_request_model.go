// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UploadDocumentRequest
	GetData() *string
	SetFileName(v string) *UploadDocumentRequest
	GetFileName() *string
	SetFileUrl(v string) *UploadDocumentRequest
	GetFileUrl() *string
	SetLibraryId(v string) *UploadDocumentRequest
	GetLibraryId() *string
}

type UploadDocumentRequest struct {
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://oss-xxx.hangzhou.com/test.pdf
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdhbcsj
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UploadDocumentRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentRequest) GetData() *string {
	return s.Data
}

func (s *UploadDocumentRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadDocumentRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadDocumentRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *UploadDocumentRequest) SetData(v string) *UploadDocumentRequest {
	s.Data = &v
	return s
}

func (s *UploadDocumentRequest) SetFileName(v string) *UploadDocumentRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentRequest) SetFileUrl(v string) *UploadDocumentRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadDocumentRequest) SetLibraryId(v string) *UploadDocumentRequest {
	s.LibraryId = &v
	return s
}

func (s *UploadDocumentRequest) Validate() error {
	return dara.Validate(s)
}
