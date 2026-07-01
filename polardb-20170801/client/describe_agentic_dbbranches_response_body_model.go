// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAgenticDBBranchesResponseBodyItems) *DescribeAgenticDBBranchesResponseBody
	GetItems() []*DescribeAgenticDBBranchesResponseBodyItems
	SetPageNumber(v int32) *DescribeAgenticDBBranchesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBBranchesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAgenticDBBranchesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAgenticDBBranchesResponseBody
	GetTotalCount() *int32
}

type DescribeAgenticDBBranchesResponseBody struct {
	Items []*DescribeAgenticDBBranchesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// B2C3D4E5-F6A7-8901-BCDE-2345678901BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAgenticDBBranchesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchesResponseBody) GetItems() []*DescribeAgenticDBBranchesResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBBranchesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBBranchesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBBranchesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBBranchesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAgenticDBBranchesResponseBody) SetItems(v []*DescribeAgenticDBBranchesResponseBodyItems) *DescribeAgenticDBBranchesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBody) SetPageNumber(v int32) *DescribeAgenticDBBranchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBody) SetPageSize(v int32) *DescribeAgenticDBBranchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBody) SetRequestId(v string) *DescribeAgenticDBBranchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBody) SetTotalCount(v int32) *DescribeAgenticDBBranchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAgenticDBBranchesResponseBodyItems struct {
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
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBBranchesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetBranchComputeClusterId() *string {
	return s.BranchComputeClusterId
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetBranchComputeNodeCount() *int32 {
	return s.BranchComputeNodeCount
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetIsDefaultBranch() *bool {
	return s.IsDefaultBranch
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetLastActivatedAt() *string {
	return s.LastActivatedAt
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetMaxCU() *string {
	return s.MaxCU
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetMinCU() *string {
	return s.MinCU
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetBranchComputeClusterId(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.BranchComputeClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetBranchComputeNodeCount(v int32) *DescribeAgenticDBBranchesResponseBodyItems {
	s.BranchComputeNodeCount = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetBranchId(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetBranchName(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.BranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetCreateTime(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetDescription(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetIsDefaultBranch(v bool) *DescribeAgenticDBBranchesResponseBodyItems {
	s.IsDefaultBranch = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetLastActivatedAt(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.LastActivatedAt = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetMaxCU(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.MaxCU = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetMinCU(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.MinCU = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetParentBranchId(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.ParentBranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetParentBranchName(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.ParentBranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetProjectId(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetProjectName(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetStatus(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) SetTenantId(v string) *DescribeAgenticDBBranchesResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
