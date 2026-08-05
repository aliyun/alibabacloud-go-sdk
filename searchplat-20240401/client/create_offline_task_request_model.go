// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMeta(v *CreateOfflineTaskRequestMeta) *CreateOfflineTaskRequest
	GetMeta() *CreateOfflineTaskRequestMeta
	SetParameters(v map[string]interface{}) *CreateOfflineTaskRequest
	GetParameters() map[string]interface{}
	SetProcessors(v []*CreateOfflineTaskRequestProcessors) *CreateOfflineTaskRequest
	GetProcessors() []*CreateOfflineTaskRequestProcessors
	SetSink(v []*CreateOfflineTaskRequestSink) *CreateOfflineTaskRequest
	GetSink() []*CreateOfflineTaskRequestSink
	SetSource(v []*CreateOfflineTaskRequestSource) *CreateOfflineTaskRequest
	GetSource() []*CreateOfflineTaskRequestSource
	SetStatus(v *CreateOfflineTaskRequestStatus) *CreateOfflineTaskRequest
	GetStatus() *CreateOfflineTaskRequestStatus
	SetDraft(v bool) *CreateOfflineTaskRequest
	GetDraft() *bool
	SetDryRun(v bool) *CreateOfflineTaskRequest
	GetDryRun() *bool
	SetRegionId(v string) *CreateOfflineTaskRequest
	GetRegionId() *string
}

type CreateOfflineTaskRequest struct {
	// The task metadata.
	Meta *CreateOfflineTaskRequestMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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
	// The processing pipeline operators.
	Processors []*CreateOfflineTaskRequestProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*CreateOfflineTaskRequestSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*CreateOfflineTaskRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The task status.
	Status *CreateOfflineTaskRequestStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// Specifies whether the task is a draft.
	//
	// example:
	//
	// false
	Draft *bool `json:"draft,omitempty" xml:"draft,omitempty"`
	// Specifies whether to validate the parameters without creating the task.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CreateOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskRequest) GetMeta() *CreateOfflineTaskRequestMeta {
	return s.Meta
}

func (s *CreateOfflineTaskRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateOfflineTaskRequest) GetProcessors() []*CreateOfflineTaskRequestProcessors {
	return s.Processors
}

func (s *CreateOfflineTaskRequest) GetSink() []*CreateOfflineTaskRequestSink {
	return s.Sink
}

func (s *CreateOfflineTaskRequest) GetSource() []*CreateOfflineTaskRequestSource {
	return s.Source
}

func (s *CreateOfflineTaskRequest) GetStatus() *CreateOfflineTaskRequestStatus {
	return s.Status
}

func (s *CreateOfflineTaskRequest) GetDraft() *bool {
	return s.Draft
}

func (s *CreateOfflineTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOfflineTaskRequest) SetMeta(v *CreateOfflineTaskRequestMeta) *CreateOfflineTaskRequest {
	s.Meta = v
	return s
}

func (s *CreateOfflineTaskRequest) SetParameters(v map[string]interface{}) *CreateOfflineTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskRequest) SetProcessors(v []*CreateOfflineTaskRequestProcessors) *CreateOfflineTaskRequest {
	s.Processors = v
	return s
}

func (s *CreateOfflineTaskRequest) SetSink(v []*CreateOfflineTaskRequestSink) *CreateOfflineTaskRequest {
	s.Sink = v
	return s
}

func (s *CreateOfflineTaskRequest) SetSource(v []*CreateOfflineTaskRequestSource) *CreateOfflineTaskRequest {
	s.Source = v
	return s
}

func (s *CreateOfflineTaskRequest) SetStatus(v *CreateOfflineTaskRequestStatus) *CreateOfflineTaskRequest {
	s.Status = v
	return s
}

func (s *CreateOfflineTaskRequest) SetDraft(v bool) *CreateOfflineTaskRequest {
	s.Draft = &v
	return s
}

