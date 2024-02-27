// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApproveProvisionedProductPlanRequest struct {
	// The review action. Valid values:
	//
	// *   Approve
	// *   Reject
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// The review description.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s ApproveProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *ApproveProvisionedProductPlanRequest) SetApprovalAction(v string) *ApproveProvisionedProductPlanRequest {
	s.ApprovalAction = &v
	return s
}

func (s *ApproveProvisionedProductPlanRequest) SetComment(v string) *ApproveProvisionedProductPlanRequest {
	s.Comment = &v
	return s
}

func (s *ApproveProvisionedProductPlanRequest) SetPlanId(v string) *ApproveProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

type ApproveProvisionedProductPlanResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveProvisionedProductPlanResponseBody) SetRequestId(v string) *ApproveProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

type ApproveProvisionedProductPlanResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *ApproveProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *ApproveProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *ApproveProvisionedProductPlanResponse) SetStatusCode(v int32) *ApproveProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveProvisionedProductPlanResponse) SetBody(v *ApproveProvisionedProductPlanResponseBody) *ApproveProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type AssociatePrincipalWithPortfolioRequest struct {
	PortfolioId   *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	PrincipalId   *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s AssociatePrincipalWithPortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociatePrincipalWithPortfolioRequest) GoString() string {
	return s.String()
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPortfolioId(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPrincipalId(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PrincipalId = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPrincipalType(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PrincipalType = &v
	return s
}

type AssociatePrincipalWithPortfolioResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociatePrincipalWithPortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociatePrincipalWithPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *AssociatePrincipalWithPortfolioResponseBody) SetRequestId(v string) *AssociatePrincipalWithPortfolioResponseBody {
	s.RequestId = &v
	return s
}

type AssociatePrincipalWithPortfolioResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociatePrincipalWithPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociatePrincipalWithPortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociatePrincipalWithPortfolioResponse) GoString() string {
	return s.String()
}

func (s *AssociatePrincipalWithPortfolioResponse) SetHeaders(v map[string]*string) *AssociatePrincipalWithPortfolioResponse {
	s.Headers = v
	return s
}

func (s *AssociatePrincipalWithPortfolioResponse) SetStatusCode(v int32) *AssociatePrincipalWithPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioResponse) SetBody(v *AssociatePrincipalWithPortfolioResponseBody) *AssociatePrincipalWithPortfolioResponse {
	s.Body = v
	return s
}

type AssociateProductWithPortfolioRequest struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s AssociateProductWithPortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateProductWithPortfolioRequest) GoString() string {
	return s.String()
}

func (s *AssociateProductWithPortfolioRequest) SetPortfolioId(v string) *AssociateProductWithPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *AssociateProductWithPortfolioRequest) SetProductId(v string) *AssociateProductWithPortfolioRequest {
	s.ProductId = &v
	return s
}

type AssociateProductWithPortfolioResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateProductWithPortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateProductWithPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateProductWithPortfolioResponseBody) SetRequestId(v string) *AssociateProductWithPortfolioResponseBody {
	s.RequestId = &v
	return s
}

type AssociateProductWithPortfolioResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateProductWithPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateProductWithPortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateProductWithPortfolioResponse) GoString() string {
	return s.String()
}

func (s *AssociateProductWithPortfolioResponse) SetHeaders(v map[string]*string) *AssociateProductWithPortfolioResponse {
	s.Headers = v
	return s
}

func (s *AssociateProductWithPortfolioResponse) SetStatusCode(v int32) *AssociateProductWithPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateProductWithPortfolioResponse) SetBody(v *AssociateProductWithPortfolioResponseBody) *AssociateProductWithPortfolioResponse {
	s.Body = v
	return s
}

type AssociateTagOptionWithResourceRequest struct {
	// The ID of the resource with which the tag option is associated.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s AssociateTagOptionWithResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s AssociateTagOptionWithResourceRequest) GoString() string {
	return s.String()
}

func (s *AssociateTagOptionWithResourceRequest) SetResourceId(v string) *AssociateTagOptionWithResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *AssociateTagOptionWithResourceRequest) SetTagOptionId(v string) *AssociateTagOptionWithResourceRequest {
	s.TagOptionId = &v
	return s
}

type AssociateTagOptionWithResourceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateTagOptionWithResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AssociateTagOptionWithResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateTagOptionWithResourceResponseBody) SetRequestId(v string) *AssociateTagOptionWithResourceResponseBody {
	s.RequestId = &v
	return s
}

type AssociateTagOptionWithResourceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateTagOptionWithResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateTagOptionWithResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s AssociateTagOptionWithResourceResponse) GoString() string {
	return s.String()
}

func (s *AssociateTagOptionWithResourceResponse) SetHeaders(v map[string]*string) *AssociateTagOptionWithResourceResponse {
	s.Headers = v
	return s
}

func (s *AssociateTagOptionWithResourceResponse) SetStatusCode(v int32) *AssociateTagOptionWithResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateTagOptionWithResourceResponse) SetBody(v *AssociateTagOptionWithResourceResponseBody) *AssociateTagOptionWithResourceResponse {
	s.Body = v
	return s
}

type CancelProvisionedProductPlanRequest struct {
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s CancelProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *CancelProvisionedProductPlanRequest) SetPlanId(v string) *CancelProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

type CancelProvisionedProductPlanResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CancelProvisionedProductPlanResponseBody) SetRequestId(v string) *CancelProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

type CancelProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *CancelProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *CancelProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *CancelProvisionedProductPlanResponse) SetStatusCode(v int32) *CancelProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelProvisionedProductPlanResponse) SetBody(v *CancelProvisionedProductPlanResponseBody) *CancelProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type CopyProductRequest struct {
	// The Alibaba Cloud Resource Name (ARN) of the source product.
	//
	// > The source product can belong to the current account or belong to a product portfolio that is shared by another account.
	SourceProductArn *string `json:"SourceProductArn,omitempty" xml:"SourceProductArn,omitempty"`
	// The name of the destination product.
	TargetProductName *string `json:"TargetProductName,omitempty" xml:"TargetProductName,omitempty"`
}

func (s CopyProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyProductRequest) GoString() string {
	return s.String()
}

func (s *CopyProductRequest) SetSourceProductArn(v string) *CopyProductRequest {
	s.SourceProductArn = &v
	return s
}

func (s *CopyProductRequest) SetTargetProductName(v string) *CopyProductRequest {
	s.TargetProductName = &v
	return s
}

type CopyProductResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyProductResponseBody) GoString() string {
	return s.String()
}

func (s *CopyProductResponseBody) SetRequestId(v string) *CopyProductResponseBody {
	s.RequestId = &v
	return s
}

type CopyProductResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyProductResponse) GoString() string {
	return s.String()
}

func (s *CopyProductResponse) SetHeaders(v map[string]*string) *CopyProductResponse {
	s.Headers = v
	return s
}

func (s *CopyProductResponse) SetStatusCode(v int32) *CopyProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyProductResponse) SetBody(v *CopyProductResponseBody) *CopyProductResponse {
	s.Body = v
	return s
}

type CreateConstraintRequest struct {
	// The configuration of the constraint.
	//
	// Format: { "LocalRoleName": "\<role_name>" }.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The type of the constraint.
	//
	// The value is fixed as Launch, which specifies the launch constraint.
	ConstraintType *string `json:"ConstraintType,omitempty" xml:"ConstraintType,omitempty"`
	// The description of the constraint.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the product portfolio to which the constraint belongs.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product for which the constraint is created.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s CreateConstraintRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConstraintRequest) GoString() string {
	return s.String()
}

func (s *CreateConstraintRequest) SetConfig(v string) *CreateConstraintRequest {
	s.Config = &v
	return s
}

func (s *CreateConstraintRequest) SetConstraintType(v string) *CreateConstraintRequest {
	s.ConstraintType = &v
	return s
}

func (s *CreateConstraintRequest) SetDescription(v string) *CreateConstraintRequest {
	s.Description = &v
	return s
}

func (s *CreateConstraintRequest) SetPortfolioId(v string) *CreateConstraintRequest {
	s.PortfolioId = &v
	return s
}

func (s *CreateConstraintRequest) SetProductId(v string) *CreateConstraintRequest {
	s.ProductId = &v
	return s
}

type CreateConstraintResponseBody struct {
	// The ID of the constraint.
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConstraintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConstraintResponseBody) SetConstraintId(v string) *CreateConstraintResponseBody {
	s.ConstraintId = &v
	return s
}

func (s *CreateConstraintResponseBody) SetRequestId(v string) *CreateConstraintResponseBody {
	s.RequestId = &v
	return s
}

type CreateConstraintResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConstraintResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConstraintResponse) GoString() string {
	return s.String()
}

func (s *CreateConstraintResponse) SetHeaders(v map[string]*string) *CreateConstraintResponse {
	s.Headers = v
	return s
}

func (s *CreateConstraintResponse) SetStatusCode(v int32) *CreateConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConstraintResponse) SetBody(v *CreateConstraintResponseBody) *CreateConstraintResponse {
	s.Body = v
	return s
}

type CreatePortfolioRequest struct {
	// The description of the product portfolio.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the product portfolio.
	//
	// The value must be 1 to 128 characters in length.
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// The provider of the product portfolio.
	//
	// The value must be 1 to 128 characters in length.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s CreatePortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePortfolioRequest) GoString() string {
	return s.String()
}

func (s *CreatePortfolioRequest) SetDescription(v string) *CreatePortfolioRequest {
	s.Description = &v
	return s
}

func (s *CreatePortfolioRequest) SetPortfolioName(v string) *CreatePortfolioRequest {
	s.PortfolioName = &v
	return s
}

func (s *CreatePortfolioRequest) SetProviderName(v string) *CreatePortfolioRequest {
	s.ProviderName = &v
	return s
}

type CreatePortfolioResponseBody struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePortfolioResponseBody) SetPortfolioId(v string) *CreatePortfolioResponseBody {
	s.PortfolioId = &v
	return s
}

func (s *CreatePortfolioResponseBody) SetRequestId(v string) *CreatePortfolioResponseBody {
	s.RequestId = &v
	return s
}

type CreatePortfolioResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePortfolioResponse) GoString() string {
	return s.String()
}

func (s *CreatePortfolioResponse) SetHeaders(v map[string]*string) *CreatePortfolioResponse {
	s.Headers = v
	return s
}

func (s *CreatePortfolioResponse) SetStatusCode(v int32) *CreatePortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePortfolioResponse) SetBody(v *CreatePortfolioResponseBody) *CreatePortfolioResponse {
	s.Body = v
	return s
}

type CreateProductRequest struct {
	// The description of the product.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the product.
	//
	// The value must be 1 to 128 characters in length.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// Set the value to Ros, which specifies Resource Orchestration Service (ROS).
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The information about the product version.
	ProductVersionParameters *CreateProductRequestProductVersionParameters `json:"ProductVersionParameters,omitempty" xml:"ProductVersionParameters,omitempty" type:"Struct"`
	// The provider of the product.
	//
	// The value must be 1 to 128 characters in length.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequest) GoString() string {
	return s.String()
}

func (s *CreateProductRequest) SetDescription(v string) *CreateProductRequest {
	s.Description = &v
	return s
}

func (s *CreateProductRequest) SetProductName(v string) *CreateProductRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductRequest) SetProductType(v string) *CreateProductRequest {
	s.ProductType = &v
	return s
}

func (s *CreateProductRequest) SetProductVersionParameters(v *CreateProductRequestProductVersionParameters) *CreateProductRequest {
	s.ProductVersionParameters = v
	return s
}

func (s *CreateProductRequest) SetProviderName(v string) *CreateProductRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateProductRequest) SetTemplateType(v string) *CreateProductRequest {
	s.TemplateType = &v
	return s
}

type CreateProductRequestProductVersionParameters struct {
	// Specifies whether to enable the product version. Valid values:
	//
	// *   true: enables the product version. This is the default value.
	// *   false: disables the product version.
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the product version.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// *   Default: No recommendation information is provided. This is the default value.
	// *   Recommended: the recommended version.
	// *   Latest: the latest version.
	// *   Deprecated: the version that is about to be discontinued.
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The name of the product version.
	//
	// The value must be 1 to 128 characters in length.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the template.
	//
	// Set the value to RosTerraformTemplate, which specifies the Terraform template that is supported by ROS.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	//
	// For more information about how to obtain the URL of a template, see [CreateTemplate](~~CreateTemplate~~).
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s CreateProductRequestProductVersionParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateProductRequestProductVersionParameters) GoString() string {
	return s.String()
}

func (s *CreateProductRequestProductVersionParameters) SetActive(v bool) *CreateProductRequestProductVersionParameters {
	s.Active = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetDescription(v string) *CreateProductRequestProductVersionParameters {
	s.Description = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetGuidance(v string) *CreateProductRequestProductVersionParameters {
	s.Guidance = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetProductVersionName(v string) *CreateProductRequestProductVersionParameters {
	s.ProductVersionName = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetTemplateType(v string) *CreateProductRequestProductVersionParameters {
	s.TemplateType = &v
	return s
}

func (s *CreateProductRequestProductVersionParameters) SetTemplateUrl(v string) *CreateProductRequestProductVersionParameters {
	s.TemplateUrl = &v
	return s
}

type CreateProductShrinkRequest struct {
	// The description of the product.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the product.
	//
	// The value must be 1 to 128 characters in length.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// Set the value to Ros, which specifies Resource Orchestration Service (ROS).
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The information about the product version.
	ProductVersionParametersShrink *string `json:"ProductVersionParameters,omitempty" xml:"ProductVersionParameters,omitempty"`
	// The provider of the product.
	//
	// The value must be 1 to 128 characters in length.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateProductShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProductShrinkRequest) SetDescription(v string) *CreateProductShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProductName(v string) *CreateProductShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProductType(v string) *CreateProductShrinkRequest {
	s.ProductType = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProductVersionParametersShrink(v string) *CreateProductShrinkRequest {
	s.ProductVersionParametersShrink = &v
	return s
}

func (s *CreateProductShrinkRequest) SetProviderName(v string) *CreateProductShrinkRequest {
	s.ProviderName = &v
	return s
}

func (s *CreateProductShrinkRequest) SetTemplateType(v string) *CreateProductShrinkRequest {
	s.TemplateType = &v
	return s
}

type CreateProductResponseBody struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductResponseBody) SetProductId(v string) *CreateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *CreateProductResponseBody) SetProductVersionId(v string) *CreateProductResponseBody {
	s.ProductVersionId = &v
	return s
}

func (s *CreateProductResponseBody) SetRequestId(v string) *CreateProductResponseBody {
	s.RequestId = &v
	return s
}

type CreateProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductResponse) GoString() string {
	return s.String()
}

func (s *CreateProductResponse) SetHeaders(v map[string]*string) *CreateProductResponse {
	s.Headers = v
	return s
}

func (s *CreateProductResponse) SetStatusCode(v int32) *CreateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductResponse) SetBody(v *CreateProductResponseBody) *CreateProductResponse {
	s.Body = v
	return s
}

type CreateProductVersionRequest struct {
	// Specifies whether the product version is active. Valid values:
	//
	// *   true: The product version is active. This is the default value.
	// *   false: The product version is inactive.
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the product version.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// *   Default: No recommendation information is provided. This is the default value.
	// *   Recommended: the recommendation version.
	// *   Latest: the latest version.
	// *   Deprecated: the version that is about to be discontinued.
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product to which the product version belongs.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product version.
	//
	// The value must be 1 to 128 characters in length.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the template.
	//
	// The value is fixed as RosTerraformTemplate, which specifies the Terraform template that is supported by Resource Orchestration Service (ROS).
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	//
	// For more information about how to obtain the URL of a template, see [CreateTemplate](~~CreateTemplate~~).
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s CreateProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateProductVersionRequest) SetActive(v bool) *CreateProductVersionRequest {
	s.Active = &v
	return s
}

func (s *CreateProductVersionRequest) SetDescription(v string) *CreateProductVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateProductVersionRequest) SetGuidance(v string) *CreateProductVersionRequest {
	s.Guidance = &v
	return s
}

func (s *CreateProductVersionRequest) SetProductId(v string) *CreateProductVersionRequest {
	s.ProductId = &v
	return s
}

func (s *CreateProductVersionRequest) SetProductVersionName(v string) *CreateProductVersionRequest {
	s.ProductVersionName = &v
	return s
}

func (s *CreateProductVersionRequest) SetTemplateType(v string) *CreateProductVersionRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateProductVersionRequest) SetTemplateUrl(v string) *CreateProductVersionRequest {
	s.TemplateUrl = &v
	return s
}

type CreateProductVersionResponseBody struct {
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponseBody) SetProductVersionId(v string) *CreateProductVersionResponseBody {
	s.ProductVersionId = &v
	return s
}

func (s *CreateProductVersionResponseBody) SetRequestId(v string) *CreateProductVersionResponseBody {
	s.RequestId = &v
	return s
}

type CreateProductVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateProductVersionResponse) SetHeaders(v map[string]*string) *CreateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *CreateProductVersionResponse) SetStatusCode(v int32) *CreateProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProductVersionResponse) SetBody(v *CreateProductVersionResponseBody) *CreateProductVersionResponse {
	s.Body = v
	return s
}

type CreateProvisionedProductPlanRequest struct {
	// The description of the plan.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the operation that you want the plan to perform. Valid values:
	//
	// *   LaunchProduct: launches the product. This is the default value.
	// *   UpdateProvisionedProduct: updates the information about the product instance.
	// *   TerminateProvisionedProduct: terminates the product instance.
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
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The plan type.
	//
	// Set the value to Ros, which specifies Resource Orchestration Service (ROS).
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The product portfolio ID.
	//
	// > If PortfolioId is not required, you do not need to specify PortfolioId. If PortfolioId is required, you must specify PortfolioId. For more information about how to obtain the value of PortfolioId, see [ListLaunchOptions](~~ListLaunchOptions~~).
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The product ID.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The product version ID.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The product instance name.
	//
	// The value must be 1 to 128 characters in length.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	//
	// For more information about how to obtain the regions that are supported by ROS, see [DescribeRegions](~~131035~~).
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// An array that consists of custom tags.
	//
	// Maximum value of N: 20.
	//
	// >
	// *   If you specify Tags, you must specify Tags.N.Key and Tags.N.Value.
	// *   The tag of a stack is propagated to each resource that supports the tag feature in the stack.
	Tags []*CreateProvisionedProductPlanRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProvisionedProductPlanRequest) GoString() string {
	return s.String()
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

type CreateProvisionedProductPlanRequestParameters struct {
	// The name of the parameter in the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter in the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateProvisionedProductPlanRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateProvisionedProductPlanRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanRequestParameters) SetParameterKey(v string) *CreateProvisionedProductPlanRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateProvisionedProductPlanRequestParameters) SetParameterValue(v string) *CreateProvisionedProductPlanRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateProvisionedProductPlanRequestTags struct {
	// The key of the custom tag.
	//
	// The key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	//
	// The value can be up to 128 characters in length, and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProvisionedProductPlanRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateProvisionedProductPlanRequestTags) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanRequestTags) SetKey(v string) *CreateProvisionedProductPlanRequestTags {
	s.Key = &v
	return s
}

func (s *CreateProvisionedProductPlanRequestTags) SetValue(v string) *CreateProvisionedProductPlanRequestTags {
	s.Value = &v
	return s
}

type CreateProvisionedProductPlanResponseBody struct {
	// The plan ID.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The product instance ID.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanResponseBody) SetPlanId(v string) *CreateProvisionedProductPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *CreateProvisionedProductPlanResponseBody) SetProvisionedProductId(v string) *CreateProvisionedProductPlanResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *CreateProvisionedProductPlanResponseBody) SetRequestId(v string) *CreateProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

type CreateProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *CreateProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateProvisionedProductPlanResponse) SetStatusCode(v int32) *CreateProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProvisionedProductPlanResponse) SetBody(v *CreateProvisionedProductPlanResponseBody) *CreateProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type CreateTagOptionRequest struct {
	// The key of the tag option.
	//
	// The key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag option.
	//
	// The value can be up to 128 characters in length. It cannot start with `acs:`and cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTagOptionRequest) GoString() string {
	return s.String()
}

