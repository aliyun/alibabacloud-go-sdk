// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveFingerprintFromOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) *RemoveFingerprintFromOIDCProviderResponseBody
	GetOIDCProvider() *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *RemoveFingerprintFromOIDCProviderResponseBody
	GetRequestId() *string
}

type RemoveFingerprintFromOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C90CE971-4C7F-5D93-BD3E-2D0E79D03C01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveFingerprintFromOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) GetOIDCProvider() *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) SetOIDCProvider(v *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) *RemoveFingerprintFromOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) SetRequestId(v string) *RemoveFingerprintFromOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client. If multiple client IDs are returned, the client IDs are separated by commas (,).
	//
	// example:
	//
	// 0oa4u6l8x5WoaVbd****
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
	// 5938fd4d98bab03faadb97b34396831e3780****
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

func (s RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
