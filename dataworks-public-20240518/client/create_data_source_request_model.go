// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionProperties(v string) *CreateDataSourceRequest
	GetConnectionProperties() *string
	SetConnectionPropertiesMode(v string) *CreateDataSourceRequest
	GetConnectionPropertiesMode() *string
	SetDescription(v string) *CreateDataSourceRequest
	GetDescription() *string
	SetName(v string) *CreateDataSourceRequest
	GetName() *string
	SetProjectId(v int64) *CreateDataSourceRequest
	GetProjectId() *int64
	SetType(v string) *CreateDataSourceRequest
	GetType() *string
}

type CreateDataSourceRequest struct {
	// The connection properties of the data source, including the endpoint, access identity, and environment context. The envType property is a member of this object and specifies the data source environment. Valid values:
	//
	// - Dev: development environment.
	//
	// - Prod: production environment.
	//
	// Different data source types have different property specifications under different connection patterns (ConnectionPropertiesMode). For more information, see [Data source connection properties ConnectionProperties](https://help.aliyun.com/document_detail/2852465.html).
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
	// The connection mode of the data source. Different types have different subtypes with different parameter constraints. For example, a MySQL data source supports the following modes:
	//
	// - InstanceMode (instance mode)
	//
	// - UrlMode (connection string mode)
	//
	// This parameter is required.
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the data source. The description can be up to 3,000 characters in length.
	//
	// example:
	//
	// this is a holo datasource
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the data source. The name can contain letters, digits, and underscores (_), and cannot start with a digit or underscore. The name can be up to 255 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo_holo_datasource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/overview) and go to the Workspace Management page to obtain the ID.
	//
	// This parameter specifies the DataWorks workspace for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the data source. More than 70 data source types are supported. For more information about the enumerated data source types, refer to References: [Data source type list](https://help.aliyun.com/document_detail/2852465.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) GetConnectionProperties() *string {
	return s.ConnectionProperties
}

func (s *CreateDataSourceRequest) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *CreateDataSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataSourceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataSourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDataSourceRequest) SetConnectionProperties(v string) *CreateDataSourceRequest {
	s.ConnectionProperties = &v
	return s
}

func (s *CreateDataSourceRequest) SetConnectionPropertiesMode(v string) *CreateDataSourceRequest {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetProjectId(v int64) *CreateDataSourceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
