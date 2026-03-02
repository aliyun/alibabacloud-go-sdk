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
	OIDCProvider *GetOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
