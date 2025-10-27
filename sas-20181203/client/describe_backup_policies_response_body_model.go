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
	// The details of the anti-ransomware policy.
	Policies []*DescribeBackupPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
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
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of anti-ransomware policies returned.
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
	// The number of the servers on which the anti-ransomware agent is in an abnormal state.
	//
	// example:
	//
	// 2
	ClientErrorCount *int32 `json:"ClientErrorCount,omitempty" xml:"ClientErrorCount,omitempty"`
	// The UUIDs of the servers on which the anti-ransomware agent is in an **abnormal*	- state.
	ClientErrorUuidList []*string `json:"ClientErrorUuidList,omitempty" xml:"ClientErrorUuidList,omitempty" type:"Repeated"`
	// The status of the anti-ransomware agent. Valid values:
	//
	// 	- **running**: normal
	//
	// 	- **exception**: abnormal
	//
	// example:
	//
	// running
	ClientStatus *string `json:"ClientStatus,omitempty" xml:"ClientStatus,omitempty"`
	// The number of the servers on which the anti-ransomware agent is in a normal state.
	//
	// example:
	//
	// 2
	HealthClientCount *int32 `json:"HealthClientCount,omitempty" xml:"HealthClientCount,omitempty"`
	// The UUIDs of the servers on which the anti-ransomware agent is in a **normal*	- state.
	HealthClientUuidList []*string `json:"HealthClientUuidList,omitempty" xml:"HealthClientUuidList,omitempty" type:"Repeated"`
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the anti-ransomware policy was last updated. Unit: milliseconds.
	//
	// example:
	//
	// 1719488535027
	LastStatusSyncTime *int64 `json:"LastStatusSyncTime,omitempty" xml:"LastStatusSyncTime,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// SecurityStrategy-20200303
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The configurations of the anti-ransomware policy. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **IsDefault**: the type of the anti-ransomware policy. Valid values:
	//
	//     	- **1**: recommended policy
	//
	//     	- **0**: custom policy
	//
	// 	- **Include**: the format of the files that are protected. If the value of this field is [], all formats of files are protected.
	//
	// 	- **Source**: the directory that is protected. If the value of this field is [], all directories are protected.
	//
	// 	- **ExcludeSystemPath**: indicates whether a specified directory is excluded from the anti-ransomware policy. If the value of this field is **true**, a directory is excluded. If this field is left empty, no directories are excluded.
	//
	// 	- **Exclude**: the directory that is excluded from the anti-ransomware policy. If the value of this field is [], no directories are excluded.
	//
	// 	- **Schedule**: the start time and interval of a data backup task. We recommend that you specify a start time that begins during off-peak hours but does not start on the hour. Examples:
	//
	//     	- If the value of this field is I|1583216092|P21D, the data backup task starts from 2020-03-03 14:14:52, and the task is run at an interval of three weeks.
	//
	//     	- If the value of this field is I|1583216092|PT24H, the data backup task starts from 2020-03-03 14:14:52, and the task is run at an interval of 24 hours.
	//
	// 	- **Retention**: the period during which backup data is retained. Unit: days. If the value of this field is 7, backup data is retained for a week. If the value of this field is 365, backup data is retained for a year. If the value of this field is -1, backup data is permanently retained.
	//
	// 	- **SpeedLimiter**: the limit on the network bandwidth for data backup tasks. If the value of this field is 0:24:30720, the maximum bandwidth for a data backup task is 30 MB/s from 00:00 to 24:00.
	//
	// 	- **UseVss**: indicates whether the VSS feature is enabled. The feature is available only for Windows servers. Valid values:
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// >  The VSS feature is available only if you create the anti-ransomware policy for Windows servers. After you enable the feature, the number of backup failures due to running processes is significantly reduced. We recommend that you enable the VSS feature. After you enable the feature, the data of disks that are in the exFAT and FAT32 formats cannot be backed up.
	//
	// example:
	//
	// {"Exclude":["/bin/","/usr/bin/","/sbin/","/boot/","/proc/","/sys/","/srv/","/lib/","/selinux/","/usr/sbin/","/run/","/lib32/","/lib64/","/lost+found/","/var/lib/kubelet/","/var/lib/ntp/proc","/var/lib/container","Windows\\\\","Python27\\\\","Program Files (x86)\\\\","Program Files\\\\","Boot\\\\","$RECYCLE.BIN","System Volume Information\\\\","Users\\\\Administrator\\\\NTUSER.DAT*","ProgramData\\\\","pagefile.sys","Users\\\\Default\\\\NTUSER.DAT*","Users\\\\Administrator\\\\ntuser.*"],"ExcludeSystemPath":true,"Include":[],"IsDefault":1,"Retention":7,"Schedule":"I|1630689360|PT24H","Source":[],"SpeedLimiter":"","UseVss":true}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the region that you specified for data backup when you installed the anti-ransomware agent for the server not deployed on Alibaba Cloud.
	//
	// example:
	//
	// ch-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the anti-ransomware policy. Valid values:
	//
	// 	- 1.0.0
	//
	// 	- 2.0.0
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The previous status of the anti-ransomware policy. Valid values:
	//
	// 	- **enabled**: The anti-ransomware policy is manually enabled.
	//
	// 	- **disabled**: The anti-ransomware policy is manually disabled. After an anti-ransomware policy is disabled, the data backup task that is running based on the policy stops.
	//
	// 	- **closed**: The anti-ransomware policy automatically stops because the anti-ransomware capacity is insufficient.
	//
	// example:
	//
	// disabled
	PreStatus *string `json:"PreStatus,omitempty" xml:"PreStatus,omitempty"`
	// The UUIDs that are returned based on the value of the MachineRemark request parameter.
	RemarkedUuidList []*string `json:"RemarkedUuidList,omitempty" xml:"RemarkedUuidList,omitempty" type:"Repeated"`
	// The type of the server. Valid values:
	//
	// 	- **OUT_CLOUD**: server not deployed on Alibaba Cloud
	//
	// 	- **ALIYUN**: Elastic Compute Service (ECS) instance
	//
	// 	- **TRIPARTITE**: simple application server
	//
	// example:
	//
	// OUT_CLOUD
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The number of servers on which data backup is exceptional.
	//
	// example:
	//
	// 2
	ServiceErrorCount *int32 `json:"ServiceErrorCount,omitempty" xml:"ServiceErrorCount,omitempty"`
	// The UUIDs of the servers on which data backup is exceptional.
	ServiceErrorUuidList []*string `json:"ServiceErrorUuidList,omitempty" xml:"ServiceErrorUuidList,omitempty" type:"Repeated"`
	// The status of the anti-ransomware policy. Valid values:
	//
	// 	- **enabled**: The anti-ransomware policy is manually enabled.
	//
	// 	- **disabled**: The anti-ransomware policy is manually disabled. After an anti-ransomware policy is disabled, the data backup task that is running based on the policy stops.
	//
	// 	- **closed**: The anti-ransomware policy automatically stops because the anti-ransomware capacity is insufficient.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The upgrade status of the anti-ransomware policy. Valid values:
	//
	// 	- **NotUpgraded**
	//
	// 	- **Upgrading**
	//
	// 	- **UpgradeFailed**
	//
	// 	- **UpgradeSuccess**
	//
	// example:
	//
	// Upgrading
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	// The UUIDs of the servers to which the anti-ransomware policy is applied.
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
