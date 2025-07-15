// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePatchBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalRules(v string) *CreatePatchBaselineRequest
	GetApprovalRules() *string
	SetApprovedPatches(v []*string) *CreatePatchBaselineRequest
	GetApprovedPatches() []*string
	SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineRequest
	GetApprovedPatchesEnableNonSecurity() *bool
	SetClientToken(v string) *CreatePatchBaselineRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePatchBaselineRequest
	GetDescription() *string
	SetName(v string) *CreatePatchBaselineRequest
	GetName() *string
	SetOperationSystem(v string) *CreatePatchBaselineRequest
	GetOperationSystem() *string
	SetRegionId(v string) *CreatePatchBaselineRequest
	GetRegionId() *string
	SetRejectedPatches(v []*string) *CreatePatchBaselineRequest
	GetRejectedPatches() []*string
	SetRejectedPatchesAction(v string) *CreatePatchBaselineRequest
	GetRejectedPatchesAction() *string
	SetResourceGroupId(v string) *CreatePatchBaselineRequest
	GetResourceGroupId() *string
	SetSources(v []*string) *CreatePatchBaselineRequest
	GetSources() []*string
	SetTags(v []*CreatePatchBaselineRequestTags) *CreatePatchBaselineRequest
	GetTags() []*CreatePatchBaselineRequestTags
}

type CreatePatchBaselineRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
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
	Tags []*CreatePatchBaselineRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreatePatchBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineRequest) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *CreatePatchBaselineRequest) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *CreatePatchBaselineRequest) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *CreatePatchBaselineRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePatchBaselineRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePatchBaselineRequest) GetName() *string {
	return s.Name
}

func (s *CreatePatchBaselineRequest) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *CreatePatchBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePatchBaselineRequest) GetRejectedPatches() []*string {
	return s.RejectedPatches
}

func (s *CreatePatchBaselineRequest) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *CreatePatchBaselineRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePatchBaselineRequest) GetSources() []*string {
	return s.Sources
}

func (s *CreatePatchBaselineRequest) GetTags() []*CreatePatchBaselineRequestTags {
	return s.Tags
}

func (s *CreatePatchBaselineRequest) SetApprovalRules(v string) *CreatePatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetApprovedPatches(v []*string) *CreatePatchBaselineRequest {
	s.ApprovedPatches = v
	return s
}

func (s *CreatePatchBaselineRequest) SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetClientToken(v string) *CreatePatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetDescription(v string) *CreatePatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetName(v string) *CreatePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetOperationSystem(v string) *CreatePatchBaselineRequest {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetRegionId(v string) *CreatePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetRejectedPatches(v []*string) *CreatePatchBaselineRequest {
	s.RejectedPatches = v
	return s
}

func (s *CreatePatchBaselineRequest) SetRejectedPatchesAction(v string) *CreatePatchBaselineRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetResourceGroupId(v string) *CreatePatchBaselineRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetSources(v []*string) *CreatePatchBaselineRequest {
	s.Sources = v
	return s
}

func (s *CreatePatchBaselineRequest) SetTags(v []*CreatePatchBaselineRequestTags) *CreatePatchBaselineRequest {
	s.Tags = v
	return s
}

func (s *CreatePatchBaselineRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePatchBaselineRequestTags struct {
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

func (s CreatePatchBaselineRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreatePatchBaselineRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreatePatchBaselineRequestTags) SetKey(v string) *CreatePatchBaselineRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePatchBaselineRequestTags) SetValue(v string) *CreatePatchBaselineRequestTags {
	s.Value = &v
	return s
}

func (s *CreatePatchBaselineRequestTags) Validate() error {
	return dara.Validate(s)
}
