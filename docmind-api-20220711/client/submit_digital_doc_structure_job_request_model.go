// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDigitalDocStructureJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDigitalDocStructureJobRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDigitalDocStructureJobRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDigitalDocStructureJobRequest
	GetFileUrl() *string
	SetImageStrategy(v string) *SubmitDigitalDocStructureJobRequest
	GetImageStrategy() *string
	SetOssBucket(v string) *SubmitDigitalDocStructureJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDigitalDocStructureJobRequest
	GetOssEndpoint() *string
	SetRevealMarkdown(v bool) *SubmitDigitalDocStructureJobRequest
	GetRevealMarkdown() *bool
	SetUseUrlResponseBody(v bool) *SubmitDigitalDocStructureJobRequest
	GetUseUrlResponseBody() *bool
}

type SubmitDigitalDocStructureJobRequest struct {
	// example:
	//
	// docStructure.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// pdf
	FileNameExtension *string `json:"FileNameExtension,omitempty" xml:"FileNameExtension,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl            *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImageStrategy      *string `json:"ImageStrategy,omitempty" xml:"ImageStrategy,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	RevealMarkdown     *bool   `json:"RevealMarkdown,omitempty" xml:"RevealMarkdown,omitempty"`
	UseUrlResponseBody *bool   `json:"UseUrlResponseBody,omitempty" xml:"UseUrlResponseBody,omitempty"`
}

func (s SubmitDigitalDocStructureJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalDocStructureJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDigitalDocStructureJobRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDigitalDocStructureJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDigitalDocStructureJobRequest) GetImageStrategy() *string {
	return s.ImageStrategy
}

func (s *SubmitDigitalDocStructureJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDigitalDocStructureJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDigitalDocStructureJobRequest) GetRevealMarkdown() *bool {
	return s.RevealMarkdown
}

func (s *SubmitDigitalDocStructureJobRequest) GetUseUrlResponseBody() *bool {
	return s.UseUrlResponseBody
}

func (s *SubmitDigitalDocStructureJobRequest) SetFileName(v string) *SubmitDigitalDocStructureJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetFileNameExtension(v string) *SubmitDigitalDocStructureJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetFileUrl(v string) *SubmitDigitalDocStructureJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetImageStrategy(v string) *SubmitDigitalDocStructureJobRequest {
	s.ImageStrategy = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetOssBucket(v string) *SubmitDigitalDocStructureJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetOssEndpoint(v string) *SubmitDigitalDocStructureJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetRevealMarkdown(v bool) *SubmitDigitalDocStructureJobRequest {
	s.RevealMarkdown = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) SetUseUrlResponseBody(v bool) *SubmitDigitalDocStructureJobRequest {
	s.UseUrlResponseBody = &v
	return s
}

func (s *SubmitDigitalDocStructureJobRequest) Validate() error {
	return dara.Validate(s)
}
