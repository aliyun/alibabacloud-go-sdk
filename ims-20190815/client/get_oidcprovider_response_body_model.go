// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *GetOIDCProviderResponseBodyOIDCProvider) *GetOIDCProviderResponseBody
	GetOIDCProvider() *GetOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *GetOIDCProviderResponseBody
	GetRequestId() *string
}

type GetOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *GetOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E5E1A300-279D-5FBD-A8CF-F4EDC20C4896
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderResponseBody) GetOIDCProvider() *GetOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *GetOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOIDCProviderResponseBody) SetOIDCProvider(v *GetOIDCProviderResponseBodyOIDCProvider) *GetOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *GetOIDCProviderResponseBody) SetRequestId(v string) *GetOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOIDCProviderResponseBody) Validate() error {
	if s.OIDCProvider != nil {
		if err := s.OIDCProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOIDCProviderResponseBodyOIDCProvider struct {
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

func (s GetOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s GetOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *GetOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *GetOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *GetOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
