// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTranslateCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateType(v string) *TranslateCertificateRequest
	GetCertificateType() *string
	SetImageUrl(v string) *TranslateCertificateRequest
	GetImageUrl() *string
	SetResultType(v string) *TranslateCertificateRequest
	GetResultType() *string
	SetSourceLanguage(v string) *TranslateCertificateRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TranslateCertificateRequest
	GetTargetLanguage() *string
}

type TranslateCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// driving_license
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://imageurl
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// text
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
}

func (s TranslateCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateCertificateRequest) GoString() string {
	return s.String()
}

func (s *TranslateCertificateRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *TranslateCertificateRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *TranslateCertificateRequest) GetResultType() *string {
	return s.ResultType
}

func (s *TranslateCertificateRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateCertificateRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateCertificateRequest) SetCertificateType(v string) *TranslateCertificateRequest {
	s.CertificateType = &v
	return s
}

func (s *TranslateCertificateRequest) SetImageUrl(v string) *TranslateCertificateRequest {
	s.ImageUrl = &v
	return s
}

func (s *TranslateCertificateRequest) SetResultType(v string) *TranslateCertificateRequest {
	s.ResultType = &v
	return s
}

func (s *TranslateCertificateRequest) SetSourceLanguage(v string) *TranslateCertificateRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateCertificateRequest) SetTargetLanguage(v string) *TranslateCertificateRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateCertificateRequest) Validate() error {
	return dara.Validate(s)
}
