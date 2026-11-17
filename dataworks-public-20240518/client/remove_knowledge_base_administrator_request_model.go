// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveKnowledgeBaseAdministratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministratorUserId(v string) *RemoveKnowledgeBaseAdministratorRequest
	GetAdministratorUserId() *string
	SetKnowledgeBaseId(v string) *RemoveKnowledgeBaseAdministratorRequest
	GetKnowledgeBaseId() *string
}

type RemoveKnowledgeBaseAdministratorRequest struct {
	// The DataWorks user ID of the administrator to be removed. This is not the identity of the caller.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000000000000
	AdministratorUserId *string `json:"AdministratorUserId,omitempty" xml:"AdministratorUserId,omitempty"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
}

func (s RemoveKnowledgeBaseAdministratorRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveKnowledgeBaseAdministratorRequest) GoString() string {
	return s.String()
}

func (s *RemoveKnowledgeBaseAdministratorRequest) GetAdministratorUserId() *string {
	return s.AdministratorUserId
}

func (s *RemoveKnowledgeBaseAdministratorRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RemoveKnowledgeBaseAdministratorRequest) SetAdministratorUserId(v string) *RemoveKnowledgeBaseAdministratorRequest {
	s.AdministratorUserId = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorRequest) SetKnowledgeBaseId(v string) *RemoveKnowledgeBaseAdministratorRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RemoveKnowledgeBaseAdministratorRequest) Validate() error {
	return dara.Validate(s)
}
