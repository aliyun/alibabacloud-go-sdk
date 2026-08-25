// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyBindingList(v []*CreatePolicyBindingsRequestPolicyBindingList) *CreatePolicyBindingsRequest
	GetPolicyBindingList() []*CreatePolicyBindingsRequestPolicyBindingList
	SetPolicyId(v string) *CreatePolicyBindingsRequest
	GetPolicyId() *string
}

type CreatePolicyBindingsRequest struct {
	// The list of policy bindings.
	PolicyBindingList []*CreatePolicyBindingsRequestPolicyBindingList `json:"PolicyBindingList,omitempty" xml:"PolicyBindingList,omitempty" type:"Repeated"`
	// The policy ID.
	//
	// example:
	//
	// po-000************8ep
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s CreatePolicyBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequest) GetPolicyBindingList() []*CreatePolicyBindingsRequestPolicyBindingList {
	return s.PolicyBindingList
}

func (s *CreatePolicyBindingsRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *CreatePolicyBindingsRequest) SetPolicyBindingList(v []*CreatePolicyBindingsRequestPolicyBindingList) *CreatePolicyBindingsRequest {
	s.PolicyBindingList = v
	return s
}

func (s *CreatePolicyBindingsRequest) SetPolicyId(v string) *CreatePolicyBindingsRequest {
	s.PolicyId = &v
	return s
}

