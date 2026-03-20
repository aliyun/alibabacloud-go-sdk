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
	// example:
	//
	// WorkBench
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
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
