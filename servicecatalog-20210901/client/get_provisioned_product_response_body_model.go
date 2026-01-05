// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionedProductResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProvisionedProductDetail(v *GetProvisionedProductResponseBodyProvisionedProductDetail) *GetProvisionedProductResponseBody
	GetProvisionedProductDetail() *GetProvisionedProductResponseBodyProvisionedProductDetail
	SetRequestId(v string) *GetProvisionedProductResponseBody
	GetRequestId() *string
}

type GetProvisionedProductResponseBody struct {
	// The details of the product instance.
	ProvisionedProductDetail *GetProvisionedProductResponseBodyProvisionedProductDetail `json:"ProvisionedProductDetail,omitempty" xml:"ProvisionedProductDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProvisionedProductResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductResponseBody) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductResponseBody) GetProvisionedProductDetail() *GetProvisionedProductResponseBodyProvisionedProductDetail {
	return s.ProvisionedProductDetail
}

func (s *GetProvisionedProductResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProvisionedProductResponseBody) SetProvisionedProductDetail(v *GetProvisionedProductResponseBodyProvisionedProductDetail) *GetProvisionedProductResponseBody {
	s.ProvisionedProductDetail = v
	return s
}

func (s *GetProvisionedProductResponseBody) SetRequestId(v string) *GetProvisionedProductResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProvisionedProductResponseBody) Validate() error {
	if s.ProvisionedProductDetail != nil {
		if err := s.ProvisionedProductDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProvisionedProductResponseBodyProvisionedProductDetail struct {
	// The time when the product instance was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-05-23T09:46:27Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task that was last run on the product instance.
	//
	// The task can be one of the following types:
	//
	// 	- LaunchProduct: a task that launches the product.
	//
	// 	- UpdateProvisionedProduct: a task that updates the information about the product instance.
	//
	// 	- TerminateProvisionedProduct: a task that terminates the product instance.
	//
	// example:
	//
	// task-bp1dmg242c****
	LastProvisioningTaskId *string `json:"LastProvisioningTaskId,omitempty" xml:"LastProvisioningTaskId,omitempty"`
	// The ID of the last task that was successfully run on the product instance.
	//
	// The task can be one of the following types:
	//
	// 	- LaunchProduct: a task that launches the product.
	//
	// 	- UpdateProvisionedProduct: a task that updates the information about the product instance.
	//
	// 	- TerminateProvisionedProduct: a task that terminates the product instance.
	//
	// example:
	//
	// task-bp1dmg242c****
	LastSuccessfulProvisioningTaskId *string `json:"LastSuccessfulProvisioningTaskId,omitempty" xml:"LastSuccessfulProvisioningTaskId,omitempty"`
	// The ID of the task that was last run.
	//
	// example:
	//
	// task-bp1dmg242c****
	LastTaskId *string `json:"LastTaskId,omitempty" xml:"LastTaskId,omitempty"`
	// The ID of the RAM entity to which the product instance belongs.
	//
	// example:
	//
	// 24477111603637****
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The type of the Resource Access Management (RAM) entity to which the product instance belongs. Valid values:
	//
	// 	- RamUser: a RAM user
	//
	// 	- RamRole: a RAM role
	//
	// example:
	//
	// RamUser
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
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
	// The name of the product.
	//
	// example:
	//
	// DEMO-Create an ECS instance
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	//
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product instance.
	//
	// example:
	//
	// acs:servicecatalog:cn-hangzhou:146611588617****:provisionedproduct/pp-bp1ddg1n2a****
	ProvisionedProductArn *string `json:"ProvisionedProductArn,omitempty" xml:"ProvisionedProductArn,omitempty"`
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
	// The type of the product instance.
	//
	// The value is fixed as RosStack, which indicates an ROS stack.
	//
	// example:
	//
	// RosStack
	ProvisionedProductType *string `json:"ProvisionedProductType,omitempty" xml:"ProvisionedProductType,omitempty"`
	// The ID of the Resource Orchestration Service (ROS) stack.
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
	// The state of the product instance. Valid values:
	//
	// 	- Available: The product instance was available.
	//
	// 	- UnderChange: The information about the product instance was being changed.
	//
	// 	- Error: An exception occurred on the product instance.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message that is returned for the status of the product instance.
	//
	// > This parameter is returned only when Error is returned for the Status parameter.
	//
	// example:
	//
	// Resource CREATE failed: terraform stack sc-146611588617****-pp-bp1ddg1n2a***	- failure...
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s GetProvisionedProductResponseBodyProvisionedProductDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductResponseBodyProvisionedProductDetail) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetLastProvisioningTaskId() *string {
	return s.LastProvisioningTaskId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetLastSuccessfulProvisioningTaskId() *string {
	return s.LastSuccessfulProvisioningTaskId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetLastTaskId() *string {
	return s.LastTaskId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetOwnerPrincipalId() *string {
	return s.OwnerPrincipalId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetOwnerPrincipalType() *string {
	return s.OwnerPrincipalType
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProductName() *string {
	return s.ProductName
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProvisionedProductArn() *string {
	return s.ProvisionedProductArn
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetProvisionedProductType() *string {
	return s.ProvisionedProductType
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetStackId() *string {
	return s.StackId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetStackRegionId() *string {
	return s.StackRegionId
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetStatus() *string {
	return s.Status
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetCreateTime(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetLastProvisioningTaskId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.LastProvisioningTaskId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetLastSuccessfulProvisioningTaskId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.LastSuccessfulProvisioningTaskId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetLastTaskId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.LastTaskId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetOwnerPrincipalId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.OwnerPrincipalId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetOwnerPrincipalType(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.OwnerPrincipalType = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetPortfolioId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProductId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProductId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProductName(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProductName = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProductVersionId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProductVersionId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProductVersionName(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProductVersionName = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProvisionedProductArn(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProvisionedProductArn = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProvisionedProductId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProvisionedProductId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProvisionedProductName(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProvisionedProductName = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetProvisionedProductType(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.ProvisionedProductType = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetStackId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.StackId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetStackRegionId(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.StackRegionId = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetStatus(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.Status = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) SetStatusMessage(v string) *GetProvisionedProductResponseBodyProvisionedProductDetail {
	s.StatusMessage = &v
	return s
}

func (s *GetProvisionedProductResponseBodyProvisionedProductDetail) Validate() error {
	return dara.Validate(s)
}
