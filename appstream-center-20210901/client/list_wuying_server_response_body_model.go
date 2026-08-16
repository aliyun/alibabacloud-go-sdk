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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The list of workstation information.
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
	if s.WuyingServerList != nil {
		for _, item := range s.WuyingServerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWuyingServerResponseBodyWuyingServerList struct {
	// The status of adding to the virtual node pool.
	//
	// example:
	//
	// Added
	AddVirtualNodePoolStatus *string `json:"AddVirtualNodePoolStatus,omitempty" xml:"AddVirtualNodePoolStatus,omitempty"`
	// The tenant UID.
	//
	// example:
	//
	// 1234567890123456
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The bandwidth size. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-08-02T16:52:11.000+00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of data cloud disks.
	DataDisk []*ListWuyingServerResponseBodyWuyingServerListDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The maximum number of private IP addresses per ENI, including the primary IP address.
	//
	// example:
	//
	// 10
	EniPrivateIpAddressQuantity *int32 `json:"EniPrivateIpAddressQuantity,omitempty" xml:"EniPrivateIpAddressQuantity,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2025-09-03T16:00:00.000+00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The FOTA version number.
	//
	// example:
	//
	// 2.0.0
	FotaVersion *string `json:"FotaVersion,omitempty" xml:"FotaVersion,omitempty"`
	// The image ID.
	//
	// example:
	//
	// imgc-06****oagaev
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Alibaba Cloud Linux 3.2104 LTS 64位
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The list of workspace instance information.
	InstanceInfoList []*ListWuyingServerResponseBodyWuyingServerListInstanceInfoList `json:"InstanceInfoList,omitempty" xml:"InstanceInfoList,omitempty" type:"Repeated"`
	// The maximum price of the spot instance.
	//
	// example:
	//
	// 0.5
	MaxPrice *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	// The internal IP address.
	//
	// example:
	//
	// 10.80.21.149
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-1b****ayv2
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// exampleOfficeSite
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The office network type.
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The operating system type.
	//
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The list of policy group IDs.
	PolicyGroupIdList []*string `json:"PolicyGroupIdList,omitempty" xml:"PolicyGroupIdList,omitempty" type:"Repeated"`
	// The set of private IP addresses, including the primary IP address and secondary IP addresses.
	PrivateIpSets []*ListWuyingServerResponseBodyWuyingServerListPrivateIpSets `json:"PrivateIpSets,omitempty" xml:"PrivateIpSets,omitempty" type:"Repeated"`
	// The resource session status.
	//
	// example:
	//
	// Connected
	ResourceSessionStatus *string `json:"ResourceSessionStatus,omitempty" xml:"ResourceSessionStatus,omitempty"`
	// The list of security group IDs.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The instance type information.
	ServerInstanceTypeInfo *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo `json:"ServerInstanceTypeInfo,omitempty" xml:"ServerInstanceTypeInfo,omitempty" type:"Struct"`
	// The list of sessions.
	Sessions []*ListWuyingServerResponseBodyWuyingServerListSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The workstation status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-payment type.
	//
	// example:
	//
	// spot
	SubPayType *string `json:"SubPayType,omitempty" xml:"SubPayType,omitempty"`
	// The system cloud disk type.
	//
	// example:
	//
	// cloud_auto
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The ID of the system cloud disk.
	//
	// example:
	//
	// d-bp1234567890abcde
	SystemDiskId *string `json:"SystemDiskId,omitempty" xml:"SystemDiskId,omitempty"`
	// The system cloud disk performance level.
	//
	// example:
	//
	// PL0
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	// The system cloud disk size. Unit: GB.
	//
	// example:
	//
	// 100
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The ID of the timer group.
	//
	// example:
	//
	// tg-bp1234567890abcde
	TimerGroupId *string `json:"TimerGroupId,omitempty" xml:"TimerGroupId,omitempty"`
	// The list of authorized users.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	// The IP address of the Virtual Kubelet node.
	//
	// example:
	//
	// 10.0.0.100
	VirtualKubeletIp *string `json:"VirtualKubeletIp,omitempty" xml:"VirtualKubeletIp,omitempty"`
	// The virtual node pool ID.
	//
	// example:
	//
	// vnp-bp1234567890abcde
	VirtualNodePoolId *string `json:"VirtualNodePoolId,omitempty" xml:"VirtualNodePoolId,omitempty"`
	// Indicates whether the Virtual Kubelet needs to be upgraded.
	VkUpgradeNeeded *bool `json:"VkUpgradeNeeded,omitempty" xml:"VkUpgradeNeeded,omitempty"`
	// The Virtual Kubelet version.
	//
	// example:
	//
	// 1.0.0
	VkVersion *string `json:"VkVersion,omitempty" xml:"VkVersion,omitempty"`
	// The workstation ID.
	//
	// example:
	//
	// ws-0byd****8wn2lwi
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
	// The workstation name.
	//
	// example:
	//
	// exampleServerName
	WuyingServerName *string `json:"WuyingServerName,omitempty" xml:"WuyingServerName,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerList) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerList) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetAddVirtualNodePoolStatus() *string {
	return s.AddVirtualNodePoolStatus
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetBandwidth() *int32 {
	return s.Bandwidth
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

func (s *ListWuyingServerResponseBodyWuyingServerList) GetEniPrivateIpAddressQuantity() *int32 {
	return s.EniPrivateIpAddressQuantity
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetFotaVersion() *string {
	return s.FotaVersion
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

func (s *ListWuyingServerResponseBodyWuyingServerList) GetMaxPrice() *float32 {
	return s.MaxPrice
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

func (s *ListWuyingServerResponseBodyWuyingServerList) GetPolicyGroupIdList() []*string {
	return s.PolicyGroupIdList
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetPrivateIpSets() []*ListWuyingServerResponseBodyWuyingServerListPrivateIpSets {
	return s.PrivateIpSets
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetResourceSessionStatus() *string {
	return s.ResourceSessionStatus
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetServerInstanceTypeInfo() *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	return s.ServerInstanceTypeInfo
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSessions() []*ListWuyingServerResponseBodyWuyingServerListSessions {
	return s.Sessions
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetStatus() *string {
	return s.Status
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSubPayType() *string {
	return s.SubPayType
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskId() *string {
	return s.SystemDiskId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetTimerGroupId() *string {
	return s.TimerGroupId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetUsers() []*string {
	return s.Users
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetVirtualKubeletIp() *string {
	return s.VirtualKubeletIp
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetVirtualNodePoolId() *string {
	return s.VirtualNodePoolId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetVkUpgradeNeeded() *bool {
	return s.VkUpgradeNeeded
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetVkVersion() *string {
	return s.VkVersion
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetWuyingServerName() *string {
	return s.WuyingServerName
}

func (s *ListWuyingServerResponseBodyWuyingServerList) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetAddVirtualNodePoolStatus(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.AddVirtualNodePoolStatus = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetAliUid(v int64) *ListWuyingServerResponseBodyWuyingServerList {
	s.AliUid = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetBandwidth(v int32) *ListWuyingServerResponseBodyWuyingServerList {
	s.Bandwidth = &v
	return s
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

func (s *ListWuyingServerResponseBodyWuyingServerList) SetEniPrivateIpAddressQuantity(v int32) *ListWuyingServerResponseBodyWuyingServerList {
	s.EniPrivateIpAddressQuantity = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetExpiredTime(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ExpiredTime = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetFotaVersion(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.FotaVersion = &v
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

func (s *ListWuyingServerResponseBodyWuyingServerList) SetMaxPrice(v float32) *ListWuyingServerResponseBodyWuyingServerList {
	s.MaxPrice = &v
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

func (s *ListWuyingServerResponseBodyWuyingServerList) SetPolicyGroupIdList(v []*string) *ListWuyingServerResponseBodyWuyingServerList {
	s.PolicyGroupIdList = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetPrivateIpSets(v []*ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) *ListWuyingServerResponseBodyWuyingServerList {
	s.PrivateIpSets = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetResourceSessionStatus(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ResourceSessionStatus = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSecurityGroupIds(v []*string) *ListWuyingServerResponseBodyWuyingServerList {
	s.SecurityGroupIds = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetServerInstanceTypeInfo(v *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) *ListWuyingServerResponseBodyWuyingServerList {
	s.ServerInstanceTypeInfo = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSessions(v []*ListWuyingServerResponseBodyWuyingServerListSessions) *ListWuyingServerResponseBodyWuyingServerList {
	s.Sessions = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetStatus(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.Status = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSubPayType(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.SubPayType = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSystemDiskCategory(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.SystemDiskCategory = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetSystemDiskId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.SystemDiskId = &v
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

func (s *ListWuyingServerResponseBodyWuyingServerList) SetTimerGroupId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.TimerGroupId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetUsers(v []*string) *ListWuyingServerResponseBodyWuyingServerList {
	s.Users = v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetVirtualKubeletIp(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.VirtualKubeletIp = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetVirtualNodePoolId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.VirtualNodePoolId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetVkUpgradeNeeded(v bool) *ListWuyingServerResponseBodyWuyingServerList {
	s.VkUpgradeNeeded = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) SetVkVersion(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.VkVersion = &v
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

func (s *ListWuyingServerResponseBodyWuyingServerList) SetZoneId(v string) *ListWuyingServerResponseBodyWuyingServerList {
	s.ZoneId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerList) Validate() error {
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.InstanceInfoList != nil {
		for _, item := range s.InstanceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PrivateIpSets != nil {
		for _, item := range s.PrivateIpSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServerInstanceTypeInfo != nil {
		if err := s.ServerInstanceTypeInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWuyingServerResponseBodyWuyingServerListDataDisk struct {
	// The data cloud disk type.
	//
	// example:
	//
	// cloud_auto
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The data cloud disk ID.
	//
	// example:
	//
	// d-bp1234567890abcde
	DataDiskId *string `json:"DataDiskId,omitempty" xml:"DataDiskId,omitempty"`
	// The data cloud disk sequence number.
	//
	// example:
	//
	// 1
	DataDiskNo *string `json:"DataDiskNo,omitempty" xml:"DataDiskNo,omitempty"`
	// The data cloud disk performance level.
	//
	// example:
	//
	// PL0
	DataDiskPerformanceLevel *string `json:"DataDiskPerformanceLevel,omitempty" xml:"DataDiskPerformanceLevel,omitempty"`
	// The data cloud disk size. Unit: GB.
	//
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

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) GetDataDiskId() *string {
	return s.DataDiskId
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) GetDataDiskNo() *string {
	return s.DataDiskNo
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

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) SetDataDiskId(v string) *ListWuyingServerResponseBodyWuyingServerListDataDisk {
	s.DataDiskId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListDataDisk) SetDataDiskNo(v string) *ListWuyingServerResponseBodyWuyingServerListDataDisk {
	s.DataDiskNo = &v
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
	// The instance ID.
	//
	// example:
	//
	// p-0ceitx****c5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network interface controller (NIC).
	//
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

type ListWuyingServerResponseBodyWuyingServerListPrivateIpSets struct {
	// Indicates whether the IP address is the primary private IP address. A value of true indicates the primary private IP address. A value of false indicates a secondary private IP address.
	//
	// example:
	//
	// true
	Primary *bool `json:"Primary,omitempty" xml:"Primary,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 10.0.0.1
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) GetPrimary() *bool {
	return s.Primary
}

func (s *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) SetPrimary(v bool) *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets {
	s.Primary = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) SetPrivateIpAddress(v string) *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListPrivateIpSets) Validate() error {
	return dara.Validate(s)
}

type ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo struct {
	// The number of vCPUs.
	//
	// example:
	//
	// 96
	Cpu *string `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of GPUs.
	//
	// example:
	//
	// 4
	Gpu *string `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// The GPU memory size. Unit: MB.
	//
	// example:
	//
	// 196,608
	GpuMemory *int32 `json:"GpuMemory,omitempty" xml:"GpuMemory,omitempty"`
	// The GPU specification description.
	//
	// example:
	//
	// NVIDIA T4
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The memory size. Unit: MB.
	//
	// example:
	//
	// 393,216
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The workstation instance type.
	//
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

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) GetGpuSpec() *string {
	return s.GpuSpec
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

func (s *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo) SetGpuSpec(v string) *ListWuyingServerResponseBodyWuyingServerListServerInstanceTypeInfo {
	s.GpuSpec = &v
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

type ListWuyingServerResponseBodyWuyingServerListSessions struct {
	// The start time of the session.
	//
	// example:
	//
	// 2026-01-01T08:00:00Z
	ResourceSessionStartTime *string `json:"ResourceSessionStartTime,omitempty" xml:"ResourceSessionStartTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// user1
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListWuyingServerResponseBodyWuyingServerListSessions) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerResponseBodyWuyingServerListSessions) GoString() string {
	return s.String()
}

func (s *ListWuyingServerResponseBodyWuyingServerListSessions) GetResourceSessionStartTime() *string {
	return s.ResourceSessionStartTime
}

func (s *ListWuyingServerResponseBodyWuyingServerListSessions) GetUserId() *string {
	return s.UserId
}

func (s *ListWuyingServerResponseBodyWuyingServerListSessions) SetResourceSessionStartTime(v string) *ListWuyingServerResponseBodyWuyingServerListSessions {
	s.ResourceSessionStartTime = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListSessions) SetUserId(v string) *ListWuyingServerResponseBodyWuyingServerListSessions {
	s.UserId = &v
	return s
}

func (s *ListWuyingServerResponseBodyWuyingServerListSessions) Validate() error {
	return dara.Validate(s)
}
