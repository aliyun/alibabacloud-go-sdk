// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocumentExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDocumentExtractJobRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocumentExtractJobRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDocumentExtractJobRequest
	GetFileUrl() *string
	SetOssBucket(v string) *SubmitDocumentExtractJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocumentExtractJobRequest
	GetOssEndpoint() *string
}

type SubmitDocumentExtractJobRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl     *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitDocumentExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentExtractJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocumentExtractJobRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocumentExtractJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocumentExtractJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocumentExtractJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocumentExtractJobRequest) SetFileName(v string) *SubmitDocumentExtractJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) SetFileNameExtension(v string) *SubmitDocumentExtractJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) SetFileUrl(v string) *SubmitDocumentExtractJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) SetOssBucket(v string) *SubmitDocumentExtractJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) SetOssEndpoint(v string) *SubmitDocumentExtractJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocumentExtractJobRequest) Validate() error {
	return dara.Validate(s)
}
