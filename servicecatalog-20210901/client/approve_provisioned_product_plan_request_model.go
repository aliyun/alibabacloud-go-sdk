// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveProvisionedProductPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApprovalAction(v string) *ApproveProvisionedProductPlanRequest
	GetApprovalAction() *string
	SetComment(v string) *ApproveProvisionedProductPlanRequest
	GetComment() *string
	SetPlanId(v string) *ApproveProvisionedProductPlanRequest
	GetPlanId() *string
}

type ApproveProvisionedProductPlanRequest struct {
	// The review action. Valid values:
	//
	// 	- Approve
	//
	// 	- Reject
	//
	// This parameter is required.
	//
	// example:
	//
	// Approve
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// The review description.
	//
	// example:
	//
	// Approved. You can create a resource.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s ApproveProvisionedProductPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ApproveProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *ApproveProvisionedProductPlanRequest) GetApprovalAction() *string {
	return s.ApprovalAction
}

func (s *ApproveProvisionedProductPlanRequest) GetComment() *string {
	return s.Comment
}

func (s *ApproveProvisionedProductPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *ApproveProvisionedProductPlanRequest) SetApprovalAction(v string) *ApproveProvisionedProductPlanRequest {
	s.ApprovalAction = &v
	return s
}

func (s *ApproveProvisionedProductPlanRequest) SetComment(v string) *ApproveProvisionedProductPlanRequest {
	s.Comment = &v
	return s
}

func (s *ApproveProvisionedProductPlanRequest) SetPlanId(v string) *ApproveProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

func (s *ApproveProvisionedProductPlanRequest) Validate() error {
	return dara.Validate(s)
}
