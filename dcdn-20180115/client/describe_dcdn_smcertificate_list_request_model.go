// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSMCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnSMCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeDcdnSMCertificateListRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnSMCertificateListRequest
	GetSecurityToken() *string
}

type DescribeDcdnSMCertificateListRequest struct {
	// The accelerated domain name whose SM certificates you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnSMCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSMCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSMCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnSMCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnSMCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnSMCertificateListRequest) SetDomainName(v string) *DescribeDcdnSMCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnSMCertificateListRequest) SetOwnerId(v int64) *DescribeDcdnSMCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnSMCertificateListRequest) SetSecurityToken(v string) *DescribeDcdnSMCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnSMCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
