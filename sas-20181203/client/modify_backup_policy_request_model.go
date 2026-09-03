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
	// The ID of the anti-ransomware policy to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anti-ransomware policy to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_name_A
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the policy to modify. The value is a JSON format character string that contains the following fields:
	//
	// - **Source**: The server folder to protect. To protect all folders, set this field to [].
	//
	// - **Include**: The file types to protect. Examples: "\\*.jpg" and "\\*.doc".
	//
	// - **Exclude**: The custom folders to exclude. For example, "/home/user" excludes the /home/user folder. Invoke the DescribeExcludeSystemPath operation to obtain all folders, and then add the folders that you want to exclude.
	//
	// - **Schedule**: The start time and interval of the data backup node. Specify a non-hourly time during off-peak hours.
	//
	//     - Example 1: I|1583216092|P21D indicates that the execute start time is 2020-03-03 14:14:52 and the interval is 3 weeks.
	//
	//     - Example 2: I|1583216092|PT24H indicates that the execute start time is 2020-03-03 14:14:52 and the interval is 24 hours.
	//
	// - **Retention**: The retention period of backup data. Unit: days. 7 indicates 1 week, 365 indicates 1 year, and -1 indicates permanent retention.
	//
	// - **SpeedLimiter**: The network bandwidth throttling for backup. For example, 12:15:15360|6:12:5120 indicates 15 MB from 12:00 to 15:00 and 5 MB from 6:00 to 12:00.
	//
	// For cloud-based servers connected to the internal network, do not limit the backup network bandwidth. To remove the network bandwidth throttling, set this parameter to an empty character string ("").
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Source":["home","admin"]}
	Policy map[string]interface{} `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region of the server for which you want to modify the policy.
	//
	// You can invoke the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the regions supported by the anti-ransomware feature.
	//
	// example:
	//
	// cn-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the policy. You can invoke the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the version.
	//
	// - **1.0.0**
	//
	// - **2.0.0**
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The method used to select assets. Valid values:
	//
	// - **ALL_MACHINE**: all assets
	//
	// >To cover all assets of the specified type, set this parameter to **ALL_MACHINE**. In this case, **UuidList*	- is invalid. Only one policy that covers all assets can exist for each server type.
	//
	// example:
	//
	// ALL_MACHINE
	SelectType *string `json:"SelectType,omitempty" xml:"SelectType,omitempty"`
	// The server type. Valid values:
	//
	// - **ALIYUN**: Alibaba Cloud server
	//
	// - **OUT_CLOUD**: non-Alibaba Cloud server
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
