// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecutorGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateExecutorGroupRequest
	GetApiKey() *string
	SetAuthType(v string) *CreateExecutorGroupRequest
	GetAuthType() *string
	SetAutoScale(v bool) *CreateExecutorGroupRequest
	GetAutoScale() *bool
	SetClusterId(v string) *CreateExecutorGroupRequest
	GetClusterId() *string
	SetCmsWorkspaceId(v string) *CreateExecutorGroupRequest
	GetCmsWorkspaceId() *string
	SetDescription(v string) *CreateExecutorGroupRequest
	GetDescription() *string
	SetIntegrationType(v string) *CreateExecutorGroupRequest
	GetIntegrationType() *string
	SetName(v string) *CreateExecutorGroupRequest
	GetName() *string
	SetNetwork(v string) *CreateExecutorGroupRequest
	GetNetwork() *string
	SetProtocol(v string) *CreateExecutorGroupRequest
	GetProtocol() *string
	SetWorkerType(v string) *CreateExecutorGroupRequest
	GetWorkerType() *string
	SetWorkers(v string) *CreateExecutorGroupRequest
	GetWorkers() *string
	SetXAttrs(v string) *CreateExecutorGroupRequest
	GetXAttrs() *string
}

type CreateExecutorGroupRequest struct {
	// The client secret.
	//
	// example:
	//
	// HfGUG/Qf8qRCCWlv5RT6WA==
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The authentication type.
	//
	// example:
	//
	// api_key
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The auto scaling configuration.
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The CloudMonitor Service (CMS) workspace ID.
	CmsWorkspaceId *string `json:"CmsWorkspaceId,omitempty" xml:"CmsWorkspaceId,omitempty"`
	// The description of the worker group.
	//
	// example:
	//
	// group1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The integration method of the executor. Valid values: manual (indicates manual entry) and acs (indicates ACS integration).
	//
	// example:
	//
	// manual
	IntegrationType *string `json:"IntegrationType,omitempty" xml:"IntegrationType,omitempty"`
	// The name of the worker group.
	//
	// This parameter is required.
	//
	// example:
	//
	// group1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The network type. Valid values: public and private.
	//
	// example:
	//
	// public
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// The backend protocol. Valid values:
	//
	// - **HTTP*	- (default): Supports association with HTTPS, HTTP, and QUIC listeners.
	//
	// - **HTTPS**: Supports association with HTTPS listeners.
	//
	// - **gRPC**: Supports association with HTTPS and QUIC listeners.
	//
	// > If **ServerGroupType*	- is set to **Fc**, you do not need to configure the backend protocol.
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The worker type.
	//
	// This parameter is required.
	//
	// example:
	//
	// openclaw
	WorkerType *string `json:"WorkerType,omitempty" xml:"WorkerType,omitempty"`
	// The list of workers in JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"address":"http://47.111.188.191:18789"}]
	Workers *string `json:"Workers,omitempty" xml:"Workers,omitempty"`
	// The extended configuration of the executor group.
	//
	// example:
	//
	// Not publicly available
	XAttrs *string `json:"XAttrs,omitempty" xml:"XAttrs,omitempty"`
}

func (s CreateExecutorGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExecutorGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateExecutorGroupRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateExecutorGroupRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateExecutorGroupRequest) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *CreateExecutorGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateExecutorGroupRequest) GetCmsWorkspaceId() *string {
	return s.CmsWorkspaceId
}

func (s *CreateExecutorGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExecutorGroupRequest) GetIntegrationType() *string {
	return s.IntegrationType
}

func (s *CreateExecutorGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateExecutorGroupRequest) GetNetwork() *string {
	return s.Network
}

func (s *CreateExecutorGroupRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateExecutorGroupRequest) GetWorkerType() *string {
	return s.WorkerType
}

func (s *CreateExecutorGroupRequest) GetWorkers() *string {
	return s.Workers
}

func (s *CreateExecutorGroupRequest) GetXAttrs() *string {
	return s.XAttrs
}

func (s *CreateExecutorGroupRequest) SetApiKey(v string) *CreateExecutorGroupRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetAuthType(v string) *CreateExecutorGroupRequest {
	s.AuthType = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetAutoScale(v bool) *CreateExecutorGroupRequest {
	s.AutoScale = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetClusterId(v string) *CreateExecutorGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetCmsWorkspaceId(v string) *CreateExecutorGroupRequest {
	s.CmsWorkspaceId = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetDescription(v string) *CreateExecutorGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetIntegrationType(v string) *CreateExecutorGroupRequest {
	s.IntegrationType = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetName(v string) *CreateExecutorGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetNetwork(v string) *CreateExecutorGroupRequest {
	s.Network = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetProtocol(v string) *CreateExecutorGroupRequest {
	s.Protocol = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetWorkerType(v string) *CreateExecutorGroupRequest {
	s.WorkerType = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetWorkers(v string) *CreateExecutorGroupRequest {
	s.Workers = &v
	return s
}

func (s *CreateExecutorGroupRequest) SetXAttrs(v string) *CreateExecutorGroupRequest {
	s.XAttrs = &v
	return s
}

func (s *CreateExecutorGroupRequest) Validate() error {
	return dara.Validate(s)
}
