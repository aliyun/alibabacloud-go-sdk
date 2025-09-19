// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUniBackupPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyUniBackupPolicyRequest
	GetAccountName() *string
	SetAccountPassword(v string) *ModifyUniBackupPolicyRequest
	GetAccountPassword() *string
	SetFullPlan(v map[string]interface{}) *ModifyUniBackupPolicyRequest
	GetFullPlan() map[string]interface{}
	SetIncPlan(v map[string]interface{}) *ModifyUniBackupPolicyRequest
	GetIncPlan() map[string]interface{}
	SetPolicyId(v int64) *ModifyUniBackupPolicyRequest
	GetPolicyId() *int64
	SetPolicyName(v string) *ModifyUniBackupPolicyRequest
	GetPolicyName() *string
	SetPolicyStatus(v string) *ModifyUniBackupPolicyRequest
	GetPolicyStatus() *string
	SetRetention(v int32) *ModifyUniBackupPolicyRequest
	GetRetention() *int32
	SetSpeedLimiter(v int64) *ModifyUniBackupPolicyRequest
	GetSpeedLimiter() *int64
}

type ModifyUniBackupPolicyRequest struct {
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
	FullPlan map[string]interface{} `json:"FullPlan,omitempty" xml:"FullPlan,omitempty"`
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
	IncPlan map[string]interface{} `json:"IncPlan,omitempty" xml:"IncPlan,omitempty"`
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

func (s ModifyUniBackupPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUniBackupPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyUniBackupPolicyRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyUniBackupPolicyRequest) GetAccountPassword() *string {
	return s.AccountPassword
}

func (s *ModifyUniBackupPolicyRequest) GetFullPlan() map[string]interface{} {
	return s.FullPlan
}

func (s *ModifyUniBackupPolicyRequest) GetIncPlan() map[string]interface{} {
	return s.IncPlan
}

func (s *ModifyUniBackupPolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *ModifyUniBackupPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ModifyUniBackupPolicyRequest) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *ModifyUniBackupPolicyRequest) GetRetention() *int32 {
	return s.Retention
}

func (s *ModifyUniBackupPolicyRequest) GetSpeedLimiter() *int64 {
	return s.SpeedLimiter
}

func (s *ModifyUniBackupPolicyRequest) SetAccountName(v string) *ModifyUniBackupPolicyRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetAccountPassword(v string) *ModifyUniBackupPolicyRequest {
	s.AccountPassword = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetFullPlan(v map[string]interface{}) *ModifyUniBackupPolicyRequest {
	s.FullPlan = v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetIncPlan(v map[string]interface{}) *ModifyUniBackupPolicyRequest {
	s.IncPlan = v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetPolicyId(v int64) *ModifyUniBackupPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetPolicyName(v string) *ModifyUniBackupPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetPolicyStatus(v string) *ModifyUniBackupPolicyRequest {
	s.PolicyStatus = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetRetention(v int32) *ModifyUniBackupPolicyRequest {
	s.Retention = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) SetSpeedLimiter(v int64) *ModifyUniBackupPolicyRequest {
	s.SpeedLimiter = &v
	return s
}

func (s *ModifyUniBackupPolicyRequest) Validate() error {
	return dara.Validate(s)
}
