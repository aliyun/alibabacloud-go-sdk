// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetName() *string
	SetParentDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseDirectoryRequest struct {
	// 新分类描述；不传表示不更新
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 待修改的分类 ID（必传）
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 新分类名称；不传表示不更新
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 新父分类 ID；不传表示不移动，传入时必须是当前租户下已存在的企业知识库目录 ID
	//
	// example:
	//
	// exampleParentDirectoryId
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// 租户ID，公共参数，缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateKnowledgeBaseDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetDescription(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetName(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetParentDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetTenantId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
