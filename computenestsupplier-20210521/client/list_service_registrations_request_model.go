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
	// The filter.
	Filter []*ListServiceRegistrationsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set it to the NextToken value returned from the previous API call.
	//
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
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

type ListServiceRegistrationsRequestFilter struct {
	// The name of the filter field. Valid values:
	//
	// - ServiceId: The service ID.
	//
	// - RegistrationId: The review request ID.
	//
	// - Status: The service status. Valid values: Submitted, Approved, Rejected, Canceled, and Executed.
	//
	// example:
	//
	// Canceled
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter values.
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
