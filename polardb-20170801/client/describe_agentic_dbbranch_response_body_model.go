// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBranchComputeClusterId(v string) *DescribeAgenticDBBranchResponseBody
	GetBranchComputeClusterId() *string
	SetBranchComputeNodeCount(v int32) *DescribeAgenticDBBranchResponseBody
	GetBranchComputeNodeCount() *int32
	SetBranchId(v string) *DescribeAgenticDBBranchResponseBody
	GetBranchId() *string
	SetBranchName(v string) *DescribeAgenticDBBranchResponseBody
	GetBranchName() *string
	SetCreateTime(v string) *DescribeAgenticDBBranchResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeAgenticDBBranchResponseBody
	GetDescription() *string
	SetIsDefaultBranch(v bool) *DescribeAgenticDBBranchResponseBody
	GetIsDefaultBranch() *bool
	SetLastActivatedAt(v string) *DescribeAgenticDBBranchResponseBody
	GetLastActivatedAt() *string
	SetMaxCU(v string) *DescribeAgenticDBBranchResponseBody
	GetMaxCU() *string
	SetMinCU(v string) *DescribeAgenticDBBranchResponseBody
	GetMinCU() *string
	SetParentBranchId(v string) *DescribeAgenticDBBranchResponseBody
	GetParentBranchId() *string
	SetParentBranchName(v string) *DescribeAgenticDBBranchResponseBody
	GetParentBranchName() *string
	SetProjectId(v string) *DescribeAgenticDBBranchResponseBody
	GetProjectId() *string
	SetProjectName(v string) *DescribeAgenticDBBranchResponseBody
	GetProjectName() *string
	SetRequestId(v string) *DescribeAgenticDBBranchResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeAgenticDBBranchResponseBody
	GetStatus() *string
	SetTenantId(v string) *DescribeAgenticDBBranchResponseBody
	GetTenantId() *string
}

type DescribeAgenticDBBranchResponseBody struct {
	// example:
	//
	// pc-g0lsayq8c5qe
	BranchComputeClusterId *string `json:"BranchComputeClusterId,omitempty" xml:"BranchComputeClusterId,omitempty"`
	// example:
	//
	// 1
	BranchComputeNodeCount *int32 `json:"BranchComputeNodeCount,omitempty" xml:"BranchComputeNodeCount,omitempty"`
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// example:
	//
	// feature-analytics
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// example:
	//
	// 2026-06-10T11:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// Feature branch for analytics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	IsDefaultBranch *bool `json:"IsDefaultBranch,omitempty" xml:"IsDefaultBranch,omitempty"`
	// example:
	//
	// 2026-06-10T12:30:00Z
	LastActivatedAt *string `json:"LastActivatedAt,omitempty" xml:"LastActivatedAt,omitempty"`
	// example:
	//
	// 2
	MaxCU *string `json:"MaxCU,omitempty" xml:"MaxCU,omitempty"`
	// example:
	//
	// 0.25
	MinCU *string `json:"MinCU,omitempty" xml:"MinCU,omitempty"`
	// example:
	//
	// br-1a2b3c4d5e6f
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// example:
	//
	// main
	ParentBranchName *string `json:"ParentBranchName,omitempty" xml:"ParentBranchName,omitempty"`
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
	// C3D4E5F6-A7B8-9012-CDEF-345678901CDE
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

func (s DescribeAgenticDBBranchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchResponseBody) GetBranchComputeClusterId() *string {
	return s.BranchComputeClusterId
}

func (s *DescribeAgenticDBBranchResponseBody) GetBranchComputeNodeCount() *int32 {
	return s.BranchComputeNodeCount
}

func (s *DescribeAgenticDBBranchResponseBody) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchResponseBody) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeAgenticDBBranchResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBBranchResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticDBBranchResponseBody) GetIsDefaultBranch() *bool {
	return s.IsDefaultBranch
}

func (s *DescribeAgenticDBBranchResponseBody) GetLastActivatedAt() *string {
	return s.LastActivatedAt
}

func (s *DescribeAgenticDBBranchResponseBody) GetMaxCU() *string {
	return s.MaxCU
}

func (s *DescribeAgenticDBBranchResponseBody) GetMinCU() *string {
	return s.MinCU
}

func (s *DescribeAgenticDBBranchResponseBody) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *DescribeAgenticDBBranchResponseBody) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *DescribeAgenticDBBranchResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBBranchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBBranchResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBBranchResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchResponseBody) SetBranchComputeClusterId(v string) *DescribeAgenticDBBranchResponseBody {
	s.BranchComputeClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetBranchComputeNodeCount(v int32) *DescribeAgenticDBBranchResponseBody {
	s.BranchComputeNodeCount = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetBranchId(v string) *DescribeAgenticDBBranchResponseBody {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetBranchName(v string) *DescribeAgenticDBBranchResponseBody {
	s.BranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetCreateTime(v string) *DescribeAgenticDBBranchResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetDescription(v string) *DescribeAgenticDBBranchResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetIsDefaultBranch(v bool) *DescribeAgenticDBBranchResponseBody {
	s.IsDefaultBranch = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetLastActivatedAt(v string) *DescribeAgenticDBBranchResponseBody {
	s.LastActivatedAt = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetMaxCU(v string) *DescribeAgenticDBBranchResponseBody {
	s.MaxCU = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetMinCU(v string) *DescribeAgenticDBBranchResponseBody {
	s.MinCU = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetParentBranchId(v string) *DescribeAgenticDBBranchResponseBody {
	s.ParentBranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetParentBranchName(v string) *DescribeAgenticDBBranchResponseBody {
	s.ParentBranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetProjectId(v string) *DescribeAgenticDBBranchResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetProjectName(v string) *DescribeAgenticDBBranchResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetRequestId(v string) *DescribeAgenticDBBranchResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetStatus(v string) *DescribeAgenticDBBranchResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) SetTenantId(v string) *DescribeAgenticDBBranchResponseBody {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchResponseBody) Validate() error {
	return dara.Validate(s)
}
