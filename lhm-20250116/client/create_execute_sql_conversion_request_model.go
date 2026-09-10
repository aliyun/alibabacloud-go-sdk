// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExecuteSqlConversionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceDialect(v string) *CreateExecuteSqlConversionRequest
	GetSourceDialect() *string
	SetSourceSqlScript(v []*CreateExecuteSqlConversionRequestSourceSqlScript) *CreateExecuteSqlConversionRequest
	GetSourceSqlScript() []*CreateExecuteSqlConversionRequestSourceSqlScript
	SetTargetDialect(v string) *CreateExecuteSqlConversionRequest
	GetTargetDialect() *string
	SetTaskDescription(v string) *CreateExecuteSqlConversionRequest
	GetTaskDescription() *string
	SetTaskName(v string) *CreateExecuteSqlConversionRequest
	GetTaskName() *string
	SetType(v int32) *CreateExecuteSqlConversionRequest
	GetType() *int32
}

type CreateExecuteSqlConversionRequest struct {
	// The source dialect.
	//
	// example:
	//
	// hive
	SourceDialect *string `json:"sourceDialect,omitempty" xml:"sourceDialect,omitempty"`
	// The list of source SQL scripts.
	SourceSqlScript []*CreateExecuteSqlConversionRequestSourceSqlScript `json:"sourceSqlScript,omitempty" xml:"sourceSqlScript,omitempty" type:"Repeated"`
	// The target dialect.
	//
	// example:
	//
	// hive
	TargetDialect *string `json:"targetDialect,omitempty" xml:"targetDialect,omitempty"`
	// The task description.
	//
	// example:
	//
	// Data validation task description
	TaskDescription *string `json:"taskDescription,omitempty" xml:"taskDescription,omitempty"`
	// The task name.
	//
	// This parameter is required.
	//
	// example:
	//
	// data_check_task_demo
	TaskName *string `json:"taskName,omitempty" xml:"taskName,omitempty"`
	// The script type. Valid values: 0 (DDL) and 1 (DQL).
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateExecuteSqlConversionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExecuteSqlConversionRequest) GoString() string {
	return s.String()
}

func (s *CreateExecuteSqlConversionRequest) GetSourceDialect() *string {
	return s.SourceDialect
}

func (s *CreateExecuteSqlConversionRequest) GetSourceSqlScript() []*CreateExecuteSqlConversionRequestSourceSqlScript {
	return s.SourceSqlScript
}

func (s *CreateExecuteSqlConversionRequest) GetTargetDialect() *string {
	return s.TargetDialect
}

func (s *CreateExecuteSqlConversionRequest) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *CreateExecuteSqlConversionRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateExecuteSqlConversionRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateExecuteSqlConversionRequest) SetSourceDialect(v string) *CreateExecuteSqlConversionRequest {
	s.SourceDialect = &v
	return s
}

func (s *CreateExecuteSqlConversionRequest) SetSourceSqlScript(v []*CreateExecuteSqlConversionRequestSourceSqlScript) *CreateExecuteSqlConversionRequest {
	s.SourceSqlScript = v
	return s
}

func (s *CreateExecuteSqlConversionRequest) SetTargetDialect(v string) *CreateExecuteSqlConversionRequest {
	s.TargetDialect = &v
	return s
}

func (s *CreateExecuteSqlConversionRequest) SetTaskDescription(v string) *CreateExecuteSqlConversionRequest {
	s.TaskDescription = &v
	return s
}

func (s *CreateExecuteSqlConversionRequest) SetTaskName(v string) *CreateExecuteSqlConversionRequest {
	s.TaskName = &v
	return s
}

func (s *CreateExecuteSqlConversionRequest) SetType(v int32) *CreateExecuteSqlConversionRequest {
	s.Type = &v
	return s
}

