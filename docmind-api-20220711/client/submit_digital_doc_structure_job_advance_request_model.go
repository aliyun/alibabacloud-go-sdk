// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitDigitalDocStructureJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitDigitalDocStructureJobAdvanceRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDigitalDocStructureJobAdvanceRequest
	GetFileNameExtension() *string
	SetFileUrlObject(v io.Reader) *SubmitDigitalDocStructureJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetImageStrategy(v string) *SubmitDigitalDocStructureJobAdvanceRequest
	GetImageStrategy() *string
	SetOssBucket(v string) *SubmitDigitalDocStructureJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDigitalDocStructureJobAdvanceRequest
	GetOssEndpoint() *string
	SetRevealMarkdown(v bool) *SubmitDigitalDocStructureJobAdvanceRequest
	GetRevealMarkdown() *bool
	SetUseUrlResponseBody(v bool) *SubmitDigitalDocStructureJobAdvanceRequest
	GetUseUrlResponseBody() *bool
}

type SubmitDigitalDocStructureJobAdvanceRequest struct {
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
	FileUrlObject      io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ImageStrategy      *string   `json:"ImageStrategy,omitempty" xml:"ImageStrategy,omitempty"`
	OssBucket          *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	RevealMarkdown     *bool     `json:"RevealMarkdown,omitempty" xml:"RevealMarkdown,omitempty"`
	UseUrlResponseBody *bool     `json:"UseUrlResponseBody,omitempty" xml:"UseUrlResponseBody,omitempty"`
}

func (s SubmitDigitalDocStructureJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDigitalDocStructureJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetImageStrategy() *string {
	return s.ImageStrategy
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetRevealMarkdown() *bool {
	return s.RevealMarkdown
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) GetUseUrlResponseBody() *bool {
	return s.UseUrlResponseBody
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetFileName(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetImageStrategy(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.ImageStrategy = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetOssBucket(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetOssEndpoint(v string) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetRevealMarkdown(v bool) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.RevealMarkdown = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) SetUseUrlResponseBody(v bool) *SubmitDigitalDocStructureJobAdvanceRequest {
	s.UseUrlResponseBody = &v
	return s
}

func (s *SubmitDigitalDocStructureJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
