// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatchBaselineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalRules(v string) *UpdatePatchBaselineShrinkRequest
	GetApprovalRules() *string
	SetApprovedPatchesShrink(v string) *UpdatePatchBaselineShrinkRequest
	GetApprovedPatchesShrink() *string
	SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineShrinkRequest
	GetApprovedPatchesEnableNonSecurity() *bool
	SetClientToken(v string) *UpdatePatchBaselineShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdatePatchBaselineShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdatePatchBaselineShrinkRequest
	GetName() *string
	SetRegionId(v string) *UpdatePatchBaselineShrinkRequest
	GetRegionId() *string
	SetRejectedPatchesShrink(v string) *UpdatePatchBaselineShrinkRequest
	GetRejectedPatchesShrink() *string
	SetRejectedPatchesAction(v string) *UpdatePatchBaselineShrinkRequest
	GetRejectedPatchesAction() *string
	SetResourceGroupId(v string) *UpdatePatchBaselineShrinkRequest
	GetResourceGroupId() *string
	SetSourcesShrink(v string) *UpdatePatchBaselineShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *UpdatePatchBaselineShrinkRequest
	GetTagsShrink() *string
}

type UpdatePatchBaselineShrinkRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatchesShrink *string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty"`
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
	RejectedPatchesShrink *string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty"`
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
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdatePatchBaselineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineShrinkRequest) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *UpdatePatchBaselineShrinkRequest) GetApprovedPatchesShrink() *string {
	return s.ApprovedPatchesShrink
}

func (s *UpdatePatchBaselineShrinkRequest) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *UpdatePatchBaselineShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdatePatchBaselineShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePatchBaselineShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePatchBaselineShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePatchBaselineShrinkRequest) GetRejectedPatchesShrink() *string {
	return s.RejectedPatchesShrink
}

func (s *UpdatePatchBaselineShrinkRequest) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *UpdatePatchBaselineShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePatchBaselineShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *UpdatePatchBaselineShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdatePatchBaselineShrinkRequest) SetApprovalRules(v string) *UpdatePatchBaselineShrinkRequest {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetApprovedPatchesShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.ApprovedPatchesShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineShrinkRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetClientToken(v string) *UpdatePatchBaselineShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetDescription(v string) *UpdatePatchBaselineShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetName(v string) *UpdatePatchBaselineShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetRegionId(v string) *UpdatePatchBaselineShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetRejectedPatchesShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.RejectedPatchesShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetRejectedPatchesAction(v string) *UpdatePatchBaselineShrinkRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetResourceGroupId(v string) *UpdatePatchBaselineShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetSourcesShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetTagsShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
