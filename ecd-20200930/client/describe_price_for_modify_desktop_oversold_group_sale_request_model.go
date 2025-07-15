// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForModifyDesktopOversoldGroupSaleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrenceCount(v int32) *DescribePriceForModifyDesktopOversoldGroupSaleRequest
	GetConcurrenceCount() *int32
	SetOversoldGroupId(v string) *DescribePriceForModifyDesktopOversoldGroupSaleRequest
	GetOversoldGroupId() *string
	SetOversoldUserCount(v int32) *DescribePriceForModifyDesktopOversoldGroupSaleRequest
	GetOversoldUserCount() *int32
}

type DescribePriceForModifyDesktopOversoldGroupSaleRequest struct {
	ConcurrenceCount  *int32  `json:"ConcurrenceCount,omitempty" xml:"ConcurrenceCount,omitempty"`
	OversoldGroupId   *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	OversoldUserCount *int32  `json:"OversoldUserCount,omitempty" xml:"OversoldUserCount,omitempty"`
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForModifyDesktopOversoldGroupSaleRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) GetConcurrenceCount() *int32 {
	return s.ConcurrenceCount
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) GetOversoldUserCount() *int32 {
	return s.OversoldUserCount
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) SetConcurrenceCount(v int32) *DescribePriceForModifyDesktopOversoldGroupSaleRequest {
	s.ConcurrenceCount = &v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) SetOversoldGroupId(v string) *DescribePriceForModifyDesktopOversoldGroupSaleRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) SetOversoldUserCount(v int32) *DescribePriceForModifyDesktopOversoldGroupSaleRequest {
	s.OversoldUserCount = &v
	return s
}

func (s *DescribePriceForModifyDesktopOversoldGroupSaleRequest) Validate() error {
	return dara.Validate(s)
}
