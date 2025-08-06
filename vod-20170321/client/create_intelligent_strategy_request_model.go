// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntelligentStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v string) *CreateIntelligentStrategyRequest
	GetConditions() *string
	SetExecuteParams(v string) *CreateIntelligentStrategyRequest
	GetExecuteParams() *string
	SetName(v string) *CreateIntelligentStrategyRequest
	GetName() *string
	SetPriority(v int32) *CreateIntelligentStrategyRequest
	GetPriority() *int32
	SetStartTime(v string) *CreateIntelligentStrategyRequest
	GetStartTime() *string
	SetState(v string) *CreateIntelligentStrategyRequest
	GetState() *string
	SetType(v string) *CreateIntelligentStrategyRequest
	GetType() *string
}

type CreateIntelligentStrategyRequest struct {
	// This parameter is required.
	Conditions *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	// This parameter is required.
	ExecuteParams *string `json:"ExecuteParams,omitempty" xml:"ExecuteParams,omitempty"`
	// This parameter is required.
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateIntelligentStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntelligentStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateIntelligentStrategyRequest) GetConditions() *string {
	return s.Conditions
}

func (s *CreateIntelligentStrategyRequest) GetExecuteParams() *string {
	return s.ExecuteParams
}

func (s *CreateIntelligentStrategyRequest) GetName() *string {
	return s.Name
}

func (s *CreateIntelligentStrategyRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *CreateIntelligentStrategyRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateIntelligentStrategyRequest) GetState() *string {
	return s.State
}

func (s *CreateIntelligentStrategyRequest) GetType() *string {
	return s.Type
}

func (s *CreateIntelligentStrategyRequest) SetConditions(v string) *CreateIntelligentStrategyRequest {
	s.Conditions = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) SetExecuteParams(v string) *CreateIntelligentStrategyRequest {
	s.ExecuteParams = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) SetName(v string) *CreateIntelligentStrategyRequest {
	s.Name = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) SetPriority(v int32) *CreateIntelligentStrategyRequest {
	s.Priority = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) SetStartTime(v string) *CreateIntelligentStrategyRequest {
	s.StartTime = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) SetState(v string) *CreateIntelligentStrategyRequest {
	s.State = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) SetType(v string) *CreateIntelligentStrategyRequest {
	s.Type = &v
	return s
}

func (s *CreateIntelligentStrategyRequest) Validate() error {
	return dara.Validate(s)
}
