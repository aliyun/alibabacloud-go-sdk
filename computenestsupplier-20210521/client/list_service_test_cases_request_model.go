// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceTestCasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*ListServiceTestCasesRequestFilters) *ListServiceTestCasesRequest
	GetFilters() []*ListServiceTestCasesRequestFilters
	SetMaxResults(v int32) *ListServiceTestCasesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListServiceTestCasesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListServiceTestCasesRequest
	GetRegionId() *string
	SetServiceId(v string) *ListServiceTestCasesRequest
	GetServiceId() *string
	SetServiceVersion(v string) *ListServiceTestCasesRequest
	GetServiceVersion() *string
}

type ListServiceTestCasesRequest struct {
	// The filters.
	Filters []*ListServiceTestCasesRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The number of entries to return on each page. The maximum value is 100. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
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
	// The service ID.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ListServiceTestCasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestCasesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesRequest) GetFilters() []*ListServiceTestCasesRequestFilters {
	return s.Filters
}

func (s *ListServiceTestCasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceTestCasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceTestCasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListServiceTestCasesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListServiceTestCasesRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *ListServiceTestCasesRequest) SetFilters(v []*ListServiceTestCasesRequestFilters) *ListServiceTestCasesRequest {
	s.Filters = v
	return s
}

func (s *ListServiceTestCasesRequest) SetMaxResults(v int32) *ListServiceTestCasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetNextToken(v string) *ListServiceTestCasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetRegionId(v string) *ListServiceTestCasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetServiceId(v string) *ListServiceTestCasesRequest {
	s.ServiceId = &v
	return s
}

func (s *ListServiceTestCasesRequest) SetServiceVersion(v string) *ListServiceTestCasesRequest {
	s.ServiceVersion = &v
	return s
}

func (s *ListServiceTestCasesRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceTestCasesRequestFilters struct {
	// The name of the filter. You can specify one or more filter names. Valid values:
	//
	// **Status**: Filters by status. **TaskId**: Filters by task ID.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceTestCasesRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s ListServiceTestCasesRequestFilters) GoString() string {
	return s.String()
}

func (s *ListServiceTestCasesRequestFilters) GetName() *string {
	return s.Name
}

func (s *ListServiceTestCasesRequestFilters) GetValue() []*string {
	return s.Value
}

func (s *ListServiceTestCasesRequestFilters) SetName(v string) *ListServiceTestCasesRequestFilters {
	s.Name = &v
	return s
}

func (s *ListServiceTestCasesRequestFilters) SetValue(v []*string) *ListServiceTestCasesRequestFilters {
	s.Value = v
	return s
}

func (s *ListServiceTestCasesRequestFilters) Validate() error {
	return dara.Validate(s)
}
