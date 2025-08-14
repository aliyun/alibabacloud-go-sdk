// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToWordJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToWordJobRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitConvertPdfToWordJobRequest
	GetFileUrl() *string
	SetForceExportInnerImage(v bool) *SubmitConvertPdfToWordJobRequest
	GetForceExportInnerImage() *bool
	SetFormulaEnhancement(v bool) *SubmitConvertPdfToWordJobRequest
	GetFormulaEnhancement() *bool
	SetOption(v string) *SubmitConvertPdfToWordJobRequest
	GetOption() *string
	SetOssBucket(v string) *SubmitConvertPdfToWordJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToWordJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToWordJobRequest struct {
	// example:
	//
	// covertPdfToWord.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl               *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool   `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
	FormulaEnhancement    *bool   `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	Option                *string `json:"Option,omitempty" xml:"Option,omitempty"`
	OssBucket             *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint           *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToWordJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToWordJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToWordJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToWordJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitConvertPdfToWordJobRequest) GetForceExportInnerImage() *bool {
	return s.ForceExportInnerImage
}

func (s *SubmitConvertPdfToWordJobRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitConvertPdfToWordJobRequest) GetOption() *string {
	return s.Option
}

func (s *SubmitConvertPdfToWordJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToWordJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToWordJobRequest) SetFileName(v string) *SubmitConvertPdfToWordJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetFileUrl(v string) *SubmitConvertPdfToWordJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToWordJobRequest {
	s.ForceExportInnerImage = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetFormulaEnhancement(v bool) *SubmitConvertPdfToWordJobRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetOption(v string) *SubmitConvertPdfToWordJobRequest {
	s.Option = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetOssBucket(v string) *SubmitConvertPdfToWordJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) SetOssEndpoint(v string) *SubmitConvertPdfToWordJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToWordJobRequest) Validate() error {
	return dara.Validate(s)
}
