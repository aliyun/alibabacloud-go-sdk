// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *CreateRestoreJobShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateRestoreJobShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateRestoreJobShrinkRequest
	GetCrossAccountUserId() *int64
	SetExclude(v string) *CreateRestoreJobShrinkRequest
	GetExclude() *string
	SetFailbackDetailShrink(v string) *CreateRestoreJobShrinkRequest
	GetFailbackDetailShrink() *string
	SetInclude(v string) *CreateRestoreJobShrinkRequest
	GetInclude() *string
	SetInitiatedByAck(v bool) *CreateRestoreJobShrinkRequest
	GetInitiatedByAck() *bool
	SetOptions(v string) *CreateRestoreJobShrinkRequest
	GetOptions() *string
	SetOtsDetailShrink(v string) *CreateRestoreJobShrinkRequest
	GetOtsDetailShrink() *string
	SetRestoreType(v string) *CreateRestoreJobShrinkRequest
	GetRestoreType() *string
	SetSnapshotHash(v string) *CreateRestoreJobShrinkRequest
	GetSnapshotHash() *string
	SetSnapshotId(v string) *CreateRestoreJobShrinkRequest
	GetSnapshotId() *string
	SetSourceType(v string) *CreateRestoreJobShrinkRequest
	GetSourceType() *string
	SetTargetBucket(v string) *CreateRestoreJobShrinkRequest
	GetTargetBucket() *string
	SetTargetContainer(v string) *CreateRestoreJobShrinkRequest
	GetTargetContainer() *string
	SetTargetContainerClusterId(v string) *CreateRestoreJobShrinkRequest
	GetTargetContainerClusterId() *string
	SetTargetCreateTime(v int64) *CreateRestoreJobShrinkRequest
	GetTargetCreateTime() *int64
	SetTargetFileSystemId(v string) *CreateRestoreJobShrinkRequest
	GetTargetFileSystemId() *string
	SetTargetInstanceId(v string) *CreateRestoreJobShrinkRequest
	GetTargetInstanceId() *string
	SetTargetInstanceName(v string) *CreateRestoreJobShrinkRequest
	GetTargetInstanceName() *string
	SetTargetPath(v string) *CreateRestoreJobShrinkRequest
	GetTargetPath() *string
	SetTargetPrefix(v string) *CreateRestoreJobShrinkRequest
	GetTargetPrefix() *string
	SetTargetTableName(v string) *CreateRestoreJobShrinkRequest
	GetTargetTableName() *string
	SetTargetTime(v int64) *CreateRestoreJobShrinkRequest
	GetTargetTime() *int64
	SetUdmDetailShrink(v string) *CreateRestoreJobShrinkRequest
	GetUdmDetailShrink() *string
	SetUdmRegionId(v string) *CreateRestoreJobShrinkRequest
	GetUdmRegionId() *string
	SetVaultId(v string) *CreateRestoreJobShrinkRequest
	GetVaultId() *string
}

