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
	OIDCProvider *UpdateOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.OIDCProvider != nil {
		if err := s.OIDCProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateOIDCProviderResponseBodyOIDCProvider struct {
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
