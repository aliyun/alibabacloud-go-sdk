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
	SetSource(v string) *CreateStackRequest
	GetSource() *string
	SetSourcePath(v string) *CreateStackRequest
	GetSourcePath() *string
	SetWorkingDirectory(v string) *CreateStackRequest
	GetWorkingDirectory() *string
}

type CreateStackRequest struct {
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. We recommend that you use a UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the resource stack. The value cannot exceed 256 characters.
	//
	// example:
	//
	// Stack to create ecs and related resource for multiple enviroments.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The resource stack name. The name must be 2 to 128 characters in length and can contain letters, digits, Chinese characters, hyphens (-), underscores (_), and periods (.). The name cannot start or end with a hyphen, underscore, or period.
	//
	// example:
	//
	// stack-test
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// The RAM role to be assigned to the task. This role is used to automatically continue the execution of scheduled tasks during automatic triggers or offline scenarios.
	//
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The creation source. Valid values:
	//
	// - OSS: a template from OSS.
	//
	// - IAC_SERVICE_MODULE: a template created in the automation service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path of the configuration source. The value cannot exceed 1000 characters.
	//
	// - If the source is OSS, the value is in the format oss::<file link> and must be a zip file. Example: oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	//
	// - If the source is IAC_SERVICE_MODULE, the value is a template ID. Example: mod-xxxxx
	//
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The working directory where the configuration file is located. Enter / if it is in the root directory. Example: config/ or /
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
