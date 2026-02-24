// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveredResourceCountsGroupByRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveredResourceCountsSummary(v []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) *GetDiscoveredResourceCountsGroupByRegionResponseBody
	GetDiscoveredResourceCountsSummary() []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary
	SetRequestId(v string) *GetDiscoveredResourceCountsGroupByRegionResponseBody
	GetRequestId() *string
}

type GetDiscoveredResourceCountsGroupByRegionResponseBody struct {
	DiscoveredResourceCountsSummary []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	RequestId                       *string                                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) GetDiscoveredResourceCountsSummary() []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	return s.DiscoveredResourceCountsSummary
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) SetDiscoveredResourceCountsSummary(v []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) *GetDiscoveredResourceCountsGroupByRegionResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) SetRequestId(v string) *GetDiscoveredResourceCountsGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) Validate() error {
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

type GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary struct {
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCount *int64  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GetGroupName() *string {
	return s.GroupName
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GetRegion() *string {
	return s.Region
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetGroupName(v string) *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.GroupName = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetRegion(v string) *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) Validate() error {
	return dara.Validate(s)
}
