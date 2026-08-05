// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v *GetOfflineTaskResponseBodyResult) *GetOfflineTaskResponseBody
	GetResult() *GetOfflineTaskResponseBodyResult
}

type GetOfflineTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 5950143C-B8F0-5758-A08A-66F302FD587F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	Result *GetOfflineTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOfflineTaskResponseBody) GetResult() *GetOfflineTaskResponseBodyResult {
	return s.Result
}

func (s *GetOfflineTaskResponseBody) SetRequestId(v string) *GetOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfflineTaskResponseBody) SetResult(v *GetOfflineTaskResponseBodyResult) *GetOfflineTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetOfflineTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOfflineTaskResponseBodyResult struct {
	// The node metadata.
	Meta *GetOfflineTaskResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The node processing parameters.
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
	// The processing flow operators.
	Processors []*GetOfflineTaskResponseBodyResultProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*GetOfflineTaskResponseBodyResultSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*GetOfflineTaskResponseBodyResultSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The node status.
	Status *GetOfflineTaskResponseBodyResultStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s GetOfflineTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBodyResult) GetMeta() *GetOfflineTaskResponseBodyResultMeta {
	return s.Meta
}

func (s *GetOfflineTaskResponseBodyResult) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetOfflineTaskResponseBodyResult) GetProcessors() []*GetOfflineTaskResponseBodyResultProcessors {
	return s.Processors
}

func (s *GetOfflineTaskResponseBodyResult) GetSink() []*GetOfflineTaskResponseBodyResultSink {
	return s.Sink
}

func (s *GetOfflineTaskResponseBodyResult) GetSource() []*GetOfflineTaskResponseBodyResultSource {
	return s.Source
}

func (s *GetOfflineTaskResponseBodyResult) GetStatus() *GetOfflineTaskResponseBodyResultStatus {
	return s.Status
}

func (s *GetOfflineTaskResponseBodyResult) SetMeta(v *GetOfflineTaskResponseBodyResultMeta) *GetOfflineTaskResponseBodyResult {
	s.Meta = v
	return s
}

func (s *GetOfflineTaskResponseBodyResult) SetParameters(v map[string]interface{}) *GetOfflineTaskResponseBodyResult {
	s.Parameters = v
	return s
}

func (s *GetOfflineTaskResponseBodyResult) SetProcessors(v []*GetOfflineTaskResponseBodyResultProcessors) *GetOfflineTaskResponseBodyResult {
	s.Processors = v
	return s
}

func (s *GetOfflineTaskResponseBodyResult) SetSink(v []*GetOfflineTaskResponseBodyResultSink) *GetOfflineTaskResponseBodyResult {
	s.Sink = v
	return s
}

func (s *GetOfflineTaskResponseBodyResult) SetSource(v []*GetOfflineTaskResponseBodyResultSource) *GetOfflineTaskResponseBodyResult {
	s.Source = v
	return s
}

func (s *GetOfflineTaskResponseBodyResult) SetStatus(v *GetOfflineTaskResponseBodyResultStatus) *GetOfflineTaskResponseBodyResult {
	s.Status = v
	return s
}

