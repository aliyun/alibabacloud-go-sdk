// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABMetricGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricGroups(v []*ListABMetricGroupsResponseBodyABMetricGroups) *ListABMetricGroupsResponseBody
	GetABMetricGroups() []*ListABMetricGroupsResponseBodyABMetricGroups
	SetRequestId(v string) *ListABMetricGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListABMetricGroupsResponseBody
	GetTotalCount() *int64
}

type ListABMetricGroupsResponseBody struct {
	// The A/B metric groups.
	ABMetricGroups []*ListABMetricGroupsResponseBodyABMetricGroups `json:"ABMetricGroups,omitempty" xml:"ABMetricGroups,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E15A1443-7917-5BE0-AE70-25538ECF398D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of A/B metric groups.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListABMetricGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsResponseBody) GetABMetricGroups() []*ListABMetricGroupsResponseBodyABMetricGroups {
	return s.ABMetricGroups
}

func (s *ListABMetricGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListABMetricGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListABMetricGroupsResponseBody) SetABMetricGroups(v []*ListABMetricGroupsResponseBodyABMetricGroups) *ListABMetricGroupsResponseBody {
	s.ABMetricGroups = v
	return s
}

func (s *ListABMetricGroupsResponseBody) SetRequestId(v string) *ListABMetricGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABMetricGroupsResponseBody) SetTotalCount(v int64) *ListABMetricGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListABMetricGroupsResponseBody) Validate() error {
	if s.ABMetricGroups != nil {
		for _, item := range s.ABMetricGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListABMetricGroupsResponseBodyABMetricGroups struct {
	// The A/B metric group ID.
	//
	// example:
	//
	// 1
	ABMetricGroupId *string `json:"ABMetricGroupId,omitempty" xml:"ABMetricGroupId,omitempty"`
	// The comma-separated IDs of the A/B metrics in the group.
	//
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// The comma-separated names of the A/B metrics in the group.
	//
	// example:
	//
	// pv,uv
	ABMetricNames *string `json:"ABMetricNames,omitempty" xml:"ABMetricNames,omitempty"`
	// The description of the A/B metric group.
	//
	// example:
	//
	// 访问量相关指标
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The A/B metric group name.
	//
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner.
	//
	// example:
	//
	// 2799614***
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Indicates whether the A/B metric group is a real-time group.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListABMetricGroupsResponseBodyABMetricGroups) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricGroupsResponseBodyABMetricGroups) GoString() string {
	return s.String()
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetABMetricGroupId() *string {
	return s.ABMetricGroupId
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetABMetricIds() *string {
	return s.ABMetricIds
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetABMetricNames() *string {
	return s.ABMetricNames
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetDescription() *string {
	return s.Description
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetName() *string {
	return s.Name
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetOwner() *string {
	return s.Owner
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetRealtime() *bool {
	return s.Realtime
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) GetSceneId() *string {
	return s.SceneId
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetABMetricGroupId(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.ABMetricGroupId = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetABMetricIds(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.ABMetricIds = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetABMetricNames(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.ABMetricNames = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetDescription(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Description = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetName(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Name = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetOwner(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Owner = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetRealtime(v bool) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.Realtime = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) SetSceneId(v string) *ListABMetricGroupsResponseBodyABMetricGroups {
	s.SceneId = &v
	return s
}

func (s *ListABMetricGroupsResponseBodyABMetricGroups) Validate() error {
	return dara.Validate(s)
}
