// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserVisibleKnowledgeBasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListUserVisibleKnowledgeBasesRequest
	GetKeyword() *string
	SetTenantId(v string) *ListUserVisibleKnowledgeBasesRequest
	GetTenantId() *string
}

type ListUserVisibleKnowledgeBasesRequest struct {
	// 知识库名称或描述关键词；不传时返回全部可见知识库
	//
	// example:
	//
	// 产品知识
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListUserVisibleKnowledgeBasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserVisibleKnowledgeBasesRequest) GoString() string {
	return s.String()
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListUserVisibleKnowledgeBasesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetKeyword(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.Keyword = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) SetTenantId(v string) *ListUserVisibleKnowledgeBasesRequest {
	s.TenantId = &v
	return s
}

func (s *ListUserVisibleKnowledgeBasesRequest) Validate() error {
	return dara.Validate(s)
}
