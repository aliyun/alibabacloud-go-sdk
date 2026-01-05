// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductVersionDetail(v *GetProductVersionResponseBodyProductVersionDetail) *GetProductVersionResponseBody
	GetProductVersionDetail() *GetProductVersionResponseBodyProductVersionDetail
	SetRequestId(v string) *GetProductVersionResponseBody
	GetRequestId() *string
}

type GetProductVersionResponseBody struct {
	// The details of the product version.
	ProductVersionDetail *GetProductVersionResponseBodyProductVersionDetail `json:"ProductVersionDetail,omitempty" xml:"ProductVersionDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponseBody) GetProductVersionDetail() *GetProductVersionResponseBodyProductVersionDetail {
	return s.ProductVersionDetail
}

func (s *GetProductVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProductVersionResponseBody) SetProductVersionDetail(v *GetProductVersionResponseBodyProductVersionDetail) *GetProductVersionResponseBody {
	s.ProductVersionDetail = v
	return s
}

func (s *GetProductVersionResponseBody) SetRequestId(v string) *GetProductVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductVersionResponseBody) Validate() error {
	if s.ProductVersionDetail != nil {
		if err := s.ProductVersionDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProductVersionResponseBodyProductVersionDetail struct {
	// Indicates whether the product version is visible to end users. Valid values:
	//
	// - true: The product version is visible to end users. This is the default value.
	//
	// - false: The product version is invisible to end users.
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
	// The recommendation information. Valid values:
	//
	// - Default: No recommendation information is provided. This is the default value.
	//
	// - Recommended: the recommendation version.
	//
	// - Latest: the latest version.
	//
	// - Deprecated: the version that is about to be deprecated.
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
	// The name of the product version.
	//
	// example:
	//
	// 1.0
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the product template. Valid values:
	//
	// 	- RosTerraformTemplate: The Terraform template that is supported by Resource Orchestration Service (ROS).
	//
	// 	- RosStandardTemplate: The standard ROS template.
	//
	// example:
	//
	// RosTerraformTemplate
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s GetProductVersionResponseBodyProductVersionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetProductVersionResponseBodyProductVersionDetail) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetActive() *bool {
	return s.Active
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetDescription() *string {
	return s.Description
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetGuidance() *string {
	return s.Guidance
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetProductVersionId() *string {
	return s.ProductVersionId
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetProductVersionName() *string {
	return s.ProductVersionName
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetProductVersionResponseBodyProductVersionDetail) GetTemplateUrl() *string {
	return s.TemplateUrl
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetActive(v bool) *GetProductVersionResponseBodyProductVersionDetail {
	s.Active = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetCreateTime(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetDescription(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.Description = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetGuidance(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.Guidance = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetProductId(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.ProductId = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetProductVersionId(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.ProductVersionId = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetProductVersionName(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.ProductVersionName = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetTemplateType(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.TemplateType = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) SetTemplateUrl(v string) *GetProductVersionResponseBodyProductVersionDetail {
	s.TemplateUrl = &v
	return s
}

func (s *GetProductVersionResponseBodyProductVersionDetail) Validate() error {
	return dara.Validate(s)
}
