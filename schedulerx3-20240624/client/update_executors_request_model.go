// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateExecutorsRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateExecutorsRequest
	GetClusterId() *string
	SetWorkerType(v string) *UpdateExecutorsRequest
	GetWorkerType() *string
	SetWorkers(v string) *UpdateExecutorsRequest
	GetWorkers() *string
}

type UpdateExecutorsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s UpdateExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorsRequest) GoString() string {
	return s.String()
}

func (s *UpdateExecutorsRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateExecutorsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateExecutorsRequest) GetWorkerType() *string {
	return s.WorkerType
}

func (s *UpdateExecutorsRequest) GetWorkers() *string {
	return s.Workers
}

func (s *UpdateExecutorsRequest) SetAppName(v string) *UpdateExecutorsRequest {
	s.AppName = &v
	return s
}

func (s *UpdateExecutorsRequest) SetClusterId(v string) *UpdateExecutorsRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateExecutorsRequest) SetWorkerType(v string) *UpdateExecutorsRequest {
	s.WorkerType = &v
	return s
}

func (s *UpdateExecutorsRequest) SetWorkers(v string) *UpdateExecutorsRequest {
	s.Workers = &v
	return s
}

func (s *UpdateExecutorsRequest) Validate() error {
	return dara.Validate(s)
}
