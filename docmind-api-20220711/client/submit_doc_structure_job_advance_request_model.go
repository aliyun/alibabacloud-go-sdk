// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSubmitDocStructureJobAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowPptFormat(v bool) *SubmitDocStructureJobAdvanceRequest
	GetAllowPptFormat() *bool
	SetFileName(v string) *SubmitDocStructureJobAdvanceRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocStructureJobAdvanceRequest
	GetFileNameExtension() *string
	SetFileUrlObject(v io.Reader) *SubmitDocStructureJobAdvanceRequest
	GetFileUrlObject() io.Reader
	SetFormulaEnhancement(v bool) *SubmitDocStructureJobAdvanceRequest
	GetFormulaEnhancement() *bool
	SetOssBucket(v string) *SubmitDocStructureJobAdvanceRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocStructureJobAdvanceRequest
	GetOssEndpoint() *string
	SetPageIndex(v string) *SubmitDocStructureJobAdvanceRequest
	GetPageIndex() *string
	SetStructureType(v string) *SubmitDocStructureJobAdvanceRequest
	GetStructureType() *string
}

type SubmitDocStructureJobAdvanceRequest struct {
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
	FileUrlObject      io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	FormulaEnhancement *bool     `json:"FormulaEnhancement,omitempty" xml:"FormulaEnhancement,omitempty"`
	OssBucket          *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint        *string   `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	PageIndex          *string   `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	StructureType      *string   `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s SubmitDocStructureJobAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocStructureJobAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobAdvanceRequest) GetAllowPptFormat() *bool {
	return s.AllowPptFormat
}

func (s *SubmitDocStructureJobAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocStructureJobAdvanceRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocStructureJobAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *SubmitDocStructureJobAdvanceRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitDocStructureJobAdvanceRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocStructureJobAdvanceRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocStructureJobAdvanceRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *SubmitDocStructureJobAdvanceRequest) GetStructureType() *string {
	return s.StructureType
}

func (s *SubmitDocStructureJobAdvanceRequest) SetAllowPptFormat(v bool) *SubmitDocStructureJobAdvanceRequest {
	s.AllowPptFormat = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFileName(v string) *SubmitDocStructureJobAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFileNameExtension(v string) *SubmitDocStructureJobAdvanceRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFileUrlObject(v io.Reader) *SubmitDocStructureJobAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetFormulaEnhancement(v bool) *SubmitDocStructureJobAdvanceRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetOssBucket(v string) *SubmitDocStructureJobAdvanceRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetOssEndpoint(v string) *SubmitDocStructureJobAdvanceRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetPageIndex(v string) *SubmitDocStructureJobAdvanceRequest {
	s.PageIndex = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) SetStructureType(v string) *SubmitDocStructureJobAdvanceRequest {
	s.StructureType = &v
	return s
}

func (s *SubmitDocStructureJobAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
