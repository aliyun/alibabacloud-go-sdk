// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatRefDocPageInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAngle(v float64) *ChatRefDocPageInfo
	GetAngle() *float64
	SetExcelParseResult(v string) *ChatRefDocPageInfo
	GetExcelParseResult() *string
	SetImageHeight(v int32) *ChatRefDocPageInfo
	GetImageHeight() *int32
	SetImageUrl(v string) *ChatRefDocPageInfo
	GetImageUrl() *string
	SetImageWidth(v int32) *ChatRefDocPageInfo
	GetImageWidth() *int32
	SetPageIdCurDoc(v int32) *ChatRefDocPageInfo
	GetPageIdCurDoc() *int32
	SetPdfParseResult(v string) *ChatRefDocPageInfo
	GetPdfParseResult() *string
	SetWordParseResult(v string) *ChatRefDocPageInfo
	GetWordParseResult() *string
}

type ChatRefDocPageInfo struct {
	Angle            *float64 `json:"angle,omitempty" xml:"angle,omitempty"`
	ExcelParseResult *string  `json:"excelParseResult,omitempty" xml:"excelParseResult,omitempty"`
	ImageHeight      *int32   `json:"imageHeight,omitempty" xml:"imageHeight,omitempty"`
	ImageUrl         *string  `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	ImageWidth       *int32   `json:"imageWidth,omitempty" xml:"imageWidth,omitempty"`
	PageIdCurDoc     *int32   `json:"pageIdCurDoc,omitempty" xml:"pageIdCurDoc,omitempty"`
	PdfParseResult   *string  `json:"pdfParseResult,omitempty" xml:"pdfParseResult,omitempty"`
	WordParseResult  *string  `json:"wordParseResult,omitempty" xml:"wordParseResult,omitempty"`
}

func (s ChatRefDocPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatRefDocPageInfo) GoString() string {
	return s.String()
}

func (s *ChatRefDocPageInfo) GetAngle() *float64 {
	return s.Angle
}

func (s *ChatRefDocPageInfo) GetExcelParseResult() *string {
	return s.ExcelParseResult
}

func (s *ChatRefDocPageInfo) GetImageHeight() *int32 {
	return s.ImageHeight
}

func (s *ChatRefDocPageInfo) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ChatRefDocPageInfo) GetImageWidth() *int32 {
	return s.ImageWidth
}

func (s *ChatRefDocPageInfo) GetPageIdCurDoc() *int32 {
	return s.PageIdCurDoc
}

func (s *ChatRefDocPageInfo) GetPdfParseResult() *string {
	return s.PdfParseResult
}

func (s *ChatRefDocPageInfo) GetWordParseResult() *string {
	return s.WordParseResult
}

func (s *ChatRefDocPageInfo) SetAngle(v float64) *ChatRefDocPageInfo {
	s.Angle = &v
	return s
}

func (s *ChatRefDocPageInfo) SetExcelParseResult(v string) *ChatRefDocPageInfo {
	s.ExcelParseResult = &v
	return s
}

func (s *ChatRefDocPageInfo) SetImageHeight(v int32) *ChatRefDocPageInfo {
	s.ImageHeight = &v
	return s
}

func (s *ChatRefDocPageInfo) SetImageUrl(v string) *ChatRefDocPageInfo {
	s.ImageUrl = &v
	return s
}

func (s *ChatRefDocPageInfo) SetImageWidth(v int32) *ChatRefDocPageInfo {
	s.ImageWidth = &v
	return s
}

func (s *ChatRefDocPageInfo) SetPageIdCurDoc(v int32) *ChatRefDocPageInfo {
	s.PageIdCurDoc = &v
	return s
}

func (s *ChatRefDocPageInfo) SetPdfParseResult(v string) *ChatRefDocPageInfo {
	s.PdfParseResult = &v
	return s
}

func (s *ChatRefDocPageInfo) SetWordParseResult(v string) *ChatRefDocPageInfo {
	s.WordParseResult = &v
	return s
}

func (s *ChatRefDocPageInfo) Validate() error {
	return dara.Validate(s)
}
