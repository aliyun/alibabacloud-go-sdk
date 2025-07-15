// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSecretParametersRequest
	GetMaxResults() *int32
	SetName(v string) *ListSecretParametersRequest
	GetName() *string
	SetNextToken(v string) *ListSecretParametersRequest
	GetNextToken() *string
	SetPath(v string) *ListSecretParametersRequest
	GetPath() *string
	SetRecursive(v bool) *ListSecretParametersRequest
	GetRecursive() *bool
	SetRegionId(v string) *ListSecretParametersRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListSecretParametersRequest
	GetResourceGroupId() *string
	SetSortField(v string) *ListSecretParametersRequest
	GetSortField() *string
	SetSortOrder(v string) *ListSecretParametersRequest
	GetSortOrder() *string
	SetTags(v map[string]interface{}) *ListSecretParametersRequest
	GetTags() map[string]interface{}
}

type ListSecretParametersRequest struct {
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
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListSecretParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretParametersRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParametersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSecretParametersRequest) GetName() *string {
	return s.Name
}

func (s *ListSecretParametersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSecretParametersRequest) GetPath() *string {
	return s.Path
}

func (s *ListSecretParametersRequest) GetRecursive() *bool {
	return s.Recursive
}

func (s *ListSecretParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSecretParametersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSecretParametersRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListSecretParametersRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListSecretParametersRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListSecretParametersRequest) SetMaxResults(v int32) *ListSecretParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersRequest) SetName(v string) *ListSecretParametersRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParametersRequest) SetNextToken(v string) *ListSecretParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersRequest) SetPath(v string) *ListSecretParametersRequest {
	s.Path = &v
	return s
}

func (s *ListSecretParametersRequest) SetRecursive(v bool) *ListSecretParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListSecretParametersRequest) SetRegionId(v string) *ListSecretParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParametersRequest) SetResourceGroupId(v string) *ListSecretParametersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersRequest) SetSortField(v string) *ListSecretParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListSecretParametersRequest) SetSortOrder(v string) *ListSecretParametersRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSecretParametersRequest) SetTags(v map[string]interface{}) *ListSecretParametersRequest {
	s.Tags = v
	return s
}

func (s *ListSecretParametersRequest) Validate() error {
	return dara.Validate(s)
}
