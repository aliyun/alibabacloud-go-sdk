// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFiltersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultFilterName(v string) *ListFiltersResponseBody
	GetDefaultFilterName() *string
	SetFilters(v []*ListFiltersResponseBodyFilters) *ListFiltersResponseBody
	GetFilters() []*ListFiltersResponseBodyFilters
	SetRequestId(v string) *ListFiltersResponseBody
	GetRequestId() *string
}

type ListFiltersResponseBody struct {
	// The name of the default filter.
	//
	// example:
	//
	// My Filters
	DefaultFilterName *string `json:"DefaultFilterName,omitempty" xml:"DefaultFilterName,omitempty"`
	// The configurations of the filter.
	Filters []*ListFiltersResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AA39FB9C-CB74-5E73-8DFE-3A2B096F0759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFiltersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFiltersResponseBody) GoString() string {
	return s.String()
}

func (s *ListFiltersResponseBody) GetDefaultFilterName() *string {
	return s.DefaultFilterName
}

func (s *ListFiltersResponseBody) GetFilters() []*ListFiltersResponseBodyFilters {
	return s.Filters
}

func (s *ListFiltersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFiltersResponseBody) SetDefaultFilterName(v string) *ListFiltersResponseBody {
	s.DefaultFilterName = &v
	return s
}

func (s *ListFiltersResponseBody) SetFilters(v []*ListFiltersResponseBodyFilters) *ListFiltersResponseBody {
	s.Filters = v
	return s
}

func (s *ListFiltersResponseBody) SetRequestId(v string) *ListFiltersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFiltersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFiltersResponseBodyFilters struct {
	// The configurations of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "regions": [],
	//
	//   "tagFilters": [
	//
	//     [{ "type": "notContainTagKey", "tagKey": "xxx", "tagValue": "" }],
	//
	//     [{ "tagKey": "xxx", "tagValue": "xxx" }]
	//
	//   ],
	//
	//   "resourceTypes": [
	//
	//     "ACS::ECS::AutoSnapshotPolicy"
	//
	//   ]
	//
	// }
	FilterConfiguration *string `json:"FilterConfiguration,omitempty" xml:"FilterConfiguration,omitempty"`
	// The name of the filter.
	//
	// example:
	//
	// My devices
	FilterName *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
}

func (s ListFiltersResponseBodyFilters) String() string {
	return dara.Prettify(s)
}

func (s ListFiltersResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *ListFiltersResponseBodyFilters) GetFilterConfiguration() *string {
	return s.FilterConfiguration
}

func (s *ListFiltersResponseBodyFilters) GetFilterName() *string {
	return s.FilterName
}

func (s *ListFiltersResponseBodyFilters) SetFilterConfiguration(v string) *ListFiltersResponseBodyFilters {
	s.FilterConfiguration = &v
	return s
}

func (s *ListFiltersResponseBodyFilters) SetFilterName(v string) *ListFiltersResponseBodyFilters {
	s.FilterName = &v
	return s
}

func (s *ListFiltersResponseBodyFilters) Validate() error {
	return dara.Validate(s)
}
