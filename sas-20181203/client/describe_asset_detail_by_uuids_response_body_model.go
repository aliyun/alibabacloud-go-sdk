// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetDetailByUuidsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*DescribeAssetDetailByUuidsResponseBodyAssetList) *DescribeAssetDetailByUuidsResponseBody
	GetAssetList() []*DescribeAssetDetailByUuidsResponseBodyAssetList
	SetRequestId(v string) *DescribeAssetDetailByUuidsResponseBody
	GetRequestId() *string
}

type DescribeAssetDetailByUuidsResponseBody struct {
	// An array that consists of the details of the instances.
	AssetList []*DescribeAssetDetailByUuidsResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 92016EC8-D52D-49D8-9FF7-9EA340A950B9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssetDetailByUuidsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsResponseBody) GetAssetList() []*DescribeAssetDetailByUuidsResponseBodyAssetList {
	return s.AssetList
}

func (s *DescribeAssetDetailByUuidsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetDetailByUuidsResponseBody) SetAssetList(v []*DescribeAssetDetailByUuidsResponseBodyAssetList) *DescribeAssetDetailByUuidsResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBody) SetRequestId(v string) *DescribeAssetDetailByUuidsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBody) Validate() error {
	if s.AssetList != nil {
		for _, item := range s.AssetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAssetDetailByUuidsResponseBodyAssetList struct {
	// The type of the asset.
	//
	// The value is fixed as **0**, which indicates ECS instances.
	//
	// example:
	//
	// 0
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The timestamp when Security Center is authorized to protect the instance. Unit: milliseconds.
	//
	// example:
	//
	// 1627974044000
	AuthModifyTime *int64 `json:"AuthModifyTime,omitempty" xml:"AuthModifyTime,omitempty"`
	// The edition of Security Center that is authorized to protect the instance. Valid values:
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
	// Indicates whether Security Center is authorized to protect the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The status of the Security Center agent. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
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
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The details of the CPU.
	//
	// example:
	//
	// Intel(R) Xeon(R) Platinum 8163 CPU @ 2.50GHz
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// The timestamp when Security Center records the details of the instance. Unit: milliseconds.
	//
	// example:
	//
	// 1603863599000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// An array that consists of the information about the disk.
	DiskInfoList []*string `json:"DiskInfoList,omitempty" xml:"DiskInfoList,omitempty" type:"Repeated"`
	// The type of the asset by source. Valid values:
	//
	// 	- **0**: The asset is provided by Alibaba Cloud.
	//
	// 	- **1**: The asset is not provided by Alibaba Cloud.
	//
	// 	- **2**: The asset resides in a data center.
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset.
	//
	// 	- **8**: light-weight assets.
	//
	// example:
	//
	// 0
	Flag *int32 `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// The group to which the instance belongs. By default, the instances that are not grouped belong to the **Default*	- group.
	//
	// example:
	//
	// default
	GroupTrace *string `json:"GroupTrace,omitempty" xml:"GroupTrace,omitempty"`
	// The hostname.
	//
	// example:
	//
	// test
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-rj9gda4wolo0zixi****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the ECS instance.
	//
	// example:
	//
	// TestInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the ECS instance.
	//
	// example:
	//
	// 10.10.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the ECS instance.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address of the ECS instance.
	//
	// >  If the ECS instance has a public IP address, the value of this parameter is the public IP address of the ECS instance. If the ECS instance does not have a public IP address, the value of this parameter is the private IP address of the ECS instance.
	//
	// example:
	//
	// 10.10.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP addresses of the instances.
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The kernel version of the operating system.
	//
	// example:
	//
	// 4.18.0-80.11.2.el8_0.x86_64
	Kernel *string `json:"Kernel,omitempty" xml:"Kernel,omitempty"`
	// The media access control (MAC) addresses of the instances.
	MacList []*string `json:"MacList,omitempty" xml:"MacList,omitempty" type:"Repeated"`
	// The memory size of the instance. Unit: GB.
	//
	// example:
	//
	// 4
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The memory size of the instance. Unit: MB.
	//
	// example:
	//
	// 1024
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The operating system of the ECS instance.
	//
	// example:
	//
	// Linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The operating system version of the instance.
	//
	// example:
	//
	// Linux 64bit
	OsDetail *string `json:"OsDetail,omitempty" xml:"OsDetail,omitempty"`
	// The name of the operating system run by the ECS instance.
	//
	// example:
	//
	// CentOS 7.6 64-bit
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The region in which the ECS instance resides.
	//
	// example:
	//
	// cn-guangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region in which the ECS instance resides.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region in which the ECS instance resides.
	//
	// example:
	//
	// cn-shenzhen
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The operating system information about the instance.
	//
	// example:
	//
	// CentOS Linux 8.0.1905
	SysInfo *string `json:"SysInfo,omitempty" xml:"SysInfo,omitempty"`
	// The tag added to the instance.
	//
	// example:
	//
	// test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The UUID of the ECS instance.
	//
	// example:
	//
	// 2a98f149-0256-414c-a29a-a69f8a75****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// 13231-331331
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeAssetDetailByUuidsResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetDetailByUuidsResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetAuthModifyTime() *int64 {
	return s.AuthModifyTime
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetBind() *bool {
	return s.Bind
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetDiskInfoList() []*string {
	return s.DiskInfoList
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetFlag() *int32 {
	return s.Flag
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetGroupTrace() *string {
	return s.GroupTrace
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetHostName() *string {
	return s.HostName
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetIp() *string {
	return s.Ip
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetIpList() []*string {
	return s.IpList
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetKernel() *string {
	return s.Kernel
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetMacList() []*string {
	return s.MacList
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetMem() *int32 {
	return s.Mem
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetOs() *string {
	return s.Os
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetOsDetail() *string {
	return s.OsDetail
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetOsName() *string {
	return s.OsName
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetRegion() *string {
	return s.Region
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetSysInfo() *string {
	return s.SysInfo
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetTag() *string {
	return s.Tag
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetAssetType(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.AssetType = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetAuthModifyTime(v int64) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.AuthModifyTime = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetAuthVersion(v int32) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.AuthVersion = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetBind(v bool) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Bind = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetClientStatus(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.ClientStatus = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetClientVersion(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.ClientVersion = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetCpu(v int32) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Cpu = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetCpuInfo(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.CpuInfo = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetCreateTime(v int64) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.CreateTime = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetDiskInfoList(v []*string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.DiskInfoList = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetFlag(v int32) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Flag = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetGroupTrace(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.GroupTrace = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetHostName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.HostName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetInstanceId(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.InstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetInstanceName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.InstanceName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetInternetIp(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.InternetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetIntranetIp(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetIp(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Ip = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetIpList(v []*string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.IpList = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetKernel(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Kernel = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetMacList(v []*string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.MacList = v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetMem(v int32) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Mem = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetMemory(v int64) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Memory = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetOs(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Os = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetOsDetail(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.OsDetail = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetOsName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.OsName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetRegion(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Region = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetRegionId(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.RegionId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetRegionName(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.RegionName = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetSysInfo(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.SysInfo = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetTag(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Tag = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetUuid(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.Uuid = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) SetVpcInstanceId(v string) *DescribeAssetDetailByUuidsResponseBodyAssetList {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeAssetDetailByUuidsResponseBodyAssetList) Validate() error {
	return dara.Validate(s)
}
