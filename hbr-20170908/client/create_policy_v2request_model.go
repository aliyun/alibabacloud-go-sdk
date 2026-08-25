// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyDescription(v string) *CreatePolicyV2Request
	GetPolicyDescription() *string
	SetPolicyName(v string) *CreatePolicyV2Request
	GetPolicyName() *string
	SetPolicyType(v string) *CreatePolicyV2Request
	GetPolicyType() *string
	SetRules(v []*CreatePolicyV2RequestRules) *CreatePolicyV2Request
	GetRules() []*CreatePolicyV2RequestRules
}

type CreatePolicyV2Request struct {
	// The policy description.
	//
	// example:
	//
	// Backup once every day at 10:00 AM, with cross-region backup to Shanghai.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Daily local backup + geo-redundancy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// - **STANDARD**: general backup policy. Supports backing up data sources other than ECS instances.
	//
	// - **UDM_ECS_ONLY**: ECS instance backup policy. Supports backing up only ECS instances.
	//
	// If you do not specify the policy type, Cloud Backup automatically sets the policy type based on whether a backup vault is specified in the policy rules:
	//
	// - A backup vault is specified in the policy rules: **STANDARD**
	//
	// - No backup vault is specified in the policy rules: **UDM_ECS_ONLY**
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The list of policy rules.
	Rules []*CreatePolicyV2RequestRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s CreatePolicyV2Request) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2Request) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2Request) GetPolicyDescription() *string {
	return s.PolicyDescription
}

func (s *CreatePolicyV2Request) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyV2Request) GetPolicyType() *string {
	return s.PolicyType
}

func (s *CreatePolicyV2Request) GetRules() []*CreatePolicyV2RequestRules {
	return s.Rules
}

func (s *CreatePolicyV2Request) SetPolicyDescription(v string) *CreatePolicyV2Request {
	s.PolicyDescription = &v
	return s
}

func (s *CreatePolicyV2Request) SetPolicyName(v string) *CreatePolicyV2Request {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyV2Request) SetPolicyType(v string) *CreatePolicyV2Request {
	s.PolicyType = &v
	return s
}

func (s *CreatePolicyV2Request) SetRules(v []*CreatePolicyV2RequestRules) *CreatePolicyV2Request {
	s.Rules = v
	return s
}

