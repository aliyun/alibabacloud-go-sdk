// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSAMLProviderResponseBody
	GetRequestId() *string
	SetSAMLProvider(v *GetSAMLProviderResponseBodySAMLProvider) *GetSAMLProviderResponseBody
	GetSAMLProvider() *GetSAMLProviderResponseBodySAMLProvider
}

type GetSAMLProviderResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SAML provider.
	SAMLProvider *GetSAMLProviderResponseBodySAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" type:"Struct"`
}

func (s GetSAMLProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSAMLProviderResponseBody) GetSAMLProvider() *GetSAMLProviderResponseBodySAMLProvider {
	return s.SAMLProvider
}

func (s *GetSAMLProviderResponseBody) SetRequestId(v string) *GetSAMLProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSAMLProviderResponseBody) SetSAMLProvider(v *GetSAMLProviderResponseBodySAMLProvider) *GetSAMLProviderResponseBody {
	s.SAMLProvider = v
	return s
}

func (s *GetSAMLProviderResponseBody) Validate() error {
	if s.SAMLProvider != nil {
		if err := s.SAMLProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSAMLProviderResponseBodySAMLProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the SAML provider.
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The signature algorithm supported by the Alibaba Cloud service provider (SP). Valid values:
	//
	// - rsa-sha256
	//
	// - rsa-sha1
	AuthnSignAlgo *string `json:"AuthnSignAlgo,omitempty" xml:"AuthnSignAlgo,omitempty"`
	// The time when the SAML provider was created.
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Base64-encoded metadata file.
	EncodedSAMLMetadataDocument *string `json:"EncodedSAMLMetadataDocument,omitempty" xml:"EncodedSAMLMetadataDocument,omitempty"`
	// The name of the SAML provider.
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty"`
	// The time when the SAML provider was last updated.
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetSAMLProviderResponseBodySAMLProvider) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLProviderResponseBodySAMLProvider) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetArn() *string {
	return s.Arn
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetAuthnSignAlgo() *string {
	return s.AuthnSignAlgo
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetDescription() *string {
	return s.Description
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetEncodedSAMLMetadataDocument() *string {
	return s.EncodedSAMLMetadataDocument
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetSAMLProviderName() *string {
	return s.SAMLProviderName
}

func (s *GetSAMLProviderResponseBodySAMLProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetArn(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.Arn = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetAuthnSignAlgo(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.AuthnSignAlgo = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetCreateDate(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.CreateDate = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetDescription(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.Description = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetEncodedSAMLMetadataDocument(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetSAMLProviderName(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) SetUpdateDate(v string) *GetSAMLProviderResponseBodySAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *GetSAMLProviderResponseBodySAMLProvider) Validate() error {
	return dara.Validate(s)
}
