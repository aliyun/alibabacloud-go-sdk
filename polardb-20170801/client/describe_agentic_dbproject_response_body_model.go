// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeAgenticDBProjectResponseBody
	GetCreateTime() *string
	SetDefaultBranchId(v string) *DescribeAgenticDBProjectResponseBody
	GetDefaultBranchId() *string
	SetDefaultBranchName(v string) *DescribeAgenticDBProjectResponseBody
	GetDefaultBranchName() *string
	SetDescription(v string) *DescribeAgenticDBProjectResponseBody
	GetDescription() *string
	SetProjectId(v string) *DescribeAgenticDBProjectResponseBody
	GetProjectId() *string
	SetProjectName(v string) *DescribeAgenticDBProjectResponseBody
	GetProjectName() *string
	SetRequestId(v string) *DescribeAgenticDBProjectResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAgenticDBProjectResponseBody
	GetStatus() *string
	SetTenantId(v string) *DescribeAgenticDBProjectResponseBody
	GetTenantId() *string
}

type DescribeAgenticDBProjectResponseBody struct {
	// example:
	//
	// 2026-06-10T11:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// br-1a2b3c4d5e6f
	DefaultBranchId *string `json:"DefaultBranchId,omitempty" xml:"DefaultBranchId,omitempty"`
	// example:
	//
	// main
	DefaultBranchName *string `json:"DefaultBranchName,omitempty" xml:"DefaultBranchName,omitempty"`
	// example:
	//
	// Production analytics database
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// I9J0K1L2-M3N4-5678-IJKL-789012345678
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBProjectResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBProjectResponseBody) GetDefaultBranchId() *string {
	return s.DefaultBranchId
}

func (s *DescribeAgenticDBProjectResponseBody) GetDefaultBranchName() *string {
	return s.DefaultBranchName
}

func (s *DescribeAgenticDBProjectResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticDBProjectResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBProjectResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBProjectResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBProjectResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBProjectResponseBody) SetCreateTime(v string) *DescribeAgenticDBProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetDefaultBranchId(v string) *DescribeAgenticDBProjectResponseBody {
	s.DefaultBranchId = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetDefaultBranchName(v string) *DescribeAgenticDBProjectResponseBody {
	s.DefaultBranchName = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetDescription(v string) *DescribeAgenticDBProjectResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetProjectId(v string) *DescribeAgenticDBProjectResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetProjectName(v string) *DescribeAgenticDBProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetRequestId(v string) *DescribeAgenticDBProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetStatus(v string) *DescribeAgenticDBProjectResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) SetTenantId(v string) *DescribeAgenticDBProjectResponseBody {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
