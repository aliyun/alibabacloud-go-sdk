// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestDiagnosisResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRequestDiagnosisResultResponseBody
	GetCode() *string
	SetData(v *GetRequestDiagnosisResultResponseBodyData) *GetRequestDiagnosisResultResponseBody
	GetData() *GetRequestDiagnosisResultResponseBodyData
	SetMessage(v string) *GetRequestDiagnosisResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRequestDiagnosisResultResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetRequestDiagnosisResultResponseBody
	GetSuccess() *string
}

type GetRequestDiagnosisResultResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetRequestDiagnosisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// > This parameter returns `Successful` if the request succeeds. If the request fails, it returns an error message, which may include an error code.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 800FBAF5-A539-5B97-A09E-C63AB2F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request succeeded.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRequestDiagnosisResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRequestDiagnosisResultResponseBody) GetData() *GetRequestDiagnosisResultResponseBodyData {
	return s.Data
}

func (s *GetRequestDiagnosisResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRequestDiagnosisResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRequestDiagnosisResultResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetRequestDiagnosisResultResponseBody) SetCode(v string) *GetRequestDiagnosisResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetData(v *GetRequestDiagnosisResultResponseBodyData) *GetRequestDiagnosisResultResponseBody {
	s.Data = v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetMessage(v string) *GetRequestDiagnosisResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetRequestId(v string) *GetRequestDiagnosisResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) SetSuccess(v string) *GetRequestDiagnosisResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRequestDiagnosisResultResponseBodyData struct {
	// The user ID.
	//
	// example:
	//
	// 2093****
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The database name.
	//
	// example:
	//
	// das
	DbSchema *string `json:"dbSchema,omitempty" xml:"dbSchema,omitempty"`
	// The database engine. Valid values:
	//
	// - **MySQL**
	//
	// - **PostgreSQL**
	//
	// - **SQL Server**
	//
	// - **PolarDB-X**
	//
	// - **PolarDB for Oracle**
	//
	// - **MongoDB**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
	// The creation time of the SQL diagnosis, provided as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1633071840000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The last modification time of the SQL diagnosis, provided as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1633071850000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The unique ID of the diagnosis.
	//
	// example:
	//
	// 61820b594664275c4429****
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// Additional information.
	//
	// example:
	//
	// {"":""}
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	// The details of the SQL diagnosis result, returned as a JSON-formatted string.
	//
	// - **endTime**: The end time of the SQL diagnosis.
	//
	// - **errorCode**: The error code.
	//
	//   - **0001**: The diagnosis was successful.
	//
	//   - **0003**: The diagnosis failed.
	//
	// - **errorMessage**: The error message.
	//
	// - **estimateCost**: The estimated cost.
	//
	//   - **cpu**: The estimated CPU cost of the query.
	//
	//   - **io**: The estimated I/O cost of the query.
	//
	//   - **rows**: The estimated number of rows returned by the query.
	//
	// - **improvement**: The performance improvement ratio.
	//
	// - **indexAdvices**: The index suggestions.
	//
	//   - **columns**: The index columns.
	//
	//   - **ddlAddIndex**: The DDL statement for creating the index.
	//
	//   - **indexName**: The index name.
	//
	//   - **schemaName**: The schema name.
	//
	//   - **tableName**: The table name.
	//
	//   - **unique**: Indicates whether the index is a unique index.
	//
	// - **ip**: The instance IP address.
	//
	// - **messageId**: The diagnosis ID.
	//
	// - **port**: The instance port.
	//
	// - **sqlTag**: The SQL tags.
	//
	//   - **PRED_EQUAL**: Equality predicate.
	//
	//   - **CNT_QB**: Number of query blocks.
	//
	//   - **CNT_TB**: Number of tables.
	//
	//   - **JOIN_LEFT**: Left join.
	//
	//   - **SEL_SMALL**: Small result set selection.
	//
	//   - **AGGR_SEL**: Aggregate selection.
	//
	//   - **PRED_LT_EQ / PRED_GT_EQ**: Less-than-or-equal-to / greater-than-or-equal-to predicate.
	//
	//   - **PRED_LIKE_PREFIX**: LIKE prefix match.
	//
	//   - **ORDER_BY**: Contains an ORDER BY clause.
	//
	//   - **LIMIT**: Contains a LIMIT clause.
	//
	//   - **GROUP_BY**: Contains a GROUP BY clause.
	//
	//   - **JOIN_INNER**: Inner join.
	//
	//   - **JOIN_RIGHT**: Right join.
	//
	//   - **HAVING**: Contains a HAVING clause.
	//
	//   - **UNION**: Contains a UNION operation.
	//
	// - **startTime**: The start time of the SQL diagnosis.
	//
	// - **success**: Indicates whether the diagnosis was successful.
	//
	// - **support**: Indicates whether the SQL statement can be diagnosed.
	//
	//   - **true**: Supported.
	//
	//   - **false**: Not supported.
	//
	// - **tuningAdvices**: The SQL rewrite suggestions.
	//
	// example:
	//
	// { "endTime":1636354256000, "errorCode":"0001", "errorMessage":"TFX成功", "estimateCost":{ "cpu":1.7878745150389268, "io":9.948402604746128, "rows":8.889372575194633 }, "improvement":12933.97, "indexAdvices":[ { "columns":[ "work_no" ], "ddlAddIndex":"ALTER TABLE `test`.`work_order` ADD INDEX `idx_workno` (`work_no`)", "indexName":"idx_workno", "schemaName":"test", "tableName":"work_order", "unique":false } ], "ip":"****.mysql.rds.aliyuncs.com", "messageId":"6188c8cb2f1365b16aee****", "port":3306, "sqlTag":"{\\"PRED_EQUAL\\":\\"Y\\",\\"CNT_QB\\":\\"1\\",\\"CNT_TB\\":\\"1\\"}", "startTime":1636354252000, "success":true, "support":true, "tuningAdvices":[ ] }
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 0c95dae3afef77be06572612df9b****
	SqlId *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	// The diagnosis status. Valid values:
	//
	// - **0**: In progress.
	//
	// - **1**: Diagnosis error.
	//
	// - **2**: Completed.
	//
	// - **3**: SQL error.
	//
	// - **4**: Engine error.
	//
	// example:
	//
	// 2
	State *int32 `json:"state,omitempty" xml:"state,omitempty"`
	// The unique identifier of the diagnosed instance.
	//
	// example:
	//
	// hdm_51fe9bc19ec413f4d530431af87a****
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

func (s GetRequestDiagnosisResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetDbSchema() *string {
	return s.DbSchema
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetMessageId() *string {
	return s.MessageId
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetParam() *string {
	return s.Param
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetResult() *string {
	return s.Result
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetSqlId() *string {
	return s.SqlId
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetState() *int32 {
	return s.State
}

func (s *GetRequestDiagnosisResultResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetAccountId(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetDbSchema(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.DbSchema = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetEngine(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Engine = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetGmtCreate(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetGmtModified(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetMessageId(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.MessageId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetParam(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Param = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetResult(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Result = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetSqlId(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetState(v int32) *GetRequestDiagnosisResultResponseBodyData {
	s.State = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) SetUuid(v string) *GetRequestDiagnosisResultResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetRequestDiagnosisResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
