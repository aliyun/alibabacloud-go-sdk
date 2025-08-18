// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int32) *CreateWuyingServerRequest
	GetAmount() *int32
	SetAutoPay(v bool) *CreateWuyingServerRequest
	GetAutoPay() *bool
	SetAutoRenew(v bool) *CreateWuyingServerRequest
	GetAutoRenew() *bool
	SetBizRegionId(v string) *CreateWuyingServerRequest
	GetBizRegionId() *string
	SetChargeType(v string) *CreateWuyingServerRequest
	GetChargeType() *string
	SetDataDisk(v []*CreateWuyingServerRequestDataDisk) *CreateWuyingServerRequest
	GetDataDisk() []*CreateWuyingServerRequestDataDisk
	SetImageId(v string) *CreateWuyingServerRequest
	GetImageId() *string
	SetOfficeSiteId(v string) *CreateWuyingServerRequest
	GetOfficeSiteId() *string
	SetPassword(v string) *CreateWuyingServerRequest
	GetPassword() *string
	SetPeriod(v int32) *CreateWuyingServerRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateWuyingServerRequest
	GetPeriodUnit() *string
	SetPromotionId(v string) *CreateWuyingServerRequest
	GetPromotionId() *string
	SetServerInstanceType(v string) *CreateWuyingServerRequest
	GetServerInstanceType() *string
	SetSystemDiskCategory(v string) *CreateWuyingServerRequest
	GetSystemDiskCategory() *string
	SetSystemDiskPerformanceLevel(v string) *CreateWuyingServerRequest
	GetSystemDiskPerformanceLevel() *string
	SetSystemDiskSize(v int32) *CreateWuyingServerRequest
	GetSystemDiskSize() *int32
	SetVSwitchIds(v []*string) *CreateWuyingServerRequest
	GetVSwitchIds() []*string
	SetWuyingServerName(v string) *CreateWuyingServerRequest
	GetWuyingServerName() *string
}

type CreateWuyingServerRequest struct {
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// example:
	//
	// PrePaid
	ChargeType *string                              `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DataDisk   []*CreateWuyingServerRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// example:
	//
	// img-bp13mu****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-643067****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// YourPassword123
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// example:
	//
	// eds.proworkstation_flagship_elite_ne.96c384g.192g4x
	ServerInstanceType *string `json:"ServerInstanceType,omitempty" xml:"ServerInstanceType,omitempty"`
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// example:
	//
	// PL0
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// example:
	//
	// 100
	SystemDiskSize *int32    `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	VSwitchIds     []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// exampleServerName
	WuyingServerName *string `json:"WuyingServerName,omitempty" xml:"WuyingServerName,omitempty"`
}

func (s CreateWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *CreateWuyingServerRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *CreateWuyingServerRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateWuyingServerRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateWuyingServerRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateWuyingServerRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateWuyingServerRequest) GetDataDisk() []*CreateWuyingServerRequestDataDisk {
	return s.DataDisk
}

func (s *CreateWuyingServerRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateWuyingServerRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateWuyingServerRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateWuyingServerRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateWuyingServerRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateWuyingServerRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *CreateWuyingServerRequest) GetServerInstanceType() *string {
	return s.ServerInstanceType
}

func (s *CreateWuyingServerRequest) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *CreateWuyingServerRequest) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateWuyingServerRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateWuyingServerRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateWuyingServerRequest) GetWuyingServerName() *string {
	return s.WuyingServerName
}

func (s *CreateWuyingServerRequest) SetAmount(v int32) *CreateWuyingServerRequest {
	s.Amount = &v
	return s
}

func (s *CreateWuyingServerRequest) SetAutoPay(v bool) *CreateWuyingServerRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateWuyingServerRequest) SetAutoRenew(v bool) *CreateWuyingServerRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateWuyingServerRequest) SetBizRegionId(v string) *CreateWuyingServerRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateWuyingServerRequest) SetChargeType(v string) *CreateWuyingServerRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateWuyingServerRequest) SetDataDisk(v []*CreateWuyingServerRequestDataDisk) *CreateWuyingServerRequest {
	s.DataDisk = v
	return s
}

func (s *CreateWuyingServerRequest) SetImageId(v string) *CreateWuyingServerRequest {
	s.ImageId = &v
	return s
}

func (s *CreateWuyingServerRequest) SetOfficeSiteId(v string) *CreateWuyingServerRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateWuyingServerRequest) SetPassword(v string) *CreateWuyingServerRequest {
	s.Password = &v
	return s
}

func (s *CreateWuyingServerRequest) SetPeriod(v int32) *CreateWuyingServerRequest {
	s.Period = &v
	return s
}

func (s *CreateWuyingServerRequest) SetPeriodUnit(v string) *CreateWuyingServerRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateWuyingServerRequest) SetPromotionId(v string) *CreateWuyingServerRequest {
	s.PromotionId = &v
	return s
}

func (s *CreateWuyingServerRequest) SetServerInstanceType(v string) *CreateWuyingServerRequest {
	s.ServerInstanceType = &v
	return s
}

func (s *CreateWuyingServerRequest) SetSystemDiskCategory(v string) *CreateWuyingServerRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateWuyingServerRequest) SetSystemDiskPerformanceLevel(v string) *CreateWuyingServerRequest {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateWuyingServerRequest) SetSystemDiskSize(v int32) *CreateWuyingServerRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateWuyingServerRequest) SetVSwitchIds(v []*string) *CreateWuyingServerRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateWuyingServerRequest) SetWuyingServerName(v string) *CreateWuyingServerRequest {
	s.WuyingServerName = &v
	return s
}

func (s *CreateWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}

type CreateWuyingServerRequestDataDisk struct {
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// example:
	//
	// PL0
	DataDiskPerformanceLevel *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	// example:
	//
	// 100
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s CreateWuyingServerRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateWuyingServerRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateWuyingServerRequestDataDisk) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *CreateWuyingServerRequestDataDisk) GetDataDiskPerformanceLevel() *string {
	return s.DataDiskPerformanceLevel
}

func (s *CreateWuyingServerRequestDataDisk) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *CreateWuyingServerRequestDataDisk) SetDataDiskCategory(v string) *CreateWuyingServerRequestDataDisk {
	s.DataDiskCategory = &v
	return s
}

func (s *CreateWuyingServerRequestDataDisk) SetDataDiskPerformanceLevel(v string) *CreateWuyingServerRequestDataDisk {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *CreateWuyingServerRequestDataDisk) SetDataDiskSize(v int32) *CreateWuyingServerRequestDataDisk {
	s.DataDiskSize = &v
	return s
}

func (s *CreateWuyingServerRequestDataDisk) Validate() error {
	return dara.Validate(s)
}
