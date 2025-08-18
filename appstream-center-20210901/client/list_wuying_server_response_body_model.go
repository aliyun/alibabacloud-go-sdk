// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListWuyingServerResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWuyingServerResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListWuyingServerResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListWuyingServerResponseBody
	GetTotalCount() *int32
	SetWuyingServerList(v []*ListWuyingServerResponseBodyWuyingServerList) *ListWuyingServerResponseBody
	GetWuyingServerList() []*ListWuyingServerResponseBodyWuyingServerList
}

type ListWuyingServerResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount       *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WuyingServerList []*ListWuyingServerResponseBodyWuyingServerList `json:"WuyingServerList,omitempty" xml:"WuyingServerList,omitempty" type:"Repeated"`
}

func (s ListWuyingServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBody) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWuyingServerResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWuyingServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWuyingServerResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListWuyingServerResponseBody) GetWuyingServerList() []*ListWuyingServerResponseBodyWuyingServerList {
	return s.WuyingServerList
}

func (s *ListWuyingServerResponseBody) SetPageNumber(v int32) *ListWuyingServerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListWuyingServerResponseBody) SetPageSize(v int32) *ListWuyingServerResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWuyingServerResponseBody) SetRequestId(v string) *ListWuyingServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWuyingServerResponseBody) SetTotalCount(v int32) *ListWuyingServerResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWuyingServerResponseBody) SetWuyingServerList(v []*ListWuyingServerResponseBodyWuyingServerList) *ListWuyingServerResponseBody {
	s.WuyingServerList = v
	return s
}

func (s *ListWuyingServerResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListWuyingServerResponseBodyWuyingServerList struct {
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
	// 2025-08-02T16:52:11.000+00:00
	CreateTime *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataDisk   []*ListWuyingServerResponseBodyWuyingServerListDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-09-03T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// imgc-06****oagaev
	ImageId          *string                                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName        *string                                                         `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceInfoList []*ListWuyingServerResponseBodyWuyingServerListInstanceInfoList `json:"InstanceInfoList,omitempty" xml:"InstanceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 10.80.21.149
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-1b****ayv2
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// exampleOfficeSite
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// example:
	//
	// Linux
	OsType                 *string                                                             `json:"OsType,omitempty" xml:"OsType,omitempty"`
	ServerInstanceTypeInfo *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo `json:"ServerInstanceTypeInfo,omitempty" xml:"ServerInstanceTypeInfo,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// example:
	//
	// ws-0byd****8wn2lwi
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
	// example:
	//
	// exampleServerName
	WuyingServerName *string `json:"WuyingServerName,omitempty" xml:"WuyingServerName,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerList) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerList) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetDataDisk() []*ListWuyingServerResponseBodyWuyingServerListDataDisk {
	return s.DataDisk
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetImageId() *string {
	return s.ImageId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetImageName() *string {
	return s.ImageName
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetInstanceInfoList() []*ListWuyingServerResponseBodyWuyingServerListInstanceInfoList {
	return s.InstanceInfoList
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetOsType() *string {
	return s.OsType
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetServerInstanceTypeInfo() *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	return s.ServerInstanceTypeInfo
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetStatus() *string {
	return s.Status
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetWuyingServerName() *string {
	return s.WuyingServerName
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetBizRegionId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.BizRegionId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetChargeType(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ChargeType = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetCreateTime(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.CreateTime = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetDataDisk(v []*ListWuyingServerResponseBodyWuyingServerListDataDisk) *ListWuyingServerResponseBodyWuyingServerList {
	s.DataDisk = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetExpiredTime(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ExpiredTime = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetImageId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ImageId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetImageName(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ImageName = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetInstanceInfoList(v []*ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) *ListWuyingServerResponseBodyWuyingServerList {
	s.InstanceInfoList = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetNetworkInterfaceIp(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetOfficeSiteId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.OfficeSiteId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetOfficeSiteName(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.OfficeSiteName = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetOfficeSiteType(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.OfficeSiteType = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetOsType(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.OsType = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetServerInstanceTypeInfo(v *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) *ListWuyingServerResponseBodyWuyingServerList {
	s.ServerInstanceTypeInfo = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetStatus(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.Status = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSystemDiskCategory(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.SystemDiskCategory = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSystemDiskPerformanceLevel(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSystemDiskSize(v int32) *ListWuyingServerResponseBodyWuyingServerList {
	s.SystemDiskSize = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetWuyingServerId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.WuyingServerId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetWuyingServerName(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.WuyingServerName = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) Validate() error {
	return dara.Validate(s)
}

type ListWuyingServerResponseBodyWuyingServerListDataDisk struct {
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
	// 200
	DataDiskSize *int32 `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerListDataDisk) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerListDataDisk) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) GetDataDiskPerformanceLevel() *string {
	return s.DataDiskPerformanceLevel
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) SetDataDiskCategory(v string) *ListWuyingServerResponseBodyWuyingServerListDataDisk {
	s.DataDiskCategory = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) SetDataDiskPerformanceLevel(v string) *ListWuyingServerResponseBodyWuyingServerListDataDisk {
	s.DataDiskPerformanceLevel = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) SetDataDiskSize(v int32) *ListWuyingServerResponseBodyWuyingServerListDataDisk {
	s.DataDiskSize = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) Validate() error {
	return dara.Validate(s)
}

type ListWuyingServerResponseBodyWuyingServerListInstanceInfoList struct {
	// example:
	//
	// p-0ceitx****c5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// eni-uf65b****dfnt3wb
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) SetInstanceId(v string) *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList {
	s.InstanceId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) SetNetworkInterfaceId(v string) *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList {
	s.NetworkInterfaceId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListInstanceInfoList) Validate() error {
	return dara.Validate(s)
}

type ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo struct {
	// example:
	//
	// 96
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// 4
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// example:
	//
	// 196,608
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// example:
	//
	// 393,216
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// eds.proworkstation_flagship_elite_ne.96c384g.192g4x
	ServerInstanceType *string `json:"ServerInstanceType,omitempty" xml:"ServerInstanceType,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GetCpu() *string {
	return s.Cpu
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GetGpu() *string {
	return s.Gpu
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GetGpuMemory() *int32 {
	return s.GpuMemory
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GetMemory() *int32 {
	return s.Memory
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GetServerInstanceType() *string {
	return s.ServerInstanceType
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) SetCpu(v string) *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	s.Cpu = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) SetGpu(v string) *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	s.Gpu = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) SetGpuMemory(v int32) *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	s.GpuMemory = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) SetMemory(v int32) *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	s.Memory = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) SetServerInstanceType(v string) *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	s.ServerInstanceType = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) Validate() error {
	return dara.Validate(s)
}
