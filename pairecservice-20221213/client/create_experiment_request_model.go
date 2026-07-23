// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateExperimentRequest
	GetConfig() *string
	SetDebugCrowdId(v string) *CreateExperimentRequest
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *CreateExperimentRequest
	GetDebugUsers() *string
	SetDescription(v string) *CreateExperimentRequest
	GetDescription() *string
	SetExperimentGroupId(v string) *CreateExperimentRequest
	GetExperimentGroupId() *string
	SetFlowPercent(v int32) *CreateExperimentRequest
	GetFlowPercent() *int32
	SetInstanceId(v string) *CreateExperimentRequest
	GetInstanceId() *string
	SetName(v string) *CreateExperimentRequest
	GetName() *string
	SetType(v string) *CreateExperimentRequest
	GetType() *string
}

type CreateExperimentRequest struct {
	// The experiment configuration.
	//
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The ID of the debug crowd. Call the ListCrowds operation to obtain this ID.
	//
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// The UIDs of Alibaba Cloud accounts or RAM users for debugging. Separate multiple UIDs with a comma.
	//
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// The experiment description.
	//
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the experiment group. Call the ListExperimentGroups operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// The traffic distribution percentage.
	//
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
	// The instance ID. Call the ListInstances operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The experiment name.
	//
	// This parameter is required.
	//
	// example:
	//
	// experiment_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The experiment type. Valid values:<br>● `Baseline`: Indicates a baseline experiment.<br>● `Normal`: Indicates a normal experiment.<br><br>
	//
	// This parameter is required.
	//
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateExperimentRequest) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *CreateExperimentRequest) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *CreateExperimentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExperimentRequest) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *CreateExperimentRequest) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *CreateExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExperimentRequest) GetName() *string {
	return s.Name
}

func (s *CreateExperimentRequest) GetType() *string {
	return s.Type
}

func (s *CreateExperimentRequest) SetConfig(v string) *CreateExperimentRequest {
	s.Config = &v
	return s
}

func (s *CreateExperimentRequest) SetDebugCrowdId(v string) *CreateExperimentRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *CreateExperimentRequest) SetDebugUsers(v string) *CreateExperimentRequest {
	s.DebugUsers = &v
	return s
}

func (s *CreateExperimentRequest) SetDescription(v string) *CreateExperimentRequest {
	s.Description = &v
	return s
}

func (s *CreateExperimentRequest) SetExperimentGroupId(v string) *CreateExperimentRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *CreateExperimentRequest) SetFlowPercent(v int32) *CreateExperimentRequest {
	s.FlowPercent = &v
	return s
}

func (s *CreateExperimentRequest) SetInstanceId(v string) *CreateExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExperimentRequest) SetName(v string) *CreateExperimentRequest {
	s.Name = &v
	return s
}

func (s *CreateExperimentRequest) SetType(v string) *CreateExperimentRequest {
	s.Type = &v
	return s
}

func (s *CreateExperimentRequest) Validate() error {
	return dara.Validate(s)
}
