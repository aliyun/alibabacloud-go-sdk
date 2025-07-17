// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataArchiveOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDataArchiveOrderRequest
	GetComment() *string
	SetParam(v *CreateDataArchiveOrderRequestParam) *CreateDataArchiveOrderRequest
	GetParam() *CreateDataArchiveOrderRequestParam
	SetParentId(v int64) *CreateDataArchiveOrderRequest
	GetParentId() *int64
	SetPluginType(v string) *CreateDataArchiveOrderRequest
	GetPluginType() *string
	SetRelatedUserList(v []*string) *CreateDataArchiveOrderRequest
	GetRelatedUserList() []*string
	SetTid(v int64) *CreateDataArchiveOrderRequest
	GetTid() *int64
}

type CreateDataArchiveOrderRequest struct {
	// The description of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters for archiving data.
	//
	// This parameter is required.
	Param *CreateDataArchiveOrderRequestParam `json:"Param,omitempty" xml:"Param,omitempty" type:"Struct"`
	// The ID of the parent ticket. A parent ticket is generated only when a child ticket is created.
	//
	// example:
	//
	// 123****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the plug-in. Default value: DATA_ARCHIVE.
	//
	// example:
	//
	// DATA_ARCHIVE
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The list of the related users.
	RelatedUserList []*string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty" type:"Repeated"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataArchiveOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataArchiveOrderRequest) GetParam() *CreateDataArchiveOrderRequestParam {
	return s.Param
}

func (s *CreateDataArchiveOrderRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDataArchiveOrderRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *CreateDataArchiveOrderRequest) GetRelatedUserList() []*string {
	return s.RelatedUserList
}

func (s *CreateDataArchiveOrderRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataArchiveOrderRequest) SetComment(v string) *CreateDataArchiveOrderRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataArchiveOrderRequest) SetParam(v *CreateDataArchiveOrderRequestParam) *CreateDataArchiveOrderRequest {
	s.Param = v
	return s
}

func (s *CreateDataArchiveOrderRequest) SetParentId(v int64) *CreateDataArchiveOrderRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataArchiveOrderRequest) SetPluginType(v string) *CreateDataArchiveOrderRequest {
	s.PluginType = &v
	return s
}

func (s *CreateDataArchiveOrderRequest) SetRelatedUserList(v []*string) *CreateDataArchiveOrderRequest {
	s.RelatedUserList = v
	return s
}

