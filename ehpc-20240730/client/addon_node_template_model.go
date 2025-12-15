// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddonNodeTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *AddonNodeTemplate
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *AddonNodeTemplate
	GetAutoRenewPeriod() *int32
	SetDataDisks(v []*AddonNodeTemplateDataDisks) *AddonNodeTemplate
	GetDataDisks() []*AddonNodeTemplateDataDisks
	SetDuration(v int32) *AddonNodeTemplate
	GetDuration() *int32
	SetEnableHT(v bool) *AddonNodeTemplate
	GetEnableHT() *bool
	SetImageId(v string) *AddonNodeTemplate
	GetImageId() *string
	SetInstanceChargeType(v string) *AddonNodeTemplate
	GetInstanceChargeType() *string
	SetInstanceId(v string) *AddonNodeTemplate
	GetInstanceId() *string
	SetInstanceType(v string) *AddonNodeTemplate
	GetInstanceType() *string
	SetOsName(v string) *AddonNodeTemplate
	GetOsName() *string
	SetOsNameEN(v string) *AddonNodeTemplate
	GetOsNameEN() *string
	SetPeriod(v int32) *AddonNodeTemplate
	GetPeriod() *int32
	SetPeriodUnit(v string) *AddonNodeTemplate
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *AddonNodeTemplate
	GetPrivateIpAddress() *string
	SetSpotPriceLimit(v float32) *AddonNodeTemplate
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *AddonNodeTemplate
	GetSpotStrategy() *string
	SetSystemDisk(v *AddonNodeTemplateSystemDisk) *AddonNodeTemplate
	GetSystemDisk() *AddonNodeTemplateSystemDisk
}

type AddonNodeTemplate struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *int32                        `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataDisks       []*AddonNodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aliyun_3_x64_20G_alibase_20221102.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ecs.c7.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CentOS  7.6 64 位
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CentOS  7.6 64 bit
	OsNameEN *string `json:"OsNameEN,omitempty" xml:"OsNameEN,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// Month
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string                      `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk   *AddonNodeTemplateSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s AddonNodeTemplate) String() string {
	return dara.Prettify(s)
}

func (s AddonNodeTemplate) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplate) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *AddonNodeTemplate) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *AddonNodeTemplate) GetDataDisks() []*AddonNodeTemplateDataDisks {
	return s.DataDisks
}

func (s *AddonNodeTemplate) GetDuration() *int32 {
	return s.Duration
}

func (s *AddonNodeTemplate) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *AddonNodeTemplate) GetImageId() *string {
	return s.ImageId
}

func (s *AddonNodeTemplate) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *AddonNodeTemplate) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddonNodeTemplate) GetInstanceType() *string {
	return s.InstanceType
}

func (s *AddonNodeTemplate) GetOsName() *string {
	return s.OsName
}

func (s *AddonNodeTemplate) GetOsNameEN() *string {
	return s.OsNameEN
}

func (s *AddonNodeTemplate) GetPeriod() *int32 {
	return s.Period
}

func (s *AddonNodeTemplate) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *AddonNodeTemplate) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *AddonNodeTemplate) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *AddonNodeTemplate) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *AddonNodeTemplate) GetSystemDisk() *AddonNodeTemplateSystemDisk {
	return s.SystemDisk
}

func (s *AddonNodeTemplate) SetAutoRenew(v bool) *AddonNodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *AddonNodeTemplate) SetAutoRenewPeriod(v int32) *AddonNodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddonNodeTemplate) SetDataDisks(v []*AddonNodeTemplateDataDisks) *AddonNodeTemplate {
	s.DataDisks = v
	return s
}

func (s *AddonNodeTemplate) SetDuration(v int32) *AddonNodeTemplate {
	s.Duration = &v
	return s
}

func (s *AddonNodeTemplate) SetEnableHT(v bool) *AddonNodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *AddonNodeTemplate) SetImageId(v string) *AddonNodeTemplate {
	s.ImageId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceChargeType(v string) *AddonNodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceId(v string) *AddonNodeTemplate {
	s.InstanceId = &v
	return s
}

func (s *AddonNodeTemplate) SetInstanceType(v string) *AddonNodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *AddonNodeTemplate) SetOsName(v string) *AddonNodeTemplate {
	s.OsName = &v
	return s
}

func (s *AddonNodeTemplate) SetOsNameEN(v string) *AddonNodeTemplate {
	s.OsNameEN = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriod(v int32) *AddonNodeTemplate {
	s.Period = &v
	return s
}

func (s *AddonNodeTemplate) SetPeriodUnit(v string) *AddonNodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *AddonNodeTemplate) SetPrivateIpAddress(v string) *AddonNodeTemplate {
	s.PrivateIpAddress = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotPriceLimit(v float32) *AddonNodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *AddonNodeTemplate) SetSpotStrategy(v string) *AddonNodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *AddonNodeTemplate) SetSystemDisk(v *AddonNodeTemplateSystemDisk) *AddonNodeTemplate {
	s.SystemDisk = v
	return s
}

func (s *AddonNodeTemplate) Validate() error {
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddonNodeTemplateDataDisks struct {
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// false
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateDataDisks) String() string {
	return dara.Prettify(s)
}

func (s AddonNodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateDataDisks) GetCategory() *string {
	return s.Category
}

func (s *AddonNodeTemplateDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *AddonNodeTemplateDataDisks) GetLevel() *string {
	return s.Level
}

func (s *AddonNodeTemplateDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *AddonNodeTemplateDataDisks) SetCategory(v string) *AddonNodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetDeleteWithInstance(v bool) *AddonNodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetLevel(v string) *AddonNodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) SetSize(v int32) *AddonNodeTemplateDataDisks {
	s.Size = &v
	return s
}

func (s *AddonNodeTemplateDataDisks) Validate() error {
	return dara.Validate(s)
}

type AddonNodeTemplateSystemDisk struct {
	// example:
	//
	// cloud_auto
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddonNodeTemplateSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s AddonNodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *AddonNodeTemplateSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *AddonNodeTemplateSystemDisk) GetLevel() *string {
	return s.Level
}

func (s *AddonNodeTemplateSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *AddonNodeTemplateSystemDisk) SetCategory(v string) *AddonNodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetLevel(v string) *AddonNodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) SetSize(v int32) *AddonNodeTemplateSystemDisk {
	s.Size = &v
	return s
}

func (s *AddonNodeTemplateSystemDisk) Validate() error {
	return dara.Validate(s)
}
