// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPatchBaselinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPatchBaselinesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPatchBaselinesResponseBody
	GetNextToken() *string
	SetPatchBaselines(v []*ListPatchBaselinesResponseBodyPatchBaselines) *ListPatchBaselinesResponseBody
	GetPatchBaselines() []*ListPatchBaselinesResponseBodyPatchBaselines
	SetRequestId(v string) *ListPatchBaselinesResponseBody
	GetRequestId() *string
}

type ListPatchBaselinesResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// The number of entries returned on each page.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The patch baselines.
	PatchBaselines []*ListPatchBaselinesResponseBodyPatchBaselines `json:"PatchBaselines,omitempty" xml:"PatchBaselines,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 432996A1-03C0-5C4C-A8E6-66C4110765B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPatchBaselinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPatchBaselinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPatchBaselinesResponseBody) GetPatchBaselines() []*ListPatchBaselinesResponseBodyPatchBaselines {
	return s.PatchBaselines
}

func (s *ListPatchBaselinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPatchBaselinesResponseBody) SetMaxResults(v int32) *ListPatchBaselinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetNextToken(v string) *ListPatchBaselinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetPatchBaselines(v []*ListPatchBaselinesResponseBodyPatchBaselines) *ListPatchBaselinesResponseBody {
	s.PatchBaselines = v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetRequestId(v string) *ListPatchBaselinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPatchBaselinesResponseBodyPatchBaselines struct {
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The user who created the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-08T03:41:23Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// ListPatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-c2838b5d89b540e19ee6
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the patch baseline is set as the default patch baseline.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// AliyunLinux
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek256ia6vhsndy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The configurations of patch sources.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags of the patch baseline.
	Tags []*ListPatchBaselinesResponseBodyPatchBaselinesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user who last updated the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was updated.
	//
	// example:
	//
	// 2021-09-08T03:44:34Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListPatchBaselinesResponseBodyPatchBaselines) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesResponseBodyPatchBaselines) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetDescription() *string {
	return s.Description
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetId() *string {
	return s.Id
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetName() *string {
	return s.Name
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetShareType() *string {
	return s.ShareType
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetSources() []*string {
	return s.Sources
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetTags() []*ListPatchBaselinesResponseBodyPatchBaselinesTags {
	return s.Tags
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetApprovedPatches(v []*string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ApprovedPatches = v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetCreatedBy(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedBy = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetCreatedDate(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedDate = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetDescription(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Description = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetId(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Id = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetIsDefault(v bool) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.IsDefault = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetName(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetOperationSystem(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetResourceGroupId(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetShareType(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetSources(v []*string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Sources = v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetTags(v []*ListPatchBaselinesResponseBodyPatchBaselinesTags) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Tags = v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetUpdatedBy(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedBy = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetUpdatedDate(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedDate = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) Validate() error {
	return dara.Validate(s)
}

type ListPatchBaselinesResponseBodyPatchBaselinesTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPatchBaselinesResponseBodyPatchBaselinesTags) String() string {
	return dara.Prettify(s)
}

func (s ListPatchBaselinesResponseBodyPatchBaselinesTags) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) SetTagKey(v string) *ListPatchBaselinesResponseBodyPatchBaselinesTags {
	s.TagKey = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) SetTagValue(v string) *ListPatchBaselinesResponseBodyPatchBaselinesTags {
	s.TagValue = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) Validate() error {
	return dara.Validate(s)
}
