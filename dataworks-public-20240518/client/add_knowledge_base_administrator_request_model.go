// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddKnowledgeBaseAdministratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministratorUserId(v string) *AddKnowledgeBaseAdministratorRequest
	GetAdministratorUserId() *string
	SetKnowledgeBaseId(v string) *AddKnowledgeBaseAdministratorRequest
	GetKnowledgeBaseId() *string
}

type AddKnowledgeBaseAdministratorRequest struct {
	// The DataWorks user ID of the administrator to be added. This is not the identity of the caller.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000000000000
	AdministratorUserId *string `json:"AdministratorUserId,omitempty" xml:"AdministratorUserId,omitempty"`
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
}

func (s AddKnowledgeBaseAdministratorRequest) String() string {
	return dara.Prettify(s)
}

func (s AddKnowledgeBaseAdministratorRequest) GoString() string {
	return s.String()
}

func (s *AddKnowledgeBaseAdministratorRequest) GetAdministratorUserId() *string {
	return s.AdministratorUserId
}

func (s *AddKnowledgeBaseAdministratorRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *AddKnowledgeBaseAdministratorRequest) SetAdministratorUserId(v string) *AddKnowledgeBaseAdministratorRequest {
	s.AdministratorUserId = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorRequest) SetKnowledgeBaseId(v string) *AddKnowledgeBaseAdministratorRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *AddKnowledgeBaseAdministratorRequest) Validate() error {
	return dara.Validate(s)
}