func (s *CreateTagOptionRequest) SetKey(v string) *CreateTagOptionRequest {
	s.Key = &v
	return s
}

func (s *CreateTagOptionRequest) SetValue(v string) *CreateTagOptionRequest {
	s.Value = &v
	return s
}

type CreateTagOptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the tag option.
	TagOptionDetail *CreateTagOptionResponseBodyTagOptionDetail `json:"TagOptionDetail,omitempty" xml:"TagOptionDetail,omitempty" type:"Struct"`
}

func (s CreateTagOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTagOptionResponseBody) SetRequestId(v string) *CreateTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTagOptionResponseBody) SetTagOptionDetail(v *CreateTagOptionResponseBodyTagOptionDetail) *CreateTagOptionResponseBody {
	s.TagOptionDetail = v
	return s
}

type CreateTagOptionResponseBodyTagOptionDetail struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// *   true (default)
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	//
	// The key must be 1 to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the Alibaba Cloud account to which the tag option belongs.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// The value must be 1 to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTagOptionResponseBodyTagOptionDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateTagOptionResponseBodyTagOptionDetail) GoString() string {
	return s.String()
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetActive(v bool) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Active = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetKey(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Key = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetOwner(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Owner = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetTagOptionId(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.TagOptionId = &v
	return s
}

func (s *CreateTagOptionResponseBodyTagOptionDetail) SetValue(v string) *CreateTagOptionResponseBodyTagOptionDetail {
	s.Value = &v
	return s
}

type CreateTagOptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTagOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTagOptionResponse) GoString() string {
	return s.String()
}

func (s *CreateTagOptionResponse) SetHeaders(v map[string]*string) *CreateTagOptionResponse {
	s.Headers = v
	return s
}

func (s *CreateTagOptionResponse) SetStatusCode(v int32) *CreateTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTagOptionResponse) SetBody(v *CreateTagOptionResponseBody) *CreateTagOptionResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// The content of the template.
	//
	// For more information about the template syntax, see [Structure of Terraform templates](~~184397~~).
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The type of the template.
	//
	// Set the value to RosTerraformTemplate, which specifies the Terraform template that is supported by Resource Orchestration Service (ROS).
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The variable settings of the Terraform template. You can configure the variables in a structured manner. Service Catalog applies the variable settings to the template.
	//
	// > The variables must be defined in the Terraform template.
	TerraformVariables []*CreateTemplateRequestTerraformVariables `json:"TerraformVariables,omitempty" xml:"TerraformVariables,omitempty" type:"Repeated"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetTemplateBody(v string) *CreateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateType(v string) *CreateTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateRequest) SetTerraformVariables(v []*CreateTemplateRequestTerraformVariables) *CreateTemplateRequest {
	s.TerraformVariables = v
	return s
}

type CreateTemplateRequestTerraformVariables struct {
	// The description of the variable.
	//
	// For more information about the format of variable descriptions, see [Methods and suggestions for Terraform code development](~~322216~~).
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the variable.
	VariableName *string `json:"VariableName,omitempty" xml:"VariableName,omitempty"`
}

func (s CreateTemplateRequestTerraformVariables) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequestTerraformVariables) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestTerraformVariables) SetDescription(v string) *CreateTemplateRequestTerraformVariables {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequestTerraformVariables) SetVariableName(v string) *CreateTemplateRequestTerraformVariables {
	s.VariableName = &v
	return s
}

type CreateTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The URL of the template.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateUrl(v string) *CreateTemplateResponseBody {
	s.TemplateUrl = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteConstraintRequest struct {
	// The ID of the constraint.
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
}

func (s DeleteConstraintRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConstraintRequest) GoString() string {
	return s.String()
}

func (s *DeleteConstraintRequest) SetConstraintId(v string) *DeleteConstraintRequest {
	s.ConstraintId = &v
	return s
}

type DeleteConstraintResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConstraintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConstraintResponseBody) SetRequestId(v string) *DeleteConstraintResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConstraintResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConstraintResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConstraintResponse) GoString() string {
	return s.String()
}

func (s *DeleteConstraintResponse) SetHeaders(v map[string]*string) *DeleteConstraintResponse {
	s.Headers = v
	return s
}

func (s *DeleteConstraintResponse) SetStatusCode(v int32) *DeleteConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConstraintResponse) SetBody(v *DeleteConstraintResponseBody) *DeleteConstraintResponse {
	s.Body = v
	return s
}

type DeletePortfolioRequest struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
}

func (s DeletePortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePortfolioRequest) GoString() string {
	return s.String()
}

func (s *DeletePortfolioRequest) SetPortfolioId(v string) *DeletePortfolioRequest {
	s.PortfolioId = &v
	return s
}

type DeletePortfolioResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePortfolioResponseBody) SetRequestId(v string) *DeletePortfolioResponseBody {
	s.RequestId = &v
	return s
}

type DeletePortfolioResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePortfolioResponse) GoString() string {
	return s.String()
}

func (s *DeletePortfolioResponse) SetHeaders(v map[string]*string) *DeletePortfolioResponse {
	s.Headers = v
	return s
}

func (s *DeletePortfolioResponse) SetStatusCode(v int32) *DeletePortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePortfolioResponse) SetBody(v *DeletePortfolioResponseBody) *DeletePortfolioResponse {
	s.Body = v
	return s
}

type DeleteProductRequest struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DeleteProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) SetProductId(v string) *DeleteProductRequest {
	s.ProductId = &v
	return s
}

type DeleteProductResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductResponseBody) SetRequestId(v string) *DeleteProductResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductResponse) SetHeaders(v map[string]*string) *DeleteProductResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductResponse) SetStatusCode(v int32) *DeleteProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductResponse) SetBody(v *DeleteProductResponseBody) *DeleteProductResponse {
	s.Body = v
	return s
}

type DeleteProductVersionRequest struct {
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
}

func (s DeleteProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionRequest) SetProductVersionId(v string) *DeleteProductVersionRequest {
	s.ProductVersionId = &v
	return s
}

type DeleteProductVersionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionResponseBody) SetRequestId(v string) *DeleteProductVersionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProductVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProductVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteProductVersionResponse) SetHeaders(v map[string]*string) *DeleteProductVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteProductVersionResponse) SetStatusCode(v int32) *DeleteProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProductVersionResponse) SetBody(v *DeleteProductVersionResponseBody) *DeleteProductVersionResponse {
	s.Body = v
	return s
}

type DeleteProvisionedProductPlanRequest struct {
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteProvisionedProductPlanRequest) SetPlanId(v string) *DeleteProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

type DeleteProvisionedProductPlanResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProvisionedProductPlanResponseBody) SetRequestId(v string) *DeleteProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

type DeleteProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *DeleteProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteProvisionedProductPlanResponse) SetStatusCode(v int32) *DeleteProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProvisionedProductPlanResponse) SetBody(v *DeleteProvisionedProductPlanResponseBody) *DeleteProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type DeleteTagOptionRequest struct {
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s DeleteTagOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagOptionRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagOptionRequest) SetTagOptionId(v string) *DeleteTagOptionRequest {
	s.TagOptionId = &v
	return s
}

type DeleteTagOptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagOptionResponseBody) SetRequestId(v string) *DeleteTagOptionResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTagOptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTagOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTagOptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteTagOptionResponse) SetHeaders(v map[string]*string) *DeleteTagOptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteTagOptionResponse) SetStatusCode(v int32) *DeleteTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTagOptionResponse) SetBody(v *DeleteTagOptionResponseBody) *DeleteTagOptionResponse {
	s.Body = v
	return s
}

type DisAssociateTagOptionFromResourceRequest struct {
	// The ID of the resource with which the tag option is associated.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s DisAssociateTagOptionFromResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DisAssociateTagOptionFromResourceRequest) GoString() string {
	return s.String()
}

func (s *DisAssociateTagOptionFromResourceRequest) SetResourceId(v string) *DisAssociateTagOptionFromResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DisAssociateTagOptionFromResourceRequest) SetTagOptionId(v string) *DisAssociateTagOptionFromResourceRequest {
	s.TagOptionId = &v
	return s
}

type DisAssociateTagOptionFromResourceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisAssociateTagOptionFromResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisAssociateTagOptionFromResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DisAssociateTagOptionFromResourceResponseBody) SetRequestId(v string) *DisAssociateTagOptionFromResourceResponseBody {
	s.RequestId = &v
	return s
}

type DisAssociateTagOptionFromResourceResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisAssociateTagOptionFromResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisAssociateTagOptionFromResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisAssociateTagOptionFromResourceResponse) GoString() string {
	return s.String()
}

func (s *DisAssociateTagOptionFromResourceResponse) SetHeaders(v map[string]*string) *DisAssociateTagOptionFromResourceResponse {
	s.Headers = v
	return s
}

func (s *DisAssociateTagOptionFromResourceResponse) SetStatusCode(v int32) *DisAssociateTagOptionFromResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisAssociateTagOptionFromResourceResponse) SetBody(v *DisAssociateTagOptionFromResourceResponseBody) *DisAssociateTagOptionFromResourceResponse {
	s.Body = v
	return s
}

type DisassociatePrincipalFromPortfolioRequest struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the RAM entity.
	//
	// For more information about how to obtain the ID of a RAM user, see [GetUser](~~28681~~).
	//
	// For more information about how to obtain the ID of a RAM role, see [GetRole](~~28711~~).
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the Resource Access Management (RAM) entity. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s DisassociatePrincipalFromPortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociatePrincipalFromPortfolioRequest) GoString() string {
	return s.String()
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPortfolioId(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPrincipalId(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PrincipalId = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPrincipalType(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PrincipalType = &v
	return s
}

type DisassociatePrincipalFromPortfolioResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociatePrincipalFromPortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociatePrincipalFromPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociatePrincipalFromPortfolioResponseBody) SetRequestId(v string) *DisassociatePrincipalFromPortfolioResponseBody {
	s.RequestId = &v
	return s
}

type DisassociatePrincipalFromPortfolioResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociatePrincipalFromPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociatePrincipalFromPortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociatePrincipalFromPortfolioResponse) GoString() string {
	return s.String()
}

func (s *DisassociatePrincipalFromPortfolioResponse) SetHeaders(v map[string]*string) *DisassociatePrincipalFromPortfolioResponse {
	s.Headers = v
	return s
}

func (s *DisassociatePrincipalFromPortfolioResponse) SetStatusCode(v int32) *DisassociatePrincipalFromPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioResponse) SetBody(v *DisassociatePrincipalFromPortfolioResponseBody) *DisassociatePrincipalFromPortfolioResponse {
	s.Body = v
	return s
}

type DisassociateProductFromPortfolioRequest struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s DisassociateProductFromPortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s DisassociateProductFromPortfolioRequest) GoString() string {
	return s.String()
}

func (s *DisassociateProductFromPortfolioRequest) SetPortfolioId(v string) *DisassociateProductFromPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *DisassociateProductFromPortfolioRequest) SetProductId(v string) *DisassociateProductFromPortfolioRequest {
	s.ProductId = &v
	return s
}

type DisassociateProductFromPortfolioResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateProductFromPortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisassociateProductFromPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateProductFromPortfolioResponseBody) SetRequestId(v string) *DisassociateProductFromPortfolioResponseBody {
	s.RequestId = &v
	return s
}

type DisassociateProductFromPortfolioResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateProductFromPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateProductFromPortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s DisassociateProductFromPortfolioResponse) GoString() string {
	return s.String()
}

func (s *DisassociateProductFromPortfolioResponse) SetHeaders(v map[string]*string) *DisassociateProductFromPortfolioResponse {
	s.Headers = v
	return s
}

func (s *DisassociateProductFromPortfolioResponse) SetStatusCode(v int32) *DisassociateProductFromPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateProductFromPortfolioResponse) SetBody(v *DisassociateProductFromPortfolioResponseBody) *DisassociateProductFromPortfolioResponse {
	s.Body = v
	return s
}

type ExecuteProvisionedProductPlanRequest struct {
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s ExecuteProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *ExecuteProvisionedProductPlanRequest) SetPlanId(v string) *ExecuteProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

type ExecuteProvisionedProductPlanResponseBody struct {
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteProvisionedProductPlanResponseBody) SetPlanId(v string) *ExecuteProvisionedProductPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *ExecuteProvisionedProductPlanResponseBody) SetRequestId(v string) *ExecuteProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteProvisionedProductPlanResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *ExecuteProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *ExecuteProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *ExecuteProvisionedProductPlanResponse) SetStatusCode(v int32) *ExecuteProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteProvisionedProductPlanResponse) SetBody(v *ExecuteProvisionedProductPlanResponseBody) *ExecuteProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type GetConstraintRequest struct {
	// The ID of the constraint.
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
}

func (s GetConstraintRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConstraintRequest) GoString() string {
	return s.String()
}

func (s *GetConstraintRequest) SetConstraintId(v string) *GetConstraintRequest {
	s.ConstraintId = &v
	return s
}

type GetConstraintResponseBody struct {
	// The details of the constraint.
	ConstraintDetail *GetConstraintResponseBodyConstraintDetail `json:"ConstraintDetail,omitempty" xml:"ConstraintDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConstraintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *GetConstraintResponseBody) SetConstraintDetail(v *GetConstraintResponseBodyConstraintDetail) *GetConstraintResponseBody {
	s.ConstraintDetail = v
	return s
}

func (s *GetConstraintResponseBody) SetRequestId(v string) *GetConstraintResponseBody {
	s.RequestId = &v
	return s
}

type GetConstraintResponseBodyConstraintDetail struct {
	// The configuration of the constraint.
	//
	// Format: { "LocalRoleName": "\<role_name>" }.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the constraint.
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The type of the constraint.
	//
	// The value is fixed as Launch, which indicates the launch constraint.
	ConstraintType *string `json:"ConstraintType,omitempty" xml:"ConstraintType,omitempty"`
	// The time when the constraint was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the constraint.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the product portfolio to which the constraint belongs.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product for which the constraint is created.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
}

func (s GetConstraintResponseBodyConstraintDetail) String() string {
	return tea.Prettify(s)
}

func (s GetConstraintResponseBodyConstraintDetail) GoString() string {
	return s.String()
}

