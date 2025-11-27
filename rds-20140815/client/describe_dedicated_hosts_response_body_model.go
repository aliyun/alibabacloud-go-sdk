// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBody
	GetDedicatedHostGroupId() *string
	SetDedicatedHosts(v *DescribeDedicatedHostsResponseBodyDedicatedHosts) *DescribeDedicatedHostsResponseBody
	GetDedicatedHosts() *DescribeDedicatedHostsResponseBodyDedicatedHosts
	SetRequestId(v string) *DescribeDedicatedHostsResponseBody
	GetRequestId() *string
}

type DescribeDedicatedHostsResponseBody struct {
	// The host group ID.
	//
	// example:
	//
	// dhg-7a9xxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The host information.
	DedicatedHosts *DescribeDedicatedHostsResponseBodyDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C860658E-68A6-46C1-AF6E-3AE7C4D3CACF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBody) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDedicatedHostsResponseBody) GetDedicatedHosts() *DescribeDedicatedHostsResponseBodyDedicatedHosts {
	return s.DedicatedHosts
}

func (s *DescribeDedicatedHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetDedicatedHosts(v *DescribeDedicatedHostsResponseBodyDedicatedHosts) *DescribeDedicatedHostsResponseBody {
	s.DedicatedHosts = v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) SetRequestId(v string) *DescribeDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBody) Validate() error {
	if s.DedicatedHosts != nil {
		if err := s.DedicatedHosts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHosts struct {
	DedicatedHosts []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts `json:"DedicatedHosts,omitempty" xml:"DedicatedHosts,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) GetDedicatedHosts() []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	return s.DedicatedHosts
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) SetDedicatedHosts(v []*DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) *DescribeDedicatedHostsResponseBodyDedicatedHosts {
	s.DedicatedHosts = v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHosts) Validate() error {
	if s.DedicatedHosts != nil {
		for _, item := range s.DedicatedHosts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts struct {
	// The host account. You can call the [CreateDedicatedHostAccount](https://help.aliyun.com/document_detail/196877.html) operation to create a host account.
	//
	// example:
	//
	// test123
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Specifies whether instances can be deployed on the host. Valid values:
	//
	// 	- **0**: Instances cannot be deployed on the host.
	//
	// 	- **1**: Instances can be deployed on the host.
	//
	// example:
	//
	// 1
	AllocationStatus *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	// The bastion host ID.
	//
	// example:
	//
	// bastionhost-cn-m7xxxxxxxx
	BastionInstanceId *string `json:"BastionInstanceId,omitempty" xml:"BastionInstanceId,omitempty"`
	// The core overcommitment ratio of the dedicated cluster. Unit: percentage. For more information about the core overcommitment ratio, see [Manage a dedicated cluster](https://help.aliyun.com/document_detail/182328.html).
	//
	// example:
	//
	// 200
	CPUAllocationRatio *string `json:"CPUAllocationRatio,omitempty" xml:"CPUAllocationRatio,omitempty"`
	// The number of used CPU cores on the host. Unit: cores.
	//
	// example:
	//
	// 4
	CpuUsed *string `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	// The time when the host was created.
	//
	// example:
	//
	// 2021-03-25 17:29:06.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The dedicated cluster ID.
	//
	// example:
	//
	// dhg-7a9xxxxxxxx
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The host ID.
	//
	// example:
	//
	// i-bpxxxxxxx
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The disk overcommitment ratio of the dedicated cluster. Unit: percentage. For more information about the core overcommitment ratio, see [Manage a dedicated cluster](https://help.aliyun.com/document_detail/182328.html).
	//
	// example:
	//
	// 200
	DiskAllocationRatio *string `json:"DiskAllocationRatio,omitempty" xml:"DiskAllocationRatio,omitempty"`
	// The time when the host expires.
	//
	// example:
	//
	// 2021-04-25T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of instances that are created on the host.
	//
	// example:
	//
	// mysql
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The total number of CPU cores that are configured for the host. Unit: cores.
	//
	// example:
	//
	// 8
	HostCPU *string `json:"HostCPU,omitempty" xml:"HostCPU,omitempty"`
	// The instance type of the host.
	//
	// example:
	//
	// ecs.i2.16xlarge
	HostClass *string `json:"HostClass,omitempty" xml:"HostClass,omitempty"`
	// The total memory space of the host. Unit: MB.
	//
	// example:
	//
	// 32238
	HostMem *string `json:"HostMem,omitempty" xml:"HostMem,omitempty"`
	// The host name.
	//
	// example:
	//
	// testHost1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The status of the host. Valid values:
	//
	// 	- **0**: creating
	//
	// 	- **1**: running
	//
	// 	- **2**: faulty
	//
	// 	- **3**: being replaced
	//
	// 	- **4**: deprecated
	//
	// 	- **5**: deleting
	//
	// 	- **6**: restarting
	//
	// example:
	//
	// 1
	HostStatus *string `json:"HostStatus,omitempty" xml:"HostStatus,omitempty"`
	// The storage capacity of the host. Unit: MB.
	//
	// example:
	//
	// 2097152
	HostStorage *string `json:"HostStorage,omitempty" xml:"HostStorage,omitempty"`
	// The storage type of the host. Valid values:
	//
	// 	- **dhg_cloud_ssd**: ESSD
	//
	// 	- **dhg_local_ssd**: local SSD
	//
	// example:
	//
	// dhg_cloud_ssd
	HostType *string `json:"HostType,omitempty" xml:"HostType,omitempty"`
	// The internal IP address of the host.
	//
	// example:
	//
	// 192.xx.xx.xx
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The host image. This parameter is returned only when the **Engine*	- parameter is set to **mssql**. Valid values:
	//
	// 	- **WindowsWithMssqlStdLicense**: a Windows image that contains the licenses of SQL Server Standard Edition
	//
	// 	- **WindowsWithMssqlEntLisence**: a Windows image that contains the licenses of SQL Server Enterprise Edition
	//
	// 	- **WindowsWithMssqlWebLisence**: a Windows image that contains the licenses of SQL Server Web Edition
	//
	// example:
	//
	// WindowsWithMssqlStdLicense
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// The total number of instances that are created on the host.
	//
	// example:
	//
	// 4
	InstanceNumber *string `json:"InstanceNumber,omitempty" xml:"InstanceNumber,omitempty"`
	// The maximum memory usage per host in the dedicated cluster.
	//
	// example:
	//
	// 90
	MemAllocationRatio *string `json:"MemAllocationRatio,omitempty" xml:"MemAllocationRatio,omitempty"`
	// The size of the used memory. Unit: MB.
	//
	// example:
	//
	// 16384
	MemoryUsed *string `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	// Indicates whether the feature that allows you to have the OS permissions on the host is enabled. Valid values:
	//
	// 	- **0*	- or **null**: The permissions cannot be granted.
	//
	// 	- **1**: The permissions can be granted.
	//
	// 	- **3**: The permissions have been granted.
	//
	// example:
	//
	// 3
	OpenPermission *string `json:"OpenPermission,omitempty" xml:"OpenPermission,omitempty"`
	// The amount of used storage space on the host.
	//
	// example:
	//
	// 0
	StorageUsed *string `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the host belongs.
	//
	// example:
	//
	// vpc-bpxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch associated with the specified VPC.
	//
	// example:
	//
	// vsw-bpxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the host.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetAccountName() *string {
	return s.AccountName
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetAllocationStatus() *string {
	return s.AllocationStatus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetBastionInstanceId() *string {
	return s.BastionInstanceId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetCPUAllocationRatio() *string {
	return s.CPUAllocationRatio
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetCpuUsed() *string {
	return s.CpuUsed
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetDiskAllocationRatio() *string {
	return s.DiskAllocationRatio
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostCPU() *string {
	return s.HostCPU
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostClass() *string {
	return s.HostClass
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostMem() *string {
	return s.HostMem
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostName() *string {
	return s.HostName
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostStatus() *string {
	return s.HostStatus
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostStorage() *string {
	return s.HostStorage
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetHostType() *string {
	return s.HostType
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetInstanceNumber() *string {
	return s.InstanceNumber
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetMemAllocationRatio() *string {
	return s.MemAllocationRatio
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetMemoryUsed() *string {
	return s.MemoryUsed
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetOpenPermission() *string {
	return s.OpenPermission
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetStorageUsed() *string {
	return s.StorageUsed
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAccountName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AccountName = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetAllocationStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.AllocationStatus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetBastionInstanceId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.BastionInstanceId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCPUAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CPUAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCpuUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CpuUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetCreatedTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDedicatedHostGroupId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDedicatedHostId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetDiskAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.DiskAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEndTime(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.EndTime = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetEngine(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.Engine = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostCPU(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostCPU = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostClass(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostClass = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostMem(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostMem = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostName(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostName = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostStatus(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostStatus = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostStorage(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostStorage = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetHostType(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.HostType = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetIPAddress(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.IPAddress = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetImageCategory(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.ImageCategory = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetInstanceNumber(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.InstanceNumber = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMemAllocationRatio(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MemAllocationRatio = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetMemoryUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.MemoryUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetOpenPermission(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.OpenPermission = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetStorageUsed(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetVPCId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.VPCId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetVSwitchId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) SetZoneId(v string) *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsResponseBodyDedicatedHostsDedicatedHosts) Validate() error {
	return dara.Validate(s)
}
