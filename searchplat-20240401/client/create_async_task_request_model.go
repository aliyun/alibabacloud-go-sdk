// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAsyncTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v int64) *CreateAsyncTaskRequest
	GetDataId() *int64
	SetId(v string) *CreateAsyncTaskRequest
	GetId() *string
	SetName(v string) *CreateAsyncTaskRequest
	GetName() *string
	SetServiceId(v string) *CreateAsyncTaskRequest
	GetServiceId() *string
	SetServiceType(v string) *CreateAsyncTaskRequest
	GetServiceType() *string
	SetDryRun(v bool) *CreateAsyncTaskRequest
	GetDryRun() *bool
}

type CreateAsyncTaskRequest struct {
	// The playground data ID.
	//
	// example:
	//
	// 12323
	DataId *int64 `json:"dataId,omitempty" xml:"dataId,omitempty"`
	// The asynchronous task ID.
	//
	// example:
	//
	// fae9bcc5-949f-4c31-b9b7-a273bf891699
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The task name.
	//
	// example:
	//
	// 文档解析任务
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The service ID.
	//
	// example:
	//
	// ops-document-analyze-001
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The service type.
	//
	// example:
	//
	// document-analyze
	ServiceType *string `json:"serviceType,omitempty" xml:"serviceType,omitempty"`
	// Specifies whether to perform a dry run request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAsyncTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateAsyncTaskRequest) GetDataId() *int64 {
	return s.DataId
}

func (s *CreateAsyncTaskRequest) GetId() *string {
	return s.Id
}

func (s *CreateAsyncTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateAsyncTaskRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateAsyncTaskRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *CreateAsyncTaskRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateAsyncTaskRequest) SetDataId(v int64) *CreateAsyncTaskRequest {
	s.DataId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetId(v string) *CreateAsyncTaskRequest {
	s.Id = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetName(v string) *CreateAsyncTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetServiceId(v string) *CreateAsyncTaskRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetServiceType(v string) *CreateAsyncTaskRequest {
	s.ServiceType = &v
	return s
}

func (s *CreateAsyncTaskRequest) SetDryRun(v bool) *CreateAsyncTaskRequest {
	s.DryRun = &v
	return s
}

func (s *CreateAsyncTaskRequest) Validate() error {
	return dara.Validate(s)
}
