// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeBaseRequest
	GetClientToken() *string
	SetDescription(v string) *CreateKnowledgeBaseRequest
	GetDescription() *string
	SetName(v string) *CreateKnowledgeBaseRequest
	GetName() *string
	SetType(v string) *CreateKnowledgeBaseRequest
	GetType() *string
}

type CreateKnowledgeBaseRequest struct {
	// The idempotency token for the creation request. This parameter is optional. It accepts up to 64 characters and must contain at least one non-whitespace character. Use the same value when the same caller retries the same creation request in the same region. Reusing a token with different request content returns a conflict. If you omit the token, repeated requests may create multiple knowledge bases.
	//
	// example:
	//
	// kb-create-sales-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the knowledge base. This parameter is optional. The maximum length is 1024 characters.
	//
	// example:
	//
	// Sales metrics and business policies
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the knowledge base. This parameter cannot be empty. The maximum length is 128 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sales knowledge base
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the knowledge base. PERSONAL: personal knowledge base. ENTERPRISE: enterprise knowledge base. If this parameter is omitted, the default value is PERSONAL. Creating an enterprise knowledge base requires the corresponding enterprise knowledge base capabilities.
	//
	// example:
	//
	// PERSONAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKnowledgeBaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeBaseRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeBaseRequest) SetClientToken(v string) *CreateKnowledgeBaseRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetDescription(v string) *CreateKnowledgeBaseRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetName(v string) *CreateKnowledgeBaseRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) SetType(v string) *CreateKnowledgeBaseRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeBaseRequest) Validate() error {
	return dara.Validate(s)
}
