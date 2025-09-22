// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSAMLProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthnSignAlgo(v string) *CreateSAMLProviderRequest
	GetAuthnSignAlgo() *string
	SetDescription(v string) *CreateSAMLProviderRequest
	GetDescription() *string
	SetEncodedSAMLMetadataDocument(v string) *CreateSAMLProviderRequest
	GetEncodedSAMLMetadataDocument() *string
	SetSAMLProviderName(v string) *CreateSAMLProviderRequest
	GetSAMLProviderName() *string
}

type CreateSAMLProviderRequest struct {
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata file which is Base64-encoded.
	//
	// The file is provided by an IdP that supports Security Assertion Markup Language (SAML) 2.0.
	//
	// example:
	//
	// PD94bWwgdmVy****
	EncodedSAMLMetadataDocument *string `json:"EncodedSAMLMetadataDocument,omitempty" xml:"EncodedSAMLMetadataDocument,omitempty"`
	// The name of the IdP.
	//
	// The name can be up to 128 characters in length. The name can contain letters, digits, `periods (.), hyphens (-), and underscores (_)`. The name cannot start or end with `periods (.), hyphens (-), or underscores (_)`.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s CreateSAMLProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderRequest) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *CreateSAMLProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSAMLProviderRequest) GetEncodedSAMLMetadataDocument() *string {
	return s.EncodedSAMLMetadataDocument
}

func (s *CreateSAMLProviderRequest) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *CreateSAMLProviderRequest) SetAuthnSignAlgo(v string) *CreateSAMLProviderRequest {
	s.AuthnSignAlgo = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetDescription(v string) *CreateSAMLProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetEncodedSAMLMetadataDocument(v string) *CreateSAMLProviderRequest {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetSAMLProviderName(v string) *CreateSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

func (s *CreateSAMLProviderRequest) Validate() error {
	return dara.Validate(s)
}
