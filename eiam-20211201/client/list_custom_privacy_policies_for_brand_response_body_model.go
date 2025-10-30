// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomPrivacyPoliciesForBrandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrandCustomPrivacyPolicies(v []*ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) *ListCustomPrivacyPoliciesForBrandResponseBody
	GetBrandCustomPrivacyPolicies() []*ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies
	SetMaxResults(v int64) *ListCustomPrivacyPoliciesForBrandResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListCustomPrivacyPoliciesForBrandResponseBody
	GetNextToken() *string
	SetPreviousToken(v string) *ListCustomPrivacyPoliciesForBrandResponseBody
	GetPreviousToken() *string
	SetRequestId(v string) *ListCustomPrivacyPoliciesForBrandResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCustomPrivacyPoliciesForBrandResponseBody
	GetTotalCount() *int64
}

type ListCustomPrivacyPoliciesForBrandResponseBody struct {
	BrandCustomPrivacyPolicies []*ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies `json:"BrandCustomPrivacyPolicies,omitempty" xml:"BrandCustomPrivacyPolicies,omitempty" type:"Repeated"`
	// 分页查询时每页行数。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于上一次翻页查询。
	//
	// example:
	//
	// PTxxxexample
	PreviousToken *string `json:"PreviousToken,omitempty" xml:"PreviousToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomPrivacyPoliciesForBrandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesForBrandResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) GetBrandCustomPrivacyPolicies() []*ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies {
	return s.BrandCustomPrivacyPolicies
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) GetPreviousToken() *string {
	return s.PreviousToken
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) SetBrandCustomPrivacyPolicies(v []*ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) *ListCustomPrivacyPoliciesForBrandResponseBody {
	s.BrandCustomPrivacyPolicies = v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) SetMaxResults(v int64) *ListCustomPrivacyPoliciesForBrandResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) SetNextToken(v string) *ListCustomPrivacyPoliciesForBrandResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) SetPreviousToken(v string) *ListCustomPrivacyPoliciesForBrandResponseBody {
	s.PreviousToken = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) SetRequestId(v string) *ListCustomPrivacyPoliciesForBrandResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) SetTotalCount(v int64) *ListCustomPrivacyPoliciesForBrandResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBody) Validate() error {
	if s.BrandCustomPrivacyPolicies != nil {
		for _, item := range s.BrandCustomPrivacyPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies struct {
	// 条款ID
	//
	// example:
	//
	// pp_xxxx
	CustomPrivacyPolicyId *string `json:"CustomPrivacyPolicyId,omitempty" xml:"CustomPrivacyPolicyId,omitempty"`
}

func (s ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) GoString() string {
	return s.String()
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) GetCustomPrivacyPolicyId() *string {
	return s.CustomPrivacyPolicyId
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) SetCustomPrivacyPolicyId(v string) *ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies {
	s.CustomPrivacyPolicyId = &v
	return s
}

func (s *ListCustomPrivacyPoliciesForBrandResponseBodyBrandCustomPrivacyPolicies) Validate() error {
	return dara.Validate(s)
}
