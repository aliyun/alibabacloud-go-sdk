// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParametersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSecretParametersShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListSecretParametersShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListSecretParametersShrinkRequest
	GetNextToken() *string
	SetPath(v string) *ListSecretParametersShrinkRequest
	GetPath() *string
	SetRecursive(v bool) *ListSecretParametersShrinkRequest
	GetRecursive() *bool
	SetRegionId(v string) *ListSecretParametersShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListSecretParametersShrinkRequest
	GetResourceGroupId() *string
	SetSortField(v string) *ListSecretParametersShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListSecretParametersShrinkRequest
	GetSortOrder() *string
	SetTagsShrink(v string) *ListSecretParametersShrinkRequest
	GetTagsShrink() *string
}

type ListSecretParametersShrinkRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the parameter. **You can enter a keyword to query parameter names in fuzzy match mode.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// sPH90GZOVGC6KPDUL0FIIbEtMQHq_19S6_4O_XqA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// example:
	//
	// /path1/path2/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to query parameters from all levels of directories in the specified path. Default value: false.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The ID of the region.
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
	// The field used to sort the query results. Valid values:
	//
	// 	- Name
	//
	// 	- CreatedDate
	//
	// example:
	//
	// Name
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- Ascending
	//
	// 	- Descending (Default)
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListSecretParametersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParametersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecretParametersShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListSecretParametersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecretParametersShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *ListSecretParametersShrinkRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *ListSecretParametersShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSecretParametersShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSecretParametersShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListSecretParametersShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListSecretParametersShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListSecretParametersShrinkRequest) SetMaxResults(v int32) *ListSecretParametersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetName(v string) *ListSecretParametersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetNextToken(v string) *ListSecretParametersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetPath(v string) *ListSecretParametersShrinkRequest {
	s.Path = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetRecursive(v bool) *ListSecretParametersShrinkRequest {
	s.Recursive = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetRegionId(v string) *ListSecretParametersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetResourceGroupId(v string) *ListSecretParametersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetSortField(v string) *ListSecretParametersShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetSortOrder(v string) *ListSecretParametersShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetTagsShrink(v string) *ListSecretParametersShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
