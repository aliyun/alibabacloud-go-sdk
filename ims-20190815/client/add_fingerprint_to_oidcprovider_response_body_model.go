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
	OIDCProvider *AddFingerprintToOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	RequestId    *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
