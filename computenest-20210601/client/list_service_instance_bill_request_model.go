// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingCycle(v string) *ListServiceInstanceBillRequest
	GetBillingCycle() *string
	SetBillingDate(v string) *ListServiceInstanceBillRequest
	GetBillingDate() *string
	SetGranularity(v string) *ListServiceInstanceBillRequest
	GetGranularity() *string
	SetMaxResults(v int32) *ListServiceInstanceBillRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceInstanceBillRequest
	GetNextToken() *string
	SetServiceInstanceId(v string) *ListServiceInstanceBillRequest
	GetServiceInstanceId() *string
}

type ListServiceInstanceBillRequest struct {
	// The billing cycle in YYYY-MM format. Only billing cycles in the last 18 months are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when **Granularity*	- is set to DAILY. The format is YYYY-MM-DD.
	//
	// example:
	//
	// 2025-04-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity of the bills to query. Valid values:
	//
	// - MONTHLY: monthly. The data is consistent with the bills aggregated by billing cycle in the User Center.
	//
	// - DAILY: daily. The data is consistent with the bills aggregated by day in the User Center.
	//
	// If you set this parameter to DAILY, you must also specify **BillingDate**.
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-49793f3bfa034ec6a990
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceBillRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *ListServiceInstanceBillRequest) GetBillingDate() *string {
	return s.BillingDate
}

func (s *ListServiceInstanceBillRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *ListServiceInstanceBillRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceInstanceBillRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceInstanceBillRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstanceBillRequest) SetBillingCycle(v string) *ListServiceInstanceBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetBillingDate(v string) *ListServiceInstanceBillRequest {
	s.BillingDate = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetGranularity(v string) *ListServiceInstanceBillRequest {
	s.Granularity = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetMaxResults(v int32) *ListServiceInstanceBillRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetNextToken(v string) *ListServiceInstanceBillRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceInstanceId(v string) *ListServiceInstanceBillRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceBillRequest) Validate() error {
	return dara.Validate(s)
}
