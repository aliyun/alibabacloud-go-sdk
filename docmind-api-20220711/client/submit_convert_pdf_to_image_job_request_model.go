// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToImageJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToImageJobRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitConvertPdfToImageJobRequest
	GetFileUrl() *string
	SetOssBucket(v string) *SubmitConvertPdfToImageJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToImageJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToImageJobRequest struct {
	// example:
	//
	// convertPdfToImage.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl     *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToImageJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToImageJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToImageJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitConvertPdfToImageJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToImageJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToImageJobRequest) SetFileName(v string) *SubmitConvertPdfToImageJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToImageJobRequest) SetFileUrl(v string) *SubmitConvertPdfToImageJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitConvertPdfToImageJobRequest) SetOssBucket(v string) *SubmitConvertPdfToImageJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToImageJobRequest) SetOssEndpoint(v string) *SubmitConvertPdfToImageJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToImageJobRequest) Validate() error {
	return dara.Validate(s)
}
