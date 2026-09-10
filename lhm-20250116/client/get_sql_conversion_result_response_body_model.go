// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlConversionResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetSqlConversionResultResponseBodyData) *GetSqlConversionResultResponseBody
	GetData() []*GetSqlConversionResultResponseBodyData
	SetErrCode(v string) *GetSqlConversionResultResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetSqlConversionResultResponseBody
	GetErrMessage() *string
	SetPageIndex(v int32) *GetSqlConversionResultResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *GetSqlConversionResultResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetSqlConversionResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSqlConversionResultResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetSqlConversionResultResponseBody
	GetTotalCount() *int32
}

type GetSqlConversionResultResponseBody struct {
	// The data list returned by the operation. For the structure of each element, see the child parameters.
	Data []*GetSqlConversionResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size, which is the number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues with this call.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records that meet the query conditions. This value is used for pagination.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetSqlConversionResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlConversionResultResponseBody) GetData() []*GetSqlConversionResultResponseBodyData {
	return s.Data
}

func (s *GetSqlConversionResultResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetSqlConversionResultResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetSqlConversionResultResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetSqlConversionResultResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSqlConversionResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlConversionResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSqlConversionResultResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetSqlConversionResultResponseBody) SetData(v []*GetSqlConversionResultResponseBodyData) *GetSqlConversionResultResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetErrCode(v string) *GetSqlConversionResultResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetErrMessage(v string) *GetSqlConversionResultResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetPageIndex(v int32) *GetSqlConversionResultResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetPageSize(v int32) *GetSqlConversionResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetRequestId(v string) *GetSqlConversionResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetSuccess(v bool) *GetSqlConversionResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) SetTotalCount(v int32) *GetSqlConversionResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetSqlConversionResultResponseBody) Validate() error {
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

type GetSqlConversionResultResponseBodyData struct {
	// The error reason.
	//
	// example:
	//
	// connection timeout
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The completion time.
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
	// The script conversion status. In conversion job scenarios: pass for conversion succeeded, turning for converting, and fail for conversion failed. In some scenarios: success for succeeded, failed for failed, and skipped for skipped.
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
	// The table name mapping.
	TableMappingList []*GetSqlConversionResultResponseBodyDataTableMappingList `json:"tableMappingList,omitempty" xml:"tableMappingList,omitempty" type:"Repeated"`
}

func (s GetSqlConversionResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlConversionResultResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSqlConversionResultResponseBodyData) GetFinishTime() *string {
	return s.FinishTime
}

func (s *GetSqlConversionResultResponseBodyData) GetScriptId() *int64 {
	return s.ScriptId
}

func (s *GetSqlConversionResultResponseBodyData) GetScriptName() *string {
	return s.ScriptName
}

func (s *GetSqlConversionResultResponseBodyData) GetScriptTransformStatus() *string {
	return s.ScriptTransformStatus
}

func (s *GetSqlConversionResultResponseBodyData) GetSqlResultContent() *string {
	return s.SqlResultContent
}

func (s *GetSqlConversionResultResponseBodyData) GetSqlSourceContent() *string {
	return s.SqlSourceContent
}

func (s *GetSqlConversionResultResponseBodyData) GetTableMappingList() []*GetSqlConversionResultResponseBodyDataTableMappingList {
	return s.TableMappingList
}

func (s *GetSqlConversionResultResponseBodyData) SetErrorMessage(v string) *GetSqlConversionResultResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetFinishTime(v string) *GetSqlConversionResultResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetScriptId(v int64) *GetSqlConversionResultResponseBodyData {
	s.ScriptId = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetScriptName(v string) *GetSqlConversionResultResponseBodyData {
	s.ScriptName = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetScriptTransformStatus(v string) *GetSqlConversionResultResponseBodyData {
	s.ScriptTransformStatus = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetSqlResultContent(v string) *GetSqlConversionResultResponseBodyData {
	s.SqlResultContent = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetSqlSourceContent(v string) *GetSqlConversionResultResponseBodyData {
	s.SqlSourceContent = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) SetTableMappingList(v []*GetSqlConversionResultResponseBodyDataTableMappingList) *GetSqlConversionResultResponseBodyData {
	s.TableMappingList = v
	return s
}

func (s *GetSqlConversionResultResponseBodyData) Validate() error {
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

type GetSqlConversionResultResponseBodyDataTableMappingList struct {
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

func (s GetSqlConversionResultResponseBodyDataTableMappingList) String() string {
	return dara.Prettify(s)
}

func (s GetSqlConversionResultResponseBodyDataTableMappingList) GoString() string {
	return s.String()
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetId() *int64 {
	return s.Id
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetSourceSchema() *string {
	return s.SourceSchema
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetTargetTableName() *string {
	return s.TargetTableName
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetTargetType() *string {
	return s.TargetType
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) GetUid() *string {
	return s.Uid
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetId(v int64) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.Id = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetSourceSchema(v string) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.SourceSchema = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetSourceTableName(v string) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.SourceTableName = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetTargetTableName(v string) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.TargetTableName = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetTargetType(v string) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.TargetType = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetTaskId(v int64) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.TaskId = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetTenantId(v string) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.TenantId = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) SetUid(v string) *GetSqlConversionResultResponseBodyDataTableMappingList {
	s.Uid = &v
	return s
}

func (s *GetSqlConversionResultResponseBodyDataTableMappingList) Validate() error {
	return dara.Validate(s)
}
