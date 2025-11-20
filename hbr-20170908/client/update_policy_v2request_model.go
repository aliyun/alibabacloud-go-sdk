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
	// po-000************viy
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the backup policy.
	//
	// example:
	//
	// Daily Local Backup + Remote Backup
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The rules in the backup policy.
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
	// This parameter is required only if the **RuleType*	- parameter is set to **TRANSITION**. This parameter specifies the time when data is dumped from a backup vault to an archive vault. Unit: days.
	//
	// example:
	//
	// 90
	ArchiveDays *int64 `json:"ArchiveDays,omitempty" xml:"ArchiveDays,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **BACKUP**. This parameter specifies the backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TRANSITION**. This parameter specifies the time when data is dumped from a backup vault to a cold archive vault. Unit: days.
	//
	// example:
	//
	// 365
	ColdArchiveDays *int64 `json:"ColdArchiveDays,omitempty" xml:"ColdArchiveDays,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TAG**. This parameter specifies the data source filter rule.
	DataSourceFilters []*UpdatePolicyV2RequestRulesDataSourceFilters `json:"DataSourceFilters,omitempty" xml:"DataSourceFilters,omitempty" type:"Repeated"`
	// This parameter is required only if the **PolicyType*	- parameter is set to **UDM_ECS_ONLY**. This parameter specifies whether to enable the immutable backup feature.
	//
	// example:
	//
	// true
	Immutable *bool `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	// Specifies whether to enable the feature of keeping at least one backup version. Valid values:
	//
	// 	- 0: The feature is disabled.
	//
	// 	- 1: The feature is enabled.
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **REPLICATION**. This parameter specifies the ID of the destination region.
	//
	// example:
	//
	// cn-shanghai
	ReplicationRegionId *string `json:"ReplicationRegionId,omitempty" xml:"ReplicationRegionId,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TRANSITION*	- or **REPLICATION**.
	//
	// 	- If the **RuleType*	- parameter is set to **TRANSITION**, this parameter specifies the retention period of the backup data. Minimum value: 1. Unit: days.
	//
	// 	- If the **RuleType*	- parameter is set to **REPLICATION**, this parameter specifies the retention period of remote backups. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TRANSITION**. This parameter specifies the special retention rules.
	RetentionRules []*UpdatePolicyV2RequestRulesRetentionRules `json:"RetentionRules,omitempty" xml:"RetentionRules,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// rule-000************rof
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
	// This parameter is required only if the **RuleType*	- parameter is set to **BACKUP**. This parameter specifies the backup schedule settings. Formats:
	//
	// 	- `I|{startTime}|{interval}`: The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. For example, `I|1631685600|P1D` indicates that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	//     	- startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds.
	//
	//     	- interval: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, `PT1H` specifies an interval of 1 hour. `P1D` specifies an interval of one day.
	//
	// 	- `C|{startTime}|{crontab}`: The system runs backup jobs at a point in time that is specified in the {startTime} parameter based on the {crontab} expression. For example, C|1631685600|0 0 2 ?\\	- 3,5,7 indicates that the system runs backup jobs at 02:00:00 every Tuesday, Thursday, and Saturday from14:00:00 on September 15, 2021.``
	//
	//     	- startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds.
	//
	//     	- crontab: the crontab expression. For example, 0 0 2 ?\\	- 3,5,7 indicates 02:00:00 every Tuesday, Thursday, and Saturday.``
	//
	// The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed.
	//
	// example:
	//
	// I|1648647166|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TAG**. This parameter specifies the resource tag filter rule.
	TagFilters []*UpdatePolicyV2RequestRulesTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
	// This parameter is required only if the RuleType parameter is set to BACKUP. The ID of the backup vault.
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
	// This parameter is deprecated.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: Elastic Compute Service (ECS) instance This type of data source is supported only if the **RuleType*	- parameter is set to **UDM_ECS_ONLY**.
	//
	// 	- **OSS**: Object Storage Service (OSS) bucket This type of data source is supported only if the **RuleType*	- parameter is set to **STANDARD**.
	//
	// 	- **NAS**: File Storage NAS (NAS) file system This type of data source is supported only if the **RuleType*	- parameter is set to **STANDARD**.
	//
	// 	- **ECS_FILE**: ECS file This type of data source is supported only if the **RuleType*	- parameter is set to **STANDARD**.
	//
	// 	- **OTS**: Tablestore instance This type of data source is supported only if the **RuleType*	- parameter is set to **STANDARD**.
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

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *UpdatePolicyV2RequestRulesDataSourceFilters) GetSourceType() *string {
	return s.SourceType
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
	return dara.Validate(s)
}

type UpdatePolicyV2RequestRulesRetentionRules struct {
	// The type of the special retention rule. Valid values:
	//
	// 	- **WEEKLY**: retains weekly backups
	//
	// 	- **MONTHLY**: retains monthly backups
	//
	// 	- **YEARLY**: retains yearly backups
	//
	// example:
	//
	// YEARLY
	AdvancedRetentionType *string `json:"AdvancedRetentionType,omitempty" xml:"AdvancedRetentionType,omitempty"`
	// The special retention period of backups. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 365
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// Specifies which backup is retained based on the special retention rule. Only the first backup can be retained.
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
	// The tag-based matching rule. Valid values:
	//
	// 	- **EQUAL**: Both the tag key and tag value are matched.
	//
	// 	- **NOT**: The tag key is matched and the tag value is not matched.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The tag value. If you leave this parameter empty, the value is any value.
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
