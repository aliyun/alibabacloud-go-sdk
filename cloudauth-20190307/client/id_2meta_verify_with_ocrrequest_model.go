// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId2MetaVerifyWithOCRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertFile(v string) *Id2MetaVerifyWithOCRRequest
	GetCertFile() *string
	SetCertNationalFile(v string) *Id2MetaVerifyWithOCRRequest
	GetCertNationalFile() *string
	SetCertNationalUrl(v string) *Id2MetaVerifyWithOCRRequest
	GetCertNationalUrl() *string
	SetCertUrl(v string) *Id2MetaVerifyWithOCRRequest
	GetCertUrl() *string
}

type Id2MetaVerifyWithOCRRequest struct {
	// Input stream for the portrait side of the ID card image.
	//
	// Choose one between CertUrl and CertFile.
	//
	// example:
	//
	// 无
	CertFile *string `json:"CertFile,omitempty" xml:"CertFile,omitempty"`
	// National emblem side of the ID card image address.
	//
	// Choose one between CertNationalUrl and CertNationalFile, or omit both.
	//
	// example:
	//
	// 无
	CertNationalFile *string `json:"CertNationalFile,omitempty" xml:"CertNationalFile,omitempty"`
	// National emblem side of the ID card image URL. National emblem side
	//
	// A publicly accessible HTTP or HTTPS link.
	//
	// Choose one between CertNationalUrl and CertNationalFile, or omit both.
	//
	// example:
	//
	// https://www.aliyun.com/cert.jpeg
	CertNationalUrl *string `json:"CertNationalUrl,omitempty" xml:"CertNationalUrl,omitempty"`
	// Portrait side of the ID card image.
	//
	// A publicly accessible HTTP or HTTPS link.
	//
	// Choose one between CertUrl and CertFile.
	//
	// example:
	//
	// https://www.aliyun.com/cert.jpeg
	CertUrl *string `json:"CertUrl,omitempty" xml:"CertUrl,omitempty"`
}

func (s Id2MetaVerifyWithOCRRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyWithOCRRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyWithOCRRequest) GetCertFile() *string {
	return s.CertFile
}

func (s *Id2MetaVerifyWithOCRRequest) GetCertNationalFile() *string {
	return s.CertNationalFile
}

func (s *Id2MetaVerifyWithOCRRequest) GetCertNationalUrl() *string {
	return s.CertNationalUrl
}

func (s *Id2MetaVerifyWithOCRRequest) GetCertUrl() *string {
	return s.CertUrl
}

func (s *Id2MetaVerifyWithOCRRequest) SetCertFile(v string) *Id2MetaVerifyWithOCRRequest {
	s.CertFile = &v
	return s
}

func (s *Id2MetaVerifyWithOCRRequest) SetCertNationalFile(v string) *Id2MetaVerifyWithOCRRequest {
	s.CertNationalFile = &v
	return s
}

func (s *Id2MetaVerifyWithOCRRequest) SetCertNationalUrl(v string) *Id2MetaVerifyWithOCRRequest {
	s.CertNationalUrl = &v
	return s
}

func (s *Id2MetaVerifyWithOCRRequest) SetCertUrl(v string) *Id2MetaVerifyWithOCRRequest {
	s.CertUrl = &v
	return s
}

func (s *Id2MetaVerifyWithOCRRequest) Validate() error {
	return dara.Validate(s)
}
