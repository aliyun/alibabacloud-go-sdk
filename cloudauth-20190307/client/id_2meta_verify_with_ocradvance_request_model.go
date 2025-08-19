// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iId2MetaVerifyWithOCRAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertFileObject(v io.Reader) *Id2MetaVerifyWithOCRAdvanceRequest
	GetCertFileObject() io.Reader
	SetCertNationalFileObject(v io.Reader) *Id2MetaVerifyWithOCRAdvanceRequest
	GetCertNationalFileObject() io.Reader
	SetCertNationalUrl(v string) *Id2MetaVerifyWithOCRAdvanceRequest
	GetCertNationalUrl() *string
	SetCertUrl(v string) *Id2MetaVerifyWithOCRAdvanceRequest
	GetCertUrl() *string
}

type Id2MetaVerifyWithOCRAdvanceRequest struct {
	// Input stream for the portrait side of the ID card image.
	//
	// Choose one between CertUrl and CertFile.
	//
	// example:
	//
	// 无
	CertFileObject io.Reader `json:"CertFile,omitempty" xml:"CertFile,omitempty"`
	// National emblem side of the ID card image address.
	//
	// Choose one between CertNationalUrl and CertNationalFile, or omit both.
	//
	// example:
	//
	// 无
	CertNationalFileObject io.Reader `json:"CertNationalFile,omitempty" xml:"CertNationalFile,omitempty"`
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

func (s Id2MetaVerifyWithOCRAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s Id2MetaVerifyWithOCRAdvanceRequest) GoString() string {
	return s.String()
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) GetCertFileObject() io.Reader {
	return s.CertFileObject
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) GetCertNationalFileObject() io.Reader {
	return s.CertNationalFileObject
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) GetCertNationalUrl() *string {
	return s.CertNationalUrl
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) GetCertUrl() *string {
	return s.CertUrl
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) SetCertFileObject(v io.Reader) *Id2MetaVerifyWithOCRAdvanceRequest {
	s.CertFileObject = v
	return s
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) SetCertNationalFileObject(v io.Reader) *Id2MetaVerifyWithOCRAdvanceRequest {
	s.CertNationalFileObject = v
	return s
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) SetCertNationalUrl(v string) *Id2MetaVerifyWithOCRAdvanceRequest {
	s.CertNationalUrl = &v
	return s
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) SetCertUrl(v string) *Id2MetaVerifyWithOCRAdvanceRequest {
	s.CertUrl = &v
	return s
}

func (s *Id2MetaVerifyWithOCRAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
