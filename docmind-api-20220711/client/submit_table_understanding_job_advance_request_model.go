// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitTableUnderstandingJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitTableUnderstandingJobAdvanceRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitTableUnderstandingJobAdvanceRequest
	GetFileNameExtension() *string
	SetFileUrlObject(v io.Reader) *SubmitTableUnderstandingJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetOssBucket(v string) *SubmitTableUnderstandingJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitTableUnderstandingJobAdvanceRequest
	GetOssEndpoint() *string
}

type SubmitTableUnderstandingJobAdvanceRequest struct {
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

func (s SubmitTableUnderstandingJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitTableUnderstandingJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetFileName(v string) *SubmitTableUnderstandingJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetFileNameExtension(v string) *SubmitTableUnderstandingJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitTableUnderstandingJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetOssBucket(v string) *SubmitTableUnderstandingJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) SetOssEndpoint(v string) *SubmitTableUnderstandingJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitTableUnderstandingJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
