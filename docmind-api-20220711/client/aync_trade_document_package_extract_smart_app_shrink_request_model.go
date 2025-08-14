// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAyncTradeDocumentPackageExtractSmartAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomExtractionRangeShrink(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest
	GetCustomExtractionRangeShrink() *string
	SetFileName(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest
	GetFileName() *string
	SetFileUrl(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest
	GetFileUrl() *string
	SetOption(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest
	GetOption() *string
	SetTemplateName(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest
	GetTemplateName() *string
}

type AyncTradeDocumentPackageExtractSmartAppShrinkRequest struct {
	CustomExtractionRangeShrink *string `json:"CustomExtractionRange,omitempty" xml:"CustomExtractionRange,omitempty"`
	FileName                    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Option       *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GetCustomExtractionRangeShrink() *string {
	return s.CustomExtractionRangeShrink
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GetOption() *string {
	return s.Option
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetCustomExtractionRangeShrink(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.CustomExtractionRangeShrink = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetFileName(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.FileName = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetFileUrl(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetOption(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.Option = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) SetTemplateName(v string) *AyncTradeDocumentPackageExtractSmartAppShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
