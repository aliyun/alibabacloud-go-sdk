// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v []*ListOfflineTaskResponseBodyResult) *ListOfflineTaskResponseBody
	GetResult() []*ListOfflineTaskResponseBodyResult
	SetTotalCount(v int64) *ListOfflineTaskResponseBody
	GetTotalCount() *int64
}

type ListOfflineTaskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 786CC01F-0F1D-5FB5-8BFF-B0F3DB289772
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result []*ListOfflineTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 7
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOfflineTaskResponseBody) GetResult() []*ListOfflineTaskResponseBodyResult {
	return s.Result
}

func (s *ListOfflineTaskResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListOfflineTaskResponseBody) SetRequestId(v string) *ListOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOfflineTaskResponseBody) SetResult(v []*ListOfflineTaskResponseBodyResult) *ListOfflineTaskResponseBody {
	s.Result = v
	return s
}

func (s *ListOfflineTaskResponseBody) SetTotalCount(v int64) *ListOfflineTaskResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOfflineTaskResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOfflineTaskResponseBodyResult struct {
	// The task metadata.
	Meta *ListOfflineTaskResponseBodyResultMeta `json:"Meta,omitempty" xml:"Meta,omitempty" type:"Struct"`
	// The processing pipeline operators.
	Processors []*ListOfflineTaskResponseBodyResultProcessors `json:"Processors,omitempty" xml:"Processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*ListOfflineTaskResponseBodyResultSink `json:"Sink,omitempty" xml:"Sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*ListOfflineTaskResponseBodyResultSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Repeated"`
	// The task status.
	Status *ListOfflineTaskResponseBodyResultStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s ListOfflineTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBodyResult) GetMeta() *ListOfflineTaskResponseBodyResultMeta {
	return s.Meta
}

func (s *ListOfflineTaskResponseBodyResult) GetProcessors() []*ListOfflineTaskResponseBodyResultProcessors {
	return s.Processors
}

func (s *ListOfflineTaskResponseBodyResult) GetSink() []*ListOfflineTaskResponseBodyResultSink {
	return s.Sink
}

func (s *ListOfflineTaskResponseBodyResult) GetSource() []*ListOfflineTaskResponseBodyResultSource {
	return s.Source
}

func (s *ListOfflineTaskResponseBodyResult) GetStatus() *ListOfflineTaskResponseBodyResultStatus {
	return s.Status
}

func (s *ListOfflineTaskResponseBodyResult) SetMeta(v *ListOfflineTaskResponseBodyResultMeta) *ListOfflineTaskResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListOfflineTaskResponseBodyResult) SetProcessors(v []*ListOfflineTaskResponseBodyResultProcessors) *ListOfflineTaskResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ListOfflineTaskResponseBodyResult) SetSink(v []*ListOfflineTaskResponseBodyResultSink) *ListOfflineTaskResponseBodyResult {
	s.Sink = v
	return s
}

func (s *ListOfflineTaskResponseBodyResult) SetSource(v []*ListOfflineTaskResponseBodyResultSource) *ListOfflineTaskResponseBodyResult {
	s.Source = v
	return s
}

func (s *ListOfflineTaskResponseBodyResult) SetStatus(v *ListOfflineTaskResponseBodyResultStatus) *ListOfflineTaskResponseBodyResult {
	s.Status = v
	return s
}

func (s *ListOfflineTaskResponseBodyResult) Validate() error {
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	if s.Processors != nil {
		for _, item := range s.Processors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Sink != nil {
		for _, item := range s.Sink {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Source != nil {
		for _, item := range s.Source {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOfflineTaskResponseBodyResultMeta struct {
	// The list of labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The region ID of the task.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task name.
	//
	// example:
	//
	// jly-fesOffline-172.16.8.133-20912
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 2192861158
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListOfflineTaskResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBodyResultMeta) GetLabels() []*string {
	return s.Labels
}

func (s *ListOfflineTaskResponseBodyResultMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOfflineTaskResponseBodyResultMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *ListOfflineTaskResponseBodyResultMeta) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListOfflineTaskResponseBodyResultMeta) SetLabels(v []*string) *ListOfflineTaskResponseBodyResultMeta {
	s.Labels = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultMeta) SetRegionId(v string) *ListOfflineTaskResponseBodyResultMeta {
	s.RegionId = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultMeta) SetTaskName(v string) *ListOfflineTaskResponseBodyResultMeta {
	s.TaskName = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultMeta) SetWorkspaceId(v string) *ListOfflineTaskResponseBodyResultMeta {
	s.WorkspaceId = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}

type ListOfflineTaskResponseBodyResultProcessors struct {
	// The input parameters.
	Input map[string]*string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name.
	//
	// example:
	//
	// processor1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output parameters.
	Output map[string]*string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The processor parameters.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The type.
	//
	// example:
	//
	// document-analyze
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOfflineTaskResponseBodyResultProcessors) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBodyResultProcessors) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBodyResultProcessors) GetInput() map[string]*string {
	return s.Input
}

func (s *ListOfflineTaskResponseBodyResultProcessors) GetName() *string {
	return s.Name
}

func (s *ListOfflineTaskResponseBodyResultProcessors) GetOutput() map[string]*string {
	return s.Output
}

func (s *ListOfflineTaskResponseBodyResultProcessors) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ListOfflineTaskResponseBodyResultProcessors) GetType() *string {
	return s.Type
}

func (s *ListOfflineTaskResponseBodyResultProcessors) SetInput(v map[string]*string) *ListOfflineTaskResponseBodyResultProcessors {
	s.Input = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultProcessors) SetName(v string) *ListOfflineTaskResponseBodyResultProcessors {
	s.Name = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultProcessors) SetOutput(v map[string]*string) *ListOfflineTaskResponseBodyResultProcessors {
	s.Output = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultProcessors) SetParameters(v map[string]*string) *ListOfflineTaskResponseBodyResultProcessors {
	s.Parameters = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultProcessors) SetType(v string) *ListOfflineTaskResponseBodyResultProcessors {
	s.Type = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultProcessors) Validate() error {
	return dara.Validate(s)
}

type ListOfflineTaskResponseBodyResultSink struct {
	// The data sink name.
	//
	// example:
	//
	// table2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data sink configuration parameters, which are determined by the type.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The primary key field of the data sink.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The data sink schema.
	Schema []map[string]*string `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
	// The data sink type.
	//
	// example:
	//
	// swift
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOfflineTaskResponseBodyResultSink) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBodyResultSink) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBodyResultSink) GetName() *string {
	return s.Name
}

func (s *ListOfflineTaskResponseBodyResultSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ListOfflineTaskResponseBodyResultSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ListOfflineTaskResponseBodyResultSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *ListOfflineTaskResponseBodyResultSink) GetType() *string {
	return s.Type
}

func (s *ListOfflineTaskResponseBodyResultSink) SetName(v string) *ListOfflineTaskResponseBodyResultSink {
	s.Name = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSink) SetParameters(v map[string]*string) *ListOfflineTaskResponseBodyResultSink {
	s.Parameters = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSink) SetPrimaryKey(v string) *ListOfflineTaskResponseBodyResultSink {
	s.PrimaryKey = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSink) SetSchema(v []map[string]*string) *ListOfflineTaskResponseBodyResultSink {
	s.Schema = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSink) SetType(v string) *ListOfflineTaskResponseBodyResultSink {
	s.Type = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSink) Validate() error {
	return dara.Validate(s)
}

type ListOfflineTaskResponseBodyResultSource struct {
	// The data source name.
	//
	// example:
	//
	// table
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data source configuration parameters.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The primary key field of the data source.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The data source schema.
	Schema []map[string]*string `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
	// The data source type.
	//
	// example:
	//
	// rds
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOfflineTaskResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBodyResultSource) GetName() *string {
	return s.Name
}

func (s *ListOfflineTaskResponseBodyResultSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ListOfflineTaskResponseBodyResultSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ListOfflineTaskResponseBodyResultSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *ListOfflineTaskResponseBodyResultSource) GetType() *string {
	return s.Type
}

func (s *ListOfflineTaskResponseBodyResultSource) SetName(v string) *ListOfflineTaskResponseBodyResultSource {
	s.Name = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSource) SetParameters(v map[string]*string) *ListOfflineTaskResponseBodyResultSource {
	s.Parameters = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSource) SetPrimaryKey(v string) *ListOfflineTaskResponseBodyResultSource {
	s.PrimaryKey = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSource) SetSchema(v []map[string]*string) *ListOfflineTaskResponseBodyResultSource {
	s.Schema = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSource) SetType(v string) *ListOfflineTaskResponseBodyResultSource {
	s.Type = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}

type ListOfflineTaskResponseBodyResultStatus struct {
	// The time when the task was created.
	//
	// example:
	//
	// 1744941600000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The monitoring information.
	MetricData map[string]*string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	// The task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the task was last modified.
	//
	// example:
	//
	// 1744941600000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListOfflineTaskResponseBodyResultStatus) String() string {
	return dara.Prettify(s)
}

func (s ListOfflineTaskResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *ListOfflineTaskResponseBodyResultStatus) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListOfflineTaskResponseBodyResultStatus) GetMetricData() map[string]*string {
	return s.MetricData
}

func (s *ListOfflineTaskResponseBodyResultStatus) GetStatus() *string {
	return s.Status
}

func (s *ListOfflineTaskResponseBodyResultStatus) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListOfflineTaskResponseBodyResultStatus) SetCreateTime(v int64) *ListOfflineTaskResponseBodyResultStatus {
	s.CreateTime = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultStatus) SetMetricData(v map[string]*string) *ListOfflineTaskResponseBodyResultStatus {
	s.MetricData = v
	return s
}

func (s *ListOfflineTaskResponseBodyResultStatus) SetStatus(v string) *ListOfflineTaskResponseBodyResultStatus {
	s.Status = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultStatus) SetUpdateTime(v int64) *ListOfflineTaskResponseBodyResultStatus {
	s.UpdateTime = &v
	return s
}

func (s *ListOfflineTaskResponseBodyResultStatus) Validate() error {
	return dara.Validate(s)
}
