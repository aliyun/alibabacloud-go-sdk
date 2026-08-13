// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateKnowledgeBaseDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreateKnowledgeBaseDirectoryRequest
	GetName() *string
	SetParentDirectoryId(v string) *CreateKnowledgeBaseDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreateKnowledgeBaseDirectoryRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseDirectoryRequest struct {
	// 分类描述（可选）
	//
	// example:
	//
	// 示例描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 分类名称（企业知识库目录名称）
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 父分类 ID；不传时新分类挂在企业知识库根目录下，传入时必须是当前租户下已存在的企业知识库目录 ID
	//
	// example:
	//
	// exampleParentDirectoryId
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetDescription(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetName(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetParentDirectoryId(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetTenantId(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
