// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationByUser(v bool) *GetABMetricResponseBody
	GetAggregationByUser() *bool
	SetDefinition(v string) *GetABMetricResponseBody
	GetDefinition() *string
	SetDenominator(v string) *GetABMetricResponseBody
	GetDenominator() *string
	SetDescription(v string) *GetABMetricResponseBody
	GetDescription() *string
	SetIsBinomialDistribution(v bool) *GetABMetricResponseBody
	GetIsBinomialDistribution() *bool
	SetLeftMetricId(v string) *GetABMetricResponseBody
	GetLeftMetricId() *string
	SetName(v string) *GetABMetricResponseBody
	GetName() *string
	SetNeedSignificance(v bool) *GetABMetricResponseBody
	GetNeedSignificance() *bool
	SetNumerator(v string) *GetABMetricResponseBody
	GetNumerator() *string
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
	// Specifies whether to aggregate metrics by user.
	AggregationByUser *bool `json:"AggregationByUser,omitempty" xml:"AggregationByUser,omitempty"`
	// The metric definition.
	//
	// example:
	//
	// sum(click_cnt)
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The denominator of the derived metric formula.
	//
	// example:
	//
	// pv
	Denominator *string `json:"Denominator,omitempty" xml:"Denominator,omitempty"`
	// The metric description.
	//
	// example:
	//
	// 页面访问次数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the metric follows a binomial distribution. This affects how significance is calculated.
	IsBinomialDistribution *bool `json:"IsBinomialDistribution,omitempty" xml:"IsBinomialDistribution,omitempty"`
	// The ID of the left operand metric for a derived metric.
	//
	// example:
	//
	// 3
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// The metric name.
	//
	// example:
	//
	// pv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether significance testing is required for this metric.
	NeedSignificance *bool `json:"NeedSignificance,omitempty" xml:"NeedSignificance,omitempty"`
	// The numerator of the derived metric formula.
	//
	// example:
	//
	// click
	Numerator *string `json:"Numerator,omitempty" xml:"Numerator,omitempty"`
	// The operator for the derived metric. Valid values:
	//
	// - `Plus`: addition
	//
	// - `Minus`: subtraction
	//
	// - `Multiplication`: multiplication
	//
	// - `Division`: division
	//
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Specifies whether the metric is real-time. Valid values:
	//
	// - `true`: The metric is real-time.
	//
	// - `false`: The metric is not real-time.
	//
	// example:
	//
	// false
	Realtime *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the data source for the results table.
	//
	// example:
	//
	// 5
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// The ID of the results table.
	//
	// example:
	//
	// 3
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	// The ID of the right operand metric for a derived metric.
	//
	// example:
	//
	// 2
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// home_feed
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	// The statistics cycle.
	//
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// The ID of the source table.
	//
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// The type of the metric. Valid values:
	//
	// - `Single`: a single metric.
	//
	// - `Derived`: a derived metric.
	//
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

func (s *GetABMetricResponseBody) GetAggregationByUser() *bool {
	return s.AggregationByUser
}

func (s *GetABMetricResponseBody) GetDefinition() *string {
	return s.Definition
}

func (s *GetABMetricResponseBody) GetDenominator() *string {
	return s.Denominator
}

func (s *GetABMetricResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetABMetricResponseBody) GetIsBinomialDistribution() *bool {
	return s.IsBinomialDistribution
}

func (s *GetABMetricResponseBody) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *GetABMetricResponseBody) GetName() *string {
	return s.Name
}

func (s *GetABMetricResponseBody) GetNeedSignificance() *bool {
	return s.NeedSignificance
}

func (s *GetABMetricResponseBody) GetNumerator() *string {
	return s.Numerator
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

func (s *GetABMetricResponseBody) SetAggregationByUser(v bool) *GetABMetricResponseBody {
	s.AggregationByUser = &v
	return s
}

func (s *GetABMetricResponseBody) SetDefinition(v string) *GetABMetricResponseBody {
	s.Definition = &v
	return s
}

func (s *GetABMetricResponseBody) SetDenominator(v string) *GetABMetricResponseBody {
	s.Denominator = &v
	return s
}

func (s *GetABMetricResponseBody) SetDescription(v string) *GetABMetricResponseBody {
	s.Description = &v
	return s
}

func (s *GetABMetricResponseBody) SetIsBinomialDistribution(v bool) *GetABMetricResponseBody {
	s.IsBinomialDistribution = &v
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

func (s *GetABMetricResponseBody) SetNeedSignificance(v bool) *GetABMetricResponseBody {
	s.NeedSignificance = &v
	return s
}

func (s *GetABMetricResponseBody) SetNumerator(v string) *GetABMetricResponseBody {
	s.Numerator = &v
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