func (s *CreateDataArchiveOrderRequest) SetTid(v int64) *CreateDataArchiveOrderRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataArchiveOrderRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataArchiveOrderRequestParam struct {
	// The archiving destination to which you want to archive data. Valid values:
	//
	// >  If you set ArchiveMethod to a value other than inner_oss, you must register the corresponding destination database with Data Management (DMS) before you create the data archiving ticket. After the database is registered with DMS, the database is displayed in the Instances Connected section of the DMS console.
	//
	// 	- **inner_oss**: dedicated storage, which is a built-in Object Storage Service (OSS) bucket.
	//
	// 	- **oss_userself**: OSS bucket of the user.
	//
	// 	- **mysql**: ApsaraDB RDS for MySQL instance.
	//
	// 	- **polardb**: PolarDB for MySQL cluster.
	//
	// 	- **adb_mysql**: AnalyticDB for MySQL V3.0 cluster.
	//
	// 	- **lindorm**: Lindorm instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	ArchiveMethod *string `json:"ArchiveMethod,omitempty" xml:"ArchiveMethod,omitempty"`
	// A crontab expression that specifies the scheduling cycle of the data archiving task. For more information, see the [Crontab expressions](https://help.aliyun.com/document_detail/206581.html) section of the "Create shadow tables for synchronization" topic. You must specify this parameter if you set RunMethod to schedule.
	//
	// example:
	//
	// 00 05 11 	- 	- ?
	CronStr *string `json:"CronStr,omitempty" xml:"CronStr,omitempty"`
	// The database ID. If the database is a self-managed database or a third-party cloud database, you can call the [GetDatabase](https://help.aliyun.com/document_detail/465856.html) operation to query the database ID. If the database is an Alibaba Cloud database, ignore this parameter.
	//
	// example:
	//
	// 1***
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// Specifies whether the database is a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The post behaviors.
	OrderAfter []*string `json:"OrderAfter,omitempty" xml:"OrderAfter,omitempty" type:"Repeated"`
	// The method that is used to run the data archiving task. Valid values:
	//
	// 	- **schedule**: The data archiving task is periodically scheduled.
	//
	// 	- **now**: The data archiving task is immediately run.
	//
	// This parameter is required.
	//
	// example:
	//
	// now
	RunMethod *string `json:"RunMethod,omitempty" xml:"RunMethod,omitempty"`
	// The catalog of the source database. Valid values:
	//
	// 	- **def**: Set this parameter to def if the source database is of the two-layer logical schema, such as a MySQL database, a PolarDB for MySQL cluster, or an AnalyticDB for MySQL instance.
	//
	// 	- **Empty string**: Set this parameter to an empty string if the source database is a Lindorm or ApsaraDB for MongoDB instance.
	//
	// 	- **Catalog name**: Set this parameter to the catalog name of the source database if the source database is of the three-layer logical schema, such as a PostgreSQL database.
	//
	// This parameter is required.
	//
	// example:
	//
	// def
	SourceCatalogName *string `json:"SourceCatalogName,omitempty" xml:"SourceCatalogName,omitempty"`
	// The name of the source instance. If the database instance is a self-managed database or a third-party cloud database, you can call the [GetInstance](https://help.aliyun.com/document_detail/465826.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-bp1*******
	SourceInstanceName *string `json:"SourceInstanceName,omitempty" xml:"SourceInstanceName,omitempty"`
	// The schema name of the source database. The schema name of the source database is the same as that of the destination database. If the source database is a MySQL database, this parameter specifies the name of the source database. If the source database is a PostgreSQL database, this parameter specifies the schema name of the source database.
	//
	// This parameter is required.
	//
	// example:
	//
	// schema_test
	SourceSchemaName *string `json:"SourceSchemaName,omitempty" xml:"SourceSchemaName,omitempty"`
	// The collection of tables to be archived.
	//
	// This parameter is required.
	TableIncludes []*CreateDataArchiveOrderRequestParamTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
	// The table names mapped to the destination database. This parameter is not required and the default value is used.
	TableMapping []*string `json:"TableMapping,omitempty" xml:"TableMapping,omitempty" type:"Repeated"`
	// The host of the destination instance. If the destination instance can be accessed over an internal network or the Internet, preferentially set the value to the internal endpoint of the destination instance.
	//
	// 	- If data is archived in an OSS bucket, set the value to the name of the bucket.
	//
	// 	- If data is archived in dedicated storage space, set the value to inner_oss.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1*********.ads.aliyuncs.com
	TargetInstanceHost *string `json:"TargetInstanceHost,omitempty" xml:"TargetInstanceHost,omitempty"`
	// The configuration of archiving variables. You can use a time variable as a filter condition for archiving data. Each variable has two attributes: name and pattern.
	Variables []*CreateDataArchiveOrderRequestParamVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s CreateDataArchiveOrderRequestParam) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderRequestParam) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderRequestParam) GetArchiveMethod() *string {
	return s.ArchiveMethod
}

func (s *CreateDataArchiveOrderRequestParam) GetCronStr() *string {
	return s.CronStr
}

func (s *CreateDataArchiveOrderRequestParam) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *CreateDataArchiveOrderRequestParam) GetLogic() *bool {
	return s.Logic
}

func (s *CreateDataArchiveOrderRequestParam) GetOrderAfter() []*string {
	return s.OrderAfter
}

func (s *CreateDataArchiveOrderRequestParam) GetRunMethod() *string {
	return s.RunMethod
}

func (s *CreateDataArchiveOrderRequestParam) GetSourceCatalogName() *string {
	return s.SourceCatalogName
}

func (s *CreateDataArchiveOrderRequestParam) GetSourceInstanceName() *string {
	return s.SourceInstanceName
}

func (s *CreateDataArchiveOrderRequestParam) GetSourceSchemaName() *string {
	return s.SourceSchemaName
}

func (s *CreateDataArchiveOrderRequestParam) GetTableIncludes() []*CreateDataArchiveOrderRequestParamTableIncludes {
	return s.TableIncludes
}

