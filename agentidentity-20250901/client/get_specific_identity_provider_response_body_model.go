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
	RequestId                             *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	IDPMetadata          *string `json:"IDPMetadata,omitempty" xml:"IDPMetadata,omitempty"`
	IdentityProviderType *string `json:"IdentityProviderType,omitempty" xml:"IdentityProviderType,omitempty"`
	SSOStatus            *string `json:"SSOStatus,omitempty" xml:"SSOStatus,omitempty"`
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
