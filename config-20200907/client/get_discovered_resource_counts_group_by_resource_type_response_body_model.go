// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceCountsGroupByResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceCountsSummary(v []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody
	GetDiscoveredResourceCountsSummary() []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary
	SetRequestId(v string) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody
	GetRequestId() *string
}

type GetDiscoveredResourceCountsGroupByResourceTypeResponseBody struct {
	// The statistics on the resources.
	DiscoveredResourceCountsSummary []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// AC9BD94C-D20C-4D27-88D4-89E8D75C051B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) GetDiscoveredResourceCountsSummary() []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	return s.DiscoveredResourceCountsSummary
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) SetDiscoveredResourceCountsSummary(v []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) SetRequestId(v string) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) Validate() error {
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

type GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary struct {
	// The resource type by which the statistics are collected.
	//
	// > We recommend that you use the `ResourceType` parameter.
	//
	// example:
	//
	// ACS::ECS::Instance
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The total number of resources.
	//
	// example:
	//
	// 10
	ResourceCount *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The resource type by which the statistics are collected.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetGroupName(v string) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.GroupName = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceType(v string) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceType = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) Validate() error {
	return dara.Validate(s)
}
