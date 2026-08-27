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
	// The new name of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	NewName *string `json:"newName,omitempty" xml:"newName,omitempty"`
	// The data source ID, which is unique within the tenant.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
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
