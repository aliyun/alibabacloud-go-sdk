// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceCountsGroupByResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceCountsSummary(v []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) *GetAggregateResourceCountsGroupByResourceTypeResponseBody
	GetDiscoveredResourceCountsSummary() []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary
	SetRequestId(v string) *GetAggregateResourceCountsGroupByResourceTypeResponseBody
	GetRequestId() *string
}

type GetAggregateResourceCountsGroupByResourceTypeResponseBody struct {
	// The resource type by which the statistics are collected.
	DiscoveredResourceCountsSummary []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 99114B22-1EFF-47DF-B906-1CCE82FF9D60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) GetDiscoveredResourceCountsSummary() []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	return s.DiscoveredResourceCountsSummary
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) SetDiscoveredResourceCountsSummary(v []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) *GetAggregateResourceCountsGroupByResourceTypeResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) SetRequestId(v string) *GetAggregateResourceCountsGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) Validate() error {
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

type GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary struct {
	// This parameter is expired. The resource type by which statistics are collected.
	//
	// example:
	//
	// ACS::RAM::Role
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The total number of resources in the region.
	//
	// example:
	//
	// 7
	ResourceCount *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The resource type by which statistics are collected.
	//
	// example:
	//
	// ACS::RAM::Role
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetGroupName(v string) *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.GroupName = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceType(v string) *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) Validate() error {
	return dara.Validate(s)
}
