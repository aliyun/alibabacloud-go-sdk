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
	// The data cloud disks. If DataDisk.1.Size is empty and this field is not empty, this field takes precedence.
	DataDisks []*DescribePriceRequestDataDisks `json:"DataDisks,omitempty" xml:"DataDisks,omitempty" type:"Repeated"`
	// The edge zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-unicom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ens.sn1.tiny
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The Billable methods for ENS bandwidth. Valid values:
	//
	// - BandwidthByDay: daily peak bandwidth.
	//
	// - 95BandwidthByMonth: monthly 95th percentile bandwidth.
	//
	// - PayByBandwidth4thMonth: monthly fourth peak bandwidth.
	//
	// - PayByBandwidth: bandwidth with defined specifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95BandwidthByMonth
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The subscription duration of the resource.
	//
	// - If PeriodUnit is not specified, the default billing cycle is month. Only Day and Month are supported.
	//
	// - If PeriodUnit is set to Day, Period can only be set to 3.
	//
	// - If PeriodUnit is set to Month, Period can be set to 1 to 9 or 12.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The billing cycle for querying the prices of ENS resources. Valid values:
	//
	// - Month (default): monthly billing cycle.
	//
	// - Day: daily billing cycle.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The quantity.
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
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
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
	if s.DataDisks != nil {
		for _, item := range s.DataDisks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePriceRequestDataDisk struct {
	// The size of the data cloud disk. Unit: GB. If this field is not empty, it takes precedence.
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
	// The size of the system cloud disk. Unit: GB.
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
	// The cloud disk category.
	//
	// example:
	//
	// cloud_efficiency
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The size of the data cloud disk. Unit: GB.
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
