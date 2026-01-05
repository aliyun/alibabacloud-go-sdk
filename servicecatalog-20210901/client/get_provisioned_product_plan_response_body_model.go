// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionedProductPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPlanDetail(v *GetProvisionedProductPlanResponseBodyPlanDetail) *GetProvisionedProductPlanResponseBody
	GetPlanDetail() *GetProvisionedProductPlanResponseBodyPlanDetail
	SetProductDetail(v *GetProvisionedProductPlanResponseBodyProductDetail) *GetProvisionedProductPlanResponseBody
	GetProductDetail() *GetProvisionedProductPlanResponseBodyProductDetail
	SetProductVersionDetail(v *GetProvisionedProductPlanResponseBodyProductVersionDetail) *GetProvisionedProductPlanResponseBody
	GetProductVersionDetail() *GetProvisionedProductPlanResponseBodyProductVersionDetail
	SetRequestId(v string) *GetProvisionedProductPlanResponseBody
	GetRequestId() *string
	SetResourceChanges(v []*GetProvisionedProductPlanResponseBodyResourceChanges) *GetProvisionedProductPlanResponseBody
	GetResourceChanges() []*GetProvisionedProductPlanResponseBodyResourceChanges
}

type GetProvisionedProductPlanResponseBody struct {
	// The details of the plan.
	PlanDetail *GetProvisionedProductPlanResponseBodyPlanDetail `json:"PlanDetail,omitempty" xml:"PlanDetail,omitempty" type:"Struct"`
	// The details of the product.
	ProductDetail *GetProvisionedProductPlanResponseBodyProductDetail `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty" type:"Struct"`
	// The details of the product version.
	ProductVersionDetail *GetProvisionedProductPlanResponseBodyProductVersionDetail `json:"ProductVersionDetail,omitempty" xml:"ProductVersionDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the resources to be changed in the plan.
	ResourceChanges []*GetProvisionedProductPlanResponseBodyResourceChanges `json:"ResourceChanges,omitempty" xml:"ResourceChanges,omitempty" type:"Repeated"`
}

func (s GetProvisionedProductPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBody) GetPlanDetail() *GetProvisionedProductPlanResponseBodyPlanDetail {
	return s.PlanDetail
}

func (s *GetProvisionedProductPlanResponseBody) GetProductDetail() *GetProvisionedProductPlanResponseBodyProductDetail {
	return s.ProductDetail
}

func (s *GetProvisionedProductPlanResponseBody) GetProductVersionDetail() *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	return s.ProductVersionDetail
}

func (s *GetProvisionedProductPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProvisionedProductPlanResponseBody) GetResourceChanges() []*GetProvisionedProductPlanResponseBodyResourceChanges {
	return s.ResourceChanges
}

func (s *GetProvisionedProductPlanResponseBody) SetPlanDetail(v *GetProvisionedProductPlanResponseBodyPlanDetail) *GetProvisionedProductPlanResponseBody {
	s.PlanDetail = v
	return s
}

func (s *GetProvisionedProductPlanResponseBody) SetProductDetail(v *GetProvisionedProductPlanResponseBodyProductDetail) *GetProvisionedProductPlanResponseBody {
	s.ProductDetail = v
	return s
}

func (s *GetProvisionedProductPlanResponseBody) SetProductVersionDetail(v *GetProvisionedProductPlanResponseBodyProductVersionDetail) *GetProvisionedProductPlanResponseBody {
	s.ProductVersionDetail = v
	return s
}

func (s *GetProvisionedProductPlanResponseBody) SetRequestId(v string) *GetProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBody) SetResourceChanges(v []*GetProvisionedProductPlanResponseBodyResourceChanges) *GetProvisionedProductPlanResponseBody {
	s.ResourceChanges = v
	return s
}

func (s *GetProvisionedProductPlanResponseBody) Validate() error {
	if s.PlanDetail != nil {
		if err := s.PlanDetail.Validate(); err != nil {
			return err
		}
	}
	if s.ProductDetail != nil {
		if err := s.ProductDetail.Validate(); err != nil {
			return err
		}
	}
	if s.ProductVersionDetail != nil {
		if err := s.ProductVersionDetail.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceChanges != nil {
		for _, item := range s.ResourceChanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProvisionedProductPlanResponseBodyPlanDetail struct {
	// The approval details of the plan.
	ApprovalDetail *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty" type:"Struct"`
	// An array that consists of reviewers.
	AssignedApprovers []*GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers `json:"AssignedApprovers,omitempty" xml:"AssignedApprovers,omitempty" type:"Repeated"`
	// The time when the plan was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-23T09:46:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the plan.
	//
	// example:
	//
	// Create an ECS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The purpose of the plan. Valid values:
	//
	// 	- LaunchProduct: launches the product.
	//
	// 	- UpdateProvisionedProduct: updates the information about the product instance.
	//
	// 	- TerminateProvisionedProduct: terminates the product instance.
	//
	// example:
	//
	// LaunchProduct
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The ID of the RAM entity to which the plan belongs.
	//
	// example:
	//
	// 27740196382623****
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The name of the RAM entity to which the plan belongs.
	//
	// example:
	//
	// endUser
	OwnerPrincipalName *string `json:"OwnerPrincipalName,omitempty" xml:"OwnerPrincipalName,omitempty"`
	// The type of the RAM entity to which the plan belongs. Valid values:
	//
	// 	- RamUser: a RAM user
	//
	// 	- RamRole: a RAM role
	//
	// example:
	//
	// RamUser
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// An array that consists of the parameters in the template.
	Parameters []*GetProvisionedProductPlanResponseBodyPlanDetailParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the plan.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	//
	// example:
	//
	// DEMO-ECS instance
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The type of the plan.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	//
	// example:
	//
	// Ros
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The ID of the product portfolio.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the product instance.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	//
	// example:
	//
	// DEMO-ECS instance
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the ROS stack.
	//
	// example:
	//
	// 137e31df-3754-40b4-be2f-c793ad84****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	//
	// example:
	//
	// cn-hangzhou
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The state of the plan. Valid values:
	//
	// 	- PreviewInProgress: The plan is being prechecked.
	//
	// 	- PreviewSuccess: The precheck is successful.
	//
	// 	- PreviewFailed: The precheck fails.
	//
	// 	- ApplicationInProgress: The plan is being approved.
	//
	// 	- ApplicationApproved: The plan is approved.
	//
	// 	- ApplicationRejected: The approval is rejected.
	//
	// 	- ExecuteInProgress: The plan is being run.
	//
	// 	- ExecuteSuccess: The plan is run.
	//
	// 	- ExecuteFailed: The plan fails to be run.
	//
	// 	- Canceled: The plan is canceled.
	//
	// example:
	//
	// PreviewSuccess
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message returned for the state.
	//
	// > This parameter is returned only when PreviewFailed or ExecuteFailed is returned for Status.
	//
	// example:
	//
	// Create stack failed: Resource CREATE failed: terraform stack sc-146611588617****-pp-bp1ddg1n2a***	- failure...
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// An array that consists of custom tags.
	Tags []*GetProvisionedProductPlanResponseBodyPlanDetailTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the plan belongs.
	//
	// example:
	//
	// 146611588617****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The last time when the task was modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-23T09:47:29Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetail) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetApprovalDetail() *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail {
	return s.ApprovalDetail
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetAssignedApprovers() []*GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers {
	return s.AssignedApprovers
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetDescription() *string {
	return s.Description
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetOperationType() *string {
	return s.OperationType
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetOwnerPrincipalId() *string {
	return s.OwnerPrincipalId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetOwnerPrincipalName() *string {
	return s.OwnerPrincipalName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetOwnerPrincipalType() *string {
	return s.OwnerPrincipalType
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetParameters() []*GetProvisionedProductPlanResponseBodyPlanDetailParameters {
	return s.Parameters
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetPlanId() *string {
	return s.PlanId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetPlanName() *string {
	return s.PlanName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetPlanType() *string {
	return s.PlanType
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetStackId() *string {
	return s.StackId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetStackRegionId() *string {
	return s.StackRegionId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetStatus() *string {
	return s.Status
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetTags() []*GetProvisionedProductPlanResponseBodyPlanDetailTags {
	return s.Tags
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetUid() *string {
	return s.Uid
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetApprovalDetail(v *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.ApprovalDetail = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetAssignedApprovers(v []*GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.AssignedApprovers = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetCreateTime(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetDescription(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.Description = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetOperationType(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.OperationType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetOwnerPrincipalId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.OwnerPrincipalId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetOwnerPrincipalName(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.OwnerPrincipalName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetOwnerPrincipalType(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.OwnerPrincipalType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetParameters(v []*GetProvisionedProductPlanResponseBodyPlanDetailParameters) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.Parameters = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetPlanId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.PlanId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetPlanName(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.PlanName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetPlanType(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.PlanType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetPortfolioId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetProductId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.ProductId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetProductVersionId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.ProductVersionId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetProvisionedProductId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.ProvisionedProductId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetProvisionedProductName(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.ProvisionedProductName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetStackId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.StackId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetStackRegionId(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.StackRegionId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetStatus(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.Status = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetStatusMessage(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.StatusMessage = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetTags(v []*GetProvisionedProductPlanResponseBodyPlanDetailTags) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.Tags = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetUid(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.Uid = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) SetUpdateTime(v string) *GetProvisionedProductPlanResponseBodyPlanDetail {
	s.UpdateTime = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetail) Validate() error {
	if s.ApprovalDetail != nil {
		if err := s.ApprovalDetail.Validate(); err != nil {
			return err
		}
	}
	if s.AssignedApprovers != nil {
		for _, item := range s.AssignedApprovers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail struct {
	// The operation records.
	OperationRecords []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords `json:"OperationRecords,omitempty" xml:"OperationRecords,omitempty" type:"Repeated"`
	// The operations that are being performed by the plan.
	TodoTaskActivities []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities `json:"TodoTaskActivities,omitempty" xml:"TodoTaskActivities,omitempty" type:"Repeated"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) GetOperationRecords() []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords {
	return s.OperationRecords
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) GetTodoTaskActivities() []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities {
	return s.TodoTaskActivities
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) SetOperationRecords(v []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail {
	s.OperationRecords = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) SetTodoTaskActivities(v []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail {
	s.TodoTaskActivities = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) Validate() error {
	if s.OperationRecords != nil {
		for _, item := range s.OperationRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TodoTaskActivities != nil {
		for _, item := range s.TodoTaskActivities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords struct {
	// The operation that is performed by the operator on the plan. Valid values:
	//
	// 	- Submit: submits the plan.
	//
	// 	- Cancel: cancels the plan.
	//
	// 	- Approve: approves the plan.
	//
	// 	- reject: rejectes the plan.
	//
	// example:
	//
	// Approve
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// The approval comment of the operator.
	//
	// example:
	//
	// Agreed.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the operation was performed.
	//
	// example:
	//
	// 2022-03-22T05:56:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The RAM entities that have performed operations on the plan.
	Operator *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Struct"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) GetApprovalAction() *string {
	return s.ApprovalAction
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) GetComment() *string {
	return s.Comment
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) GetOperator() *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator {
	return s.Operator
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) SetApprovalAction(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords {
	s.ApprovalAction = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) SetComment(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords {
	s.Comment = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) SetCreateTime(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords {
	s.CreateTime = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) SetOperator(v *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords {
	s.Operator = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) Validate() error {
	if s.Operator != nil {
		if err := s.Operator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator struct {
	// The ID of the RAM entity.
	//
	// example:
	//
	// 277401963826235***
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name RAM entity that servers as the operator.
	//
	// example:
	//
	// approver
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the RAM entity. Valid values:
	//
	// 	- RamUser: a RAM user.
	//
	// 	- RamRole: a RAM role.
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) SetPrincipalId(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator {
	s.PrincipalId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) SetPrincipalName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator {
	s.PrincipalName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) SetPrincipalType(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator {
	s.PrincipalType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities struct {
	// The name of the operation that is being performed by the plan.
	//
	// example:
	//
	// ApproverNode-1
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The tasks that are pending for review.
	Tasks []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) GetActivityName() *string {
	return s.ActivityName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) GetTasks() []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks {
	return s.Tasks
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) SetActivityName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities {
	s.ActivityName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) SetTasks(v []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities {
	s.Tasks = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks struct {
	// The RAM entities that can perform operations on the plan.
	Operator *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Struct"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) GetOperator() *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator {
	return s.Operator
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) SetOperator(v *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks {
	s.Operator = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) Validate() error {
	if s.Operator != nil {
		if err := s.Operator.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator struct {
	// The name of the RAM entity.
	//
	// example:
	//
	// approver
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the RAM entity. Valid values:
	//
	// 	- RamUser: a RAM user.
	//
	// 	- RamRole: a RAM role.
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) SetPrincipalName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator {
	s.PrincipalName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) SetPrincipalType(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator {
	s.PrincipalType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers struct {
	// The name of the RAM entity of the plan approver.
	//
	// example:
	//
	// approver
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the Resource Access Management (RAM) entity of the plan approver. Valid values:
	//
	// 	- RamUser: a RAM user.
	//
	// 	- RamRole: a RAM role.
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) SetPrincipalName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers {
	s.PrincipalName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) SetPrincipalType(v string) *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers {
	s.PrincipalType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyPlanDetailParameters struct {
	// The name of the input parameter for the template.
	//
	// example:
	//
	// instance_type
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the input parameter for the template.
	//
	// example:
	//
	// ecs.s6-c1m1.small
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailParameters) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailParameters) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) SetParameterKey(v string) *GetProvisionedProductPlanResponseBodyPlanDetailParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) SetParameterValue(v string) *GetProvisionedProductPlanResponseBodyPlanDetailParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyPlanDetailTags struct {
	// The key of the custom tag.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailTags) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailTags) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) GetKey() *string {
	return s.Key
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) GetValue() *string {
	return s.Value
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) SetKey(v string) *GetProvisionedProductPlanResponseBodyPlanDetailTags {
	s.Key = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) SetValue(v string) *GetProvisionedProductPlanResponseBodyPlanDetailTags {
	s.Value = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyProductDetail struct {
	// The creation time.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-12T06:10:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	//
	// example:
	//
	// This is a product description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	//
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:product/prod-bp18r7q127****
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// DEMO-Create an ECS instance
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates ROS.
	//
	// example:
	//
	// Ros
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	//
	// example:
	//
	// IT team
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyProductDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyProductDetail) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetDescription() *string {
	return s.Description
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetProductArn() *string {
	return s.ProductArn
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetProductName() *string {
	return s.ProductName
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetProductType() *string {
	return s.ProductType
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) GetProviderName() *string {
	return s.ProviderName
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetCreateTime(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetDescription(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.Description = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetProductArn(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.ProductArn = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetProductId(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.ProductId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetProductName(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.ProductName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetProductType(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.ProductType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) SetProviderName(v string) *GetProvisionedProductPlanResponseBodyProductDetail {
	s.ProviderName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductDetail) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyProductVersionDetail struct {
	// Indicates whether the product version is visible to end users. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The time when the product version was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-04-12T06:10:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product version.
	//
	// example:
	//
	// The description of the product version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// 	- Default: No recommendation information is provided. This is the default value.
	//
	// 	- Recommended: the recommendation version.
	//
	// 	- Latest: the latest version.
	//
	// 	- Deprecated: the version that is about to be deprecated.
	//
	// example:
	//
	// Default
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product to which the product version belongs.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name for the version of the product.
	//
	// example:
	//
	// 1.0.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the template.
	//
	// The value is fixed as RosTerraformTemplate, which indicates that the Terraform template is supported by ROS.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyProductVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyProductVersionDetail) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetActive() *bool {
	return s.Active
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetDescription() *string {
	return s.Description
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetGuidance() *string {
	return s.Guidance
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetActive(v bool) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.Active = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetCreateTime(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetDescription(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.Description = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetGuidance(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.Guidance = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetProductId(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.ProductId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetProductVersionId(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.ProductVersionId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetProductVersionName(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.ProductVersionName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetTemplateType(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.TemplateType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) SetTemplateUrl(v string) *GetProvisionedProductPlanResponseBodyProductVersionDetail {
	s.TemplateUrl = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyProductVersionDetail) Validate() error {
	return dara.Validate(s)
}

type GetProvisionedProductPlanResponseBodyResourceChanges struct {
	// The action that is performed on the resource. Valid values:
	//
	// 	- Add
	//
	// 	- Modify
	//
	// 	- Remove
	//
	// 	- None
	//
	// example:
	//
	// Add
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The logical ID of the resource.
	//
	// example:
	//
	// instance
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// >  This parameter is returned if the value of Action is Modify or Remove.
	//
	// example:
	//
	// i-bp13lmam3qd9q6il****
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// Indicates whether a replacement update is performed on the template. Valid values:
	//
	// 	- True: A replacement update is performed on the template.
	//
	// 	- False: A change is made on the template.
	//
	// 	- Conditional: A replacement update may be performed on the template. You can check whether a replacement update is performed when the template is in use.
	//
	// >  This parameter is returned if the value of Action is Modify.
	//
	// example:
	//
	// True
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
	// The resource type.
	//
	// example:
	//
	// alicloud_instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyResourceChanges) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyResourceChanges) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) GetAction() *string {
	return s.Action
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) GetPhysicalResourceId() *string {
	return s.PhysicalResourceId
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) GetReplacement() *string {
	return s.Replacement
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) SetAction(v string) *GetProvisionedProductPlanResponseBodyResourceChanges {
	s.Action = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) SetLogicalResourceId(v string) *GetProvisionedProductPlanResponseBodyResourceChanges {
	s.LogicalResourceId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) SetPhysicalResourceId(v string) *GetProvisionedProductPlanResponseBodyResourceChanges {
	s.PhysicalResourceId = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) SetReplacement(v string) *GetProvisionedProductPlanResponseBodyResourceChanges {
	s.Replacement = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) SetResourceType(v string) *GetProvisionedProductPlanResponseBodyResourceChanges {
	s.ResourceType = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyResourceChanges) Validate() error {
	return dara.Validate(s)
}
