// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBase interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *KnowledgeBase
	GetCreatedAt() *string
	SetCredentialName(v string) *KnowledgeBase
	GetCredentialName() *string
	SetDescription(v string) *KnowledgeBase
	GetDescription() *string
	SetKnowledgeBaseId(v string) *KnowledgeBase
	GetKnowledgeBaseId() *string
	SetKnowledgeBaseName(v string) *KnowledgeBase
	GetKnowledgeBaseName() *string
	SetLastUpdatedAt(v string) *KnowledgeBase
	GetLastUpdatedAt() *string
	SetProvider(v string) *KnowledgeBase
	GetProvider() *string
	SetProviderSettings(v map[string]interface{}) *KnowledgeBase
	GetProviderSettings() map[string]interface{}
	SetRetrieveSettings(v map[string]interface{}) *KnowledgeBase
	GetRetrieveSettings() map[string]interface{}
}

type KnowledgeBase struct {
	CreatedAt         *string                `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	CredentialName    *string                `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	Description       *string                `json:"description,omitempty" xml:"description,omitempty"`
	KnowledgeBaseId   *string                `json:"knowledgeBaseId,omitempty" xml:"knowledgeBaseId,omitempty"`
	KnowledgeBaseName *string                `json:"knowledgeBaseName,omitempty" xml:"knowledgeBaseName,omitempty"`
	LastUpdatedAt     *string                `json:"lastUpdatedAt,omitempty" xml:"lastUpdatedAt,omitempty"`
	Provider          *string                `json:"provider,omitempty" xml:"provider,omitempty"`
	ProviderSettings  map[string]interface{} `json:"providerSettings,omitempty" xml:"providerSettings,omitempty"`
	RetrieveSettings  map[string]interface{} `json:"retrieveSettings,omitempty" xml:"retrieveSettings,omitempty"`
}

func (s KnowledgeBase) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBase) GoString() string {
	return s.String()
}

func (s *KnowledgeBase) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KnowledgeBase) GetCredentialName() *string {
	return s.CredentialName
}

func (s *KnowledgeBase) GetDescription() *string {
	return s.Description
}

func (s *KnowledgeBase) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *KnowledgeBase) GetKnowledgeBaseName() *string {
	return s.KnowledgeBaseName
}

func (s *KnowledgeBase) GetLastUpdatedAt() *string {
	return s.LastUpdatedAt
}

func (s *KnowledgeBase) GetProvider() *string {
	return s.Provider
}

func (s *KnowledgeBase) GetProviderSettings() map[string]interface{} {
	return s.ProviderSettings
}

func (s *KnowledgeBase) GetRetrieveSettings() map[string]interface{} {
	return s.RetrieveSettings
}

func (s *KnowledgeBase) SetCreatedAt(v string) *KnowledgeBase {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeBase) SetCredentialName(v string) *KnowledgeBase {
	s.CredentialName = &v
	return s
}

func (s *KnowledgeBase) SetDescription(v string) *KnowledgeBase {
	s.Description = &v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseId(v string) *KnowledgeBase {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseName(v string) *KnowledgeBase {
	s.KnowledgeBaseName = &v
	return s
}

func (s *KnowledgeBase) SetLastUpdatedAt(v string) *KnowledgeBase {
	s.LastUpdatedAt = &v
	return s
}

func (s *KnowledgeBase) SetProvider(v string) *KnowledgeBase {
	s.Provider = &v
	return s
}

func (s *KnowledgeBase) SetProviderSettings(v map[string]interface{}) *KnowledgeBase {
	s.ProviderSettings = v
	return s
}

func (s *KnowledgeBase) SetRetrieveSettings(v map[string]interface{}) *KnowledgeBase {
	s.RetrieveSettings = v
	return s
}

func (s *KnowledgeBase) Validate() error {
	return dara.Validate(s)
}
