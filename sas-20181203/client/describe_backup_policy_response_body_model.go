// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPolicyDetail(v *DescribeBackupPolicyResponseBodyBackupPolicyDetail) *DescribeBackupPolicyResponseBody
	GetBackupPolicyDetail() *DescribeBackupPolicyResponseBodyBackupPolicyDetail
	SetRequestId(v string) *DescribeBackupPolicyResponseBody
	GetRequestId() *string
}

type DescribeBackupPolicyResponseBody struct {
	// The details of the anti-ransomware policy.
	BackupPolicyDetail *DescribeBackupPolicyResponseBodyBackupPolicyDetail `json:"BackupPolicyDetail,omitempty" xml:"BackupPolicyDetail,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBackupPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBody) GetBackupPolicyDetail() *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	return s.BackupPolicyDetail
}

func (s *DescribeBackupPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupPolicyResponseBody) SetBackupPolicyDetail(v *DescribeBackupPolicyResponseBodyBackupPolicyDetail) *DescribeBackupPolicyResponseBody {
	s.BackupPolicyDetail = v
	return s
}

func (s *DescribeBackupPolicyResponseBody) SetRequestId(v string) *DescribeBackupPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBody) Validate() error {
	if s.BackupPolicyDetail != nil {
		if err := s.BackupPolicyDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBackupPolicyResponseBodyBackupPolicyDetail struct {
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
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 51880
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// testBuckUp
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 	- **IsDefault**: the type of the anti-ransomware policy. Valid values:
	//
	//     	- **1**: recommended policy
	//
	//     	- **0**: custom policy
	//
	// 	- **Include**: the format of the files that you want to protect. If you want to protect the files in all formats, set this field to [].
	//
	// 	- **Source**: the directory that you want to protect. If you want to protect all directories, set this field to [].
	//
	// 	- **ExcludeSystemPath**: specifies whether to exclude a specific directory from the anti-ransomware policy. If you want to exclude a directory, set this field to **true**. If you do not want to exclude a directory, leave this field empty.
	//
	// 	- **Exclude**: the directory that you want to exclude from the anti-ransomware policy. If you do not want to exclude a directory, set this field to [].
	//
	// 	- **Schedule**: the start time and interval of a data backup task. We recommend that you specify a start time that begins during off-peak hours but does not start on the hour. Examples:
	//
	//     	- If you set this field to I|1583216092|P21D, the data backup task starts from 2020-03-03 14:14:52, and the task is run at an interval of three weeks.
	//
	//     	- If you set this field to I|1583216092|PT24H, the data backup task starts from 2020-03-03 14:14:52, and the task is run at an interval of 24 hours.
	//
	// 	- **Retention**: the period during which backup data is retained. Unit: days. If you set this field to 7, backup data is retained for a week. If you set this field to 365, backup data is retained for a year. If you set this field to -1, backup data is permanently retained.
	//
	// 	- **SpeedLimiter**: the limit on the network bandwidth for data backup tasks. If you set this field to 0:24:30720, the maximum bandwidth for a data backup task is 30 MB/s from 00:00 to 24:00.
	//
	// 	- **UseVss**: specifies whether to enable the VSS feature. The feature is available only for Windows servers. Valid values:
	//
	//     	- **true**: yes
	//
	//     	- **false**: no
	//
	// >  The VSS feature is available only if you create the anti-ransomware policy for Windows servers. After you enable the feature, the number of backup failures due to running processes is significantly reduced. We recommend that you enable the VSS feature. After you enable the feature, the data of disks that are in the exFAT and FAT32 formats cannot be backed up.
	//
	// example:
	//
	// {
	//
	//       "Exclude": [
	//
	//             "/bin/",
	//
	//             "/usr/bin/",
	//
	//             "/sbin/",
	//
	//             "/boot/",
	//
	//             "/proc/",
	//
	//             "/sys/",
	//
	//             "/srv/",
	//
	//             "/lib/",
	//
	//             "/selinux/",
	//
	//             "/usr/sbin/",
	//
	//             "/run/",
	//
	//             "/lib32/",
	//
	//             "/lib64/",
	//
	//             "/lost+found/",
	//
	//             "/var/lib/kubelet/",
	//
	//             "/var/lib/ntp/proc",
	//
	//             "/var/lib/container",
	//
	//             "Windows\\\\",
	//
	//             "Python27\\\\",
	//
	//             "Program Files (x86)\\\\",
	//
	//             "Program Files\\\\",
	//
	//             "Boot\\\\",
	//
	//             "$RECYCLE.BIN",
	//
	//             "System Volume Information\\\\",
	//
	//             "Users\\\\Administrator\\\\NTUSER.DAT*",
	//
	//             "ProgramData\\\\",
	//
	//             "pagefile.sys",
	//
	//             "Users\\\\Default\\\\NTUSER.DAT*",
	//
	//             "Users\\\\Administrator\\\\ntuser.*"
	//
	//       ],
	//
	//       "ExcludeSystemPath": true,
	//
	//       "Include": [],
	//
	//       "IsDefault": 1,
	//
	//       "Retention": 7,
	//
	//       "Schedule": "I|1630689360|PT24H",
	//
	//       "Source": [],
	//
	//       "SpeedLimiter": "",
	//
	//       "UseVss": true
	//
	// }
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The version of the anti-ransomware policy.
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The ID of the region in which backup data is stored.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
	// closed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// An array consisting of the UUIDs of the servers to which the anti-ransomware policy is applied.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupPolicyResponseBodyBackupPolicyDetail) GoString() string {
	return s.String()
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetClientStatus() *string {
	return s.ClientStatus
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetId() *int64 {
	return s.Id
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetName() *string {
	return s.Name
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) GetUuidList() []*string {
	return s.UuidList
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetClientStatus(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.ClientStatus = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetId(v int64) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.Id = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetName(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.Name = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetPolicy(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.Policy = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetPolicyVersion(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.PolicyVersion = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetRegionId(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetStatus(v string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.Status = &v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) SetUuidList(v []*string) *DescribeBackupPolicyResponseBodyBackupPolicyDetail {
	s.UuidList = v
	return s
}

func (s *DescribeBackupPolicyResponseBodyBackupPolicyDetail) Validate() error {
	return dara.Validate(s)
}
