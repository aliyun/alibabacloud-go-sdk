// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *CreateOIDCProviderResponseBodyOIDCProvider) *CreateOIDCProviderResponseBody
	GetOIDCProvider() *CreateOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *CreateOIDCProviderResponseBody
	GetRequestId() *string
}

type CreateOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *CreateOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 64B11B41-636D-51E3-A39B-C8703CD2218C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderResponseBody) GetOIDCProvider() *CreateOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *CreateOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOIDCProviderResponseBody) SetOIDCProvider(v *CreateOIDCProviderResponseBodyOIDCProvider) *CreateOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *CreateOIDCProviderResponseBody) SetRequestId(v string) *CreateOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOIDCProviderResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateOIDCProviderResponseBodyOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client.
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
	// This is an OIDC Provider.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The fingerprint of the HTTPS certificate.
	//
	// example:
	//
	// 902ef2deeb3c5b13ea4c3d5193629309e231****
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
	// 1636613763000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The earliest time when an external IdP can issue an ID token. If the value of the iat field in the ID token is later than the current time, the request is rejected. Unit: hours. Valid values: 1 to 168.
	//
	// example:
	//
	// 6
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
	// 2021-11-11T06:56:03Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s CreateOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *CreateOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *CreateOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
