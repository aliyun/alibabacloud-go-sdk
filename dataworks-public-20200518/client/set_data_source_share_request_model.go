// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDataSourceShareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasourceName(v string) *SetDataSourceShareRequest
	GetDatasourceName() *string
	SetEnvType(v string) *SetDataSourceShareRequest
	GetEnvType() *string
	SetProjectId(v int64) *SetDataSourceShareRequest
	GetProjectId() *int64
	SetProjectPermissions(v string) *SetDataSourceShareRequest
	GetProjectPermissions() *string
	SetUserPermissions(v string) *SetDataSourceShareRequest
	GetUserPermissions() *string
}

type SetDataSourceShareRequest struct {
	// The name of the data source that you want to share.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql_name
	DatasourceName *string `json:"DatasourceName,omitempty" xml:"DatasourceName,omitempty"`
	// The environment in which the data source is used. Valid values:
	//
	// 	- 0: development environment
	//
	// 	- 1: production environment
	//
	// example:
	//
	// 1
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the DataWorks workspace to which the data source belongs. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The DataWorks workspace to which you want to share the data source. If you configure this parameter, all members of the specified DataWorks workspace can view and use the data source. The value of this parameter is a JSON array. Example: [{"projectId":1000,"permission":"WRITE","sharedName":"PX_DATAHUB1.shared_name"}], Parameter description:
	//
	// 	- projectId: the ID of the DataWorks workspace to which you want to share the data source.
	//
	// 	- permission: the mode in which the data source is shared. Valid values: READ and WRITE. The value READ indicates that all members of the specified workspace can read data from the data source, but cannot modify the data. The value WRITE indicates that all members of the specified workspace can modify the data in the data source.
	//
	// 	- sharedName: the name of the data source that you want to share.
	//
	// example:
	//
	// [{"projectId":1000,"permission":"WRITE","sharedName":"PX_DATAHUB1.shared_name"}]
	ProjectPermissions *string `json:"ProjectPermissions,omitempty" xml:"ProjectPermissions,omitempty"`
	// The user to whom you want to share the data source. If you configure this parameter, the specified user can view or use the data source. The value of this parameter is a JSON array. Example: [{"projectId":10000,"users":[{"userId":"276184575345452131","permission":"WRITE"}],"sharedName":"PX_DATAHUB1.shared_name"}], Parameter description:
	//
	// 	- projectId: the ID of the DataWorks workspace. If you configure the UserPermissions parameter, the specified user can view or use the data source only in the specified DataWorks workspace.
	//
	// 	- userId: the ID of the user to whom you want to share the data source.
	//
	// 	- permission: the mode in which the data source is shared. Valid values: READ and WRITE. The value READ indicates that the specified user can read data from the data source, but cannot modify the data. The value WRITE indicates that the specified user can modify the data in the data source.
	//
	// 	- sharedName: the name of the data source that you want to share.
	//
	// If the ProjectPermissions and UserPermissions parameters are both left empty, the specified data source is not shared to any DataWorks workspace or user. If neither of the parameters is left empty, both parameters take effect.
	//
	// example:
	//
	// [{"projectId":10000,"users":[{"userId":"276184575345452131","permission":"WRITE"}],"sharedName":"PX_DATAHUB1.shared_name"}]
	UserPermissions *string `json:"UserPermissions,omitempty" xml:"UserPermissions,omitempty"`
}

func (s SetDataSourceShareRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDataSourceShareRequest) GoString() string {
	return s.String()
}

func (s *SetDataSourceShareRequest) GetDatasourceName() *string {
	return s.DatasourceName
}

func (s *SetDataSourceShareRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *SetDataSourceShareRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SetDataSourceShareRequest) GetProjectPermissions() *string {
	return s.ProjectPermissions
}

func (s *SetDataSourceShareRequest) GetUserPermissions() *string {
	return s.UserPermissions
}

func (s *SetDataSourceShareRequest) SetDatasourceName(v string) *SetDataSourceShareRequest {
	s.DatasourceName = &v
	return s
}

func (s *SetDataSourceShareRequest) SetEnvType(v string) *SetDataSourceShareRequest {
	s.EnvType = &v
	return s
}

func (s *SetDataSourceShareRequest) SetProjectId(v int64) *SetDataSourceShareRequest {
	s.ProjectId = &v
	return s
}

func (s *SetDataSourceShareRequest) SetProjectPermissions(v string) *SetDataSourceShareRequest {
	s.ProjectPermissions = &v
	return s
}

func (s *SetDataSourceShareRequest) SetUserPermissions(v string) *SetDataSourceShareRequest {
	s.UserPermissions = &v
	return s
}

func (s *SetDataSourceShareRequest) Validate() error {
	return dara.Validate(s)
}
