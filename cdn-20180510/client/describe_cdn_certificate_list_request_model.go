// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeCdnCertificateListRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnCertificateListRequest
	GetSecurityToken() *string
}

type DescribeCdnCertificateListRequest struct {
	// The accelerated domain name. Separate multiple accelerated domain names with commas (,).
	//
	// If you do not specify an accelerated domain name, SSL certificates of all your accelerated domain names are queried.
	//
	// example:
	//
	// example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnCertificateListRequest) SetDomainName(v string) *DescribeCdnCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnCertificateListRequest) SetOwnerId(v int64) *DescribeCdnCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnCertificateListRequest) SetSecurityToken(v string) *DescribeCdnCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
