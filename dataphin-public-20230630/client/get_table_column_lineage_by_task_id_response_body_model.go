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
	// example:
	//
	// OK
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetTableColumnLineageByTaskIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type GetTableColumnLineageByTaskIdResponseBodyData struct {
	// example:
	//
	// 123
	InputBizUnitId *int64 `json:"InputBizUnitId,omitempty" xml:"InputBizUnitId,omitempty"`
	// example:
	//
	// odps.123.test_project.input_table.id
	InputColumnId *string `json:"InputColumnId,omitempty" xml:"InputColumnId,omitempty"`
	// example:
	//
	// id
	InputColumnName *string `json:"InputColumnName,omitempty" xml:"InputColumnName,omitempty"`
	// example:
	//
	// 123
	InputDataSourceId *int64 `json:"InputDataSourceId,omitempty" xml:"InputDataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	InputDataSourceType *string `json:"InputDataSourceType,omitempty" xml:"InputDataSourceType,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	InputDbType *string `json:"InputDbType,omitempty" xml:"InputDbType,omitempty"`
	// example:
	//
	// DEV
	InputEnv *string `json:"InputEnv,omitempty" xml:"InputEnv,omitempty"`
	// example:
	//
	// 123
	InputProjectId    *int64 `json:"InputProjectId,omitempty" xml:"InputProjectId,omitempty"`
	InputTableDeleted *bool  `json:"InputTableDeleted,omitempty" xml:"InputTableDeleted,omitempty"`
	// example:
	//
	// odps.123.test_project.order
	InputTableId *string `json:"InputTableId,omitempty" xml:"InputTableId,omitempty"`
	// example:
	//
	// order
	InputTableName *string `json:"InputTableName,omitempty" xml:"InputTableName,omitempty"`
	// example:
	//
	// PHYSICAL_TABLE
	InputTableType *string `json:"InputTableType,omitempty" xml:"InputTableType,omitempty"`
	// example:
	//
	// 123
	OutputBizUnitId *int64 `json:"OutputBizUnitId,omitempty" xml:"OutputBizUnitId,omitempty"`
	// example:
	//
	// odps.123.test_project.input_table.id
	OutputColumnId *string `json:"OutputColumnId,omitempty" xml:"OutputColumnId,omitempty"`
	// example:
	//
	// id
	OutputColumnName *string `json:"OutputColumnName,omitempty" xml:"OutputColumnName,omitempty"`
	// example:
	//
	// 123
	OutputDataSourceId *int64 `json:"OutputDataSourceId,omitempty" xml:"OutputDataSourceId,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	OutputDataSourceType *string `json:"OutputDataSourceType,omitempty" xml:"OutputDataSourceType,omitempty"`
	// example:
	//
	// MAX_COMPUTE
	OutputDbType *string `json:"OutputDbType,omitempty" xml:"OutputDbType,omitempty"`
	// example:
	//
	// DEV
	OutputEnv *string `json:"OutputEnv,omitempty" xml:"OutputEnv,omitempty"`
	// example:
	//
	// 123
	OutputProjectId    *int64 `json:"OutputProjectId,omitempty" xml:"OutputProjectId,omitempty"`
	OutputTableDeleted *bool  `json:"OutputTableDeleted,omitempty" xml:"OutputTableDeleted,omitempty"`
	// example:
	//
	// odps.123.test_project.order
	OutputTableId *string `json:"OutputTableId,omitempty" xml:"OutputTableId,omitempty"`
	// example:
	//
	// order
	OutputTableName *string `json:"OutputTableName,omitempty" xml:"OutputTableName,omitempty"`
	// example:
	//
	// PHYSICAL_TABLE
	OutputTableType *string `json:"OutputTableType,omitempty" xml:"OutputTableType,omitempty"`
	// example:
	//
	// DEV
	TaskEnv *string `json:"TaskEnv,omitempty" xml:"TaskEnv,omitempty"`
	// example:
	//
	// n_123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
