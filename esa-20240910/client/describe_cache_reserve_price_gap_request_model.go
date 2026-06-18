// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCacheReservePriceGapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCacheReservePriceGapRequest
	GetInstanceId() *string
	SetTargetQuotaGb(v int64) *DescribeCacheReservePriceGapRequest
	GetTargetQuotaGb() *int64
}

type DescribeCacheReservePriceGapRequest struct {
	// example:
	//
	// esa-cr-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1000
	TargetQuotaGb *int64 `json:"TargetQuotaGb,omitempty" xml:"TargetQuotaGb,omitempty"`
}

func (s DescribeCacheReservePriceGapRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCacheReservePriceGapRequest) GoString() string {
	return s.String()
}

func (s *DescribeCacheReservePriceGapRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCacheReservePriceGapRequest) GetTargetQuotaGb() *int64 {
	return s.TargetQuotaGb
}

func (s *DescribeCacheReservePriceGapRequest) SetInstanceId(v string) *DescribeCacheReservePriceGapRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCacheReservePriceGapRequest) SetTargetQuotaGb(v int64) *DescribeCacheReservePriceGapRequest {
	s.TargetQuotaGb = &v
	return s
}

func (s *DescribeCacheReservePriceGapRequest) Validate() error {
	return dara.Validate(s)
}
