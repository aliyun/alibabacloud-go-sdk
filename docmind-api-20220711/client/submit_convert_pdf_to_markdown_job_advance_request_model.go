// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitConvertPdfToMarkdownJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitConvertPdfToMarkdownJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOssBucket(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToMarkdownJobAdvanceRequest struct {
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OssBucket     *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint   *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToMarkdownJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToMarkdownJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToMarkdownJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) SetOssBucket(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) SetOssEndpoint(v string) *SubmitConvertPdfToMarkdownJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToMarkdownJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
