// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*GetResourceCountsResponseBodyFilters) *GetResourceCountsResponseBody
	GetFilters() []*GetResourceCountsResponseBodyFilters
	SetGroupByKey(v string) *GetResourceCountsResponseBody
	GetGroupByKey() *string
	SetRequestId(v string) *GetResourceCountsResponseBody
	GetRequestId() *string
	SetResourceCounts(v []*GetResourceCountsResponseBodyResourceCounts) *GetResourceCountsResponseBody
	GetResourceCounts() []*GetResourceCountsResponseBodyResourceCounts
}

type GetResourceCountsResponseBody struct {
	// The filter conditions.
	Filters []*GetResourceCountsResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The dimension by which resources are queried.
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D98D9B0-318D-56A4-910C-93B5F945AF2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The numbers of resources.
	ResourceCounts []*GetResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
}

func (s GetResourceCountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBody) GetFilters() []*GetResourceCountsResponseBodyFilters {
	return s.Filters
}

func (s *GetResourceCountsResponseBody) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetResourceCountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceCountsResponseBody) GetResourceCounts() []*GetResourceCountsResponseBodyResourceCounts {
	return s.ResourceCounts
}

func (s *GetResourceCountsResponseBody) SetFilters(v []*GetResourceCountsResponseBodyFilters) *GetResourceCountsResponseBody {
	s.Filters = v
	return s
}

func (s *GetResourceCountsResponseBody) SetGroupByKey(v string) *GetResourceCountsResponseBody {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceCountsResponseBody) SetRequestId(v string) *GetResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceCountsResponseBody) SetResourceCounts(v []*GetResourceCountsResponseBodyResourceCounts) *GetResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

func (s *GetResourceCountsResponseBody) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceCounts != nil {
		for _, item := range s.ResourceCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceCountsResponseBodyFilters struct {
	// The key of the filter condition.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the filter condition.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetResourceCountsResponseBodyFilters) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBodyFilters) GetKey() *string {
	return s.Key
}

func (s *GetResourceCountsResponseBodyFilters) GetValues() []*string {
	return s.Values
}

func (s *GetResourceCountsResponseBodyFilters) SetKey(v string) *GetResourceCountsResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *GetResourceCountsResponseBodyFilters) SetValues(v []*string) *GetResourceCountsResponseBodyFilters {
	s.Values = v
	return s
}

func (s *GetResourceCountsResponseBodyFilters) Validate() error {
	return dara.Validate(s)
}

type GetResourceCountsResponseBodyResourceCounts struct {
	// The number of resources.
	//
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The group name.
	//
	// example:
	//
	// ACS::ECS::NetworkInterface
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetResourceCountsResponseBodyResourceCounts) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetResourceCountsResponseBodyResourceCounts) GetCount() *int64 {
	return s.Count
}

func (s *GetResourceCountsResponseBodyResourceCounts) GetGroupName() *string {
	return s.GroupName
}

func (s *GetResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetResourceCountsResponseBodyResourceCounts) SetGroupName(v string) *GetResourceCountsResponseBodyResourceCounts {
	s.GroupName = &v
	return s
}

func (s *GetResourceCountsResponseBodyResourceCounts) Validate() error {
	return dara.Validate(s)
}
