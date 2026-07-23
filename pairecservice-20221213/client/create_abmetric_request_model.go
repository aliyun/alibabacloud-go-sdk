// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregationByUser(v bool) *CreateABMetricRequest
	GetAggregationByUser() *bool
	SetDefinition(v string) *CreateABMetricRequest
	GetDefinition() *string
	SetDenominator(v string) *CreateABMetricRequest
	GetDenominator() *string
	SetDescription(v string) *CreateABMetricRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateABMetricRequest
	GetInstanceId() *string
	SetIsBinomialDistribution(v bool) *CreateABMetricRequest
	GetIsBinomialDistribution() *bool
	SetLeftMetricId(v string) *CreateABMetricRequest
	GetLeftMetricId() *string
	SetName(v string) *CreateABMetricRequest
	GetName() *string
	SetNeedSignificance(v bool) *CreateABMetricRequest
	GetNeedSignificance() *bool
	SetNumerator(v string) *CreateABMetricRequest
	GetNumerator() *string
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
	AggregationByUser *bool `json:"AggregationByUser,omitempty" xml:"AggregationByUser,omitempty"`
	// The metric definition.
	//
	// This parameter is required.
	//
	// example:
	//
	// sum(click_cnt)
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// example:
	//
	// pv
	Denominator *string `json:"Denominator,omitempty" xml:"Denominator,omitempty"`
	// The metric description.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv指标
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID. Call the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsBinomialDistribution *bool   `json:"IsBinomialDistribution,omitempty" xml:"IsBinomialDistribution,omitempty"`
	// The ID of the left metric used to calculate the derived metric.
	//
	// example:
	//
	// 2
	LeftMetricId *string `json:"LeftMetricId,omitempty" xml:"LeftMetricId,omitempty"`
	// The metric name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pv
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedSignificance *bool   `json:"NeedSignificance,omitempty" xml:"NeedSignificance,omitempty"`
	// example:
	//
	// click
	Numerator *string `json:"Numerator,omitempty" xml:"Numerator,omitempty"`
	// The operator used to calculate the derived metric. Valid values:
	//
	// - `Plus`: Addition
	//
	// - `Minus`: Subtraction
	//
	// - `Multiplication`: Multiplication
	//
	// - `Division`: Division
	//
	// example:
	//
	// Division
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// Specifies whether the metric is a real-time metric. Valid values:
	//
	// - `true`: The metric is real-time.
	//
	// - `false`: The metric is not real-time.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// The data source ID to which the results are written back. Call the ListInstanceResources operation to obtain this ID.
	//
	// example:
	//
	// 3
	ResultResourceId *string `json:"ResultResourceId,omitempty" xml:"ResultResourceId,omitempty"`
	// The ID of the right metric used to calculate the derived metric.
	//
	// example:
	//
	// 3
	RightMetricId *string `json:"RightMetricId,omitempty" xml:"RightMetricId,omitempty"`
	// The scene ID. Call the [ListScenes](https://help.aliyun.com/document_detail/2402581.html) operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The statistics cycle.
	//
	// example:
	//
	// 1
	StatisticsCycle *int32 `json:"StatisticsCycle,omitempty" xml:"StatisticsCycle,omitempty"`
	// The data table ID. Call the ListTableMetas operation to obtain this ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	TableMetaId *string `json:"TableMetaId,omitempty" xml:"TableMetaId,omitempty"`
	// The metric type. Valid values:
	//
	// - `Single`: A single metric.
	//
	// - `Derived`: A derived metric.
	//
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

func (s *CreateABMetricRequest) GetAggregationByUser() *bool {
	return s.AggregationByUser
}

func (s *CreateABMetricRequest) GetDefinition() *string {
	return s.Definition
}

func (s *CreateABMetricRequest) GetDenominator() *string {
	return s.Denominator
}

func (s *CreateABMetricRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateABMetricRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateABMetricRequest) GetIsBinomialDistribution() *bool {
	return s.IsBinomialDistribution
}

func (s *CreateABMetricRequest) GetLeftMetricId() *string {
	return s.LeftMetricId
}

func (s *CreateABMetricRequest) GetName() *string {
	return s.Name
}

func (s *CreateABMetricRequest) GetNeedSignificance() *bool {
	return s.NeedSignificance
}

func (s *CreateABMetricRequest) GetNumerator() *string {
	return s.Numerator
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

func (s *CreateABMetricRequest) SetAggregationByUser(v bool) *CreateABMetricRequest {
	s.AggregationByUser = &v
	return s
}

func (s *CreateABMetricRequest) SetDefinition(v string) *CreateABMetricRequest {
	s.Definition = &v
	return s
}

func (s *CreateABMetricRequest) SetDenominator(v string) *CreateABMetricRequest {
	s.Denominator = &v
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

func (s *CreateABMetricRequest) SetIsBinomialDistribution(v bool) *CreateABMetricRequest {
	s.IsBinomialDistribution = &v
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

func (s *CreateABMetricRequest) SetNeedSignificance(v bool) *CreateABMetricRequest {
	s.NeedSignificance = &v
	return s
}

func (s *CreateABMetricRequest) SetNumerator(v string) *CreateABMetricRequest {
	s.Numerator = &v
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
