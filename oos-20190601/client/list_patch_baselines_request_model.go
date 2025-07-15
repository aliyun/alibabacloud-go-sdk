// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatchBaselinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovedPatches(v []*string) *ListPatchBaselinesRequest
	GetApprovedPatches() []*string
	SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesRequest
	GetApprovedPatchesEnableNonSecurity() *bool
	SetMaxResults(v int32) *ListPatchBaselinesRequest
	GetMaxResults() *int32
	SetName(v string) *ListPatchBaselinesRequest
	GetName() *string
	SetNextToken(v string) *ListPatchBaselinesRequest
	GetNextToken() *string
	SetOperationSystem(v string) *ListPatchBaselinesRequest
	GetOperationSystem() *string
	SetRegionId(v string) *ListPatchBaselinesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListPatchBaselinesRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListPatchBaselinesRequest
	GetShareType() *string
	SetSources(v []*string) *ListPatchBaselinesRequest
	GetSources() []*string
	SetTags(v []*ListPatchBaselinesRequestTags) *ListPatchBaselinesRequest
	GetTags() []*ListPatchBaselinesRequestTags
}

type ListPatchBaselinesRequest struct {
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
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
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*ListPatchBaselinesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPatchBaselinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesRequest) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *ListPatchBaselinesRequest) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *ListPatchBaselinesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPatchBaselinesRequest) GetName() *string {
	return s.Name
}

func (s *ListPatchBaselinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPatchBaselinesRequest) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *ListPatchBaselinesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPatchBaselinesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPatchBaselinesRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListPatchBaselinesRequest) GetSources() []*string {
	return s.Sources
}

func (s *ListPatchBaselinesRequest) GetTags() []*ListPatchBaselinesRequestTags {
	return s.Tags
}

func (s *ListPatchBaselinesRequest) SetApprovedPatches(v []*string) *ListPatchBaselinesRequest {
	s.ApprovedPatches = v
	return s
}

func (s *ListPatchBaselinesRequest) SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetMaxResults(v int32) *ListPatchBaselinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetName(v string) *ListPatchBaselinesRequest {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetNextToken(v string) *ListPatchBaselinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetOperationSystem(v string) *ListPatchBaselinesRequest {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetRegionId(v string) *ListPatchBaselinesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetResourceGroupId(v string) *ListPatchBaselinesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetShareType(v string) *ListPatchBaselinesRequest {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetSources(v []*string) *ListPatchBaselinesRequest {
	s.Sources = v
	return s
}

func (s *ListPatchBaselinesRequest) SetTags(v []*ListPatchBaselinesRequestTags) *ListPatchBaselinesRequest {
	s.Tags = v
	return s
}

func (s *ListPatchBaselinesRequest) Validate() error {
	return dara.Validate(s)
}

type ListPatchBaselinesRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPatchBaselinesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesRequestTags) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListPatchBaselinesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListPatchBaselinesRequestTags) SetKey(v string) *ListPatchBaselinesRequestTags {
	s.Key = &v
	return s
}

func (s *ListPatchBaselinesRequestTags) SetValue(v string) *ListPatchBaselinesRequestTags {
	s.Value = &v
	return s
}

func (s *ListPatchBaselinesRequestTags) Validate() error {
	return dara.Validate(s)
}