func (s *CreatePolicyV2Request) Validate() error {
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

type CreatePolicyV2RequestRules struct {
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The number of days after which a backup is automatically moved to the archive tier. Backups must be retained in the standard tier for at least 30 days and in the archive tier for at least 60 days. Unit: days.
	//
	// example:
	//
	// 90
	ArchiveDays *int64 `json:"ArchiveDays,omitempty" xml:"ArchiveDays,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **BACKUP**. The backup type. Set the value to **COMPLETE**, which specifies full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. The data source filter rules.
	DataSourceFilters []*CreatePolicyV2RequestRulesDataSourceFilters `json:"DataSourceFilters,omitempty" xml:"DataSourceFilters,omitempty" type:"Repeated"`
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
	// This parameter is required only when **RuleType*	- is set to **BACKUP**, **TRANSITION**, or **REPLICATION**.
	//
	// - If **RuleType*	- is set to **BACKUP**: the retention period of backups. The priority of this parameter is lower than the Retention parameter of the rule whose **RuleType*	- is **TRANSITION**. Minimum value: 1. Maximum value: 364635. Unit: days.
	//
	// - If **RuleType*	- is set to **TRANSITION**: the retention period of backups. Minimum value: 1. Maximum value: 364635. Unit: days.
	//
	// - If **RuleType*	- is set to **REPLICATION**: the retention period of cross-region backups. Minimum value: 1. Maximum value: 364635. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TRANSITION**. The special retention rules.
	RetentionRules []*CreatePolicyV2RequestRulesRetentionRules `json:"RetentionRules,omitempty" xml:"RetentionRules,omitempty" type:"Repeated"`
	// The rule type. Each policy must have at least one **BACKUP*	- rule and exactly one **TRANSITION*	- rule. Valid values:
	//
	// - **BACKUP**: backup rule.
	//
	// - **TRANSITION**: lifecycle rule.
	//
	// - **REPLICATION**: replication rule.
	//
	// - **TAG**: tag-based resource association rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// BACKUP
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **BACKUP**. The backup schedule settings. Supported formats:
	//
	// - `I|{startTime}|{interval}`: specifies that a backup job is executed at the specified interval starting from {startTime}. For example, `I|1631685600|P1D` specifies that a backup job is executed once a day starting from 2021-09-15 14:00:00.
	//
	//   	- startTime: the start time of the backup. This value is a UNIX timestamp. Unit: seconds.
	//
	//   	- interval: the ISO 8601 time interval. For example, `PT1H` specifies an interval of one hour. `P1D` specifies an interval of one day.
	//
	// - `C|{startTime}|{crontab}`: specifies that a backup job is executed based on the {crontab} expression starting from {startTime}. For example, `C|1631685600|0 0 2 ? 	- 3,5,7` specifies that a backup job is executed at 02:00:00 every Tuesday, Thursday, and Saturday starting from 2021-09-15 14:00:00.
	//
	//   	- startTime: the start time of the backup. This value is a UNIX timestamp. Unit: seconds.
	//
	//   	- crontab: the crontab expression. For example, `0 0 2 ? 	- 3,5,7` specifies every Tuesday, Thursday, and Saturday at 02:00:00.
	//
	// Backup jobs that are missed are not compensated. If the previous backup job is not complete, the next backup job is not triggered.
	//
	// example:
	//
	// I|1648647166|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// This parameter is required only when **RuleType*	- is set to **TAG**. The resource tag filter rules.
	TagFilters []*CreatePolicyV2RequestRulesTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
	// This parameter is required only when RuleType is set to BACKUP. The backup vault ID.
	//
	// example:
	//
	// v-0001************aseg
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreatePolicyV2RequestRules) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2RequestRules) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2RequestRules) GetArchiveDays() *int64 {
	return s.ArchiveDays
}

func (s *CreatePolicyV2RequestRules) GetBackupType() *string {
	return s.BackupType
}

func (s *CreatePolicyV2RequestRules) GetDataSourceFilters() []*CreatePolicyV2RequestRulesDataSourceFilters {
	return s.DataSourceFilters
}

func (s *CreatePolicyV2RequestRules) GetImmutable() *bool {
	return s.Immutable
}

func (s *CreatePolicyV2RequestRules) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *CreatePolicyV2RequestRules) GetReplicationRegionId() *string {
	return s.ReplicationRegionId
}

func (s *CreatePolicyV2RequestRules) GetRetention() *int64 {
	return s.Retention
}

func (s *CreatePolicyV2RequestRules) GetRetentionRules() []*CreatePolicyV2RequestRulesRetentionRules {
	return s.RetentionRules
}

func (s *CreatePolicyV2RequestRules) GetRuleType() *string {
	return s.RuleType
}

func (s *CreatePolicyV2RequestRules) GetSchedule() *string {
	return s.Schedule
}

func (s *CreatePolicyV2RequestRules) GetTagFilters() []*CreatePolicyV2RequestRulesTagFilters {
	return s.TagFilters
}

func (s *CreatePolicyV2RequestRules) GetVaultId() *string {
	return s.VaultId
}

func (s *CreatePolicyV2RequestRules) SetArchiveDays(v int64) *CreatePolicyV2RequestRules {
	s.ArchiveDays = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetBackupType(v string) *CreatePolicyV2RequestRules {
	s.BackupType = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetDataSourceFilters(v []*CreatePolicyV2RequestRulesDataSourceFilters) *CreatePolicyV2RequestRules {
	s.DataSourceFilters = v
	return s
}

func (s *CreatePolicyV2RequestRules) SetImmutable(v bool) *CreatePolicyV2RequestRules {
	s.Immutable = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetKeepLatestSnapshots(v int64) *CreatePolicyV2RequestRules {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetReplicationRegionId(v string) *CreatePolicyV2RequestRules {
	s.ReplicationRegionId = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetRetention(v int64) *CreatePolicyV2RequestRules {
	s.Retention = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetRetentionRules(v []*CreatePolicyV2RequestRulesRetentionRules) *CreatePolicyV2RequestRules {
	s.RetentionRules = v
	return s
}

func (s *CreatePolicyV2RequestRules) SetRuleType(v string) *CreatePolicyV2RequestRules {
	s.RuleType = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetSchedule(v string) *CreatePolicyV2RequestRules {
	s.Schedule = &v
	return s
}

func (s *CreatePolicyV2RequestRules) SetTagFilters(v []*CreatePolicyV2RequestRulesTagFilters) *CreatePolicyV2RequestRules {
	s.TagFilters = v
	return s
}

func (s *CreatePolicyV2RequestRules) SetVaultId(v string) *CreatePolicyV2RequestRules {
	s.VaultId = &v
	return s
}

func (s *CreatePolicyV2RequestRules) Validate() error {
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

type CreatePolicyV2RequestRulesDataSourceFilters struct {
	AccountScope *string                                                `json:"AccountScope,omitempty" xml:"AccountScope,omitempty"`
	Accounts     []*CreatePolicyV2RequestRulesDataSourceFiltersAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// Deprecated
	//
	// Deprecated.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS instance backup. This data source type is supported only when **PolicyType*	- is set to **UDM_ECS_ONLY**.
	//
	// - **OSS**: OSS backup. This data source type is supported only when **PolicyType*	- is set to **STANDARD**.
	//
	// - **NAS**: Alibaba Cloud NAS backup. This data source type is supported only when **PolicyType*	- is set to **STANDARD**.
	//
	// - **ECS_FILE**: ECS File Backup Essential Edition. This data source type is supported only when **PolicyType*	- is set to **STANDARD**.
	//
	// - **OTS**: Tablestore backup. This data source type is supported only when **PolicyType*	- is set to **STANDARD**.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreatePolicyV2RequestRulesDataSourceFilters) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2RequestRulesDataSourceFilters) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) GetAccountScope() *string {
	return s.AccountScope
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) GetAccounts() []*CreatePolicyV2RequestRulesDataSourceFiltersAccounts {
	return s.Accounts
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) GetSourceType() *string {
	return s.SourceType
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) SetAccountScope(v string) *CreatePolicyV2RequestRulesDataSourceFilters {
	s.AccountScope = &v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) SetAccounts(v []*CreatePolicyV2RequestRulesDataSourceFiltersAccounts) *CreatePolicyV2RequestRulesDataSourceFilters {
	s.Accounts = v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) SetDataSourceIds(v []*string) *CreatePolicyV2RequestRulesDataSourceFilters {
	s.DataSourceIds = v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) SetSourceType(v string) *CreatePolicyV2RequestRulesDataSourceFilters {
	s.SourceType = &v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) Validate() error {
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

type CreatePolicyV2RequestRulesDataSourceFiltersAccounts struct {
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	CrossAccountType     *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	CrossAccountUserId   *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s CreatePolicyV2RequestRulesDataSourceFiltersAccounts) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2RequestRulesDataSourceFiltersAccounts) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) SetCrossAccountRoleName(v string) *CreatePolicyV2RequestRulesDataSourceFiltersAccounts {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) SetCrossAccountType(v string) *CreatePolicyV2RequestRulesDataSourceFiltersAccounts {
	s.CrossAccountType = &v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) SetCrossAccountUserId(v int64) *CreatePolicyV2RequestRulesDataSourceFiltersAccounts {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreatePolicyV2RequestRulesDataSourceFiltersAccounts) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyV2RequestRulesRetentionRules struct {
	// The type of the special retention rule. Valid values:
	//
	// - **DAILY**: daily backup.
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
	// The backup to which the rule applies. Currently, only the first backup is supported. Set the value to 1.
	//
	// example:
	//
	// 1
	WhichSnapshot *int64 `json:"WhichSnapshot,omitempty" xml:"WhichSnapshot,omitempty"`
}

func (s CreatePolicyV2RequestRulesRetentionRules) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2RequestRulesRetentionRules) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2RequestRulesRetentionRules) GetAdvancedRetentionType() *string {
	return s.AdvancedRetentionType
}

func (s *CreatePolicyV2RequestRulesRetentionRules) GetRetention() *int64 {
	return s.Retention
}

func (s *CreatePolicyV2RequestRulesRetentionRules) GetWhichSnapshot() *int64 {
	return s.WhichSnapshot
}

func (s *CreatePolicyV2RequestRulesRetentionRules) SetAdvancedRetentionType(v string) *CreatePolicyV2RequestRulesRetentionRules {
	s.AdvancedRetentionType = &v
	return s
}

func (s *CreatePolicyV2RequestRulesRetentionRules) SetRetention(v int64) *CreatePolicyV2RequestRulesRetentionRules {
	s.Retention = &v
	return s
}

func (s *CreatePolicyV2RequestRulesRetentionRules) SetWhichSnapshot(v int64) *CreatePolicyV2RequestRulesRetentionRules {
	s.WhichSnapshot = &v
	return s
}

func (s *CreatePolicyV2RequestRulesRetentionRules) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyV2RequestRulesTagFilters struct {
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

func (s CreatePolicyV2RequestRulesTagFilters) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyV2RequestRulesTagFilters) GoString() string {
	return s.String()
}

func (s *CreatePolicyV2RequestRulesTagFilters) GetKey() *string {
	return s.Key
}

func (s *CreatePolicyV2RequestRulesTagFilters) GetOperator() *string {
	return s.Operator
}

func (s *CreatePolicyV2RequestRulesTagFilters) GetValue() *string {
	return s.Value
}

func (s *CreatePolicyV2RequestRulesTagFilters) SetKey(v string) *CreatePolicyV2RequestRulesTagFilters {
	s.Key = &v
	return s
}

func (s *CreatePolicyV2RequestRulesTagFilters) SetOperator(v string) *CreatePolicyV2RequestRulesTagFilters {
	s.Operator = &v
	return s
}

func (s *CreatePolicyV2RequestRulesTagFilters) SetValue(v string) *CreatePolicyV2RequestRulesTagFilters {
	s.Value = &v
	return s
}

func (s *CreatePolicyV2RequestRulesTagFilters) Validate() error {
	return dara.Validate(s)
}
