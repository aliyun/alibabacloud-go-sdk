// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iTranslateCertificateAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateType(v string) *TranslateCertificateAdvanceRequest
	GetCertificateType() *string
	SetImageUrlObject(v io.Reader) *TranslateCertificateAdvanceRequest
	GetImageUrlObject() io.Reader
	SetResultType(v string) *TranslateCertificateAdvanceRequest
	GetResultType() *string
	SetSourceLanguage(v string) *TranslateCertificateAdvanceRequest
	GetSourceLanguage() *string
	SetTargetLanguage(v string) *TranslateCertificateAdvanceRequest
	GetTargetLanguage() *string
}

type TranslateCertificateAdvanceRequest struct {
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
	ImageUrlObject io.Reader `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
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

func (s TranslateCertificateAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s TranslateCertificateAdvanceRequest) GoString() string {
	return s.String()
}

func (s *TranslateCertificateAdvanceRequest) GetCertificateType() *string {
	return s.CertificateType
}

func (s *TranslateCertificateAdvanceRequest) GetImageUrlObject() io.Reader {
	return s.ImageUrlObject
}

func (s *TranslateCertificateAdvanceRequest) GetResultType() *string {
	return s.ResultType
}

func (s *TranslateCertificateAdvanceRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *TranslateCertificateAdvanceRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *TranslateCertificateAdvanceRequest) SetCertificateType(v string) *TranslateCertificateAdvanceRequest {
	s.CertificateType = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetImageUrlObject(v io.Reader) *TranslateCertificateAdvanceRequest {
	s.ImageUrlObject = v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetResultType(v string) *TranslateCertificateAdvanceRequest {
	s.ResultType = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetSourceLanguage(v string) *TranslateCertificateAdvanceRequest {
	s.SourceLanguage = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) SetTargetLanguage(v string) *TranslateCertificateAdvanceRequest {
	s.TargetLanguage = &v
	return s
}

func (s *TranslateCertificateAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
