// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktops(v []*DescribeDesktopsResponseBodyDesktops) *DescribeDesktopsResponseBody
	GetDesktops() []*DescribeDesktopsResponseBodyDesktops
	SetNextToken(v string) *DescribeDesktopsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDesktopsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDesktopsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDesktopsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDesktopsResponseBody
	GetTotalCount() *int32
}

type DescribeDesktopsResponseBody struct {
	// The cloud computers.
	Desktops []*DescribeDesktopsResponseBodyDesktops `json:"Desktops,omitempty" xml:"Desktops,omitempty" type:"Repeated"`
	// The token that is used for the next query. If this parameter is left empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 484256DA-D816-44D2-9D86-B6EE4D5BA78C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud computers.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBody) GetDesktops() []*DescribeDesktopsResponseBodyDesktops {
	return s.Desktops
}

func (s *DescribeDesktopsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDesktopsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDesktopsResponseBody) SetDesktops(v []*DescribeDesktopsResponseBodyDesktops) *DescribeDesktopsResponseBody {
	s.Desktops = v
	return s
}

func (s *DescribeDesktopsResponseBody) SetNextToken(v string) *DescribeDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetPageNumber(v int32) *DescribeDesktopsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetPageSize(v int32) *DescribeDesktopsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetRequestId(v string) *DescribeDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetTotalCount(v int32) *DescribeDesktopsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDesktopsResponseBody) Validate() error {
	if s.Desktops != nil {
		for _, item := range s.Desktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopsResponseBodyDesktops struct {
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The number of concurrent sessions of each cloud computer in a multi-session cloud computer pool.
	//
	// example:
	//
	// 10
	BindAmount *int32 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// The ID of the template used to create the cloud computer.
	//
	// example:
	//
	// b-2g65ljy4291vl****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The name of the template used to create the cloud computer.
	//
	// example:
	//
	// Name
	BundleName *string `json:"BundleName,omitempty" xml:"BundleName,omitempty"`
	// The billing method of the cloud computer.
	//
	// Valid values:
	//
	// 	- Postpaid (default): pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The connection status of the end user.
	//
	// Valid values:
	//
	// 	- Unknown
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Connected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Disconnected
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Disconnected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The number of vCPUs.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the cloud computer was created.
	//
	// example:
	//
	// 2020-11-06T08:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	DataDiskSize        *string                                                    `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	DesktopDurationList []*DescribeDesktopsResponseBodyDesktopsDesktopDurationList `json:"DesktopDurationList,omitempty" xml:"DesktopDurationList,omitempty" type:"Repeated"`
	// The ID of the cloud computer pool to which cloud computers belong. Default value: null.``
	//
	// example:
	//
	// null
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// testDesktopName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The cloud computer status.
	//
	// example:
	//
	// Running
	DesktopStatus *string `json:"DesktopStatus,omitempty" xml:"DesktopStatus,omitempty"`
	// The cloud computer type.
	//
	// example:
	//
	// ecd.basic.large
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The directory ID, which is the same as the office network ID (OfficeSiteId).
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// SIMPLE
	DirectoryType *string `json:"DirectoryType,omitempty" xml:"DirectoryType,omitempty"`
	// The information about the disks.
	Disks      []*DescribeDesktopsResponseBodyDesktopsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	DomainType *string                                      `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// The number of times for which the cloud desktop can be downgraded.
	//
	// example:
	//
	// 3
	DowngradeQuota *int64 `json:"DowngradeQuota,omitempty" xml:"DowngradeQuota,omitempty"`
	// The number of times for which the cloud desktop has been downgraded.
	//
	// example:
	//
	// 0
	DowngradedTimes *int64 `json:"DowngradedTimes,omitempty" xml:"DowngradedTimes,omitempty"`
	// The end user IDs.
	EndUserIds      []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	EntraDomainName *string   `json:"EntraDomainName,omitempty" xml:"EntraDomainName,omitempty"`
	EnvId           *string   `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	EnvType         *string   `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The time when a subscription cloud computer expired.
	//
	// example:
	//
	// 2021-12-31T15:59Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The information about the image version of the cloud computer.
	FotaUpdate *DescribeDesktopsResponseBodyDesktopsFotaUpdate `json:"FotaUpdate,omitempty" xml:"FotaUpdate,omitempty" type:"Struct"`
	// Indicates whether the cloud computer uses GPUs.
	//
	// example:
	//
	// 0
	GpuCategory *int64 `json:"GpuCategory,omitempty" xml:"GpuCategory,omitempty"`
	// The number of GPU cores.
	//
	// example:
	//
	// 1
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The GPU driver version used by the cloud computer.
	//
	// example:
	//
	// null
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The GPU Specifications.
	//
	// example:
	//
	// NVIDIA T4
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// true
	HibernationBeta *bool `json:"HibernationBeta,omitempty" xml:"HibernationBeta,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// true
	HibernationOptionsConfigured *bool `json:"HibernationOptionsConfigured,omitempty" xml:"HibernationOptionsConfigured,omitempty"`
	// The hostname of the cloud desktop.
	//
	// example:
	//
	// testName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-4zfb6zj728hhr****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	IsLdap  *bool   `json:"IsLdap,omitempty" xml:"IsLdap,omitempty"`
	// The flag that is used to manage the cloud computer.
	//
	// Valid values:
	//
	// 	- Migrating: The cloud computer is being migrated.
	//
	// 	- Updating: The configurations of the cloud computer are being updated.
	//
	// 	- NoFlag: No flags are available.
	//
	// example:
	//
	// NoFlag
	ManagementFlag *string `json:"ManagementFlag,omitempty" xml:"ManagementFlag,omitempty"`
	// The flags that are used to manage the cloud computers.
	ManagementFlags []*string `json:"ManagementFlags,omitempty" xml:"ManagementFlags,omitempty" type:"Repeated"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 4096
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The ID of the supplementary network interface controller (NIC) created by EDS within an RAM user or Active Directory (AD) user. You cannot modify the ID.
	//
	// example:
	//
	// 123456
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The IP address of the supplementary NIC created by EDS within an RAM or AD user.
	//
	// example:
	//
	// 192.168.74.165
	NetworkInterfaceIp *string `json:"NetworkInterfaceIp,omitempty" xml:"NetworkInterfaceIp,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The office network name.
	//
	// example:
	//
	// test
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The account type of the office network.
	//
	// Valid values:
	//
	// 	- SIMPLE: convenience account
	//
	// 	- AD_CONNECTOR: enterprise AD account
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The VPC type of the office network.
	//
	// Valid values:
	//
	// 	- standard
	//
	// 	- customized
	//
	// 	- basic
	//
	// example:
	//
	// basic
	OfficeSiteVpcType *string `json:"OfficeSiteVpcType,omitempty" xml:"OfficeSiteVpcType,omitempty"`
	// The OS that is defined in the desktop template.
	//
	// example:
	//
	// Windows
	OsType   *string                                       `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OsUpdate *DescribeDesktopsResponseBodyDesktopsOsUpdate `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty" type:"Struct"`
	// The information about the OS platform.
	//
	// Valid values:
	//
	// 	- Ubuntu
	//
	// 	- Windows Server 2022
	//
	// 	- UOS
	//
	// 	- CentOS
	//
	// 	- Windows Server 2019
	//
	// 	- Windows Server 2016
	//
	// example:
	//
	// Ubuntu
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// system-all-enabled-policy
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The IDs of the cloud computer policies.
	PolicyGroupIdList []*string `json:"PolicyGroupIdList,omitempty" xml:"PolicyGroupIdList,omitempty" type:"Repeated"`
	// The policy name.
	//
	// example:
	//
	// test
	PolicyGroupName *string `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	// The names of the cloud computer policies.
	PolicyGroupNameList []*string `json:"PolicyGroupNameList,omitempty" xml:"PolicyGroupNameList,omitempty" type:"Repeated"`
	// The progress of creating the cloud computer.
	//
	// example:
	//
	// 100%
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The protocol.
	//
	// Valid values:
	//
	// 	- HDX
	//
	// 	- ASP
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The information about the enterprise resource groups.
	ResourceGroups []*DescribeDesktopsResponseBodyDesktopsResourceGroups `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	SerialNumber   *string                                               `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The type of the session.
	//
	// Valid values:
	//
	// 	- SINGLE_SESSION
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- MULTIPLE_SESSION
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SINGLE_SESSION
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// The session information about cloud computers connected by end users.
	Sessions []*DescribeDesktopsResponseBodyDesktopsSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The snapshot policy ID.
	//
	// example:
	//
	// sp-gi007jgyc3kcey2bb
	SnapshotPolicyId *string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty"`
	// The name of the snapshot policy.
	//
	// example:
	//
	// testSnapshotName
	SnapshotPolicyName *string `json:"SnapshotPolicyName,omitempty" xml:"SnapshotPolicyName,omitempty"`
	// The standard start time.
	//
	// example:
	//
	// 2025-02-24T06:38:02Z
	StandardStartTime *string `json:"StandardStartTime,omitempty" xml:"StandardStartTime,omitempty"`
	// The time when the cloud computer was first started.
	//
	// example:
	//
	// 2020-11-06T08:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Indicates whether the cloud desktop supports hibernation.
	//
	// example:
	//
	// true
	SupportHibernation *bool `json:"SupportHibernation,omitempty" xml:"SupportHibernation,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 0
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// Details about the tags.
	Tags []*DescribeDesktopsResponseBodyDesktopsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key that is used when disk encryption is enabled. You can call the [ListKeys](https://help.aliyun.com/document_detail/28951.html) operation to query the list of KMS keys.
	//
	// example:
	//
	// 08c33a6f-4e0a-4a1b-a3fa-7ddfa1d4****
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
	// The zone type. Default value: `AvailabilityZone`. This value indicates Alibaba Cloud zones.
	//
	// example:
	//
	// AvailabilityZone
	ZoneType *string `json:"ZoneType,omitempty" xml:"ZoneType,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktops) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktops) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetBindAmount() *int32 {
	return s.BindAmount
}

func (s *DescribeDesktopsResponseBodyDesktops) GetBundleId() *string {
	return s.BundleId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetBundleName() *string {
	return s.BundleName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeDesktopsResponseBodyDesktops) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeDesktopsResponseBodyDesktops) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDataDiskSize() *string {
	return s.DataDiskSize
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDesktopDurationList() []*DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	return s.DesktopDurationList
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDesktopStatus() *string {
	return s.DesktopStatus
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDirectoryType() *string {
	return s.DirectoryType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDisks() []*DescribeDesktopsResponseBodyDesktopsDisks {
	return s.Disks
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDomainType() *string {
	return s.DomainType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDowngradeQuota() *int64 {
	return s.DowngradeQuota
}

func (s *DescribeDesktopsResponseBodyDesktops) GetDowngradedTimes() *int64 {
	return s.DowngradedTimes
}

func (s *DescribeDesktopsResponseBodyDesktops) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeDesktopsResponseBodyDesktops) GetEntraDomainName() *string {
	return s.EntraDomainName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetEnvId() *string {
	return s.EnvId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetEnvType() *string {
	return s.EnvType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDesktopsResponseBodyDesktops) GetFotaUpdate() *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	return s.FotaUpdate
}

func (s *DescribeDesktopsResponseBodyDesktops) GetGpuCategory() *int64 {
	return s.GpuCategory
}

func (s *DescribeDesktopsResponseBodyDesktops) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *DescribeDesktopsResponseBodyDesktops) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeDesktopsResponseBodyDesktops) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeDesktopsResponseBodyDesktops) GetHibernationBeta() *bool {
	return s.HibernationBeta
}

func (s *DescribeDesktopsResponseBodyDesktops) GetHibernationOptionsConfigured() *bool {
	return s.HibernationOptionsConfigured
}

func (s *DescribeDesktopsResponseBodyDesktops) GetHostName() *string {
	return s.HostName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetIsLdap() *bool {
	return s.IsLdap
}

func (s *DescribeDesktopsResponseBodyDesktops) GetManagementFlag() *string {
	return s.ManagementFlag
}

func (s *DescribeDesktopsResponseBodyDesktops) GetManagementFlags() []*string {
	return s.ManagementFlags
}

func (s *DescribeDesktopsResponseBodyDesktops) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeDesktopsResponseBodyDesktops) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetNetworkInterfaceIp() *string {
	return s.NetworkInterfaceIp
}

func (s *DescribeDesktopsResponseBodyDesktops) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetOfficeSiteVpcType() *string {
	return s.OfficeSiteVpcType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetOsType() *string {
	return s.OsType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetOsUpdate() *DescribeDesktopsResponseBodyDesktopsOsUpdate {
	return s.OsUpdate
}

func (s *DescribeDesktopsResponseBodyDesktops) GetPlatform() *string {
	return s.Platform
}

func (s *DescribeDesktopsResponseBodyDesktops) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetPolicyGroupIdList() []*string {
	return s.PolicyGroupIdList
}

func (s *DescribeDesktopsResponseBodyDesktops) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetPolicyGroupNameList() []*string {
	return s.PolicyGroupNameList
}

func (s *DescribeDesktopsResponseBodyDesktops) GetProgress() *string {
	return s.Progress
}

func (s *DescribeDesktopsResponseBodyDesktops) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetResourceGroups() []*DescribeDesktopsResponseBodyDesktopsResourceGroups {
	return s.ResourceGroups
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSessionType() *string {
	return s.SessionType
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSessions() []*DescribeDesktopsResponseBodyDesktopsSessions {
	return s.Sessions
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSnapshotPolicyId() *string {
	return s.SnapshotPolicyId
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSnapshotPolicyName() *string {
	return s.SnapshotPolicyName
}

func (s *DescribeDesktopsResponseBodyDesktops) GetStandardStartTime() *string {
	return s.StandardStartTime
}

func (s *DescribeDesktopsResponseBodyDesktops) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSupportHibernation() *bool {
	return s.SupportHibernation
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeDesktopsResponseBodyDesktops) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeDesktopsResponseBodyDesktops) GetTags() []*DescribeDesktopsResponseBodyDesktopsTags {
	return s.Tags
}

func (s *DescribeDesktopsResponseBodyDesktops) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *DescribeDesktopsResponseBodyDesktops) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *DescribeDesktopsResponseBodyDesktops) GetZoneType() *string {
	return s.ZoneType
}

func (s *DescribeDesktopsResponseBodyDesktops) SetAccountType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.AccountType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetBindAmount(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.BindAmount = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetBundleId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.BundleId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetBundleName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.BundleName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetChargeType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ChargeType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetConnectionStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCpu(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetCreationTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDataDiskSize(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopDurationList(v []*DescribeDesktopsResponseBodyDesktopsDesktopDurationList) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopDurationList = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopGroupId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopStatus(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDesktopType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDirectoryType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DirectoryType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDisks(v []*DescribeDesktopsResponseBodyDesktopsDisks) *DescribeDesktopsResponseBodyDesktops {
	s.Disks = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDomainType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.DomainType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDowngradeQuota(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.DowngradeQuota = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetDowngradedTimes(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.DowngradedTimes = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEndUserIds(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.EndUserIds = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEntraDomainName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.EntraDomainName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEnvId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.EnvId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetEnvType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.EnvType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetExpiredTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetFotaUpdate(v *DescribeDesktopsResponseBodyDesktopsFotaUpdate) *DescribeDesktopsResponseBodyDesktops {
	s.FotaUpdate = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuCategory(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.GpuCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuCount(v float32) *DescribeDesktopsResponseBodyDesktops {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuDriverVersion(v string) *DescribeDesktopsResponseBodyDesktops {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetGpuSpec(v string) *DescribeDesktopsResponseBodyDesktops {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetHibernationBeta(v bool) *DescribeDesktopsResponseBodyDesktops {
	s.HibernationBeta = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetHibernationOptionsConfigured(v bool) *DescribeDesktopsResponseBodyDesktops {
	s.HibernationOptionsConfigured = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetHostName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.HostName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetImageId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetIsLdap(v bool) *DescribeDesktopsResponseBodyDesktops {
	s.IsLdap = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetManagementFlag(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ManagementFlag = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetManagementFlags(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.ManagementFlags = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetMemory(v int64) *DescribeDesktopsResponseBodyDesktops {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetNetworkInterfaceIp(v string) *DescribeDesktopsResponseBodyDesktops {
	s.NetworkInterfaceIp = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOfficeSiteVpcType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OfficeSiteVpcType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOsType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetOsUpdate(v *DescribeDesktopsResponseBodyDesktopsOsUpdate) *DescribeDesktopsResponseBodyDesktops {
	s.OsUpdate = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPlatform(v string) *DescribeDesktopsResponseBodyDesktops {
	s.Platform = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupIdList(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupIdList = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetPolicyGroupNameList(v []*string) *DescribeDesktopsResponseBodyDesktops {
	s.PolicyGroupNameList = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetProgress(v string) *DescribeDesktopsResponseBodyDesktops {
	s.Progress = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetProtocolType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetResourceGroups(v []*DescribeDesktopsResponseBodyDesktopsResourceGroups) *DescribeDesktopsResponseBodyDesktops {
	s.ResourceGroups = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSerialNumber(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSessionType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SessionType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSessions(v []*DescribeDesktopsResponseBodyDesktopsSessions) *DescribeDesktopsResponseBodyDesktops {
	s.Sessions = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSnapshotPolicyId(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SnapshotPolicyId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSnapshotPolicyName(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SnapshotPolicyName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetStandardStartTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.StandardStartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetStartTime(v string) *DescribeDesktopsResponseBodyDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSupportHibernation(v bool) *DescribeDesktopsResponseBodyDesktops {
	s.SupportHibernation = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskCategory(v string) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetSystemDiskSize(v int32) *DescribeDesktopsResponseBodyDesktops {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetTags(v []*DescribeDesktopsResponseBodyDesktopsTags) *DescribeDesktopsResponseBodyDesktops {
	s.Tags = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetVolumeEncryptionEnabled(v bool) *DescribeDesktopsResponseBodyDesktops {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetVolumeEncryptionKey(v string) *DescribeDesktopsResponseBodyDesktops {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) SetZoneType(v string) *DescribeDesktopsResponseBodyDesktops {
	s.ZoneType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktops) Validate() error {
	if s.DesktopDurationList != nil {
		for _, item := range s.DesktopDurationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Disks != nil {
		for _, item := range s.Disks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FotaUpdate != nil {
		if err := s.FotaUpdate.Validate(); err != nil {
			return err
		}
	}
	if s.OsUpdate != nil {
		if err := s.OsUpdate.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceGroups != nil {
		for _, item := range s.ResourceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopsResponseBodyDesktopsDesktopDurationList struct {
	OrderInstanceId       *string  `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	PackageCreationTime   *string  `json:"PackageCreationTime,omitempty" xml:"PackageCreationTime,omitempty"`
	PackageExpiredTime    *string  `json:"PackageExpiredTime,omitempty" xml:"PackageExpiredTime,omitempty"`
	PackageId             *string  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	PackageStatus         *string  `json:"PackageStatus,omitempty" xml:"PackageStatus,omitempty"`
	PackageType           *string  `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	PackageUsedUpStrategy *string  `json:"PackageUsedUpStrategy,omitempty" xml:"PackageUsedUpStrategy,omitempty"`
	PeriodEndTime         *string  `json:"PeriodEndTime,omitempty" xml:"PeriodEndTime,omitempty"`
	PeriodStartTime       *string  `json:"PeriodStartTime,omitempty" xml:"PeriodStartTime,omitempty"`
	PostPaidLimitFee      *float32 `json:"PostPaidLimitFee,omitempty" xml:"PostPaidLimitFee,omitempty"`
	TotalDuration         *int64   `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	UsedDuration          *int64   `json:"UsedDuration,omitempty" xml:"UsedDuration,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsDesktopDurationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPackageCreationTime() *string {
	return s.PackageCreationTime
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPackageExpiredTime() *string {
	return s.PackageExpiredTime
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPackageStatus() *string {
	return s.PackageStatus
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPackageUsedUpStrategy() *string {
	return s.PackageUsedUpStrategy
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPeriodEndTime() *string {
	return s.PeriodEndTime
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPeriodStartTime() *string {
	return s.PeriodStartTime
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetPostPaidLimitFee() *float32 {
	return s.PostPaidLimitFee
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetTotalDuration() *int64 {
	return s.TotalDuration
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) GetUsedDuration() *int64 {
	return s.UsedDuration
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetOrderInstanceId(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.OrderInstanceId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPackageCreationTime(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PackageCreationTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPackageExpiredTime(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PackageExpiredTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPackageId(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PackageId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPackageStatus(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PackageStatus = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPackageType(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PackageType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPackageUsedUpStrategy(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PackageUsedUpStrategy = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPeriodEndTime(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PeriodEndTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPeriodStartTime(v string) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PeriodStartTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetPostPaidLimitFee(v float32) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.PostPaidLimitFee = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetTotalDuration(v int64) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.TotalDuration = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) SetUsedDuration(v int64) *DescribeDesktopsResponseBodyDesktopsDesktopDurationList {
	s.UsedDuration = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDesktopDurationList) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsResponseBodyDesktopsDisks struct {
	// The type of the disk. Valid values:
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_auto: standard SSD.
	//
	// 	- cloud_essd: enhanced SSD (ESSD).
	//
	// example:
	//
	// cloud_auto
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk ID.
	//
	// example:
	//
	// d-jedbpr4sl9l37****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The disk size. Unit: GiB.
	//
	// example:
	//
	// 80
	DiskSize *int32 `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	// The type of the disk.
	//
	// Valid values:
	//
	// 	- SYSTEM: system disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DATA: data disk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// SYSTEM
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The performance level (PL) of the disk when an enterprise SSD (ESSD) is used.
	//
	// For more information about the differences among enterprise SSDs (ESSDs) at different PLs, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// Valid values:
	//
	// 	- PL1
	//
	// 	- PL0
	//
	// 	- PL3
	//
	// 	- PL2
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsDisks) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) GetDiskSize() *int32 {
	return s.DiskSize
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskCategory(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskCategory = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskId(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskSize(v int32) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskSize = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetDiskType(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.DiskType = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) SetPerformanceLevel(v string) *DescribeDesktopsResponseBodyDesktopsDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsDisks) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsResponseBodyDesktopsFotaUpdate struct {
	// The current image version of the cloud computer.
	//
	// example:
	//
	// 0.0.0-D-20220102.000000
	CurrentAppVersion *string `json:"CurrentAppVersion,omitempty" xml:"CurrentAppVersion,omitempty"`
	// The version number to which the image of the cloud computer can be updated.
	//
	// example:
	//
	// 0.0.0-R-20220307.190736
	NewAppVersion *string `json:"NewAppVersion,omitempty" xml:"NewAppVersion,omitempty"`
	// The description of the version to which the image of the cloud computer can be updated.
	//
	// example:
	//
	// Upgrade package for testing 03-07
	ReleaseNote *string `json:"ReleaseNote,omitempty" xml:"ReleaseNote,omitempty"`
	// The English description of the version to which the image of the cloud computer can be updated.
	//
	// example:
	//
	// Release note
	ReleaseNoteEn *string `json:"ReleaseNoteEn,omitempty" xml:"ReleaseNoteEn,omitempty"`
	// The Japanese description of the image version to which the cloud desktop can be updated.
	//
	// example:
	//
	// リリースノート
	ReleaseNoteJp *string `json:"ReleaseNoteJp,omitempty" xml:"ReleaseNoteJp,omitempty"`
	// The size of the installation package for the image to which the cloud desktop can be updated. Unit: KB.
	//
	// example:
	//
	// 108815097
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsFotaUpdate) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsFotaUpdate) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) GetCurrentAppVersion() *string {
	return s.CurrentAppVersion
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) GetNewAppVersion() *string {
	return s.NewAppVersion
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) GetReleaseNote() *string {
	return s.ReleaseNote
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) GetReleaseNoteEn() *string {
	return s.ReleaseNoteEn
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) GetReleaseNoteJp() *string {
	return s.ReleaseNoteJp
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) GetSize() *int64 {
	return s.Size
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetCurrentAppVersion(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.CurrentAppVersion = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetNewAppVersion(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.NewAppVersion = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNote(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNote = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteEn(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteEn = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetReleaseNoteJp(v string) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.ReleaseNoteJp = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) SetSize(v int64) *DescribeDesktopsResponseBodyDesktopsFotaUpdate {
	s.Size = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsFotaUpdate) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsResponseBodyDesktopsOsUpdate struct {
	CheckId      *string                                                 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	PackageCount *int32                                                  `json:"PackageCount,omitempty" xml:"PackageCount,omitempty"`
	Packages     []*DescribeDesktopsResponseBodyDesktopsOsUpdatePackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
}

func (s DescribeDesktopsResponseBodyDesktopsOsUpdate) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsOsUpdate) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) GetPackageCount() *int32 {
	return s.PackageCount
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) GetPackages() []*DescribeDesktopsResponseBodyDesktopsOsUpdatePackages {
	return s.Packages
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) SetCheckId(v string) *DescribeDesktopsResponseBodyDesktopsOsUpdate {
	s.CheckId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) SetPackageCount(v int32) *DescribeDesktopsResponseBodyDesktopsOsUpdate {
	s.PackageCount = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) SetPackages(v []*DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) *DescribeDesktopsResponseBodyDesktopsOsUpdate {
	s.Packages = v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdate) Validate() error {
	if s.Packages != nil {
		for _, item := range s.Packages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopsResponseBodyDesktopsOsUpdatePackages struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Kb          *string `json:"Kb,omitempty" xml:"Kb,omitempty"`
	Title       *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) GetDescription() *string {
	return s.Description
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) GetKb() *string {
	return s.Kb
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) GetTitle() *string {
	return s.Title
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) SetDescription(v string) *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages {
	s.Description = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) SetKb(v string) *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages {
	s.Kb = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) SetTitle(v string) *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages {
	s.Title = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsOsUpdatePackages) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsResponseBodyDesktopsResourceGroups struct {
	// The ID of the enterprise resource group.
	//
	// example:
	//
	// rg-4hsvzbbmqdzu3s****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the enterprise resource group.
	//
	// example:
	//
	// Resource group 01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsResourceGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsResourceGroups) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsResourceGroups) GetId() *string {
	return s.Id
}

func (s *DescribeDesktopsResponseBodyDesktopsResourceGroups) GetName() *string {
	return s.Name
}

func (s *DescribeDesktopsResponseBodyDesktopsResourceGroups) SetId(v string) *DescribeDesktopsResponseBodyDesktopsResourceGroups {
	s.Id = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsResourceGroups) SetName(v string) *DescribeDesktopsResponseBodyDesktopsResourceGroups {
	s.Name = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsResourceGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsResponseBodyDesktopsSessions struct {
	// The ID of the end user that connects to the cloud computer.
	//
	// example:
	//
	// 29615820929547****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the cloud computer session was established.
	//
	// example:
	//
	// 2021-03-07T08:23Z
	EstablishmentTime *string `json:"EstablishmentTime,omitempty" xml:"EstablishmentTime,omitempty"`
	// The name of the external user.
	//
	// example:
	//
	// Testname
	ExternalUserName *string `json:"ExternalUserName,omitempty" xml:"ExternalUserName,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsSessions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsSessions) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) GetEstablishmentTime() *string {
	return s.EstablishmentTime
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) GetExternalUserName() *string {
	return s.ExternalUserName
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetEndUserId(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.EndUserId = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetEstablishmentTime(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.EstablishmentTime = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) SetExternalUserName(v string) *DescribeDesktopsResponseBodyDesktopsSessions {
	s.ExternalUserName = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsSessions) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopsResponseBodyDesktopsTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDesktopsResponseBodyDesktopsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBodyDesktopsTags) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) SetKey(v string) *DescribeDesktopsResponseBodyDesktopsTags {
	s.Key = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) SetValue(v string) *DescribeDesktopsResponseBodyDesktopsTags {
	s.Value = &v
	return s
}

func (s *DescribeDesktopsResponseBodyDesktopsTags) Validate() error {
	return dara.Validate(s)
}
