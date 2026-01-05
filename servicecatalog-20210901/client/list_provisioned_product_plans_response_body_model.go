// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProvisionedProductPlansResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProvisionedProductPlansResponseBody
	GetPageSize() *int32
	SetPlanDetails(v []*ListProvisionedProductPlansResponseBodyPlanDetails) *ListProvisionedProductPlansResponseBody
	GetPlanDetails() []*ListProvisionedProductPlansResponseBodyPlanDetails
	SetRequestId(v string) *ListProvisionedProductPlansResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProvisionedProductPlansResponseBody
	GetTotalCount() *int32
}

type ListProvisionedProductPlansResponseBody struct {
	// The page number of the returned page.
	//
	// Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array that consists of plans.
	PlanDetails []*ListProvisionedProductPlansResponseBodyPlanDetails `json:"PlanDetails,omitempty" xml:"PlanDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProvisionedProductPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProvisionedProductPlansResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProvisionedProductPlansResponseBody) GetPlanDetails() []*ListProvisionedProductPlansResponseBodyPlanDetails {
	return s.PlanDetails
}

func (s *ListProvisionedProductPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProvisionedProductPlansResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProvisionedProductPlansResponseBody) SetPageNumber(v int32) *ListProvisionedProductPlansResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBody) SetPageSize(v int32) *ListProvisionedProductPlansResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBody) SetPlanDetails(v []*ListProvisionedProductPlansResponseBodyPlanDetails) *ListProvisionedProductPlansResponseBody {
	s.PlanDetails = v
	return s
}

func (s *ListProvisionedProductPlansResponseBody) SetRequestId(v string) *ListProvisionedProductPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBody) SetTotalCount(v int32) *ListProvisionedProductPlansResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBody) Validate() error {
	if s.PlanDetails != nil {
		for _, item := range s.PlanDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProvisionedProductPlansResponseBodyPlanDetails struct {
	// An array that consists of reviewers.
	AssignedApprovers []*ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers `json:"AssignedApprovers,omitempty" xml:"AssignedApprovers,omitempty" type:"Repeated"`
	// The time when the plan was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-09-13T02:01:22Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the plan.
	//
	// example:
	//
	// For development team.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The purpose of the plan. Valid values:
	//
	// 	- LaunchProduct: launches the product. This is the default value.
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
	// 24477111603637****
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The name of the RAM entity to which the plan belongs.
	//
	// example:
	//
	// enduser
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
	Parameters []*ListProvisionedProductPlansResponseBodyPlanDetailsParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the plan.
	//
	// example:
	//
	// plan-bp18mmdh2u****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	//
	// example:
	//
	// DEMO-Create an ECS instance-637\\*\\*\\*\\*
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
	// port-bp1438kf2j****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	//
	// example:
	//
	// prod-bp1rtrnh2c****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	//
	// example:
	//
	// Create an ECS instance
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp19udk22v****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the product instance.
	//
	// example:
	//
	// pp-bp1c35162d****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	//
	// example:
	//
	// rds-MYSQL-875****
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the ROS stack.
	//
	// example:
	//
	// 2599090a-309e-4306-b989-17ba66a9****
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
	// 	- ApplicationInProgress: The plan is being reviewed.
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
	// > This parameter is returned only when PreviewFailed or ExecuteFailed is returned for the Status parameter.
	//
	// example:
	//
	// Create stack failed: Resource CREATE failed: terraform stack sc-146611588617****-pp-bp1ddg1n2a***	- failure...
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// An array that consists of custom tags.
	Tags []*ListProvisionedProductPlansResponseBodyPlanDetailsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	// 2022-07-18T06:02:35.075Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetails) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetails) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetAssignedApprovers() []*ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers {
	return s.AssignedApprovers
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetDescription() *string {
	return s.Description
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetOperationType() *string {
	return s.OperationType
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetOwnerPrincipalId() *string {
	return s.OwnerPrincipalId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetOwnerPrincipalName() *string {
	return s.OwnerPrincipalName
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetOwnerPrincipalType() *string {
	return s.OwnerPrincipalType
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetParameters() []*ListProvisionedProductPlansResponseBodyPlanDetailsParameters {
	return s.Parameters
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetPlanId() *string {
	return s.PlanId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetPlanName() *string {
	return s.PlanName
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetPlanType() *string {
	return s.PlanType
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetProductId() *string {
	return s.ProductId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetProductName() *string {
	return s.ProductName
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetStackId() *string {
	return s.StackId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetStackRegionId() *string {
	return s.StackRegionId
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetStatus() *string {
	return s.Status
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetTags() []*ListProvisionedProductPlansResponseBodyPlanDetailsTags {
	return s.Tags
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetUid() *string {
	return s.Uid
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetAssignedApprovers(v []*ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.AssignedApprovers = v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetCreateTime(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.CreateTime = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetDescription(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.Description = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetOperationType(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.OperationType = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetOwnerPrincipalId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.OwnerPrincipalId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetOwnerPrincipalName(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.OwnerPrincipalName = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetOwnerPrincipalType(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.OwnerPrincipalType = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetParameters(v []*ListProvisionedProductPlansResponseBodyPlanDetailsParameters) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.Parameters = v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetPlanId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.PlanId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetPlanName(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.PlanName = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetPlanType(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.PlanType = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetPortfolioId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.PortfolioId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetProductId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.ProductId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetProductName(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.ProductName = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetProductVersionId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.ProductVersionId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetProvisionedProductId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetProvisionedProductName(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.ProvisionedProductName = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetStackId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.StackId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetStackRegionId(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.StackRegionId = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetStatus(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.Status = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetStatusMessage(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.StatusMessage = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetTags(v []*ListProvisionedProductPlansResponseBodyPlanDetailsTags) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.Tags = v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetUid(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.Uid = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) SetUpdateTime(v string) *ListProvisionedProductPlansResponseBodyPlanDetails {
	s.UpdateTime = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetails) Validate() error {
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

type ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers struct {
	// The RAM entity name of the reviewer.
	//
	// example:
	//
	// endUser
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the RAM entity of the reviewer. Valid values:
	//
	// 	- RamUser: a RAM user
	//
	// 	- RamRole: a RAM role
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) SetPrincipalName(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers {
	s.PrincipalName = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) SetPrincipalType(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers {
	s.PrincipalType = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) Validate() error {
	return dara.Validate(s)
}

type ListProvisionedProductPlansResponseBodyPlanDetailsParameters struct {
	// The name of the parameter in the template.
	//
	// example:
	//
	// role_name
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter in the template.
	//
	// example:
	//
	// Test-8
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsParameters) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsParameters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) SetParameterKey(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsParameters {
	s.ParameterKey = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) SetParameterValue(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsParameters {
	s.ParameterValue = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) Validate() error {
	return dara.Validate(s)
}

type ListProvisionedProductPlansResponseBodyPlanDetailsTags struct {
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

func (s ListProvisionedProductPlansResponseBodyPlanDetailsTags) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsTags) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) GetKey() *string {
	return s.Key
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) GetValue() *string {
	return s.Value
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) SetKey(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsTags {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) SetValue(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsTags {
	s.Value = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) Validate() error {
	return dara.Validate(s)
}
