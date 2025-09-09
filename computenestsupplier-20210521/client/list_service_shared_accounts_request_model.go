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
	// The filters.
	Filter []*ListServiceSharedAccountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAR130adlM4fHHVSWpTca/t4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The permissions on the service. Valid values:
	//
	// 	- Deployable: Permissions to deploy the service.
	//
	// 	- Accessible: Permissions to access the service.
	//
	// example:
	//
	// Accessible
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The region ID where the service instance resides.
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
	return dara.Validate(s)
}

type ListServiceSharedAccountsRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- Name: the name of the service.
	//
	// example:
	//
	// UserAliUid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter value N of the filter. Valid values of N: 1 to 10.
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
