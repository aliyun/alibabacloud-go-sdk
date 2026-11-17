// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeKnowledgeBaseRetrievalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseId(v string) *RevokeKnowledgeBaseRetrievalRequest
	GetKnowledgeBaseId() *string
	SetPrincipalId(v string) *RevokeKnowledgeBaseRetrievalRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *RevokeKnowledgeBaseRetrievalRequest
	GetPrincipalType() *string
}

type RevokeKnowledgeBaseRetrievalRequest struct {
	// The knowledge base ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The DataWorks user ID of the target user. This is neither the RAM UID nor the caller identity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000000000000000
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the principal that receives the share. This parameter is optional. If omitted, the default value is USER. Only USER is supported in the initial phase.
	//
	// example:
	//
	// USER
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s RevokeKnowledgeBaseRetrievalRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeKnowledgeBaseRetrievalRequest) GoString() string {
	return s.String()
}

func (s *RevokeKnowledgeBaseRetrievalRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RevokeKnowledgeBaseRetrievalRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *RevokeKnowledgeBaseRetrievalRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *RevokeKnowledgeBaseRetrievalRequest) SetKnowledgeBaseId(v string) *RevokeKnowledgeBaseRetrievalRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalRequest) SetPrincipalId(v string) *RevokeKnowledgeBaseRetrievalRequest {
	s.PrincipalId = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalRequest) SetPrincipalType(v string) *RevokeKnowledgeBaseRetrievalRequest {
	s.PrincipalType = &v
	return s
}

func (s *RevokeKnowledgeBaseRetrievalRequest) Validate() error {
	return dara.Validate(s)
}