func (s *GetConstraintResponseBodyConstraintDetail) SetConfig(v string) *GetConstraintResponseBodyConstraintDetail {
	s.Config = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetConstraintId(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ConstraintId = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetConstraintType(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ConstraintType = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetCreateTime(v string) *GetConstraintResponseBodyConstraintDetail {
	s.CreateTime = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetDescription(v string) *GetConstraintResponseBodyConstraintDetail {
	s.Description = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetPortfolioId(v string) *GetConstraintResponseBodyConstraintDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetProductId(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ProductId = &v
	return s
}

func (s *GetConstraintResponseBodyConstraintDetail) SetProductName(v string) *GetConstraintResponseBodyConstraintDetail {
	s.ProductName = &v
	return s
}

type GetConstraintResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConstraintResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConstraintResponse) GoString() string {
	return s.String()
}

func (s *GetConstraintResponse) SetHeaders(v map[string]*string) *GetConstraintResponse {
	s.Headers = v
	return s
}

func (s *GetConstraintResponse) SetStatusCode(v int32) *GetConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConstraintResponse) SetBody(v *GetConstraintResponseBody) *GetConstraintResponse {
	s.Body = v
	return s
}

type GetPortfolioRequest struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
}

func (s GetPortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPortfolioRequest) GoString() string {
	return s.String()
}

func (s *GetPortfolioRequest) SetPortfolioId(v string) *GetPortfolioRequest {
	s.PortfolioId = &v
	return s
}

type GetPortfolioResponseBody struct {
	// The details of the product portfolio.
	PortfolioDetail *GetPortfolioResponseBodyPortfolioDetail `json:"PortfolioDetail,omitempty" xml:"PortfolioDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag options associated with the service portfolio.
	TagOptions []*GetPortfolioResponseBodyTagOptions `json:"TagOptions,omitempty" xml:"TagOptions,omitempty" type:"Repeated"`
}

func (s GetPortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponseBody) SetPortfolioDetail(v *GetPortfolioResponseBodyPortfolioDetail) *GetPortfolioResponseBody {
	s.PortfolioDetail = v
	return s
}

func (s *GetPortfolioResponseBody) SetRequestId(v string) *GetPortfolioResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPortfolioResponseBody) SetTagOptions(v []*GetPortfolioResponseBodyTagOptions) *GetPortfolioResponseBody {
	s.TagOptions = v
	return s
}

type GetPortfolioResponseBodyPortfolioDetail struct {
	// The time when the product portfolio was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product portfolio.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product portfolio.
	PortfolioArn *string `json:"PortfolioArn,omitempty" xml:"PortfolioArn,omitempty"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The name of the product portfolio.
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// The provider of the product portfolio.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s GetPortfolioResponseBodyPortfolioDetail) String() string {
	return tea.Prettify(s)
}

func (s GetPortfolioResponseBodyPortfolioDetail) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetCreateTime(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.CreateTime = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetDescription(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.Description = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetPortfolioArn(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.PortfolioArn = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetPortfolioId(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetPortfolioName(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.PortfolioName = &v
	return s
}

func (s *GetPortfolioResponseBodyPortfolioDetail) SetProviderName(v string) *GetPortfolioResponseBodyPortfolioDetail {
	s.ProviderName = &v
	return s
}

type GetPortfolioResponseBodyTagOptions struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// - true (default)
	// - false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the owner of the tag option.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPortfolioResponseBodyTagOptions) String() string {
	return tea.Prettify(s)
}

func (s GetPortfolioResponseBodyTagOptions) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponseBodyTagOptions) SetActive(v bool) *GetPortfolioResponseBodyTagOptions {
	s.Active = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetKey(v string) *GetPortfolioResponseBodyTagOptions {
	s.Key = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetOwner(v string) *GetPortfolioResponseBodyTagOptions {
	s.Owner = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetTagOptionId(v string) *GetPortfolioResponseBodyTagOptions {
	s.TagOptionId = &v
	return s
}

func (s *GetPortfolioResponseBodyTagOptions) SetValue(v string) *GetPortfolioResponseBodyTagOptions {
	s.Value = &v
	return s
}

type GetPortfolioResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPortfolioResponse) GoString() string {
	return s.String()
}

func (s *GetPortfolioResponse) SetHeaders(v map[string]*string) *GetPortfolioResponse {
	s.Headers = v
	return s
}

func (s *GetPortfolioResponse) SetStatusCode(v int32) *GetPortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPortfolioResponse) SetBody(v *GetPortfolioResponseBody) *GetPortfolioResponse {
	s.Body = v
	return s
}

type GetProductAsAdminRequest struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s GetProductAsAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsAdminRequest) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminRequest) SetProductId(v string) *GetProductAsAdminRequest {
	s.ProductId = &v
	return s
}

type GetProductAsAdminResponseBody struct {
	// The information about the product.
	ProductDetail *GetProductAsAdminResponseBodyProductDetail `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag options associated with the product.
	TagOptions []*GetProductAsAdminResponseBodyTagOptions `json:"TagOptions,omitempty" xml:"TagOptions,omitempty" type:"Repeated"`
}

func (s GetProductAsAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsAdminResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponseBody) SetProductDetail(v *GetProductAsAdminResponseBodyProductDetail) *GetProductAsAdminResponseBody {
	s.ProductDetail = v
	return s
}

func (s *GetProductAsAdminResponseBody) SetRequestId(v string) *GetProductAsAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProductAsAdminResponseBody) SetTagOptions(v []*GetProductAsAdminResponseBodyTagOptions) *GetProductAsAdminResponseBody {
	s.TagOptions = v
	return s
}

type GetProductAsAdminResponseBodyProductDetail struct {
	// The creation time.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetProductAsAdminResponseBodyProductDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsAdminResponseBodyProductDetail) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetCreateTime(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.CreateTime = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetDescription(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.Description = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductArn(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductArn = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductId(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductId = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductName(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductName = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProductType(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProductType = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetProviderName(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.ProviderName = &v
	return s
}

func (s *GetProductAsAdminResponseBodyProductDetail) SetTemplateType(v string) *GetProductAsAdminResponseBodyProductDetail {
	s.TemplateType = &v
	return s
}

type GetProductAsAdminResponseBodyTagOptions struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// - true (default)
	// - false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the owner of the tag option.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProductAsAdminResponseBodyTagOptions) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsAdminResponseBodyTagOptions) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetActive(v bool) *GetProductAsAdminResponseBodyTagOptions {
	s.Active = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetKey(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.Key = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetOwner(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.Owner = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetTagOptionId(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.TagOptionId = &v
	return s
}

func (s *GetProductAsAdminResponseBodyTagOptions) SetValue(v string) *GetProductAsAdminResponseBodyTagOptions {
	s.Value = &v
	return s
}

type GetProductAsAdminResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductAsAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductAsAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsAdminResponse) GoString() string {
	return s.String()
}

func (s *GetProductAsAdminResponse) SetHeaders(v map[string]*string) *GetProductAsAdminResponse {
	s.Headers = v
	return s
}

func (s *GetProductAsAdminResponse) SetStatusCode(v int32) *GetProductAsAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductAsAdminResponse) SetBody(v *GetProductAsAdminResponseBody) *GetProductAsAdminResponse {
	s.Body = v
	return s
}

type GetProductAsEndUserRequest struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s GetProductAsEndUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsEndUserRequest) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserRequest) SetProductId(v string) *GetProductAsEndUserRequest {
	s.ProductId = &v
	return s
}

type GetProductAsEndUserResponseBody struct {
	// The information about the product.
	ProductSummary *GetProductAsEndUserResponseBodyProductSummary `json:"ProductSummary,omitempty" xml:"ProductSummary,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductAsEndUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserResponseBody) SetProductSummary(v *GetProductAsEndUserResponseBodyProductSummary) *GetProductAsEndUserResponseBody {
	s.ProductSummary = v
	return s
}

func (s *GetProductAsEndUserResponseBody) SetRequestId(v string) *GetProductAsEndUserResponseBody {
	s.RequestId = &v
	return s
}

type GetProductAsEndUserResponseBodyProductSummary struct {
	// The time when the product was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the default launch option exists. Valid values:
	//
	// *   true: The default launch option exists. In this case, the PortfolioId parameter is not required when the product is launched or when the information about the product instance is updated.
	// *   false: The default launch option does not exist. In this case, the PortfolioId parameter is required when the product is launched or when the information about the product instance is updated. For more information about how to obtain the value of the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
	//
	// > If the product is added to only one product portfolio, the default launch option exists. If the product is added to multiple product portfolios, multiple launch options exist at the same time. However, no default launch options exist.
	HasDefaultLaunchOption *bool `json:"HasDefaultLaunchOption,omitempty" xml:"HasDefaultLaunchOption,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GetProductAsEndUserResponseBodyProductSummary) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsEndUserResponseBodyProductSummary) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetCreateTime(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.CreateTime = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetDescription(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.Description = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetHasDefaultLaunchOption(v bool) *GetProductAsEndUserResponseBodyProductSummary {
	s.HasDefaultLaunchOption = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductArn(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductArn = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductId(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductId = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductName(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductName = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProductType(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProductType = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetProviderName(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.ProviderName = &v
	return s
}

func (s *GetProductAsEndUserResponseBodyProductSummary) SetTemplateType(v string) *GetProductAsEndUserResponseBodyProductSummary {
	s.TemplateType = &v
	return s
}

type GetProductAsEndUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductAsEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductAsEndUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductAsEndUserResponse) GoString() string {
	return s.String()
}

func (s *GetProductAsEndUserResponse) SetHeaders(v map[string]*string) *GetProductAsEndUserResponse {
	s.Headers = v
	return s
}

func (s *GetProductAsEndUserResponse) SetStatusCode(v int32) *GetProductAsEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductAsEndUserResponse) SetBody(v *GetProductAsEndUserResponseBody) *GetProductAsEndUserResponse {
	s.Body = v
	return s
}

type GetProductVersionRequest struct {
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
}

func (s GetProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionRequest) GoString() string {
	return s.String()
}

func (s *GetProductVersionRequest) SetProductVersionId(v string) *GetProductVersionRequest {
	s.ProductVersionId = &v
	return s
}

type GetProductVersionResponseBody struct {
	ProductVersionDetail *GetProductVersionResponseBodyProductVersionDetail `json:"ProductVersionDetail,omitempty" xml:"ProductVersionDetail,omitempty" type:"Struct"`
	RequestId            *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponseBody) SetProductVersionDetail(v *GetProductVersionResponseBodyProductVersionDetail) *GetProductVersionResponseBody {
	s.ProductVersionDetail = v
	return s
}

func (s *GetProductVersionResponseBody) SetRequestId(v string) *GetProductVersionResponseBody {
	s.RequestId = &v
	return s
}

type GetProductVersionResponseBodyProductVersionDetail struct {
	Active             *bool   `json:"Active,omitempty" xml:"Active,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Guidance           *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	ProductId          *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersionId   *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	TemplateType       *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TemplateUrl        *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s GetProductVersionResponseBodyProductVersionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResponseBodyProductVersionDetail) GoString() string {
	return s.String()
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

type GetProductVersionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProductVersionResponse) GoString() string {
	return s.String()
}

func (s *GetProductVersionResponse) SetHeaders(v map[string]*string) *GetProductVersionResponse {
	s.Headers = v
	return s
}

func (s *GetProductVersionResponse) SetStatusCode(v int32) *GetProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProductVersionResponse) SetBody(v *GetProductVersionResponseBody) *GetProductVersionResponse {
	s.Body = v
	return s
}

type GetProvisionedProductRequest struct {
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
}

func (s GetProvisionedProductRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductRequest) SetProvisionedProductId(v string) *GetProvisionedProductRequest {
	s.ProvisionedProductId = &v
	return s
}

type GetProvisionedProductResponseBody struct {
	// The details of the product instance.
	ProvisionedProductDetail *GetProvisionedProductResponseBodyProvisionedProductDetail `json:"ProvisionedProductDetail,omitempty" xml:"ProvisionedProductDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProvisionedProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductResponseBody) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductResponseBody) SetProvisionedProductDetail(v *GetProvisionedProductResponseBodyProvisionedProductDetail) *GetProvisionedProductResponseBody {
	s.ProvisionedProductDetail = v
	return s
}

func (s *GetProvisionedProductResponseBody) SetRequestId(v string) *GetProvisionedProductResponseBody {
	s.RequestId = &v
	return s
}

type GetProvisionedProductResponseBodyProvisionedProductDetail struct {
	// The time when the product instance was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task that was last run on the product instance.
	//
	// The task can be one of the following types:
	//
	// *   LaunchProduct: a task that launches the product.
	// *   UpdateProvisionedProduct: a task that updates the information about the product instance.
	// *   TerminateProvisionedProduct: a task that terminates the product instance.
	LastProvisioningTaskId *string `json:"LastProvisioningTaskId,omitempty" xml:"LastProvisioningTaskId,omitempty"`
	// The ID of the last task that was successfully run on the product instance.
	//
	// The task can be one of the following types:
	//
	// *   LaunchProduct: a task that launches the product.
	// *   UpdateProvisionedProduct: a task that updates the information about the product instance.
	// *   TerminateProvisionedProduct: a task that terminates the product instance.
	LastSuccessfulProvisioningTaskId *string `json:"LastSuccessfulProvisioningTaskId,omitempty" xml:"LastSuccessfulProvisioningTaskId,omitempty"`
	// The ID of the task that was last run.
	LastTaskId *string `json:"LastTaskId,omitempty" xml:"LastTaskId,omitempty"`
	// The ID of the RAM entity to which the product instance belongs.
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The type of the Resource Access Management (RAM) entity to which the product instance belongs. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product instance.
	ProvisionedProductArn *string `json:"ProvisionedProductArn,omitempty" xml:"ProvisionedProductArn,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The type of the product instance.
	//
	// The value is fixed as RosStack, which indicates an ROS stack.
	ProvisionedProductType *string `json:"ProvisionedProductType,omitempty" xml:"ProvisionedProductType,omitempty"`
	// The ID of the Resource Orchestration Service (ROS) stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The state of the product instance. Valid values:
	//
	// *   Available: The product instance was available.
	// *   UnderChange: The information about the product instance was being changed.
	// *   Error: An exception occurred on the product instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message that is returned for the status of the product instance.
	//
	// > This parameter is returned only when Error is returned for the Status parameter.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s GetProvisionedProductResponseBodyProvisionedProductDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductResponseBodyProvisionedProductDetail) GoString() string {
	return s.String()
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

type GetProvisionedProductResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProvisionedProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProvisionedProductResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductResponse) SetHeaders(v map[string]*string) *GetProvisionedProductResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionedProductResponse) SetStatusCode(v int32) *GetProvisionedProductResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProvisionedProductResponse) SetBody(v *GetProvisionedProductResponseBody) *GetProvisionedProductResponse {
	s.Body = v
	return s
}

type GetProvisionedProductPlanRequest struct {
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s GetProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanRequest) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanRequest) SetPlanId(v string) *GetProvisionedProductPlanRequest {
	s.PlanId = &v
	return s
}

