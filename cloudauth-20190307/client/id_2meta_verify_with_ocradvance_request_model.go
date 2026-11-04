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
	// The input stream of the image for the portrait side of the ID card.
	//
	// Specify either CertUrl or CertFile.
	//
	// example:
	//
	// None
	CertFileObject io.Reader `json:"CertFile,omitempty" xml:"CertFile,omitempty"`
	// The URL of the image of the national emblem side of the ID card. You can specify either CertNationalUrl or CertNationalFile. You can also specify neither.
	//
	// example:
	//
	// None
	CertNationalFileObject io.Reader `json:"CertNationalFile,omitempty" xml:"CertNationalFile,omitempty"`
	// The URL of the image for the national emblem side of the ID card.
	//
	// The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// Specify either CertNationalUrl or CertNationalFile, or leave both empty.
	//
	// example:
	//
	// https://www.aliyun.com/cert.jpeg
	CertNationalUrl *string `json:"CertNationalUrl,omitempty" xml:"CertNationalUrl,omitempty"`
	// The URL of the image for the portrait side of the ID card.
	//
	// The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// Specify either CertUrl or CertFile.
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
