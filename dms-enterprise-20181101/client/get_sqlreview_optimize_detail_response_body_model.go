// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSQLReviewOptimizeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetSQLReviewOptimizeDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetSQLReviewOptimizeDetailResponseBody
	GetErrorMessage() *string
	SetOptimizeDetail(v *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) *GetSQLReviewOptimizeDetailResponseBody
	GetOptimizeDetail() *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail
	SetRequestId(v string) *GetSQLReviewOptimizeDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSQLReviewOptimizeDetailResponseBody
	GetSuccess() *bool
}

type GetSQLReviewOptimizeDetailResponseBody struct {
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
	// The details of optimization suggestions for SQL statements.
	OptimizeDetail *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail `json:"OptimizeDetail,omitempty" xml:"OptimizeDetail,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetSQLReviewOptimizeDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSQLReviewOptimizeDetailResponseBody) GetOptimizeDetail() *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	return s.OptimizeDetail
}

func (s *GetSQLReviewOptimizeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSQLReviewOptimizeDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetErrorCode(v string) *GetSQLReviewOptimizeDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetErrorMessage(v string) *GetSQLReviewOptimizeDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetOptimizeDetail(v *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) *GetSQLReviewOptimizeDetailResponseBody {
	s.OptimizeDetail = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetRequestId(v string) *GetSQLReviewOptimizeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) SetSuccess(v bool) *GetSQLReviewOptimizeDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBody) Validate() error {
	if s.OptimizeDetail != nil {
		if err := s.OptimizeDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail struct {
	// The ID of the database.
	//
	// example:
	//
	// 111222
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The ID of the instance to which the database belongs.
	//
	// example:
	//
	// 123321
	InstanceId *int32 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The quality of the SQL statement.
	QualityResult *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult `json:"QualityResult,omitempty" xml:"QualityResult,omitempty" type:"Struct"`
	// The key that is used to query the details of optimization suggestions.
	//
	// example:
	//
	// a57e54ec5433475ea3082d882fdb****
	QueryKey *string `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
	// The type of the SQL statement. Valid values: DELETE, UPDATE, and ALTER_TABLE.
	//
	// example:
	//
	// UPDATE
	SqlType *string `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GetDbId() *int32 {
	return s.DbId
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GetInstanceId() *int32 {
	return s.InstanceId
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GetQualityResult() *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	return s.QualityResult
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GetQueryKey() *string {
	return s.QueryKey
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) GetSqlType() *string {
	return s.SqlType
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetDbId(v int32) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.DbId = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetInstanceId(v int32) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.InstanceId = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetQualityResult(v *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.QualityResult = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetQueryKey(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.QueryKey = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) SetSqlType(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail {
	s.SqlType = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetail) Validate() error {
	if s.QualityResult != nil {
		if err := s.QualityResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult struct {
	// The error message returned.
	//
	// example:
	//
	// syntax error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether an error occurs. Valid values:
	//
	// 	- **true**: An error occurs.
	//
	// 	- **false**: No error occurs.
	//
	// example:
	//
	// false
	OccurError *bool `json:"OccurError,omitempty" xml:"OccurError,omitempty"`
	// The review results based on rules.
	Results []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) GetOccurError() *bool {
	return s.OccurError
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) GetResults() []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	return s.Results
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) SetErrorMessage(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	s.ErrorMessage = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) SetOccurError(v bool) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	s.OccurError = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) SetResults(v []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult {
	s.Results = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResult) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults struct {
	// The comment that is specified when you create the SQL review rule. For more information, see [SQL review optimization](https://help.aliyun.com/document_detail/194114.html).
	//
	// example:
	//
	// xxx business rule: the query must have a where condition.
	Comments *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// The optimization suggestion for the SQL statement. Valid values:
	//
	// 	- **MUST_IMPROVE**: The SQL statement must be improved.
	//
	// 	- **POTENTIAL_ISSUE**: The SQL statement contains potential issues.
	//
	// 	- **SUGGEST_IMPROVE**: We recommend that you improve the SQL statement.
	//
	// 	- **USEDMSTOOLKIT**: We recommend that you change schemas without locking tables.
	//
	// 	- **USEDMSDML_UNLOCK**: We recommend that you change data without locking tables.
	//
	// 	- **TABLEINDEXSUGGEST**: We recommend that you use SQL statements that use indexes.
	//
	// example:
	//
	// MUST_IMPROVE
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The review results.
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The name of the rule. For more information, see [SQL review optimization](https://help.aliyun.com/document_detail/194114.html).
	//
	// example:
	//
	// SELECT_SUGGEST_ASSIGN_WHERE
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the SQL review rule. Valid values:
	//
	// 	- **REVIEW**: a rule that is used to review SQL statements based on standards.
	//
	// 	- **OPTIMIZE**: a rule that is used to provide optimization suggestions.
	//
	// example:
	//
	// REVIEW
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The SQL script for data changes.
	Scripts []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts `json:"Scripts,omitempty" xml:"Scripts,omitempty" type:"Repeated"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GetComments() *string {
	return s.Comments
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GetFeedback() *string {
	return s.Feedback
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GetMessages() []*string {
	return s.Messages
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GetRuleName() *string {
	return s.RuleName
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GetRuleType() *string {
	return s.RuleType
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) GetScripts() []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	return s.Scripts
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetComments(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Comments = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetFeedback(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Feedback = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetMessages(v []*string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Messages = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetRuleName(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.RuleName = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetRuleType(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.RuleType = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) SetScripts(v []*GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults {
	s.Scripts = v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResults) Validate() error {
	if s.Scripts != nil {
		for _, item := range s.Scripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts struct {
	// The content of the SQL script.
	//
	// example:
	//
	// alter table xxx add index idx_xx(yyy);
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The purpose of the SQL script. The value is set to AddIndex.
	//
	// example:
	//
	// AddIndex
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// xxx
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) String() string {
	return dara.Prettify(s)
}

func (s GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) GoString() string {
	return s.String()
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) GetContent() *string {
	return s.Content
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) GetOpType() *string {
	return s.OpType
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) GetTableName() *string {
	return s.TableName
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) SetContent(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	s.Content = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) SetOpType(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	s.OpType = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) SetTableName(v string) *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts {
	s.TableName = &v
	return s
}

func (s *GetSQLReviewOptimizeDetailResponseBodyOptimizeDetailQualityResultResultsScripts) Validate() error {
	return dara.Validate(s)
}
