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
	// The specific connection configuration information of the compute resource, including the endpoint, access identity, and environment context. The EnvType property of the compute resource environment is a member attribute of this object. Valid values include DEV (development environment) and PROD (production environment). The value is case-insensitive.
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
	// The category of the added compute resource. Different types have different subtypes, which correspond to different parameter schema constraints. Examples: InstanceMode for instance mode and UrlMode for connection string mode.
	//
	// example:
	//
	// InstanceMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the compute resource. The description can be up to 3,000 characters in length.
	//
	// example:
	//
	// Hierarchical description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The compute resource ID.
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
