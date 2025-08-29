// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListABMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetABMetrics(v []*ListABMetricsResponseBodyABMetrics) *ListABMetricsResponseBody
	GetABMetrics() []*ListABMetricsResponseBodyABMetrics
	SetRequestId(v string) *ListABMetricsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListABMetricsResponseBody
	GetTotalCount() *int64
}

type ListABMetricsResponseBody struct {
	ABMetrics []*ListABMetricsResponseBodyABMetrics `json:"ABMetrics,omitempty" xml:"ABMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListABMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponseBody) GetABMetrics() []*ListABMetricsResponseBodyABMetrics {
	return s.ABMetrics
}

func (s *ListABMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListABMetricsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListABMetricsResponseBody) SetABMetrics(v []*ListABMetricsResponseBodyABMetrics) *ListABMetricsResponseBody {
	s.ABMetrics = v
	return s
}

func (s *ListABMetricsResponseBody) SetRequestId(v string) *ListABMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABMetricsResponseBody) SetTotalCount(v int64) *ListABMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListABMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListABMetricsResponseBodyABMetrics struct {
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
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
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// example:
	//
	// 2
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
	// 1
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// example:
	//
	// Single
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListABMetricsResponseBodyABMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricsResponseBodyABMetrics) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponseBodyABMetrics) GetABMetricId() *string {
	return s.ABMetricId
}

func (s *ListABMetricsResponseBodyABMetrics) GetDefinition() *string {
	return s.Definition
}

func (s *ListABMetricsResponseBodyABMetrics) GetDescription() *string {
	return s.Description
}

func (s *ListABMetricsResponseBodyABMetrics) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *ListABMetricsResponseBodyABMetrics) GetName() *string {
	return s.Name
}

func (s *ListABMetricsResponseBodyABMetrics) GetOperator() *string {
	return s.Operator
}

func (s *ListABMetricsResponseBodyABMetrics) GetRealtime() *string {
	return s.Realtime
}

func (s *ListABMetricsResponseBodyABMetrics) GetResultResourceId() *string {
	return s.ResultResourceId
}

func (s *ListABMetricsResponseBodyABMetrics) GetResultTableMetaId() *string {
	return s.ResultTableMetaId
}

func (s *ListABMetricsResponseBodyABMetrics) GetRightMetricId() *string {
	return s.RightMetricId
}

func (s *ListABMetricsResponseBodyABMetrics) GetSceneId() *string {
	return s.SceneId
}

func (s *ListABMetricsResponseBodyABMetrics) GetSceneName() *string {
	return s.SceneName
}

func (s *ListABMetricsResponseBodyABMetrics) GetStatisticsCycle() *int32 {
	return s.StatisticsCycle
}

func (s *ListABMetricsResponseBodyABMetrics) GetTableMetaId() *string {
	return s.TableMetaId
}

func (s *ListABMetricsResponseBodyABMetrics) GetType() *string {
	return s.Type
}

func (s *ListABMetricsResponseBodyABMetrics) SetABMetricId(v string) *ListABMetricsResponseBodyABMetrics {
	s.ABMetricId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDefinition(v string) *ListABMetricsResponseBodyABMetrics {
	s.Definition = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDescription(v string) *ListABMetricsResponseBodyABMetrics {
	s.Description = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetLeftMetricId(v string) *ListABMetricsResponseBodyABMetrics {
	s.LeftMetricId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetName(v string) *ListABMetricsResponseBodyABMetrics {
	s.Name = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetOperator(v string) *ListABMetricsResponseBodyABMetrics {
	s.Operator = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetRealtime(v string) *ListABMetricsResponseBodyABMetrics {
	s.Realtime = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetResultResourceId(v string) *ListABMetricsResponseBodyABMetrics {
	s.ResultResourceId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetResultTableMetaId(v string) *ListABMetricsResponseBodyABMetrics {
	s.ResultTableMetaId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetRightMetricId(v string) *ListABMetricsResponseBodyABMetrics {
	s.RightMetricId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetSceneId(v string) *ListABMetricsResponseBodyABMetrics {
	s.SceneId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetSceneName(v string) *ListABMetricsResponseBodyABMetrics {
	s.SceneName = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetStatisticsCycle(v int32) *ListABMetricsResponseBodyABMetrics {
	s.StatisticsCycle = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetTableMetaId(v string) *ListABMetricsResponseBodyABMetrics {
	s.TableMetaId = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetType(v string) *ListABMetricsResponseBodyABMetrics {
	s.Type = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) Validate() error {
	return dara.Validate(s)
}
