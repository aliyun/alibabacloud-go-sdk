// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAyncTradeDocumentPackageExtractSmartAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomExtractionRange(v []*string) *AyncTradeDocumentPackageExtractSmartAppRequest
	GetCustomExtractionRange() []*string
	SetFileName(v string) *AyncTradeDocumentPackageExtractSmartAppRequest
	GetFileName() *string
	SetFileUrl(v string) *AyncTradeDocumentPackageExtractSmartAppRequest
	GetFileUrl() *string
	SetOption(v string) *AyncTradeDocumentPackageExtractSmartAppRequest
	GetOption() *string
	SetTemplateName(v string) *AyncTradeDocumentPackageExtractSmartAppRequest
	GetTemplateName() *string
}

type AyncTradeDocumentPackageExtractSmartAppRequest struct {
	CustomExtractionRange []*string `json:"CustomExtractionRange,omitempty" xml:"CustomExtractionRange,omitempty" type:"Repeated"`
	FileName              *string   `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	FileUrl      *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Option       *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s AyncTradeDocumentPackageExtractSmartAppRequest) String() string {
	return dara.Prettify(s)
}

func (s AyncTradeDocumentPackageExtractSmartAppRequest) GoString() string {
	return s.String()
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) GetCustomExtractionRange() []*string {
	return s.CustomExtractionRange
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) GetFileName() *string {
	return s.FileName
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) GetOption() *string {
	return s.Option
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetCustomExtractionRange(v []*string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.CustomExtractionRange = v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetFileName(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.FileName = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetFileUrl(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.FileUrl = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetOption(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.Option = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) SetTemplateName(v string) *AyncTradeDocumentPackageExtractSmartAppRequest {
	s.TemplateName = &v
	return s
}

func (s *AyncTradeDocumentPackageExtractSmartAppRequest) Validate() error {
	return dara.Validate(s)
}
