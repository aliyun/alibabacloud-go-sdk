// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatchBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPatchBaseline(v *UpdatePatchBaselineResponseBodyPatchBaseline) *UpdatePatchBaselineResponseBody
	GetPatchBaseline() *UpdatePatchBaselineResponseBodyPatchBaseline
	SetRequestId(v string) *UpdatePatchBaselineResponseBody
	GetRequestId() *string
}

type UpdatePatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *UpdatePatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1457F46C-7AAE-59FA-BD12-0BDB3751E6F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePatchBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBody) GetPatchBaseline() *UpdatePatchBaselineResponseBodyPatchBaseline {
	return s.PatchBaseline
}

func (s *UpdatePatchBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePatchBaselineResponseBody) SetPatchBaseline(v *UpdatePatchBaselineResponseBodyPatchBaseline) *UpdatePatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *UpdatePatchBaselineResponseBody) SetRequestId(v string) *UpdatePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePatchBaselineResponseBody) Validate() error {
	if s.PatchBaseline != nil {
		if err := s.PatchBaseline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePatchBaselineResponseBodyPatchBaseline struct {
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
	// The creator of the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-07T03:42:56Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// UpdatePatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-445340b5c6504a85a300
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
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
	// rg-acfmy2zdbbjplii
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*UpdatePatchBaselineResponseBodyPatchBaselineTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user who updated the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was updated.
	//
	// example:
	//
	// 2021-09-08T07:26:37Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdatePatchBaselineResponseBodyPatchBaseline) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetDescription() *string {
	return s.Description
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetId() *string {
	return s.Id
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetName() *string {
	return s.Name
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetRejectedPatches() []*string {
	return s.RejectedPatches
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetShareType() *string {
	return s.ShareType
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetSources() []*string {
	return s.Sources
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetTags() []*UpdatePatchBaselineResponseBodyPatchBaselineTags {
	return s.Tags
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatches(v []*string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatches = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetId(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetName(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatches(v []*string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatches = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatchesAction(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatchesAction = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetSources(v []*string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Sources = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetTags(v []*UpdatePatchBaselineResponseBodyPatchBaselineTags) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Tags = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePatchBaselineResponseBodyPatchBaselineTags struct {
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

func (s UpdatePatchBaselineResponseBodyPatchBaselineTags) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineResponseBodyPatchBaselineTags) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) SetTagKey(v string) *UpdatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagKey = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) SetTagValue(v string) *UpdatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagValue = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) Validate() error {
	return dara.Validate(s)
}
