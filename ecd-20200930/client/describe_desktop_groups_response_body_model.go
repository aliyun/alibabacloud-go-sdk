// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroups(v []*DescribeDesktopGroupsResponseBodyDesktopGroups) *DescribeDesktopGroupsResponseBody
	GetDesktopGroups() []*DescribeDesktopGroupsResponseBodyDesktopGroups
	SetNextToken(v string) *DescribeDesktopGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopGroupsResponseBody
	GetRequestId() *string
}

type DescribeDesktopGroupsResponseBody struct {
	// The cloud computer shares.
	DesktopGroups []*DescribeDesktopGroupsResponseBodyDesktopGroups `json:"DesktopGroups,omitempty" xml:"DesktopGroups,omitempty" type:"Repeated"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDesktopGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseBody) GetDesktopGroups() []*DescribeDesktopGroupsResponseBodyDesktopGroups {
	return s.DesktopGroups
}

func (s *DescribeDesktopGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopGroupsResponseBody) SetDesktopGroups(v []*DescribeDesktopGroupsResponseBodyDesktopGroups) *DescribeDesktopGroupsResponseBody {
	s.DesktopGroups = v
	return s
}

func (s *DescribeDesktopGroupsResponseBody) SetNextToken(v string) *DescribeDesktopGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBody) SetRequestId(v string) *DescribeDesktopGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBody) Validate() error {
	if s.DesktopGroups != nil {
		for _, item := range s.DesktopGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDesktopGroupsResponseBodyDesktopGroups struct {
	// The number of concurrent sessions allowed for each cloud computer within the multi-session many-to-many share.
	//
	// example:
	//
	// 1
	BindAmount *int64 `json:"BindAmount,omitempty" xml:"BindAmount,omitempty"`
	// This parameter is applicable only to subscription cloud computer shares. It defines the initial number of cloud computers that are purchased. Valid values: 0 to 200.
	//
	// example:
	//
	// 5
	BuyDesktopsCount *int32 `json:"BuyDesktopsCount,omitempty" xml:"BuyDesktopsCount,omitempty"`
	// The remarks.
	//
	// example:
	//
	// test
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The maximum period of time during which a session is connected. When the specified maximum period of time is reached, the session is automatically disconnected. Unit: milliseconds.
	//
	// example:
	//
	// 90000
	ConnectDuration *int64 `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// The number of cloud computers in each state.
	CountPerStatus []*DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus `json:"CountPerStatus,omitempty" xml:"CountPerStatus,omitempty" type:"Repeated"`
	// The number of vCPUs.
	//
	// example:
	//
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The time when the cloud computer pool was created.
	//
	// example:
	//
	// 2022-02-17T14:51:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account that creates the cloud computer pool.
	//
	// example:
	//
	// 1007214305******
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The category of the user disk.
	//
	// Valid values:
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: enhanced SSD (ESSD)
	//
	// example:
	//
	// cloud_ssd
	DataDiskCategory *string `json:"DataDiskCategory,omitempty" xml:"DataDiskCategory,omitempty"`
	// The user disk capacity. Unit: GiB.
	//
	// example:
	//
	// 50
	DataDiskSize *string `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	// The number of cloud computers that are created.
	//
	// example:
	//
	// 2
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The ID of the cloud computer share.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The name of the cloud computer share.
	//
	// example:
	//
	// test1
	DesktopGroupName *string `json:"DesktopGroupName,omitempty" xml:"DesktopGroupName,omitempty"`
	// The cloud computer type. You can call the [DescribeDesktopTypes](https://help.aliyun.com/document_detail/188882.html) operation to query the IDs of the cloud computer types supported by WUYING Workspace.
	//
	// example:
	//
	// eds.enterprise_office.4c4g
	DesktopType *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	// The number of users who can access the cloud computer share.
	//
	// example:
	//
	// 1
	EndUserCount *int32 `json:"EndUserCount,omitempty" xml:"EndUserCount,omitempty"`
	// The expiration date of the subscription cloud computer share.
	//
	// example:
	//
	// 2022-03-17T16:00:00Z
	ExpiredTime  *string   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ExpiredTimes []*string `json:"ExpiredTimes,omitempty" xml:"ExpiredTimes,omitempty" type:"Repeated"`
	// The number of GPUs.
	//
	// example:
	//
	// 1
	GpuCount *float32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// The version of the GPU driver.
	//
	// example:
	//
	// 12
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// The GPU memory.
	//
	// example:
	//
	// 16 GiB
	GpuSpec *string `json:"GpuSpec,omitempty" xml:"GpuSpec,omitempty"`
	// The period of time after which a session is closed. After an end user connects to a cloud computer, the session is established. If the system does not detect inputs from the keyboard or mouse within the specified period of time, the session is closed. Unit: milliseconds.
	//
	// example:
	//
	// 90000
	IdleDisconnectDuration *int64 `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-gq15cq5ydlvwn****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	IsLdap  *bool   `json:"IsLdap,omitempty" xml:"IsLdap,omitempty"`
	// The keep-alive duration of a session after the session is disconnected. Valid values: 180000 (3 minutes) to 345600000 (4 days). Unit: milliseconds. If you set this parameter to 0, the session is permanently retained after it is disconnected.
	//
	// When a session is disconnected, take note of the following situations: If an end user does not resume the session within the specified duration, the session is closed and all unsaved data is cleared. If the end user resumes the session within the specified duration, the end user can continue to access data of the session.
	//
	// example:
	//
	// 1000
	KeepDuration *int64 `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	// The load balancing policy for the multi-session many-to-many share.
	//
	// Valid values:
	//
	// 	- 0: depth-first
	//
	// 	- 1: breadth-first
	//
	// example:
	//
	// 1
	LoadPolicy *int64 `json:"LoadPolicy,omitempty" xml:"LoadPolicy,omitempty"`
	// 	- For pay-as-you-go cloud computer shares, this parameter defines the maximum number of cloud computers allowed.
	//
	// 	- For subscription cloud computer shares, this parameter defines the total number of cloud computers, including both the initially purchased cloud computers (`BuyDesktopsCount`) and those that can be auto-created.
	//
	// example:
	//
	// 10
	MaxDesktopsCount *int32 `json:"MaxDesktopsCount,omitempty" xml:"MaxDesktopsCount,omitempty"`
	// The memory size. Unit: MiB.
	//
	// example:
	//
	// 16384
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 	- For pay-as-you-go cloud computer shares, this parameter defines the minimum number of cloud computers allowed.
	//
	// 	- For subscription cloud computer shares, this parameter defines the number of cloud computers that are initially purchased (`BuyDesktopsCount`).
	//
	// example:
	//
	// 1
	MinDesktopsCount *int32 `json:"MinDesktopsCount,omitempty" xml:"MinDesktopsCount,omitempty"`
	// The ID of the office network in which the cloud computer network resides.
	//
	// example:
	//
	// cn-hangzhou+dir-467671****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the office network in which the cloud computer share resides.
	//
	// example:
	//
	// testName
	OfficeSiteName *string `json:"OfficeSiteName,omitempty" xml:"OfficeSiteName,omitempty"`
	// The account type of the office network.
	//
	// Valid values:
	//
	// 	- PERSONAL: individual office network
	//
	// 	- SIMPLE: convenience office network
	//
	// 	- AD_CONNECTOR: enterprise Active Directory (AD) office network
	//
	// 	- RAM: Resource Access Management (RAM)-based office network
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The OS.
	//
	// Valid values:
	//
	// 	- Linux
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Windows
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The ID of the cloud computer template.
	//
	// example:
	//
	// bundle_eds_general_4c8g_s8d5_win2019
	OwnBundleId *string `json:"OwnBundleId,omitempty" xml:"OwnBundleId,omitempty"`
	// The name of the cloud computer template.
	//
	// example:
	//
	// test
	OwnBundleName *string `json:"OwnBundleName,omitempty" xml:"OwnBundleName,omitempty"`
	// The type of the cloud computer share.
	//
	// Valid values:
	//
	// 	- 0: a single-session many-to-many share.
	//
	// 	- 1: a multi-session many-to-many share.
	//
	// example:
	//
	// 0
	OwnType *int64 `json:"OwnType,omitempty" xml:"OwnType,omitempty"`
	// The billing method of the cloud computer pool.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go
	//
	// 	- PrePaid: subscription
	//
	// example:
	//
	// PrePaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the applied policy.
	//
	// example:
	//
	// pg-53iyi2aar0nd6c8qj
	PolicyGroupId     *string   `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	PolicyGroupIdList []*string `json:"PolicyGroupIdList,omitempty" xml:"PolicyGroupIdList,omitempty" type:"Repeated"`
	// The name of the applied policy.
	//
	// example:
	//
	// test-policy
	PolicyGroupName     *string   `json:"PolicyGroupName,omitempty" xml:"PolicyGroupName,omitempty"`
	PolicyGroupNameList []*string `json:"PolicyGroupNameList,omitempty" xml:"PolicyGroupNameList,omitempty" type:"Repeated"`
	// The protocol type.
	//
	// Valid values:
	//
	// 	- HDX
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- ASP
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// ASP
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The threshold for the ratio of connected sessions, which triggers automatic scaling of cloud computers within the multi-session many-to-many share. To calculate the ratio of connected sessions, use the following formula:
	//
	// `Ratio of connected sessions = Number of connected sessions/(Total number of cloud computers × Maximum number of sessions allowed for each cloud computer) × 100%`
	//
	// When the specified threshold is reached, new cloud computers are automatically created. When the specified threshold is not reached, idle cloud computers are released.
	//
	// example:
	//
	// 0.85
	RatioThreshold *float32 `json:"RatioThreshold,omitempty" xml:"RatioThreshold,omitempty"`
	// The reset option for the cloud computer share.
	//
	// Valid values:
	//
	// 	- 0: does not reset any disk.
	//
	// 	- 1: resets only the system disk.
	//
	// 	- 2: resets only the data disk.
	//
	// 	- 3: resets the system disk and data disk.
	//
	// example:
	//
	// 0
	ResetType         *int64  `json:"ResetType,omitempty" xml:"ResetType,omitempty"`
	SimpleUserGroupId *string `json:"SimpleUserGroupId,omitempty" xml:"SimpleUserGroupId,omitempty"`
	// The status of the cloud computer share.
	//
	// Valid values:
	//
	// 	- 0: The cloud computer share is unpaid.
	//
	// 	- 1: The cloud computer share is normal.
	//
	// 	- 2: The cloud computer share expired, or your account has an overdue payment.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The period of time after which an idle cloud computer is stopped. When the specified period of time is reached, the cloud computer is automatically stopped. If an end user connects to the stopped cloud computer, the cloud computer is automatically started. Unit: milliseconds.
	//
	// example:
	//
	// 900000
	StopDuration *int64 `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	// The ID of the subnet.
	//
	// example:
	//
	// vsw-uf63bb6*****8gfytm
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The category of the system disk.
	//
	// Valid values:
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: enhanced SSD (ESSD)
	//
	// example:
	//
	// cloud_ssd
	SystemDiskCategory *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	// The system disk capacity. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The tags.
	Tags          []*DescribeDesktopGroupsResponseBodyDesktopGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UserGroupName *string                                               `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	// The user\\"s organizational unit path.
	//
	// example:
	//
	// example.com
	UserOuPath *string `json:"UserOuPath,omitempty" xml:"UserOuPath,omitempty"`
	// The version number of the cloud computer share.
	//
	// example:
	//
	// 2
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
	// Indicates whether disk encryption is enabled.
	//
	// example:
	//
	// false
	VolumeEncryptionEnabled *bool `json:"VolumeEncryptionEnabled,omitempty" xml:"VolumeEncryptionEnabled,omitempty"`
	// The ID of the Key Management Service (KMS) key for disk encryption.
	//
	// example:
	//
	// e5409ada-xxxx-xxxx-xxxx-89e31e23e993
	VolumeEncryptionKey *string `json:"VolumeEncryptionKey,omitempty" xml:"VolumeEncryptionKey,omitempty"`
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroups) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetBindAmount() *int64 {
	return s.BindAmount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetBuyDesktopsCount() *int32 {
	return s.BuyDesktopsCount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetComments() *string {
	return s.Comments
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetConnectDuration() *int64 {
	return s.ConnectDuration
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetCountPerStatus() []*DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus {
	return s.CountPerStatus
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetCpu() *int32 {
	return s.Cpu
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetCreator() *string {
	return s.Creator
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetDataDiskCategory() *string {
	return s.DataDiskCategory
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetDataDiskSize() *string {
	return s.DataDiskSize
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetDesktopCount() *int32 {
	return s.DesktopCount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetDesktopGroupName() *string {
	return s.DesktopGroupName
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetDesktopType() *string {
	return s.DesktopType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetEndUserCount() *int32 {
	return s.EndUserCount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetExpiredTimes() []*string {
	return s.ExpiredTimes
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetGpuCount() *float32 {
	return s.GpuCount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetGpuDriverVersion() *string {
	return s.GpuDriverVersion
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetGpuSpec() *string {
	return s.GpuSpec
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetIdleDisconnectDuration() *int64 {
	return s.IdleDisconnectDuration
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetIsLdap() *bool {
	return s.IsLdap
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetKeepDuration() *int64 {
	return s.KeepDuration
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetLoadPolicy() *int64 {
	return s.LoadPolicy
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetMaxDesktopsCount() *int32 {
	return s.MaxDesktopsCount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetMemory() *int64 {
	return s.Memory
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetMinDesktopsCount() *int32 {
	return s.MinDesktopsCount
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOfficeSiteName() *string {
	return s.OfficeSiteName
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOsType() *string {
	return s.OsType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOwnBundleId() *string {
	return s.OwnBundleId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOwnBundleName() *string {
	return s.OwnBundleName
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetOwnType() *int64 {
	return s.OwnType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetPolicyGroupIdList() []*string {
	return s.PolicyGroupIdList
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetPolicyGroupName() *string {
	return s.PolicyGroupName
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetPolicyGroupNameList() []*string {
	return s.PolicyGroupNameList
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetRatioThreshold() *float32 {
	return s.RatioThreshold
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetResetType() *int64 {
	return s.ResetType
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetSimpleUserGroupId() *string {
	return s.SimpleUserGroupId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetStopDuration() *int64 {
	return s.StopDuration
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetSystemDiskCategory() *string {
	return s.SystemDiskCategory
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetTags() []*DescribeDesktopGroupsResponseBodyDesktopGroupsTags {
	return s.Tags
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetUserOuPath() *string {
	return s.UserOuPath
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetVolumeEncryptionEnabled() *bool {
	return s.VolumeEncryptionEnabled
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) GetVolumeEncryptionKey() *string {
	return s.VolumeEncryptionKey
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetBindAmount(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.BindAmount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetBuyDesktopsCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.BuyDesktopsCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetComments(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Comments = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetConnectDuration(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ConnectDuration = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCountPerStatus(v []*DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.CountPerStatus = v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCpu(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Cpu = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCreateTime(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.CreateTime = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetCreator(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Creator = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDataDiskCategory(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DataDiskCategory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDataDiskSize(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DataDiskSize = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDesktopCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DesktopCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDesktopGroupId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DesktopGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDesktopGroupName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DesktopGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetDesktopType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.DesktopType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetEndUserCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.EndUserCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetExpiredTime(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetExpiredTimes(v []*string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ExpiredTimes = v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetGpuCount(v float32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.GpuCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetGpuDriverVersion(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.GpuDriverVersion = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetGpuSpec(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.GpuSpec = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetIdleDisconnectDuration(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetImageId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ImageId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetIsLdap(v bool) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.IsLdap = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetKeepDuration(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.KeepDuration = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetLoadPolicy(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.LoadPolicy = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetMaxDesktopsCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.MaxDesktopsCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetMemory(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Memory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetMinDesktopsCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.MinDesktopsCount = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOfficeSiteId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOfficeSiteName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OfficeSiteName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOfficeSiteType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOsType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OsType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOwnBundleId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OwnBundleId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOwnBundleName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OwnBundleName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetOwnType(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.OwnType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPayType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PayType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPolicyGroupId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPolicyGroupIdList(v []*string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PolicyGroupIdList = v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPolicyGroupName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PolicyGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetPolicyGroupNameList(v []*string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.PolicyGroupNameList = v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetProtocolType(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ProtocolType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetRatioThreshold(v float32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.RatioThreshold = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetResetType(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.ResetType = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetSimpleUserGroupId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.SimpleUserGroupId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetStatus(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Status = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetStopDuration(v int64) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.StopDuration = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetSubnetId(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.SubnetId = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetSystemDiskCategory(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetSystemDiskSize(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetTags(v []*DescribeDesktopGroupsResponseBodyDesktopGroupsTags) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Tags = v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetUserGroupName(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.UserGroupName = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetUserOuPath(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.UserOuPath = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetVersion(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.Version = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetVolumeEncryptionEnabled(v bool) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.VolumeEncryptionEnabled = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) SetVolumeEncryptionKey(v string) *DescribeDesktopGroupsResponseBodyDesktopGroups {
	s.VolumeEncryptionKey = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroups) Validate() error {
	if s.CountPerStatus != nil {
		for _, item := range s.CountPerStatus {
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

type DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus struct {
	// The total number of cloud computers.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The status of the cloud computer.
	//
	// Valid values:
	//
	// 	- Stopped
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Starting
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Rebuilding
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Stopping
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Expired
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Deleted
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Pending
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) SetCount(v int32) *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus {
	s.Count = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) SetStatus(v string) *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus {
	s.Status = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsCountPerStatus) Validate() error {
	return dara.Validate(s)
}

type DescribeDesktopGroupsResponseBodyDesktopGroupsTags struct {
	// The tag key.
	//
	// example:
	//
	// desktop_group_name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ds-dq2mybjr23yw*****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopGroupsResponseBodyDesktopGroupsTags) GoString() string {
	return s.String()
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsTags) SetKey(v string) *DescribeDesktopGroupsResponseBodyDesktopGroupsTags {
	s.Key = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsTags) SetValue(v string) *DescribeDesktopGroupsResponseBodyDesktopGroupsTags {
	s.Value = &v
	return s
}

func (s *DescribeDesktopGroupsResponseBodyDesktopGroupsTags) Validate() error {
	return dara.Validate(s)
}
