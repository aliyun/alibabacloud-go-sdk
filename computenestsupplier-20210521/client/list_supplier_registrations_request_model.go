// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupplierRegistrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListSupplierRegistrationsRequestFilter) *ListSupplierRegistrationsRequest
	GetFilter() []*ListSupplierRegistrationsRequestFilter
	SetMaxResults(v int32) *ListSupplierRegistrationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupplierRegistrationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListSupplierRegistrationsRequest
	GetRegionId() *string
}

type ListSupplierRegistrationsRequest struct {
	// The filter.
	Filter []*ListSupplierRegistrationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// Number of items per page in a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAYChudnQUoBH+mGWFpb6oP0=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListSupplierRegistrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupplierRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsRequest) GetFilter() []*ListSupplierRegistrationsRequestFilter {
	return s.Filter
}

func (s *ListSupplierRegistrationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupplierRegistrationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupplierRegistrationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupplierRegistrationsRequest) SetFilter(v []*ListSupplierRegistrationsRequestFilter) *ListSupplierRegistrationsRequest {
	s.Filter = v
	return s
}

func (s *ListSupplierRegistrationsRequest) SetMaxResults(v int32) *ListSupplierRegistrationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupplierRegistrationsRequest) SetNextToken(v string) *ListSupplierRegistrationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupplierRegistrationsRequest) SetRegionId(v string) *ListSupplierRegistrationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSupplierRegistrationsRequest) Validate() error {
	return dara.Validate(s)
}

type ListSupplierRegistrationsRequestFilter struct {
	// Name of the filter field. Allowed values:
	//
	// - SupplierUid: The aliUid of supplier.
	//
	// - SupplierName: The name of supplier.
	//
	// - RegistrationId: Registration ID.
	//
	// - Status: Registration status. Allowed values: Submitted, Approved, Rejected.
	//
	// example:
	//
	// SupplierUid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Filter value.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListSupplierRegistrationsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListSupplierRegistrationsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListSupplierRegistrationsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListSupplierRegistrationsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListSupplierRegistrationsRequestFilter) SetName(v string) *ListSupplierRegistrationsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListSupplierRegistrationsRequestFilter) SetValue(v []*string) *ListSupplierRegistrationsRequestFilter {
	s.Value = v
	return s
}

func (s *ListSupplierRegistrationsRequestFilter) Validate() error {
	return dara.Validate(s)
}
