// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeRegionResourceResponseBodyData) *DescribeRegionResourceResponseBody
	GetData() []*DescribeRegionResourceResponseBodyData
	SetDesc(v string) *DescribeRegionResourceResponseBody
	GetDesc() *string
	SetMsg(v string) *DescribeRegionResourceResponseBody
	GetMsg() *string
	SetPager(v *DescribeRegionResourceResponseBodyPager) *DescribeRegionResourceResponseBody
	GetPager() *DescribeRegionResourceResponseBodyPager
	SetRequestId(v string) *DescribeRegionResourceResponseBody
	GetRequestId() *string
}

type DescribeRegionResourceResponseBody struct {
	Data      []*DescribeRegionResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Desc      *string                                   `json:"Desc,omitempty" xml:"Desc,omitempty"`
	Msg       *string                                   `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Pager     *DescribeRegionResourceResponseBodyPager  `json:"Pager,omitempty" xml:"Pager,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBody) GetData() []*DescribeRegionResourceResponseBodyData {
	return s.Data
}

func (s *DescribeRegionResourceResponseBody) GetDesc() *string {
	return s.Desc
}

func (s *DescribeRegionResourceResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeRegionResourceResponseBody) GetPager() *DescribeRegionResourceResponseBodyPager {
	return s.Pager
}

func (s *DescribeRegionResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionResourceResponseBody) SetData(v []*DescribeRegionResourceResponseBodyData) *DescribeRegionResourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRegionResourceResponseBody) SetDesc(v string) *DescribeRegionResourceResponseBody {
	s.Desc = &v
	return s
}

func (s *DescribeRegionResourceResponseBody) SetMsg(v string) *DescribeRegionResourceResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeRegionResourceResponseBody) SetPager(v *DescribeRegionResourceResponseBodyPager) *DescribeRegionResourceResponseBody {
	s.Pager = v
	return s
}

