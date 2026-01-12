// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *CreateKnowledgeBaseInput
	GetCredentialName() *string
	SetDescription(v string) *CreateKnowledgeBaseInput
	GetDescription() *string
	SetKnowledgeBaseName(v string) *CreateKnowledgeBaseInput
	GetKnowledgeBaseName() *string
	SetProvider(v string) *CreateKnowledgeBaseInput
	GetProvider() *string
	SetProviderSettings(v map[string]interface{}) *CreateKnowledgeBaseInput
	GetProviderSettings() map[string]interface{}
	SetRetrieveSettings(v map[string]interface{}) *CreateKnowledgeBaseInput
	GetRetrieveSettings() map[string]interface{}
}

type CreateKnowledgeBaseInput struct {
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	KnowledgeBaseName *string `json:"knowledgeBaseName,omitempty" xml:"knowledgeBaseName,omitempty"`
	// This parameter is required.
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// This parameter is required.
	ProviderSettings map[string]interface{} `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	RetrieveSettings map[string]interface{} `json:"retrieveSettings,omitempty" xml:"retrieveSettings,omitempty"`
}

func (s CreateKnowledgeBaseInput) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseInput) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateKnowledgeBaseInput) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseInput) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *CreateKnowledgeBaseInput) GetProvider() *string {
	return s.Provider
}

func (s *CreateKnowledgeBaseInput) GetProviderSettings() map[string]interface{} {
	return s.ProviderSettings
}

func (s *CreateKnowledgeBaseInput) GetRetrieveSettings() map[string]interface{} {
	return s.RetrieveSettings
}

func (s *CreateKnowledgeBaseInput) SetCredentialName(v string) *CreateKnowledgeBaseInput {
	s.CredentialName = &v
	return s
}

func (s *CreateKnowledgeBaseInput) SetDescription(v string) *CreateKnowledgeBaseInput {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseInput) SetKnowledgeBaseName(v string) *CreateKnowledgeBaseInput {
	s.KnowledgeBaseName = &v
	return s
}

func (s *CreateKnowledgeBaseInput) SetProvider(v string) *CreateKnowledgeBaseInput {
	s.Provider = &v
	return s
}

func (s *CreateKnowledgeBaseInput) SetProviderSettings(v map[string]interface{}) *CreateKnowledgeBaseInput {
	s.ProviderSettings = v
	return s
}

func (s *CreateKnowledgeBaseInput) SetRetrieveSettings(v map[string]interface{}) *CreateKnowledgeBaseInput {
	s.RetrieveSettings = v
	return s
}

func (s *CreateKnowledgeBaseInput) Validate() error {
	return dara.Validate(s)
}
