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
	SetAutoScale(v bool) *UpdateExecutorGroupRequest
	GetAutoScale() *bool
	SetClusterId(v string) *UpdateExecutorGroupRequest
	GetClusterId() *string
	SetCmsWorkspaceId(v string) *UpdateExecutorGroupRequest
	GetCmsWorkspaceId() *string
	SetDescription(v string) *UpdateExecutorGroupRequest
	GetDescription() *string
	SetId(v string) *UpdateExecutorGroupRequest
	GetId() *string
	SetName(v string) *UpdateExecutorGroupRequest
	GetName() *string
	SetNetwork(v string) *UpdateExecutorGroupRequest
	GetNetwork() *string
	SetProtocol(v string) *UpdateExecutorGroupRequest
	GetProtocol() *string
	SetWorkerType(v string) *UpdateExecutorGroupRequest
	GetWorkerType() *string
	SetWorkers(v string) *UpdateExecutorGroupRequest
	GetWorkers() *string
	SetXAttrs(v string) *UpdateExecutorGroupRequest
	GetXAttrs() *string
}

type UpdateExecutorGroupRequest struct {
	// example:
	//
	// K4KWGINZVY9JwcZT0/vpSg==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// APP
	AuthType  *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	AutoScale *bool   `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-d6a5243b6fa
	ClusterId      *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CmsWorkspaceId *string `json:"CmsWorkspaceId,omitempty" xml:"CmsWorkspaceId,omitempty"`
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	//
	// ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 110176
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	XAttrs  *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
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

func (s *UpdateExecutorGroupRequest) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *UpdateExecutorGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateExecutorGroupRequest) GetCmsWorkspaceId() *string {
	return s.CmsWorkspaceId
}

func (s *UpdateExecutorGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateExecutorGroupRequest) GetId() *string {
	return s.Id
}

func (s *UpdateExecutorGroupRequest) GetName() *string {
	return s.Name
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

func (s *UpdateExecutorGroupRequest) GetXAttrs() *string {
	return s.XAttrs
}

func (s *UpdateExecutorGroupRequest) SetApiKey(v string) *UpdateExecutorGroupRequest {
	s.ApiKey = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetAuthType(v string) *UpdateExecutorGroupRequest {
	s.AuthType = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetAutoScale(v bool) *UpdateExecutorGroupRequest {
	s.AutoScale = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetClusterId(v string) *UpdateExecutorGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateExecutorGroupRequest) SetCmsWorkspaceId(v string) *UpdateExecutorGroupRequest {
	s.CmsWorkspaceId = &v
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

func (s *UpdateExecutorGroupRequest) SetName(v string) *UpdateExecutorGroupRequest {
	s.Name = &v
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

func (s *UpdateExecutorGroupRequest) SetXAttrs(v string) *UpdateExecutorGroupRequest {
	s.XAttrs = &v
	return s
}

func (s *UpdateExecutorGroupRequest) Validate() error {
	return dara.Validate(s)
}
