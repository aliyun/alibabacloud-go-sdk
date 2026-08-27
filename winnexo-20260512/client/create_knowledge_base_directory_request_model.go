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
	// The description of the AI assistant.
	//
	// example:
	//
	// recorder function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// oklabs_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// wd-lxykjnnw4lyl9eq
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
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
