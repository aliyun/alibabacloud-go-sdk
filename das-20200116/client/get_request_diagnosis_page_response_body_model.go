// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRequestDiagnosisPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRequestDiagnosisPageResponseBody
	GetCode() *string
	SetData(v *GetRequestDiagnosisPageResponseBodyData) *GetRequestDiagnosisPageResponseBody
	GetData() *GetRequestDiagnosisPageResponseBodyData
	SetMessage(v string) *GetRequestDiagnosisPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRequestDiagnosisPageResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetRequestDiagnosisPageResponseBody
	GetSuccess() *string
}

type GetRequestDiagnosisPageResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetRequestDiagnosisPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message that contains information such as an error code is returned.
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

func (s GetRequestDiagnosisPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRequestDiagnosisPageResponseBody) GetData() *GetRequestDiagnosisPageResponseBodyData {
	return s.Data
}

func (s *GetRequestDiagnosisPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRequestDiagnosisPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRequestDiagnosisPageResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetRequestDiagnosisPageResponseBody) SetCode(v string) *GetRequestDiagnosisPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetData(v *GetRequestDiagnosisPageResponseBodyData) *GetRequestDiagnosisPageResponseBody {
	s.Data = v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetMessage(v string) *GetRequestDiagnosisPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetRequestId(v string) *GetRequestDiagnosisPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) SetSuccess(v string) *GetRequestDiagnosisPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRequestDiagnosisPageResponseBodyData struct {
	// Additional information.
	//
	// example:
	//
	// {"":""}
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The SQL diagnostics records returned.
	List []*GetRequestDiagnosisPageResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// The number of entries per page. The value must be a positive integer. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetRequestDiagnosisPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponseBodyData) GetExtra() *string {
	return s.Extra
}

func (s *GetRequestDiagnosisPageResponseBodyData) GetList() []*GetRequestDiagnosisPageResponseBodyDataList {
	return s.List
}

func (s *GetRequestDiagnosisPageResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetRequestDiagnosisPageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRequestDiagnosisPageResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetExtra(v string) *GetRequestDiagnosisPageResponseBodyData {
	s.Extra = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetList(v []*GetRequestDiagnosisPageResponseBodyDataList) *GetRequestDiagnosisPageResponseBodyData {
	s.List = v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetPageNo(v int64) *GetRequestDiagnosisPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetPageSize(v int64) *GetRequestDiagnosisPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) SetTotal(v int64) *GetRequestDiagnosisPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRequestDiagnosisPageResponseBodyDataList struct {
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
	// The time when the SQL diagnostics task was created. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
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
	// Additional information.
	//
	// example:
	//
	// {"":""}
	Param *string `json:"param,omitempty" xml:"param,omitempty"`
	// The result of the SQL diagnostics task. The result includes the following information:
	//
	// 	- **endTime**: the end time of the SQL diagnostics task.
	//
	// 	- **errorCode**: indicates whether the SQL diagnostics task is complete. Valid values:
	//
	//   	- **0001**: The SQL diagnostics task is complete.
	//
	//   	- **0003**: The SQL diagnostics task failed.
	//
	// 	- **errorMessage**: the error message.
	//
	// 	- **estimateCost**: the estimated cost.
	//
	//   	- **cpu**: the estimated CPU utilization of the index.
	//
	//   	- **io**: the estimated I/O usage of the index.
	//
	//   	- **rows**: the estimated values of the rows returned for the index.
	//
	// 	- **improvement**: the performance improvement ratio.
	//
	// 	- **indexAdvices**: the index recommendations, which include the following information:
	//
	//   	- **columns**: the index columns.
	//
	//   	- **ddlAddIndex**: the DDL statement for the index.
	//
	//   	- **indexName**: the name of the index.
	//
	//   	- **schemaName**: the name of the database.
	//
	//   	- **tableName**: the name of the table.
	//
	//   	- **unique**: indicates whether the index is unique.
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
	//   	- **true**: The SQL statement can be diagnosed.
	//
	//   	- **false**: The SQL statement cannot be diagnosed.
	//
	// 	- **tuningAdvices**: the SQL rewrite suggestions.
	//
	// example:
	//
	// { "endTime":1636354256000, "errorCode":"0001", "errorMessage":"TFX Successful", "estimateCost":{ "cpu":1.7878745150389268, "io":9.948402604746128, "rows":8.889372575194633 }, "improvement":12933.97, "indexAdvices":[ { "columns":[ "work_no" ], "ddlAddIndex":"ALTER TABLE `test`.`work_order` ADD INDEX `idx_workno` (`work_no`)", "indexName":"idx_workno", "schemaName":"test", "tableName":"work_order", "unique":false } ], "ip":"****.mysql.rds.aliyuncs.com", "messageId":"6188c8cb2f1365b16aee****", "port":3306, "sqlTag":"{\\"PRED_EQUAL\\":\\"Y\\",\\"CNT_QB\\":\\"1\\",\\"CNT_TB\\":\\"1\\"}", "startTime":1636354252000, "success":true, "support":true, "tuningAdvices":[ ] }
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The SQL template ID.
	//
	// example:
	//
	// 0c95dae3afef77be06572612df9b****
	SqlId *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	// The status of the diagnostics task. Valid values:
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

func (s GetRequestDiagnosisPageResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetRequestDiagnosisPageResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetAccountId() *string {
	return s.AccountId
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetDbSchema() *string {
	return s.DbSchema
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetEngine() *string {
	return s.Engine
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetMessageId() *string {
	return s.MessageId
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetParam() *string {
	return s.Param
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetResult() *string {
	return s.Result
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetSqlId() *string {
	return s.SqlId
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetState() *int32 {
	return s.State
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) GetUuid() *string {
	return s.Uuid
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetAccountId(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.AccountId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetDbSchema(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.DbSchema = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetEngine(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Engine = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetGmtCreate(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetGmtModified(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetMessageId(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.MessageId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetParam(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Param = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetResult(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Result = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetSqlId(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.SqlId = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetState(v int32) *GetRequestDiagnosisPageResponseBodyDataList {
	s.State = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) SetUuid(v string) *GetRequestDiagnosisPageResponseBodyDataList {
	s.Uuid = &v
	return s
}

func (s *GetRequestDiagnosisPageResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
