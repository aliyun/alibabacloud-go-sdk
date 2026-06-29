// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnLineageByTaskIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableColumnLineageByTaskIdResponseBody
	GetCode() *string
	SetData(v []*GetTableColumnLineageByTaskIdResponseBodyData) *GetTableColumnLineageByTaskIdResponseBody
	GetData() []*GetTableColumnLineageByTaskIdResponseBodyData
	SetHttpStatusCode(v int32) *GetTableColumnLineageByTaskIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTableColumnLineageByTaskIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableColumnLineageByTaskIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableColumnLineageByTaskIdResponseBody
	GetSuccess() *bool
}

type GetTableColumnLineageByTaskIdResponseBody struct {
	// Error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Query results.
	Data []*GetTableColumnLineageByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableColumnLineageByTaskIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineageByTaskIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineageByTaskIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableColumnLineageByTaskIdResponseBody) GetData() []*GetTableColumnLineageByTaskIdResponseBodyData {
	return s.Data
}

func (s *GetTableColumnLineageByTaskIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTableColumnLineageByTaskIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableColumnLineageByTaskIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableColumnLineageByTaskIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableColumnLineageByTaskIdResponseBody) SetCode(v string) *GetTableColumnLineageByTaskIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBody) SetData(v []*GetTableColumnLineageByTaskIdResponseBodyData) *GetTableColumnLineageByTaskIdResponseBody {
	s.Data = v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBody) SetHttpStatusCode(v int32) *GetTableColumnLineageByTaskIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBody) SetMessage(v string) *GetTableColumnLineageByTaskIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBody) SetRequestId(v string) *GetTableColumnLineageByTaskIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBody) SetSuccess(v bool) *GetTableColumnLineageByTaskIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableColumnLineageByTaskIdResponseBodyData struct {
	// Business unit ID of the input table.
	//
	// example:
	//
	// 123
	InputBizUnitId *int64 `json:"InputBizUnitId,omitempty" xml:"InputBizUnitId,omitempty"`
	// GUID of the input column.
	//
	// example:
	//
	// odps.123.test_project.input_table.id
	InputColumnId *string `json:"InputColumnId,omitempty" xml:"InputColumnId,omitempty"`
	// Input column name.
	//
	// example:
	//
	// id
	InputColumnName *string `json:"InputColumnName,omitempty" xml:"InputColumnName,omitempty"`
	// Data source ID of the input table.
	//
	// example:
	//
	// 123
	InputDataSourceId *int64 `json:"InputDataSourceId,omitempty" xml:"InputDataSourceId,omitempty"`
	// Storage type of the input table.
	//
	// example:
	//
	// MAX_COMPUTE
	InputDataSourceType *string `json:"InputDataSourceType,omitempty" xml:"InputDataSourceType,omitempty"`
	// Database type of the input table.
	//
	// example:
	//
	// MAX_COMPUTE
	InputDbType *string `json:"InputDbType,omitempty" xml:"InputDbType,omitempty"`
	// Environment of the input table: DEV or PROD.
	//
	// example:
	//
	// DEV
	InputEnv *string `json:"InputEnv,omitempty" xml:"InputEnv,omitempty"`
	// Project ID of the input table.
	//
	// example:
	//
	// 123
	InputProjectId *int64 `json:"InputProjectId,omitempty" xml:"InputProjectId,omitempty"`
	// Indicates whether the input table is deleted.
	InputTableDeleted *bool `json:"InputTableDeleted,omitempty" xml:"InputTableDeleted,omitempty"`
	// GUID of the input table. Each asset has a unique GUID in the following format:
	//
	// - Logical table: dp_table.[TenantId].[BizUnitName].[TableName]
	//
	// - Compute source physical table: [EngineType].[TenantId].[ProjectName].[TableName]
	//
	// - Data source table: dp_ds_table.[TenantId].[DataSourceId].[SchemaName].[TableName]
	//
	// example:
	//
	// odps.123.test_project.order
	InputTableId *string `json:"InputTableId,omitempty" xml:"InputTableId,omitempty"`
	// Input table name.
	//
	// example:
	//
	// order
	InputTableName *string `json:"InputTableName,omitempty" xml:"InputTableName,omitempty"`
	// Input table type. Valid values:
	//
	// - PHYSICAL_TABLE: Physical table (compute source)
	//
	// - DIM_LOGIC_TABLE: Dimension logical table
	//
	// - FACT_LOGIC_TABLE: Fact logical table
	//
	// - SUM_LOGIC_TABLE: Summary logical table
	//
	// - REAL_TIME_LOGIC_TABLE: Real-time meta table
	//
	// - REAL_TIME_MIRROR_TABLE: Real-time mirror table
	//
	// - PHYSICAL_VIEW: Physical view
	//
	// - LOGICAL_VIEW: Logical view
	//
	// - DATA_SOURCE_PHYSICAL_TABLE: Data source table
	//
	// - DATA_SOURCE_VIEW: Data source view
	//
	// - DATA_SOURCE_MATERIALIZED_VIEW: Data source materialized view
	//
	// example:
	//
	// PHYSICAL_TABLE
	InputTableType *string `json:"InputTableType,omitempty" xml:"InputTableType,omitempty"`
	// Business unit ID of the output table.
	//
	// example:
	//
	// 123
	OutputBizUnitId *int64 `json:"OutputBizUnitId,omitempty" xml:"OutputBizUnitId,omitempty"`
	// GUID of the output column.
	//
	// example:
	//
	// odps.123.test_project.input_table.id
	OutputColumnId *string `json:"OutputColumnId,omitempty" xml:"OutputColumnId,omitempty"`
	// Output column name.
	//
	// example:
	//
	// id
	OutputColumnName *string `json:"OutputColumnName,omitempty" xml:"OutputColumnName,omitempty"`
	// Data source ID of the output table.
	//
	// example:
	//
	// 123
	OutputDataSourceId *int64 `json:"OutputDataSourceId,omitempty" xml:"OutputDataSourceId,omitempty"`
	// Storage type of the output table.
	//
	// example:
	//
	// MAX_COMPUTE
	OutputDataSourceType *string `json:"OutputDataSourceType,omitempty" xml:"OutputDataSourceType,omitempty"`
	// Database type of the output table.
	//
	// example:
	//
	// MAX_COMPUTE
	OutputDbType *string `json:"OutputDbType,omitempty" xml:"OutputDbType,omitempty"`
	// Environment of the output table: DEV or PROD.
	//
	// example:
	//
	// DEV
	OutputEnv *string `json:"OutputEnv,omitempty" xml:"OutputEnv,omitempty"`
	// Project ID of the output table.
	//
	// example:
	//
	// 123
	OutputProjectId *int64 `json:"OutputProjectId,omitempty" xml:"OutputProjectId,omitempty"`
	// Indicates whether the output table is deleted.
	OutputTableDeleted *bool `json:"OutputTableDeleted,omitempty" xml:"OutputTableDeleted,omitempty"`
	// GUID of the output table. Each asset has a unique GUID. For the format, see InputTableId.
	//
	// example:
	//
	// odps.123.test_project.order
	OutputTableId *string `json:"OutputTableId,omitempty" xml:"OutputTableId,omitempty"`
	// Output table name.
	//
	// example:
	//
	// order
	OutputTableName *string `json:"OutputTableName,omitempty" xml:"OutputTableName,omitempty"`
	// Output table type. For valid values, see InputTableType.
	//
	// example:
	//
	// PHYSICAL_TABLE
	OutputTableType *string `json:"OutputTableType,omitempty" xml:"OutputTableType,omitempty"`
	// Environment of the task (node) associated with the lineage: DEV or PROD.
	//
	// example:
	//
	// DEV
	TaskEnv *string `json:"TaskEnv,omitempty" xml:"TaskEnv,omitempty"`
	// Task (node) ID associated with the lineage.
	//
	// example:
	//
	// n_123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// 12345
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetTableColumnLineageByTaskIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnLineageByTaskIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputBizUnitId() *int64 {
	return s.InputBizUnitId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputColumnId() *string {
	return s.InputColumnId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputColumnName() *string {
	return s.InputColumnName
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputDataSourceId() *int64 {
	return s.InputDataSourceId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputDataSourceType() *string {
	return s.InputDataSourceType
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputDbType() *string {
	return s.InputDbType
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputEnv() *string {
	return s.InputEnv
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputProjectId() *int64 {
	return s.InputProjectId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputTableDeleted() *bool {
	return s.InputTableDeleted
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputTableId() *string {
	return s.InputTableId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputTableName() *string {
	return s.InputTableName
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetInputTableType() *string {
	return s.InputTableType
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputBizUnitId() *int64 {
	return s.OutputBizUnitId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputColumnId() *string {
	return s.OutputColumnId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputColumnName() *string {
	return s.OutputColumnName
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputDataSourceId() *int64 {
	return s.OutputDataSourceId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputDataSourceType() *string {
	return s.OutputDataSourceType
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputDbType() *string {
	return s.OutputDbType
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputEnv() *string {
	return s.OutputEnv
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputProjectId() *int64 {
	return s.OutputProjectId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputTableDeleted() *bool {
	return s.OutputTableDeleted
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputTableId() *string {
	return s.OutputTableId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputTableName() *string {
	return s.OutputTableName
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetOutputTableType() *string {
	return s.OutputTableType
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetTaskEnv() *string {
	return s.TaskEnv
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputBizUnitId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputBizUnitId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputColumnId(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputColumnId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputColumnName(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputColumnName = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputDataSourceId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputDataSourceId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputDataSourceType(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputDataSourceType = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputDbType(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputDbType = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputEnv(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputEnv = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputProjectId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputProjectId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputTableDeleted(v bool) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputTableDeleted = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputTableId(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputTableId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputTableName(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputTableName = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetInputTableType(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.InputTableType = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputBizUnitId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputBizUnitId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputColumnId(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputColumnId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputColumnName(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputColumnName = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputDataSourceId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputDataSourceId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputDataSourceType(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputDataSourceType = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputDbType(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputDbType = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputEnv(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputEnv = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputProjectId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputProjectId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputTableDeleted(v bool) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputTableDeleted = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputTableId(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputTableId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputTableName(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputTableName = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetOutputTableType(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.OutputTableType = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetTaskEnv(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.TaskEnv = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetTaskId(v string) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) SetTenantId(v int64) *GetTableColumnLineageByTaskIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetTableColumnLineageByTaskIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
