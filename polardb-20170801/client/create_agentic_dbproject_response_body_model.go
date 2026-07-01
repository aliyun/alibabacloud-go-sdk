// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticDBProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchComputeClusterId(v string) *CreateAgenticDBProjectResponseBody
	GetBranchComputeClusterId() *string
	SetCreateTime(v string) *CreateAgenticDBProjectResponseBody
	GetCreateTime() *string
	SetDefaultBranchId(v string) *CreateAgenticDBProjectResponseBody
	GetDefaultBranchId() *string
	SetDefaultBranchName(v string) *CreateAgenticDBProjectResponseBody
	GetDefaultBranchName() *string
	SetProjectId(v string) *CreateAgenticDBProjectResponseBody
	GetProjectId() *string
	SetProjectName(v string) *CreateAgenticDBProjectResponseBody
	GetProjectName() *string
	SetRequestId(v string) *CreateAgenticDBProjectResponseBody
	GetRequestId() *string
	SetTenantId(v string) *CreateAgenticDBProjectResponseBody
	GetTenantId() *string
}

type CreateAgenticDBProjectResponseBody struct {
	// The ID of the compute instance associated with the default branch.
	//
	// example:
	//
	// pc-g0lsayq8c5qe
	BranchComputeClusterId *string `json:"BranchComputeClusterId,omitempty" xml:"BranchComputeClusterId,omitempty"`
	// The time when the project was created.
	//
	// example:
	//
	// 2026-06-10T11:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default branch ID.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	DefaultBranchId *string `json:"DefaultBranchId,omitempty" xml:"DefaultBranchId,omitempty"`
	// The default branch name.
	//
	// example:
	//
	// main
	DefaultBranchName *string `json:"DefaultBranchName,omitempty" xml:"DefaultBranchName,omitempty"`
	// The ID of the new project.
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E5F6A7B8-C9D0-1234-EFAB-345678901234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateAgenticDBProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticDBProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgenticDBProjectResponseBody) GetBranchComputeClusterId() *string {
	return s.BranchComputeClusterId
}

func (s *CreateAgenticDBProjectResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAgenticDBProjectResponseBody) GetDefaultBranchId() *string {
	return s.DefaultBranchId
}

func (s *CreateAgenticDBProjectResponseBody) GetDefaultBranchName() *string {
	return s.DefaultBranchName
}

func (s *CreateAgenticDBProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateAgenticDBProjectResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateAgenticDBProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgenticDBProjectResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAgenticDBProjectResponseBody) SetBranchComputeClusterId(v string) *CreateAgenticDBProjectResponseBody {
	s.BranchComputeClusterId = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetCreateTime(v string) *CreateAgenticDBProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetDefaultBranchId(v string) *CreateAgenticDBProjectResponseBody {
	s.DefaultBranchId = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetDefaultBranchName(v string) *CreateAgenticDBProjectResponseBody {
	s.DefaultBranchName = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetProjectId(v string) *CreateAgenticDBProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetProjectName(v string) *CreateAgenticDBProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetRequestId(v string) *CreateAgenticDBProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) SetTenantId(v string) *CreateAgenticDBProjectResponseBody {
	s.TenantId = &v
	return s
}

func (s *CreateAgenticDBProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
