// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlStatementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSqlStatementResponseBodyData) *GetSqlStatementResponseBody
	GetData() *GetSqlStatementResponseBodyData
	SetRequestId(v string) *GetSqlStatementResponseBody
	GetRequestId() *string
}

type GetSqlStatementResponseBody struct {
	// The response parameters.
	Data *GetSqlStatementResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSqlStatementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSqlStatementResponseBody) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponseBody) GetData() *GetSqlStatementResponseBodyData {
	return s.Data
}

func (s *GetSqlStatementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSqlStatementResponseBody) SetData(v *GetSqlStatementResponseBodyData) *GetSqlStatementResponseBody {
	s.Data = v
	return s
}

func (s *GetSqlStatementResponseBody) SetRequestId(v string) *GetSqlStatementResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSqlStatementResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSqlStatementResponseBodyData struct {
	// The list of time that is consumed by SQL queries.
	ExecutionTime []*int64 `json:"executionTime,omitempty" xml:"executionTime,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// ERROR-102
	SqlErrorCode *string `json:"sqlErrorCode,omitempty" xml:"sqlErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// error message
	SqlErrorMessage *string `json:"sqlErrorMessage,omitempty" xml:"sqlErrorMessage,omitempty"`
	// The query results.
	SqlOutputs []*GetSqlStatementResponseBodyDataSqlOutputs `json:"sqlOutputs,omitempty" xml:"sqlOutputs,omitempty" type:"Repeated"`
	// The query status.
	//
	// Valid values:
	//
	// 	- running
	//
	// 	- available
	//
	// 	- cancelled
	//
	// 	- error
	//
	// 	- cancelling
	//
	// example:
	//
	// running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The query ID.
	//
	// example:
	//
	// st-1231311abadfaa
	StatementId *string `json:"statementId,omitempty" xml:"statementId,omitempty"`
}

func (s GetSqlStatementResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSqlStatementResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponseBodyData) GetExecutionTime() []*int64 {
	return s.ExecutionTime
}

func (s *GetSqlStatementResponseBodyData) GetSqlErrorCode() *string {
	return s.SqlErrorCode
}

func (s *GetSqlStatementResponseBodyData) GetSqlErrorMessage() *string {
	return s.SqlErrorMessage
}

func (s *GetSqlStatementResponseBodyData) GetSqlOutputs() []*GetSqlStatementResponseBodyDataSqlOutputs {
	return s.SqlOutputs
}

func (s *GetSqlStatementResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetSqlStatementResponseBodyData) GetStatementId() *string {
	return s.StatementId
}

func (s *GetSqlStatementResponseBodyData) SetExecutionTime(v []*int64) *GetSqlStatementResponseBodyData {
	s.ExecutionTime = v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetSqlErrorCode(v string) *GetSqlStatementResponseBodyData {
	s.SqlErrorCode = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetSqlErrorMessage(v string) *GetSqlStatementResponseBodyData {
	s.SqlErrorMessage = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetSqlOutputs(v []*GetSqlStatementResponseBodyDataSqlOutputs) *GetSqlStatementResponseBodyData {
	s.SqlOutputs = v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetState(v string) *GetSqlStatementResponseBodyData {
	s.State = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) SetStatementId(v string) *GetSqlStatementResponseBodyData {
	s.StatementId = &v
	return s
}

func (s *GetSqlStatementResponseBodyData) Validate() error {
	if s.SqlOutputs != nil {
		for _, item := range s.SqlOutputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSqlStatementResponseBodyDataSqlOutputs struct {
	// The queried data, which is a string in the JSON format.
	//
	// example:
	//
	// [{\\"values\\":[\\"test_db\\",\\"test_table\\",false]}
	Rows         *string `json:"rows,omitempty" xml:"rows,omitempty"`
	RowsFilePath *string `json:"rowsFilePath,omitempty" xml:"rowsFilePath,omitempty"`
	// The information about the schema, which is a string in the JSON format.
	//
	// example:
	//
	// {\\"type\\":\\"struct\\",\\"fields\\":[{\\"name\\":\\"namespace\\",\\"type\\":\\"string\\",\\"nullable\\":false,\\"metadata\\":{}},{\\"name\\":\\"tableName\\",\\"type\\":\\"string\\",\\"nullable\\":false,\\"metadata\\":{}},{\\"name\\":\\"isTemporary\\",\\"type\\":\\"boolean\\",\\"nullable\\":false,\\"metadata\\":{}}]}
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s GetSqlStatementResponseBodyDataSqlOutputs) String() string {
	return dara.Prettify(s)
}

func (s GetSqlStatementResponseBodyDataSqlOutputs) GoString() string {
	return s.String()
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) GetRows() *string {
	return s.Rows
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) GetRowsFilePath() *string {
	return s.RowsFilePath
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) GetSchema() *string {
	return s.Schema
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) SetRows(v string) *GetSqlStatementResponseBodyDataSqlOutputs {
	s.Rows = &v
	return s
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) SetRowsFilePath(v string) *GetSqlStatementResponseBodyDataSqlOutputs {
	s.RowsFilePath = &v
	return s
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) SetSchema(v string) *GetSqlStatementResponseBodyDataSqlOutputs {
	s.Schema = &v
	return s
}

func (s *GetSqlStatementResponseBodyDataSqlOutputs) Validate() error {
	return dara.Validate(s)
}