func (s *CreateExecuteSqlConversionRequest) Validate() error {
	if s.SourceSqlScript != nil {
		for _, item := range s.SourceSqlScript {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateExecuteSqlConversionRequestSourceSqlScript struct {
	// The error reason.
	//
	// example:
	//
	// connection timeout
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The time when the conversion is completed.
	//
	// example:
	//
	// 2026-01-16T10:00:00Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 1234567890
	ScriptId *int64 `json:"scriptId,omitempty" xml:"scriptId,omitempty"`
	// The script name.
	//
	// example:
	//
	// node_script_demo
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The script conversion status. In conversion job scenarios, valid values: pass (conversion succeeded), turning (conversion in progress), fail (conversion failed). In some scenarios, the following values are used: success (succeeded), failed (failed), skipped (skipped).
	//
	// example:
	//
	// pass
	ScriptTransformStatus *string `json:"scriptTransformStatus,omitempty" xml:"scriptTransformStatus,omitempty"`
	// The converted script content.
	//
	// example:
	//
	// SELECT 	- FROM t;
	SqlResultContent *string `json:"sqlResultContent,omitempty" xml:"sqlResultContent,omitempty"`
	// The original script content.
	//
	// example:
	//
	// SELECT 	- FROM t;
	SqlSourceContent *string `json:"sqlSourceContent,omitempty" xml:"sqlSourceContent,omitempty"`
	// The table name mappings for conversion.
	TableMappingList []*CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList `json:"tableMappingList,omitempty" xml:"tableMappingList,omitempty" type:"Repeated"`
}

func (s CreateExecuteSqlConversionRequestSourceSqlScript) String() string {
	return dara.Prettify(s)
}

func (s CreateExecuteSqlConversionRequestSourceSqlScript) GoString() string {
	return s.String()
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetFinishTime() *string {
	return s.FinishTime
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetScriptId() *int64 {
	return s.ScriptId
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetScriptTransformStatus() *string {
	return s.ScriptTransformStatus
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetSqlResultContent() *string {
	return s.SqlResultContent
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetSqlSourceContent() *string {
	return s.SqlSourceContent
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) GetTableMappingList() []*CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	return s.TableMappingList
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetErrorMessage(v string) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.ErrorMessage = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetFinishTime(v string) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.FinishTime = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetScriptId(v int64) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.ScriptId = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetScriptName(v string) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.ScriptName = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetScriptTransformStatus(v string) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.ScriptTransformStatus = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetSqlResultContent(v string) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.SqlResultContent = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetSqlSourceContent(v string) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.SqlSourceContent = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) SetTableMappingList(v []*CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) *CreateExecuteSqlConversionRequestSourceSqlScript {
	s.TableMappingList = v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScript) Validate() error {
	if s.TableMappingList != nil {
		for _, item := range s.TableMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList struct {
	// The primary key.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The source type. Valid values: DB and Schema.
	//
	// example:
	//
	// db_demo
	SourceSchema *string `json:"sourceSchema,omitempty" xml:"sourceSchema,omitempty"`
	// The source table name.
	//
	// example:
	//
	// table_demo
	SourceTableName *string `json:"sourceTableName,omitempty" xml:"sourceTableName,omitempty"`
	// The target table name.
	//
	// example:
	//
	// table_demo
	TargetTableName *string `json:"targetTableName,omitempty" xml:"targetTableName,omitempty"`
	// The target type. Valid values: DB and Schema.
	//
	// example:
	//
	// hive
	TargetType *string `json:"targetType,omitempty" xml:"targetType,omitempty"`
	// The SQL conversion task ID.
	//
	// example:
	//
	// 10001
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 10001
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) String() string {
	return dara.Prettify(s)
}

func (s CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GoString() string {
	return s.String()
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetId() *int64 {
	return s.Id
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetSourceSchema() *string {
	return s.SourceSchema
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) GetUid() *string {
	return s.Uid
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetId(v int64) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.Id = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetSourceSchema(v string) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.SourceSchema = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetSourceTableName(v string) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.SourceTableName = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetTargetTableName(v string) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.TargetTableName = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetTargetType(v string) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.TargetType = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetTaskId(v int64) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.TaskId = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetTenantId(v string) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.TenantId = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) SetUid(v string) *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList {
	s.Uid = &v
	return s
}

func (s *CreateExecuteSqlConversionRequestSourceSqlScriptTableMappingList) Validate() error {
	return dara.Validate(s)
}
