// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest
	GetDataDisk() []*DescribePriceRequestDataDisk
	SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest
	GetSystemDisk() *DescribePriceRequestSystemDisk
	SetDataDisks(v []*DescribePriceRequestDataDisks) *DescribePriceRequest
	GetDataDisks() []*DescribePriceRequestDataDisks
	SetEnsRegionId(v string) *DescribePriceRequest
	GetEnsRegionId() *string
	SetInstanceType(v string) *DescribePriceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *DescribePriceRequest
	GetInternetChargeType() *string
	SetPeriod(v int32) *DescribePriceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *DescribePriceRequest
	GetPeriodUnit() *string
	SetQuantity(v int32) *DescribePriceRequest
	GetQuantity() *int32
}

type DescribePriceRequest struct {
	DataDisk   []*DescribePriceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	SystemDisk *DescribePriceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// If you leave DataDisk.1.Size empty, the value that you specified for this parameter is used.
	DataDisks []*DescribePriceRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
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

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetDataDisk() []*DescribePriceRequestDataDisk {
	return s.DataDisk
}

func (s *DescribePriceRequest) GetSystemDisk() *DescribePriceRequestSystemDisk {
	return s.SystemDisk
}

func (s *DescribePriceRequest) GetDataDisks() []*DescribePriceRequestDataDisks {
	return s.DataDisks
}

func (s *DescribePriceRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribePriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribePriceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribePriceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribePriceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribePriceRequest) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribePriceRequest) SetDataDisk(v []*DescribePriceRequestDataDisk) *DescribePriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribePriceRequest) SetSystemDisk(v *DescribePriceRequestSystemDisk) *DescribePriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribePriceRequest) SetDataDisks(v []*DescribePriceRequestDataDisks) *DescribePriceRequest {
	s.DataDisks = v
	return s
}

func (s *DescribePriceRequest) SetEnsRegionId(v string) *DescribePriceRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribePriceRequest) SetInstanceType(v string) *DescribePriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequest) SetInternetChargeType(v string) *DescribePriceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetPeriod(v int32) *DescribePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribePriceRequest) SetPeriodUnit(v string) *DescribePriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *DescribePriceRequest) SetQuantity(v int32) *DescribePriceRequest {
	s.Quantity = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestDataDisk struct {
	// The size of the data disk. Unit: GB. If you specify this parameter, this parameter takes precedence over the Size property in DataDisks.
	//
	// example:
	//
	// 50
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribePriceRequestDataDisk) SetSize(v int32) *DescribePriceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestSystemDisk struct {
	// The size of the system disk. Unit: GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribePriceRequestSystemDisk) SetSize(v int32) *DescribePriceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribePriceRequestDataDisks struct {
	// The category of the disk.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The size of the data disk. Unit: GB.
	//
	// example:
	//
	// 50
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribePriceRequestDataDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequestDataDisks) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestDataDisks) GetCategory() *string {
	return s.Category
}

func (s *DescribePriceRequestDataDisks) GetSize() *int64 {
	return s.Size
}

func (s *DescribePriceRequestDataDisks) SetCategory(v string) *DescribePriceRequestDataDisks {
	s.Category = &v
	return s
}

func (s *DescribePriceRequestDataDisks) SetSize(v int64) *DescribePriceRequestDataDisks {
	s.Size = &v
	return s
}

func (s *DescribePriceRequestDataDisks) Validate() error {
	return dara.Validate(s)
}
