// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportModelFeatureTrainingSetTableRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFeatureViewConfig(v map[string]*FeatureViewConfigValue) *ExportModelFeatureTrainingSetTableRequest
  GetFeatureViewConfig() map[string]*FeatureViewConfigValue 
  SetLabelInputConfig(v *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) *ExportModelFeatureTrainingSetTableRequest
  GetLabelInputConfig() *ExportModelFeatureTrainingSetTableRequestLabelInputConfig 
  SetRealTimeIterateInterval(v int64) *ExportModelFeatureTrainingSetTableRequest
  GetRealTimeIterateInterval() *int64 
  SetRealTimePartitionCountValue(v int64) *ExportModelFeatureTrainingSetTableRequest
  GetRealTimePartitionCountValue() *int64 
  SetTrainingSetConfig(v *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) *ExportModelFeatureTrainingSetTableRequest
  GetTrainingSetConfig() *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig 
}

type ExportModelFeatureTrainingSetTableRequest struct {
  FeatureViewConfig map[string]*FeatureViewConfigValue `json:"FeatureViewConfig,omitempty" xml:"FeatureViewConfig,omitempty"`
  LabelInputConfig *ExportModelFeatureTrainingSetTableRequestLabelInputConfig `json:"LabelInputConfig,omitempty" xml:"LabelInputConfig,omitempty" type:"Struct"`
  RealTimeIterateInterval *int64 `json:"RealTimeIterateInterval,omitempty" xml:"RealTimeIterateInterval,omitempty"`
  RealTimePartitionCountValue *int64 `json:"RealTimePartitionCountValue,omitempty" xml:"RealTimePartitionCountValue,omitempty"`
  TrainingSetConfig *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig `json:"TrainingSetConfig,omitempty" xml:"TrainingSetConfig,omitempty" type:"Struct"`
}

func (s ExportModelFeatureTrainingSetTableRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableRequest) GoString() string {
  return s.String()
}

func (s *ExportModelFeatureTrainingSetTableRequest) GetFeatureViewConfig() map[string]*FeatureViewConfigValue  {
  return s.FeatureViewConfig
}

func (s *ExportModelFeatureTrainingSetTableRequest) GetLabelInputConfig() *ExportModelFeatureTrainingSetTableRequestLabelInputConfig  {
  return s.LabelInputConfig
}

func (s *ExportModelFeatureTrainingSetTableRequest) GetRealTimeIterateInterval() *int64  {
  return s.RealTimeIterateInterval
}

func (s *ExportModelFeatureTrainingSetTableRequest) GetRealTimePartitionCountValue() *int64  {
  return s.RealTimePartitionCountValue
}

func (s *ExportModelFeatureTrainingSetTableRequest) GetTrainingSetConfig() *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig  {
  return s.TrainingSetConfig
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetFeatureViewConfig(v map[string]*FeatureViewConfigValue) *ExportModelFeatureTrainingSetTableRequest {
  s.FeatureViewConfig = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetLabelInputConfig(v *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) *ExportModelFeatureTrainingSetTableRequest {
  s.LabelInputConfig = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetRealTimeIterateInterval(v int64) *ExportModelFeatureTrainingSetTableRequest {
  s.RealTimeIterateInterval = &v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetRealTimePartitionCountValue(v int64) *ExportModelFeatureTrainingSetTableRequest {
  s.RealTimePartitionCountValue = &v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) SetTrainingSetConfig(v *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) *ExportModelFeatureTrainingSetTableRequest {
  s.TrainingSetConfig = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequest) Validate() error {
  return dara.Validate(s)
}

type ExportModelFeatureTrainingSetTableRequestLabelInputConfig struct {
  // example:
  // 
  // 2022-07-02 00:00:00
  EventTime *string `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
  Partitions map[string]map[string]interface{} `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableRequestLabelInputConfig) String() string {
  return dara.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableRequestLabelInputConfig) GoString() string {
  return s.String()
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) GetEventTime() *string  {
  return s.EventTime
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) GetPartitions() map[string]map[string]interface{}  {
  return s.Partitions
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) SetEventTime(v string) *ExportModelFeatureTrainingSetTableRequestLabelInputConfig {
  s.EventTime = &v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) SetPartitions(v map[string]map[string]interface{}) *ExportModelFeatureTrainingSetTableRequestLabelInputConfig {
  s.Partitions = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequestLabelInputConfig) Validate() error {
  return dara.Validate(s)
}

type ExportModelFeatureTrainingSetTableRequestTrainingSetConfig struct {
  Partitions map[string]map[string]interface{} `json:"Partitions,omitempty" xml:"Partitions,omitempty"`
}

func (s ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) String() string {
  return dara.Prettify(s)
}

func (s ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) GoString() string {
  return s.String()
}

func (s *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) GetPartitions() map[string]map[string]interface{}  {
  return s.Partitions
}

func (s *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) SetPartitions(v map[string]map[string]interface{}) *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig {
  s.Partitions = v
  return s
}

func (s *ExportModelFeatureTrainingSetTableRequestTrainingSetConfig) Validate() error {
  return dara.Validate(s)
}

