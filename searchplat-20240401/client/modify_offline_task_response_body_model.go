// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v *ModifyOfflineTaskResponseBodyResult) *ModifyOfflineTaskResponseBody
	GetResult() *ModifyOfflineTaskResponseBodyResult
}

type ModifyOfflineTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0abb793917165176014887584e28d9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *ModifyOfflineTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyOfflineTaskResponseBody) GetResult() *ModifyOfflineTaskResponseBodyResult {
	return s.Result
}

func (s *ModifyOfflineTaskResponseBody) SetRequestId(v string) *ModifyOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyOfflineTaskResponseBody) SetResult(v *ModifyOfflineTaskResponseBodyResult) *ModifyOfflineTaskResponseBody {
	s.Result = v
	return s
}

func (s *ModifyOfflineTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyOfflineTaskResponseBodyResult struct {
	// The metadata.
	Meta *ModifyOfflineTaskResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The task processing parameters.
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
	Processors []*ModifyOfflineTaskResponseBodyResultProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*ModifyOfflineTaskResponseBodyResultSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*ModifyOfflineTaskResponseBodyResultSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The task status.
	Status *ModifyOfflineTaskResponseBodyResultStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s ModifyOfflineTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBodyResult) GetMeta() *ModifyOfflineTaskResponseBodyResultMeta {
	return s.Meta
}

func (s *ModifyOfflineTaskResponseBodyResult) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ModifyOfflineTaskResponseBodyResult) GetProcessors() []*ModifyOfflineTaskResponseBodyResultProcessors {
	return s.Processors
}

func (s *ModifyOfflineTaskResponseBodyResult) GetSink() []*ModifyOfflineTaskResponseBodyResultSink {
	return s.Sink
}

func (s *ModifyOfflineTaskResponseBodyResult) GetSource() []*ModifyOfflineTaskResponseBodyResultSource {
	return s.Source
}

func (s *ModifyOfflineTaskResponseBodyResult) GetStatus() *ModifyOfflineTaskResponseBodyResultStatus {
	return s.Status
}

func (s *ModifyOfflineTaskResponseBodyResult) SetMeta(v *ModifyOfflineTaskResponseBodyResultMeta) *ModifyOfflineTaskResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResult) SetParameters(v map[string]interface{}) *ModifyOfflineTaskResponseBodyResult {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResult) SetProcessors(v []*ModifyOfflineTaskResponseBodyResultProcessors) *ModifyOfflineTaskResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResult) SetSink(v []*ModifyOfflineTaskResponseBodyResultSink) *ModifyOfflineTaskResponseBodyResult {
	s.Sink = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResult) SetSource(v []*ModifyOfflineTaskResponseBodyResultSource) *ModifyOfflineTaskResponseBodyResult {
	s.Source = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResult) SetStatus(v *ModifyOfflineTaskResponseBodyResultStatus) *ModifyOfflineTaskResponseBodyResult {
	s.Status = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResult) Validate() error {
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

type ModifyOfflineTaskResponseBodyResultMeta struct {
	// The billing specification.
	//
	// example:
	//
	// small
	ComputeResource *string `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The task name.
	//
	// example:
	//
	// test
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ModifyOfflineTaskResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBodyResultMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *ModifyOfflineTaskResponseBodyResultMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyOfflineTaskResponseBodyResultMeta) SetComputeResource(v string) *ModifyOfflineTaskResponseBodyResultMeta {
	s.ComputeResource = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultMeta) SetTaskName(v string) *ModifyOfflineTaskResponseBodyResultMeta {
	s.TaskName = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskResponseBodyResultProcessors struct {
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
	// processor1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The processor parameters.
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
	// The type.
	//
	// example:
	//
	// document-analyze
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyOfflineTaskResponseBodyResultProcessors) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBodyResultProcessors) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) GetName() *string {
	return s.Name
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) GetType() *string {
	return s.Type
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) SetInput(v map[string]interface{}) *ModifyOfflineTaskResponseBodyResultProcessors {
	s.Input = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) SetName(v string) *ModifyOfflineTaskResponseBodyResultProcessors {
	s.Name = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) SetParameters(v map[string]interface{}) *ModifyOfflineTaskResponseBodyResultProcessors {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) SetType(v string) *ModifyOfflineTaskResponseBodyResultProcessors {
	s.Type = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultProcessors) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskResponseBodyResultSink struct {
	// The data sink name.
	//
	// example:
	//
	// table2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data sink configuration parameters, which are determined by the type.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The primary key field of the data sink.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// The data sink schema.
	Schema []map[string]*string `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	// The data sink type.
	//
	// example:
	//
	// swift
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyOfflineTaskResponseBodyResultSink) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBodyResultSink) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBodyResultSink) GetName() *string {
	return s.Name
}