type GetProvisionedProductPlanResponseBody struct {
	// The details of the plan.
	PlanDetail *GetProvisionedProductPlanResponseBodyPlanDetail `json:"PlanDetail,omitempty" xml:"PlanDetail,omitempty" type:"Struct"`
	// The details of the product.
	ProductDetail *GetProvisionedProductPlanResponseBodyProductDetail `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty" type:"Struct"`
	// The details of the product version.
	ProductVersionDetail *GetProvisionedProductPlanResponseBodyProductVersionDetail `json:"ProductVersionDetail,omitempty" xml:"ProductVersionDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the resources to be changed in the plan.
	ResourceChanges []*GetProvisionedProductPlanResponseBodyResourceChanges `json:"ResourceChanges,omitempty" xml:"ResourceChanges,omitempty" type:"Repeated"`
}

func (s GetProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponseBodyPlanDetail struct {
	// The review details of the plan.
	ApprovalDetail *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail `json:"ApprovalDetail,omitempty" xml:"ApprovalDetail,omitempty" type:"Struct"`
	// An array that consists of reviewers.
	AssignedApprovers []*GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers `json:"AssignedApprovers,omitempty" xml:"AssignedApprovers,omitempty" type:"Repeated"`
	// The time when the plan was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The purpose of the plan. Valid values:
	//
	// *   LaunchProduct: launches the product.
	// *   UpdateProvisionedProduct: updates the information about the product instance.
	// *   TerminateProvisionedProduct: terminates the product instance.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The ID of the RAM entity to which the plan belongs.
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The name of the RAM entity to which the plan belongs.
	OwnerPrincipalName *string `json:"OwnerPrincipalName,omitempty" xml:"OwnerPrincipalName,omitempty"`
	// The type of the RAM entity to which the plan belongs. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// An array that consists of the parameters in the template.
	Parameters []*GetProvisionedProductPlanResponseBodyPlanDetailParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The type of the plan.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the ROS stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The state of the plan. Valid values:
	//
	// *   PreviewInProgress: The plan is being prechecked.
	// *   PreviewSuccess: The precheck is successful.
	// *   PreviewFailed: The precheck fails.
	// *   ExecuteInProgress: The plan is being run.
	// *   ExecuteSuccess: The plan is run.
	// *   ExecuteFailed: The plan fails to be run.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message returned for the state.
	//
	// > This parameter is returned only when PreviewFailed or ExecuteFailed is returned for Status.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// An array that consists of custom tags.
	Tags []*GetProvisionedProductPlanResponseBodyPlanDetailTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the plan belongs.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The last time when the task was modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetail) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail struct {
	// An array that consists of operations that are performed by the operator.
	OperationRecords []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords `json:"OperationRecords,omitempty" xml:"OperationRecords,omitempty" type:"Repeated"`
	// An array that consists of operations that are being performed by the plan.
	TodoTaskActivities []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities `json:"TodoTaskActivities,omitempty" xml:"TodoTaskActivities,omitempty" type:"Repeated"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) SetOperationRecords(v []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail {
	s.OperationRecords = v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail) SetTodoTaskActivities(v []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetail {
	s.TodoTaskActivities = v
	return s
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords struct {
	// The operation that is performed by the operator on the plan. Valid values:
	//
	// *   Submit: submits the plan.
	// *   Cancel: cancels the plan.
	// *   Approve: approves the plan.
	// *   reject: rejects the plan.
	ApprovalAction *string `json:"ApprovalAction,omitempty" xml:"ApprovalAction,omitempty"`
	// The review comment of the operator.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The time when the operation was performed.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The operator who performs operations on the plan.
	Operator *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Struct"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecords) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator struct {
	// The RAM entity ID of the operator.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The RAM entity name of the operator.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The RAM entity type of the operator. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailOperationRecordsOperator) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities struct {
	// The name of the operation that is being performed by the plan.
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// An array consisting of tasks that are pending for review.
	Tasks []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) SetActivityName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities {
	s.ActivityName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities) SetTasks(v []*GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivities {
	s.Tasks = v
	return s
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks struct {
	// The operator who performs operations on the plan.
	Operator *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator `json:"Operator,omitempty" xml:"Operator,omitempty" type:"Struct"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks) SetOperator(v *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasks {
	s.Operator = v
	return s
}

type GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator struct {
	// The RAM entity name of the operator.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The RAM entity type of the operator. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) SetPrincipalName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator {
	s.PrincipalName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator) SetPrincipalType(v string) *GetProvisionedProductPlanResponseBodyPlanDetailApprovalDetailTodoTaskActivitiesTasksOperator {
	s.PrincipalType = &v
	return s
}

type GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers struct {
	// The RAM entity name of the reviewer.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the Resource Access Management (RAM) entity of the reviewer. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) SetPrincipalName(v string) *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers {
	s.PrincipalName = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers) SetPrincipalType(v string) *GetProvisionedProductPlanResponseBodyPlanDetailAssignedApprovers {
	s.PrincipalType = &v
	return s
}

type GetProvisionedProductPlanResponseBodyPlanDetailParameters struct {
	// The name of the parameter in the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter in the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailParameters) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailParameters) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) SetParameterKey(v string) *GetProvisionedProductPlanResponseBodyPlanDetailParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailParameters) SetParameterValue(v string) *GetProvisionedProductPlanResponseBodyPlanDetailParameters {
	s.ParameterValue = &v
	return s
}

type GetProvisionedProductPlanResponseBodyPlanDetailTags struct {
	// The key of the custom tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailTags) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyPlanDetailTags) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) SetKey(v string) *GetProvisionedProductPlanResponseBodyPlanDetailTags {
	s.Key = &v
	return s
}

func (s *GetProvisionedProductPlanResponseBodyPlanDetailTags) SetValue(v string) *GetProvisionedProductPlanResponseBodyPlanDetailTags {
	s.Value = &v
	return s
}

type GetProvisionedProductPlanResponseBodyProductDetail struct {
	// The creation time.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates ROS.
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyProductDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyProductDetail) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponseBodyProductVersionDetail struct {
	// Indicates whether the product version is visible to end users. Valid values:
	//
	// *   true (default)
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The time when the product version was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// *   Default: No recommendation information is provided. This is the default value.
	// *   Recommended: the recommendation version.
	// *   Latest: the latest version.
	// *   Deprecated: the version that is about to be deprecated.
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product to which the product version belongs.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name for the version of the product.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the template.
	//
	// The value is fixed as RosTerraformTemplate, which indicates that the Terraform template is supported by ROS.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyProductVersionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyProductVersionDetail) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponseBodyResourceChanges struct {
	// The action that is performed on the resource. Valid values:
	//
	// *   Add
	// *   Modify
	// *   Remove
	// *   None
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The logical ID of the resource.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// > This parameter is returned only when Action is set to Modify or Remove.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// Indicates whether a replacement update is performed on the template. Valid values:
	//
	// *   True: A replacement update is performed on the template.
	// *   False: A change is made on the template.
	// *   Conditional: A replacement update may be performed on the template. You can check whether a replacement update is performed when the template is in use.
	//
	// > This parameter is returned only when Action is set to Modify.
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetProvisionedProductPlanResponseBodyResourceChanges) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponseBodyResourceChanges) GoString() string {
	return s.String()
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

type GetProvisionedProductPlanResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *GetProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionedProductPlanResponse) SetStatusCode(v int32) *GetProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProvisionedProductPlanResponse) SetBody(v *GetProvisionedProductPlanResponseBody) *GetProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type GetTagOptionRequest struct {
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s GetTagOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTagOptionRequest) GoString() string {
	return s.String()
}

func (s *GetTagOptionRequest) SetTagOptionId(v string) *GetTagOptionRequest {
	s.TagOptionId = &v
	return s
}

type GetTagOptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the tag option.
	TagOptionDetail *GetTagOptionResponseBodyTagOptionDetail `json:"TagOptionDetail,omitempty" xml:"TagOptionDetail,omitempty" type:"Struct"`
}

func (s GetTagOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetTagOptionResponseBody) SetRequestId(v string) *GetTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTagOptionResponseBody) SetTagOptionDetail(v *GetTagOptionResponseBodyTagOptionDetail) *GetTagOptionResponseBody {
	s.TagOptionDetail = v
	return s
}

type GetTagOptionResponseBodyTagOptionDetail struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// *   true
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the Alibaba Cloud account to which the tag option belongs.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTagOptionResponseBodyTagOptionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetTagOptionResponseBodyTagOptionDetail) GoString() string {
	return s.String()
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetActive(v bool) *GetTagOptionResponseBodyTagOptionDetail {
	s.Active = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetKey(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.Key = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetOwner(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.Owner = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetTagOptionId(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.TagOptionId = &v
	return s
}

func (s *GetTagOptionResponseBodyTagOptionDetail) SetValue(v string) *GetTagOptionResponseBodyTagOptionDetail {
	s.Value = &v
	return s
}

type GetTagOptionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTagOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTagOptionResponse) GoString() string {
	return s.String()
}

func (s *GetTagOptionResponse) SetHeaders(v map[string]*string) *GetTagOptionResponse {
	s.Headers = v
	return s
}

func (s *GetTagOptionResponse) SetStatusCode(v int32) *GetTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTagOptionResponse) SetBody(v *GetTagOptionResponseBody) *GetTagOptionResponse {
	s.Body = v
	return s
}

type GetTaskRequest struct {
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

type GetTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the task.
	TaskDetail *GetTaskResponseBodyTaskDetail `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTaskDetail(v *GetTaskResponseBodyTaskDetail) *GetTaskResponseBody {
	s.TaskDetail = v
	return s
}

type GetTaskResponseBodyTaskDetail struct {
	// The time when the task was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The logs of the instance.
	Log *GetTaskResponseBodyTaskDetailLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The output parameters of the template.
	Outputs []*GetTaskResponseBodyTaskDetailOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters that are specified in the template.
	Parameters []*GetTaskResponseBodyTaskDetailParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The state of the task. Valid values:
	//
	// *   Succeeded: The task was successful.
	// *   InProgress: The task was in progress.
	// *   Failed: The task failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message that is returned for the status of the task.
	//
	// > This parameter is returned only when Failed is returned for the Status parameter.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The ID of the task.
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The custom tags.
	TaskTags []*GetTaskResponseBodyTaskDetailTaskTags `json:"TaskTags,omitempty" xml:"TaskTags,omitempty" type:"Repeated"`
	// The type of the task. Valid values:
	//
	// *   LaunchProduct: a task that launches the product.
	// *   UpdateProvisionedProduct: a task that updates the information about the product instance.
	// *   TerminateProvisionedProduct: a task that terminates the product instance.
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The time when the task was last modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTaskResponseBodyTaskDetail) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetail) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetail) SetCreateTime(v string) *GetTaskResponseBodyTaskDetail {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetLog(v *GetTaskResponseBodyTaskDetailLog) *GetTaskResponseBodyTaskDetail {
	s.Log = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetOutputs(v []*GetTaskResponseBodyTaskDetailOutputs) *GetTaskResponseBodyTaskDetail {
	s.Outputs = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetParameters(v []*GetTaskResponseBodyTaskDetailParameters) *GetTaskResponseBodyTaskDetail {
	s.Parameters = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetPortfolioId(v string) *GetTaskResponseBodyTaskDetail {
	s.PortfolioId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductId(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductName(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductName = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductVersionId(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductVersionId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProductVersionName(v string) *GetTaskResponseBodyTaskDetail {
	s.ProductVersionName = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProvisionedProductId(v string) *GetTaskResponseBodyTaskDetail {
	s.ProvisionedProductId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetProvisionedProductName(v string) *GetTaskResponseBodyTaskDetail {
	s.ProvisionedProductName = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetStatus(v string) *GetTaskResponseBodyTaskDetail {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetStatusMessage(v string) *GetTaskResponseBodyTaskDetail {
	s.StatusMessage = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetTaskId(v string) *GetTaskResponseBodyTaskDetail {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetTaskTags(v []*GetTaskResponseBodyTaskDetailTaskTags) *GetTaskResponseBodyTaskDetail {
	s.TaskTags = v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetTaskType(v string) *GetTaskResponseBodyTaskDetail {
	s.TaskType = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetail) SetUpdateTime(v string) *GetTaskResponseBodyTaskDetail {
	s.UpdateTime = &v
	return s
}

type GetTaskResponseBodyTaskDetailLog struct {
	// The Terraform logs.
	TerraformLogs []*GetTaskResponseBodyTaskDetailLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s GetTaskResponseBodyTaskDetailLog) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailLog) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailLog) SetTerraformLogs(v []*GetTaskResponseBodyTaskDetailLogTerraformLogs) *GetTaskResponseBodyTaskDetailLog {
	s.TerraformLogs = v
	return s
}

type GetTaskResponseBodyTaskDetailLogTerraformLogs struct {
	// The name of the Terraform command that is run. Valid values:
	//
	// *   apply
	// *   plan
	// *   destroy
	// *   version
	//
	// For more information about Terraform commands, see [Basic CLI Features](https://www.terraform.io/cli/commands).
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The content of the output stream that is returned after the command is run.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The output stream. Valid values:
	//
	// *   stdout: a standard output stream
	// *   stderr: a standard error stream
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailLogTerraformLogs) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) SetCommand(v string) *GetTaskResponseBodyTaskDetailLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) SetContent(v string) *GetTaskResponseBodyTaskDetailLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailLogTerraformLogs) SetStream(v string) *GetTaskResponseBodyTaskDetailLogTerraformLogs {
	s.Stream = &v
	return s
}

type GetTaskResponseBodyTaskDetailOutputs struct {
	// The description of the output parameter for the template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the output parameter for the template.
	OutputKey *string `json:"OutputKey,omitempty" xml:"OutputKey,omitempty"`
	// The value of the output parameter for the template.
	OutputValue *string `json:"OutputValue,omitempty" xml:"OutputValue,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailOutputs) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailOutputs) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailOutputs) SetDescription(v string) *GetTaskResponseBodyTaskDetailOutputs {
	s.Description = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailOutputs) SetOutputKey(v string) *GetTaskResponseBodyTaskDetailOutputs {
	s.OutputKey = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailOutputs) SetOutputValue(v string) *GetTaskResponseBodyTaskDetailOutputs {
	s.OutputValue = &v
	return s
}

type GetTaskResponseBodyTaskDetailParameters struct {
	// The name of the input parameter for the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the input parameter for the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailParameters) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailParameters) SetParameterKey(v string) *GetTaskResponseBodyTaskDetailParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailParameters) SetParameterValue(v string) *GetTaskResponseBodyTaskDetailParameters {
	s.ParameterValue = &v
	return s
}

type GetTaskResponseBodyTaskDetailTaskTags struct {
	// The custom tag key.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The custom tag value.
	//
	// The value must be 1 to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTaskResponseBodyTaskDetailTaskTags) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponseBodyTaskDetailTaskTags) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) SetKey(v string) *GetTaskResponseBodyTaskDetailTaskTags {
	s.Key = &v
	return s
}

func (s *GetTaskResponseBodyTaskDetailTaskTags) SetValue(v string) *GetTaskResponseBodyTaskDetailTaskTags {
	s.Value = &v
	return s
}

type GetTaskResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

func (s *GetTaskResponse) SetStatusCode(v int32) *GetTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskResponse) SetBody(v *GetTaskResponseBody) *GetTaskResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetProductId(v string) *GetTemplateRequest {
	s.ProductId = &v
	return s
}

func (s *GetTemplateRequest) SetProductVersionId(v string) *GetTemplateRequest {
	s.ProductVersionId = &v
	return s
}

type GetTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The content of the template.
	//
	// For more information about the template syntax, see [Structure of Terraform templates](~~184397~~).
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateBody(v string) *GetTemplateResponseBody {
	s.TemplateBody = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
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
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product instance.
	//
	// The value must be 1 to 128 characters in length.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the region to which the Resource Orchestration Service (ROS) stack belongs.
	//
	// For more information about how to obtain the regions that are supported by ROS, see [DescribeRegions](~~131035~~).
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The custom tags that are specified by the end user.
	//
	// Maximum value of N: 20.
	//
	// >
	//
	// *   The Tags parameter is optional. If you specify the Tags parameter, you must specify the Tags.N.Key and Tags.N.Value parameters.
	//
	// *   The tag is propagated to each stack resource that supports the tag feature.
	Tags []*LaunchProductRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s LaunchProductRequest) String() string {
	return tea.Prettify(s)
}

func (s LaunchProductRequest) GoString() string {
	return s.String()
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

type LaunchProductRequestParameters struct {
	// The name of the input parameter for the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the input parameter for the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s LaunchProductRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s LaunchProductRequestParameters) GoString() string {
	return s.String()
}

func (s *LaunchProductRequestParameters) SetParameterKey(v string) *LaunchProductRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *LaunchProductRequestParameters) SetParameterValue(v string) *LaunchProductRequestParameters {
	s.ParameterValue = &v
	return s
}

type LaunchProductRequestTags struct {
	// The tag key of the custom tag.
	//
	// The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the custom tag.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LaunchProductRequestTags) String() string {
	return tea.Prettify(s)
}

func (s LaunchProductRequestTags) GoString() string {
	return s.String()
}

func (s *LaunchProductRequestTags) SetKey(v string) *LaunchProductRequestTags {
	s.Key = &v
	return s
}

func (s *LaunchProductRequestTags) SetValue(v string) *LaunchProductRequestTags {
	s.Value = &v
	return s
}

type LaunchProductResponseBody struct {
	// The ID of the instance
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LaunchProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LaunchProductResponseBody) GoString() string {
	return s.String()
}

func (s *LaunchProductResponseBody) SetProvisionedProductId(v string) *LaunchProductResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *LaunchProductResponseBody) SetRequestId(v string) *LaunchProductResponseBody {
	s.RequestId = &v
	return s
}

type LaunchProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchProductResponse) String() string {
	return tea.Prettify(s)
}

func (s LaunchProductResponse) GoString() string {
	return s.String()
}

func (s *LaunchProductResponse) SetHeaders(v map[string]*string) *LaunchProductResponse {
	s.Headers = v
	return s
}

func (s *LaunchProductResponse) SetStatusCode(v int32) *LaunchProductResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchProductResponse) SetBody(v *LaunchProductResponseBody) *LaunchProductResponse {
	s.Body = v
	return s
}

type ListLaunchOptionsRequest struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListLaunchOptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLaunchOptionsRequest) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsRequest) SetProductId(v string) *ListLaunchOptionsRequest {
	s.ProductId = &v
	return s
}

type ListLaunchOptionsResponseBody struct {
	// The launch options.
	LaunchOptionSummaries []*ListLaunchOptionsResponseBodyLaunchOptionSummaries `json:"LaunchOptionSummaries,omitempty" xml:"LaunchOptionSummaries,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLaunchOptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLaunchOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponseBody) SetLaunchOptionSummaries(v []*ListLaunchOptionsResponseBodyLaunchOptionSummaries) *ListLaunchOptionsResponseBody {
	s.LaunchOptionSummaries = v
	return s
}

func (s *ListLaunchOptionsResponseBody) SetRequestId(v string) *ListLaunchOptionsResponseBody {
	s.RequestId = &v
	return s
}

type ListLaunchOptionsResponseBodyLaunchOptionSummaries struct {
	// The constraints.
	ConstraintSummaries []*ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries `json:"ConstraintSummaries,omitempty" xml:"ConstraintSummaries,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The name of the product portfolio.
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummaries) String() string {
	return tea.Prettify(s)
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummaries) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) SetConstraintSummaries(v []*ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) *ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	s.ConstraintSummaries = v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) SetPortfolioId(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	s.PortfolioId = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummaries) SetPortfolioName(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummaries {
	s.PortfolioName = &v
	return s
}

type ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries struct {
	// The type of the constraint. Valid values:
	//
	// 1.  Launch
	// 2.  Template
	// 3.  Approval
	// 4.  StackRegion
	ConstraintType *string `json:"ConstraintType,omitempty" xml:"ConstraintType,omitempty"`
	// The description of the constraint.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The types of operations that require approval. If the type of the constraint is Approval, this parameter is not empty.
	OperationTypes []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	// The regions in which users can launch the service. If the type of the constraint is StackRegion, this parameter is not empty.
	StackRegions []*string `json:"StackRegions,omitempty" xml:"StackRegions,omitempty" type:"Repeated"`
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) String() string {
	return tea.Prettify(s)
}

func (s ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetConstraintType(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.ConstraintType = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetDescription(v string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.Description = &v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetOperationTypes(v []*string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.OperationTypes = v
	return s
}

func (s *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries) SetStackRegions(v []*string) *ListLaunchOptionsResponseBodyLaunchOptionSummariesConstraintSummaries {
	s.StackRegions = v
	return s
}

type ListLaunchOptionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLaunchOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLaunchOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLaunchOptionsResponse) GoString() string {
	return s.String()
}

func (s *ListLaunchOptionsResponse) SetHeaders(v map[string]*string) *ListLaunchOptionsResponse {
	s.Headers = v
	return s
}

func (s *ListLaunchOptionsResponse) SetStatusCode(v int32) *ListLaunchOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLaunchOptionsResponse) SetBody(v *ListLaunchOptionsResponseBody) *ListLaunchOptionsResponse {
	s.Body = v
	return s
}

type ListPortfoliosRequest struct {
	Filters    []*ListPortfoliosRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	PageNumber *int32                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProductId  *string                         `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	Scope      *string                         `json:"Scope,omitempty" xml:"Scope,omitempty"`
	SortBy     *string                         `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortOrder  *string                         `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListPortfoliosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPortfoliosRequest) GoString() string {
	return s.String()
}

func (s *ListPortfoliosRequest) SetFilters(v []*ListPortfoliosRequestFilters) *ListPortfoliosRequest {
	s.Filters = v
	return s
}

func (s *ListPortfoliosRequest) SetPageNumber(v int32) *ListPortfoliosRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPortfoliosRequest) SetPageSize(v int32) *ListPortfoliosRequest {
	s.PageSize = &v
	return s
}

func (s *ListPortfoliosRequest) SetProductId(v string) *ListPortfoliosRequest {
	s.ProductId = &v
	return s
}

func (s *ListPortfoliosRequest) SetScope(v string) *ListPortfoliosRequest {
	s.Scope = &v
	return s
}

func (s *ListPortfoliosRequest) SetSortBy(v string) *ListPortfoliosRequest {
	s.SortBy = &v
	return s
}

func (s *ListPortfoliosRequest) SetSortOrder(v string) *ListPortfoliosRequest {
	s.SortOrder = &v
	return s
}

type ListPortfoliosRequestFilters struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPortfoliosRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListPortfoliosRequestFilters) GoString() string {
	return s.String()
}

func (s *ListPortfoliosRequestFilters) SetKey(v string) *ListPortfoliosRequestFilters {
	s.Key = &v
	return s
}

func (s *ListPortfoliosRequestFilters) SetValue(v string) *ListPortfoliosRequestFilters {
	s.Value = &v
	return s
}

type ListPortfoliosResponseBody struct {
	PageNumber       *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PortfolioDetails []*ListPortfoliosResponseBodyPortfolioDetails `json:"PortfolioDetails,omitempty" xml:"PortfolioDetails,omitempty" type:"Repeated"`
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount       *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPortfoliosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPortfoliosResponseBody) GoString() string {
	return s.String()
}

func (s *ListPortfoliosResponseBody) SetPageNumber(v int32) *ListPortfoliosResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPortfoliosResponseBody) SetPageSize(v int32) *ListPortfoliosResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPortfoliosResponseBody) SetPortfolioDetails(v []*ListPortfoliosResponseBodyPortfolioDetails) *ListPortfoliosResponseBody {
	s.PortfolioDetails = v
	return s
}

func (s *ListPortfoliosResponseBody) SetRequestId(v string) *ListPortfoliosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPortfoliosResponseBody) SetTotalCount(v int32) *ListPortfoliosResponseBody {
	s.TotalCount = &v
	return s
}

type ListPortfoliosResponseBodyPortfolioDetails struct {
	// 代表创建时间的资源属性字段
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 产品组合描述
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PortfolioArn *string `json:"PortfolioArn,omitempty" xml:"PortfolioArn,omitempty"`
	// 代表资源一级ID的资源属性字段
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// 代表资源名称的资源属性字段
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// 产品组合提供方
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s ListPortfoliosResponseBodyPortfolioDetails) String() string {
	return tea.Prettify(s)
}

func (s ListPortfoliosResponseBodyPortfolioDetails) GoString() string {
	return s.String()
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetCreateTime(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.CreateTime = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetDescription(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.Description = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetPortfolioArn(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.PortfolioArn = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetPortfolioId(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.PortfolioId = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetPortfolioName(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.PortfolioName = &v
	return s
}

func (s *ListPortfoliosResponseBodyPortfolioDetails) SetProviderName(v string) *ListPortfoliosResponseBodyPortfolioDetails {
	s.ProviderName = &v
	return s
}

type ListPortfoliosResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPortfoliosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPortfoliosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPortfoliosResponse) GoString() string {
	return s.String()
}

func (s *ListPortfoliosResponse) SetHeaders(v map[string]*string) *ListPortfoliosResponse {
	s.Headers = v
	return s
}

func (s *ListPortfoliosResponse) SetStatusCode(v int32) *ListPortfoliosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPortfoliosResponse) SetBody(v *ListPortfoliosResponseBody) *ListPortfoliosResponse {
	s.Body = v
	return s
}

type ListPrincipalsRequest struct {
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
}

func (s ListPrincipalsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrincipalsRequest) GoString() string {
	return s.String()
}

func (s *ListPrincipalsRequest) SetPortfolioId(v string) *ListPrincipalsRequest {
	s.PortfolioId = &v
	return s
}

type ListPrincipalsResponseBody struct {
	// The RAM entities.
	Principals []*ListPrincipalsResponseBodyPrincipals `json:"Principals,omitempty" xml:"Principals,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrincipalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrincipalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrincipalsResponseBody) SetPrincipals(v []*ListPrincipalsResponseBodyPrincipals) *ListPrincipalsResponseBody {
	s.Principals = v
	return s
}

func (s *ListPrincipalsResponseBody) SetRequestId(v string) *ListPrincipalsResponseBody {
	s.RequestId = &v
	return s
}

type ListPrincipalsResponseBodyPrincipals struct {
	// The ID of the RAM entity.
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the RAM entity. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListPrincipalsResponseBodyPrincipals) String() string {
	return tea.Prettify(s)
}

func (s ListPrincipalsResponseBodyPrincipals) GoString() string {
	return s.String()
}

func (s *ListPrincipalsResponseBodyPrincipals) SetPrincipalId(v string) *ListPrincipalsResponseBodyPrincipals {
	s.PrincipalId = &v
	return s
}

func (s *ListPrincipalsResponseBodyPrincipals) SetPrincipalType(v string) *ListPrincipalsResponseBodyPrincipals {
	s.PrincipalType = &v
	return s
}

type ListPrincipalsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrincipalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrincipalsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrincipalsResponse) GoString() string {
	return s.String()
}

func (s *ListPrincipalsResponse) SetHeaders(v map[string]*string) *ListPrincipalsResponse {
	s.Headers = v
	return s
}

func (s *ListPrincipalsResponse) SetStatusCode(v int32) *ListPrincipalsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrincipalsResponse) SetBody(v *ListPrincipalsResponseBody) *ListPrincipalsResponse {
	s.Body = v
	return s
}

type ListProductVersionsRequest struct {
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
}

func (s ListProductVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListProductVersionsRequest) SetProductId(v string) *ListProductVersionsRequest {
	s.ProductId = &v
	return s
}

type ListProductVersionsResponseBody struct {
	// The versions of the product.
	ProductVersionDetails []*ListProductVersionsResponseBodyProductVersionDetails `json:"ProductVersionDetails,omitempty" xml:"ProductVersionDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProductVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponseBody) SetProductVersionDetails(v []*ListProductVersionsResponseBodyProductVersionDetails) *ListProductVersionsResponseBody {
	s.ProductVersionDetails = v
	return s
}

func (s *ListProductVersionsResponseBody) SetRequestId(v string) *ListProductVersionsResponseBody {
	s.RequestId = &v
	return s
}

type ListProductVersionsResponseBodyProductVersionDetails struct {
	// Indicates whether the product version is enabled. Valid values:
	//
	// true: The product version is enabled. This is the default value. false: The product version is disabled.
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The time when the product version was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// *   Default: No recommendation information is provided. This is the default value.
	// *   Recommended: the recommended version.
	// *   Latest: the latest version.
	// *   Deprecated: the version that is about to be discontinued.
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product to which the product version belongs.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The type of the template.
	//
	// The value is fixed as RosTerraformTemplate, which indicates the Terraform template that is supported by Resource Orchestration Service (ROS).
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The URL of the template.
	TemplateUrl *string `json:"TemplateUrl,omitempty" xml:"TemplateUrl,omitempty"`
}

func (s ListProductVersionsResponseBodyProductVersionDetails) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponseBodyProductVersionDetails) GoString() string {
	return s.String()
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

type ListProductVersionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductVersionsResponse) SetHeaders(v map[string]*string) *ListProductVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductVersionsResponse) SetStatusCode(v int32) *ListProductVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductVersionsResponse) SetBody(v *ListProductVersionsResponseBody) *ListProductVersionsResponse {
	s.Body = v
	return s
}

type ListProductsAsAdminRequest struct {
	// The filter conditions.
	Filters []*ListProductsAsAdminRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The query scope. Valid values:
	//
	// *   Local: the products that are created by using the current account. This is the default value.
	// *   Import: the products that are imported from other accounts.
	// *   All: all available products.
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The field that is used to sort the queried data.
	//
	// Set the value to CreateTime, which specifies the time when the product was created.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the queried data. Valid values:
	//
	// *   Asc: the ascending order
	// *   Desc: the descending order
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductsAsAdminRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsAdminRequest) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminRequest) SetFilters(v []*ListProductsAsAdminRequestFilters) *ListProductsAsAdminRequest {
	s.Filters = v
	return s
}

func (s *ListProductsAsAdminRequest) SetPageNumber(v int32) *ListProductsAsAdminRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetPageSize(v int32) *ListProductsAsAdminRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetPortfolioId(v string) *ListProductsAsAdminRequest {
	s.PortfolioId = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetScope(v string) *ListProductsAsAdminRequest {
	s.Scope = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetSortBy(v string) *ListProductsAsAdminRequest {
	s.SortBy = &v
	return s
}

func (s *ListProductsAsAdminRequest) SetSortOrder(v string) *ListProductsAsAdminRequest {
	s.SortOrder = &v
	return s
}

type ListProductsAsAdminRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// *   ProductName: performs exact matches by product name. Product names are not case-sensitive.
	// *   FullTextSearch: performs full-text searches by product name, product provider, or product description. Fuzzy match is supported.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductsAsAdminRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsAdminRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminRequestFilters) SetKey(v string) *ListProductsAsAdminRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProductsAsAdminRequestFilters) SetValue(v string) *ListProductsAsAdminRequestFilters {
	s.Value = &v
	return s
}

type ListProductsAsAdminResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The products.
	ProductDetails []*ListProductsAsAdminResponseBodyProductDetails `json:"ProductDetails,omitempty" xml:"ProductDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsAsAdminResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsAdminResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminResponseBody) SetPageNumber(v int32) *ListProductsAsAdminResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetPageSize(v int32) *ListProductsAsAdminResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetProductDetails(v []*ListProductsAsAdminResponseBodyProductDetails) *ListProductsAsAdminResponseBody {
	s.ProductDetails = v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetRequestId(v string) *ListProductsAsAdminResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsAsAdminResponseBody) SetTotalCount(v int32) *ListProductsAsAdminResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductsAsAdminResponseBodyProductDetails struct {
	// The time when the product was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListProductsAsAdminResponseBodyProductDetails) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsAdminResponseBodyProductDetails) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetCreateTime(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.CreateTime = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetDescription(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.Description = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductArn(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductArn = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductId(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductId = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductName(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductName = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProductType(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProductType = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetProviderName(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.ProviderName = &v
	return s
}

func (s *ListProductsAsAdminResponseBodyProductDetails) SetTemplateType(v string) *ListProductsAsAdminResponseBodyProductDetails {
	s.TemplateType = &v
	return s
}

type ListProductsAsAdminResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsAsAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsAsAdminResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsAdminResponse) GoString() string {
	return s.String()
}

func (s *ListProductsAsAdminResponse) SetHeaders(v map[string]*string) *ListProductsAsAdminResponse {
	s.Headers = v
	return s
}

func (s *ListProductsAsAdminResponse) SetStatusCode(v int32) *ListProductsAsAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsAsAdminResponse) SetBody(v *ListProductsAsAdminResponseBody) *ListProductsAsAdminResponse {
	s.Body = v
	return s
}

type ListProductsAsEndUserRequest struct {
	// The filter conditions.
	Filters []*ListProductsAsEndUserRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field that is used to sort the queried data.
	//
	// Set the value to CreateTime, which specifies the time when the product was created.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the queried data. Valid values:
	//
	// *   Asc: the ascending order
	// *   Desc: the descending order
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProductsAsEndUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsEndUserRequest) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserRequest) SetFilters(v []*ListProductsAsEndUserRequestFilters) *ListProductsAsEndUserRequest {
	s.Filters = v
	return s
}

func (s *ListProductsAsEndUserRequest) SetPageNumber(v int32) *ListProductsAsEndUserRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsEndUserRequest) SetPageSize(v int32) *ListProductsAsEndUserRequest {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsEndUserRequest) SetSortBy(v string) *ListProductsAsEndUserRequest {
	s.SortBy = &v
	return s
}

func (s *ListProductsAsEndUserRequest) SetSortOrder(v string) *ListProductsAsEndUserRequest {
	s.SortOrder = &v
	return s
}

type ListProductsAsEndUserRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// *   ProductName: performs exact matches by product name. Product names are not case-sensitive.
	// *   FullTextSearch: performs full-text searches by product name, product provider, or product description. Fuzzy match is supported.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProductsAsEndUserRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsEndUserRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserRequestFilters) SetKey(v string) *ListProductsAsEndUserRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProductsAsEndUserRequestFilters) SetValue(v string) *ListProductsAsEndUserRequestFilters {
	s.Value = &v
	return s
}

type ListProductsAsEndUserResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The products.
	ProductSummaries []*ListProductsAsEndUserResponseBodyProductSummaries `json:"ProductSummaries,omitempty" xml:"ProductSummaries,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProductsAsEndUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsEndUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserResponseBody) SetPageNumber(v int32) *ListProductsAsEndUserResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetPageSize(v int32) *ListProductsAsEndUserResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetProductSummaries(v []*ListProductsAsEndUserResponseBodyProductSummaries) *ListProductsAsEndUserResponseBody {
	s.ProductSummaries = v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetRequestId(v string) *ListProductsAsEndUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductsAsEndUserResponseBody) SetTotalCount(v int32) *ListProductsAsEndUserResponseBody {
	s.TotalCount = &v
	return s
}

type ListProductsAsEndUserResponseBodyProductSummaries struct {
	// The time when the product was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the product.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the default launch option exists. Valid values:
	//
	// *   true: The default launch option exists. In this case, the PortfolioId parameter is not required when the product is launched or when the information about the product instance is updated.
	// *   false: The default launch option does not exist. In this case, the PortfolioId parameter is required when the product is launched or when the information about the product instance is updated. For more information about how to obtain the value of the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
	//
	// > If the product is added to only one product portfolio, the default launch option exists. If the product is added to multiple product portfolios, multiple launch options exist at the same time. However, no default launch options exist.
	HasDefaultLaunchOption *bool `json:"HasDefaultLaunchOption,omitempty" xml:"HasDefaultLaunchOption,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product.
	ProductArn *string `json:"ProductArn,omitempty" xml:"ProductArn,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The type of the product.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The provider of the product.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListProductsAsEndUserResponseBodyProductSummaries) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsEndUserResponseBodyProductSummaries) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetCreateTime(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.CreateTime = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetDescription(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.Description = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetHasDefaultLaunchOption(v bool) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.HasDefaultLaunchOption = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductArn(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductArn = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductId(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductId = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductName(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductName = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProductType(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProductType = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetProviderName(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.ProviderName = &v
	return s
}

func (s *ListProductsAsEndUserResponseBodyProductSummaries) SetTemplateType(v string) *ListProductsAsEndUserResponseBodyProductSummaries {
	s.TemplateType = &v
	return s
}

type ListProductsAsEndUserResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductsAsEndUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductsAsEndUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProductsAsEndUserResponse) GoString() string {
	return s.String()
}

func (s *ListProductsAsEndUserResponse) SetHeaders(v map[string]*string) *ListProductsAsEndUserResponse {
	s.Headers = v
	return s
}

func (s *ListProductsAsEndUserResponse) SetStatusCode(v int32) *ListProductsAsEndUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductsAsEndUserResponse) SetBody(v *ListProductsAsEndUserResponseBody) *ListProductsAsEndUserResponse {
	s.Body = v
	return s
}

type ListProvisionedProductPlanApproversRequest struct {
	// The access filter. Valid values:
	//
	// *   User (default): queries the plans that are created by the current requester.
	// *   Account: queries the plans that belong to the current Alibaba Cloud account.
	// *   ResourceDirectory: queries the plans that belong to the current resource directory.
	//
	// >  You must specify one of the `ApprovalFilter` and `AccessLevelFilter` parameters, but not both.
	AccessLevelFilter *string `json:"AccessLevelFilter,omitempty" xml:"AccessLevelFilter,omitempty"`
	// The access filter of the review dimension. Valid values:
	//
	// *   AccountRequests: queries all reviewed plans that belong to the current Alibaba Cloud account.
	// *   ResourceDirectoryRequests: queries all plans that belong to the current resource directory.
	//
	// >  You must specify one of the `ApprovalFilter` and `AccessLevelFilter` parameters, but not both.
	ApprovalFilter *string `json:"ApprovalFilter,omitempty" xml:"ApprovalFilter,omitempty"`
	// An array that consists of filter conditions.
	Filters []*ListProvisionedProductPlanApproversRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
}

func (s ListProvisionedProductPlanApproversRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlanApproversRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversRequest) SetAccessLevelFilter(v string) *ListProvisionedProductPlanApproversRequest {
	s.AccessLevelFilter = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequest) SetApprovalFilter(v string) *ListProvisionedProductPlanApproversRequest {
	s.ApprovalFilter = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequest) SetFilters(v []*ListProvisionedProductPlanApproversRequestFilters) *ListProvisionedProductPlanApproversRequest {
	s.Filters = v
	return s
}

type ListProvisionedProductPlanApproversRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// *   ProvisionedProductPlanApproverName: performs fuzzy match by reviewer name.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductPlanApproversRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlanApproversRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversRequestFilters) SetKey(v string) *ListProvisionedProductPlanApproversRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductPlanApproversRequestFilters) SetValue(v string) *ListProvisionedProductPlanApproversRequestFilters {
	s.Value = &v
	return s
}

type ListProvisionedProductPlanApproversResponseBody struct {
	// An array that consists of reviewers.
	Approvers []*ListProvisionedProductPlanApproversResponseBodyApprovers `json:"Approvers,omitempty" xml:"Approvers,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProvisionedProductPlanApproversResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlanApproversResponseBody) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversResponseBody) SetApprovers(v []*ListProvisionedProductPlanApproversResponseBodyApprovers) *ListProvisionedProductPlanApproversResponseBody {
	s.Approvers = v
	return s
}

func (s *ListProvisionedProductPlanApproversResponseBody) SetRequestId(v string) *ListProvisionedProductPlanApproversResponseBody {
	s.RequestId = &v
	return s
}

type ListProvisionedProductPlanApproversResponseBodyApprovers struct {
	// The name of the reviewer.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the Resource Access Management (RAM) entity of the reviewer. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListProvisionedProductPlanApproversResponseBodyApprovers) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlanApproversResponseBodyApprovers) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) SetPrincipalName(v string) *ListProvisionedProductPlanApproversResponseBodyApprovers {
	s.PrincipalName = &v
	return s
}

func (s *ListProvisionedProductPlanApproversResponseBodyApprovers) SetPrincipalType(v string) *ListProvisionedProductPlanApproversResponseBodyApprovers {
	s.PrincipalType = &v
	return s
}

type ListProvisionedProductPlanApproversResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionedProductPlanApproversResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionedProductPlanApproversResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlanApproversResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlanApproversResponse) SetHeaders(v map[string]*string) *ListProvisionedProductPlanApproversResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionedProductPlanApproversResponse) SetStatusCode(v int32) *ListProvisionedProductPlanApproversResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionedProductPlanApproversResponse) SetBody(v *ListProvisionedProductPlanApproversResponseBody) *ListProvisionedProductPlanApproversResponse {
	s.Body = v
	return s
}

type ListProvisionedProductPlansRequest struct {
	// The access filter. Valid values:
	//
	// *   User (default): queries the plans that are created by the current requester.
	// *   Account: queries the plans that belong to the current Alibaba Cloud account.
	// *   ResourceDirectory: queries the plans that belong to the current resource directory.
	AccessLevelFilter *string `json:"AccessLevelFilter,omitempty" xml:"AccessLevelFilter,omitempty"`
	// The access filter of the review dimension. Valid values:
	//
	// *   ReceivedRequests: queries plans that are pending for review.
	// *   ApprovalHistory: queries review history.
	// *   AccountRequests: queries all plans that belong to the current Alibaba Cloud account.
	// *   AccountRequests: queries all plans that belong to the current Alibaba Cloud account.
	ApprovalFilter *string `json:"ApprovalFilter,omitempty" xml:"ApprovalFilter,omitempty"`
	// An array that consists of filter conditions.
	Filters []*ListProvisionedProductPlansRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The information based on which you want to sort the query results.
	//
	// Set the value to CreateTime, which specifies the creation time of plans.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// *   Asc: the ascending order
	// *   Desc (default): the descending order.
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProvisionedProductPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansRequest) SetAccessLevelFilter(v string) *ListProvisionedProductPlansRequest {
	s.AccessLevelFilter = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetApprovalFilter(v string) *ListProvisionedProductPlansRequest {
	s.ApprovalFilter = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetFilters(v []*ListProvisionedProductPlansRequestFilters) *ListProvisionedProductPlansRequest {
	s.Filters = v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetPageNumber(v int32) *ListProvisionedProductPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetPageSize(v int32) *ListProvisionedProductPlansRequest {
	s.PageSize = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetProvisionedProductId(v string) *ListProvisionedProductPlansRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetSortBy(v string) *ListProvisionedProductPlansRequest {
	s.SortBy = &v
	return s
}

func (s *ListProvisionedProductPlansRequest) SetSortOrder(v string) *ListProvisionedProductPlansRequest {
	s.SortOrder = &v
	return s
}

type ListProvisionedProductPlansRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// *   ProvisionedProductPlanName: performs exact matches by plan name. Plan names are not case-sensitive.
	// *   ProvisionedProductPlanApprover: performs exact matches by reviewer. You must specify a reviewer in the `RamUser/RamRole:<Name of the reviewer>` format. You can specify multiple reviewers.
	// *   ProvisionedProductPlanApproverName: performs matches by reviewer name. You must specify the Resource Access Management (RAM) entity name of the reviewer. You can specify multiple reviewer names.
	// *   ProvisionedProductPlanStatus: performs matches by plan status. You must specify the state of the plan. You can specify multiple states.
	// *   ProvisionedProductPlanOwnerUid: performs exact matches by ID of Alibaba Cloud account to which a plan belongs.
	// *   FullTextSearch: performs fuzzy full-text searches by plan name.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductPlansRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansRequestFilters) SetKey(v string) *ListProvisionedProductPlansRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductPlansRequestFilters) SetValue(v string) *ListProvisionedProductPlansRequestFilters {
	s.Value = &v
	return s
}

type ListProvisionedProductPlansResponseBody struct {
	// The page number of the returned page.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// An array that consists of plans.
	PlanDetails []*ListProvisionedProductPlansResponseBodyPlanDetails `json:"PlanDetails,omitempty" xml:"PlanDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProvisionedProductPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBody) GoString() string {
	return s.String()
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

type ListProvisionedProductPlansResponseBodyPlanDetails struct {
	// An array that consists of reviewers.
	AssignedApprovers []*ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers `json:"AssignedApprovers,omitempty" xml:"AssignedApprovers,omitempty" type:"Repeated"`
	// The time when the plan was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The purpose of the plan. Valid values:
	//
	// *   LaunchProduct: launches the product. This is the default value.
	// *   UpdateProvisionedProduct: updates the information about the product instance.
	// *   TerminateProvisionedProduct: terminates the product instance.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The ID of the RAM entity to which the plan belongs.
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The name of the RAM entity to which the plan belongs.
	OwnerPrincipalName *string `json:"OwnerPrincipalName,omitempty" xml:"OwnerPrincipalName,omitempty"`
	// The type of the RAM entity to which the plan belongs. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// An array that consists of the parameters in the template.
	Parameters []*ListProvisionedProductPlansResponseBodyPlanDetailsParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The name of the plan.
	PlanName *string `json:"PlanName,omitempty" xml:"PlanName,omitempty"`
	// The type of the plan.
	//
	// The value is fixed as Ros, which indicates Resource Orchestration Service (ROS).
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The ID of the ROS stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The state of the plan. Valid values:
	//
	// *   PreviewInProgress: The plan is being prechecked.
	// *   PreviewSuccess: The precheck is successful.
	// *   PreviewFailed: The precheck fails.
	// *   ApplicationInProgress: The plan is being reviewed.
	// *   ApplicationApproved: The plan is approved.
	// *   ApplicationRejected: The approval is rejected.
	// *   ExecuteInProgress: The plan is being run.
	// *   ExecuteSuccess: The plan is run.
	// *   ExecuteFailed: The plan fails to be run.
	// *   Canceled: The plan is canceled.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message returned for the state.
	//
	// > This parameter is returned only when PreviewFailed or ExecuteFailed is returned for the Status parameter.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// An array that consists of custom tags.
	Tags []*ListProvisionedProductPlansResponseBodyPlanDetailsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the plan belongs.
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// The last time when the task was modified.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetails) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetails) GoString() string {
	return s.String()
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

type ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers struct {
	// The RAM entity name of the reviewer.
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the RAM entity of the reviewer. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) SetPrincipalName(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers {
	s.PrincipalName = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers) SetPrincipalType(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsAssignedApprovers {
	s.PrincipalType = &v
	return s
}

type ListProvisionedProductPlansResponseBodyPlanDetailsParameters struct {
	// The name of the parameter in the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter in the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsParameters) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsParameters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) SetParameterKey(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsParameters {
	s.ParameterKey = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsParameters) SetParameterValue(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsParameters {
	s.ParameterValue = &v
	return s
}

type ListProvisionedProductPlansResponseBodyPlanDetailsTags struct {
	// The key of the custom tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsTags) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansResponseBodyPlanDetailsTags) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) SetKey(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsTags {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductPlansResponseBodyPlanDetailsTags) SetValue(v string) *ListProvisionedProductPlansResponseBodyPlanDetailsTags {
	s.Value = &v
	return s
}

type ListProvisionedProductPlansResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionedProductPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionedProductPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductPlansResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductPlansResponse) SetHeaders(v map[string]*string) *ListProvisionedProductPlansResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionedProductPlansResponse) SetStatusCode(v int32) *ListProvisionedProductPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionedProductPlansResponse) SetBody(v *ListProvisionedProductPlansResponseBody) *ListProvisionedProductPlansResponse {
	s.Body = v
	return s
}

type ListProvisionedProductsRequest struct {
	// The access filter. Valid values:
	//
	// *   User: queries the product instances that are created by the current requester. This is the default value.
	// *   Account: queries the product instances that belong to the current Alibaba Cloud account.
	AccessLevelFilter *string `json:"AccessLevelFilter,omitempty" xml:"AccessLevelFilter,omitempty"`
	// The filter conditions.
	Filters []*ListProvisionedProductsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field that is used to sort the queried data.
	//
	// Set the value to CreateTime, which specifies the time when the product instance was created.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting method. Valid values:
	//
	// *   Asc: the ascending order.
	// *   Desc (default): the descending order.
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListProvisionedProductsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductsRequest) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsRequest) SetAccessLevelFilter(v string) *ListProvisionedProductsRequest {
	s.AccessLevelFilter = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetFilters(v []*ListProvisionedProductsRequestFilters) *ListProvisionedProductsRequest {
	s.Filters = v
	return s
}

func (s *ListProvisionedProductsRequest) SetPageNumber(v int32) *ListProvisionedProductsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetPageSize(v int32) *ListProvisionedProductsRequest {
	s.PageSize = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetSortBy(v string) *ListProvisionedProductsRequest {
	s.SortBy = &v
	return s
}

func (s *ListProvisionedProductsRequest) SetSortOrder(v string) *ListProvisionedProductsRequest {
	s.SortOrder = &v
	return s
}

type ListProvisionedProductsRequestFilters struct {
	// The name of the filter condition. Valid values:
	//
	// *   ProvisionedProductName: performs exact matches by product instance name. Product instance names are not case-sensitive.
	// *   FullTextSearch: performs full-text searches by product instance name. Fuzzy match is supported.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the filter condition.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListProvisionedProductsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsRequestFilters) SetKey(v string) *ListProvisionedProductsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListProvisionedProductsRequestFilters) SetValue(v string) *ListProvisionedProductsRequestFilters {
	s.Value = &v
	return s
}

type ListProvisionedProductsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product instances.
	ProvisionedProductDetails []*ListProvisionedProductsResponseBodyProvisionedProductDetails `json:"ProvisionedProductDetails,omitempty" xml:"ProvisionedProductDetails,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProvisionedProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductsResponseBody) GoString() string {
	return s.String()
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

type ListProvisionedProductsResponseBodyProvisionedProductDetails struct {
	// The time when the product instance was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the task that was last run on the product instance.
	//
	// The task can be one of the following types:
	//
	// *   LaunchProduct: a task that launches the product.
	// *   UpdateProvisionedProduct: a task that updates the information about the product instance.
	// *   TerminateProvisionedProduct: a task that terminates the product instance.
	LastProvisioningTaskId *string `json:"LastProvisioningTaskId,omitempty" xml:"LastProvisioningTaskId,omitempty"`
	// The ID of the last task that was successfully run on the product instance.
	//
	// The task can be one of the following types:
	//
	// *   LaunchProduct: a task that launches the product.
	// *   UpdateProvisionedProduct: a task that updates the information about the product instance.
	// *   TerminateProvisionedProduct: a task that terminates the product instance.
	LastSuccessfulProvisioningTaskId *string `json:"LastSuccessfulProvisioningTaskId,omitempty" xml:"LastSuccessfulProvisioningTaskId,omitempty"`
	// The ID of the task that was last run.
	LastTaskId *string `json:"LastTaskId,omitempty" xml:"LastTaskId,omitempty"`
	// The ID of the RAM entity to which the product instance belongs.
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The type of the Resource Access Management (RAM) entity to which the product instance belongs. Valid values:
	//
	// *   RamUser: a RAM user
	// *   RamRole: a RAM role
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// The ID of the product portfolio.
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The name of the product.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the product instance.
	ProvisionedProductArn *string `json:"ProvisionedProductArn,omitempty" xml:"ProvisionedProductArn,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The name of the product instance.
	ProvisionedProductName *string `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	// The type of the product instance.
	//
	// The value is fixed as RosStack, which indicates an ROS stack.
	ProvisionedProductType *string `json:"ProvisionedProductType,omitempty" xml:"ProvisionedProductType,omitempty"`
	// The ID of the Resource Orchestration Service (ROS) stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the region to which the ROS stack belongs.
	StackRegionId *string `json:"StackRegionId,omitempty" xml:"StackRegionId,omitempty"`
	// The state of the product instance. Valid values:
	//
	// *   Available: The product instance was available.
	// *   UnderChange: The information about the product instance was being changed.
	// *   Error: An exception occurred on the product instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The message that is returned for the status of the product instance.
	//
	// > This parameter is returned only when Error is returned for the Status parameter.
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
}

func (s ListProvisionedProductsResponseBodyProvisionedProductDetails) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductsResponseBodyProvisionedProductDetails) GoString() string {
	return s.String()
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

type ListProvisionedProductsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProvisionedProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProvisionedProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProvisionedProductsResponse) GoString() string {
	return s.String()
}

func (s *ListProvisionedProductsResponse) SetHeaders(v map[string]*string) *ListProvisionedProductsResponse {
	s.Headers = v
	return s
}

func (s *ListProvisionedProductsResponse) SetStatusCode(v int32) *ListProvisionedProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProvisionedProductsResponse) SetBody(v *ListProvisionedProductsResponseBody) *ListProvisionedProductsResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	// The details of regions.
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListResourcesForTagOptionRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100 Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of resource that is associated with the tag option. Valid values:
	//
	// *   product: product
	// *   Portfolio: product portfolio
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
}

func (s ListResourcesForTagOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesForTagOptionRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionRequest) SetPageNumber(v int32) *ListResourcesForTagOptionRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) SetPageSize(v int32) *ListResourcesForTagOptionRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) SetResourceType(v string) *ListResourcesForTagOptionRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourcesForTagOptionRequest) SetTagOptionId(v string) *ListResourcesForTagOptionRequest {
	s.TagOptionId = &v
	return s
}

type ListResourcesForTagOptionResponseBody struct {
	// The page number of the returned page.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: 1 to 100 Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the associated resources.
	ResourceDetails []*ListResourcesForTagOptionResponseBodyResourceDetails `json:"ResourceDetails,omitempty" xml:"ResourceDetails,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesForTagOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesForTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionResponseBody) SetPageNumber(v int32) *ListResourcesForTagOptionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetPageSize(v int32) *ListResourcesForTagOptionResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetRequestId(v string) *ListResourcesForTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetResourceDetails(v []*ListResourcesForTagOptionResponseBodyResourceDetails) *ListResourcesForTagOptionResponseBody {
	s.ResourceDetails = v
	return s
}

func (s *ListResourcesForTagOptionResponseBody) SetTotalCount(v int32) *ListResourcesForTagOptionResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesForTagOptionResponseBodyResourceDetails struct {
	// The time when the resource was created.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the resource.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the resource.
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The ID of the resource with which the tag option is associated.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s ListResourcesForTagOptionResponseBodyResourceDetails) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesForTagOptionResponseBodyResourceDetails) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetCreateTime(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.CreateTime = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetDescription(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.Description = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetResourceArn(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.ResourceArn = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetResourceId(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.ResourceId = &v
	return s
}

func (s *ListResourcesForTagOptionResponseBodyResourceDetails) SetResourceName(v string) *ListResourcesForTagOptionResponseBodyResourceDetails {
	s.ResourceName = &v
	return s
}

type ListResourcesForTagOptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourcesForTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourcesForTagOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesForTagOptionResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesForTagOptionResponse) SetHeaders(v map[string]*string) *ListResourcesForTagOptionResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesForTagOptionResponse) SetStatusCode(v int32) *ListResourcesForTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesForTagOptionResponse) SetBody(v *ListResourcesForTagOptionResponseBody) *ListResourcesForTagOptionResponse {
	s.Body = v
	return s
}

type ListTagOptionsRequest struct {
	// The filter condition.
	Filters *ListTagOptionsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information based on which you want to sort the query results.
	//
	// Set the value to CreateTime, which specifies the creation time of tag options.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// *   Asc: the ascending order
	// *   Desc (default): the descending order
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTagOptionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagOptionsRequest) GoString() string {
	return s.String()
}

func (s *ListTagOptionsRequest) SetFilters(v *ListTagOptionsRequestFilters) *ListTagOptionsRequest {
	s.Filters = v
	return s
}

func (s *ListTagOptionsRequest) SetPageNumber(v int32) *ListTagOptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagOptionsRequest) SetPageSize(v int32) *ListTagOptionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagOptionsRequest) SetSortBy(v string) *ListTagOptionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTagOptionsRequest) SetSortOrder(v string) *ListTagOptionsRequest {
	s.SortOrder = &v
	return s
}

type ListTagOptionsRequestFilters struct {
	// Specifies whether to enable the tag option. Valid values:
	//
	// *   true (default)
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The full-text search method.
	FullTextSearch *string `json:"FullTextSearch,omitempty" xml:"FullTextSearch,omitempty"`
	// The key of the tag option.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagOptionsRequestFilters) String() string {
	return tea.Prettify(s)
}

func (s ListTagOptionsRequestFilters) GoString() string {
	return s.String()
}

func (s *ListTagOptionsRequestFilters) SetActive(v bool) *ListTagOptionsRequestFilters {
	s.Active = &v
	return s
}

func (s *ListTagOptionsRequestFilters) SetFullTextSearch(v string) *ListTagOptionsRequestFilters {
	s.FullTextSearch = &v
	return s
}

func (s *ListTagOptionsRequestFilters) SetKey(v string) *ListTagOptionsRequestFilters {
	s.Key = &v
	return s
}

func (s *ListTagOptionsRequestFilters) SetValue(v string) *ListTagOptionsRequestFilters {
	s.Value = &v
	return s
}

type ListTagOptionsShrinkRequest struct {
	// The filter condition.
	FiltersShrink *string `json:"Filters,omitempty" xml:"Filters,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information based on which you want to sort the query results.
	//
	// Set the value to CreateTime, which specifies the creation time of tag options.
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
	//
	// *   Asc: the ascending order
	// *   Desc (default): the descending order
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTagOptionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagOptionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagOptionsShrinkRequest) SetFiltersShrink(v string) *ListTagOptionsShrinkRequest {
	s.FiltersShrink = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetPageNumber(v int32) *ListTagOptionsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetPageSize(v int32) *ListTagOptionsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetSortBy(v string) *ListTagOptionsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListTagOptionsShrinkRequest) SetSortOrder(v string) *ListTagOptionsShrinkRequest {
	s.SortOrder = &v
	return s
}

type ListTagOptionsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// Valid values: 1 to 100. Minimum value: 1. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the tag option.
	TagOptionDetails []*ListTagOptionsResponseBodyTagOptionDetails `json:"TagOptionDetails,omitempty" xml:"TagOptionDetails,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagOptionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagOptionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagOptionsResponseBody) SetPageNumber(v int32) *ListTagOptionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetPageSize(v int32) *ListTagOptionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetRequestId(v string) *ListTagOptionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagOptionsResponseBody) SetTagOptionDetails(v []*ListTagOptionsResponseBodyTagOptionDetails) *ListTagOptionsResponseBody {
	s.TagOptionDetails = v
	return s
}

func (s *ListTagOptionsResponseBody) SetTotalCount(v int32) *ListTagOptionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagOptionsResponseBodyTagOptionDetails struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// *   true
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the Alibaba Cloud account to which the tag option belongs.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagOptionsResponseBodyTagOptionDetails) String() string {
	return tea.Prettify(s)
}

func (s ListTagOptionsResponseBodyTagOptionDetails) GoString() string {
	return s.String()
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetActive(v bool) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Active = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetKey(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Key = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetOwner(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Owner = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetTagOptionId(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.TagOptionId = &v
	return s
}

func (s *ListTagOptionsResponseBodyTagOptionDetails) SetValue(v string) *ListTagOptionsResponseBodyTagOptionDetails {
	s.Value = &v
	return s
}

type ListTagOptionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagOptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagOptionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagOptionsResponse) GoString() string {
	return s.String()
}

func (s *ListTagOptionsResponse) SetHeaders(v map[string]*string) *ListTagOptionsResponse {
	s.Headers = v
	return s
}

func (s *ListTagOptionsResponse) SetStatusCode(v int32) *ListTagOptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagOptionsResponse) SetBody(v *ListTagOptionsResponseBody) *ListTagOptionsResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	SortBy               *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SortOrder            *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProvisionedProductId(v string) *ListTasksRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListTasksRequest) SetSortBy(v string) *ListTasksRequest {
	s.SortBy = &v
	return s
}

func (s *ListTasksRequest) SetSortOrder(v string) *ListTasksRequest {
	s.SortOrder = &v
	return s
}

type ListTasksResponseBody struct {
	PageNumber  *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskDetails []*ListTasksResponseBodyTaskDetails `json:"TaskDetails,omitempty" xml:"TaskDetails,omitempty" type:"Repeated"`
	TotalCount  *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTaskDetails(v []*ListTasksResponseBodyTaskDetails) *ListTasksResponseBody {
	s.TaskDetails = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListTasksResponseBodyTaskDetails struct {
	CreateTime             *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Log                    *ListTasksResponseBodyTaskDetailsLog          `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	Outputs                []*ListTasksResponseBodyTaskDetailsOutputs    `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	Parameters             []*ListTasksResponseBodyTaskDetailsParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	PortfolioId            *string                                       `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	ProductId              *string                                       `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName            *string                                       `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductVersionId       *string                                       `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	ProductVersionName     *string                                       `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
	ProvisionedProductId   *string                                       `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	ProvisionedProductName *string                                       `json:"ProvisionedProductName,omitempty" xml:"ProvisionedProductName,omitempty"`
	Status                 *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusMessage          *string                                       `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// 代表资源名称的资源属性字段
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType   *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTasksResponseBodyTaskDetails) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetails) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetails) SetCreateTime(v string) *ListTasksResponseBodyTaskDetails {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetLog(v *ListTasksResponseBodyTaskDetailsLog) *ListTasksResponseBodyTaskDetails {
	s.Log = v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetOutputs(v []*ListTasksResponseBodyTaskDetailsOutputs) *ListTasksResponseBodyTaskDetails {
	s.Outputs = v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetParameters(v []*ListTasksResponseBodyTaskDetailsParameters) *ListTasksResponseBodyTaskDetails {
	s.Parameters = v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetPortfolioId(v string) *ListTasksResponseBodyTaskDetails {
	s.PortfolioId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductId(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductName(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductName = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductVersionId(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductVersionId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProductVersionName(v string) *ListTasksResponseBodyTaskDetails {
	s.ProductVersionName = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProvisionedProductId(v string) *ListTasksResponseBodyTaskDetails {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetProvisionedProductName(v string) *ListTasksResponseBodyTaskDetails {
	s.ProvisionedProductName = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetStatus(v string) *ListTasksResponseBodyTaskDetails {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetStatusMessage(v string) *ListTasksResponseBodyTaskDetails {
	s.StatusMessage = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetTaskId(v string) *ListTasksResponseBodyTaskDetails {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetTaskType(v string) *ListTasksResponseBodyTaskDetails {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetails) SetUpdateTime(v string) *ListTasksResponseBodyTaskDetails {
	s.UpdateTime = &v
	return s
}

type ListTasksResponseBodyTaskDetailsLog struct {
	TerraformLogs []*ListTasksResponseBodyTaskDetailsLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s ListTasksResponseBodyTaskDetailsLog) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsLog) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsLog) SetTerraformLogs(v []*ListTasksResponseBodyTaskDetailsLogTerraformLogs) *ListTasksResponseBodyTaskDetailsLog {
	s.TerraformLogs = v
	return s
}

type ListTasksResponseBodyTaskDetailsLogTerraformLogs struct {
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Stream  *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s ListTasksResponseBodyTaskDetailsLogTerraformLogs) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) SetCommand(v string) *ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) SetContent(v string) *ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsLogTerraformLogs) SetStream(v string) *ListTasksResponseBodyTaskDetailsLogTerraformLogs {
	s.Stream = &v
	return s
}

type ListTasksResponseBodyTaskDetailsOutputs struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OutputKey   *string `json:"OutputKey,omitempty" xml:"OutputKey,omitempty"`
	OutputValue *string `json:"OutputValue,omitempty" xml:"OutputValue,omitempty"`
}

func (s ListTasksResponseBodyTaskDetailsOutputs) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsOutputs) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) SetDescription(v string) *ListTasksResponseBodyTaskDetailsOutputs {
	s.Description = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) SetOutputKey(v string) *ListTasksResponseBodyTaskDetailsOutputs {
	s.OutputKey = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsOutputs) SetOutputValue(v string) *ListTasksResponseBodyTaskDetailsOutputs {
	s.OutputValue = &v
	return s
}

type ListTasksResponseBodyTaskDetailsParameters struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ListTasksResponseBodyTaskDetailsParameters) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTaskDetailsParameters) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTaskDetailsParameters) SetParameterKey(v string) *ListTasksResponseBodyTaskDetailsParameters {
	s.ParameterKey = &v
	return s
}

func (s *ListTasksResponseBodyTaskDetailsParameters) SetParameterValue(v string) *ListTasksResponseBodyTaskDetailsParameters {
	s.ParameterValue = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type TerminateProvisionedProductRequest struct {
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
}

func (s TerminateProvisionedProductRequest) String() string {
	return tea.Prettify(s)
}

func (s TerminateProvisionedProductRequest) GoString() string {
	return s.String()
}

func (s *TerminateProvisionedProductRequest) SetProvisionedProductId(v string) *TerminateProvisionedProductRequest {
	s.ProvisionedProductId = &v
	return s
}

type TerminateProvisionedProductResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TerminateProvisionedProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TerminateProvisionedProductResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateProvisionedProductResponseBody) SetRequestId(v string) *TerminateProvisionedProductResponseBody {
	s.RequestId = &v
	return s
}

type TerminateProvisionedProductResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateProvisionedProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateProvisionedProductResponse) String() string {
	return tea.Prettify(s)
}

func (s TerminateProvisionedProductResponse) GoString() string {
	return s.String()
}

func (s *TerminateProvisionedProductResponse) SetHeaders(v map[string]*string) *TerminateProvisionedProductResponse {
	s.Headers = v
	return s
}

func (s *TerminateProvisionedProductResponse) SetStatusCode(v int32) *TerminateProvisionedProductResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateProvisionedProductResponse) SetBody(v *TerminateProvisionedProductResponseBody) *TerminateProvisionedProductResponse {
	s.Body = v
	return s
}

type UpdateConstraintRequest struct {
	// The configurations of the constraint.
	//
	// Format: { "LocalRoleName": "\<role_name>" }.
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the constraint.
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The description of the constraint.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateConstraintRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConstraintRequest) GoString() string {
	return s.String()
}

func (s *UpdateConstraintRequest) SetConfig(v string) *UpdateConstraintRequest {
	s.Config = &v
	return s
}

func (s *UpdateConstraintRequest) SetConstraintId(v string) *UpdateConstraintRequest {
	s.ConstraintId = &v
	return s
}

func (s *UpdateConstraintRequest) SetDescription(v string) *UpdateConstraintRequest {
	s.Description = &v
	return s
}

type UpdateConstraintResponseBody struct {
	// The ID of the constraint.
	ConstraintId *string `json:"ConstraintId,omitempty" xml:"ConstraintId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConstraintResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConstraintResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConstraintResponseBody) SetConstraintId(v string) *UpdateConstraintResponseBody {
	s.ConstraintId = &v
	return s
}

func (s *UpdateConstraintResponseBody) SetRequestId(v string) *UpdateConstraintResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConstraintResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConstraintResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConstraintResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConstraintResponse) GoString() string {
	return s.String()
}

func (s *UpdateConstraintResponse) SetHeaders(v map[string]*string) *UpdateConstraintResponse {
	s.Headers = v
	return s
}

func (s *UpdateConstraintResponse) SetStatusCode(v int32) *UpdateConstraintResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConstraintResponse) SetBody(v *UpdateConstraintResponseBody) *UpdateConstraintResponse {
	s.Body = v
	return s
}

type UpdatePortfolioRequest struct {
	// 产品组合描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// 代表资源名称的资源属性字段
	PortfolioName *string `json:"PortfolioName,omitempty" xml:"PortfolioName,omitempty"`
	// 产品组合提供方
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s UpdatePortfolioRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePortfolioRequest) GoString() string {
	return s.String()
}

func (s *UpdatePortfolioRequest) SetDescription(v string) *UpdatePortfolioRequest {
	s.Description = &v
	return s
}

func (s *UpdatePortfolioRequest) SetPortfolioId(v string) *UpdatePortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *UpdatePortfolioRequest) SetPortfolioName(v string) *UpdatePortfolioRequest {
	s.PortfolioName = &v
	return s
}

func (s *UpdatePortfolioRequest) SetProviderName(v string) *UpdatePortfolioRequest {
	s.ProviderName = &v
	return s
}

type UpdatePortfolioResponseBody struct {
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePortfolioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePortfolioResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePortfolioResponseBody) SetPortfolioId(v string) *UpdatePortfolioResponseBody {
	s.PortfolioId = &v
	return s
}

func (s *UpdatePortfolioResponseBody) SetRequestId(v string) *UpdatePortfolioResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePortfolioResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePortfolioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePortfolioResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePortfolioResponse) GoString() string {
	return s.String()
}

