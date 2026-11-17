// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantKnowledgeBaseRetrievalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *GrantKnowledgeBaseRetrievalRequest
	GetKnowledgeBaseId() *string
	SetPrincipalId(v string) *GrantKnowledgeBaseRetrievalRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *GrantKnowledgeBaseRetrievalRequest
	GetPrincipalType() *string
}

type GrantKnowledgeBaseRetrievalRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The DataWorks user ID of the target user. This is not the RAM UID or the caller identity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000000000000
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the principal that receives the share. This parameter is optional and defaults to USER. Currently, only USER is supported.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s GrantKnowledgeBaseRetrievalRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantKnowledgeBaseRetrievalRequest) GoString() string {
	return s.String()
}

func (s *GrantKnowledgeBaseRetrievalRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GrantKnowledgeBaseRetrievalRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GrantKnowledgeBaseRetrievalRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GrantKnowledgeBaseRetrievalRequest) SetKnowledgeBaseId(v string) *GrantKnowledgeBaseRetrievalRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalRequest) SetPrincipalId(v string) *GrantKnowledgeBaseRetrievalRequest {
	s.PrincipalId = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalRequest) SetPrincipalType(v string) *GrantKnowledgeBaseRetrievalRequest {
	s.PrincipalType = &v
	return s
}

func (s *GrantKnowledgeBaseRetrievalRequest) Validate() error {
	return dara.Validate(s)
}
