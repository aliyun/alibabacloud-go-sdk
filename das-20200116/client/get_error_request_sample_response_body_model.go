// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorRequestSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetErrorRequestSampleResponseBody
	GetCode() *int64
	SetData(v []*GetErrorRequestSampleResponseBodyData) *GetErrorRequestSampleResponseBody
	GetData() []*GetErrorRequestSampleResponseBodyData
	SetMessage(v string) *GetErrorRequestSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetErrorRequestSampleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetErrorRequestSampleResponseBody
	GetSuccess() *bool
}

type GetErrorRequestSampleResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// [         {             "sqlId": "2cd4432556c3dab9d825ba363637****",             "database": "dbgateway",             "originHost": "172.16.1****",             "tables": [                 "meter_****"             ],             "instanceId": "rm-2ze8g2am97624****",             "errorCode": "1062",             "user": "dbgat****",             "sql": "insert into meter_****\\n        ( \\n        ****\\n     )\\n        values (now(), now(), \\"bbbc8624-5e19-455a-9714-8466f688****\\", \\"2022-02-10 14:00:00\\", \\"{\\"endTime\\":\\"2022-02-10 14:00:00\\",\\"endTimestamp\\":1644472800,\\"startTime\\":\\"2022-02-10 13:00:00\\",\\"startTimestamp\\":1644469200}\\", null, null)",             "timestamp": 1644476100435         }]
	Data []*GetErrorRequestSampleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7172BECE-588A-5961-8126-C216E16B****
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
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetErrorRequestSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErrorRequestSampleResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetErrorRequestSampleResponseBody) GetData() []*GetErrorRequestSampleResponseBodyData {
	return s.Data
}

func (s *GetErrorRequestSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetErrorRequestSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErrorRequestSampleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetErrorRequestSampleResponseBody) SetCode(v int64) *GetErrorRequestSampleResponseBody {
	s.Code = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetData(v []*GetErrorRequestSampleResponseBodyData) *GetErrorRequestSampleResponseBody {
	s.Data = v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetMessage(v string) *GetErrorRequestSampleResponseBody {
	s.Message = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetRequestId(v string) *GetErrorRequestSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) SetSuccess(v bool) *GetErrorRequestSampleResponseBody {
	s.Success = &v
	return s
}

func (s *GetErrorRequestSampleResponseBody) Validate() error {
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

type GetErrorRequestSampleResponseBodyData struct {
	// The database name.
	//
	// example:
	//
	// dbgateway
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// 1062
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The IP address of the client that executes the SQL statement.
	//
	// example:
	//
	// 172.16.1****
	OriginHost *string `json:"originHost,omitempty" xml:"originHost,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// insert into meter_****
	Sql *string `json:"sql,omitempty" xml:"sql,omitempty"`
	// The SQL query ID.
	//
	// example:
	//
	// 2cd4432556c3dab9d825ba363637****
	SqlId *string `json:"sqlId,omitempty" xml:"sqlId,omitempty"`
	// The table information.
	Tables []*string `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// The time when the SQL query was executed. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1644476100435
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// The username of the account that is used to log on to the database.
	//
	// example:
	//
	// dbgat****
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s GetErrorRequestSampleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetErrorRequestSampleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetErrorRequestSampleResponseBodyData) GetDatabase() *string {
	return s.Database
}

func (s *GetErrorRequestSampleResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetErrorRequestSampleResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetErrorRequestSampleResponseBodyData) GetOriginHost() *string {
	return s.OriginHost
}

func (s *GetErrorRequestSampleResponseBodyData) GetSql() *string {
	return s.Sql
}

func (s *GetErrorRequestSampleResponseBodyData) GetSqlId() *string {
	return s.SqlId
}

func (s *GetErrorRequestSampleResponseBodyData) GetTables() []*string {
	return s.Tables
}

func (s *GetErrorRequestSampleResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetErrorRequestSampleResponseBodyData) GetUser() *string {
	return s.User
}

func (s *GetErrorRequestSampleResponseBodyData) SetDatabase(v string) *GetErrorRequestSampleResponseBodyData {
	s.Database = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetErrorCode(v string) *GetErrorRequestSampleResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetInstanceId(v string) *GetErrorRequestSampleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetOriginHost(v string) *GetErrorRequestSampleResponseBodyData {
	s.OriginHost = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetSql(v string) *GetErrorRequestSampleResponseBodyData {
	s.Sql = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetSqlId(v string) *GetErrorRequestSampleResponseBodyData {
	s.SqlId = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetTables(v []*string) *GetErrorRequestSampleResponseBodyData {
	s.Tables = v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetTimestamp(v int64) *GetErrorRequestSampleResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) SetUser(v string) *GetErrorRequestSampleResponseBodyData {
	s.User = &v
	return s
}

func (s *GetErrorRequestSampleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
