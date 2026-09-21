// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgType(v string) *ApplyCertificateRequest
	GetAlgType() *string
	SetDomains(v string) *ApplyCertificateRequest
	GetDomains() *string
	SetSiteId(v int64) *ApplyCertificateRequest
	GetSiteId() *int64
	SetType(v string) *ApplyCertificateRequest
	GetType() *string
}

type ApplyCertificateRequest struct {
	// The algorithm type.
	AlgType *string `json:"AlgType,omitempty" xml:"AlgType,omitempty"`
	// The list of domain names, separated by commas.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com,blog.example.com
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// The site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The certificate type. Valid values:
	//
	// - lets_encrypt: Let\\"s Encrypt certificate.
	//
	// - digicert_single: DigiCert single-domain certificate.
	//
	// - digicert_wildcard: DigiCert wildcard domain certificate.
	//
	// example:
	//
	// lets_encrypt
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ApplyCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateRequest) GoString() string {
	return s.String()
}

func (s *ApplyCertificateRequest) GetAlgType() *string {
	return s.AlgType
}

func (s *ApplyCertificateRequest) GetDomains() *string {
	return s.Domains
}

func (s *ApplyCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ApplyCertificateRequest) GetType() *string {
	return s.Type
}

func (s *ApplyCertificateRequest) SetAlgType(v string) *ApplyCertificateRequest {
	s.AlgType = &v
	return s
}

func (s *ApplyCertificateRequest) SetDomains(v string) *ApplyCertificateRequest {
	s.Domains = &v
	return s
}

func (s *ApplyCertificateRequest) SetSiteId(v int64) *ApplyCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *ApplyCertificateRequest) SetType(v string) *ApplyCertificateRequest {
	s.Type = &v
	return s
}

func (s *ApplyCertificateRequest) Validate() error {
	return dara.Validate(s)
}
