// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateCommand(v *CreateBatchTaskRequestCreateCommand) *CreateBatchTaskRequest
	GetCreateCommand() *CreateBatchTaskRequestCreateCommand
	SetOpTenantId(v int64) *CreateBatchTaskRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *CreateBatchTaskRequest
	GetOpUserId() *string
}

type CreateBatchTaskRequest struct {
	// The create command.
	//
	// This parameter is required.
	CreateCommand *CreateBatchTaskRequestCreateCommand `json:"CreateCommand,omitempty" xml:"CreateCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s CreateBatchTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequest) GetCreateCommand() *CreateBatchTaskRequestCreateCommand {
	return s.CreateCommand
}

func (s *CreateBatchTaskRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateBatchTaskRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *CreateBatchTaskRequest) SetCreateCommand(v *CreateBatchTaskRequestCreateCommand) *CreateBatchTaskRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBatchTaskRequest) SetOpTenantId(v int64) *CreateBatchTaskRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateBatchTaskRequest) SetOpUserId(v string) *CreateBatchTaskRequest {
	s.OpUserId = &v
	return s
}

func (s *CreateBatchTaskRequest) Validate() error {
	if s.CreateCommand != nil {
		if err := s.CreateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchTaskRequestCreateCommand struct {
	// The ID of the offline code template.
	//
	// example:
	//
	// st_123456
	CodeTemplateId *string `json:"CodeTemplateId,omitempty" xml:"CodeTemplateId,omitempty"`
	// The version of the offline code template.
	//
	// example:
	//
	// 1
	CodeTemplateVersion *int32 `json:"CodeTemplateVersion,omitempty" xml:"CodeTemplateVersion,omitempty"`
	// The catalog for database SQL tasks. This parameter takes effect only for data source types that require a catalog, such as Presto.
	//
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// The ID of the data source for database SQL tasks.
	//
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The schema for database SQL tasks. This parameter takes effect only for data source types that require a schema, such as Oracle.
	//
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// The Python connection configuration.
	DatabricksPythonConnectInfo *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo `json:"DatabricksPythonConnectInfo,omitempty" xml:"DatabricksPythonConnectInfo,omitempty" type:"Struct"`
	// The description.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the development owners.
	DevelopOwnerIdList []*string `json:"DevelopOwnerIdList,omitempty" xml:"DevelopOwnerIdList,omitempty" type:"Repeated"`
	// The directory in the menu tree to which the task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// /a/b
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The task execution engine, such as Python tasks. Valid values:
	//
	// - PYTHON2_7
	//
	// - PYTHON3_7
	//
	// - PYTHON3_11
	//
	// example:
	//
	// PYTHON3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The name of the offline task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the O&M owners.
	OpsOwnerIdList []*string `json:"OpsOwnerIdList,omitempty" xml:"OpsOwnerIdList,omitempty" type:"Repeated"`
	// The PAI configuration.
	PaiConfig *CreateBatchTaskRequestCreateCommandPaiConfig `json:"PaiConfig,omitempty" xml:"PaiConfig,omitempty" type:"Struct"`
	// The ID of the project to which the task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The third-party Python packages that the task depends on.
	PythonModuleList []*string `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The scheduling type. Valid values: 1 (periodic task) and 3 (manual task).
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// Specifies whether to enable dynamic resource configuration for Spark. This parameter takes effect only for Spark SQL tasks. If you do not specify this parameter, dynamic resource configuration is enabled by default for Spark SQL tasks.
	//
	// example:
	//
	// true
	SparkDynamicResourceSwitch *bool `json:"SparkDynamicResourceSwitch,omitempty" xml:"SparkDynamicResourceSwitch,omitempty"`
	// The effective scope of Spark resources. Valid values: ONLY_PROD and PROD_AND_DEV. If you do not specify this parameter, the default value ONLY_PROD is used for Spark SQL tasks.
	//
	// example:
	//
	// ONLY_PROD
	SparkResourceEffectScope *string `json:"SparkResourceEffectScope,omitempty" xml:"SparkResourceEffectScope,omitempty"`
	// The resource configuration for the Spark task. The value is a JSON string, such as {"spark.executor.memory":"2g"}.
	//
	// example:
	//
	// {"spark.executor.memory":"2g"}
	SparkSqlResourceConfigMap *string `json:"SparkSqlResourceConfigMap,omitempty" xml:"SparkSqlResourceConfigMap,omitempty"`
	// The MaxCompute Spark version. Valid values: MAX_COMPUTE_23X, MAX_COMPUTE_245X, and MAX_COMPUTE_311X. For Apsara Stack, only MAX_COMPUTE_APSARA is supported. Display names such as Spark 2.3.0, Spark 2.4.5, and Spark 3.1.1 are also accepted.
	//
	// example:
	//
	// Spark 3.1.1
	SparkVersion *string `json:"SparkVersion,omitempty" xml:"SparkVersion,omitempty"`
	// The task type. Valid values:
	//
	// - 1: Hive_SQL.
	//
	// - 101: Hive_SQL_23X.
	//
	// - 111: HIVE_SQL_FUSION_INSIGHT_80X.
	//
	// - 131: COMMON_HIVE_SQL.
	//
	// - 2: HADOOP_MR.
	//
	// - 5: MaxCompute_SQL.
	//
	// - 6: MaxCompute_MR.
	//
	// - 7: SPARK_SQL_ON_MAX_COMPUTE.
	//
	// - 8: SPARK_JAR_ON_MAX_COMPUTE.
	//
	// - 17: SPARK_SQL_ON_HIVE.
	//
	// - 18: Spark_JAR_ON_HIVE.
	//
	// - 10: Shell.
	//
	// - 71: PAI_DESIGNER.
	//
	// - 15: DataX.
	//
	// - 16: Merge.
	//
	// - 21: Python.
	//
	// - 22: Python37x.
	//
	// - 23: Perl.
	//
	// - 24: Python311x.
	//
	// - 25: OneService_SQL.
	//
	// - 26: ONE_SERVICE_SQL_ADB_FOR_PG.
	//
	// - 27: OneService_SQL_Hive11x.
	//
	// - 29: OneService_SQL_Hive23x.
	//
	// - 75: ONE_SERVICE_SQL_TDH_INCEPTOR.
	//
	// - 91: ONE_SERVICE_SQL_HIVE_CDP.
	//
	// - 92: ONE_SERVICE_SQL_HIVE_ASIA_INFO_DP_53X.
	//
	// - 30: Dlink.
	//
	// - 33: ONE_SERVICE_SQL_ADB_FOR_MYSQL.
	//
	// - 31: Logical.
	//
	// - 41: Flink_Streaming.
	//
	// - 42: Flink_Batch.
	//
	// - 51: ADB_FOR_PG.
	//
	// - 100: DryRun.
	//
	// - 902: CHECK.
	//
	// - 999: VIRTUAL.
	//
	// - 10000: INCEPTOR_SQL.
	//
	// - 28: HOLOGRES_SQL.
	//
	// - 76: ARGODB_SQL.
	//
	// - 78: IMPALA_SQL.
	//
	// - 79: STARROCKS_SQL.
	//
	// - 80: SPARK_SQL.
	//
	// - 81: GAUSS_SQL.
	//
	// - 998: DATABASE_SQL.
	//
	// - 997: EXTERNAL_TRIGGER.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateBatchTaskRequestCreateCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequestCreateCommand) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequestCreateCommand) GetCodeTemplateId() *string {
	return s.CodeTemplateId
}

