// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPlanShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeListPath(v string) *UpdateBackupPlanShrinkRequest
	GetChangeListPath() *string
	SetDetailShrink(v string) *UpdateBackupPlanShrinkRequest
	GetDetailShrink() *string
	SetExclude(v string) *UpdateBackupPlanShrinkRequest
	GetExclude() *string
	SetInclude(v string) *UpdateBackupPlanShrinkRequest
	GetInclude() *string
	SetKeepLatestSnapshots(v int64) *UpdateBackupPlanShrinkRequest
	GetKeepLatestSnapshots() *int64
	SetOptions(v string) *UpdateBackupPlanShrinkRequest
	GetOptions() *string
	SetOtsDetailShrink(v string) *UpdateBackupPlanShrinkRequest
	GetOtsDetailShrink() *string
	SetPath(v []*string) *UpdateBackupPlanShrinkRequest
	GetPath() []*string
	SetPlanId(v string) *UpdateBackupPlanShrinkRequest
	GetPlanId() *string
	SetPlanName(v string) *UpdateBackupPlanShrinkRequest
	GetPlanName() *string
	SetPrefix(v string) *UpdateBackupPlanShrinkRequest
	GetPrefix() *string
	SetRetention(v int64) *UpdateBackupPlanShrinkRequest
	GetRetention() *int64
	SetRule(v []*UpdateBackupPlanShrinkRequestRule) *UpdateBackupPlanShrinkRequest
	GetRule() []*UpdateBackupPlanShrinkRequestRule
	SetSchedule(v string) *UpdateBackupPlanShrinkRequest
	GetSchedule() *string
	SetSourceType(v string) *UpdateBackupPlanShrinkRequest
	GetSourceType() *string
	SetSpeedLimit(v string) *UpdateBackupPlanShrinkRequest
	GetSpeedLimit() *string
	SetUpdatePaths(v bool) *UpdateBackupPlanShrinkRequest
	GetUpdatePaths() *bool
	SetVaultId(v string) *UpdateBackupPlanShrinkRequest
	GetVaultId() *string
}

type UpdateBackupPlanShrinkRequest struct {
	// The configurations of the incremental file synchronization. This parameter is required for data synchronization only.
	//
	// example:
	//
	// {"dataSourceId": "ds-123456789", "path": "/changelist"}
	ChangeListPath *string `json:"ChangeListPath,omitempty" xml:"ChangeListPath,omitempty"`
	// The details about ECS instance backup. The value is a JSON string.
	//
	// 	- snapshotGroup: specifies whether to use a snapshot-consistent group. This parameter is valid only if all disks of the ECS instance are enhanced SSDs (ESSDs).
	//
	// 	- appConsistent: specifies whether to enable application consistency. If you set this parameter to true, you must also specify the preScriptPath and postScriptPath parameters.
	//
	// 	- preScriptPath: the path to the pre-freeze scripts.
	//
	// 	- postScriptPath: the path to the post-thaw scripts.
	//
	// example:
	//
	// {\\"EnableFsFreeze\\":true,\\"appConsistent\\":false,\\"postScriptPath\\":\\"\\",\\"preScriptPath\\":\\"\\",\\"snapshotGroup\\":true,\\"timeoutInSeconds\\":60}
	DetailShrink *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// This parameter is required only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter specifies the paths to the files that are excluded from the backup job. The value must be 1 to 255 characters in length.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is required only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter specifies the paths to the files that you want to back up. The value must be 1 to 255 characters in length.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
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
	// This parameter is required only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter specifies whether to use Windows Volume Shadow Copy Service (VSS) to define a source path.
	//
	// 	- This parameter is available only for Windows ECS instances.
	//
	// 	- If data changes occur in the backup source, the source data must be the same as the data to be backed up before you can set this parameter to `["UseVSS":true]`.
	//
	// 	- If you use VSS, you cannot back up data from multiple directories.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details about the Tablestore instance.
	OtsDetailShrink *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// The source paths.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
	// The ID of the backup plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-20211***735
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the backup plan.
	//
	// example:
	//
	// planname
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// This parameter is required only if the **SourceType*	- parameter is set to **OSS**. This parameter specifies the prefix of objects that you want to back up. After a prefix is specified, only objects whose names start with the prefix are backed up.
	//
	// example:
	//
	// oss-prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The retention period of the backup data. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The rule of the backup plan.
	Rule []*UpdateBackupPlanShrinkRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	// The backup policy. Format: `I|{startTime}|{interval}`. The system runs the first backup job at a point in time that is specified in the `{startTime}` parameter and the subsequent backup jobs at an interval that is specified in the `{interval}` parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, `I|1631685600|P1D` specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// 	- **startTime**: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds.
	//
	// 	- **interval**: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **ECS_FILE**: Elastic Compute Service (ECS) files
	//
	// 	- **OSS**: Object Storage Service (OSS) buckets
	//
	// 	- **NAS**: Apsara File Storage NAS file systems
	//
	// 	- **OTS**: Tablestore instances
	//
	// 	- **UDM_ECS**: ECS instances
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only if the **SourceType*	- parameter is set to **ECS_FILE**. This parameter specifies the throttling rules. To ensure business continuity, you can limit the bandwidth that is used for file backup during peak hours. Format: `{start}|{end}|{bandwidth}`. Separate multiple throttling rules with vertical bars (|). A specified time range cannot overlap with another time range.
	//
	// 	- **start**: the start hour
	//
	// 	- **end**: the end hour.
	//
	// 	- **bandwidth**: the bandwidth. Unit: KB/s.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// Specifies whether to update the source path if the backup source is empty. Valid values:
	//
	// 	- true: The system replaces the original source path with the specified source path.
	//
	// 	- false: The system does not update the original source path. The system backs up data based on the source path that you specified when you created the backup plan.
	//
	// example:
	//
	// false
	UpdatePaths *bool `json:"UpdatePaths,omitempty" xml:"UpdatePaths,omitempty"`
	// The ID of the backup vault.
	//
	// example:
	//
	// v-0006******q
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s UpdateBackupPlanShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanShrinkRequest) GetChangeListPath() *string {
	return s.ChangeListPath
}

