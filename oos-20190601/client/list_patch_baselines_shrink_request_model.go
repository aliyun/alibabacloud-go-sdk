// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatchBaselinesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovedPatchesShrink(v string) *ListPatchBaselinesShrinkRequest
	GetApprovedPatchesShrink() *string
	SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesShrinkRequest
	GetApprovedPatchesEnableNonSecurity() *bool
	SetMaxResults(v int32) *ListPatchBaselinesShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListPatchBaselinesShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListPatchBaselinesShrinkRequest
	GetNextToken() *string
	SetOperationSystem(v string) *ListPatchBaselinesShrinkRequest
	GetOperationSystem() *string
	SetRegionId(v string) *ListPatchBaselinesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListPatchBaselinesShrinkRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListPatchBaselinesShrinkRequest
	GetShareType() *string
	SetSourcesShrink(v string) *ListPatchBaselinesShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *ListPatchBaselinesShrinkRequest
	GetTagsShrink() *string
}

type ListPatchBaselinesShrinkRequest struct {
	// The approved patches.
	ApprovedPatchesShrink *string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty"`
	// Specifies whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- Windows
	//
	// 	- Ubuntu
	//
	// 	- CentOS
	//
	// 	- Debian
	//
	// 	- AliyunLinux
	//
	// 	- RedhatEnterpriseLinux
	//
	// 	- Anolis
	//
	// 	- AlmaLinux
	//
	// example:
	//
	// AliyunLinux
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListPatchBaselinesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesShrinkRequest) GetApprovedPatchesShrink() *string {
	return s.ApprovedPatchesShrink
}

func (s *ListPatchBaselinesShrinkRequest) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *ListPatchBaselinesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPatchBaselinesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListPatchBaselinesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPatchBaselinesShrinkRequest) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *ListPatchBaselinesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPatchBaselinesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPatchBaselinesShrinkRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListPatchBaselinesShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *ListPatchBaselinesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListPatchBaselinesShrinkRequest) SetApprovedPatchesShrink(v string) *ListPatchBaselinesShrinkRequest {
	s.ApprovedPatchesShrink = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesShrinkRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetMaxResults(v int32) *ListPatchBaselinesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetName(v string) *ListPatchBaselinesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetNextToken(v string) *ListPatchBaselinesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetOperationSystem(v string) *ListPatchBaselinesShrinkRequest {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetRegionId(v string) *ListPatchBaselinesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetResourceGroupId(v string) *ListPatchBaselinesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetShareType(v string) *ListPatchBaselinesShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetSourcesShrink(v string) *ListPatchBaselinesShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetTagsShrink(v string) *ListPatchBaselinesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
