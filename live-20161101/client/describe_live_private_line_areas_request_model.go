// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePrivateLineAreasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeLivePrivateLineAreasRequest
	GetDomainName() *string
	SetOwnerId(v int64) *DescribeLivePrivateLineAreasRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLivePrivateLineAreasRequest
	GetRegionId() *string
}

type DescribeLivePrivateLineAreasRequest struct {
	// The main streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLivePrivateLineAreasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePrivateLineAreasRequest) GoString() string {
	return s.String()
}

func (s *DescribeLivePrivateLineAreasRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLivePrivateLineAreasRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLivePrivateLineAreasRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLivePrivateLineAreasRequest) SetDomainName(v string) *DescribeLivePrivateLineAreasRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeLivePrivateLineAreasRequest) SetOwnerId(v int64) *DescribeLivePrivateLineAreasRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLivePrivateLineAreasRequest) SetRegionId(v string) *DescribeLivePrivateLineAreasRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLivePrivateLineAreasRequest) Validate() error {
	return dara.Validate(s)
}
