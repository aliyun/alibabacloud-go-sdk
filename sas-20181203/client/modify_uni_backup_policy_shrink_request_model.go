// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUniBackupPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyUniBackupPolicyShrinkRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ModifyUniBackupPolicyShrinkRequest
	GetAccountPassword() *string
	SetFullPlanShrink(v string) *ModifyUniBackupPolicyShrinkRequest
	GetFullPlanShrink() *string
	SetIncPlanShrink(v string) *ModifyUniBackupPolicyShrinkRequest
	GetIncPlanShrink() *string
	SetPolicyId(v int64) *ModifyUniBackupPolicyShrinkRequest
	GetPolicyId() *int64
	SetPolicyName(v string) *ModifyUniBackupPolicyShrinkRequest
	GetPolicyName() *string
	SetPolicyStatus(v string) *ModifyUniBackupPolicyShrinkRequest
	GetPolicyStatus() *string
	SetRetention(v int32) *ModifyUniBackupPolicyShrinkRequest
	GetRetention() *int32
	SetSpeedLimiter(v int64) *ModifyUniBackupPolicyShrinkRequest
	GetSpeedLimiter() *int64
}

type ModifyUniBackupPolicyShrinkRequest struct {
	// The name of the database account.
	//
	// example:
	//
	// sa
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The password of the database account.
	//
	// example:
	//
	// Sa@****
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty"`
	// The policy for full backup. The value of this parameter is a JSON string that contains the following fields:
	//
	// 	- **start**: the start time of a backup task
	//
	// 	- **interval**: the interval of backup tasks
	//
	// 	- **type**: the unit of the interval
	//
	// 	- **days**: the days of a week on which a backup task is performed
	//
	// example:
	//
	// {"days":[4],"interval":1,"planType":"weekly","startTime":"22:00:00"}
	FullPlanShrink *string `json:"FullPlan,omitempty" xml:"FullPlan,omitempty"`
	// The policy for incremental backup. The value of this parameter is a JSON string that contains the following fields:
	//
	// 	- **start**: the start time of a backup task
	//
	// 	- **interval**: the interval of backup tasks
	//
	// 	- **type**: the unit of the interval
	//
	// 	- **days**: the days of a week on which a backup task is performed
	//
	// example:
	//
	// {"interval":1,"planType":"daily","startTime":"23:30:00"}
	IncPlanShrink *string `json:"IncPlan,omitempty" xml:"IncPlan,omitempty"`
	// The ID of the anti-ransomware policy.
	//
	// > You can call the [DescribeUniBackupPolicies](~~DescribeUniBackupPolicies~~) operation to query the IDs of anti-ransomware policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// databak
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the anti-ransomware policy. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// example:
	//
	// enabled
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The retention period of the backup snapshot.
	//
	// example:
	//
	// 7
	Retention *int32 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The maximum network bandwidth that is allowed during data backup. Unit: bytes.
	//
	// example:
	//
	// 1048576
	SpeedLimiter *int64 `json:"SpeedLimiter,omitempty" xml:"SpeedLimiter,omitempty"`
}

func (s ModifyUniBackupPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUniBackupPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetFullPlanShrink() *string {
	return s.FullPlanShrink
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetIncPlanShrink() *string {
	return s.IncPlanShrink
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifyUniBackupPolicyShrinkRequest) GetSpeedLimiter() *int64 {
	return s.SpeedLimiter
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetAccountName(v string) *ModifyUniBackupPolicyShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetAccountPassword(v string) *ModifyUniBackupPolicyShrinkRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetFullPlanShrink(v string) *ModifyUniBackupPolicyShrinkRequest {
	s.FullPlanShrink = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetIncPlanShrink(v string) *ModifyUniBackupPolicyShrinkRequest {
	s.IncPlanShrink = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetPolicyId(v int64) *ModifyUniBackupPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetPolicyName(v string) *ModifyUniBackupPolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetPolicyStatus(v string) *ModifyUniBackupPolicyShrinkRequest {
	s.PolicyStatus = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetRetention(v int32) *ModifyUniBackupPolicyShrinkRequest {
	s.Retention = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) SetSpeedLimiter(v int64) *ModifyUniBackupPolicyShrinkRequest {
	s.SpeedLimiter = &v
	return s
}

func (s *ModifyUniBackupPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
