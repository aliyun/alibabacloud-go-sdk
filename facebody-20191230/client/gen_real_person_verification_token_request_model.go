// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenRealPersonVerificationTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateName(v string) *GenRealPersonVerificationTokenRequest
	GetCertificateName() *string
	SetCertificateNumber(v string) *GenRealPersonVerificationTokenRequest
	GetCertificateNumber() *string
	SetMetaInfo(v string) *GenRealPersonVerificationTokenRequest
	GetMetaInfo() *string
}

type GenRealPersonVerificationTokenRequest struct {
	// This parameter is required.
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 33010219001123123X
	CertificateNumber *string `json:"CertificateNumber,omitempty" xml:"CertificateNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"dwe":"ew4e"...}
	MetaInfo *string `json:"MetaInfo,omitempty" xml:"MetaInfo,omitempty"`
}

func (s GenRealPersonVerificationTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenRealPersonVerificationTokenRequest) GoString() string {
	return s.String()
}

func (s *GenRealPersonVerificationTokenRequest) GetCertificateName() *string {
	return s.CertificateName
}

func (s *GenRealPersonVerificationTokenRequest) GetCertificateNumber() *string {
	return s.CertificateNumber
}

func (s *GenRealPersonVerificationTokenRequest) GetMetaInfo() *string {
	return s.MetaInfo
}

func (s *GenRealPersonVerificationTokenRequest) SetCertificateName(v string) *GenRealPersonVerificationTokenRequest {
	s.CertificateName = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) SetCertificateNumber(v string) *GenRealPersonVerificationTokenRequest {
	s.CertificateNumber = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) SetMetaInfo(v string) *GenRealPersonVerificationTokenRequest {
	s.MetaInfo = &v
	return s
}

func (s *GenRealPersonVerificationTokenRequest) Validate() error {
	return dara.Validate(s)
}
