// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProvisionedProductPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProvisionedProductPlanRequest
	GetDescription() *string
	SetParameters(v []*UpdateProvisionedProductPlanRequestParameters) *UpdateProvisionedProductPlanRequest
	GetParameters() []*UpdateProvisionedProductPlanRequestParameters
	SetPlanId(v string) *UpdateProvisionedProductPlanRequest
	GetPlanId() *string
	SetPortfolioId(v string) *UpdateProvisionedProductPlanRequest
	GetPortfolioId() *string
	SetProductId(v string) *UpdateProvisionedProductPlanRequest
	GetProductId() *string
	SetProductVersionId(v string) *UpdateProvisionedProductPlanRequest
	GetProductVersionId() *string
	SetTags(v []*UpdateProvisionedProductPlanRequestTags) *UpdateProvisionedProductPlanRequest
	GetTags() []*UpdateProvisionedProductPlanRequestTags
}

type UpdateProvisionedProductPlanRequest struct {
	// The description of the plan.
	//
	// example:
	//
	// Create an ECS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// An array that consists of the parameters in the template.
	//
	// Maximum value of N: 200.
	//
	// > If you specify Parameters, you must specify ParameterKey and ParameterValue.
	Parameters []*UpdateProvisionedProductPlanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the plan.
	//
	// This parameter is required.
	//
	// example:
	//
	// plan-bp1jvmdk2k****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the product portfolio.
	//
	// > If the default launch option exists, you do not need to specify PortfolioId. If the default launch option does not exist, you must specify PortfolioId. For more information about how to obtain the value of PortfolioId, see [ListLaunchOptions](~~ListLaunchOptions~~).
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-bp18r7q127****
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// An array that consists of custom tags.
	//
	// Maximum value of N: 20.
	//
	// >
	//
	// 	- If you specify Tags, you must specify Tags.N.Key and Tags.N.Value.
	//
	// 	- The tag of a stack is propagated to each resource that supports the tag feature in the stack.
	Tags []*UpdateProvisionedProductPlanRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateProvisionedProductPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProvisionedProductPlanRequest) GetParameters() []*UpdateProvisionedProductPlanRequestParameters {
	return s.Parameters
}

func (s *UpdateProvisionedProductPlanRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *UpdateProvisionedProductPlanRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *UpdateProvisionedProductPlanRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateProvisionedProductPlanRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *UpdateProvisionedProductPlanRequest) GetTags() []*UpdateProvisionedProductPlanRequestTags {
	return s.Tags
}

func (s *UpdateProvisionedProductPlanRequest) SetDescription(v string) *UpdateProvisionedProductPlanRequest {
	s.Description = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) SetParameters(v []*UpdateProvisionedProductPlanRequestParameters) *UpdateProvisionedProductPlanRequest {
	s.Parameters = v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) SetPlanId(v string) *UpdateProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) SetPortfolioId(v string) *UpdateProvisionedProductPlanRequest {
	s.PortfolioId = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) SetProductId(v string) *UpdateProvisionedProductPlanRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) SetProductVersionId(v string) *UpdateProvisionedProductPlanRequest {
	s.ProductVersionId = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) SetTags(v []*UpdateProvisionedProductPlanRequestTags) *UpdateProvisionedProductPlanRequest {
	s.Tags = v
	return s
}

func (s *UpdateProvisionedProductPlanRequest) Validate() error {
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

type UpdateProvisionedProductPlanRequestParameters struct {
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

func (s UpdateProvisionedProductPlanRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductPlanRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateProvisionedProductPlanRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateProvisionedProductPlanRequestParameters) SetParameterKey(v string) *UpdateProvisionedProductPlanRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequestParameters) SetParameterValue(v string) *UpdateProvisionedProductPlanRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequestParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateProvisionedProductPlanRequestTags struct {
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

func (s UpdateProvisionedProductPlanRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductPlanRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateProvisionedProductPlanRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateProvisionedProductPlanRequestTags) SetKey(v string) *UpdateProvisionedProductPlanRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequestTags) SetValue(v string) *UpdateProvisionedProductPlanRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequestTags) Validate() error {
	return dara.Validate(s)
}
