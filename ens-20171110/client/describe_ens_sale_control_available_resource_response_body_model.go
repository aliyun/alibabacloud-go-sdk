// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsSaleControlAvailableResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEnsSaleControlAvailableResourceResponseBody
	GetRequestId() *string
	SetSaleControlAvailableResource(v []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) *DescribeEnsSaleControlAvailableResourceResponseBody
	GetSaleControlAvailableResource() []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource
}

type DescribeEnsSaleControlAvailableResourceResponseBody struct {
	RequestId                    *string                                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SaleControlAvailableResource []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource `json:"SaleControlAvailableResource,omitempty" xml:"SaleControlAvailableResource,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBody) GetSaleControlAvailableResource() []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	return s.SaleControlAvailableResource
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBody) SetRequestId(v string) *DescribeEnsSaleControlAvailableResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBody) SetSaleControlAvailableResource(v []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) *DescribeEnsSaleControlAvailableResourceResponseBody {
	s.SaleControlAvailableResource = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBody) Validate() error {
	if s.SaleControlAvailableResource != nil {
		for _, item := range s.SaleControlAvailableResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource struct {
	AvailableDiskType    []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType  `json:"AvailableDiskType,omitempty" xml:"AvailableDiskType,omitempty" type:"Repeated"`
	AvailableRegion      []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion    `json:"AvailableRegion,omitempty" xml:"AvailableRegion,omitempty" type:"Repeated"`
	AvailableSpec        []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec      `json:"AvailableSpec,omitempty" xml:"AvailableSpec,omitempty" type:"Repeated"`
	AvailableStorageType *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType `json:"AvailableStorageType,omitempty" xml:"AvailableStorageType,omitempty" type:"Struct"`
	CommodityCode        *string                                                                                              `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	OrderType            *string                                                                                              `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GetAvailableDiskType() []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType {
	return s.AvailableDiskType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GetAvailableRegion() []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	return s.AvailableRegion
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GetAvailableSpec() []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec {
	return s.AvailableSpec
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GetAvailableStorageType() *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType {
	return s.AvailableStorageType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) SetAvailableDiskType(v []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	s.AvailableDiskType = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) SetAvailableRegion(v []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	s.AvailableRegion = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) SetAvailableSpec(v []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	s.AvailableSpec = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) SetAvailableStorageType(v *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	s.AvailableStorageType = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) SetCommodityCode(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	s.CommodityCode = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) SetOrderType(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource {
	s.OrderType = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResource) Validate() error {
	if s.AvailableDiskType != nil {
		for _, item := range s.AvailableDiskType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AvailableRegion != nil {
		for _, item := range s.AvailableRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AvailableSpec != nil {
		for _, item := range s.AvailableSpec {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AvailableStorageType != nil {
		if err := s.AvailableStorageType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType struct {
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) SetDiskName(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType {
	s.DiskName = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) SetDiskType(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType {
	s.DiskType = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableDiskType) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion struct {
	Area          *string `json:"Area,omitempty" xml:"Area,omitempty"`
	City          *string `json:"City,omitempty" xml:"City,omitempty"`
	Country       *string `json:"Country,omitempty" xml:"Country,omitempty"`
	EnsRegionId   *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	EnsRegionName *string `json:"EnsRegionName,omitempty" xml:"EnsRegionName,omitempty"`
	Isp           *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Province      *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetArea() *string {
	return s.Area
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetCity() *string {
	return s.City
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetCountry() *string {
	return s.Country
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetEnsRegionName() *string {
	return s.EnsRegionName
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetIsp() *string {
	return s.Isp
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) GetProvince() *string {
	return s.Province
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetArea(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.Area = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetCity(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.City = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetCountry(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.Country = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetEnsRegionId(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetEnsRegionName(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.EnsRegionName = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetIsp(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.Isp = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) SetProvince(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion {
	s.Province = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableRegion) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec struct {
	Cores     *string `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Memory    *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	SpecName  *string `json:"SpecName,omitempty" xml:"SpecName,omitempty"`
	SpecValue *string `json:"SpecValue,omitempty" xml:"SpecValue,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) GetCores() *string {
	return s.Cores
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) GetMemory() *string {
	return s.Memory
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) GetSpecName() *string {
	return s.SpecName
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) GetSpecValue() *string {
	return s.SpecValue
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) SetCores(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec {
	s.Cores = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) SetMemory(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec {
	s.Memory = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) SetSpecName(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec {
	s.SpecName = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) SetSpecValue(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec {
	s.SpecValue = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType struct {
	AvailableDefaultStorageType []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType   `json:"AvailableDefaultStorageType,omitempty" xml:"AvailableDefaultStorageType,omitempty" type:"Repeated"`
	AvailableSpecialStorageType [][]*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType `json:"AvailableSpecialStorageType,omitempty" xml:"AvailableSpecialStorageType,omitempty" type:"Repeated"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) GetAvailableDefaultStorageType() []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType {
	return s.AvailableDefaultStorageType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) GetAvailableSpecialStorageType() [][]*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType {
	return s.AvailableSpecialStorageType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) SetAvailableDefaultStorageType(v []*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType {
	s.AvailableDefaultStorageType = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) SetAvailableSpecialStorageType(v [][]*DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType {
	s.AvailableSpecialStorageType = v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageType) Validate() error {
	if s.AvailableDefaultStorageType != nil {
		for _, item := range s.AvailableDefaultStorageType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType struct {
	StorageName *string `json:"StorageName,omitempty" xml:"StorageName,omitempty"`
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) GetStorageName() *string {
	return s.StorageName
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) SetStorageName(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType {
	s.StorageName = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) SetStorageType(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType {
	s.StorageType = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableDefaultStorageType) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType struct {
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StorageName *string `json:"StorageName,omitempty" xml:"StorageName,omitempty"`
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) GoString() string {
	return s.String()
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) GetStorageName() *string {
	return s.StorageName
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) SetStorageType(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType {
	s.StorageType = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) SetStorageName(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType {
	s.StorageName = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) SetEnsRegionId(v string) *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsSaleControlAvailableResourceResponseBodySaleControlAvailableResourceAvailableStorageTypeAvailableSpecialStorageType) Validate() error {
	return dara.Validate(s)
}
