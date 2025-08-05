// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupType(v string) *CreateBackupJobShrinkRequest
	GetBackupType() *string
	SetClusterId(v string) *CreateBackupJobShrinkRequest
	GetClusterId() *string
	SetContainerClusterId(v string) *CreateBackupJobShrinkRequest
	GetContainerClusterId() *string
	SetContainerResources(v string) *CreateBackupJobShrinkRequest
	GetContainerResources() *string
	SetCrossAccountRoleName(v string) *CreateBackupJobShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateBackupJobShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateBackupJobShrinkRequest
	GetCrossAccountUserId() *int64
	SetDetailShrink(v string) *CreateBackupJobShrinkRequest
	GetDetailShrink() *string
	SetExclude(v string) *CreateBackupJobShrinkRequest
	GetExclude() *string
	SetInclude(v string) *CreateBackupJobShrinkRequest
	GetInclude() *string
	SetInitiatedByAck(v bool) *CreateBackupJobShrinkRequest
	GetInitiatedByAck() *bool
	SetInstanceId(v string) *CreateBackupJobShrinkRequest
	GetInstanceId() *string
	SetJobName(v string) *CreateBackupJobShrinkRequest
	GetJobName() *string
	SetOptions(v string) *CreateBackupJobShrinkRequest
	GetOptions() *string
	SetRetention(v int64) *CreateBackupJobShrinkRequest
	GetRetention() *int64
	SetSourceType(v string) *CreateBackupJobShrinkRequest
	GetSourceType() *string
	SetSpeedLimit(v string) *CreateBackupJobShrinkRequest
	GetSpeedLimit() *string
	SetVaultId(v string) *CreateBackupJobShrinkRequest
	GetVaultId() *string
}

