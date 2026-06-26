// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSpecificIdentityProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSpecificIdentityProviderResponseBody
	GetRequestId() *string
	SetSpecificIdentityProviderConfiguration(v *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) *SetSpecificIdentityProviderResponseBody
	GetSpecificIdentityProviderConfiguration() *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration
}

type SetSpecificIdentityProviderResponseBody struct {
	RequestId                             *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SpecificIdentityProviderConfiguration *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration `json:"SpecificIdentityProviderConfiguration,omitempty" xml:"SpecificIdentityProviderConfiguration,omitempty" type:"Struct"`
}

func (s SetSpecificIdentityProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSpecificIdentityProviderResponseBody) GoString() string {
	return s.String()
}

func (s *SetSpecificIdentityProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSpecificIdentityProviderResponseBody) GetSpecificIdentityProviderConfiguration() *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	return s.SpecificIdentityProviderConfiguration
}

func (s *SetSpecificIdentityProviderResponseBody) SetRequestId(v string) *SetSpecificIdentityProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSpecificIdentityProviderResponseBody) SetSpecificIdentityProviderConfiguration(v *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) *SetSpecificIdentityProviderResponseBody {
	s.SpecificIdentityProviderConfiguration = v
	return s
}

func (s *SetSpecificIdentityProviderResponseBody) Validate() error {
	if s.SpecificIdentityProviderConfiguration != nil {
		if err := s.SpecificIdentityProviderConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration struct {
	IDPMetadata          *string `json:"IDPMetadata,omitempty" xml:"IDPMetadata,omitempty"`
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	SSOStatus            *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
}

func (s SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) String() string {
	return dara.Prettify(s)
}

func (s SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GoString() string {
	return s.String()
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GetIDPMetadata() *string {
	return s.IDPMetadata
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GetIdentityProviderType() *string {
	return s.IdentityProviderType
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) GetSSOStatus() *string {
	return s.SSOStatus
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) SetIDPMetadata(v string) *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	s.IDPMetadata = &v
	return s
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) SetIdentityProviderType(v string) *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	s.IdentityProviderType = &v
	return s
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) SetSSOStatus(v string) *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration {
	s.SSOStatus = &v
	return s
}

func (s *SetSpecificIdentityProviderResponseBodySpecificIdentityProviderConfiguration) Validate() error {
	return dara.Validate(s)
}
