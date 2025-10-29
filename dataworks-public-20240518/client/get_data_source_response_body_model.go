// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSource(v *GetDataSourceResponseBodyDataSource) *GetDataSourceResponseBody
	GetDataSource() *GetDataSourceResponseBodyDataSource
	SetRequestId(v string) *GetDataSourceResponseBody
	GetRequestId() *string
}

type GetDataSourceResponseBody struct {
	// The information about the data source.
	DataSource *GetDataSourceResponseBodyDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9252F32F-D855-549E-8898-61CF5A733050
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) GetDataSource() *GetDataSourceResponseBodyDataSource {
	return s.DataSource
}

func (s *GetDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceResponseBody) SetDataSource(v *GetDataSourceResponseBodyDataSource) *GetDataSourceResponseBody {
	s.DataSource = v
	return s
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceResponseBody) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataSourceResponseBodyDataSource struct {
	// The connection configurations of the data source, including the connection address, access identity, and environment information. The envType parameter specifies the environment in which the data source is used. Valid values of the envType parameter:
	//
	// 	- Dev: development environment
	//
	// 	- Prod: production environment
	//
	// The parameters that you need to configure for the data source vary based on the mode in which the data source is added. For more information, see [Data source connection information (ConnectionProperties)](https://help.aliyun.com/document_detail/2852465.html).
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
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The mode in which the data source is added. The mode varies based on the data source type. Valid values:
	//
	// 	- InstanceMode: instance mode
	//
	// 	- UrlMode: connection string mode
	//
	// 	- CdhMode: CDH cluster mode
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The time when the data source was added. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698286929333
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who adds the data source.
	//
	// example:
	//
	// 1107550004253538
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 16738
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the data source was last modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1698286929333
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The ID of the user who modifies the data source.
	//
	// example:
	//
	// 1107550004253538
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the workspace with which the data source is associated.
	//
	// example:
	//
	// 52660
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The unique business key of the data source. For example, the unique business key of a Hologres data source is in the `${tenantOwnerId}:${regionId}:${type}:${instanceId}:${database}` format.
	//
	// example:
	//
	// 1107550004253538:cn-beijing:holo:hgprecn-cn-x0r3oun4k001:testdb
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataSourceResponseBodyDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceResponseBodyDataSource) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBodyDataSource) GetConnectionProperties() interface{} {
	return s.ConnectionProperties
}

func (s *GetDataSourceResponseBodyDataSource) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *GetDataSourceResponseBodyDataSource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDataSourceResponseBodyDataSource) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetDataSourceResponseBodyDataSource) GetDescription() *string {
	return s.Description
}

func (s *GetDataSourceResponseBodyDataSource) GetId() *int64 {
	return s.Id
}

func (s *GetDataSourceResponseBodyDataSource) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetDataSourceResponseBodyDataSource) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *GetDataSourceResponseBodyDataSource) GetName() *string {
	return s.Name
}

func (s *GetDataSourceResponseBodyDataSource) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataSourceResponseBodyDataSource) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *GetDataSourceResponseBodyDataSource) GetType() *string {
	return s.Type
}

func (s *GetDataSourceResponseBodyDataSource) SetConnectionProperties(v interface{}) *GetDataSourceResponseBodyDataSource {
	s.ConnectionProperties = v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetConnectionPropertiesMode(v string) *GetDataSourceResponseBodyDataSource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetCreateTime(v int64) *GetDataSourceResponseBodyDataSource {
	s.CreateTime = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetCreateUser(v string) *GetDataSourceResponseBodyDataSource {
	s.CreateUser = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetDescription(v string) *GetDataSourceResponseBodyDataSource {
	s.Description = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetId(v int64) *GetDataSourceResponseBodyDataSource {
	s.Id = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetModifyTime(v int64) *GetDataSourceResponseBodyDataSource {
	s.ModifyTime = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetModifyUser(v string) *GetDataSourceResponseBodyDataSource {
	s.ModifyUser = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetName(v string) *GetDataSourceResponseBodyDataSource {
	s.Name = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetProjectId(v int64) *GetDataSourceResponseBodyDataSource {
	s.ProjectId = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetQualifiedName(v string) *GetDataSourceResponseBodyDataSource {
	s.QualifiedName = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) SetType(v string) *GetDataSourceResponseBodyDataSource {
	s.Type = &v
	return s
}

func (s *GetDataSourceResponseBodyDataSource) Validate() error {
	return dara.Validate(s)
}
