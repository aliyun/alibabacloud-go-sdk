// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v *StopOfflineTaskResponseBodyResult) *StopOfflineTaskResponseBody
	GetResult() *StopOfflineTaskResponseBodyResult
}

type StopOfflineTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5D56E988-A189-53A4-A0A6-C8D744B59775
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *StopOfflineTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s StopOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopOfflineTaskResponseBody) GetResult() *StopOfflineTaskResponseBodyResult {
	return s.Result
}

func (s *StopOfflineTaskResponseBody) SetRequestId(v string) *StopOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopOfflineTaskResponseBody) SetResult(v *StopOfflineTaskResponseBodyResult) *StopOfflineTaskResponseBody {
	s.Result = v
	return s
}

func (s *StopOfflineTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopOfflineTaskResponseBodyResult struct {
	// The metadata.
	Meta *StopOfflineTaskResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The node processing parameters.
	//
	// example:
	//
	// {
	//
	// "parameter1": {
	//
	// "key": "value"
	//
	// },
	//
	// "parameter2": {
	//
	// "key": "value"
	//
	// }
	//
	// }
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The processing flow operators.
	Processors []*StopOfflineTaskResponseBodyResultProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*StopOfflineTaskResponseBodyResultSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*StopOfflineTaskResponseBodyResultSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The node status.
	Status *StopOfflineTaskResponseBodyResultStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s StopOfflineTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBodyResult) GetMeta() *StopOfflineTaskResponseBodyResultMeta {
	return s.Meta
}

func (s *StopOfflineTaskResponseBodyResult) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *StopOfflineTaskResponseBodyResult) GetProcessors() []*StopOfflineTaskResponseBodyResultProcessors {
	return s.Processors
}

func (s *StopOfflineTaskResponseBodyResult) GetSink() []*StopOfflineTaskResponseBodyResultSink {
	return s.Sink
}

func (s *StopOfflineTaskResponseBodyResult) GetSource() []*StopOfflineTaskResponseBodyResultSource {
	return s.Source
}

func (s *StopOfflineTaskResponseBodyResult) GetStatus() *StopOfflineTaskResponseBodyResultStatus {
	return s.Status
}

func (s *StopOfflineTaskResponseBodyResult) SetMeta(v *StopOfflineTaskResponseBodyResultMeta) *StopOfflineTaskResponseBodyResult {
	s.Meta = v
	return s
}

func (s *StopOfflineTaskResponseBodyResult) SetParameters(v map[string]interface{}) *StopOfflineTaskResponseBodyResult {
	s.Parameters = v
	return s
}

func (s *StopOfflineTaskResponseBodyResult) SetProcessors(v []*StopOfflineTaskResponseBodyResultProcessors) *StopOfflineTaskResponseBodyResult {
	s.Processors = v
	return s
}

func (s *StopOfflineTaskResponseBodyResult) SetSink(v []*StopOfflineTaskResponseBodyResultSink) *StopOfflineTaskResponseBodyResult {
	s.Sink = v
	return s
}

func (s *StopOfflineTaskResponseBodyResult) SetSource(v []*StopOfflineTaskResponseBodyResultSource) *StopOfflineTaskResponseBodyResult {
	s.Source = v
	return s
}

func (s *StopOfflineTaskResponseBodyResult) SetStatus(v *StopOfflineTaskResponseBodyResultStatus) *StopOfflineTaskResponseBodyResult {
	s.Status = v
	return s
}

func (s *StopOfflineTaskResponseBodyResult) Validate() error {
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

type StopOfflineTaskResponseBodyResultMeta struct {
	// The billing specification.
	//
	// example:
	//
	// small
	ComputeResource *string `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// test
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s StopOfflineTaskResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBodyResultMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *StopOfflineTaskResponseBodyResultMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *StopOfflineTaskResponseBodyResultMeta) SetComputeResource(v string) *StopOfflineTaskResponseBodyResultMeta {
	s.ComputeResource = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultMeta) SetTaskName(v string) *StopOfflineTaskResponseBodyResultMeta {
	s.TaskName = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}

type StopOfflineTaskResponseBodyResultProcessors struct {
	// The input parameters.
	//
	// example:
	//
	// {
	//
	//   "key": "content"
	//
	// }
	Input map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
	// The name.
	//
	// example:
	//
	// "processor1"
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The processor processing parameters.
	//
	// example:
	//
	// {
	//
	//   "service_id": "xxx"
	//
	// }
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The type.
	//
	// example:
	//
	// document-analyze
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StopOfflineTaskResponseBodyResultProcessors) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBodyResultProcessors) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBodyResultProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *StopOfflineTaskResponseBodyResultProcessors) GetName() *string {
	return s.Name
}

func (s *StopOfflineTaskResponseBodyResultProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *StopOfflineTaskResponseBodyResultProcessors) GetType() *string {
	return s.Type
}

func (s *StopOfflineTaskResponseBodyResultProcessors) SetInput(v map[string]interface{}) *StopOfflineTaskResponseBodyResultProcessors {
	s.Input = v
	return s
}

func (s *StopOfflineTaskResponseBodyResultProcessors) SetName(v string) *StopOfflineTaskResponseBodyResultProcessors {
	s.Name = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultProcessors) SetParameters(v map[string]interface{}) *StopOfflineTaskResponseBodyResultProcessors {
	s.Parameters = v
	return s
}

func (s *StopOfflineTaskResponseBodyResultProcessors) SetType(v string) *StopOfflineTaskResponseBodyResultProcessors {
	s.Type = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultProcessors) Validate() error {
	return dara.Validate(s)
}

type StopOfflineTaskResponseBodyResultSink struct {
	// The name of the data sink.
	//
	// example:
	//
	// table2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configuration parameters of the data sink.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The primary key field of the data sink.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// The schema of the data sink.
	Schema []map[string]*string `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	// The type of the data sink.
	//
	// example:
	//
	// swift
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StopOfflineTaskResponseBodyResultSink) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBodyResultSink) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBodyResultSink) GetName() *string {
	return s.Name
}

func (s *StopOfflineTaskResponseBodyResultSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *StopOfflineTaskResponseBodyResultSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *StopOfflineTaskResponseBodyResultSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *StopOfflineTaskResponseBodyResultSink) GetType() *string {
	return s.Type
}

func (s *StopOfflineTaskResponseBodyResultSink) SetName(v string) *StopOfflineTaskResponseBodyResultSink {
	s.Name = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSink) SetParameters(v map[string]*string) *StopOfflineTaskResponseBodyResultSink {
	s.Parameters = v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSink) SetPrimaryKey(v string) *StopOfflineTaskResponseBodyResultSink {
	s.PrimaryKey = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSink) SetSchema(v []map[string]*string) *StopOfflineTaskResponseBodyResultSink {
	s.Schema = v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSink) SetType(v string) *StopOfflineTaskResponseBodyResultSink {
	s.Type = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSink) Validate() error {
	return dara.Validate(s)
}

type StopOfflineTaskResponseBodyResultSource struct {
	// The name of the data source.
	//
	// example:
	//
	// table1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The datasource config parameters.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The primary key field of the data source.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// The schema of the data source.
	Schema []map[string]*string `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	// The type of the data source.
	//
	// example:
	//
	// swift
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StopOfflineTaskResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBodyResultSource) GetName() *string {
	return s.Name
}

func (s *StopOfflineTaskResponseBodyResultSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *StopOfflineTaskResponseBodyResultSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *StopOfflineTaskResponseBodyResultSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *StopOfflineTaskResponseBodyResultSource) GetType() *string {
	return s.Type
}

func (s *StopOfflineTaskResponseBodyResultSource) SetName(v string) *StopOfflineTaskResponseBodyResultSource {
	s.Name = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSource) SetParameters(v map[string]*string) *StopOfflineTaskResponseBodyResultSource {
	s.Parameters = v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSource) SetPrimaryKey(v string) *StopOfflineTaskResponseBodyResultSource {
	s.PrimaryKey = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSource) SetSchema(v []map[string]*string) *StopOfflineTaskResponseBodyResultSource {
	s.Schema = v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSource) SetType(v string) *StopOfflineTaskResponseBodyResultSource {
	s.Type = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}

type StopOfflineTaskResponseBodyResultStatus struct {
	// The time when the node was started.
	//
	// example:
	//
	// 1744941600000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the node was stopped.
	//
	// example:
	//
	// 1744941600000
	DeleteTime *int64 `json:"deleteTime,omitempty" xml:"deleteTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The node status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StopOfflineTaskResponseBodyResultStatus) String() string {
	return dara.Prettify(s)
}

func (s StopOfflineTaskResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *StopOfflineTaskResponseBodyResultStatus) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *StopOfflineTaskResponseBodyResultStatus) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *StopOfflineTaskResponseBodyResultStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopOfflineTaskResponseBodyResultStatus) GetStatus() *string {
	return s.Status
}

func (s *StopOfflineTaskResponseBodyResultStatus) SetCreateTime(v int64) *StopOfflineTaskResponseBodyResultStatus {
	s.CreateTime = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultStatus) SetDeleteTime(v int64) *StopOfflineTaskResponseBodyResultStatus {
	s.DeleteTime = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultStatus) SetErrorMessage(v string) *StopOfflineTaskResponseBodyResultStatus {
	s.ErrorMessage = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultStatus) SetStatus(v string) *StopOfflineTaskResponseBodyResultStatus {
	s.Status = &v
	return s
}

func (s *StopOfflineTaskResponseBodyResultStatus) Validate() error {
	return dara.Validate(s)
}
