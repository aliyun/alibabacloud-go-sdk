// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *ListEngineConfigsRequest
	GetEnvironment() *string
	SetInstanceId(v string) *ListEngineConfigsRequest
	GetInstanceId() *string
	SetName(v string) *ListEngineConfigsRequest
	GetName() *string
	SetPageNumber(v int32) *ListEngineConfigsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEngineConfigsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListEngineConfigsRequest
	GetStatus() *string
	SetVersion(v string) *ListEngineConfigsRequest
	GetVersion() *string
}

type ListEngineConfigsRequest struct {
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-***test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// latest
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListEngineConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEngineConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListEngineConfigsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEngineConfigsRequest) GetName() *string {
	return s.Name
}

func (s *ListEngineConfigsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEngineConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEngineConfigsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEngineConfigsRequest) GetVersion() *string {
	return s.Version
}

func (s *ListEngineConfigsRequest) SetEnvironment(v string) *ListEngineConfigsRequest {
	s.Environment = &v
	return s
}

func (s *ListEngineConfigsRequest) SetInstanceId(v string) *ListEngineConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListEngineConfigsRequest) SetName(v string) *ListEngineConfigsRequest {
	s.Name = &v
	return s
}

func (s *ListEngineConfigsRequest) SetPageNumber(v int32) *ListEngineConfigsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEngineConfigsRequest) SetPageSize(v int32) *ListEngineConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEngineConfigsRequest) SetStatus(v string) *ListEngineConfigsRequest {
	s.Status = &v
	return s
}

func (s *ListEngineConfigsRequest) SetVersion(v string) *ListEngineConfigsRequest {
	s.Version = &v
	return s
}

func (s *ListEngineConfigsRequest) Validate() error {
	return dara.Validate(s)
}
