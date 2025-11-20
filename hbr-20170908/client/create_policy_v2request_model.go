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
	// The description of the backup policy.
	//
	// example:
	//
	// Data is backed up at 10:00:00 every day and replicated to the China (Shanghai) region for geo-redundancy.
	PolicyDescription *string `json:"PolicyDescription,omitempty" xml:"PolicyDescription,omitempty"`
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
	// 	- **UDM_ECS_ONLY**: This type of policy applies only to ECS instance backup.
	//
	// If the policy type is not specified, Cloud Backup automatically sets the policy type based on whether the backup vault is specified in the rules of the policy:
	//
	// 	- If the backup vault is specified, Cloud Backup sets the policy type to **STANDARD**.
	//
	// 	- If the backup vault is not specified, Cloud Backup sets the policy type to **UDM_ECS_ONLY**.
	//
	// example:
	//
	// STANDARD
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The rules in the backup policy.
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
	// This parameter is required only if the **RuleType*	- parameter is set to **BACKUP**. This parameter specifies the backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TAG**. This parameter specifies the data source filter rule.
	DataSourceFilters []*CreatePolicyV2RequestRulesDataSourceFilters `json:"DataSourceFilters,omitempty" xml:"DataSourceFilters,omitempty" type:"Repeated"`
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
	// This parameter is required only if the **RuleType*	- parameter is set to **BACKUP**, **TRANSITION**, or **REPLICATION**.
	//
	// 	- If the **RuleType*	- parameter is set to **BACKUP**, this parameter specifies the retention period of the backup data. The priority is lower than the retention period when the **RuleType*	- parameter is set to **TRANSITION**. Minimum value: 1. Maximum value: 364635. Unit: days.
	//
	// 	- If the **RuleType*	- parameter is set to **TRANSITION**, this parameter specifies the retention period of the backup data. Minimum value: 1. Maximum value: 364635. Unit: days.
	//
	// 	- If the **RuleType*	- parameter is set to **REPLICATION**, this parameter specifies the retention period of remote backups. Minimum value: 1. Maximum value: 364635. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// This parameter is required only if the **RuleType*	- parameter is set to **TRANSITION**. This parameter specifies the special retention rules.
	RetentionRules []*CreatePolicyV2RequestRulesRetentionRules `json:"RetentionRules,omitempty" xml:"RetentionRules,omitempty" type:"Repeated"`
	// The type of the rule. Each backup policy must have at least one rule of the **BACKUP*	- type and only one rule of the **TRANSITION*	- type. Valid values:
	//
	// 	- **BACKUP**: backup rule
	//
	// 	- **TRANSITION**: lifecycle rule
	//
	// 	- **REPLICATION**: replication rule
	//
	// 	- **TAG**: tag-based resource association rule
	//
	// This parameter is required.
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
	TagFilters []*CreatePolicyV2RequestRulesTagFilters `json:"TagFilters,omitempty" xml:"TagFilters,omitempty" type:"Repeated"`
	// This parameter is required only if the RuleType parameter is set to BACKUP. The ID of the backup vault.
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
	// This parameter is deprecated.
	DataSourceIds []*string `json:"DataSourceIds,omitempty" xml:"DataSourceIds,omitempty" type:"Repeated"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: Elastic Compute Service (ECS) instance This type of data source is supported only if the **PolicyType*	- parameter is set to **UDM_ECS_ONLY**.
	//
	// 	- **OSS**: Object Storage Service (OSS) bucket This type of data source is supported only if the **PolicyType*	- parameter is set to **STANDARD**.
	//
	// 	- **NAS**: File Storage NAS (NAS) file system This type of data source is supported only if the **PolicyType*	- parameter is set to **STANDARD**.
	//
	// 	- **ECS_FILE**: ECS file This type of data source is supported only if the **PolicyType*	- parameter is set to **STANDARD**.
	//
	// 	- **OTS**: Tablestore instance This type of data source is supported only if the **PolicyType*	- parameter is set to **STANDARD**.
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

func (s *CreatePolicyV2RequestRulesDataSourceFilters) GetDataSourceIds() []*string {
	return s.DataSourceIds
}

func (s *CreatePolicyV2RequestRulesDataSourceFilters) GetSourceType() *string {
	return s.SourceType
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
	return dara.Validate(s)
}

type CreatePolicyV2RequestRulesRetentionRules struct {
	// The type of the special retention rule. Valid values:
	//
	// 	- **DAILY**: retains daily backups
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
	// 730
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// Specifies which backup is retained based on the special retention rule. Only the first backup can be retained.
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
