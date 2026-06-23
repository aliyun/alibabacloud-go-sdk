// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadDocumentAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UploadDocumentAdvanceRequest
	GetData() *string
	SetFileName(v string) *UploadDocumentAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *UploadDocumentAdvanceRequest
	GetFileUrlObject() io.Reader
	SetLibraryId(v string) *UploadDocumentAdvanceRequest
	GetLibraryId() *string
}

type UploadDocumentAdvanceRequest struct {
	// File metadata. You can use this to filter results during retrieval.
	//
	// example:
	//
	// {\\"cateogry\\": \\"报告\\"}
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The full file name, including the extension. Supported formats are PDF, DOC, DOCX, Markdown, PPT, and PPTX. File size must not exceed 100 MB. PDF, DOC, DOCX, PPT, and PPTX files must not exceed 500 pages.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The OSS URL of the file. If the file is not publicly readable, include a signature in the URL.
	//
	// If you use the SDK to upload files, upload the file directly. You do not need to provide an OSS URL. For more information, see the SDK documentation.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://oss-xxx.hangzhou.com/test.pdf
	FileUrlObject io.Reader `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The document library ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sjdhbcsj
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
}

func (s UploadDocumentAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentAdvanceRequest) GetData() *string {
	return s.Data
}

func (s *UploadDocumentAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadDocumentAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *UploadDocumentAdvanceRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *UploadDocumentAdvanceRequest) SetData(v string) *UploadDocumentAdvanceRequest {
	s.Data = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetFileName(v string) *UploadDocumentAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetFileUrlObject(v io.Reader) *UploadDocumentAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UploadDocumentAdvanceRequest) SetLibraryId(v string) *UploadDocumentAdvanceRequest {
	s.LibraryId = &v
	return s
}

func (s *UploadDocumentAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
