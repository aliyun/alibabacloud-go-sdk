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
	PolicyShrink *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
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

func (s *ModifyBackupPolicyShrinkRequest) SetUuidList(v []*string) *ModifyBackupPolicyShrinkRequest {
	s.UuidList = v
	return s
}

func (s *ModifyBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
