// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseSourceContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetContent() *string
	SetForceSync(v bool) *UpdateKnowledgeBaseSourceContentRequest
	GetForceSync() *bool
	SetSourceId(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetSourceId() *string
	SetTenantId(v string) *UpdateKnowledgeBaseSourceContentRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseSourceContentRequest struct {
	// 更新后的完整正文；允许空字符串
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例内容
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// 是否同步等待重新解析完成；默认 false，异步入队
	//
	// example:
	//
	// false
	ForceSync *bool `json:"forceSync,omitempty" xml:"forceSync,omitempty"`
	// 待编辑的企业知识库数据源 ID
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

func (s UpdateKnowledgeBaseSourceContentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseSourceContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetForceSync() *bool {
	return s.ForceSync
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeBaseSourceContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetContent(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetForceSync(v bool) *UpdateKnowledgeBaseSourceContentRequest {
	s.ForceSync = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetSourceId(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) SetTenantId(v string) *UpdateKnowledgeBaseSourceContentRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseSourceContentRequest) Validate() error {
	return dara.Validate(s)
}
