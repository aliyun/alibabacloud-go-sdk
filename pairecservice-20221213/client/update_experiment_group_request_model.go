// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateExperimentGroupRequest
	GetConfig() *string
	SetCrowdId(v string) *UpdateExperimentGroupRequest
	GetCrowdId() *string
	SetCrowdTargetType(v string) *UpdateExperimentGroupRequest
	GetCrowdTargetType() *string
	SetDebugCrowdId(v string) *UpdateExperimentGroupRequest
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *UpdateExperimentGroupRequest
	GetDebugUsers() *string
	SetDescription(v string) *UpdateExperimentGroupRequest
	GetDescription() *string
	SetDistributionTimeDuration(v int32) *UpdateExperimentGroupRequest
	GetDistributionTimeDuration() *int32
	SetDistributionType(v string) *UpdateExperimentGroupRequest
	GetDistributionType() *string
	SetFilter(v string) *UpdateExperimentGroupRequest
	GetFilter() *string
	SetInstanceId(v string) *UpdateExperimentGroupRequest
	GetInstanceId() *string
	SetLayerId(v string) *UpdateExperimentGroupRequest
	GetLayerId() *string
	SetName(v string) *UpdateExperimentGroupRequest
	GetName() *string
	SetNeedAA(v bool) *UpdateExperimentGroupRequest
	GetNeedAA() *bool
	SetRandomFlow(v int64) *UpdateExperimentGroupRequest
	GetRandomFlow() *int64
	SetReservcedBuckets(v string) *UpdateExperimentGroupRequest
	GetReservcedBuckets() *string
}

type UpdateExperimentGroupRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	CrowdId         *string `json:"CrowdId,omitempty" xml:"CrowdId,omitempty"`
	CrowdTargetType *string `json:"CrowdTargetType,omitempty" xml:"CrowdTargetType,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// user1,user2,user3
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// This parameter is required.
	//
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
	// pairec-cn-abcdefg1234
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
	// experiment_group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NeedAA     *bool  `json:"NeedAA,omitempty" xml:"NeedAA,omitempty"`
	RandomFlow *int64 `json:"RandomFlow,omitempty" xml:"RandomFlow,omitempty"`
	// example:
	//
	// 1,2,3
	ReservcedBuckets *string `json:"ReservcedBuckets,omitempty" xml:"ReservcedBuckets,omitempty"`
}

func (s UpdateExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentGroupRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateExperimentGroupRequest) GetCrowdId() *string {
	return s.CrowdId
}

func (s *UpdateExperimentGroupRequest) GetCrowdTargetType() *string {
	return s.CrowdTargetType
}

func (s *UpdateExperimentGroupRequest) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *UpdateExperimentGroupRequest) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *UpdateExperimentGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExperimentGroupRequest) GetDistributionTimeDuration() *int32 {
	return s.DistributionTimeDuration
}

func (s *UpdateExperimentGroupRequest) GetDistributionType() *string {
	return s.DistributionType
}

func (s *UpdateExperimentGroupRequest) GetFilter() *string {
	return s.Filter
}

func (s *UpdateExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateExperimentGroupRequest) GetLayerId() *string {
	return s.LayerId
}

func (s *UpdateExperimentGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateExperimentGroupRequest) GetNeedAA() *bool {
	return s.NeedAA
}

func (s *UpdateExperimentGroupRequest) GetRandomFlow() *int64 {
	return s.RandomFlow
}

func (s *UpdateExperimentGroupRequest) GetReservcedBuckets() *string {
	return s.ReservcedBuckets
}

func (s *UpdateExperimentGroupRequest) SetConfig(v string) *UpdateExperimentGroupRequest {
	s.Config = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetCrowdId(v string) *UpdateExperimentGroupRequest {
	s.CrowdId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetCrowdTargetType(v string) *UpdateExperimentGroupRequest {
	s.CrowdTargetType = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDebugCrowdId(v string) *UpdateExperimentGroupRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDebugUsers(v string) *UpdateExperimentGroupRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDescription(v string) *UpdateExperimentGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDistributionTimeDuration(v int32) *UpdateExperimentGroupRequest {
	s.DistributionTimeDuration = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetDistributionType(v string) *UpdateExperimentGroupRequest {
	s.DistributionType = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetFilter(v string) *UpdateExperimentGroupRequest {
	s.Filter = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetInstanceId(v string) *UpdateExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetLayerId(v string) *UpdateExperimentGroupRequest {
	s.LayerId = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetName(v string) *UpdateExperimentGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetNeedAA(v bool) *UpdateExperimentGroupRequest {
	s.NeedAA = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetRandomFlow(v int64) *UpdateExperimentGroupRequest {
	s.RandomFlow = &v
	return s
}

func (s *UpdateExperimentGroupRequest) SetReservcedBuckets(v string) *UpdateExperimentGroupRequest {
	s.ReservcedBuckets = &v
	return s
}

func (s *UpdateExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
