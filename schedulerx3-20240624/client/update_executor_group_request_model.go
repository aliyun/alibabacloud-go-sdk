// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *UpdateExecutorGroupRequest
	GetApiKey() *string
	SetAuthType(v string) *UpdateExecutorGroupRequest
	GetAuthType() *string
	SetClusterId(v string) *UpdateExecutorGroupRequest
	GetClusterId() *string
	SetDescription(v string) *UpdateExecutorGroupRequest
	GetDescription() *string
	SetId(v string) *UpdateExecutorGroupRequest
	GetId() *string
	SetNetwork(v string) *UpdateExecutorGroupRequest
	GetNetwork() *string
	SetProtocol(v string) *UpdateExecutorGroupRequest
	GetProtocol() *string
	SetWorkerType(v string) *UpdateExecutorGroupRequest
	GetWorkerType() *string
	SetWorkers(v string) *UpdateExecutorGroupRequest
	GetWorkers() *string
}

type UpdateExecutorGroupRequest struct {
	// example:
	//
	// K4KWGINZVY9JwcZT0/vpSg==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// APP
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 110176
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// public
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// openclaw
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// example:
	//
	// [{"address":"http://47.111.188.191:18789"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
}

func (s UpdateExecutorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateExecutorGroupRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *UpdateExecutorGroupRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *UpdateExecutorGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateExecutorGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExecutorGroupRequest) GetId() *string {
	return s.Id
}

func (s *UpdateExecutorGroupRequest) GetNetwork() *string {
	return s.Network
}

func (s *UpdateExecutorGroupRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateExecutorGroupRequest) GetWorkerType() *string {
	return s.WorkerType
}

func (s *UpdateExecutorGroupRequest) GetWorkers() *string {
	return s.Workers
}

func (s *UpdateExecutorGroupRequest) SetApiKey(v string) *UpdateExecutorGroupRequest {
	s.ApiKey = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetAuthType(v string) *UpdateExecutorGroupRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetClusterId(v string) *UpdateExecutorGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetDescription(v string) *UpdateExecutorGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetId(v string) *UpdateExecutorGroupRequest {
	s.Id = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetNetwork(v string) *UpdateExecutorGroupRequest {
	s.Network = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetProtocol(v string) *UpdateExecutorGroupRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetWorkerType(v string) *UpdateExecutorGroupRequest {
	s.WorkerType = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetWorkers(v string) *UpdateExecutorGroupRequest {
	s.Workers = &v
	return s
}

func (s *UpdateExecutorGroupRequest) Validate() error {
	return dara.Validate(s)
}
