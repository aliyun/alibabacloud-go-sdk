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
	// The data sources that you want to bind to the backup policy.
	PolicyBindingList []*CreatePolicyBindingsRequestPolicyBindingList `json:"PolicyBindingList,omitempty" xml:"PolicyBindingList,omitempty" type:"Repeated"`
	// The ID of the backup policy.
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
	// The name of the RAM role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether to back up and restore data within the same Alibaba Cloud account or across Alibaba Cloud accounts. Default value: SELF_ACCOUNT. Valid values:
	//
	// 	- **SELF_ACCOUNT**: backs up data within the same Alibaba Cloud account.
	//
	// 	- **CROSS_ACCOUNT**: backs up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// SELF_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 144**********732
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The ID of the data source. The meaning of this parameter depends on the **SourceType*	- parameter. Valid values:
	//
	// 	- **UDM_ECS**: the ID of the Elastic Compute Service (ECS) instance
	//
	// 	- **OSS**: the name of the Object Storage Service (OSS) bucket
	//
	// 	- **NAS**: the ID of the File Storage NAS (NAS) file system
	//
	// 	- **COMMON_NAS**: the ID of the on-premises NAS file system
	//
	// 	- **ECS_FILE**: the ID of the ECS instance
	//
	// 	- **File**: the ID of the Cloud Backup client
	//
	// 	- **COMMON_FILE_SYSTEM**: the ID of the Cloud Parallel File Storage (CPFS) backup data source
	//
	// example:
	//
	// i-bp1************dl8
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Specifies whether to disable the backup policy for the data source. Valid values:
	//
	// 	- true: disables the backup policy for the data source
	//
	// 	- false: enables the backup policy for the data source
	//
	// example:
	//
	// true
	Disabled *string `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **ECS_FILE**, **File**, **NAS**, **COMMON_NAS**, or **COMMON_FILE_SYSTEM**. This parameter specifies the type of files that do not need to be backed up. No files of the specified type are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **ECS_FILE**, **File**, **NAS**, **COMMON_NAS**, or **COMMON_FILE_SYSTEM**. This parameter specifies the type of files to be backed up. All files of the specified type are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The description of the association.
	//
	// example:
	//
	// Bind data sources to a backup policy
	PolicyBindingDescription *string `json:"PolicyBindingDescription,omitempty" xml:"PolicyBindingDescription,omitempty"`
	// 	- If the SourceType parameter is set to **OSS**, set the Source parameter to the prefix of the path to the folder that you want to back up. If you do not specify the Source parameter, the entire bucket (root directory) is backed up.
	//
	// 	- If the SourceType parameter is set to **ECS_FILE*	- or **File**, set the Source parameter to the path to the files that you want to back up. If you do not specify the Source parameter, all paths are backed up.
	//
	// 	- This parameter is required if the SourceType parameter is set to **COMMON_FILE_SYSTEM**. This parameter specifies the path to be backed up. To back up the /src path, enter ["/src"]. To back up the root path, enter ["/"].
	//
	// example:
	//
	// backup/
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: ECS instance
	//
	// 	- **OSS**: OSS bucket
	//
	// 	- **NAS**: NAS file system
	//
	// 	- **COMMON_NAS**: on-premises NAS file system
	//
	// 	- **ECS_FILE**: ECS file
	//
	// 	- **File**: on-premises file
	//
	// 	- **COMMON_FILE_SYSTEM**: CPFS file system
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **ECS_FILE*	- or **File**. This parameter specifies the throttling rules. Format: `{start}{end}{bandwidth}`. Separate multiple throttling rules with vertical bars (|). The time ranges of the throttling rules cannot overlap.
	//
	// 	- **start**: the start hour.
	//
	// 	- **end**: the end hour.
	//
	// 	- **bandwidth**: the bandwidth. Unit: KB/s.
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
	// The advanced options for Object Storage Service (OSS) backup.
	OssDetail *CreatePolicyBindingsRequestPolicyBindingListAdvancedOptionsOssDetail `json:"OssDetail,omitempty" xml:"OssDetail,omitempty" type:"Struct"`
	// The advanced options for ECS instance backup.
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
	// The size of backup shards (the number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Specifies whether the system performs full backup if incremental backup fails. Valid values:
	//
	// 	- **true**: The system performs full backup if incremental backup fails.
	//
	// 	- **false**: The system does not perform full backup if incremental backup fails.
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
	// The ID of the backup client group. When you perform on-premises NAS backup, Cloud Backup selects clients from the specified backup client group.
	//
	// example:
	//
	// cl-000**************ggu
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The size of backup shards (the number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Specifies whether the system performs full backup if incremental backup fails. Valid values:
	//
	// 	- **true**: The system performs full backup if incremental backup fails.
	//
	// 	- **false**: The system does not perform full backup if incremental backup fails.
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
	// 	- **true**: uses the advanced policy.
	//
	// 	- **false**: does not use the advanced policy.
	//
	// example:
	//
	// true
	AdvPolicy *bool `json:"AdvPolicy,omitempty" xml:"AdvPolicy,omitempty"`
	// Specifies whether to enable the Volume Shadow Copy Service (VSS) feature. Valid values:
	//
	// 	- **true**: enables the feature.
	//
	// 	- **false**: disables the feature.
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
	// Do not prompt for archival type objects in task statistics and failed file lists.
	//
	// example:
	//
	// true
	IgnoreArchiveObject *bool `json:"IgnoreArchiveObject,omitempty" xml:"IgnoreArchiveObject,omitempty"`
	// Specifies whether the system deletes the inventory lists when a backup is completed. This parameter is valid only when OSS inventories are used. Valid values:
	//
	// 	- **NO_CLEANUP**: does not delete inventory lists.
	//
	// 	- **DELETE_CURRENT**: deletes the current inventory list.
	//
	// 	- **DELETE_CURRENT_AND_PREVIOUS**: deletes all inventory lists.
	//
	// example:
	//
	// NO_CLEANUP
	InventoryCleanupPolicy *string `json:"InventoryCleanupPolicy,omitempty" xml:"InventoryCleanupPolicy,omitempty"`
	// The name of the OSS inventory. If this parameter is not empty, the OSS inventory is used for performance optimization.
	//
	// 	- If you want to back up more than 100 million OSS objects, we recommend that you use inventory lists to accelerate incremental backup. Storage fees for inventory lists are included into your OSS bills.
	//
	// 	- A certain amount of time is required for OSS to generate inventory lists. Before inventory lists are generated, OSS objects may fail to be backed up. In this case, you can back up the OSS objects in the next backup cycle.
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
	// Specifies whether to enable application consistency. You can enable application consistency only if all disks are ESSDs.
	//
	// example:
	//
	// false
	AppConsistent *bool `json:"AppConsistent,omitempty" xml:"AppConsistent,omitempty"`
	// The IDs of the disks that need to be protected. If all disks need to be protected, this parameter is empty.
	DiskIdList []*string `json:"DiskIdList,omitempty" xml:"DiskIdList,omitempty" type:"Repeated"`
	// This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies whether to enable Linux fsfreeze to put file systems into the read-only state before application-consistent snapshots are created. Default value: true.
	//
	// example:
	//
	// true
	EnableFsFreeze *bool `json:"EnableFsFreeze,omitempty" xml:"EnableFsFreeze,omitempty"`
	// This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies whether to create application-consistent snapshots. Valid values:
	//
	// 	- true: creates application-consistent snapshots.
	//
	// 	- false: creates file system-consistent snapshots.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	EnableWriters *bool `json:"EnableWriters,omitempty" xml:"EnableWriters,omitempty"`
	// The IDs of the disks that do not need to be protected. If the DiskIdList parameter is not empty, this parameter is ignored.
	ExcludeDiskIdList []*string `json:"ExcludeDiskIdList,omitempty" xml:"ExcludeDiskIdList,omitempty" type:"Repeated"`
	// This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies the path of the post-thaw scripts that are executed after application-consistent snapshots are created.
	//
	// example:
	//
	// /tmp/postscript.sh
	PostScriptPath *string `json:"PostScriptPath,omitempty" xml:"PostScriptPath,omitempty"`
	// This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies the path of the pre-freeze scripts that are executed before application-consistent snapshots are created.
	//
	// example:
	//
	// /tmp/prescript.sh
	PreScriptPath *string `json:"PreScriptPath,omitempty" xml:"PreScriptPath,omitempty"`
	// This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies the name of the Resource Access Management (RAM) role that is required to create application-consistent snapshots.
	//
	// example:
	//
	// AliyunECSInstanceForHbrRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// Specifies whether to create a snapshot-consistent group. You can create a snapshot-consistent group only if all disks are Enterprise SSDs (ESSDs).
	//
	// example:
	//
	// true
	SnapshotGroup *bool `json:"SnapshotGroup,omitempty" xml:"SnapshotGroup,omitempty"`
	// This parameter is required only if you set the **AppConsistent*	- parameter to **true**. This parameter specifies the I/O freeze timeout period. Default value: 30. Unit: seconds.
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
