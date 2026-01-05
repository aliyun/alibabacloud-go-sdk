// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProvisionedProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v []*UpdateProvisionedProductRequestParameters) *UpdateProvisionedProductRequest
	GetParameters() []*UpdateProvisionedProductRequestParameters
	SetPortfolioId(v string) *UpdateProvisionedProductRequest
	GetPortfolioId() *string
	SetProductId(v string) *UpdateProvisionedProductRequest
	GetProductId() *string
	SetProductVersionId(v string) *UpdateProvisionedProductRequest
	GetProductVersionId() *string
	SetProvisionedProductId(v string) *UpdateProvisionedProductRequest
	GetProvisionedProductId() *string
	SetTags(v []*UpdateProvisionedProductRequestTags) *UpdateProvisionedProductRequest
	GetTags() []*UpdateProvisionedProductRequestTags
}

type UpdateProvisionedProductRequest struct {
	// The input parameters of the template.
	//
	// You can specify up to 200 parameters.
	//
	// > - This parameter is optional. If you specify the Parameters parameter, you must specify the ParameterKey and ParameterValue parameters.
	//
	// > - If the values of the ProductVersionId and Parameters parameters are not changed, you are not allowed to update the information about the product instance.
	Parameters []*UpdateProvisionedProductRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	//
	// > The PortfolioId parameter is not required if the default launch option exists. The PortfolioId parameter is required if the default launch option does not exist. For more information about how to obtain the value of the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
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
	// > If the values of the ProductVersionId and Parameters parameters are not changed, the information about the product instance cannot be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv-bp15e79d26****
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the product instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The input custom tags.
	//
	// Maximum value of N: 20.
	//
	// > - The Tags parameter is optional. If you need to specify the Tags parameter, you must specify the Tags.N.Key and Tags.N.Value parameters.
	//
	// > - The tag is propagated to each stack resource that supports the tag feature.
	Tags []*UpdateProvisionedProductRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateProvisionedProductRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductRequest) GetParameters() []*UpdateProvisionedProductRequestParameters {
	return s.Parameters
}

func (s *UpdateProvisionedProductRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *UpdateProvisionedProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *UpdateProvisionedProductRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *UpdateProvisionedProductRequest) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *UpdateProvisionedProductRequest) GetTags() []*UpdateProvisionedProductRequestTags {
	return s.Tags
}

func (s *UpdateProvisionedProductRequest) SetParameters(v []*UpdateProvisionedProductRequestParameters) *UpdateProvisionedProductRequest {
	s.Parameters = v
	return s
}

func (s *UpdateProvisionedProductRequest) SetPortfolioId(v string) *UpdateProvisionedProductRequest {
	s.PortfolioId = &v
	return s
}

func (s *UpdateProvisionedProductRequest) SetProductId(v string) *UpdateProvisionedProductRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateProvisionedProductRequest) SetProductVersionId(v string) *UpdateProvisionedProductRequest {
	s.ProductVersionId = &v
	return s
}

func (s *UpdateProvisionedProductRequest) SetProvisionedProductId(v string) *UpdateProvisionedProductRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *UpdateProvisionedProductRequest) SetTags(v []*UpdateProvisionedProductRequestTags) *UpdateProvisionedProductRequest {
	s.Tags = v
	return s
}

func (s *UpdateProvisionedProductRequest) Validate() error {
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

type UpdateProvisionedProductRequestParameters struct {
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

func (s UpdateProvisionedProductRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateProvisionedProductRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateProvisionedProductRequestParameters) SetParameterKey(v string) *UpdateProvisionedProductRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateProvisionedProductRequestParameters) SetParameterValue(v string) *UpdateProvisionedProductRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateProvisionedProductRequestParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateProvisionedProductRequestTags struct {
	// The tag key of the custom tag.
	//
	// The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the custom tag.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateProvisionedProductRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateProvisionedProductRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateProvisionedProductRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateProvisionedProductRequestTags) SetKey(v string) *UpdateProvisionedProductRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateProvisionedProductRequestTags) SetValue(v string) *UpdateProvisionedProductRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateProvisionedProductRequestTags) Validate() error {
	return dara.Validate(s)
}
