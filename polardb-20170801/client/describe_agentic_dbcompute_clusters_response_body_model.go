// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBComputeClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeAgenticDBComputeClustersResponseBodyItems) *DescribeAgenticDBComputeClustersResponseBody
	GetItems() []*DescribeAgenticDBComputeClustersResponseBodyItems
	SetPageNumber(v int32) *DescribeAgenticDBComputeClustersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBComputeClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAgenticDBComputeClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAgenticDBComputeClustersResponseBody
	GetTotalCount() *int32
}

type DescribeAgenticDBComputeClustersResponseBody struct {
	// The list of compute instances.
	Items []*DescribeAgenticDBComputeClustersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F6A7B8C9-D0E1-2345-FABC-678901234FAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAgenticDBComputeClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBComputeClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBComputeClustersResponseBody) GetItems() []*DescribeAgenticDBComputeClustersResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBComputeClustersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBComputeClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBComputeClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBComputeClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAgenticDBComputeClustersResponseBody) SetItems(v []*DescribeAgenticDBComputeClustersResponseBodyItems) *DescribeAgenticDBComputeClustersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBody) SetPageNumber(v int32) *DescribeAgenticDBComputeClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBody) SetPageSize(v int32) *DescribeAgenticDBComputeClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBody) SetRequestId(v string) *DescribeAgenticDBComputeClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBody) SetTotalCount(v int32) *DescribeAgenticDBComputeClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBody) Validate() error {
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

type DescribeAgenticDBComputeClustersResponseBodyItems struct {
	// The associated branch ID.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The associated branch name.
	//
	// example:
	//
	// feature-analytics
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The compute instance ID.
	//
	// example:
	//
	// pc-g0lsayq8c5qe
	ComputeClusterId *string `json:"ComputeClusterId,omitempty" xml:"ComputeClusterId,omitempty"`
	// The number of compute nodes (fixed to 1 in the first phase).
	//
	// example:
	//
	// 1
	ComputeNodeCount *int32 `json:"ComputeNodeCount,omitempty" xml:"ComputeNodeCount,omitempty"`
	// The creation time. The time is in ISO 8601 format and uses the UTC time zone (format: YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2026-06-10T11:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Compute cluster for analytics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether it is the default branch.
	//
	// example:
	//
	// false
	IsDefaultBranch *bool `json:"IsDefaultBranch,omitempty" xml:"IsDefaultBranch,omitempty"`
	// The last activation time.
	//
	// The time is in ISO 8601 format and uses the UTC time zone (format: YYYY-MM-DDTHH:mm:ssZ).
	//
	// example:
	//
	// 2026-06-10T12:30:00Z
	LastActivatedAt *string `json:"LastActivatedAt,omitempty" xml:"LastActivatedAt,omitempty"`
	// The maximum compute unit.
	//
	// example:
	//
	// 2
	MaxCU *string `json:"MaxCU,omitempty" xml:"MaxCU,omitempty"`
	// The minimum compute unit.
	//
	// example:
	//
	// 0.25
	MinCU *string `json:"MinCU,omitempty" xml:"MinCU,omitempty"`
	// The operator type.
	//
	// example:
	//
	// Neon
	OperatorType *string `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The parent branch ID. An empty string indicates the root branch.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The parent branch name. An empty string indicates the root branch.
	//
	// example:
	//
	// main
	ParentBranchName *string `json:"ParentBranchName,omitempty" xml:"ParentBranchName,omitempty"`
	// The associated project ID.
	//
	// example:
	//
	// proj-a1b2c3d4e5f6
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The associated project name.
	//
	// example:
	//
	// analytics-prod
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The used data volume (in bytes).
	//
	// example:
	//
	// 1073741824
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// The associated tenant ID.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The associated tenant name.
	//
	// example:
	//
	// my-saas-app
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeAgenticDBComputeClustersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBComputeClustersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetComputeClusterId() *string {
	return s.ComputeClusterId
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetComputeNodeCount() *int32 {
	return s.ComputeNodeCount
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetIsDefaultBranch() *bool {
	return s.IsDefaultBranch
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetLastActivatedAt() *string {
	return s.LastActivatedAt
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetMaxCU() *string {
	return s.MaxCU
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetMinCU() *string {
	return s.MinCU
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetOperatorType() *string {
	return s.OperatorType
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetBranchId(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetBranchName(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.BranchName = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetComputeClusterId(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.ComputeClusterId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetComputeNodeCount(v int32) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.ComputeNodeCount = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetCreateTime(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetDescription(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetIsDefaultBranch(v bool) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.IsDefaultBranch = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetLastActivatedAt(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.LastActivatedAt = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetMaxCU(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.MaxCU = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetMinCU(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.MinCU = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetOperatorType(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.OperatorType = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetParentBranchId(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.ParentBranchId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetParentBranchName(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.ParentBranchName = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetProjectId(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetProjectName(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetStatus(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetStorageSize(v int64) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.StorageSize = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetTenantId(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) SetTenantName(v string) *DescribeAgenticDBComputeClustersResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeAgenticDBComputeClustersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