func (s *DescribeRegionResourceResponseBody) SetRequestId(v string) *DescribeRegionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionResourceResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pager != nil {
		if err := s.Pager.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionResourceResponseBodyData struct {
	AreaCode       *string                                             `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	AreaName       *string                                             `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	ArmCard        *DescribeRegionResourceResponseBodyDataArmCard      `json:"ArmCard,omitempty" xml:"ArmCard,omitempty" type:"Struct"`
	Attributes     []*string                                           `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Bandwidth      *DescribeRegionResourceResponseBodyDataBandwidth    `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" type:"Struct"`
	BlockStorage   *DescribeRegionResourceResponseBodyDataBlockStorage `json:"BlockStorage,omitempty" xml:"BlockStorage,omitempty" type:"Struct"`
	CountryCode    *string                                             `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CountryName    *string                                             `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	Cpu            *DescribeRegionResourceResponseBodyDataCpu          `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	Gpu            *DescribeRegionResourceResponseBodyDataGpu          `json:"Gpu,omitempty" xml:"Gpu,omitempty" type:"Struct"`
	Hdd            *DescribeRegionResourceResponseBodyDataHdd          `json:"Hdd,omitempty" xml:"Hdd,omitempty" type:"Struct"`
	HouseId        *string                                             `json:"HouseId,omitempty" xml:"HouseId,omitempty"`
	Ipv4s          []*DescribeRegionResourceResponseBodyDataIpv4s      `json:"Ipv4s,omitempty" xml:"Ipv4s,omitempty" type:"Repeated"`
	Ipv6s          []*DescribeRegionResourceResponseBodyDataIpv6s      `json:"Ipv6s,omitempty" xml:"Ipv6s,omitempty" type:"Repeated"`
	IspTypes       []*string                                           `json:"IspTypes,omitempty" xml:"IspTypes,omitempty" type:"Repeated"`
	Memory         *DescribeRegionResourceResponseBodyDataMemory       `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	Name           *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Nvme           *DescribeRegionResourceResponseBodyDataNvme         `json:"Nvme,omitempty" xml:"Nvme,omitempty" type:"Struct"`
	OssStorage     *DescribeRegionResourceResponseBodyDataOssStorage   `json:"OssStorage,omitempty" xml:"OssStorage,omitempty" type:"Struct"`
	Pangu          *DescribeRegionResourceResponseBodyDataPangu        `json:"Pangu,omitempty" xml:"Pangu,omitempty" type:"Struct"`
	PcfarmNum      *DescribeRegionResourceResponseBodyDataPcfarmNum    `json:"PcfarmNum,omitempty" xml:"PcfarmNum,omitempty" type:"Struct"`
	Poc            *bool                                               `json:"Poc,omitempty" xml:"Poc,omitempty"`
	ProvinceCode   *string                                             `json:"ProvinceCode,omitempty" xml:"ProvinceCode,omitempty"`
	ProvinceName   *string                                             `json:"ProvinceName,omitempty" xml:"ProvinceName,omitempty"`
	ReserveDisable *bool                                               `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	Ssd            *DescribeRegionResourceResponseBodyDataSsd          `json:"Ssd,omitempty" xml:"Ssd,omitempty" type:"Struct"`
	StatusDisable  *bool                                               `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	Type           *string                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	Uuid           *string                                             `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	Virtual        *string                                             `json:"Virtual,omitempty" xml:"Virtual,omitempty"`
}

func (s DescribeRegionResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyData) GetAreaCode() *string {
	return s.AreaCode
}

func (s *DescribeRegionResourceResponseBodyData) GetAreaName() *string {
	return s.AreaName
}

func (s *DescribeRegionResourceResponseBodyData) GetArmCard() *DescribeRegionResourceResponseBodyDataArmCard {
	return s.ArmCard
}

func (s *DescribeRegionResourceResponseBodyData) GetAttributes() []*string {
	return s.Attributes
}

func (s *DescribeRegionResourceResponseBodyData) GetBandwidth() *DescribeRegionResourceResponseBodyDataBandwidth {
	return s.Bandwidth
}

func (s *DescribeRegionResourceResponseBodyData) GetBlockStorage() *DescribeRegionResourceResponseBodyDataBlockStorage {
	return s.BlockStorage
}

func (s *DescribeRegionResourceResponseBodyData) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeRegionResourceResponseBodyData) GetCountryName() *string {
	return s.CountryName
}

func (s *DescribeRegionResourceResponseBodyData) GetCpu() *DescribeRegionResourceResponseBodyDataCpu {
	return s.Cpu
}

func (s *DescribeRegionResourceResponseBodyData) GetGpu() *DescribeRegionResourceResponseBodyDataGpu {
	return s.Gpu
}

func (s *DescribeRegionResourceResponseBodyData) GetHdd() *DescribeRegionResourceResponseBodyDataHdd {
	return s.Hdd
}

func (s *DescribeRegionResourceResponseBodyData) GetHouseId() *string {
	return s.HouseId
}

func (s *DescribeRegionResourceResponseBodyData) GetIpv4s() []*DescribeRegionResourceResponseBodyDataIpv4s {
	return s.Ipv4s
}

func (s *DescribeRegionResourceResponseBodyData) GetIpv6s() []*DescribeRegionResourceResponseBodyDataIpv6s {
	return s.Ipv6s
}

func (s *DescribeRegionResourceResponseBodyData) GetIspTypes() []*string {
	return s.IspTypes
}

func (s *DescribeRegionResourceResponseBodyData) GetMemory() *DescribeRegionResourceResponseBodyDataMemory {
	return s.Memory
}

func (s *DescribeRegionResourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeRegionResourceResponseBodyData) GetNvme() *DescribeRegionResourceResponseBodyDataNvme {
	return s.Nvme
}

func (s *DescribeRegionResourceResponseBodyData) GetOssStorage() *DescribeRegionResourceResponseBodyDataOssStorage {
	return s.OssStorage
}

func (s *DescribeRegionResourceResponseBodyData) GetPangu() *DescribeRegionResourceResponseBodyDataPangu {
	return s.Pangu
}

func (s *DescribeRegionResourceResponseBodyData) GetPcfarmNum() *DescribeRegionResourceResponseBodyDataPcfarmNum {
	return s.PcfarmNum
}

func (s *DescribeRegionResourceResponseBodyData) GetPoc() *bool {
	return s.Poc
}

func (s *DescribeRegionResourceResponseBodyData) GetProvinceCode() *string {
	return s.ProvinceCode
}

func (s *DescribeRegionResourceResponseBodyData) GetProvinceName() *string {
	return s.ProvinceName
}

func (s *DescribeRegionResourceResponseBodyData) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyData) GetSsd() *DescribeRegionResourceResponseBodyDataSsd {
	return s.Ssd
}

func (s *DescribeRegionResourceResponseBodyData) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeRegionResourceResponseBodyData) GetVirtual() *string {
	return s.Virtual
}

func (s *DescribeRegionResourceResponseBodyData) SetAreaCode(v string) *DescribeRegionResourceResponseBodyData {
	s.AreaCode = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetAreaName(v string) *DescribeRegionResourceResponseBodyData {
	s.AreaName = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetArmCard(v *DescribeRegionResourceResponseBodyDataArmCard) *DescribeRegionResourceResponseBodyData {
	s.ArmCard = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetAttributes(v []*string) *DescribeRegionResourceResponseBodyData {
	s.Attributes = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetBandwidth(v *DescribeRegionResourceResponseBodyDataBandwidth) *DescribeRegionResourceResponseBodyData {
	s.Bandwidth = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetBlockStorage(v *DescribeRegionResourceResponseBodyDataBlockStorage) *DescribeRegionResourceResponseBodyData {
	s.BlockStorage = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetCountryCode(v string) *DescribeRegionResourceResponseBodyData {
	s.CountryCode = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetCountryName(v string) *DescribeRegionResourceResponseBodyData {
	s.CountryName = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetCpu(v *DescribeRegionResourceResponseBodyDataCpu) *DescribeRegionResourceResponseBodyData {
	s.Cpu = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetGpu(v *DescribeRegionResourceResponseBodyDataGpu) *DescribeRegionResourceResponseBodyData {
	s.Gpu = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetHdd(v *DescribeRegionResourceResponseBodyDataHdd) *DescribeRegionResourceResponseBodyData {
	s.Hdd = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetHouseId(v string) *DescribeRegionResourceResponseBodyData {
	s.HouseId = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetIpv4s(v []*DescribeRegionResourceResponseBodyDataIpv4s) *DescribeRegionResourceResponseBodyData {
	s.Ipv4s = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetIpv6s(v []*DescribeRegionResourceResponseBodyDataIpv6s) *DescribeRegionResourceResponseBodyData {
	s.Ipv6s = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetIspTypes(v []*string) *DescribeRegionResourceResponseBodyData {
	s.IspTypes = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetMemory(v *DescribeRegionResourceResponseBodyDataMemory) *DescribeRegionResourceResponseBodyData {
	s.Memory = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetName(v string) *DescribeRegionResourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetNvme(v *DescribeRegionResourceResponseBodyDataNvme) *DescribeRegionResourceResponseBodyData {
	s.Nvme = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetOssStorage(v *DescribeRegionResourceResponseBodyDataOssStorage) *DescribeRegionResourceResponseBodyData {
	s.OssStorage = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetPangu(v *DescribeRegionResourceResponseBodyDataPangu) *DescribeRegionResourceResponseBodyData {
	s.Pangu = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetPcfarmNum(v *DescribeRegionResourceResponseBodyDataPcfarmNum) *DescribeRegionResourceResponseBodyData {
	s.PcfarmNum = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetPoc(v bool) *DescribeRegionResourceResponseBodyData {
	s.Poc = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetProvinceCode(v string) *DescribeRegionResourceResponseBodyData {
	s.ProvinceCode = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetProvinceName(v string) *DescribeRegionResourceResponseBodyData {
	s.ProvinceName = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyData {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetSsd(v *DescribeRegionResourceResponseBodyDataSsd) *DescribeRegionResourceResponseBodyData {
	s.Ssd = v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyData {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetType(v string) *DescribeRegionResourceResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetUuid(v string) *DescribeRegionResourceResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) SetVirtual(v string) *DescribeRegionResourceResponseBodyData {
	s.Virtual = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyData) Validate() error {
	if s.ArmCard != nil {
		if err := s.ArmCard.Validate(); err != nil {
			return err
		}
	}
	if s.Bandwidth != nil {
		if err := s.Bandwidth.Validate(); err != nil {
			return err
		}
	}
	if s.BlockStorage != nil {
		if err := s.BlockStorage.Validate(); err != nil {
			return err
		}
	}
	if s.Cpu != nil {
		if err := s.Cpu.Validate(); err != nil {
			return err
		}
	}
	if s.Gpu != nil {
		if err := s.Gpu.Validate(); err != nil {
			return err
		}
	}
	if s.Hdd != nil {
		if err := s.Hdd.Validate(); err != nil {
			return err
		}
	}
	if s.Ipv4s != nil {
		for _, item := range s.Ipv4s {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Ipv6s != nil {
		for _, item := range s.Ipv6s {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Memory != nil {
		if err := s.Memory.Validate(); err != nil {
			return err
		}
	}
	if s.Nvme != nil {
		if err := s.Nvme.Validate(); err != nil {
			return err
		}
	}
	if s.OssStorage != nil {
		if err := s.OssStorage.Validate(); err != nil {
			return err
		}
	}
	if s.Pangu != nil {
		if err := s.Pangu.Validate(); err != nil {
			return err
		}
	}
	if s.PcfarmNum != nil {
		if err := s.PcfarmNum.Validate(); err != nil {
			return err
		}
	}
	if s.Ssd != nil {
		if err := s.Ssd.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionResourceResponseBodyDataArmCard struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataArmCard) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataArmCard) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataArmCard {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataArmCard {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataArmCard {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetType(v string) *DescribeRegionResourceResponseBodyDataArmCard {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataArmCard {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataArmCard) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataBandwidth struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataBandwidth) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataBandwidth) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetType(v string) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataBandwidth {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBandwidth) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataBlockStorage struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataBlockStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataBlockStorage) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetType(v string) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataBlockStorage {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataBlockStorage) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataCpu struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataCpu) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataCpu) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataCpu) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataCpu {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataCpu {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataCpu {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetType(v string) *DescribeRegionResourceResponseBodyDataCpu {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataCpu {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataCpu) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataGpu struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataGpu) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataGpu) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataGpu) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataGpu {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataGpu {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataGpu {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetType(v string) *DescribeRegionResourceResponseBodyDataGpu {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataGpu {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataGpu) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataHdd struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataHdd) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataHdd) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataHdd) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataHdd {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataHdd {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataHdd {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetType(v string) *DescribeRegionResourceResponseBodyDataHdd {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataHdd {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataHdd) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataIpv4s struct {
	Display             *string `json:"Display,omitempty" xml:"Display,omitempty"`
	Isp                 *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
	Vlan                *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataIpv4s) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataIpv4s) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetDisplay() *string {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetIsp() *string {
	return s.Isp
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) GetVlan() *string {
	return s.Vlan
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetDisplay(v string) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetIsp(v string) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Isp = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetType(v string) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) SetVlan(v string) *DescribeRegionResourceResponseBodyDataIpv4s {
	s.Vlan = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv4s) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataIpv6s struct {
	Display             *string `json:"Display,omitempty" xml:"Display,omitempty"`
	Isp                 *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
	Vlan                *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataIpv6s) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataIpv6s) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetDisplay() *string {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetIsp() *string {
	return s.Isp
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) GetVlan() *string {
	return s.Vlan
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetDisplay(v string) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetIsp(v string) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Isp = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetType(v string) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) SetVlan(v string) *DescribeRegionResourceResponseBodyDataIpv6s {
	s.Vlan = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataIpv6s) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataMemory struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataMemory) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataMemory) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataMemory) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataMemory {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataMemory {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataMemory {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetType(v string) *DescribeRegionResourceResponseBodyDataMemory {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataMemory {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataMemory) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataNvme struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataNvme) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataNvme) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataNvme) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataNvme {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataNvme {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataNvme {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetType(v string) *DescribeRegionResourceResponseBodyDataNvme {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataNvme {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataNvme) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataOssStorage struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataOssStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataOssStorage) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetType(v string) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataOssStorage {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataOssStorage) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataPangu struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataPangu) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataPangu) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataPangu) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataPangu {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataPangu {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataPangu {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetType(v string) *DescribeRegionResourceResponseBodyDataPangu {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataPangu {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPangu) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataPcfarmNum struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataPcfarmNum) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataPcfarmNum) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetType(v string) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataPcfarmNum {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataPcfarmNum) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyDataSsd struct {
	Display             *bool   `json:"Display,omitempty" xml:"Display,omitempty"`
	OversellRatio       *int64  `json:"OversellRatio,omitempty" xml:"OversellRatio,omitempty"`
	Remain              *int64  `json:"Remain,omitempty" xml:"Remain,omitempty"`
	ReserveDisable      *bool   `json:"ReserveDisable,omitempty" xml:"ReserveDisable,omitempty"`
	ReserveDisableTotal *int64  `json:"ReserveDisableTotal,omitempty" xml:"ReserveDisableTotal,omitempty"`
	Reserved            *int64  `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	StatusDisable       *bool   `json:"StatusDisable,omitempty" xml:"StatusDisable,omitempty"`
	StatusDisableTotal  *int64  `json:"StatusDisableTotal,omitempty" xml:"StatusDisableTotal,omitempty"`
	Total               *int64  `json:"Total,omitempty" xml:"Total,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Used                *int64  `json:"Used,omitempty" xml:"Used,omitempty"`
	UsedRatio           *int64  `json:"UsedRatio,omitempty" xml:"UsedRatio,omitempty"`
}

func (s DescribeRegionResourceResponseBodyDataSsd) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyDataSsd) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetDisplay() *bool {
	return s.Display
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetOversellRatio() *int64 {
	return s.OversellRatio
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetRemain() *int64 {
	return s.Remain
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetReserveDisable() *bool {
	return s.ReserveDisable
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetReserveDisableTotal() *int64 {
	return s.ReserveDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetReserved() *int64 {
	return s.Reserved
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetStatusDisable() *bool {
	return s.StatusDisable
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetStatusDisableTotal() *int64 {
	return s.StatusDisableTotal
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetType() *string {
	return s.Type
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetUsed() *int64 {
	return s.Used
}

func (s *DescribeRegionResourceResponseBodyDataSsd) GetUsedRatio() *int64 {
	return s.UsedRatio
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetDisplay(v bool) *DescribeRegionResourceResponseBodyDataSsd {
	s.Display = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetOversellRatio(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.OversellRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetRemain(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.Remain = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetReserveDisable(v bool) *DescribeRegionResourceResponseBodyDataSsd {
	s.ReserveDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetReserveDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.ReserveDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetReserved(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.Reserved = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetStatusDisable(v bool) *DescribeRegionResourceResponseBodyDataSsd {
	s.StatusDisable = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetStatusDisableTotal(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.StatusDisableTotal = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetTotal(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetType(v string) *DescribeRegionResourceResponseBodyDataSsd {
	s.Type = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetUsed(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.Used = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) SetUsedRatio(v int64) *DescribeRegionResourceResponseBodyDataSsd {
	s.UsedRatio = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyDataSsd) Validate() error {
	return dara.Validate(s)
}

type DescribeRegionResourceResponseBodyPager struct {
	Page  *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	Size  *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeRegionResourceResponseBodyPager) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionResourceResponseBodyPager) GoString() string {
	return s.String()
}

func (s *DescribeRegionResourceResponseBodyPager) GetPage() *int64 {
	return s.Page
}

func (s *DescribeRegionResourceResponseBodyPager) GetSize() *int64 {
	return s.Size
}

func (s *DescribeRegionResourceResponseBodyPager) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeRegionResourceResponseBodyPager) SetPage(v int64) *DescribeRegionResourceResponseBodyPager {
	s.Page = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyPager) SetSize(v int64) *DescribeRegionResourceResponseBodyPager {
	s.Size = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyPager) SetTotal(v int64) *DescribeRegionResourceResponseBodyPager {
	s.Total = &v
	return s
}

func (s *DescribeRegionResourceResponseBodyPager) Validate() error {
	return dara.Validate(s)
}
