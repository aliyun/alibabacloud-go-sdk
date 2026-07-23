// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportABMetricGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseExperimentId(v string) *ReportABMetricGroupRequest
	GetBaseExperimentId() *string
	SetDimensionFields(v string) *ReportABMetricGroupRequest
	GetDimensionFields() *string
	SetEndDate(v string) *ReportABMetricGroupRequest
	GetEndDate() *string
	SetExperimentGroupId(v string) *ReportABMetricGroupRequest
	GetExperimentGroupId() *string
	SetExperimentIds(v string) *ReportABMetricGroupRequest
	GetExperimentIds() *string
	SetInstanceId(v string) *ReportABMetricGroupRequest
	GetInstanceId() *string
	SetReportType(v string) *ReportABMetricGroupRequest
	GetReportType() *string
	SetSceneId(v string) *ReportABMetricGroupRequest
	GetSceneId() *string
	SetStartDate(v string) *ReportABMetricGroupRequest
	GetStartDate() *string
	SetTimeStatisticsMethod(v string) *ReportABMetricGroupRequest
	GetTimeStatisticsMethod() *string
}

type ReportABMetricGroupRequest struct {
	// The ID of the base experiment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	BaseExperimentId *string `json:"BaseExperimentId,omitempty" xml:"BaseExperimentId,omitempty"`
	// The dimension fields.
	//
	// example:
	//
	// {"gender":"man"}
	DimensionFields *string `json:"DimensionFields,omitempty" xml:"DimensionFields,omitempty"`
	// The end date.
	//
	// example:
	//
	// 2021-07-01
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The ID of the experiment group. You can call the [ListExperimentGroups](https://help.aliyun.com/document_detail/2402573.html) operation to obtain the ID of the experiment group. This parameter is required for offline and real-time reports.
	//
	// example:
	//
	// 3
	ExperimentGroupId *string `json:"ExperimentGroupId,omitempty" xml:"ExperimentGroupId,omitempty"`
	// The IDs of the experiments. You can call the [ListExperiments](https://help.aliyun.com/document_detail/2402582.html) operation to obtain the IDs of the experiments.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3,4,5
	ExperimentIds *string `json:"ExperimentIds,omitempty" xml:"ExperimentIds,omitempty"`
	// The instance ID. You can call the [ListInstances](https://help.aliyun.com/document_detail/2411819.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the report.
	//
	// - Realtime: real-time report.
	//
	// - Offline: offline report.
	//
	// - Market: dashboard report.
	//
	// This parameter is required.
	//
	// example:
	//
	// Offline
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	// The ID of the scene. You can call the [ListScenes](https://help.aliyun.com/document_detail/2402581.html) operation to obtain the scene ID.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The start date.
	//
	// example:
	//
	// 2021-07-01
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The time statistics method. This parameter is required for real-time reports.
	//
	// - Hour: The statistics are collected by hour.
	//
	// - Day: The statistics are accumulated from 00:00 on the current day.
	//
	// example:
	//
	// Hour
	TimeStatisticsMethod *string `json:"TimeStatisticsMethod,omitempty" xml:"TimeStatisticsMethod,omitempty"`
}

func (s ReportABMetricGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *ReportABMetricGroupRequest) GetBaseExperimentId() *string {
	return s.BaseExperimentId
}

func (s *ReportABMetricGroupRequest) GetDimensionFields() *string {
	return s.DimensionFields
}

func (s *ReportABMetricGroupRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ReportABMetricGroupRequest) GetExperimentGroupId() *string {
	return s.ExperimentGroupId
}

func (s *ReportABMetricGroupRequest) GetExperimentIds() *string {
	return s.ExperimentIds
}

func (s *ReportABMetricGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReportABMetricGroupRequest) GetReportType() *string {
	return s.ReportType
}

func (s *ReportABMetricGroupRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ReportABMetricGroupRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ReportABMetricGroupRequest) GetTimeStatisticsMethod() *string {
	return s.TimeStatisticsMethod
}

func (s *ReportABMetricGroupRequest) SetBaseExperimentId(v string) *ReportABMetricGroupRequest {
	s.BaseExperimentId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetDimensionFields(v string) *ReportABMetricGroupRequest {
	s.DimensionFields = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetEndDate(v string) *ReportABMetricGroupRequest {
	s.EndDate = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetExperimentGroupId(v string) *ReportABMetricGroupRequest {
	s.ExperimentGroupId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetExperimentIds(v string) *ReportABMetricGroupRequest {
	s.ExperimentIds = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetInstanceId(v string) *ReportABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetReportType(v string) *ReportABMetricGroupRequest {
	s.ReportType = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetSceneId(v string) *ReportABMetricGroupRequest {
	s.SceneId = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetStartDate(v string) *ReportABMetricGroupRequest {
	s.StartDate = &v
	return s
}

func (s *ReportABMetricGroupRequest) SetTimeStatisticsMethod(v string) *ReportABMetricGroupRequest {
	s.TimeStatisticsMethod = &v
	return s
}

func (s *ReportABMetricGroupRequest) Validate() error {
	return dara.Validate(s)
}
