// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListTemplatesShrinkRequest
	GetCategory() *string
	SetCreatedBy(v string) *ListTemplatesShrinkRequest
	GetCreatedBy() *string
	SetCreatedDateAfter(v string) *ListTemplatesShrinkRequest
	GetCreatedDateAfter() *string
	SetCreatedDateBefore(v string) *ListTemplatesShrinkRequest
	GetCreatedDateBefore() *string
	SetHasTrigger(v bool) *ListTemplatesShrinkRequest
	GetHasTrigger() *bool
	SetIsExample(v bool) *ListTemplatesShrinkRequest
	GetIsExample() *bool
	SetIsFavorite(v bool) *ListTemplatesShrinkRequest
	GetIsFavorite() *bool
	SetMaxResults(v int32) *ListTemplatesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTemplatesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTemplatesShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTemplatesShrinkRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListTemplatesShrinkRequest
	GetShareType() *string
	SetSortField(v string) *ListTemplatesShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListTemplatesShrinkRequest
	GetSortOrder() *string
	SetTagsShrink(v string) *ListTemplatesShrinkRequest
	GetTagsShrink() *string
	SetTemplateFormat(v string) *ListTemplatesShrinkRequest
	GetTemplateFormat() *string
	SetTemplateName(v string) *ListTemplatesShrinkRequest
	GetTemplateName() *string
	SetTemplateType(v string) *ListTemplatesShrinkRequest
	GetTemplateType() *string
}

type ListTemplatesShrinkRequest struct {
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
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s ListTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesShrinkRequest) GetCategory() *string {
	return s.Category
}

func (s *ListTemplatesShrinkRequest) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListTemplatesShrinkRequest) GetCreatedDateAfter() *string {
	return s.CreatedDateAfter
}

func (s *ListTemplatesShrinkRequest) GetCreatedDateBefore() *string {
	return s.CreatedDateBefore
}

func (s *ListTemplatesShrinkRequest) GetHasTrigger() *bool {
	return s.HasTrigger
}

func (s *ListTemplatesShrinkRequest) GetIsExample() *bool {
	return s.IsExample
}

func (s *ListTemplatesShrinkRequest) GetIsFavorite() *bool {
	return s.IsFavorite
}

func (s *ListTemplatesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTemplatesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTemplatesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTemplatesShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplatesShrinkRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListTemplatesShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListTemplatesShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTemplatesShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListTemplatesShrinkRequest) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *ListTemplatesShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplatesShrinkRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplatesShrinkRequest) SetCategory(v string) *ListTemplatesShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedBy(v string) *ListTemplatesShrinkRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedDateAfter(v string) *ListTemplatesShrinkRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedDateBefore(v string) *ListTemplatesShrinkRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetHasTrigger(v bool) *ListTemplatesShrinkRequest {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetIsExample(v bool) *ListTemplatesShrinkRequest {
	s.IsExample = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetIsFavorite(v bool) *ListTemplatesShrinkRequest {
	s.IsFavorite = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetMaxResults(v int32) *ListTemplatesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetNextToken(v string) *ListTemplatesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetRegionId(v string) *ListTemplatesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetResourceGroupId(v string) *ListTemplatesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetShareType(v string) *ListTemplatesShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSortField(v string) *ListTemplatesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSortOrder(v string) *ListTemplatesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTagsShrink(v string) *ListTemplatesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateFormat(v string) *ListTemplatesShrinkRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateName(v string) *ListTemplatesShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateType(v string) *ListTemplatesShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
