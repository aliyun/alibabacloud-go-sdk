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
	OIDCProvider *RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	RequestId    *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.OIDCProvider != nil {
		if err := s.OIDCProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveFingerprintFromOIDCProviderResponseBodyOIDCProvider struct {
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
