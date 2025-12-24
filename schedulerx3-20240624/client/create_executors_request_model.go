// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateExecutorsRequest
	GetAppName() *string
	SetClusterId(v string) *CreateExecutorsRequest
	GetClusterId() *string
	SetWorkerType(v string) *CreateExecutorsRequest
	GetWorkerType() *string
	SetWorkers(v string) *CreateExecutorsRequest
	GetWorkers() *string
}

type CreateExecutorsRequest struct {
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
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// k8s_service
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"cluster":"xxxxx","namespace":"xxxxx","service":"xxxxx"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s CreateExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorsRequest) GoString() string {
	return s.String()
}

func (s *CreateExecutorsRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateExecutorsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateExecutorsRequest) GetWorkerType() *string {
	return s.WorkerType
}

func (s *CreateExecutorsRequest) GetWorkers() *string {
	return s.Workers
}

func (s *CreateExecutorsRequest) SetAppName(v string) *CreateExecutorsRequest {
	s.AppName = &v
	return s
}

func (s *CreateExecutorsRequest) SetClusterId(v string) *CreateExecutorsRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateExecutorsRequest) SetWorkerType(v string) *CreateExecutorsRequest {
	s.WorkerType = &v
	return s
}

func (s *CreateExecutorsRequest) SetWorkers(v string) *CreateExecutorsRequest {
	s.Workers = &v
	return s
}

func (s *CreateExecutorsRequest) Validate() error {
	return dara.Validate(s)
}