type CreateBackupJobShrinkRequest struct {
	// The backup type. This parameter is required only if you set the SourceType parameter to UDM_ECS.
	//
	// 	- **COMPLETE**: full backup
	//
	// example:
	//
	// INCREMENTAL
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// You do not need to specify this parameter.
	//
	// example:
	//
	// cl-00068btz******oku
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// You do not need to specify this parameter.
	//
	// example:
	//
	// cc-000xxxxxxxxxxxxxxi00
	ContainerClusterId *string `json:"ContainerClusterId,omitempty" xml:"ContainerClusterId,omitempty"`
	// You do not need to specify this parameter.
	//
	// example:
	//
	// [{\\"resourceType\\":\\"PV\\",\\"backupMethod\\":\\"FILE\\",\\"resourceId\\":\\"674dac6d-74cd-47e9-a675-09e2f10d2c45\\",\\"resourceInfo\\":\\"{\\\\\\"pv_name\\\\\\":\\\\\\"nas-650dac6d-74cd-47e9-a675-09e2f10d2c45\\\\\\",\\\\\\"pv_size\\\\\\":\\\\\\"8Gi\\\\\\",\\\\\\"storage_class\\\\\\":\\\\\\"alibabacloud-cnfs-nas\\\\\\",\\\\\\"pvc_name\\\\\\":\\\\\\"data-postgresql-default-0\\\\\\",\\\\\\"namespace\\\\\\":\\\\\\"database\\\\\\"}\\",\\"host\\":\\"cn-huhehaote.192.168.13.133\\",\\"hostPrefix\\":\\"6f5e758e-8d35-4584-b9ce-8333adfc7547/volumes/kubernetes.io~csi/nas-670dac6d-74cd-47e9-a675-09e2f10d2c45/mount\\",\\"pvPath\\":\\"/\\"}]
	ContainerResources *string `json:"ContainerResources,omitempty" xml:"ContainerResources,omitempty"`
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 158975xxxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **UDM_ECS**. The value is a JSON string. Valid values:
	//
	// 	- doCopy: specifies whether to enable remote replication.
	//
	// 	- destinationRegionId: the destination region for remote replication.
	//
	// 	- destinationRetention: the retention period of the backup point for remote replication.
	//
	// 	- diskIdList: the IDs of the disks that are to be backed up. If this parameter is left empty, all disks are backed up.
	//
	// 	- snapshotGroup: specifies whether to use a snapshot-consistent group. This parameter is valid only if all disks of the ECS instance are Enterprise SSDs (ESSDs).
	//
	// 	- appConsistent: specifies whether to use the application-consistent backup feature. This parameter must be used with the preScriptPath and postScriptPath parameters.
	//
	// 	- preScriptPath: the path to the pre-freeze scripts.
	//
	// 	- postScriptPath: the path to the post-thaw scripts.
	//
	// 	- enableWriters: This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies whether to create application-consistent snapshots.
	//
	//     	- true: creates application-consistent snapshots.
	//
	//     	- false: creates file system-consistent snapshots.
	//
	// 	- enableFsFreeze: This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies whether to enable Linux fsfreeze to put file systems into the read-only state before application-consistent snapshots are created. Default value: true.
	//
	// 	- timeoutSeconds: This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies the I/O freeze timeout period. Default value: 30. Unit: seconds.
	//
	// example:
	//
	// {
	//
	//     "doCopy": false,
	//
	//     "destinationRegionId": "",
	//
	//     "destinationRetention": null,
	//
	//     "diskIdList": [],
	//
	//     "snapshotGroup": false,
	//
	//     "appConsistent": false,
	//
	//     "enableWriters": true,
	//
	//     "preScriptPath": "",
	//
	//     "postScriptPath": "",
	//
	//     "enableFsFreeze": true,
	//
	//     "timeoutInSeconds": 60
	//
	// }
	DetailShrink *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// This parameter does not take effect if you set the **SourceType*	- parameter to **UDM_ECS**. This parameter specifies the paths to the files that are excluded from the backup job. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter does not take effect if you set the **SourceType*	- parameter to **UDM_ECS**. This parameter specifies the paths to the files that are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// false or left empty
	//
	// example:
	//
	// false
	InitiatedByAck *bool `json:"InitiatedByAck,omitempty" xml:"InitiatedByAck,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **UDM_ECS**. This parameter specifies the ID of the ECS instance.
	//
	// example:
	//
	// i-bp1xxxxxxxxxxxxxxysm
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the backup job.
	//
	// example:
	//
	// k8s-backup-infra-20220131150046-hbr
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// You do not need to specify this parameter.
	//
	// example:
	//
	// {"UseVSS":false}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The retention period of the backup data. Unit: days.
	//
	// example:
	//
	// 15
	Retention *int64 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: Elastic Compute Service (ECS) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// CONTAINER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter does not take effect if you set the **SourceType*	- parameter to **UDM_ECS**. This parameter specifies the throttling rules. Format: `{start}|{end}|{bandwidth}`. Separate multiple throttling rules with vertical bars (|). A specified time range cannot overlap with another time range.
	//
	// 	- **start**: the start hour.
	//
	// 	- **end**: the end hour.
	//
	// 	- **bandwidth**: the bandwidth. Unit: KB/s.
	//
	// example:
	//
	// 0:24:NaN
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// The ID of the backup vault. This parameter is not required if you set the SourceType parameter to UDM_ECS.
	//
	// example:
	//
	// v-000xxxxxxxxxxxxxxy1v
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateBackupJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupJobShrinkRequest) GetBackupType() *string {
	return s.BackupType
}

func (s *CreateBackupJobShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateBackupJobShrinkRequest) GetContainerClusterId() *string {
	return s.ContainerClusterId
}

func (s *CreateBackupJobShrinkRequest) GetContainerResources() *string {
	return s.ContainerResources
}

func (s *CreateBackupJobShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateBackupJobShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateBackupJobShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateBackupJobShrinkRequest) GetDetailShrink() *string {
	return s.DetailShrink
}

func (s *CreateBackupJobShrinkRequest) GetExclude() *string {
	return s.Exclude
}

func (s *CreateBackupJobShrinkRequest) GetInclude() *string {
	return s.Include
}

func (s *CreateBackupJobShrinkRequest) GetInitiatedByAck() *bool {
	return s.InitiatedByAck
}

func (s *CreateBackupJobShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateBackupJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateBackupJobShrinkRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateBackupJobShrinkRequest) GetRetention() *int64 {
	return s.Retention
}

func (s *CreateBackupJobShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateBackupJobShrinkRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *CreateBackupJobShrinkRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateBackupJobShrinkRequest) SetBackupType(v string) *CreateBackupJobShrinkRequest {
	s.BackupType = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetClusterId(v string) *CreateBackupJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetContainerClusterId(v string) *CreateBackupJobShrinkRequest {
	s.ContainerClusterId = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetContainerResources(v string) *CreateBackupJobShrinkRequest {
	s.ContainerResources = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetCrossAccountRoleName(v string) *CreateBackupJobShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetCrossAccountType(v string) *CreateBackupJobShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetCrossAccountUserId(v int64) *CreateBackupJobShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetDetailShrink(v string) *CreateBackupJobShrinkRequest {
	s.DetailShrink = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetExclude(v string) *CreateBackupJobShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetInclude(v string) *CreateBackupJobShrinkRequest {
	s.Include = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetInitiatedByAck(v bool) *CreateBackupJobShrinkRequest {
	s.InitiatedByAck = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetInstanceId(v string) *CreateBackupJobShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetJobName(v string) *CreateBackupJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetOptions(v string) *CreateBackupJobShrinkRequest {
	s.Options = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetRetention(v int64) *CreateBackupJobShrinkRequest {
	s.Retention = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetSourceType(v string) *CreateBackupJobShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetSpeedLimit(v string) *CreateBackupJobShrinkRequest {
	s.SpeedLimit = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) SetVaultId(v string) *CreateBackupJobShrinkRequest {
	s.VaultId = &v
	return s
}

func (s *CreateBackupJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
