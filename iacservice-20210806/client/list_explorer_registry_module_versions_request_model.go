// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModuleVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListExplorerRegistryModuleVersionsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListExplorerRegistryModuleVersionsRequest
	GetMaxResults() *int32
	SetModuleName(v string) *ListExplorerRegistryModuleVersionsRequest
	GetModuleName() *string
	SetModuleVersion(v string) *ListExplorerRegistryModuleVersionsRequest
	GetModuleVersion() *string
	SetNamespaceName(v string) *ListExplorerRegistryModuleVersionsRequest
	GetNamespaceName() *string
	SetNextToken(v string) *ListExplorerRegistryModuleVersionsRequest
	GetNextToken() *string
}

type ListExplorerRegistryModuleVersionsRequest struct {
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// vpc
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// 1.11.0
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// alibaba
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// lJTuhMWkNH89zZWyYM9GjpAbgRb+bPPPwN0Q3pclzKI=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListExplorerRegistryModuleVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleVersionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListExplorerRegistryModuleVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExplorerRegistryModuleVersionsRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListExplorerRegistryModuleVersionsRequest) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListExplorerRegistryModuleVersionsRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListExplorerRegistryModuleVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExplorerRegistryModuleVersionsRequest) SetKeyword(v string) *ListExplorerRegistryModuleVersionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsRequest) SetMaxResults(v int32) *ListExplorerRegistryModuleVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsRequest) SetModuleName(v string) *ListExplorerRegistryModuleVersionsRequest {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsRequest) SetModuleVersion(v string) *ListExplorerRegistryModuleVersionsRequest {
	s.ModuleVersion = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsRequest) SetNamespaceName(v string) *ListExplorerRegistryModuleVersionsRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsRequest) SetNextToken(v string) *ListExplorerRegistryModuleVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsRequest) Validate() error {
	return dara.Validate(s)
}
