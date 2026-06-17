// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransportCertificateInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *TransportCertificateInfo
	GetAlgorithm() *string
	SetCertIdentifier(v string) *TransportCertificateInfo
	GetCertIdentifier() *string
	SetCertName(v string) *TransportCertificateInfo
	GetCertName() *string
	SetCertificateMatchStatus(v string) *TransportCertificateInfo
	GetCertificateMatchStatus() *string
	SetCommonName(v string) *TransportCertificateInfo
	GetCommonName() *string
	SetCoveredDomains(v []*string) *TransportCertificateInfo
	GetCoveredDomains() []*string
	SetIssuer(v string) *TransportCertificateInfo
	GetIssuer() *string
	SetMatchedDomains(v []*string) *TransportCertificateInfo
	GetMatchedDomains() []*string
	SetNotAfterTimestamp(v int64) *TransportCertificateInfo
	GetNotAfterTimestamp() *int64
	SetNotBeforeTimestamp(v int64) *TransportCertificateInfo
	GetNotBeforeTimestamp() *int64
	SetSans(v string) *TransportCertificateInfo
	GetSans() *string
}

type TransportCertificateInfo struct {
	Algorithm              *string   `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	CertIdentifier         *string   `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	CertName               *string   `json:"certName,omitempty" xml:"certName,omitempty"`
	CertificateMatchStatus *string   `json:"certificateMatchStatus,omitempty" xml:"certificateMatchStatus,omitempty"`
	CommonName             *string   `json:"commonName,omitempty" xml:"commonName,omitempty"`
	CoveredDomains         []*string `json:"coveredDomains,omitempty" xml:"coveredDomains,omitempty" type:"Repeated"`
	Issuer                 *string   `json:"issuer,omitempty" xml:"issuer,omitempty"`
	MatchedDomains         []*string `json:"matchedDomains,omitempty" xml:"matchedDomains,omitempty" type:"Repeated"`
	NotAfterTimestamp      *int64    `json:"notAfterTimestamp,omitempty" xml:"notAfterTimestamp,omitempty"`
	NotBeforeTimestamp     *int64    `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	Sans                   *string   `json:"sans,omitempty" xml:"sans,omitempty"`
}

func (s TransportCertificateInfo) String() string {
	return dara.Prettify(s)
}

func (s TransportCertificateInfo) GoString() string {
	return s.String()
}

func (s *TransportCertificateInfo) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *TransportCertificateInfo) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *TransportCertificateInfo) GetCertName() *string {
	return s.CertName
}

func (s *TransportCertificateInfo) GetCertificateMatchStatus() *string {
	return s.CertificateMatchStatus
}

func (s *TransportCertificateInfo) GetCommonName() *string {
	return s.CommonName
}

func (s *TransportCertificateInfo) GetCoveredDomains() []*string {
	return s.CoveredDomains
}

func (s *TransportCertificateInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *TransportCertificateInfo) GetMatchedDomains() []*string {
	return s.MatchedDomains
}

func (s *TransportCertificateInfo) GetNotAfterTimestamp() *int64 {
	return s.NotAfterTimestamp
}

func (s *TransportCertificateInfo) GetNotBeforeTimestamp() *int64 {
	return s.NotBeforeTimestamp
}

func (s *TransportCertificateInfo) GetSans() *string {
	return s.Sans
}

func (s *TransportCertificateInfo) SetAlgorithm(v string) *TransportCertificateInfo {
	s.Algorithm = &v
	return s
}

func (s *TransportCertificateInfo) SetCertIdentifier(v string) *TransportCertificateInfo {
	s.CertIdentifier = &v
	return s
}

func (s *TransportCertificateInfo) SetCertName(v string) *TransportCertificateInfo {
	s.CertName = &v
	return s
}

func (s *TransportCertificateInfo) SetCertificateMatchStatus(v string) *TransportCertificateInfo {
	s.CertificateMatchStatus = &v
	return s
}

func (s *TransportCertificateInfo) SetCommonName(v string) *TransportCertificateInfo {
	s.CommonName = &v
	return s
}

func (s *TransportCertificateInfo) SetCoveredDomains(v []*string) *TransportCertificateInfo {
	s.CoveredDomains = v
	return s
}

func (s *TransportCertificateInfo) SetIssuer(v string) *TransportCertificateInfo {
	s.Issuer = &v
	return s
}

func (s *TransportCertificateInfo) SetMatchedDomains(v []*string) *TransportCertificateInfo {
	s.MatchedDomains = v
	return s
}

func (s *TransportCertificateInfo) SetNotAfterTimestamp(v int64) *TransportCertificateInfo {
	s.NotAfterTimestamp = &v
	return s
}

func (s *TransportCertificateInfo) SetNotBeforeTimestamp(v int64) *TransportCertificateInfo {
	s.NotBeforeTimestamp = &v
	return s
}

func (s *TransportCertificateInfo) SetSans(v string) *TransportCertificateInfo {
	s.Sans = &v
	return s
}

func (s *TransportCertificateInfo) Validate() error {
	return dara.Validate(s)
}
