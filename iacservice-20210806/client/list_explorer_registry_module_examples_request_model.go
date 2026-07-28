// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModuleExamplesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExampleName(v string) *ListExplorerRegistryModuleExamplesRequest
	GetExampleName() *string
	SetKeyword(v string) *ListExplorerRegistryModuleExamplesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListExplorerRegistryModuleExamplesRequest
	GetMaxResults() *int32
	SetModuleName(v string) *ListExplorerRegistryModuleExamplesRequest
	GetModuleName() *string
	SetModuleVersion(v string) *ListExplorerRegistryModuleExamplesRequest
	GetModuleVersion() *string
	SetNamespaceName(v string) *ListExplorerRegistryModuleExamplesRequest
	GetNamespaceName() *string
	SetNextToken(v string) *ListExplorerRegistryModuleExamplesRequest
	GetNextToken() *string
}

type ListExplorerRegistryModuleExamplesRequest struct {
	// The example name of the module.
	//
	// example:
	//
	// complete
	ExampleName *string `json:"exampleName,omitempty" xml:"exampleName,omitempty"`
	// The search keyword. Supports fuzzy match based on the module name or module example name.
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
	// example:
	//
	// sls
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The version of the module.
	//
	// example:
	//
	// 1.11.0
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// The name of the workspace to which the module belongs.
	//
	// example:
	//
	// alibaba
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// The pagination token for the next page of results.
	//
	// If the total number of entries exceeds the maxResults limit, the data is truncated. You can use nextToken to query the next page of data.
	//
	// example:
	//
	// IbuvZbAXFOiB4nKg8iOH447bhHWDavGTOMijI2Jep7c=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExplorerRegistryModuleExamplesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleExamplesRequest) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetExampleName() *string {
	return s.ExampleName
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListExplorerRegistryModuleExamplesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetExampleName(v string) *ListExplorerRegistryModuleExamplesRequest {
	s.ExampleName = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetKeyword(v string) *ListExplorerRegistryModuleExamplesRequest {
	s.Keyword = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetMaxResults(v int32) *ListExplorerRegistryModuleExamplesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetModuleName(v string) *ListExplorerRegistryModuleExamplesRequest {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetModuleVersion(v string) *ListExplorerRegistryModuleExamplesRequest {
	s.ModuleVersion = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetNamespaceName(v string) *ListExplorerRegistryModuleExamplesRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) SetNextToken(v string) *ListExplorerRegistryModuleExamplesRequest {
	s.NextToken = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesRequest) Validate() error {
	return dara.Validate(s)
}
