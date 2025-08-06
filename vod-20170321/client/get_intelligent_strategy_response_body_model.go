// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntelligentStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntelligentStrategy(v *GetIntelligentStrategyResponseBodyIntelligentStrategy) *GetIntelligentStrategyResponseBody
	GetIntelligentStrategy() *GetIntelligentStrategyResponseBodyIntelligentStrategy
	SetRequestId(v string) *GetIntelligentStrategyResponseBody
	GetRequestId() *string
}

type GetIntelligentStrategyResponseBody struct {
	IntelligentStrategy *GetIntelligentStrategyResponseBodyIntelligentStrategy `json:"IntelligentStrategy,omitempty" xml:"IntelligentStrategy,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIntelligentStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntelligentStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntelligentStrategyResponseBody) GetIntelligentStrategy() *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	return s.IntelligentStrategy
}

func (s *GetIntelligentStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntelligentStrategyResponseBody) SetIntelligentStrategy(v *GetIntelligentStrategyResponseBodyIntelligentStrategy) *GetIntelligentStrategyResponseBody {
	s.IntelligentStrategy = v
	return s
}

func (s *GetIntelligentStrategyResponseBody) SetRequestId(v string) *GetIntelligentStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntelligentStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIntelligentStrategyResponseBodyIntelligentStrategy struct {
	Conditions    *string `json:"Conditions,omitempty" xml:"Conditions,omitempty"`
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExecuteParams *string `json:"ExecuteParams,omitempty" xml:"ExecuteParams,omitempty"`
	ModifiedTime  *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority      *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	StrategyId    *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetIntelligentStrategyResponseBodyIntelligentStrategy) String() string {
	return dara.Prettify(s)
}

func (s GetIntelligentStrategyResponseBodyIntelligentStrategy) GoString() string {
	return s.String()
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetConditions() *string {
	return s.Conditions
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetExecuteParams() *string {
	return s.ExecuteParams
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetName() *string {
	return s.Name
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetPriority() *int32 {
	return s.Priority
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetStartTime() *string {
	return s.StartTime
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetState() *string {
	return s.State
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) GetType() *string {
	return s.Type
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetConditions(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.Conditions = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetCreationTime(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.CreationTime = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetExecuteParams(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.ExecuteParams = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetModifiedTime(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.ModifiedTime = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetName(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.Name = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetPriority(v int32) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.Priority = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetStartTime(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.StartTime = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetState(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.State = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetStrategyId(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.StrategyId = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) SetType(v string) *GetIntelligentStrategyResponseBodyIntelligentStrategy {
	s.Type = &v
	return s
}

func (s *GetIntelligentStrategyResponseBodyIntelligentStrategy) Validate() error {
	return dara.Validate(s)
}
