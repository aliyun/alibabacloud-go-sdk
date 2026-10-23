// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *DescribeCloudPhoneNodePriceRequest
	GetBizRegionId() *string
	SetChargeType(v string) *DescribeCloudPhoneNodePriceRequest
	GetChargeType() *string
	SetCount(v string) *DescribeCloudPhoneNodePriceRequest
	GetCount() *string
	SetIsSingleImgDisk(v bool) *DescribeCloudPhoneNodePriceRequest
	GetIsSingleImgDisk() *bool
	SetPeriod(v string) *DescribeCloudPhoneNodePriceRequest
	GetPeriod() *string
	SetPeriodUnit(v string) *DescribeCloudPhoneNodePriceRequest
	GetPeriodUnit() *string
	SetPhoneCount(v int32) *DescribeCloudPhoneNodePriceRequest
	GetPhoneCount() *int32
	SetPhoneDataVolume(v int32) *DescribeCloudPhoneNodePriceRequest
	GetPhoneDataVolume() *int32
	SetPromotionId(v string) *DescribeCloudPhoneNodePriceRequest
	GetPromotionId() *string
	SetServerShareDataVolume(v int32) *DescribeCloudPhoneNodePriceRequest
	GetServerShareDataVolume() *int32
	SetServerType(v string) *DescribeCloudPhoneNodePriceRequest
	GetServerType() *string
	SetSwapSize(v int32) *DescribeCloudPhoneNodePriceRequest
	GetSwapSize() *int32
}

type DescribeCloudPhoneNodePriceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// false
	IsSingleImgDisk *bool `json:"IsSingleImgDisk,omitempty" xml:"IsSingleImgDisk,omitempty"`
	// example:
	//
	// 1
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 40
	PhoneCount *int32 `json:"PhoneCount,omitempty" xml:"PhoneCount,omitempty"`
	// example:
	//
	// 20
	PhoneDataVolume *int32 `json:"PhoneDataVolume,omitempty" xml:"PhoneDataVolume,omitempty"`
	// example:
	//
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// example:
	//
	// 200
	ServerShareDataVolume *int32 `json:"ServerShareDataVolume,omitempty" xml:"ServerShareDataVolume,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cpm.gx7.10xlarge
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// example:
	//
	// 10
	SwapSize *int32 `json:"SwapSize,omitempty" xml:"SwapSize,omitempty"`
}

func (s DescribeCloudPhoneNodePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeCloudPhoneNodePriceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCloudPhoneNodePriceRequest) GetCount() *string {
	return s.Count
}

func (s *DescribeCloudPhoneNodePriceRequest) GetIsSingleImgDisk() *bool {
	return s.IsSingleImgDisk
}

func (s *DescribeCloudPhoneNodePriceRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeCloudPhoneNodePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeCloudPhoneNodePriceRequest) GetPhoneCount() *int32 {
	return s.PhoneCount
}

func (s *DescribeCloudPhoneNodePriceRequest) GetPhoneDataVolume() *int32 {
	return s.PhoneDataVolume
}

func (s *DescribeCloudPhoneNodePriceRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCloudPhoneNodePriceRequest) GetServerShareDataVolume() *int32 {
	return s.ServerShareDataVolume
}

func (s *DescribeCloudPhoneNodePriceRequest) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeCloudPhoneNodePriceRequest) GetSwapSize() *int32 {
	return s.SwapSize
}

func (s *DescribeCloudPhoneNodePriceRequest) SetBizRegionId(v string) *DescribeCloudPhoneNodePriceRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetChargeType(v string) *DescribeCloudPhoneNodePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetCount(v string) *DescribeCloudPhoneNodePriceRequest {
	s.Count = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetIsSingleImgDisk(v bool) *DescribeCloudPhoneNodePriceRequest {
	s.IsSingleImgDisk = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetPeriod(v string) *DescribeCloudPhoneNodePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetPeriodUnit(v string) *DescribeCloudPhoneNodePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetPhoneCount(v int32) *DescribeCloudPhoneNodePriceRequest {
	s.PhoneCount = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetPhoneDataVolume(v int32) *DescribeCloudPhoneNodePriceRequest {
	s.PhoneDataVolume = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetPromotionId(v string) *DescribeCloudPhoneNodePriceRequest {
	s.PromotionId = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetServerShareDataVolume(v int32) *DescribeCloudPhoneNodePriceRequest {
	s.ServerShareDataVolume = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetServerType(v string) *DescribeCloudPhoneNodePriceRequest {
	s.ServerType = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) SetSwapSize(v int32) *DescribeCloudPhoneNodePriceRequest {
	s.SwapSize = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceRequest) Validate() error {
	return dara.Validate(s)
}
