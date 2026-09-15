// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateBackupPolicyRequest
	GetName() *string
	SetPolicy(v map[string]interface{}) *CreateBackupPolicyRequest
	GetPolicy() map[string]interface{}
	SetPolicyRegionId(v string) *CreateBackupPolicyRequest
	GetPolicyRegionId() *string
	SetPolicyVersion(v string) *CreateBackupPolicyRequest
	GetPolicyVersion() *string
	SetSelectType(v string) *CreateBackupPolicyRequest
	GetSelectType() *string
	SetServerType(v string) *CreateBackupPolicyRequest
	GetServerType() *string
	SetUuidList(v []*string) *CreateBackupPolicyRequest
	GetUuidList() []*string
}

type CreateBackupPolicyRequest struct {
	// The name of the protection policy to create. Set this parameter to the desired policy name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ServerBackUpPolicy01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// - **IsDefault**: The type of the protection policy to create. Valid values:
	//
	//     - **1**: Recommended policy.
	//
	//     - **0**: Custom policy.
	//
	// - **Include**: The file types to protect. To protect all file types, set this parameter to [].
	//
	// - **Source**: The server folders to protect. To protect all folders, set this parameter to [].
	//
	// - **ExcludeSystemPath**: Specifies whether to exclude specified folders. Set this parameter to **true*	- to exclude folders. If you do not want to exclude folders, you do not need to set this parameter.
	//
	// - **Exclude**: The specified protection folder addresses. If you do not want to specify protection folder addresses, set this parameter to [].
	//
	// - **Schedule**: The time and interval at which the data backup task is scheduled to run. Specify a non-peak hour that is not on the hour. Examples:
	//
	//     - Example 1: I|1583216092|P21D indicates that the data backup starts at 2020-03-03 14:14:52 and the backup policy runs at an interval of 3 weeks.
	//
	//     - Example 2: I|1583216092|PT24H indicates that the data backup starts at 2020-03-03 14:14:52 and the backup policy runs at an interval of 24 hours.
	//
	// - **Retention**: The retention period of backup data, in days. The value 7 indicates 1 week, 365 indicates 1 year, and -1 indicates permanent retention.
	//
	// - **SpeedLimiter**: The network bandwidth throttling for backup. For example, 0:24:30720 indicates that the network bandwidth throttling for backup is 30 MB/s from 00:00 to 24:00.
	//
	// - **UseVss**: Specifies whether to enable the Volume Shadow Copy Service (VSS) feature for Windows. Valid values:
	//
	//     - **true**: Enabled.
	//
	//     - **false**: Not enabled.
	//
	// > The VSS (Windows) feature applies only to Windows systems. After this feature is enabled, the issue of individual file backup failures due to process occupation is effectively reduced. We recommend that you enable this feature. After this feature is enabled, file backup for exFAT and FAT32 disk formats is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// "{"IsDefault":1,"Include":[],"Source":[],"Schedule":"I|1648061040|PT24H","Retention":7,"SpeedLimiter":"","ExcludeSystemPath":true,"Exclude":["/bin/","/usr/bin/","/sbin/","/boot/","/proc/","/sys/","/srv/","/lib/","/selinux/","/usr/sbin/","/run/","/lib32/","/lib64/","/lost+found/","/var/lib/kubelet/","/var/lib/ntp/proc","/var/lib/container","Windows","Python27","Program Files (x86)","Program Files","Boot","$RECYCLE.BIN","System Volume Information","Users\\Administrator\\NTUSER.DAT*","ProgramData","pagefile.sys","Users\\Default\\NTUSER.DAT*","Users\\Administrator\\ntuser.*"],"UseVss":true}"
	Policy map[string]interface{} `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region ID of the non-Alibaba Cloud server.
	//
	// > Call the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the regions supported by the anti-ransomware feature, and then select the supported region closest to the region where your non-Alibaba Cloud server resides.
	//
	// example:
	//
	// ch-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the protection policy. Set the value to **2.0.0**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The method used to cover assets. Valid values:
	//
	// - **ALL_MACHINE**: All assets.
	//
	// > To cover all assets of this type, set this parameter to **ALL_MACHINE**. In this case, **UuidList*	- is invalid. Only one policy that covers all assets can exist for each server type.
	//
	// example:
	//
	// ALL_MACHINE
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The server type. Valid values:
	//
	// - **ALIYUN**: Alibaba Cloud server.
	//
	// - **OUT_CLOUD**: Non-Alibaba Cloud server.
	//
	// - **TRIPARTITE**: Simple application server.
	//
	// example:
	//
	// ALIYUN
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The UUIDs of the servers to protect.
	//
	// example:
	//
	// ["3bb30859-b3b5-4f28-868f-b0892c98****", "3bb30859-b3b5-4f28-868f-b0892c98****"]
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s CreateBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateBackupPolicyRequest) GetPolicy() map[string]interface{} {
	return s.Policy
}

func (s *CreateBackupPolicyRequest) GetPolicyRegionId() *string {
	return s.PolicyRegionId
}

func (s *CreateBackupPolicyRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *CreateBackupPolicyRequest) GetSelectType() *string {
	return s.SelectType
}

func (s *CreateBackupPolicyRequest) GetServerType() *string {
	return s.ServerType
}

func (s *CreateBackupPolicyRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *CreateBackupPolicyRequest) SetName(v string) *CreateBackupPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetPolicy(v map[string]interface{}) *CreateBackupPolicyRequest {
	s.Policy = v
	return s
}

func (s *CreateBackupPolicyRequest) SetPolicyRegionId(v string) *CreateBackupPolicyRequest {
	s.PolicyRegionId = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetPolicyVersion(v string) *CreateBackupPolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetSelectType(v string) *CreateBackupPolicyRequest {
	s.SelectType = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetServerType(v string) *CreateBackupPolicyRequest {
	s.ServerType = &v
	return s
}

func (s *CreateBackupPolicyRequest) SetUuidList(v []*string) *CreateBackupPolicyRequest {
	s.UuidList = v
	return s
}

func (s *CreateBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
