// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDelegatedServicesForAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *ListDelegatedServicesForAccountRequest
	GetAccountId() *string
	SetMaxResults(v int32) *ListDelegatedServicesForAccountRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDelegatedServicesForAccountRequest
	GetNextToken() *string
}

type ListDelegatedServicesForAccountRequest struct {
	// The Alibaba Cloud account ID of the member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 138660628348****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDelegatedServicesForAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDelegatedServicesForAccountRequest) GoString() string {
	return s.String()
}

func (s *ListDelegatedServicesForAccountRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDelegatedServicesForAccountRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDelegatedServicesForAccountRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDelegatedServicesForAccountRequest) SetAccountId(v string) *ListDelegatedServicesForAccountRequest {
	s.AccountId = &v
	return s
}

func (s *ListDelegatedServicesForAccountRequest) SetMaxResults(v int32) *ListDelegatedServicesForAccountRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDelegatedServicesForAccountRequest) SetNextToken(v string) *ListDelegatedServicesForAccountRequest {
	s.NextToken = &v
	return s
}

func (s *ListDelegatedServicesForAccountRequest) Validate() error {
	return dara.Validate(s)
}
