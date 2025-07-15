// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListParametersShrinkRequest
	GetMaxResults() *int32
	SetName(v string) *ListParametersShrinkRequest
	GetName() *string
	SetNextToken(v string) *ListParametersShrinkRequest
	GetNextToken() *string
	SetPath(v string) *ListParametersShrinkRequest
	GetPath() *string
	SetRecursive(v bool) *ListParametersShrinkRequest
	GetRecursive() *bool
	SetRegionId(v string) *ListParametersShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListParametersShrinkRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListParametersShrinkRequest
	GetShareType() *string
	SetSortField(v string) *ListParametersShrinkRequest
	GetSortField() *string
	SetSortOrder(v string) *ListParametersShrinkRequest
	GetSortOrder() *string
	SetTagsShrink(v string) *ListParametersShrinkRequest
	GetTagsShrink() *string
	SetType(v string) *ListParametersShrinkRequest
	GetType() *string
}

type ListParametersShrinkRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter. Valid values:
	//
	// 	- Public
	//
	// 	- Private
	//
	// Default value: Private.
	//
	// example:
	//
	// ‘Private’
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
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
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParametersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListParametersShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListParametersShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListParametersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListParametersShrinkRequest) GetPath() *string {
	return s.Path
}

func (s *ListParametersShrinkRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *ListParametersShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListParametersShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListParametersShrinkRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListParametersShrinkRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListParametersShrinkRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListParametersShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *ListParametersShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListParametersShrinkRequest) SetMaxResults(v int32) *ListParametersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParametersShrinkRequest) SetName(v string) *ListParametersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListParametersShrinkRequest) SetNextToken(v string) *ListParametersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListParametersShrinkRequest) SetPath(v string) *ListParametersShrinkRequest {
	s.Path = &v
	return s
}

func (s *ListParametersShrinkRequest) SetRecursive(v bool) *ListParametersShrinkRequest {
	s.Recursive = &v
	return s
}

func (s *ListParametersShrinkRequest) SetRegionId(v string) *ListParametersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListParametersShrinkRequest) SetResourceGroupId(v string) *ListParametersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersShrinkRequest) SetShareType(v string) *ListParametersShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortField(v string) *ListParametersShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortOrder(v string) *ListParametersShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListParametersShrinkRequest) SetTagsShrink(v string) *ListParametersShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListParametersShrinkRequest) SetType(v string) *ListParametersShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListParametersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