type CreateRestoreJobShrinkRequest struct {
	// The name of the role created in the RAM of the original account for cross-account backup managed by the current account.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Cross-account backup type. Supported values:
	//
	// - SELF_ACCOUNT: Backup within the same account
	//
	// - CROSS_ACCOUNT: Cross-account backup
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The original account ID managed by the current account for cross-account backup.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The path not to be restored. All documents under this path will not be restored. Maximum length is 255 characters.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// Details of restoring to the local environment.
	FailbackDetailShrink *string `json:"FailbackDetail,omitempty" xml:"FailbackDetail,omitempty"`
	// The path to be restored. All documents under this path will be restored. Maximum length is 255 characters.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// Indicates whether it is called by the container service. Default is false.
	//
	// example:
	//
	// false
	InitiatedByAck *bool `json:"InitiatedByAck,omitempty" xml:"InitiatedByAck,omitempty"`
	// Parameters for the restore job.
	//
	// example:
	//
	// {\\"includes\\":[],\\"excludes\\":[],\\"conflictPolicy\\":\\"OVERWRITE_EXISTING\\"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// Details of the Table Store instance.
	OtsDetailShrink *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// The type of the restore destination data source. Possible values:
	//
	//   - **ECS_FILE**: Restore to ECS file.
	//
	//   - **OSS**: Restore to Alibaba Cloud OSS.
	//
	//   - **NAS**: Restore to Alibaba Cloud NAS.
	//
	//   - **OTS_TABLE**: Restore to Alibaba Cloud OTS.
	//
	//   - **UDM_ECS_ROLLBACK**: Restore to Alibaba Cloud ECS whole machine.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The HASH value of the backup snapshot.
	//
	// example:
	//
	// f2fe...
	SnapshotHash *string `json:"SnapshotHash,omitempty" xml:"SnapshotHash,omitempty"`
	// The ID of the backup snapshot.
	//
	// example:
	//
	// s-********************
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The type of the data source. Possible values:
	//
	//   - **ECS_FILE**: Restore ECS file.
	//
	//   - **OSS**: Restore Alibaba Cloud OSS.
	//
	//   - **NAS**: Restore Alibaba Cloud NAS.
	//
	//   - **OTS_TABLE**: Restore to Alibaba Cloud OTS.
	//
	//   - **UDM_ECS**: Restore to Alibaba Cloud ECS whole machine.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Valid only when **RestoreType*	- is **OSS**. Indicates the name of the OSS bucket at the restore destination.
	//
	// example:
	//
	// hbr-backup-oss
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// Details of the target container.
	//
	// example:
	//
	// {\\"host\\":\\"k8s-node1\\",\\"hostPrefix\\":\\"/var/lib/kubelet/pods/4acb31fe-8577-40ff-bc8c-eccabd835f73/volumes/kubernetes.io~csi/pvc-b050b00e-ef17-4792-aab1-1642355cf1f4/mount\\",\\"pvPath\\":\\"/\\"}
	TargetContainer *string `json:"TargetContainer,omitempty" xml:"TargetContainer,omitempty"`
	// The ID of the target container cluster.
	//
	// example:
	//
	// cc-000amjsc7o1h9506oob7
	TargetContainerClusterId *string `json:"TargetContainerClusterId,omitempty" xml:"TargetContainerClusterId,omitempty"`
	// Valid only when **RestoreType*	- is **NAS**. Indicates the creation time of the file system at the restore destination.
	//
	// example:
	//
	// 1554347313
	TargetCreateTime *int64 `json:"TargetCreateTime,omitempty" xml:"TargetCreateTime,omitempty"`
	// Valid only when **RestoreType*	- is **NAS**. Indicates the ID of the file system at the restore destination.
	//
	// example:
	//
	// 005494
	TargetFileSystemId *string `json:"TargetFileSystemId,omitempty" xml:"TargetFileSystemId,omitempty"`
	// Valid only when **RestoreType*	- is **ECS_FILE**. Indicates the ECS instance ID at the restore destination.
	//
	// example:
	//
	// i-*********************
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The name of the target Table Store instance.
	//
	// example:
	//
	// instancename
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// Valid only when **RestoreType*	- is **ECS_FILE**. Indicates the file path at the restore destination.
	//
	// example:
	//
	// C:\\
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// Valid only when **RestoreType*	- is **OSS**. Indicates the object prefix at the restore destination.
	//
	// example:
	//
	// hbr
	TargetPrefix *string `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
	// The name of the data table in the target Table Store.
	//
	// example:
	//
	// tablename
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The time of the Table Store to be restored. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1642496881
	TargetTime *int64 `json:"TargetTime,omitempty" xml:"TargetTime,omitempty"`
	// The parameter is valid only when the SourceType is set to UDM_ECS. It represents the details of the entire machine backup and is a JSON string. Depending on the value of RestoreType, different details must be passed as follows:
	//
	// - **UDM_ECS_DISK**: ECS disk cloning.
	//
	//   - **targetInstanceId**: string (required). Specifies the target ECS instance ID to which the cloned disk will be attached.
	//
	//   - **diskCategory**: string (required). Specifies the type of the target disk.
	//
	//   - **diskPerformanceLevel**: string. When diskCategory is "essd", this indicates the disk performance level, supporting PL0, PL1, PL2, and PL3, with PL1 as the default.
	//
	// - **UDM_ECS_DISK_ROLLBACK**: ECS disk rollback.
	//
	//   - **sourceInstanceId**: string (required). Specifies the source ECS instance ID.
	//
	//   - **forceRestore**: bool (default: false). Indicates whether to force restore. NOTE: If forceRestore is set to true, the disk restoration will proceed even if the backup disk has been unmounted from the original ECS instance or mounted to another instance. Exercise caution when using this option.
	//
	//   - **bootAfterRestore**: bool (default: false). Indicates whether to start the ECS instance after restoration.
	//
	// - **UDM_ECS**: Full ECS cloning.
	//
	//   - **bootAfterRestore**: bool (default: false). Indicates whether to start the ECS instance after restoration.
	//
	//   - **diskCategory**: string (required). Specifies the type of the target disk.
	//
	//   - **diskPerformanceLevel**: string. When diskCategory is "essd", this indicates the disk performance level (PL0/PL1/PL2/PL3), defaulting to PL1.
	//
	//   - **instanceType**: string (required). Specifies the specification of the target ECS instance.
	//
	//   - **restoredNetwork**: string (required). Specifies the vSwitch ID for the target ECS instance.
	//
	//   - **securityGroup**: string (required). Specifies the security group ID for the target ECS instance.
	//
	//   - **restoredName:*	- string (required). Specifies the instance name of the target ECS instance.
	//
	//   - **restoredHostName**: string (required). Specifies the host name of the target ECS instance.
	//
	//   - **allocatePublicIp**: bool (default: false). Indicates whether to assign a public IP to the target ECS instance.
	//
	//   - **privateIpAddress**: string. Specifies the internal IP address of the target ECS instance. If not specified, an IP will be assigned via DHCP.
	//
	// - **UDM_ECS_ROLLBACK**: Full ECS rollback.
	//
	//   - **sourceInstanceId**: string (required). Specifies the source ECS instance ID.
	//
	//   - **forceRestore**: bool (default: false). Indicates whether to force restore. NOTE: If forceRestore is set to true, the disk restoration will proceed even if the backup disk has been unmounted from the original ECS instance or mounted to another instance. Exercise caution when using this option.
	//
	//   - **bootAfterRestore**: bool (default: false). Indicates whether to start the ECS instance after restoration.
	//
	// example:
	//
	// {\\"sourceInstanceId\\":\\"i-uf62te6pm3iwsyxyz66q\\",\\"bootAfterRestore\\":false}
	UdmDetailShrink *string `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
	// Valid only when **SourceType*	- is **UDM_ECS**. Indicates the target region for the restore.
	//
	// example:
	//
	// cn-shanghai
	UdmRegionId *string `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
	// The ID of the backup vault that the snapshot belongs to.
	//
	// example:
	//
	// v-*********************
	VaultId *string `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
}

