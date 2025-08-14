// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitConvertPdfToWordJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToWordJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitConvertPdfToWordJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetForceExportInnerImage(v bool) *SubmitConvertPdfToWordJobAdvanceRequest
	GetForceExportInnerImage() *bool
	SetFormulaEnhancement(v bool) *SubmitConvertPdfToWordJobAdvanceRequest
	GetFormulaEnhancement() *bool
	SetOption(v string) *SubmitConvertPdfToWordJobAdvanceRequest
	GetOption() *string
	SetOssBucket(v string) *SubmitConvertPdfToWordJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToWordJobAdvanceRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToWordJobAdvanceRequest struct {
	// example:
	//
	// covertPdfToWord.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject         io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool     `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
	FormulaEnhancement    *bool     `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	Option                *string   `json:"Option,omitempty" xml:"Option,omitempty"`
	OssBucket             *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint           *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToWordJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToWordJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetForceExportInnerImage() *bool {
	return s.ForceExportInnerImage
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetOption() *string {
	return s.Option
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.ForceExportInnerImage = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetFormulaEnhancement(v bool) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetOption(v string) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.Option = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetOssBucket(v string) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) SetOssEndpoint(v string) *SubmitConvertPdfToWordJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToWordJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
