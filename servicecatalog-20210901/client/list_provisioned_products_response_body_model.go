// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProvisionedProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProvisionedProductsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProvisionedProductsResponseBody
	GetPageSize() *int32
	SetProvisionedProductDetails(v []*ListProvisionedProductsResponseBodyProvisionedProductDetails) *ListProvisionedProductsResponseBody
	GetProvisionedProductDetails() []*ListProvisionedProductsResponseBodyProvisionedProductDetails
	SetRequestId(v string) *ListProvisionedProductsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProvisionedProductsResponseBody
	GetTotalCount() *int32
}

type ListProvisionedProductsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product instances.
	ProvisionedProductDetails []*ListProvisionedProductsResponseBodyProvisionedProductDetails `json:"ProvisionedProductDetails,omitempty" xml:"ProvisionedProductDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProvisionedProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProvisionedProductsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProvisionedProductsResponseBody) GetProvisionedProductDetails() []*ListProvisionedProductsResponseBodyProvisionedProductDetails {
	return s.ProvisionedProductDetails
}

func (s *ListProvisionedProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProvisionedProductsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProvisionedProductsResponseBody) SetPageNumber(v int32) *ListProvisionedProductsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProvisionedProductsResponseBody) SetPageSize(v int32) *ListProvisionedProductsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProvisionedProductsResponseBody) SetProvisionedProductDetails(v []*ListProvisionedProductsResponseBodyProvisionedProductDetails) *ListProvisionedProductsResponseBody {
	s.ProvisionedProductDetails = v
	return s
}

func (s *ListProvisionedProductsResponseBody) SetRequestId(v string) *ListProvisionedProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProvisionedProductsResponseBody) SetTotalCount(v int32) *ListProvisionedProductsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProvisionedProductsResponseBody) Validate() error {
	if s.ProvisionedProductDetails != nil {
		for _, item := range s.ProvisionedProductDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProvisionedProductsResponseBodyProvisionedProductDetails struct {
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

func (s ListProvisionedProductsResponseBodyProvisionedProductDetails) String() string {
	return dara.Prettify(s)
}

func (s ListProvisionedProductsResponseBodyProvisionedProductDetails) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetLastProvisioningTaskId() *string {
	return s.LastProvisioningTaskId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetLastSuccessfulProvisioningTaskId() *string {
	return s.LastSuccessfulProvisioningTaskId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetLastTaskId() *string {
	return s.LastTaskId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetOwnerPrincipalId() *string {
	return s.OwnerPrincipalId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetOwnerPrincipalType() *string {
	return s.OwnerPrincipalType
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProductId() *string {
	return s.ProductId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProductName() *string {
	return s.ProductName
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProvisionedProductArn() *string {
	return s.ProvisionedProductArn
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetProvisionedProductType() *string {
	return s.ProvisionedProductType
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetStackId() *string {
	return s.StackId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetStackRegionId() *string {
	return s.StackRegionId
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetStatus() *string {
	return s.Status
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetCreateTime(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.CreateTime = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetLastProvisioningTaskId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.LastProvisioningTaskId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetLastSuccessfulProvisioningTaskId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.LastSuccessfulProvisioningTaskId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetLastTaskId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.LastTaskId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetOwnerPrincipalId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.OwnerPrincipalId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetOwnerPrincipalType(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.OwnerPrincipalType = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetPortfolioId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.PortfolioId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProductId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProductId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProductName(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProductName = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProductVersionId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProductVersionId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProductVersionName(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProductVersionName = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProvisionedProductArn(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProvisionedProductArn = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProvisionedProductId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProvisionedProductName(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProvisionedProductName = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetProvisionedProductType(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.ProvisionedProductType = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetStackId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.StackId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetStackRegionId(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.StackRegionId = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetStatus(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.Status = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) SetStatusMessage(v string) *ListProvisionedProductsResponseBodyProvisionedProductDetails {
	s.StatusMessage = &v
	return s
}

func (s *ListProvisionedProductsResponseBodyProvisionedProductDetails) Validate() error {
	return dara.Validate(s)
}
