// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseInput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *UpdateKnowledgeBaseInput
	GetCredentialName() *string
	SetDescription(v string) *UpdateKnowledgeBaseInput
	GetDescription() *string
	SetProviderSettings(v map[string]interface{}) *UpdateKnowledgeBaseInput
	GetProviderSettings() map[string]interface{}
	SetRetrieveSettings(v map[string]interface{}) *UpdateKnowledgeBaseInput
	GetRetrieveSettings() map[string]interface{}
}

type UpdateKnowledgeBaseInput struct {
	CredentialName   *string                `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description      *string                `json:"description,omitempty" xml:"description,omitempty"`
	ProviderSettings map[string]interface{} `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	RetrieveSettings map[string]interface{} `json:"retrieveSettings,omitempty" xml:"retrieveSettings,omitempty"`
}

func (s UpdateKnowledgeBaseInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseInput) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseInput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *UpdateKnowledgeBaseInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseInput) GetProviderSettings() map[string]interface{} {
	return s.ProviderSettings
}

func (s *UpdateKnowledgeBaseInput) GetRetrieveSettings() map[string]interface{} {
	return s.RetrieveSettings
}

func (s *UpdateKnowledgeBaseInput) SetCredentialName(v string) *UpdateKnowledgeBaseInput {
	s.CredentialName = &v
	return s
}

func (s *UpdateKnowledgeBaseInput) SetDescription(v string) *UpdateKnowledgeBaseInput {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseInput) SetProviderSettings(v map[string]interface{}) *UpdateKnowledgeBaseInput {
	s.ProviderSettings = v
	return s
}

func (s *UpdateKnowledgeBaseInput) SetRetrieveSettings(v map[string]interface{}) *UpdateKnowledgeBaseInput {
	s.RetrieveSettings = v
	return s
}

func (s *UpdateKnowledgeBaseInput) Validate() error {
	return dara.Validate(s)
}
