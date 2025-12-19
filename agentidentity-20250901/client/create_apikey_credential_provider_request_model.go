// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAPIKeyCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *CreateAPIKeyCredentialProviderRequest
	GetAPIKey() *string
	SetAPIKeyCredentialProviderName(v string) *CreateAPIKeyCredentialProviderRequest
	GetAPIKeyCredentialProviderName() *string
	SetDescription(v string) *CreateAPIKeyCredentialProviderRequest
	GetDescription() *string
}

type CreateAPIKeyCredentialProviderRequest struct {
	// example:
	//
	// example api key
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// example:
	//
	// api-key-dash-scope
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
	// example:
	//
	// example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAPIKeyCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAPIKeyCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateAPIKeyCredentialProviderRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *CreateAPIKeyCredentialProviderRequest) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *CreateAPIKeyCredentialProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAPIKeyCredentialProviderRequest) SetAPIKey(v string) *CreateAPIKeyCredentialProviderRequest {
	s.APIKey = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderRequest) SetAPIKeyCredentialProviderName(v string) *CreateAPIKeyCredentialProviderRequest {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderRequest) SetDescription(v string) *CreateAPIKeyCredentialProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateAPIKeyCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
