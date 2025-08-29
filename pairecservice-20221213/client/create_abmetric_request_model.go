// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *CreateABMetricRequest
	GetDefinition() *string
	SetDescription(v string) *CreateABMetricRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateABMetricRequest
	GetInstanceId() *string
	SetLeftMetricId(v string) *CreateABMetricRequest
	GetLeftMetricId() *string
	SetName(v string) *CreateABMetricRequest
	GetName() *string
	SetOperator(v string) *CreateABMetricRequest
	GetOperator() *string
	SetRealtime(v bool) *CreateABMetricRequest
	GetRealtime() *bool
	SetResultResourceId(v string) *CreateABMetricRequest
	GetResultResourceId() *string
	SetRightMetricId(v string) *CreateABMetricRequest
	GetRightMetricId() *string
	SetSceneId(v string) *CreateABMetricRequest
	GetSceneId() *string
	SetStatisticsCycle(v int32) *CreateABMetricRequest
	GetStatisticsCycle() *int32
	SetTableMetaId(v string) *CreateABMetricRequest
	GetTableMetaId() *string
	SetType(v string) *CreateABMetricRequest
	GetType() *string
}

type CreateABMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sum(click_cnt)
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 2
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 3
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateABMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateABMetricRequest) GoString() string {
	return s.String()
}

func (s *CreateABMetricRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateABMetricRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateABMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateABMetricRequest) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *CreateABMetricRequest) GetName() *string {
	return s.Name
}

func (s *CreateABMetricRequest) GetOperator() *string {
	return s.Operator
}

func (s *CreateABMetricRequest) GetRealtime() *bool {
	return s.Realtime
}

func (s *CreateABMetricRequest) GetResultResourceId() *string {
	return s.ResultResourceId
}

func (s *CreateABMetricRequest) GetRightMetricId() *string {
	return s.RightMetricId
}

func (s *CreateABMetricRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateABMetricRequest) GetStatisticsCycle() *int32 {
	return s.StatisticsCycle
}

func (s *CreateABMetricRequest) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *CreateABMetricRequest) GetType() *string {
	return s.Type
}

func (s *CreateABMetricRequest) SetDefinition(v string) *CreateABMetricRequest {
	s.Definition = &v
	return s
}

func (s *CreateABMetricRequest) SetDescription(v string) *CreateABMetricRequest {
	s.Description = &v
	return s
}

func (s *CreateABMetricRequest) SetInstanceId(v string) *CreateABMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateABMetricRequest) SetLeftMetricId(v string) *CreateABMetricRequest {
	s.LeftMetricId = &v
	return s
}

func (s *CreateABMetricRequest) SetName(v string) *CreateABMetricRequest {
	s.Name = &v
	return s
}

func (s *CreateABMetricRequest) SetOperator(v string) *CreateABMetricRequest {
	s.Operator = &v
	return s
}

func (s *CreateABMetricRequest) SetRealtime(v bool) *CreateABMetricRequest {
	s.Realtime = &v
	return s
}

func (s *CreateABMetricRequest) SetResultResourceId(v string) *CreateABMetricRequest {
	s.ResultResourceId = &v
	return s
}

func (s *CreateABMetricRequest) SetRightMetricId(v string) *CreateABMetricRequest {
	s.RightMetricId = &v
	return s
}

func (s *CreateABMetricRequest) SetSceneId(v string) *CreateABMetricRequest {
	s.SceneId = &v
	return s
}

func (s *CreateABMetricRequest) SetStatisticsCycle(v int32) *CreateABMetricRequest {
	s.StatisticsCycle = &v
	return s
}

func (s *CreateABMetricRequest) SetTableMetaId(v string) *CreateABMetricRequest {
	s.TableMetaId = &v
	return s
}

func (s *CreateABMetricRequest) SetType(v string) *CreateABMetricRequest {
	s.Type = &v
	return s
}

func (s *CreateABMetricRequest) Validate() error {
	return dara.Validate(s)
}
