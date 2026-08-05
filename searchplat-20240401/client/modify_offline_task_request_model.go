// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMeta(v *ModifyOfflineTaskRequestMeta) *ModifyOfflineTaskRequest
	GetMeta() *ModifyOfflineTaskRequestMeta
	SetParameters(v map[string]interface{}) *ModifyOfflineTaskRequest
	GetParameters() map[string]interface{}
	SetProcessors(v []*ModifyOfflineTaskRequestProcessors) *ModifyOfflineTaskRequest
	GetProcessors() []*ModifyOfflineTaskRequestProcessors
	SetSink(v []*ModifyOfflineTaskRequestSink) *ModifyOfflineTaskRequest
	GetSink() []*ModifyOfflineTaskRequestSink
	SetSource(v []*ModifyOfflineTaskRequestSource) *ModifyOfflineTaskRequest
	GetSource() []*ModifyOfflineTaskRequestSource
	SetStatus(v *ModifyOfflineTaskRequestStatus) *ModifyOfflineTaskRequest
	GetStatus() *ModifyOfflineTaskRequestStatus
	SetDryRun(v bool) *ModifyOfflineTaskRequest
	GetDryRun() *bool
	SetRegionId(v string) *ModifyOfflineTaskRequest
	GetRegionId() *string
}

type ModifyOfflineTaskRequest struct {
	// The task metadata.
	Meta *ModifyOfflineTaskRequestMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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
	// The processing flow operators.
	Processors []*ModifyOfflineTaskRequestProcessors `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The data sink information.
	Sink []*ModifyOfflineTaskRequestSink `json:"sink,omitempty" xml:"sink,omitempty" type:"Repeated"`
	// The data source information.
	Source []*ModifyOfflineTaskRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Repeated"`
	// The task status.
	Status *ModifyOfflineTaskRequestStatus `json:"status,omitempty" xml:"status,omitempty" type:"Struct"`
	// Specifies whether to validate the request parameters without performing the actual operation. Default value: false.
	//
	// Valid values:
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ModifyOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskRequest) GetMeta() *ModifyOfflineTaskRequestMeta {
	return s.Meta
}

func (s *ModifyOfflineTaskRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ModifyOfflineTaskRequest) GetProcessors() []*ModifyOfflineTaskRequestProcessors {
	return s.Processors
}

func (s *ModifyOfflineTaskRequest) GetSink() []*ModifyOfflineTaskRequestSink {
	return s.Sink
}

func (s *ModifyOfflineTaskRequest) GetSource() []*ModifyOfflineTaskRequestSource {
	return s.Source
}

func (s *ModifyOfflineTaskRequest) GetStatus() *ModifyOfflineTaskRequestStatus {
	return s.Status
}

func (s *ModifyOfflineTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfflineTaskRequest) SetMeta(v *ModifyOfflineTaskRequestMeta) *ModifyOfflineTaskRequest {
	s.Meta = v
	return s
}

func (s *ModifyOfflineTaskRequest) SetParameters(v map[string]interface{}) *ModifyOfflineTaskRequest {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskRequest) SetProcessors(v []*ModifyOfflineTaskRequestProcessors) *ModifyOfflineTaskRequest {
	s.Processors = v
	return s
}

func (s *ModifyOfflineTaskRequest) SetSink(v []*ModifyOfflineTaskRequestSink) *ModifyOfflineTaskRequest {
	s.Sink = v
	return s
}

func (s *ModifyOfflineTaskRequest) SetSource(v []*ModifyOfflineTaskRequestSource) *ModifyOfflineTaskRequest {
	s.Source = v
	return s
}

func (s *ModifyOfflineTaskRequest) SetStatus(v *ModifyOfflineTaskRequestStatus) *ModifyOfflineTaskRequest {
	s.Status = v
	return s
}

func (s *ModifyOfflineTaskRequest) SetDryRun(v bool) *ModifyOfflineTaskRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyOfflineTaskRequest) SetRegionId(v string) *ModifyOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfflineTaskRequest) Validate() error {
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

type ModifyOfflineTaskRequestMeta struct {
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
	// The list of task labels.
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
	// syh
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
}

func (s ModifyOfflineTaskRequestMeta) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskRequestMeta) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskRequestMeta) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModifyOfflineTaskRequestMeta) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *ModifyOfflineTaskRequestMeta) GetLabels() []*string {
	return s.Labels
}

func (s *ModifyOfflineTaskRequestMeta) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfflineTaskRequestMeta) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyOfflineTaskRequestMeta) SetApiKey(v string) *ModifyOfflineTaskRequestMeta {
	s.ApiKey = &v
	return s
}

func (s *ModifyOfflineTaskRequestMeta) SetComputeResource(v string) *ModifyOfflineTaskRequestMeta {
	s.ComputeResource = &v
	return s
}

func (s *ModifyOfflineTaskRequestMeta) SetLabels(v []*string) *ModifyOfflineTaskRequestMeta {
	s.Labels = v
	return s
}

func (s *ModifyOfflineTaskRequestMeta) SetRegionId(v string) *ModifyOfflineTaskRequestMeta {
	s.RegionId = &v
	return s
}

func (s *ModifyOfflineTaskRequestMeta) SetTaskName(v string) *ModifyOfflineTaskRequestMeta {
	s.TaskName = &v
	return s
}

