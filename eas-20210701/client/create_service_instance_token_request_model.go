// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *CreateServiceInstanceTokenRequest
	GetActionType() *string
	SetWorkerName(v string) *CreateServiceInstanceTokenRequest
	GetWorkerName() *string
}

type CreateServiceInstanceTokenRequest struct {
	// The type of URL to retrieve. Valid values:
	//
	// - **WorkBench**: Logs on to the container through the Workbench page.
	//
	// - **WebShell**: Logs on to the container through a WebSocket URL. This value is not supported for processor-type services or services created before June 2026.
	//
	// example:
	//
	// WorkBench
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The name of the service worker. You can call the ListServiceContainers operation to obtain the worker name.
	//
	// example:
	//
	// worker0
	WorkerName *string `json:"WorkerName,omitempty" xml:"WorkerName,omitempty"`
}

func (s CreateServiceInstanceTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceTokenRequest) GetActionType() *string {
	return s.ActionType
}

func (s *CreateServiceInstanceTokenRequest) GetWorkerName() *string {
	return s.WorkerName
}

func (s *CreateServiceInstanceTokenRequest) SetActionType(v string) *CreateServiceInstanceTokenRequest {
	s.ActionType = &v
	return s
}

func (s *CreateServiceInstanceTokenRequest) SetWorkerName(v string) *CreateServiceInstanceTokenRequest {
	s.WorkerName = &v
	return s
}

func (s *CreateServiceInstanceTokenRequest) Validate() error {
	return dara.Validate(s)
}
