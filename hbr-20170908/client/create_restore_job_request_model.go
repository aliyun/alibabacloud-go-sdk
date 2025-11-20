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
	CrossAccountUserId *int64  `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	Edition            *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The path not to be restored. All documents under this path will not be restored. Maximum length is 255 characters.
	//
	// example:
	//
	// ["/var", "/proc"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// Details of restoring to the local environment.
	FailbackDetail map[string]interface{} `json:"FailbackDetail,omitempty" xml:"FailbackDetail,omitempty"`
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
	OtsDetail *OtsTableRestoreDetail `json:"OtsDetail,omitempty" xml:"OtsDetail,omitempty"`
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
	UdmDetail map[string]interface{} `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty"`
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
