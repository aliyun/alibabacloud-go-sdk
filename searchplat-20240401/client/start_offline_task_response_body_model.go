// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v *StartOfflineTaskResponseBodyResult) *StartOfflineTaskResponseBody
	GetResult() *StartOfflineTaskResponseBodyResult
}

type StartOfflineTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1CC93E65-6734-5060-BEF7-0EB0A4862BCF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *StartOfflineTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s StartOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartOfflineTaskResponseBody) GetResult() *StartOfflineTaskResponseBodyResult {
	return s.Result
}

func (s *StartOfflineTaskResponseBody) SetRequestId(v string) *StartOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartOfflineTaskResponseBody) SetResult(v *StartOfflineTaskResponseBodyResult) *StartOfflineTaskResponseBody {
	s.Result = v
	return s
}

func (s *StartOfflineTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartOfflineTaskResponseBodyResult struct {
	// The metadata.
	Meta *StartOfflineTaskResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The task processing parameters.
	//
	// example:
	//
	// {
	//
	//   "parameter1": {
	//
	//        "key": "value"
	//
	//     },
	//
	//     "parameter2": {
	//
	//          "key": "value"
	//
	//      }
	//
	// }
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The processing operators.
	Processors []*StartOfflineTaskResponseBodyResultProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*StartOfflineTaskResponseBodyResultSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The source.
	Source []*StartOfflineTaskResponseBodyResultSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The task status. Valid values:
	//
	// - PENDING: In progress.
	//
	// - SUCCESS: Parsing succeeded.
	//
	// - FAILED: Parsing failed.
	Status *StartOfflineTaskResponseBodyResultStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s StartOfflineTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBodyResult) GetMeta() *StartOfflineTaskResponseBodyResultMeta {
	return s.Meta
}

func (s *StartOfflineTaskResponseBodyResult) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *StartOfflineTaskResponseBodyResult) GetProcessors() []*StartOfflineTaskResponseBodyResultProcessors {
	return s.Processors
}

func (s *StartOfflineTaskResponseBodyResult) GetSink() []*StartOfflineTaskResponseBodyResultSink {
	return s.Sink
}

func (s *StartOfflineTaskResponseBodyResult) GetSource() []*StartOfflineTaskResponseBodyResultSource {
	return s.Source
}

func (s *StartOfflineTaskResponseBodyResult) GetStatus() *StartOfflineTaskResponseBodyResultStatus {
	return s.Status
}

func (s *StartOfflineTaskResponseBodyResult) SetMeta(v *StartOfflineTaskResponseBodyResultMeta) *StartOfflineTaskResponseBodyResult {
	s.Meta = v
	return s
}

func (s *StartOfflineTaskResponseBodyResult) SetParameters(v map[string]interface{}) *StartOfflineTaskResponseBodyResult {
	s.Parameters = v
	return s
}

func (s *StartOfflineTaskResponseBodyResult) SetProcessors(v []*StartOfflineTaskResponseBodyResultProcessors) *StartOfflineTaskResponseBodyResult {
	s.Processors = v
	return s
}

func (s *StartOfflineTaskResponseBodyResult) SetSink(v []*StartOfflineTaskResponseBodyResultSink) *StartOfflineTaskResponseBodyResult {
	s.Sink = v
	return s
}

func (s *StartOfflineTaskResponseBodyResult) SetSource(v []*StartOfflineTaskResponseBodyResultSource) *StartOfflineTaskResponseBodyResult {
	s.Source = v
	return s
}

func (s *StartOfflineTaskResponseBodyResult) SetStatus(v *StartOfflineTaskResponseBodyResultStatus) *StartOfflineTaskResponseBodyResult {
	s.Status = v
	return s
}

func (s *StartOfflineTaskResponseBodyResult) Validate() error {
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

type StartOfflineTaskResponseBodyResultMeta struct {
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
	// taskName
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s StartOfflineTaskResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBodyResultMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *StartOfflineTaskResponseBodyResultMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *StartOfflineTaskResponseBodyResultMeta) SetComputeResource(v string) *StartOfflineTaskResponseBodyResultMeta {
	s.ComputeResource = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultMeta) SetTaskName(v string) *StartOfflineTaskResponseBodyResultMeta {
	s.TaskName = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}

type StartOfflineTaskResponseBodyResultProcessors struct {
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
	// The data source name.
	//
	// example:
	//
	// processor1
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
	// The data sink type.
	//
	// example:
	//
	// document-analyze
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StartOfflineTaskResponseBodyResultProcessors) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBodyResultProcessors) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBodyResultProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *StartOfflineTaskResponseBodyResultProcessors) GetName() *string {
	return s.Name
}

func (s *StartOfflineTaskResponseBodyResultProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *StartOfflineTaskResponseBodyResultProcessors) GetType() *string {
	return s.Type
}

func (s *StartOfflineTaskResponseBodyResultProcessors) SetInput(v map[string]interface{}) *StartOfflineTaskResponseBodyResultProcessors {
	s.Input = v
	return s
}

func (s *StartOfflineTaskResponseBodyResultProcessors) SetName(v string) *StartOfflineTaskResponseBodyResultProcessors {
	s.Name = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultProcessors) SetParameters(v map[string]interface{}) *StartOfflineTaskResponseBodyResultProcessors {
	s.Parameters = v
	return s
}

func (s *StartOfflineTaskResponseBodyResultProcessors) SetType(v string) *StartOfflineTaskResponseBodyResultProcessors {
	s.Type = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultProcessors) Validate() error {
	return dara.Validate(s)
}

type StartOfflineTaskResponseBodyResultSink struct {
	// The task name.
	//
	// example:
	//
	// milvus-ali-cn-hangzhou-1
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
	// The type. Valid values:
	//
	// - standard.
	//
	// example:
	//
	// ModuleRelation
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s StartOfflineTaskResponseBodyResultSink) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBodyResultSink) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBodyResultSink) GetName() *string {
	return s.Name
}

func (s *StartOfflineTaskResponseBodyResultSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *StartOfflineTaskResponseBodyResultSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *StartOfflineTaskResponseBodyResultSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *StartOfflineTaskResponseBodyResultSink) GetType() *string {
	return s.Type
}

func (s *StartOfflineTaskResponseBodyResultSink) SetName(v string) *StartOfflineTaskResponseBodyResultSink {
	s.Name = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSink) SetParameters(v map[string]*string) *StartOfflineTaskResponseBodyResultSink {
	s.Parameters = v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSink) SetPrimaryKey(v string) *StartOfflineTaskResponseBodyResultSink {
	s.PrimaryKey = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSink) SetSchema(v []map[string]*string) *StartOfflineTaskResponseBodyResultSink {
	s.Schema = v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSink) SetType(v string) *StartOfflineTaskResponseBodyResultSink {
	s.Type = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSink) Validate() error {
	return dara.Validate(s)
}

type StartOfflineTaskResponseBodyResultSource struct {
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

func (s StartOfflineTaskResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBodyResultSource) GetName() *string {
	return s.Name
}

func (s *StartOfflineTaskResponseBodyResultSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *StartOfflineTaskResponseBodyResultSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *StartOfflineTaskResponseBodyResultSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *StartOfflineTaskResponseBodyResultSource) GetType() *string {
	return s.Type
}

func (s *StartOfflineTaskResponseBodyResultSource) SetName(v string) *StartOfflineTaskResponseBodyResultSource {
	s.Name = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSource) SetParameters(v map[string]*string) *StartOfflineTaskResponseBodyResultSource {
	s.Parameters = v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSource) SetPrimaryKey(v string) *StartOfflineTaskResponseBodyResultSource {
	s.PrimaryKey = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSource) SetSchema(v []map[string]*string) *StartOfflineTaskResponseBodyResultSource {
	s.Schema = v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSource) SetType(v string) *StartOfflineTaskResponseBodyResultSource {
	s.Type = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}

type StartOfflineTaskResponseBodyResultStatus struct {
	// The task start time.
	//
	// example:
	//
	// 1744941600000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The task stop time.
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
	// The request status.
	//
	// example:
	//
	// OK
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StartOfflineTaskResponseBodyResultStatus) String() string {
	return dara.Prettify(s)
}

func (s StartOfflineTaskResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *StartOfflineTaskResponseBodyResultStatus) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *StartOfflineTaskResponseBodyResultStatus) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *StartOfflineTaskResponseBodyResultStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartOfflineTaskResponseBodyResultStatus) GetStatus() *string {
	return s.Status
}

func (s *StartOfflineTaskResponseBodyResultStatus) SetCreateTime(v int64) *StartOfflineTaskResponseBodyResultStatus {
	s.CreateTime = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultStatus) SetDeleteTime(v int64) *StartOfflineTaskResponseBodyResultStatus {
	s.DeleteTime = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultStatus) SetErrorMessage(v string) *StartOfflineTaskResponseBodyResultStatus {
	s.ErrorMessage = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultStatus) SetStatus(v string) *StartOfflineTaskResponseBodyResultStatus {
	s.Status = &v
	return s
}

func (s *StartOfflineTaskResponseBodyResultStatus) Validate() error {
	return dara.Validate(s)
}
