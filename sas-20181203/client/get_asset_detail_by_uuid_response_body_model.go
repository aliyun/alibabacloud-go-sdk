// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetDetailByUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetDetail(v *GetAssetDetailByUuidResponseBodyAssetDetail) *GetAssetDetailByUuidResponseBody
	GetAssetDetail() *GetAssetDetailByUuidResponseBodyAssetDetail
	SetRequestId(v string) *GetAssetDetailByUuidResponseBody
	GetRequestId() *string
}

type GetAssetDetailByUuidResponseBody struct {
	// The details of the server.
	AssetDetail *GetAssetDetailByUuidResponseBodyAssetDetail `json:"AssetDetail,omitempty" xml:"AssetDetail,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4892B68B-47BC-5E56-B327-9C2ACC6C1C09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAssetDetailByUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAssetDetailByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *GetAssetDetailByUuidResponseBody) GetAssetDetail() *GetAssetDetailByUuidResponseBodyAssetDetail {
	return s.AssetDetail
}

func (s *GetAssetDetailByUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAssetDetailByUuidResponseBody) SetAssetDetail(v *GetAssetDetailByUuidResponseBodyAssetDetail) *GetAssetDetailByUuidResponseBody {
	s.AssetDetail = v
	return s
}

func (s *GetAssetDetailByUuidResponseBody) SetRequestId(v string) *GetAssetDetailByUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAssetDetailByUuidResponseBodyAssetDetail struct {
	// The type of the server. Valid values:
	//
	// 	- **0**: ECS instance
	//
	// 	- **1**: Server Load Balancer (SLB) instance
	//
	// 	- **2**: NAT gateway
	//
	// 	- **3**: ApsaraDB RDS instance
	//
	// 	- **4**: ApsaraDB for MongoDB instance
	//
	// 	- **5**: ApsaraDB for Redis instance
	//
	// 	- **6**: image
	//
	// 	- **7**: container
	//
	// example:
	//
	// 0
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The timestamp when Security Center is authorized to protect the asset. Unit: milliseconds.
	//
	// example:
	//
	// 1627974044000
	AuthModifyTime *int64 `json:"AuthModifyTime,omitempty" xml:"AuthModifyTime,omitempty"`
	// The edition of Security Center that is authorized to protect the server. Valid values:
	//
	// 	- **1**: Basic (Unauthorized).
	//
	// 	- **6**: Anti-virus.
	//
	// 	- **5**: Advanced.
	//
	// 	- **3**: Enterprise.
	//
	// 	- **7**: Ultimate.
	//
	// example:
	//
	// 7
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Indicates whether Security Center is authorized to protect the asset. Valid values:
	//
	// 	- **true**: Security Center is authorized to protect the asset.
	//
	// 	- **false**: Security Center is not authorized to protect the asset.
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The status of the Security Center agent. Valid values:
	//
	// 	- **pause**: The Security Center agent stops protecting your server.
	//
	// 	- **online**: The Security Center agent is protecting your server.
	//
	// 	- **offline**: The Security Center agent does not protect your server.
	//
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The sub-status of the Security Center agent on the server. Valid values:
	//
	// 	- **online**: The Security Center agent on the asset is **enabled**.
	//
	// 	- **offline**: The Security Center agent on the asset is **disabled**.
	//
	// 	- **pause**: The Security Center agent is **suspended**.
	//
	// 	- **uninstalled**: The Security Center agent is **not installed**.
	//
	// 	- **stopped**: The asset is **shut down**.
	//
	// example:
	//
	// online
	ClientSubStatus *string `json:"ClientSubStatus,omitempty" xml:"ClientSubStatus,omitempty"`
	// The version of the Security Center agent.
	//
	// example:
	//
	// 2.0.0
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The details of the CPU.
	//
	// example:
	//
	// Intel(R) Xeon(R) Platinum 8163 CPU @ 2.50GHz
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// The timestamp when Security Center records the details of the server. Unit: milliseconds.
	//
	// example:
	//
	// 1603863599000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of information about the disk.
	DiskInfoList []*GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList `json:"DiskInfoList,omitempty" xml:"DiskInfoList,omitempty" type:"Repeated"`
	// Indicates whether the asset is provided by Alibaba Cloud. Valid values:
	//
	// 	- **0**: The server is provided by Alibaba Cloud.
	//
	// 	- **1**: The server is not provided by Alibaba Cloud.
	//
	// example:
	//
	// 0
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The group to which the server belongs. By default, the servers that are not grouped belong to the **Default*	- group.
	//
	// example:
	//
	// default
	GroupTrace *string `json:"GroupTrace,omitempty" xml:"GroupTrace,omitempty"`
	// The name of the host.
	//
	// example:
	//
	// qewrqwerqs****
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-uf6h7p2fgk6rkk0g****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// i-fasdfasdfadfafa****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 120.47.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address that is assigned to the Elastic Compute Service (ECS) instance.
	//
	// example:
	//
	// 120.47.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// An array that consists of the IP addresses of the server.
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The kernel version of the operating system.
	//
	// example:
	//
	// 4.18.0-80.11.2.el8_0.x86_64
	Kernel *string `json:"Kernel,omitempty" xml:"Kernel,omitempty"`
	// An array that consists of the media access control (MAC) addresses of the server.
	MacList []*string `json:"MacList,omitempty" xml:"MacList,omitempty" type:"Repeated"`
	// The memory size of the server. Unit: GB.
	//
	// example:
	//
	// 16
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The memory size of the server. Unit: MB.
	//
	// example:
	//
	// 16384
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The operating system type of the server.
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The operating system version of the server.
	//
	// example:
	//
	// Linux 64bit
	OsDetail *string `json:"OsDetail,omitempty" xml:"OsDetail,omitempty"`
	// The name of the operating system that the server runs.
	//
	// example:
	//
	// CentOS  7.4 64bit
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the region in which the asset resides.
	//
	// example:
	//
	// cn-hanghzou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region in which the server resides.
	//
	// example:
	//
	// China (Hohhot)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The operating system information about the server.
	//
	// example:
	//
	// CentOS Linux 8.0.1905
	SysInfo *string `json:"SysInfo,omitempty" xml:"SysInfo,omitempty"`
	// The tag that is added to the server.
	//
	// example:
	//
	// InternetIp
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 9e6cad93-a379-46fd-a701-9bbf02f4****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The account to which the AccessKey pair belongs.
	//
	// >  This parameter is returned only by third-party cloud servers. If the parameter value is empty, it will not be returned.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the server resides.
	//
	// example:
	//
	// vpc-bp1fs3bwonlfq503w****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s GetAssetDetailByUuidResponseBodyAssetDetail) String() string {
	return dara.Prettify(s)
}

func (s GetAssetDetailByUuidResponseBodyAssetDetail) GoString() string {
	return s.String()
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetAssetType() *string {
	return s.AssetType
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetAuthModifyTime() *int64 {
	return s.AuthModifyTime
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetBind() *bool {
	return s.Bind
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetClientSubStatus() *string {
	return s.ClientSubStatus
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetDiskInfoList() []*GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList {
	return s.DiskInfoList
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetFlag() *int32 {
	return s.Flag
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetGroupTrace() *string {
	return s.GroupTrace
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetHostName() *string {
	return s.HostName
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetIp() *string {
	return s.Ip
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetIpList() []*string {
	return s.IpList
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetKernel() *string {
	return s.Kernel
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetMacList() []*string {
	return s.MacList
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetMem() *int32 {
	return s.Mem
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetMemory() *int64 {
	return s.Memory
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetOs() *string {
	return s.Os
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetOsDetail() *string {
	return s.OsDetail
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetOsName() *string {
	return s.OsName
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetRegion() *string {
	return s.Region
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetRegionName() *string {
	return s.RegionName
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetSysInfo() *string {
	return s.SysInfo
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetTag() *string {
	return s.Tag
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetUuid() *string {
	return s.Uuid
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetAssetType(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.AssetType = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetAuthModifyTime(v int64) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.AuthModifyTime = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetAuthVersion(v int32) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.AuthVersion = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetBind(v bool) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Bind = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetClientStatus(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.ClientStatus = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetClientSubStatus(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.ClientSubStatus = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetClientVersion(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.ClientVersion = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetCpu(v int32) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Cpu = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetCpuInfo(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.CpuInfo = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetCreateTime(v int64) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.CreateTime = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetDiskInfoList(v []*GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.DiskInfoList = v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetFlag(v int32) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Flag = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetGroupTrace(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.GroupTrace = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetHostName(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.HostName = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetInstanceId(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceId = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetInstanceName(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceName = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetInternetIp(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.InternetIp = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetIntranetIp(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.IntranetIp = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetIp(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Ip = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetIpList(v []*string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.IpList = v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetKernel(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Kernel = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetMacList(v []*string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.MacList = v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetMem(v int32) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Mem = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetMemory(v int64) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Memory = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetOs(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Os = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetOsDetail(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.OsDetail = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetOsName(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.OsName = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetRegion(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Region = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetRegionId(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.RegionId = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetRegionName(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.RegionName = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetSysInfo(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.SysInfo = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetTag(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Tag = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetUuid(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.Uuid = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetVendorAuthAlias(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.VendorAuthAlias = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) SetVpcInstanceId(v string) *GetAssetDetailByUuidResponseBodyAssetDetail {
	s.VpcInstanceId = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetail) Validate() error {
	return dara.Validate(s)
}

type GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList struct {
	// The name of the disk.
	//
	// example:
	//
	// /dev/vda1
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The total disk space. Unit: GB.
	//
	// example:
	//
	// 40
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
	// The total disk space. Unit: bytes.
	//
	// example:
	//
	// 42140479488
	TotalSizeByte *int64 `json:"TotalSizeByte,omitempty" xml:"TotalSizeByte,omitempty"`
	// The amount of the used disk space. Unit: GB.
	//
	// example:
	//
	// 2
	UseSize *int64 `json:"UseSize,omitempty" xml:"UseSize,omitempty"`
	// The amount of the used disk space. Unit: bytes.
	//
	// example:
	//
	// 2998996992
	UseSizeByte *int64 `json:"UseSizeByte,omitempty" xml:"UseSizeByte,omitempty"`
}

func (s GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) GoString() string {
	return s.String()
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) GetDiskName() *string {
	return s.DiskName
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) GetTotalSizeByte() *int64 {
	return s.TotalSizeByte
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) GetUseSize() *int64 {
	return s.UseSize
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) GetUseSizeByte() *int64 {
	return s.UseSizeByte
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) SetDiskName(v string) *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList {
	s.DiskName = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) SetTotalSize(v int64) *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList {
	s.TotalSize = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) SetTotalSizeByte(v int64) *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList {
	s.TotalSizeByte = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) SetUseSize(v int64) *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList {
	s.UseSize = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) SetUseSizeByte(v int64) *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList {
	s.UseSizeByte = &v
	return s
}

func (s *GetAssetDetailByUuidResponseBodyAssetDetailDiskInfoList) Validate() error {
	return dara.Validate(s)
}
