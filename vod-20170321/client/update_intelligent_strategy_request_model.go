// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntelligentStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v string) *UpdateIntelligentStrategyRequest
	GetConditions() *string
	SetExecuteParams(v string) *UpdateIntelligentStrategyRequest
	GetExecuteParams() *string
	SetName(v string) *UpdateIntelligentStrategyRequest
	GetName() *string
	SetPriority(v int32) *UpdateIntelligentStrategyRequest
	GetPriority() *int32
	SetStartTime(v string) *UpdateIntelligentStrategyRequest
	GetStartTime() *string
	SetState(v string) *UpdateIntelligentStrategyRequest
	GetState() *string
	SetStrategyId(v string) *UpdateIntelligentStrategyRequest
	GetStrategyId() *string
}

type UpdateIntelligentStrategyRequest struct {
	Conditions    *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	ExecuteParams *string `json:"ExecuteParams,omitempty" xml:"ExecuteParams,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority      *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	// This parameter is required.
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
}

func (s UpdateIntelligentStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntelligentStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateIntelligentStrategyRequest) GetConditions() *string {
	return s.Conditions
}

func (s *UpdateIntelligentStrategyRequest) GetExecuteParams() *string {
	return s.ExecuteParams
}

func (s *UpdateIntelligentStrategyRequest) GetName() *string {
	return s.Name
}

func (s *UpdateIntelligentStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateIntelligentStrategyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateIntelligentStrategyRequest) GetState() *string {
	return s.State
}

func (s *UpdateIntelligentStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateIntelligentStrategyRequest) SetConditions(v string) *UpdateIntelligentStrategyRequest {
	s.Conditions = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) SetExecuteParams(v string) *UpdateIntelligentStrategyRequest {
	s.ExecuteParams = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) SetName(v string) *UpdateIntelligentStrategyRequest {
	s.Name = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) SetPriority(v int32) *UpdateIntelligentStrategyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) SetStartTime(v string) *UpdateIntelligentStrategyRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) SetState(v string) *UpdateIntelligentStrategyRequest {
	s.State = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) SetStrategyId(v string) *UpdateIntelligentStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateIntelligentStrategyRequest) Validate() error {
	return dara.Validate(s)
}
