// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v string) *GetABMetricResponseBody
	GetDefinition() *string
	SetDescription(v string) *GetABMetricResponseBody
	GetDescription() *string
	SetLeftMetricId(v string) *GetABMetricResponseBody
	GetLeftMetricId() *string
	SetName(v string) *GetABMetricResponseBody
	GetName() *string
	SetOperator(v string) *GetABMetricResponseBody
	GetOperator() *string
	SetRealtime(v string) *GetABMetricResponseBody
	GetRealtime() *string
	SetRequestId(v string) *GetABMetricResponseBody
	GetRequestId() *string
	SetResultResourceId(v string) *GetABMetricResponseBody
	GetResultResourceId() *string
	SetResultTableMetaId(v string) *GetABMetricResponseBody
	GetResultTableMetaId() *string
	SetRightMetricId(v string) *GetABMetricResponseBody
	GetRightMetricId() *string
	SetSceneId(v string) *GetABMetricResponseBody
	GetSceneId() *string
	SetSceneName(v string) *GetABMetricResponseBody
	GetSceneName() *string
	SetStatisticsCycle(v int32) *GetABMetricResponseBody
	GetStatisticsCycle() *int32
	SetTableMetaId(v string) *GetABMetricResponseBody
	GetTableMetaId() *string
	SetType(v string) *GetABMetricResponseBody
	GetType() *string
}

type GetABMetricResponseBody struct {
	// example:
	//
	// sum(click_cnt)
	Definition  *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// false
	Realtime *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 3
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	// example:
	//
	// 2
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// home_feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetABMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetABMetricResponseBody) GoString() string {
	return s.String()
}

func (s *GetABMetricResponseBody) GetDefinition() *string {
	return s.Definition
}

func (s *GetABMetricResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetABMetricResponseBody) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *GetABMetricResponseBody) GetName() *string {
	return s.Name
}

func (s *GetABMetricResponseBody) GetOperator() *string {
	return s.Operator
}

func (s *GetABMetricResponseBody) GetRealtime() *string {
	return s.Realtime
}

func (s *GetABMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetABMetricResponseBody) GetResultResourceId() *string {
	return s.ResultResourceId
}

func (s *GetABMetricResponseBody) GetResultTableMetaId() *string {
	return s.ResultTableMetaId
}

func (s *GetABMetricResponseBody) GetRightMetricId() *string {
	return s.RightMetricId
}

func (s *GetABMetricResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetABMetricResponseBody) GetSceneName() *string {
	return s.SceneName
}

func (s *GetABMetricResponseBody) GetStatisticsCycle() *int32 {
	return s.StatisticsCycle
}

func (s *GetABMetricResponseBody) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *GetABMetricResponseBody) GetType() *string {
	return s.Type
}

func (s *GetABMetricResponseBody) SetDefinition(v string) *GetABMetricResponseBody {
	s.Definition = &v
	return s
}

func (s *GetABMetricResponseBody) SetDescription(v string) *GetABMetricResponseBody {
	s.Description = &v
	return s
}

func (s *GetABMetricResponseBody) SetLeftMetricId(v string) *GetABMetricResponseBody {
	s.LeftMetricId = &v
	return s
}

func (s *GetABMetricResponseBody) SetName(v string) *GetABMetricResponseBody {
	s.Name = &v
	return s
}

func (s *GetABMetricResponseBody) SetOperator(v string) *GetABMetricResponseBody {
	s.Operator = &v
	return s
}

func (s *GetABMetricResponseBody) SetRealtime(v string) *GetABMetricResponseBody {
	s.Realtime = &v
	return s
}

func (s *GetABMetricResponseBody) SetRequestId(v string) *GetABMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetABMetricResponseBody) SetResultResourceId(v string) *GetABMetricResponseBody {
	s.ResultResourceId = &v
	return s
}

func (s *GetABMetricResponseBody) SetResultTableMetaId(v string) *GetABMetricResponseBody {
	s.ResultTableMetaId = &v
	return s
}

func (s *GetABMetricResponseBody) SetRightMetricId(v string) *GetABMetricResponseBody {
	s.RightMetricId = &v
	return s
}

func (s *GetABMetricResponseBody) SetSceneId(v string) *GetABMetricResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetABMetricResponseBody) SetSceneName(v string) *GetABMetricResponseBody {
	s.SceneName = &v
	return s
}

func (s *GetABMetricResponseBody) SetStatisticsCycle(v int32) *GetABMetricResponseBody {
	s.StatisticsCycle = &v
	return s
}

func (s *GetABMetricResponseBody) SetTableMetaId(v string) *GetABMetricResponseBody {
	s.TableMetaId = &v
	return s
}

func (s *GetABMetricResponseBody) SetType(v string) *GetABMetricResponseBody {
	s.Type = &v
	return s
}

func (s *GetABMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
