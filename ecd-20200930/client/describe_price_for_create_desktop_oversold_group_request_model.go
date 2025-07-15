// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceForCreateDesktopOversoldGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrenceCount(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetConcurrenceCount() *int32
	SetDataDiskSize(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetDataDiskSize() *int32
	SetDesktopType(v string) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetDesktopType() *string
	SetOversoldUserCount(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetOversoldUserCount() *int32
	SetPeriod(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetPeriodUnit() *string
	SetSystemDiskSize(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest
	GetSystemDiskSize() *int32
}

type DescribePriceForCreateDesktopOversoldGroupRequest struct {
	ConcurrenceCount  *int32  `json:"ConcurrenceCount,omitempty" xml:"ConcurrenceCount,omitempty"`
	DataDiskSize      *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopType       *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	OversoldUserCount *int32  `json:"OversoldUserCount,omitempty" xml:"OversoldUserCount,omitempty"`
	Period            *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit        *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	SystemDiskSize    *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribePriceForCreateDesktopOversoldGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceForCreateDesktopOversoldGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetConcurrenceCount() *int32 {
	return s.ConcurrenceCount
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetOversoldUserCount() *int32 {
	return s.OversoldUserCount
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetConcurrenceCount(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.ConcurrenceCount = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetDataDiskSize(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.DataDiskSize = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetDesktopType(v string) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.DesktopType = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetOversoldUserCount(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.OversoldUserCount = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetPeriod(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetPeriodUnit(v string) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) SetSystemDiskSize(v int32) *DescribePriceForCreateDesktopOversoldGroupRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribePriceForCreateDesktopOversoldGroupRequest) Validate() error {
	return dara.Validate(s)
}