func (s *UpdateBackupPlanShrinkRequest) GetDetailShrink() *string {
	return s.DetailShrink
}

func (s *UpdateBackupPlanShrinkRequest) GetExclude() *string {
	return s.Exclude
}

func (s *UpdateBackupPlanShrinkRequest) GetInclude() *string {
	return s.Include
}

func (s *UpdateBackupPlanShrinkRequest) GetKeepLatestSnapshots() *int64 {
	return s.KeepLatestSnapshots
}

func (s *UpdateBackupPlanShrinkRequest) GetOptions() *string {
	return s.Options
}

func (s *UpdateBackupPlanShrinkRequest) GetOtsDetailShrink() *string {
	return s.OtsDetailShrink
}

func (s *UpdateBackupPlanShrinkRequest) GetPath() []*string {
	return s.Path
}

func (s *UpdateBackupPlanShrinkRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateBackupPlanShrinkRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *UpdateBackupPlanShrinkRequest) GetPrefix() *string {
	return s.Prefix
}

func (s *UpdateBackupPlanShrinkRequest) GetRetention() *int64 {
	return s.Retention
}

func (s *UpdateBackupPlanShrinkRequest) GetRule() []*UpdateBackupPlanShrinkRequestRule {
	return s.Rule
}

func (s *UpdateBackupPlanShrinkRequest) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateBackupPlanShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateBackupPlanShrinkRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *UpdateBackupPlanShrinkRequest) GetUpdatePaths() *bool {
	return s.UpdatePaths
}

func (s *UpdateBackupPlanShrinkRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *UpdateBackupPlanShrinkRequest) SetChangeListPath(v string) *UpdateBackupPlanShrinkRequest {
	s.ChangeListPath = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetDetailShrink(v string) *UpdateBackupPlanShrinkRequest {
	s.DetailShrink = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetExclude(v string) *UpdateBackupPlanShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetInclude(v string) *UpdateBackupPlanShrinkRequest {
	s.Include = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetKeepLatestSnapshots(v int64) *UpdateBackupPlanShrinkRequest {
	s.KeepLatestSnapshots = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetOptions(v string) *UpdateBackupPlanShrinkRequest {
	s.Options = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetOtsDetailShrink(v string) *UpdateBackupPlanShrinkRequest {
	s.OtsDetailShrink = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPath(v []*string) *UpdateBackupPlanShrinkRequest {
	s.Path = v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPlanId(v string) *UpdateBackupPlanShrinkRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPlanName(v string) *UpdateBackupPlanShrinkRequest {
	s.PlanName = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetPrefix(v string) *UpdateBackupPlanShrinkRequest {
	s.Prefix = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetRetention(v int64) *UpdateBackupPlanShrinkRequest {
	s.Retention = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetRule(v []*UpdateBackupPlanShrinkRequestRule) *UpdateBackupPlanShrinkRequest {
	s.Rule = v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetSchedule(v string) *UpdateBackupPlanShrinkRequest {
	s.Schedule = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetSourceType(v string) *UpdateBackupPlanShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetSpeedLimit(v string) *UpdateBackupPlanShrinkRequest {
	s.SpeedLimit = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetUpdatePaths(v bool) *UpdateBackupPlanShrinkRequest {
	s.UpdatePaths = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) SetVaultId(v string) *UpdateBackupPlanShrinkRequest {
	s.VaultId = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateBackupPlanShrinkRequestRule struct {
	// The backup type. Valid value: **COMPLETE**, which indicates full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the region where the remote backup vault resides.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The retention period of the backup data. Unit: days.
	//
	// example:
	//
	// 7
	DestinationRetention *int64 `json:"DestinationRetention,omitempty" xml:"DestinationRetention,omitempty"`
	// Specifies whether to disable the policy.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// Specifies whether to enable remote replication.
	//
	// example:
	//
	// false
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// The retention period of the backup data. Minimum value: 1. Unit: days.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The name of the backup policy.
	//
	// example:
	//
	// rule-test-name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The backup policy. Format: I|{startTime}|{interval}. The system runs the first backup job at a point in time that is specified in the {startTime} parameter and the subsequent backup jobs at an interval that is specified in the {interval} parameter. The system does not run a backup job before the specified point in time. Each backup job, except the first one, starts only after the previous backup job is completed. For example, I|1631685600|P1D specifies that the system runs the first backup job at 14:00:00 on September 15, 2021 and the subsequent backup jobs once a day.
	//
	// startTime: the time at which the system starts to run a backup job. The time must follow the UNIX time format. Unit: seconds. interval: the interval at which the system runs a backup job. The interval must follow the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
	//
	// example:
	//
	// I|1631685600|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s UpdateBackupPlanShrinkRequestRule) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPlanShrinkRequestRule) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanShrinkRequestRule) GetBackupType() *string {
	return s.BackupType
}

func (s *UpdateBackupPlanShrinkRequestRule) GetDestinationRegionId() *string {
	return s.DestinationRegionId
}

func (s *UpdateBackupPlanShrinkRequestRule) GetDestinationRetention() *int64 {
	return s.DestinationRetention
}

func (s *UpdateBackupPlanShrinkRequestRule) GetDisabled() *bool {
	return s.Disabled
}

func (s *UpdateBackupPlanShrinkRequestRule) GetDoCopy() *bool {
	return s.DoCopy
}

func (s *UpdateBackupPlanShrinkRequestRule) GetRetention() *int64 {
	return s.Retention
}

func (s *UpdateBackupPlanShrinkRequestRule) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateBackupPlanShrinkRequestRule) GetSchedule() *string {
	return s.Schedule
}

func (s *UpdateBackupPlanShrinkRequestRule) SetBackupType(v string) *UpdateBackupPlanShrinkRequestRule {
	s.BackupType = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDestinationRegionId(v string) *UpdateBackupPlanShrinkRequestRule {
	s.DestinationRegionId = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDestinationRetention(v int64) *UpdateBackupPlanShrinkRequestRule {
	s.DestinationRetention = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDisabled(v bool) *UpdateBackupPlanShrinkRequestRule {
	s.Disabled = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetDoCopy(v bool) *UpdateBackupPlanShrinkRequestRule {
	s.DoCopy = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetRetention(v int64) *UpdateBackupPlanShrinkRequestRule {
	s.Retention = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetRuleName(v string) *UpdateBackupPlanShrinkRequestRule {
	s.RuleName = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) SetSchedule(v string) *UpdateBackupPlanShrinkRequestRule {
	s.Schedule = &v
	return s
}

func (s *UpdateBackupPlanShrinkRequestRule) Validate() error {
	return dara.Validate(s)
}
