// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceRegistrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceRegistrationsRequestFilter) *ListServiceRegistrationsRequest
	GetFilter() []*ListServiceRegistrationsRequestFilter
	SetMaxResults(v int32) *ListServiceRegistrationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceRegistrationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceRegistrationsRequest
	GetRegionId() *string
}

type ListServiceRegistrationsRequest struct {
	// Filter.
	Filter []*ListServiceRegistrationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of items to return per page during a paginated query. The maximum is 100, and the default is 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListServiceRegistrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsRequest) GetFilter() []*ListServiceRegistrationsRequestFilter {
	return s.Filter
}

func (s *ListServiceRegistrationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceRegistrationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceRegistrationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceRegistrationsRequest) SetFilter(v []*ListServiceRegistrationsRequestFilter) *ListServiceRegistrationsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceRegistrationsRequest) SetMaxResults(v int32) *ListServiceRegistrationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetNextToken(v string) *ListServiceRegistrationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceRegistrationsRequest) SetRegionId(v string) *ListServiceRegistrationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceRegistrationsRequest) Validate() error {
	return dara.Validate(s)
}

type ListServiceRegistrationsRequestFilter struct {
	// Name of the filter field. Allowed values:
	//
	// - ServiceId: Service ID.
	//
	// - RegistrationId: Registration ID.
	//
	// - Status: Registration status. Allowed values: Submitted, Approved, Rejected, Canceled, and Executed.
	//
	// example:
	//
	// Canceled
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceRegistrationsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceRegistrationsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceRegistrationsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceRegistrationsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceRegistrationsRequestFilter) SetName(v string) *ListServiceRegistrationsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceRegistrationsRequestFilter) SetValue(v []*string) *ListServiceRegistrationsRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceRegistrationsRequestFilter) Validate() error {
	return dara.Validate(s)
}
