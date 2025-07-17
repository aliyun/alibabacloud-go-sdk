// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataImportSQLPreCheckDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataImportSQLPreCheckDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataImportSQLPreCheckDetailResponseBody
	GetErrorMessage() *string
	SetPreCheckSQLDetailList(v []*ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) *ListDataImportSQLPreCheckDetailResponseBody
	GetPreCheckSQLDetailList() []*ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList
	SetRequestId(v string) *ListDataImportSQLPreCheckDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataImportSQLPreCheckDetailResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListDataImportSQLPreCheckDetailResponseBody
	GetTotalCount() *int64
}

type ListDataImportSQLPreCheckDetailResponseBody struct {
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
	// The precheck information of SQL statements.
	PreCheckSQLDetailList []*ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList `json:"PreCheckSQLDetailList,omitempty" xml:"PreCheckSQLDetailList,omitempty" type:"Repeated"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 31853A2B-DC9D-5B39-8492-D2AC8BCF550E
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
	// The number of SQL statements.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataImportSQLPreCheckDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLPreCheckDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) GetPreCheckSQLDetailList() []*ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList {
	return s.PreCheckSQLDetailList
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) SetErrorCode(v string) *ListDataImportSQLPreCheckDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) SetErrorMessage(v string) *ListDataImportSQLPreCheckDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) SetPreCheckSQLDetailList(v []*ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) *ListDataImportSQLPreCheckDetailResponseBody {
	s.PreCheckSQLDetailList = v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) SetRequestId(v string) *ListDataImportSQLPreCheckDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) SetSuccess(v bool) *ListDataImportSQLPreCheckDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) SetTotalCount(v int64) *ListDataImportSQLPreCheckDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList struct {
	// Indicates whether the precheck of the SQL statement was skipped. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Skip *bool `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The SQL ID, which indicates the sequence number of the SQL statement. The number starts with 1.
	//
	// example:
	//
	// 1
	SqlId *int64 `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	// The type of the SQL statement, such as DELETE, UPDATE, or ALTER_TABLE.
	//
	// example:
	//
	// INSERT
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The state of the ticket. Valid values:
	//
	// 	- **INIT**: The ticket was being initialized.
	//
	// 	- **RUNNING**: The ticket was in progress.
	//
	// 	- **SUCCESS**: The ticket was complete.
	//
	// 	- **TIMEOUT**: The ticket was skipped due to timeout.
	//
	// 	- **FAIL**: The ticket failed.
	//
	// example:
	//
	// SUCCESS
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) GetSkip() *bool {
	return s.Skip
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) GetSqlId() *int64 {
	return s.SqlId
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) GetSqlType() *string {
	return s.SqlType
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) SetSkip(v bool) *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList {
	s.Skip = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) SetSqlId(v int64) *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList {
	s.SqlId = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) SetSqlType(v string) *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList {
	s.SqlType = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) SetStatusCode(v string) *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList {
	s.StatusCode = &v
	return s
}

func (s *ListDataImportSQLPreCheckDetailResponseBodyPreCheckSQLDetailList) Validate() error {
	return dara.Validate(s)
}
