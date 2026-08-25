// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRestoreJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *CreateRestoreJobRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *CreateRestoreJobRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *CreateRestoreJobRequest
	GetCrossAccountUserId() *int64
	SetEdition(v string) *CreateRestoreJobRequest
	GetEdition() *string
	SetExclude(v string) *CreateRestoreJobRequest
	GetExclude() *string
	SetFailbackDetail(v map[string]interface{}) *CreateRestoreJobRequest
	GetFailbackDetail() map[string]interface{}
	SetInclude(v string) *CreateRestoreJobRequest
	GetInclude() *string
	SetInitiatedByAck(v bool) *CreateRestoreJobRequest
	GetInitiatedByAck() *bool
	SetOptions(v string) *CreateRestoreJobRequest
	GetOptions() *string
	SetOtsDetail(v *OtsTableRestoreDetail) *CreateRestoreJobRequest
	GetOtsDetail() *OtsTableRestoreDetail
	SetRestoreType(v string) *CreateRestoreJobRequest
	GetRestoreType() *string
	SetSnapshotHash(v string) *CreateRestoreJobRequest
	GetSnapshotHash() *string
	SetSnapshotId(v string) *CreateRestoreJobRequest
	GetSnapshotId() *string
	SetSourceType(v string) *CreateRestoreJobRequest
	GetSourceType() *string
	SetTargetBucket(v string) *CreateRestoreJobRequest
	GetTargetBucket() *string
	SetTargetContainer(v string) *CreateRestoreJobRequest
	GetTargetContainer() *string
	SetTargetContainerClusterId(v string) *CreateRestoreJobRequest
	GetTargetContainerClusterId() *string
	SetTargetCreateTime(v int64) *CreateRestoreJobRequest
	GetTargetCreateTime() *int64
	SetTargetFileSystemId(v string) *CreateRestoreJobRequest
	GetTargetFileSystemId() *string
	SetTargetInstanceId(v string) *CreateRestoreJobRequest
	GetTargetInstanceId() *string
	SetTargetInstanceName(v string) *CreateRestoreJobRequest
	GetTargetInstanceName() *string
	SetTargetPath(v string) *CreateRestoreJobRequest
	GetTargetPath() *string
	SetTargetPrefix(v string) *CreateRestoreJobRequest
	GetTargetPrefix() *string
	SetTargetTableName(v string) *CreateRestoreJobRequest
	GetTargetTableName() *string
	SetTargetTime(v int64) *CreateRestoreJobRequest
	GetTargetTime() *int64
	SetUdmDetail(v map[string]interface{}) *CreateRestoreJobRequest
	GetUdmDetail() map[string]interface{}
	SetUdmRegionId(v string) *CreateRestoreJobRequest
	GetUdmRegionId() *string
	SetVaultId(v string) *CreateRestoreJobRequest
	GetVaultId() *string
}

