// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryModulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListRegistryModulesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListRegistryModulesRequest
	GetMaxResults() *int32
	SetNamespaceName(v string) *ListRegistryModulesRequest
	GetNamespaceName() *string
	SetNextToken(v string) *ListRegistryModulesRequest
	GetNextToken() *string
	SetStatus(v string) *ListRegistryModulesRequest
	GetStatus() *string
	SetType(v string) *ListRegistryModulesRequest
	GetType() *string
}

type ListRegistryModulesRequest struct {
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
	// test_namespace
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// NFzbQCa7/yd7rAuSo5xZb54dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// system
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListRegistryModulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModulesRequest) GoString() string {
	return s.String()
}

func (s *ListRegistryModulesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListRegistryModulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegistryModulesRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListRegistryModulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegistryModulesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRegistryModulesRequest) GetType() *string {
	return s.Type
}

func (s *ListRegistryModulesRequest) SetKeyword(v string) *ListRegistryModulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListRegistryModulesRequest) SetMaxResults(v int32) *ListRegistryModulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRegistryModulesRequest) SetNamespaceName(v string) *ListRegistryModulesRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListRegistryModulesRequest) SetNextToken(v string) *ListRegistryModulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListRegistryModulesRequest) SetStatus(v string) *ListRegistryModulesRequest {
	s.Status = &v
	return s
}

func (s *ListRegistryModulesRequest) SetType(v string) *ListRegistryModulesRequest {
	s.Type = &v
	return s
}

func (s *ListRegistryModulesRequest) Validate() error {
	return dara.Validate(s)
}
