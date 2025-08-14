// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitConvertPdfToExcelJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *SubmitConvertPdfToExcelJobAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *SubmitConvertPdfToExcelJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetForceExportInnerImage(v bool) *SubmitConvertPdfToExcelJobAdvanceRequest
	GetForceExportInnerImage() *bool
	SetForceMergeExcel(v bool) *SubmitConvertPdfToExcelJobAdvanceRequest
	GetForceMergeExcel() *bool
	SetOssBucket(v string) *SubmitConvertPdfToExcelJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitConvertPdfToExcelJobAdvanceRequest
	GetOssEndpoint() *string
}

type SubmitConvertPdfToExcelJobAdvanceRequest struct {
	// example:
	//
	// convertPdfToExcel.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// https://gw.alipayobjects.com/os/basement_prod/598b9edf-5287-4065-9e36-464305c60698.pdf
	FileUrlObject         io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	ForceExportInnerImage *bool     `json:"ForceExportInnerImage,omitempty" xml:"ForceExportInnerImage,omitempty"`
	ForceMergeExcel       *bool     `json:"ForceMergeExcel,omitempty" xml:"ForceMergeExcel,omitempty"`
	OssBucket             *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint           *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
}

func (s SubmitConvertPdfToExcelJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitConvertPdfToExcelJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) GetForceExportInnerImage() *bool {
	return s.ForceExportInnerImage
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) GetForceMergeExcel() *bool {
	return s.ForceMergeExcel
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetFileName(v string) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetForceExportInnerImage(v bool) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.ForceExportInnerImage = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetForceMergeExcel(v bool) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.ForceMergeExcel = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetOssBucket(v string) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) SetOssEndpoint(v string) *SubmitConvertPdfToExcelJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitConvertPdfToExcelJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
