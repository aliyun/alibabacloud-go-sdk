// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *UpdateABMetricRequest
	GetDefinition() *string
	SetDescription(v string) *UpdateABMetricRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateABMetricRequest
	GetInstanceId() *string
	SetLeftMetricId(v string) *UpdateABMetricRequest
	GetLeftMetricId() *string
	SetName(v string) *UpdateABMetricRequest
	GetName() *string
	SetOperator(v string) *UpdateABMetricRequest
	GetOperator() *string
	SetRealtime(v bool) *UpdateABMetricRequest
	GetRealtime() *bool
	SetResultResourceId(v string) *UpdateABMetricRequest
	GetResultResourceId() *string
	SetRightMetricId(v string) *UpdateABMetricRequest
	GetRightMetricId() *string
	SetSceneId(v string) *UpdateABMetricRequest
	GetSceneId() *string
	SetStatisticsCycle(v int32) *UpdateABMetricRequest
	GetStatisticsCycle() *int32
	SetTableMetaId(v string) *UpdateABMetricRequest
	GetTableMetaId() *string
	SetType(v string) *UpdateABMetricRequest
	GetType() *string
}

type UpdateABMetricRequest struct {
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
	// pairec-cn-test123
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

func (s UpdateABMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateABMetricRequest) GoString() string {
	return s.String()
}

func (s *UpdateABMetricRequest) GetDefinition() *string {
	return s.Definition
}

func (s *UpdateABMetricRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateABMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateABMetricRequest) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *UpdateABMetricRequest) GetName() *string {
	return s.Name
}

func (s *UpdateABMetricRequest) GetOperator() *string {
	return s.Operator
}

func (s *UpdateABMetricRequest) GetRealtime() *bool {
	return s.Realtime
}

func (s *UpdateABMetricRequest) GetResultResourceId() *string {
	return s.ResultResourceId
}

func (s *UpdateABMetricRequest) GetRightMetricId() *string {
	return s.RightMetricId
}

func (s *UpdateABMetricRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateABMetricRequest) GetStatisticsCycle() *int32 {
	return s.StatisticsCycle
}

func (s *UpdateABMetricRequest) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *UpdateABMetricRequest) GetType() *string {
	return s.Type
}

func (s *UpdateABMetricRequest) SetDefinition(v string) *UpdateABMetricRequest {
	s.Definition = &v
	return s
}

func (s *UpdateABMetricRequest) SetDescription(v string) *UpdateABMetricRequest {
	s.Description = &v
	return s
}

func (s *UpdateABMetricRequest) SetInstanceId(v string) *UpdateABMetricRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateABMetricRequest) SetLeftMetricId(v string) *UpdateABMetricRequest {
	s.LeftMetricId = &v
	return s
}

func (s *UpdateABMetricRequest) SetName(v string) *UpdateABMetricRequest {
	s.Name = &v
	return s
}

func (s *UpdateABMetricRequest) SetOperator(v string) *UpdateABMetricRequest {
	s.Operator = &v
	return s
}

func (s *UpdateABMetricRequest) SetRealtime(v bool) *UpdateABMetricRequest {
	s.Realtime = &v
	return s
}

func (s *UpdateABMetricRequest) SetResultResourceId(v string) *UpdateABMetricRequest {
	s.ResultResourceId = &v
	return s
}

func (s *UpdateABMetricRequest) SetRightMetricId(v string) *UpdateABMetricRequest {
	s.RightMetricId = &v
	return s
}

func (s *UpdateABMetricRequest) SetSceneId(v string) *UpdateABMetricRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateABMetricRequest) SetStatisticsCycle(v int32) *UpdateABMetricRequest {
	s.StatisticsCycle = &v
	return s
}

func (s *UpdateABMetricRequest) SetTableMetaId(v string) *UpdateABMetricRequest {
	s.TableMetaId = &v
	return s
}

func (s *UpdateABMetricRequest) SetType(v string) *UpdateABMetricRequest {
	s.Type = &v
	return s
}

func (s *UpdateABMetricRequest) Validate() error {
	return dara.Validate(s)
}
