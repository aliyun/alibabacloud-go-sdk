// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExperimentGroups(v []*ListExperimentGroupsResponseBodyExperimentGroups) *ListExperimentGroupsResponseBody
	GetExperimentGroups() []*ListExperimentGroupsResponseBodyExperimentGroups
	SetRequestId(v string) *ListExperimentGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListExperimentGroupsResponseBody
	GetTotalCount() *int64
}

type ListExperimentGroupsResponseBody struct {
	// The list of experiment groups.
	ExperimentGroups []*ListExperimentGroupsResponseBodyExperimentGroups `json:"ExperimentGroups,omitempty" xml:"ExperimentGroups,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsResponseBody) GetExperimentGroups() []*ListExperimentGroupsResponseBodyExperimentGroups {
	return s.ExperimentGroups
}

func (s *ListExperimentGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExperimentGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExperimentGroupsResponseBody) SetExperimentGroups(v []*ListExperimentGroupsResponseBodyExperimentGroups) *ListExperimentGroupsResponseBody {
	s.ExperimentGroups = v
	return s
}

func (s *ListExperimentGroupsResponseBody) SetRequestId(v string) *ListExperimentGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentGroupsResponseBody) SetTotalCount(v int64) *ListExperimentGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExperimentGroupsResponseBody) Validate() error {
	if s.ExperimentGroups != nil {
		for _, item := range s.ExperimentGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExperimentGroupsResponseBodyExperimentGroups struct {
	// The configuration for the experiment group, in JSON format.
	//
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The crowd ID.
	//
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// The traffic targeting method for the experiment group. Valid values:
	//
	// - `All`: All traffic.
	//
	// - `Filter`: Traffic that matches the filter.
	//
	// - `CrowdId`: Traffic from a specified crowd.
	//
	// - `Random`: A random percentage of traffic.
	//
	// example:
	//
	// All
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// The debug crowd ID.
	//
	// example:
	//
	// 4
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// The IDs of debug users, separated by commas.
	//
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// The experiment group description.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The distribution duration. This parameter is required only when `DistributionType` is `TimeDuration`.
	//
	// example:
	//
	// 5
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// The traffic distribution method.<br>● `UserId`: by user ID<br>● `TimeDuration`: by time duration<br><br>
	//
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The experiment group ID.
	//
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// The filter condition.
	//
	// example:
	//
	// gender=female
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// If `CrowdTargetType` is set to `Random`, this parameter returns the comma-separated IDs of buckets allocated based on the `RandomFlow` value.
	//
	// example:
	//
	// 1,2,3,4
	HoldingBuckets *string `json:"HoldingBuckets,omitempty" xml:"HoldingBuckets,omitempty"`
	// The laboratory ID.
	//
	// example:
	//
	// 4
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// The layer ID.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// The experiment group name.
	//
	// example:
	//
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable A/B testing for the experiment group.
	//
	// example:
	//
	// true
	NeedAA *bool `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	// The owner of the experiment group.
	//
	// example:
	//
	// 1124512470******
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// If `CrowdTargetType` is set to `Random`, this parameter specifies the percentage of traffic (an integer from 0 to 100) randomly allocated to the experiment group.
	//
	// example:
	//
	// 20
	RandomFlow *int64 `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// The IDs of reserved buckets, separated by commas.
	//
	// example:
	//
	// 1,2,3,4
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The status of the experiment group. Valid values:<br>● `Offline`: The experiment group is inactive.<br>● `Online`: The experiment group is active.<br>● `Pushed`: The experiment group is fully rolled out.<br><br><br>
	//
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListExperimentGroupsResponseBodyExperimentGroups) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentGroupsResponseBodyExperimentGroups) GoString() string {
	return s.String()
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetConfig() *string {
	return s.Config
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetCrowdId() *string {
	return s.CrowdId
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetCrowdTargetType() *string {
	return s.CrowdTargetType
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetDescription() *string {
	return s.Description
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetDistributionTimeDuration() *int32 {
	return s.DistributionTimeDuration
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetDistributionType() *string {
	return s.DistributionType
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetFilter() *string {
	return s.Filter
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetHoldingBuckets() *string {
	return s.HoldingBuckets
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetLayerId() *string {
	return s.LayerId
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetName() *string {
	return s.Name
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetNeedAA() *bool {
	return s.NeedAA
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetOwner() *string {
	return s.Owner
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetRandomFlow() *int64 {
	return s.RandomFlow
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetReservedBuckets() *string {
	return s.ReservedBuckets
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetSceneId() *string {
	return s.SceneId
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) GetStatus() *string {
	return s.Status
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetConfig(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Config = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetCrowdId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.CrowdId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetCrowdTargetType(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.CrowdTargetType = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDebugCrowdId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DebugCrowdId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDebugUsers(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DebugUsers = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDescription(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Description = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDistributionTimeDuration(v int32) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DistributionTimeDuration = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetDistributionType(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.DistributionType = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetExperimentGroupId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.ExperimentGroupId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetFilter(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Filter = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetHoldingBuckets(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.HoldingBuckets = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetLaboratoryId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.LaboratoryId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetLayerId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.LayerId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetName(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Name = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetNeedAA(v bool) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.NeedAA = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetOwner(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Owner = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetRandomFlow(v int64) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.RandomFlow = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetReservedBuckets(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.ReservedBuckets = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetSceneId(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.SceneId = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) SetStatus(v string) *ListExperimentGroupsResponseBodyExperimentGroups {
	s.Status = &v
	return s
}

func (s *ListExperimentGroupsResponseBodyExperimentGroups) Validate() error {
	return dara.Validate(s)
}