func (s *CreateBatchTaskRequestCreateCommand) GetCodeTemplateVersion() *int32 {
	return s.CodeTemplateVersion
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *CreateBatchTaskRequestCreateCommand) GetDatabricksPythonConnectInfo() *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo {
	return s.DatabricksPythonConnectInfo
}

func (s *CreateBatchTaskRequestCreateCommand) GetDescription() *string {
	return s.Description
}

func (s *CreateBatchTaskRequestCreateCommand) GetDevelopOwnerIdList() []*string {
	return s.DevelopOwnerIdList
}

func (s *CreateBatchTaskRequestCreateCommand) GetDirectory() *string {
	return s.Directory
}

func (s *CreateBatchTaskRequestCreateCommand) GetEngine() *string {
	return s.Engine
}

func (s *CreateBatchTaskRequestCreateCommand) GetName() *string {
	return s.Name
}

func (s *CreateBatchTaskRequestCreateCommand) GetOpsOwnerIdList() []*string {
	return s.OpsOwnerIdList
}

func (s *CreateBatchTaskRequestCreateCommand) GetPaiConfig() *CreateBatchTaskRequestCreateCommandPaiConfig {
	return s.PaiConfig
}

func (s *CreateBatchTaskRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBatchTaskRequestCreateCommand) GetPythonModuleList() []*string {
	return s.PythonModuleList
}

func (s *CreateBatchTaskRequestCreateCommand) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *CreateBatchTaskRequestCreateCommand) GetSparkDynamicResourceSwitch() *bool {
	return s.SparkDynamicResourceSwitch
}

func (s *CreateBatchTaskRequestCreateCommand) GetSparkResourceEffectScope() *string {
	return s.SparkResourceEffectScope
}

func (s *CreateBatchTaskRequestCreateCommand) GetSparkSqlResourceConfigMap() *string {
	return s.SparkSqlResourceConfigMap
}

func (s *CreateBatchTaskRequestCreateCommand) GetSparkVersion() *string {
	return s.SparkVersion
}

func (s *CreateBatchTaskRequestCreateCommand) GetTaskType() *int32 {
	return s.TaskType
}

