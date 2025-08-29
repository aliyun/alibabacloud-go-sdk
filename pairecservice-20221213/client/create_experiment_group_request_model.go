// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateExperimentGroupRequest
	GetConfig() *string
	SetCrowdId(v string) *CreateExperimentGroupRequest
	GetCrowdId() *string
	SetCrowdTargetType(v string) *CreateExperimentGroupRequest
	GetCrowdTargetType() *string
	SetDebugCrowdId(v string) *CreateExperimentGroupRequest
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *CreateExperimentGroupRequest
	GetDebugUsers() *string
	SetDescription(v string) *CreateExperimentGroupRequest
	GetDescription() *string
	SetDistributionTimeDuration(v int32) *CreateExperimentGroupRequest
	GetDistributionTimeDuration() *int32
	SetDistributionType(v string) *CreateExperimentGroupRequest
	GetDistributionType() *string
	SetFilter(v string) *CreateExperimentGroupRequest
	GetFilter() *string
	SetInstanceId(v string) *CreateExperimentGroupRequest
	GetInstanceId() *string
	SetLayerId(v string) *CreateExperimentGroupRequest
	GetLayerId() *string
	SetName(v string) *CreateExperimentGroupRequest
	GetName() *string
	SetNeedAA(v bool) *CreateExperimentGroupRequest
	GetNeedAA() *bool
	SetRandomFlow(v int64) *CreateExperimentGroupRequest
	GetRandomFlow() *int64
	SetReservedBuckets(v string) *CreateExperimentGroupRequest
	GetReservedBuckets() *string
}

type CreateExperimentGroupRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 1
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 3
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
	// 3
	DistributionTimeDuration *int32 `json:"DistributionTimeDuration,omitempty" xml:"DistributionTimeDuration,omitempty"`
	// example:
	//
	// UserId
	DistributionType *string `json:"DistributionType,omitempty" xml:"DistributionType,omitempty"`
	// example:
	//
	// gender=male
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// experiment_group_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedAA     *bool  `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	RandomFlow *int64 `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// example:
	//
	// 1,2,3
	ReservedBuckets *string `json:"ReservedBuckets,omitempty" xml:"ReservedBuckets,omitempty"`
}

func (s CreateExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentGroupRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateExperimentGroupRequest) GetCrowdId() *string {
	return s.CrowdId
}

func (s *CreateExperimentGroupRequest) GetCrowdTargetType() *string {
	return s.CrowdTargetType
}

func (s *CreateExperimentGroupRequest) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *CreateExperimentGroupRequest) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *CreateExperimentGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExperimentGroupRequest) GetDistributionTimeDuration() *int32 {
	return s.DistributionTimeDuration
}

func (s *CreateExperimentGroupRequest) GetDistributionType() *string {
	return s.DistributionType
}

func (s *CreateExperimentGroupRequest) GetFilter() *string {
	return s.Filter
}

func (s *CreateExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExperimentGroupRequest) GetLayerId() *string {
	return s.LayerId
}

func (s *CreateExperimentGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateExperimentGroupRequest) GetNeedAA() *bool {
	return s.NeedAA
}

func (s *CreateExperimentGroupRequest) GetRandomFlow() *int64 {
	return s.RandomFlow
}

func (s *CreateExperimentGroupRequest) GetReservedBuckets() *string {
	return s.ReservedBuckets
}

func (s *CreateExperimentGroupRequest) SetConfig(v string) *CreateExperimentGroupRequest {
	s.Config = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetCrowdId(v string) *CreateExperimentGroupRequest {
	s.CrowdId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetCrowdTargetType(v string) *CreateExperimentGroupRequest {
	s.CrowdTargetType = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDebugCrowdId(v string) *CreateExperimentGroupRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDebugUsers(v string) *CreateExperimentGroupRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDescription(v string) *CreateExperimentGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDistributionTimeDuration(v int32) *CreateExperimentGroupRequest {
	s.DistributionTimeDuration = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetDistributionType(v string) *CreateExperimentGroupRequest {
	s.DistributionType = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetFilter(v string) *CreateExperimentGroupRequest {
	s.Filter = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetInstanceId(v string) *CreateExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetLayerId(v string) *CreateExperimentGroupRequest {
	s.LayerId = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetName(v string) *CreateExperimentGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetNeedAA(v bool) *CreateExperimentGroupRequest {
	s.NeedAA = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetRandomFlow(v int64) *CreateExperimentGroupRequest {
	s.RandomFlow = &v
	return s
}

func (s *CreateExperimentGroupRequest) SetReservedBuckets(v string) *CreateExperimentGroupRequest {
	s.ReservedBuckets = &v
	return s
}

func (s *CreateExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
