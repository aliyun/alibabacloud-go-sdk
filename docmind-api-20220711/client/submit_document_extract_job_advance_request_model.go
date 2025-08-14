// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitDocumentExtractJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDocumentExtractJobAdvanceRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocumentExtractJobAdvanceRequest
	GetFileNameExtension() *string
	SetFileUrlObject(v io.Reader) *SubmitDocumentExtractJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOssBucket(v string) *SubmitDocumentExtractJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocumentExtractJobAdvanceRequest
	GetOssEndpoint() *string
}

type SubmitDocumentExtractJobAdvanceRequest struct {
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	OssBucket     *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint   *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitDocumentExtractJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocumentExtractJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocumentExtractJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocumentExtractJobAdvanceRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocumentExtractJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitDocumentExtractJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocumentExtractJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetFileName(v string) *SubmitDocumentExtractJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDocumentExtractJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocumentExtractJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetOssBucket(v string) *SubmitDocumentExtractJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) SetOssEndpoint(v string) *SubmitDocumentExtractJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocumentExtractJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
