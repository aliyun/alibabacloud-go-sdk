// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserCertificate interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *UserCertificate
	GetCertId() *string
	SetCertificate(v string) *UserCertificate
	GetCertificate() *string
	SetDescription(v string) *UserCertificate
	GetDescription() *string
	SetDnsNames(v []*string) *UserCertificate
	GetDnsNames() []*string
	SetExpTimeUnix(v int64) *UserCertificate
	GetExpTimeUnix() *int64
	SetGmtCreateUnix(v int64) *UserCertificate
	GetGmtCreateUnix() *int64
	SetGmtModifiedUnix(v int64) *UserCertificate
	GetGmtModifiedUnix() *int64
	SetName(v string) *UserCertificate
	GetName() *string
	SetPrivateKey(v string) *UserCertificate
	GetPrivateKey() *string
}

type UserCertificate struct {
	// example:
	//
	// cert-aabbccdd
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// 用于测试
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// example:
	//
	// xxxx
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DnsNames    []*string `json:"DnsNames,omitempty" xml:"DnsNames,omitempty" type:"Repeated"`
	// example:
	//
	// 1734492686
	ExpTimeUnix *int64 `json:"ExpTimeUnix,omitempty" xml:"ExpTimeUnix,omitempty"`
	// example:
	//
	// 1734523812
	GmtCreateUnix *int64 `json:"GmtCreateUnix,omitempty" xml:"GmtCreateUnix,omitempty"`
	// example:
	//
	// 1734523812
	GmtModifiedUnix *int64 `json:"GmtModifiedUnix,omitempty" xml:"GmtModifiedUnix,omitempty"`
	// example:
	//
	// 证书1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// xxxx
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
}

func (s UserCertificate) String() string {
	return dara.Prettify(s)
}

func (s UserCertificate) GoString() string {
	return s.String()
}

func (s *UserCertificate) GetCertId() *string {
	return s.CertId
}

func (s *UserCertificate) GetCertificate() *string {
	return s.Certificate
}

func (s *UserCertificate) GetDescription() *string {
	return s.Description
}

func (s *UserCertificate) GetDnsNames() []*string {
	return s.DnsNames
}

func (s *UserCertificate) GetExpTimeUnix() *int64 {
	return s.ExpTimeUnix
}

func (s *UserCertificate) GetGmtCreateUnix() *int64 {
	return s.GmtCreateUnix
}

func (s *UserCertificate) GetGmtModifiedUnix() *int64 {
	return s.GmtModifiedUnix
}

func (s *UserCertificate) GetName() *string {
	return s.Name
}

func (s *UserCertificate) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *UserCertificate) SetCertId(v string) *UserCertificate {
	s.CertId = &v
	return s
}

func (s *UserCertificate) SetCertificate(v string) *UserCertificate {
	s.Certificate = &v
	return s
}

func (s *UserCertificate) SetDescription(v string) *UserCertificate {
	s.Description = &v
	return s
}

func (s *UserCertificate) SetDnsNames(v []*string) *UserCertificate {
	s.DnsNames = v
	return s
}

func (s *UserCertificate) SetExpTimeUnix(v int64) *UserCertificate {
	s.ExpTimeUnix = &v
	return s
}

func (s *UserCertificate) SetGmtCreateUnix(v int64) *UserCertificate {
	s.GmtCreateUnix = &v
	return s
}

func (s *UserCertificate) SetGmtModifiedUnix(v int64) *UserCertificate {
	s.GmtModifiedUnix = &v
	return s
}

func (s *UserCertificate) SetName(v string) *UserCertificate {
	s.Name = &v
	return s
}

func (s *UserCertificate) SetPrivateKey(v string) *UserCertificate {
	s.PrivateKey = &v
	return s
}

func (s *UserCertificate) Validate() error {
	return dara.Validate(s)
}
