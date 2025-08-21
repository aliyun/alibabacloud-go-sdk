// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOIDCProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListOIDCProvidersResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListOIDCProvidersResponseBody
	GetMarker() *string
	SetOIDCProviders(v *ListOIDCProvidersResponseBodyOIDCProviders) *ListOIDCProvidersResponseBody
	GetOIDCProviders() *ListOIDCProvidersResponseBodyOIDCProviders
	SetRequestId(v string) *ListOIDCProvidersResponseBody
	GetRequestId() *string
}

type ListOIDCProvidersResponseBody struct {
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The `marker`. This parameter is returned only if the value of `IsTruncated` is `true`. If the parameter is returned, you can call this operation again and set this parameter to obtain the truncated part.``
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The information about the OIDC IdP.
	OIDCProviders *ListOIDCProvidersResponseBodyOIDCProviders `json:"OIDCProviders,omitempty" xml:"OIDCProviders,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D2148337-B86A-57F0-8B31-EB7BE0125226
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOIDCProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOIDCProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListOIDCProvidersResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListOIDCProvidersResponseBody) GetOIDCProviders() *ListOIDCProvidersResponseBodyOIDCProviders {
	return s.OIDCProviders
}

func (s *ListOIDCProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOIDCProvidersResponseBody) SetIsTruncated(v bool) *ListOIDCProvidersResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListOIDCProvidersResponseBody) SetMarker(v string) *ListOIDCProvidersResponseBody {
	s.Marker = &v
	return s
}

func (s *ListOIDCProvidersResponseBody) SetOIDCProviders(v *ListOIDCProvidersResponseBodyOIDCProviders) *ListOIDCProvidersResponseBody {
	s.OIDCProviders = v
	return s
}

func (s *ListOIDCProvidersResponseBody) SetRequestId(v string) *ListOIDCProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOIDCProvidersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOIDCProvidersResponseBodyOIDCProviders struct {
	OIDCProvider []*ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Repeated"`
}

func (s ListOIDCProvidersResponseBodyOIDCProviders) String() string {
	return dara.Prettify(s)
}

func (s ListOIDCProvidersResponseBodyOIDCProviders) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponseBodyOIDCProviders) GetOIDCProvider() []*ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	return s.OIDCProvider
}

func (s *ListOIDCProvidersResponseBodyOIDCProviders) SetOIDCProvider(v []*ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) *ListOIDCProvidersResponseBodyOIDCProviders {
	s.OIDCProvider = v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProviders) Validate() error {
	return dara.Validate(s)
}

type ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider struct {
	// The Alibaba Cloud Resource Name (ARN) of the OIDC IdP.
	//
	// example:
	//
	// acs:ram::177242285274****:oidc-provider/TestOIDCProvider
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the client, If you want to specify multiple client IDs, separate the client IDs with commas (,).
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

func (s ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GoString() string {
	return s.String()
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetArn(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.Arn = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetClientIds(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetCreateDate(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetDescription(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.Description = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetFingerprints(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetGmtCreate(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetGmtModified(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetIssuanceLimitTime(v int64) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetIssuerUrl(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetOIDCProviderName(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) SetUpdateDate(v string) *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider) Validate() error {
	return dara.Validate(s)
}
