// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAPIKeyCredentialProvidersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAPIKeyCredentialProvidersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAPIKeyCredentialProvidersRequest
	GetNextToken() *string
	SetTokenVaultName(v string) *ListAPIKeyCredentialProvidersRequest
	GetTokenVaultName() *string
}

type ListAPIKeyCredentialProvidersRequest struct {
	// The maximum number of results to return per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token used to retrieve the next page of results. Leave this parameter empty for the first request. For subsequent requests, specify the NextToken value returned in the previous response.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l8i017XhgKLf/HqWDGD375
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TokenVaultName *string `json:"TokenVaultName,omitempty" xml:"TokenVaultName,omitempty"`
}

func (s ListAPIKeyCredentialProvidersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAPIKeyCredentialProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListAPIKeyCredentialProvidersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAPIKeyCredentialProvidersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAPIKeyCredentialProvidersRequest) GetTokenVaultName() *string {
	return s.TokenVaultName
}

func (s *ListAPIKeyCredentialProvidersRequest) SetMaxResults(v int32) *ListAPIKeyCredentialProvidersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersRequest) SetNextToken(v string) *ListAPIKeyCredentialProvidersRequest {
	s.NextToken = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersRequest) SetTokenVaultName(v string) *ListAPIKeyCredentialProvidersRequest {
	s.TokenVaultName = &v
	return s
}

func (s *ListAPIKeyCredentialProvidersRequest) Validate() error {
	return dara.Validate(s)
}