func (s CreateRestoreJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateRestoreJobShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateRestoreJobShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateRestoreJobShrinkRequest) GetExclude() *string {
	return s.Exclude
}

func (s *CreateRestoreJobShrinkRequest) GetFailbackDetailShrink() *string {
	return s.FailbackDetailShrink
}

func (s *CreateRestoreJobShrinkRequest) GetInclude() *string {
	return s.Include
}

func (s *CreateRestoreJobShrinkRequest) GetInitiatedByAck() *bool {
	return s.InitiatedByAck
}

func (s *CreateRestoreJobShrinkRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateRestoreJobShrinkRequest) GetOtsDetailShrink() *string {
	return s.OtsDetailShrink
}

func (s *CreateRestoreJobShrinkRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateRestoreJobShrinkRequest) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *CreateRestoreJobShrinkRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateRestoreJobShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRestoreJobShrinkRequest) GetTargetBucket() *string {
	return s.TargetBucket
}

func (s *CreateRestoreJobShrinkRequest) GetTargetContainer() *string {
	return s.TargetContainer
}

func (s *CreateRestoreJobShrinkRequest) GetTargetContainerClusterId() *string {
	return s.TargetContainerClusterId
}

func (s *CreateRestoreJobShrinkRequest) GetTargetCreateTime() *int64 {
	return s.TargetCreateTime
}

