// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLimitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLiveDomainLimitRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLiveDomainLimitRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveDomainLimitRequest
	GetRegionId() *string
}

type DescribeLiveDomainLimitRequest struct {
	// The name of the main streaming domain.
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

func (s DescribeLiveDomainLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLimitRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainLimitRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveDomainLimitRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveDomainLimitRequest) SetDomainName(v string) *DescribeLiveDomainLimitRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainLimitRequest) SetOwnerId(v int64) *DescribeLiveDomainLimitRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveDomainLimitRequest) SetRegionId(v string) *DescribeLiveDomainLimitRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveDomainLimitRequest) Validate() error {
	return dara.Validate(s)
}
