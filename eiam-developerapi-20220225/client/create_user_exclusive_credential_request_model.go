// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserExclusiveCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialContent(v *CreateUserExclusiveCredentialRequestCredentialContent) *CreateUserExclusiveCredentialRequest
	GetCredentialContent() *CreateUserExclusiveCredentialRequestCredentialContent
	SetCredentialIdentifier(v string) *CreateUserExclusiveCredentialRequest
	GetCredentialIdentifier() *string
	SetCredentialName(v string) *CreateUserExclusiveCredentialRequest
	GetCredentialName() *string
	SetCredentialScenarioLabel(v string) *CreateUserExclusiveCredentialRequest
	GetCredentialScenarioLabel() *string
	SetCredentialType(v string) *CreateUserExclusiveCredentialRequest
	GetCredentialType() *string
	SetDescription(v string) *CreateUserExclusiveCredentialRequest
	GetDescription() *string
}

type CreateUserExclusiveCredentialRequest struct {
	// This parameter is required.
	CredentialContent *CreateUserExclusiveCredentialRequestCredentialContent `json:"credentialContent,omitempty" xml:"credentialContent,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// credential_identifier_test
	CredentialIdentifier *string `json:"credentialIdentifier,omitempty" xml:"credentialIdentifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// credential_name
	CredentialName *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	// example:
	//
	// llm
	CredentialScenarioLabel *string `json:"credentialScenarioLabel,omitempty" xml:"credentialScenarioLabel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// api_key
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// example:
	//
	// credential_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s CreateUserExclusiveCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserExclusiveCredentialRequest) GoString() string {
	return s.String()
}

func (s *CreateUserExclusiveCredentialRequest) GetCredentialContent() *CreateUserExclusiveCredentialRequestCredentialContent {
	return s.CredentialContent
}

func (s *CreateUserExclusiveCredentialRequest) GetCredentialIdentifier() *string {
	return s.CredentialIdentifier
}

func (s *CreateUserExclusiveCredentialRequest) GetCredentialName() *string {
	return s.CredentialName
}

func (s *CreateUserExclusiveCredentialRequest) GetCredentialScenarioLabel() *string {
	return s.CredentialScenarioLabel
}

func (s *CreateUserExclusiveCredentialRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateUserExclusiveCredentialRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserExclusiveCredentialRequest) SetCredentialContent(v *CreateUserExclusiveCredentialRequestCredentialContent) *CreateUserExclusiveCredentialRequest {
	s.CredentialContent = v
	return s
}

func (s *CreateUserExclusiveCredentialRequest) SetCredentialIdentifier(v string) *CreateUserExclusiveCredentialRequest {
	s.CredentialIdentifier = &v
	return s
}

func (s *CreateUserExclusiveCredentialRequest) SetCredentialName(v string) *CreateUserExclusiveCredentialRequest {
	s.CredentialName = &v
	return s
}

func (s *CreateUserExclusiveCredentialRequest) SetCredentialScenarioLabel(v string) *CreateUserExclusiveCredentialRequest {
	s.CredentialScenarioLabel = &v
	return s
}

func (s *CreateUserExclusiveCredentialRequest) SetCredentialType(v string) *CreateUserExclusiveCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *CreateUserExclusiveCredentialRequest) SetDescription(v string) *CreateUserExclusiveCredentialRequest {
	s.Description = &v
	return s
}

func (s *CreateUserExclusiveCredentialRequest) Validate() error {
	if s.CredentialContent != nil {
		if err := s.CredentialContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserExclusiveCredentialRequestCredentialContent struct {
	ApiKeyContent *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent `json:"apiKeyContent,omitempty" xml:"apiKeyContent,omitempty" type:"Struct"`
}

func (s CreateUserExclusiveCredentialRequestCredentialContent) String() string {
	return dara.Prettify(s)
}

func (s CreateUserExclusiveCredentialRequestCredentialContent) GoString() string {
	return s.String()
}

func (s *CreateUserExclusiveCredentialRequestCredentialContent) GetApiKeyContent() *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent {
	return s.ApiKeyContent
}

func (s *CreateUserExclusiveCredentialRequestCredentialContent) SetApiKeyContent(v *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent) *CreateUserExclusiveCredentialRequestCredentialContent {
	s.ApiKeyContent = v
	return s
}

func (s *CreateUserExclusiveCredentialRequestCredentialContent) Validate() error {
	if s.ApiKeyContent != nil {
		if err := s.ApiKeyContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent struct {
	// This parameter is required.
	//
	// example:
	//
	// sk-nsklncmwizncxxxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
}

func (s CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent) GoString() string {
	return s.String()
}

func (s *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent) SetApiKey(v string) *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent {
	s.ApiKey = &v
	return s
}

func (s *CreateUserExclusiveCredentialRequestCredentialContentApiKeyContent) Validate() error {
	return dara.Validate(s)
}
