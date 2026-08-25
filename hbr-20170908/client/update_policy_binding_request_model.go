// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedOptions(v *UpdatePolicyBindingRequestAdvancedOptions) *UpdatePolicyBindingRequest
	GetAdvancedOptions() *UpdatePolicyBindingRequestAdvancedOptions
	SetDataSourceId(v string) *UpdatePolicyBindingRequest
	GetDataSourceId() *string
	SetDisabled(v bool) *UpdatePolicyBindingRequest
	GetDisabled() *bool
	SetExclude(v string) *UpdatePolicyBindingRequest
	GetExclude() *string
	SetInclude(v string) *UpdatePolicyBindingRequest
	GetInclude() *string
	SetPolicyBindingDescription(v string) *UpdatePolicyBindingRequest
	GetPolicyBindingDescription() *string
	SetPolicyId(v string) *UpdatePolicyBindingRequest
	GetPolicyId() *string
	SetSource(v string) *UpdatePolicyBindingRequest
	GetSource() *string
	SetSourceType(v string) *UpdatePolicyBindingRequest
	GetSourceType() *string
	SetSpeedLimit(v string) *UpdatePolicyBindingRequest
	GetSpeedLimit() *string
}

type UpdatePolicyBindingRequest struct {
	// The advanced options.
	AdvancedOptions *UpdatePolicyBindingRequestAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1************dtv
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Specifies whether the policy is suspended for the data source.
	//
	// - true: Suspended.
	//
	// - false: Not suspended.
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**, **File**, **NAS**, **COMMON_NAS**, or **COMMON_FILE_SYSTEM**. Specifies the file types to back up. All files of these types are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is valid only when **SourceType*	- is set to **ECS_FILE**, **File**, **NAS**, **COMMON_NAS**, or **COMMON_FILE_SYSTEM**. Specifies the file types to back up. All files of these types are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The description of the policy binding.
	//
	// example:
	//
	// po-000************5xx-i-2ze************nw4
	PolicyBindingDescription *string `json:"PolicyBindingDescription,omitempty" xml:"PolicyBindingDescription,omitempty"`
	// The policy ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// po-000************ky9
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The meaning varies depending on the SourceType value:
	//
	// - **OSS**: The prefix to back up. If not specified, the entire root directory of the bucket is backed up. Only a single prefix is supported. To back up /backup, set this parameter to /backup.
	//
	// - **ECS_FILE**: The file directories to back up. If not specified, all directories are backed up. Multiple directories are supported. To back up files in /a and /b, set this parameter to ["/a", "/b"].
	//
	// - **File**: The file directories to back up. If not specified, all directories are backed up. Multiple directories are supported. To back up files in /a and /b, set this parameter to ["/a", "/b"].
	//
	// - **COMMON_FILE_SYSTEM**: Required. The source paths to back up. Multiple paths are supported. To back up /a and /b, set this parameter to ["/a", "/b"]. To back up the root path, set this parameter to ["/"].
	//
	// - **COMMON_NAS**: Required. The source path to back up. Only a single path is supported. To back up /a, set this parameter to ["/a"]. To back up the root path, set this parameter to ["/"].
	//
	// - **OTS**: The list of data tables to back up. If not specified, all data tables are backed up. Multiple data tables are supported. To back up data tables a and b, set this parameter to ["a", "b"].
	//
	// example:
	//
	// backup/
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS instance backup.
	//
	// - **OSS**: OSS backup.
	//
	// - **NAS**: Alibaba Cloud NAS backup.
	//
	// - **COMMON_NAS**: On-premises NAS backup.
	//
	// - **ECS_FILE**: ECS File Backup Essential Edition.
	//
	// - **File**: On-premises file backup.
	//
	// - **COMMON_FILE_SYSTEM**: CPFS backup.
	//
	// - **OTS**: Tablestore backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE*	- or **File**. Specifies the backup traffic control. The format is `{start}{end}{bandwidth}`. Multiple traffic control configurations are separated by delimiters, and the time ranges cannot overlap.
	//
	// - **start**: The start hour.
	//
	// - **end**: The end hour.
	//
	// - **bandwidth**: The rate limit, in KB/s.
	//
	// example:
	//
	// 0:24:5120
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
}

func (s UpdatePolicyBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingRequest) GetAdvancedOptions() *UpdatePolicyBindingRequestAdvancedOptions {
	return s.AdvancedOptions
}

func (s *UpdatePolicyBindingRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdatePolicyBindingRequest) GetDisabled() *bool {
	return s.Disabled
}

func (s *UpdatePolicyBindingRequest) GetExclude() *string {
	return s.Exclude
}

func (s *UpdatePolicyBindingRequest) GetInclude() *string {
	return s.Include
}

func (s *UpdatePolicyBindingRequest) GetPolicyBindingDescription() *string {
	return s.PolicyBindingDescription
}

func (s *UpdatePolicyBindingRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdatePolicyBindingRequest) GetSource() *string {
	return s.Source
}

