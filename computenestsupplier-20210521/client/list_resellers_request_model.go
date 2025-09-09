// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResellersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListResellersRequestFilter) *ListResellersRequest
	GetFilter() []*ListResellersRequestFilter
	SetMaxResults(v int32) *ListResellersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResellersRequest
	GetNextToken() *string
	SetRegionId(v string) *ListResellersRequest
	GetRegionId() *string
}

type ListResellersRequest struct {
	// The filters.
	Filter []*ListResellersRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	// AAAAAWVmrOoWHbw/80lX0TWxe/s=
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

func (s ListResellersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResellersRequest) GoString() string {
	return s.String()
}

func (s *ListResellersRequest) GetFilter() []*ListResellersRequestFilter {
	return s.Filter
}

func (s *ListResellersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResellersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResellersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResellersRequest) SetFilter(v []*ListResellersRequestFilter) *ListResellersRequest {
	s.Filter = v
	return s
}

func (s *ListResellersRequest) SetMaxResults(v int32) *ListResellersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResellersRequest) SetNextToken(v string) *ListResellersRequest {
	s.NextToken = &v
	return s
}

func (s *ListResellersRequest) SetRegionId(v string) *ListResellersRequest {
	s.RegionId = &v
	return s
}

func (s *ListResellersRequest) Validate() error {
	return dara.Validate(s)
}

type ListResellersRequestFilter struct {
	// The parameter name of the filter. You can specify one or more parameter names to query services. Valid values:
	//
	// 	- ResellerUid: the uid of the distributor.
	//
	// 	- Name: the name of the distributor.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Filter value array.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListResellersRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListResellersRequestFilter) GoString() string {
	return s.String()
}

func (s *ListResellersRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListResellersRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListResellersRequestFilter) SetName(v string) *ListResellersRequestFilter {
	s.Name = &v
	return s
}

func (s *ListResellersRequestFilter) SetValue(v []*string) *ListResellersRequestFilter {
	s.Value = v
	return s
}

func (s *ListResellersRequestFilter) Validate() error {
	return dara.Validate(s)
}
