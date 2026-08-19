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
	// The connection configuration of the compute resource, including the endpoint, access identity, and environment context. The EnvType field is a member property of this object and specifies the environment of the compute resource. Valid values: DEV (development environment) and PROD (production environment). The EnvType value is case-insensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "EndpointMode": "custom",     "Database": "testdb",     "TaskSubmitter": "1107550004253538",     "InstanceId": "hgprecn-cn-x0r3oun4k001",     "SecurityProtocol": "authTypeNone",     "RegionId": "cn-beijing",     "EnvType": "Prod",     "AuthType": "Executor" }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The category for adding the compute resource. Different types have different subtypes with different parameter constraints. For example, a Hologres compute resource supports InstanceMode (instance mode) and UrlMode (connection string mode).
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the compute resource. The description can be up to 3,000 characters in length.
	//
	// example:
	//
	// demo_holo_cs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the compute resource. The name can contain letters, digits, and underscores (_), and cannot start with a digit or underscore. The name can be up to 255 characters in length.
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
	// The type of the compute resource. Multiple compute resource types are supported.
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
