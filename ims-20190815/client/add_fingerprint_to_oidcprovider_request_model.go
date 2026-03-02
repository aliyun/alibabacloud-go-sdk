// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFingerprintToOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFingerprint(v string) *AddFingerprintToOIDCProviderRequest
	GetFingerprint() *string
	SetOIDCProviderName(v string) *AddFingerprintToOIDCProviderRequest
	GetOIDCProviderName() *string
}

type AddFingerprintToOIDCProviderRequest struct {
	Fingerprint      *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s AddFingerprintToOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFingerprintToOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderRequest) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *AddFingerprintToOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *AddFingerprintToOIDCProviderRequest) SetFingerprint(v string) *AddFingerprintToOIDCProviderRequest {
	s.Fingerprint = &v
	return s
}

func (s *AddFingerprintToOIDCProviderRequest) SetOIDCProviderName(v string) *AddFingerprintToOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *AddFingerprintToOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