func (s *CreatePolicyBindingsRequest) Validate() error {
	if s.PolicyBindingList != nil {
		for _, item := range s.PolicyBindingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePolicyBindingsRequestPolicyBindingList struct {
	// The advanced options.
	AdvancedOptions *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	// The RAM role name created in the source account for cross-account backup.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The cross-account backup type. Default value: SELF_ACCOUNT. Valid values:
	//
	// - **SELF_ACCOUNT**: Backup within the same account.
	//
	// - **CROSS_ACCOUNT**: Cross-account backup.
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The Alibaba Cloud UID of the source account for cross-account backup.
	//
	// example:
	//
	// 144**********732
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The data source ID. The value has different meanings depending on the **SourceType*	- field:
	//
	// - **UDM_ECS**: The ECS instance ID.
	//
	// - **OSS**: The OSS bucket name.
	//
	// - **NAS**: The Alibaba Cloud NAS file system ID.
	//
	// - **COMMON_NAS**: The on-premises NAS instance ID.
	//
	// - **ECS_FILE**: The ECS instance ID.
	//
	// - **File**: The Cloud Backup client ID.
	//
	// - **COMMON_FILE_SYSTEM**: The CPFS backup data source ID.
	//
	// example:
	//
	// i-bp1************dl8
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Specifies whether the policy is suspended for this data source.
	//
	// - true: Suspended.
	//
	// - false: Not suspended.
	//
	// example:
	//
	// true
	Disabled *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter can be configured when **SourceType*	- is set to **ECS_FILE**, **File**, **NAS**, **COMMON_NAS**, or **COMMON_FILE_SYSTEM**. Specifies the file types to exclude from the backup. All files of these types are not backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter can be configured when **SourceType*	- is set to **ECS_FILE**, **File**, **NAS**, **COMMON_NAS**, or **COMMON_FILE_SYSTEM**. Specifies the file types to include in the backup. All files of these types are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The description of the policy binding.
	//
	// example:
	//
	// This is a description of the policy binding
	PolicyBindingDescription *string `json:"PolicyBindingDescription,omitempty" xml:"PolicyBindingDescription,omitempty"`
	// The value has different meanings depending on the SourceType value:
	//
	// - **OSS**: The prefix to back up. If not specified, the entire bucket root directory is backed up. Only a single prefix is supported. To back up /backup, specify /backup.
	//
	// - **ECS_FILE**: The file directories to back up. If not specified, all directories are backed up. Multiple directories are supported. To back up files under /a and /b, specify ["/a", "/b"].
	//
	// - **File**: The file directories to back up. If not specified, all directories are backed up. Multiple directories are supported. To back up files under /a and /b, specify ["/a", "/b"].
	//
	// - **COMMON_FILE_SYSTEM**: Required. The source paths to back up. Multiple paths are supported. To back up /a and /b, specify ["/a", "/b"]. To back up the root path, specify ["/"].
	//
	// - **COMMON_NAS**: Required. The source path to back up. Only a single path is supported. To back up /a, specify ["/a"]. To back up the root path, specify ["/"].
	//
	// - **OTS**: The list of data tables to back up. If not specified, all data tables are backed up. Multiple data tables are supported. To back up tables a and b, specify ["a", "b"].
	//
	// example:
	//
	// backup/
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The data source type. Valid values:
	//
	// - **UDM_ECS**: ECS full server backup.
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
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE*	- or **File**. Specifies the backup traffic control. The format is `{start}{end}{bandwidth}`. Multiple traffic control configurations are separated by delimiters, and the time ranges must not overlap.
	//
	// - **start**: The start hour.
	//
	// - **end**: The end hour.
	//
	// - **bandwidth**: The rate limit, in KB/s.
	//
	// example:
	//
	// 0:24:1024
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
}

func (s CreatePolicyBindingsRequestPolicyBindingList) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingList) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetAdvancedOptions() *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions {
	return s.AdvancedOptions
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetDisabled() *string {
	return s.Disabled
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetExclude() *string {
	return s.Exclude
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetInclude() *string {
	return s.Include
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetPolicyBindingDescription() *string {
	return s.PolicyBindingDescription
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetSource() *string {
	return s.Source
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetSourceType() *string {
	return s.SourceType
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetAdvancedOptions(v *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) *CreatePolicyBindingsRequestPolicyBindingList {
	s.AdvancedOptions = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetCrossAccountRoleName(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.CrossAccountRoleName = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetCrossAccountType(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.CrossAccountType = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetCrossAccountUserId(v int64) *CreatePolicyBindingsRequestPolicyBindingList {
	s.CrossAccountUserId = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetDataSourceId(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.DataSourceId = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetDisabled(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.Disabled = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetExclude(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.Exclude = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetInclude(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.Include = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetPolicyBindingDescription(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.PolicyBindingDescription = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetSource(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.Source = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetSourceType(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.SourceType = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) SetSpeedLimit(v string) *CreatePolicyBindingsRequestPolicyBindingList {
	s.SpeedLimit = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingList) Validate() error {
	if s.AdvancedOptions != nil {
		if err := s.AdvancedOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions struct {
	// The advanced options for CPFS backup.
	CommonFileSystemDetail *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail `json:"CommonFileSystemDetail,omitempty" xml:"CommonFileSystemDetail,omitempty" type:"Struct"`
	// The advanced options for on-premises NAS backup.
	CommonNasDetail *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail `json:"CommonNasDetail,omitempty" xml:"CommonNasDetail,omitempty" type:"Struct"`
	// The advanced options for file backup.
	FileDetail *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail `json:"FileDetail,omitempty" xml:"FileDetail,omitempty" type:"Struct"`
	// The advanced options for OSS backup.
	OssDetail *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail `json:"OssDetail,omitempty" xml:"OssDetail,omitempty" type:"Struct"`
	// The advanced options for ECS full server backup.
	UdmDetail *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty" type:"Struct"`
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) GetCommonFileSystemDetail() *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail {
	return s.CommonFileSystemDetail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) GetCommonNasDetail() *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail {
	return s.CommonNasDetail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) GetFileDetail() *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail {
	return s.FileDetail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) GetOssDetail() *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail {
	return s.OssDetail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) GetUdmDetail() *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	return s.UdmDetail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) SetCommonFileSystemDetail(v *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions {
	s.CommonFileSystemDetail = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) SetCommonNasDetail(v *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions {
	s.CommonNasDetail = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) SetFileDetail(v *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions {
	s.FileDetail = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) SetOssDetail(v *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions {
	s.OssDetail = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) SetUdmDetail(v *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions {
	s.UdmDetail = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptions) Validate() error {
	if s.CommonFileSystemDetail != nil {
		if err := s.CommonFileSystemDetail.Validate(); err != nil {
			return err
		}
	}
	if s.CommonNasDetail != nil {
		if err := s.CommonNasDetail.Validate(); err != nil {
			return err
		}
	}
	if s.FileDetail != nil {
		if err := s.FileDetail.Validate(); err != nil {
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

type CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail struct {
	// The sub-task slice size (number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Specifies whether to switch to a full backup when an incremental backup fails. Valid values:
	//
	// - **true**: Switch to a full backup on failure.
	//
	// - **false**: Do not switch to a full backup on failure.
	//
	// example:
	//
	// true
	FullOnIncrementFail *bool `json:"FullOnIncrementFail,omitempty" xml:"FullOnIncrementFail,omitempty"`
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) GetFetchSliceSize() *int64 {
	return s.FetchSliceSize
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) GetFullOnIncrementFail() *bool {
	return s.FullOnIncrementFail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) SetFetchSliceSize(v int64) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail {
	s.FetchSliceSize = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) SetFullOnIncrementFail(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail {
	s.FullOnIncrementFail = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonFileSystemDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail struct {
	// The backup client group ID. On-premises NAS backup selects a client from the backup client group to perform the backup.
	//
	// example:
	//
	// cl-000**************ggu
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The sub-task slice size (number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Specifies whether to switch to a full backup when an incremental backup fails. Valid values:
	//
	// - **true**: Switch to a full backup on failure.
	//
	// - **false**: Do not switch to a full backup on failure.
	//
	// example:
	//
	// true
	FullOnIncrementFail *bool `json:"FullOnIncrementFail,omitempty" xml:"FullOnIncrementFail,omitempty"`
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) GetFetchSliceSize() *int64 {
	return s.FetchSliceSize
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) GetFullOnIncrementFail() *bool {
	return s.FullOnIncrementFail
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) SetClusterId(v string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail {
	s.ClusterId = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) SetFetchSliceSize(v int64) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail {
	s.FetchSliceSize = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) SetFullOnIncrementFail(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail {
	s.FullOnIncrementFail = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsCommonNasDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail struct {
	// Specifies whether to use an advanced policy. Valid values:
	//
	// - **true**: Use.
	//
	// - **false**: Do not use.
	//
	// example:
	//
	// true
	AdvPolicy *bool `json:"AdvPolicy,omitempty" xml:"AdvPolicy,omitempty"`
	// Specifies whether to enable the Volume Shadow Copy Service (VSS) feature (Windows). Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// true
	UseVSS *bool `json:"UseVSS,omitempty" xml:"UseVSS,omitempty"`
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) GetAdvPolicy() *bool {
	return s.AdvPolicy
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) GetUseVSS() *bool {
	return s.UseVSS
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) SetAdvPolicy(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail {
	s.AdvPolicy = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) SetUseVSS(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail {
	s.UseVSS = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsFileDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail struct {
	// Specifies whether to exclude archive objects from job statistics and failed file lists.
	//
	// example:
	//
	// true
	IgnoreArchiveObject *bool `json:"IgnoreArchiveObject,omitempty" xml:"IgnoreArchiveObject,omitempty"`
	// Specifies whether to delete inventory files after backup. This parameter takes effect only when OSS inventory is used. Valid values:
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
	// - Using an inventory is recommended for backing up more than 100 million OSS objects to improve incremental performance. Storage fees generated by inventory files are charged separately by OSS.
	//
	// - OSS inventory files take time to generate. Backup jobs may fail before the inventory files are generated. Wait for the next backup cycle.
	//
	// example:
	//
	// oss-inventory-default
	InventoryId *string `json:"InventoryId,omitempty" xml:"InventoryId,omitempty"`
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) GetIgnoreArchiveObject() *bool {
	return s.IgnoreArchiveObject
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) GetInventoryCleanupPolicy() *string {
	return s.InventoryCleanupPolicy
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) GetInventoryId() *string {
	return s.InventoryId
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) SetIgnoreArchiveObject(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail {
	s.IgnoreArchiveObject = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) SetInventoryCleanupPolicy(v string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail {
	s.InventoryCleanupPolicy = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) SetInventoryId(v string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail {
	s.InventoryId = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail) Validate() error {
	return dara.Validate(s)
}

type CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail struct {
	// Specifies whether to create an application-consistent snapshot. Application-consistent snapshots are supported only when all cloud disk types are ESSD.
	//
	// example:
	//
	// false
	AppConsistent *bool `json:"AppConsistent,omitempty" xml:"AppConsistent,omitempty"`
	// The list of cloud disk IDs that need to be protected. Leave this value empty to protect all cloud disks.
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
	// This parameter is required only when **AppConsistent*	- is set to **true**. The I/O freeze timeout period, in seconds. Default value: 30.
	//
	// example:
	//
	// 30
	TimeoutInSeconds *int64 `json:"TimeoutInSeconds,omitempty" xml:"TimeoutInSeconds,omitempty"`
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GoString() string {
	return s.String()
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetAppConsistent() *bool {
	return s.AppConsistent
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetDiskIdList() []*string {
	return s.DiskIdList
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetEnableFsFreeze() *bool {
	return s.EnableFsFreeze
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetEnableWriters() *bool {
	return s.EnableWriters
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetExcludeDiskIdList() []*string {
	return s.ExcludeDiskIdList
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetPostScriptPath() *string {
	return s.PostScriptPath
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetPreScriptPath() *string {
	return s.PreScriptPath
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetSnapshotGroup() *bool {
	return s.SnapshotGroup
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) GetTimeoutInSeconds() *int64 {
	return s.TimeoutInSeconds
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetAppConsistent(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.AppConsistent = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetDiskIdList(v []*string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.DiskIdList = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetEnableFsFreeze(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.EnableFsFreeze = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetEnableWriters(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.EnableWriters = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetExcludeDiskIdList(v []*string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.ExcludeDiskIdList = v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetPostScriptPath(v string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.PostScriptPath = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetPreScriptPath(v string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.PreScriptPath = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetRamRoleName(v string) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.RamRoleName = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetSnapshotGroup(v bool) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.SnapshotGroup = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) SetTimeoutInSeconds(v int64) *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail {
	s.TimeoutInSeconds = &v
	return s
}

func (s *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsUdmDetail) Validate() error {
	return dara.Validate(s)
}
