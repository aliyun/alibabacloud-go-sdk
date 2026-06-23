// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPrivacyPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrivacyPolicyNameStartsWith(v string) *ListCustomPrivacyPoliciesRequest
	GetCustomPrivacyPolicyNameStartsWith() *string
	SetInstanceId(v string) *ListCustomPrivacyPoliciesRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListCustomPrivacyPoliciesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListCustomPrivacyPoliciesRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListCustomPrivacyPoliciesRequest
	GetPreviousToken() *string
}

type ListCustomPrivacyPoliciesRequest struct {
	// The custom term name. Left fuzzy match is supported.
	//
	// example:
	//
	// Custom
	CustomPrivacyPolicyNameStartsWith *string `json:"CustomPrivacyPolicyNameStartsWith,omitempty" xml:"CustomPrivacyPolicyNameStartsWith,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token for the previous page.
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListCustomPrivacyPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesRequest) GetCustomPrivacyPolicyNameStartsWith() *string {
	return s.CustomPrivacyPolicyNameStartsWith
}

func (s *ListCustomPrivacyPoliciesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomPrivacyPoliciesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCustomPrivacyPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomPrivacyPoliciesRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListCustomPrivacyPoliciesRequest) SetCustomPrivacyPolicyNameStartsWith(v string) *ListCustomPrivacyPoliciesRequest {
	s.CustomPrivacyPolicyNameStartsWith = &v
	return s
}

func (s *ListCustomPrivacyPoliciesRequest) SetInstanceId(v string) *ListCustomPrivacyPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesRequest) SetMaxResults(v int64) *ListCustomPrivacyPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomPrivacyPoliciesRequest) SetNextToken(v string) *ListCustomPrivacyPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesRequest) SetPreviousToken(v string) *ListCustomPrivacyPoliciesRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