func (s *UpdatePortfolioResponse) SetHeaders(v map[string]*string) *UpdatePortfolioResponse {
	s.Headers = v
	return s
}

func (s *UpdatePortfolioResponse) SetStatusCode(v int32) *UpdatePortfolioResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePortfolioResponse) SetBody(v *UpdatePortfolioResponseBody) *UpdatePortfolioResponse {
	s.Body = v
	return s
}

type UpdateProductRequest struct {
	// 产品描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代表资源一级ID的资源属性字段
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// 代表资源名称的资源属性字段
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// 产品提供方
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
}

func (s UpdateProductRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductRequest) SetDescription(v string) *UpdateProductRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductRequest) SetProductId(v string) *UpdateProductRequest {
	s.ProductId = &v
	return s
}

func (s *UpdateProductRequest) SetProductName(v string) *UpdateProductRequest {
	s.ProductName = &v
	return s
}

func (s *UpdateProductRequest) SetProviderName(v string) *UpdateProductRequest {
	s.ProviderName = &v
	return s
}

type UpdateProductResponseBody struct {
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductResponseBody) SetProductId(v string) *UpdateProductResponseBody {
	s.ProductId = &v
	return s
}

func (s *UpdateProductResponseBody) SetRequestId(v string) *UpdateProductResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProductResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProductResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductResponse) SetHeaders(v map[string]*string) *UpdateProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductResponse) SetStatusCode(v int32) *UpdateProductResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProductResponse) SetBody(v *UpdateProductResponseBody) *UpdateProductResponse {
	s.Body = v
	return s
}

