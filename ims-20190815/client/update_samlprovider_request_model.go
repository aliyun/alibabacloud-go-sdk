// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSAMLProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthnSignAlgo(v string) *UpdateSAMLProviderRequest
	GetAuthnSignAlgo() *string
	SetNewDescription(v string) *UpdateSAMLProviderRequest
	GetNewDescription() *string
	SetNewEncodedSAMLMetadataDocument(v string) *UpdateSAMLProviderRequest
	GetNewEncodedSAMLMetadataDocument() *string
	SetSAMLProviderName(v string) *UpdateSAMLProviderRequest
	GetSAMLProviderName() *string
}

type UpdateSAMLProviderRequest struct {
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The new description.
	//
	// >  You must specify at least one of the `NewDescription` and `NewEncodedSAMLMetadataDocument` parameters.
	//
	// example:
	//
	// This is a new provider.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new metadata file.
	//
	// >  You must specify at least one of the `NewDescription` and `NewEncodedSAMLMetadataDocument` parameters.
	//
	// example:
	//
	// PD94bWwgdmVy****
	NewEncodedSAMLMetadataDocument *string `json:"NewEncodedSAMLMetadataDocument,omitempty" xml:"NewEncodedSAMLMetadataDocument,omitempty"`
	// The name of the IdP whose information you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-provider
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
}

func (s UpdateSAMLProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderRequest) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *UpdateSAMLProviderRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateSAMLProviderRequest) GetNewEncodedSAMLMetadataDocument() *string {
	return s.NewEncodedSAMLMetadataDocument
}

func (s *UpdateSAMLProviderRequest) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *UpdateSAMLProviderRequest) SetAuthnSignAlgo(v string) *UpdateSAMLProviderRequest {
	s.AuthnSignAlgo = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetNewDescription(v string) *UpdateSAMLProviderRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetNewEncodedSAMLMetadataDocument(v string) *UpdateSAMLProviderRequest {
	s.NewEncodedSAMLMetadataDocument = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetSAMLProviderName(v string) *UpdateSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

func (s *UpdateSAMLProviderRequest) Validate() error {
	return dara.Validate(s)
}
