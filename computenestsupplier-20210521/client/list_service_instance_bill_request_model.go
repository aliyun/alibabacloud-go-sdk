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
	SetServiceId(v string) *ListServiceInstanceBillRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *ListServiceInstanceBillRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *ListServiceInstanceBillRequest
	GetServiceVersion() *string
}

type ListServiceInstanceBillRequest struct {
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the **Granularity*	- parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2024-12-05
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
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-6121296da4f44e469519
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-0d0d7bc9accc4e2e8a8f
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
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

func (s *ListServiceInstanceBillRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceInstanceBillRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *ListServiceInstanceBillRequest) GetServiceVersion() *string {
	return s.ServiceVersion
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

func (s *ListServiceInstanceBillRequest) SetServiceId(v string) *ListServiceInstanceBillRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceInstanceId(v string) *ListServiceInstanceBillRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceBillRequest) SetServiceVersion(v string) *ListServiceInstanceBillRequest {
	s.ServiceVersion = &v
	return s
}

func (s *ListServiceInstanceBillRequest) Validate() error {
	return dara.Validate(s)
}
