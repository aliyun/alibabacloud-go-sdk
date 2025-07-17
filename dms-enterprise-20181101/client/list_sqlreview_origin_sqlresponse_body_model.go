// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSQLReviewOriginSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListSQLReviewOriginSQLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListSQLReviewOriginSQLResponseBody
	GetErrorMessage() *string
	SetOriginSQLList(v []*ListSQLReviewOriginSQLResponseBodyOriginSQLList) *ListSQLReviewOriginSQLResponseBody
	GetOriginSQLList() []*ListSQLReviewOriginSQLResponseBodyOriginSQLList
	SetRequestId(v string) *ListSQLReviewOriginSQLResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSQLReviewOriginSQLResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSQLReviewOriginSQLResponseBody
	GetTotalCount() *int32
}

type ListSQLReviewOriginSQLResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The information about the parsed SQL statements.
	OriginSQLList []*ListSQLReviewOriginSQLResponseBodyOriginSQLList `json:"OriginSQLList,omitempty" xml:"OriginSQLList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of SQL statements in the file.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSQLReviewOriginSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLResponseBody) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSQLReviewOriginSQLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListSQLReviewOriginSQLResponseBody) GetOriginSQLList() []*ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	return s.OriginSQLList
}

func (s *ListSQLReviewOriginSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSQLReviewOriginSQLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSQLReviewOriginSQLResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSQLReviewOriginSQLResponseBody) SetErrorCode(v string) *ListSQLReviewOriginSQLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetErrorMessage(v string) *ListSQLReviewOriginSQLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetOriginSQLList(v []*ListSQLReviewOriginSQLResponseBodyOriginSQLList) *ListSQLReviewOriginSQLResponseBody {
	s.OriginSQLList = v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetRequestId(v string) *ListSQLReviewOriginSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetSuccess(v bool) *ListSQLReviewOriginSQLResponseBody {
	s.Success = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) SetTotalCount(v int32) *ListSQLReviewOriginSQLResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSQLReviewOriginSQLResponseBodyOriginSQLList struct {
	// The review status of the SQL statement. Valid values:
	//
	// 	- **new**: The SQL statement was waiting to be reviewed.
	//
	// 	- **unknown**: The SQL statement cannot be parsed.
	//
	// 	- **check_not_pass**: The SQL statement failed to pass the review.
	//
	// 	- **check_pass**: The SQL statement passed the review.
	//
	// 	- **force_pass**: The SQL statement passed the manual review.
	//
	// 	- **force_not_pass**: The SQL statement failed to pass the manual review.
	//
	// example:
	//
	// check_pass
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The time when the SQL statement was reviewed.
	//
	// example:
	//
	// 2021-06-09 21:07:00
	CheckedTime *string `json:"CheckedTime,omitempty" xml:"CheckedTime,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 123321
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the file.
	//
	// example:
	//
	// test.sql
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The statistics on the optimization suggestions for SQL statements. The value is a JSON string. Valid values:
	//
	// 	- **MUST_IMPROVE**: The SQL statements must be optimized.
	//
	// 	- **POTENTIAL_ISSUE**: The SQL statements contain potential issues.
	//
	// 	- **SUGGEST_IMPROVE**: We recommend that you optimize the SQL statements.
	//
	// 	- **USEDMSTOOLKIT**: We recommend that you change schemas without locking tables.
	//
	// 	- **USEDMSDML_UNLOCK**: We recommend that you change data without locking tables.
	//
	// 	- **TABLEINDEXSUGGEST**: We recommend that you optimize indexes for the SQL statements.
	//
	// example:
	//
	// {"POTENTIAL_ISSUE":1,"SUGGEST_IMPROVE":1}
	ReviewSummary *string `json:"ReviewSummary,omitempty" xml:"ReviewSummary,omitempty"`
	// The SQL statement in the file.
	//
	// example:
	//
	// select id from table_name
	SQLContent *string `json:"SQLContent,omitempty" xml:"SQLContent,omitempty"`
	// The ID of the SQL statement.
	//
	// example:
	//
	// 1111
	SQLId *int64 `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	// The name of the SQL statement.
	//
	// example:
	//
	// getByPk
	SQLName *string `json:"SQLName,omitempty" xml:"SQLName,omitempty"`
	// The key that is used to query the information about optimization suggestions. You can call the [GetSQLReviewOptimizeDetail](https://help.aliyun.com/document_detail/465919.html) operation to query the details based on this key.
	//
	// example:
	//
	// a57e54ec5433475ea3082d882fdb89c5
	SQLReviewQueryKey *string `json:"SQLReviewQueryKey,omitempty" xml:"SQLReviewQueryKey,omitempty"`
	// The MD5 hash value that is obtained after the SQL statement is calculated by using a hash algorithm.
	//
	// example:
	//
	// 95adb6e77a0884d9e50232cb8c5c969d
	SqlHash *string `json:"SqlHash,omitempty" xml:"SqlHash,omitempty"`
	// The description of the review status.
	//
	// example:
	//
	// passed the test
	StatusDesc *string `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
}

func (s ListSQLReviewOriginSQLResponseBodyOriginSQLList) String() string {
	return dara.Prettify(s)
}

func (s ListSQLReviewOriginSQLResponseBodyOriginSQLList) GoString() string {
	return s.String()
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetCheckedTime() *string {
	return s.CheckedTime
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetFileId() *int64 {
	return s.FileId
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetFileName() *string {
	return s.FileName
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetReviewSummary() *string {
	return s.ReviewSummary
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetSQLContent() *string {
	return s.SQLContent
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetSQLId() *int64 {
	return s.SQLId
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetSQLName() *string {
	return s.SQLName
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetSQLReviewQueryKey() *string {
	return s.SQLReviewQueryKey
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetSqlHash() *string {
	return s.SqlHash
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetCheckStatus(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.CheckStatus = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetCheckedTime(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.CheckedTime = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetFileId(v int64) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.FileId = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetFileName(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.FileName = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetReviewSummary(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.ReviewSummary = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLContent(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLContent = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLId(v int64) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLId = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLName(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLName = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSQLReviewQueryKey(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SQLReviewQueryKey = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetSqlHash(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.SqlHash = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) SetStatusDesc(v string) *ListSQLReviewOriginSQLResponseBodyOriginSQLList {
	s.StatusDesc = &v
	return s
}

func (s *ListSQLReviewOriginSQLResponseBodyOriginSQLList) Validate() error {
	return dara.Validate(s)
}
