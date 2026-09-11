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
	// 关键词，匹配 graphName / displayName（可选，忽略大小写）
	//
	// example:
	//
	// crm
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 语义标签过滤（可选，命中任一标签即保留）
	//
	// example:
	//
	// ["销售"]
	SemanticTagsShrink *string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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
