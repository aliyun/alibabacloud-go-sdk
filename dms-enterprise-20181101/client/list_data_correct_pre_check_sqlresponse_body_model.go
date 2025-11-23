// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCorrectPreCheckSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataCorrectPreCheckSQLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataCorrectPreCheckSQLResponseBody
	GetErrorMessage() *string
	SetPreCheckSQLList(v []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) *ListDataCorrectPreCheckSQLResponseBody
	GetPreCheckSQLList() []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList
	SetRequestId(v string) *ListDataCorrectPreCheckSQLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataCorrectPreCheckSQLResponseBody
	GetSuccess() *bool
}

type ListDataCorrectPreCheckSQLResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The precheck information about SQL statements.
	PreCheckSQLList []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList `json:"PreCheckSQLList,omitempty" xml:"PreCheckSQLList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 31853A2B-DC9D-5B39-8492-D2AC8BCF550E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataCorrectPreCheckSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataCorrectPreCheckSQLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataCorrectPreCheckSQLResponseBody) GetPreCheckSQLList() []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	return s.PreCheckSQLList
}

func (s *ListDataCorrectPreCheckSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCorrectPreCheckSQLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetErrorCode(v string) *ListDataCorrectPreCheckSQLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetErrorMessage(v string) *ListDataCorrectPreCheckSQLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetPreCheckSQLList(v []*ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) *ListDataCorrectPreCheckSQLResponseBody {
	s.PreCheckSQLList = v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetRequestId(v string) *ListDataCorrectPreCheckSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) SetSuccess(v bool) *ListDataCorrectPreCheckSQLResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBody) Validate() error {
	if s.PreCheckSQLList != nil {
		for _, item := range s.PreCheckSQLList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList struct {
	// The estimated number of affected rows.
	//
	// example:
	//
	// 0
	AffectRows *int64 `json:"AffectRows,omitempty" xml:"AffectRows,omitempty"`
	// The SQL statement.
	//
	// example:
	//
	// OPTIMIZE TABLE `Text_TableNames`
	CheckSQL *string `json:"CheckSQL,omitempty" xml:"CheckSQL,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 1930****
	DbId *int64 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The key that is used to query the details of optimization suggestions. You can call the [GetSQLReviewOptimizeDetail](https://help.aliyun.com/document_detail/265977.html) operation to query the details of optimization suggestions based on the key.
	//
	// example:
	//
	// b9e771fc6ec247dea6d06a32c777****
	SQLReviewQueryKey *string `json:"SQLReviewQueryKey,omitempty" xml:"SQLReviewQueryKey,omitempty"`
	// The review status of the SQL statement. Valid values:
	//
	// 	- **WAITING**: The SQL statement is pending for review.
	//
	// 	- **RUNNING**: The SQL statement is being reviewed.
	//
	// 	- **IGNORE**: The SQL statement review is skipped.
	//
	// 	- **PASS**: The SQL statement passed the review.
	//
	// 	- **BLOCK**: The SQL statement failed the review.
	//
	// example:
	//
	// WAITING
	SqlReviewStatus *string `json:"SqlReviewStatus,omitempty" xml:"SqlReviewStatus,omitempty"`
	// The type of the SQL statement, such as DELETE, UPDATE, or ALTER_TABLE.
	//
	// example:
	//
	// OPTIMIZE
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	// The name of the table whose data is changed.
	//
	// example:
	//
	// Text_TableNames
	TableNames *string `json:"TableNames,omitempty" xml:"TableNames,omitempty"`
}

func (s ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetAffectRows() *int64 {
	return s.AffectRows
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetCheckSQL() *string {
	return s.CheckSQL
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetDbId() *int64 {
	return s.DbId
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetSQLReviewQueryKey() *string {
	return s.SQLReviewQueryKey
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetSqlReviewStatus() *string {
	return s.SqlReviewStatus
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetSqlType() *string {
	return s.SqlType
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) GetTableNames() *string {
	return s.TableNames
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetAffectRows(v int64) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.AffectRows = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetCheckSQL(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.CheckSQL = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetDbId(v int64) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.DbId = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetSQLReviewQueryKey(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.SQLReviewQueryKey = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetSqlReviewStatus(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.SqlReviewStatus = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetSqlType(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.SqlType = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) SetTableNames(v string) *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList {
	s.TableNames = &v
	return s
}

func (s *ListDataCorrectPreCheckSQLResponseBodyPreCheckSQLList) Validate() error {
	return dara.Validate(s)
}
