// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAgenticDBProjectRequest
	GetDBClusterId() *string
	SetDefaultBranchName(v string) *CreateAgenticDBProjectRequest
	GetDefaultBranchName() *string
	SetDescription(v string) *CreateAgenticDBProjectRequest
	GetDescription() *string
	SetProjectName(v string) *CreateAgenticDBProjectRequest
	GetProjectName() *string
	SetRegionId(v string) *CreateAgenticDBProjectRequest
	GetRegionId() *string
	SetTenantId(v string) *CreateAgenticDBProjectRequest
	GetTenantId() *string
}

type CreateAgenticDBProjectRequest struct {
	// The AgenticDB cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The default branch name. Default value: main.
	//
	// example:
	//
	// main
	DefaultBranchName *string `json:"DefaultBranchName,omitempty" xml:"DefaultBranchName,omitempty"`
	// The description of the project.
	//
	// example:
	//
	// Production analytics database
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The project name. The name must be unique within the same tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the tenant to which the project belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateAgenticDBProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBProjectRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAgenticDBProjectRequest) GetDefaultBranchName() *string {
	return s.DefaultBranchName
}

func (s *CreateAgenticDBProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAgenticDBProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateAgenticDBProjectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAgenticDBProjectRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAgenticDBProjectRequest) SetDBClusterId(v string) *CreateAgenticDBProjectRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetDefaultBranchName(v string) *CreateAgenticDBProjectRequest {
	s.DefaultBranchName = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetDescription(v string) *CreateAgenticDBProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetProjectName(v string) *CreateAgenticDBProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetRegionId(v string) *CreateAgenticDBProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) SetTenantId(v string) *CreateAgenticDBProjectRequest {
	s.TenantId = &v
	return s
}

func (s *CreateAgenticDBProjectRequest) Validate() error {
	return dara.Validate(s)
}