func (s *UpdatePolicyBindingRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdatePolicyBindingRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *UpdatePolicyBindingRequest) SetAdvancedOptions(v *UpdatePolicyBindingRequestAdvancedOptions) *UpdatePolicyBindingRequest {
	s.AdvancedOptions = v
	return s
}

func (s *UpdatePolicyBindingRequest) SetDataSourceId(v string) *UpdatePolicyBindingRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetDisabled(v bool) *UpdatePolicyBindingRequest {
	s.Disabled = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetExclude(v string) *UpdatePolicyBindingRequest {
	s.Exclude = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetInclude(v string) *UpdatePolicyBindingRequest {
	s.Include = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetPolicyBindingDescription(v string) *UpdatePolicyBindingRequest {
	s.PolicyBindingDescription = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetPolicyId(v string) *UpdatePolicyBindingRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetSource(v string) *UpdatePolicyBindingRequest {
	s.Source = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetSourceType(v string) *UpdatePolicyBindingRequest {
	s.SourceType = &v
	return s
}

func (s *UpdatePolicyBindingRequest) SetSpeedLimit(v string) *UpdatePolicyBindingRequest {
	s.SpeedLimit = &v
	return s
}

func (s *UpdatePolicyBindingRequest) Validate() error {
	if s.AdvancedOptions != nil {
		if err := s.AdvancedOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolicyBindingRequestAdvancedOptions struct {
	// The large-scale file system backup details.
	CommonFileSystemDetail *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail `json:"CommonFileSystemDetail,omitempty" xml:"CommonFileSystemDetail,omitempty" type:"Struct"`
	// The OSS backup details.
	OssDetail *UpdatePolicyBindingRequestAdvancedOptionsOssDetail `json:"OssDetail,omitempty" xml:"OssDetail,omitempty" type:"Struct"`
	// The ECS instance backup details.
	UdmDetail *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty" type:"Struct"`
}

func (s UpdatePolicyBindingRequestAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingRequestAdvancedOptions) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) GetCommonFileSystemDetail() *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail {
	return s.CommonFileSystemDetail
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) GetOssDetail() *UpdatePolicyBindingRequestAdvancedOptionsOssDetail {
	return s.OssDetail
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) GetUdmDetail() *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	return s.UdmDetail
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) SetCommonFileSystemDetail(v *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) *UpdatePolicyBindingRequestAdvancedOptions {
	s.CommonFileSystemDetail = v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) SetOssDetail(v *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) *UpdatePolicyBindingRequestAdvancedOptions {
	s.OssDetail = v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) SetUdmDetail(v *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) *UpdatePolicyBindingRequestAdvancedOptions {
	s.UdmDetail = v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptions) Validate() error {
	if s.CommonFileSystemDetail != nil {
		if err := s.CommonFileSystemDetail.Validate(); err != nil {
			return err
		}
	}
	if s.OssDetail != nil {
		if err := s.OssDetail.Validate(); err != nil {
			return err
		}
	}
	if s.UdmDetail != nil {
		if err := s.UdmDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail struct {
	// The sub-task slice size (number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Specifies whether to switch to a full backup when an incremental backup fails. Valid values:
	//
	// - **true**: Switches to a full backup upon failure.
	//
	// - **false**: Does not switch to a full backup upon failure.
	//
	// example:
	//
	// true
	FullOnIncrementFail *bool `json:"FullOnIncrementFail,omitempty" xml:"FullOnIncrementFail,omitempty"`
}

func (s UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) GetFetchSliceSize() *int64 {
	return s.FetchSliceSize
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) GetFullOnIncrementFail() *bool {
	return s.FullOnIncrementFail
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) SetFetchSliceSize(v int64) *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail {
	s.FetchSliceSize = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) SetFullOnIncrementFail(v bool) *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail {
	s.FullOnIncrementFail = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyBindingRequestAdvancedOptionsOssDetail struct {
	// Specifies whether to exclude archive objects from job statistics and failed file lists.
	//
	// example:
	//
	// true
	IgnoreArchiveObject *bool `json:"IgnoreArchiveObject,omitempty" xml:"IgnoreArchiveObject,omitempty"`
	// Specifies whether to delete inventory files after backup. This parameter is valid only when OSS inventory is used. Valid values:
	//
	// - **NO_CLEANUP**: Do not delete.
	//
	// - **DELETE_CURRENT**: Delete the current file.
	//
	// - **DELETE_CURRENT_AND_PREVIOUS**: Delete all files.
	//
	// example:
	//
	// NO_CLEANUP
	InventoryCleanupPolicy *string `json:"InventoryCleanupPolicy,omitempty" xml:"InventoryCleanupPolicy,omitempty"`
	// The OSS inventory name. When this value is not empty, the OSS inventory is used for performance optimization.
	//
	// - For more than 100 million OSS objects, use an inventory to improve incremental performance. Storage fees generated by inventory files are charged separately by OSS.
	//
	// - OSS inventory files take time to generate. Backup jobs may fail before the inventory files are generated. Wait for the next cycle to execute.
	//
	// example:
	//
	// 30663060
	InventoryId *string `json:"InventoryId,omitempty" xml:"InventoryId,omitempty"`
}

func (s UpdatePolicyBindingRequestAdvancedOptionsOssDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingRequestAdvancedOptionsOssDetail) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) GetIgnoreArchiveObject() *bool {
	return s.IgnoreArchiveObject
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) GetInventoryCleanupPolicy() *string {
	return s.InventoryCleanupPolicy
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) GetInventoryId() *string {
	return s.InventoryId
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) SetIgnoreArchiveObject(v bool) *UpdatePolicyBindingRequestAdvancedOptionsOssDetail {
	s.IgnoreArchiveObject = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) SetInventoryCleanupPolicy(v string) *UpdatePolicyBindingRequestAdvancedOptionsOssDetail {
	s.InventoryCleanupPolicy = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) SetInventoryId(v string) *UpdatePolicyBindingRequestAdvancedOptionsOssDetail {
	s.InventoryId = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsOssDetail) Validate() error {
	return dara.Validate(s)
}

type UpdatePolicyBindingRequestAdvancedOptionsUdmDetail struct {
	// Specifies whether to create an application-consistent snapshot. Application-consistent snapshots are supported only when all cloud disk types are ESSD.
	//
	// example:
	//
	// false
	AppConsistent *bool `json:"AppConsistent,omitempty" xml:"AppConsistent,omitempty"`
	// The list of cloud disk IDs that need to be protected. This value is empty when all cloud disks are protected.
	DiskIdList []*string `json:"DiskIdList,omitempty" xml:"DiskIdList,omitempty" type:"Repeated"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. Specifies whether to use the Linux FsFreeze mechanism to ensure the file system is in read consistency before creating an application-consistent snapshot. Default value: true.
	//
	// example:
	//
	// true
	EnableFsFreeze *bool `json:"EnableFsFreeze,omitempty" xml:"EnableFsFreeze,omitempty"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. Specifies whether to create an application-consistent snapshot:
	//
	// - true: Creates an application-consistent snapshot.
	//
	// - false: Creates a file system-consistent snapshot.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableWriters *bool `json:"EnableWriters,omitempty" xml:"EnableWriters,omitempty"`
	// The list of cloud disk IDs that do not need to be protected. This parameter is ignored when DiskIdList is not empty.
	ExcludeDiskIdList []*string `json:"ExcludeDiskIdList,omitempty" xml:"ExcludeDiskIdList,omitempty" type:"Repeated"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. The path of the post-thaw script to run after creating an application-consistent snapshot.
	//
	// example:
	//
	// /tmp/postscript.sh
	PostScriptPath *string `json:"PostScriptPath,omitempty" xml:"PostScriptPath,omitempty"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. The path of the pre-freeze script to run before creating an application-consistent snapshot.
	//
	// example:
	//
	// /tmp/prescript.sh
	PreScriptPath *string `json:"PreScriptPath,omitempty" xml:"PreScriptPath,omitempty"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. The RAM role name required for creating application-consistent snapshots.
	//
	// example:
	//
	// AliyunECSInstanceForHbrRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// Specifies whether to create a snapshot-consistent group. Snapshot-consistent groups are supported only when all cloud disk types are ESSD.
	//
	// example:
	//
	// true
	SnapshotGroup *bool `json:"SnapshotGroup,omitempty" xml:"SnapshotGroup,omitempty"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. The I/O freeze timeout period. Unit: seconds. Default value: 30.
	//
	// example:
	//
	// 30
	TimeoutInSeconds *int64 `json:"TimeoutInSeconds,omitempty" xml:"TimeoutInSeconds,omitempty"`
}

func (s UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetAppConsistent() *bool {
	return s.AppConsistent
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetDiskIdList() []*string {
	return s.DiskIdList
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetEnableFsFreeze() *bool {
	return s.EnableFsFreeze
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetEnableWriters() *bool {
	return s.EnableWriters
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetExcludeDiskIdList() []*string {
	return s.ExcludeDiskIdList
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetPostScriptPath() *string {
	return s.PostScriptPath
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetPreScriptPath() *string {
	return s.PreScriptPath
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetSnapshotGroup() *bool {
	return s.SnapshotGroup
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) GetTimeoutInSeconds() *int64 {
	return s.TimeoutInSeconds
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetAppConsistent(v bool) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.AppConsistent = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetDiskIdList(v []*string) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.DiskIdList = v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetEnableFsFreeze(v bool) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.EnableFsFreeze = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetEnableWriters(v bool) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.EnableWriters = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetExcludeDiskIdList(v []*string) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.ExcludeDiskIdList = v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetPostScriptPath(v string) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.PostScriptPath = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetPreScriptPath(v string) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.PreScriptPath = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetRamRoleName(v string) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.RamRoleName = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetSnapshotGroup(v bool) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.SnapshotGroup = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) SetTimeoutInSeconds(v int64) *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail {
	s.TimeoutInSeconds = &v
	return s
}

func (s *UpdatePolicyBindingRequestAdvancedOptionsUdmDetail) Validate() error {
	return dara.Validate(s)
}
