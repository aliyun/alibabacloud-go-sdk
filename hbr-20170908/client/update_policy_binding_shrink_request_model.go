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
