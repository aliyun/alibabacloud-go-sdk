// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitConvertPdfToExcelJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToExcelJobRequest
	GetFileName() *string
	SetFileUrl(v string) *SubmitConvertPdfToExcelJobRequest
	GetFileUrl() *string
	SetForceExportInnerImage(v bool) *SubmitConvertPdfToExcelJobRequest
	GetForceExportInnerImage() *bool
	SetForceMergeExcel(v bool) *SubmitConvertPdfToExcelJobRequest
	GetForceMergeExcel() *bool
	SetOssBucket(v string) *SubmitConvertPdfToExcelJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToExcelJobRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToExcelJobRequest struct {
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrl               *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool   `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
	ForceMergeExcel       *bool   `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	OssBucket             *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint           *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToExcelJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToExcelJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitConvertPdfToExcelJobRequest) GetForceExportInnerImage() *bool {
	return s.ForceExportInnerImage
}

func (s *SubmitConvertPdfToExcelJobRequest) GetForceMergeExcel() *bool {
	return s.ForceMergeExcel
}

func (s *SubmitConvertPdfToExcelJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToExcelJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToExcelJobRequest) SetFileName(v string) *SubmitConvertPdfToExcelJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetFileUrl(v string) *SubmitConvertPdfToExcelJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToExcelJobRequest {
	s.ForceExportInnerImage = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetForceMergeExcel(v bool) *SubmitConvertPdfToExcelJobRequest {
	s.ForceMergeExcel = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetOssBucket(v string) *SubmitConvertPdfToExcelJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) SetOssEndpoint(v string) *SubmitConvertPdfToExcelJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobRequest) Validate() error {
	return dara.Validate(s)
}
