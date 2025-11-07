// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePatchBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPatchBaseline(v *CreatePatchBaselineResponseBodyPatchBaseline) *CreatePatchBaselineResponseBody
	GetPatchBaseline() *CreatePatchBaselineResponseBodyPatchBaseline
	SetRequestId(v string) *CreatePatchBaselineResponseBody
	GetRequestId() *string
}

type CreatePatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *CreatePatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A5173FF6-D10D-5E8C-8F71-943C2A3E25C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePatchBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBody) GetPatchBaseline() *CreatePatchBaselineResponseBodyPatchBaseline {
	return s.PatchBaseline
}

func (s *CreatePatchBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePatchBaselineResponseBody) SetPatchBaseline(v *CreatePatchBaselineResponseBodyPatchBaseline) *CreatePatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *CreatePatchBaselineResponseBody) SetRequestId(v string) *CreatePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePatchBaselineResponseBody) Validate() error {
	if s.PatchBaseline != nil {
		if err := s.PatchBaseline.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePatchBaselineResponseBodyPatchBaseline struct {
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
	// 2021-09-08T06:25:41Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// PatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-0a0aeda72ed147eb97ea
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// "ALLOW_AS_DEPENDENCY"
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3comlufxpva
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
	Tags []*CreatePatchBaselineResponseBodyPatchBaselineTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that last modified the information about the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the information about the patch baseline was last modified.
	//
	// example:
	//
	// 2021-09-08T06:25:41Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreatePatchBaselineResponseBodyPatchBaseline) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetDescription() *string {
	return s.Description
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetId() *string {
	return s.Id
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetName() *string {
	return s.Name
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetRejectedPatches() []*string {
	return s.RejectedPatches
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetShareType() *string {
	return s.ShareType
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetSources() []*string {
	return s.Sources
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetTags() []*CreatePatchBaselineResponseBodyPatchBaselineTags {
	return s.Tags
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatches(v []*string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatches = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetId(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetName(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatches(v []*string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatches = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatchesAction(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatchesAction = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetSources(v []*string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Sources = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetTags(v []*CreatePatchBaselineResponseBodyPatchBaselineTags) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Tags = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) Validate() error {
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

type CreatePatchBaselineResponseBodyPatchBaselineTags struct {
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

func (s CreatePatchBaselineResponseBodyPatchBaselineTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineResponseBodyPatchBaselineTags) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) GetTagKey() *string {
	return s.TagKey
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) GetTagValue() *string {
	return s.TagValue
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) SetTagKey(v string) *CreatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagKey = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) SetTagValue(v string) *CreatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagValue = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) Validate() error {
	return dara.Validate(s)
}
