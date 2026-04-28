// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCertBody(v string) *CertInfo
	GetCertBody() *string
	SetCertName(v string) *CertInfo
	GetCertName() *string
	SetCertPrivatekey(v string) *CertInfo
	GetCertPrivatekey() *string
}

type CertInfo struct {
	// example:
	//
	// xxx
	CertBody *string `json:"cert_body,omitempty" xml:"cert_body,omitempty"`
	// example:
	//
	// xxx
	CertName *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	// example:
	//
	// xxx
	CertPrivatekey *string `json:"cert_privatekey,omitempty" xml:"cert_privatekey,omitempty"`
}

func (s CertInfo) String() string {
	return dara.Prettify(s)
}

func (s CertInfo) GoString() string {
	return s.String()
}

func (s *CertInfo) GetCertBody() *string {
	return s.CertBody
}

func (s *CertInfo) GetCertName() *string {
	return s.CertName
}

func (s *CertInfo) GetCertPrivatekey() *string {
	return s.CertPrivatekey
}

func (s *CertInfo) SetCertBody(v string) *CertInfo {
	s.CertBody = &v
	return s
}

func (s *CertInfo) SetCertName(v string) *CertInfo {
	s.CertName = &v
	return s
}

func (s *CertInfo) SetCertPrivatekey(v string) *CertInfo {
	s.CertPrivatekey = &v
	return s
}

func (s *CertInfo) Validate() error {
	return dara.Validate(s)
}
