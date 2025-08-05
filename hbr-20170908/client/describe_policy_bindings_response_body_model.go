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
	// Return code, 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of results per query.
	//
	// Range: 10~100. Default: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Description of the return message. A successful response usually returns \\"successful\\", while an error will return a corresponding error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The token required to fetch the next page of policy and data source bindings.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// List of bound policies.
	PolicyBindings []*DescribePolicyBindingsResponseBodyPolicyBindings `json:"PolicyBindings,omitempty" xml:"PolicyBindings,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 5225929A-4EBD-55EE-9FE1-4A130E582A76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// - true: Success
	//
	// - false: Failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of records.
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
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindings struct {
	// Advanced options.
	AdvancedOptions *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty" type:"Struct"`
	// Whether the resource is automatically associated through the backup policy resource tag.
	//
	// example:
	//
	// false
	CreatedByTag *bool `json:"CreatedByTag,omitempty" xml:"CreatedByTag,omitempty"`
	// Creation time. UNIX timestamp, in seconds.
	//
	// example:
	//
	// 1661399570
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the role created in the RAM of the original account for cross-account backup.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Cross-account backup type. Supported values:
	//
	// - SELF_ACCOUNT: Backup within the same account
	//
	// - CROSS_ACCOUNT: Cross-account backup
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the original account for cross-account backup.
	//
	// example:
	//
	// 1480************
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// Data source ID.
	//
	// example:
	//
	// i-8vb************5ly
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Whether the policy is disbaled for this data source.
	//
	// - true: disabled
	//
	// - false: Not disabled
	//
	// example:
	//
	// true
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// This parameter is required only when **SourceType*	- is **ECS_FILE*	- or **File**. It specifies the file types that should not be backed up, and all files of these types will be excluded. Supports up to 255 characters.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Exclude *string `json:"Exclude,omitempty" xml:"Exclude,omitempty"`
	// Hit tag rules.
	HitTags []*DescribePolicyBindingsResponseBodyPolicyBindingsHitTags `json:"HitTags,omitempty" xml:"HitTags,omitempty" type:"Repeated"`
	// This parameter is required only when **SourceType*	- is **ECS_FILE*	- or **File**. It specifies the file types to be backed up, and all files of these types will be backed up. Supports up to 255 characters.
	//
	// example:
	//
	// [\\"*.doc\\",\\"*.xltm\\"]
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// Bound policy description.
	//
	// example:
	//
	// po-000************eslc-i-uf6************y5g
	PolicyBindingDescription *string `json:"PolicyBindingDescription,omitempty" xml:"PolicyBindingDescription,omitempty"`
	// Bound policy ID.
	//
	// example:
	//
	// pd-000************slc
	PolicyBindingId *string `json:"PolicyBindingId,omitempty" xml:"PolicyBindingId,omitempty"`
	// Policy ID.
	//
	// example:
	//
	// po-000************56y
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// - When **SourceType*	- is **OSS**, it indicates the prefix to be backed up. If not specified, it means backing up the entire root directory of the Bucket.
	//
	// - When **SourceType*	- is **ECS_FILE*	- or **File**, it indicates the file directory to be backed up. If not specified, it means backing up all directories.
	//
	// example:
	//
	// backup/
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Data source type, with the value range:
	//
	// - **UDM_ECS**: indicates ECS full machine backup
	//
	// example:
	//
	// UDM_ECS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required only when **SourceType*	- is **ECS_FILE*	- or **File**. It specifies the backup traffic control. The format is `{start}{end}{bandwidth}`. Multiple traffic control configurations are separated by commas, and the configured times must not overlap.
	//
	// - **start**: Start hour.
	//
	// - **end**: End hour.
	//
	// - **bandwidth**: Limit rate, in KB/s.
	//
	// example:
	//
	// 0:24:10240
	SpeedLimit *string `json:"SpeedLimit,omitempty" xml:"SpeedLimit,omitempty"`
	// Update time. UNIX timestamp, in seconds.
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
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptions struct {
	// Advanced options for large-scale file system backup.
	CommonFileSystemDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail `json:"CommonFileSystemDetail,omitempty" xml:"CommonFileSystemDetail,omitempty" type:"Struct"`
	// Advanced options for local NAS.
	CommonNasDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonNasDetail `json:"CommonNasDetail,omitempty" xml:"CommonNasDetail,omitempty" type:"Struct"`
	// Advanced options for file backup.
	FileDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsFileDetail `json:"FileDetail,omitempty" xml:"FileDetail,omitempty" type:"Struct"`
	// Advanced options for OSS backup.
	OssDetail *DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsOssDetail `json:"OssDetail,omitempty" xml:"OssDetail,omitempty" type:"Struct"`
	// Advanced options for full machine backup.
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
	return dara.Validate(s)
}

