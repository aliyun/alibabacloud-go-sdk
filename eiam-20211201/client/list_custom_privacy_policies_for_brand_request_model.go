// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPrivacyPoliciesForBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *ListCustomPrivacyPoliciesForBrandRequest
	GetBrandId() *string
	SetInstanceId(v string) *ListCustomPrivacyPoliciesForBrandRequest
	GetInstanceId() *string
	SetMaxResults(v int64) *ListCustomPrivacyPoliciesForBrandRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListCustomPrivacyPoliciesForBrandRequest
	GetNextToken() *string
	SetPreviousToken(v string) *ListCustomPrivacyPoliciesForBrandRequest
	GetPreviousToken() *string
}

type ListCustomPrivacyPoliciesForBrandRequest struct {
	// The brand ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The token that is used to retrieve the previous page of results.
	//
	// example:
	//
	// PTxxxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
}

func (s ListCustomPrivacyPoliciesForBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesForBrandRequest) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) SetBrandId(v string) *ListCustomPrivacyPoliciesForBrandRequest {
	s.BrandId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) SetInstanceId(v string) *ListCustomPrivacyPoliciesForBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) SetMaxResults(v int64) *ListCustomPrivacyPoliciesForBrandRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) SetNextToken(v string) *ListCustomPrivacyPoliciesForBrandRequest {
	s.NextToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) SetPreviousToken(v string) *ListCustomPrivacyPoliciesForBrandRequest {
	s.PreviousToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandRequest) Validate() error {
	return dara.Validate(s)
}
