// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilters(v []*GetMultiAccountResourceCountsResponseBodyFilters) *GetMultiAccountResourceCountsResponseBody
	GetFilters() []*GetMultiAccountResourceCountsResponseBodyFilters
	SetGroupByKey(v string) *GetMultiAccountResourceCountsResponseBody
	GetGroupByKey() *string
	SetRequestId(v string) *GetMultiAccountResourceCountsResponseBody
	GetRequestId() *string
	SetResourceCounts(v []*GetMultiAccountResourceCountsResponseBodyResourceCounts) *GetMultiAccountResourceCountsResponseBody
	GetResourceCounts() []*GetMultiAccountResourceCountsResponseBodyResourceCounts
	SetScope(v string) *GetMultiAccountResourceCountsResponseBody
	GetScope() *string
}

type GetMultiAccountResourceCountsResponseBody struct {
	Filters []*GetMultiAccountResourceCountsResponseBodyFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// example:
	//
	// EFA806B9-7F36-55AB-8B7A-D680C2C5EE57
	RequestId      *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceCounts []*GetMultiAccountResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s GetMultiAccountResourceCountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBody) GetFilters() []*GetMultiAccountResourceCountsResponseBodyFilters {
	return s.Filters
}

func (s *GetMultiAccountResourceCountsResponseBody) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetMultiAccountResourceCountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiAccountResourceCountsResponseBody) GetResourceCounts() []*GetMultiAccountResourceCountsResponseBodyResourceCounts {
	return s.ResourceCounts
}

func (s *GetMultiAccountResourceCountsResponseBody) GetScope() *string {
	return s.Scope
}

func (s *GetMultiAccountResourceCountsResponseBody) SetFilters(v []*GetMultiAccountResourceCountsResponseBodyFilters) *GetMultiAccountResourceCountsResponseBody {
	s.Filters = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetGroupByKey(v string) *GetMultiAccountResourceCountsResponseBody {
	s.GroupByKey = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetRequestId(v string) *GetMultiAccountResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetResourceCounts(v []*GetMultiAccountResourceCountsResponseBodyResourceCounts) *GetMultiAccountResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) SetScope(v string) *GetMultiAccountResourceCountsResponseBody {
	s.Scope = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountResourceCountsResponseBodyFilters struct {
	// example:
	//
	// RegionId
	Key    *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetMultiAccountResourceCountsResponseBodyFilters) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBodyFilters) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) GetKey() *string {
	return s.Key
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) GetValues() []*string {
	return s.Values
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) SetKey(v string) *GetMultiAccountResourceCountsResponseBodyFilters {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) SetValues(v []*string) *GetMultiAccountResourceCountsResponseBodyFilters {
	s.Values = v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyFilters) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountResourceCountsResponseBodyResourceCounts struct {
	// example:
	//
	// 2
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// ACS::ECS::NetworkInterface
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetMultiAccountResourceCountsResponseBodyResourceCounts) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) GetCount() *int64 {
	return s.Count
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) GetGroupName() *string {
	return s.GroupName
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetMultiAccountResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) SetGroupName(v string) *GetMultiAccountResourceCountsResponseBodyResourceCounts {
	s.GroupName = &v
	return s
}

func (s *GetMultiAccountResourceCountsResponseBodyResourceCounts) Validate() error {
	return dara.Validate(s)
}
