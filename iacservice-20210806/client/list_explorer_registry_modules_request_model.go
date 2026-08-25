// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListExplorerRegistryModulesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListExplorerRegistryModulesRequest
	GetMaxResults() *int32
	SetModuleName(v string) *ListExplorerRegistryModulesRequest
	GetModuleName() *string
	SetNextToken(v string) *ListExplorerRegistryModulesRequest
	GetNextToken() *string
	SetSort(v string) *ListExplorerRegistryModulesRequest
	GetSort() *string
}

type ListExplorerRegistryModulesRequest struct {
	// The search keyword for the module name. Fuzzy matching is performed based on `moduleName`.
	//
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 0 to 200.
	//
	// Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the module.
	//
	// Note: The module name is not necessarily the same as the product name or resource name.
	//
	// example:
	//
	// vpc
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The pagination token for the next page of results.
	//
	// If the total number of entries exceeds the maxResults limit, the data is truncated. You can use nextToken to query the next page of data.
	//
	// example:
	//
	// DxEkv+3w0EDAQgcRFBp8Ep4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The sort order of the returned results. Valid values:
	//
	// - Normal (default): returns results in normal order.
	//
	// - Top: returns results sorted by popularity.
	//
	// example:
	//
	// Normal
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
}

func (s ListExplorerRegistryModulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModulesRequest) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListExplorerRegistryModulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExplorerRegistryModulesRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListExplorerRegistryModulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExplorerRegistryModulesRequest) GetSort() *string {
	return s.Sort
}

func (s *ListExplorerRegistryModulesRequest) SetKeyword(v string) *ListExplorerRegistryModulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListExplorerRegistryModulesRequest) SetMaxResults(v int32) *ListExplorerRegistryModulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerRegistryModulesRequest) SetModuleName(v string) *ListExplorerRegistryModulesRequest {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerRegistryModulesRequest) SetNextToken(v string) *ListExplorerRegistryModulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExplorerRegistryModulesRequest) SetSort(v string) *ListExplorerRegistryModulesRequest {
	s.Sort = &v
	return s
}

func (s *ListExplorerRegistryModulesRequest) Validate() error {
	return dara.Validate(s)
}