type UpdateProductVersionRequest struct {
	// Specifies whether to enable the product version. Valid values:
	//
	// *   true: enables the product version. This is the default value.
	// *   false: disables the product version.
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The description of the product version.
	//
	// The value must be 1 to 128 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The recommendation information. Valid values:
	//
	// *   Default: No recommendation information is provided. This is the default value.
	// *   Recommended: the recommended version.
	// *   Latest: the latest version.
	// *   Deprecated: the version that is about to be discontinued.
	Guidance *string `json:"Guidance,omitempty" xml:"Guidance,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The name of the product version.
	//
	// The value must be 1 to 128 characters in length.
	ProductVersionName *string `json:"ProductVersionName,omitempty" xml:"ProductVersionName,omitempty"`
}

func (s UpdateProductVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionRequest) SetActive(v bool) *UpdateProductVersionRequest {
	s.Active = &v
	return s
}

func (s *UpdateProductVersionRequest) SetDescription(v string) *UpdateProductVersionRequest {
	s.Description = &v
	return s
}

func (s *UpdateProductVersionRequest) SetGuidance(v string) *UpdateProductVersionRequest {
	s.Guidance = &v
	return s
}

func (s *UpdateProductVersionRequest) SetProductVersionId(v string) *UpdateProductVersionRequest {
	s.ProductVersionId = &v
	return s
}

func (s *UpdateProductVersionRequest) SetProductVersionName(v string) *UpdateProductVersionRequest {
	s.ProductVersionName = &v
	return s
}

type UpdateProductVersionResponseBody struct {
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProductVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponseBody) SetProductVersionId(v string) *UpdateProductVersionResponseBody {
	s.ProductVersionId = &v
	return s
}

func (s *UpdateProductVersionResponseBody) SetRequestId(v string) *UpdateProductVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProductVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProductVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProductVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProductVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateProductVersionResponse) SetHeaders(v map[string]*string) *UpdateProductVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateProductVersionResponse) SetStatusCode(v int32) *UpdateProductVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProductVersionResponse) SetBody(v *UpdateProductVersionResponseBody) *UpdateProductVersionResponse {
	s.Body = v
	return s
}

type UpdateProvisionedProductRequest struct {
	// The input parameters of the template.
	//
	// You can specify up to 200 parameters.
	//
	// > - This parameter is optional. If you specify the Parameters parameter, you must specify the ParameterKey and ParameterValue parameters.
	// > - If the values of the ProductVersionId and Parameters parameters are not changed, you are not allowed to update the information about the product instance.
	Parameters []*UpdateProvisionedProductRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the product portfolio.
	//
	// > The PortfolioId parameter is not required if the default launch option exists. The PortfolioId parameter is required if the default launch option does not exist. For more information about how to obtain the value of the PortfolioId parameter, see [ListLaunchOptions](~~ListLaunchOptions~~).
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	//
	// > If the values of the ProductVersionId and Parameters parameters are not changed, the information about the product instance cannot be updated.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The input custom tags.
	//
	// Maximum value of N: 20.
	//
	// > - The Tags parameter is optional. If you need to specify the Tags parameter, you must specify the Tags.N.Key and Tags.N.Value parameters.
	// > - The tag is propagated to each stack resource that supports the tag feature.
	Tags []*UpdateProvisionedProductRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateProvisionedProductRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductRequest) GoString() string {
	return s.String()
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

type UpdateProvisionedProductRequestParameters struct {
	// The name of the input parameter for the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the input parameter for the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateProvisionedProductRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductRequestParameters) SetParameterKey(v string) *UpdateProvisionedProductRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateProvisionedProductRequestParameters) SetParameterValue(v string) *UpdateProvisionedProductRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateProvisionedProductRequestTags struct {
	// The tag key of the custom tag.
	//
	// The tag key must be 1 to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the custom tag.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:`. It cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateProvisionedProductRequestTags) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductRequestTags) SetKey(v string) *UpdateProvisionedProductRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateProvisionedProductRequestTags) SetValue(v string) *UpdateProvisionedProductRequestTags {
	s.Value = &v
	return s
}

