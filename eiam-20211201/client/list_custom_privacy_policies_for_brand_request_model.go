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
	// 品牌化Id
	//
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 分页查询时每页行数。默认值为20，最大值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 查询凭证（Token），取值为上一次API调用返回的NextToken参数值。
	//
	// example:
	//
	// NTxxxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 查询上一页凭证（Token），取值为上一次API调用返回的previousToken参数值。
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
