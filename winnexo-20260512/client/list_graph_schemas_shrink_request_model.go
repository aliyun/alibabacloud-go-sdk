// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphSchemasShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListGraphSchemasShrinkRequest
	GetKeyword() *string
	SetSemanticTagsShrink(v string) *ListGraphSchemasShrinkRequest
	GetSemanticTagsShrink() *string
	SetTenantId(v string) *ListGraphSchemasShrinkRequest
	GetTenantId() *string
}

type ListGraphSchemasShrinkRequest struct {
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
	SemanticTagsShrink *string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListGraphSchemasShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGraphSchemasShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGraphSchemasShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGraphSchemasShrinkRequest) GetSemanticTagsShrink() *string {
	return s.SemanticTagsShrink
}

func (s *ListGraphSchemasShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListGraphSchemasShrinkRequest) SetKeyword(v string) *ListGraphSchemasShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetSemanticTagsShrink(v string) *ListGraphSchemasShrinkRequest {
	s.SemanticTagsShrink = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) SetTenantId(v string) *ListGraphSchemasShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ListGraphSchemasShrinkRequest) Validate() error {
	return dara.Validate(s)
}
