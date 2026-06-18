// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheReservePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrRegion(v string) *DescribeCacheReservePriceRequest
	GetCrRegion() *string
	SetPeriod(v int32) *DescribeCacheReservePriceRequest
	GetPeriod() *int32
	SetQuotaGb(v int64) *DescribeCacheReservePriceRequest
	GetQuotaGb() *int64
}

type DescribeCacheReservePriceRequest struct {
	// The cache reserve region.
	//
	// - HK: Hong Kong (China)
	//
	// - CN: the Chinese mainland.
	//
	// example:
	//
	// HK
	CrRegion *string `json:"CrRegion,omitempty" xml:"CrRegion,omitempty"`
	// The purchase period. Unit: months.
	//
	// example:
	//
	// 2
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The cache reserve specification. Unit: GB.
	//
	// example:
	//
	// 512000
	QuotaGb *int64 `json:"QuotaGb,omitempty" xml:"QuotaGb,omitempty"`
}

func (s DescribeCacheReservePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceRequest) GetCrRegion() *string {
	return s.CrRegion
}

func (s *DescribeCacheReservePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeCacheReservePriceRequest) GetQuotaGb() *int64 {
	return s.QuotaGb
}

func (s *DescribeCacheReservePriceRequest) SetCrRegion(v string) *DescribeCacheReservePriceRequest {
	s.CrRegion = &v
	return s
}

func (s *DescribeCacheReservePriceRequest) SetPeriod(v int32) *DescribeCacheReservePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeCacheReservePriceRequest) SetQuotaGb(v int64) *DescribeCacheReservePriceRequest {
	s.QuotaGb = &v
	return s
}

func (s *DescribeCacheReservePriceRequest) Validate() error {
	return dara.Validate(s)
}
