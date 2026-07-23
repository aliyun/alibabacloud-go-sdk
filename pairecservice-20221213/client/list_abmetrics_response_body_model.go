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
	// The list of AB metrics.
	ABMetrics []*ListABMetricsResponseBodyABMetrics `json:"ABMetrics,omitempty" xml:"ABMetrics,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of AB metrics returned.
	//
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
	if s.ABMetrics != nil {
		for _, item := range s.ABMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListABMetricsResponseBodyABMetrics struct {
	// The AB metric ID.
	//
	// example:
	//
	// 1
	ABMetricId *string `json:"ABMetricId,omitempty" xml:"ABMetricId,omitempty"`
	// Indicates whether to aggregate data by user.
	AggregationByUser *bool `json:"AggregationByUser,omitempty" xml:"AggregationByUser,omitempty"`
	// The metric definition.
	//
	// example:
	//
	// sum(click_cnt)
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The denominator of the derived metric.
	//
	// example:
	//
	// pv
	Denominator *string `json:"Denominator,omitempty" xml:"Denominator,omitempty"`
	// The metric description.
	//
	// example:
	//
	// pv指标
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the metric follows a binomial distribution.
	IsBinomialDistribution *bool `json:"IsBinomialDistribution,omitempty" xml:"IsBinomialDistribution,omitempty"`
	// The ID of the left-hand metric in the formula for the derived metric.
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
	// Indicates whether significance calculation is required.
	NeedSignificance *bool `json:"NeedSignificance,omitempty" xml:"NeedSignificance,omitempty"`
	// The numerator of the derived metric.
	//
	// example:
	//
	// click
	Numerator *string `json:"Numerator,omitempty" xml:"Numerator,omitempty"`
	// The operator used to calculate the derived metric. Valid values:
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
	// Indicates whether the metric is a real-time metric.
	//
	// - `true`: The metric is calculated in real time.
	//
	// - `false`: The metric is not calculated in real time.
	//
	// example:
	//
	// false
	Realtime *string `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// The ID of the data source for the data table to which the results are written.
	//
	// example:
	//
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// The ID of the data table to which the results are written.
	//
	// example:
	//
	// 2
	ResultTableMetaId *string `json:"ResultTableMetaId,omitempty" xml:"ResultTableMetaId,omitempty"`
	// The ID of the right-hand metric in the formula for the derived metric.
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
	// The statistical period.
	//
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// The data table ID.
	//
	// example:
	//
	// 1
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// The metric type. Valid values:
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

func (s ListABMetricsResponseBodyABMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListABMetricsResponseBodyABMetrics) GoString() string {
	return s.String()
}

func (s *ListABMetricsResponseBodyABMetrics) GetABMetricId() *string {
	return s.ABMetricId
}

func (s *ListABMetricsResponseBodyABMetrics) GetAggregationByUser() *bool {
	return s.AggregationByUser
}

func (s *ListABMetricsResponseBodyABMetrics) GetDefinition() *string {
	return s.Definition
}

func (s *ListABMetricsResponseBodyABMetrics) GetDenominator() *string {
	return s.Denominator
}

func (s *ListABMetricsResponseBodyABMetrics) GetDescription() *string {
	return s.Description
}

func (s *ListABMetricsResponseBodyABMetrics) GetIsBinomialDistribution() *bool {
	return s.IsBinomialDistribution
}

func (s *ListABMetricsResponseBodyABMetrics) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *ListABMetricsResponseBodyABMetrics) GetName() *string {
	return s.Name
}

func (s *ListABMetricsResponseBodyABMetrics) GetNeedSignificance() *bool {
	return s.NeedSignificance
}

func (s *ListABMetricsResponseBodyABMetrics) GetNumerator() *string {
	return s.Numerator
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

func (s *ListABMetricsResponseBodyABMetrics) SetAggregationByUser(v bool) *ListABMetricsResponseBodyABMetrics {
	s.AggregationByUser = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDefinition(v string) *ListABMetricsResponseBodyABMetrics {
	s.Definition = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDenominator(v string) *ListABMetricsResponseBodyABMetrics {
	s.Denominator = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetDescription(v string) *ListABMetricsResponseBodyABMetrics {
	s.Description = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetIsBinomialDistribution(v bool) *ListABMetricsResponseBodyABMetrics {
	s.IsBinomialDistribution = &v
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

func (s *ListABMetricsResponseBodyABMetrics) SetNeedSignificance(v bool) *ListABMetricsResponseBodyABMetrics {
	s.NeedSignificance = &v
	return s
}

func (s *ListABMetricsResponseBodyABMetrics) SetNumerator(v string) *ListABMetricsResponseBodyABMetrics {
	s.Numerator = &v
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
