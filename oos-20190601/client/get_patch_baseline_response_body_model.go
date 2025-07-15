// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatchBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPatchBaseline(v *GetPatchBaselineResponseBodyPatchBaseline) *GetPatchBaselineResponseBody
	GetPatchBaseline() *GetPatchBaselineResponseBodyPatchBaseline
	SetRequestId(v string) *GetPatchBaselineResponseBody
	GetRequestId() *string
}

type GetPatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *GetPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2C630E64-7273-57AC-A598-1B2B8B35CEA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPatchBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBody) GetPatchBaseline() *GetPatchBaselineResponseBodyPatchBaseline {
	return s.PatchBaseline
}

func (s *GetPatchBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPatchBaselineResponseBody) SetPatchBaseline(v *GetPatchBaselineResponseBodyPatchBaseline) *GetPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *GetPatchBaselineResponseBody) SetRequestId(v string) *GetPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPatchBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPatchBaselineResponseBodyPatchBaseline struct {
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
	// MypatchBaseline
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
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzmhzoaad5oq
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
	Tags []*GetPatchBaselineResponseBodyPatchBaselineTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user who last modified the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was last modified.
	//
	// example:
	//
	// 2021-09-08T07:26:38Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetPatchBaselineResponseBodyPatchBaseline) String() string {
	return dara.Prettify(s)
}

func (s GetPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetApprovedPatches() []*string {
	return s.ApprovedPatches
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetApprovedPatchesEnableNonSecurity() *bool {
	return s.ApprovedPatchesEnableNonSecurity
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetDescription() *string {
	return s.Description
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetId() *string {
	return s.Id
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetName() *string {
	return s.Name
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetRejectedPatches() []*string {
	return s.RejectedPatches
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetRejectedPatchesAction() *string {
	return s.RejectedPatchesAction
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetShareType() *string {
	return s.ShareType
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetSources() []*string {
	return s.Sources
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetTags() []*GetPatchBaselineResponseBodyPatchBaselineTags {
	return s.Tags
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovedPatches(v []*string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatches = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovedPatchesEnableNonSecurity(v bool) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetId(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetIsDefault(v bool) *GetPatchBaselineResponseBodyPatchBaseline {
	s.IsDefault = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetName(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetRejectedPatches(v []*string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatches = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetRejectedPatchesAction(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatchesAction = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetSources(v []*string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Sources = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetTags(v []*GetPatchBaselineResponseBodyPatchBaselineTags) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Tags = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) Validate() error {
	return dara.Validate(s)
}

type GetPatchBaselineResponseBodyPatchBaselineTags struct {
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

func (s GetPatchBaselineResponseBodyPatchBaselineTags) String() string {
	return dara.Prettify(s)
}

func (s GetPatchBaselineResponseBodyPatchBaselineTags) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) SetTagKey(v string) *GetPatchBaselineResponseBodyPatchBaselineTags {
	s.TagKey = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) SetTagValue(v string) *GetPatchBaselineResponseBodyPatchBaselineTags {
	s.TagValue = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) Validate() error {
	return dara.Validate(s)
}
