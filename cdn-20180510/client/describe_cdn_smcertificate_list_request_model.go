// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSMCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnSMCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeCdnSMCertificateListRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnSMCertificateListRequest
	GetSecurityToken() *string
}

type DescribeCdnSMCertificateListRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
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

func (s DescribeCdnSMCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSMCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnSMCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnSMCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnSMCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnSMCertificateListRequest) SetDomainName(v string) *DescribeCdnSMCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnSMCertificateListRequest) SetOwnerId(v int64) *DescribeCdnSMCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnSMCertificateListRequest) SetSecurityToken(v string) *DescribeCdnSMCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnSMCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
