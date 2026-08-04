// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpIdaas2SubConfig interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *IdpIdaas2SubConfig
	GetApplicationId() *string
	SetClientId(v string) *IdpIdaas2SubConfig
	GetClientId() *string
	SetClientSecret(v string) *IdpIdaas2SubConfig
	GetClientSecret() *string
	SetEventAesKey(v string) *IdpIdaas2SubConfig
	GetEventAesKey() *string
	SetEventLabel(v string) *IdpIdaas2SubConfig
	GetEventLabel() *string
	SetInstanceId(v string) *IdpIdaas2SubConfig
	GetInstanceId() *string
	SetPublicKeyEndpoint(v string) *IdpIdaas2SubConfig
	GetPublicKeyEndpoint() *string
	SetRegion(v string) *IdpIdaas2SubConfig
	GetRegion() *string
	SetSamlMetadata(v string) *IdpIdaas2SubConfig
	GetSamlMetadata() *string
}

type IdpIdaas2SubConfig struct {
	// The unique identifier of the application within the IDaaS instance.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client ID of the application registered with the identity provider.
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The client secret used to authenticate the application with the identity provider.
	ClientSecret *string `json:"ClientSecret,omitempty" xml:"ClientSecret,omitempty"`
	// The AES encryption key for securing event data.
	EventAesKey *string `json:"EventAesKey,omitempty" xml:"EventAesKey,omitempty"`
	// A label that identifies the event subscription.
	EventLabel *string `json:"EventLabel,omitempty" xml:"EventLabel,omitempty"`
	// The unique identifier of the IDaaS instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The URL of the endpoint providing the public key for token signature verification.
	PublicKeyEndpoint *string `json:"PublicKeyEndpoint,omitempty" xml:"PublicKeyEndpoint,omitempty"`
	// The deployment region of the IDaaS instance.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The SAML metadata in XML format. It specifies the identity provider\\"s configuration, including endpoints and certificates.
	SamlMetadata *string `json:"SamlMetadata,omitempty" xml:"SamlMetadata,omitempty"`
}

func (s IdpIdaas2SubConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpIdaas2SubConfig) GoString() string {
	return s.String()
}

func (s *IdpIdaas2SubConfig) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *IdpIdaas2SubConfig) GetClientId() *string {
	return s.ClientId
}

func (s *IdpIdaas2SubConfig) GetClientSecret() *string {
	return s.ClientSecret
}

func (s *IdpIdaas2SubConfig) GetEventAesKey() *string {
	return s.EventAesKey
}

func (s *IdpIdaas2SubConfig) GetEventLabel() *string {
	return s.EventLabel
}

func (s *IdpIdaas2SubConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *IdpIdaas2SubConfig) GetPublicKeyEndpoint() *string {
	return s.PublicKeyEndpoint
}

func (s *IdpIdaas2SubConfig) GetRegion() *string {
	return s.Region
}

func (s *IdpIdaas2SubConfig) GetSamlMetadata() *string {
	return s.SamlMetadata
}

func (s *IdpIdaas2SubConfig) SetApplicationId(v string) *IdpIdaas2SubConfig {
	s.ApplicationId = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetClientId(v string) *IdpIdaas2SubConfig {
	s.ClientId = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetClientSecret(v string) *IdpIdaas2SubConfig {
	s.ClientSecret = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetEventAesKey(v string) *IdpIdaas2SubConfig {
	s.EventAesKey = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetEventLabel(v string) *IdpIdaas2SubConfig {
	s.EventLabel = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetInstanceId(v string) *IdpIdaas2SubConfig {
	s.InstanceId = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetPublicKeyEndpoint(v string) *IdpIdaas2SubConfig {
	s.PublicKeyEndpoint = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetRegion(v string) *IdpIdaas2SubConfig {
	s.Region = &v
	return s
}

func (s *IdpIdaas2SubConfig) SetSamlMetadata(v string) *IdpIdaas2SubConfig {
	s.SamlMetadata = &v
	return s
}

func (s *IdpIdaas2SubConfig) Validate() error {
	return dara.Validate(s)
}
