// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOIDCProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIds(v string) *CreateOIDCProviderRequest
	GetClientIds() *string
	SetDescription(v string) *CreateOIDCProviderRequest
	GetDescription() *string
	SetFingerprints(v string) *CreateOIDCProviderRequest
	GetFingerprints() *string
	SetIssuanceLimitTime(v int64) *CreateOIDCProviderRequest
	GetIssuanceLimitTime() *int64
	SetIssuerUrl(v string) *CreateOIDCProviderRequest
	GetIssuerUrl() *string
	SetOIDCProviderName(v string) *CreateOIDCProviderRequest
	GetOIDCProviderName() *string
}

type CreateOIDCProviderRequest struct {
	// The ID of the client, which is provided by the external IdP. If you want to specify multiple client IDs, separate the client IDs with commas (,).
	//
	// The client ID can contain letters, digits, and special characters and cannot start with the special characters. The special characters are `periods (.), hyphens (-), underscores (_), colons (:), and forward slashes (/)`.``
	//
	// The client ID can be up to 128 characters in length.
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The description of the OIDC IdP.
	//
	// The description can be up to 256 characters in length.
	//
	// example:
	//
	// This is an OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS CA certificate, which is provided by the external IdP. If you want to specify multiple fingerprints, separate the fingerprints with commas (,).
	//
	// The fingerprint can contain letters and digits.
	//
	// The fingerprint can be up to 128 characters in length.
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer, which is provided by the external IdP. The URL of the issuer must be unique within an Alibaba Cloud account.
	//
	// The URL of the issuer must start with `https` and be in the valid URL format. The URL cannot contain query parameters that follow a question mark (`?`) or logon information that is identified by at signs (`@`). The URL cannot be a fragment URL that contains number signs (`#`).
	//
	// The URL can be up to 255 characters in length.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// The name can contain letters, digits, and special characters and cannot start or end with the special characters. The special characters are `periods, (.), hyphens (-), and underscores (_)`.``
	//
	// The name can be up to 128 characters in length.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
}

func (s CreateOIDCProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOIDCProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderRequest) GetClientIds() *string {
	return s.ClientIds
}

func (s *CreateOIDCProviderRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateOIDCProviderRequest) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *CreateOIDCProviderRequest) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *CreateOIDCProviderRequest) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *CreateOIDCProviderRequest) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *CreateOIDCProviderRequest) SetClientIds(v string) *CreateOIDCProviderRequest {
	s.ClientIds = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetDescription(v string) *CreateOIDCProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetFingerprints(v string) *CreateOIDCProviderRequest {
	s.Fingerprints = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetIssuanceLimitTime(v int64) *CreateOIDCProviderRequest {
	s.IssuanceLimitTime = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetIssuerUrl(v string) *CreateOIDCProviderRequest {
	s.IssuerUrl = &v
	return s
}

func (s *CreateOIDCProviderRequest) SetOIDCProviderName(v string) *CreateOIDCProviderRequest {
	s.OIDCProviderName = &v
	return s
}

func (s *CreateOIDCProviderRequest) Validate() error {
	return dara.Validate(s)
}
