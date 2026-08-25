// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyBindingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedOptionsShrink(v string) *UpdatePolicyBindingShrinkRequest
	GetAdvancedOptionsShrink() *string
	SetDataSourceId(v string) *UpdatePolicyBindingShrinkRequest
	GetDataSourceId() *string
	SetDisabled(v bool) *UpdatePolicyBindingShrinkRequest
	GetDisabled() *bool
	SetExclude(v string) *UpdatePolicyBindingShrinkRequest
	GetExclude() *string
	SetInclude(v string) *UpdatePolicyBindingShrinkRequest
	GetInclude() *string
	SetPolicyBindingDescription(v string) *UpdatePolicyBindingShrinkRequest
	GetPolicyBindingDescription() *string
	SetPolicyId(v string) *UpdatePolicyBindingShrinkRequest
	GetPolicyId() *string
	SetSource(v string) *UpdatePolicyBindingShrinkRequest
	GetSource() *string
	SetSourceType(v string) *UpdatePolicyBindingShrinkRequest
	GetSourceType() *string
	SetSpeedLimit(v string) *UpdatePolicyBindingShrinkRequest
	GetSpeedLimit() *string
}

type UpdatePolicyBindingShrinkRequest struct {
	// The advanced options.
	AdvancedOptionsShrink *string `json:"AdvancedOptions,omitempty" xml:"AdvancedOptions,omitempty"`
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

func (s UpdatePolicyBindingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyBindingShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyBindingShrinkRequest) GetAdvancedOptionsShrink() *string {
	return s.AdvancedOptionsShrink
}

func (s *UpdatePolicyBindingShrinkRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *UpdatePolicyBindingShrinkRequest) GetDisabled() *bool {
	return s.Disabled
}

func (s *UpdatePolicyBindingShrinkRequest) GetExclude() *string {
	return s.Exclude
}

func (s *UpdatePolicyBindingShrinkRequest) GetInclude() *string {
	return s.Include
}

func (s *UpdatePolicyBindingShrinkRequest) GetPolicyBindingDescription() *string {
	return s.PolicyBindingDescription
}

func (s *UpdatePolicyBindingShrinkRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *UpdatePolicyBindingShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *UpdatePolicyBindingShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdatePolicyBindingShrinkRequest) GetSpeedLimit() *string {
	return s.SpeedLimit
}

func (s *UpdatePolicyBindingShrinkRequest) SetAdvancedOptionsShrink(v string) *UpdatePolicyBindingShrinkRequest {
	s.AdvancedOptionsShrink = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetDataSourceId(v string) *UpdatePolicyBindingShrinkRequest {
	s.DataSourceId = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetDisabled(v bool) *UpdatePolicyBindingShrinkRequest {
	s.Disabled = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetExclude(v string) *UpdatePolicyBindingShrinkRequest {
	s.Exclude = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetInclude(v string) *UpdatePolicyBindingShrinkRequest {
	s.Include = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetPolicyBindingDescription(v string) *UpdatePolicyBindingShrinkRequest {
	s.PolicyBindingDescription = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetPolicyId(v string) *UpdatePolicyBindingShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetSource(v string) *UpdatePolicyBindingShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetSourceType(v string) *UpdatePolicyBindingShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) SetSpeedLimit(v string) *UpdatePolicyBindingShrinkRequest {
	s.SpeedLimit = &v
	return s
}

func (s *UpdatePolicyBindingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
