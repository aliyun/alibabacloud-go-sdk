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
	Marker        *string                                     `json:"Marker,omitempty" xml:"Marker,omitempty"`
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
	if s.OIDCProviders != nil {
		if err := s.OIDCProviders.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.OIDCProvider != nil {
		for _, item := range s.OIDCProvider {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOIDCProvidersResponseBodyOIDCProvidersOIDCProvider struct {
	Arn               *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	ClientIds         *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	CreateDate        *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Fingerprints      *string `json:"Fingerprints,omitempty" xml:"Fingerprints,omitempty"`
	GmtCreate         *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified       *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IssuanceLimitTime *int64  `json:"IssuanceLimitTime,omitempty" xml:"IssuanceLimitTime,omitempty"`
	IssuerUrl         *string `json:"IssuerUrl,omitempty" xml:"IssuerUrl,omitempty"`
	OIDCProviderName  *string `json:"OIDCProviderName,omitempty" xml:"OIDCProviderName,omitempty"`
	UpdateDate        *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
