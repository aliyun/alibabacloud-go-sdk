// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceUsagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListServiceUsagesRequestFilter) *ListServiceUsagesRequest
	GetFilter() []*ListServiceUsagesRequestFilter
	SetMaxResults(v int32) *ListServiceUsagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceUsagesRequest
	GetNextToken() *string
}

type ListServiceUsagesRequest struct {
	// The filters.
	Filter []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListServiceUsagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceUsagesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequest) GetFilter() []*ListServiceUsagesRequestFilter {
	return s.Filter
}

func (s *ListServiceUsagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceUsagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceUsagesRequest) SetFilter(v []*ListServiceUsagesRequestFilter) *ListServiceUsagesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceUsagesRequest) SetMaxResults(v int32) *ListServiceUsagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceUsagesRequest) SetNextToken(v string) *ListServiceUsagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceUsagesRequest) Validate() error {
	return dara.Validate(s)
}

type ListServiceUsagesRequestFilter struct {
	// The parameter name of the filter. You can specify one or more filters. Valid values:
	//
	// 	- ServiceId: the ID of the service.
	//
	// 	- ServiceName: the service name.
	//
	// 	- Status: the state of the service.
	//
	// 	- SupplierName: the name of the service provider.
	//
	// example:
	//
	// ServiceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceUsagesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListServiceUsagesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceUsagesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListServiceUsagesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListServiceUsagesRequestFilter) SetName(v string) *ListServiceUsagesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceUsagesRequestFilter) SetValue(v []*string) *ListServiceUsagesRequestFilter {
	s.Value = v
	return s
}

func (s *ListServiceUsagesRequestFilter) Validate() error {
	return dara.Validate(s)
}
