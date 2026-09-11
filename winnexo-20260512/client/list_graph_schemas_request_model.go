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
	SemanticTags []*string `json:"semanticTags,omitempty" xml:"semanticTags,omitempty" type:"Repeated"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
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
