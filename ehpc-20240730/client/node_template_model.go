// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenew(v bool) *NodeTemplate
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *NodeTemplate
	GetAutoRenewPeriod() *int32
	SetDataDisks(v []*NodeTemplateDataDisks) *NodeTemplate
	GetDataDisks() []*NodeTemplateDataDisks
	SetDuration(v int32) *NodeTemplate
	GetDuration() *int32
	SetEnableHT(v bool) *NodeTemplate
	GetEnableHT() *bool
	SetImageId(v string) *NodeTemplate
	GetImageId() *string
	SetInstanceChargeType(v string) *NodeTemplate
	GetInstanceChargeType() *string
	SetInstanceType(v string) *NodeTemplate
	GetInstanceType() *string
	SetPeriod(v int32) *NodeTemplate
	GetPeriod() *int32
	SetPeriodUnit(v string) *NodeTemplate
	GetPeriodUnit() *string
	SetSpotPriceLimit(v float32) *NodeTemplate
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *NodeTemplate
	GetSpotStrategy() *string
	SetSystemDisk(v *NodeTemplateSystemDisk) *NodeTemplate
	GetSystemDisk() *NodeTemplateSystemDisk
}

type NodeTemplate struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	AutoRenewPeriod *int32                   `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	DataDisks       []*NodeTemplateDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// true
	EnableHT *bool `json:"EnableHT,omitempty" xml:"EnableHT,omitempty"`
	// example:
	//
	// aliyun_3_x64_20G_alibase_20221102.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// example:
	//
	// ecs.c7.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
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
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// example:
	//
	// NoSpot
	SpotStrategy *string                 `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SystemDisk   *NodeTemplateSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
}

func (s NodeTemplate) String() string {
	return dara.Prettify(s)
}

func (s NodeTemplate) GoString() string {
	return s.String()
}

func (s *NodeTemplate) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *NodeTemplate) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *NodeTemplate) GetDataDisks() []*NodeTemplateDataDisks {
	return s.DataDisks
}

func (s *NodeTemplate) GetDuration() *int32 {
	return s.Duration
}

func (s *NodeTemplate) GetEnableHT() *bool {
	return s.EnableHT
}

func (s *NodeTemplate) GetImageId() *string {
	return s.ImageId
}

func (s *NodeTemplate) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *NodeTemplate) GetInstanceType() *string {
	return s.InstanceType
}

func (s *NodeTemplate) GetPeriod() *int32 {
	return s.Period
}

func (s *NodeTemplate) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *NodeTemplate) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *NodeTemplate) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *NodeTemplate) GetSystemDisk() *NodeTemplateSystemDisk {
	return s.SystemDisk
}

func (s *NodeTemplate) SetAutoRenew(v bool) *NodeTemplate {
	s.AutoRenew = &v
	return s
}

func (s *NodeTemplate) SetAutoRenewPeriod(v int32) *NodeTemplate {
	s.AutoRenewPeriod = &v
	return s
}

func (s *NodeTemplate) SetDataDisks(v []*NodeTemplateDataDisks) *NodeTemplate {
	s.DataDisks = v
	return s
}

func (s *NodeTemplate) SetDuration(v int32) *NodeTemplate {
	s.Duration = &v
	return s
}

func (s *NodeTemplate) SetEnableHT(v bool) *NodeTemplate {
	s.EnableHT = &v
	return s
}

func (s *NodeTemplate) SetImageId(v string) *NodeTemplate {
	s.ImageId = &v
	return s
}

func (s *NodeTemplate) SetInstanceChargeType(v string) *NodeTemplate {
	s.InstanceChargeType = &v
	return s
}

func (s *NodeTemplate) SetInstanceType(v string) *NodeTemplate {
	s.InstanceType = &v
	return s
}

func (s *NodeTemplate) SetPeriod(v int32) *NodeTemplate {
	s.Period = &v
	return s
}

func (s *NodeTemplate) SetPeriodUnit(v string) *NodeTemplate {
	s.PeriodUnit = &v
	return s
}

func (s *NodeTemplate) SetSpotPriceLimit(v float32) *NodeTemplate {
	s.SpotPriceLimit = &v
	return s
}

func (s *NodeTemplate) SetSpotStrategy(v string) *NodeTemplate {
	s.SpotStrategy = &v
	return s
}

func (s *NodeTemplate) SetSystemDisk(v *NodeTemplateSystemDisk) *NodeTemplate {
	s.SystemDisk = v
	return s
}

func (s *NodeTemplate) Validate() error {
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

type NodeTemplateDataDisks struct {
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
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// example:
	//
	// PL0
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// example:
	//
	// /data1
	MountDir *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// s-bp1ei2b44ripxuo46hym
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s NodeTemplateDataDisks) String() string {
	return dara.Prettify(s)
}

func (s NodeTemplateDataDisks) GoString() string {
	return s.String()
}

func (s *NodeTemplateDataDisks) GetCategory() *string {
	return s.Category
}

func (s *NodeTemplateDataDisks) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *NodeTemplateDataDisks) GetDevice() *string {
	return s.Device
}

func (s *NodeTemplateDataDisks) GetLevel() *string {
	return s.Level
}

func (s *NodeTemplateDataDisks) GetMountDir() *string {
	return s.MountDir
}

func (s *NodeTemplateDataDisks) GetSize() *int32 {
	return s.Size
}

func (s *NodeTemplateDataDisks) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *NodeTemplateDataDisks) SetCategory(v string) *NodeTemplateDataDisks {
	s.Category = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDeleteWithInstance(v bool) *NodeTemplateDataDisks {
	s.DeleteWithInstance = &v
	return s
}

func (s *NodeTemplateDataDisks) SetDevice(v string) *NodeTemplateDataDisks {
	s.Device = &v
	return s
}

func (s *NodeTemplateDataDisks) SetLevel(v string) *NodeTemplateDataDisks {
	s.Level = &v
	return s
}

func (s *NodeTemplateDataDisks) SetMountDir(v string) *NodeTemplateDataDisks {
	s.MountDir = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSize(v int32) *NodeTemplateDataDisks {
	s.Size = &v
	return s
}

func (s *NodeTemplateDataDisks) SetSnapshotId(v string) *NodeTemplateDataDisks {
	s.SnapshotId = &v
	return s
}

func (s *NodeTemplateDataDisks) Validate() error {
	return dara.Validate(s)
}

type NodeTemplateSystemDisk struct {
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

func (s NodeTemplateSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s NodeTemplateSystemDisk) GoString() string {
	return s.String()
}

func (s *NodeTemplateSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *NodeTemplateSystemDisk) GetLevel() *string {
	return s.Level
}

func (s *NodeTemplateSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *NodeTemplateSystemDisk) SetCategory(v string) *NodeTemplateSystemDisk {
	s.Category = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetLevel(v string) *NodeTemplateSystemDisk {
	s.Level = &v
	return s
}

func (s *NodeTemplateSystemDisk) SetSize(v int32) *NodeTemplateSystemDisk {
	s.Size = &v
	return s
}

func (s *NodeTemplateSystemDisk) Validate() error {
	return dara.Validate(s)
}
