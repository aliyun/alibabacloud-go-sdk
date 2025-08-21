// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFingerprintFromOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFingerprint(v string) *RemoveFingerprintFromOIDCProviderRequest
	GetFingerprint() *string
	SetOIDCProviderName(v string) *RemoveFingerprintFromOIDCProviderRequest
	GetOIDCProviderName() *string
}

type RemoveFingerprintFromOIDCProviderRequest struct {
	// The fingerprint that you want to remove.
	//
	// example:
	//
	// 6938fd4d98bab03faadb97b34396831e3780****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderRequest) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *RemoveFingerprintFromOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *RemoveFingerprintFromOIDCProviderRequest) SetFingerprint(v string) *RemoveFingerprintFromOIDCProviderRequest {
	s.Fingerprint = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderRequest) SetOIDCProviderName(v string) *RemoveFingerprintFromOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
