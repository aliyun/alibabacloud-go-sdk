// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDataTrackResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SearchDataTrackResultResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchDataTrackResultResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchDataTrackResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchDataTrackResultResponseBody
	GetSuccess() *bool
	SetTrackResult(v *SearchDataTrackResultResponseBodyTrackResult) *SearchDataTrackResultResponseBody
	GetTrackResult() *SearchDataTrackResultResponseBodyTrackResult
}

type SearchDataTrackResultResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The parsing result of the data tracking task.
	TrackResult *SearchDataTrackResultResponseBodyTrackResult `json:"TrackResult,omitempty" xml:"TrackResult,omitempty" type:"Struct"`
}

func (s SearchDataTrackResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchDataTrackResultResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchDataTrackResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchDataTrackResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchDataTrackResultResponseBody) GetTrackResult() *SearchDataTrackResultResponseBodyTrackResult {
	return s.TrackResult
}

func (s *SearchDataTrackResultResponseBody) SetErrorCode(v string) *SearchDataTrackResultResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchDataTrackResultResponseBody) SetErrorMessage(v string) *SearchDataTrackResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchDataTrackResultResponseBody) SetRequestId(v string) *SearchDataTrackResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDataTrackResultResponseBody) SetSuccess(v bool) *SearchDataTrackResultResponseBody {
	s.Success = &v
	return s
}

func (s *SearchDataTrackResultResponseBody) SetTrackResult(v *SearchDataTrackResultResponseBodyTrackResult) *SearchDataTrackResultResponseBody {
	s.TrackResult = v
	return s
}

