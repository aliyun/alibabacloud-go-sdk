// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateStackRequest
	GetClientToken() *string
	SetDescription(v string) *CreateStackRequest
	GetDescription() *string
	SetName(v string) *CreateStackRequest
	GetName() *string
	SetParameterSetIds(v []*string) *CreateStackRequest
	GetParameterSetIds() []*string
	SetRamRole(v string) *CreateStackRequest
	GetRamRole() *string
	SetSkipPropertyValidation(v bool) *CreateStackRequest
	GetSkipPropertyValidation() *bool
	SetSource(v string) *CreateStackRequest
	GetSource() *string
	SetSourcePath(v string) *CreateStackRequest
	GetSourcePath() *string
	SetWorkingDirectory(v string) *CreateStackRequest
	GetWorkingDirectory() *string
}

type CreateStackRequest struct {
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. A UUID is recommended.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the stack. The description can be up to 256 characters in length.
	//
	// example:
	//
	// Stack to create ecs and related resource for multiple enviroments.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the stack. The name must be 2 to 128 characters in length and can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). It cannot start or end with a hyphen, underscore, or period.
	//
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The associated parameter sets.
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// The RAM role authorized for the task. This role is used to automatically resume the execution of scheduled tasks in auto-triggered or offline scenarios.
	//
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// Specifies whether to skip resource property validation for the provider. Default value: false.
	SkipPropertyValidation *bool `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	// The source of the stack. Valid values:
	//
	// - OSS: a template stored in Object Storage Service (OSS).
	//
	// - IAC_SERVICE_MODULE: a template created in the automation service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path of the source. The value can be up to 1,000 characters in length.
	//
	// - If the source is OSS, the value must be in the oss::<file URL> format and the file must be a ZIP file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If the source is IAC_SERVICE_MODULE, the value must be the template ID. Example: mod-xxxxx.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The working directory where the configuration file is located. If the file is in the root directory, enter /. Examples: config/ or /.
	//
	// example:
	//
	// /stack
	WorkingDirectory *string `json:"workingDirectory,omitempty" xml:"workingDirectory,omitempty"`
}

func (s CreateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStackRequest) GoString() string {
	return s.String()
}

func (s *CreateStackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStackRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStackRequest) GetName() *string {
	return s.Name
}

func (s *CreateStackRequest) GetParameterSetIds() []*string {
	return s.ParameterSetIds
}

func (s *CreateStackRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *CreateStackRequest) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *CreateStackRequest) GetSource() *string {
	return s.Source
}

func (s *CreateStackRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *CreateStackRequest) GetWorkingDirectory() *string {
	return s.WorkingDirectory
}

func (s *CreateStackRequest) SetClientToken(v string) *CreateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackRequest) SetDescription(v string) *CreateStackRequest {
	s.Description = &v
	return s
}

func (s *CreateStackRequest) SetName(v string) *CreateStackRequest {
	s.Name = &v
	return s
}

func (s *CreateStackRequest) SetParameterSetIds(v []*string) *CreateStackRequest {
	s.ParameterSetIds = v
	return s
}

func (s *CreateStackRequest) SetRamRole(v string) *CreateStackRequest {
	s.RamRole = &v
	return s
}

func (s *CreateStackRequest) SetSkipPropertyValidation(v bool) *CreateStackRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *CreateStackRequest) SetSource(v string) *CreateStackRequest {
	s.Source = &v
	return s
}

func (s *CreateStackRequest) SetSourcePath(v string) *CreateStackRequest {
	s.SourcePath = &v
	return s
}

func (s *CreateStackRequest) SetWorkingDirectory(v string) *CreateStackRequest {
	s.WorkingDirectory = &v
	return s
}

func (s *CreateStackRequest) Validate() error {
	return dara.Validate(s)
}
