// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeBackupPoliciesResponseBodyPageInfo) *DescribeBackupPoliciesResponseBody
	GetPageInfo() *DescribeBackupPoliciesResponseBodyPageInfo
	SetPolicies(v []*DescribeBackupPoliciesResponseBodyPolicies) *DescribeBackupPoliciesResponseBody
	GetPolicies() []*DescribeBackupPoliciesResponseBodyPolicies
	SetRequestId(v string) *DescribeBackupPoliciesResponseBody
	GetRequestId() *string
}

type DescribeBackupPoliciesResponseBody struct {
	// The pagination information.
	PageInfo *DescribeBackupPoliciesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The details of the mitigation policies.
	Policies []*DescribeBackupPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request and can be used to troubleshoot issues.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesResponseBody) GetPageInfo() *DescribeBackupPoliciesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeBackupPoliciesResponseBody) GetPolicies() []*DescribeBackupPoliciesResponseBodyPolicies {
	return s.Policies
}

func (s *DescribeBackupPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPoliciesResponseBody) SetPageInfo(v *DescribeBackupPoliciesResponseBodyPageInfo) *DescribeBackupPoliciesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetPolicies(v []*DescribeBackupPoliciesResponseBodyPolicies) *DescribeBackupPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) SetRequestId(v string) *DescribeBackupPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupPoliciesResponseBodyPageInfo struct {
	// The number of entries on the current page in a paged query.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in the returned data.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of backup policies per page in a paged query. Default value: 10, which indicates that each page contains 10 backup policies.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of backup policies in the returned data.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBackupPoliciesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) SetCount(v int32) *DescribeBackupPoliciesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeBackupPoliciesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) SetPageSize(v int32) *DescribeBackupPoliciesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeBackupPoliciesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupPoliciesResponseBodyPolicies struct {
	// The number of errors reported by the anti-ransomware client.
	//
	// example:
	//
	// 2
	ClientErrorCount *int32 `json:"ClientErrorCount,omitempty" xml:"ClientErrorCount,omitempty"`
	// The list of UUIDs of clients in the **abnormal*	- state.
	ClientErrorUuidList []*string `json:"ClientErrorUuidList,omitempty" xml:"ClientErrorUuidList,omitempty" type:"Repeated"`
	// The status of the anti-ransomware client. Valid values:
	//
	// - **running**: Normal.
	//
	// - **exception**: Abnormal.
	//
	// example:
	//
	// running
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The number of clients in the Normal state.
	//
	// example:
	//
	// 2
	HealthClientCount *int32 `json:"HealthClientCount,omitempty" xml:"HealthClientCount,omitempty"`
	// The list of UUIDs of clients in the **healthy*	- state.
	HealthClientUuidList []*string `json:"HealthClientUuidList,omitempty" xml:"HealthClientUuidList,omitempty" type:"Repeated"`
	// The ID of the anti-ransomware mitigation policy.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The most recent time when the anti-ransomware mitigation policy status was updated, in milliseconds.
	//
	// example:
	//
	// 1719488535027
	LastStatusSyncTime *int64 `json:"LastStatusSyncTime,omitempty" xml:"LastStatusSyncTime,omitempty"`
	// The name of the anti-ransomware mitigation policy.
	//
	// example:
	//
	// SecurityStrategy-20200303
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the anti-ransomware mitigation policy. This parameter is in JSON format. The following fields are included:
	//
	// - **IsDefault**: The type of the mitigation policy. Valid values:
	//
	//     - **1**: recommended policy
	//
	//     - **0**: custom policy
	//
	// - **Include**: The file types to protect. If all file types are protected, this parameter is set to [].
	//
	// - **Source**: The server folders to protect. If all folders need to be protected, this parameter is set to [].
	//
	// - **ExcludeSystemPath**: Specifies whether to exclude specified folders. To exclude folders, set this parameter to **true**. If you do not want to exclude folders, you do not need to set this parameter.
	//
	// - **Exclude**: The specified protection folder addresses. If no specific protection folder address is set, this parameter is set to [].
	//
	// - **Schedule**: The execution time and interval of the data backup node. Specify a non-peak hour that is not on the hour. Examples:
	//
	//     - Example 1: I|1583216092|P21D indicates that data backup starts at 2020-03-03 14:14:52, and the backup policy executes at an interval of 3 weeks.
	//
	//     - Example 2: I|1583216092|PT24H indicates that data backup starts at 2020-03-03 14:14:52, and the backup policy executes at an interval of 24 hours.
	//
	// - **Retention**: The retention period of backup data, in days. 7 indicates 1 week, 365 indicates 1 year, and -1 indicates permanent retention.
	//
	// - **SpeedLimiter**: The backup network bandwidth throttling. For example, 0:24:30720 indicates that the backup network bandwidth throttling is 30 MB/s from 00:00 to 24:00.
	//
	// - **UseVss**: Specifies whether to enable the VSS (Windows) feature. Valid values:
	//
	//     - **true**: enabled
	//
	//     - **false**: not enabled
	//
	// > The VSS (Windows) feature is available only for Windows systems. After this feature is enabled, it effectively reduces the issue of individual file backup failures caused by process occupation. Enable this feature. After this feature is enabled, file backup for exFAT and FAT32 disk formats is not supported.
	//
	// example:
	//
	// {"Exclude":["/bin/","/usr/bin/","/sbin/","/boot/","/proc/","/sys/","/srv/","/lib/","/selinux/","/usr/sbin/","/run/","/lib32/","/lib64/","/lost+found/","/var/lib/kubelet/","/var/lib/ntp/proc","/var/lib/container","Windows\\\\","Python27\\\\","Program Files (x86)\\\\","Program Files\\\\","Boot\\\\","$RECYCLE.BIN","System Volume Information\\\\","Users\\\\Administrator\\\\NTUSER.DAT*","ProgramData\\\\","pagefile.sys","Users\\\\Default\\\\NTUSER.DAT*","Users\\\\Administrator\\\\ntuser.*"],"ExcludeSystemPath":true,"Include":[],"IsDefault":1,"Retention":7,"Schedule":"I|1630689360|PT24H","Source":[],"SpeedLimiter":"","UseVss":true}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region ID of the backup service selected when the anti-ransomware client is installed on a non-Alibaba Cloud server.
	//
	// example:
	//
	// ch-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the mitigation policy. Valid values:
	//
	// - 1.0.0.
	//
	// - 2.0.0.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The previous status of the anti-ransomware mitigation policy.
	//
	// - **enabled**: The policy was manually enabled.
	//
	// - **disabled**: The policy was manually disabled. After the policy is disabled, running backup nodes are stopped.
	//
	// - **closed**: The anti-ransomware capacity was exceeded, and the system disabled the policy.
	//
	// example:
	//
	// disabled
	PreStatus *string `json:"PreStatus,omitempty" xml:"PreStatus,omitempty"`
	// The list of UUIDs of servers returned after the search by the MachineRemark request parameter.
	RemarkedUuidList []*string `json:"RemarkedUuidList,omitempty" xml:"RemarkedUuidList,omitempty" type:"Repeated"`
	// The method used to select covered assets. Valid values:
	//
	// - **ALL_MACHINE**: All assets.
	//
	// > If the policy covers **all assets**, this property value is **ALL_MACHINE**.
	//
	// example:
	//
	// ALL_MACHINE
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The server type. Valid values:
	//
	// - **OUT_CLOUD**: Non-Alibaba Cloud server.
	//
	// - **ALIYUN**: Alibaba Cloud server.
	//
	// - **TRIPARTITE**: Lightweight application server.
	//
	// example:
	//
	// OUT_CLOUD
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The number of servers with data backup exceptions.
	//
	// example:
	//
	// 2
	ServiceErrorCount *int32 `json:"ServiceErrorCount,omitempty" xml:"ServiceErrorCount,omitempty"`
	// The list of UUIDs of servers with data backup exceptions.
	ServiceErrorUuidList []*string `json:"ServiceErrorUuidList,omitempty" xml:"ServiceErrorUuidList,omitempty" type:"Repeated"`
	// The status of the anti-ransomware mitigation policy.
	//
	// - **enabled**: The policy is manually enabled.
	//
	// - **disabled**: The policy is manually disabled. After the policy is disabled, running backup nodes are stopped.
	//
	// - **closed**: The anti-ransomware capacity is exceeded, and the system disables the policy.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The upgrade status of the policy. Valid values:
	//
	// - **NotUpgraded**: Not upgraded.
	//
	// - **Upgrading**: Upgrading.
	//
	// - **UpgradeFailed**: Upgrade failed.
	//
	// - **UpgradeSuccess**: Upgrade succeeded.
	//
	// example:
	//
	// Upgrading
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	// The list of UUIDs of servers protected by the anti-ransomware mitigation policy.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DescribeBackupPoliciesResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetClientErrorCount() *int32 {
	return s.ClientErrorCount
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetClientErrorUuidList() []*string {
	return s.ClientErrorUuidList
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetHealthClientCount() *int32 {
	return s.HealthClientCount
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetHealthClientUuidList() []*string {
	return s.HealthClientUuidList
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetId() *int64 {
	return s.Id
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetLastStatusSyncTime() *int64 {
	return s.LastStatusSyncTime
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetPolicyRegionId() *string {
	return s.PolicyRegionId
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetPreStatus() *string {
	return s.PreStatus
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetRemarkedUuidList() []*string {
	return s.RemarkedUuidList
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetSelectType() *string {
	return s.SelectType
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetServiceErrorCount() *int32 {
	return s.ServiceErrorCount
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetServiceErrorUuidList() []*string {
	return s.ServiceErrorUuidList
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) GetUuidList() []*string {
	return s.UuidList
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetClientErrorCount(v int32) *DescribeBackupPoliciesResponseBodyPolicies {
	s.ClientErrorCount = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetClientErrorUuidList(v []*string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.ClientErrorUuidList = v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetClientStatus(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.ClientStatus = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetHealthClientCount(v int32) *DescribeBackupPoliciesResponseBodyPolicies {
	s.HealthClientCount = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetHealthClientUuidList(v []*string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.HealthClientUuidList = v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetId(v int64) *DescribeBackupPoliciesResponseBodyPolicies {
	s.Id = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetLastStatusSyncTime(v int64) *DescribeBackupPoliciesResponseBodyPolicies {
	s.LastStatusSyncTime = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetName(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetPolicy(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.Policy = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetPolicyRegionId(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.PolicyRegionId = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetPolicyVersion(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.PolicyVersion = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetPreStatus(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.PreStatus = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetRemarkedUuidList(v []*string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.RemarkedUuidList = v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetSelectType(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.SelectType = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetServerType(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.ServerType = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetServiceErrorCount(v int32) *DescribeBackupPoliciesResponseBodyPolicies {
	s.ServiceErrorCount = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetServiceErrorUuidList(v []*string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.ServiceErrorUuidList = v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetStatus(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.Status = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetUpgradeStatus(v string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.UpgradeStatus = &v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) SetUuidList(v []*string) *DescribeBackupPoliciesResponseBodyPolicies {
	s.UuidList = v
	return s
}

func (s *DescribeBackupPoliciesResponseBodyPolicies) Validate() error {
	return dara.Validate(s)
}
