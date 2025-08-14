// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTableUnderstandingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitTableUnderstandingJobRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitTableUnderstandingJobRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitTableUnderstandingJobRequest
	GetFileUrl() *string
	SetOssBucket(v string) *SubmitTableUnderstandingJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitTableUnderstandingJobRequest
	GetOssEndpoint() *string
}

type SubmitTableUnderstandingJobRequest struct {
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

func (s SubmitTableUnderstandingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTableUnderstandingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitTableUnderstandingJobRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitTableUnderstandingJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitTableUnderstandingJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitTableUnderstandingJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitTableUnderstandingJobRequest) SetFileName(v string) *SubmitTableUnderstandingJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) SetFileNameExtension(v string) *SubmitTableUnderstandingJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) SetFileUrl(v string) *SubmitTableUnderstandingJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) SetOssBucket(v string) *SubmitTableUnderstandingJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) SetOssEndpoint(v string) *SubmitTableUnderstandingJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitTableUnderstandingJobRequest) Validate() error {
	return dara.Validate(s)
}
