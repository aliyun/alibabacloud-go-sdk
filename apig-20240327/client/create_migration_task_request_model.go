// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateMigrationTaskRequest
	GetClusterId() *string
	SetDescription(v string) *CreateMigrationTaskRequest
	GetDescription() *string
	SetEnvironmentId(v string) *CreateMigrationTaskRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *CreateMigrationTaskRequest
	GetGatewayId() *string
	SetHttpApiId(v string) *CreateMigrationTaskRequest
	GetHttpApiId() *string
	SetIngressClass(v string) *CreateMigrationTaskRequest
	GetIngressClass() *string
	SetMigrationType(v string) *CreateMigrationTaskRequest
	GetMigrationType() *string
	SetWatchNamespace(v string) *CreateMigrationTaskRequest
	GetWatchNamespace() *string
}

type CreateMigrationTaskRequest struct {
	// example:
	//
	// c-xxxxxx
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// example:
	//
	// migration from Nginx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// env-xxxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// api-xxxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// example:
	//
	// nginx
	IngressClass *string `json:"ingressClass,omitempty" xml:"ingressClass,omitempty"`
	// example:
	//
	// Nginx Ingress
	MigrationType *string `json:"migrationType,omitempty" xml:"migrationType,omitempty"`
	// example:
	//
	// default
	WatchNamespace *string `json:"watchNamespace,omitempty" xml:"watchNamespace,omitempty"`
}

func (s CreateMigrationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrationTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateMigrationTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMigrationTaskRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateMigrationTaskRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateMigrationTaskRequest) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *CreateMigrationTaskRequest) GetIngressClass() *string {
	return s.IngressClass
}

func (s *CreateMigrationTaskRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *CreateMigrationTaskRequest) GetWatchNamespace() *string {
	return s.WatchNamespace
}

func (s *CreateMigrationTaskRequest) SetClusterId(v string) *CreateMigrationTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetDescription(v string) *CreateMigrationTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetEnvironmentId(v string) *CreateMigrationTaskRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetGatewayId(v string) *CreateMigrationTaskRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetHttpApiId(v string) *CreateMigrationTaskRequest {
	s.HttpApiId = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetIngressClass(v string) *CreateMigrationTaskRequest {
	s.IngressClass = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetMigrationType(v string) *CreateMigrationTaskRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateMigrationTaskRequest) SetWatchNamespace(v string) *CreateMigrationTaskRequest {
	s.WatchNamespace = &v
	return s
}

func (s *CreateMigrationTaskRequest) Validate() error {
	return dara.Validate(s)
}