func (s *SearchDataTrackResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchDataTrackResultResponseBodyTrackResult struct {
	// The details of the event logs.
	EventList []*SearchDataTrackResultResponseBodyTrackResultEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// The metadata of tables for which you track data operations.
	TableInfoList []*SearchDataTrackResultResponseBodyTrackResultTableInfoList `json:"TableInfoList,omitempty" xml:"TableInfoList,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 109
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchDataTrackResultResponseBodyTrackResult) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultResponseBodyTrackResult) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultResponseBodyTrackResult) GetEventList() []*SearchDataTrackResultResponseBodyTrackResultEventList {
	return s.EventList
}

func (s *SearchDataTrackResultResponseBodyTrackResult) GetTableInfoList() []*SearchDataTrackResultResponseBodyTrackResultTableInfoList {
	return s.TableInfoList
}

func (s *SearchDataTrackResultResponseBodyTrackResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchDataTrackResultResponseBodyTrackResult) SetEventList(v []*SearchDataTrackResultResponseBodyTrackResultEventList) *SearchDataTrackResultResponseBodyTrackResult {
	s.EventList = v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResult) SetTableInfoList(v []*SearchDataTrackResultResponseBodyTrackResultTableInfoList) *SearchDataTrackResultResponseBodyTrackResult {
	s.TableInfoList = v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResult) SetTotalCount(v int64) *SearchDataTrackResultResponseBodyTrackResult {
	s.TotalCount = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResult) Validate() error {
	return dara.Validate(s)
}

type SearchDataTrackResultResponseBodyTrackResultEventList struct {
	// The data records after you perform data operations in the database.
	DataAfter []*string `json:"DataAfter,omitempty" xml:"DataAfter,omitempty" type:"Repeated"`
	// The data records before you perform data operations in the database.
	DataBefore []*string `json:"DataBefore,omitempty" xml:"DataBefore,omitempty" type:"Repeated"`
	// The ID of the event.
	//
	// example:
	//
	// 1
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The length of the event content. Unit: bytes.
	//
	// example:
	//
	// 4324
	EventLength *int64 `json:"EventLength,omitempty" xml:"EventLength,omitempty"`
	// The event time.
	//
	// example:
	//
	// 2023-04-23 10:25:47
	EventTimestamp *string `json:"EventTimestamp,omitempty" xml:"EventTimestamp,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- **WRITE_ROWS**: indicates an INSERT operation.
	//
	// 	- **UPDATE_ROWS**: indicates an UPDATE operation.
	//
	// 	- **DELETE_ROWS**: indicates a DELETE operation.
	//
	// 	- **EXT_WRITE_ROWS**: indicates an INSERT operation, which is equivalent to WRITE_ROWS.
	//
	// 	- **EXT_UPDATE_ROWS**: indicates an UPDATE operation, which is equivalent to UPDATE_ROWS.
	//
	// 	- **EXT_DELETE_ROWS**: indicates a DELETE operation, which is equivalent to DELETE_ROWS.
	//
	// example:
	//
	// UPDATE_ROWS
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The SQL statements used to roll back the data change.
	//
	// example:
	//
	// -- Timestamp:2023-04-23 10:25:47 #1\\r\\nUPDATE `dc_test`.`tb_chunk_dml` SET `id`=1 , `gmt_create`=\\"2021-09-30T00:00:00\\" , `content`=\\"2023-03-30 14:51:50\\" , `c1`=\\"2023-04-17 13:42:03\\" , `c_id`=1 , `c13425`=\\"b\\\\\\"\\" , `c432532535`= null , `c1432`= null , `c143243253`= null , `c1432535`= null , `c43125325`= null , `c3425325`= null WHERE  (`id`=1)"
	RollSQL *string `json:"RollSQL,omitempty" xml:"RollSQL,omitempty"`
}

func (s SearchDataTrackResultResponseBodyTrackResultEventList) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultResponseBodyTrackResultEventList) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetDataAfter() []*string {
	return s.DataAfter
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetDataBefore() []*string {
	return s.DataBefore
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetEventId() *int64 {
	return s.EventId
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetEventLength() *int64 {
	return s.EventLength
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetEventTimestamp() *string {
	return s.EventTimestamp
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetEventType() *string {
	return s.EventType
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) GetRollSQL() *string {
	return s.RollSQL
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetDataAfter(v []*string) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.DataAfter = v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetDataBefore(v []*string) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.DataBefore = v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetEventId(v int64) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.EventId = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetEventLength(v int64) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.EventLength = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetEventTimestamp(v string) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.EventTimestamp = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetEventType(v string) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.EventType = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) SetRollSQL(v string) *SearchDataTrackResultResponseBodyTrackResultEventList {
	s.RollSQL = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultEventList) Validate() error {
	return dara.Validate(s)
}

type SearchDataTrackResultResponseBodyTrackResultTableInfoList struct {
	// The information about columns.
	Columns []*SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The description of the column.
	//
	// example:
	//
	// auto-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// prod_eb_vas
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// import_table_test1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s SearchDataTrackResultResponseBodyTrackResultTableInfoList) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultResponseBodyTrackResultTableInfoList) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) GetColumns() []*SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns {
	return s.Columns
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) GetDescription() *string {
	return s.Description
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) GetSchemaName() *string {
	return s.SchemaName
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) GetTableName() *string {
	return s.TableName
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) SetColumns(v []*SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) *SearchDataTrackResultResponseBodyTrackResultTableInfoList {
	s.Columns = v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) SetDescription(v string) *SearchDataTrackResultResponseBodyTrackResultTableInfoList {
	s.Description = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) SetSchemaName(v string) *SearchDataTrackResultResponseBodyTrackResultTableInfoList {
	s.SchemaName = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) SetTableName(v string) *SearchDataTrackResultResponseBodyTrackResultTableInfoList {
	s.TableName = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoList) Validate() error {
	return dara.Validate(s)
}

type SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns struct {
	// The name of the column.
	//
	// example:
	//
	// basic_platform
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The position of the column.
	//
	// example:
	//
	// 1
	ColumnPosition *int32 `json:"ColumnPosition,omitempty" xml:"ColumnPosition,omitempty"`
	// The data type of the column. Examples: BIGINT, INT, and VARCHAR.
	//
	// example:
	//
	// BIGINT
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// Indicates whether the column is a virtual column. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Fictive *bool `json:"Fictive,omitempty" xml:"Fictive,omitempty"`
}

func (s SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) String() string {
	return dara.Prettify(s)
}

func (s SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) GoString() string {
	return s.String()
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) GetColumnPosition() *int32 {
	return s.ColumnPosition
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) GetColumnType() *string {
	return s.ColumnType
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) GetFictive() *bool {
	return s.Fictive
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) SetColumnName(v string) *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns {
	s.ColumnName = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) SetColumnPosition(v int32) *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns {
	s.ColumnPosition = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) SetColumnType(v string) *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns {
	s.ColumnType = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) SetFictive(v bool) *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns {
	s.Fictive = &v
	return s
}

func (s *SearchDataTrackResultResponseBodyTrackResultTableInfoListColumns) Validate() error {
	return dara.Validate(s)
}