func (s *CreateBatchTaskRequestCreateCommand) SetCodeTemplateId(v string) *CreateBatchTaskRequestCreateCommand {
	s.CodeTemplateId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetCodeTemplateVersion(v int32) *CreateBatchTaskRequestCreateCommand {
	s.CodeTemplateVersion = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDataSourceCatalog(v string) *CreateBatchTaskRequestCreateCommand {
	s.DataSourceCatalog = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDataSourceId(v string) *CreateBatchTaskRequestCreateCommand {
	s.DataSourceId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDataSourceSchema(v string) *CreateBatchTaskRequestCreateCommand {
	s.DataSourceSchema = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDatabricksPythonConnectInfo(v *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) *CreateBatchTaskRequestCreateCommand {
	s.DatabricksPythonConnectInfo = v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDescription(v string) *CreateBatchTaskRequestCreateCommand {
	s.Description = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDevelopOwnerIdList(v []*string) *CreateBatchTaskRequestCreateCommand {
	s.DevelopOwnerIdList = v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetDirectory(v string) *CreateBatchTaskRequestCreateCommand {
	s.Directory = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetEngine(v string) *CreateBatchTaskRequestCreateCommand {
	s.Engine = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetName(v string) *CreateBatchTaskRequestCreateCommand {
	s.Name = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetOpsOwnerIdList(v []*string) *CreateBatchTaskRequestCreateCommand {
	s.OpsOwnerIdList = v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetPaiConfig(v *CreateBatchTaskRequestCreateCommandPaiConfig) *CreateBatchTaskRequestCreateCommand {
	s.PaiConfig = v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetProjectId(v int64) *CreateBatchTaskRequestCreateCommand {
	s.ProjectId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetPythonModuleList(v []*string) *CreateBatchTaskRequestCreateCommand {
	s.PythonModuleList = v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetScheduleType(v int32) *CreateBatchTaskRequestCreateCommand {
	s.ScheduleType = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetSparkDynamicResourceSwitch(v bool) *CreateBatchTaskRequestCreateCommand {
	s.SparkDynamicResourceSwitch = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetSparkResourceEffectScope(v string) *CreateBatchTaskRequestCreateCommand {
	s.SparkResourceEffectScope = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetSparkSqlResourceConfigMap(v string) *CreateBatchTaskRequestCreateCommand {
	s.SparkSqlResourceConfigMap = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetSparkVersion(v string) *CreateBatchTaskRequestCreateCommand {
	s.SparkVersion = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) SetTaskType(v int32) *CreateBatchTaskRequestCreateCommand {
	s.TaskType = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) Validate() error {
	if s.DatabricksPythonConnectInfo != nil {
		if err := s.DatabricksPythonConnectInfo.Validate(); err != nil {
			return err
		}
	}
	if s.PaiConfig != nil {
		if err := s.PaiConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo struct {
	// The HTTP path for the development environment.
	//
	// example:
	//
	// /sql/1.0/warehouses/dev
	DevHttpPath *string `json:"DevHttpPath,omitempty" xml:"DevHttpPath,omitempty"`
	// Specifies whether to enable Databricks Python connect.
	//
	// example:
	//
	// true
	EnableDatabricksPythonConnect *bool `json:"EnableDatabricksPythonConnect,omitempty" xml:"EnableDatabricksPythonConnect,omitempty"`
	// The HTTP path for the production environment.
	//
	// example:
	//
	// /sql/1.0/warehouses/prod
	ProdHttpPath *string `json:"ProdHttpPath,omitempty" xml:"ProdHttpPath,omitempty"`
}

func (s CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) GetDevHttpPath() *string {
	return s.DevHttpPath
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) GetEnableDatabricksPythonConnect() *bool {
	return s.EnableDatabricksPythonConnect
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) GetProdHttpPath() *string {
	return s.ProdHttpPath
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) SetDevHttpPath(v string) *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo {
	s.DevHttpPath = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) SetEnableDatabricksPythonConnect(v bool) *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo {
	s.EnableDatabricksPythonConnect = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) SetProdHttpPath(v string) *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo {
	s.ProdHttpPath = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommandDatabricksPythonConnectInfo) Validate() error {
	return dara.Validate(s)
}

type CreateBatchTaskRequestCreateCommandPaiConfig struct {
	// The ID of the PAI experiment.
	//
	// example:
	//
	// 789012
	ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// The PAI region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the PAI workspace.
	//
	// example:
	//
	// 123456
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" xml:"WorkSpaceId,omitempty"`
}

func (s CreateBatchTaskRequestCreateCommandPaiConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchTaskRequestCreateCommandPaiConfig) GoString() string {
	return s.String()
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) GetExperimentId() *string {
	return s.ExperimentId
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) GetRegion() *string {
	return s.Region
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) GetWorkSpaceId() *string {
	return s.WorkSpaceId
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) SetExperimentId(v string) *CreateBatchTaskRequestCreateCommandPaiConfig {
	s.ExperimentId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) SetRegion(v string) *CreateBatchTaskRequestCreateCommandPaiConfig {
	s.Region = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) SetWorkSpaceId(v string) *CreateBatchTaskRequestCreateCommandPaiConfig {
	s.WorkSpaceId = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommandPaiConfig) Validate() error {
	return dara.Validate(s)
}
