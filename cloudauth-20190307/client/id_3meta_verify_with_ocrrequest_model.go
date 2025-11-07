// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyWithOCRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertFile(v string) *Id3MetaVerifyWithOCRRequest
	GetCertFile() *string
	SetCertNationalFile(v string) *Id3MetaVerifyWithOCRRequest
	GetCertNationalFile() *string
	SetCertNationalUrl(v string) *Id3MetaVerifyWithOCRRequest
	GetCertNationalUrl() *string
	SetCertUrl(v string) *Id3MetaVerifyWithOCRRequest
	GetCertUrl() *string
}

type Id3MetaVerifyWithOCRRequest struct {
	// Input stream for the portrait side of the ID card image. Choose either CertUrl or CertFile.
	//
	// example:
	//
	// 无
	CertFile *string `json:"CertFile,omitempty" xml:"CertFile,omitempty"`
	// URL for the national emblem side of the ID card image. Choose either CertNationalUrl or CertNationalFile, or omit both.
	//
	// example:
	//
	// 无
	CertNationalFile *string `json:"CertNationalFile,omitempty" xml:"CertNationalFile,omitempty"`
	// National emblem side of the ID card image URL. A publicly accessible HTTP or HTTPS link. You can choose either CertNationalUrl or CertNationalFile, or omit both.
	//
	// example:
	//
	// https://www.aliyun.com/cert.jpeg
	CertNationalUrl *string `json:"CertNationalUrl,omitempty" xml:"CertNationalUrl,omitempty"`
	// Portrait side of the ID card image. A publicly accessible HTTP or HTTPS link. Choose either CertUrl or CertFile.
	//
	// example:
	//
	// https://www.aliyun.com/cert.jpeg
	CertUrl *string `json:"CertUrl,omitempty" xml:"CertUrl,omitempty"`
}

func (s Id3MetaVerifyWithOCRRequest) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyWithOCRRequest) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyWithOCRRequest) GetCertFile() *string {
	return s.CertFile
}

func (s *Id3MetaVerifyWithOCRRequest) GetCertNationalFile() *string {
	return s.CertNationalFile
}

func (s *Id3MetaVerifyWithOCRRequest) GetCertNationalUrl() *string {
	return s.CertNationalUrl
}

func (s *Id3MetaVerifyWithOCRRequest) GetCertUrl() *string {
	return s.CertUrl
}

func (s *Id3MetaVerifyWithOCRRequest) SetCertFile(v string) *Id3MetaVerifyWithOCRRequest {
	s.CertFile = &v
	return s
}

func (s *Id3MetaVerifyWithOCRRequest) SetCertNationalFile(v string) *Id3MetaVerifyWithOCRRequest {
	s.CertNationalFile = &v
	return s
}

func (s *Id3MetaVerifyWithOCRRequest) SetCertNationalUrl(v string) *Id3MetaVerifyWithOCRRequest {
	s.CertNationalUrl = &v
	return s
}

func (s *Id3MetaVerifyWithOCRRequest) SetCertUrl(v string) *Id3MetaVerifyWithOCRRequest {
	s.CertUrl = &v
	return s
}

func (s *Id3MetaVerifyWithOCRRequest) Validate() error {
	return dara.Validate(s)
}
