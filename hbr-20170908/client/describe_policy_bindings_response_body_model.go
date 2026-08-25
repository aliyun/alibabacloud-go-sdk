// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyBindingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePolicyBindingsResponseBody
	GetCode() *string
	SetMaxResults(v int32) *DescribePolicyBindingsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribePolicyBindingsResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribePolicyBindingsResponseBody
	GetNextToken() *string
	SetPolicyBindings(v []*DescribePolicyBindingsResponseBodyPolicyBindings) *DescribePolicyBindingsResponseBody
	GetPolicyBindings() []*DescribePolicyBindingsResponseBodyPolicyBindings
	SetRequestId(v string) *DescribePolicyBindingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribePolicyBindingsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *DescribePolicyBindingsResponseBody
	GetTotalCount() *int64
}

type DescribePolicyBindingsResponseBody struct {
	// The response code. 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of results for each query.
	//
	// Valid values: 10 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned message. "successful" is returned for success. An error message is returned for failure.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token required to obtain the next page of policy-data source bindings.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of policy bindings.
	PolicyBindings []*DescribePolicyBindingsResponseBodyPolicyBindings `json:"PolicyBindings,omitempty" xml:"PolicyBindings,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5225929A-4EBD-55EE-9FE1-4A130E582A76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 38
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePolicyBindingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePolicyBindingsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePolicyBindingsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePolicyBindingsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePolicyBindingsResponseBody) GetPolicyBindings() []*DescribePolicyBindingsResponseBodyPolicyBindings {
	return s.PolicyBindings
}

func (s *DescribePolicyBindingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolicyBindingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribePolicyBindingsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePolicyBindingsResponseBody) SetCode(v string) *DescribePolicyBindingsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetMaxResults(v int32) *DescribePolicyBindingsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetMessage(v string) *DescribePolicyBindingsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetNextToken(v string) *DescribePolicyBindingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetPolicyBindings(v []*DescribePolicyBindingsResponseBodyPolicyBindings) *DescribePolicyBindingsResponseBody {
	s.PolicyBindings = v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetRequestId(v string) *DescribePolicyBindingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetSuccess(v bool) *DescribePolicyBindingsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) SetTotalCount(v int64) *DescribePolicyBindingsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePolicyBindingsResponseBody) Validate() error {
	if s.PolicyBindings != nil {
		for _, item := range s.PolicyBindings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyBindingsResponseBodyPolicyBindings struct {
	// The advanced options.
	AdvancedOptions *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	// Indicates whether the resource is automatically associated through a backup policy resource tag.
	//
	// example:
	//
	// false
	CreatedByTag *bool `json:"CreatedByTag,omitempty" xml:"CreatedByTag,omitempty"`
	// The creation time. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1661399570
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The RAM role name created in the source account for cross-account backup.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The cross-account backup type. Valid values:
	//
	// - SELF_ACCOUNT: backup within the current account.
	//
	// - CROSS_ACCOUNT: cross-account backup.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source account for cross-account backup.
	//
	// example:
	//
	// 1480************
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// i-8vb************5ly
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Indicates whether the policy is effective for the data source.
	//
	// - true: paused.
	//
	// - false: not paused.
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE*	- or **File**. Specifies the file types to exclude from backup. All files of these types are not backed up. Maximum of 255 characters.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// The matched tag rules.
	HitTags []*DescribePolicyBindingsResponseBodyPolicyBindingsHitTags `json:"HitTags,omitempty" xml:"HitTags,omitempty" type:"Repeated"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE*	- or **File**. Specifies the file types to back up. All files of these types are backed up. Maximum of 255 characters.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// The description of the policy binding.
	//
	// example:
	//
	// po-000************eslc-i-uf6************y5g
	PolicyBindingDescription *string `json:"PolicyBindingDescription,omitempty" xml:"PolicyBindingDescription,omitempty"`
	// The policy binding ID.
	//
	// example:
	//
	// pd-000************slc
	PolicyBindingId *string `json:"PolicyBindingId,omitempty" xml:"PolicyBindingId,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// po-000************56y
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// - If SourceType is set to **OSS**, this parameter specifies the prefix to back up. If not specified, the entire Bucket root directory is backed up.
	//
	// - If SourceType is set to **ECS_FILE*	- or **File**, this parameter specifies the file directory to back up. If not specified, all directories are backed up.
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
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only when **SourceType*	- is set to **ECS_FILE*	- or **File**. Specifies the backup traffic control. Format: `{start}{end}{bandwidth}`. Multiple traffic control configurations are separated by delimiters, and configuration times cannot overlap.
	//
	// - **start**: start hour.
	//
	// - **end**: end hour.
	//
	// - **bandwidth**: rate limit, in KB/s.
	//
	// example:
	//
	// 0:24:10240
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// The update time. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1653611573
	UpdatedTime *int64 `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s DescribePolicyBindingsResponseBodyPolicyBindings) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindings) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetAdvancedOptions() *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions {
	return s.AdvancedOptions
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetCreatedByTag() *bool {
	return s.CreatedByTag
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetExclude() *string {
	return s.Exclude
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetHitTags() []*DescribePolicyBindingsResponseBodyPolicyBindingsHitTags {
	return s.HitTags
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetInclude() *string {
	return s.Include
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetPolicyBindingDescription() *string {
	return s.PolicyBindingDescription
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetPolicyBindingId() *string {
	return s.PolicyBindingId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetSource() *string {
	return s.Source
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetAdvancedOptions(v *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.AdvancedOptions = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetCreatedByTag(v bool) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.CreatedByTag = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetCreatedTime(v int64) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.CreatedTime = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetCrossAccountRoleName(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetCrossAccountType(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.CrossAccountType = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetCrossAccountUserId(v int64) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetDataSourceId(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.DataSourceId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetDisabled(v bool) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.Disabled = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetExclude(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.Exclude = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetHitTags(v []*DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.HitTags = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetInclude(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.Include = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetPolicyBindingDescription(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.PolicyBindingDescription = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetPolicyBindingId(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.PolicyBindingId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetPolicyId(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.PolicyId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetSource(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.Source = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetSourceType(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.SourceType = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetSpeedLimit(v string) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.SpeedLimit = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) SetUpdatedTime(v int64) *DescribePolicyBindingsResponseBodyPolicyBindings {
	s.UpdatedTime = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindings) Validate() error {
	if s.AdvancedOptions != nil {
		if err := s.AdvancedOptions.Validate(); err != nil {
			return err
		}
	}
	if s.HitTags != nil {
		for _, item := range s.HitTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions struct {
	// The advanced options for large-scale file system backup.
	CommonFileSystemDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail `json:"CommonFileSystemDetail,omitempty" xml:"CommonFileSystemDetail,omitempty" type:"Struct"`
	// The advanced options for on-premises NAS.
	CommonNasDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail `json:"CommonNasDetail,omitempty" xml:"CommonNasDetail,omitempty" type:"Struct"`
	// The advanced options for file backup.
	FileDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail `json:"FileDetail,omitempty" xml:"FileDetail,omitempty" type:"Struct"`
	// The advanced options for OSS backup.
	OssDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail `json:"OssDetail,omitempty" xml:"OssDetail,omitempty" type:"Struct"`
	// The advanced options for ECS instance backup.
	UdmDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail `json:"UdmDetail,omitempty" xml:"UdmDetail,omitempty" type:"Struct"`
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) GetCommonFileSystemDetail() *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail {
	return s.CommonFileSystemDetail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) GetCommonNasDetail() *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail {
	return s.CommonNasDetail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) GetFileDetail() *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail {
	return s.FileDetail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) GetOssDetail() *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail {
	return s.OssDetail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) GetUdmDetail() *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	return s.UdmDetail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) SetCommonFileSystemDetail(v *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions {
	s.CommonFileSystemDetail = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) SetCommonNasDetail(v *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions {
	s.CommonNasDetail = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) SetFileDetail(v *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions {
	s.FileDetail = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) SetOssDetail(v *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions {
	s.OssDetail = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) SetUdmDetail(v *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions {
	s.UdmDetail = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions) Validate() error {
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

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail struct {
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

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) GetFetchSliceSize() *int64 {
	return s.FetchSliceSize
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) GetFullOnIncrementFail() *bool {
	return s.FullOnIncrementFail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) SetFetchSliceSize(v int64) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail {
	s.FetchSliceSize = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) SetFullOnIncrementFail(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail {
	s.FullOnIncrementFail = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail struct {
	// The backup client ID.
	//
	// example:
	//
	// c-0001eg6mcvjs93f46s2d
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client group ID.
	//
	// example:
	//
	// cl-000gkcofngi04j6k680a
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) GetClientId() *string {
	return s.ClientId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) GetFetchSliceSize() *int64 {
	return s.FetchSliceSize
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) GetFullOnIncrementFail() *bool {
	return s.FullOnIncrementFail
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) SetClientId(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail {
	s.ClientId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) SetClusterId(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail {
	s.ClusterId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) SetFetchSliceSize(v int64) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail {
	s.FetchSliceSize = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) SetFullOnIncrementFail(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail {
	s.FullOnIncrementFail = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail struct {
	// Specifies whether to use an advanced policy. Valid values:
	//
	// - **true**: Used.
	//
	// - **false**: Not used.
	//
	// example:
	//
	// false
	AdvPolicy *bool `json:"AdvPolicy,omitempty" xml:"AdvPolicy,omitempty"`
	// Specifies whether to enable the Volume Shadow Copy Service (VSS) feature (Windows). Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Disabled.
	//
	// example:
	//
	// false
	UseVSS *bool `json:"UseVSS,omitempty" xml:"UseVSS,omitempty"`
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) GetAdvPolicy() *bool {
	return s.AdvPolicy
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) GetUseVSS() *bool {
	return s.UseVSS
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) SetAdvPolicy(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail {
	s.AdvPolicy = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) SetUseVSS(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail {
	s.UseVSS = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail struct {
	// Specifies whether to exclude archive objects from task statistics and failed file lists.
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
	// DELETE_CURRENT_AND_PREVIOUS
	InventoryCleanupPolicy *string `json:"InventoryCleanupPolicy,omitempty" xml:"InventoryCleanupPolicy,omitempty"`
	// The OSS inventory name. If this value is not empty, the OSS inventory is used for performance tuning.
	//
	// - Using an inventory to improve incremental performance is recommended when backing up more than 100 million OSS objects. Storage fees generated by inventory files are charged separately by OSS.
	//
	// - OSS inventory files take time to generate. Backup may fail before the OSS inventory file is generated. Wait for the next cycle to execute.
	//
	// example:
	//
	// inventory_test
	InventoryId *string `json:"InventoryId,omitempty" xml:"InventoryId,omitempty"`
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) GetIgnoreArchiveObject() *bool {
	return s.IgnoreArchiveObject
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) GetInventoryCleanupPolicy() *string {
	return s.InventoryCleanupPolicy
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) GetInventoryId() *string {
	return s.InventoryId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) SetIgnoreArchiveObject(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail {
	s.IgnoreArchiveObject = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) SetInventoryCleanupPolicy(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail {
	s.InventoryCleanupPolicy = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) SetInventoryId(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail {
	s.InventoryId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail struct {
	// Specifies whether to create an application-consistent snapshot. Creating an application-consistent snapshot is supported only when all cloud disk types are ESSD.
	//
	// example:
	//
	// false
	AppConsistent *bool `json:"AppConsistent,omitempty" xml:"AppConsistent,omitempty"`
	// The custom KMS key ID in the destination region. If this field is not empty and cross-region replication is enabled, this key is used to encrypt the cross-region replication.
	//
	// example:
	//
	// 4ed37b1e-da51-4187-aceb-9db4f9b7148b
	DestinationKmsKeyId *string `json:"DestinationKmsKeyId,omitempty" xml:"DestinationKmsKeyId,omitempty"`
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
	// This parameter is required only when **AppConsistent*	- is set to **true**. The path of the post-thaw script to execute after creating an application-consistent snapshot.
	//
	// example:
	//
	// /tmp/postscript.sh
	PostScriptPath *string `json:"PostScriptPath,omitempty" xml:"PostScriptPath,omitempty"`
	// This parameter is required only when **AppConsistent*	- is set to **true**. The path of the pre-freeze script to execute before creating an application-consistent snapshot.
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
	// Specifies whether to create a snapshot-consistent group. Creating a snapshot-consistent group is supported only when all cloud disk types are ESSD.
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

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetAppConsistent() *bool {
	return s.AppConsistent
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetDestinationKmsKeyId() *string {
	return s.DestinationKmsKeyId
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetDiskIdList() []*string {
	return s.DiskIdList
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetEnableFsFreeze() *bool {
	return s.EnableFsFreeze
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetEnableWriters() *bool {
	return s.EnableWriters
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetExcludeDiskIdList() []*string {
	return s.ExcludeDiskIdList
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetPostScriptPath() *string {
	return s.PostScriptPath
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetPreScriptPath() *string {
	return s.PreScriptPath
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetSnapshotGroup() *bool {
	return s.SnapshotGroup
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) GetTimeoutInSeconds() *int64 {
	return s.TimeoutInSeconds
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetAppConsistent(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.AppConsistent = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetDestinationKmsKeyId(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.DestinationKmsKeyId = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetDiskIdList(v []*string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.DiskIdList = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetEnableFsFreeze(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.EnableFsFreeze = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetEnableWriters(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.EnableWriters = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetExcludeDiskIdList(v []*string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.ExcludeDiskIdList = v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetPostScriptPath(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.PostScriptPath = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetPreScriptPath(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.PreScriptPath = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetRamRoleName(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.RamRoleName = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetSnapshotGroup(v bool) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.SnapshotGroup = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) SetTimeoutInSeconds(v int64) *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail {
	s.TimeoutInSeconds = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsUdmDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsHitTags struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag matching rule.
	//
	// - **EQUAL**: Matches both the tag key and tag value.
	//
	// - **NOT**: Matches the tag key but does not match the tag value.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) GoString() string {
	return s.String()
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) GetKey() *string {
	return s.Key
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) GetOperator() *string {
	return s.Operator
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) GetValue() *string {
	return s.Value
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) SetKey(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags {
	s.Key = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) SetOperator(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags {
	s.Operator = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) SetValue(v string) *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags {
	s.Value = &v
	return s
}

func (s *DescribePolicyBindingsResponseBodyPolicyBindingsHitTags) Validate() error {
	return dara.Validate(s)
}
