// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePatchBaselineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalRules(v string) *CreatePatchBaselineShrinkRequest
	GetApprovalRules() *string
	SetApprovedPatchesShrink(v string) *CreatePatchBaselineShrinkRequest
	GetApprovedPatchesShrink() *string
	SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineShrinkRequest
	GetApprovedPatchesEnableNonSecurity() *bool
	SetClientToken(v string) *CreatePatchBaselineShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePatchBaselineShrinkRequest
	GetDescription() *string
	SetName(v string) *CreatePatchBaselineShrinkRequest
	GetName() *string
	SetOperationSystem(v string) *CreatePatchBaselineShrinkRequest
	GetOperationSystem() *string
	SetRegionId(v string) *CreatePatchBaselineShrinkRequest
	GetRegionId() *string
	SetRejectedPatchesShrink(v string) *CreatePatchBaselineShrinkRequest
	GetRejectedPatchesShrink() *string
	SetRejectedPatchesAction(v string) *CreatePatchBaselineShrinkRequest
	GetRejectedPatchesAction() *string
	SetResourceGroupId(v string) *CreatePatchBaselineShrinkRequest
	GetResourceGroupId() *string
	SetSourcesShrink(v string) *CreatePatchBaselineShrinkRequest
	GetSourcesShrink() *string
	SetTagsShrink(v string) *CreatePatchBaselineShrinkRequest
	GetTagsShrink() *string
}

type CreatePatchBaselineShrinkRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatchesShrink *string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty"`
	// Specifies whether the approved patch involves updates other than security-related updates.
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
	// PatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the region in which you want to create a patch baseline.
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

func (s CreatePatchBaselineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineShrinkRequest) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *CreatePatchBaselineShrinkRequest) GetApprovedPatchesShrink() *string {
	return s.ApprovedPatchesShrink
}

func (s *CreatePatchBaselineShrinkRequest) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *CreatePatchBaselineShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePatchBaselineShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePatchBaselineShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreatePatchBaselineShrinkRequest) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *CreatePatchBaselineShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePatchBaselineShrinkRequest) GetRejectedPatchesShrink() *string {
	return s.RejectedPatchesShrink
}

func (s *CreatePatchBaselineShrinkRequest) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *CreatePatchBaselineShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePatchBaselineShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreatePatchBaselineShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreatePatchBaselineShrinkRequest) SetApprovalRules(v string) *CreatePatchBaselineShrinkRequest {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetApprovedPatchesShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.ApprovedPatchesShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineShrinkRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetClientToken(v string) *CreatePatchBaselineShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetDescription(v string) *CreatePatchBaselineShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetName(v string) *CreatePatchBaselineShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetOperationSystem(v string) *CreatePatchBaselineShrinkRequest {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetRegionId(v string) *CreatePatchBaselineShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetRejectedPatchesShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.RejectedPatchesShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetRejectedPatchesAction(v string) *CreatePatchBaselineShrinkRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetResourceGroupId(v string) *CreatePatchBaselineShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetSourcesShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetTagsShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
