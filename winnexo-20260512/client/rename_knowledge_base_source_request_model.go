// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameKnowledgeBaseSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewName(v string) *RenameKnowledgeBaseSourceRequest
	GetNewName() *string
	SetSourceId(v string) *RenameKnowledgeBaseSourceRequest
	GetSourceId() *string
	SetTenantId(v string) *RenameKnowledgeBaseSourceRequest
	GetTenantId() *string
}

type RenameKnowledgeBaseSourceRequest struct {
	// 新的数据源名称
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// 数据源 ID（租户内唯一）
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

func (s RenameKnowledgeBaseSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s RenameKnowledgeBaseSourceRequest) GoString() string {
	return s.String()
}

func (s *RenameKnowledgeBaseSourceRequest) GetNewName() *string {
	return s.NewName
}

func (s *RenameKnowledgeBaseSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *RenameKnowledgeBaseSourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RenameKnowledgeBaseSourceRequest) SetNewName(v string) *RenameKnowledgeBaseSourceRequest {
	s.NewName = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetSourceId(v string) *RenameKnowledgeBaseSourceRequest {
	s.SourceId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) SetTenantId(v string) *RenameKnowledgeBaseSourceRequest {
	s.TenantId = &v
	return s
}

func (s *RenameKnowledgeBaseSourceRequest) Validate() error {
	return dara.Validate(s)
}
