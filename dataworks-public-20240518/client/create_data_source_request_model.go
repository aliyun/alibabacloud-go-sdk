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
	// The connection configurations of the data source, including the connection address, access identity, and environment information. The envType parameter specifies the environment in which the data source is used. Valid values of the envType parameter:
	//
	// 	- Dev: development environment
	//
	// 	- Prod: production environment
	//
	// The parameters that you need to configure for the data source vary based on the mode in which the data source is added. For more information, see [Data source connection information (ConnectionProperties)](https://help.aliyun.com/document_detail/2852465.html).
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
	// The mode in which you want to add the data source. The mode varies based on the data source type. Valid values for MySQL data sources:
	//
	// 	- InstanceMode: instance mode
	//
	// 	- UrlMode: connection string mode
	//
	// This parameter is required.
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The description of the data source. The description cannot exceed 3,000 characters in length.
	//
	// example:
	//
	// this is a holo datasource
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the data source. The name can be up to 255 characters in length and can contain letters, digits, and underscores (_). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo_holo_datasource
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://dataworks.console.aliyun.com/overview) and go to the workspace management page to obtain the ID.
	//
	// This parameter is used to determine the DataWorks workspaces used for this API call.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the data source. More than 70 types of data sources are supported in DataWorks. For more information, see [Data source types](https://help.aliyun.com/document_detail/2852465.html).
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
