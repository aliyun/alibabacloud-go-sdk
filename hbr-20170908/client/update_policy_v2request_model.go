// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyDescription(v string) *UpdatePolicyV2Request
	GetPolicyDescription() *string
	SetPolicyId(v string) *UpdatePolicyV2Request
	GetPolicyId() *string
	SetPolicyName(v string) *UpdatePolicyV2Request
	GetPolicyName() *string
	SetRules(v []*UpdatePolicyV2RequestRules) *UpdatePolicyV2Request
	GetRules() []*UpdatePolicyV2RequestRules
}

type UpdatePolicyV2Request struct {
	// The policy description.
	//
	// example:
	//
	// Back up once every day at 10:00 AM, with cross-region backup to Shanghai.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// po-000************viy
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Daily backup + geo-redundancy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The list of policy rules.
	Rules []*UpdatePolicyV2RequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s UpdatePolicyV2Request) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2Request) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2Request) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *UpdatePolicyV2Request) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdatePolicyV2Request) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyV2Request) GetRules() []*UpdatePolicyV2RequestRules {
	return s.Rules
}

func (s *UpdatePolicyV2Request) SetPolicyDescription(v string) *UpdatePolicyV2Request {
	s.PolicyDescription = &v
	return s
}

func (s *UpdatePolicyV2Request) SetPolicyId(v string) *UpdatePolicyV2Request {
	s.PolicyId = &v
	return s
}

func (s *UpdatePolicyV2Request) SetPolicyName(v string) *UpdatePolicyV2Request {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyV2Request) SetRules(v []*UpdatePolicyV2RequestRules) *UpdatePolicyV2Request {
	s.Rules = v
	return s
}

