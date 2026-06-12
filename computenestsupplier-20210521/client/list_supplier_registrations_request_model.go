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
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
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
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSupplierRegistrationsRequestFilter struct {
	// The name of the filter. You can query by one or more filter names. Valid values:
	//
	// - SupplierUid: The UID of the service provider.
	//
	// - RegistrationId: The review ID.
	//
	// - Status: The review status.
	//
	// - SupplierName: The name of the service provider.
	//
	// example:
	//
	// SupplierUid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the filter condition.
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