func (s *ModifyOfflineTaskRequestMeta) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskRequestProcessors struct {
	// The input parameters.
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

func (s ModifyOfflineTaskRequestProcessors) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskRequestProcessors) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskRequestProcessors) GetInput() map[string]interface{} {
	return s.Input
}

func (s *ModifyOfflineTaskRequestProcessors) GetName() *string {
	return s.Name
}

func (s *ModifyOfflineTaskRequestProcessors) GetOutput() map[string]interface{} {
	return s.Output
}

func (s *ModifyOfflineTaskRequestProcessors) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ModifyOfflineTaskRequestProcessors) GetType() *string {
	return s.Type
}

func (s *ModifyOfflineTaskRequestProcessors) SetInput(v map[string]interface{}) *ModifyOfflineTaskRequestProcessors {
	s.Input = v
	return s
}

func (s *ModifyOfflineTaskRequestProcessors) SetName(v string) *ModifyOfflineTaskRequestProcessors {
	s.Name = &v
	return s
}

func (s *ModifyOfflineTaskRequestProcessors) SetOutput(v map[string]interface{}) *ModifyOfflineTaskRequestProcessors {
	s.Output = v
	return s
}

func (s *ModifyOfflineTaskRequestProcessors) SetParameters(v map[string]interface{}) *ModifyOfflineTaskRequestProcessors {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskRequestProcessors) SetType(v string) *ModifyOfflineTaskRequestProcessors {
	s.Type = &v
	return s
}

func (s *ModifyOfflineTaskRequestProcessors) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskRequestSink struct {
	// The data sink name.
	//
	// example:
	//
	// default
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

func (s ModifyOfflineTaskRequestSink) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskRequestSink) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskRequestSink) GetName() *string {
	return s.Name
}

func (s *ModifyOfflineTaskRequestSink) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ModifyOfflineTaskRequestSink) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ModifyOfflineTaskRequestSink) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *ModifyOfflineTaskRequestSink) GetType() *string {
	return s.Type
}

func (s *ModifyOfflineTaskRequestSink) SetName(v string) *ModifyOfflineTaskRequestSink {
	s.Name = &v
	return s
}

func (s *ModifyOfflineTaskRequestSink) SetParameters(v map[string]*string) *ModifyOfflineTaskRequestSink {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskRequestSink) SetPrimaryKey(v string) *ModifyOfflineTaskRequestSink {
	s.PrimaryKey = &v
	return s
}

func (s *ModifyOfflineTaskRequestSink) SetSchema(v []map[string]*string) *ModifyOfflineTaskRequestSink {
	s.Schema = v
	return s
}

func (s *ModifyOfflineTaskRequestSink) SetType(v string) *ModifyOfflineTaskRequestSink {
	s.Type = &v
	return s
}

func (s *ModifyOfflineTaskRequestSink) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskRequestSource struct {
	// The data source name.
	//
	// example:
	//
	// tmp-exec-pop-eTcMpC
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

func (s ModifyOfflineTaskRequestSource) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskRequestSource) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskRequestSource) GetName() *string {
	return s.Name
}

func (s *ModifyOfflineTaskRequestSource) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *ModifyOfflineTaskRequestSource) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ModifyOfflineTaskRequestSource) GetSchema() []map[string]*string {
	return s.Schema
}

func (s *ModifyOfflineTaskRequestSource) GetType() *string {
	return s.Type
}

func (s *ModifyOfflineTaskRequestSource) SetName(v string) *ModifyOfflineTaskRequestSource {
	s.Name = &v
	return s
}

func (s *ModifyOfflineTaskRequestSource) SetParameters(v map[string]*string) *ModifyOfflineTaskRequestSource {
	s.Parameters = v
	return s
}

func (s *ModifyOfflineTaskRequestSource) SetPrimaryKey(v string) *ModifyOfflineTaskRequestSource {
	s.PrimaryKey = &v
	return s
}

func (s *ModifyOfflineTaskRequestSource) SetSchema(v []map[string]*string) *ModifyOfflineTaskRequestSource {
	s.Schema = v
	return s
}

func (s *ModifyOfflineTaskRequestSource) SetType(v string) *ModifyOfflineTaskRequestSource {
	s.Type = &v
	return s
}

func (s *ModifyOfflineTaskRequestSource) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskRequestStatus struct {
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
	// “”
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The task status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModifyOfflineTaskRequestStatus) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskRequestStatus) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskRequestStatus) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ModifyOfflineTaskRequestStatus) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *ModifyOfflineTaskRequestStatus) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyOfflineTaskRequestStatus) GetStatus() *string {
	return s.Status
}

func (s *ModifyOfflineTaskRequestStatus) SetCreateTime(v int64) *ModifyOfflineTaskRequestStatus {
	s.CreateTime = &v
	return s
}

func (s *ModifyOfflineTaskRequestStatus) SetDeleteTime(v int64) *ModifyOfflineTaskRequestStatus {
	s.DeleteTime = &v
	return s
}

func (s *ModifyOfflineTaskRequestStatus) SetErrorMessage(v string) *ModifyOfflineTaskRequestStatus {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyOfflineTaskRequestStatus) SetStatus(v string) *ModifyOfflineTaskRequestStatus {
	s.Status = &v
	return s
}

func (s *ModifyOfflineTaskRequestStatus) Validate() error {
	return dara.Validate(s)
}
