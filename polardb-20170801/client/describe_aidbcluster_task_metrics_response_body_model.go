// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTaskMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAIDBClusterTaskMetricsResponseBody
	GetDBClusterId() *string
	SetEndTime(v string) *DescribeAIDBClusterTaskMetricsResponseBody
	GetEndTime() *string
	SetItems(v *DescribeAIDBClusterTaskMetricsResponseBodyItems) *DescribeAIDBClusterTaskMetricsResponseBody
	GetItems() *DescribeAIDBClusterTaskMetricsResponseBodyItems
	SetMetricType(v string) *DescribeAIDBClusterTaskMetricsResponseBody
	GetMetricType() *string
	SetPageNumber(v int64) *DescribeAIDBClusterTaskMetricsResponseBody
	GetPageNumber() *int64
	SetPageRecordCount(v int32) *DescribeAIDBClusterTaskMetricsResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int64) *DescribeAIDBClusterTaskMetricsResponseBody
	GetPageSize() *int64
	SetRelativeDbClusterId(v string) *DescribeAIDBClusterTaskMetricsResponseBody
	GetRelativeDbClusterId() *string
	SetRequestId(v string) *DescribeAIDBClusterTaskMetricsResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeAIDBClusterTaskMetricsResponseBody
	GetStartTime() *string
}

type DescribeAIDBClusterTaskMetricsResponseBody struct {
	// example:
	//
	// pm-2zejpr***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 2026-01-15T15:00:00Z
	EndTime *string                                          `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Items   *DescribeAIDBClusterTaskMetricsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// all
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 5
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pc-2zejpr***
	RelativeDbClusterId *string `json:"RelativeDbClusterId,omitempty" xml:"RelativeDbClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5E71541A-6007-4DCC-A38A-F872C31FEB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2026-01-15T14:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAIDBClusterTaskMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetItems() *DescribeAIDBClusterTaskMetricsResponseBodyItems {
	return s.Items
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetRelativeDbClusterId() *string {
	return s.RelativeDbClusterId
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetDBClusterId(v string) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetEndTime(v string) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetItems(v *DescribeAIDBClusterTaskMetricsResponseBodyItems) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetMetricType(v string) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.MetricType = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetPageNumber(v int64) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetPageRecordCount(v int32) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetPageSize(v int64) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetRelativeDbClusterId(v string) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.RelativeDbClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetRequestId(v string) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) SetStartTime(v string) *DescribeAIDBClusterTaskMetricsResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAIDBClusterTaskMetricsResponseBodyItems struct {
	SlsMetricsItems []*DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems `json:"SlsMetricsItems,omitempty" xml:"SlsMetricsItems,omitempty" type:"Repeated"`
}

func (s DescribeAIDBClusterTaskMetricsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskMetricsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItems) GetSlsMetricsItems() []*DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	return s.SlsMetricsItems
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItems) SetSlsMetricsItems(v []*DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) *DescribeAIDBClusterTaskMetricsResponseBodyItems {
	s.SlsMetricsItems = v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItems) Validate() error {
	if s.SlsMetricsItems != nil {
		for _, item := range s.SlsMetricsItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems struct {
	// example:
	//
	// 21
	CurrentStep *int32 `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// example:
	//
	// 0.23
	Epoch *float64 `json:"Epoch,omitempty" xml:"Epoch,omitempty"`
	// example:
	//
	// 90
	GlobalStep *int32 `json:"GlobalStep,omitempty" xml:"GlobalStep,omitempty"`
	// example:
	//
	// 2026-01-15T14:16:49.52140317Z
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// example:
	//
	// {"memory(GiB)":"xxx"}
	Metric map[string]interface{} `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// train
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// 1742090703
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetCurrentStep() *int32 {
	return s.CurrentStep
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetEpoch() *float64 {
	return s.Epoch
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetGlobalStep() *int32 {
	return s.GlobalStep
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetLogTime() *string {
	return s.LogTime
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetMetric() map[string]interface{} {
	return s.Metric
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetCurrentStep(v int32) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.CurrentStep = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetEpoch(v float64) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.Epoch = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetGlobalStep(v int32) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.GlobalStep = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetLogTime(v string) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.LogTime = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetMetric(v map[string]interface{}) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.Metric = v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetMetricType(v string) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.MetricType = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) SetTimestamp(v string) *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems {
	s.Timestamp = &v
	return s
}

func (s *DescribeAIDBClusterTaskMetricsResponseBodyItemsSlsMetricsItems) Validate() error {
	return dara.Validate(s)
}
