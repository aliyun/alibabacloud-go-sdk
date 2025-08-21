// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDisk(v []*DescribePriceShrinkRequestDataDisk) *DescribePriceShrinkRequest
	GetDataDisk() []*DescribePriceShrinkRequestDataDisk
	SetSystemDisk(v *DescribePriceShrinkRequestSystemDisk) *DescribePriceShrinkRequest
	GetSystemDisk() *DescribePriceShrinkRequestSystemDisk
	SetDataDisksShrink(v string) *DescribePriceShrinkRequest
	GetDataDisksShrink() *string
	SetEnsRegionId(v string) *DescribePriceShrinkRequest
	GetEnsRegionId() *string
	SetInstanceType(v string) *DescribePriceShrinkRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *DescribePriceShrinkRequest
	GetInternetChargeType() *string
	SetPeriod(v int32) *DescribePriceShrinkRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribePriceShrinkRequest
	GetPeriodUnit() *string
	SetQuantity(v int32) *DescribePriceShrinkRequest
	GetQuantity() *int32
}

type DescribePriceShrinkRequest struct {
	DataDisk   []*DescribePriceShrinkRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SystemDisk *DescribePriceShrinkRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// If you leave DataDisk.1.Size empty, the value that you specified for this parameter is used.
	DataDisksShrink *string `json:"DataDisks,omitempty" xml:"DataDisks,omitempty"`
	// The ID of the ENS node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specifications of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ens.sn1.tiny
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The bandwidth metering method of the instance. Valid values:
	//
	// 	- BandwidthByDay: Pay by daily peak bandwidth
	//
	// 	- 95BandwidthByMonth: Pay by monthly 95th percentile bandwidth
	//
	// 	- PayByBandwidth4thMonth: Pay by monthly fourth peak bandwidth
	//
	// 	- PayByBandwidth: Pay by fixed bandwidth
	//
	// This parameter is required.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The subscription duration of the instance.
	//
	// 	- If you leave the PeriodUnit parameter empty, the instance is purchased on a monthly basis. Valid values: Day and Month.
	//
	// 	- If you set PeriodUnit to Day, you can set Period only to 3.
	//
	// 	- If you set PeriodUnit to Month, you can set Period to a number from 1 to 9, or set Period to 12.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle of the ENS instance. Valid values:
	//
	// 	- Month (default):
	//
	// 	- Day
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The number of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s DescribePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceShrinkRequest) GetDataDisk() []*DescribePriceShrinkRequestDataDisk {
	return s.DataDisk
}

func (s *DescribePriceShrinkRequest) GetSystemDisk() *DescribePriceShrinkRequestSystemDisk {
	return s.SystemDisk
}

func (s *DescribePriceShrinkRequest) GetDataDisksShrink() *string {
	return s.DataDisksShrink
}

func (s *DescribePriceShrinkRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribePriceShrinkRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePriceShrinkRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribePriceShrinkRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceShrinkRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribePriceShrinkRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribePriceShrinkRequest) SetDataDisk(v []*DescribePriceShrinkRequestDataDisk) *DescribePriceShrinkRequest {
	s.DataDisk = v
	return s
}

func (s *DescribePriceShrinkRequest) SetSystemDisk(v *DescribePriceShrinkRequestSystemDisk) *DescribePriceShrinkRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribePriceShrinkRequest) SetDataDisksShrink(v string) *DescribePriceShrinkRequest {
	s.DataDisksShrink = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetEnsRegionId(v string) *DescribePriceShrinkRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetInstanceType(v string) *DescribePriceShrinkRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetInternetChargeType(v string) *DescribePriceShrinkRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetPeriod(v int32) *DescribePriceShrinkRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetPeriodUnit(v string) *DescribePriceShrinkRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceShrinkRequest) SetQuantity(v int32) *DescribePriceShrinkRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type DescribePriceShrinkRequestDataDisk struct {
	// The size of the data disk. Unit: GB. If you specify this parameter, this parameter takes precedence over the Size property in DataDisks.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceShrinkRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceShrinkRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceShrinkRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribePriceShrinkRequestDataDisk) SetSize(v int32) *DescribePriceShrinkRequestDataDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceShrinkRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceShrinkRequestSystemDisk struct {
	// The size of the system disk. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceShrinkRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceShrinkRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceShrinkRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribePriceShrinkRequestSystemDisk) SetSize(v int32) *DescribePriceShrinkRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceShrinkRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}
