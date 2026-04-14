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
	SetSourcePath(v string) *UpdateStackRequest
	GetSourcePath() *string
	SetWorkingDirectory(v string) *UpdateStackRequest
	GetWorkingDirectory() *string
}

type UpdateStackRequest struct {
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// example:
	//
	// description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// stack-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// TestIacRole
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// example:
	//
	// oss::https://terraform-pipeline.oss-eu-central-1.aliyuncs.com/code.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
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
