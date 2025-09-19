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
	// The configurations of the anti-ransomware policy that you want to modify. The value is a JSON string that contains the following fields:
	//
	// 	- **Source**: the directory that you want to protect. If you want to protect all directories, set this field to brackets [].
	//
	// 	- **Include**: the format of the file that you want to protect. Examples: \\*.jpg and \\*.doc.
	//
	// 	- **Exclude**: the directory that you want to exclude from the anti-ransomware policy. You can call the DescribeExcludeSystemPath operation to query all directories and then specify the directory that you want to exclude. Example: /home/user.
	//
	// 	- **Schedule**: the start time and interval of a data backup task. We recommend that you specify a start time that begins during off-peak hours but does not start on the hour.
	//
	//     	- If you set this field to I|1583216092|P21D, the data backup task starts from 2020-03-03 14:14:52, and the task is executed at an interval of three weeks.
	//
	//     	- If you set this field to I|1583216092|PT24H, the data backup task starts from 2020-03-03 14:14:52, and the task is executed at an interval of 24 hours.
	//
	// 	- **Retention**: the period during which backup data is retained. Unit: day. If you set this field to 7, backup data is retained for a week. If you set this field to 365, backup data is retained for a year. If you set this field to -1, backup data is permanently retained.
	//
	// 	- **SpeedLimiter**: the limit on the network bandwidth for data backup tasks. If you set this field to 12:15:15360|6:12:5120, the maximum bandwidth for a data backup task is 15 Mbit/s from 12:00 to 15:00 and 5 Mbit/s from 06:00 to 12:00.
	//
	// If you back up data on an Elastic Compute Service (ECS) instance that is connected over an internal network, we recommend that you leave this field empty. If this field is left empty, the bandwidth for data backup tasks is unlimited.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Source":["home","admin"]}
	Policy map[string]interface{} `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region ID of the server to which the anti-ransomware policy is applied.
	//
	// You can call the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the regions in which the anti-ransomware feature is supported.
	//
	// example:
	//
	// cn-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the anti-ransomware policy. You can call the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the versions of anti-ransomware policies.
	//
	// 	- **1.0.0**
	//
	// 	- **2.0.0**
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUIDs of the servers to which the anti-ransomware policy is applied.
	//
	// This parameter is required.
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

func (s *ModifyBackupPolicyRequest) SetUuidList(v []*string) *ModifyBackupPolicyRequest {
	s.UuidList = v
	return s
}

func (s *ModifyBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
