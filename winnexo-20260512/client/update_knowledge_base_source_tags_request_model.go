// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetSourceId() *string
	SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetSourceTags() *string
	SetTenantId(v string) *UpdateKnowledgeBaseSourceTagsRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseSourceTagsRequest struct {
	// 数据源 ID（租户内唯一）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 资源标签（JSON 字符串列表，如 ["tagA","tagB"]；传 null 表示清空标签）
	//
	// example:
	//
	// string_value
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateKnowledgeBaseSourceTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceTagsRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetSourceId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetSourceTags(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.SourceTags = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) SetTenantId(v string) *UpdateKnowledgeBaseSourceTagsRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceTagsRequest) Validate() error {
	return dara.Validate(s)
}
