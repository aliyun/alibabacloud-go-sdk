// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocStructureJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowPptFormat(v bool) *SubmitDocStructureJobRequest
	GetAllowPptFormat() *bool
	SetFileName(v string) *SubmitDocStructureJobRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocStructureJobRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDocStructureJobRequest
	GetFileUrl() *string
	SetFormulaEnhancement(v bool) *SubmitDocStructureJobRequest
	GetFormulaEnhancement() *bool
	SetOssBucket(v string) *SubmitDocStructureJobRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocStructureJobRequest
	GetOssEndpoint() *string
	SetPageIndex(v string) *SubmitDocStructureJobRequest
	GetPageIndex() *string
	SetStructureType(v string) *SubmitDocStructureJobRequest
	GetStructureType() *string
}

type SubmitDocStructureJobRequest struct {
	AllowPptFormat *bool `json:"AllowPptFormat,omitempty" xml:"AllowPptFormat,omitempty"`
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
	FormulaEnhancement *bool   `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	OssBucket          *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	PageIndex          *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	StructureType      *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s SubmitDocStructureJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocStructureJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobRequest) GetAllowPptFormat() *bool {
	return s.AllowPptFormat
}

func (s *SubmitDocStructureJobRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocStructureJobRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocStructureJobRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocStructureJobRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitDocStructureJobRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocStructureJobRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocStructureJobRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *SubmitDocStructureJobRequest) GetStructureType() *string {
	return s.StructureType
}

func (s *SubmitDocStructureJobRequest) SetAllowPptFormat(v bool) *SubmitDocStructureJobRequest {
	s.AllowPptFormat = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetFileName(v string) *SubmitDocStructureJobRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetFileNameExtension(v string) *SubmitDocStructureJobRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetFileUrl(v string) *SubmitDocStructureJobRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetFormulaEnhancement(v bool) *SubmitDocStructureJobRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetOssBucket(v string) *SubmitDocStructureJobRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetOssEndpoint(v string) *SubmitDocStructureJobRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetPageIndex(v string) *SubmitDocStructureJobRequest {
	s.PageIndex = &v
	return s
}

func (s *SubmitDocStructureJobRequest) SetStructureType(v string) *SubmitDocStructureJobRequest {
	s.StructureType = &v
	return s
}

func (s *SubmitDocStructureJobRequest) Validate() error {
	return dara.Validate(s)
}
