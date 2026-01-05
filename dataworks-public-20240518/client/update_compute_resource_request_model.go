// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionProperties(v string) *UpdateComputeResourceRequest
	GetConnectionProperties() *string
	SetConnectionPropertiesMode(v string) *UpdateComputeResourceRequest
	GetConnectionPropertiesMode() *string
	SetDescription(v string) *UpdateComputeResourceRequest
	GetDescription() *string
	SetId(v int64) *UpdateComputeResourceRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateComputeResourceRequest
	GetProjectId() *int64
}

type UpdateComputeResourceRequest struct {
	// The specific connection configuration of the computing resource, including the connection address, access identity, and environment information. The environment type (EnvType) of the computing resource is a member attribute of this object, including DEV (development environment) and PROD (production environment). The value is not case-sensitive.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//     "envType": "Prod",
	//
	//     "regionId": "cn-beijing",
	//
	//     "instanceId": "hgprecn-cn-x0r3oun4k001",
	//
	//     "database": "testdb",
	//
	//     "securityProtocol": "authTypeNone",
	//
	//     "authType": "Executor",
	//
	//     "authIdentity": "1107550004253538"
	//
	// }
	ConnectionProperties *string `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The category of the computing resource to be added. Different types have different subtypes and corresponding parameter schema constraints. Examples: InstanceMode and UrlMode.
	//
	// example:
	//
	// InstanceMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the computing resource. The maximum length is 3000 characters.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the computing resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateComputeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceRequest) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *UpdateComputeResourceRequest) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *UpdateComputeResourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateComputeResourceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateComputeResourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateComputeResourceRequest) SetConnectionProperties(v string) *UpdateComputeResourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetConnectionPropertiesMode(v string) *UpdateComputeResourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetDescription(v string) *UpdateComputeResourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetId(v int64) *UpdateComputeResourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateComputeResourceRequest) SetProjectId(v int64) *UpdateComputeResourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateComputeResourceRequest) Validate() error {
	return dara.Validate(s)
}
