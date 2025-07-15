// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCertificateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveCertificateListRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveCertificateListRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeLiveCertificateListRequest
	GetSecurityToken() *string
}

type DescribeLiveCertificateListRequest struct {
	// The ingest domain or streaming domain. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// demo.aliyundoc.com,example.com
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLiveCertificateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveCertificateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveCertificateListRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLiveCertificateListRequest) SetDomainName(v string) *DescribeLiveCertificateListRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveCertificateListRequest) SetOwnerId(v int64) *DescribeLiveCertificateListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveCertificateListRequest) SetSecurityToken(v string) *DescribeLiveCertificateListRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLiveCertificateListRequest) Validate() error {
	return dara.Validate(s)
}
