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
	SetRegionId(v string) *ListServiceUsagesRequest
	GetRegionId() *string
}

type ListServiceUsagesRequest struct {
	// The filter.
	Filter []*ListServiceUsagesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. This value is the NextToken value returned from the previous API call.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *ListServiceUsagesRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *ListServiceUsagesRequest) SetRegionId(v string) *ListServiceUsagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceUsagesRequest) Validate() error {
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

type ListServiceUsagesRequestFilter struct {
	// The name of the filter. You can query by one or more names. Valid values:
	//
	// - ServiceId: The service ID.
	//
	// - ServiceName: The service name.
	//
	// - Status: The service status.
	//
	// - SupplierName: The name of the supplier.
	//
	// example:
	//
	// ServiceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A list of filter values.
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
