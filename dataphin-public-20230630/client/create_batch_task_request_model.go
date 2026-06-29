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

func (s *CreateBatchTaskRequest) SetCreateCommand(v *CreateBatchTaskRequestCreateCommand) *CreateBatchTaskRequest {
	s.CreateCommand = v
	return s
}

func (s *CreateBatchTaskRequest) SetOpTenantId(v int64) *CreateBatchTaskRequest {
	s.OpTenantId = &v
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
	// The catalog for a database SQL node. This parameter takes effect only for data source types that require a catalog, such as Presto.
	//
	// example:
	//
	// mysql_catalog
	DataSourceCatalog *string `json:"DataSourceCatalog,omitempty" xml:"DataSourceCatalog,omitempty"`
	// The data source ID for a database SQL node.
	//
	// example:
	//
	// 12131111
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The schema for a database SQL node. This parameter takes effect only for data source types that require a schema, such as Oracle.
	//
	// example:
	//
	// erp
	DataSourceSchema *string `json:"DataSourceSchema,omitempty" xml:"DataSourceSchema,omitempty"`
	// The description.
	//
	// This parameter is required.
	//
	// example:
	//
	// xx test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The folder path in the menu tree to which the node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// /a/b
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The execution engine for the node, such as a Python node. Valid values:
	//
	// - 1: PYTHON2_7
	//
	// - 2: PYTHON3_7
	//
	// - 3: PYTHON3_11.
	//
	// example:
	//
	// PYTHON3_7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The name of the batch task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the project to which the node belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of third-party Python packages that the node depends on.
	PythonModuleList []*string `json:"PythonModuleList,omitempty" xml:"PythonModuleList,omitempty" type:"Repeated"`
	// The scheduling type. Valid values:
	//
	// - 1: periodic node.
	//
	// - 3: manual node.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The node type. Valid values:
	//
	// - Hive_SQL: 1
	//
	// - Hive_SQL_23X: 101
	//
	// - HIVE_SQL_FUSION_INSIGHT_80X: 111
	//
	// - COMMON_HIVE_SQL: 131
	//
	// - HADOOP_MR: 2
	//
	// - MaxCompute_SQL: 5
	//
	// - MaxCompute_MR: 6
	//
	// - SPARK_SQL_ON_MAX_COMPUTE: 7
	//
	// - SPARK_JAR_ON_MAX_COMPUTE: 8
	//
	// - SPARK_SQL_ON_HIVE: 17
	//
	// - Spark_JAR_ON_HIVE: 18
	//
	// - Shell: 10
	//
	// - PAI_DESIGNER: 71
	//
	// - DataX: 15
	//
	// - Merge: 16
	//
	// - Python: 21
	//
	// - Python37x: 22
	//
	// - Perl: 23
	//
	// - Python311x: 24
	//
	// - OneService_SQL: 25
	//
	// - ONE_SERVICE_SQL_ADB_FOR_PG: 26
	//
	// - OneService_SQL_Hive11x: 27
	//
	// - OneService_SQL_Hive23x: 29
	//
	// - ONE_SERVICE_SQL_TDH_INCEPTOR: 75
	//
	// - ONE_SERVICE_SQL_HIVE_CDP: 91
	//
	// - ONE_SERVICE_SQL_HIVE_ASIA_INFO_DP_53X: 92
	//
	// - Dlink: 30
	//
	// - ONE_SERVICE_SQL_ADB_FOR_MYSQL: 33
	//
	// - Logical: 31
	//
	// - Flink_Streaming: 41
	//
	// - Flink_Batch: 42
	//
	// - ADB_FOR_PG: 51
	//
	// - DryRun: 100
	//
	// - CHECK: 902
	//
	// - VIRTUAL: 999
	//
	// - INCEPTOR_SQL: 10000
	//
	// - HOLOGRES_SQL: 28
	//
	// - ARGODB_SQL: 76
	//
	// - IMPALA_SQL: 78
	//
	// - STARROCKS_SQL: 79
	//
	// - SPARK_SQL: 80
	//
	// - GAUSS_SQL: 81
	//
	// - DATABASE_SQL: 998
	//
	// - EXTERNAL_TRIGGER: 997.
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

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceCatalog() *string {
	return s.DataSourceCatalog
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateBatchTaskRequestCreateCommand) GetDataSourceSchema() *string {
	return s.DataSourceSchema
}

func (s *CreateBatchTaskRequestCreateCommand) GetDescription() *string {
	return s.Description
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

func (s *CreateBatchTaskRequestCreateCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateBatchTaskRequestCreateCommand) GetPythonModuleList() []*string {
	return s.PythonModuleList
}

func (s *CreateBatchTaskRequestCreateCommand) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *CreateBatchTaskRequestCreateCommand) GetTaskType() *int32 {
	return s.TaskType
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

func (s *CreateBatchTaskRequestCreateCommand) SetDescription(v string) *CreateBatchTaskRequestCreateCommand {
	s.Description = &v
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

func (s *CreateBatchTaskRequestCreateCommand) SetTaskType(v int32) *CreateBatchTaskRequestCreateCommand {
	s.TaskType = &v
	return s
}

func (s *CreateBatchTaskRequestCreateCommand) Validate() error {
	return dara.Validate(s)
}
