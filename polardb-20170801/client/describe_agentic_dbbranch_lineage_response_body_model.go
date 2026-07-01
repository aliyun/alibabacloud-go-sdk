// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBBranchLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnchorBranchId(v string) *DescribeAgenticDBBranchLineageResponseBody
	GetAnchorBranchId() *string
	SetItems(v []*DescribeAgenticDBBranchLineageResponseBodyItems) *DescribeAgenticDBBranchLineageResponseBody
	GetItems() []*DescribeAgenticDBBranchLineageResponseBodyItems
	SetNodeCount(v int32) *DescribeAgenticDBBranchLineageResponseBody
	GetNodeCount() *int32
	SetProjectId(v string) *DescribeAgenticDBBranchLineageResponseBody
	GetProjectId() *string
	SetProjectName(v string) *DescribeAgenticDBBranchLineageResponseBody
	GetProjectName() *string
	SetRequestId(v string) *DescribeAgenticDBBranchLineageResponseBody
	GetRequestId() *string
	SetRootBranchId(v string) *DescribeAgenticDBBranchLineageResponseBody
	GetRootBranchId() *string
	SetTenantId(v string) *DescribeAgenticDBBranchLineageResponseBody
	GetTenantId() *string
}

type DescribeAgenticDBBranchLineageResponseBody struct {
	// The anchor branch ID.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	AnchorBranchId *string `json:"AnchorBranchId,omitempty" xml:"AnchorBranchId,omitempty"`
	// The list of lineage nodes.
	Items []*DescribeAgenticDBBranchLineageResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The total number of returned nodes.
	//
	// example:
	//
	// 6
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The project ID.
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
	// B2C3D4E5-F6A7-8901-BCDE-2345678901BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The primary branch ID.
	//
	// example:
	//
	// br-root
	RootBranchId *string `json:"RootBranchId,omitempty" xml:"RootBranchId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// t-4b83e0da66674951
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAgenticDBBranchLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchLineageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetAnchorBranchId() *string {
	return s.AnchorBranchId
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetItems() []*DescribeAgenticDBBranchLineageResponseBodyItems {
	return s.Items
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetRootBranchId() *string {
	return s.RootBranchId
}

func (s *DescribeAgenticDBBranchLineageResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetAnchorBranchId(v string) *DescribeAgenticDBBranchLineageResponseBody {
	s.AnchorBranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetItems(v []*DescribeAgenticDBBranchLineageResponseBodyItems) *DescribeAgenticDBBranchLineageResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetNodeCount(v int32) *DescribeAgenticDBBranchLineageResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetProjectId(v string) *DescribeAgenticDBBranchLineageResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetProjectName(v string) *DescribeAgenticDBBranchLineageResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetRequestId(v string) *DescribeAgenticDBBranchLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetRootBranchId(v string) *DescribeAgenticDBBranchLineageResponseBody {
	s.RootBranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) SetTenantId(v string) *DescribeAgenticDBBranchLineageResponseBody {
	s.TenantId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBody) Validate() error {
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

type DescribeAgenticDBBranchLineageResponseBodyItems struct {
	// The compute cluster ID of the branch.
	//
	// example:
	//
	// pc-g0lsayq8c5qe
	BranchComputeClusterId *string `json:"BranchComputeClusterId,omitempty" xml:"BranchComputeClusterId,omitempty"`
	// The branch description.
	//
	// example:
	//
	// Feature branch for analytics
	BranchDescription *string `json:"BranchDescription,omitempty" xml:"BranchDescription,omitempty"`
	// The branch ID.
	//
	// example:
	//
	// br-7g8h9i0j1k2l
	BranchId *string `json:"BranchId,omitempty" xml:"BranchId,omitempty"`
	// The branch name.
	//
	// example:
	//
	// feature-analytics
	BranchName *string `json:"BranchName,omitempty" xml:"BranchName,omitempty"`
	// The time when the branch was created.
	//
	// example:
	//
	// 2026-06-10T11:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The depth relative to the anchor branch. The anchor branch has a depth of 0. Ancestor branches have negative values. Descendant branches have positive values.
	//
	// example:
	//
	// 0
	Depth *int32 `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The total number of direct child branches.
	//
	// example:
	//
	// 3
	DirectChildCount *int32 `json:"DirectChildCount,omitempty" xml:"DirectChildCount,omitempty"`
	// Indicates whether more ancestor nodes exist but are not returned.
	//
	// example:
	//
	// false
	HasMoreAncestors *bool `json:"HasMoreAncestors,omitempty" xml:"HasMoreAncestors,omitempty"`
	// Indicates whether more child nodes exist but are not returned.
	//
	// example:
	//
	// false
	HasMoreChildren *bool `json:"HasMoreChildren,omitempty" xml:"HasMoreChildren,omitempty"`
	// Indicates whether the branch is the anchor branch.
	//
	// example:
	//
	// true
	IsAnchor *bool `json:"IsAnchor,omitempty" xml:"IsAnchor,omitempty"`
	// example:
	//
	// false
	IsDefaultBranch *bool `json:"IsDefaultBranch,omitempty" xml:"IsDefaultBranch,omitempty"`
	// Indicates whether the branch is the primary branch.
	//
	// example:
	//
	// false
	IsRoot *bool `json:"IsRoot,omitempty" xml:"IsRoot,omitempty"`
	// The parent branch ID.
	//
	// example:
	//
	// br-1a2b3c4d5e6f
	ParentBranchId *string `json:"ParentBranchId,omitempty" xml:"ParentBranchId,omitempty"`
	// The parent branch name.
	//
	// example:
	//
	// main
	ParentBranchName *string `json:"ParentBranchName,omitempty" xml:"ParentBranchName,omitempty"`
	// The branch status. Valid values:
	//
	// - Active
	//
	// - Destroying
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAgenticDBBranchLineageResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBBranchLineageResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetBranchComputeClusterId() *string {
	return s.BranchComputeClusterId
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetBranchDescription() *string {
	return s.BranchDescription
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetBranchId() *string {
	return s.BranchId
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetBranchName() *string {
	return s.BranchName
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetDepth() *int32 {
	return s.Depth
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetDirectChildCount() *int32 {
	return s.DirectChildCount
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetHasMoreAncestors() *bool {
	return s.HasMoreAncestors
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetHasMoreChildren() *bool {
	return s.HasMoreChildren
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetIsAnchor() *bool {
	return s.IsAnchor
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetIsDefaultBranch() *bool {
	return s.IsDefaultBranch
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetIsRoot() *bool {
	return s.IsRoot
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetParentBranchId() *string {
	return s.ParentBranchId
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetParentBranchName() *string {
	return s.ParentBranchName
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetBranchComputeClusterId(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.BranchComputeClusterId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetBranchDescription(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.BranchDescription = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetBranchId(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.BranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetBranchName(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.BranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetCreateTime(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetDepth(v int32) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.Depth = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetDirectChildCount(v int32) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.DirectChildCount = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetHasMoreAncestors(v bool) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.HasMoreAncestors = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetHasMoreChildren(v bool) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.HasMoreChildren = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetIsAnchor(v bool) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.IsAnchor = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetIsDefaultBranch(v bool) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.IsDefaultBranch = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetIsRoot(v bool) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.IsRoot = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetParentBranchId(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.ParentBranchId = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetParentBranchName(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.ParentBranchName = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) SetStatus(v string) *DescribeAgenticDBBranchLineageResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeAgenticDBBranchLineageResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
