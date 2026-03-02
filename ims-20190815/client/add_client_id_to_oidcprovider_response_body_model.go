// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientIdToOIDCProviderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOIDCProvider(v *AddClientIdToOIDCProviderResponseBodyOIDCProvider) *AddClientIdToOIDCProviderResponseBody
	GetOIDCProvider() *AddClientIdToOIDCProviderResponseBodyOIDCProvider
	SetRequestId(v string) *AddClientIdToOIDCProviderResponseBody
	GetRequestId() *string
}

type AddClientIdToOIDCProviderResponseBody struct {
	OIDCProvider *AddClientIdToOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddClientIdToOIDCProviderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddClientIdToOIDCProviderResponseBody) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderResponseBody) GetOIDCProvider() *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	return s.OIDCProvider
}

func (s *AddClientIdToOIDCProviderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddClientIdToOIDCProviderResponseBody) SetOIDCProvider(v *AddClientIdToOIDCProviderResponseBodyOIDCProvider) *AddClientIdToOIDCProviderResponseBody {
	s.OIDCProvider = v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBody) SetRequestId(v string) *AddClientIdToOIDCProviderResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBody) Validate() error {
	if s.OIDCProvider != nil {
		if err := s.OIDCProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddClientIdToOIDCProviderResponseBodyOIDCProvider struct {
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

func (s AddClientIdToOIDCProviderResponseBodyOIDCProvider) String() string {
	return dara.Prettify(s)
}

func (s AddClientIdToOIDCProviderResponseBodyOIDCProvider) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetArn() *string {
	return s.Arn
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetClientIds() *string {
	return s.ClientIds
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetCreateDate() *string {
	return s.CreateDate
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetDescription() *string {
	return s.Description
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetFingerprints() *string {
	return s.Fingerprints
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetIssuanceLimitTime() *int64 {
	return s.IssuanceLimitTime
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetIssuerUrl() *string {
	return s.IssuerUrl
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetOIDCProviderName() *string {
	return s.OIDCProviderName
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetArn(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.Arn = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetClientIds(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.ClientIds = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetCreateDate(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.CreateDate = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetDescription(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.Description = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetFingerprints(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.Fingerprints = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetGmtCreate(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.GmtCreate = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetGmtModified(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.GmtModified = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetIssuanceLimitTime(v int64) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.IssuanceLimitTime = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetIssuerUrl(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.IssuerUrl = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetOIDCProviderName(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.OIDCProviderName = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) SetUpdateDate(v string) *AddClientIdToOIDCProviderResponseBodyOIDCProvider {
	s.UpdateDate = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponseBodyOIDCProvider) Validate() error {
	return dara.Validate(s)
}
