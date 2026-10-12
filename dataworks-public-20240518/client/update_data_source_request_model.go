// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionProperties(v string) *UpdateDataSourceRequest
	GetConnectionProperties() *string
	SetConnectionPropertiesMode(v string) *UpdateDataSourceRequest
	GetConnectionPropertiesMode() *string
	SetDescription(v string) *UpdateDataSourceRequest
	GetDescription() *string
	SetId(v int64) *UpdateDataSourceRequest
	GetId() *int64
	SetProjectId(v int64) *UpdateDataSourceRequest
	GetProjectId() *int64
}

type UpdateDataSourceRequest struct {
	// The specific connection configuration of the data source, including the endpoint, access identity, and environment context. The EnvType property of the data source environment is a member attribute of this object. Valid values:
	//
	// - Dev: development environment.
	//
	// - Prod: production environment.
	//
	// Different types of data sources have different attribute specifications in different configuration patterns (ConnectionPropertiesMode). For more information, see [Connection properties](https://help.aliyun.com/document_detail/2852465.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"envType": "Prod",
	//
	// 	"regionId": "cn-beijing",
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
	// The category used to add the data source. Different types have different subtypes and corresponding parameter constraints. Examples:
	//
	// - InstanceMode: instance mode.
	//
	// - UrlMode: connection string mode.
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the data source. The description can be up to 3,000 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5678
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSourceRequest) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *UpdateDataSourceRequest) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *UpdateDataSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataSourceRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataSourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataSourceRequest) SetConnectionProperties(v string) *UpdateDataSourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *UpdateDataSourceRequest) SetConnectionPropertiesMode(v string) *UpdateDataSourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *UpdateDataSourceRequest) SetDescription(v string) *UpdateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataSourceRequest) SetId(v int64) *UpdateDataSourceRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataSourceRequest) SetProjectId(v int64) *UpdateDataSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
