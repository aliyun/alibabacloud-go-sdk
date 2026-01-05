// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProvisionedProductPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProvisionedProductPlanRequest
	GetDescription() *string
	SetOperationType(v string) *CreateProvisionedProductPlanRequest
	GetOperationType() *string
	SetParameters(v []*CreateProvisionedProductPlanRequestParameters) *CreateProvisionedProductPlanRequest
	GetParameters() []*CreateProvisionedProductPlanRequestParameters
	SetPlanName(v string) *CreateProvisionedProductPlanRequest
	GetPlanName() *string
	SetPlanType(v string) *CreateProvisionedProductPlanRequest
	GetPlanType() *string
	SetPortfolioId(v string) *CreateProvisionedProductPlanRequest
	GetPortfolioId() *string
	SetProductId(v string) *CreateProvisionedProductPlanRequest
	GetProductId() *string
	SetProductVersionId(v string) *CreateProvisionedProductPlanRequest
	GetProductVersionId() *string
	SetProvisionedProductName(v string) *CreateProvisionedProductPlanRequest
	GetProvisionedProductName() *string
	SetStackRegionId(v string) *CreateProvisionedProductPlanRequest
	GetStackRegionId() *string
	SetTags(v []*CreateProvisionedProductPlanRequestTags) *CreateProvisionedProductPlanRequest
	GetTags() []*CreateProvisionedProductPlanRequestTags
}

type CreateProvisionedProductPlanRequest struct {
	// The description of the plan.
	//
	// The value must be 1 to 128 characters in length.
	//
	// example:
	//
	// Create an ECS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the operation that you want the plan to perform. Valid values:
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
	// An array that consists of the parameters in the template.
	//
	// You can specify up to 200 parameters.
	//
	// > If you specify Parameters, you must specify ParameterKey and ParameterValue.
	Parameters []*CreateProvisionedProductPlanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The plan name.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEMO-ECS instance
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The plan type.
	//
	// Set the value to Ros, which specifies Resource Orchestration Service (ROS).
	//
	// This parameter is required.
	//
	// example:
	//
	// Ros
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The product portfolio ID.
	//
	// > If PortfolioId is not required, you do not need to specify PortfolioId. If PortfolioId is required, you must specify PortfolioId. For more information about how to obtain the value of PortfolioId, see [ListLaunchOptions](~~ListLaunchOptions~~).
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The product ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The product version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The product instance name.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEMO-ECS instance
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	//
	// For more information about how to obtain the regions that are supported by ROS, see [DescribeRegions](https://help.aliyun.com/document_detail/131035.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// An array that consists of custom tags.
	//
	// Maximum value of N: 20.
	//
	// >
	//
	// 	- If you specify Tags, you must specify Tags.N.Key and Tags.N.Value.
	//
	// 	- The tag of a stack is propagated to each resource that supports the tag feature in the stack.
	Tags []*CreateProvisionedProductPlanRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateProvisionedProductPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProvisionedProductPlanRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *CreateProvisionedProductPlanRequest) GetParameters() []*CreateProvisionedProductPlanRequestParameters {
	return s.Parameters
}

func (s *CreateProvisionedProductPlanRequest) GetPlanName() *string {
	return s.PlanName
}

func (s *CreateProvisionedProductPlanRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *CreateProvisionedProductPlanRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *CreateProvisionedProductPlanRequest) GetProductId() *string {
	return s.ProductId
}

func (s *CreateProvisionedProductPlanRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *CreateProvisionedProductPlanRequest) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *CreateProvisionedProductPlanRequest) GetStackRegionId() *string {
	return s.StackRegionId
}

func (s *CreateProvisionedProductPlanRequest) GetTags() []*CreateProvisionedProductPlanRequestTags {
	return s.Tags
}

func (s *CreateProvisionedProductPlanRequest) SetDescription(v string) *CreateProvisionedProductPlanRequest {
	s.Description = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetOperationType(v string) *CreateProvisionedProductPlanRequest {
	s.OperationType = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetParameters(v []*CreateProvisionedProductPlanRequestParameters) *CreateProvisionedProductPlanRequest {
	s.Parameters = v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetPlanName(v string) *CreateProvisionedProductPlanRequest {
	s.PlanName = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetPlanType(v string) *CreateProvisionedProductPlanRequest {
	s.PlanType = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetPortfolioId(v string) *CreateProvisionedProductPlanRequest {
	s.PortfolioId = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetProductId(v string) *CreateProvisionedProductPlanRequest {
	s.ProductId = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetProductVersionId(v string) *CreateProvisionedProductPlanRequest {
	s.ProductVersionId = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetProvisionedProductName(v string) *CreateProvisionedProductPlanRequest {
	s.ProvisionedProductName = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetStackRegionId(v string) *CreateProvisionedProductPlanRequest {
	s.StackRegionId = &v
	return s
}

func (s *CreateProvisionedProductPlanRequest) SetTags(v []*CreateProvisionedProductPlanRequestTags) *CreateProvisionedProductPlanRequest {
	s.Tags = v
	return s
}

func (s *CreateProvisionedProductPlanRequest) Validate() error {
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

type CreateProvisionedProductPlanRequestParameters struct {
	// The name of the parameter in the template.
	//
	// example:
	//
	// instance_type
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter in the template.
	//
	// example:
	//
	// ecs.s6-c1m1.small
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateProvisionedProductPlanRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateProvisionedProductPlanRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateProvisionedProductPlanRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateProvisionedProductPlanRequestParameters) SetParameterKey(v string) *CreateProvisionedProductPlanRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateProvisionedProductPlanRequestParameters) SetParameterValue(v string) *CreateProvisionedProductPlanRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateProvisionedProductPlanRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateProvisionedProductPlanRequestTags struct {
	// The key of the custom tag.
	//
	// The key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	//
	// The value can be up to 128 characters in length, and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProvisionedProductPlanRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateProvisionedProductPlanRequestTags) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateProvisionedProductPlanRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateProvisionedProductPlanRequestTags) SetKey(v string) *CreateProvisionedProductPlanRequestTags {
	s.Key = &v
	return s
}

func (s *CreateProvisionedProductPlanRequestTags) SetValue(v string) *CreateProvisionedProductPlanRequestTags {
	s.Value = &v
	return s
}

func (s *CreateProvisionedProductPlanRequestTags) Validate() error {
	return dara.Validate(s)
}
