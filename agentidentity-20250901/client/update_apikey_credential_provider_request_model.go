// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAPIKeyCredentialProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKey(v string) *UpdateAPIKeyCredentialProviderRequest
	GetAPIKey() *string
	SetAPIKeyCredentialProviderName(v string) *UpdateAPIKeyCredentialProviderRequest
	GetAPIKeyCredentialProviderName() *string
	SetDescription(v string) *UpdateAPIKeyCredentialProviderRequest
	GetDescription() *string
}

type UpdateAPIKeyCredentialProviderRequest struct {
	// example:
	//
	// new example api key
	APIKey *string `json:"APIKey,omitempty" xml:"APIKey,omitempty"`
	// example:
	//
	// api-key-dash-scope
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
	// example:
	//
	// new example provider
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateAPIKeyCredentialProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAPIKeyCredentialProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateAPIKeyCredentialProviderRequest) GetAPIKey() *string {
	return s.APIKey
}

func (s *UpdateAPIKeyCredentialProviderRequest) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *UpdateAPIKeyCredentialProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAPIKeyCredentialProviderRequest) SetAPIKey(v string) *UpdateAPIKeyCredentialProviderRequest {
	s.APIKey = &v
	return s
}

func (s *UpdateAPIKeyCredentialProviderRequest) SetAPIKeyCredentialProviderName(v string) *UpdateAPIKeyCredentialProviderRequest {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *UpdateAPIKeyCredentialProviderRequest) SetDescription(v string) *UpdateAPIKeyCredentialProviderRequest {
	s.Description = &v
	return s
}

func (s *UpdateAPIKeyCredentialProviderRequest) Validate() error {
	return dara.Validate(s)
}
