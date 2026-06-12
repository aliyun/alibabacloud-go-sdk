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
	// The billing cycle in the YYYY-MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when **Granularity*	- is set to DAILY. The format is YYYY-MM-DD.
	//
	// example:
	//
	// 2024-12-05
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which you want to query bills. Valid values:
	//
	// - MONTHLY: by month. The bill details are consistent with the bills on the By Billing Cycle tab of the Bill Details page in User Center.
	//
	// - DAILY: by day. The bill details are consistent with the bills on the By Day tab of the Bill Details page in User Center.
	//
	// > If you set this parameter to DAILY, you must specify BillingDate.
	//
	// example:
	//
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query. Valid values:
	//
	// - If **NextToken*	- is empty, no more results exist.
	//
	// - If **NextToken*	- has a value, the value is the token that is used to start the next query.
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
	// You can call the [ListServiceInstances](https://help.aliyun.com/document_detail/396200.html) operation to obtain the service instance ID.
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
