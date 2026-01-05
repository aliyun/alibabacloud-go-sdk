// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateComputeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionProperties(v string) *CreateComputeResourceRequest
	GetConnectionProperties() *string
	SetConnectionPropertiesMode(v string) *CreateComputeResourceRequest
	GetConnectionPropertiesMode() *string
	SetDescription(v string) *CreateComputeResourceRequest
	GetDescription() *string
	SetName(v string) *CreateComputeResourceRequest
	GetName() *string
	SetProjectId(v int64) *CreateComputeResourceRequest
	GetProjectId() *int64
	SetType(v string) *CreateComputeResourceRequest
	GetType() *string
}

type CreateComputeResourceRequest struct {
	// The specific connection configuration of the computing resource, including the connection address, access identity, and environment information. The environment for the computing resource is a member attribute of this object. It includes DEV (development environment) and PROD (production environment). The value for EnvType is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "EndpointMode": "custom",     "Database": "testdb",     "TaskSubmitter": "1107550004253538",     "InstanceId": "hgprecn-cn-x0r3oun4k001",     "SecurityProtocol": "authTypeNone",     "RegionId": "cn-beijing",     "EnvType": "Prod",     "AuthType": "Executor" }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// Depending on the added source category, different types may have specific subtypes subject to distinct parameter constraints. For example, Hologres computing resources support InstanceMode and UrlMode.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the computing resource. The maximum length is 3000 characters.
	//
	// example:
	//
	// demo_holo_cs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the computing resource. It can contain letters, digits, and underscores (_), cannot start with a digit or an underscore, and has a maximum length of 255 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo_holo_cs
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the computing resource. Multiple computing resource types are currently supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateComputeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeResourceRequest) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *CreateComputeResourceRequest) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *CreateComputeResourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateComputeResourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateComputeResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateComputeResourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateComputeResourceRequest) SetConnectionProperties(v string) *CreateComputeResourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *CreateComputeResourceRequest) SetConnectionPropertiesMode(v string) *CreateComputeResourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *CreateComputeResourceRequest) SetDescription(v string) *CreateComputeResourceRequest {
	s.Description = &v
	return s
}

func (s *CreateComputeResourceRequest) SetName(v string) *CreateComputeResourceRequest {
	s.Name = &v
	return s
}

func (s *CreateComputeResourceRequest) SetProjectId(v int64) *CreateComputeResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateComputeResourceRequest) SetType(v string) *CreateComputeResourceRequest {
	s.Type = &v
	return s
}

func (s *CreateComputeResourceRequest) Validate() error {
	return dara.Validate(s)
}
