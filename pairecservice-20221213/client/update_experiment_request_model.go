// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *UpdateExperimentRequest
	GetConfig() *string
	SetDebugCrowdId(v string) *UpdateExperimentRequest
	GetDebugCrowdId() *string
	SetDebugUsers(v string) *UpdateExperimentRequest
	GetDebugUsers() *string
	SetDescription(v string) *UpdateExperimentRequest
	GetDescription() *string
	SetFlowPercent(v int32) *UpdateExperimentRequest
	GetFlowPercent() *int32
	SetInstanceId(v string) *UpdateExperimentRequest
	GetInstanceId() *string
	SetName(v string) *UpdateExperimentRequest
	GetName() *string
	SetType(v string) *UpdateExperimentRequest
	GetType() *string
}

type UpdateExperimentRequest struct {
	// example:
	//
	// {"RankBy": "Score"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DebugCrowdId *string `json:"DebugCrowdId,omitempty" xml:"DebugCrowdId,omitempty"`
	// example:
	//
	// 1124512470******,1124512471******,1124512472******
	DebugUsers *string `json:"DebugUsers,omitempty" xml:"DebugUsers,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 100
	FlowPercent *int32 `json:"FlowPercent,omitempty" xml:"FlowPercent,omitempty"`
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
	// experiment_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Baseline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentRequest) GetConfig() *string {
	return s.Config
}

func (s *UpdateExperimentRequest) GetDebugCrowdId() *string {
	return s.DebugCrowdId
}

func (s *UpdateExperimentRequest) GetDebugUsers() *string {
	return s.DebugUsers
}

func (s *UpdateExperimentRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExperimentRequest) GetFlowPercent() *int32 {
	return s.FlowPercent
}

func (s *UpdateExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateExperimentRequest) GetName() *string {
	return s.Name
}

func (s *UpdateExperimentRequest) GetType() *string {
	return s.Type
}

func (s *UpdateExperimentRequest) SetConfig(v string) *UpdateExperimentRequest {
	s.Config = &v
	return s
}

func (s *UpdateExperimentRequest) SetDebugCrowdId(v string) *UpdateExperimentRequest {
	s.DebugCrowdId = &v
	return s
}

func (s *UpdateExperimentRequest) SetDebugUsers(v string) *UpdateExperimentRequest {
	s.DebugUsers = &v
	return s
}

func (s *UpdateExperimentRequest) SetDescription(v string) *UpdateExperimentRequest {
	s.Description = &v
	return s
}

func (s *UpdateExperimentRequest) SetFlowPercent(v int32) *UpdateExperimentRequest {
	s.FlowPercent = &v
	return s
}

func (s *UpdateExperimentRequest) SetInstanceId(v string) *UpdateExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateExperimentRequest) SetName(v string) *UpdateExperimentRequest {
	s.Name = &v
	return s
}

func (s *UpdateExperimentRequest) SetType(v string) *UpdateExperimentRequest {
	s.Type = &v
	return s
}

func (s *UpdateExperimentRequest) Validate() error {
	return dara.Validate(s)
}
