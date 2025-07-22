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
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetRequestDiagnosisResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
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
	// 	- **true**
	//
	// 	- **false**
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
	return dara.Validate(s)
}

type GetRequestDiagnosisResultResponseBodyData struct {
	// The user ID.
	//
	// example:
	//
	// 2093****
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// das
	DbSchema *string `json:"dbSchema,omitempty" xml:"dbSchema,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// 	- **PolarDBMySQL**
	//
	// 	- **PolarDBOracle**
	//
	// 	- **MongoDB**
	//
	// example:
	//
	// MySQL
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
	// The time when the SQL diagnostics task was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1633071840000
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the SQL diagnostics task was modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1633071850000
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The unique ID of the diagnostics task.
	//
	// example:
	//
	// 61820b594664275c4429****
	MessageId *string `json:"messageId,omitempty" xml:"messageId,omitempty"`
	// The additional information.
	//
	// example:
	//
	// {"":""}
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	// The result of the SQL diagnostics task. The result includes the following information:
	//
	// 	- **endTime**: the end time of the SQL diagnostics task.
	//
	// 	- **errorCode**: the error code.
	//
	//     	- **0001**: The SQL diagnostics task is complete.
	//
	//     	- **0003**: The SQL diagnostics task failed.
	//
	// 	- **errorMessage**: the error message.
	//
	// 	- **estimateCost**: the estimated cost.
	//
	//     	- **cpu**: the estimated CPU utilization of the index.
	//
	//     	- **io**: the estimated I/O usage of the index.
	//
	//     	- **rows**: the estimated values of the rows returned for the index.
	//
	// 	- **improvement**: the performance improvement ratio.
	//
	// 	- **indexAdvices**: the index recommendations, which include the following information:
	//
	//     	- **columns**: the index columns.
	//
	//     	- **ddlAddIndex**: the DDL statement for the index.
	//
	//     	- **indexName**: the name of the index.
	//
	//     	- **schemaName**: the name of the database.
	//
	//     	- **tableName**: the name of the table.
	//
	//     	- **unique**: indicates whether the index is unique.
	//
	// 	- **ip**: the IP address of the instance.
	//
	// 	- **messageId**: the ID of the diagnostics task.
	//
	// 	- **port**: the port used to connect to the instance.
	//
	// 	- **sqlTag**: the SQL tag.
	//
	// 	- **startTime**: the start time of the SQL diagnostics task.
	//
	// 	- **success**: indicates whether the request was successful.
	//
	// 	- **support**: indicates whether the SQL statement can be diagnosed. Valid values:
	//
	//     	- **true**
	//
	//     	- **false**
	//
	// 	- **tuningAdvices*	- : the SQL rewrite suggestions.
	//
	// example:
	//
	// { "endTime":1636354256000, "errorCode":"0001", "errorMessage":"TFX succeeded", "estimateCost":{ "cpu":1.7878745150389268, "io":9.948402604746128, "rows":8.889372575194633 }, "improvement":12933.97, "indexAdvices":[ { "columns":[ "work_no" ], "ddlAddIndex":"ALTER TABLE `test`.`work_order` ADD INDEX `idx_workno` (`work_no`)", "indexName":"idx_workno", "schemaName":"test", "tableName":"work_order", "unique":false } ], "ip":"****.mysql.rds.aliyuncs.com", "messageId":"6188c8cb2f1365b16aee****", "port":3306, "sqlTag":"{\\"PRED_EQUAL\\":\\"Y\\",\\"CNT_QB\\":\\"1\\",\\"CNT_TB\\":\\"1\\"}", "startTime":1636354252000, "success":true, "support":true, "tuningAdvices":[ ] }
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 0c95dae3afef77be06572612df9b****
	SqlId *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	// The state of the diagnostics task. Valid values:
	//
	// 	- **0**: The diagnostics task is in progress.
	//
	// 	- **1**: A diagnostics error occurred.
	//
	// 	- **2**: The diagnostics task is complete.
	//
	// 	- **3**: An SQL error occurred.
	//
	// 	- **4**: An engine error occurred.
	//
	// example:
	//
	// 2
	State *int32 `json:"state,omitempty" xml:"state,omitempty"`
	// The unique ID of the diagnostics instance.
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
