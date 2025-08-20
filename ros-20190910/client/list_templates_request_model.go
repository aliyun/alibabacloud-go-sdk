// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListTemplatesRequestFilters) *ListTemplatesRequest
	GetFilters() []*ListTemplatesRequestFilters
	SetIncludeTags(v string) *ListTemplatesRequest
	GetIncludeTags() *string
	SetPageNumber(v int64) *ListTemplatesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListTemplatesRequest
	GetPageSize() *int64
	SetResourceGroupId(v string) *ListTemplatesRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListTemplatesRequest
	GetShareType() *string
	SetTag(v []*ListTemplatesRequestTag) *ListTemplatesRequest
	GetTag() []*ListTemplatesRequestTag
	SetTemplateName(v string) *ListTemplatesRequest
	GetTemplateName() *string
}

type ListTemplatesRequest struct {
	// Filter.
	Filters []*ListTemplatesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// Whether to query tag information. Values:
	//
	// - Enabled: Query.
	//
	// - Disabled (default): Do not query.
	//
	// example:
	//
	// Enabled
	IncludeTags *string `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The page number of the template list.
	//
	// Start value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paginated query.
	//
	// Value range: 1~50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.
	//
	// For more information about resource groups, see [What is a Resource Group](https://help.aliyun.com/document_detail/94475.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing type of the template.
	//
	// Values:
	//
	// - Private (default): The template is owned by the user.
	//
	// - Shared: The template is shared by other users.
	//
	// - Official: The template is officially shared.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// Tags. A maximum of 20 tags are supported.
	Tag []*ListTemplatesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template. This parameter is effective only when ShareType is Private.
	//
	// The length must not exceed 255 characters and must start with a digit or a letter. It can contain digits, letters, hyphens (-), and underscores (_).
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetFilters() []*ListTemplatesRequestFilters {
	return s.Filters
}

func (s *ListTemplatesRequest) GetIncludeTags() *string {
	return s.IncludeTags
}

func (s *ListTemplatesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplatesRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListTemplatesRequest) GetTag() []*ListTemplatesRequestTag {
	return s.Tag
}

func (s *ListTemplatesRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListTemplatesRequest) SetFilters(v []*ListTemplatesRequestFilters) *ListTemplatesRequest {
	s.Filters = v
	return s
}

func (s *ListTemplatesRequest) SetIncludeTags(v string) *ListTemplatesRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNumber(v int64) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int64) *ListTemplatesRequest {
	s.PageSize = &v
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

func (s *ListTemplatesRequest) SetTag(v []*ListTemplatesRequestTag) *ListTemplatesRequest {
	s.Tag = v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesRequestFilters struct {
	// The name of the filter. You can choose one or more names for the query. Value range:
	//
	// - Categories: Template categories
	//
	// - DeployTypes: Deployment types
	//
	// - ApplicationScenes: Application scenarios
	//
	// - BasicServices: Basic services
	//
	// - ResourceTypes: Resource types
	//
	// - TemplateNames: Template names
	//
	// example:
	//
	// Categories
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTemplatesRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequestFilters) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequestFilters) GetName() *string {
	return s.Name
}

func (s *ListTemplatesRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *ListTemplatesRequestFilters) SetName(v string) *ListTemplatesRequestFilters {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequestFilters) SetValues(v []*string) *ListTemplatesRequestFilters {
	s.Values = v
	return s
}

func (s *ListTemplatesRequestFilters) Validate() error {
	return dara.Validate(s)
}

type ListTemplatesRequestTag struct {
	// The key of the tag. This parameter is effective only when ShareType is Private.
	//
	// A maximum of 20 tag keys are supported.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. This parameter is effective only when ShareType is Private.
	//
	// A maximum of 20 tag values are supported.
	//
	// example:
	//
	// deploy
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplatesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTemplatesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTemplatesRequestTag) SetKey(v string) *ListTemplatesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTemplatesRequestTag) SetValue(v string) *ListTemplatesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTemplatesRequestTag) Validate() error {
	return dara.Validate(s)
}
