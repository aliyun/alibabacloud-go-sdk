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
	OIDCProvider *RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider `json:"OIDCProvider,omitempty" xml:"OIDCProvider,omitempty" type:"Struct"`
	RequestId    *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.OIDCProvider != nil {
		if err := s.OIDCProvider.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RemoveClientIdFromOIDCProviderResponseBodyOIDCProvider struct {
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