func (s *GetOfflineTaskResponseBodyResult) Validate() error {
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

type GetOfflineTaskResponseBodyResultMeta struct {
	// The billing specification.
	//
	// example:
	//
	// small
	ComputeResource *string `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The node name.
	//
	// example:
	//
	// test
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s GetOfflineTaskResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBodyResultMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *GetOfflineTaskResponseBodyResultMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *GetOfflineTaskResponseBodyResultMeta) SetComputeResource(v string) *GetOfflineTaskResponseBodyResultMeta {
	s.ComputeResource = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultMeta) SetTaskName(v string) *GetOfflineTaskResponseBodyResultMeta {
	s.TaskName = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}

type GetOfflineTaskResponseBodyResultProcessors struct {
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

func (s GetOfflineTaskResponseBodyResultProcessors) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBodyResultProcessors) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBodyResultProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *GetOfflineTaskResponseBodyResultProcessors) GetName() *string {
	return s.Name
}

func (s *GetOfflineTaskResponseBodyResultProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetOfflineTaskResponseBodyResultProcessors) GetType() *string {
	return s.Type
}

func (s *GetOfflineTaskResponseBodyResultProcessors) SetInput(v map[string]interface{}) *GetOfflineTaskResponseBodyResultProcessors {
	s.Input = v
	return s
}

func (s *GetOfflineTaskResponseBodyResultProcessors) SetName(v string) *GetOfflineTaskResponseBodyResultProcessors {
	s.Name = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultProcessors) SetParameters(v map[string]interface{}) *GetOfflineTaskResponseBodyResultProcessors {
	s.Parameters = v
	return s
}

func (s *GetOfflineTaskResponseBodyResultProcessors) SetType(v string) *GetOfflineTaskResponseBodyResultProcessors {
	s.Type = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultProcessors) Validate() error {
	return dara.Validate(s)
}

type GetOfflineTaskResponseBodyResultSink struct {
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

func (s GetOfflineTaskResponseBodyResultSink) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBodyResultSink) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBodyResultSink) GetName() *string {
	return s.Name
}

func (s *GetOfflineTaskResponseBodyResultSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *GetOfflineTaskResponseBodyResultSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetOfflineTaskResponseBodyResultSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *GetOfflineTaskResponseBodyResultSink) GetType() *string {
	return s.Type
}

func (s *GetOfflineTaskResponseBodyResultSink) SetName(v string) *GetOfflineTaskResponseBodyResultSink {
	s.Name = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSink) SetParameters(v map[string]*string) *GetOfflineTaskResponseBodyResultSink {
	s.Parameters = v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSink) SetPrimaryKey(v string) *GetOfflineTaskResponseBodyResultSink {
	s.PrimaryKey = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSink) SetSchema(v []map[string]*string) *GetOfflineTaskResponseBodyResultSink {
	s.Schema = v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSink) SetType(v string) *GetOfflineTaskResponseBodyResultSink {
	s.Type = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSink) Validate() error {
	return dara.Validate(s)
}

type GetOfflineTaskResponseBodyResultSource struct {
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

func (s GetOfflineTaskResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBodyResultSource) GetName() *string {
	return s.Name
}

func (s *GetOfflineTaskResponseBodyResultSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *GetOfflineTaskResponseBodyResultSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetOfflineTaskResponseBodyResultSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *GetOfflineTaskResponseBodyResultSource) GetType() *string {
	return s.Type
}

func (s *GetOfflineTaskResponseBodyResultSource) SetName(v string) *GetOfflineTaskResponseBodyResultSource {
	s.Name = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSource) SetParameters(v map[string]*string) *GetOfflineTaskResponseBodyResultSource {
	s.Parameters = v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSource) SetPrimaryKey(v string) *GetOfflineTaskResponseBodyResultSource {
	s.PrimaryKey = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSource) SetSchema(v []map[string]*string) *GetOfflineTaskResponseBodyResultSource {
	s.Schema = v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSource) SetType(v string) *GetOfflineTaskResponseBodyResultSource {
	s.Type = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}

type GetOfflineTaskResponseBodyResultStatus struct {
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
	// “”
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The node status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetOfflineTaskResponseBodyResultStatus) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskResponseBodyResultStatus) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetOfflineTaskResponseBodyResultStatus) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *GetOfflineTaskResponseBodyResultStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOfflineTaskResponseBodyResultStatus) GetStatus() *string {
	return s.Status
}

func (s *GetOfflineTaskResponseBodyResultStatus) SetCreateTime(v int64) *GetOfflineTaskResponseBodyResultStatus {
	s.CreateTime = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultStatus) SetDeleteTime(v int64) *GetOfflineTaskResponseBodyResultStatus {
	s.DeleteTime = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultStatus) SetErrorMessage(v string) *GetOfflineTaskResponseBodyResultStatus {
	s.ErrorMessage = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultStatus) SetStatus(v string) *GetOfflineTaskResponseBodyResultStatus {
	s.Status = &v
	return s
}

func (s *GetOfflineTaskResponseBodyResultStatus) Validate() error {
	return dara.Validate(s)
}
