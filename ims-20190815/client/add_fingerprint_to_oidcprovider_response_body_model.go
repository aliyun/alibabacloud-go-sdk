// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFingerprintToOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) *AddFingerprintToOIDCProviderResponseBody
	GetOIDCProvider() *AddFingerprintToOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *AddFingerprintToOIDCProviderResponseBody
	GetRequestId() *string
}

type AddFingerprintToOIDCProviderResponseBody struct {
	// The name of the OIDC IdP.
	OIDCProvider *AddFingerprintToOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4B809BBC-0E78-544A-A91A-648926412E3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFingerprintToOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFingerprintToOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderResponseBody) GetOIDCProvider() *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *AddFingerprintToOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFingerprintToOIDCProviderResponseBody) SetOIDCProvider(v *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) *AddFingerprintToOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBody) SetRequestId(v string) *AddFingerprintToOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBody) Validate() error {
	if s.OIDCProvider != nil {
		if err := s.OIDCProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFingerprintToOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/OIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 498469743454717****
	ClientIds *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The time when the OIDC IdP was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-11T06:56:03Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the OIDC IdP.
	//
	// example:
	//
	// This is a new OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate. If multiple fingerprints are returned, the fingerprints are separated by commas (,).
	//
	// example:
	//
	// 502ef2deeb3c5b13ea4c3d5193629309e231****,902ef2deeb3c5b13ea4c3d5193629309e231****
	Fingerprints *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	// The timestamp when the OIDC IdP was created.
	//
	// example:
	//
	// 1636613763000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The timestamp when the OIDC IdP was modified.
	//
	// example:
	//
	// 1636706309000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 12
	IssuanceLimitTime *int64 `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	// The URL of the issuer.
	//
	// example:
	//
	// https://dev-xxxxxx.okta.com
	IssuerUrl *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	// The name of the OIDC IdP.
	//
	// example:
	//
	// TestOIDCProvider
	OIDCProviderName *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	// The time when the OIDC IdP was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s AddFingerprintToOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *AddFingerprintToOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *AddFingerprintToOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
