// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyBackupPolicyShrinkRequest
	GetId() *int64
	SetName(v string) *ModifyBackupPolicyShrinkRequest
	GetName() *string
	SetPolicyShrink(v string) *ModifyBackupPolicyShrinkRequest
	GetPolicyShrink() *string
	SetPolicyRegionId(v string) *ModifyBackupPolicyShrinkRequest
	GetPolicyRegionId() *string
	SetPolicyVersion(v string) *ModifyBackupPolicyShrinkRequest
	GetPolicyVersion() *string
	SetSelectType(v string) *ModifyBackupPolicyShrinkRequest
	GetSelectType() *string
	SetServerType(v string) *ModifyBackupPolicyShrinkRequest
	GetServerType() *string
	SetUuidList(v []*string) *ModifyBackupPolicyShrinkRequest
	GetUuidList() []*string
}

type ModifyBackupPolicyShrinkRequest struct {
	// The ID of the anti-ransomware policy that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anti-ransomware policy that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_name_A
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the policy that you want to modify. The value is a JSON string that contains the following fields:
	//
	// - **Source**: the server directories to be protected. To protect all directories, set this field to [].
	//
	// - **Include**: the types of files to be protected. Examples: "\\*.jpg" and "\\*.doc".
	//
	// - **Exclude**: the custom directories to be excluded. To exclude the "/home/user" directory, call the DescribeExcludeSystemPath operation to obtain all directories and then add the directories to be excluded. Example: exclude the "/home/user" directory.
	//
	// - **Schedule**: the start time and interval of the data backup task. Specify a time during off-peak hours and avoid exact hours.
	//
	//     - Example 1: I|1583216092|P21D indicates that the start time is 2020-03-03 14:14:52 and the interval is 3 weeks.
	//
	//     - Example 2: I|1583216092|PT24H indicates that the start time is 2020-03-03 14:14:52 and the interval is 24 hours.
	//
	// - **Retention**: the retention period of backup data. Unit: days. A value of 7 indicates 1 week, 365 indicates 1 year, and -1 indicates permanent retention.
	//
	// - **SpeedLimiter**: the bandwidth limit for the backup network. Example: 12:15:15360|6:12:5120 indicates a limit of 15 MB from 12:00 to 15:00 and 5 MB from 06:00 to 12:00.
	//
	// If the cloud server is connected to the internal network, do not limit the backup network bandwidth. To remove the bandwidth limit, set this parameter to an empty string ("").
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Source":["home","admin"]}
	PolicyShrink *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region where the server to which the policy applies resides.
	//
	// You can call the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the regions supported by the anti-ransomware feature.
	//
	// example:
	//
	// cn-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the policy. You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the version.
	//
	// - **1.0.0**
	//
	// - **2.0.0**
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The method to cover assets. Valid values:
	//
	// - **ALL_MACHINE**: all assets
	//
	// > To cover all assets of this type, set this parameter to **ALL_MACHINE**. In this case, **UuidList*	- does not take effect. Only one policy that covers all assets can exist for the same server type.
	//
	// example:
	//
	// ALL_MACHINE
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The server type. Valid values:
	//
	// - **ALIYUN**: Alibaba Cloud server
	//
	// - **OUT_CLOUD**: third-party server
	//
	// - **TRIPARTITE**: simple application server
	//
	// example:
	//
	// ALIYUN
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
	// The list of UUIDs of the servers protected by the policy.
	//
	// example:
	//
	// ["3bb30859-b3b5-4f28-868f-b0892c98****", "3bb30859-b3b5-4f28-868f-b0892c98****"]
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s ModifyBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyBackupPolicyShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyBackupPolicyShrinkRequest) GetPolicyShrink() *string {
	return s.PolicyShrink
}

func (s *ModifyBackupPolicyShrinkRequest) GetPolicyRegionId() *string {
	return s.PolicyRegionId
}

func (s *ModifyBackupPolicyShrinkRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *ModifyBackupPolicyShrinkRequest) GetSelectType() *string {
	return s.SelectType
}

func (s *ModifyBackupPolicyShrinkRequest) GetServerType() *string {
	return s.ServerType
}

func (s *ModifyBackupPolicyShrinkRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ModifyBackupPolicyShrinkRequest) SetId(v int64) *ModifyBackupPolicyShrinkRequest {
	s.Id = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetName(v string) *ModifyBackupPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetPolicyShrink(v string) *ModifyBackupPolicyShrinkRequest {
	s.PolicyShrink = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetPolicyRegionId(v string) *ModifyBackupPolicyShrinkRequest {
	s.PolicyRegionId = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetPolicyVersion(v string) *ModifyBackupPolicyShrinkRequest {
	s.PolicyVersion = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetSelectType(v string) *ModifyBackupPolicyShrinkRequest {
	s.SelectType = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetServerType(v string) *ModifyBackupPolicyShrinkRequest {
	s.ServerType = &v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) SetUuidList(v []*string) *ModifyBackupPolicyShrinkRequest {
	s.UuidList = v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
