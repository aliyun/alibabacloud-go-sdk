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
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token that is used to obtain the next page of backup policies.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The backup policies.
	Policies []*DescribePoliciesV2ResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
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
	return dara.Validate(s)
}

type DescribePoliciesV2ResponseBodyPolicies struct {
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the backup policy was created. The value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1650248136
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The number of data sources that are bound to the backup policy.
	//
	// example:
	//
	// 5
	PolicyBindingCount *int64 `json:"PolicyBindingCount,omitempty" xml:"PolicyBindingCount,omitempty"`
	// The description of the backup policy.
	//
	// example:
	//
	// Data is backed up at 10:00:00 every day and replicated to the China (Shanghai) region for geo-redundancy.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The ID of the backup policy.
	//
	// example:
	//
	// po-000************bkz
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the backup policy.
	//
	// example:
	//
	// Daily Local Backup + Remote Backup
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// 	- **STANDARD**: the general backup policy. This type of policy applies to backups other than Elastic Compute Service (ECS) instance backup.
	//
	// 	- **UDM_ECS_ONLY**: the ECS instance backup policy. This type of policy applies only to ECS instance backup.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The rules in the backup policy.
	Rules []*DescribePoliciesV2ResponseBodyPoliciesRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The time when the backup policy was updated. The value is a UNIX timestamp. Unit: seconds.
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
	return dara.Validate(s)
}

type DescribePoliciesV2ResponseBodyPoliciesRules struct {
	// This parameter is returned only if the value of the **RuleType*	- parameter is **TRANSITION**. This parameter indicates the time when data is dumped from a backup vault to an archive vault. Unit: days.
	//
	// example:
	//
	// 30
	ArchiveDays *int64 `json:"ArchiveDays,omitempty" xml:"ArchiveDays,omitempty"`
	// This parameter is returned only if the value of the **RuleType*	- parameter is **BACKUP**. This parameter indicates the backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. It defines the data source filtering rule.
	DataSourceFilters []*DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters `json:"DataSourceFilters,omitempty" xml:"DataSourceFilters,omitempty" type:"Repeated"`
	// This parameter is returned only if the **PolicyType*	- is **UDM_ECS_ONLY**. This parameter indicates whether the immutable backup feature is enabled.
	//
	// example:
	//
	// true
	Immutable *bool `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	// Indicates whether the feature of keeping at least one backup version is enabled. Valid values:
	//
	// 	- **0**: The feature is disabled.
	//
	// 	- **1**: The feature is enabled.
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is returned only if the value of the **RuleType*	- parameter is **REPLICATION**. This parameter indicates the ID of the destination region.
	//
	// example:
	//
	// cn-shanghai
	ReplicationRegionId *string `json:"ReplicationRegionId,omitempty" xml:"ReplicationRegionId,omitempty"`
	// This parameter is returned only if the value of the **RuleType*	- parameter is **TRANSITION*	- or **REPLICATION**.
	//
	// 	- If the value of the **RuleType*	- parameter is **TRANSITION**, this parameter indicates the retention period of the backup data. Minimum value: 1. Unit: days.
	//
	// 	- If the value of the **RuleType*	- parameter is **REPLICATION**, this parameter indicates the retention period of remote backups. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// This parameter is returned only if the value of the **RuleType*	- parameter is **TRANSITION**. This parameter indicates the special retention rules.
	RetentionRules []*DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules `json:"RetentionRules,omitempty" xml:"RetentionRules,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// rule-000************f1e
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The type of the rule. Each backup policy must have at least one rule of the **BACKUP*	- type and only one rule of the **TRANSITION*	- type. Valid values:
	//
	// 	- **BACKUP**: backup rule
	//
	// 	- **TRANSITION**: lifecycle rule
	//
	// 	- **REPLICATION**: replication rule
	//
	// example:
	//
	// BACKUP
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is returned only if the value of the **RuleType*	- parameter is **BACKUP**. This parameter indicates the backup schedule settings. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- startTime: the time at which the system starts to run a backup job. The time follows the UNIX time format. Unit: seconds.
	//
	// 	- interval: the interval at which the system runs a backup job. The interval follows the ISO 8601 standard. For example, PT1H indicates an interval of 1 hour. P1D indicates an interval of one day.
	//
	// example:
	//
	// I|1648647166|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. It defines the resource tag filtering rule.
	TagFilters []*DescribePoliciesV2ResponseBodyPoliciesRulesTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
	// This parameter is returned only if the value of the RuleType parameter is BACKUP. The ID of the backup vault.
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
	return dara.Validate(s)
}

type DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters struct {
	// Deprecated.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// Data source type. The value range is as follows:
	//
	// - **UDM_ECS**: Indicates ECS server backup.
	//
	// - **OSS**: Indicates OSS backup.
	//
	// - **NAS**: Indicates Alibaba Cloud NAS backup.
	//
	// - **ECS_FILE**: Indicates ECS file backup.
	//
	// - **OTS**: Indicates Tablestore backup.
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

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *DescribePoliciesV2ResponseBodyPoliciesRulesDataSourceFilters) GetSourceType() *string {
	return s.SourceType
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
	return dara.Validate(s)
}

type DescribePoliciesV2ResponseBodyPoliciesRulesRetentionRules struct {
	// The type of the special retention rule. Valid values:
	//
	// 	- **WEEKLY**: weekly backups
	//
	// 	- **MONTHLY**: monthly backups
	//
	// 	- **YEARLY**: yearly backups
	//
	// example:
	//
	// YEARLY
	AdvancedRetentionType *string `json:"AdvancedRetentionType,omitempty" xml:"AdvancedRetentionType,omitempty"`
	// The special retention period of backups. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 730
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// Indicates which backup is retained based on the special retention rule. Only the first backup can be retained.
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
	// Tag key
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag matching rules, supporting: - **EQUAL**: Matches both the tag key and tag value. - **NOT**: Matches the tag key but not the tag value.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Tag value.
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