type DescribePolicyBindingsResponseBodyPolicyBindingsAdvancedOptionsCommonFileSystemDetail struct {
	// Backup shard size (number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Whether to switch to a full backup when an incremental backup fails. Values:
	//
	// - **true**: Switch to full backup on failure.
	//
	// - **false**: Do not switch to full backup on failure.
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
	// Backup client ID.
	//
	// example:
	//
	// c-0001eg6mcvjs93f46s2d
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Client group ID.
	//
	// example:
	//
	// cl-000gkcofngi04j6k680a
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Backup slice size (number of files).
	//
	// example:
	//
	// 100000
	FetchSliceSize *int64 `json:"FetchSliceSize,omitempty" xml:"FetchSliceSize,omitempty"`
	// Whether to switch to a full backup when an incremental backup fails. Values:
	//
	// - **true**: Switch to full backup on failure.
	//
	// - **false**: Do not switch to full backup on failure.
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
	// Whether to use advanced policies. Values:
	//
	// - **true**: Use.
	//
	// - **false**: Do not use.
	//
	// example:
	//
	// false
	AdvPolicy *bool `json:"AdvPolicy,omitempty" xml:"AdvPolicy,omitempty"`
	// Whether to enable VSS (Windows) functionality. Values:
	//
	// - **true**: Enable.
	//
	// - **false**: Disable.
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
	// Do not prompt for archive-type objects in the task statistics and failed file list.
	//
	// example:
	//
	// true
	IgnoreArchiveObject *bool `json:"IgnoreArchiveObject,omitempty" xml:"IgnoreArchiveObject,omitempty"`
	// Whether to delete the inventory file after the backup. This is only effective when using an OSS inventory. Supported values:
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
	// The name of the OSS inventory. If this value is not empty, the OSS inventory will be used for performance optimization.
	//
	// - It is recommended to use an inventory for backing up more than 100 million OSS objects to improve incremental performance. Storage costs for the inventory files are charged separately by the OSS service.
	//
	// - The generation of the OSS inventory file takes time, and the backup may fail before the inventory file is generated. You can wait for the next cycle to execute.
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
	// 是否创建应用一致性。仅云盘类型全部为ESSD时，支持创建快照应用一致性。
	//
	// example:
	//
	// false
	AppConsistent *bool `json:"AppConsistent,omitempty" xml:"AppConsistent,omitempty"`
	// The custom KMS key ID in the destination region. When this field is not empty and cross-region replication is enabled, the specified key will be used for encrypting the cross-region replication.
	//
	// example:
	//
	// 4ed37b1e-da51-4187-aceb-9db4f9b7148b
	DestinationKmsKeyId *string `json:"DestinationKmsKeyId,omitempty" xml:"DestinationKmsKeyId,omitempty"`
	// List of disk IDs that need protection. This value is empty when protecting all disks.
	DiskIdList []*string `json:"DiskIdList,omitempty" xml:"DiskIdList,omitempty" type:"Repeated"`
	// This parameter is required when **AppConsistent*	- is **true**. It indicates whether to use the Linux FsFreeze mechanism to ensure the file system is in a read-only consistent state before creating an application-consistent snapshot. The default value is true.
	//
	// example:
	//
	// true
	EnableFsFreeze *bool `json:"EnableFsFreeze,omitempty" xml:"EnableFsFreeze,omitempty"`
	// This parameter is required when **AppConsistent*	- is **true**. It determines whether to set an application-consistent snapshot:
	//
	// - **true**: Create an application-consistent snapshot
	//
	// - **false**: Create a file system-consistent snapshot
	//
	// The default value is true.
	//
	// example:
	//
	// true
	EnableWriters *bool `json:"EnableWriters,omitempty" xml:"EnableWriters,omitempty"`
	// List of disk IDs that do not need protection. This parameter is ignored if DiskIdList is not empty.
	ExcludeDiskIdList []*string `json:"ExcludeDiskIdList,omitempty" xml:"ExcludeDiskIdList,omitempty" type:"Repeated"`
	// This parameter is required when **AppConsistent*	- is **true**. It specifies the path of the unfreeze script to be executed after creating an application-consistent snapshot.
	//
	// example:
	//
	// /tmp/postscript.sh
	PostScriptPath *string `json:"PostScriptPath,omitempty" xml:"PostScriptPath,omitempty"`
	// This parameter is required when **AppConsistent*	- is **true**. It specifies the path of the freeze script to be executed before creating an application-consistent snapshot.
	//
	// example:
	//
	// /tmp/prescript.sh
	PreScriptPath *string `json:"PreScriptPath,omitempty" xml:"PreScriptPath,omitempty"`
	// This parameter is required when **AppConsistent*	- is **true**. It specifies the RAM role name needed for creating an application-consistent snapshot.
	//
	// example:
	//
	// AliyunECSInstanceForHbrRole
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// Indicates whether to create a snapshot consistency group. Only supported when all disk types are ESSD.
	//
	// example:
	//
	// true
	SnapshotGroup *bool `json:"SnapshotGroup,omitempty" xml:"SnapshotGroup,omitempty"`
	// This parameter is required when **AppConsistent*	- is **true**. It specifies the IO freeze timeout duration. The default value is 30 seconds.
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
	// Tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag matching rule.
	//
	// - **EQUAL**: Matches both the tag key and tag value.
	//
	// - **NOT**: Matches the tag key but not the tag value.
	//
	// example:
	//
	// EQUAL
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Tag value.
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
