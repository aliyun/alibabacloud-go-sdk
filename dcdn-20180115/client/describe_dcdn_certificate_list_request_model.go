// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeDcdnCertificateListRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnCertificateListRequest
	GetSecurityToken() *string
}

type DescribeDcdnCertificateListRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// By default, this operation queries the certificates of all accelerated domain names.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnCertificateListRequest) SetDomainName(v string) *DescribeDcdnCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnCertificateListRequest) SetOwnerId(v int64) *DescribeDcdnCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnCertificateListRequest) SetSecurityToken(v string) *DescribeDcdnCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
