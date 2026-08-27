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
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleParentDirectoryId
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID to take effect.
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
