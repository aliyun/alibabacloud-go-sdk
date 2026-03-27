// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocStructureJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowPptFormat(v bool) *SubmitDocStructureJobShrinkRequest
	GetAllowPptFormat() *bool
	SetEnableEventCallback(v bool) *SubmitDocStructureJobShrinkRequest
	GetEnableEventCallback() *bool
	SetFileName(v string) *SubmitDocStructureJobShrinkRequest
	GetFileName() *string
	SetFileNameExtension(v string) *SubmitDocStructureJobShrinkRequest
	GetFileNameExtension() *string
	SetFileUrl(v string) *SubmitDocStructureJobShrinkRequest
	GetFileUrl() *string
	SetFormulaEnhancement(v bool) *SubmitDocStructureJobShrinkRequest
	GetFormulaEnhancement() *bool
	SetOssBucket(v string) *SubmitDocStructureJobShrinkRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *SubmitDocStructureJobShrinkRequest
	GetOssEndpoint() *string
	SetOutputFormatShrink(v string) *SubmitDocStructureJobShrinkRequest
	GetOutputFormatShrink() *string
	SetPageIndex(v string) *SubmitDocStructureJobShrinkRequest
	GetPageIndex() *string
	SetStructureType(v string) *SubmitDocStructureJobShrinkRequest
	GetStructureType() *string
}

type SubmitDocStructureJobShrinkRequest struct {
	AllowPptFormat      *bool `json:"AllowPptFormat,omitempty" xml:"AllowPptFormat,omitempty"`
	EnableEventCallback *bool `json:"EnableEventCallback,omitempty" xml:"EnableEventCallback,omitempty"`
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
	OutputFormatShrink *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	PageIndex          *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	StructureType      *string `json:"StructureType,omitempty" xml:"StructureType,omitempty"`
}

func (s SubmitDocStructureJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocStructureJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDocStructureJobShrinkRequest) GetAllowPptFormat() *bool {
	return s.AllowPptFormat
}

func (s *SubmitDocStructureJobShrinkRequest) GetEnableEventCallback() *bool {
	return s.EnableEventCallback
}

func (s *SubmitDocStructureJobShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitDocStructureJobShrinkRequest) GetFileNameExtension() *string {
	return s.FileNameExtension
}

func (s *SubmitDocStructureJobShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *SubmitDocStructureJobShrinkRequest) GetFormulaEnhancement() *bool {
	return s.FormulaEnhancement
}

func (s *SubmitDocStructureJobShrinkRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *SubmitDocStructureJobShrinkRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *SubmitDocStructureJobShrinkRequest) GetOutputFormatShrink() *string {
	return s.OutputFormatShrink
}

func (s *SubmitDocStructureJobShrinkRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *SubmitDocStructureJobShrinkRequest) GetStructureType() *string {
	return s.StructureType
}

func (s *SubmitDocStructureJobShrinkRequest) SetAllowPptFormat(v bool) *SubmitDocStructureJobShrinkRequest {
	s.AllowPptFormat = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetEnableEventCallback(v bool) *SubmitDocStructureJobShrinkRequest {
	s.EnableEventCallback = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetFileName(v string) *SubmitDocStructureJobShrinkRequest {
	s.FileName = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetFileNameExtension(v string) *SubmitDocStructureJobShrinkRequest {
	s.FileNameExtension = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetFileUrl(v string) *SubmitDocStructureJobShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetFormulaEnhancement(v bool) *SubmitDocStructureJobShrinkRequest {
	s.FormulaEnhancement = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetOssBucket(v string) *SubmitDocStructureJobShrinkRequest {
	s.OssBucket = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetOssEndpoint(v string) *SubmitDocStructureJobShrinkRequest {
	s.OssEndpoint = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetOutputFormatShrink(v string) *SubmitDocStructureJobShrinkRequest {
	s.OutputFormatShrink = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetPageIndex(v string) *SubmitDocStructureJobShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) SetStructureType(v string) *SubmitDocStructureJobShrinkRequest {
	s.StructureType = &v
	return s
}

func (s *SubmitDocStructureJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