type UpdateProvisionedProductResponseBody struct {
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProvisionedProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductResponseBody) SetProvisionedProductId(v string) *UpdateProvisionedProductResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *UpdateProvisionedProductResponseBody) SetRequestId(v string) *UpdateProvisionedProductResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProvisionedProductResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProvisionedProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProvisionedProductResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductResponse) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductResponse) SetHeaders(v map[string]*string) *UpdateProvisionedProductResponse {
	s.Headers = v
	return s
}

func (s *UpdateProvisionedProductResponse) SetStatusCode(v int32) *UpdateProvisionedProductResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProvisionedProductResponse) SetBody(v *UpdateProvisionedProductResponseBody) *UpdateProvisionedProductResponse {
	s.Body = v
	return s
}

type UpdateProvisionedProductPlanRequest struct {
	// The description of the plan.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// An array that consists of the parameters in the template.
	//
	// Maximum value of N: 200.
	//
	// > If you specify Parameters, you must specify ParameterKey and ParameterValue.
	Parameters []*UpdateProvisionedProductPlanRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the product portfolio.
	//
	// > If the default launch option exists, you do not need to specify PortfolioId. If the default launch option does not exist, you must specify PortfolioId. For more information about how to obtain the value of PortfolioId, see [ListLaunchOptions](~~ListLaunchOptions~~).
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the product.
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The ID of the product version.
	ProductVersionId *string `json:"ProductVersionId,omitempty" xml:"ProductVersionId,omitempty"`
	// An array that consists of custom tags.
	//
	// Maximum value of N: 20.
	//
	// >
	// *   If you specify Tags, you must specify Tags.N.Key and Tags.N.Value.
	// *   The tag of a stack is propagated to each resource that supports the tag feature in the stack.
	Tags []*UpdateProvisionedProductPlanRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdateProvisionedProductPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductPlanRequest) GoString() string {
	return s.String()
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

type UpdateProvisionedProductPlanRequestParameters struct {
	// The name of the parameter in the template.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter in the template.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateProvisionedProductPlanRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductPlanRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanRequestParameters) SetParameterKey(v string) *UpdateProvisionedProductPlanRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequestParameters) SetParameterValue(v string) *UpdateProvisionedProductPlanRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateProvisionedProductPlanRequestTags struct {
	// The key of the custom tag.
	//
	// The key can be up to 128 characters in length, and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom tag.
	//
	// The value can be up to 128 characters in length, and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateProvisionedProductPlanRequestTags) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductPlanRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanRequestTags) SetKey(v string) *UpdateProvisionedProductPlanRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateProvisionedProductPlanRequestTags) SetValue(v string) *UpdateProvisionedProductPlanRequestTags {
	s.Value = &v
	return s
}

type UpdateProvisionedProductPlanResponseBody struct {
	// The ID of the plan.
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the product instance.
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateProvisionedProductPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanResponseBody) SetPlanId(v string) *UpdateProvisionedProductPlanResponseBody {
	s.PlanId = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponseBody) SetProvisionedProductId(v string) *UpdateProvisionedProductPlanResponseBody {
	s.ProvisionedProductId = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponseBody) SetRequestId(v string) *UpdateProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

type UpdateProvisionedProductPlanResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProvisionedProductPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProvisionedProductPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateProvisionedProductPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateProvisionedProductPlanResponse) SetHeaders(v map[string]*string) *UpdateProvisionedProductPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateProvisionedProductPlanResponse) SetStatusCode(v int32) *UpdateProvisionedProductPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProvisionedProductPlanResponse) SetBody(v *UpdateProvisionedProductPlanResponseBody) *UpdateProvisionedProductPlanResponse {
	s.Body = v
	return s
}

type UpdateTagOptionRequest struct {
	// Specifies whether to enable the tag option. Valid values:
	//
	// *   true (default)
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// The value can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTagOptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTagOptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionRequest) SetActive(v bool) *UpdateTagOptionRequest {
	s.Active = &v
	return s
}

func (s *UpdateTagOptionRequest) SetTagOptionId(v string) *UpdateTagOptionRequest {
	s.TagOptionId = &v
	return s
}

func (s *UpdateTagOptionRequest) SetValue(v string) *UpdateTagOptionRequest {
	s.Value = &v
	return s
}

type UpdateTagOptionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the tag option.
	TagOptionDetail *UpdateTagOptionResponseBodyTagOptionDetail `json:"TagOptionDetail,omitempty" xml:"TagOptionDetail,omitempty" type:"Struct"`
}

func (s UpdateTagOptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionResponseBody) SetRequestId(v string) *UpdateTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTagOptionResponseBody) SetTagOptionDetail(v *UpdateTagOptionResponseBodyTagOptionDetail) *UpdateTagOptionResponseBody {
	s.TagOptionDetail = v
	return s
}

type UpdateTagOptionResponseBodyTagOptionDetail struct {
	// Indicates whether the tag option is enabled. Valid values:
	//
	// *   true (default)
	// *   false
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The key of the tag option.
	//
	// The key must be 1 to 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the Alibaba Cloud account to which the tag option belongs.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the tag option.
	TagOptionId *string `json:"TagOptionId,omitempty" xml:"TagOptionId,omitempty"`
	// The value of the tag option.
	//
	// The value must be 1 to 128 characters in length. It cannot contain `http://` or `https://` and cannot start with `acs:` or `aliyun`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateTagOptionResponseBodyTagOptionDetail) String() string {
	return tea.Prettify(s)
}

