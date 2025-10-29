// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceTypeResponseBody
	GetRequestId() *string
	SetResourceType(v *GetResourceTypeResponseBodyResourceType) *GetResourceTypeResponseBody
	GetResourceType() *GetResourceTypeResponseBodyResourceType
}

type GetResourceTypeResponseBody struct {
	// example:
	//
	// 9bcaac3c-420d-4303-87ab-7638c07b0a0b
	RequestId    *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceType *GetResourceTypeResponseBodyResourceType `json:"resourceType,omitempty" xml:"resourceType,omitempty" type:"Struct"`
}

func (s GetResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceTypeResponseBody) GetResourceType() *GetResourceTypeResponseBodyResourceType {
	return s.ResourceType
}

func (s *GetResourceTypeResponseBody) SetRequestId(v string) *GetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetResourceType(v *GetResourceTypeResponseBodyResourceType) *GetResourceTypeResponseBody {
	s.ResourceType = v
	return s
}

func (s *GetResourceTypeResponseBody) Validate() error {
	if s.ResourceType != nil {
		if err := s.ResourceType.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceTypeResponseBodyResourceType struct {
	// example:
	//
	// test
	Description *string                                              `json:"description,omitempty" xml:"description,omitempty"`
	Operations  []*GetResourceTypeResponseBodyResourceTypeOperations `json:"operations,omitempty" xml:"operations,omitempty" type:"Repeated"`
	// example:
	//
	// ECS
	Product     *string `json:"product,omitempty" xml:"product,omitempty"`
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// example:
	//
	// vpc
	ProductNameEn *string `json:"productNameEn,omitempty" xml:"productNameEn,omitempty"`
	// example:
	//
	// {}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// https://vpc.console.aliyun.com/vpc/${RegionId}/route-tables/${RouteTableId}
	ResourceDetailPageUrl *string `json:"resourceDetailPageUrl,omitempty" xml:"resourceDetailPageUrl,omitempty"`
	// example:
	//
	// https://vpc.console.aliyun.com/vpc/${RegionId}/route-tables
	ResourceListPageUrl *string `json:"resourceListPageUrl,omitempty" xml:"resourceListPageUrl,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1.227.0
	StatusStartVersion *string `json:"statusStartVersion,omitempty" xml:"statusStartVersion,omitempty"`
	// example:
	//
	// network
	Subcategory *string `json:"subcategory,omitempty" xml:"subcategory,omitempty"`
	// example:
	//
	// true
	SupportExported *bool `json:"supportExported,omitempty" xml:"supportExported,omitempty"`
	// example:
	//
	// 1.227.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// alicloud_vpc
	TerraformResourceType *string `json:"terraformResourceType,omitempty" xml:"terraformResourceType,omitempty"`
	Title                 *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s GetResourceTypeResponseBodyResourceType) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceType) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceType) GetDescription() *string {
	return s.Description
}

func (s *GetResourceTypeResponseBodyResourceType) GetOperations() []*GetResourceTypeResponseBodyResourceTypeOperations {
	return s.Operations
}

func (s *GetResourceTypeResponseBodyResourceType) GetProduct() *string {
	return s.Product
}

func (s *GetResourceTypeResponseBodyResourceType) GetProductName() *string {
	return s.ProductName
}

func (s *GetResourceTypeResponseBodyResourceType) GetProductNameEn() *string {
	return s.ProductNameEn
}

func (s *GetResourceTypeResponseBodyResourceType) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *GetResourceTypeResponseBodyResourceType) GetResourceDetailPageUrl() *string {
	return s.ResourceDetailPageUrl
}

func (s *GetResourceTypeResponseBodyResourceType) GetResourceListPageUrl() *string {
	return s.ResourceListPageUrl
}

func (s *GetResourceTypeResponseBodyResourceType) GetStatus() *string {
	return s.Status
}

func (s *GetResourceTypeResponseBodyResourceType) GetStatusStartVersion() *string {
	return s.StatusStartVersion
}

func (s *GetResourceTypeResponseBodyResourceType) GetSubcategory() *string {
	return s.Subcategory
}

func (s *GetResourceTypeResponseBodyResourceType) GetSupportExported() *bool {
	return s.SupportExported
}

func (s *GetResourceTypeResponseBodyResourceType) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *GetResourceTypeResponseBodyResourceType) GetTerraformResourceType() *string {
	return s.TerraformResourceType
}

func (s *GetResourceTypeResponseBodyResourceType) GetTitle() *string {
	return s.Title
}

func (s *GetResourceTypeResponseBodyResourceType) SetDescription(v string) *GetResourceTypeResponseBodyResourceType {
	s.Description = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetOperations(v []*GetResourceTypeResponseBodyResourceTypeOperations) *GetResourceTypeResponseBodyResourceType {
	s.Operations = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProduct(v string) *GetResourceTypeResponseBodyResourceType {
	s.Product = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProductName(v string) *GetResourceTypeResponseBodyResourceType {
	s.ProductName = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProductNameEn(v string) *GetResourceTypeResponseBodyResourceType {
	s.ProductNameEn = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetProperties(v map[string]interface{}) *GetResourceTypeResponseBodyResourceType {
	s.Properties = v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetResourceDetailPageUrl(v string) *GetResourceTypeResponseBodyResourceType {
	s.ResourceDetailPageUrl = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetResourceListPageUrl(v string) *GetResourceTypeResponseBodyResourceType {
	s.ResourceListPageUrl = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetStatus(v string) *GetResourceTypeResponseBodyResourceType {
	s.Status = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetStatusStartVersion(v string) *GetResourceTypeResponseBodyResourceType {
	s.StatusStartVersion = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetSubcategory(v string) *GetResourceTypeResponseBodyResourceType {
	s.Subcategory = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetSupportExported(v bool) *GetResourceTypeResponseBodyResourceType {
	s.SupportExported = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetTerraformProviderVersion(v string) *GetResourceTypeResponseBodyResourceType {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetTerraformResourceType(v string) *GetResourceTypeResponseBodyResourceType {
	s.TerraformResourceType = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) SetTitle(v string) *GetResourceTypeResponseBodyResourceType {
	s.Title = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceType) Validate() error {
	if s.Operations != nil {
		for _, item := range s.Operations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceTypeResponseBodyResourceTypeOperations struct {
	// example:
	//
	// CreateVSwitch
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// example:
	//
	// 2016-04-28
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// example:
	//
	// Write
	OperationType *string `json:"operationType,omitempty" xml:"operationType,omitempty"`
	// serviceCode
	//
	// example:
	//
	// Vpc
	ServiceCode *string `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
}

func (s GetResourceTypeResponseBodyResourceTypeOperations) String() string {
	return dara.Prettify(s)
}

func (s GetResourceTypeResponseBodyResourceTypeOperations) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) GetApiName() *string {
	return s.ApiName
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) GetOperationType() *string {
	return s.OperationType
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) SetApiName(v string) *GetResourceTypeResponseBodyResourceTypeOperations {
	s.ApiName = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) SetApiVersion(v string) *GetResourceTypeResponseBodyResourceTypeOperations {
	s.ApiVersion = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) SetOperationType(v string) *GetResourceTypeResponseBodyResourceTypeOperations {
	s.OperationType = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) SetServiceCode(v string) *GetResourceTypeResponseBodyResourceTypeOperations {
	s.ServiceCode = &v
	return s
}

func (s *GetResourceTypeResponseBodyResourceTypeOperations) Validate() error {
	return dara.Validate(s)
}
