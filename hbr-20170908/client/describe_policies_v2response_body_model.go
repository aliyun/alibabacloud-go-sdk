// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePoliciesV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePoliciesV2ResponseBody
	GetCode() *string
	SetMaxResults(v int32) *DescribePoliciesV2ResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribePoliciesV2ResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribePoliciesV2ResponseBody
	GetNextToken() *string
	SetPolicies(v []*DescribePoliciesV2ResponseBodyPolicies) *DescribePoliciesV2ResponseBody
	GetPolicies() []*DescribePoliciesV2ResponseBodyPolicies
	SetRequestId(v string) *DescribePoliciesV2ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePoliciesV2ResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribePoliciesV2ResponseBody
	GetTotalCount() *int64
}

type DescribePoliciesV2ResponseBody struct {
	// The response code. 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of results per query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message. The value "successful" is returned for a successful request. An error message is returned for a failed request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token required to retrieve the next page of policies.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of policies.
	Policies []*DescribePoliciesV2ResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePoliciesV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePoliciesV2ResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePoliciesV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePoliciesV2ResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePoliciesV2ResponseBody) GetPolicies() []*DescribePoliciesV2ResponseBodyPolicies {
	return s.Policies
}

func (s *DescribePoliciesV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePoliciesV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePoliciesV2ResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePoliciesV2ResponseBody) SetCode(v string) *DescribePoliciesV2ResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetMaxResults(v int32) *DescribePoliciesV2ResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetMessage(v string) *DescribePoliciesV2ResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetNextToken(v string) *DescribePoliciesV2ResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetPolicies(v []*DescribePoliciesV2ResponseBodyPolicies) *DescribePoliciesV2ResponseBody {
	s.Policies = v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetRequestId(v string) *DescribePoliciesV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetSuccess(v bool) *DescribePoliciesV2ResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) SetTotalCount(v int64) *DescribePoliciesV2ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePoliciesV2ResponseBody) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePoliciesV2ResponseBodyPolicies struct {
	// The user business status.
	//
	// example:
	//
	// ACTIVE
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The creation time. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1650248136
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The number of data sources bound to the policy.
	//
	// example:
	//
	// 5
	PolicyBindingCount *int64 `json:"PolicyBindingCount,omitempty" xml:"PolicyBindingCount,omitempty"`
	// The policy description.
	//
	// example:
	//
	// Back up every day at 10:00 AM and replicate to Shanghai
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// po-000************bkz
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Daily backup + geo-redundancy backup
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// - **STANDARD**: general backup policy. Supports backing up data sources other than ECS instance backup.
	//
	// - **UDM_ECS_ONLY**: ECS instance backup policy. Supports backing up only ECS instances.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of policy rules.
	Rules []*DescribePoliciesV2ResponseBodyPoliciesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The update time. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1662080404
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribePoliciesV2ResponseBodyPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetPolicyBindingCount() *int64 {
	return s.PolicyBindingCount
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetRules() []*DescribePoliciesV2ResponseBodyPoliciesRules {
	return s.Rules
}

func (s *DescribePoliciesV2ResponseBodyPolicies) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetBusinessStatus(v string) *DescribePoliciesV2ResponseBodyPolicies {
	s.BusinessStatus = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetCreatedTime(v int64) *DescribePoliciesV2ResponseBodyPolicies {
	s.CreatedTime = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetPolicyBindingCount(v int64) *DescribePoliciesV2ResponseBodyPolicies {
	s.PolicyBindingCount = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetPolicyDescription(v string) *DescribePoliciesV2ResponseBodyPolicies {
	s.PolicyDescription = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetPolicyId(v string) *DescribePoliciesV2ResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetPolicyName(v string) *DescribePoliciesV2ResponseBodyPolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetPolicyType(v string) *DescribePoliciesV2ResponseBodyPolicies {
	s.PolicyType = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetRules(v []*DescribePoliciesV2ResponseBodyPoliciesRules) *DescribePoliciesV2ResponseBodyPolicies {
	s.Rules = v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) SetUpdatedTime(v int64) *DescribePoliciesV2ResponseBodyPolicies {
	s.UpdatedTime = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPolicies) Validate() error {
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePoliciesV2ResponseBodyPoliciesRules struct {
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The number of days after which the backup is converted to archive storage. Unit: days.
	//
	// example:
	//
	// 30
	ArchiveDays *int64 `json:"ArchiveDays,omitempty" xml:"ArchiveDays,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **BACKUP**. The backup type. The value is **COMPLETE**, which indicates a full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. The data source filter rules.
	DataSourceFilters []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters `json:"DataSourceFilters,omitempty" xml:"DataSourceFilters,omitempty" type:"Repeated"`
	// This parameter is valid only when **PolicyType*	- is set to **UDM_ECS_ONLY**. Specifies whether to enable backup locking.
	//
	// example:
	//
	// true
	Immutable *bool `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	// Specifies whether to retain at least one backup version. Valid values:
	//
	// - **0**: Do not retain.
	//
	// - **1**: Retain.
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **REPLICATION**. The destination region ID for replication.
	//
	// example:
	//
	// cn-shanghai
	ReplicationRegionId *string `json:"ReplicationRegionId,omitempty" xml:"ReplicationRegionId,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TRANSITION*	- or **REPLICATION**.
	//
	// - **RuleType*	- is set to **TRANSITION**: the retention period of the backup. Minimum value: 1. Unit: days.
	//
	// - **RuleType*	- is set to **REPLICATION**: the retention period of the geo-redundancy backup. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The list of special retention rules.
	RetentionRules []*DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules `json:"RetentionRules,omitempty" xml:"RetentionRules,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// rule-000************f1e
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule type. Each policy must have at least one **BACKUP*	- rule and exactly one **TRANSITION*	- rule. Valid values:
	//
	// - **BACKUP**: backup rule.
	//
	// - **TRANSITION**: lifecycle rule.
	//
	// - **REPLICATION**: replication rule.
	//
	// example:
	//
	// BACKUP
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **BACKUP**. The backup schedule. Optional format: `I|{startTime}|{interval}`. This indicates that a backup job is executed at every {interval} starting from {startTime}. Backup jobs for past time periods are not compensated. If the previous backup job is not completed, the next backup job is not triggered. For example, `I|1631685600|P1D` indicates that a backup is performed once a day starting from 2021-09-15 14:00:00.
	//
	// 	- startTime: the start time of the backup. UNIX timestamp, in seconds.
	//
	// 	- interval: the ISO 8601 time interval. For example, PT1H indicates an interval of one hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1648647166|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. The resource tag filter rules.
	TagFilters []*DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
	// This parameter is required only when RuleType is set to BACKUP. The backup vault ID.
	//
	// example:
	//
	// v-000**************kgm
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s DescribePoliciesV2ResponseBodyPoliciesRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBodyPoliciesRules) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetArchiveDays() *int64 {
	return s.ArchiveDays
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetDataSourceFilters() []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters {
	return s.DataSourceFilters
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetImmutable() *bool {
	return s.Immutable
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetReplicationRegionId() *string {
	return s.ReplicationRegionId
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetRetentionRules() []*DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules {
	return s.RetentionRules
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetSchedule() *string {
	return s.Schedule
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetTagFilters() []*DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters {
	return s.TagFilters
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetArchiveDays(v int64) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.ArchiveDays = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetBackupType(v string) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.BackupType = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetDataSourceFilters(v []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.DataSourceFilters = v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetImmutable(v bool) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.Immutable = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetKeepLatestSnapshots(v int64) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetReplicationRegionId(v string) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.ReplicationRegionId = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetRetention(v int64) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.Retention = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetRetentionRules(v []*DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.RetentionRules = v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetRuleId(v string) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.RuleId = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetRuleType(v string) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.RuleType = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetSchedule(v string) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.Schedule = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetTagFilters(v []*DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.TagFilters = v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) SetVaultId(v string) *DescribePoliciesV2ResponseBodyPoliciesRules {
	s.VaultId = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRules) Validate() error {
	if s.DataSourceFilters != nil {
		for _, item := range s.DataSourceFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RetentionRules != nil {
		for _, item := range s.RetentionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagFilters != nil {
		for _, item := range s.TagFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters struct {
	AccountScope *string                                                                 `json:"AccountScope,omitempty" xml:"AccountScope,omitempty"`
	Accounts     []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS instance backup.
	//
	// - **OSS**: OSS backup.
	//
	// - **NAS**: Alibaba Cloud NAS backup.
	//
	// - **ECS_FILE**: ECS File Backup Essential Edition.
	//
	// - **OTS**: Tablestore backup.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GetAccountScope() *string {
	return s.AccountScope
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GetAccounts() []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts {
	return s.Accounts
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) SetAccountScope(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters {
	s.AccountScope = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) SetAccounts(v []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters {
	s.Accounts = v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) SetDataSourceIds(v []*string) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters {
	s.DataSourceIds = v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) SetSourceType(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters {
	s.SourceType = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) SetCrossAccountRoleName(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) SetCrossAccountType(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts {
	s.CrossAccountType = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) SetCrossAccountUserId(v int64) *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFiltersAccounts) Validate() error {
	return dara.Validate(s)
}

type DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules struct {
	// The type of the special retention rule. Valid values:
	//
	// - **WEEKLY**: weekly backup.
	//
	// - **MONTHLY**: monthly backup.
	//
	// - **YEARLY**: yearly backup.
	//
	// example:
	//
	// YEARLY
	AdvancedRetentionType *string `json:"AdvancedRetentionType,omitempty" xml:"AdvancedRetentionType,omitempty"`
	// The special retention period of the backup. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 730
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The backup to which the rule applies. Currently, only the first backup is supported. The value is 1.
	//
	// example:
	//
	// 1
	WhichSnapshot *int64 `json:"WhichSnapshot,omitempty" xml:"WhichSnapshot,omitempty"`
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) GetAdvancedRetentionType() *string {
	return s.AdvancedRetentionType
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) GetWhichSnapshot() *int64 {
	return s.WhichSnapshot
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) SetAdvancedRetentionType(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules {
	s.AdvancedRetentionType = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) SetRetention(v int64) *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules {
	s.Retention = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) SetWhichSnapshot(v int64) *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules {
	s.WhichSnapshot = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules) Validate() error {
	return dara.Validate(s)
}

type DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag matching rule. Valid values:
	//
	// - **EQUAL**: matches both the tag key and the tag value.
	//
	// - **NOT**: matches the tag key but not the tag value.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The tag value. An empty value indicates any value.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) GoString() string {
	return s.String()
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) GetKey() *string {
	return s.Key
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) GetOperator() *string {
	return s.Operator
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) GetValue() *string {
	return s.Value
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) SetKey(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters {
	s.Key = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) SetOperator(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters {
	s.Operator = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) SetValue(v string) *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters {
	s.Value = &v
	return s
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters) Validate() error {
	return dara.Validate(s)
}
