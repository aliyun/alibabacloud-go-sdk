// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListGraphSchemasRequest
	GetKeyword() *string
	SetSemanticTags(v []*string) *ListGraphSchemasRequest
	GetSemanticTags() []*string
	SetTenantId(v string) *ListGraphSchemasRequest
	GetTenantId() *string
}

type ListGraphSchemasRequest struct {
	// The keyword for fuzzy match of component data in the form.
	//
	// example:
	//
	// crm
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The semantic tags used for filtering. A graph is retained if any tag matches.
	//
	// example:
	//
	// ["Sales"]
	SemanticTags []*string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty" type:"Repeated"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGraphSchemasRequest) GetSemanticTags() []*string {
	return s.SemanticTags
}

func (s *ListGraphSchemasRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphSchemasRequest) SetKeyword(v string) *ListGraphSchemasRequest {
	s.Keyword = &v
	return s
}

func (s *ListGraphSchemasRequest) SetSemanticTags(v []*string) *ListGraphSchemasRequest {
	s.SemanticTags = v
	return s
}

func (s *ListGraphSchemasRequest) SetTenantId(v string) *ListGraphSchemasRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphSchemasRequest) Validate() error {
	return dara.Validate(s)
}