func (s *UpdatePolicyV2Request) Validate() error {
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

type UpdatePolicyV2RequestRules struct {
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The number of days after which the backup is converted to archive storage. Unit: days.
	//
	// example:
	//
	// 90
	ArchiveDays *int64 `json:"ArchiveDays,omitempty" xml:"ArchiveDays,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **BACKUP**. The backup type. Set the value to **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The number of days after which the backup is converted to cold archive storage. Unit: days.
	//
	// example:
	//
	// 365
	ColdArchiveDays *int64 `json:"ColdArchiveDays,omitempty" xml:"ColdArchiveDays,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. The data source filter rules.
	DataSourceFilters []*UpdatePolicyV2RequestRulesDataSourceFilters `json:"DataSourceFilters,omitempty" xml:"DataSourceFilters,omitempty" type:"Repeated"`
	// This parameter is required only when **PolicyType*	- is set to **UDM_ECS_ONLY*	- and **RuleType*	- is set to **SECURITY**. Specifies whether to enable backup locking.
	//
	// example:
	//
	// true
	Immutable *bool `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	// Specifies whether to retain at least one backup version. Valid values:
	//
	// - 0: do not retain.
	//
	// - 1: retain.
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **REPLICATION**. The ID of the destination region for replication.
	//
	// example:
	//
	// cn-shanghai
	ReplicationRegionId *string `json:"ReplicationRegionId,omitempty" xml:"ReplicationRegionId,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TRANSITION*	- or **REPLICATION**.
	//
	// - If **RuleType*	- is set to **TRANSITION**: the retention period of the backup. Minimum value: 1. Unit: days.
	//
	// - If **RuleType*	- is set to **REPLICATION**: the retention period of the cross-region backup. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The special retention rules.
	RetentionRules []*UpdatePolicyV2RequestRulesRetentionRules `json:"RetentionRules,omitempty" xml:"RetentionRules,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// rule-000************rof
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
	// This parameter is required only when **RuleType*	- is set to **BACKUP**. The backup schedule settings. Supported formats:
	//
	// - `I|{startTime}|{interval}`: specifies that a backup job is run at the {interval} from the {startTime}. Example: `I|1631685600|P1D` specifies that a backup job is run once a day starting from 2021-09-15 14:00:00.
	//
	//   	- startTime: the start time of the backup. This value is a UNIX timestamp. Unit: seconds.
	//
	//   	- interval: the ISO 8601 time interval. Example: `PT1H` specifies an interval of one hour. `P1D` specifies an interval of one day.
	//
	// - `C|{startTime}|{crontab}`: specifies that a backup job is run based on the {crontab} expression from the {startTime}. Example: `C|1631685600|0 0 2 ? 	- 3,5,7` specifies that a backup job is run at 02:00:00 every Tuesday, Thursday, and Saturday starting from 2021-09-15 14:00:00.
	//
	//   	- startTime: the start time of the backup. This value is a UNIX timestamp. Unit: seconds.
	//
	//   	- crontab: the crontab expression. Example: `0 0 2 ? 	- 3,5,7` specifies every Tuesday, Thursday, and Saturday at 02:00:00.
	//
	// Backup jobs for elapsed time periods are not compensated. If the previous backup job is not completed, the next backup job is not triggered.
	//
	// example:
	//
	// I|1648647166|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. The resource tag filter rules.
	TagFilters []*UpdatePolicyV2RequestRulesTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
	// This parameter is required only when RuleType is set to BACKUP. The backup vault ID.
	//
	// example:
	//
	// v-0001************aseg
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdatePolicyV2RequestRules) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2RequestRules) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2RequestRules) GetArchiveDays() *int64 {
	return s.ArchiveDays
}

func (s *UpdatePolicyV2RequestRules) GetBackupType() *string {
	return s.BackupType
}

func (s *UpdatePolicyV2RequestRules) GetColdArchiveDays() *int64 {
	return s.ColdArchiveDays
}

func (s *UpdatePolicyV2RequestRules) GetDataSourceFilters() []*UpdatePolicyV2RequestRulesDataSourceFilters {
	return s.DataSourceFilters
}

func (s *UpdatePolicyV2RequestRules) GetImmutable() *bool {
	return s.Immutable
}

func (s *UpdatePolicyV2RequestRules) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *UpdatePolicyV2RequestRules) GetReplicationRegionId() *string {
	return s.ReplicationRegionId
}

func (s *UpdatePolicyV2RequestRules) GetRetention() *int64 {
	return s.Retention
}

func (s *UpdatePolicyV2RequestRules) GetRetentionRules() []*UpdatePolicyV2RequestRulesRetentionRules {
	return s.RetentionRules
}

func (s *UpdatePolicyV2RequestRules) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdatePolicyV2RequestRules) GetRuleType() *string {
	return s.RuleType
}

func (s *UpdatePolicyV2RequestRules) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdatePolicyV2RequestRules) GetTagFilters() []*UpdatePolicyV2RequestRulesTagFilters {
	return s.TagFilters
}

func (s *UpdatePolicyV2RequestRules) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdatePolicyV2RequestRules) SetArchiveDays(v int64) *UpdatePolicyV2RequestRules {
	s.ArchiveDays = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetBackupType(v string) *UpdatePolicyV2RequestRules {
	s.BackupType = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetColdArchiveDays(v int64) *UpdatePolicyV2RequestRules {
	s.ColdArchiveDays = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetDataSourceFilters(v []*UpdatePolicyV2RequestRulesDataSourceFilters) *UpdatePolicyV2RequestRules {
	s.DataSourceFilters = v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetImmutable(v bool) *UpdatePolicyV2RequestRules {
	s.Immutable = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetKeepLatestSnapshots(v int64) *UpdatePolicyV2RequestRules {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetReplicationRegionId(v string) *UpdatePolicyV2RequestRules {
	s.ReplicationRegionId = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetRetention(v int64) *UpdatePolicyV2RequestRules {
	s.Retention = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetRetentionRules(v []*UpdatePolicyV2RequestRulesRetentionRules) *UpdatePolicyV2RequestRules {
	s.RetentionRules = v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetRuleId(v string) *UpdatePolicyV2RequestRules {
	s.RuleId = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetRuleType(v string) *UpdatePolicyV2RequestRules {
	s.RuleType = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetSchedule(v string) *UpdatePolicyV2RequestRules {
	s.Schedule = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetTagFilters(v []*UpdatePolicyV2RequestRulesTagFilters) *UpdatePolicyV2RequestRules {
	s.TagFilters = v
	return s
}

func (s *UpdatePolicyV2RequestRules) SetVaultId(v string) *UpdatePolicyV2RequestRules {
	s.VaultId = &v
	return s
}

func (s *UpdatePolicyV2RequestRules) Validate() error {
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

type UpdatePolicyV2RequestRulesDataSourceFilters struct {
	AccountScope *string                                                `json:"AccountScope,omitempty" xml:"AccountScope,omitempty"`
	Accounts     []*UpdatePolicyV2RequestRulesDataSourceFiltersAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS instance backup. This data source type is supported only when **RuleType*	- is set to **UDM_ECS_ONLY**.
	//
	// - **OSS**: OSS backup. This data source type is supported only when **RuleType*	- is set to **STANDARD**.
	//
	// - **NAS**: Alibaba Cloud NAS backup. This data source type is supported only when **RuleType*	- is set to **STANDARD**.
	//
	// - **ECS_FILE**: ECS File Backup Essential Edition. This data source type is supported only when **RuleType*	- is set to **STANDARD**.
	//
	// - **OTS**: Tablestore backup. This data source type is supported only when **RuleType*	- is set to **STANDARD**.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdatePolicyV2RequestRulesDataSourceFilters) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2RequestRulesDataSourceFilters) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) GetAccountScope() *string {
	return s.AccountScope
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) GetAccounts() []*UpdatePolicyV2RequestRulesDataSourceFiltersAccounts {
	return s.Accounts
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) SetAccountScope(v string) *UpdatePolicyV2RequestRulesDataSourceFilters {
	s.AccountScope = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) SetAccounts(v []*UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) *UpdatePolicyV2RequestRulesDataSourceFilters {
	s.Accounts = v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) SetDataSourceIds(v []*string) *UpdatePolicyV2RequestRulesDataSourceFilters {
	s.DataSourceIds = v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) SetSourceType(v string) *UpdatePolicyV2RequestRulesDataSourceFilters {
	s.SourceType = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) Validate() error {
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

type UpdatePolicyV2RequestRulesDataSourceFiltersAccounts struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) SetCrossAccountRoleName(v string) *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts {
	s.CrossAccountRoleName = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) SetCrossAccountType(v string) *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts {
	s.CrossAccountType = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) SetCrossAccountUserId(v int64) *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts {
	s.CrossAccountUserId = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesDataSourceFiltersAccounts) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyV2RequestRulesRetentionRules struct {
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
	// 365
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The backup to which the rule applies. Currently, only the first backup is supported. Set the value to 1.
	//
	// example:
	//
	// 1
	WhichSnapshot *int64 `json:"WhichSnapshot,omitempty" xml:"WhichSnapshot,omitempty"`
}

func (s UpdatePolicyV2RequestRulesRetentionRules) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2RequestRulesRetentionRules) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) GetAdvancedRetentionType() *string {
	return s.AdvancedRetentionType
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) GetRetention() *int64 {
	return s.Retention
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) GetWhichSnapshot() *int64 {
	return s.WhichSnapshot
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) SetAdvancedRetentionType(v string) *UpdatePolicyV2RequestRulesRetentionRules {
	s.AdvancedRetentionType = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) SetRetention(v int64) *UpdatePolicyV2RequestRulesRetentionRules {
	s.Retention = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) SetWhichSnapshot(v int64) *UpdatePolicyV2RequestRulesRetentionRules {
	s.WhichSnapshot = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesRetentionRules) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyV2RequestRulesTagFilters struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag matching rule. Valid values:
	//
	// - **EQUAL**: matches both the tag key and tag value.
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

func (s UpdatePolicyV2RequestRulesTagFilters) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyV2RequestRulesTagFilters) GoString() string {
	return s.String()
}

func (s *UpdatePolicyV2RequestRulesTagFilters) GetKey() *string {
	return s.Key
}

func (s *UpdatePolicyV2RequestRulesTagFilters) GetOperator() *string {
	return s.Operator
}

func (s *UpdatePolicyV2RequestRulesTagFilters) GetValue() *string {
	return s.Value
}

func (s *UpdatePolicyV2RequestRulesTagFilters) SetKey(v string) *UpdatePolicyV2RequestRulesTagFilters {
	s.Key = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesTagFilters) SetOperator(v string) *UpdatePolicyV2RequestRulesTagFilters {
	s.Operator = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesTagFilters) SetValue(v string) *UpdatePolicyV2RequestRulesTagFilters {
	s.Value = &v
	return s
}

func (s *UpdatePolicyV2RequestRulesTagFilters) Validate() error {
	return dara.Validate(s)
}
