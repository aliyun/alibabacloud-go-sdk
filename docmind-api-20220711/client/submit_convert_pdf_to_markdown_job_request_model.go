// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToMarkdownJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToMarkdownJobRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitConvertPdfToMarkdownJobRequest
	GetFileUrl() *string
	SetOssBucket(v string) *SubmitConvertPdfToMarkdownJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToMarkdownJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToMarkdownJobRequest struct {
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl     *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToMarkdownJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitConvertPdfToMarkdownJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToMarkdownJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToMarkdownJobRequest) SetFileName(v string) *SubmitConvertPdfToMarkdownJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobRequest) SetFileUrl(v string) *SubmitConvertPdfToMarkdownJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobRequest) SetOssBucket(v string) *SubmitConvertPdfToMarkdownJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobRequest) SetOssEndpoint(v string) *SubmitConvertPdfToMarkdownJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobRequest) Validate() error {
	return dara.Validate(s)
}
