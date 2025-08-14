// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitConvertPdfToImageJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToImageJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitConvertPdfToImageJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOssBucket(v string) *SubmitConvertPdfToImageJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToImageJobAdvanceRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToImageJobAdvanceRequest struct {
	// example:
	//
	// convertPdfToImage.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OssBucket     *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint   *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToImageJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToImageJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToImageJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToImageJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) SetOssBucket(v string) *SubmitConvertPdfToImageJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) SetOssEndpoint(v string) *SubmitConvertPdfToImageJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToImageJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
