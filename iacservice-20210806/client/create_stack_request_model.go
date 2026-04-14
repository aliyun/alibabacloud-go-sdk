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
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// Stack to create ecs and related resource for multiple enviroments.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
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
