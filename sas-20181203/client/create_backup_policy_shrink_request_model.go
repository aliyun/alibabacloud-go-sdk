// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateBackupPolicyShrinkRequest
	GetName() *string
	SetPolicyShrink(v string) *CreateBackupPolicyShrinkRequest
	GetPolicyShrink() *string
	SetPolicyRegionId(v string) *CreateBackupPolicyShrinkRequest
	GetPolicyRegionId() *string
	SetPolicyVersion(v string) *CreateBackupPolicyShrinkRequest
	GetPolicyVersion() *string
	SetUuidList(v []*string) *CreateBackupPolicyShrinkRequest
	GetUuidList() []*string
}

type CreateBackupPolicyShrinkRequest struct {
	// The name of the anti-ransomware policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// ServerBackUpPolicy01
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
	// This parameter is required.
	//
	// example:
	//
	// "{"IsDefault":1,"Include":[],"Source":[],"Schedule":"I|1648061040|PT24H","Retention":7,"SpeedLimiter":"","ExcludeSystemPath":true,"Exclude":["/bin/","/usr/bin/","/sbin/","/boot/","/proc/","/sys/","/srv/","/lib/","/selinux/","/usr/sbin/","/run/","/lib32/","/lib64/","/lost+found/","/var/lib/kubelet/","/var/lib/ntp/proc","/var/lib/container","Windows","Python27","Program Files (x86)","Program Files","Boot","$RECYCLE.BIN","System Volume Information","Users\\Administrator\\NTUSER.DAT*","ProgramData","pagefile.sys","Users\\Default\\NTUSER.DAT*","Users\\Administrator\\ntuser.*"],"UseVss":true}"
	PolicyShrink *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region ID of the server that is not deployed on Alibaba Cloud.
	//
	// >  We recommend that you specify the ID of the supported region that is the nearest to the location of the server. You can call the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the supported regions of the anti-ransomware feature.
	//
	// example:
	//
	// ch-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the anti-ransomware policy. Set the value to **2.0.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUIDs of the servers that you want to protect.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["3bb30859-b3b5-4f28-868f-b0892c98****", "3bb30859-b3b5-4f28-868f-b0892c98****"]
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s CreateBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateBackupPolicyShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *CreateBackupPolicyShrinkRequest) GetPolicyRegionId() *string {
	return s.PolicyRegionId
}

func (s *CreateBackupPolicyShrinkRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *CreateBackupPolicyShrinkRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *CreateBackupPolicyShrinkRequest) SetName(v string) *CreateBackupPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateBackupPolicyShrinkRequest) SetPolicyShrink(v string) *CreateBackupPolicyShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *CreateBackupPolicyShrinkRequest) SetPolicyRegionId(v string) *CreateBackupPolicyShrinkRequest {
	s.PolicyRegionId = &v
	return s
}

func (s *CreateBackupPolicyShrinkRequest) SetPolicyVersion(v string) *CreateBackupPolicyShrinkRequest {
	s.PolicyVersion = &v
	return s
}

func (s *CreateBackupPolicyShrinkRequest) SetUuidList(v []*string) *CreateBackupPolicyShrinkRequest {
	s.UuidList = v
	return s
}

func (s *CreateBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