func (s *CreateRestoreJobShrinkRequest) GetTargetFileSystemId() *string {
	return s.TargetFileSystemId
}

func (s *CreateRestoreJobShrinkRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *CreateRestoreJobShrinkRequest) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *CreateRestoreJobShrinkRequest) GetTargetPath() *string {
	return s.TargetPath
}

func (s *CreateRestoreJobShrinkRequest) GetTargetPrefix() *string {
	return s.TargetPrefix
}

func (s *CreateRestoreJobShrinkRequest) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *CreateRestoreJobShrinkRequest) GetTargetTime() *int64 {
	return s.TargetTime
}

func (s *CreateRestoreJobShrinkRequest) GetUdmDetailShrink() *string {
	return s.UdmDetailShrink
}

func (s *CreateRestoreJobShrinkRequest) GetUdmRegionId() *string {
	return s.UdmRegionId
}

func (s *CreateRestoreJobShrinkRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateRestoreJobShrinkRequest) SetCrossAccountRoleName(v string) *CreateRestoreJobShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetCrossAccountType(v string) *CreateRestoreJobShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetCrossAccountUserId(v int64) *CreateRestoreJobShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetExclude(v string) *CreateRestoreJobShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetFailbackDetailShrink(v string) *CreateRestoreJobShrinkRequest {
	s.FailbackDetailShrink = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetInclude(v string) *CreateRestoreJobShrinkRequest {
	s.Include = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetInitiatedByAck(v bool) *CreateRestoreJobShrinkRequest {
	s.InitiatedByAck = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetOptions(v string) *CreateRestoreJobShrinkRequest {
	s.Options = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetOtsDetailShrink(v string) *CreateRestoreJobShrinkRequest {
	s.OtsDetailShrink = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetRestoreType(v string) *CreateRestoreJobShrinkRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetSnapshotHash(v string) *CreateRestoreJobShrinkRequest {
	s.SnapshotHash = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetSnapshotId(v string) *CreateRestoreJobShrinkRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetSourceType(v string) *CreateRestoreJobShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetBucket(v string) *CreateRestoreJobShrinkRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetContainer(v string) *CreateRestoreJobShrinkRequest {
	s.TargetContainer = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetContainerClusterId(v string) *CreateRestoreJobShrinkRequest {
	s.TargetContainerClusterId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetCreateTime(v int64) *CreateRestoreJobShrinkRequest {
	s.TargetCreateTime = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetFileSystemId(v string) *CreateRestoreJobShrinkRequest {
	s.TargetFileSystemId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetInstanceId(v string) *CreateRestoreJobShrinkRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetInstanceName(v string) *CreateRestoreJobShrinkRequest {
	s.TargetInstanceName = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetPath(v string) *CreateRestoreJobShrinkRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetPrefix(v string) *CreateRestoreJobShrinkRequest {
	s.TargetPrefix = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetTableName(v string) *CreateRestoreJobShrinkRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetTargetTime(v int64) *CreateRestoreJobShrinkRequest {
	s.TargetTime = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetUdmDetailShrink(v string) *CreateRestoreJobShrinkRequest {
	s.UdmDetailShrink = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetUdmRegionId(v string) *CreateRestoreJobShrinkRequest {
	s.UdmRegionId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) SetVaultId(v string) *CreateRestoreJobShrinkRequest {
	s.VaultId = &v
	return s
}

func (s *CreateRestoreJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
