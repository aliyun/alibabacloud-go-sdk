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
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1************dtv
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
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **ECS_FILE*	- or **File**. This parameter specifies the type of files that do not need to be backed up. No files of the specified type are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// This parameter is required only if you set the **SourceType*	- parameter to **ECS_FILE*	- or **File**. This parameter specifies the type of files to be backed up. All files of the specified type are backed up. The value can be up to 255 characters in length.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The description of the association.
	//
	// example:
	//
	// po-000************5xx-i-2ze************nw4
	PolicyBindingDescription *string `json:"PolicyBindingDescription,omitempty" xml:"PolicyBindingDescription,omitempty"`
	// The ID of the backup policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// po-000************ky9
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// 	- If the SourceType parameter is set to **OSS**, set the Source parameter to the prefix of the path to the folder that you want to back up. If you do not specify the Source parameter, the entire bucket (root directory) is backed up.
	//
	// 	- If the SourceType parameter is set to **ECS_FILE*	- or **File**, set the Source parameter to the path to the files that you want to back up. If you do not specify the Source parameter, all paths backed up.
	//
	// example:
	//
	// backup/
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- **UDM_ECS**: ECS instance backup
	//
	// This parameter is required.
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
	return dara.Validate(s)
}

type UpdatePolicyBindingRequestAdvancedOptions struct {
	// The details about large-scale file system backup.
	CommonFileSystemDetail *UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail `json:"CommonFileSystemDetail,omitempty" xml:"CommonFileSystemDetail,omitempty" type:"Struct"`
	// The details about Object Storage Service (OSS) backup.
	OssDetail *UpdatePolicyBindingRequestAdvancedOptionsOssDetail `json:"OssDetail,omitempty" xml:"OssDetail,omitempty" type:"Struct"`
	// The details about Elastic Compute Service (ECS) instance backup.
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
	return dara.Validate(s)
}

type UpdatePolicyBindingRequestAdvancedOptionsCommonFileSystemDetail struct {
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
	// 	- true: creates application-consistent snapshots
	//
	// 	- false: creates file system-consistent snapshots
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
