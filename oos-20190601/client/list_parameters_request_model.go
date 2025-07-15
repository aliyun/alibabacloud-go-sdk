// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListParametersRequest
	GetMaxResults() *int32
	SetName(v string) *ListParametersRequest
	GetName() *string
	SetNextToken(v string) *ListParametersRequest
	GetNextToken() *string
	SetPath(v string) *ListParametersRequest
	GetPath() *string
	SetRecursive(v bool) *ListParametersRequest
	GetRecursive() *bool
	SetRegionId(v string) *ListParametersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListParametersRequest
	GetResourceGroupId() *string
	SetShareType(v string) *ListParametersRequest
	GetShareType() *string
	SetSortField(v string) *ListParametersRequest
	GetSortField() *string
	SetSortOrder(v string) *ListParametersRequest
	GetSortOrder() *string
	SetTags(v map[string]interface{}) *ListParametersRequest
	GetTags() map[string]interface{}
	SetType(v string) *ListParametersRequest
	GetType() *string
}

type ListParametersRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParametersRequest) GoString() string {
	return s.String()
}

func (s *ListParametersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListParametersRequest) GetName() *string {
	return s.Name
}

func (s *ListParametersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListParametersRequest) GetPath() *string {
	return s.Path
}

func (s *ListParametersRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *ListParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListParametersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListParametersRequest) GetShareType() *string {
	return s.ShareType
}

func (s *ListParametersRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListParametersRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListParametersRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListParametersRequest) GetType() *string {
	return s.Type
}

func (s *ListParametersRequest) SetMaxResults(v int32) *ListParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParametersRequest) SetName(v string) *ListParametersRequest {
	s.Name = &v
	return s
}

func (s *ListParametersRequest) SetNextToken(v string) *ListParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListParametersRequest) SetPath(v string) *ListParametersRequest {
	s.Path = &v
	return s
}

func (s *ListParametersRequest) SetRecursive(v bool) *ListParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListParametersRequest) SetRegionId(v string) *ListParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListParametersRequest) SetResourceGroupId(v string) *ListParametersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersRequest) SetShareType(v string) *ListParametersRequest {
	s.ShareType = &v
	return s
}

func (s *ListParametersRequest) SetSortField(v string) *ListParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListParametersRequest) SetSortOrder(v string) *ListParametersRequest {
	s.SortOrder = &v
	return s
}

func (s *ListParametersRequest) SetTags(v map[string]interface{}) *ListParametersRequest {
	s.Tags = v
	return s
}

func (s *ListParametersRequest) SetType(v string) *ListParametersRequest {
	s.Type = &v
	return s
}

func (s *ListParametersRequest) Validate() error {
	return dara.Validate(s)
}
