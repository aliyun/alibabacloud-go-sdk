// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryModuleVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListRegistryModuleVersionsRequest
	GetMaxResults() *int32
	SetModuleName(v string) *ListRegistryModuleVersionsRequest
	GetModuleName() *string
	SetNamespaceName(v string) *ListRegistryModuleVersionsRequest
	GetNamespaceName() *string
	SetNextToken(v string) *ListRegistryModuleVersionsRequest
	GetNextToken() *string
}

type ListRegistryModuleVersionsRequest struct {
	// The number of entries per page in a paged query. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The Registry template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ModuleName
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// The workspace name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MamespaceName
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// TRkuCaTw/VsEHrnCZgrBA0ftQSEJU/lzo2ei7MJjplg=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListRegistryModuleVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModuleVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegistryModuleVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegistryModuleVersionsRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListRegistryModuleVersionsRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListRegistryModuleVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegistryModuleVersionsRequest) SetMaxResults(v int32) *ListRegistryModuleVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRegistryModuleVersionsRequest) SetModuleName(v string) *ListRegistryModuleVersionsRequest {
	s.ModuleName = &v
	return s
}

func (s *ListRegistryModuleVersionsRequest) SetNamespaceName(v string) *ListRegistryModuleVersionsRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListRegistryModuleVersionsRequest) SetNextToken(v string) *ListRegistryModuleVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRegistryModuleVersionsRequest) Validate() error {
	return dara.Validate(s)
}
