// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetDetailByUuidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetDetail(v *DescribeAssetDetailByUuidResponseBodyAssetDetail) *DescribeAssetDetailByUuidResponseBody
	GetAssetDetail() *DescribeAssetDetailByUuidResponseBodyAssetDetail
	SetRequestId(v string) *DescribeAssetDetailByUuidResponseBody
	GetRequestId() *string
}

type DescribeAssetDetailByUuidResponseBody struct {
	// The details of the server.
	AssetDetail *DescribeAssetDetailByUuidResponseBodyAssetDetail `json:"AssetDetail,omitempty" xml:"AssetDetail,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 39031E31-6BBA-5C99-A870-D807E78918CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssetDetailByUuidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidResponseBody) GetAssetDetail() *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	return s.AssetDetail
}

func (s *DescribeAssetDetailByUuidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetDetailByUuidResponseBody) SetAssetDetail(v *DescribeAssetDetailByUuidResponseBodyAssetDetail) *DescribeAssetDetailByUuidResponseBody {
	s.AssetDetail = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBody) SetRequestId(v string) *DescribeAssetDetailByUuidResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBody) Validate() error {
	if s.AssetDetail != nil {
		if err := s.AssetDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAssetDetailByUuidResponseBodyAssetDetail struct {
	// The type of the asset. Valid values:
	//
	// 	- **0**: ECS instance
	//
	// 	- **1**: Server Load Balancer (SLB) instance
	//
	// 	- **2**: NAT gateway
	//
	// 	- **3**: ApsaraDB RDS database
	//
	// 	- **4**: ApsaraDB for MongoDB database
	//
	// 	- **5**: ApsaraDB for Redis database
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
	// The edition of Security Center that is authorized to protect the asset. Valid values:
	//
	// 	- **1**: Basic edition (Unauthorized)
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **7**: Ultimate edition
	//
	// example:
	//
	// 7
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Indicates whether Security Center is authorized to protect the asset. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The status of the Security Center agent. Valid values:
	//
	// 	- **pause**: The Security Center agent suspends protection for your server.
	//
	// 	- **online**: The Security Center agent is protecting your server.
	//
	// 	- **offline**: The Security Center agent does not protect your server.
	//
	// example:
	//
	// online
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
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
	// An array that consists of the information about the disk.
	DiskInfoList []*string `json:"DiskInfoList,omitempty" xml:"DiskInfoList,omitempty" type:"Repeated"`
	// Indicates whether the asset is provided by Alibaba Cloud. Valid values:
	//
	// 	- **0**: yes
	//
	// 	- **1**: no
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
	// 192.168.XX.XX
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
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP addresses of the server.
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The kernel version of the operating system.
	//
	// example:
	//
	// 4.18.0-80.11.2.el8_0.x86_64
	Kernel *string `json:"Kernel,omitempty" xml:"Kernel,omitempty"`
	// The media access control (MAC) addresses of the server.
	MacList []*string `json:"MacList,omitempty" xml:"MacList,omitempty" type:"Repeated"`
	// The memory size of the server. Unit: GB.
	//
	// example:
	//
	// 32
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The memory size of the server. Unit: MB.
	//
	// example:
	//
	// 512
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
	// The name of the operating system.
	//
	// example:
	//
	// -
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou-dg-a01
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
	// The ID of the virtual private cloud (VPC) in which the server resides.
	//
	// example:
	//
	// vpc-bp1fs3bwonlfq503w****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeAssetDetailByUuidResponseBodyAssetDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidResponseBodyAssetDetail) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetAuthModifyTime() *int64 {
	return s.AuthModifyTime
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetBind() *bool {
	return s.Bind
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetDiskInfoList() []*string {
	return s.DiskInfoList
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetFlag() *int32 {
	return s.Flag
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetGroupTrace() *string {
	return s.GroupTrace
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetHostName() *string {
	return s.HostName
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetIp() *string {
	return s.Ip
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetIpList() []*string {
	return s.IpList
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetKernel() *string {
	return s.Kernel
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetMacList() []*string {
	return s.MacList
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetMem() *int32 {
	return s.Mem
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetOs() *string {
	return s.Os
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetOsDetail() *string {
	return s.OsDetail
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetOsName() *string {
	return s.OsName
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetRegion() *string {
	return s.Region
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetSysInfo() *string {
	return s.SysInfo
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetTag() *string {
	return s.Tag
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetAssetType(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.AssetType = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetAuthModifyTime(v int64) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.AuthModifyTime = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetAuthVersion(v int32) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.AuthVersion = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetBind(v bool) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Bind = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetClientStatus(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.ClientStatus = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetClientVersion(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.ClientVersion = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetCpu(v int32) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Cpu = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetCpuInfo(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.CpuInfo = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetCreateTime(v int64) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.CreateTime = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetDiskInfoList(v []*string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.DiskInfoList = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetFlag(v int32) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Flag = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetGroupTrace(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.GroupTrace = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetHostName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.HostName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInstanceId(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInstanceName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InstanceName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetInternetIp(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.InternetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetIntranetIp(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetIp(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Ip = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetIpList(v []*string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.IpList = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetKernel(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Kernel = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetMacList(v []*string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.MacList = v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetMem(v int32) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Mem = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetMemory(v int64) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Memory = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetOs(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Os = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetOsDetail(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.OsDetail = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetOsName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.OsName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetRegion(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Region = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetRegionId(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetRegionName(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.RegionName = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetSysInfo(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.SysInfo = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetTag(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Tag = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetUuid(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) SetVpcInstanceId(v string) *DescribeAssetDetailByUuidResponseBodyAssetDetail {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidResponseBodyAssetDetail) Validate() error {
	return dara.Validate(s)
}
