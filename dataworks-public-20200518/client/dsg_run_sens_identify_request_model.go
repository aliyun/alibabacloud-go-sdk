// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgRunSensIdentifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEsMetaParams(v []*DsgRunSensIdentifyRequestEsMetaParams) *DsgRunSensIdentifyRequest
	GetEsMetaParams() []*DsgRunSensIdentifyRequestEsMetaParams
	SetTenantId(v string) *DsgRunSensIdentifyRequest
	GetTenantId() *string
}

type DsgRunSensIdentifyRequest struct {
	// The parameters that you need to configure to scan specified metadata.
	EsMetaParams []*DsgRunSensIdentifyRequestEsMetaParams `json:"EsMetaParams,omitempty" xml:"EsMetaParams,omitempty" type:"Repeated"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DsgRunSensIdentifyRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgRunSensIdentifyRequest) GoString() string {
	return s.String()
}

func (s *DsgRunSensIdentifyRequest) GetEsMetaParams() []*DsgRunSensIdentifyRequestEsMetaParams {
	return s.EsMetaParams
}

func (s *DsgRunSensIdentifyRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DsgRunSensIdentifyRequest) SetEsMetaParams(v []*DsgRunSensIdentifyRequestEsMetaParams) *DsgRunSensIdentifyRequest {
	s.EsMetaParams = v
	return s
}

func (s *DsgRunSensIdentifyRequest) SetTenantId(v string) *DsgRunSensIdentifyRequest {
	s.TenantId = &v
	return s
}

func (s *DsgRunSensIdentifyRequest) Validate() error {
	return dara.Validate(s)
}

type DsgRunSensIdentifyRequestEsMetaParams struct {
	// The cluster ID. You can obtain the ID based on the data source you use.
	//
	// example:
	//
	// 101010
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- ODPS.ODPS
	//
	// 	- EMR
	//
	// 	- HOLO.POSTGRES
	//
	// example:
	//
	// ODPS.ODPS
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The instance ID. You can obtain the ID based on the data source you use.
	//
	// example:
	//
	// hgprecn-cn-9lb37k181024
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to query the name.
	//
	// example:
	//
	// 1666676756691024
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// default
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The names of the tables.
	TableNameList []*string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty" type:"Repeated"`
	// The username of the operator. We recommend that you enter the username of your Alibaba Cloud account.
	//
	// example:
	//
	// xxx-uat
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DsgRunSensIdentifyRequestEsMetaParams) String() string {
	return dara.Prettify(s)
}

func (s DsgRunSensIdentifyRequestEsMetaParams) GoString() string {
	return s.String()
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetClusterId() *string {
	return s.ClusterId
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetDbType() *string {
	return s.DbType
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetTableName() *string {
	return s.TableName
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetTableNameList() []*string {
	return s.TableNameList
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) GetUser() *string {
	return s.User
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetClusterId(v string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.ClusterId = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetDbType(v string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.DbType = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetInstanceId(v int64) *DsgRunSensIdentifyRequestEsMetaParams {
	s.InstanceId = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetProjectName(v string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.ProjectName = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetSchemaName(v string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.SchemaName = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetTableName(v string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.TableName = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetTableNameList(v []*string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.TableNameList = v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) SetUser(v string) *DsgRunSensIdentifyRequestEsMetaParams {
	s.User = &v
	return s
}

func (s *DsgRunSensIdentifyRequestEsMetaParams) Validate() error {
	return dara.Validate(s)
}
