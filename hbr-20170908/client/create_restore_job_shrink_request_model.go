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
	SetEdition(v string) *CreateRestoreJobShrinkRequest
	GetEdition() *string
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
	// The name of the RAM role created in the source account for cross-account backup managed by the current account.
	//
	// example:
	//
	// BackupRole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The cross-account backup type. Valid values:
	//
	// - SELF_ACCOUNT: backup within the current account.
	//
	// - CROSS_ACCOUNT: cross-account backup.
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source account for cross-account backup managed by the current account.
	//
	// example:
	//
	// 158975xxxxx4625
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The Cloud Backup feature edition. Valid values:
	//
	// - **STANDARD**: Standard Edition. This is the default value.
	//
	// - **BASIC**: Essential Edition. Currently, only ECS File Backup Essential Edition is supported.
	//
	// example:
	//
	// STANDARD
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The path to exclude from restoration. All files under this path are not restored. Maximum length: 255 characters.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// The details of the restoration to the local host.
	FailbackDetailShrink *string `json:"FailbackDetail,omitempty" xml:"FailbackDetail,omitempty"`
	// The path to restore. All files under this path are restored. Maximum length: 255 characters.
	//
	// example:
	//
	// ["/home/alice/*.pdf", "/home/bob/*.txt"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// Specifies whether the operation is invoked by Container Service. Default value: false.
	//
	// example:
	//
	// false
	InitiatedByAck *bool `json:"InitiatedByAck,omitempty" xml:"InitiatedByAck,omitempty"`
	// The restore job parameters.
	//
	// example:
	//
	// {\\"includes\\":[],\\"excludes\\":[],\\"conflictPolicy\\":\\"OVERWRITE_EXISTING\\"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The details of the Tablestore instance.
	OtsDetailShrink *string `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
	// The data source type of the restore destination. Valid values:
	//
	//   - **ECS_FILE**: restores to an ECS file.
	//
	//   - **OSS**: restores to Alibaba Cloud OSS.
	//
	//   - **NAS**: restores to Alibaba Cloud NAS.
	//
	//   - **COMMON_FILE_SYSTEM**: restores to CPFS.
	//
	//   - **OTS_TABLE**: restores to Alibaba Cloud OTS.
	//
	//   - **UDM_ECS_ROLLBACK**: restores to an Alibaba Cloud ECS instance (full-copy migration).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	RestoreType *string `json:"RestoreType,omitempty" xml:"RestoreType,omitempty"`
	// The hash value of the backup snapshot.
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
	// The data source type. Valid values:
	//
	//   - **ECS_FILE**: restores ECS files.
	//
	//   - **OSS**: restores Alibaba Cloud OSS.
	//
	//   - **NAS**: restores Alibaba Cloud NAS.
	//
	//   - **COMMON_FILE_SYSTEM**: restores to CPFS.
	//
	//   - **OTS_TABLE**: restores to Alibaba Cloud OTS.
	//
	//   - **UDM_ECS**: restores to an Alibaba Cloud ECS instance (full-copy migration).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is valid only when **RestoreType*	- is set to **OSS**. The name of the destination OSS bucket.
	//
	// example:
	//
	// hbr-backup-oss
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	// The details of the target container for restoration.
	//
	// example:
	//
	// {\\"host\\":\\"k8s-node1\\",\\"hostPrefix\\":\\"/var/lib/kubelet/pods/4acb31fe-8577-40ff-bc8c-eccabd835f73/volumes/kubernetes.io~csi/pvc-b050b00e-ef17-4792-aab1-1642355cf1f4/mount\\",\\"pvPath\\":\\"/\\"}
	TargetContainer *string `json:"TargetContainer,omitempty" xml:"TargetContainer,omitempty"`
	// The ID of the target container cluster for restoration.
	//
	// example:
	//
	// cc-000amjsc7o1h9506oob7
	TargetContainerClusterId *string `json:"TargetContainerClusterId,omitempty" xml:"TargetContainerClusterId,omitempty"`
	// This parameter is valid only when **RestoreType*	- is set to **NAS**. The creation time of the destination file system. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1554347313
	TargetCreateTime *int64 `json:"TargetCreateTime,omitempty" xml:"TargetCreateTime,omitempty"`
	// This parameter is valid only when **RestoreType*	- is set to **NAS**. The file system ID of the restore destination.
	//
	// example:
	//
	// 005494
	TargetFileSystemId *string `json:"TargetFileSystemId,omitempty" xml:"TargetFileSystemId,omitempty"`
	// This parameter is valid only when **RestoreType*	- is set to **ECS_FILE**. The ECS instance ID of the restore destination.
	//
	// example:
	//
	// i-*********************
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The name of the target Tablestore instance for restoration.
	//
	// example:
	//
	// instancename
	TargetInstanceName *string `json:"TargetInstanceName,omitempty" xml:"TargetInstanceName,omitempty"`
	// This parameter is valid only when **RestoreType*	- is set to **ECS_FILE**. The file path of the restore destination.
	//
	// example:
	//
	// C:\\
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// This parameter is valid only when **RestoreType*	- is set to **OSS**. The object prefix of the restore destination.
	//
	// example:
	//
	// hbr
	TargetPrefix *string `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
	// The name of the target data table in Tablestore for restoration.
	//
	// example:
	//
	// tablename
	TargetTableName *string `json:"TargetTableName,omitempty" xml:"TargetTableName,omitempty"`
	// The point in time to which the Tablestore data is restored. This value is a UNIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1642496881
	TargetTime *int64 `json:"TargetTime,omitempty" xml:"TargetTime,omitempty"`
	// This parameter is valid only when SourceType is set to UDM_ECS. The details of the full-copy migration backup. This parameter is a JSON string. The details vary depending on the value of RestoreType:
	//
	// - **UDM_ECS_DISK**: ECS cloud disk clone.
	//
	//   - **targetInstanceId**: string type, required. Instance ID of the target ECS instance to which the cloned cloud disk is attached.
	//
	//   - **diskCategory**: string type, required. The type of the target cloud disk.
	//
	//   - **diskPerformanceLevel**: string type. If diskCategory is set to essd, this parameter specifies the performance level (PL) of the cloud disk. Valid values: PL0, PL1, PL2, and PL3. Default value: PL1.
	//
	// - **UDM_ECS_DISK_ROLLBACK**: ECS cloud disk restoration.
	//
	//   - **sourceInstanceId**: string type, required. Instance ID of the source ECS instance.
	//
	//   - **foreceRestore**: bool type. Default value: false. Specifies whether to forcibly restore. If foreceRestore is set to true, the restore job still restores the cloud disk even if the backed-up cloud disk has been unmounted from the original ECS instance or attached to a new ECS instance. Proceed with caution.
	//
	//   - **bootAfterRestore**: bool type. Default value: false. Specifies whether to start the ECS instance after restoration.
	//
	// - **UDM_ECS**: ECS full-copy clone.
	//
	//   - **bootAfterRestore**: bool type. Default value: false. Specifies whether to start the ECS instance after restoration.
	//
	//   - **diskCategory**: string type, required. The type of the target cloud disk.
	//
	//   - **diskPerformanceLevel**: string type. If diskCategory is set to essd, this parameter specifies the performance level (PL) of the cloud disk. Valid values: PL0, PL1, PL2, and PL3. Default value: PL1.
	//
	//   - **instanceType**: string type, required. The instance type of the target ECS instance.
	//
	//   - **restoredNetwork**: string type, required. The vSwitch ID of the target ECS instance.
	//
	//   - **securityGroup**: string type, required. The security group ID of the target ECS instance.
	//
	//   - **restoredName**: string type, required. The instance name of the target ECS instance.
	//
	//   - **restoredHostName**: string type, required. The hostname of the target ECS instance.
	//
	//   - **allocatePublicIp**: bool type. Default value: false. Specifies whether to assign a public IP address to the target ECS instance.
	//
	//   - **privateIpAddress**: string type. The internal IP address of the target ECS instance. If this parameter is not specified, DHCP is used to randomly assign an IP address.
	//
	// - **UDM_ECS_ROLLBACK**: ECS full-copy restoration.
	//
	//   - **sourceInstanceId**: string type, required. Instance ID of the source ECS instance.
	//
	//   - **forceRestore**: bool type. Default value: false. Specifies whether to forcibly restore. If foreceRestore is set to true, the restore job still restores the cloud disk even if the backed-up cloud disk has been unmounted from the original ECS instance or attached to a new ECS instance. Proceed with caution.
	//
	//   - **bootAfterRestore**: bool type. Default value: false. Specifies whether to start the ECS instance after restoration.
	//
	// example:
	//
	// {\\"sourceInstanceId\\":\\"i-uf62te6pm3iwsyxyz66q\\",\\"bootAfterRestore\\":false}
	UdmDetailShrink *string `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **UDM_ECS**. The destination region for restoration.
	//
	// example:
	//
	// cn-shanghai
	UdmRegionId *string `json:"UdmRegionId,omitempty" xml:"UdmRegionId,omitempty"`
	// The ID of the backup vault to which the backup snapshot belongs.
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

func (s *CreateRestoreJobShrinkRequest) GetEdition() *string {
	return s.Edition
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

func (s *CreateRestoreJobShrinkRequest) SetEdition(v string) *CreateRestoreJobShrinkRequest {
	s.Edition = &v
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
