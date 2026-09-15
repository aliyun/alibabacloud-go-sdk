// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomDomainResp interface {
	dara.Model
	String() string
	GoString() string
	SetCertConfig(v *CertConfig) *CustomDomainResp
	GetCertConfig() *CertConfig
	SetCreatedAt(v int64) *CustomDomainResp
	GetCreatedAt() *int64
	SetDescription(v string) *CustomDomainResp
	GetDescription() *string
	SetDomainName(v string) *CustomDomainResp
	GetDomainName() *string
	SetTlsConfig(v *TLSConfig) *CustomDomainResp
	GetTlsConfig() *TLSConfig
	SetUpdatedAt(v int64) *CustomDomainResp
	GetUpdatedAt() *int64
}

type CustomDomainResp struct {
	CertConfig  *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedAt   *int64      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	DomainName  *string     `json:"domainName,omitempty" xml:"domainName,omitempty"`
	TlsConfig   *TLSConfig  `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	UpdatedAt   *int64      `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CustomDomainResp) String() string {
	return dara.Prettify(s)
}

func (s CustomDomainResp) GoString() string {
	return s.String()
}

func (s *CustomDomainResp) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CustomDomainResp) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *CustomDomainResp) GetDescription() *string {
	return s.Description
}

func (s *CustomDomainResp) GetDomainName() *string {
	return s.DomainName
}

func (s *CustomDomainResp) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *CustomDomainResp) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *CustomDomainResp) SetCertConfig(v *CertConfig) *CustomDomainResp {
	s.CertConfig = v
	return s
}

func (s *CustomDomainResp) SetCreatedAt(v int64) *CustomDomainResp {
	s.CreatedAt = &v
	return s
}

func (s *CustomDomainResp) SetDescription(v string) *CustomDomainResp {
	s.Description = &v
	return s
}

func (s *CustomDomainResp) SetDomainName(v string) *CustomDomainResp {
	s.DomainName = &v
	return s
}

func (s *CustomDomainResp) SetTlsConfig(v *TLSConfig) *CustomDomainResp {
	s.TlsConfig = v
	return s
}

func (s *CustomDomainResp) SetUpdatedAt(v int64) *CustomDomainResp {
	s.UpdatedAt = &v
	return s
}

func (s *CustomDomainResp) Validate() error {
	if s.CertConfig != nil {
		if err := s.CertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TlsConfig != nil {
		if err := s.TlsConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
