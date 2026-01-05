// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameters(v []*LaunchProductRequestParameters) *LaunchProductRequest
	GetParameters() []*LaunchProductRequestParameters
	SetPortfolioId(v string) *LaunchProductRequest
	GetPortfolioId() *string
	SetProductId(v string) *LaunchProductRequest
	GetProductId() *string
	SetProductVersionId(v string) *LaunchProductRequest
	GetProductVersionId() *string
	SetProvisionedProductName(v string) *LaunchProductRequest
	GetProvisionedProductName() *string
	SetStackRegionId(v string) *LaunchProductRequest
	GetStackRegionId() *string
	SetTags(v []*LaunchProductRequestTags) *LaunchProductRequest
	GetTags() []*LaunchProductRequestTags
}

type LaunchProductRequest struct {
	// The input parameters of the template.
	//
	// You can specify up to 200 parameters.
	//
	// > This parameter is optional. If you specify the Parameters parameter, you must specify the ParameterKey and ParameterValue parameters.
	Parameters []*LaunchProductRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	//
	// > If the PortfolioId parameter is not required, you do not need to specify the PortfolioId parameter. If the PortfolioId parameter is required, you must specify the PortfolioId parameter. For more information about how to obtain the value of the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
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
	// The name of the product instance.
	//
	// The value must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEMO-ECS instance
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the region to which the Resource Orchestration Service (ROS) stack belongs.
	//
	// For more information about how to obtain the regions that are supported by ROS, see [DescribeRegions](https://help.aliyun.com/document_detail/131035.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The custom tags that are specified by the end user.
	//
	// Maximum value of N: 20.
	//
	// >
	//
	// 	- The Tags parameter is optional. If you specify the Tags parameter, you must specify the Tags.N.Key and Tags.N.Value parameters.
	//
	// 	- The tag is propagated to each stack resource that supports the tag feature.
	Tags []*LaunchProductRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s LaunchProductRequest) String() string {
	return dara.Prettify(s)
}

func (s LaunchProductRequest) GoString() string {
	return s.String()
}

func (s *LaunchProductRequest) GetParameters() []*LaunchProductRequestParameters {
	return s.Parameters
}

func (s *LaunchProductRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *LaunchProductRequest) GetProductId() *string {
	return s.ProductId
}

func (s *LaunchProductRequest) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *LaunchProductRequest) GetProvisionedProductName() *string {
	return s.ProvisionedProductName
}

func (s *LaunchProductRequest) GetStackRegionId() *string {
	return s.StackRegionId
}

func (s *LaunchProductRequest) GetTags() []*LaunchProductRequestTags {
	return s.Tags
}

func (s *LaunchProductRequest) SetParameters(v []*LaunchProductRequestParameters) *LaunchProductRequest {
	s.Parameters = v
	return s
}

func (s *LaunchProductRequest) SetPortfolioId(v string) *LaunchProductRequest {
	s.PortfolioId = &v
	return s
}

func (s *LaunchProductRequest) SetProductId(v string) *LaunchProductRequest {
	s.ProductId = &v
	return s
}

func (s *LaunchProductRequest) SetProductVersionId(v string) *LaunchProductRequest {
	s.ProductVersionId = &v
	return s
}

func (s *LaunchProductRequest) SetProvisionedProductName(v string) *LaunchProductRequest {
	s.ProvisionedProductName = &v
	return s
}

func (s *LaunchProductRequest) SetStackRegionId(v string) *LaunchProductRequest {
	s.StackRegionId = &v
	return s
}

func (s *LaunchProductRequest) SetTags(v []*LaunchProductRequestTags) *LaunchProductRequest {
	s.Tags = v
	return s
}

func (s *LaunchProductRequest) Validate() error {
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

type LaunchProductRequestParameters struct {
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

func (s LaunchProductRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s LaunchProductRequestParameters) GoString() string {
	return s.String()
}

func (s *LaunchProductRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *LaunchProductRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *LaunchProductRequestParameters) SetParameterKey(v string) *LaunchProductRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *LaunchProductRequestParameters) SetParameterValue(v string) *LaunchProductRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *LaunchProductRequestParameters) Validate() error {
	return dara.Validate(s)
}

type LaunchProductRequestTags struct {
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

func (s LaunchProductRequestTags) String() string {
	return dara.Prettify(s)
}

func (s LaunchProductRequestTags) GoString() string {
	return s.String()
}

func (s *LaunchProductRequestTags) GetKey() *string {
	return s.Key
}

func (s *LaunchProductRequestTags) GetValue() *string {
	return s.Value
}

func (s *LaunchProductRequestTags) SetKey(v string) *LaunchProductRequestTags {
	s.Key = &v
	return s
}

func (s *LaunchProductRequestTags) SetValue(v string) *LaunchProductRequestTags {
	s.Value = &v
	return s
}

func (s *LaunchProductRequestTags) Validate() error {
	return dara.Validate(s)
}
