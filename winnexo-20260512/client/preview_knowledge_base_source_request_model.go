// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewKnowledgeBaseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceId(v string) *PreviewKnowledgeBaseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *PreviewKnowledgeBaseSourceRequest
	GetTenantId() *string
}

type PreviewKnowledgeBaseSourceRequest struct {
	// 知识 ID（数据源唯一标识）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s PreviewKnowledgeBaseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewKnowledgeBaseSourceRequest) GoString() string {
	return s.String()
}

func (s *PreviewKnowledgeBaseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *PreviewKnowledgeBaseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *PreviewKnowledgeBaseSourceRequest) SetSourceId(v string) *PreviewKnowledgeBaseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) SetTenantId(v string) *PreviewKnowledgeBaseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *PreviewKnowledgeBaseSourceRequest) Validate() error {
	return dara.Validate(s)
}
