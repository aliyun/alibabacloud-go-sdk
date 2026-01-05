// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductVersionDetails(v []*ListProductVersionsResponseBodyProductVersionDetails) *ListProductVersionsResponseBody
	GetProductVersionDetails() []*ListProductVersionsResponseBodyProductVersionDetails
	SetRequestId(v string) *ListProductVersionsResponseBody
	GetRequestId() *string
}

type ListProductVersionsResponseBody struct {
	// The versions of the product.
	ProductVersionDetails []*ListProductVersionsResponseBodyProductVersionDetails `json:"ProductVersionDetails,omitempty" xml:"ProductVersionDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProductVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBody) GetProductVersionDetails() []*ListProductVersionsResponseBodyProductVersionDetails {
	return s.ProductVersionDetails
}

func (s *ListProductVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductVersionsResponseBody) SetProductVersionDetails(v []*ListProductVersionsResponseBodyProductVersionDetails) *ListProductVersionsResponseBody {
	s.ProductVersionDetails = v
	return s
}

func (s *ListProductVersionsResponseBody) SetRequestId(v string) *ListProductVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductVersionsResponseBody) Validate() error {
	if s.ProductVersionDetails != nil {
		for _, item := range s.ProductVersionDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProductVersionsResponseBodyProductVersionDetails struct {
	// Indicates whether the product version is enabled. Valid values:
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
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ssZ format. The time is displayed in UTC.
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
	// The recommended product version. Valid values:
	//
	// 	- Default (default): No product version is recommended.
	//
	// 	- Recommended: the stable version.
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
	// The product version ID.
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
	// The type of the product template. Valid values:
	//
	// 	- RosTerraformTemplate: the Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// 	- RosStandardTemplate: the standard ROS template.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s ListProductVersionsResponseBodyProductVersionDetails) String() string {
	return dara.Prettify(s)
}

func (s ListProductVersionsResponseBodyProductVersionDetails) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetActive() *bool {
	return s.Active
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetDescription() *string {
	return s.Description
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetGuidance() *string {
	return s.Guidance
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetProductId() *string {
	return s.ProductId
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetActive(v bool) *ListProductVersionsResponseBodyProductVersionDetails {
	s.Active = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetCreateTime(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.CreateTime = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetDescription(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.Description = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetGuidance(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.Guidance = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetProductId(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.ProductId = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetProductVersionId(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.ProductVersionId = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetProductVersionName(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.ProductVersionName = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetTemplateType(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.TemplateType = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) SetTemplateUrl(v string) *ListProductVersionsResponseBodyProductVersionDetails {
	s.TemplateUrl = &v
	return s
}

func (s *ListProductVersionsResponseBodyProductVersionDetails) Validate() error {
	return dara.Validate(s)
}
