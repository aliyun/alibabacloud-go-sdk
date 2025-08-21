// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveClientIdFromOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) *RemoveClientIdFromOIDCProviderResponseBody
	GetOIDCProvider() *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *RemoveClientIdFromOIDCProviderResponseBody
	GetRequestId() *string
}

type RemoveClientIdFromOIDCProviderResponseBody struct {
	// The information about the OIDC IdP.
	OIDCProvider *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EC9A8F3B-AFA5-5C8F-999D-F97BC7CF1FC5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveClientIdFromOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) GetOIDCProvider() *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) SetOIDCProvider(v *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) *RemoveClientIdFromOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) SetRequestId(v string) *RemoveClientIdFromOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider struct {
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
	// 598469743454717****
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

func (s RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
