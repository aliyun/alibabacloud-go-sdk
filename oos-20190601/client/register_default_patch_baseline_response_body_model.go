// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDefaultPatchBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPatchBaseline(v *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) *RegisterDefaultPatchBaselineResponseBody
	GetPatchBaseline() *RegisterDefaultPatchBaselineResponseBodyPatchBaseline
	SetRequestId(v string) *RegisterDefaultPatchBaselineResponseBody
	GetRequestId() *string
}

type RegisterDefaultPatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *RegisterDefaultPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D6850689-348D-59FC-AE13-BB0EDB7C4BE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponseBody) GetPatchBaseline() *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	return s.PatchBaseline
}

func (s *RegisterDefaultPatchBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterDefaultPatchBaselineResponseBody) SetPatchBaseline(v *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) *RegisterDefaultPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBody) SetRequestId(v string) *RegisterDefaultPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}

type RegisterDefaultPatchBaselineResponseBodyPatchBaseline struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
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
	// 2021-09-07T03:42:56Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// RegisterPatchBaseline
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4dpaq2yox6q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The user who last updated the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was last updated.
	//
	// example:
	//
	// 2021-09-07T03:42:56Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponseBodyPatchBaseline) String() string {
	return dara.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetApprovalRules() *string {
	return s.ApprovalRules
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetDescription() *string {
	return s.Description
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetId() *string {
	return s.Id
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetName() *string {
	return s.Name
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetOperationSystem() *string {
	return s.OperationSystem
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetShareType() *string {
	return s.ShareType
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetId(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetName(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) Validate() error {
	return dara.Validate(s)
}
