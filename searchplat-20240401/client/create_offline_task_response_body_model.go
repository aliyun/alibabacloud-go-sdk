// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfflineTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOfflineTaskResponseBody
	GetRequestId() *string
	SetResult(v *CreateOfflineTaskResponseBodyResult) *CreateOfflineTaskResponseBody
	GetResult() *CreateOfflineTaskResponseBodyResult
}

type CreateOfflineTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B7C901ED-2BC1-5CFB-BE23-242DE5E3BA5C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response result.
	Result *CreateOfflineTaskResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateOfflineTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOfflineTaskResponseBody) GetResult() *CreateOfflineTaskResponseBodyResult {
	return s.Result
}

func (s *CreateOfflineTaskResponseBody) SetRequestId(v string) *CreateOfflineTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOfflineTaskResponseBody) SetResult(v *CreateOfflineTaskResponseBodyResult) *CreateOfflineTaskResponseBody {
	s.Result = v
	return s
}

func (s *CreateOfflineTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOfflineTaskResponseBodyResult struct {
	// The task metadata.
	Meta *CreateOfflineTaskResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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
	// The processing pipeline operators.
	Processors []*CreateOfflineTaskResponseBodyResultProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The output destination information.
	Sink []*CreateOfflineTaskResponseBodyResultSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*CreateOfflineTaskResponseBodyResultSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The task status.
	Status *CreateOfflineTaskResponseBodyResultStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
}

func (s CreateOfflineTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBodyResult) GetMeta() *CreateOfflineTaskResponseBodyResultMeta {
	return s.Meta
}

func (s *CreateOfflineTaskResponseBodyResult) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateOfflineTaskResponseBodyResult) GetProcessors() []*CreateOfflineTaskResponseBodyResultProcessors {
	return s.Processors
}

func (s *CreateOfflineTaskResponseBodyResult) GetSink() []*CreateOfflineTaskResponseBodyResultSink {
	return s.Sink
}

func (s *CreateOfflineTaskResponseBodyResult) GetSource() []*CreateOfflineTaskResponseBodyResultSource {
	return s.Source
}

func (s *CreateOfflineTaskResponseBodyResult) GetStatus() *CreateOfflineTaskResponseBodyResultStatus {
	return s.Status
}

func (s *CreateOfflineTaskResponseBodyResult) SetMeta(v *CreateOfflineTaskResponseBodyResultMeta) *CreateOfflineTaskResponseBodyResult {
	s.Meta = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResult) SetParameters(v map[string]interface{}) *CreateOfflineTaskResponseBodyResult {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResult) SetProcessors(v []*CreateOfflineTaskResponseBodyResultProcessors) *CreateOfflineTaskResponseBodyResult {
	s.Processors = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResult) SetSink(v []*CreateOfflineTaskResponseBodyResultSink) *CreateOfflineTaskResponseBodyResult {
	s.Sink = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResult) SetSource(v []*CreateOfflineTaskResponseBodyResultSource) *CreateOfflineTaskResponseBodyResult {
	s.Source = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResult) SetStatus(v *CreateOfflineTaskResponseBodyResultStatus) *CreateOfflineTaskResponseBodyResult {
	s.Status = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResult) Validate() error {
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

type CreateOfflineTaskResponseBodyResultMeta struct {
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

func (s CreateOfflineTaskResponseBodyResultMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBodyResultMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *CreateOfflineTaskResponseBodyResultMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateOfflineTaskResponseBodyResultMeta) SetComputeResource(v string) *CreateOfflineTaskResponseBodyResultMeta {
	s.ComputeResource = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultMeta) SetTaskName(v string) *CreateOfflineTaskResponseBodyResultMeta {
	s.TaskName = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultMeta) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskResponseBodyResultProcessors struct {
	// The input parameters.
	Input map[string]interface{} `json:"input,omitempty" xml:"input,omitempty"`
	// The name.
	//
	// example:
	//
	// "processor1"
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The processor parameters.
	//
	// example:
	//
	// {
	//
	// "service_id": "xxx"
	//
	// }
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The type.
	//
	// example:
	//
	// "document-analyze"
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateOfflineTaskResponseBodyResultProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBodyResultProcessors) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) GetName() *string {
	return s.Name
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) GetType() *string {
	return s.Type
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) SetInput(v map[string]interface{}) *CreateOfflineTaskResponseBodyResultProcessors {
	s.Input = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) SetName(v string) *CreateOfflineTaskResponseBodyResultProcessors {
	s.Name = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) SetParameters(v map[string]interface{}) *CreateOfflineTaskResponseBodyResultProcessors {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) SetType(v string) *CreateOfflineTaskResponseBodyResultProcessors {
	s.Type = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultProcessors) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskResponseBodyResultSink struct {
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

func (s CreateOfflineTaskResponseBodyResultSink) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBodyResultSink) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBodyResultSink) GetName() *string {
	return s.Name
}

func (s *CreateOfflineTaskResponseBodyResultSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *CreateOfflineTaskResponseBodyResultSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *CreateOfflineTaskResponseBodyResultSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *CreateOfflineTaskResponseBodyResultSink) GetType() *string {
	return s.Type
}

func (s *CreateOfflineTaskResponseBodyResultSink) SetName(v string) *CreateOfflineTaskResponseBodyResultSink {
	s.Name = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSink) SetParameters(v map[string]*string) *CreateOfflineTaskResponseBodyResultSink {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSink) SetPrimaryKey(v string) *CreateOfflineTaskResponseBodyResultSink {
	s.PrimaryKey = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSink) SetSchema(v []map[string]*string) *CreateOfflineTaskResponseBodyResultSink {
	s.Schema = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSink) SetType(v string) *CreateOfflineTaskResponseBodyResultSink {
	s.Type = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSink) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskResponseBodyResultSource struct {
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

func (s CreateOfflineTaskResponseBodyResultSource) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBodyResultSource) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBodyResultSource) GetName() *string {
	return s.Name
}

func (s *CreateOfflineTaskResponseBodyResultSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *CreateOfflineTaskResponseBodyResultSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *CreateOfflineTaskResponseBodyResultSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *CreateOfflineTaskResponseBodyResultSource) GetType() *string {
	return s.Type
}

func (s *CreateOfflineTaskResponseBodyResultSource) SetName(v string) *CreateOfflineTaskResponseBodyResultSource {
	s.Name = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSource) SetParameters(v map[string]*string) *CreateOfflineTaskResponseBodyResultSource {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSource) SetPrimaryKey(v string) *CreateOfflineTaskResponseBodyResultSource {
	s.PrimaryKey = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSource) SetSchema(v []map[string]*string) *CreateOfflineTaskResponseBodyResultSource {
	s.Schema = v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSource) SetType(v string) *CreateOfflineTaskResponseBodyResultSource {
	s.Type = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultSource) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskResponseBodyResultStatus struct {
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

func (s CreateOfflineTaskResponseBodyResultStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskResponseBodyResultStatus) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskResponseBodyResultStatus) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateOfflineTaskResponseBodyResultStatus) GetDeleteTime() *string {
	return s.DeleteTime
}

func (s *CreateOfflineTaskResponseBodyResultStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateOfflineTaskResponseBodyResultStatus) GetStatus() *string {
	return s.Status
}

func (s *CreateOfflineTaskResponseBodyResultStatus) SetCreateTime(v string) *CreateOfflineTaskResponseBodyResultStatus {
	s.CreateTime = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultStatus) SetDeleteTime(v string) *CreateOfflineTaskResponseBodyResultStatus {
	s.DeleteTime = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultStatus) SetErrorMessage(v string) *CreateOfflineTaskResponseBodyResultStatus {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultStatus) SetStatus(v string) *CreateOfflineTaskResponseBodyResultStatus {
	s.Status = &v
	return s
}

func (s *CreateOfflineTaskResponseBodyResultStatus) Validate() error {
	return dara.Validate(s)
}
