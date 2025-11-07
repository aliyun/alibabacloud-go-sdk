// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iId3MetaVerifyWithOCRAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertFileObject(v io.Reader) *Id3MetaVerifyWithOCRAdvanceRequest
	GetCertFileObject() io.Reader
	SetCertNationalFileObject(v io.Reader) *Id3MetaVerifyWithOCRAdvanceRequest
	GetCertNationalFileObject() io.Reader
	SetCertNationalUrl(v string) *Id3MetaVerifyWithOCRAdvanceRequest
	GetCertNationalUrl() *string
	SetCertUrl(v string) *Id3MetaVerifyWithOCRAdvanceRequest
	GetCertUrl() *string
}

type Id3MetaVerifyWithOCRAdvanceRequest struct {
	// Input stream for the portrait side of the ID card image. Choose either CertUrl or CertFile.
	//
	// example:
	//
	// 无
	CertFileObject io.Reader `json:"CertFile,omitempty" xml:"CertFile,omitempty"`
	// URL for the national emblem side of the ID card image. Choose either CertNationalUrl or CertNationalFile, or omit both.
	//
	// example:
	//
	// 无
	CertNationalFileObject io.Reader `json:"CertNationalFile,omitempty" xml:"CertNationalFile,omitempty"`
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

func (s Id3MetaVerifyWithOCRAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyWithOCRAdvanceRequest) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) GetCertFileObject() io.Reader {
	return s.CertFileObject
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) GetCertNationalFileObject() io.Reader {
	return s.CertNationalFileObject
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) GetCertNationalUrl() *string {
	return s.CertNationalUrl
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) GetCertUrl() *string {
	return s.CertUrl
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) SetCertFileObject(v io.Reader) *Id3MetaVerifyWithOCRAdvanceRequest {
	s.CertFileObject = v
	return s
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) SetCertNationalFileObject(v io.Reader) *Id3MetaVerifyWithOCRAdvanceRequest {
	s.CertNationalFileObject = v
	return s
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) SetCertNationalUrl(v string) *Id3MetaVerifyWithOCRAdvanceRequest {
	s.CertNationalUrl = &v
	return s
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) SetCertUrl(v string) *Id3MetaVerifyWithOCRAdvanceRequest {
	s.CertUrl = &v
	return s
}

func (s *Id3MetaVerifyWithOCRAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
