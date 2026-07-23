// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *GetExperimentGroupResponseBody
	GetConfig() *string
	SetCrowdId(v string) *GetExperimentGroupResponseBody
	GetCrowdId() *string
	SetCrowdTargetType(v string) *GetExperimentGroupResponseBody
	GetCrowdTargetType() *string
	SetDebugCrowdId(v string) *GetExperimentGroupResponseBody
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *GetExperimentGroupResponseBody
	GetDebugUsers() *string
	SetDescription(v string) *GetExperimentGroupResponseBody
	GetDescription() *string
	SetDistributionTimeDuration(v int32) *GetExperimentGroupResponseBody
	GetDistributionTimeDuration() *int32
	SetDistributionType(v string) *GetExperimentGroupResponseBody
	GetDistributionType() *string
	SetFilter(v string) *GetExperimentGroupResponseBody
	GetFilter() *string
	SetHoldingBuckets(v string) *GetExperimentGroupResponseBody
	GetHoldingBuckets() *string
	SetLaboratoryId(v string) *GetExperimentGroupResponseBody
	GetLaboratoryId() *string
	SetLayerId(v string) *GetExperimentGroupResponseBody
	GetLayerId() *string
	SetName(v string) *GetExperimentGroupResponseBody
	GetName() *string
	SetNeedAA(v bool) *GetExperimentGroupResponseBody
	GetNeedAA() *bool
	SetOwner(v string) *GetExperimentGroupResponseBody
	GetOwner() *string
	SetRandomFlow(v int64) *GetExperimentGroupResponseBody
	GetRandomFlow() *int64
	SetRequestId(v string) *GetExperimentGroupResponseBody
	GetRequestId() *string
	SetReservedBuckets(v string) *GetExperimentGroupResponseBody
	GetReservedBuckets() *string
	SetSceneId(v string) *GetExperimentGroupResponseBody
	GetSceneId() *string
	SetStatus(v string) *GetExperimentGroupResponseBody
	GetStatus() *string
}

type GetExperimentGroupResponseBody struct {
	// The configuration of the experiment group.
	//
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the crowd.
	//
	// example:
	//
	// 3
	CrowdId *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	// The method for targeting traffic to the experiment group. Valid values:
	//
	// - `All`: all traffic
	//
	// - `Filter`: traffic that matches the filter
	//
	// - `CrowdId`: traffic from the specified crowd
	//
	// - `Random`: a random percentage of traffic
	//
	// example:
	//
	// All
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// The ID of the debug crowd.
	//
	// example:
	//
	// 4
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// The user IDs for debugging, separated by commas.
	//
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// The description of the experiment group.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The distribution duration.
	//
	// example:
	//
	// 5
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// The traffic distribution method.
	//
	// ● `UserId`: Distributes traffic by user ID.
	//
	// ● `TimeDuration`: Distributes traffic by time period.
	//
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// The filter.
	//
	// example:
	//
	// gender=female
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The bucket IDs that correspond to the `RandomFlow` percentage. This parameter is returned only when `CrowdTargetType` is set to `Random`.
	//
	// example:
	//
	// 1,2,3,4
	HoldingBuckets *string `json:"HoldingBuckets,omitempty" xml:"HoldingBuckets,omitempty"`
	// The ID of the laboratory.
	//
	// example:
	//
	// 4
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// The ID of the layer.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// The name of the experiment group.
	//
	// example:
	//
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether an A/B test is required.
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
	// The percentage of traffic from 0 to 100 that is randomly allocated to this experiment group. This parameter is returned only when `CrowdTargetType` is set to `Random`.
	//
	// example:
	//
	// 20
	RandomFlow *int64 `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BDB621CB-A81E-5D39-8793-39A365CBCC74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The reserved bucket IDs.
	//
	// example:
	//
	// 1,2,3,4
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
	// The ID of the scene.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The status of the experiment group. Valid values:
	//
	// ● `Offline`: The experiment group is not launched.
	//
	// ● `Online`: The experiment group is launched.
	//
	// ● `Pushed`: The experiment group is fully launched.
	//
	// example:
	//
	// Offline
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentGroupResponseBody) GetConfig() *string {
	return s.Config
}

func (s *GetExperimentGroupResponseBody) GetCrowdId() *string {
	return s.CrowdId
}

func (s *GetExperimentGroupResponseBody) GetCrowdTargetType() *string {
	return s.CrowdTargetType
}

func (s *GetExperimentGroupResponseBody) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *GetExperimentGroupResponseBody) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *GetExperimentGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetExperimentGroupResponseBody) GetDistributionTimeDuration() *int32 {
	return s.DistributionTimeDuration
}

func (s *GetExperimentGroupResponseBody) GetDistributionType() *string {
	return s.DistributionType
}

func (s *GetExperimentGroupResponseBody) GetFilter() *string {
	return s.Filter
}

func (s *GetExperimentGroupResponseBody) GetHoldingBuckets() *string {
	return s.HoldingBuckets
}

func (s *GetExperimentGroupResponseBody) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *GetExperimentGroupResponseBody) GetLayerId() *string {
	return s.LayerId
}

func (s *GetExperimentGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *GetExperimentGroupResponseBody) GetNeedAA() *bool {
	return s.NeedAA
}

func (s *GetExperimentGroupResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetExperimentGroupResponseBody) GetRandomFlow() *int64 {
	return s.RandomFlow
}

func (s *GetExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExperimentGroupResponseBody) GetReservedBuckets() *string {
	return s.ReservedBuckets
}

func (s *GetExperimentGroupResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetExperimentGroupResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetExperimentGroupResponseBody) SetConfig(v string) *GetExperimentGroupResponseBody {
	s.Config = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetCrowdId(v string) *GetExperimentGroupResponseBody {
	s.CrowdId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetCrowdTargetType(v string) *GetExperimentGroupResponseBody {
	s.CrowdTargetType = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDebugCrowdId(v string) *GetExperimentGroupResponseBody {
	s.DebugCrowdId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDebugUsers(v string) *GetExperimentGroupResponseBody {
	s.DebugUsers = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDescription(v string) *GetExperimentGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDistributionTimeDuration(v int32) *GetExperimentGroupResponseBody {
	s.DistributionTimeDuration = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetDistributionType(v string) *GetExperimentGroupResponseBody {
	s.DistributionType = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetFilter(v string) *GetExperimentGroupResponseBody {
	s.Filter = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetHoldingBuckets(v string) *GetExperimentGroupResponseBody {
	s.HoldingBuckets = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetLaboratoryId(v string) *GetExperimentGroupResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetLayerId(v string) *GetExperimentGroupResponseBody {
	s.LayerId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetName(v string) *GetExperimentGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetNeedAA(v bool) *GetExperimentGroupResponseBody {
	s.NeedAA = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetOwner(v string) *GetExperimentGroupResponseBody {
	s.Owner = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetRandomFlow(v int64) *GetExperimentGroupResponseBody {
	s.RandomFlow = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetRequestId(v string) *GetExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetReservedBuckets(v string) *GetExperimentGroupResponseBody {
	s.ReservedBuckets = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetSceneId(v string) *GetExperimentGroupResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetExperimentGroupResponseBody) SetStatus(v string) *GetExperimentGroupResponseBody {
	s.Status = &v
	return s
}

func (s *GetExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
