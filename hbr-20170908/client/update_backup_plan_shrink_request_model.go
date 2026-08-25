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
	SetEdition(v string) *UpdateBackupPlanShrinkRequest
	GetEdition() *string
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
	// The configuration for the incremental file synchronization list. (This parameter is required only for file synchronization.)
	//
	// example:
	//
	// {"dataSourceId": "ds-123456789", "path": "/changelist"}
	ChangeListPath *string `json:"ChangeListPath,omitempty" xml:"ChangeListPath,omitempty"`
	// The details of the ECS instance backup. This is a JSON string.
	//
	// - snapshotGroup: Specifies whether to use a snapshot-consistent group. This feature is available only when all disks of the instance are Enhanced Solid-State Drives (ESSDs).
	//
	// - appConsistent: Specifies whether to enable application consistency. You must also configure the preScriptPath and postScriptPath parameters.
	//
	// - preScriptPath: The path to the pre-freeze script.
	//
	// - postScriptPath: The path to the post-thaw script.
	//
	// example:
	//
	// {\\"EnableFsFreeze\\":true,\\"appConsistent\\":false,\\"postScriptPath\\":\\"\\",\\"preScriptPath\\":\\"\\",\\"snapshotGroup\\":true,\\"timeoutInSeconds\\":60}
	DetailShrink *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The edition. Valid values are BASIC and STANDARD. The default value is STANDARD.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. This parameter specifies the paths to the files to exclude from the backup. All files in the specified paths are not backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. This parameter specifies the paths to the files to back up. All files in the specified paths are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// Specifies whether to permanently retain the latest backup version.
	//
	// - 0: No
	//
	// - 1: Yes
	//
	// example:
	//
	// 1
	KeepLatestSnapshots *int64 `json:"KeepLatestSnapshots,omitempty" xml:"KeepLatestSnapshots,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. This parameter specifies whether to use Volume Shadow Copy Service (VSS) to define the backup path.
	//
	// - This feature is available only for Windows ECS instances.
	//
	// - If data changes occur in the backup source, set this parameter to `["UseVSS":true]` to ensure data consistency.
	//
	// - If you enable VSS, you cannot back up multiple file directories at the same time.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details of the Tablestore instance.
	OtsDetailShrink *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// The backup paths.
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
	// This parameter is required only when **SourceType*	- is set to **OSS**. This parameter specifies the prefix of objects to back up. After you specify a prefix, only objects that match the prefix are backed up.
	//
	// example:
	//
	// oss-prefix
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	// The number of days to retain backups. The minimum value is 1.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The rules of the backup plan.
	Rule []*UpdateBackupPlanShrinkRequestRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
	// The backup policy. Use the `I|{startTime}|{interval}` format. This specifies that a backup job runs at a recurring interval. The `{startTime}` is when the backup starts. The `{interval}` is the time between jobs. HBR does not run overdue backup jobs. If the previous backup job is not finished, the next one does not start. For example, `I|1631685600|P1D` means the backup runs once a day, starting at 14:00:00 on September 15, 2021.
	//
	// - **startTime**: The start time of the backup. This is a UNIX timestamp in seconds.
	//
	// - **interval**: The time interval. Use the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
	//
	// example:
	//
	// I|1602673264|P1D
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// The type of the data source. Valid values:
	//
	// - **ECS_FILE**: Backs up ECS files.
	//
	// - **OSS**: Backs up Alibaba Cloud OSS.
	//
	// - **NAS**: Backs up Alibaba Cloud NAS.
	//
	// - **OTS**: Backs up Alibaba Cloud Tablestore.
	//
	// - **UDM_ECS**: Backs up an entire ECS instance.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE**. This parameter specifies traffic shaping for backups. Traffic shaping helps you control backup traffic during peak business hours to avoid affecting your services. The format is `{start}|{end}|{bandwidth}`. You can specify multiple traffic shaping rules. Separate them with vertical bars (|). The time ranges of the rules cannot overlap.
	//
	// - **start**: The start hour.
	//
	// - **end**: The end hour.
	//
	// - **bandwidth**: The maximum speed. Unit: KB/s.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// Specifies whether to update the backup paths if the Path parameter is empty.
	//
	// - true: Updates the backup paths based on the paths specified in this call.
	//
	// - false: Does not update the backup paths. The backup paths that were configured when the backup plan was created are used.
	//
	// example:
	//
	// false
	UpdatePaths *bool `json:"UpdatePaths,omitempty" xml:"UpdatePaths,omitempty"`
	// The ID of the backup repository.
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

func (s *UpdateBackupPlanShrinkRequest) GetEdition() *string {
	return s.Edition
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

func (s *UpdateBackupPlanShrinkRequest) SetEdition(v string) *UpdateBackupPlanShrinkRequest {
	s.Edition = &v
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
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateBackupPlanShrinkRequestRule struct {
	// The backup type. Set the value to **COMPLETE**. This indicates a full backup.
	//
	// example:
	//
	// COMPLETE
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The ID of the destination region for the geo-redundant backup.
	//
	// example:
	//
	// cn-shanghai
	DestinationRegionId *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	// The number of days to retain the geo-redundant backup.
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
	// Specifies whether to enable geo-redundant replication.
	//
	// example:
	//
	// false
	DoCopy *bool `json:"DoCopy,omitempty" xml:"DoCopy,omitempty"`
	// The number of days to retain backups. The minimum value is 1.
	//
	// example:
	//
	// 7
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// rule-test-name
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The backup policy. Use the I|{startTime}|{interval} format. This specifies that a backup job runs at a recurring interval. The {startTime} is when the backup starts. The {interval} is the time between jobs. HBR does not run overdue backup jobs. If the previous backup job is not finished, the next one does not start. For example, I|1631685600|P1D means the backup runs once a day, starting at 14:00:00 on September 15, 2021.
	//
	// startTime: The start time of the backup. This is a UNIX timestamp in seconds. interval: The time interval. Use the ISO 8601 standard. For example, PT1H specifies an interval of one hour. P1D specifies an interval of one day.
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
