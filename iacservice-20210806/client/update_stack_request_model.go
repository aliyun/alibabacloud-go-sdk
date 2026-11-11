// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateStackRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateStackRequest
	GetDescription() *string
	SetName(v string) *UpdateStackRequest
	GetName() *string
	SetRamRole(v string) *UpdateStackRequest
	GetRamRole() *string
	SetSkipPropertyValidation(v bool) *UpdateStackRequest
	GetSkipPropertyValidation() *bool
	SetSourcePath(v string) *UpdateStackRequest
	GetSourcePath() *string
	SetWorkingDirectory(v string) *UpdateStackRequest
	GetWorkingDirectory() *string
}

type UpdateStackRequest struct {
	// The idempotency token. The format is [0-9a-zA-Z-]{1,64}. A UUID is recommended.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the stack.
	//
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the stack. The name must be 2 to 128 characters in length, and can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). The name cannot start or end with a hyphen, underscore, or period.
	//
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The Resource Access Management (RAM) role authorized for the task. This role is used to automatically resume the execution of scheduled tasks in auto-triggered or offline scenarios.
	//
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// Specifies whether to skip the property validation of the provider. Default value: false.
	SkipPropertyValidation *bool `json:"skipPropertyValidation,omitempty" xml:"skipPropertyValidation,omitempty"`
	// The path value of the configuration source. The value cannot exceed 1,000 characters.
	//
	// - If the source is OSS, the value is in the oss::<file URL> format. The file must be a ZIP file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip.
	//
	// - If the source is IAC_SERVICE_MODULE, the value is the template ID. Example: mod-xxxxx.
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The working directory where the configuration file is located. If the file is in the root directory, enter /. Examples: config/ or /.
	//
	// example:
	//
	// /
	WorkingDirectory *string `json:"workingDirectory,omitempty" xml:"workingDirectory,omitempty"`
}

func (s UpdateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateStackRequest) GetName() *string {
	return s.Name
}

func (s *UpdateStackRequest) GetRamRole() *string {
	return s.RamRole
}

func (s *UpdateStackRequest) GetSkipPropertyValidation() *bool {
	return s.SkipPropertyValidation
}

func (s *UpdateStackRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *UpdateStackRequest) GetWorkingDirectory() *string {
	return s.WorkingDirectory
}

func (s *UpdateStackRequest) SetClientToken(v string) *UpdateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackRequest) SetDescription(v string) *UpdateStackRequest {
	s.Description = &v
	return s
}

func (s *UpdateStackRequest) SetName(v string) *UpdateStackRequest {
	s.Name = &v
	return s
}

func (s *UpdateStackRequest) SetRamRole(v string) *UpdateStackRequest {
	s.RamRole = &v
	return s
}

func (s *UpdateStackRequest) SetSkipPropertyValidation(v bool) *UpdateStackRequest {
	s.SkipPropertyValidation = &v
	return s
}

func (s *UpdateStackRequest) SetSourcePath(v string) *UpdateStackRequest {
	s.SourcePath = &v
	return s
}

func (s *UpdateStackRequest) SetWorkingDirectory(v string) *UpdateStackRequest {
	s.WorkingDirectory = &v
	return s
}

func (s *UpdateStackRequest) Validate() error {
	return dara.Validate(s)
}
