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
	// The ID of the anti-ransomware mitigation policy to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the anti-ransomware mitigation policy to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy_name_A
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the mitigation policy to modify. The value is a JSON-format string that contains the following fields:
	//
	// - **Source**: The server folder to protect. To protect all folders, set this field to [].
	//
	// - **Include**: The file types to protect. Examples: "\\*.jpg" and "\\*.doc".
	//
	// - **Exclude**: The custom folders to exclude. For example, exclude the folder "/home/user". Invoke the DescribeExcludeSystemPath operation to obtain all folders, and then add the folders that you want to exclude.
	//
	// - **Schedule**: The start time and interval of the data backup task. Specify a non-hourly time during off-peak hours.
	//
	//     - Example 1: I|1583216092|P21D indicates that the start time is 2020-03-03 14:14:52 and the interval is 3 weeks.
	//
	//     - Example 2: I|1583216092|PT24H indicates that the start time is 2020-03-03 14:14:52 and the interval is 24 hours.
	//
	// - **Retention**: The retention period of backup data. Unit: days. 7 indicates 1 week, 365 indicates 1 year, and -1 indicates permanent retention.
	//
	// - **SpeedLimiter**: The network bandwidth throttling for backup. Example: 12:15:15360|6:12:5120 indicates 15 MB from 12:00 to 15:00 and 5 MB from 6:00 to 12:00.
	//
	// Cloud-hosted servers connect through the internal network. Do not limit the backup network bandwidth. To remove the bandwidth limit, set this parameter to an empty string ("").
	//
	// This parameter is required.
	//
	// example:
	//
	// {"Source":["home","admin"]}
	PolicyShrink *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region of the server for which you want to modify the mitigation policy.
	//
	// You can invoke the [DescribeSupportRegion](~~DescribeSupportRegion~~) operation to query the regions supported by the anti-ransomware feature.
	//
	// example:
	//
	// cn-hangzhou
	PolicyRegionId *string `json:"PolicyRegionId,omitempty" xml:"PolicyRegionId,omitempty"`
	// The version of the mitigation policy. You can invoke the [DescribeBackupPolicies](~~DescribeBackupPolicies~~) operation to query the version.
	//
	// - **1.0.0**
	//
	// - **2.0.0**
	//
	// example:
	//
	// 2.0.0
	PolicyVersion *string `json:"PolicyVersion,omitempty" xml:"PolicyVersion,omitempty"`
	// The UUIDs of the servers protected by the mitigation policy.
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
