// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCondition interface {
	dara.Model
	String() string
	GoString() string
	SetComparisonOperator(v string) *TriggerCondition
	GetComparisonOperator() *string
	SetMetricName(v string) *TriggerCondition
	GetMetricName() *string
	SetStatistics(v string) *TriggerCondition
	GetStatistics() *string
	SetTags(v []*Tag) *TriggerCondition
	GetTags() []*Tag
	SetThreshold(v float64) *TriggerCondition
	GetThreshold() *float64
}

type TriggerCondition struct {
	// The comparison operator. This parameter is required. Valid values:
	//
	// 	- EQ: equal to
	//
	// 	- NE: not equal to
	//
	// 	- GT: greater than
	//
	// 	- LT: less than
	//
	// 	- GE: greater than or equal to
	//
	// 	- LE: less than or equal to
	//
	// This parameter is required.
	//
	// example:
	//
	// LT
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" xml:"ComparisonOperator,omitempty"`
	// The name of the metric. This parameter is required and cannot be an empty string. You can view the metric name in [Add Auto Scaling Rules](https://help.aliyun.com/document_detail/445658.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// yarn_resourcemanager_queue_PendingVCores
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The name of the statistic. This parameter is required. Valid values:
	//
	// 	- MAX
	//
	// 	- MIN
	//
	// 	- AVG
	//
	// This parameter is required.
	//
	// example:
	//
	// AVG
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The tags for the metrics of a partition. This parameter is available for only metrics that contain ByPartition. For other metrics, leave this parameter empty.
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The trigger threshold. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.5
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s TriggerCondition) String() string {
	return dara.Prettify(s)
}

func (s TriggerCondition) GoString() string {
	return s.String()
}

func (s *TriggerCondition) GetComparisonOperator() *string {
	return s.ComparisonOperator
}

func (s *TriggerCondition) GetMetricName() *string {
	return s.MetricName
}

func (s *TriggerCondition) GetStatistics() *string {
	return s.Statistics
}

func (s *TriggerCondition) GetTags() []*Tag {
	return s.Tags
}

func (s *TriggerCondition) GetThreshold() *float64 {
	return s.Threshold
}

func (s *TriggerCondition) SetComparisonOperator(v string) *TriggerCondition {
	s.ComparisonOperator = &v
	return s
}

func (s *TriggerCondition) SetMetricName(v string) *TriggerCondition {
	s.MetricName = &v
	return s
}

func (s *TriggerCondition) SetStatistics(v string) *TriggerCondition {
	s.Statistics = &v
	return s
}

func (s *TriggerCondition) SetTags(v []*Tag) *TriggerCondition {
	s.Tags = v
	return s
}

func (s *TriggerCondition) SetThreshold(v float64) *TriggerCondition {
	s.Threshold = &v
	return s
}

func (s *TriggerCondition) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
