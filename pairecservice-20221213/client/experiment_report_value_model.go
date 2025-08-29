// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperimentReportValue interface {
  dara.Model
  String() string
  GoString() string
  SetBaseline(v bool) *ExperimentReportValue
  GetBaseline() *bool 
  SetMetricResults(v map[string]map[string]interface{}) *ExperimentReportValue
  GetMetricResults() map[string]map[string]interface{} 
}

type ExperimentReportValue struct {
  // example:
  // 
  // true
  Baseline *bool `json:"Baseline,omitempty" xml:"Baseline,omitempty"`
  MetricResults map[string]map[string]interface{} `json:"MetricResults,omitempty" xml:"MetricResults,omitempty"`
}

func (s ExperimentReportValue) String() string {
  return dara.Prettify(s)
}

func (s ExperimentReportValue) GoString() string {
  return s.String()
}

func (s *ExperimentReportValue) GetBaseline() *bool  {
  return s.Baseline
}

func (s *ExperimentReportValue) GetMetricResults() map[string]map[string]interface{}  {
  return s.MetricResults
}

func (s *ExperimentReportValue) SetBaseline(v bool) *ExperimentReportValue {
  s.Baseline = &v
  return s
}

func (s *ExperimentReportValue) SetMetricResults(v map[string]map[string]interface{}) *ExperimentReportValue {
  s.MetricResults = v
  return s
}

func (s *ExperimentReportValue) Validate() error {
  return dara.Validate(s)
}

