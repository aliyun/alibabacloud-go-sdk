// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ModifyBackupPolicyRequest
	GetId() *int64
	SetName(v string) *ModifyBackupPolicyRequest
	GetName() *string
	SetPolicy(v map[string]interface{}) *ModifyBackupPolicyRequest
	GetPolicy() map[string]interface{}
	SetPolicyRegionId(v string) *ModifyBackupPolicyRequest
	GetPolicyRegionId() *string
	SetPolicyVersion(v string) *ModifyBackupPolicyRequest
	GetPolicyVersion() *string
	SetSelectType(v string) *ModifyBackupPolicyRequest
	GetSelectType() *string
	SetServerType(v string) *ModifyBackupPolicyRequest
	GetServerType() *string
	SetUuidList(v []*string) *ModifyBackupPolicyRequest
	GetUuidList() []*string
}

type ModifyBackupPolicyRequest struct {
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
	Policy map[string]interface{} `json:"Policy,omitempty" xml:"Policy,omitempty"`
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

func (s ModifyBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPolicyRequest) GetId() *int64 {
	return s.Id
}

func (s *ModifyBackupPolicyRequest) GetName() *string {
	return s.Name
}

func (s *ModifyBackupPolicyRequest) GetPolicy() map[string]interface{} {
	return s.Policy
}

func (s *ModifyBackupPolicyRequest) GetPolicyRegionId() *string {
	return s.PolicyRegionId
}

func (s *ModifyBackupPolicyRequest) GetPolicyVersion() *string {
	return s.PolicyVersion
}

func (s *ModifyBackupPolicyRequest) GetSelectType() *string {
	return s.SelectType
}

func (s *ModifyBackupPolicyRequest) GetServerType() *string {
	return s.ServerType
}

func (s *ModifyBackupPolicyRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *ModifyBackupPolicyRequest) SetId(v int64) *ModifyBackupPolicyRequest {
	s.Id = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetName(v string) *ModifyBackupPolicyRequest {
	s.Name = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPolicy(v map[string]interface{}) *ModifyBackupPolicyRequest {
	s.Policy = v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPolicyRegionId(v string) *ModifyBackupPolicyRequest {
	s.PolicyRegionId = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetPolicyVersion(v string) *ModifyBackupPolicyRequest {
	s.PolicyVersion = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetSelectType(v string) *ModifyBackupPolicyRequest {
	s.SelectType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetServerType(v string) *ModifyBackupPolicyRequest {
	s.ServerType = &v
	return s
}

func (s *ModifyBackupPolicyRequest) SetUuidList(v []*string) *ModifyBackupPolicyRequest {
	s.UuidList = v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