func (s *ModifyOfflineTaskResponseBodyResultSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ModifyOfflineTaskResponseBodyResultSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ModifyOfflineTaskResponseBodyResultSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *ModifyOfflineTaskResponseBodyResultSink) GetType() *string {
	return s.Type
}

func (s *ModifyOfflineTaskResponseBodyResultSink) SetName(v string) *ModifyOfflineTaskResponseBodyResultSink {
	s.Name = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSink) SetParameters(v map[string]*string) *ModifyOfflineTaskResponseBodyResultSink {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSink) SetPrimaryKey(v string) *ModifyOfflineTaskResponseBodyResultSink {
	s.PrimaryKey = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSink) SetSchema(v []map[string]*string) *ModifyOfflineTaskResponseBodyResultSink {
	s.Schema = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSink) SetType(v string) *ModifyOfflineTaskResponseBodyResultSink {
	s.Type = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSink) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskResponseBodyResultSource struct {
	// The data source name.
	//
	// example:
	//
	// table1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The datasource config parameters, which are determined by the type.
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The primary key field of the data source.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// The data source schema.
	Schema []map[string]*string `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	// The data source type.
	//
	// example:
	//
	// swift
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyOfflineTaskResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBodyResultSource) GetName() *string {
	return s.Name
}

func (s *ModifyOfflineTaskResponseBodyResultSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ModifyOfflineTaskResponseBodyResultSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ModifyOfflineTaskResponseBodyResultSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *ModifyOfflineTaskResponseBodyResultSource) GetType() *string {
	return s.Type
}

func (s *ModifyOfflineTaskResponseBodyResultSource) SetName(v string) *ModifyOfflineTaskResponseBodyResultSource {
	s.Name = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSource) SetParameters(v map[string]*string) *ModifyOfflineTaskResponseBodyResultSource {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSource) SetPrimaryKey(v string) *ModifyOfflineTaskResponseBodyResultSource {
	s.PrimaryKey = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSource) SetSchema(v []map[string]*string) *ModifyOfflineTaskResponseBodyResultSource {
	s.Schema = v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSource) SetType(v string) *ModifyOfflineTaskResponseBodyResultSource {
	s.Type = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskResponseBodyResultStatus struct {
	// The task start time.
	//
	// example:
	//
	// 1744941600000
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The task stop time.
	//
	// example:
	//
	// 1744941600000
	DeleteTime *string `json:"deleteTime,omitempty" xml:"deleteTime,omitempty"`
	// The error message.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModifyOfflineTaskResponseBodyResultStatus) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) GetDeleteTime() *string {
	return s.DeleteTime
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) GetStatus() *string {
	return s.Status
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) SetCreateTime(v string) *ModifyOfflineTaskResponseBodyResultStatus {
	s.CreateTime = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) SetDeleteTime(v string) *ModifyOfflineTaskResponseBodyResultStatus {
	s.DeleteTime = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) SetErrorMessage(v string) *ModifyOfflineTaskResponseBodyResultStatus {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) SetStatus(v string) *ModifyOfflineTaskResponseBodyResultStatus {
	s.Status = &v
	return s
}

func (s *ModifyOfflineTaskResponseBodyResultStatus) Validate() error {
	return dara.Validate(s)
}
