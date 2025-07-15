// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainCertificateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainCertificateInfoRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainCertificateInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainCertificateInfoRequest
	GetRegionId() *string
}

type DescribeLiveDomainCertificateInfoRequest struct {
	// The streaming domain or ingest domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveDomainCertificateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainCertificateInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainCertificateInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainCertificateInfoRequest) SetDomainName(v string) *DescribeLiveDomainCertificateInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoRequest) SetOwnerId(v int64) *DescribeLiveDomainCertificateInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoRequest) SetRegionId(v string) *DescribeLiveDomainCertificateInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoRequest) Validate() error {
	return dara.Validate(s)
}