func (s *CreateOfflineTaskRequest) SetDryRun(v bool) *CreateOfflineTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateOfflineTaskRequest) SetRegionId(v string) *CreateOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOfflineTaskRequest) Validate() error {
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

type CreateOfflineTaskRequestMeta struct {
	// The access credential.
	//
	// example:
	//
	// OS-xxx
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The billing specification.
	//
	// example:
	//
	// small
	ComputeResource *string `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The list of labels.
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The task name.
	//
	// example:
	//
	// test
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s CreateOfflineTaskRequestMeta) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskRequestMeta) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskRequestMeta) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateOfflineTaskRequestMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *CreateOfflineTaskRequestMeta) GetLabels() []*string {
	return s.Labels
}

func (s *CreateOfflineTaskRequestMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOfflineTaskRequestMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateOfflineTaskRequestMeta) SetApiKey(v string) *CreateOfflineTaskRequestMeta {
	s.ApiKey = &v
	return s
}

func (s *CreateOfflineTaskRequestMeta) SetComputeResource(v string) *CreateOfflineTaskRequestMeta {
	s.ComputeResource = &v
	return s
}

func (s *CreateOfflineTaskRequestMeta) SetLabels(v []*string) *CreateOfflineTaskRequestMeta {
	s.Labels = v
	return s
}

func (s *CreateOfflineTaskRequestMeta) SetRegionId(v string) *CreateOfflineTaskRequestMeta {
	s.RegionId = &v
	return s
}

func (s *CreateOfflineTaskRequestMeta) SetTaskName(v string) *CreateOfflineTaskRequestMeta {
	s.TaskName = &v
	return s
}

func (s *CreateOfflineTaskRequestMeta) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskRequestProcessors struct {
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
	// The output parameters.
	//
	// example:
	//
	// {
	//
	// "vector": "syh_image_uri_multi-modal-embedding"
	//
	// }
	Output map[string]interface{} `json:"output,omitempty" xml:"output,omitempty"`
	// The processor parameters.
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

func (s CreateOfflineTaskRequestProcessors) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskRequestProcessors) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskRequestProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *CreateOfflineTaskRequestProcessors) GetName() *string {
	return s.Name
}

func (s *CreateOfflineTaskRequestProcessors) GetOutput() map[string]interface{} {
	return s.Output
}

func (s *CreateOfflineTaskRequestProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateOfflineTaskRequestProcessors) GetType() *string {
	return s.Type
}

func (s *CreateOfflineTaskRequestProcessors) SetInput(v map[string]interface{}) *CreateOfflineTaskRequestProcessors {
	s.Input = v
	return s
}

func (s *CreateOfflineTaskRequestProcessors) SetName(v string) *CreateOfflineTaskRequestProcessors {
	s.Name = &v
	return s
}

func (s *CreateOfflineTaskRequestProcessors) SetOutput(v map[string]interface{}) *CreateOfflineTaskRequestProcessors {
	s.Output = v
	return s
}

func (s *CreateOfflineTaskRequestProcessors) SetParameters(v map[string]interface{}) *CreateOfflineTaskRequestProcessors {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskRequestProcessors) SetType(v string) *CreateOfflineTaskRequestProcessors {
	s.Type = &v
	return s
}

func (s *CreateOfflineTaskRequestProcessors) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskRequestSink struct {
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

func (s CreateOfflineTaskRequestSink) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskRequestSink) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskRequestSink) GetName() *string {
	return s.Name
}

func (s *CreateOfflineTaskRequestSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *CreateOfflineTaskRequestSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *CreateOfflineTaskRequestSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *CreateOfflineTaskRequestSink) GetType() *string {
	return s.Type
}

func (s *CreateOfflineTaskRequestSink) SetName(v string) *CreateOfflineTaskRequestSink {
	s.Name = &v
	return s
}

func (s *CreateOfflineTaskRequestSink) SetParameters(v map[string]*string) *CreateOfflineTaskRequestSink {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskRequestSink) SetPrimaryKey(v string) *CreateOfflineTaskRequestSink {
	s.PrimaryKey = &v
	return s
}

func (s *CreateOfflineTaskRequestSink) SetSchema(v []map[string]*string) *CreateOfflineTaskRequestSink {
	s.Schema = v
	return s
}

func (s *CreateOfflineTaskRequestSink) SetType(v string) *CreateOfflineTaskRequestSink {
	s.Type = &v
	return s
}

func (s *CreateOfflineTaskRequestSink) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskRequestSource struct {
	// **The data source name.**.
	//
	// example:
	//
	// table1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// **The datasource config parameters, which are determined by the type.**.
	//
	// example:
	//
	// {
	//
	//   "key1": "value1",
	//
	//    "key2": "value2"
	//
	// }
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The primary key field of the data source.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// **The data source schema.**.
	Schema []map[string]*string `json:"schema,omitempty" xml:"schema,omitempty" type:"Repeated"`
	// **The data source type.**.
	//
	// example:
	//
	// swift
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateOfflineTaskRequestSource) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskRequestSource) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskRequestSource) GetName() *string {
	return s.Name
}

func (s *CreateOfflineTaskRequestSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *CreateOfflineTaskRequestSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *CreateOfflineTaskRequestSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *CreateOfflineTaskRequestSource) GetType() *string {
	return s.Type
}

func (s *CreateOfflineTaskRequestSource) SetName(v string) *CreateOfflineTaskRequestSource {
	s.Name = &v
	return s
}

func (s *CreateOfflineTaskRequestSource) SetParameters(v map[string]*string) *CreateOfflineTaskRequestSource {
	s.Parameters = v
	return s
}

func (s *CreateOfflineTaskRequestSource) SetPrimaryKey(v string) *CreateOfflineTaskRequestSource {
	s.PrimaryKey = &v
	return s
}

func (s *CreateOfflineTaskRequestSource) SetSchema(v []map[string]*string) *CreateOfflineTaskRequestSource {
	s.Schema = v
	return s
}

func (s *CreateOfflineTaskRequestSource) SetType(v string) *CreateOfflineTaskRequestSource {
	s.Type = &v
	return s
}

func (s *CreateOfflineTaskRequestSource) Validate() error {
	return dara.Validate(s)
}

type CreateOfflineTaskRequestStatus struct {
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
	// The task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateOfflineTaskRequestStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateOfflineTaskRequestStatus) GoString() string {
	return s.String()
}

func (s *CreateOfflineTaskRequestStatus) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateOfflineTaskRequestStatus) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *CreateOfflineTaskRequestStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateOfflineTaskRequestStatus) GetStatus() *string {
	return s.Status
}

func (s *CreateOfflineTaskRequestStatus) SetCreateTime(v int64) *CreateOfflineTaskRequestStatus {
	s.CreateTime = &v
	return s
}

func (s *CreateOfflineTaskRequestStatus) SetDeleteTime(v int64) *CreateOfflineTaskRequestStatus {
	s.DeleteTime = &v
	return s
}

func (s *CreateOfflineTaskRequestStatus) SetErrorMessage(v string) *CreateOfflineTaskRequestStatus {
	s.ErrorMessage = &v
	return s
}

func (s *CreateOfflineTaskRequestStatus) SetStatus(v string) *CreateOfflineTaskRequestStatus {
	s.Status = &v
	return s
}

func (s *CreateOfflineTaskRequestStatus) Validate() error {
	return dara.Validate(s)
}
