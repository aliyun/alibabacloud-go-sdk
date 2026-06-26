// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAPIKeyCredentialProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAPIKeyCredentialProviders(v []*ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) *ListAPIKeyCredentialProvidersResponseBody
	GetAPIKeyCredentialProviders() []*ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders
	SetMaxResults(v int32) *ListAPIKeyCredentialProvidersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAPIKeyCredentialProvidersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAPIKeyCredentialProvidersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAPIKeyCredentialProvidersResponseBody
	GetTotalCount() *int32
}

type ListAPIKeyCredentialProvidersResponseBody struct {
	APIKeyCredentialProviders []*ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders `json:"APIKeyCredentialProviders,omitempty" xml:"APIKeyCredentialProviders,omitempty" type:"Repeated"`
	MaxResults                *int32                                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                 *string                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId                 *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount                *int32                                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAPIKeyCredentialProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAPIKeyCredentialProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAPIKeyCredentialProvidersResponseBody) GetAPIKeyCredentialProviders() []*ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	return s.APIKeyCredentialProviders
}

func (s *ListAPIKeyCredentialProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAPIKeyCredentialProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAPIKeyCredentialProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAPIKeyCredentialProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAPIKeyCredentialProvidersResponseBody) SetAPIKeyCredentialProviders(v []*ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) *ListAPIKeyCredentialProvidersResponseBody {
	s.APIKeyCredentialProviders = v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBody) SetMaxResults(v int32) *ListAPIKeyCredentialProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBody) SetNextToken(v string) *ListAPIKeyCredentialProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBody) SetRequestId(v string) *ListAPIKeyCredentialProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBody) SetTotalCount(v int32) *ListAPIKeyCredentialProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBody) Validate() error {
	if s.APIKeyCredentialProviders != nil {
		for _, item := range s.APIKeyCredentialProviders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders struct {
	APIKeyCredentialProviderName *string `json:"APIKeyCredentialProviderName,omitempty" xml:"APIKeyCredentialProviderName,omitempty"`
	CreateTime                   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CredentialProviderArn        *string `json:"CredentialProviderArn,omitempty" xml:"CredentialProviderArn,omitempty"`
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TokenVaultName               *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
	UpdateTime                   *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) String() string {
	return dara.Prettify(s)
}

func (s ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GoString() string {
	return s.String()
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GetAPIKeyCredentialProviderName() *string {
	return s.APIKeyCredentialProviderName
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GetCredentialProviderArn() *string {
	return s.CredentialProviderArn
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GetDescription() *string {
	return s.Description
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) SetAPIKeyCredentialProviderName(v string) *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	s.APIKeyCredentialProviderName = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) SetCreateTime(v string) *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	s.CreateTime = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) SetCredentialProviderArn(v string) *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	s.CredentialProviderArn = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) SetDescription(v string) *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	s.Description = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) SetTokenVaultName(v string) *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	s.TokenVaultName = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) SetUpdateTime(v string) *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders {
	s.UpdateTime = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersResponseBodyAPIKeyCredentialProviders) Validate() error {
	return dara.Validate(s)
}
