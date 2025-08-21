// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *UpdateOIDCProviderResponseBodyOIDCProvider) *UpdateOIDCProviderResponseBody
	GetOIDCProvider() *UpdateOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *UpdateOIDCProviderResponseBody
	GetRequestId() *string
}

type UpdateOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *UpdateOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4C4D1BD-2558-5BD1-8C26-A5D7FB174A55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderResponseBody) GetOIDCProvider() *UpdateOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *UpdateOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOIDCProviderResponseBody) SetOIDCProvider(v *UpdateOIDCProviderResponseBodyOIDCProvider) *UpdateOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *UpdateOIDCProviderResponseBody) SetRequestId(v string) *UpdateOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOIDCProviderResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateOIDCProviderResponseBodyOIDCProvider struct {
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
	// 1636706309000
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
	// 2021-11-12T08:38:29Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s UpdateOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *UpdateOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *UpdateOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