func (s UpdateTagOptionResponseBodyTagOptionDetail) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetActive(v bool) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Active = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetKey(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Key = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetOwner(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Owner = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetTagOptionId(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.TagOptionId = &v
	return s
}

func (s *UpdateTagOptionResponseBodyTagOptionDetail) SetValue(v string) *UpdateTagOptionResponseBodyTagOptionDetail {
	s.Value = &v
	return s
}

type UpdateTagOptionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTagOptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTagOptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTagOptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateTagOptionResponse) SetHeaders(v map[string]*string) *UpdateTagOptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateTagOptionResponse) SetStatusCode(v int32) *UpdateTagOptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTagOptionResponse) SetBody(v *UpdateTagOptionResponseBody) *UpdateTagOptionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("servicecatalog"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApproveProvisionedProductPlanWithOptions(request *ApproveProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *ApproveProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalAction)) {
		body["ApprovalAction"] = request.ApprovalAction
	}

	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		body["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApproveProvisionedProductPlan(request *ApproveProvisionedProductPlanRequest) (_result *ApproveProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveProvisionedProductPlanResponse{}
	_body, _err := client.ApproveProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociatePrincipalWithPortfolioWithOptions(request *AssociatePrincipalWithPortfolioRequest, runtime *util.RuntimeOptions) (_result *AssociatePrincipalWithPortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		body["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		body["PrincipalType"] = request.PrincipalType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociatePrincipalWithPortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociatePrincipalWithPortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociatePrincipalWithPortfolio(request *AssociatePrincipalWithPortfolioRequest) (_result *AssociatePrincipalWithPortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociatePrincipalWithPortfolioResponse{}
	_body, _err := client.AssociatePrincipalWithPortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateProductWithPortfolioWithOptions(request *AssociateProductWithPortfolioRequest, runtime *util.RuntimeOptions) (_result *AssociateProductWithPortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateProductWithPortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateProductWithPortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateProductWithPortfolio(request *AssociateProductWithPortfolioRequest) (_result *AssociateProductWithPortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateProductWithPortfolioResponse{}
	_body, _err := client.AssociateProductWithPortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AssociateTagOptionWithResourceWithOptions(request *AssociateTagOptionWithResourceRequest, runtime *util.RuntimeOptions) (_result *AssociateTagOptionWithResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TagOptionId)) {
		body["TagOptionId"] = request.TagOptionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssociateTagOptionWithResource"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AssociateTagOptionWithResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AssociateTagOptionWithResource(request *AssociateTagOptionWithResourceRequest) (_result *AssociateTagOptionWithResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AssociateTagOptionWithResourceResponse{}
	_body, _err := client.AssociateTagOptionWithResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelProvisionedProductPlanWithOptions(request *CancelProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *CancelProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelProvisionedProductPlan(request *CancelProvisionedProductPlanRequest) (_result *CancelProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelProvisionedProductPlanResponse{}
	_body, _err := client.CancelProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyProductWithOptions(request *CopyProductRequest, runtime *util.RuntimeOptions) (_result *CopyProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SourceProductArn)) {
		body["SourceProductArn"] = request.SourceProductArn
	}

	if !tea.BoolValue(util.IsUnset(request.TargetProductName)) {
		body["TargetProductName"] = request.TargetProductName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyProduct(request *CopyProductRequest) (_result *CopyProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyProductResponse{}
	_body, _err := client.CopyProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConstraintWithOptions(request *CreateConstraintRequest, runtime *util.RuntimeOptions) (_result *CreateConstraintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.ConstraintType)) {
		body["ConstraintType"] = request.ConstraintType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConstraint"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConstraintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConstraint(request *CreateConstraintRequest) (_result *CreateConstraintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConstraintResponse{}
	_body, _err := client.CreateConstraintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePortfolioWithOptions(request *CreatePortfolioRequest, runtime *util.RuntimeOptions) (_result *CreatePortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioName)) {
		body["PortfolioName"] = request.PortfolioName
	}

	if !tea.BoolValue(util.IsUnset(request.ProviderName)) {
		body["ProviderName"] = request.ProviderName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePortfolio(request *CreatePortfolioRequest) (_result *CreatePortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePortfolioResponse{}
	_body, _err := client.CreatePortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call the CreateProduct operation, you must call the [CreateTemplate](~~CreateTemplate~~) operation to create a template.
 *
 * @param tmpReq CreateProductRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateProductResponse
 */
func (client *Client) CreateProductWithOptions(tmpReq *CreateProductRequest, runtime *util.RuntimeOptions) (_result *CreateProductResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateProductShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ProductVersionParameters)) {
		request.ProductVersionParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProductVersionParameters, tea.String("ProductVersionParameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionParametersShrink)) {
		body["ProductVersionParameters"] = request.ProductVersionParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ProviderName)) {
		body["ProviderName"] = request.ProviderName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call the CreateProduct operation, you must call the [CreateTemplate](~~CreateTemplate~~) operation to create a template.
 *
 * @param request CreateProductRequest
 * @return CreateProductResponse
 */
func (client *Client) CreateProduct(request *CreateProductRequest) (_result *CreateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProductResponse{}
	_body, _err := client.CreateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call the CreateProductVersion operation, you must call the [CreateTemplate](~~CreateTemplate~~) operation to create a template.
 *
 * @param request CreateProductVersionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateProductVersionResponse
 */
func (client *Client) CreateProductVersionWithOptions(request *CreateProductVersionRequest, runtime *util.RuntimeOptions) (_result *CreateProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Guidance)) {
		body["Guidance"] = request.Guidance
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionName)) {
		body["ProductVersionName"] = request.ProductVersionName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateUrl)) {
		body["TemplateUrl"] = request.TemplateUrl
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProductVersion"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call the CreateProductVersion operation, you must call the [CreateTemplate](~~CreateTemplate~~) operation to create a template.
 *
 * @param request CreateProductVersionRequest
 * @return CreateProductVersionResponse
 */
func (client *Client) CreateProductVersion(request *CreateProductVersionRequest) (_result *CreateProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProductVersionResponse{}
	_body, _err := client.CreateProductVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProvisionedProductPlanWithOptions(request *CreateProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *CreateProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		body["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PlanName)) {
		body["PlanName"] = request.PlanName
	}

	if !tea.BoolValue(util.IsUnset(request.PlanType)) {
		body["PlanType"] = request.PlanType
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		body["ProductVersionId"] = request.ProductVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductName)) {
		body["ProvisionedProductName"] = request.ProvisionedProductName
	}

	if !tea.BoolValue(util.IsUnset(request.StackRegionId)) {
		body["StackRegionId"] = request.StackRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProvisionedProductPlan(request *CreateProvisionedProductPlanRequest) (_result *CreateProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProvisionedProductPlanResponse{}
	_body, _err := client.CreateProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTagOptionWithOptions(request *CreateTagOptionRequest, runtime *util.RuntimeOptions) (_result *CreateTagOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		body["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTagOption"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTagOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTagOption(request *CreateTagOptionRequest) (_result *CreateTagOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTagOptionResponse{}
	_body, _err := client.CreateTagOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		body["TemplateType"] = request.TemplateType
	}

	if !tea.BoolValue(util.IsUnset(request.TerraformVariables)) {
		body["TerraformVariables"] = request.TerraformVariables
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConstraintWithOptions(request *DeleteConstraintRequest, runtime *util.RuntimeOptions) (_result *DeleteConstraintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConstraintId)) {
		body["ConstraintId"] = request.ConstraintId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConstraint"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConstraintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConstraint(request *DeleteConstraintRequest) (_result *DeleteConstraintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConstraintResponse{}
	_body, _err := client.DeleteConstraintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePortfolioWithOptions(request *DeletePortfolioRequest, runtime *util.RuntimeOptions) (_result *DeletePortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePortfolio(request *DeletePortfolioRequest) (_result *DeletePortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePortfolioResponse{}
	_body, _err := client.DeletePortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductWithOptions(request *DeleteProductRequest, runtime *util.RuntimeOptions) (_result *DeleteProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProduct(request *DeleteProductRequest) (_result *DeleteProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProductResponse{}
	_body, _err := client.DeleteProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProductVersionWithOptions(request *DeleteProductVersionRequest, runtime *util.RuntimeOptions) (_result *DeleteProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		body["ProductVersionId"] = request.ProductVersionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProductVersion"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProductVersion(request *DeleteProductVersionRequest) (_result *DeleteProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProductVersionResponse{}
	_body, _err := client.DeleteProductVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProvisionedProductPlanWithOptions(request *DeleteProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProvisionedProductPlan(request *DeleteProvisionedProductPlanRequest) (_result *DeleteProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteProvisionedProductPlanResponse{}
	_body, _err := client.DeleteProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTagOptionWithOptions(request *DeleteTagOptionRequest, runtime *util.RuntimeOptions) (_result *DeleteTagOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagOptionId)) {
		body["TagOptionId"] = request.TagOptionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTagOption"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTagOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTagOption(request *DeleteTagOptionRequest) (_result *DeleteTagOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTagOptionResponse{}
	_body, _err := client.DeleteTagOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisAssociateTagOptionFromResourceWithOptions(request *DisAssociateTagOptionFromResourceRequest, runtime *util.RuntimeOptions) (_result *DisAssociateTagOptionFromResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TagOptionId)) {
		body["TagOptionId"] = request.TagOptionId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisAssociateTagOptionFromResource"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisAssociateTagOptionFromResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisAssociateTagOptionFromResource(request *DisAssociateTagOptionFromResourceRequest) (_result *DisAssociateTagOptionFromResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisAssociateTagOptionFromResourceResponse{}
	_body, _err := client.DisAssociateTagOptionFromResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisassociatePrincipalFromPortfolioWithOptions(request *DisassociatePrincipalFromPortfolioRequest, runtime *util.RuntimeOptions) (_result *DisassociatePrincipalFromPortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalId)) {
		body["PrincipalId"] = request.PrincipalId
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalType)) {
		body["PrincipalType"] = request.PrincipalType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisassociatePrincipalFromPortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociatePrincipalFromPortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisassociatePrincipalFromPortfolio(request *DisassociatePrincipalFromPortfolioRequest) (_result *DisassociatePrincipalFromPortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociatePrincipalFromPortfolioResponse{}
	_body, _err := client.DisassociatePrincipalFromPortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisassociateProductFromPortfolioWithOptions(request *DisassociateProductFromPortfolioRequest, runtime *util.RuntimeOptions) (_result *DisassociateProductFromPortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisassociateProductFromPortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisassociateProductFromPortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisassociateProductFromPortfolio(request *DisassociateProductFromPortfolioRequest) (_result *DisassociateProductFromPortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisassociateProductFromPortfolioResponse{}
	_body, _err := client.DisassociateProductFromPortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteProvisionedProductPlanWithOptions(request *ExecuteProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *ExecuteProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteProvisionedProductPlan(request *ExecuteProvisionedProductPlanRequest) (_result *ExecuteProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteProvisionedProductPlanResponse{}
	_body, _err := client.ExecuteProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConstraintWithOptions(request *GetConstraintRequest, runtime *util.RuntimeOptions) (_result *GetConstraintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConstraintId)) {
		query["ConstraintId"] = request.ConstraintId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConstraint"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConstraintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConstraint(request *GetConstraintRequest) (_result *GetConstraintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConstraintResponse{}
	_body, _err := client.GetConstraintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPortfolioWithOptions(request *GetPortfolioRequest, runtime *util.RuntimeOptions) (_result *GetPortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		query["PortfolioId"] = request.PortfolioId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPortfolio(request *GetPortfolioRequest) (_result *GetPortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPortfolioResponse{}
	_body, _err := client.GetPortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductAsAdminWithOptions(request *GetProductAsAdminRequest, runtime *util.RuntimeOptions) (_result *GetProductAsAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductAsAdmin"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductAsAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductAsAdmin(request *GetProductAsAdminRequest) (_result *GetProductAsAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductAsAdminResponse{}
	_body, _err := client.GetProductAsAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Make sure that you are granted the permissions to manage relevant products as a user by an administrator. For more information, see [Manage access permissions](~~405233~~).
 *
 * @param request GetProductAsEndUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetProductAsEndUserResponse
 */
func (client *Client) GetProductAsEndUserWithOptions(request *GetProductAsEndUserRequest, runtime *util.RuntimeOptions) (_result *GetProductAsEndUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductAsEndUser"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductAsEndUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Make sure that you are granted the permissions to manage relevant products as a user by an administrator. For more information, see [Manage access permissions](~~405233~~).
 *
 * @param request GetProductAsEndUserRequest
 * @return GetProductAsEndUserResponse
 */
func (client *Client) GetProductAsEndUser(request *GetProductAsEndUserRequest) (_result *GetProductAsEndUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductAsEndUserResponse{}
	_body, _err := client.GetProductAsEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProductVersionWithOptions(request *GetProductVersionRequest, runtime *util.RuntimeOptions) (_result *GetProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		query["ProductVersionId"] = request.ProductVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProductVersion"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProductVersion(request *GetProductVersionRequest) (_result *GetProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProductVersionResponse{}
	_body, _err := client.GetProductVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProvisionedProductWithOptions(request *GetProvisionedProductRequest, runtime *util.RuntimeOptions) (_result *GetProvisionedProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductId)) {
		query["ProvisionedProductId"] = request.ProvisionedProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProvisionedProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProvisionedProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProvisionedProduct(request *GetProvisionedProductRequest) (_result *GetProvisionedProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProvisionedProductResponse{}
	_body, _err := client.GetProvisionedProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetProvisionedProductPlanWithOptions(request *GetProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *GetProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProvisionedProductPlan(request *GetProvisionedProductPlanRequest) (_result *GetProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProvisionedProductPlanResponse{}
	_body, _err := client.GetProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTagOptionWithOptions(request *GetTagOptionRequest, runtime *util.RuntimeOptions) (_result *GetTagOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTagOption"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTagOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTagOption(request *GetTagOptionRequest) (_result *GetTagOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTagOptionResponse{}
	_body, _err := client.GetTagOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(request *GetTaskRequest, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTask"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTask(request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		query["ProductVersionId"] = request.ProductVersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LaunchProductWithOptions(request *LaunchProductRequest, runtime *util.RuntimeOptions) (_result *LaunchProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		body["ProductVersionId"] = request.ProductVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductName)) {
		body["ProvisionedProductName"] = request.ProvisionedProductName
	}

	if !tea.BoolValue(util.IsUnset(request.StackRegionId)) {
		body["StackRegionId"] = request.StackRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LaunchProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LaunchProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LaunchProduct(request *LaunchProductRequest) (_result *LaunchProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LaunchProductResponse{}
	_body, _err := client.LaunchProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLaunchOptionsWithOptions(request *ListLaunchOptionsRequest, runtime *util.RuntimeOptions) (_result *ListLaunchOptionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLaunchOptions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLaunchOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLaunchOptions(request *ListLaunchOptionsRequest) (_result *ListLaunchOptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLaunchOptionsResponse{}
	_body, _err := client.ListLaunchOptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPortfoliosWithOptions(request *ListPortfoliosRequest, runtime *util.RuntimeOptions) (_result *ListPortfoliosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPortfolios"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPortfoliosResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPortfolios(request *ListPortfoliosRequest) (_result *ListPortfoliosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPortfoliosResponse{}
	_body, _err := client.ListPortfoliosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPrincipalsWithOptions(request *ListPrincipalsRequest, runtime *util.RuntimeOptions) (_result *ListPrincipalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		query["PortfolioId"] = request.PortfolioId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrincipals"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrincipalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPrincipals(request *ListPrincipalsRequest) (_result *ListPrincipalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrincipalsResponse{}
	_body, _err := client.ListPrincipalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductVersionsWithOptions(request *ListProductVersionsRequest, runtime *util.RuntimeOptions) (_result *ListProductVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		query["ProductId"] = request.ProductId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductVersions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductVersions(request *ListProductVersionsRequest) (_result *ListProductVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductVersionsResponse{}
	_body, _err := client.ListProductVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProductsAsAdminWithOptions(request *ListProductsAsAdminRequest, runtime *util.RuntimeOptions) (_result *ListProductsAsAdminResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		query["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductsAsAdmin"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsAsAdminResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProductsAsAdmin(request *ListProductsAsAdminRequest) (_result *ListProductsAsAdminResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductsAsAdminResponse{}
	_body, _err := client.ListProductsAsAdminWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Make sure that you are granted the permissions to manage relevant products as a user by an administrator. For more information, see [Manage access permissions](~~405233~~).
 *
 * @param request ListProductsAsEndUserRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListProductsAsEndUserResponse
 */
func (client *Client) ListProductsAsEndUserWithOptions(request *ListProductsAsEndUserRequest, runtime *util.RuntimeOptions) (_result *ListProductsAsEndUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProductsAsEndUser"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProductsAsEndUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Make sure that you are granted the permissions to manage relevant products as a user by an administrator. For more information, see [Manage access permissions](~~405233~~).
 *
 * @param request ListProductsAsEndUserRequest
 * @return ListProductsAsEndUserResponse
 */
func (client *Client) ListProductsAsEndUser(request *ListProductsAsEndUserRequest) (_result *ListProductsAsEndUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProductsAsEndUserResponse{}
	_body, _err := client.ListProductsAsEndUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProvisionedProductPlanApproversWithOptions(request *ListProvisionedProductPlanApproversRequest, runtime *util.RuntimeOptions) (_result *ListProvisionedProductPlanApproversResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProvisionedProductPlanApprovers"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProvisionedProductPlanApproversResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProvisionedProductPlanApprovers(request *ListProvisionedProductPlanApproversRequest) (_result *ListProvisionedProductPlanApproversResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProvisionedProductPlanApproversResponse{}
	_body, _err := client.ListProvisionedProductPlanApproversWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProvisionedProductPlansWithOptions(request *ListProvisionedProductPlansRequest, runtime *util.RuntimeOptions) (_result *ListProvisionedProductPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLevelFilter)) {
		query["AccessLevelFilter"] = request.AccessLevelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovalFilter)) {
		query["ApprovalFilter"] = request.ApprovalFilter
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductId)) {
		query["ProvisionedProductId"] = request.ProvisionedProductId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProvisionedProductPlans"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProvisionedProductPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProvisionedProductPlans(request *ListProvisionedProductPlansRequest) (_result *ListProvisionedProductPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProvisionedProductPlansResponse{}
	_body, _err := client.ListProvisionedProductPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProvisionedProductsWithOptions(request *ListProvisionedProductsRequest, runtime *util.RuntimeOptions) (_result *ListProvisionedProductsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccessLevelFilter)) {
		query["AccessLevelFilter"] = request.AccessLevelFilter
	}

	if !tea.BoolValue(util.IsUnset(request.Filters)) {
		query["Filters"] = request.Filters
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProvisionedProducts"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProvisionedProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProvisionedProducts(request *ListProvisionedProductsRequest) (_result *ListProvisionedProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListProvisionedProductsResponse{}
	_body, _err := client.ListProvisionedProductsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesForTagOptionWithOptions(request *ListResourcesForTagOptionRequest, runtime *util.RuntimeOptions) (_result *ListResourcesForTagOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourcesForTagOption"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesForTagOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourcesForTagOption(request *ListResourcesForTagOptionRequest) (_result *ListResourcesForTagOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourcesForTagOptionResponse{}
	_body, _err := client.ListResourcesForTagOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagOptionsWithOptions(tmpReq *ListTagOptionsRequest, runtime *util.RuntimeOptions) (_result *ListTagOptionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagOptionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Filters)) {
		request.FiltersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Filters, tea.String("Filters"), tea.String("json"))
	}

	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagOptions"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagOptionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagOptions(request *ListTagOptionsRequest) (_result *ListTagOptionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagOptionsResponse{}
	_body, _err := client.ListTagOptionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductId)) {
		query["ProvisionedProductId"] = request.ProvisionedProductId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * After a product instance is terminated, the product instance is deleted from the product instance list. End users cannot manage the product instance throughout its lifecycle. Proceed with caution.
 *
 * @param request TerminateProvisionedProductRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TerminateProvisionedProductResponse
 */
func (client *Client) TerminateProvisionedProductWithOptions(request *TerminateProvisionedProductRequest, runtime *util.RuntimeOptions) (_result *TerminateProvisionedProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductId)) {
		body["ProvisionedProductId"] = request.ProvisionedProductId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TerminateProvisionedProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TerminateProvisionedProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * After a product instance is terminated, the product instance is deleted from the product instance list. End users cannot manage the product instance throughout its lifecycle. Proceed with caution.
 *
 * @param request TerminateProvisionedProductRequest
 * @return TerminateProvisionedProductResponse
 */
func (client *Client) TerminateProvisionedProduct(request *TerminateProvisionedProductRequest) (_result *TerminateProvisionedProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TerminateProvisionedProductResponse{}
	_body, _err := client.TerminateProvisionedProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConstraintWithOptions(request *UpdateConstraintRequest, runtime *util.RuntimeOptions) (_result *UpdateConstraintResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Config)) {
		body["Config"] = request.Config
	}

	if !tea.BoolValue(util.IsUnset(request.ConstraintId)) {
		body["ConstraintId"] = request.ConstraintId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConstraint"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConstraintResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConstraint(request *UpdateConstraintRequest) (_result *UpdateConstraintResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConstraintResponse{}
	_body, _err := client.UpdateConstraintWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePortfolioWithOptions(request *UpdatePortfolioRequest, runtime *util.RuntimeOptions) (_result *UpdatePortfolioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioName)) {
		body["PortfolioName"] = request.PortfolioName
	}

	if !tea.BoolValue(util.IsUnset(request.ProviderName)) {
		body["ProviderName"] = request.ProviderName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePortfolio"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePortfolioResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePortfolio(request *UpdatePortfolioRequest) (_result *UpdatePortfolioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePortfolioResponse{}
	_body, _err := client.UpdatePortfolioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductWithOptions(request *UpdateProductRequest, runtime *util.RuntimeOptions) (_result *UpdateProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductName)) {
		body["ProductName"] = request.ProductName
	}

	if !tea.BoolValue(util.IsUnset(request.ProviderName)) {
		body["ProviderName"] = request.ProviderName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProduct(request *UpdateProductRequest) (_result *UpdateProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProductResponse{}
	_body, _err := client.UpdateProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProductVersionWithOptions(request *UpdateProductVersionRequest, runtime *util.RuntimeOptions) (_result *UpdateProductVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Guidance)) {
		body["Guidance"] = request.Guidance
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		body["ProductVersionId"] = request.ProductVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionName)) {
		body["ProductVersionName"] = request.ProductVersionName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProductVersion"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProductVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProductVersion(request *UpdateProductVersionRequest) (_result *UpdateProductVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProductVersionResponse{}
	_body, _err := client.UpdateProductVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProvisionedProductWithOptions(request *UpdateProvisionedProductRequest, runtime *util.RuntimeOptions) (_result *UpdateProvisionedProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		body["ProductVersionId"] = request.ProductVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.ProvisionedProductId)) {
		body["ProvisionedProductId"] = request.ProvisionedProductId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProvisionedProduct"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProvisionedProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProvisionedProduct(request *UpdateProvisionedProductRequest) (_result *UpdateProvisionedProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProvisionedProductResponse{}
	_body, _err := client.UpdateProvisionedProductWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateProvisionedProductPlanWithOptions(request *UpdateProvisionedProductPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateProvisionedProductPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		body["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PortfolioId)) {
		body["PortfolioId"] = request.PortfolioId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductId)) {
		body["ProductId"] = request.ProductId
	}

	if !tea.BoolValue(util.IsUnset(request.ProductVersionId)) {
		body["ProductVersionId"] = request.ProductVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		body["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProvisionedProductPlan"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateProvisionedProductPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateProvisionedProductPlan(request *UpdateProvisionedProductPlanRequest) (_result *UpdateProvisionedProductPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateProvisionedProductPlanResponse{}
	_body, _err := client.UpdateProvisionedProductPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTagOptionWithOptions(request *UpdateTagOptionRequest, runtime *util.RuntimeOptions) (_result *UpdateTagOptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Active)) {
		body["Active"] = request.Active
	}

	if !tea.BoolValue(util.IsUnset(request.TagOptionId)) {
		body["TagOptionId"] = request.TagOptionId
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		body["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTagOption"),
		Version:     tea.String("2021-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTagOptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTagOption(request *UpdateTagOptionRequest) (_result *UpdateTagOptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTagOptionResponse{}
	_body, _err := client.UpdateTagOptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
