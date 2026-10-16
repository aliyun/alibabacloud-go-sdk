// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecificIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetSpecificIdentityProviderResponseBody
	GetRequestId() *string
	SetSpecificIdentityProviderConfiguration(v *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) *GetSpecificIdentityProviderResponseBody
	GetSpecificIdentityProviderConfiguration() *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration
}

type GetSpecificIdentityProviderResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// AABD6E03-4B3A-5687-88FF-72232670ED0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The identity provider configuration information.
	SpecificIdentityProviderConfiguration *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration `json:"SpecificIdentityProviderConfiguration,omitempty" xml:"SpecificIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s GetSpecificIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetSpecificIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSpecificIdentityProviderResponseBody) GetSpecificIdentityProviderConfiguration() *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	return s.SpecificIdentityProviderConfiguration
}

func (s *GetSpecificIdentityProviderResponseBody) SetRequestId(v string) *GetSpecificIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSpecificIdentityProviderResponseBody) SetSpecificIdentityProviderConfiguration(v *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) *GetSpecificIdentityProviderResponseBody {
	s.SpecificIdentityProviderConfiguration = v
	return s
}

func (s *GetSpecificIdentityProviderResponseBody) Validate() error {
	if s.SpecificIdentityProviderConfiguration != nil {
		if err := s.SpecificIdentityProviderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration struct {
	// The IdP metadata content.
	//
	// example:
	//
	// {"corpId":"dingxxxxxxxxxx","appKey":"dingxxxxxxxxxx","syncCallbackStatus":"Enabled"}
	IDPMetadata *string `json:"IDPMetadata,omitempty" xml:"IDPMetadata,omitempty"`
	// The type of the identity provider.
	//
	// example:
	//
	// DingTalk
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	// The SSO enabling status.
	//
	// example:
	//
	// Enabled
	SSOStatus *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
}

func (s GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GetIDPMetadata() *string {
	return s.IDPMetadata
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) SetIDPMetadata(v string) *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	s.IDPMetadata = &v
	return s
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) SetIdentityProviderType(v string) *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	s.IdentityProviderType = &v
	return s
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) SetSSOStatus(v string) *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *GetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) Validate() error {
	return dara.Validate(s)
}
