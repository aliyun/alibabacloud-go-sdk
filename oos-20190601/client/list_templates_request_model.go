// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListTemplatesRequest
	GetCategory() *string
	SetCreatedBy(v string) *ListTemplatesRequest
	GetCreatedBy() *string
	SetCreatedDateAfter(v string) *ListTemplatesRequest
	GetCreatedDateAfter() *string
	SetCreatedDateBefore(v string) *ListTemplatesRequest
	GetCreatedDateBefore() *string
	SetHasTrigger(v bool) *ListTemplatesRequest
	GetHasTrigger() *bool
	SetIsExample(v bool) *ListTemplatesRequest
	GetIsExample() *bool
	SetIsFavorite(v bool) *ListTemplatesRequest
	GetIsFavorite() *bool
	SetMaxResults(v int32) *ListTemplatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTemplatesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTemplatesRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListTemplatesRequest
	GetShareType() *string
	SetSortField(v string) *ListTemplatesRequest
	GetSortField() *string
	SetSortOrder(v string) *ListTemplatesRequest
	GetSortOrder() *string
	SetTags(v map[string]interface{}) *ListTemplatesRequest
	GetTags() map[string]interface{}
	SetTemplateFormat(v string) *ListTemplatesRequest
	GetTemplateFormat() *string
	SetTemplateName(v string) *ListTemplatesRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ListTemplatesRequest
	GetTemplateType() *string
}

type ListTemplatesRequest struct {
	// The type of the template. Valid values include TimerTrigger, EventTrigger, AlarmTrigger, and Other.
	//
	// example:
	//
	// TimerTrigger
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creator of the template.
	//
	// 	- To query the template provided by Alibaba Cloud, set this parameter to **ACS**.
	//
	// 	- To query the template created by a user, set this parameter to the **ID*	- of the template or the **name of the user*	- who creates the template.
	//
	// example:
	//
	// ACS
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// Specifies to query the template that is created at or later than the specified time.
	//
	// The value must be in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDateAfter *string `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	// Specifies to query the template that is created at or before the specified time.
	//
	// The value must be in the YYYY-MM-DDThh:mm::ssZ format.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDateBefore *string `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	// Specifies whether to query the template that is configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// Specifies whether the template is an example template.
	//
	// example:
	//
	// false
	IsExample *bool `json:"IsExample,omitempty" xml:"IsExample,omitempty"`
	// Specifies whether the template is added to favorites.
	//
	// example:
	//
	// true
	IsFavorite *bool `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which you want to query templates.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The field that is used to sort the templates to be queried. Valid values:
	//
	// 	- **TotalExecutionCount*	- (default): The system sorts the returned templates based on the total number of times that the templates are used.
	//
	// 	- **Popularity**: The system sorts the returned templates based on the popularity of the templates.
	//
	// 	- **TemplateName**: The system sorts the returned templates based on the names of the templates.
	//
	// 	- **CreatedDate**: The system sorts the returned templates based on the points in time when the templates are created.
	//
	// 	- **UpdatedDate**: The system sorts the returned templates based on the points in time when the templates are updated.
	//
	// example:
	//
	// Popularity
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. Valid values:
	//
	// 	- **JSON**
	//
	// 	- **YAML**
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are to be returned.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template. Valid values:
	//
	// 	- Automation: the template for automated tasks.
	//
	// 	- State: the template for configuration inventories.
	//
	// 	- Package: the template for software packages.
	//
	// example:
	//
	// private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListTemplatesRequest) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListTemplatesRequest) GetCreatedDateAfter() *string {
	return s.CreatedDateAfter
}

func (s *ListTemplatesRequest) GetCreatedDateBefore() *string {
	return s.CreatedDateBefore
}

func (s *ListTemplatesRequest) GetHasTrigger() *bool {
	return s.HasTrigger
}

func (s *ListTemplatesRequest) GetIsExample() *bool {
	return s.IsExample
}

func (s *ListTemplatesRequest) GetIsFavorite() *bool {
	return s.IsFavorite
}

func (s *ListTemplatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTemplatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplatesRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListTemplatesRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTemplatesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTemplatesRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListTemplatesRequest) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *ListTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplatesRequest) SetCategory(v string) *ListTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedBy(v string) *ListTemplatesRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedDateAfter(v string) *ListTemplatesRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedDateBefore(v string) *ListTemplatesRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListTemplatesRequest) SetHasTrigger(v bool) *ListTemplatesRequest {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesRequest) SetIsExample(v bool) *ListTemplatesRequest {
	s.IsExample = &v
	return s
}

func (s *ListTemplatesRequest) SetIsFavorite(v bool) *ListTemplatesRequest {
	s.IsFavorite = &v
	return s
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetRegionId(v string) *ListTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesRequest) SetResourceGroupId(v string) *ListTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesRequest) SetShareType(v string) *ListTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesRequest) SetSortField(v string) *ListTemplatesRequest {
	s.SortField = &v
	return s
}

func (s *ListTemplatesRequest) SetSortOrder(v string) *ListTemplatesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTemplatesRequest) SetTags(v map[string]interface{}) *ListTemplatesRequest {
	s.Tags = v
	return s
}

func (s *ListTemplatesRequest) SetTemplateFormat(v string) *ListTemplatesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
