// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatchBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalRules(v string) *UpdatePatchBaselineRequest
	GetApprovalRules() *string
	SetApprovedPatches(v []*string) *UpdatePatchBaselineRequest
	GetApprovedPatches() []*string
	SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineRequest
	GetApprovedPatchesEnableNonSecurity() *bool
	SetClientToken(v string) *UpdatePatchBaselineRequest
	GetClientToken() *string
	SetDescription(v string) *UpdatePatchBaselineRequest
	GetDescription() *string
	SetName(v string) *UpdatePatchBaselineRequest
	GetName() *string
	SetRegionId(v string) *UpdatePatchBaselineRequest
	GetRegionId() *string
	SetRejectedPatches(v []*string) *UpdatePatchBaselineRequest
	GetRejectedPatches() []*string
	SetRejectedPatchesAction(v string) *UpdatePatchBaselineRequest
	GetRejectedPatchesAction() *string
	SetResourceGroupId(v string) *UpdatePatchBaselineRequest
	GetResourceGroupId() *string
	SetSources(v []*string) *UpdatePatchBaselineRequest
	GetSources() []*string
	SetTags(v []*UpdatePatchBaselineRequestTags) *UpdatePatchBaselineRequest
	GetTags() []*UpdatePatchBaselineRequestTags
}

type UpdatePatchBaselineRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// UpdatePatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*UpdatePatchBaselineRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdatePatchBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineRequest) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *UpdatePatchBaselineRequest) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *UpdatePatchBaselineRequest) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *UpdatePatchBaselineRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePatchBaselineRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePatchBaselineRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePatchBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePatchBaselineRequest) GetRejectedPatches() []*string {
	return s.RejectedPatches
}

func (s *UpdatePatchBaselineRequest) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *UpdatePatchBaselineRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePatchBaselineRequest) GetSources() []*string {
	return s.Sources
}

func (s *UpdatePatchBaselineRequest) GetTags() []*UpdatePatchBaselineRequestTags {
	return s.Tags
}

func (s *UpdatePatchBaselineRequest) SetApprovalRules(v string) *UpdatePatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetApprovedPatches(v []*string) *UpdatePatchBaselineRequest {
	s.ApprovedPatches = v
	return s
}

func (s *UpdatePatchBaselineRequest) SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetClientToken(v string) *UpdatePatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetDescription(v string) *UpdatePatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetName(v string) *UpdatePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetRegionId(v string) *UpdatePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetRejectedPatches(v []*string) *UpdatePatchBaselineRequest {
	s.RejectedPatches = v
	return s
}

func (s *UpdatePatchBaselineRequest) SetRejectedPatchesAction(v string) *UpdatePatchBaselineRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetResourceGroupId(v string) *UpdatePatchBaselineRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetSources(v []*string) *UpdatePatchBaselineRequest {
	s.Sources = v
	return s
}

func (s *UpdatePatchBaselineRequest) SetTags(v []*UpdatePatchBaselineRequestTags) *UpdatePatchBaselineRequest {
	s.Tags = v
	return s
}

func (s *UpdatePatchBaselineRequest) Validate() error {
	return dara.Validate(s)
}

type UpdatePatchBaselineRequestTags struct {
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

func (s UpdatePatchBaselineRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineRequestTags) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdatePatchBaselineRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdatePatchBaselineRequestTags) SetKey(v string) *UpdatePatchBaselineRequestTags {
	s.Key = &v
	return s
}

func (s *UpdatePatchBaselineRequestTags) SetValue(v string) *UpdatePatchBaselineRequestTags {
	s.Value = &v
	return s
}

func (s *UpdatePatchBaselineRequestTags) Validate() error {
	return dara.Validate(s)
}