type CreateRestoreJobRequest struct {
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
	FailbackDetail map[string]interface{} `json:"FailbackDetail,omitempty" xml:"FailbackDetail,omitempty"`
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
	OtsDetail *OtsTableRestoreDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
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
	UdmDetail map[string]interface{} `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
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

func (s CreateRestoreJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRestoreJobRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreateRestoreJobRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreateRestoreJobRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreateRestoreJobRequest) GetEdition() *string {
	return s.Edition
}

func (s *CreateRestoreJobRequest) GetExclude() *string {
	return s.Exclude
}

func (s *CreateRestoreJobRequest) GetFailbackDetail() map[string]interface{} {
	return s.FailbackDetail
}

func (s *CreateRestoreJobRequest) GetInclude() *string {
	return s.Include
}

func (s *CreateRestoreJobRequest) GetInitiatedByAck() *bool {
	return s.InitiatedByAck
}

func (s *CreateRestoreJobRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateRestoreJobRequest) GetOtsDetail() *OtsTableRestoreDetail {
	return s.OtsDetail
}

func (s *CreateRestoreJobRequest) GetRestoreType() *string {
	return s.RestoreType
}

func (s *CreateRestoreJobRequest) GetSnapshotHash() *string {
	return s.SnapshotHash
}

func (s *CreateRestoreJobRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateRestoreJobRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRestoreJobRequest) GetTargetBucket() *string {
	return s.TargetBucket
}

func (s *CreateRestoreJobRequest) GetTargetContainer() *string {
	return s.TargetContainer
}

func (s *CreateRestoreJobRequest) GetTargetContainerClusterId() *string {
	return s.TargetContainerClusterId
}

func (s *CreateRestoreJobRequest) GetTargetCreateTime() *int64 {
	return s.TargetCreateTime
}

func (s *CreateRestoreJobRequest) GetTargetFileSystemId() *string {
	return s.TargetFileSystemId
}

func (s *CreateRestoreJobRequest) GetTargetInstanceId() *string {
	return s.TargetInstanceId
}

func (s *CreateRestoreJobRequest) GetTargetInstanceName() *string {
	return s.TargetInstanceName
}

func (s *CreateRestoreJobRequest) GetTargetPath() *string {
	return s.TargetPath
}

func (s *CreateRestoreJobRequest) GetTargetPrefix() *string {
	return s.TargetPrefix
}

func (s *CreateRestoreJobRequest) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *CreateRestoreJobRequest) GetTargetTime() *int64 {
	return s.TargetTime
}

func (s *CreateRestoreJobRequest) GetUdmDetail() map[string]interface{} {
	return s.UdmDetail
}

func (s *CreateRestoreJobRequest) GetUdmRegionId() *string {
	return s.UdmRegionId
}

func (s *CreateRestoreJobRequest) GetVaultId() *string {
	return s.VaultId
}

func (s *CreateRestoreJobRequest) SetCrossAccountRoleName(v string) *CreateRestoreJobRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreateRestoreJobRequest) SetCrossAccountType(v string) *CreateRestoreJobRequest {
	s.CrossAccountType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetCrossAccountUserId(v int64) *CreateRestoreJobRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetEdition(v string) *CreateRestoreJobRequest {
	s.Edition = &v
	return s
}

func (s *CreateRestoreJobRequest) SetExclude(v string) *CreateRestoreJobRequest {
	s.Exclude = &v
	return s
}

func (s *CreateRestoreJobRequest) SetFailbackDetail(v map[string]interface{}) *CreateRestoreJobRequest {
	s.FailbackDetail = v
	return s
}

func (s *CreateRestoreJobRequest) SetInclude(v string) *CreateRestoreJobRequest {
	s.Include = &v
	return s
}

func (s *CreateRestoreJobRequest) SetInitiatedByAck(v bool) *CreateRestoreJobRequest {
	s.InitiatedByAck = &v
	return s
}

func (s *CreateRestoreJobRequest) SetOptions(v string) *CreateRestoreJobRequest {
	s.Options = &v
	return s
}

func (s *CreateRestoreJobRequest) SetOtsDetail(v *OtsTableRestoreDetail) *CreateRestoreJobRequest {
	s.OtsDetail = v
	return s
}

func (s *CreateRestoreJobRequest) SetRestoreType(v string) *CreateRestoreJobRequest {
	s.RestoreType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotHash(v string) *CreateRestoreJobRequest {
	s.SnapshotHash = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSnapshotId(v string) *CreateRestoreJobRequest {
	s.SnapshotId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetSourceType(v string) *CreateRestoreJobRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetBucket(v string) *CreateRestoreJobRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetContainer(v string) *CreateRestoreJobRequest {
	s.TargetContainer = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetContainerClusterId(v string) *CreateRestoreJobRequest {
	s.TargetContainerClusterId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetCreateTime(v int64) *CreateRestoreJobRequest {
	s.TargetCreateTime = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetFileSystemId(v string) *CreateRestoreJobRequest {
	s.TargetFileSystemId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetInstanceId(v string) *CreateRestoreJobRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetInstanceName(v string) *CreateRestoreJobRequest {
	s.TargetInstanceName = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetPath(v string) *CreateRestoreJobRequest {
	s.TargetPath = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetPrefix(v string) *CreateRestoreJobRequest {
	s.TargetPrefix = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetTableName(v string) *CreateRestoreJobRequest {
	s.TargetTableName = &v
	return s
}

func (s *CreateRestoreJobRequest) SetTargetTime(v int64) *CreateRestoreJobRequest {
	s.TargetTime = &v
	return s
}

func (s *CreateRestoreJobRequest) SetUdmDetail(v map[string]interface{}) *CreateRestoreJobRequest {
	s.UdmDetail = v
	return s
}

func (s *CreateRestoreJobRequest) SetUdmRegionId(v string) *CreateRestoreJobRequest {
	s.UdmRegionId = &v
	return s
}

func (s *CreateRestoreJobRequest) SetVaultId(v string) *CreateRestoreJobRequest {
	s.VaultId = &v
	return s
}

func (s *CreateRestoreJobRequest) Validate() error {
	if s.OtsDetail != nil {
		if err := s.OtsDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
