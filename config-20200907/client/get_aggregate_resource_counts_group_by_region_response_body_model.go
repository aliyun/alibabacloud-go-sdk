// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceCountsSummary(v []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) *GetAggregateResourceCountsGroupByRegionResponseBody
	GetDiscoveredResourceCountsSummary() []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary
	SetRequestId(v string) *GetAggregateResourceCountsGroupByRegionResponseBody
	GetRequestId() *string
}

type GetAggregateResourceCountsGroupByRegionResponseBody struct {
	// The ID of the region by which statistics are collected.
	DiscoveredResourceCountsSummary []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 399BD94C-D20C-4D27-88D4-89E8D75C0595
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) GetDiscoveredResourceCountsSummary() []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	return s.DiscoveredResourceCountsSummary
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) SetDiscoveredResourceCountsSummary(v []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) *GetAggregateResourceCountsGroupByRegionResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) SetRequestId(v string) *GetAggregateResourceCountsGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) Validate() error {
	if s.DiscoveredResourceCountsSummary != nil {
		for _, item := range s.DiscoveredResourceCountsSummary {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary struct {
	// The dimension by which statistics are collected.
	//
	// >  In most cases, the `Region` parameter is returned instead of the GroupName parameter.
	//
	// example:
	//
	// cn-hangzhou
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the region by which statistics are collected.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The total number of resources in the region.
	//
	// example:
	//
	// 10
	ResourceCount *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GetRegion() *string {
	return s.Region
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetGroupName(v string) *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.GroupName = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetRegion(v string) *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) Validate() error {
	return dara.Validate(s)
}
