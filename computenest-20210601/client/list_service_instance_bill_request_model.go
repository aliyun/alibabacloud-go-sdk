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
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-05
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2025-04-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills by month. The data queried is consistent with the data that is displayed for the specified billing cycle on the Billing Details tab of the Bill Details page in User Center.
	//
	// 	- DAILY: queries bills by day. The data queried is consistent with the data that is displayed for the specified day on the Billing Details tab of the Bill Details page in User Center.
	//
	// You must set the **BillingDate*	- parameter before you can set the Granularity parameter to DAILY.
	//
	// example:
	//
	// DAILY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of entries page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service instance ID.
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
