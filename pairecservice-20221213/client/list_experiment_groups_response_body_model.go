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
	ExperimentGroups []*ListExperimentGroupsResponseBodyExperimentGroups `json:"ExperimentGroups,omitempty" xml:"ExperimentGroups,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ListExperimentGroupsResponseBodyExperimentGroups struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 4
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 5
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// example:
	//
	// gender=female
	Filter         *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	HoldingBuckets *string `json:"HoldingBuckets,omitempty" xml:"HoldingBuckets,omitempty"`
	// example:
	//
	// 4
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NeedAA *bool `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	// example:
	//
	// 1124512470******
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RandomFlow *int64  `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// example:
	//
	// 1,2,3,4
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
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