func (s *CreateDataArchiveOrderRequestParam) GetTableMapping() []*string {
	return s.TableMapping
}

func (s *CreateDataArchiveOrderRequestParam) GetTargetInstanceHost() *string {
	return s.TargetInstanceHost
}

func (s *CreateDataArchiveOrderRequestParam) GetVariables() []*CreateDataArchiveOrderRequestParamVariables {
	return s.Variables
}

func (s *CreateDataArchiveOrderRequestParam) SetArchiveMethod(v string) *CreateDataArchiveOrderRequestParam {
	s.ArchiveMethod = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetCronStr(v string) *CreateDataArchiveOrderRequestParam {
	s.CronStr = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetDatabaseId(v string) *CreateDataArchiveOrderRequestParam {
	s.DatabaseId = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetLogic(v bool) *CreateDataArchiveOrderRequestParam {
	s.Logic = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetOrderAfter(v []*string) *CreateDataArchiveOrderRequestParam {
	s.OrderAfter = v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetRunMethod(v string) *CreateDataArchiveOrderRequestParam {
	s.RunMethod = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetSourceCatalogName(v string) *CreateDataArchiveOrderRequestParam {
	s.SourceCatalogName = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetSourceInstanceName(v string) *CreateDataArchiveOrderRequestParam {
	s.SourceInstanceName = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetSourceSchemaName(v string) *CreateDataArchiveOrderRequestParam {
	s.SourceSchemaName = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetTableIncludes(v []*CreateDataArchiveOrderRequestParamTableIncludes) *CreateDataArchiveOrderRequestParam {
	s.TableIncludes = v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetTableMapping(v []*string) *CreateDataArchiveOrderRequestParam {
	s.TableMapping = v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetTargetInstanceHost(v string) *CreateDataArchiveOrderRequestParam {
	s.TargetInstanceHost = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) SetVariables(v []*CreateDataArchiveOrderRequestParamVariables) *CreateDataArchiveOrderRequestParam {
	s.Variables = v
	return s
}

func (s *CreateDataArchiveOrderRequestParam) Validate() error {
	return dara.Validate(s)
}

type CreateDataArchiveOrderRequestParamTableIncludes struct {
	// The table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// table1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The filter condition that is specified by the WHERE clause of the archiving configuration. If a time variable is used in the filter condition, the filter condition is specified in the following format: field name <=\\"${variable name}\\". The variable name in the filter condition must be the same as the time variable name that is specified in the Variables parameter.
	//
	// example:
	//
	// gmt_modified<\\"${time}\\"
	TableWhere *string `json:"TableWhere,omitempty" xml:"TableWhere,omitempty"`
}

func (s CreateDataArchiveOrderRequestParamTableIncludes) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderRequestParamTableIncludes) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderRequestParamTableIncludes) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataArchiveOrderRequestParamTableIncludes) GetTableWhere() *string {
	return s.TableWhere
}

func (s *CreateDataArchiveOrderRequestParamTableIncludes) SetTableName(v string) *CreateDataArchiveOrderRequestParamTableIncludes {
	s.TableName = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParamTableIncludes) SetTableWhere(v string) *CreateDataArchiveOrderRequestParamTableIncludes {
	s.TableWhere = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParamTableIncludes) Validate() error {
	return dara.Validate(s)
}

type CreateDataArchiveOrderRequestParamVariables struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pattern *string `json:"Pattern,omitempty" xml:"Pattern,omitempty"`
}

func (s CreateDataArchiveOrderRequestParamVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderRequestParamVariables) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderRequestParamVariables) GetName() *string {
	return s.Name
}

func (s *CreateDataArchiveOrderRequestParamVariables) GetPattern() *string {
	return s.Pattern
}

func (s *CreateDataArchiveOrderRequestParamVariables) SetName(v string) *CreateDataArchiveOrderRequestParamVariables {
	s.Name = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParamVariables) SetPattern(v string) *CreateDataArchiveOrderRequestParamVariables {
	s.Pattern = &v
	return s
}

func (s *CreateDataArchiveOrderRequestParamVariables) Validate() error {
	return dara.Validate(s)
}
