// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceSharedAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceSharedAccountsRequestFilter) *ListServiceSharedAccountsRequest
	GetFilter() []*ListServiceSharedAccountsRequestFilter
	SetMaxResults(v int32) *ListServiceSharedAccountsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceSharedAccountsRequest
	GetNextToken() *string
	SetPermission(v string) *ListServiceSharedAccountsRequest
	GetPermission() *string
	SetRegionId(v string) *ListServiceSharedAccountsRequest
	GetRegionId() *string
	SetServiceId(v string) *ListServiceSharedAccountsRequest
	GetServiceId() *string
}

type ListServiceSharedAccountsRequest struct {
	// The filter.
	Filter []*ListServiceSharedAccountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query. Set this parameter to the NextToken value returned from the last API call.
	//
	// example:
	//
	// AAAAAR130adlM4fHHVSWpTca/t4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The permission type. Valid values:
	//
	// - Deployable: The service is deployable.
	//
	// - Accessible: The service is accessible.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// example:
	//
	// service-e10349089de34exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListServiceSharedAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSharedAccountsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsRequest) GetFilter() []*ListServiceSharedAccountsRequestFilter {
	return s.Filter
}

func (s *ListServiceSharedAccountsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceSharedAccountsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceSharedAccountsRequest) GetPermission() *string {
	return s.Permission
}

func (s *ListServiceSharedAccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceSharedAccountsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceSharedAccountsRequest) SetFilter(v []*ListServiceSharedAccountsRequestFilter) *ListServiceSharedAccountsRequest {
	s.Filter = v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetMaxResults(v int32) *ListServiceSharedAccountsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetNextToken(v string) *ListServiceSharedAccountsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetPermission(v string) *ListServiceSharedAccountsRequest {
	s.Permission = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetRegionId(v string) *ListServiceSharedAccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) SetServiceId(v string) *ListServiceSharedAccountsRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceSharedAccountsRequest) Validate() error {
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

type ListServiceSharedAccountsRequestFilter struct {
	// The name of the filter.
	//
	// example:
	//
	// UserAliUid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceSharedAccountsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceSharedAccountsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceSharedAccountsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceSharedAccountsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceSharedAccountsRequestFilter) SetName(v string) *ListServiceSharedAccountsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceSharedAccountsRequestFilter) SetValue(v []*string) *ListServiceSharedAccountsRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceSharedAccountsRequestFilter) Validate() error {
	return dara.Validate(s)
}
